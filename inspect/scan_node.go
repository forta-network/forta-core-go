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
	// MetricContainerScanAPIAccessible can connect to node
	MetricContainerScanAPIAccessible = "container.scanapi.accessible"
	// MetricContainerScanAPIChainID which chain id the json-rpc provides
	MetricContainerScanAPIChainID = "container.scanapi.chainid"
	// MetricContainerScanAPIModuleWeb3 node supports web3 module.
	MetricContainerScanAPIModuleWeb3 = "container.scanapi.module.web3"
	// MetricContainerScanAPIModuleEth node supports eth module.
	MetricContainerScanAPIModuleEth = "container.scanapi.module.eth"
	// MetricContainerScanAPIModuleNet node supports net module.
	MetricContainerScanAPIModuleNet = "container.scanapi.module.net"
	// MetricContainerScanAPIHistorySupport the earliest supported block height. The lower is better.
	MetricContainerScanAPIHistorySupport = "container.scanapi.historysupport"
)

// RunScanNodeInspection checks given JSON-RPC node url supports web3, eth and net modules.
//
// it doesn't actually return any errors for now,
// because the point is to keep going and check if it supports the rest
// error return parameter is simply for keeping the function extensible without api changes in the future.
func RunScanNodeInspection(ctx context.Context, nodeURL string) (map[string]float64, error) {
	var (
		result    = make(map[string]float64)
		resultErr error
	)

	rpcClient, err := rpc.DialContext(ctx, nodeURL)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't dial json-rpc api %w", err))

		result[MetricContainerScanAPIAccessible] = StateError
		result[MetricContainerScanAPIModuleWeb3] = StateError
		result[MetricContainerScanAPIModuleEth] = StateError
		result[MetricContainerScanAPIModuleNet] = StateError
		result[MetricContainerScanAPIHistorySupport] = StateError
		result[MetricContainerScanAPIChainID] = StateError

		return result, resultErr
	}

	client := ethclient.NewClient(rpcClient)

	// arbitrary call to check node access
	if id, err := client.ChainID(ctx); err != nil {
		result[MetricContainerScanAPIAccessible] = StateError
		result[MetricContainerScanAPIChainID] = StateError
	} else {
		result[MetricContainerScanAPIAccessible] = StateSuccess
		result[MetricContainerScanAPIChainID] = float64(id.Uint64())
	}

	currentHeight, err := client.BlockNumber(ctx)
	if err != nil {
		result[MetricContainerScanAPIAccessible] = StateError
		result[MetricContainerScanAPIHistorySupport] = StateError
		resultErr = multierror.Append(resultErr, err)
	} else {
		// check history support
		result[MetricContainerScanAPIAccessible] = StateSuccess
		result[MetricContainerScanAPIHistorySupport] = checkHistorySupport(ctx, currentHeight, client)
	}

	err = moduleFunctionalityCheck(ctx, rpcClient, result)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("error checking module functionality %w", err))
	}

	return result, resultErr
}

// moduleFunctionalityCheck double-checks the functionality of modules that were declared as supported by
// the node.
func moduleFunctionalityCheck(
	ctx context.Context, rpcClient *rpc.Client, result map[string]float64,
) (resultError error) {
	client := ethclient.NewClient(rpcClient)

	// sends net_version under the hood. should prove the node supports net module
	_, err := client.NetworkID(ctx)
	if err != nil {
		result[MetricContainerScanAPIModuleNet] = StateError
		resultError = multierror.Append(resultError, err)
	} else {
		result[MetricContainerScanAPIModuleNet] = StateSuccess
	}

	// sends eth_chainId under the hood. should prove the node supports eth module
	_, err = client.ChainID(ctx)
	if err != nil {
		result[MetricContainerScanAPIModuleEth] = StateError
		resultError = multierror.Append(resultError, err)
	} else {
		result[MetricContainerScanAPIModuleEth] = StateSuccess
	}

	// ask for web3 client version to prove the node supports web3 module
	err = rpcClient.CallContext(ctx, nil, "web3_clientVersion")
	if err != nil {
		resultError = multierror.Append(resultError, err)
		result[MetricContainerScanAPIModuleWeb3] = StateError
	} else {
		result[MetricContainerScanAPIModuleWeb3] = StateSuccess
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
