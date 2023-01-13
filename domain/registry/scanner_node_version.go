package registry

import (
	"time"

	"github.com/forta-network/forta-core-go/contracts/merged/contract_scanner_node_version"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/domain/registry/regmsg"
	"github.com/sirupsen/logrus"
)

const ScannerNodeVersionUpdated = "ScannerNodeVersionUpdated"

type ScannerNodeVersionMessage struct {
	regmsg.Message
	NewVersion string `json:"newVersion"`
	OldVersion string `json:"oldVersion"`
}

func (snvm *ScannerNodeVersionMessage) LogFields() logrus.Fields {
	return logrus.Fields{}
}

func NewScannerNodeVersionUpdated(evt *contract_scanner_node_version.ScannerNodeVersionScannerNodeVersionUpdated, blk *domain.Block) *ScannerNodeVersionMessage {
	return &ScannerNodeVersionMessage{
		Message: regmsg.Message{
			Timestamp: time.Now().UTC(),
			Action:    ScannerNodeVersionUpdated,
			Source:    regmsg.SourceFromBlock(evt.Raw.TxHash.Hex(), blk),
		},

		NewVersion: evt.NewVersion,
		OldVersion: evt.OldVersion,
	}
}
