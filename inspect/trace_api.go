package inspect

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hashicorp/go-multierror"
)

const (
	// IndicatorTraceAccessible tells if it can connect to given API.
	IndicatorTraceAccessible = "trace-api.accessible"
	// IndicatorTraceSupported is required only for scanning some chains.
	// It is safe to ignore the value when scanning other chains.
	IndicatorTraceSupported = "trace-api.supported"
	// IndicatorTraceAPIChainID which chain id the json-rpc provides
	IndicatorTraceAPIChainID = "trace-api.chain-id"
	// IndicatorTraceAPIIsETH2 is upgraded to Ethereum 2.0.
	IndicatorTraceAPIIsETH2 = "trace-api.is-eth2"
	// IndicatorTraceAPIUpToDate can get the configured block hash. Shows that there is no offset between the inspection and trace api.
	IndicatorTraceAPIUpToDate = "trace-api.up-to-date"

	// MetadataTraceAPIBlockByNumberHash is the hash of the block data retrieved from the trace API.
	MetadataTraceAPIBlockByNumberHash = "trace-api.block-by-number.hash"
	// MetadataTraceAPITraceBlockHash is the hash of the block trace data retrieved from the trace API.
	MetadataTraceAPITraceBlockHash = "trace-api.trace-block.hash"
)

var (
	traceAPIIndicators = []string{
		IndicatorTraceAccessible, IndicatorTraceSupported, IndicatorTraceAPIIsETH2,
	}
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
	results.Indicators = defaultIndicators(traceAPIIndicators)

	if !inspectionCfg.CheckTrace {
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

	id, err := GetChainOrNetworkID(ctx, rpcClient)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to get chain id: %w", err))
		results.Indicators[IndicatorTraceAccessible] = ResultFailure
		results.Indicators[IndicatorTraceSupported] = ResultFailure
		results.Indicators[IndicatorTraceAPIChainID] = ResultFailure
	} else {
		results.Indicators[IndicatorTraceAccessible] = ResultSuccess
		results.Indicators[IndicatorTraceAPIChainID] = float64(id.Uint64())
	}

	traceCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	// checking trace capability
	hash, err := GetTraceResponseHash(traceCtx, rpcClient, inspectionCfg.BlockNumber)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to trace block %d: %w", inspectionCfg.BlockNumber, err))
		results.Indicators[IndicatorTraceSupported] = ResultFailure
		results.Indicators[IndicatorTraceAPIUpToDate] = ResultFailure
	} else {
		results.Indicators[IndicatorTraceSupported] = ResultSuccess
		results.Indicators[IndicatorTraceAPIUpToDate] = ResultSuccess
		results.Metadata[MetadataTraceAPITraceBlockHash] = hash
	}

	// get configured block and include hash of the returned as metadata
	hash, err = GetBlockResponseHash(ctx, rpcClient, inspectionCfg.BlockNumber)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to get configured block %d: %v", inspectionCfg.BlockNumber, err))
	} else {
		results.Metadata[MetadataTraceAPIBlockByNumberHash] = hash
	}

	if SupportsETH2(ctx, rpcClient) {
		results.Indicators[IndicatorTraceAPIIsETH2] = ResultSuccess
	} else {
		results.Indicators[IndicatorTraceAPIIsETH2] = ResultFailure
	}

	return
}
