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
	ProxyAPI string `envconfig:"proxy_api" default:"https://cloudflare-eth.com"`
}

func init() {
	envconfig.MustProcess("test", &testProxyEnv)
}

func TestProxyAPIInspection(t *testing.T) {
	r := require.New(t)

	inspector := &ProxyAPIInspector{}
	results, err := inspector.Inspect(context.Background(), InspectionConfig{
		ProxyAPIURL: testProxyEnv.ProxyAPI,
		BlockNumber: testProxyAPIOldestSupportedBlock,
	})
	r.NoError(err)

	r.Equal(map[string]float64{
		IndicatorProxyAPIAccessible:     ResultSuccess,
		IndicatorProxyAPIChainID:        float64(1),
		IndicatorProxyAPIModuleWeb3:     ResultSuccess,
		IndicatorProxyAPIModuleEth:      ResultSuccess,
		IndicatorProxyAPIModuleNet:      ResultSuccess,
		IndicatorProxyAPIHistorySupport: VeryOldBlockNumber,
	}, results.Indicators)

	r.Equal(map[string]string{
		MetadataProxyAPIBlockByNumberHash: "3abe2f22edf2b463cbc13343a947be9ebbf8c16c2b50b2b90e10a199a2344f65",
	}, results.Metadata)
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
