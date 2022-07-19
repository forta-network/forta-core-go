package inspect

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-multierror"
)

const InspectionFindingName = "Scan Node Inspection"

func RunAllInspections(ctx context.Context, nodeURL string) (map[string]float64, error) {
	var (
		result      = make(map[string]float64)
		resultError error
	)

	nwResults, err := RunNetworkInspection(ctx)
	if err != nil {
		resultError = multierror.Append(resultError, fmt.Errorf("network inspection error: %w", err))
	}
	copyMetrics(nwResults, result)

	rscResults, err := RunSystemResourceInspection(ctx)
	if err != nil {
		resultError = multierror.Append(resultError, fmt.Errorf("resource inspection error: %w", err))
	}
	copyMetrics(rscResults, result)

	scanResults, err := RunScanNodeInspection(ctx, nodeURL)
	if err != nil {
		resultError = multierror.Append(resultError, fmt.Errorf("scan api inspection error: %w", err))
	}
	copyMetrics(scanResults, result)

	return result, resultError
}

func copyMetrics(src, dst map[string]float64) {
	for s, f := range src {
		dst[s] = f
	}
}
