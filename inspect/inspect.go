package inspect

import (
	"context"
	"fmt"
	"net/url"

	"github.com/forta-network/forta-core-go/protocol"
	"github.com/hashicorp/go-multierror"
)

const (
	IndicatorValidAPIReferences = "api.refs.valid"
)

// InspectionConfig contains the parameters for all inspections.
type InspectionConfig struct {
	BlockNumber        uint64 `json:"blockNumber"`
	ScanAPIURL         string `json:"scanApiUrl"`
	ProxyAPIURL        string `json:"proxyApiUrl"` // JSON-RPC API to let bots make calls
	TraceAPIURL        string `json:"traceApiUrl"`
	CheckTrace         bool   `json:"checkTrace"`
	IsETH2             bool   `json:"isETH2"`
	RegistryAPIURL     string `json:"registryApiUrl"` // JSON-RPC API for loading the bots
	ENSContractAddress string `json:"ensContractAddress"`
	ScannerAddress     string `json:"scannerAddress"`
}

// InspectionResults contains inspection results.
type InspectionResults struct {
	Inputs     InspectionConfig   `json:"inputs"`
	Metadata   map[string]string  `json:"metadata"`
	Indicators map[string]float64 `json:"indicators"`
}

// ToProtoInspectionResults transforms inspection results to protobuf inspection results model.
func ToProtoInspectionResults(results *InspectionResults) *protocol.InspectionResults {
	return &protocol.InspectionResults{
		Inputs: &protocol.InspectionInputs{
			BlockNumber:        results.Inputs.BlockNumber,
			ScanApiHost:        getHost(results.Inputs.ScanAPIURL),
			ProxyApiHost:       getHost(results.Inputs.ProxyAPIURL),
			TraceApiHost:       getHost(results.Inputs.TraceAPIURL),
			CheckTrace:         results.Inputs.CheckTrace,
			RegistryApiHost:    getHost(results.Inputs.RegistryAPIURL),
			EnsContractAddress: results.Inputs.ENSContractAddress,
			ScannerAddress:     results.Inputs.ScannerAddress,
		},
		Metadata:   results.Metadata,
		Indicators: results.Indicators,
	}
}

func getHost(apiURL string) string {
	if len(apiURL) == 0 {
		return "null"
	}
	u, err := url.Parse(apiURL)
	if err != nil {
		return "invalid"
	}
	return u.Host
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
			&RegistryAPIInspector{},
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

	// keep the hash references the same as in result metadata: allows cross-checking only
	_, err := ValidateHashReferences(allResults.Metadata, allResults.Inputs, &HashReferences{
		ScanAPIBlockHash:  allResults.Metadata[MetadataScanAPIBlockByNumberHash],
		ProxyAPIBlockHash: allResults.Metadata[MetadataProxyAPIBlockByNumberHash],
		TraceAPIBlockHash: allResults.Metadata[MetadataTraceAPIBlockByNumberHash],
		TraceAPITraceHash: allResults.Metadata[MetadataTraceAPITraceBlockHash],
	})
	if err != nil {
		allResults.Indicators[IndicatorValidAPIReferences] = ResultFailure
	} else {
		allResults.Indicators[IndicatorValidAPIReferences] = ResultSuccess
	}

	return
}
