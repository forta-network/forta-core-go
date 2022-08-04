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
	IndicatorResourcesMemoryTotal      = "resources.memory.total"
	IndicatorResourcesMemoryAvailable  = "resources.memory.available"
	IndicatorResourcesStorageAvailable = "resources.storage.total"
	IndicatorResourcesStorageTotal     = "resources.storage.available"
	IndicatorResourcesCPUUsage         = "resources.cpu.usage"
	// IndicatorResourcesCPUBenchmark is the average duration of CPU stress test in nanoseconds. Less is better.
	IndicatorResourcesCPUBenchmark = "resources.cpu.benchmark"
)

var (
	resourceIndicators = []string{
		IndicatorResourcesMemoryTotal, IndicatorResourcesMemoryAvailable, IndicatorResourcesStorageAvailable,
		IndicatorResourcesStorageTotal, IndicatorResourcesCPUUsage, IndicatorResourcesCPUBenchmark,
	}
)

// SystemResourcesInspector is an inspector implementation.
type SystemResourcesInspector struct{}

// compile time check: it should implement the interface
var _ Inspector = &SystemResourcesInspector{}

// Name returns the name of the inspector.
func (sri *SystemResourcesInspector) Name() string {
	return "system-resources"
}

// Inspect inspects system resources.
func (sri *SystemResourcesInspector) Inspect(ctx context.Context, inspectionCfg InspectionConfig) (results *InspectionResults, resultErr error) {
	results = NewInspectionResults()
	results.Indicators = defaultIndicators(resourceIndicators)

	mi, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't read memory info: %w", err))
	} else {
		results.Indicators[IndicatorResourcesMemoryTotal] = float64(mi.Total)
		results.Indicators[IndicatorResourcesMemoryAvailable] = float64(mi.Available)
	}

	calcInterval := time.Second / 2

	ci, err := cpu.PercentWithContext(ctx, calcInterval, false)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't read cpu info: %w", err))
	} else {
		// cpu.Percent should return average result in a slice with a single item
		if len(ci) != 1 {
			results.Indicators[IndicatorResourcesCPUUsage] = ci[0]
		}
	}

	di, err := disk.UsageWithContext(ctx, "/")
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't read disk usage info: %w", err))
	} else {
		results.Indicators[IndicatorResourcesStorageTotal] = float64(di.Total)
		results.Indicators[IndicatorResourcesStorageAvailable] = float64(di.Free)
	}

	results.Indicators[IndicatorResourcesCPUBenchmark] = float64(cpuStressTestBenchmark())

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
