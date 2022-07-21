package inspect

import (
	"context"
	"testing"

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

	recentBlockNumber := testGetRecentBlockNumber(r, testScanEnv.ScanAPI)
	inspectionCfg := InspectionConfig{
		ScanAPIURL:  testValidateEnv.ScanAPI,
		ProxyAPIURL: testValidateEnv.ScanAPI,
		TraceAPIURL: testValidateEnv.TraceAPI,
		BlockNumber: recentBlockNumber,
		CheckTrace:  true,
	}

	// make only scan and trace api inspections using the inspection config
	results, err := inspect(ctx, []Inspector{
		&ScanAPIInspector{},
		&TraceAPIInspector{},
	}, inspectionCfg)
	r.NoError(err)

	// validate the inspection using the same config
	validator, err := NewValidator(ctx, inspectionCfg)
	r.NoError(err)
	err = validator.Validate(ctx, results)
	r.NoError(err)
}

func TestValidateInspectionFail(t *testing.T) {
	ctx := context.Background()
	r := require.New(t)

	recentBlockNumber := testGetRecentBlockNumber(r, testScanEnv.ScanAPI)
	inspectionCfg1 := InspectionConfig{
		ScanAPIURL:  testValidateEnv.ScanAPI,
		ProxyAPIURL: testValidateEnv.ScanAPI,
		TraceAPIURL: testValidateEnv.TraceAPI,
		BlockNumber: recentBlockNumber,
	}

	// make only scan api inspection
	scanResults, err := inspect(ctx, []Inspector{
		&ScanAPIInspector{},
	}, inspectionCfg1)
	r.NoError(err)

	// now let's tamper with the initial conditions so trace inspection result is different
	inspectionCfg2 := InspectionConfig{
		ScanAPIURL:  testValidateEnv.ScanAPI,
		ProxyAPIURL: testValidateEnv.ScanAPI,
		TraceAPIURL: testValidateEnv.TraceAPI,
		BlockNumber: recentBlockNumber - 10,
		CheckTrace:  true,
	}

	// make only trace api inspection
	traceResults, err := inspect(ctx, []Inspector{
		&TraceAPIInspector{},
	}, inspectionCfg2)
	r.NoError(err)

	results := scanResults
	results.CopyFrom(traceResults)

	// validate the inspection using the first config
	validator, err := NewValidator(ctx, inspectionCfg1)
	r.NoError(err)
	err = validator.Validate(ctx, results)

	// expect error(s)
	r.Error(err)

	// expect no validator-side errors (10xx, 20xx)
	r.NotContains(err.Error(), "code: 10")
	r.NotContains(err.Error(), "code: 20")

	// scan inspection should be fine
	r.NotContains(err.Error(), ErrResultScanAPIBlockMismatch.Code())

	// trace inspection should have problems
	r.Contains(err.Error(), ErrResultBlockMismatch.Code())
	r.Contains(err.Error(), ErrResultTraceAPIBlockMismatch.Code())
	r.Contains(err.Error(), ErrResultTraceAPITraceBlockMismatch.Code())
}
