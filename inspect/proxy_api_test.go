package inspect

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"testing"

	types "github.com/ethereum/go-ethereum/core/types"
	"github.com/forta-network/forta-core-go/ethereum"
	mock_ethereum "github.com/forta-network/forta-core-go/ethereum/mocks"
	"github.com/forta-network/forta-core-go/registry"
	mock_registry "github.com/forta-network/forta-core-go/registry/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProxyAPIInspection(t *testing.T) {
	r := require.New(t)
	ctx := context.Background()

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
	rpcClient.EXPECT().CallContext(gomock.Any(), gomock.Any(), "web3_clientVersion").Return(nil)
	ethClient.EXPECT().BlockNumber(gomock.Any()).Return(uint64(123), nil)
	rpcClient.EXPECT().CallContext(gomock.Any(), gomock.Any(), "eth_getBlockByNumber", gomock.Any()).
		DoAndReturn(func(ctx interface{}, result interface{}, method interface{}, args ...interface{}) error {
			json.Unmarshal([]byte(`"{}"`), result)
			return nil
		})

	// oldest supported block inspection calls
	ethClient.EXPECT().BlockByNumber(gomock.Any(), big.NewInt(VeryOldBlockNumber)).Return(&types.Block{}, nil)

	// eth2 support inspection calls
	rpcClient.EXPECT().CallContext(gomock.Any(), gomock.Any(), "eth_getBlockByNumber", "latest", true).
		DoAndReturn(func(ctx interface{}, result interface{}, method interface{}, args ...interface{}) error {
			json.Unmarshal([]byte(`{"difficulty":"0x0","nonce":"0x0000000000000000"}`), result)
			return nil
		})

	// offset inspection calls
	ethClient.EXPECT().BlockNumber(gomock.Any()).Return(uint64(123), nil).AnyTimes()
	ethClient.EXPECT().BlockByNumber(gomock.Any(), gomock.Any()).Return(&types.Block{}, nil).AnyTimes()

	inspector := &ProxyAPIInspector{}
	results, err := inspector.Inspect(
		ctx, InspectionConfig{},
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
			MetadataProxyAPIBlockByNumberHash: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		}, results.Metadata,
	)
}

func TestFindOldestSupportedBlock(t *testing.T) {
	// Create a new Gomock controller
	ctrl := gomock.NewController(t)
	mockClient := mock_ethereum.NewMockEthClient(ctrl)

	// Create a test context
	ctx := context.Background()

	// Test when the earliest block is found
	mockClient.EXPECT().BlockByNumber(ctx, gomock.Any()).Return(&types.Block{}, nil).Times(7)

	oldestBlock := findOldestSupportedBlock(ctx, mockClient, 0, 200)
	assert.Equal(t, uint64(0), oldestBlock)

	// Test when the earliest block is not found
	mockClient.EXPECT().BlockByNumber(ctx, gomock.Any()).Return(nil, errors.New("block not found")).Times(2)
	mockClient.EXPECT().BlockByNumber(ctx, gomock.Any()).Return(&types.Block{}, nil).Times(5)
	oldestBlock = findOldestSupportedBlock(ctx, mockClient, 0, 200)
	assert.Equal(t, uint64(151), oldestBlock)

	// Test when the client returns an error for all blocks
	mockClient.EXPECT().BlockByNumber(ctx, gomock.Any()).Return(nil, errors.New("block not found")).AnyTimes()

	oldestBlock = findOldestSupportedBlock(ctx, mockClient, 0, 200)
	assert.Equal(t, uint64(200), oldestBlock)
}
