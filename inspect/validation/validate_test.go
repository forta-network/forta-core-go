package validation

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/forta-network/forta-core-go/ethereum"
	mock_ethereum "github.com/forta-network/forta-core-go/ethereum/mocks"
	"github.com/forta-network/forta-core-go/inspect"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestValidateInspectionSuccess(t *testing.T) {
	ctx := context.Background()
	r := require.New(t)

	ctrl := gomock.NewController(t)
	rpcClient := mock_ethereum.NewMockRPCClient(ctrl)

	inspect.RPCDialContext = func(ctx context.Context, rawurl string) (ethereum.RPCClient, error) {
		return rpcClient, nil
	}

	inspectionCfg := inspect.InspectionConfig{
		CheckTrace: true,
	}

	expectedHash := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

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
		}).Times(3)

	// validate the inspection using the same config
	validator, err := NewValidator(ctx, inspectionCfg)
	r.NoError(err)
	_, err = validator.Validate(ctx, &inspect.InspectionResults{
		Inputs: inspect.InspectionConfig{
			BlockNumber: 10,
			CheckTrace:  true,
		},
		Metadata: map[string]string{
			inspect.MetadataScanAPIBlockByNumberHash:  expectedHash,
			inspect.MetadataProxyAPIBlockByNumberHash: expectedHash,
			inspect.MetadataTraceAPIBlockByNumberHash: expectedHash,
			inspect.MetadataTraceAPITraceBlockHash:    expectedHash,
		},
	})
	r.NoError(err)
}

func TestValidateInspectionFail(t *testing.T) {
	ctx := context.Background()
	r := require.New(t)

	ctrl := gomock.NewController(t)
	rpcClient := mock_ethereum.NewMockRPCClient(ctrl)

	inspect.RPCDialContext = func(ctx context.Context, rawurl string) (ethereum.RPCClient, error) {
		return rpcClient, nil
	}

	inspectionCfg := inspect.InspectionConfig{
		CheckTrace: true,
	}

	expectedHash := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	unexpectedHash := "foobar"

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
		}).Times(3)

	// validate the inspection using the first config
	validator, err := NewValidator(ctx, inspectionCfg)
	r.NoError(err)
	verrs, err := validator.Validate(ctx, &inspect.InspectionResults{
		Inputs: inspect.InspectionConfig{
			BlockNumber: 10,
			CheckTrace:  true,
		},
		Metadata: map[string]string{
			inspect.MetadataScanAPIBlockByNumberHash:  expectedHash,
			inspect.MetadataProxyAPIBlockByNumberHash: expectedHash,
			inspect.MetadataTraceAPIBlockByNumberHash: unexpectedHash,
			inspect.MetadataTraceAPITraceBlockHash:    unexpectedHash,
		},
	})

	// expect error(s)
	r.Error(err)

	// expect no validator-side errors (10xx, 20xx)
	r.NotContains(err.Error(), "code: 10")
	r.NotContains(err.Error(), "code: 20")

	// scan inspection should be fine
	r.False(verrs.HasCode(inspect.ErrResultScanAPIBlockMismatch.Code()))

	// trace inspection should have problems
	r.True(verrs.HasCode(inspect.ErrResultBlockMismatch.Code()))
	r.True(verrs.HasCode(inspect.ErrResultTraceAPIBlockMismatch.Code()))
	r.True(verrs.HasCode(inspect.ErrResultTraceAPITraceBlockMismatch.Code()))
}
