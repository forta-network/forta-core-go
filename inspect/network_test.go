package inspect

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNetworkInspection(t *testing.T) {
	r := require.New(t)

	inspector := &NetworkInspector{}
	results, err := inspector.Inspect(context.Background(), InspectionConfig{})
	r.NoError(err)

	zero := float64(0)
	r.Equal(ResultSuccess, results.Metrics[MetricNetworkOutboundAccess])
	r.Greater(results.Metrics[MetricNetworkDownloadSpeed], zero, "download speed")
	r.Greater(results.Metrics[MetricNetworkUploadSpeed], zero, "upload speed")
}
