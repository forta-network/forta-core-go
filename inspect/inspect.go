package inspect

import (
	"context"
	"fmt"

	"github.com/forta-network/forta-core-go/protocol"
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
	Inputs     InspectionConfig
	Metadata   map[string]string
	Indicators map[string]float64
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

// ToProto returns the protobuf payload for the results.
func (ir *InspectionResults) ToProto() *protocol.InspectionResults {
	return &protocol.InspectionResults{
		Inputs: &protocol.InspectionInputs{
			BlockNumber:  ir.Inputs.BlockNumber,
			ScanApiHost:  getHost(ir.Inputs.ScanAPIURL),
			ProxyApiHost: getHost(ir.Inputs.ProxyAPIURL),
			TraceApiHost: getHost(ir.Inputs.TraceAPIURL),
			CheckTrace:   ir.Inputs.CheckTrace,
		},
		Metadata:   ir.Metadata,
		Indicators: ir.Indicators,
	}
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
		&ProxyAPIInspector{},
		&TraceAPIInspector{},
	}, inspectionCfg)
}

func inspect(ctx context.Context, inspectors []Inspector, inspectionCfg InspectionConfig) (allResults *InspectionResults, resultError error) {
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
