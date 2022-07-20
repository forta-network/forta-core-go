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
	r.Greater(results.Metrics[MetricResourcesMemoryTotal], zero)
	r.Greater(results.Metrics[MetricResourcesMemoryAvailable], zero)
	r.Greater(results.Metrics[MetricResourcesStorageTotal], zero)
	r.Greater(results.Metrics[MetricResourcesStorageAvailable], zero)
	r.Greater(results.Metrics[MetricResourcesCPUBenchmark], zero)
	// we do not know how to utilize CPU usage result yet so we don't test it here
}
