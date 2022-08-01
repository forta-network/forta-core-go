package validation

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/forta-network/forta-core-go/inspect"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/require"
)

var testValidateEnv struct {
	ScanAPI  string `envconfig:"scan_api" default:"https://rpcapi.fantom.network"`
	TraceAPI string `envconfig:"trace_api" default:"https://rpcapi-tracing.fantom.network"`
}

func init() {
	envconfig.MustProcess("test", &testValidateEnv)
}

func TestValidateInspectionSuccess(t *testing.T) {
	ctx := context.Background()
	r := require.New(t)

	recentBlockNumber := testGetRecentBlockNumber(r, testValidateEnv.ScanAPI)
	inspectionCfg := inspect.InspectionConfig{
		ScanAPIURL:  testValidateEnv.ScanAPI,
		ProxyAPIURL: testValidateEnv.ScanAPI,
		TraceAPIURL: testValidateEnv.TraceAPI,
		BlockNumber: recentBlockNumber,
		CheckTrace:  true,
	}

	// make only scan and trace api inspections using the inspection config
	results, err := inspect.InspectAll(ctx, []inspect.Inspector{
		&inspect.ScanAPIInspector{},
		&inspect.ProxyAPIInspector{},
		&inspect.TraceAPIInspector{},
	}, inspectionCfg)
	r.NoError(err)

	// validate the inspection using the same config
	validator, err := NewValidator(ctx, inspectionCfg)
	r.NoError(err)
	_, err = validator.Validate(ctx, results)
	r.NoError(err)
}

func TestValidateInspectionFail(t *testing.T) {
	ctx := context.Background()
	r := require.New(t)

	recentBlockNumber := testGetRecentBlockNumber(r, testValidateEnv.ScanAPI)
	inspectionCfg1 := inspect.InspectionConfig{
		ScanAPIURL:  testValidateEnv.ScanAPI,
		ProxyAPIURL: testValidateEnv.ScanAPI,
		TraceAPIURL: testValidateEnv.TraceAPI,
		BlockNumber: recentBlockNumber,
		CheckTrace:  true,
	}

	// make only scan api inspection
	results, err := inspect.InspectAll(ctx, []inspect.Inspector{
		&inspect.ScanAPIInspector{},
		&inspect.ProxyAPIInspector{},
	}, inspectionCfg1)
	r.NoError(err)

	// now let's tamper with the initial conditions so trace inspection result is different
	inspectionCfg2 := inspect.InspectionConfig{
		ScanAPIURL:  testValidateEnv.ScanAPI,
		ProxyAPIURL: testValidateEnv.ScanAPI,
		TraceAPIURL: testValidateEnv.TraceAPI,
		BlockNumber: recentBlockNumber - 10,
		CheckTrace:  true,
	}

	// make only trace api inspection
	traceResults, err := inspect.InspectAll(ctx, []inspect.Inspector{
		&inspect.TraceAPIInspector{},
	}, inspectionCfg2)
	r.NoError(err)
	results.CopyFrom(traceResults)

	// validate the inspection using the first config
	validator, err := NewValidator(ctx, inspectionCfg1)
	r.NoError(err)
	verrs, err := validator.Validate(ctx, results)

	// expect error(s)
	r.Error(err)

	// expect no validator-side errors (10xx, 20xx)
	r.NotContains(err.Error(), "code: 10")
	r.NotContains(err.Error(), "code: 20")

	// scan inspection should be fine
	r.False(verrs.HasCode(ErrResultScanAPIBlockMismatch.Code()))

	// trace inspection should have problems
	r.True(verrs.HasCode(ErrResultBlockMismatch.Code()))
	r.True(verrs.HasCode(ErrResultTraceAPIBlockMismatch.Code()))
	r.True(verrs.HasCode(ErrResultTraceAPITraceBlockMismatch.Code()))
}

func testGetRecentBlockNumber(r *require.Assertions, apiURL string) uint64 {
	client, err := ethclient.Dial(apiURL)
	r.NoError(err)
	block, err := client.BlockByNumber(context.Background(), nil)
	r.NoError(err)
	return block.NumberU64() - 20
}
