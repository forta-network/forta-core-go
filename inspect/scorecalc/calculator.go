package scorecalc

import (
	"errors"
	"fmt"

	"github.com/forta-network/forta-core-go/inspect"
)

// Constants
const (
	DefaultMinDownloadSpeedInMbps = 10 // 10 mbps
	DefaultMinUploadSpeedInMbps   = 3  // 3 mbps
	DefaultEarliestBlock          = inspect.VeryOldBlockNumber
	DefaultMinTotalMemory         = 3e9 // 3 gigabytes
	DefaultMinAvailableMemory     = 5e7 // 50 megabytes
)

// Errors
var (
	ErrCalculatorNotFound  = errors.New("calculator not found")
	ErrBadInspectionResult = fmt.Errorf("bad inspection result")
)

// ChainScoreCalculator calculates score of an inspection result for given chain.
type ChainScoreCalculator interface {
	CalculateScore(chainID uint64, results *inspect.InspectionResults) (float64, error)
}

// compile time check for interface implementation
var _ ChainScoreCalculator = &scoreCalculator{}

type scoreCalculator struct {
	chainCalculators []*chainPassFailCalculator
}

// ScoreCalculatorConfig contains calculator related config per chain.
type ScoreCalculatorConfig struct {
	// ChainID required to get correct calculator
	ChainID uint64

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
		if config.ExpectedEarliestBlock < DefaultEarliestBlock {
			config.ExpectedEarliestBlock = DefaultEarliestBlock
		}

		calculators = append(calculators, &chainPassFailCalculator{config: config})
	}

	return &scoreCalculator{chainCalculators: calculators}
}

func (s *scoreCalculator) CalculateScore(chainID uint64, results *inspect.InspectionResults) (float64, error) {
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
