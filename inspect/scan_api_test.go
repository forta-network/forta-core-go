package inspect

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/forta-network/forta-core-go/ethereum"
	mock_ethereum "github.com/forta-network/forta-core-go/ethereum/mocks"
	"github.com/forta-network/forta-core-go/registry"
	mock_registry "github.com/forta-network/forta-core-go/registry/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestScanAPIInspection(t *testing.T) {
	r := require.New(t)

	ctrl := gomock.NewController(t)
	rpcClient := mock_ethereum.NewMockRPCClient(ctrl)
	ethClient := mock_ethereum.NewMockEthClient(ctrl)
	regClient := mock_registry.NewMockClient(ctrl)

	RPCDialContext = func(ctx context.Context, rawurl string) (ethereum.RPCClient, error) {
		return rpcClient, nil
	}
	EthClientDialContext = func(ctx context.Context, rawurl string) (ethereum.EthClient, error) {
		return ethClient, nil
	}
	RegistryNewClient = func(ctx context.Context, cfg registry.ClientConfig) (registry.Client, error) {
		return regClient, nil
	}

	rpcClient.EXPECT().CallContext(gomock.Any(), gomock.Any(), "net_version").
		DoAndReturn(func(ctx interface{}, result interface{}, method interface{}, args ...interface{}) error {
			json.Unmarshal([]byte(`"0x5"`), result)
			return nil
		}).AnyTimes()
	rpcClient.EXPECT().CallContext(gomock.Any(), gomock.Any(), "eth_chainId").
		DoAndReturn(func(ctx interface{}, result interface{}, method interface{}, args ...interface{}) error {
			json.Unmarshal([]byte(`"0x5"`), result)
			return nil
		}).AnyTimes()

	// block response hash inspection
	rpcClient.EXPECT().CallContext(gomock.Any(), gomock.Any(), "eth_getBlockByNumber", gomock.Any()).
		DoAndReturn(func(ctx interface{}, result interface{}, method interface{}, args ...interface{}) error {
			json.Unmarshal([]byte(`"{}"`), result)
			return nil
		})

	// eth2 support inspection calls
	rpcClient.EXPECT().CallContext(gomock.Any(), gomock.Any(), "eth_getBlockByNumber", "latest", true).
		DoAndReturn(func(ctx interface{}, result interface{}, method interface{}, args ...interface{}) error {
			json.Unmarshal([]byte(`{"difficulty":"0x0","nonce":"0x0000000000000000"}`), result)
			return nil
		})

	inspector := &ScanAPIInspector{}
	results, err := inspector.Inspect(
		context.Background(), InspectionConfig{},
	)
	r.NoError(err)

	r.Equal(
		map[string]float64{
			IndicatorScanAPIAccessible: ResultSuccess,
			IndicatorScanAPIChainID:    float64(5),
			IndicatorScanAPIModuleEth:  ResultSuccess,
			IndicatorScanAPIModuleNet:  ResultSuccess,
			IndicatorScanAPIIsETH2:     ResultSuccess,
		}, results.Indicators,
	)

	r.NotEmpty(results.Metadata[MetadataScanAPIBlockByNumberHash])
}
