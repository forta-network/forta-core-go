package registry

import (
	"time"

	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/domain/registry/regmsg"
	"github.com/sirupsen/logrus"

	"github.com/forta-network/forta-core-go/contracts/merged/contract_agent_registry"
	"github.com/forta-network/forta-core-go/utils"
)

const SaveAgent = "SaveAgent"
const DisableAgent = "DisableAgent"
const EnableAgent = "EnableAgent"

const AgentPermissionAdmin = 0
const AgentPermissionOwner = 1

type AgentMessage struct {
	regmsg.Message
	AgentID    string `json:"agentId"`
	TxHash     string `json:"txHash"`
	Permission int    `json:"permission"`
}

func (am *AgentMessage) LogFields() logrus.Fields {
	return logrus.Fields{"agentId": am.AgentID}
}

type AgentSaveMessage struct {
	AgentMessage
	AgentProperties
}

type AgentProperties struct {
	Enabled  bool    `json:"enabled"`
	Name     string  `json:"name"`
	ChainIDs []int64 `json:"chainIds"`
	Metadata string  `json:"metadata"`
	Owner    string  `json:"owner"`
}

func NewAgentMessage(evt *contract_agent_registry.AgentRegistryAgentEnabled, blk *domain.Block) *AgentMessage {
	agentID := utils.Hex(evt.AgentId)
	evtName := DisableAgent
	if evt.Enabled {
		evtName = EnableAgent
	}
	return &AgentMessage{
		Message: regmsg.Message{
			Action:    evtName,
			Timestamp: time.Now().UTC(),
			Source:    regmsg.SourceFromBlock(evt.Raw.TxHash.Hex(), blk),
		},
		AgentID:    agentID,
		Permission: int(evt.Permission),
		TxHash:     evt.Raw.TxHash.Hex(),
	}
}

func NewAgentSaveMessageFromUpdate(evt *contract_agent_registry.AgentRegistryAgentUpdated, enabled bool, blk *domain.Block) *AgentSaveMessage {
	agentID := utils.Hex(evt.AgentId)
	return &AgentSaveMessage{
		AgentMessage: AgentMessage{
			AgentID: agentID,
			Message: regmsg.Message{
				Action:    SaveAgent,
				Timestamp: time.Now().UTC(),
				Source:    regmsg.SourceFromBlock(evt.Raw.TxHash.Hex(), blk),
			},
			TxHash: evt.Raw.TxHash.Hex(),
		},
		AgentProperties: AgentProperties{
			Enabled:  enabled,
			Name:     evt.Metadata,
			ChainIDs: utils.IntArray(evt.ChainIds),
			Metadata: evt.Metadata,
			Owner:    evt.By.Hex(),
		},
	}
}

func NewAgentSaveMessageFromTransfer(evt *contract_agent_registry.AgentRegistryTransfer, props AgentProperties, blk *domain.Block) *AgentSaveMessage {
	agentID := utils.Hex(evt.TokenId)
	return &AgentSaveMessage{
		AgentMessage: AgentMessage{
			AgentID: agentID,
			Message: regmsg.Message{
				Action:    SaveAgent,
				Timestamp: time.Now().UTC(),
				Source:    regmsg.SourceFromBlock(evt.Raw.TxHash.Hex(), blk),
			},
			TxHash: evt.Raw.TxHash.Hex(),
		},
		AgentProperties: props,
	}
}
