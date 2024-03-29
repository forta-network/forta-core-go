// Code generated by go-merge-types. DO NOT EDIT.

package contract_forta

import (
	import_fmt "fmt"
	import_sync "sync"

	forta020 "github.com/forta-network/forta-core-go/contracts/generated/contract_forta_0_2_0"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// FortaCaller is a new type which can multiplex calls to different implementation types.
type FortaCaller struct {
	typ0 *forta020.FortaCaller

	currTag string
	mu      import_sync.RWMutex
	unsafe  bool // default: false
}

// NewFortaCaller creates a new merged type.
func NewFortaCaller(address common.Address, caller bind.ContractCaller) (*FortaCaller, error) {
	var (
		mergedType FortaCaller
		err        error
	)
	mergedType.currTag = "0.2.0"

	mergedType.typ0, err = forta020.NewFortaCaller(address, caller)
	if err != nil {
		return nil, import_fmt.Errorf("failed to initialize forta020.FortaCaller: %v", err)
	}

	return &mergedType, nil
}

// IsKnownTagForFortaCaller tells if given tag is a known tag.
func IsKnownTagForFortaCaller(tag string) bool {

	if tag == "0.2.0" {
		return true
	}

	return false
}

// Use sets the used implementation to given tag.
func (merged *FortaCaller) Use(tag string) (changed bool) {
	if !merged.unsafe {
		merged.mu.Lock()
		defer merged.mu.Unlock()
	}
	// use the default tag if the provided tag is unknown
	if !IsKnownTagForFortaCaller(tag) {
		tag = "0.2.0"
	}
	changed = merged.currTag != tag
	merged.currTag = tag
	return
}

// Unsafe disables the mutex.
func (merged *FortaCaller) Unsafe() {
	merged.unsafe = true
}

// Safe enables the mutex.
func (merged *FortaCaller) Safe() {
	merged.unsafe = false
}

// ADMINROLE multiplexes to different implementations of the method.
func (merged *FortaCaller) ADMINROLE(opts *bind.CallOpts) (retVal [32]byte, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.ADMINROLE(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.ADMINROLE not implemented (tag=%s)", merged.currTag)
	return
}

// DEFAULTADMINROLE multiplexes to different implementations of the method.
func (merged *FortaCaller) DEFAULTADMINROLE(opts *bind.CallOpts) (retVal [32]byte, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.DEFAULTADMINROLE(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.DEFAULTADMINROLE not implemented (tag=%s)", merged.currTag)
	return
}

// DOMAINSEPARATOR multiplexes to different implementations of the method.
func (merged *FortaCaller) DOMAINSEPARATOR(opts *bind.CallOpts) (retVal [32]byte, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.DOMAINSEPARATOR(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.DOMAINSEPARATOR not implemented (tag=%s)", merged.currTag)
	return
}

// MINTERROLE multiplexes to different implementations of the method.
func (merged *FortaCaller) MINTERROLE(opts *bind.CallOpts) (retVal [32]byte, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.MINTERROLE(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.MINTERROLE not implemented (tag=%s)", merged.currTag)
	return
}

// SUPPLY multiplexes to different implementations of the method.
func (merged *FortaCaller) SUPPLY(opts *bind.CallOpts) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.SUPPLY(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.SUPPLY not implemented (tag=%s)", merged.currTag)
	return
}

// WHITELISTERROLE multiplexes to different implementations of the method.
func (merged *FortaCaller) WHITELISTERROLE(opts *bind.CallOpts) (retVal [32]byte, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.WHITELISTERROLE(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.WHITELISTERROLE not implemented (tag=%s)", merged.currTag)
	return
}

// WHITELISTROLE multiplexes to different implementations of the method.
func (merged *FortaCaller) WHITELISTROLE(opts *bind.CallOpts) (retVal [32]byte, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.WHITELISTROLE(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.WHITELISTROLE not implemented (tag=%s)", merged.currTag)
	return
}

// Allowance multiplexes to different implementations of the method.
func (merged *FortaCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.Allowance(opts, owner, spender)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.Allowance not implemented (tag=%s)", merged.currTag)
	return
}

// BalanceOf multiplexes to different implementations of the method.
func (merged *FortaCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.BalanceOf(opts, account)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.BalanceOf not implemented (tag=%s)", merged.currTag)
	return
}

// CheckpointsOutput is a merged return type.
type CheckpointsOutput struct {
	FromBlock uint32

	Votes *big.Int
}

// Checkpoints multiplexes to different implementations of the method.
func (merged *FortaCaller) Checkpoints(opts *bind.CallOpts, account common.Address, pos uint32) (retVal *CheckpointsOutput, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	retVal = &CheckpointsOutput{}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.Checkpoints(opts, account, pos)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal.FromBlock = val.FromBlock

		retVal.Votes = val.Votes

		return
	}

	err = import_fmt.Errorf("FortaCaller.Checkpoints not implemented (tag=%s)", merged.currTag)
	return
}

// Decimals multiplexes to different implementations of the method.
func (merged *FortaCaller) Decimals(opts *bind.CallOpts) (retVal uint8, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.Decimals(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.Decimals not implemented (tag=%s)", merged.currTag)
	return
}

// Delegates multiplexes to different implementations of the method.
func (merged *FortaCaller) Delegates(opts *bind.CallOpts, account common.Address) (retVal common.Address, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.Delegates(opts, account)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.Delegates not implemented (tag=%s)", merged.currTag)
	return
}

// GetPastTotalSupply multiplexes to different implementations of the method.
func (merged *FortaCaller) GetPastTotalSupply(opts *bind.CallOpts, blockNumber *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.GetPastTotalSupply(opts, blockNumber)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.GetPastTotalSupply not implemented (tag=%s)", merged.currTag)
	return
}

// GetPastVotes multiplexes to different implementations of the method.
func (merged *FortaCaller) GetPastVotes(opts *bind.CallOpts, account common.Address, blockNumber *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.GetPastVotes(opts, account, blockNumber)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.GetPastVotes not implemented (tag=%s)", merged.currTag)
	return
}

// GetRoleAdmin multiplexes to different implementations of the method.
func (merged *FortaCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) (retVal [32]byte, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.GetRoleAdmin(opts, role)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.GetRoleAdmin not implemented (tag=%s)", merged.currTag)
	return
}

// GetVotes multiplexes to different implementations of the method.
func (merged *FortaCaller) GetVotes(opts *bind.CallOpts, account common.Address) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.GetVotes(opts, account)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.GetVotes not implemented (tag=%s)", merged.currTag)
	return
}

// HasRole multiplexes to different implementations of the method.
func (merged *FortaCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (retVal bool, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.HasRole(opts, role, account)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.HasRole not implemented (tag=%s)", merged.currTag)
	return
}

// Name multiplexes to different implementations of the method.
func (merged *FortaCaller) Name(opts *bind.CallOpts) (retVal string, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.Name(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.Name not implemented (tag=%s)", merged.currTag)
	return
}

// Nonces multiplexes to different implementations of the method.
func (merged *FortaCaller) Nonces(opts *bind.CallOpts, owner common.Address) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.Nonces(opts, owner)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.Nonces not implemented (tag=%s)", merged.currTag)
	return
}

// NumCheckpoints multiplexes to different implementations of the method.
func (merged *FortaCaller) NumCheckpoints(opts *bind.CallOpts, account common.Address) (retVal uint32, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.NumCheckpoints(opts, account)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.NumCheckpoints not implemented (tag=%s)", merged.currTag)
	return
}

// ProxiableUUID multiplexes to different implementations of the method.
func (merged *FortaCaller) ProxiableUUID(opts *bind.CallOpts) (retVal [32]byte, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.ProxiableUUID(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.ProxiableUUID not implemented (tag=%s)", merged.currTag)
	return
}

// SupportsInterface multiplexes to different implementations of the method.
func (merged *FortaCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (retVal bool, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.SupportsInterface(opts, interfaceId)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.SupportsInterface not implemented (tag=%s)", merged.currTag)
	return
}

// Symbol multiplexes to different implementations of the method.
func (merged *FortaCaller) Symbol(opts *bind.CallOpts) (retVal string, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.Symbol(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.Symbol not implemented (tag=%s)", merged.currTag)
	return
}

// TotalSupply multiplexes to different implementations of the method.
func (merged *FortaCaller) TotalSupply(opts *bind.CallOpts) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.TotalSupply(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.TotalSupply not implemented (tag=%s)", merged.currTag)
	return
}

// Version multiplexes to different implementations of the method.
func (merged *FortaCaller) Version(opts *bind.CallOpts) (retVal string, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.2.0" {
		val, methodErr := merged.typ0.Version(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("FortaCaller.Version not implemented (tag=%s)", merged.currTag)
	return
}
