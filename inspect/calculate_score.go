package inspect

import (
	"errors"
	"fmt"
)

// Constants
const (
	DefaultMinDownloadSpeedInMbps = 10 // 10 mbps
	DefaultMinUploadSpeedInMbps   = 3  // 3 mbps
	DefaultEarliestBlock          = VeryOldBlockNumber
	DefaultMinTotalMemory         = 3e9 // 3 gigabytes
	DefaultMinAvailableMemory     = 5e7 // 50 megabytes
)

// Errors
var (
	ErrCalculatorNotFound = errors.New("calculator not found")
)

// ScoreCalculator calculates score of an inspection result.
type ScoreCalculator interface {
	CalculateScore(results *InspectionResults, chainID uint64) (float64, error)
}

type scoreCalculator struct {
	chainCalculators []*chainPassFailCalculator
}

type chainPassFailCalculator struct {
	config ScoreCalculatorConfig
}

// ScoreCalculatorConfig contains calculator related config per chain.
type ScoreCalculatorConfig struct {
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

// NewScoreCalculator creates a score calculator.
func NewScoreCalculator(configs []ScoreCalculatorConfig) *scoreCalculator {
	var calculators []*chainPassFailCalculator
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
		// this prevents us from using a very old block for expected block, but no big deal
		if config.ExpectedEarliestBlock <= 0 {
			config.ExpectedEarliestBlock = DefaultEarliestBlock
		}

		calculators = append(calculators, &chainPassFailCalculator{config: config})
	}

	return &scoreCalculator{chainCalculators: calculators}
}

func (s *scoreCalculator) CalculateScore(results *InspectionResults, chainID uint64) (float64, error) {
	calculator, err := s.getChainScoreCalculator(chainID)
	if err != nil {
		return 0, err
	}

	return calculator.CalculateScore(results)
}

func (s *scoreCalculator) getChainScoreCalculator(chainID uint64) (*chainPassFailCalculator, error) {
	for _, calculator := range s.chainCalculators {
		if calculator.config.ChainID == chainID {
			return calculator, nil
		}
	}
	return nil, fmt.Errorf("%w (chain %d)", ErrCalculatorNotFound, chainID)
}

// CalculateScore calculates an inspection score by checking provided results.
// For now, it returns 0 if some of the indicators report a negative result.
func (c *chainPassFailCalculator) CalculateScore(results *InspectionResults) (float64, error) {
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
	if c.config.InspectionConfig.CheckTrace &&
		(results.Indicators[IndicatorTraceAccessible] == ResultFailure ||
			results.Indicators[IndicatorTraceSupported] == ResultFailure) {
		return 0, nil
	}

	// scan api should be provided along with required modules
	if results.Indicators[IndicatorScanAPIAccessible] == ResultFailure ||
		results.Indicators[IndicatorScanAPIModuleEth] == ResultFailure ||
		results.Indicators[IndicatorScanAPIModuleNet] == ResultFailure {
		return 0, nil
	}

	// scan api should point to correct chain id
	if results.Indicators[IndicatorScanAPIChainID] != float64(c.config.ChainID) {
		return 0, nil
	}

	// proxy api with required modules should be accessible
	// if results.Indicators[IndicatorProxyAPIAccessible] == ResultFailure || results.Indicators[IndicatorProxyAPIModuleEth] == ResultFailure || results.Indicators[IndicatorProxyAPIModuleNet] == ResultFailure || results.Indicators[IndicatorProxyAPIModuleWeb3] == ResultFailure {
	// 	return 0, nil
	// }

	// proxy should support enough history of blocks
	// if results.Indicators[IndicatorProxyAPIHistorySupport] >= c.config.ExpectedEarliestBlock {
	// 	return 0, nil
	// }

	// enough total as well as available memory should be provided to prevent OOM issues
	if results.Indicators[IndicatorResourcesMemoryTotal] < c.config.MinTotalMemory || results.Indicators[IndicatorResourcesMemoryAvailable] < c.config.MinAvailableMemory {
		return 0, nil
	}

	return 1, nil
}

// DefaultScoreCalculatorConfig returns a ScoreCalculatorConfig with the chain id and all default limits.
func DefaultScoreCalculatorConfig(chainID uint64) ScoreCalculatorConfig {
	return ScoreCalculatorConfig{
		ChainID:                chainID,
		MinDownloadSpeedInMbps: DefaultMinDownloadSpeedInMbps,
		MinUploadSpeedInMbps:   DefaultMinUploadSpeedInMbps,
		ExpectedEarliestBlock:  DefaultEarliestBlock,
		MinTotalMemory:         DefaultMinTotalMemory,
		MinAvailableMemory:     DefaultMinAvailableMemory,
	}
}
