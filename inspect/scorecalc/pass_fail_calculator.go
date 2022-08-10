package scorecalc

import (
	"fmt"

	"github.com/forta-network/forta-core-go/inspect"
)

// constants
const (
	MinTotalMemoryRequired = 8e9 // 8 gigabytes
)

type chainPassFailCalculator struct {
	config ScoreCalculatorConfig
}

// CalculateScore calculates an inspection score by checking provided results. It returns either 0 or 1
// For now, it returns 0 if some indicators report a negative result.
func (c *chainPassFailCalculator) CalculateScore(results *inspect.InspectionResults) (float64, error) {
	if results == nil {
		return 0, fmt.Errorf("nil inspection result %v", ErrBadInspectionResult)
	}

	if results.Indicators == nil {
		return 0, fmt.Errorf("inspection result has no indicators %v", ErrBadInspectionResult)
	}

	// node must provide outbound network access
	if results.Indicators[inspect.IndicatorNetworkOutboundAccess] == inspect.ResultFailure {
		return 0, nil
	}

	// if required, trace should be supported
	if results.Inputs.CheckTrace &&
		results.Indicators[inspect.IndicatorTraceSupported] == inspect.ResultFailure {
		return 0, nil
	}

	// scan api should be provided along with required modules
	if results.Indicators[inspect.IndicatorScanAPIAccessible] == inspect.ResultFailure ||
		results.Indicators[inspect.IndicatorScanAPIModuleEth] == inspect.ResultFailure {
		return 0, nil
	}

	// scan api should point to correct chain id
	if results.Indicators[inspect.IndicatorScanAPIChainID] != float64(c.config.ChainID) {
		return 0, nil
	}

	// at least 50% of the required memory limit is required
	if results.Indicators[inspect.IndicatorResourcesMemoryTotal] < MinTotalMemoryRequired {
		return 0, nil
	}

	return 1, nil
}
