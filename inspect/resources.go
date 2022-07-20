package inspect

import (
	"context"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

const (
	MetricResourcesMemoryTotal      = "resources.memory.total"
	MetricResourcesMemoryAvailable  = "resources.memory.available"
	MetricResourcesStorageAvailable = "resources.storage.total"
	MetricResourcesStorageTotal     = "resources.storage.available"
	MetricResourcesCPUUsage         = "resources.cpu.usage"
	// MetricResourcesCPUBenchmark average nanoseconds needed for cpu stress test. Less is better.
	MetricResourcesCPUBenchmark = "resources.cpu.benchmark"
)

// SystemResourcesInspector is an inspector implementation.
type SystemResourcesInspector struct{}

var _ Inspector = &SystemResourcesInspector{}

// Name returns the name of the inspector.
func (sri *SystemResourcesInspector) Name() string {
	return "system-resources"
}

// Inspect inspects system resources.
func (sri *SystemResourcesInspector) Inspect(ctx context.Context, inspectionCfg InspectionConfig) (results *InspectionResults, resultErr error) {
	results = NewInspectionResults()

	mi, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't read memory info: %w", err))
	} else {
		results.Metrics[MetricResourcesMemoryTotal] = float64(mi.Total)
		results.Metrics[MetricResourcesMemoryAvailable] = float64(mi.Available)
	}

	calcInterval := time.Second / 2

	ci, err := cpu.PercentWithContext(ctx, calcInterval, false)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't read cpu info: %w", err))
	} else {
		// cpu.Percent should return average result in a slice with a single item
		if len(ci) != 1 {
			results.Metrics[MetricResourcesCPUUsage] = ci[0]
		}
	}

	di, err := disk.UsageWithContext(ctx, "/")
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't read disk usage info: %w", err))
	} else {
		results.Metrics[MetricResourcesStorageTotal] = float64(di.Total)
		results.Metrics[MetricResourcesStorageAvailable] = float64(di.Free)
	}

	results.Metrics[MetricResourcesCPUBenchmark] = float64(cpuStressTestBenchmark())

	return
}

func cpuStressTestBenchmark() int64 {
	result := testing.Benchmark(
		func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				firstNPrimes(1e5)
			}
		},
	)

	return result.NsPerOp()
}

// firstNPrimes finds first n primes. algorithm doesn't need to be smart.
// doesn't keep track of them either, with less memory it is clearer to see the cpu load.
func firstNPrimes(n int) {
	iterator := 2
	counter := 0

	for counter <= n {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(iterator))); i++ {
			if iterator%i == 0 {
				isPrime = false

				break
			}
		}

		if isPrime {
			counter++
		}

		iterator++
	}
}
