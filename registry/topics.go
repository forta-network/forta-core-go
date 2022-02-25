package registry

import "github.com/ethereum/go-ethereum/crypto"

const evtScannerUpdated = "ScannerUpdated(uint256,uint256)"
const evtScannerEnabled = "ScannerEnabled(uint256,bool,uint8,bool)"
const evtAgentUpdated = "AgentUpdated(uint256,address,string,uint256[])"
const evtAgentEnabled = "AgentEnabled(uint256,bool,uint8,bool)"
const evtLink = "Link(uint256,uint256,bool)"

var evtScannerUpdatedTopic = crypto.Keccak256Hash([]byte(evtScannerUpdated)).Hex()
var evtScannerEnabledTopic = crypto.Keccak256Hash([]byte(evtScannerEnabled)).Hex()
var evtAgentUpdatedTopic = crypto.Keccak256Hash([]byte(evtAgentUpdated)).Hex()
var evtAgentEnabledTopic = crypto.Keccak256Hash([]byte(evtAgentEnabled)).Hex()
var evtLinkTopic = crypto.Keccak256Hash([]byte(evtLink)).Hex()
