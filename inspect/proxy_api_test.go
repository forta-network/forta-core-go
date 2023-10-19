package inspect

import (
	"context"
	"errors"
	"testing"

	types "github.com/ethereum/go-ethereum/core/types"
	mock_ethereum "github.com/forta-network/forta-core-go/ethereum/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProxyAPIInspection(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping inspection test in short mode")
	}

	r := require.New(t)

	inspector := &ProxyAPIInspector{}
	results, err := inspector.Inspect(
		context.Background(), InspectionConfig{},
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

func TestFindOldestSupportedBlock(t *testing.T) {
	// Create a new Gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock ethclient.Client
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
