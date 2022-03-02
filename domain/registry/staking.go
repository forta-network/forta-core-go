package registry

import (
	"encoding/json"
	"time"
)

const AgentStake = "AgentStake"
const ScannerStake = "ScannerStake"
const AgentStakeThreshold = "AgentStakeThreshold"
const ScannerStakeThreshold = "ScannerStakeThreshold"

const ChangeTypeDeposit = "deposit"
const ChangeTypeWithdrawal = "withdrawal"
const ChangeTypeSlash = "slash"

type AgentStakeThresholdMessage struct {
	Message
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

func NewScannerStakeMessage(changeType, scannerID string) *ScannerStakeMessage {
	return &ScannerStakeMessage{
		StakeMessage: StakeMessage{
			Message: Message{
				Action:    ScannerStake,
				Timestamp: time.Now().UTC(),
			},
			ChangeType: changeType,
		},
		ScannerID: scannerID,
	}
}

func NewAgentStakeMessage(changeType, agentID string) *AgentStakeMessage {
	return &AgentStakeMessage{
		StakeMessage: StakeMessage{
			Message: Message{
				Action:    AgentStake,
				Timestamp: time.Now().UTC(),
			},
			ChangeType: changeType,
		},
		AgentID: agentID,
	}
}

func NewAgentStakeThresholdMessage() *AgentStakeThresholdMessage {
	return &AgentStakeThresholdMessage{
		Message: Message{
			Action:    AgentStakeThreshold,
			Timestamp: time.Now().UTC(),
		},
	}
}

func NewScannerStakeThresholdMessage(chainID int64) *ScannerStakeThresholdMessage {
	return &ScannerStakeThresholdMessage{
		Message: Message{
			Action:    ScannerStakeThreshold,
			Timestamp: time.Now().UTC(),
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
