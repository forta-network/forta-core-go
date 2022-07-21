package inspect

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/forta-network/forta-core-go/ethereum"
	"github.com/forta-network/forta-core-go/utils"
	"github.com/hashicorp/go-multierror"
)

const (
	// IndicatorTraceAccessible tells if it can connect to given API.
	IndicatorTraceAccessible = "trace-api.accessible"
	// IndicatorTraceSupported is required only for scanning some chains.
	// It is safe to ignore the value when scanning other chains.
	IndicatorTraceSupported = "trace-api.supported"

	// MetadataTraceAPIBlockByNumberHash is the hash of the block data retrieved from the trace API.
	MetadataTraceAPIBlockByNumberHash = "trace-api.block-by-number.hash"
	// MetadataTraceAPITraceBlockHash is the hash of the block trace data retrieved from the trace API.
	MetadataTraceAPITraceBlockHash = "trace-api.trace-block.hash"
)

// TraceAPIInspector is an inspector implementation.
type TraceAPIInspector struct{}

// compile time check: it should implement the interface
var _ Inspector = &TraceAPIInspector{}

// Name returns the name of the inspector.
func (tai *TraceAPIInspector) Name() string {
	return "trace"
}

// Inspect inspects tracing support.
func (tai *TraceAPIInspector) Inspect(ctx context.Context, inspectionCfg InspectionConfig) (results *InspectionResults, resultErr error) {
	results = NewInspectionResults()

	if len(inspectionCfg.TraceAPIURL) == 0 {
		results.Indicators[IndicatorTraceAccessible] = ResultFailure
		results.Indicators[IndicatorTraceSupported] = ResultFailure
		return
	}

	// checking API access
	rpcClient, err := rpc.DialContext(ctx, inspectionCfg.TraceAPIURL)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to dial api: %w", err))

		results.Indicators[IndicatorTraceAccessible] = ResultFailure
		results.Indicators[IndicatorTraceSupported] = ResultFailure

		return
	}

	ethClient := ethclient.NewClient(rpcClient)

	_, err = ethClient.ChainID(ctx)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to get chain id: %w", err))
		results.Indicators[IndicatorTraceAccessible] = ResultFailure
	} else {
		results.Indicators[IndicatorTraceAccessible] = ResultSuccess
	}

	traceClient, err := ethereum.NewStreamEthClient(ctx, "trace", inspectionCfg.TraceAPIURL)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to create eth client: %w", err))
	}

	traceCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	// checking trace capability
	hash, err := getTraceResponseHash(traceCtx, traceClient, inspectionCfg.BlockNumber)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to trace block %d: %w", inspectionCfg.BlockNumber, err))
		results.Indicators[IndicatorTraceSupported] = ResultFailure
	} else {
		results.Indicators[IndicatorTraceSupported] = ResultSuccess
		results.Metadata[MetadataTraceAPITraceBlockHash] = hash
	}

	// get configured block and include hash of the returned as metadata
	hash, err = getBlockResponseHash(ctx, rpcClient, inspectionCfg.BlockNumber)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to get configured block %d: %v", inspectionCfg.BlockNumber, err))
	} else {
		results.Metadata[MetadataTraceAPIBlockByNumberHash] = hash
	}

	return
}

func getTraceResponseHash(ctx context.Context, traceClient ethereum.Client, blockNumber uint64) (string, error) {
	traces, err := traceClient.TraceBlock(ctx, big.NewInt(0).SetUint64(blockNumber))
	if err != nil {
		return "", err
	}
	return utils.HashNormalizedJSON(traces), nil
}
