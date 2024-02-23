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

	r.Equal(ResultSuccess, results.Indicators[IndicatorNetworkOutboundAccess])
	// r.Greater(results.Indicators[IndicatorNetworkDownloadSpeed], zero, "download speed")
	// r.Greater(results.Indicators[IndicatorNetworkUploadSpeed], zero, "upload speed")

	r.Equal(map[string]string{}, results.Metadata)
}
