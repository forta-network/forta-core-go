package inspect

import (
	"context"
	"errors"
	"testing"

	types "github.com/ethereum/go-ethereum/core/types"
	mock_inspect "github.com/forta-network/forta-core-go/inspect/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFindOldestSupportedBlock(t *testing.T) {
	// Create a new Gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock ethclient.Client
	mockClient := mock_inspect.NewMockProxyAPIClient(ctrl)

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
