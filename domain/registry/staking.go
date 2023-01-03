package registry

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/forta-network/forta-core-go/contracts/contract_agent_registry"
	"github.com/forta-network/forta-core-go/contracts/contract_forta_staking"
	"github.com/forta-network/forta-core-go/contracts/contract_scanner_pool_registry"
	"github.com/forta-network/forta-core-go/contracts/contract_scanner_registry"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/utils"
)

const AgentStake = "AgentStake"
const ScannerStake = "ScannerStake"
const ScannerPoolStake = "ScannerPoolStake"
const AgentStakeThreshold = "AgentStakeThreshold"
const ScannerStakeThreshold = "ScannerStakeThreshold"
const TransferShares = "TransferShares"

const ChangeTypeDeposit = "deposit"
const ChangeTypeWithdrawal = "withdrawal"
const ChangeTypeSlash = "slash"

type ThresholdMessage struct {
	Message
	Min       string `json:"min"`
	Max       string `json:"max"`
	Activated bool   `json:"activated"`
}

type AgentStakeThresholdMessage struct {
	ThresholdMessage
}

type TransferSharesMessage struct {
	Message
	ShareID   string `json:"shareId"`
	StakeType string `json:"stakeType"`
	Active    bool   `json:"active"`
	To        string `json:"to"`
	From      string `json:"from"`
	Amount    string `json:"amount"`
}

func (t TransferSharesMessage) IsBurn() bool {
	return t.To == utils.ZeroAddress
}

func (t TransferSharesMessage) IsMint() bool {
	return t.From == utils.ZeroAddress
}

type ScannerStakeThresholdMessage struct {
	ThresholdMessage
	ChainID int64 `json:"chainId"`
}

type StakeMessage struct {
	Message
	ChangeType string `json:"changeType"`
	Amount     string `json:"amount"`
	Account    string `json:"account"`
}

type AgentStakeMessage struct {
	StakeMessage
	AgentID string `json:"agentId"`
}

type ScannerStakeMessage struct {
	StakeMessage
	ScannerID string `json:"scannerId"`
}

type ScannerPoolStakeMessage struct {
	StakeMessage
	PoolID string `json:"poolId"`
}

func valueString(v *big.Int) string {
	if v == nil {
		return ""
	}
	return v.String()
}

func extractStakeType(id *big.Int) (string, error) {
	b := id.Bytes()
	lastByte := b[len(b)-1]
	switch lastByte {
	case 0:
		return ScannerStake, nil
	case 1:
		return AgentStake, nil
	case 2:
		return ScannerPoolStake, nil
	default:
		return "", fmt.Errorf("invalid stake type '%d': %s", lastByte, id.String())
	}
}

func isActive(id *big.Int) bool {
	one := big.NewInt(1)
	return 0 == big.NewInt(0).And(id, one.Lsh(one, 8)).Cmp(big.NewInt(256))
}

func TransferSharesMessageFromSingle(l types.Log, evt *contract_forta_staking.FortaStakingTransferSingle, blk *domain.Block) (*TransferSharesMessage, error) {
	st, err := extractStakeType(evt.Id)
	if err != nil {
		return nil, err
	}
	return &TransferSharesMessage{
		Message: Message{
			Action:    TransferShares,
			Timestamp: time.Now().UTC(),
			Source:    SourceFromBlock(l.TxHash.Hex(), blk),
		},
		StakeType: st,
		ShareID:   evt.Id.String(),
		Active:    isActive(evt.Id),
		To:        strings.ToLower(evt.To.Hex()),
		From:      strings.ToLower(evt.From.Hex()),
		Amount:    evt.Value.String(),
	}, nil
}

func TransferSharesMessagesFromBatch(l types.Log, evt *contract_forta_staking.FortaStakingTransferBatch, blk *domain.Block) ([]*TransferSharesMessage, error) {
	var res []*TransferSharesMessage
	for i, id := range evt.Ids {
		st, err := extractStakeType(id)
		if err != nil {
			return nil, err
		}
		res = append(res, &TransferSharesMessage{
			Message: Message{
				Action:    TransferShares,
				Timestamp: time.Now().UTC(),
				Source:    SourceFromBlock(l.TxHash.Hex(), blk),
			},
			StakeType: st,
			ShareID:   id.String(),
			Active:    isActive(id),
			To:        strings.ToLower(evt.To.Hex()),
			From:      strings.ToLower(evt.From.Hex()),
			Amount:    evt.Values[i].String(),
		})
	}
	return res, nil
}

func NewScannerStakeMessage(l types.Log, changeType string, account common.Address, scannerID string, value *big.Int, blk *domain.Block) *ScannerStakeMessage {
	return &ScannerStakeMessage{
		StakeMessage: StakeMessage{
			Message:    MessageFrom(l.TxHash.Hex(), blk, ScannerStake),
			ChangeType: changeType,
			Amount:     valueString(value),
			Account:    strings.ToLower(account.Hex()),
		},
		ScannerID: strings.ToLower(scannerID),
	}
}

func NewAgentStakeMessage(l types.Log, changeType string, account common.Address, agentID string, value *big.Int, blk *domain.Block) *AgentStakeMessage {
	return &AgentStakeMessage{
		StakeMessage: StakeMessage{
			Message:    MessageFrom(l.TxHash.Hex(), blk, AgentStake),
			ChangeType: changeType,
			Amount:     valueString(value),
			Account:    strings.ToLower(account.Hex()),
		},
		AgentID: agentID,
	}
}

func NewScannerPoolStakeMessage(l types.Log, changeType string, account common.Address, poolID string, value *big.Int, blk *domain.Block) *ScannerPoolStakeMessage {
	return &ScannerPoolStakeMessage{
		StakeMessage: StakeMessage{
			Message:    MessageFrom(l.TxHash.Hex(), blk, ScannerPoolStake),
			ChangeType: changeType,
			Amount:     valueString(value),
			Account:    strings.ToLower(account.Hex()),
		},
		PoolID: strings.ToLower(poolID),
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
			Min:       valueString(evt.Min),
			Max:       valueString(evt.Max),
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
			Min:       valueString(evt.Min),
			Max:       valueString(evt.Max),
			Activated: evt.Activated,
		},
		ChainID: evt.ChainId.Int64(),
	}
}

func NewScannerManagedStakeThresholdMessage(evt *contract_scanner_pool_registry.ScannerPoolRegistryManagedStakeThresholdChanged, l types.Log, blk *domain.Block) *ScannerStakeThresholdMessage {
	return &ScannerStakeThresholdMessage{
		ThresholdMessage: ThresholdMessage{
			Message: Message{
				Action:    ScannerStakeThreshold,
				Timestamp: time.Now().UTC(),
				Source:    SourceFromBlock(l.TxHash.Hex(), blk),
			},
			Min:       valueString(evt.Min),
			Max:       valueString(evt.Max),
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

func ParseTransferSharesMessage(msg string) (*TransferSharesMessage, error) {
	var m TransferSharesMessage
	err := json.Unmarshal([]byte(msg), &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
