package inspect

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hashicorp/go-multierror"
)

const (
	// MetricScanAPIAccessible can connect to node
	MetricScanAPIAccessible = "scan-api.accessible"
	// MetricScanAPIChainID which chain id the json-rpc provides
	MetricScanAPIChainID = "scan-api.chain-id"
	// MetricScanAPIModuleWeb3 node supports web3 module.
	MetricScanAPIModuleWeb3 = "scan-api.module.web3"
	// MetricScanAPIModuleEth node supports eth module.
	MetricScanAPIModuleEth = "scan-api.module.eth"
	// MetricScanAPIModuleNet node supports net module.
	MetricScanAPIModuleNet = "scan-api.module.net"
	// MetricScanAPIHistorySupport the earliest supported block height. The lower is better.
	MetricScanAPIHistorySupport = "scan-api.history-support"
)

// ScanAPIInspector is an inspector implementation.
type ScanAPIInspector struct{}

var _ Inspector = &ScanAPIInspector{}

// Name returns the name of the inspector.
func (sai *ScanAPIInspector) Name() string {
	return "scan-api"
}

// Inspect checks given JSON-RPC node url supports web3, eth and net modules.
//
// it doesn't actually return any errors for now,
// because the point is to keep going and check if it supports the rest
// error return parameter is simply for keeping the function extensible without api changes in the future.
func (sai *ScanAPIInspector) Inspect(ctx context.Context, inspectionCfg InspectionConfig) (results *InspectionResults, resultErr error) {
	results = NewInspectionResults()

	rpcClient, err := rpc.DialContext(ctx, inspectionCfg.ScanAPIURL)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't dial json-rpc api %w", err))

		results.Metrics[MetricScanAPIAccessible] = ResultFailure
		results.Metrics[MetricScanAPIModuleWeb3] = ResultFailure
		results.Metrics[MetricScanAPIModuleEth] = ResultFailure
		results.Metrics[MetricScanAPIModuleNet] = ResultFailure
		results.Metrics[MetricScanAPIHistorySupport] = ResultFailure
		results.Metrics[MetricScanAPIChainID] = ResultFailure

		return
	}

	client := ethclient.NewClient(rpcClient)

	// arbitrary call to check node access
	if id, err := client.ChainID(ctx); err != nil {
		results.Metrics[MetricScanAPIAccessible] = ResultFailure
		results.Metrics[MetricScanAPIChainID] = ResultFailure
	} else {
		results.Metrics[MetricScanAPIAccessible] = ResultSuccess
		results.Metrics[MetricScanAPIChainID] = float64(id.Uint64())
	}

	currentHeight, err := client.BlockNumber(ctx)
	if err != nil {
		results.Metrics[MetricScanAPIAccessible] = ResultFailure
		results.Metrics[MetricScanAPIHistorySupport] = ResultFailure
		resultErr = multierror.Append(resultErr, err)
	} else {
		// check history support
		results.Metrics[MetricScanAPIAccessible] = ResultSuccess
		results.Metrics[MetricScanAPIHistorySupport] = checkHistorySupport(ctx, currentHeight, client)
	}

	err = moduleFunctionalityCheck(ctx, rpcClient, results)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("error checking module functionality %w", err))
	}

	return
}

// moduleFunctionalityCheck double-checks the functionality of modules that were declared as supported by
// the node.
func moduleFunctionalityCheck(
	ctx context.Context, rpcClient *rpc.Client, results *InspectionResults,
) (resultError error) {
	client := ethclient.NewClient(rpcClient)

	// sends net_version under the hood. should prove the node supports net module
	_, err := client.NetworkID(ctx)
	if err != nil {
		results.Metrics[MetricScanAPIModuleNet] = ResultFailure
		resultError = multierror.Append(resultError, err)
	} else {
		results.Metrics[MetricScanAPIModuleNet] = ResultSuccess
	}

	// sends eth_chainId under the hood. should prove the node supports eth module
	_, err = client.ChainID(ctx)
	if err != nil {
		results.Metrics[MetricScanAPIModuleEth] = ResultFailure
		resultError = multierror.Append(resultError, err)
	} else {
		results.Metrics[MetricScanAPIModuleEth] = ResultSuccess
	}

	// ask for web3 client version to prove the node supports web3 module
	err = rpcClient.CallContext(ctx, nil, "web3_clientVersion")
	if err != nil {
		resultError = multierror.Append(resultError, err)
		results.Metrics[MetricScanAPIModuleWeb3] = ResultFailure
	} else {
		results.Metrics[MetricScanAPIModuleWeb3] = ResultSuccess
	}

	return resultError
}

// checkHistorySupport inspects block history supports. results earliest provided block
func checkHistorySupport(ctx context.Context, latestBlock uint64, client *ethclient.Client) float64 {
	// check for genesis block
	_, err := client.BlockByNumber(ctx, big.NewInt(0))
	if err == nil {
		return 0
	}

	// check for earliest block if genesis wasn't available
	memo := make(map[uint64]bool)

	return float64(findOldestSupportedBlock(ctx, memo, client, latestBlock, 0))
}

// findOldestSupportedBlock returns the earliest block provided by client
func findOldestSupportedBlock(
	ctx context.Context, memo map[uint64]bool, client *ethclient.Client, low, high uint64,
) uint64 {
	// terminating condition, results merged
	for low != high {
		mid := (low + high) / 2

		// memoization trick.
		_, ok := memo[mid]
		if !ok {
			block := big.NewInt(int64(mid))

			_, err := client.BlockByNumber(ctx, block)
			isProvided := err == nil

			memo[mid] = isProvided
			// terminating condition, optimum solution
			if isProvided && mid == 0 {
				return 0
			}

			// left side of mid
			if isProvided {
				high = mid - 1

				continue
			}

			// right side of mid
			low = mid + 1
		}
	}

	return low
}
