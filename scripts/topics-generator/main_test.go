package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/forta-network/forta-core-go/contracts/contract_agent_registry"
	"github.com/forta-network/forta-core-go/contracts/contract_dispatch"
	"github.com/forta-network/forta-core-go/contracts/contract_scanner_registry"
	"github.com/stretchr/testify/assert"
)

const evtScannerUpdated = "ScannerUpdated(uint256,uint256,string)"
const evtScannerEnabled = "ScannerEnabled(uint256,bool,uint8,bool)"
const evtAgentUpdated = "AgentUpdated(uint256,address,string,uint256[])"
const evtAgentEnabled = "AgentEnabled(uint256,bool,uint8,bool)"
const evtLink = "Link(uint256,uint256,bool)"

var evtScannerUpdatedTopic = crypto.Keccak256Hash([]byte(evtScannerUpdated)).Hex()
var evtScannerEnabledTopic = crypto.Keccak256Hash([]byte(evtScannerEnabled)).Hex()
var evtAgentUpdatedTopic = crypto.Keccak256Hash([]byte(evtAgentUpdated)).Hex()
var evtAgentEnabledTopic = crypto.Keccak256Hash([]byte(evtAgentEnabled)).Hex()
var evtLinkTopic = crypto.Keccak256Hash([]byte(evtLink)).Hex()

func TestTopicGeneration(t *testing.T) {
	assert.Equal(t, evtAgentEnabledTopic, contract_agent_registry.AgentEnabledTopic)
	assert.Equal(t, evtAgentUpdatedTopic, contract_agent_registry.AgentUpdatedTopic)
	assert.Equal(t, evtScannerUpdatedTopic, contract_scanner_registry.ScannerUpdatedTopic)
	assert.Equal(t, evtScannerEnabledTopic, contract_scanner_registry.ScannerEnabledTopic)
	assert.Equal(t, evtLinkTopic, contract_dispatch.LinkTopic)
}
