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
	MetricContainerResourceMemoryTotal      = "container.resource.memory.total"
	MetricContainerResourceMemoryAvailable  = "container.resource.memory.available"
	MetricContainerResourceStorageAvailable = "container.resource.storage.total"
	MetricContainerResourceStorageTotal     = "container.resource.storage.available"
	MetricContainerResourceCPUUsage         = "container.resource.cpu.usage"
	// MetricContainerResourceCPUBenchmark average nanoseconds needed for cpu stress test. Less is better.
	MetricContainerResourceCPUBenchmark = "container.resource.cpu.benchmark"
)

// RunSystemResourceInspection checks bot has enough resources.
func RunSystemResourceInspection(ctx context.Context) (map[string]float64, error) {
	resources, err := GetResources(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve system resources: %w", err)
	}

	return resources, err
}

// GetResources returns resource metrics.
func GetResources(ctx context.Context) (map[string]float64, error) {
	var (
		result      = make(map[string]float64)
		resultError error
	)

	mi, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		resultError = multierror.Append(resultError, fmt.Errorf("can't read memory info: %w", err))
	} else {
		result[MetricContainerResourceMemoryTotal] = float64(mi.Total)
		result[MetricContainerResourceMemoryAvailable] = float64(mi.Available)
	}

	calcInterval := time.Second / 2

	ci, err := cpu.PercentWithContext(ctx, calcInterval, false)
	if err != nil {
		resultError = multierror.Append(resultError, fmt.Errorf("can't read cpu info: %w", err))
	} else {
		// cpu.Percent should return average result in a slice with a single item
		if len(ci) != 1 {
			result[MetricContainerResourceCPUUsage] = ci[0]
		}
	}

	di, err := disk.UsageWithContext(ctx, "/")
	if err != nil {
		resultError = multierror.Append(resultError, fmt.Errorf("can't read disk usage info: %w", err))
	} else {
		result[MetricContainerResourceStorageTotal] = float64(di.Total)
		result[MetricContainerResourceStorageAvailable] = float64(di.Free)
	}

	result[MetricContainerResourceCPUBenchmark] = float64(cpuStressTestBenchmark())

	return result, resultError
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
