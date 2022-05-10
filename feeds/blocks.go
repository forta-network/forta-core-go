package feeds

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	eth "github.com/ethereum/go-ethereum"
	"github.com/goccy/go-json"
	log "github.com/sirupsen/logrus"

	"github.com/forta-network/forta-core-go/clients/health"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/ethereum"
	"github.com/forta-network/forta-core-go/utils"
)

var ErrEndBlockReached = errors.New("end block reached")

type bfHandler struct {
	Handler func(evt *domain.BlockEvent) error
	ErrCh   chan<- error
}

type blockFeed struct {
	start       *big.Int
	end         *big.Int
	offset      int
	ctx         context.Context
	client      ethereum.Client
	traceClient ethereum.Client
	cache       utils.Cache
	chainID     *big.Int
	tracing     bool
	started     bool
	rateLimit   *time.Ticker
	maxBlockAge *time.Duration

	lastBlock health.MessageTracker

	handlers   []bfHandler
	handlersMu sync.RWMutex
}

type BlockFeedConfig struct {
	Start               *big.Int
	End                 *big.Int
	Offset              int
	ChainID             *big.Int
	RateLimit           *time.Ticker
	Tracing             bool
	SkipBlocksOlderThan *time.Duration
}

func (bf *blockFeed) initialize() error {
	if bf.start == nil {
		res, err := bf.client.BlockByNumber(bf.ctx, nil)
		if err != nil {
			return err
		}
		log.Debugf("retrieved block number %s", res.Number)

		bf.start, err = utils.HexToBigInt(res.Number)
		if err != nil {
			log.Errorf("error converting blocknum hex to bigint: %s", err.Error())
			return nil
		}

		// should be a positive number
		if bf.start.Sign() <= 0 {
			return fmt.Errorf("got invalid block number during initialization: %d", bf.start.Uint64())
		}
	}
	log.Infof("initialized block number %d", bf.start)

	if bf.chainID == nil {
		chainID, err := bf.client.ChainID(bf.ctx)
		if err != nil {
			return err
		}
		bf.chainID = chainID
	}
	log.Infof("initialized chainId %d", bf.chainID)

	return nil
}

func (bf *blockFeed) IsStarted() bool {
	return bf.started
}

func (bf *blockFeed) Start() {
	if !bf.started {
		go bf.loop()
	}
}

//StartRange runs a specific set of blocks synchronously
func (bf *blockFeed) StartRange(start int64, end int64, rate int64) {
	if !bf.started {
		if rate > 0 {
			bf.rateLimit = time.NewTicker(time.Duration(rate) * time.Millisecond)
		}
		bf.start = big.NewInt(start)
		bf.end = big.NewInt(end)
		go bf.loop()
	}
}

func (bf *blockFeed) loop() {
	if err := bf.initialize(); err != nil {
		log.WithError(err).Panic("failed to initialize")
	}

	bf.started = true
	defer func() {
		bf.started = false
	}()
	err := bf.forEachBlock()
	if err == nil {
		return
	}
	if err != ErrEndBlockReached {
		log.WithError(err).Warn("failed while processing blocks")
	}
	for _, handler := range bf.handlers {
		handler.ErrCh <- err
	}
}

func (bf *blockFeed) Subscribe(handler func(evt *domain.BlockEvent) error) <-chan error {
	bf.handlersMu.Lock()
	defer bf.handlersMu.Unlock()

	errCh := make(chan error)
	bf.handlers = append(bf.handlers, bfHandler{
		Handler: handler,
		ErrCh:   errCh,
	})
	return errCh
}

// converts from types.Log to domain.LogEntry object
func (bf *blockFeed) logsForBlock(blockNum *big.Int) ([]domain.LogEntry, error) {
	logs, err := bf.client.GetLogs(bf.ctx, eth.FilterQuery{
		FromBlock: blockNum,
		ToBlock:   blockNum,
	})
	if err != nil {
		return nil, err
	}
	var logEntries []domain.LogEntry
	b, err := json.Marshal(logs)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &logEntries); err != nil {
		return nil, err
	}
	return logEntries, nil
}

func (bf *blockFeed) forEachBlock() error {
	increment := big.NewInt(1)
	currentBlockNum := big.NewInt(bf.start.Int64())

	for {
		if bf.ctx.Err() != nil {
			return bf.ctx.Err()
		}
		if bf.rateLimit != nil {
			<-bf.rateLimit.C
		}

		blockNumToAnalyze := new(big.Int).Sub(currentBlockNum, big.NewInt(int64(bf.offset)))
		logger := log.WithFields(log.Fields{
			"currentBlock":   currentBlockNum.Uint64(),
			"blockToAnalyze": blockNumToAnalyze.Uint64(),
		})

		if bf.end != nil && blockNumToAnalyze.Uint64() > bf.end.Uint64() {
			logger.Info("end block reached - exiting")
			return ErrEndBlockReached
		}
		if bf.cache.Exists(blockNumToAnalyze.String()) {
			logger.Info("already analyzed block - skipping")
			currentBlockNum.Add(currentBlockNum, increment)
			continue
		}

		block, err := bf.client.BlockByNumber(bf.ctx, currentBlockNum)
		if err != nil {
			logger.WithError(err).Error("error getting block")
			continue
		}

		if blockNumToAnalyze.Cmp(currentBlockNum) != 0 {
			block, err = bf.client.BlockByNumber(bf.ctx, blockNumToAnalyze)
			if err != nil {
				logger.WithError(err).Error("error getting block")
				continue
			}
		}
		logger = logger.WithFields(log.Fields{
			"blockHash":         block.Hash,
			"blockToAnalyzeHex": block.Number,
		})

		if tooOld, age := blockIsTooOld(block, bf.maxBlockAge); tooOld {
			logger.WithField("age", age).Warnf("block is older than %v - setting current block num to latest", bf.maxBlockAge)
			latestBlockNum, err := bf.client.BlockNumber(bf.ctx)
			if err != nil {
				logger.WithError(err).Error("failed to get latest block number")
				continue
			}
			currentBlockNum = latestBlockNum
			continue
		}

		bf.lastBlock.Set(blockNumToAnalyze.String())

		var traces []domain.Trace
		if bf.tracing {
			traces, err = bf.traceClient.TraceBlock(bf.ctx, blockNumToAnalyze)
			if err != nil {
				logger.WithError(err).Error("error tracing block")
			}
		}

		if len(traces) > 0 && block.Hash != utils.String(traces[0].BlockHash) {
			logger.WithFields(log.Fields{
				"traceBlockHash": utils.String(traces[0].BlockHash),
			}).Warn("trace block hash != ethereum block hash, ignoring traces")
			traces = nil
		}

		logs, err := bf.logsForBlock(blockNumToAnalyze)
		if err != nil {
			logger.WithError(err).Errorf("error getting logs for block")
			continue
		}

		blockTs, err := block.GetTimestamp()
		if err != nil {
			logger.WithError(err).Error("failed to get block timestamp")
			continue
		}

		evt := &domain.BlockEvent{
			EventType: domain.EventTypeBlock,
			Block:     block,
			ChainID:   bf.chainID,
			Traces:    traces,
			Logs:      logs,
			Timestamps: &domain.TrackingTimestamps{
				Block: *blockTs,
				Feed:  time.Now().UTC(),
			},
		}
		bf.handlersMu.RLock()
		handlers := bf.handlers
		bf.handlersMu.RUnlock()
		for _, handler := range handlers {
			if err := handler.Handler(evt); err != nil {
				return err
			}
		}
		bf.cache.Add(blockNumToAnalyze.String())

		currentBlockNum.Add(currentBlockNum, increment)
	}
}

func blockIsTooOld(block *domain.Block, maxAge *time.Duration) (bool, *time.Duration) {
	if maxAge == nil {
		return false, nil
	}
	age, err := block.Age()
	if err != nil || age == nil {
		log.WithFields(log.Fields{
			"blockHex": block.Number,
		}).WithError(err).Errorf("error getting age of block")
		return false, age
	}
	return *age > *maxAge, age
}

// Name returns the name of this implementation.
func (bf *blockFeed) Name() string {
	return "block-feed"
}

// Health implements the health.Reporter interface.
func (bf *blockFeed) Health() health.Reports {
	return health.Reports{
		bf.lastBlock.GetReport("last-block"),
	}
}

func NewBlockFeed(ctx context.Context, client ethereum.Client, traceClient ethereum.Client, cfg BlockFeedConfig) (*blockFeed, error) {
	if cfg.Offset < 0 {
		return nil, fmt.Errorf("offset cannot be below zero: offset=%d", cfg.Offset)
	}
	bf := &blockFeed{
		start:       cfg.Start,
		end:         cfg.End,
		offset:      cfg.Offset,
		ctx:         ctx,
		client:      client,
		traceClient: traceClient,
		cache:       utils.NewCache(10000),
		chainID:     cfg.ChainID,
		tracing:     cfg.Tracing,
		rateLimit:   cfg.RateLimit,
		maxBlockAge: cfg.SkipBlocksOlderThan,
	}
	return bf, nil
}
