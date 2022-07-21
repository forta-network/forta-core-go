package inspect

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/forta-network/forta-core-go/ethereum"
	"github.com/forta-network/forta-core-go/utils"
	"github.com/hashicorp/go-multierror"
)

const (
	// MetricTraceAccessible tells if it can connect to given API.
	MetricTraceAccessible = "trace-api.accessible"
	// MetricTraceSupported is required only for scanning some chains.
	// It is safe to ignore the value when scanning other chains.
	MetricTraceSupported = "trace-api.supported"

	// MetadataTraceAPIBlockByNumberHash is the hash of the block data retrieved from the trace API.
	MetadataTraceAPIBlockByNumberHash = "trace-api.block-by-number.hash"
	// MetadataTraceAPITraceBlockHash is the hash of the block trace data retrieved from the trace API.
	MetadataTraceAPITraceBlockHash = "trace-api.trace-block.hash"
)

// TraceAPIInspector is an inspector implementation.
type TraceAPIInspector struct{}

var _ Inspector = &TraceAPIInspector{}

// Name returns the name of the inspector.
func (tai *TraceAPIInspector) Name() string {
	return "trace"
}

// Inspect inspects tracing support.
func (tai *TraceAPIInspector) Inspect(ctx context.Context, inspectionCfg InspectionConfig) (results *InspectionResults, resultErr error) {
	results = NewInspectionResults()

	if len(inspectionCfg.TraceAPIURL) == 0 {
		results.Metrics[MetricTraceAccessible] = ResultFailure
		results.Metrics[MetricTraceSupported] = ResultFailure
		return
	}

	// checking API access
	rpcClient, err := rpc.DialContext(ctx, inspectionCfg.TraceAPIURL)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to dial api: %w", err))

		results.Metrics[MetricTraceAccessible] = ResultFailure
		results.Metrics[MetricTraceSupported] = ResultFailure

		return
	}

	ethClient := ethclient.NewClient(rpcClient)

	_, err = ethClient.ChainID(ctx)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to get chain id: %w", err))
		results.Metrics[MetricTraceAccessible] = ResultFailure
	} else {
		results.Metrics[MetricTraceAccessible] = ResultSuccess
	}

	streamClient, err := ethereum.NewStreamEthClient(ctx, "trace", inspectionCfg.TraceAPIURL)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to create eth client: %w", err))
	}

	traceContext, traceCancel := context.WithTimeout(ctx, time.Second*3)
	defer traceCancel()

	// checking trace capability
	traces, err := streamClient.TraceBlock(traceContext, big.NewInt(0).SetUint64(inspectionCfg.BlockNumber))
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to trace block %d: %w", inspectionCfg.BlockNumber, err))
		results.Metrics[MetricTraceSupported] = ResultFailure
	} else {
		results.Metrics[MetricTraceSupported] = ResultSuccess
		results.Metadata[MetadataTraceAPITraceBlockHash] = utils.HashNormalizedJSON(traces)
	}

	// get configured block and include hash of the returned as metadata
	var blockData json.RawMessage
	err = rpcClient.CallContext(ctx, &blockData, "eth_getBlockByNumber", hexutil.EncodeUint64(inspectionCfg.BlockNumber), true)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to get configured block %d: %v", inspectionCfg.BlockNumber, err))
	} else {
		results.Metadata[MetadataTraceAPIBlockByNumberHash] = utils.HashNormalizedJSON(blockData)
	}

	return
}
