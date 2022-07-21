package inspect

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-multierror"
)

// InspectionConfig contains the parameters for all inspections.
type InspectionConfig struct {
	ScanAPIURL  string `json:"scanApiUrl"`
	TraceAPIURL string `json:"traceApiUrl"`
	BlockNumber uint64 `json:"blockNumber"`
}

// InspectionResults contains inspection results.
type InspectionResults struct {
	Metadata   map[string]string
	Indicators map[string]float64
}

// NewInspectionResults creates new InspectionResults.
func NewInspectionResults() *InspectionResults {
	return &InspectionResults{
		Indicators: make(map[string]float64),
		Metadata:   make(map[string]string),
	}
}

// CopyFrom copies other inspection results in.
func (ir *InspectionResults) CopyFrom(from *InspectionResults) *InspectionResults {
	for k, v := range from.Indicators {
		ir.Indicators[k] = v
	}
	for k, v := range from.Metadata {
		ir.Metadata[k] = v
	}
	return ir
}

// Inspector inspects node capabilities.
type Inspector interface {
	Name() string
	Inspect(context.Context, InspectionConfig) (*InspectionResults, error)
}

// Inspect inspects node capabilities.
func Inspect(ctx context.Context, inspectionCfg InspectionConfig) (*InspectionResults, error) {
	return inspect(ctx, []Inspector{
		&NetworkInspector{},
		&SystemResourcesInspector{},
		&ScanAPIInspector{},
		&TraceAPIInspector{},
	}, inspectionCfg)
}

func inspect(ctx context.Context, inspectors []Inspector, inspectionCfg InspectionConfig) (allResults *InspectionResults, resultError error) {
	allResults = NewInspectionResults()

	for _, inspector := range inspectors {
		results, err := inspector.Inspect(ctx, inspectionCfg)
		if err != nil {
			resultError = multierror.Append(resultError, fmt.Errorf("'%s' inspection error: %w", inspector.Name(), err))
		}
		allResults.CopyFrom(results)
	}

	return
}
