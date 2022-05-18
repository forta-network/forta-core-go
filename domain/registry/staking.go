package registry

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/forta-network/forta-core-go/contracts/contract_agent_registry"
	"github.com/forta-network/forta-core-go/contracts/contract_scanner_registry"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/utils"
)

const AgentStake = "AgentStake"
const ScannerStake = "ScannerStake"
const AgentStakeThreshold = "AgentStakeThreshold"
const ScannerStakeThreshold = "ScannerStakeThreshold"

const ChangeTypeDeposit = "deposit"
const ChangeTypeWithdrawal = "withdrawal"
const ChangeTypeSlash = "slash"

type ThresholdMessage struct {
	Message
	MinHex    string  `json:"minHex"`
	MaxHex    string  `json:"maxHex"`
	Min       float64 `json:"min"`
	Max       float64 `json:"max"`
	Activated bool    `json:"activated"`
}

type AgentStakeThresholdMessage struct {
	ThresholdMessage
}

type ScannerStakeThresholdMessage struct {
	ThresholdMessage
	ChainID int64 `json:"chainId"`
}

type StakeMessage struct {
	Message
	ChangeType string  `json:"changeType"`
	AmountHex  string  `json:"amountHex"`
	Amount     float64 `json:"amount"`
}

type AgentStakeMessage struct {
	StakeMessage
	AgentID string `json:"agentId"`
}

type ScannerStakeMessage struct {
	StakeMessage
	ScannerID string `json:"scannerId"`
}

// https://github.com/ethereum/go-ethereum/issues/21221
func weiToEther(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.Ether))
}

func toEthFloat(v *big.Int) float64 {
	if v == nil {
		return 0
	}
	fl := weiToEther(v)
	res, _ := fl.Float64()
	return res
}

func hexValue(v *big.Int) string {
	if v == nil {
		return ""
	}
	return utils.BigIntToHex(v)
}

func NewScannerStakeMessage(l types.Log, changeType, scannerID string, value *big.Int, blk *domain.Block) *ScannerStakeMessage {
	return &ScannerStakeMessage{
		StakeMessage: StakeMessage{
			Message: Message{
				Action:    ScannerStake,
				Timestamp: time.Now().UTC(),
				Source:    SourceFromBlock(l.TxHash.Hex(), blk),
			},
			ChangeType: changeType,
			AmountHex:  hexValue(value),
			Amount:     toEthFloat(value),
		},
		ScannerID: scannerID,
	}
}

func NewAgentStakeMessage(l types.Log, changeType, agentID string, value *big.Int, blk *domain.Block) *AgentStakeMessage {
	return &AgentStakeMessage{
		StakeMessage: StakeMessage{
			Message: Message{
				Action:    AgentStake,
				Timestamp: time.Now().UTC(),
				Source:    SourceFromBlock(l.TxHash.Hex(), blk),
			},
			ChangeType: changeType,
			AmountHex:  hexValue(value),
			Amount:     toEthFloat(value),
		},
		AgentID: agentID,
	}
}

func NewAgentStakeThresholdMessage(evt *contract_agent_registry.AgentRegistryStakeThresholdChanged, l types.Log, blk *domain.Block) *AgentStakeThresholdMessage {
	return &AgentStakeThresholdMessage{
		ThresholdMessage: ThresholdMessage{
			Message: Message{
				Action:    AgentStakeThreshold,
				Timestamp: time.Now().UTC(),
				Source:    SourceFromBlock(l.TxHash.Hex(), blk),
			},
			MinHex:    hexValue(evt.Min),
			MaxHex:    hexValue(evt.Max),
			Min:       toEthFloat(evt.Min),
			Max:       toEthFloat(evt.Max),
			Activated: evt.Activated,
		},
	}
}

func NewScannerStakeThresholdMessage(evt *contract_scanner_registry.ScannerRegistryStakeThresholdChanged, l types.Log, blk *domain.Block) *ScannerStakeThresholdMessage {
	return &ScannerStakeThresholdMessage{
		ThresholdMessage: ThresholdMessage{
			Message: Message{
				Action:    ScannerStakeThreshold,
				Timestamp: time.Now().UTC(),
				Source:    SourceFromBlock(l.TxHash.Hex(), blk),
			},
			MinHex:    hexValue(evt.Min),
			MaxHex:    hexValue(evt.Max),
			Min:       toEthFloat(evt.Min),
			Max:       toEthFloat(evt.Max),
			Activated: evt.Activated,
		},
		ChainID: evt.ChainId.Int64(),
	}
}

func ParseAgentStakeThresholdMessage(msg string) (*AgentStakeThresholdMessage, error) {
	var m AgentStakeThresholdMessage
	err := json.Unmarshal([]byte(msg), &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func ParseScannerStakeThresholdMessage(msg string) (*ScannerStakeThresholdMessage, error) {
	var m ScannerStakeThresholdMessage
	err := json.Unmarshal([]byte(msg), &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func ParseAgentStakeMessage(msg string) (*AgentStakeMessage, error) {
	var m AgentStakeMessage
	err := json.Unmarshal([]byte(msg), &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func ParseScannerStakeMessage(msg string) (*ScannerStakeMessage, error) {
	var m ScannerStakeMessage
	err := json.Unmarshal([]byte(msg), &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
