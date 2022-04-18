package registry

import (
	"github.com/forta-network/forta-core-go/contracts/contract_scanner_registry"
)

const SubjectTypeScanner = 0
const SubjectTypeAgent = 1

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
}

type AssignmentHash struct {
	AgentLength int64  `json:"agentLength"`
	Hash        string `json:"hash"`
}

type StakingThreshold contract_scanner_registry.IStakeSubjectStakeThreshold
