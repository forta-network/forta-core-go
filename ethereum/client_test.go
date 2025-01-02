package ethereum

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/goccy/go-json"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/forta-network/forta-core-go/domain"
	mock_domain "github.com/forta-network/forta-core-go/domain/mocks"
	mocks "github.com/forta-network/forta-core-go/ethereum/mocks"
	"github.com/forta-network/forta-core-go/ethereum/provider"
)

const testBlockHash = "0x4fc0862e76691f5312964883954d5c2db35e2b8f7a4f191775a4f50c69804a8d"

var testErr = errors.New("test err")

func initClient(t *testing.T) (*streamEthClient, *mocks.MockSubscriber, context.Context) {
	r := require.New(t)

	minBackoff = 1 * time.Millisecond
	maxBackoff = 1 * time.Millisecond
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	client := mocks.NewMockSubscriber(ctrl)

	ethClient, err := NewStreamEthClientWithRPCClient(context.Background(), "", true, client)
	r.NoError(err)

	return ethClient, client, ctx
}

func TestEthClient_BlockByHash(t *testing.T) {
	r := require.New(t)

	minBackoff = 1 * time.Millisecond
	maxBackoff = 1 * time.Millisecond
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	client1 := mocks.NewMockSubscriber(ctrl)
	client2 := mocks.NewMockSubscriber(ctrl)

	ethClient := &streamEthClient{
		rpcClientProvider: provider.NewRingProvider(Subscriber(client1), Subscriber(client2)),
		retryInterval:     defaultRetryInterval,
	}
	hash := testBlockHash
	client1.EXPECT().CallContext(gomock.Any(), gomock.Any(), blocksByHash, testBlockHash).Return(testErr)
	client2.EXPECT().CallContext(gomock.Any(), gomock.Any(), blocksByHash, testBlockHash).Return(testErr)
	client1.EXPECT().CallContext(gomock.Any(), gomock.Any(), blocksByHash, testBlockHash).Return(testErr)
	client2.EXPECT().CallContext(gomock.Any(), gomock.Any(), blocksByHash, testBlockHash).DoAndReturn(func(ctx context.Context, result interface{}, method string, args ...interface{}) error {
		b, _ := json.Marshal(domain.Block{Hash: hash})
		return json.Unmarshal(b, result)
	}).Times(1)

	res, err := ethClient.BlockByHash(ctx, testBlockHash)
	r.NoError(err)
	r.Equal(hash, res.Hash)
}

func TestEthClient_SubscribeToHeader_Err(t *testing.T) {
	r := require.New(t)

	ethClient, client, ctx := initClient(t)
	sub := mock_domain.NewMockClientSubscription(gomock.NewController(t))

	client.EXPECT().Subscribe(gomock.Any(), "eth", gomock.Any(), "newHeads").Return(sub, nil).Times(2)
	errCh := make(chan error, 1)
	errCh <- errors.New("subscription encountered some error")
	sub.EXPECT().Err().Return(errCh).Times(2)

	headerCh, err := ethClient.SubscribeToHead(ctx)
	r.NoError(err)
	<-headerCh // should close for the first time

	headerCh, err = ethClient.SubscribeToHead(ctx)
	r.NoError(err)
	var blocked bool
	select {
	case <-time.After(time.Second):
		// should continue from here
		blocked = true
	case <-headerCh: // should block
	}
	r.True(blocked)
}
