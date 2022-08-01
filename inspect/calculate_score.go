package inspect

import (
	"fmt"
)

const (
	DefaultMinDownloadSpeedInMbps = 10  // 10 mbps
	DefaultMinUploadSpeedInMbps   = 3   // 3 mbps
	DefaultEarliestBlock          = 5   // block 5 in chain
	DefaultMinTotalMemory         = 3e9 // 3 gigabytes
	DefaultMinAvailableMemory     = 5e7 // 50 megabytes
)

// ScoreCalculator calculates score of an inspection result
type ScoreCalculator interface {
	CalculateSLAScore(results InspectionResults, chainID uint64) (float64, error)
	Name() string
}

// passFailCalculator returns 0 if any blocker inspections fail, else 1. See CalculateSLAScore for blocker inspections
type passFailCalculator struct {
	chainCalculators []chainPassFailCalculator
}

// chainPassFailCalculator returns 0 if any blocker inspections fail, else 1. See CalculateSLAScore for blocker inspections
type chainPassFailCalculator struct {
	config ChainPassFailCalculatorConfig
}

// ChainPassFailCalculatorConfig contains calculator related config per chain.
type ChainPassFailCalculatorConfig struct {
	// ChainID required to get correct calculator
	ChainID uint64

	// InspectionConfig currently used only for InspectionConfig.CheckTrace field
	InspectionConfig InspectionConfig

	// MinDownloadSpeedInMbps fallbacks to DefaultMinDownloadSpeedInMbps.
	MinDownloadSpeedInMbps float64
	// MinUploadSpeedInMbps fallbacks to DefaultMinUploadSpeedInMbps.
	MinUploadSpeedInMbps float64

	// ExpectedEarliestBlock fallbacks to DefaultEarliestBlock.
	ExpectedEarliestBlock float64
	// MinTotalMemory fallbacks to DefaultMinTotalMemory.
	MinTotalMemory float64
	// MinAvailableMemory fallbacks to DefaultMinAvailableMemory.
	MinAvailableMemory float64
}

// NewPassFailCalculator creates a new pass-fail calculator.
func NewPassFailCalculator(configs []ChainPassFailCalculatorConfig) *passFailCalculator {
	var calculators []chainPassFailCalculator
	for _, config := range configs {
		if config.MinDownloadSpeedInMbps == 0 {
			config.MinDownloadSpeedInMbps = DefaultMinDownloadSpeedInMbps
		}
		if config.MinUploadSpeedInMbps == 0 {
			config.MinUploadSpeedInMbps = DefaultMinUploadSpeedInMbps
		}
		if config.MinAvailableMemory == 0 {
			config.MinAvailableMemory = DefaultMinAvailableMemory
		}
		if config.MinTotalMemory == 0 {
			config.MinTotalMemory = DefaultMinTotalMemory
		}
		// this prevents us from using genesis block for expected block, but nbd
		if config.ExpectedEarliestBlock <= 0 {
			config.ExpectedEarliestBlock = DefaultEarliestBlock
		}

		calculators = append(calculators, chainPassFailCalculator{config: config})
	}

	return &passFailCalculator{chainCalculators: calculators}
}

func (s *passFailCalculator) Name() string {
	return "pass-fail"
}

func (s *passFailCalculator) CalculateSLAScore(results InspectionResults, chainID uint64) (float64, error) {
	calculator, err := s.getChainScoreCalculator(chainID)
	if err != nil {
		return 0, err
	}

	return calculator.CalculateSLAScore(results)
}

func (s *passFailCalculator) getChainScoreCalculator(chainID uint64) (chainPassFailCalculator, error) {
	for _, calculator := range s.chainCalculators {
		if calculator.config.ChainID == chainID {
			return calculator, nil
		}
	}
	return chainPassFailCalculator{}, fmt.Errorf("no calculator for %d", chainID)
}

// CalculateSLAScore returns 0 if any blocker inspections fail, else 1.
func (c *chainPassFailCalculator) CalculateSLAScore(results InspectionResults) (float64, error) {
	// any node must provide a decent network support.
	// Including:
	// Outbound Access
	if results.Indicators[IndicatorNetworkOutboundAccess] == ResultFailure {
		return 0, nil
	}
	// Download speed
	if results.Indicators[IndicatorNetworkDownloadSpeed] > 0 && results.Indicators[IndicatorNetworkDownloadSpeed] < c.config.MinDownloadSpeedInMbps {
		return 0, nil
	}
	// Upload speed
	if results.Indicators[IndicatorNetworkUploadSpeed] > 0 && results.Indicators[IndicatorNetworkUploadSpeed] < c.config.MinUploadSpeedInMbps {
		return 0, nil
	}

	// if required, trace should be supported
	if c.config.InspectionConfig.CheckTrace && (results.Indicators[IndicatorTraceAccessible] == ResultFailure || results.Indicators[IndicatorTraceSupported] == ResultFailure) {
		return 0, nil
	}

	// 	scan api should be provided along with required modules
	if results.Indicators[IndicatorScanAPIAccessible] == ResultFailure || results.Indicators[IndicatorScanAPIModuleEth] == ResultFailure || results.Indicators[IndicatorScanAPIModuleNet] == ResultFailure {
		return 0, nil
	}

	// scan api should point to correct chain id
	if results.Indicators[IndicatorScanAPIChainID] != float64(c.config.ChainID) {
		return 0, nil
	}

	// proxy api with required modules should be accessible
	if results.Indicators[IndicatorProxyAPIAccessible] == ResultFailure || results.Indicators[IndicatorProxyAPIModuleEth] == ResultFailure || results.Indicators[IndicatorProxyAPIModuleNet] == ResultFailure || results.Indicators[IndicatorProxyAPIModuleWeb3] == ResultFailure {
		return 0, nil
	}

	// proxy should support enough history of blocks
	if results.Indicators[IndicatorProxyAPIHistorySupport] >= c.config.ExpectedEarliestBlock {
		return 0, nil
	}

	// enough total as well as available memory should be provided to prevent OOM issues
	if results.Indicators[IndicatorResourcesMemoryTotal] < c.config.MinTotalMemory || results.Indicators[IndicatorResourcesMemoryAvailable] < c.config.MinAvailableMemory {
		return 0, nil
	}

	return 1, nil
}

// DefaultChainPassFailCalculatorConfig returns a ChainPassFailCalculatorConfig with the chain id and all default limits.
func DefaultChainPassFailCalculatorConfig(chainID uint64) ChainPassFailCalculatorConfig {
	return ChainPassFailCalculatorConfig{
		ChainID:                chainID,
		MinDownloadSpeedInMbps: DefaultMinDownloadSpeedInMbps,
		MinUploadSpeedInMbps:   DefaultMinUploadSpeedInMbps,
		ExpectedEarliestBlock:  DefaultEarliestBlock,
		MinTotalMemory:         DefaultMinTotalMemory,
		MinAvailableMemory:     DefaultMinAvailableMemory,
	}
}
