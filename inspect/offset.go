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
	// IndicatorOffsetScanProxySynced no offset between scan and proxy
	IndicatorOffsetScanProxySynced = "offset.synced.scan-proxy"
)

var (
	offsetIndicators = []string{
		IndicatorOffsetScanProxySynced,
	}
)

// OffsetInspector is an inspector implementation.
type OffsetInspector struct{}

// compile time check: it should implement the interface
var _ Inspector = &OffsetInspector{}

// Name returns the name of the inspector.
func (pai *OffsetInspector) Name() string {
	return "offset"
}

// Inspect checks given JSON-RPC, Scan and Trace APIs don't have an offset.
//
// it doesn't actually return any errors for now,
// because the point is to keep going and check if it supports the rest
// error return parameter is simply for keeping the function extensible without api changes in the future.
func (pai *OffsetInspector) Inspect(ctx context.Context, inspectionCfg InspectionConfig) (
	results *InspectionResults, resultErr error,
) {
	results = NewInspectionResults()
	results.Indicators = defaultIndicators(offsetIndicators)

	proxyRPCClient, err := rpc.DialContext(ctx, inspectionCfg.ProxyAPIURL)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't dial proxy json-rpc api %w", err))
		return
	}

	scanRPCClient, err := rpc.DialContext(ctx, inspectionCfg.ScanAPIURL)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't dial scan json-rpc api %w", err))
		return
	}

	proxyClient := ethclient.NewClient(proxyRPCClient)
	scanClient := ethclient.NewClient(scanRPCClient)

	currentHeight, err := scanClient.BlockNumber(ctx)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}

	// ask for the latest block of the scan api to proxy api
	blockInfo, err := proxyClient.BlockByNumber(ctx, big.NewInt(int64(currentHeight)))
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
		results.Indicators[IndicatorOffsetScanProxySynced] = ResultFailure
	} else {
		results.Indicators[IndicatorOffsetScanProxySynced] = ResultSuccess
	}

	if blockInfo == nil {
		results.Indicators[IndicatorOffsetScanProxySynced] = ResultFailure
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't get %d from proxy json-rpc", currentHeight))
	} else {
		results.Indicators[IndicatorOffsetScanProxySynced] = ResultSuccess
	}

	return
}
