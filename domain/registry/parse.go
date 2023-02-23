package registry

import (
	"encoding/json"
	"errors"

	"github.com/forta-network/forta-core-go/domain/registry/regmsg"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

// Errors
var (
	ErrMsgTypeNotFound = errors.New("message type not found for action")
)

// ParseMessage parses given message.
func ParseMessage(msg string) (regmsg.Interface, error) {
	// peek at action field quickly
	action := gjson.Get(msg, "action").String()
	if len(action) == 0 {
		log.WithField("body", msg).Error("action is not populated")
		return nil, errors.New("action is not populated")
	}

	msgTyp, ok := MessageTypeForAction(action)
	if !ok {
		log.WithField("action", action).Error(ErrMsgTypeNotFound.Error())
		return nil, ErrMsgTypeNotFound
	}

	err := json.Unmarshal([]byte(msg), &msgTyp)
	if err != nil {
		return nil, err
	}

	return msgTyp, nil
}

// MessageTypeForAction returns the correct message type for given action.
func MessageTypeForAction(action string) (regmsg.Interface, bool) {
	switch action {
	case SaveAgent:
		return &AgentSaveMessage{}, true

	case EnableAgent, DisableAgent:
		return &AgentMessage{}, true

	case Link, Unlink:
		return &DispatchMessage{}, true

	case ScannerNodeVersionUpdated:
		return &ScannerNodeVersionMessage{}, true

	case SaveScanner:
		return &ScannerSaveMessage{}, true

	case EnableScanner, DisableScanner:
		return &ScannerMessage{}, true

	case UpdateScannerPool:
		return &UpdateScannerPoolMessage{}, true

	case ScannerPoolAllocatedStake:
		return &ScannerPoolAllocationMessage{}, true

	case TransferShares:
		return &TransferSharesMessage{}, true

	case ScannerStake:
		return &ScannerStakeMessage{}, true

	case ScannerPoolStake:
		return &ScannerPoolStakeMessage{}, true

	case AgentStake:
		return &AgentStakeMessage{}, true

	case ScannerStakeThreshold:
		return &ScannerStakeThresholdMessage{}, true

	case AgentStakeThreshold:
		return &AgentStakeThresholdMessage{}, true

	case Upgrade, ConfigurationChange:
		return &UpgradeMessage{}, true
	}
	return nil, false
}
