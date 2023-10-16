package inspect

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/require"
)

const (
	testProxyAPIOldestSupportedBlock = uint64(0)
)

var testProxyEnv struct {
	ProxyAPI string `envconfig:"proxy_api" default:"https://rpc.ankr.com/eth_goerli"`
}

func init() {
	envconfig.MustProcess("test", &testProxyEnv)
}

func TestProxyAPIInspection(t *testing.T) {
	r := require.New(t)

	inspector := &ProxyAPIInspector{}
	results, err := inspector.Inspect(
		context.Background(), InspectionConfig{
			ProxyAPIURL: testProxyEnv.ProxyAPI,
			ScanAPIURL:  testProxyEnv.ProxyAPI,
			BlockNumber: testProxyAPIOldestSupportedBlock,
		},
	)
	r.NoError(err)

	r.Equal(
		map[string]float64{
			IndicatorProxyAPIAccessible:     ResultSuccess,
			IndicatorProxyAPIChainID:        float64(5),
			IndicatorProxyAPIModuleWeb3:     ResultSuccess,
			IndicatorProxyAPIModuleEth:      ResultSuccess,
			IndicatorProxyAPIModuleNet:      ResultSuccess,
			IndicatorProxyAPIHistorySupport: VeryOldBlockNumber,
			IndicatorProxyAPIIsETH2:         ResultSuccess,
			// trick to make test less flaky and ignore offset issues
			IndicatorProxyAPIOffsetScanMax:     results.Indicators[IndicatorProxyAPIOffsetScanMax],
			IndicatorProxyAPIOffsetScanMean:    results.Indicators[IndicatorProxyAPIOffsetScanMean],
			IndicatorProxyAPIOffsetScanMedian:  results.Indicators[IndicatorProxyAPIOffsetScanMedian],
			IndicatorProxyAPIOffsetScanSamples: results.Indicators[IndicatorProxyAPIOffsetScanSamples],
		}, results.Indicators,
	)

	r.Equal(
		map[string]string{
			MetadataProxyAPIBlockByNumberHash: "7232705dbb71b9d5ef65891c2c6e7020137ffb652ed938a88621b322f09ab4a4",
		}, results.Metadata,
	)
}

func Test_findOldestSupportedBlock(t *testing.T) {
	r := require.New(t)

	cli, err := ethclient.Dial(testProxyEnv.ProxyAPI)
	r.NoError(err)

	ctx := context.Background()
	latestBlockNum, err := cli.BlockNumber(ctx)
	r.NoError(err)

	result := findOldestSupportedBlock(context.Background(), cli, 0, latestBlockNum)
	r.Equal(testProxyAPIOldestSupportedBlock, result)
}

func TestProxyAPIInspector_detectOffset(t *testing.T) {
	type args struct {
		ctx           context.Context
		inspectionCfg InspectionConfig
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "free and free combo",
			args: args{
				ctx: context.Background(),
				inspectionCfg: InspectionConfig{
					ScanAPIURL:  "https://rpc.ankr.com/eth",
					ProxyAPIURL: "https://rpc.ankr.com/eth",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				pai := &ProxyAPIInspector{}
				stats, err := pai.detectOffset(tt.args.ctx, tt.args.inspectionCfg)
				if err != nil {
					t.Log(err)
				}

				t.Log(stats)
			},
		)
	}
}
