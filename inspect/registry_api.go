package inspect

import (
	"context"
	"fmt"

	"github.com/forta-network/forta-core-go/registry"
	"github.com/hashicorp/go-multierror"
)

const (
	// IndicatorRegistryAPIAccessible tells if the API is accessible.
	IndicatorRegistryAPIAccessible = "registry-api.accessible"
	// IndicatorRegistryAPIENS tells if we are able to resolve the registry contracts using ENS.
	IndicatorRegistryAPIENS = "registry-api.ens"
	// IndicatorRegistryAPIAssignments tells if bot assignments are detectable.
	IndicatorRegistryAPIAssignments = "registry-api.assignments"
)

var (
	registryAPIIndicators = []string{
		IndicatorRegistryAPIAccessible, IndicatorRegistryAPIAssignments,
	}
)

// RegistryAPIInspector is an inspector implementation.
type RegistryAPIInspector struct{}

// compile time check: it should implement the interface
var _ Inspector = &RegistryAPIInspector{}

// Name returns the name of the inspector.
func (sai *RegistryAPIInspector) Name() string {
	return "registry-api"
}

// Inspect checks given JSON-RPC node url supports eth and net modules.
func (sai *RegistryAPIInspector) Inspect(ctx context.Context, inspectionCfg InspectionConfig) (results *InspectionResults, resultErr error) {
	results = NewInspectionResults()
	results.Indicators = defaultIndicators(registryAPIIndicators)

	regEthClient, err := EthClientDialContext(ctx, inspectionCfg.RegistryAPIURL)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("can't dial json-rpc api: %w", err))

		results.Indicators[IndicatorRegistryAPIAccessible] = ResultFailure
		results.Indicators[IndicatorRegistryAPIENS] = ResultFailure
		results.Indicators[IndicatorRegistryAPIAssignments] = ResultFailure
		return
	} else {
		defer regEthClient.Close()
		results.Indicators[IndicatorRegistryAPIAccessible] = ResultSuccess
	}

	regClient, err := RegistryNewClient(ctx, registry.ClientConfig{
		JsonRpcUrl: inspectionCfg.RegistryAPIURL,
		ENSAddress: inspectionCfg.ENSContractAddress,
		Name:       "inspection-registry-client",
	})
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to initialize the registry client: %w", err))

		results.Indicators[IndicatorRegistryAPIENS] = ResultFailure
		results.Indicators[IndicatorRegistryAPIAssignments] = ResultFailure
		return
	} else {
		defer regClient.Close()
		results.Indicators[IndicatorRegistryAPIENS] = ResultSuccess
	}

	_, err = regClient.GetAssignmentHash(inspectionCfg.ScannerAddress)
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("failed to get the assignment hash: %w", err))

		results.Indicators[IndicatorRegistryAPIAssignments] = ResultFailure
		return
	} else {
		results.Indicators[IndicatorRegistryAPIAssignments] = ResultSuccess
	}

	return
}
