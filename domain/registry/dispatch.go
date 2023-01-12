package registry

import (
	"strings"
	"time"

	"github.com/forta-network/forta-core-go/domain"
	"github.com/sirupsen/logrus"

	"github.com/forta-network/forta-core-go/contracts/merged/contract_dispatch"
	"github.com/forta-network/forta-core-go/utils"
)

var Link = "Link"
var Unlink = "Unlink"

type DispatchMessage struct {
	Message
	ScannerID string `json:"scannerId"`
	AgentID   string `json:"agentId"`
}

func (dm *DispatchMessage) LogFields() logrus.Fields {
	return logrus.Fields{
		"scannerId": dm.ScannerID,
		"agentId":   dm.AgentID,
	}
}

func NewDispatchMessage(evt *contract_dispatch.DispatchLink, blk *domain.Block) *DispatchMessage {
	scannerID := utils.HexAddr(evt.ScannerId)
	agentID := utils.Hex(evt.AgentId)
	evtName := Unlink
	if evt.Enable {
		evtName = Link
	}
	return &DispatchMessage{
		Message: Message{
			Action:    evtName,
			Timestamp: time.Now().UTC(),
			Source:    SourceFromBlock(evt.Raw.TxHash.Hex(), blk),
		},
		ScannerID: strings.ToLower(scannerID),
		AgentID:   agentID,
	}
}

func NewAlreadyLinkedDispatchMessage(evt *contract_dispatch.DispatchAlreadyLinked, blk *domain.Block) *DispatchMessage {
	scannerID := utils.HexAddr(evt.ScannerId)
	agentID := utils.Hex(evt.AgentId)
	evtName := Unlink
	if evt.Enable {
		evtName = Link
	}
	return &DispatchMessage{
		Message: Message{
			Action:    evtName,
			Timestamp: time.Now().UTC(),
			Source:    SourceFromBlock(evt.Raw.TxHash.Hex(), blk),
		},
		ScannerID: strings.ToLower(scannerID),
		AgentID:   agentID,
	}
}
