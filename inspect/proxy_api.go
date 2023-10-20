package inspect

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	geth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/forta-network/forta-core-go/ethereum"
	"github.com/forta-network/forta-core-go/utils"
	"github.com/hashicorp/go-multierror"
)

const (
	// IndicatorProxyAPIAccessible can connect to node
	IndicatorProxyAPIAccessible = "proxy-api.accessible"
	// IndicatorProxyAPIChainID which chain id the json-rpc provides
	IndicatorProxyAPIChainID = "proxy-api.chain-id"
	// IndicatorProxyAPIModuleWeb3 node supports web3 module.
	IndicatorProxyAPIModuleWeb3 = "proxy-api.module.web3"
	// IndicatorProxyAPIModuleEth node supports eth module.
	IndicatorProxyAPIModuleEth = "proxy-api.module.eth"
	// IndicatorProxyAPIModuleNet node supports net module.
	IndicatorProxyAPIModuleNet = "proxy-api.module.net"
	// IndicatorProxyAPIHistorySupport the earliest supported block height. The lower is better.
	IndicatorProxyAPIHistorySupport = "proxy-api.history-support"
	// IndicatorProxyAPIIsETH2 is upgraded to Ethereum 2.0.
	IndicatorProxyAPIIsETH2 = "proxy-api.is-eth2"
	// IndicatorProxyAPIMethodEthCall indicates whether eth_call works or not.
	IndicatorProxyAPIMethodEthCall = "proxy-api.method.eth-call"
	// IndicatorProxyAPIMethodEthLogsRange indicates whether eth_logs range is wide or not.
	IndicatorProxyAPIMethodEthLogsRange = "proxy-api.method.eth-logs-range"

	// IndicatorProxyAPIOffsetScanMean offset information between scan and proxy
	IndicatorProxyAPIOffsetScanMean    = "proxy-api.offset.scan.mean"
	IndicatorProxyAPIOffsetScanMax     = "proxy-api.offset.scan.max"
	IndicatorProxyAPIOffsetScanMedian  = "proxy-api.offset.scan.median"
	IndicatorProxyAPIOffsetScanSamples = "proxy-api.offset.scan.samples"

	// MetadataProxyAPIBlockByNumberHash is the hash of the block data retrieved from the scan API.
	MetadataProxyAPIBlockByNumberHash = "proxy-api.block-by-number.hash"
)

var (
	proxyAPIIndicators = []string{
		IndicatorProxyAPIAccessible, IndicatorProxyAPIChainID, IndicatorProxyAPIModuleWeb3, IndicatorProxyAPIModuleEth, IndicatorProxyAPIModuleNet,
		IndicatorProxyAPIHistorySupport, IndicatorProxyAPIIsETH2,
	}
	ethCallCheckToAddr  = common.HexToAddress(utils.ZeroAddress)
	ethCallCheckData    = hexutil.MustDecode("0xdeadbeef")
	inspectedBlockRange = 2000
)

const (
	// VeryOldBlockNumber is the number of a block which inspection logic considers
	// as a very old block.
	VeryOldBlockNumber = 1
)

// ProxyAPIInspector is an inspector implementation.
type ProxyAPIInspector struct{}

// compile time check: it should implement the interface
var _ Inspector = &ProxyAPIInspector{}

// Name returns the name of the inspector.
func (pai *ProxyAPIInspector) Name() string {
	return "proxy-api"
}

// Inspect checks given JSON-RPC node url supports web3, eth and net modules.
//
// it doesn't actually return any errors for now,
// because the point is to keep going and check if it supports the rest
// error return parameter is simply for keeping the function extensible without api changes in the future.
func (pai *ProxyAPIInspector) Inspect(ctx context.Context, inspectionCfg InspectionConfig) (results *InspectionResults, resultErr error) {
	results = NewInspectionResults()
	results.Indicators = defaultIndicators(proxyAPIIndicators)

	proxyRPCClient, rpcErr := RPCDialContext(ctx, inspectionCfg.ProxyAPIURL)
	proxyClient, clientErr := EthClientDialContext(ctx, inspectionCfg.ProxyAPIURL)
	if rpcErr != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't dial json-rpc api: %v", rpcErr))
	}
	if clientErr != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't dial json-rpc api: %v", clientErr))
	}
	if rpcErr != nil || clientErr != nil {
		results.Indicators[IndicatorProxyAPIAccessible] = ResultFailure
		results.Indicators[IndicatorProxyAPIModuleWeb3] = ResultFailure
		results.Indicators[IndicatorProxyAPIModuleEth] = ResultFailure
		results.Indicators[IndicatorProxyAPIModuleNet] = ResultFailure
		results.Indicators[IndicatorProxyAPIHistorySupport] = ResultFailure
		results.Indicators[IndicatorProxyAPIChainID] = ResultFailure

		return
	} else {
		results.Indicators[IndicatorProxyAPIAccessible] = ResultSuccess
	}

	if id, err := GetChainOrNetworkID(ctx, proxyRPCClient); err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't query chain id: %v", err))
		results.Indicators[IndicatorProxyAPIChainID] = ResultFailure
	} else {
		results.Indicators[IndicatorProxyAPIChainID] = float64(id.Uint64())
	}

	currentHeight, err := proxyClient.BlockNumber(ctx)
	if err != nil {
		results.Indicators[IndicatorProxyAPIAccessible] = ResultFailure
		results.Indicators[IndicatorProxyAPIHistorySupport] = ResultFailure
		resultErr = multierror.Append(resultErr, err)
	} else {
		// check history support
		results.Indicators[IndicatorProxyAPIAccessible] = ResultSuccess
		results.Indicators[IndicatorProxyAPIHistorySupport] = checkHistorySupport(ctx, currentHeight, proxyClient)
	}

	err = checkSupportedModules(ctx, proxyRPCClient, results)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("error checking module functionality %w", err))
	}

	_, err = proxyClient.CallContract(ctx, geth.CallMsg{
		To:   &ethCallCheckToAddr,
		Data: ethCallCheckData,
	}, nil)
	switch {
	case err != nil && strings.Contains(err.Error(), "revert"):
		results.Indicators[IndicatorProxyAPIMethodEthCall] = ResultSuccess

	case err != nil:
		results.Indicators[IndicatorProxyAPIMethodEthCall] = ResultFailure
		resultErr = multierror.Append(resultErr, fmt.Errorf("eth_call check failed: %v", err))

	default:
		results.Indicators[IndicatorProxyAPIMethodEthCall] = ResultSuccess
	}

	_, err = proxyClient.FilterLogs(ctx, geth.FilterQuery{
		FromBlock: big.NewInt(0).SetUint64(currentHeight - uint64(inspectedBlockRange) - 1),
		ToBlock:   big.NewInt(0).SetUint64(currentHeight - 1),
	})
	if err != nil {
		results.Indicators[IndicatorProxyAPIMethodEthLogsRange] = ResultFailure
		resultErr = multierror.Append(resultErr, fmt.Errorf("eth_logs range check failed: %v", err))
	} else {
		results.Indicators[IndicatorProxyAPIMethodEthLogsRange] = ResultSuccess
	}

	// get configured block and include hash of the returned as metadata
	hash, err := GetBlockResponseHash(ctx, proxyRPCClient, inspectionCfg.BlockNumber)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to get configured block %d: %v", inspectionCfg.BlockNumber, err))
	} else {
		results.Metadata[MetadataProxyAPIBlockByNumberHash] = hash
	}

	if SupportsETH2(ctx, proxyRPCClient) {
		results.Indicators[IndicatorProxyAPIIsETH2] = ResultSuccess
	} else {
		results.Indicators[IndicatorProxyAPIIsETH2] = ResultFailure
	}

	scanClient, err := EthClientDialContext(ctx, inspectionCfg.ScanAPIURL)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't calculate scan-proxy offset because failed to dial scan api: %w", err))
		results.Indicators[IndicatorProxyAPIOffsetScanMean] = ResultUnknown
		results.Indicators[IndicatorProxyAPIOffsetScanMedian] = ResultUnknown
		results.Indicators[IndicatorProxyAPIOffsetScanMax] = ResultUnknown
		results.Indicators[IndicatorProxyAPIOffsetScanSamples] = ResultUnknown
		return // early return
	}

	stats, err := calculateOffsetStats(ctx, proxyClient, scanClient)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't calculate scan-proxy offset: %w", err))
		results.Indicators[IndicatorProxyAPIOffsetScanMean] = ResultUnknown
		results.Indicators[IndicatorProxyAPIOffsetScanMedian] = ResultUnknown
		results.Indicators[IndicatorProxyAPIOffsetScanMax] = ResultUnknown
		results.Indicators[IndicatorProxyAPIOffsetScanSamples] = ResultUnknown
	} else {
		results.Indicators[IndicatorProxyAPIOffsetScanMean] = stats.Mean
		results.Indicators[IndicatorProxyAPIOffsetScanMedian] = stats.Median
		results.Indicators[IndicatorProxyAPIOffsetScanMax] = stats.Max
		results.Indicators[IndicatorProxyAPIOffsetScanSamples] = stats.SampleCount
	}

	return
}

// checkSupportedModules double-checks the functionality of modules that were declared as supported by
// the node.
func checkSupportedModules(
	ctx context.Context, rpcClient ethereum.RPCClient, results *InspectionResults,
) (resultError error) {
	// sends net_version under the hood. should prove the node supports net module
	_, err := GetNetworkID(ctx, rpcClient)
	if err != nil {
		results.Indicators[IndicatorProxyAPIModuleNet] = ResultFailure
		resultError = multierror.Append(resultError, err)
	} else {
		results.Indicators[IndicatorProxyAPIModuleNet] = ResultSuccess
	}

	// sends eth_chainId under the hood. should prove the node supports eth module
	_, err = GetChainID(ctx, rpcClient)
	if err != nil {
		results.Indicators[IndicatorProxyAPIModuleEth] = ResultFailure
		resultError = multierror.Append(resultError, err)
	} else {
		results.Indicators[IndicatorProxyAPIModuleEth] = ResultSuccess
	}

	// ask for web3 client version to prove the node supports web3 module
	err = rpcClient.CallContext(ctx, nil, "web3_clientVersion")
	if err != nil {
		resultError = multierror.Append(resultError, err)
		results.Indicators[IndicatorProxyAPIModuleWeb3] = ResultFailure
	} else {
		results.Indicators[IndicatorProxyAPIModuleWeb3] = ResultSuccess
	}

	return resultError
}

// checkHistorySupport inspects block history supports. results earliest provided block
func checkHistorySupport(ctx context.Context, latestBlock uint64, client ethereum.EthClient) float64 {
	// check for a very old block
	_, err := client.BlockByNumber(ctx, big.NewInt(VeryOldBlockNumber))
	if err == nil {
		return VeryOldBlockNumber
	}

	// check for earliest block if genesis wasn't available
	return float64(findOldestSupportedBlock(ctx, client, 0, latestBlock))
}

// findOldestSupportedBlock returns the earliest block provided by client
func findOldestSupportedBlock(ctx context.Context, client ethereum.EthClient, low, high uint64) uint64 {
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
