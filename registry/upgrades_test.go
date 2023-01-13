package registry_test

import (
	"context"
	"math/big"
	"os"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/forta-network/forta-core-go/domain"
	regdomain "github.com/forta-network/forta-core-go/domain/registry"
	"github.com/forta-network/forta-core-go/ethereum"
	"github.com/forta-network/forta-core-go/feeds"
	"github.com/forta-network/forta-core-go/registry"
	mock_registry "github.com/forta-network/forta-core-go/registry/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestListenToUpgrades(t *testing.T) {
	r := require.New(t)

	jsonRpc := os.Getenv("POLYGON_JSON_RPC")
	if jsonRpc == "" {
		jsonRpc = "https://polygon-rpc.com"
	}
	ctx := context.Background()

	ethClient, err := ethereum.NewStreamEthClient(context.Background(), "", jsonRpc)
	r.NoError(err)

	regClient := mock_registry.NewMockClient(gomock.NewController(t))

	blockFeed, err := feeds.NewBlockFeed(context.Background(), ethClient, nil, feeds.BlockFeedConfig{
		Start: big.NewInt(27061991),
		End:   big.NewInt(27061991),
	})
	r.NoError(err)

	errCh := registry.ListenToUpgrades(ctx, regClient, blockFeed)

	// create another subscription to ensure that the handlers are called
	doneCh := make(chan struct{})
	once := sync.Once{}
	blockFeed.Subscribe(func(evt *domain.BlockEvent) error {
		once.Do(func() {
			close(doneCh)
		})
		return nil
	})

	regClient.EXPECT().Contracts().Return(&registry.Contracts{
		Addresses: regdomain.RegistryContracts{
			ScannerNodeVersion: common.HexToAddress("0x4720c872425876B6f4b4E9130CDef667aDE553b2"),
		},
	})
	regClient.EXPECT().RefreshContracts().Return(nil)
	blockFeed.Start()
	r.Equal(feeds.ErrEndBlockReached, <-errCh)
	<-doneCh
}
