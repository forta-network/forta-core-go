package inspect

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testTraceAPIURL = "https://cloudflare-eth.com"
)

func TestTraceAPIInspection(t *testing.T) {
	r := require.New(t)

	inspector := &TraceAPIInspector{}
	results, err := inspector.Inspect(context.Background(), InspectionConfig{
		TraceAPIURL: testTraceAPIURL,
	})
	r.Error(err)

	r.Equal(map[string]float64{
		MetricTraceAccessible: ResultSuccess,
		MetricTraceSupported:  ResultFailure,
	}, results.Metrics)
	r.Equal(map[string]string{}, results.Metadata)
}
