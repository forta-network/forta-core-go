package inspect

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/sync/errgroup"
)

const (
	// IndicatorOffsetScanProxySynced no offset between scan and proxy
	IndicatorOffsetScanProxySynced = "offset.synced.scan-proxy"
	IndicatorOffsetScanTraceSynced = "offset.synced.scan-trace"

	IndicatorOffsetScanBlockNumber  = "offset.scan.block-number"
	IndicatorOffsetProxyBlockNumber = "offset.proxy.block-number"
	IndicatorOffsetTraceBlockNumber = "offset.trace.block-number"
)

var (
	offsetIndicators = []string{
		IndicatorOffsetScanProxySynced, IndicatorOffsetScanTraceSynced, IndicatorOffsetScanBlockNumber,
		IndicatorOffsetProxyBlockNumber, IndicatorOffsetTraceBlockNumber,
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

	// check scan and proxy are in sync
	err = checkJsonRPCOffset(ctx, scanClient, proxyClient)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("detected offset between scan api and proxy api: %v", err))
		results.Indicators[IndicatorOffsetScanProxySynced] = ResultFailure
	} else {
		results.Indicators[IndicatorOffsetScanProxySynced] = ResultSuccess
	}

	if inspectionCfg.CheckTrace {
		traceRPCClient, err := rpc.DialContext(ctx, inspectionCfg.TraceAPIURL)
		if err != nil {
			resultErr = multierror.Append(resultErr, fmt.Errorf("can't dial trace json-rpc api %w", err))
			return
		}

		traceClient := ethclient.NewClient(traceRPCClient)

		// check scan and trace are in sync
		err = checkJsonRPCOffset(ctx, scanClient, traceClient)
		if err != nil {
			resultErr = multierror.Append(
				resultErr, fmt.Errorf("detected offset between scan api and trace api: %v", err),
			)
			results.Indicators[IndicatorOffsetScanTraceSynced] = ResultFailure
		} else {
			results.Indicators[IndicatorOffsetScanTraceSynced] = ResultSuccess
		}
	}

	scanBlock, proxyBlock, traceBlock, err := pai.fetchLatestBlocks(ctx, inspectionCfg)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	} else {
		results.Indicators[IndicatorOffsetScanBlockNumber] = float64(scanBlock)
		results.Indicators[IndicatorOffsetProxyBlockNumber] = float64(proxyBlock)
		results.Indicators[IndicatorOffsetTraceBlockNumber] = float64(traceBlock)
	}
	return
}

func checkJsonRPCOffset(
	ctx context.Context, scanClient *ethclient.Client, secondaryClient *ethclient.Client,
) error {
	interval := time.Millisecond * 300
	t := time.NewTicker(interval)

	// to catch the beginning of recently mined block, since offsets happen at that period
	ctx, cancel := context.WithTimeout(ctx, time.Second*20)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-t.C:
			currentHeight, err := scanClient.BlockNumber(ctx)
			if err != nil {
				if !errors.Is(err, context.Canceled) && !errors.Is(err, context.DeadlineExceeded) {
					return err
				}
			}

			// ask for the latest block of the scan api to api
			_, err = secondaryClient.BlockByNumber(ctx, big.NewInt(int64(currentHeight)))
			if err != nil {
				if !errors.Is(err, context.Canceled) && !errors.Is(err, context.DeadlineExceeded) {
					return fmt.Errorf("can't fetch %d: %v", currentHeight, err)
				}
			}

		}
	}
}

func (pai *OffsetInspector) fetchLatestBlocks(ctx context.Context, inspectionCfg InspectionConfig) (
	uint64, uint64, uint64, error,
) {
	var resultErr error

	proxyRPCClient, err := rpc.DialContext(ctx, inspectionCfg.ProxyAPIURL)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't dial proxy json-rpc api %w", err))
		return 0, 0, 0, resultErr
	}
	proxyClient := ethclient.NewClient(proxyRPCClient)

	scanRPCClient, err := rpc.DialContext(ctx, inspectionCfg.ScanAPIURL)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't dial scan json-rpc api %w", err))
		return 0, 0, 0, resultErr
	}
	scanClient := ethclient.NewClient(scanRPCClient)

	g, ctx := errgroup.WithContext(ctx)

	var (
		proxyBlock, scanBlock, traceBlock uint64
	)

	// fetch latest block for proxy api
	g.Go(
		func() (err error) {
			proxyBlock, err = proxyClient.BlockNumber(ctx)
			return
		},
	)

	// fetch latest block for scan api
	g.Go(
		func() (err error) {
			scanBlock, err = scanClient.BlockNumber(ctx)
			return
		},
	)

	if inspectionCfg.CheckTrace {
		traceRPCClient, err := rpc.DialContext(ctx, inspectionCfg.TraceAPIURL)
		if err != nil {
			resultErr = multierror.Append(resultErr, fmt.Errorf("can't dial trace json-rpc api %w", err))
			return 0, 0, 0, resultErr
		}
		traceClient := ethclient.NewClient(traceRPCClient)
		// fetch latest block for trace api
		g.Go(
			func() (err error) {
				traceBlock, err = traceClient.BlockNumber(ctx)
				return
			},
		)
	}

	err = g.Wait()
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}

	return scanBlock, proxyBlock, traceBlock, resultErr
}