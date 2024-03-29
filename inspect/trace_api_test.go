package inspect

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"
	"time"

	"github.com/forta-network/forta-core-go/ethereum"
	mock_ethereum "github.com/forta-network/forta-core-go/ethereum/mocks"
	"github.com/forta-network/forta-core-go/registry"
	mock_registry "github.com/forta-network/forta-core-go/registry/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestTraceAPIInspection(t *testing.T) {
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

	rpcClient.EXPECT().Close()

	rpcClient.EXPECT().CallContext(gomock.Any(), gomock.Any(), "net_version").
		DoAndReturn(func(ctx interface{}, result interface{}, method interface{}, args ...interface{}) error {
			_ = json.Unmarshal([]byte(`"4002"`), result)
			return nil
		}).AnyTimes()
	rpcClient.EXPECT().CallContext(gomock.Any(), gomock.Any(), "eth_chainId").
		DoAndReturn(func(ctx interface{}, result interface{}, method interface{}, args ...interface{}) error {
			_ = json.Unmarshal([]byte(`"4002"`), result)
			return nil
		}).AnyTimes()

	// trace response hash inspection
	rpcClient.EXPECT().CallContext(gomock.Any(), gomock.Any(), "trace_block", gomock.Any()).
		DoAndReturn(func(ctx interface{}, result interface{}, method interface{}, args ...interface{}) error {
			_ = json.Unmarshal([]byte(`"{}"`), result)
			return nil
		})

	// block response hash inspection
	rpcClient.EXPECT().CallContext(gomock.Any(), gomock.Any(), "eth_getBlockByNumber", gomock.Any()).
		DoAndReturn(func(ctx interface{}, result interface{}, method interface{}, args ...interface{}) error {
			_ = json.Unmarshal([]byte(`"{}"`), result)
			return nil
		})

	inspector := &TraceAPIInspector{}
	results, err := inspector.Inspect(
		context.Background(), InspectionConfig{
			CheckTrace: true,
		},
	)
	r.NoError(err)

	r.Equal(
		map[string]float64{
			IndicatorTraceAccessible: ResultSuccess,
			IndicatorTraceSupported:  ResultSuccess,
			IndicatorTraceAPIChainID: 4002,
			IndicatorTraceAPIIsETH2:  ResultFailure,
		}, results.Indicators,
	)

	r.NotEmpty(results.Metadata[MetadataTraceAPIBlockByNumberHash])
	r.NotEmpty(results.Metadata[MetadataTraceAPITraceBlockHash])
}

func TestTraceFailure(t *testing.T) {
	if os.Getenv("TRACE_FAILURE_TEST") != "1" {
		t.Skip("skipping trace failure test")
	}

	r := require.New(t)
	ctx := context.Background()

	url := os.Getenv("JSON_RPC_URL")

	proxyClient, err := EthClientDialContext(ctx, url)
	r.NoError(err)

	currentHeight, err := proxyClient.BlockNumber(ctx)
	r.NoError(err)

	for {
		results, resultErr := (&ScanAPIInspector{}).Inspect(ctx, InspectionConfig{
			ScanAPIURL:  url,
			ProxyAPIURL: url,
			TraceAPIURL: url,
			BlockNumber: currentHeight,
			CheckTrace:  true,
		})
		log.Println(results)
		log.Println(resultErr)

		results, resultErr = (&ProxyAPIInspector{}).Inspect(ctx, InspectionConfig{
			ScanAPIURL:  url,
			ProxyAPIURL: url,
			TraceAPIURL: url,
			BlockNumber: currentHeight,
			CheckTrace:  true,
		})
		log.Println(results)
		log.Println(resultErr)

		results, resultErr = (&TraceAPIInspector{}).Inspect(ctx, InspectionConfig{
			TraceAPIURL: url,
			BlockNumber: currentHeight,
			CheckTrace:  true,
		})
		log.Println(results)
		log.Println(resultErr)

		if results.Indicators[IndicatorTraceAPIIsETH2] == ResultFailure {
			log.Panic("trace api eth2 check failed")
		}

		time.Sleep(time.Second * 10)
	}
}
