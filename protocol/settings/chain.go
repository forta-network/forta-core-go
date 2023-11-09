package settings

var defaultRateLimiting = &RateLimit{
	Rate:  50, // 0.347, // 30k/day
	Burst: 50, // 100,
}

// ChainSettings contains chain-specific settings.
type ChainSettings struct {
	ChainID                     int
	Name                        string
	EnableTrace                 bool
	JsonRpcRateLimiting         *RateLimit
	InspectionInterval          int // in block number
	BlockThreshold              int
	JSONRPCRetryIntervalSeconds int
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
		ChainID:                     1,
		Name:                        "Ethereum Mainnet",
		EnableTrace:                 true,
		JsonRpcRateLimiting:         defaultRateLimiting,
		InspectionInterval:          50,
		BlockThreshold:              4,
		JSONRPCRetryIntervalSeconds: 8,
	},
	{
		ChainID:                     10,
		Name:                        "Optimism",
		EnableTrace:                 false,
		JsonRpcRateLimiting:         defaultRateLimiting,
		InspectionInterval:          100,
		BlockThreshold:              10,
		JSONRPCRetryIntervalSeconds: 4,
	},
	{
		ChainID:                     56,
		Name:                        "BSC",
		EnableTrace:                 false,
		JsonRpcRateLimiting:         defaultRateLimiting,
		InspectionInterval:          250,
		BlockThreshold:              10,
		JSONRPCRetryIntervalSeconds: 4,
	},
	{
		ChainID:                     137,
		Name:                        "Polygon",
		EnableTrace:                 false,
		JsonRpcRateLimiting:         defaultRateLimiting,
		InspectionInterval:          70,
		BlockThreshold:              10,
		JSONRPCRetryIntervalSeconds: 4,
	},
	{
		ChainID:                     250,
		Name:                        "Fantom",
		EnableTrace:                 true,
		JsonRpcRateLimiting:         defaultRateLimiting,
		InspectionInterval:          750,
		BlockThreshold:              10,
		JSONRPCRetryIntervalSeconds: 4,
	},
	{
		ChainID:                     42161,
		Name:                        "Arbitrum",
		EnableTrace:                 false,
		JsonRpcRateLimiting:         defaultRateLimiting,
		InspectionInterval:          1500,
		BlockThreshold:              20,
		JSONRPCRetryIntervalSeconds: 4,
	},
	{
		ChainID:                     43114,
		Name:                        "Avalanche",
		EnableTrace:                 false,
		JsonRpcRateLimiting:         defaultRateLimiting,
		InspectionInterval:          350,
		BlockThreshold:              10,
		JSONRPCRetryIntervalSeconds: 4,
	},
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
		BlockThreshold:      10,
	}
}
