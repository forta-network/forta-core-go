package inspect

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hashicorp/go-multierror"
)

const (
	// IndicatorScanAPIAccessible can connect to node
	IndicatorScanAPIAccessible = "scan-api.accessible"
	// IndicatorScanAPIChainID which chain id the json-rpc provides
	IndicatorScanAPIChainID = "scan-api.chain-id"

	// IndicatorScanAPIModuleEth node supports eth module.
	IndicatorScanAPIModuleEth = "scan-api.module.eth"
	// IndicatorScanAPIModuleNet node supports net module.
	IndicatorScanAPIModuleNet = "scan-api.module.net"

	// IndicatorScanAPIIsETH2 is upgraded to Ethereum 2.0.
	IndicatorScanAPIIsETH2 = "scan-api.is-eth2"

	// MetadataScanAPIBlockByNumberHash is the hash of the block data retrieved from the scan API.
	MetadataScanAPIBlockByNumberHash = "scan-api.block-by-number.hash"
)

var (
	scanAPIIndicators = []string{
		IndicatorScanAPIAccessible, IndicatorScanAPIChainID, IndicatorScanAPIModuleEth, IndicatorScanAPIModuleNet, IndicatorScanAPIIsETH2,
	}
)

// ScanAPIInspector is an inspector implementation.
type ScanAPIInspector struct{}

// compile time check: it should implement the interface
var _ Inspector = &ScanAPIInspector{}

// Name returns the name of the inspector.
func (sai *ScanAPIInspector) Name() string {
	return "scan-api"
}

// Inspect checks given JSON-RPC node url supports eth and net modules.
func (sai *ScanAPIInspector) Inspect(ctx context.Context, inspectionCfg InspectionConfig) (results *InspectionResults, resultErr error) {
	results = NewInspectionResults()
	results.Indicators = defaultIndicators(scanAPIIndicators)

	rpcClient, err := rpc.DialContext(ctx, inspectionCfg.ScanAPIURL)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't dial json-rpc api %w", err))

		results.Indicators[IndicatorScanAPIAccessible] = ResultFailure
		results.Indicators[IndicatorScanAPIModuleEth] = ResultFailure
		results.Indicators[IndicatorScanAPIModuleNet] = ResultFailure
		results.Indicators[IndicatorScanAPIChainID] = ResultFailure

		return
	} else {
		results.Indicators[IndicatorScanAPIAccessible] = ResultSuccess
	}

	if id, err := GetChainOrNetworkID(ctx, rpcClient); err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't query chain id: %v", err))
		results.Indicators[IndicatorScanAPIChainID] = ResultFailure
	} else {
		results.Indicators[IndicatorScanAPIChainID] = float64(id.Uint64())
	}

	err = checkSupportedScanApiModules(ctx, rpcClient, results)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("error checking module functionality %w", err))
	}

	// get configured block and include hash of the returned as metadata
	hash, err := GetBlockResponseHash(ctx, rpcClient, inspectionCfg.BlockNumber)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to get configured block %d: %v", inspectionCfg.BlockNumber, err))
	} else {
		results.Metadata[MetadataScanAPIBlockByNumberHash] = hash
	}

	if SupportsETH2(ctx, rpcClient) {
		results.Indicators[IndicatorScanAPIIsETH2] = ResultSuccess
	} else {
		results.Indicators[IndicatorScanAPIIsETH2] = ResultFailure
	}

	return
}

// checkSupportedModules double-checks the functionality of modules that were declared as supported by
// the node.
func checkSupportedScanApiModules(
	ctx context.Context, rpcClient *rpc.Client, results *InspectionResults,
) (resultError error) {
	// sends net_version under the hood. should prove the node supports net module
	_, err := GetNetworkID(ctx, rpcClient)
	if err != nil {
		results.Indicators[IndicatorScanAPIModuleNet] = ResultFailure
		resultError = multierror.Append(resultError, err)
	} else {
		results.Indicators[IndicatorScanAPIModuleNet] = ResultSuccess
	}

	// sends eth_chainId under the hood. should prove the node supports eth module
	_, err = GetChainID(ctx, rpcClient)
	if err != nil {
		resultError = multierror.Append(resultError, err)
		results.Indicators[IndicatorScanAPIModuleEth] = ResultFailure
	} else {
		results.Indicators[IndicatorScanAPIModuleEth] = ResultSuccess
	}

	return resultError
}
