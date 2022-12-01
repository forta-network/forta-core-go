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
	ProxyAPI string `envconfig:"proxy_api" default:"https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"`
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
		name        string
		args        args
	}{
		{
			name: "different apis",
			args: args{
				ctx: context.Background(),
				inspectionCfg: InspectionConfig{
					ScanAPIURL:  "http://172.31.84.3:8545",
					ProxyAPIURL: "http://172.31.71.62:8545",
				},
			},
		}, {
			name: "dns",
			args: args{
				ctx: context.Background(),
				inspectionCfg: InspectionConfig{
					ScanAPIURL:  "http://erigon-eth-prod.forta.internal:8545",
					ProxyAPIURL: "http://erigon-eth-prod.forta.internal:8545",
				},
			},
		}, {
			name: "same node",
			args: args{
				ctx: context.Background(),
				inspectionCfg: InspectionConfig{
					ScanAPIURL:  "http://172.31.84.3:8545",
					ProxyAPIURL: "http://172.31.84.3:8545",
				},
			},
		}, {
			name: "free and full combo",
			args: args{
				ctx: context.Background(),
				inspectionCfg: InspectionConfig{
					ScanAPIURL:  "https://cloudflare-eth.com",
					ProxyAPIURL: "http://erigon-eth-prod.forta.internal:8545",
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