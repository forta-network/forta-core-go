package ethereum

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	backoff "github.com/cenkalti/backoff/v4"
	"github.com/gorilla/websocket"

	"github.com/forta-network/forta-core-go/clients/health"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/utils"
	"github.com/forta-network/forta-core-go/utils/httpclient"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	log "github.com/sirupsen/logrus"
)

// EthClient is the original interface from go-ethereum.
type EthClient interface {
	ethereum.ChainReader
	ethereum.ChainStateReader
	ethereum.TransactionReader
	ethereum.LogFilterer
	ethereum.ContractCaller
	BlockNumber(ctx context.Context) (uint64, error)
	BlockByNumber(ctx context.Context, blockNum *big.Int) (*types.Block, error)
	Close()
}

// RPCClient is a wrapper implementation of the RPC client.
type RPCClient interface {
	Close()
	Call(result interface{}, method string, args ...interface{}) error
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
}

// Subscriber subscribes to Ethereum namespaces.
type Subscriber interface {
	RPCClient
	Subscribe(ctx context.Context, namespace string, channel interface{}, args ...interface{}) (domain.ClientSubscription, error)
}

// Client is an interface encompassing all ethereum actions
type Client interface {
	Close()
	SetRetryInterval(time.Duration)
	IsWebsocket() bool

	BlockByHash(ctx context.Context, hash string) (*domain.Block, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*domain.Block, error)
	BlockNumber(ctx context.Context) (*big.Int, error)
	TransactionReceipt(ctx context.Context, txHash string) (*domain.TransactionReceipt, error)
	ChainID(ctx context.Context) (*big.Int, error)
	GetTransactionCount(ctx context.Context, address string, blockNumber any) (*big.Int, error)
	TraceBlock(ctx context.Context, number *big.Int) ([]domain.Trace, error)
	DebugTraceCall(
		ctx context.Context, req *domain.TraceCallTransaction,
		block any, traceCallConfig domain.TraceCallConfig,
		result interface{},
	) error
	GetLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
	SubscribeToHead(ctx context.Context) (domain.HeaderCh, error)

	health.Reporter
}

const (
	blocksByNumber      = "eth_getBlockByNumber"
	blocksByHash        = "eth_getBlockByHash"
	blockNumber         = "eth_blockNumber"
	getLogs             = "eth_getLogs"
	transactionReceipt  = "eth_getTransactionReceipt"
	traceBlock          = "trace_block"
	debugTraceCall      = "debug_traceCall"
	chainId             = "eth_chainId"
	getTransactionCount = "eth_getTransactionCount"
)

const defaultRetryInterval = time.Second * 15

var ErrNotFound = fmt.Errorf("not found")

// any non-retriable failure errors can be listed here
var permanentErrors = []string{
	"method not found",
	"hash is not currently canonical",
	//"unknown block",
	"unable to complete request at this time",
	"503 service unavailable",
	"trace_block is not available",
	"invalid host",
	"receipt was empty",
}

var minBackoff = 1 * time.Second
var maxBackoff = 1 * time.Minute

// streamEthClient wraps a go-ethereum client purpose-built for streaming txs (with long retries/timeouts)
type streamEthClient struct {
	apiName       string
	rpcClient     RPCClient
	subscriber    Subscriber
	retryInterval time.Duration
	isWebsocket   bool

	lastBlockByNumberReq         health.TimeTracker
	lastBlockByNumberErr         health.ErrorTracker
	lastGetTransactionReceiptReq health.TimeTracker
	lastGetTransactionReceiptErr health.ErrorTracker
	lastTraceBlockReq            health.TimeTracker
	lastTraceBlockErr            health.ErrorTracker
}

type RetryOptions struct {
	MaxElapsedTime *time.Duration
	MinBackoff     *time.Duration
	MaxBackoff     *time.Duration
}

// Close invokes close on the underlying client
func (e *streamEthClient) Close() {
	e.rpcClient.Close()
}

func (e *streamEthClient) SetRetryInterval(d time.Duration) {
	e.retryInterval = d
}

func (e *streamEthClient) IsWebsocket() bool {
	return e.isWebsocket
}

func isPermanentError(err error) bool {
	if err == nil {
		return false
	}
	for _, pe := range permanentErrors {
		if strings.Contains(strings.ToLower(err.Error()), pe) {
			return true
		}
	}
	return false
}

// withBackoff wraps an operation in an exponential backoff logic
func withBackoff(
	ctx context.Context, name string, operation func(ctx context.Context) error, options RetryOptions,
	timeTracker *health.TimeTracker, errorTracker *health.ErrorTracker,
) error {
	bo := backoff.NewExponentialBackOff()
	bo.MaxInterval = maxBackoff
	bo.InitialInterval = minBackoff
	if options.MinBackoff != nil {
		bo.InitialInterval = *options.MinBackoff
	}
	if options.MaxBackoff != nil {
		bo.MaxInterval = *options.MaxBackoff
	}
	if options.MaxElapsedTime != nil {
		bo.MaxElapsedTime = *options.MaxElapsedTime
	}
	err := backoff.Retry(func() error {
		if ctx.Err() != nil {
			return backoff.Permanent(ctx.Err())
		}

		tCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
		err := operation(tCtx)
		cancel()
		if timeTracker != nil {
			timeTracker.Set()
		}
		if errorTracker != nil {
			errorTracker.Set(err)
		}
		if err == nil {
			//success, returning now avoids failing on context timeouts in certain edge cases
			return nil
		} else if isPermanentError(err) {
			log.Errorf("backoff permanent error: %s", err.Error())
			return backoff.Permanent(err)
		} else if ctx.Err() != nil {
			log.Errorf("%s context err found: %s", name, ctx.Err())
			return backoff.Permanent(ctx.Err())
		} else {
			log.Warnf("%s failed...retrying: %s", name, err.Error())
		}
		return err
	}, bo)
	if err != nil {
		log.Errorf("%s failed with error: %s", name, err.Error())
	}
	return err
}

func pointDur(d time.Duration) *time.Duration {
	return &d
}

// BlockByHash returns the block by hash
func (e *streamEthClient) BlockByHash(ctx context.Context, hash string) (*domain.Block, error) {
	name := fmt.Sprintf("%s(%s)", blocksByHash, hash)
	log.Debugf(name)
	var result domain.Block
	err := withBackoff(ctx, name, func(ctx context.Context) error {
		err := e.rpcClient.CallContext(ctx, &result, blocksByHash, hash, true)
		if err != nil {
			return err
		}
		if result.Hash == "" {
			return ErrNotFound
		}
		return nil
	}, RetryOptions{
		MinBackoff:     pointDur(5 * time.Second),
		MaxElapsedTime: pointDur(12 * time.Hour),
		MaxBackoff:     pointDur(15 * time.Second),
	}, nil, nil)
	return &result, err
}

// TraceBlock returns the traced block
func (e *streamEthClient) TraceBlock(ctx context.Context, number *big.Int) ([]domain.Trace, error) {
	name := fmt.Sprintf("%s(%s)", traceBlock, number)
	log.Debugf(name)
	var result []domain.Trace
	err := withBackoff(ctx, name, func(ctx context.Context) error {
		err := e.rpcClient.CallContext(ctx, &result, traceBlock, utils.BigIntToHex(number))
		if err != nil {
			return err
		}
		if len(result) == 0 {
			return ErrNotFound
		}
		return nil
	}, RetryOptions{
		MinBackoff:     pointDur(e.retryInterval),
		MaxElapsedTime: pointDur(1 * time.Minute),
		MaxBackoff:     pointDur(e.retryInterval),
	}, &e.lastTraceBlockReq, &e.lastTraceBlockErr)
	return result, err
}

// DebugTraceCall returns the traces of a call.
func (e *streamEthClient) DebugTraceCall(
	ctx context.Context, req *domain.TraceCallTransaction,
	block any, traceCallConfig domain.TraceCallConfig,
	result interface{},
) error {
	name := fmt.Sprintf("%s(%v)", debugTraceCall, req)
	log.Debugf(name)

	switch block.(type) {
	case string:
	case *rpc.BlockNumberOrHash:
	default:
		return errors.New("invalid block number type")
	}

	args := []interface{}{req, block, traceCallConfig}

	err := withBackoff(ctx, name, func(ctx context.Context) error {
		err := e.rpcClient.CallContext(ctx, &result, debugTraceCall, args...)
		if err != nil {
			return err
		}

		return nil
	}, RetryOptions{
		MinBackoff:     pointDur(e.retryInterval),
		MaxElapsedTime: pointDur(1 * time.Minute),
		MaxBackoff:     pointDur(e.retryInterval),
	}, nil, nil)
	return err
}

// GetLogs returns the set of logs for a block
func (e *streamEthClient) GetLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	name := fmt.Sprintf("%s(%v)", getLogs, q)
	log.Debugf(name)
	var result []types.Log

	args, err := toFilterArg(q)
	if err != nil {
		return nil, err
	}

	err = withBackoff(ctx, name, func(ctx context.Context) error {
		return e.rpcClient.CallContext(ctx, &result, getLogs, args)
	}, RetryOptions{
		MinBackoff:     pointDur(e.retryInterval),
		MaxElapsedTime: pointDur(12 * time.Hour),
		MaxBackoff:     pointDur(15 * time.Second),
	}, nil, nil)
	return result, err
}

// BlockByNumber returns the block by number
func (e *streamEthClient) BlockByNumber(ctx context.Context, number *big.Int) (*domain.Block, error) {
	var result domain.Block
	var (
		numArg     = "latest"
		numDisplay = numArg
	)
	if number != nil {
		numArg = utils.BigIntToHex(number) // hex representation
		numDisplay = number.String()       // integer representation
	}
	name := fmt.Sprintf("%s(%s)", blocksByNumber, numDisplay)
	log.Debugf(name)

	err := withBackoff(ctx, name, func(ctx context.Context) error {
		err := e.rpcClient.CallContext(ctx, &result, blocksByNumber, numArg, true)
		if err != nil {
			return err
		}
		if result.Hash == "" {
			return ErrNotFound
		}
		return nil
	}, RetryOptions{
		MinBackoff:     pointDur(e.retryInterval),
		MaxElapsedTime: pointDur(12 * time.Hour),
		MaxBackoff:     pointDur(e.retryInterval),
	}, &e.lastBlockByNumberReq, &e.lastBlockByNumberErr)
	return &result, err
}

// BlockNumber returns the latest block number
func (e *streamEthClient) BlockNumber(ctx context.Context) (*big.Int, error) {
	log.Debugf(blockNumber)
	var result string
	err := withBackoff(ctx, blockNumber, func(ctx context.Context) error {
		return e.rpcClient.CallContext(ctx, &result, blockNumber)
	}, RetryOptions{
		MaxElapsedTime: pointDur(12 * time.Hour),
	}, nil, nil)
	if err != nil {
		return nil, err
	}
	return utils.HexToBigInt(result)
}

// ChainID gets the chainID for a network
func (e *streamEthClient) ChainID(ctx context.Context) (*big.Int, error) {
	log.Debugf(chainId)
	var result string
	err := withBackoff(ctx, chainId, func(ctx context.Context) error {
		return e.rpcClient.CallContext(ctx, &result, chainId)
	}, RetryOptions{
		MaxElapsedTime: pointDur(1 * time.Minute),
	}, nil, nil)
	if err != nil {
		return nil, err
	}
	return utils.HexToBigInt(result)
}

// TransactionReceipt returns the receipt for a transaction
func (e *streamEthClient) TransactionReceipt(ctx context.Context, txHash string) (*domain.TransactionReceipt, error) {
	name := fmt.Sprintf("%s(%s)", transactionReceipt, txHash)
	log.Debugf(name)
	var result domain.TransactionReceipt
	err := withBackoff(ctx, name, func(ctx context.Context) error {
		if err := e.rpcClient.CallContext(ctx, &result, transactionReceipt, txHash); err != nil {
			return err
		}
		if result.TransactionHash == nil {
			return errors.New("receipt was empty")
		}
		return nil
	}, RetryOptions{
		MaxElapsedTime: pointDur(5 * time.Minute),
	}, &e.lastGetTransactionReceiptReq, &e.lastGetTransactionReceiptErr)
	return &result, err
}

// GetTransactionCount returns the transaction count for an address
func (e *streamEthClient) GetTransactionCount(ctx context.Context, address string, block any) (*big.Int, error) {
	name := fmt.Sprintf("%s(%s, %s)", getTransactionCount, address, block)

	switch block.(type) {
	case string:
	case *rpc.BlockNumberOrHash:
	default:
		return nil, errors.New("invalid block number type")
	}

	log.Debugf(name)
	var result string
	err := withBackoff(ctx, name, func(ctx context.Context) error {
		err := e.rpcClient.CallContext(ctx, &result, getTransactionCount, address, block)
		if err != nil {
			return err
		}
		if result == "" {
			return ErrNotFound
		}
		return nil
	}, RetryOptions{
		MinBackoff:     pointDur(e.retryInterval),
		MaxElapsedTime: pointDur(12 * time.Hour),
		MaxBackoff:     pointDur(e.retryInterval),
	}, nil, nil)
	if err != nil {
		return nil, err
	}

	return utils.HexToBigInt(result)
}

// SubscribeToHead subscribes to the blockchain head and returns a channel which provides
// the latest block headers. The channel is closed when subscription encounters an error
// or becomes inactive (e.g. due to a hanging connection).
func (e *streamEthClient) SubscribeToHead(ctx context.Context) (domain.HeaderCh, error) {
	log.Debug("subscribing to blockchain head")
	recvCh := make(chan *types.Header)
	sendCh := make(chan *types.Header)
	sub, err := e.subscriber.Subscribe(ctx, "eth", recvCh, "newHeads")
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe: %v", err)
	}
	go e.listenToSubscription(ctx, sub, recvCh, sendCh)
	return sendCh, nil
}

func (e *streamEthClient) listenToSubscription(ctx context.Context, sub domain.ClientSubscription, recvCh, sendCh chan *types.Header) {
	defer close(recvCh)
	defer close(sendCh)
	for {
		select {
		case <-ctx.Done():
			log.WithError(ctx.Err()).Info("exiting subscription")
			return

		case header := <-recvCh:
			sendCh <- header

		case <-time.After(time.Minute): // this avoids getting stuck when connection hangs
			log.Warn("subscription is inactive! exiting loop")
			sub.Unsubscribe()
			return

		case err := <-sub.Err():
			log.WithError(err).Error("subscription returned an error")
			return
		}
	}
}

// Name returns the name of this implementation.
func (e *streamEthClient) Name() string {
	return fmt.Sprintf("%s-json-rpc-client", e.apiName)
}

// Health implements the health.Reporter interface.
func (e *streamEthClient) Health() health.Reports {
	return health.Reports{
		e.lastBlockByNumberReq.GetReport("request.block-by-number.time"),
		e.lastBlockByNumberErr.GetReport("request.block-by-number.error"),
		e.lastGetTransactionReceiptReq.GetReport("request.get-transaction-receipt.time"),
		e.lastGetTransactionReceiptErr.GetReport("request.get-transaction-receipt.error"),
		e.lastTraceBlockReq.GetReport("request.trace-block.time"),
		e.lastTraceBlockErr.GetReport("request.trace-block.error"),
	}
}

type rpcClient struct {
	*rpc.Client
}

func (rc *rpcClient) Subscribe(ctx context.Context, namespace string, channel interface{}, args ...interface{}) (domain.ClientSubscription, error) {
	sub, err := rc.Client.Subscribe(ctx, namespace, channel, args...)
	return sub, err
}

var wsBufferPool = new(sync.Pool)

func NewRpcClient(ctx context.Context, url string) (*rpc.Client, error) {
	if isWebsocket(url) {
		dialer := *websocket.DefaultDialer
		dialer.WriteBufferSize = 1024
		dialer.ReadBufferSize = 1024
		dialer.WriteBufferPool = wsBufferPool
		dialer.HandshakeTimeout = time.Second * 10
		return rpc.DialWebsocketWithDialer(ctx, url, "", dialer)
	}

	client := *httpclient.Default
	client.Transport = &http.Transport{
		DialContext: (&net.Dialer{
			KeepAlive: 30 * time.Second,
			Timeout:   5 * time.Second,
		}).DialContext,
		IdleConnTimeout:       5 * time.Second,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	return rpc.DialHTTPWithClient(url, &client)
}

// NewStreamEthClient creates a new ethereum client
func NewStreamEthClient(ctx context.Context, apiName, apiURL string) (*streamEthClient, error) {
	rClient, err := NewRpcClient(ctx, apiURL)
	if err != nil {
		return nil, err
	}
	rClient.SetHeader("Content-Type", "application/json")

	return NewStreamEthClientWithRPCClient(ctx, apiName, isWebsocket(apiURL), &rpcClient{Client: rClient})
}

// NewStreamEthClientWithRPCClient creates a new ethereum client
func NewStreamEthClientWithRPCClient(ctx context.Context, apiName string, isWs bool, rpcClient Subscriber) (*streamEthClient, error) {
	return &streamEthClient{
		apiName:       apiName,
		rpcClient:     rpcClient,
		subscriber:    rpcClient,
		retryInterval: defaultRetryInterval,
		isWebsocket:   isWs,
	}, nil
}

func isWebsocket(apiURL string) bool {
	u, err := url.Parse(apiURL)
	if err != nil {
		return false
	}
	return u.Scheme == "ws" || u.Scheme == "wss"
}
