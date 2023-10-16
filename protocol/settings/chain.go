package settings

var defaultRateLimiting = &RateLimit{
	Rate:  50, // 0.347, // 30k/day
	Burst: 50, // 100,
}

// ChainSettings contains chain-specific settings.
type ChainSettings struct {
	ChainID             int
	Name                string
	EnableTrace         bool
	JsonRpcRateLimiting *RateLimit
	InspectionInterval  int // in block number

	DefaultOffset  int
	SafeOffset     int
	BlockThreshold int
}

// RateLimit is token bucket algorithm parameters.
type RateLimit struct {
	Rate  float64
	Burst int
}

// sorted by chain ID
// inspection intervals are determined to cover around 10-15m per chain
// reward thresholds are determined from network stats (at least 10)
// safe offsets are adjusted to be 5-10% of the block threshold (at least 1)
var allChainSettings = []ChainSettings{
	{
		ChainID:             1,
		Name:                "Ethereum Mainnet",
		EnableTrace:         true,
		JsonRpcRateLimiting: defaultRateLimiting,
		InspectionInterval:  50,

		DefaultOffset:  0,
		SafeOffset:     1,
		BlockThreshold: 20,
	},
	{
		ChainID:             10,
		Name:                "Optimism",
		EnableTrace:         false,
		JsonRpcRateLimiting: defaultRateLimiting,
		InspectionInterval:  5000,

		DefaultOffset:  0,
		SafeOffset:     500,
		BlockThreshold: 10000,
	},
	{
		ChainID:             56,
		Name:                "BSC",
		EnableTrace:         false,
		JsonRpcRateLimiting: defaultRateLimiting,
		InspectionInterval:  250,

		DefaultOffset:  0,
		SafeOffset:     3,
		BlockThreshold: 50,
	},
	{
		ChainID:             137,
		Name:                "Polygon",
		EnableTrace:         false,
		JsonRpcRateLimiting: defaultRateLimiting,
		InspectionInterval:  70,

		DefaultOffset:  0,
		SafeOffset:     4,
		BlockThreshold: 70,
	},
	{
		ChainID:             250,
		Name:                "Fantom",
		EnableTrace:         true,
		JsonRpcRateLimiting: defaultRateLimiting,
		InspectionInterval:  750,

		DefaultOffset:  0,
		SafeOffset:     5,
		BlockThreshold: 100,
	},
	{
		ChainID:             42161,
		Name:                "Arbitrum",
		EnableTrace:         false,
		JsonRpcRateLimiting: defaultRateLimiting,
		InspectionInterval:  1500,

		DefaultOffset:  0,
		SafeOffset:     60,
		BlockThreshold: 1200,
	},
	{
		ChainID:             43114,
		Name:                "Avalanche",
		EnableTrace:         false,
		JsonRpcRateLimiting: defaultRateLimiting,
		InspectionInterval:  350,

		DefaultOffset:  0,
		SafeOffset:     4,
		BlockThreshold: 70,
	},
}

// ValidateChainSettings validates chain settings.
func ValidateChainSettings(chainID int) bool {
	settings := GetChainSettings(chainID)
	if settings.SafeOffset < 1 {
		return false
	}
	if settings.BlockThreshold <= 0 {
		return false
	}
	safeOffsetRate := float64(settings.SafeOffset) / float64(settings.BlockThreshold)
	if safeOffsetRate < 0.05 || safeOffsetRate > 0.1 {
		return false
	}
	return true
}

// GetChainSettings returns the settings for the chain.
func GetChainSettings(chainID int) *ChainSettings {
	for _, settings := range allChainSettings {
		if settings.ChainID == chainID {
			return &settings
		}
	}
	return &ChainSettings{
		Name:                "Unknown chain",
		ChainID:             chainID,
		JsonRpcRateLimiting: defaultRateLimiting,
		InspectionInterval:  50, // arbitrary value - not reliable

		DefaultOffset:  0,
		SafeOffset:     2,
		BlockThreshold: 10,
	}
}

// GetDefaultBlockOffset returns the block offset for a chain.
func GetBlockOffset(chainID int) int {
	return GetChainSettings(chainID).DefaultOffset
}
