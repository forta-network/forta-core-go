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
	// MetricContainerTraceAccessible 1 if it can connect to given node.
	MetricContainerTraceAccessible = "container.trace.accessible"
	// MetricContainerTraceSupported Only required in Ethereum mainnet. Ignore the value on other chains.
	MetricContainerTraceSupported = "container.trace.supported"
)

// RunTraceAPIInspection tests accessibility and tracing support of given evm node.
func RunTraceAPIInspection(ctx context.Context, nodeURL string) (map[string]float64, error) {
	var (
		result      = make(map[string]float64)
		resultError error
	)

	// checking node accessibility
	rpcClient, err := rpc.DialContext(ctx, nodeURL)
	if err != nil {
		resultError = multierror.Append(resultError, fmt.Errorf("can't dial json-rpc api %w", err))

		result[MetricContainerTraceAccessible] = StateError
		result[MetricContainerTraceSupported] = StateError

		return result, resultError
	}

	ethClient := ethclient.NewClient(rpcClient)

	// trace is only required for mainnet
	chainID, err := ethClient.ChainID(ctx)
	if err != nil {
		resultError = multierror.Append(resultError, fmt.Errorf("failed to get chain id: %w", err))
		result[MetricContainerTraceAccessible] = StateError
	} else {
		result[MetricContainerTraceAccessible] = StateSuccess
	}

	// no need to execute further on chains other than ethereum mainnet
	if chainID == nil || (chainID != nil && chainID.Cmp(mainnetChainId) != 0) {
		result[MetricContainerTraceSupported] = StateError

		return result, resultError
	}

	// checking trace capability
	streamClient, err := ethereum.NewStreamEthClient(ctx, "trace", nodeURL)
	if err != nil {
		resultError = multierror.Append(resultError, fmt.Errorf("failed to create eth client: %w", err))
	}

	// Try tracing genesis block
	traceContext, traceCancel := context.WithTimeout(ctx, time.Second*3)
	defer traceCancel()

	_, err = streamClient.TraceBlock(traceContext, big.NewInt(0))
	if err != nil {
		resultError = multierror.Append(resultError, fmt.Errorf("failed to trace block 0: %w", err))
		result[MetricContainerTraceSupported] = StateError
	} else {
		result[MetricContainerTraceSupported] = StateSuccess
	}

	return result, resultError
}
