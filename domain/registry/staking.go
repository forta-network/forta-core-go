package registry

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/forta-network/forta-core-go/contracts/contract_forta_staking"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/utils"
)

const AgentStake = "AgentStake"
const ScannerStake = "ScannerStake"
const AgentStakeThreshold = "AgentStakeThreshold"
const ScannerStakeThreshold = "ScannerStakeThreshold"
const StakeTransfer = "StakeTransfer"

const ChangeTypeDeposit = "deposit"
const ChangeTypeWithdrawal = "withdrawal"
const ChangeTypeSlash = "slash"

type AgentStakeThresholdMessage struct {
	Message
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
	Message
	ChainID int64 `json:"chainId"`
}

type StakeMessage struct {
	Message
	ChangeType string `json:"changeType"`
}

type AgentStakeMessage struct {
	StakeMessage
	AgentID string `json:"agentId"`
}

type ScannerStakeMessage struct {
	StakeMessage
	ScannerID string `json:"scannerId"`
}

func isActive(id *big.Int) bool {
	one := big.NewInt(1)
	return 0 == big.NewInt(0).And(id, one.Lsh(one, 8)).Cmp(big.NewInt(256))
}

func extractStakeType(id *big.Int) (string, error) {
	b := id.Bytes()
	lastByte := b[len(b)-1]
	switch lastByte {
	case 0:
		return ScannerStake, nil
	case 1:
		return AgentStake, nil
	default:
		return "", fmt.Errorf("invalid stake type: %s", id.String())
	}
}

func TransferSharesMessageFromSingle(l types.Log, evt *contract_forta_staking.FortaStakingTransferSingle, blk *domain.Block) (*TransferSharesMessage, error) {
	st, err := extractStakeType(evt.Id)
	if err != nil {
		return nil, err
	}
	return &TransferSharesMessage{
		Message: Message{
			Action:    StakeTransfer,
			Timestamp: time.Now().UTC(),
			Source:    SourceFromBlock(l.TxHash.Hex(), blk),
		},
		StakeType: st,
		ShareID:   evt.Id.String(),
		Active:    isActive(evt.Id),
		To:        evt.To.Hex(),
		From:      evt.From.Hex(),
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
				Action:    StakeTransfer,
				Timestamp: time.Now().UTC(),
				Source:    SourceFromBlock(l.TxHash.Hex(), blk),
			},
			StakeType: st,
			ShareID:   id.String(),
			Active:    isActive(id),
			To:        evt.To.Hex(),
			From:      evt.From.Hex(),
			Amount:    evt.Values[i].String(),
		})
	}
	return res, nil
}

func NewScannerStakeMessage(l types.Log, changeType, scannerID string, blk *domain.Block) *ScannerStakeMessage {
	return &ScannerStakeMessage{
		StakeMessage: StakeMessage{
			Message: Message{
				Action:    ScannerStake,
				Timestamp: time.Now().UTC(),
				Source:    SourceFromBlock(l.TxHash.Hex(), blk),
			},
			ChangeType: changeType,
		},
		ScannerID: scannerID,
	}
}

func NewAgentStakeMessage(l types.Log, changeType, agentID string, blk *domain.Block) *AgentStakeMessage {
	return &AgentStakeMessage{
		StakeMessage: StakeMessage{
			Message: Message{
				Action:    AgentStake,
				Timestamp: time.Now().UTC(),
				Source:    SourceFromBlock(l.TxHash.Hex(), blk),
			},
			ChangeType: changeType,
		},
		AgentID: agentID,
	}
}

func NewAgentStakeThresholdMessage(l types.Log, blk *domain.Block) *AgentStakeThresholdMessage {
	return &AgentStakeThresholdMessage{
		Message: Message{
			Action:    AgentStakeThreshold,
			Timestamp: time.Now().UTC(),
			Source:    SourceFromBlock(l.TxHash.Hex(), blk),
		},
	}
}

func NewScannerStakeThresholdMessage(l types.Log, chainID int64, blk *domain.Block) *ScannerStakeThresholdMessage {
	return &ScannerStakeThresholdMessage{
		Message: Message{
			Action:    ScannerStakeThreshold,
			Timestamp: time.Now().UTC(),
			Source:    SourceFromBlock(l.TxHash.Hex(), blk),
		},
		ChainID: chainID,
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
