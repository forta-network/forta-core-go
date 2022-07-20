package inspect

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/forta-network/forta-core-go/ethereum"
	"github.com/hashicorp/go-multierror"
)

const (
	// MetricTraceAccessible tells if it can connect to given API.
	MetricTraceAccessible = "trace.accessible"
	// MetricTraceSupported is required only for scanning some chains.
	// It is safe to ignore the value when scanning other chains.
	MetricTraceSupported = "trace.supported"
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

	// checking API access
	rpcClient, err := rpc.DialContext(ctx, inspectionCfg.TraceAPIURL)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to dial api: %w", err))

		results.Metrics[MetricTraceAccessible] = ResultFailure
		results.Metrics[MetricTraceSupported] = ResultFailure

		return
	}

	ethClient := ethclient.NewClient(rpcClient)

	// trace is only required for mainnet
	chainID, err := ethClient.ChainID(ctx)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to get chain id: %w", err))
		results.Metrics[MetricTraceAccessible] = ResultFailure
	} else {
		results.Metrics[MetricTraceAccessible] = ResultSuccess
	}

	// no need to execute further on chains other than ethereum mainnet
	mainnetChainId := big.NewInt(1)
	if chainID == nil || (chainID != nil && chainID.Cmp(mainnetChainId) != 0) {
		results.Metrics[MetricTraceSupported] = ResultFailure

		return
	}

	// checking trace capability
	streamClient, err := ethereum.NewStreamEthClient(ctx, "trace", inspectionCfg.TraceAPIURL)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to create eth client: %w", err))
	}

	// Try tracing genesis block
	traceContext, traceCancel := context.WithTimeout(ctx, time.Second*3)
	defer traceCancel()

	_, err = streamClient.TraceBlock(traceContext, big.NewInt(0).SetUint64(inspectionCfg.BlockNumber))
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to trace block %d: %w", inspectionCfg.BlockNumber, err))
		results.Metrics[MetricTraceSupported] = ResultFailure
	} else {
		results.Metrics[MetricTraceSupported] = ResultSuccess
	}

	return
}
