package inspect

import (
	"context"
	"fmt"
	"testing"
	"time"

	mock_ethereum "github.com/forta-network/forta-core-go/ethereum/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCalculateOffsetStats(t *testing.T) {
	// Create a test context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctrl := gomock.NewController(t)
	primaryClient := mock_ethereum.NewMockEthClient(ctrl)
	secondaryClient := mock_ethereum.NewMockEthClient(ctrl)

	// Test when everything is successful
	primaryClient.EXPECT().BlockNumber(gomock.Any()).Return(uint64(5), nil)
	primaryClient.EXPECT().BlockByNumber(gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()
	secondaryClient.EXPECT().BlockByNumber(gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()
	stats, err := calculateOffsetStats(ctx, primaryClient, secondaryClient)
	assert.NoError(t, err)
	assert.Equal(t, stats, offsetStats{0, 0, 0, 10})

	// Test when collectOffsetData returns an error
	primaryClient.EXPECT().BlockNumber(gomock.Any()).Return(uint64(0), fmt.Errorf("error"))

	stats, err = calculateOffsetStats(ctx, primaryClient, secondaryClient)
	assert.Error(t, err)
	assert.Equal(t, stats, offsetStats{})
}
