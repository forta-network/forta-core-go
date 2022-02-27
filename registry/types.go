package registry

type Agent struct {
	AgentID  string  `json:"agentId"`
	ChainIDs []int64 `json:"chainIds"`
	Enabled  bool    `json:"enabled"`
	Manifest string  `json:"manifest"`
}

type Scanner struct {
	ScannerID string `json:"scannerId"`
	ChainID   int64  `json:"chainId"`
	Enabled   bool   `json:"enabled"`
	Manifest  string `json:"manifest"`
}

type AssignmentHash struct {
	AgentLength int64  `json:"agentLength"`
	Hash        string `json:"hash"`
}
