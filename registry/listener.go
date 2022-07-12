package registry

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"

	"github.com/forta-network/forta-core-go/contracts/contract_agent_registry"
	"github.com/forta-network/forta-core-go/contracts/contract_dispatch"
	"github.com/forta-network/forta-core-go/contracts/contract_forta_staking"
	"github.com/forta-network/forta-core-go/contracts/contract_scanner_registry"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/domain/registry"
	"github.com/forta-network/forta-core-go/ethereum"
	"github.com/forta-network/forta-core-go/feeds"
	"github.com/forta-network/forta-core-go/utils"
)

type listener struct {
	ctx  context.Context
	cfg  ListenerConfig
	logs feeds.LogFeed
	c    Client
	eth  ethereum.Client

	scannerAddr      string
	agentAddr        string
	dispatchAddr     string
	fortaStakingAddr string

	fortaStakingFilterer *contract_forta_staking.FortaStakingFilterer
	scannerFilterer      *contract_scanner_registry.ScannerRegistryFilterer
	agentsFilterer       *contract_agent_registry.AgentRegistryFilterer
	dispatchFilterer     *contract_dispatch.DispatchFilterer
}

type Handlers struct {
	AfterBlockHandler func(blk *domain.Block) error

	// registration
	SaveAgentHandler     func(logger *log.Entry, msg *registry.AgentSaveMessage) error
	AgentActionHandler   func(logger *log.Entry, msg *registry.AgentMessage) error
	SaveScannerHandler   func(logger *log.Entry, msg *registry.ScannerSaveMessage) error
	ScannerActionHandler func(logger *log.Entry, msg *registry.ScannerMessage) error

	// assignment
	DispatchHandler func(logger *log.Entry, msg *registry.DispatchMessage) error

	// staking
	AgentStakeHandler     func(logger *log.Entry, msg *registry.AgentStakeMessage) error
	ScannerStakeHandler   func(logger *log.Entry, msg *registry.ScannerStakeMessage) error
	TransferSharesHandler func(logger *log.Entry, msg *registry.TransferSharesMessage) error

	AgentStakeThresholdHandler   func(logger *log.Entry, msg *registry.AgentStakeThresholdMessage) error
	ScannerStakeThresholdHandler func(logger *log.Entry, msg *registry.ScannerStakeThresholdMessage) error
}

type ContractFilter struct {
	AgentRegistry    bool
	ScannerRegistry  bool
	DispatchRegistry bool
	FortaStaking     bool
}

type ListenerConfig struct {
	Name           string
	JsonRpcURL     string
	ENSAddress     string
	StartBlock     *big.Int
	EndBlock       *big.Int
	BlockOffset    int
	Handlers       Handlers
	ContractFilter *ContractFilter
	Topics         []string
}

type Listener interface {
	Listen() error
	ProcessLastBlocks(blocksAgo int64) error
	ProcessBlockRange(startBlock *big.Int, endBlock *big.Int) error
}

func (l *listener) handleScannerRegistryEvent(le types.Log, blk *domain.Block, logger *log.Entry) error {
	if isEvent(le, contract_scanner_registry.ScannerUpdatedTopic) {
		su, err := l.scannerFilterer.ParseScannerUpdated(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.SaveScannerHandler != nil {
			scannerID := utils.ScannerIDBigIntToHex(su.ScannerId)
			enabled, err := l.c.IsEnabledScanner(scannerID)
			if err != nil {
				return err
			}
			return l.cfg.Handlers.SaveScannerHandler(logger, registry.NewScannerSaveMessage(su, enabled, blk))
		}
	} else if isEvent(le, contract_scanner_registry.ScannerEnabledTopic) {
		se, err := l.scannerFilterer.ParseScannerEnabled(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.ScannerActionHandler != nil {
			return l.cfg.Handlers.ScannerActionHandler(logger, registry.NewScannerMessage(se, blk))
		}
	} else if isEvent(le, contract_scanner_registry.StakeThresholdChangedTopic) {
		evt, err := l.scannerFilterer.ParseStakeThresholdChanged(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.ScannerStakeThresholdHandler != nil {
			return l.cfg.Handlers.ScannerStakeThresholdHandler(logger, registry.NewScannerStakeThresholdMessage(evt, le, blk))
		}
	}
	return nil
}

func (l *listener) handleAgentRegistryEvent(le types.Log, blk *domain.Block, logger *log.Entry) error {
	if isEvent(le, contract_agent_registry.AgentUpdatedTopic) {
		au, err := l.agentsFilterer.ParseAgentUpdated(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.SaveAgentHandler != nil {
			return l.cfg.Handlers.SaveAgentHandler(logger, registry.NewAgentSaveMessage(au, blk))
		}
	} else if isEvent(le, contract_agent_registry.AgentEnabledTopic) {
		ae, err := l.agentsFilterer.ParseAgentEnabled(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.AgentActionHandler != nil {
			return l.cfg.Handlers.AgentActionHandler(logger, registry.NewAgentMessage(ae, blk))
		}
	} else if isEvent(le, contract_agent_registry.StakeThresholdChangedTopic) {
		stc, err := l.agentsFilterer.ParseStakeThresholdChanged(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.AgentStakeThresholdHandler != nil {
			return l.cfg.Handlers.AgentStakeThresholdHandler(logger, registry.NewAgentStakeThresholdMessage(stc, le, blk))
		}
	}
	return nil
}

func (l *listener) handleFortaStakingEvent(le types.Log, blk *domain.Block, logger *log.Entry) error {
	var subjectType uint8
	var subjectID *big.Int
	var changeType string
	var value *big.Int

	if isEvent(le, contract_forta_staking.StakeDepositedTopic) {
		evt, err := l.fortaStakingFilterer.ParseStakeDeposited(le)
		if err != nil {
			return err
		}
		subjectType = evt.SubjectType
		subjectID = evt.Subject
		value = evt.Amount
		changeType = registry.ChangeTypeDeposit
	} else if isEvent(le, contract_forta_staking.WithdrawalInitiatedTopic) {
		evt, err := l.fortaStakingFilterer.ParseWithdrawalInitiated(le)
		if err != nil {
			return err
		}
		subjectType = evt.SubjectType
		subjectID = evt.Subject
		changeType = registry.ChangeTypeWithdrawal
	} else if isEvent(le, contract_forta_staking.SlashedTopic) {
		evt, err := l.fortaStakingFilterer.ParseSlashed(le)
		if err != nil {
			return err
		}
		subjectType = evt.SubjectType
		subjectID = evt.Subject
		value = evt.Value
		changeType = registry.ChangeTypeSlash
	} else if isEvent(le, contract_forta_staking.TransferSingleTopic) {
		evt, err := l.fortaStakingFilterer.ParseTransferSingle(le)
		if err != nil {
			return err
		}
		m, err := registry.TransferSharesMessageFromSingle(le, evt, blk)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.TransferSharesHandler != nil {
			return l.cfg.Handlers.TransferSharesHandler(logger, m)
		}
	} else if isEvent(le, contract_forta_staking.TransferBatchTopic) {
		evt, err := l.fortaStakingFilterer.ParseTransferBatch(le)
		if err != nil {
			return err
		}
		ms, err := registry.TransferSharesMessagesFromBatch(le, evt, blk)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.TransferSharesHandler != nil {
			for _, m := range ms {
				if err := l.cfg.Handlers.TransferSharesHandler(logger, m); err != nil {
					return err
				}
			}
			return nil
		}
	} else {
		logger.Debug("unhandled topic, ignoring")
		return nil
	}

	// parse ID for agent or scanner
	if subjectType == SubjectTypeScanner {
		scannerID := utils.ScannerIDBigIntToHex(subjectID)
		if l.cfg.Handlers.ScannerStakeHandler != nil {
			return l.cfg.Handlers.ScannerStakeHandler(logger, registry.NewScannerStakeMessage(le, changeType, scannerID, value, blk))
		}
	} else if subjectType == SubjectTypeAgent {
		agentID := utils.AgentBigIntToHex(subjectID)
		if l.cfg.Handlers.AgentStakeHandler != nil {
			return l.cfg.Handlers.AgentStakeHandler(logger, registry.NewAgentStakeMessage(le, changeType, agentID, value, blk))
		}
	} else {
		logger.WithField("subjectID", subjectType).Warn("unhandled subject ID, ignoring")
	}

	return nil
}

func (l *listener) handleDispatchEvent(le types.Log, blk *domain.Block, logger *log.Entry) error {
	if isEvent(le, contract_dispatch.LinkTopic) {
		link, err := l.dispatchFilterer.ParseLink(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.DispatchHandler != nil {
			return l.cfg.Handlers.DispatchHandler(logger, registry.NewDispatchMessage(link, blk))
		}
	}
	if isEvent(le, contract_dispatch.AlreadyLinkedTopic) {
		link, err := l.dispatchFilterer.ParseAlreadyLinked(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.DispatchHandler != nil {
			return l.cfg.Handlers.DispatchHandler(logger, registry.NewAlreadyLinkedDispatchMessage(link, blk))
		}
	}
	return nil
}

func (l *listener) handleLog(blk *domain.Block, le types.Log) error {
	if l.ctx.Err() != nil {
		return l.ctx.Err()
	}
	logger := getLoggerForLog(le)
	if equalsAddress(le.Address, l.scannerAddr) {
		return l.handleScannerRegistryEvent(le, blk, logger)
	}
	if equalsAddress(le.Address, l.agentAddr) {
		return l.handleAgentRegistryEvent(le, blk, logger)
	}
	if equalsAddress(le.Address, l.dispatchAddr) {
		return l.handleDispatchEvent(le, blk, logger)
	}
	if equalsAddress(le.Address, l.fortaStakingAddr) {
		return l.handleFortaStakingEvent(le, blk, logger)
	}
	return nil
}

func (l *listener) handleAfterBlock(blk *domain.Block) error {
	if l.ctx.Err() != nil {
		return l.ctx.Err()
	}
	if l.cfg.Handlers.AfterBlockHandler != nil {
		return l.cfg.Handlers.AfterBlockHandler(blk)
	}
	return nil
}

//ProcessBlockRange pages over a range of blocks, 10k blocks per page
func (l *listener) ProcessBlockRange(startBlock *big.Int, endBlock *big.Int) error {
	start := startBlock
	pageSize := big.NewInt(10000)
	end := math.BigMin(big.NewInt(0).Add(start, pageSize), endBlock)
	for end.Cmp(endBlock) <= 0 {
		logs, err := l.logs.GetLogsForRange(start, end)
		if err != nil {
			return err
		}
		var block *domain.Block
		for _, lg := range logs {
			if block == nil || block.Number != utils.BigIntToHex(big.NewInt(int64(lg.BlockNumber))) {
				blk, err := l.eth.BlockByNumber(l.ctx, big.NewInt(int64(lg.BlockNumber)))
				if err != nil {
					return err
				}
				block = blk
			}

			if err := l.handleLog(block, lg); err != nil {
				return err
			}
		}
		if end.Cmp(endBlock) == 0 {
			return nil
		}
		start = big.NewInt(0).Add(end, big.NewInt(1))
		end = math.BigMin(big.NewInt(0).Add(start, pageSize), endBlock)
	}
	return nil
}

// ProcessLastBlocks fetches the logs in a single pass and calls handlers for them
func (l *listener) ProcessLastBlocks(blocksAgo int64) error {
	bn, err := l.eth.BlockNumber(context.Background())
	if err != nil {
		return err
	}
	if bn.Int64() == 0 {
		return errors.New("current block is unexpectedly 0")
	}
	start := bn
	end := big.NewInt(bn.Int64() - blocksAgo)
	return l.ProcessBlockRange(start, end)
}

func (l *listener) Listen() error {
	return l.logs.ForEachLog(l.handleLog, l.handleAfterBlock)
}

func NewDefaultListener(ctx context.Context, handlers Handlers) (*listener, error) {
	return NewListener(ctx, ListenerConfig{
		Name:       "listener",
		JsonRpcURL: defaultConfig.JsonRpcUrl,
		ENSAddress: defaultEnsAddress,
		Handlers:   handlers,
	})
}

func NewListener(ctx context.Context, cfg ListenerConfig) (*listener, error) {
	jsonRpc := cfg.JsonRpcURL
	if jsonRpc == "" {
		jsonRpc = defaultConfig.JsonRpcUrl
	}
	ensAddr := cfg.ENSAddress
	if ensAddr == "" {
		ensAddr = defaultEnsAddress
	}

	client, err := ethereum.NewStreamEthClient(ctx, cfg.Name, jsonRpc)
	if err != nil {
		return nil, err
	}

	c, err := NewClient(ctx, ClientConfig{
		JsonRpcUrl: jsonRpc,
		ENSAddress: ensAddr,
		Name:       "registry-listener",
	})
	if err != nil {
		return nil, err
	}

	regContracts := c.contracts

	sf, err := contract_scanner_registry.NewScannerRegistryFilterer(regContracts.ScannerRegistry, nil)
	if err != nil {
		return nil, err
	}

	af, err := contract_agent_registry.NewAgentRegistryFilterer(regContracts.AgentRegistry, nil)
	if err != nil {
		return nil, err
	}

	df, err := contract_dispatch.NewDispatchFilterer(regContracts.Dispatch, nil)
	if err != nil {
		return nil, err
	}

	stkf, err := contract_forta_staking.NewFortaStakingFilterer(regContracts.FortaStaking, nil)
	if err != nil {
		return nil, err
	}

	var addrs []string
	if cfg.ContractFilter != nil {
		if cfg.ContractFilter.AgentRegistry {
			addrs = append(addrs, regContracts.AgentRegistry.Hex())
		}
		if cfg.ContractFilter.ScannerRegistry {
			addrs = append(addrs, regContracts.ScannerRegistry.Hex())
		}
		if cfg.ContractFilter.FortaStaking {
			addrs = append(addrs, regContracts.FortaStaking.Hex())
		}
		if cfg.ContractFilter.DispatchRegistry {
			addrs = append(addrs, regContracts.Dispatch.Hex())
		}
	} else {
		addrs = []string{regContracts.AgentRegistry.Hex(), regContracts.ScannerRegistry.Hex(), regContracts.Dispatch.Hex(), regContracts.FortaStaking.Hex()}
	}

	var topics [][]string
	if len(cfg.Topics) > 0 {
		topics = [][]string{cfg.Topics}
	}

	logFeed, err := feeds.NewLogFeed(ctx, client, feeds.LogFeedConfig{
		Addresses:  addrs,
		Topics:     topics,
		StartBlock: cfg.StartBlock,
		EndBlock:   cfg.EndBlock,
		Offset:     cfg.BlockOffset,
	})

	if err != nil {
		return nil, err
	}

	return &listener{
		ctx:                  ctx,
		c:                    c,
		cfg:                  cfg,
		logs:                 logFeed,
		eth:                  client,
		scannerAddr:          regContracts.ScannerRegistry.Hex(),
		agentAddr:            regContracts.AgentRegistry.Hex(),
		dispatchAddr:         regContracts.Dispatch.Hex(),
		fortaStakingAddr:     regContracts.FortaStaking.Hex(),
		scannerFilterer:      sf,
		agentsFilterer:       af,
		dispatchFilterer:     df,
		fortaStakingFilterer: stkf,
	}, nil
}
