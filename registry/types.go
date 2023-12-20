package registry

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/forta-network/forta-core-go/contracts/merged/contract_scanner_registry"
	"github.com/forta-network/forta-core-go/security"
	"github.com/forta-network/forta-core-go/security/eip712"
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

type Assignment struct {
	AgentID       string
	AgentManifest string
	AgentOwner    string
	ScannerIndices
}

type ScannerIndices struct {
	SameChainAssignedScanners int
	SameChainScannerIndex     int
	AllChainsAssignedScanners int
	AllChainsScannerIndex     int
}

type StakingThreshold contract_scanner_registry.GetStakeThresholdOutput

const UpgradedTopic = "0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b"

type ScannerRegistrationInput struct {
	Scanner       common.Address `json:"scanner"`
	ScannerPoolId *big.Int       `json:"scannerPoolId"`
	ChainId       *big.Int       `json:"chainId"`
	Metadata      string         `json:"metadata"`
	Timestamp     *big.Int       `json:"timestamp"`
}

type ScannerRegistrationInfo struct {
	RegistrationInput *ScannerRegistrationInput `json:"registrationInput"`
	Signature         string                    `json:"signature"`
}

func MakeScannerRegistrationInfo(reg *eip712.ScannerNodeRegistration, sig []byte) (*ScannerRegistrationInfo, error) {
	var scannerRegInfo ScannerRegistrationInfo
	scannerRegInfo.RegistrationInput = (*ScannerRegistrationInput)(reg)
	var err error
	scannerRegInfo.Signature, err = security.EncodeEthereumSignature(sig)
	if err != nil {
		return nil, err
	}
	return &scannerRegInfo, nil
}

type numOutput struct {
	Num *big.Int
}

type scannerRefAtOutput struct {
	Registered bool
	ScannerId  *big.Int
	Owner      common.Address
	ChainId    *big.Int
}

type agentRefAtOutput struct {
	Registered   bool
	Owner        common.Address
	AgentId      *big.Int
	AgentVersion *big.Int
	Metadata     string
}
