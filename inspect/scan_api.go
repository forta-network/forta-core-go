package inspect

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/forta-network/forta-core-go/utils"
	"github.com/hashicorp/go-multierror"
)

const (
	// IndicatorScanAPIAccessible can connect to node
	IndicatorScanAPIAccessible = "scan-api.accessible"
	// IndicatorScanAPIChainID which chain id the json-rpc provides
	IndicatorScanAPIChainID = "scan-api.chain-id"
	// IndicatorScanAPIModuleWeb3 node supports web3 module.
	IndicatorScanAPIModuleWeb3 = "scan-api.module.web3"
	// IndicatorScanAPIModuleEth node supports eth module.
	IndicatorScanAPIModuleEth = "scan-api.module.eth"
	// IndicatorScanAPIModuleNet node supports net module.
	IndicatorScanAPIModuleNet = "scan-api.module.net"
	// IndicatorScanAPIHistorySupport the earliest supported block height. The lower is better.
	IndicatorScanAPIHistorySupport = "scan-api.history-support"

	// MetadataScanAPIBlockByNumberHash is the hash of the block data retrieved from the scan API.
	MetadataScanAPIBlockByNumberHash = "scan-api.block-by-number.hash"
)

const (
	// VeryOldBlockNumber is the number of a block which inspection logic considers
	// as a very old block.
	VeryOldBlockNumber = 5
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

		results.Indicators[IndicatorScanAPIAccessible] = ResultFailure
		results.Indicators[IndicatorScanAPIModuleWeb3] = ResultFailure
		results.Indicators[IndicatorScanAPIModuleEth] = ResultFailure
		results.Indicators[IndicatorScanAPIModuleNet] = ResultFailure
		results.Indicators[IndicatorScanAPIHistorySupport] = ResultFailure
		results.Indicators[IndicatorScanAPIChainID] = ResultFailure

		return
	}

	client := ethclient.NewClient(rpcClient)

	// arbitrary call to check node access
	if id, err := client.ChainID(ctx); err != nil {
		results.Indicators[IndicatorScanAPIAccessible] = ResultFailure
		results.Indicators[IndicatorScanAPIChainID] = ResultFailure
	} else {
		results.Indicators[IndicatorScanAPIAccessible] = ResultSuccess
		results.Indicators[IndicatorScanAPIChainID] = float64(id.Uint64())
	}

	currentHeight, err := client.BlockNumber(ctx)
	if err != nil {
		results.Indicators[IndicatorScanAPIAccessible] = ResultFailure
		results.Indicators[IndicatorScanAPIHistorySupport] = ResultFailure
		resultErr = multierror.Append(resultErr, err)
	} else {
		// check history support
		results.Indicators[IndicatorScanAPIAccessible] = ResultSuccess
		results.Indicators[IndicatorScanAPIHistorySupport] = checkHistorySupport(ctx, currentHeight, client)
	}

	err = checkSupportedModules(ctx, rpcClient, results)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("error checking module functionality %w", err))
	}

	// get configured block and include hash of the returned as metadata
	var blockData json.RawMessage
	err = rpcClient.CallContext(ctx, &blockData, "eth_getBlockByNumber", hexutil.EncodeUint64(inspectionCfg.BlockNumber), true)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to get configured block %d: %v", inspectionCfg.BlockNumber, err))
	} else {
		results.Metadata[MetadataScanAPIBlockByNumberHash] = utils.HashNormalizedJSON(blockData)
	}

	return
}

// checkSupportedModules double-checks the functionality of modules that were declared as supported by
// the node.
func checkSupportedModules(
	ctx context.Context, rpcClient *rpc.Client, results *InspectionResults,
) (resultError error) {
	client := ethclient.NewClient(rpcClient)

	// sends net_version under the hood. should prove the node supports net module
	_, err := client.NetworkID(ctx)
	if err != nil {
		results.Indicators[IndicatorScanAPIModuleNet] = ResultFailure
		resultError = multierror.Append(resultError, err)
	} else {
		results.Indicators[IndicatorScanAPIModuleNet] = ResultSuccess
	}

	// sends eth_chainId under the hood. should prove the node supports eth module
	_, err = client.ChainID(ctx)
	if err != nil {
		results.Indicators[IndicatorScanAPIModuleEth] = ResultFailure
		resultError = multierror.Append(resultError, err)
	} else {
		results.Indicators[IndicatorScanAPIModuleEth] = ResultSuccess
	}

	// ask for web3 client version to prove the node supports web3 module
	err = rpcClient.CallContext(ctx, nil, "web3_clientVersion")
	if err != nil {
		resultError = multierror.Append(resultError, err)
		results.Indicators[IndicatorScanAPIModuleWeb3] = ResultFailure
	} else {
		results.Indicators[IndicatorScanAPIModuleWeb3] = ResultSuccess
	}

	return resultError
}

// checkHistorySupport inspects block history supports. results earliest provided block
func checkHistorySupport(ctx context.Context, latestBlock uint64, client *ethclient.Client) float64 {
	// check for a very old block
	_, err := client.BlockByNumber(ctx, big.NewInt(VeryOldBlockNumber))
	if err == nil {
		return VeryOldBlockNumber
	}

	// check for earliest block if genesis wasn't available
	return float64(findOldestSupportedBlock(ctx, client, 0, latestBlock))
}

// findOldestSupportedBlock returns the earliest block provided by client
func findOldestSupportedBlock(ctx context.Context, client *ethclient.Client, low, high uint64) uint64 {
	memo := make(map[uint64]bool)

	// terminating condition, results merged
	for low < high {
		mid := (low + high) / 2

		// memoization trick.
		_, ok := memo[mid]
		if ok {
			continue
		}

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

	return low
}
