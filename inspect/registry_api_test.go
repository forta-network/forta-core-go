package inspect

import (
	"context"
	"testing"

	"github.com/forta-network/forta-core-go/ethereum"
	mock_ethereum "github.com/forta-network/forta-core-go/ethereum/mocks"
	"github.com/forta-network/forta-core-go/registry"
	mock_registry "github.com/forta-network/forta-core-go/registry/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRegistryAPIInspection(t *testing.T) {
	r := require.New(t)

	scannerAddr := "0x3DC45b47B7559Ca3b231E5384D825F9B461A0398"

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

	ethClient.EXPECT().Close()
	regClient.EXPECT().Close()

	regClient.EXPECT().GetAssignmentHash(scannerAddr).Return(&registry.AssignmentHash{}, nil)

	inspector := &RegistryAPIInspector{}
	results, err := inspector.Inspect(
		context.Background(), InspectionConfig{
			ScannerAddress: scannerAddr,
		},
	)
	r.NoError(err)

	r.Equal(
		map[string]float64{
			IndicatorRegistryAPIAccessible:  ResultSuccess,
			IndicatorRegistryAPIENS:         ResultSuccess,
			IndicatorRegistryAPIAssignments: ResultSuccess,
		}, results.Indicators,
	)
}
