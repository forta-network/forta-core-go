package inspect

import (
	"context"
	"testing"

	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/require"
)

var testScanEnv struct {
	ScanAPI string `envconfig:"scan_api" default:"https://cloudflare-eth.com"`
}

func init() {
	envconfig.MustProcess("test", &testScanEnv)
}

func TestScanAPIInspection(t *testing.T) {
	r := require.New(t)

	recentBlockNumber := testGetRecentBlockNumber(r, testScanEnv.ScanAPI)

	inspector := &ScanAPIInspector{}
	results, err := inspector.Inspect(context.Background(), InspectionConfig{
		ScanAPIURL:  testScanEnv.ScanAPI,
		BlockNumber: recentBlockNumber,
	})
	r.NoError(err)

	r.Equal(map[string]float64{
		IndicatorScanAPIAccessible: ResultSuccess,
		IndicatorScanAPIChainID:    float64(1),
		IndicatorScanAPIModuleEth:  ResultSuccess,
		IndicatorScanAPIModuleNet:  ResultSuccess,
	}, results.Indicators)

	r.NotEmpty(results.Metadata[MetadataScanAPIBlockByNumberHash])
}
