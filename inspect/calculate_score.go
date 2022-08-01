package inspect

import (
	"fmt"
)

const (
	DefaultMinDownloadSpeedInMbps = 10
	DefaultMinUploadSpeedInMbps   = 3
	DefaultEarliestBlock          = 5
	DefaultMinTotalMemory         = 3e9 // 3 gigabytes
	DefaultMinAvailableMemory     = 5e7 // 50 megabytes
)

type ScoreCalculator interface {
	CalculateSLAScore(results InspectionResults) (float64, error)
	Name() string
}

type passFailCalculator struct {
	chainCalculators []chainPassFailCalculator
}

type chainPassFailCalculator struct {
	config ChainPassFailCalculatorConfig
}

type ChainPassFailCalculatorConfig struct {
	chainID uint64

	InspectionConfig InspectionConfig

	MinDownloadSpeedInMbps float64
	MinUploadSpeedInMbps   float64

	ExpectedEarliestBlock float64

	MinTotalMemory     float64
	MinAvailableMemory float64
}

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
		if calculator.config.chainID == chainID {
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
	if results.Indicators[IndicatorScanAPIChainID] != float64(c.config.chainID) {
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
