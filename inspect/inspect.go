package inspect

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-multierror"
)

// InspectionConfig contains the parameters for all inspections.
type InspectionConfig struct {
	BlockNumber uint64 `json:"blockNumber"`
	ScanAPIURL  string `json:"scanApiUrl"`
	ProxyAPIURL string `json:"proxyApiUrl"` // JSON-RPC API for bots
	TraceAPIURL string `json:"traceApiUrl"`
	CheckTrace  bool   `json:"checkTrace"`
}

// InspectionResults contains inspection results.
type InspectionResults struct {
	Inputs     InspectionConfig   `json:"inputs"`
	Metadata   map[string]string  `json:"metadata"`
	Indicators map[string]float64 `json:"indicators"`
}

// NewInspectionResults creates new InspectionResults.
func NewInspectionResults() *InspectionResults {
	return &InspectionResults{
		Metadata:   make(map[string]string),
		Indicators: make(map[string]float64),
	}
}

// CopyFrom copies other inspection results in.
func (ir *InspectionResults) CopyFrom(from *InspectionResults) *InspectionResults {
	for k, v := range from.Metadata {
		ir.Metadata[k] = v
	}
	for k, v := range from.Indicators {
		ir.Indicators[k] = v
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
	return InspectAll(
		ctx, []Inspector{
			&NetworkInspector{},
			&SystemResourcesInspector{},
			&ScanAPIInspector{},
			&ProxyAPIInspector{},
			&TraceAPIInspector{},
		}, inspectionCfg,
	)
}

// InspectAll runs all given inspections and aggregates results.
func InspectAll(ctx context.Context, inspectors []Inspector, inspectionCfg InspectionConfig) (allResults *InspectionResults, resultError error) {
	allResults = NewInspectionResults()
	allResults.Inputs = inspectionCfg

	for _, inspector := range inspectors {
		results, err := inspector.Inspect(ctx, inspectionCfg)
		if err != nil {
			resultError = multierror.Append(resultError, fmt.Errorf("'%s' inspection error: %w", inspector.Name(), err))
		}
		allResults.CopyFrom(results)
	}

	return
}
