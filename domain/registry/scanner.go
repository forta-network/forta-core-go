package registry

import (
	"fmt"
	"strings"

	"github.com/forta-network/forta-core-go/domain"

	"github.com/forta-network/forta-core-go/contracts/contract_scanner_pool_registry"
	"github.com/forta-network/forta-core-go/contracts/contract_scanner_registry"
	"github.com/forta-network/forta-core-go/utils"
	"github.com/goccy/go-json"
)

const SaveScanner = "SaveScanner"
const EnableScanner = "EnableScanner"
const DisableScanner = "DisableScanner"
const UpdateScannerPool = "UpdateScannerPool"

const ScannerPermissionAdmin = 0
const ScannerPermissionSelf = 1
const ScannerPermissionOwner = 2
const ScannerPermissionManager = 3

type ScannerMessage struct {
	Message
	ScannerID  string `json:"scannerId"`
	Permission int    `json:"permission"`
	Sender     string `json:"sender"`
}

type ScannerSaveMessage struct {
	ScannerMessage
	ChainID int64  `json:"chainId"`
	PoolID  string `json:"poolId"`
	Enabled bool   `json:"enabled"`
}

type UpdateScannerPoolMessage struct {
	Message
	PoolID  string  `json:"poolId"`
	ChainID *int64  `json:"chainId,omitempty"`
	Owner   *string `json:"owner,omitempty"`
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

func ParseUpdateScannerPoolMessage(msg string) (*UpdateScannerPoolMessage, error) {
	var m UpdateScannerPoolMessage
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
		Message:    MessageFrom(evt.Raw.TxHash.Hex(), blk, evtName),
		ScannerID:  strings.ToLower(scannerID),
		Permission: int(evt.Permission),
	}
}

func NewScannerMessageFromPool(evt *contract_scanner_pool_registry.ScannerPoolRegistryScannerEnabled, blk *domain.Block) *ScannerMessage {
	scannerID := utils.HexAddr(evt.ScannerId)
	evtName := DisableScanner
	if evt.Enabled {
		evtName = EnableScanner
	}
	return &ScannerMessage{
		Message:   MessageFrom(evt.Raw.TxHash.Hex(), blk, evtName),
		ScannerID: strings.ToLower(scannerID),
		Sender:    evt.Sender.Hex(),
	}
}

func NewScannerSaveMessage(evt *contract_scanner_registry.ScannerRegistryScannerUpdated, enabled bool, blk *domain.Block) *ScannerSaveMessage {
	scannerID := utils.HexAddr(evt.ScannerId)
	return &ScannerSaveMessage{
		ScannerMessage: ScannerMessage{
			ScannerID: strings.ToLower(scannerID),
			Message:   MessageFrom(evt.Raw.TxHash.Hex(), blk, SaveScanner),
		},
		ChainID: evt.ChainId.Int64(),
		Enabled: enabled,
	}
}

func NewScannerSaveMessageFromPool(evt *contract_scanner_pool_registry.ScannerPoolRegistryScannerUpdated, enabled bool, blk *domain.Block) *ScannerSaveMessage {
	scannerID := utils.HexAddr(evt.ScannerId)
	return &ScannerSaveMessage{
		ScannerMessage: ScannerMessage{
			ScannerID: strings.ToLower(scannerID),
			Message:   MessageFrom(evt.Raw.TxHash.Hex(), blk, SaveScanner),
		},
		ChainID: evt.ChainId.Int64(),
		PoolID:  utils.PoolIDToString(evt.ScannerPool),
		Enabled: enabled,
	}
}

func NewScannerPoolMessageFromTransfer(evt *contract_scanner_pool_registry.ScannerPoolRegistryTransfer, blk *domain.Block) *UpdateScannerPoolMessage {
	return &UpdateScannerPoolMessage{
		Message: MessageFrom(evt.Raw.TxHash.Hex(), blk, UpdateScannerPool),
		PoolID:  utils.PoolIDToString(evt.TokenId),
		Owner:   utils.StringPtr(evt.To.Hex()),
	}
}

func NewScannerPoolMessageFromRegistration(evt *contract_scanner_pool_registry.ScannerPoolRegistryScannerPoolRegistered, owner string, blk *domain.Block) *UpdateScannerPoolMessage {
	return &UpdateScannerPoolMessage{
		Message: MessageFrom(evt.Raw.TxHash.Hex(), blk, UpdateScannerPool),
		PoolID:  utils.PoolIDToString(evt.ScannerPoolId),
		Owner:   utils.StringPtr(owner),
		ChainID: utils.Int64Ptr(evt.ChainId.Int64()),
	}
}

func NewScannerPoolMessageFromEnablement(evt *contract_scanner_pool_registry.ScannerPoolRegistryEnabledScannersChanged, blk *domain.Block) *UpdateScannerPoolMessage {
	return &UpdateScannerPoolMessage{
		Message: MessageFrom(evt.Raw.TxHash.Hex(), blk, UpdateScannerPool),
		PoolID:  utils.PoolIDToString(evt.ScannerPoolId),
	}
}
