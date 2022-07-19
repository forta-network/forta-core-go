package registry

import (
	"fmt"
	"time"

	"github.com/forta-network/forta-core-go/domain"

	"github.com/goccy/go-json"

	"github.com/forta-network/forta-core-go/contracts/contract_agent_registry"
	"github.com/forta-network/forta-core-go/utils"
)

const SaveAgent = "SaveAgent"
const DisableAgent = "DisableAgent"
const EnableAgent = "EnableAgent"

const AgentPermissionAdmin = 0
const AgentPermissionOwner = 1

type AgentMessage struct {
	Message
	AgentID    string `json:"agentId"`
	TxHash     string `json:"txHash"`
	Permission int    `json:"permission"`
}

type AgentSaveMessage struct {
	AgentMessage
	Enabled  bool    `json:"enabled"`
	Name     string  `json:"name"`
	ChainIDs []int64 `json:"chainIds"`
	Metadata string  `json:"metadata"`
	Owner    string  `json:"owner"`
}

func ParseAgentSave(msg string) (*AgentSaveMessage, error) {
	var save AgentSaveMessage
	err := json.Unmarshal([]byte(msg), &save)
	if err != nil {
		return nil, err
	}
	if save.Action != SaveAgent {
		return nil, fmt.Errorf("invalid action for AgentSave: %s", save.Action)
	}
	return &save, nil
}

func ParseAgentMessage(msg string) (*AgentMessage, error) {
	var m AgentMessage
	err := json.Unmarshal([]byte(msg), &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func NewAgentMessage(evt *contract_agent_registry.AgentRegistryAgentEnabled, blk *domain.Block) *AgentMessage {
	agentID := utils.Hex(evt.AgentId)
	evtName := DisableAgent
	if evt.Enabled {
		evtName = EnableAgent
	}
	return &AgentMessage{
		Message: Message{
			Action:    evtName,
			Timestamp: time.Now().UTC(),
			Source:    SourceFromBlock(evt.Raw.TxHash.Hex(), blk),
		},
		AgentID:    agentID,
		Permission: int(evt.Permission),
		TxHash:     evt.Raw.TxHash.Hex(),
	}
}

func NewAgentSaveMessage(evt *contract_agent_registry.AgentRegistryAgentUpdated, blk *domain.Block) *AgentSaveMessage {
	agentID := utils.Hex(evt.AgentId)
	return &AgentSaveMessage{
		AgentMessage: AgentMessage{
			AgentID: agentID,
			Message: Message{
				Action:    SaveAgent,
				Timestamp: time.Now().UTC(),
				Source:    SourceFromBlock(evt.Raw.TxHash.Hex(), blk),
			},
			TxHash: evt.Raw.TxHash.Hex(),
		},
		Enabled:  true,
		Name:     evt.Metadata,
		ChainIDs: utils.IntArray(evt.ChainIds),
		Metadata: evt.Metadata,
		Owner:    evt.By.Hex(),
	}
}
