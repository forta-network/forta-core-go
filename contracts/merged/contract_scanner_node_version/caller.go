// Code generated by go-merge-types. DO NOT EDIT.

package contract_scanner_node_version

import (
	import_fmt "fmt"
	import_sync "sync"

	scannernodeversion010 "github.com/forta-network/forta-core-go/contracts/generated/contract_scanner_node_version_0_1_0"

	scannernodeversion011 "github.com/forta-network/forta-core-go/contracts/generated/contract_scanner_node_version_0_1_1"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
)

// ScannerNodeVersionCaller is a new type which can multiplex calls to different implementation types.
type ScannerNodeVersionCaller struct {
	typ0 *scannernodeversion010.ScannerNodeVersionCaller

	typ1 *scannernodeversion011.ScannerNodeVersionCaller

	currTag string
	mu      import_sync.RWMutex
	unsafe  bool // default: false
}

// NewScannerNodeVersionCaller creates a new merged type.
func NewScannerNodeVersionCaller(address common.Address, caller bind.ContractCaller) (*ScannerNodeVersionCaller, error) {
	var (
		mergedType ScannerNodeVersionCaller
		err        error
	)
	mergedType.currTag = "0.1.1"

	mergedType.typ0, err = scannernodeversion010.NewScannerNodeVersionCaller(address, caller)
	if err != nil {
		return nil, import_fmt.Errorf("failed to initialize scannernodeversion010.ScannerNodeVersionCaller: %v", err)
	}

	mergedType.typ1, err = scannernodeversion011.NewScannerNodeVersionCaller(address, caller)
	if err != nil {
		return nil, import_fmt.Errorf("failed to initialize scannernodeversion011.ScannerNodeVersionCaller: %v", err)
	}

	return &mergedType, nil
}

// IsKnownTagForScannerNodeVersionCaller tells if given tag is a known tag.
func IsKnownTagForScannerNodeVersionCaller(tag string) bool {

	if tag == "0.1.0" {
		return true
	}

	if tag == "0.1.1" {
		return true
	}

	return false
}

// Use sets the used implementation to given tag.
func (merged *ScannerNodeVersionCaller) Use(tag string) (changed bool) {
	if !merged.unsafe {
		merged.mu.Lock()
		defer merged.mu.Unlock()
	}
	// use the default tag if the provided tag is unknown
	if !IsKnownTagForScannerNodeVersionCaller(tag) {
		tag = "0.1.1"
	}
	changed = merged.currTag != tag
	merged.currTag = tag
	return
}

// Unsafe disables the mutex.
func (merged *ScannerNodeVersionCaller) Unsafe() {
	merged.unsafe = true
}

// Safe enables the mutex.
func (merged *ScannerNodeVersionCaller) Safe() {
	merged.unsafe = false
}

// IsTrustedForwarder multiplexes to different implementations of the method.
func (merged *ScannerNodeVersionCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (retVal bool, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.IsTrustedForwarder(opts, forwarder)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ1.IsTrustedForwarder(opts, forwarder)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("ScannerNodeVersionCaller.IsTrustedForwarder not implemented (tag=%s)", merged.currTag)
	return
}

// ScannerNodeVersion multiplexes to different implementations of the method.
func (merged *ScannerNodeVersionCaller) ScannerNodeVersion(opts *bind.CallOpts) (retVal string, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.ScannerNodeVersion(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ1.ScannerNodeVersion(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("ScannerNodeVersionCaller.ScannerNodeVersion not implemented (tag=%s)", merged.currTag)
	return
}

// Version multiplexes to different implementations of the method.
func (merged *ScannerNodeVersionCaller) Version(opts *bind.CallOpts) (retVal string, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.Version(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ1.Version(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("ScannerNodeVersionCaller.Version not implemented (tag=%s)", merged.currTag)
	return
}

// ProxiableUUID multiplexes to different implementations of the method.
func (merged *ScannerNodeVersionCaller) ProxiableUUID(opts *bind.CallOpts) (retVal [32]byte, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ1.ProxiableUUID(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("ScannerNodeVersionCaller.ProxiableUUID not implemented (tag=%s)", merged.currTag)
	return
}

// ScannerNodeBetaVersion multiplexes to different implementations of the method.
func (merged *ScannerNodeVersionCaller) ScannerNodeBetaVersion(opts *bind.CallOpts) (retVal string, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.1" {
		val, methodErr := merged.typ1.ScannerNodeBetaVersion(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("ScannerNodeVersionCaller.ScannerNodeBetaVersion not implemented (tag=%s)", merged.currTag)
	return
}
