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
			IndicatorProxyAPIUpToDate:       ResultSuccess,
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
