// Code generated by go-merge-types. DO NOT EDIT.

package contract_dispatch

import (
	import_fmt "fmt"
	import_sync "sync"


	dispatch014 "github.com/forta-network/forta-core-go/contracts/generated/contract_dispatch_0_1_4"

	dispatch015 "github.com/forta-network/forta-core-go/contracts/generated/contract_dispatch_0_1_5"



	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"math/big"

	"github.com/ethereum/go-ethereum/common"

)

// DispatchCaller is a new type which can multiplex calls to different implementation types.
type DispatchCaller struct {

	typ0 *dispatch014.DispatchCaller

	typ1 *dispatch015.DispatchCaller

	currTag string
	mu import_sync.RWMutex
	unsafe bool // default: false
}

// NewDispatchCaller creates a new merged type.
func NewDispatchCaller(address common.Address, caller bind.ContractCaller) (*DispatchCaller, error) {
	var (
		mergedType DispatchCaller
		err error
	)
	mergedType.currTag = "0.1.5"


	mergedType.typ0, err = dispatch014.NewDispatchCaller(address, caller)
	if err != nil {
		return nil, import_fmt.Errorf("failed to initialize dispatch014.DispatchCaller: %v", err)
	}

	mergedType.typ1, err = dispatch015.NewDispatchCaller(address, caller)
	if err != nil {
		return nil, import_fmt.Errorf("failed to initialize dispatch015.DispatchCaller: %v", err)
	}


	return &mergedType, nil
}

// IsKnownTagForDispatchCaller tells if given tag is a known tag.
func IsKnownTagForDispatchCaller(tag string) bool {

	if tag == "0.1.4" {
		return true
	}

	if tag == "0.1.5" {
		return true
	}

	return false
}

// Use sets the used implementation to given tag.
func (merged *DispatchCaller) Use(tag string) (changed bool) {
	if !merged.unsafe {
		merged.mu.Lock()
		defer merged.mu.Unlock()
	}
	// use the default tag if the provided tag is unknown
	if !IsKnownTagForDispatchCaller(tag) {
		tag = "0.1.5"
	}
	changed = merged.currTag != tag
	merged.currTag = tag
	return
}

// Unsafe disables the mutex.
func (merged *DispatchCaller) Unsafe() {
	merged.unsafe = true
}

// Safe enables the mutex.
func (merged *DispatchCaller) Safe() {
	merged.unsafe = false
}




// AgentAt multiplexes to different implementations of the method.
func (merged *DispatchCaller) AgentAt(opts *bind.CallOpts, scannerId *big.Int, pos *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.AgentAt(opts, scannerId, pos)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.AgentAt(opts, scannerId, pos)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchCaller.AgentAt not implemented (tag=%s)", merged.currTag)
	return
}


// AgentHashOutput is a merged return type.
type AgentHashOutput struct {

	Length *big.Int

	Manifest [32]byte

}

// AgentHash multiplexes to different implementations of the method.
func (merged *DispatchCaller) AgentHash(opts *bind.CallOpts, agentId *big.Int) (retVal *AgentHashOutput, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &AgentHashOutput{}



	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.AgentHash(opts, agentId)

		if methodErr != nil {
			err = methodErr
			return
		}


		retVal.Length = val.Length

		retVal.Manifest = val.Manifest


		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.AgentHash(opts, agentId)

		if methodErr != nil {
			err = methodErr
			return
		}


		retVal.Length = val.Length

		retVal.Manifest = val.Manifest


		return
	}


	err = import_fmt.Errorf("DispatchCaller.AgentHash not implemented (tag=%s)", merged.currTag)
	return
}


// AgentRefAtOutput is a merged return type.
type AgentRefAtOutput struct {

	Registered bool

	Owner common.Address

	AgentId *big.Int

	AgentVersion *big.Int

	Metadata string

	ChainIds []*big.Int

	Enabled bool

	DisabledFlags *big.Int

}

// AgentRefAt multiplexes to different implementations of the method.
func (merged *DispatchCaller) AgentRefAt(opts *bind.CallOpts, scannerId *big.Int, pos *big.Int) (retVal *AgentRefAtOutput, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &AgentRefAtOutput{}



	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.AgentRefAt(opts, scannerId, pos)

		if methodErr != nil {
			err = methodErr
			return
		}


		retVal.Registered = val.Registered

		retVal.Owner = val.Owner

		retVal.AgentId = val.AgentId

		retVal.AgentVersion = val.AgentVersion

		retVal.Metadata = val.Metadata

		retVal.ChainIds = val.ChainIds

		retVal.Enabled = val.Enabled

		retVal.DisabledFlags = val.DisabledFlags


		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.AgentRefAt(opts, scannerId, pos)

		if methodErr != nil {
			err = methodErr
			return
		}


		retVal.Registered = val.Registered

		retVal.Owner = val.Owner

		retVal.AgentId = val.AgentId

		retVal.AgentVersion = val.AgentVersion

		retVal.Metadata = val.Metadata

		retVal.ChainIds = val.ChainIds

		retVal.Enabled = val.Enabled

		retVal.DisabledFlags = val.DisabledFlags


		return
	}


	err = import_fmt.Errorf("DispatchCaller.AgentRefAt not implemented (tag=%s)", merged.currTag)
	return
}



// AgentRegistry multiplexes to different implementations of the method.
func (merged *DispatchCaller) AgentRegistry(opts *bind.CallOpts) (retVal common.Address, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.AgentRegistry(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.AgentRegistry(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchCaller.AgentRegistry not implemented (tag=%s)", merged.currTag)
	return
}



// AreTheyLinked multiplexes to different implementations of the method.
func (merged *DispatchCaller) AreTheyLinked(opts *bind.CallOpts, agentId *big.Int, scannerId *big.Int) (retVal bool, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.AreTheyLinked(opts, agentId, scannerId)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.AreTheyLinked(opts, agentId, scannerId)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchCaller.AreTheyLinked not implemented (tag=%s)", merged.currTag)
	return
}



// IsTrustedForwarder multiplexes to different implementations of the method.
func (merged *DispatchCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (retVal bool, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.IsTrustedForwarder(opts, forwarder)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchCaller.IsTrustedForwarder not implemented (tag=%s)", merged.currTag)
	return
}



// NumAgentsFor multiplexes to different implementations of the method.
func (merged *DispatchCaller) NumAgentsFor(opts *bind.CallOpts, scannerId *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.NumAgentsFor(opts, scannerId)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.NumAgentsFor(opts, scannerId)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchCaller.NumAgentsFor not implemented (tag=%s)", merged.currTag)
	return
}



// NumScannersFor multiplexes to different implementations of the method.
func (merged *DispatchCaller) NumScannersFor(opts *bind.CallOpts, agentId *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.NumScannersFor(opts, agentId)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.NumScannersFor(opts, agentId)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchCaller.NumScannersFor not implemented (tag=%s)", merged.currTag)
	return
}



// ProxiableUUID multiplexes to different implementations of the method.
func (merged *DispatchCaller) ProxiableUUID(opts *bind.CallOpts) (retVal [32]byte, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.ProxiableUUID(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.ProxiableUUID(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchCaller.ProxiableUUID not implemented (tag=%s)", merged.currTag)
	return
}



// ScannerAt multiplexes to different implementations of the method.
func (merged *DispatchCaller) ScannerAt(opts *bind.CallOpts, agentId *big.Int, pos *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.ScannerAt(opts, agentId, pos)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.ScannerAt(opts, agentId, pos)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchCaller.ScannerAt not implemented (tag=%s)", merged.currTag)
	return
}


// ScannerHashOutput is a merged return type.
type ScannerHashOutput struct {

	Length *big.Int

	Manifest [32]byte

}

// ScannerHash multiplexes to different implementations of the method.
func (merged *DispatchCaller) ScannerHash(opts *bind.CallOpts, scannerId *big.Int) (retVal *ScannerHashOutput, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &ScannerHashOutput{}



	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.ScannerHash(opts, scannerId)

		if methodErr != nil {
			err = methodErr
			return
		}


		retVal.Length = val.Length

		retVal.Manifest = val.Manifest


		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.ScannerHash(opts, scannerId)

		if methodErr != nil {
			err = methodErr
			return
		}


		retVal.Length = val.Length

		retVal.Manifest = val.Manifest


		return
	}


	err = import_fmt.Errorf("DispatchCaller.ScannerHash not implemented (tag=%s)", merged.currTag)
	return
}


// ScannerRefAtOutput is a merged return type.
type ScannerRefAtOutput struct {

	Registered bool

	ScannerId *big.Int

	Owner common.Address

	ChainId *big.Int

	Metadata string

	Enabled bool

	DisabledFlags *big.Int

	Operational bool

	Disabled bool

}

// ScannerRefAt multiplexes to different implementations of the method.
func (merged *DispatchCaller) ScannerRefAt(opts *bind.CallOpts, agentId *big.Int, pos *big.Int) (retVal *ScannerRefAtOutput, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &ScannerRefAtOutput{}



	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.ScannerRefAt(opts, agentId, pos)

		if methodErr != nil {
			err = methodErr
			return
		}


		retVal.Registered = val.Registered

		retVal.ScannerId = val.ScannerId

		retVal.Owner = val.Owner

		retVal.ChainId = val.ChainId

		retVal.Metadata = val.Metadata

		retVal.Enabled = val.Enabled

		retVal.DisabledFlags = val.DisabledFlags


		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.ScannerRefAt(opts, agentId, pos)

		if methodErr != nil {
			err = methodErr
			return
		}


		retVal.Registered = val.Registered

		retVal.ScannerId = val.ScannerId

		retVal.Owner = val.Owner

		retVal.ChainId = val.ChainId

		retVal.Metadata = val.Metadata

		retVal.Operational = val.Operational

		retVal.Disabled = val.Disabled


		return
	}


	err = import_fmt.Errorf("DispatchCaller.ScannerRefAt not implemented (tag=%s)", merged.currTag)
	return
}



// ScannerRegistry multiplexes to different implementations of the method.
func (merged *DispatchCaller) ScannerRegistry(opts *bind.CallOpts) (retVal common.Address, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.ScannerRegistry(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.ScannerRegistry(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchCaller.ScannerRegistry not implemented (tag=%s)", merged.currTag)
	return
}



// Version multiplexes to different implementations of the method.
func (merged *DispatchCaller) Version(opts *bind.CallOpts) (retVal string, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.Version(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.Version(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchCaller.Version not implemented (tag=%s)", merged.currTag)
	return
}



// ScannerPoolRegistry multiplexes to different implementations of the method.
func (merged *DispatchCaller) ScannerPoolRegistry(opts *bind.CallOpts) (retVal common.Address, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.ScannerPoolRegistry(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchCaller.ScannerPoolRegistry not implemented (tag=%s)", merged.currTag)
	return
}