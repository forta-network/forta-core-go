package registry

import (
	"encoding/json"
	"fmt"
	"github.com/forta-network/forta-core-go/contracts/contract_scanner_node_version"
	"github.com/forta-network/forta-core-go/domain"
	"time"
)

const NodeVersionUpdated = "ScannerNodeVersionUpdated"

type ScannerNodeVersionUpdated struct {
	Message
	NewVersion string `json:"new_version"`
	OldVersion string `json:"old_version"`
}

func ParseScannerNodeVersionUpdated(msg string) (*ScannerNodeVersionUpdated, error) {
	var nodeVersionUpdated ScannerNodeVersionUpdated
	err := json.Unmarshal([]byte(msg), &nodeVersionUpdated)
	if err != nil {
		return nil, err
	}
	if nodeVersionUpdated.Action != NodeVersionUpdated {
		return nil, fmt.Errorf("invalid action for ScannerNodeVersion: %s", nodeVersionUpdated.Action)
	}
	return &nodeVersionUpdated, nil
}

func NewScannerNodeVersionUpdated(evt *contract_scanner_node_version.ScannerNodeVersionScannerNodeVersionUpdated, blk *domain.Block) *ScannerNodeVersionUpdated {

	return &ScannerNodeVersionUpdated{
		Message: Message{
			Timestamp: time.Now().UTC(),
			Action:    NodeVersionUpdated,
			Source:    SourceFromBlock(evt.Raw.TxHash.Hex(), blk),
		},

		NewVersion: evt.NewVersion,
		OldVersion: evt.OldVersion,
	}
}
