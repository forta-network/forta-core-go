package registry

import (
	"github.com/forta-network/forta-core-go/contracts/merged/contract_scanner_registry"
)

const (
	SubjectTypeScanner              = 0
	SubjectTypeAgent                = 1
	SubjectTypeScannerPool          = 2
	SubjectTypeDelegatorScannerPool = 3
)

type Agent struct {
	AgentID  string  `json:"agentId"`
	ChainIDs []int64 `json:"chainIds"`
	Enabled  bool    `json:"enabled"`
	Manifest string  `json:"manifest"`
	Owner    string  `json:"owner"`
}

type Scanner struct {
	ScannerID string `json:"scannerId"`
	ChainID   int64  `json:"chainId"`
	Enabled   bool   `json:"enabled"`
	Manifest  string `json:"manifest"`
	Owner     string `json:"owner"`
	PoolID    string `json:"poolId"`
}

type AssignmentHash struct {
	AgentLength int64  `json:"agentLength"`
	Hash        string `json:"hash"`
}

type StakingThreshold contract_scanner_registry.GetStakeThresholdOutput

const UpgradedTopic = "0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b"
