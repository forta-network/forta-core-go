package registry

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/forta-network/forta-core-go/contracts/merged/contract_scanner_node_version"
	"github.com/forta-network/forta-core-go/domain"
)

const ScannerNodeVersionUpdated = "ScannerNodeVersionUpdated"

type ScannerNodeVersionMessage struct {
	Message
	NewVersion string `json:"newVersion"`
	OldVersion string `json:"oldVersion"`
}

func ParseScannerNodeVersionUpdated(msg string) (*ScannerNodeVersionMessage, error) {
	var nodeVersionUpdated ScannerNodeVersionMessage
	err := json.Unmarshal([]byte(msg), &nodeVersionUpdated)
	if err != nil {
		return nil, err
	}
	if nodeVersionUpdated.Action != ScannerNodeVersionUpdated {
		return nil, fmt.Errorf("invalid action for ScannerNodeVersion: %s", nodeVersionUpdated.Action)
	}
	return &nodeVersionUpdated, nil
}

func NewScannerNodeVersionUpdated(evt *contract_scanner_node_version.ScannerNodeVersionUpdatedEvent, blk *domain.Block) *ScannerNodeVersionMessage {

	return &ScannerNodeVersionMessage{
		Message: Message{
			Timestamp: time.Now().UTC(),
			Action:    ScannerNodeVersionUpdated,
			Source:    SourceFromBlock(evt.Raw.TxHash.Hex(), blk),
		},

		NewVersion: evt.NewVersion,
		OldVersion: evt.OldVersion,
	}
}
