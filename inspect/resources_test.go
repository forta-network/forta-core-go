package inspect

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSystemResourcesInspection(t *testing.T) {
	r := require.New(t)

	inspector := &SystemResourcesInspector{}
	results, err := inspector.Inspect(context.Background(), InspectionConfig{})
	r.NoError(err)

	zero := float64(0)
	r.Greater(results.Indicators[IndicatorResourcesMemoryTotal], zero)
	r.Greater(results.Indicators[IndicatorResourcesMemoryAvailable], zero)
	r.Greater(results.Indicators[IndicatorResourcesStorageTotal], zero)
	r.Greater(results.Indicators[IndicatorResourcesStorageAvailable], zero)
	r.Greater(results.Indicators[IndicatorResourcesCPUBenchmark], zero)
	// we do not know how to utilize CPU usage result yet so we don't test it here

	r.Equal(map[string]string{}, results.Metadata)
}
