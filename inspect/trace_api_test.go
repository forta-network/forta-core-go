package inspect

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/require"
)

const (
	testTraceAPIURL = "https://rpcapi-tracing.fantom.network/"
)

func TestTraceAPIInspection(t *testing.T) {
	var env struct {
		TraceAPI string `envconfig:"trace_api"`
	}
	envconfig.MustProcess("test", &env)
	if len(env.TraceAPI) == 0 {
		t.Log("TEST_TRACE_API not specified - aborting")
		return
	}

	r := require.New(t)

	client, err := ethclient.Dial(testTraceAPIURL)
	r.NoError(err)
	block, err := client.BlockByNumber(context.Background(), nil)
	r.NoError(err)

	inspector := &TraceAPIInspector{}
	results, err := inspector.Inspect(context.Background(), InspectionConfig{
		TraceAPIURL: testTraceAPIURL,
		BlockNumber: block.NumberU64() - 20,
	})
	r.NoError(err)

	r.Equal(map[string]float64{
		MetricTraceAccessible: ResultSuccess,
		MetricTraceSupported:  ResultSuccess,
	}, results.Metrics)

	r.NotEmpty(results.Metadata[MetadataTraceAPIBlockByNumberHash])
	r.NotEmpty(results.Metadata[MetadataTraceAPITraceBlockHash])
}
