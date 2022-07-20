package inspect

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

const (
	testScanAPIURL                  = "https://cloudflare-eth.com"
	testScanAPIOldestSupportedBlock = uint64(0)
)

func TestScanAPIInspection(t *testing.T) {
	r := require.New(t)

	inspector := &ScanAPIInspector{}
	results, err := inspector.Inspect(context.Background(), InspectionConfig{
		ScanAPIURL: testScanAPIURL,
	})
	r.NoError(err)

	r.Equal(map[string]float64{
		MetricScanAPIAccessible:     ResultSuccess,
		MetricScanAPIChainID:        float64(1),
		MetricScanAPIModuleWeb3:     ResultSuccess,
		MetricScanAPIModuleEth:      ResultSuccess,
		MetricScanAPIModuleNet:      ResultSuccess,
		MetricScanAPIHistorySupport: VeryOldBlockNumber,
	}, results.Metrics)
}

func Test_findOldestSupportedBlock(t *testing.T) {
	r := require.New(t)

	cli, err := ethclient.Dial(testScanAPIURL)
	r.NoError(err)

	ctx := context.Background()
	latestBlockNum, err := cli.BlockNumber(ctx)
	r.NoError(err)

	result := findOldestSupportedBlock(context.Background(), cli, 0, latestBlockNum)
	r.Equal(testScanAPIOldestSupportedBlock, result)
}
