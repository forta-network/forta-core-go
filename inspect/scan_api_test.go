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
		ScanAPIURL:  testScanAPIURL,
		BlockNumber: testScanAPIOldestSupportedBlock,
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

	r.Equal(map[string]string{
		MetadataScanAPIBlockByNumberHash: "3abe2f22edf2b463cbc13343a947be9ebbf8c16c2b50b2b90e10a199a2344f65",
	}, results.Metadata)
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
