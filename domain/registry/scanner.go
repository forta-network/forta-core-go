package registry

import (
	"fmt"
	"strings"
	"time"

	"github.com/forta-network/forta-core-go/domain"

	"github.com/forta-network/forta-core-go/contracts/contract_scanner_registry"
	"github.com/forta-network/forta-core-go/utils"
	"github.com/goccy/go-json"
)

const SaveScanner = "SaveScanner"
const EnableScanner = "EnableScanner"
const DisableScanner = "DisableScanner"

const ScannerPermissionAdmin = 0
const ScannerPermissionSelf = 1
const ScannerPermissionOwner = 2
const ScannerPermissionManager = 3

type ScannerMessage struct {
	Message
	ScannerID  string `json:"scannerId"`
	Permission int    `json:"permission"`
}

type ScannerSaveMessage struct {
	ScannerMessage
	ChainID int64 `json:"chainId"`
	Enabled bool  `json:"enabled"`
}

func ParseScannerSave(msg string) (*ScannerSaveMessage, error) {
	var save ScannerSaveMessage
	err := json.Unmarshal([]byte(msg), &save)
	if err != nil {
		return nil, err
	}
	if save.Action != SaveScanner {
		return nil, fmt.Errorf("invalid action for ScannerSave: %s", save.Action)
	}
	return &save, nil
}

func ParseScannerMessage(msg string) (*ScannerMessage, error) {
	var m ScannerMessage
	err := json.Unmarshal([]byte(msg), &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func NewScannerMessage(evt *contract_scanner_registry.ScannerRegistryScannerEnabled, blk *domain.Block) *ScannerMessage {
	scannerID := utils.HexAddr(evt.ScannerId)
	evtName := DisableScanner
	if evt.Enabled {
		evtName = EnableScanner
	}
	return &ScannerMessage{
		Message: Message{
			Timestamp: time.Now().UTC(),
			Action:    evtName,
			Source:    SourceFromBlock(evt.Raw.TxHash.Hex(), blk),
		},
		ScannerID:  strings.ToLower(scannerID),
		Permission: int(evt.Permission),
	}
}

func NewScannerSaveMessage(evt *contract_scanner_registry.ScannerRegistryScannerUpdated, enabled bool, blk *domain.Block) *ScannerSaveMessage {
	scannerID := utils.HexAddr(evt.ScannerId)
	return &ScannerSaveMessage{
		ScannerMessage: ScannerMessage{
			ScannerID: strings.ToLower(scannerID),
			Message: Message{
				Timestamp: time.Now().UTC(),
				Action:    SaveScanner,
				Source:    SourceFromBlock(evt.Raw.TxHash.Hex(), blk),
			},
		},
		Enabled: enabled,
		ChainID: evt.ChainId.Int64(),
	}
}
