// Code generated by go-merge-types. DO NOT EDIT.

package contract_rewards_distributor

import (
	import_fmt "fmt"
	import_sync "sync"

	rewardsdistributor010 "github.com/forta-network/forta-core-go/contracts/generated/contract_rewards_distributor_0_1_0"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// RewardsDistributorCaller is a new type which can multiplex calls to different implementation types.
type RewardsDistributorCaller struct {
	typ0 *rewardsdistributor010.RewardsDistributorCaller

	currTag string
	mu      import_sync.RWMutex
	unsafe  bool // default: false
}

// NewRewardsDistributorCaller creates a new merged type.
func NewRewardsDistributorCaller(address common.Address, caller bind.ContractCaller) (*RewardsDistributorCaller, error) {
	var (
		mergedType RewardsDistributorCaller
		err        error
	)
	mergedType.currTag = "0.1.0"

	mergedType.typ0, err = rewardsdistributor010.NewRewardsDistributorCaller(address, caller)
	if err != nil {
		return nil, import_fmt.Errorf("failed to initialize rewardsdistributor010.RewardsDistributorCaller: %v", err)
	}

	return &mergedType, nil
}

// IsKnownTagForRewardsDistributorCaller tells if given tag is a known tag.
func IsKnownTagForRewardsDistributorCaller(tag string) bool {

	if tag == "0.1.0" {
		return true
	}

	return false
}

// Use sets the used implementation to given tag.
func (merged *RewardsDistributorCaller) Use(tag string) (changed bool) {
	if !merged.unsafe {
		merged.mu.Lock()
		defer merged.mu.Unlock()
	}
	// use the default tag if the provided tag is unknown
	if !IsKnownTagForRewardsDistributorCaller(tag) {
		tag = "0.1.0"
	}
	changed = merged.currTag != tag
	merged.currTag = tag
	return
}

// Unsafe disables the mutex.
func (merged *RewardsDistributorCaller) Unsafe() {
	merged.unsafe = true
}

// Safe enables the mutex.
func (merged *RewardsDistributorCaller) Safe() {
	merged.unsafe = false
}

// AvailableReward multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) AvailableReward(opts *bind.CallOpts, subjectType uint8, subjectId *big.Int, epochNumber *big.Int, staker common.Address) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.AvailableReward(opts, subjectType, subjectId, epochNumber, staker)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.AvailableReward not implemented (tag=%s)", merged.currTag)
	return
}

// ClaimedRewardsPerEpoch multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) ClaimedRewardsPerEpoch(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 common.Address) (retVal bool, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.ClaimedRewardsPerEpoch(opts, arg0, arg1, arg2)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.ClaimedRewardsPerEpoch not implemented (tag=%s)", merged.currTag)
	return
}

// DefaultFeeBps multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) DefaultFeeBps(opts *bind.CallOpts) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.DefaultFeeBps(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.DefaultFeeBps not implemented (tag=%s)", merged.currTag)
	return
}

// DelegationFeesOutput is a merged return type.
type DelegationFeesOutput struct {
	FeeBps uint16

	SinceEpoch *big.Int
}

// DelegationFees multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) DelegationFees(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (retVal *DelegationFeesOutput, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	retVal = &DelegationFeesOutput{}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.DelegationFees(opts, arg0, arg1)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal.FeeBps = val.FeeBps

		retVal.SinceEpoch = val.SinceEpoch

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.DelegationFees not implemented (tag=%s)", merged.currTag)
	return
}

// DelegationParamsEpochDelay multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) DelegationParamsEpochDelay(opts *bind.CallOpts) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.DelegationParamsEpochDelay(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.DelegationParamsEpochDelay not implemented (tag=%s)", merged.currTag)
	return
}

// GetCurrentEpochEndTimestamp multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) GetCurrentEpochEndTimestamp(opts *bind.CallOpts) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.GetCurrentEpochEndTimestamp(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.GetCurrentEpochEndTimestamp not implemented (tag=%s)", merged.currTag)
	return
}

// GetCurrentEpochNumber multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) GetCurrentEpochNumber(opts *bind.CallOpts) (retVal uint32, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.GetCurrentEpochNumber(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.GetCurrentEpochNumber not implemented (tag=%s)", merged.currTag)
	return
}

// GetCurrentEpochStartTimestamp multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) GetCurrentEpochStartTimestamp(opts *bind.CallOpts) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.GetCurrentEpochStartTimestamp(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.GetCurrentEpochStartTimestamp not implemented (tag=%s)", merged.currTag)
	return
}

// GetDelegatedSubjectType multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) GetDelegatedSubjectType(opts *bind.CallOpts, subjectType uint8) (retVal uint8, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.GetDelegatedSubjectType(opts, subjectType)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.GetDelegatedSubjectType not implemented (tag=%s)", merged.currTag)
	return
}

// GetDelegationFee multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) GetDelegationFee(opts *bind.CallOpts, subjectType uint8, subjectId *big.Int, epochNumber *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.GetDelegationFee(opts, subjectType, subjectId, epochNumber)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.GetDelegationFee not implemented (tag=%s)", merged.currTag)
	return
}

// GetDelegatorSubjectType multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) GetDelegatorSubjectType(opts *bind.CallOpts, subjectType uint8) (retVal uint8, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.GetDelegatorSubjectType(opts, subjectType)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.GetDelegatorSubjectType not implemented (tag=%s)", merged.currTag)
	return
}

// GetEpochEndTimestamp multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) GetEpochEndTimestamp(opts *bind.CallOpts, epochNumber *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.GetEpochEndTimestamp(opts, epochNumber)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.GetEpochEndTimestamp not implemented (tag=%s)", merged.currTag)
	return
}

// GetEpochNumber multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) GetEpochNumber(opts *bind.CallOpts, timestamp *big.Int) (retVal uint32, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.GetEpochNumber(opts, timestamp)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.GetEpochNumber not implemented (tag=%s)", merged.currTag)
	return
}

// GetEpochStartTimestamp multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) GetEpochStartTimestamp(opts *bind.CallOpts, epochNumber *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.GetEpochStartTimestamp(opts, epochNumber)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.GetEpochStartTimestamp not implemented (tag=%s)", merged.currTag)
	return
}

// GetSubjectTypeAgency multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) GetSubjectTypeAgency(opts *bind.CallOpts, subjectType uint8) (retVal uint8, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.GetSubjectTypeAgency(opts, subjectType)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.GetSubjectTypeAgency not implemented (tag=%s)", merged.currTag)
	return
}

// IsCurrentEpoch multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) IsCurrentEpoch(opts *bind.CallOpts, timestamp *big.Int) (retVal bool, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.IsCurrentEpoch(opts, timestamp)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.IsCurrentEpoch not implemented (tag=%s)", merged.currTag)
	return
}

// IsTrustedForwarder multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (retVal bool, err error) {
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

	err = import_fmt.Errorf("RewardsDistributorCaller.IsTrustedForwarder not implemented (tag=%s)", merged.currTag)
	return
}

// ProxiableUUID multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) ProxiableUUID(opts *bind.CallOpts) (retVal [32]byte, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.ProxiableUUID(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.ProxiableUUID not implemented (tag=%s)", merged.currTag)
	return
}

// RewardedEpochs multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) RewardedEpochs(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (retVal bool, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.RewardedEpochs(opts, arg0, arg1)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.RewardedEpochs not implemented (tag=%s)", merged.currTag)
	return
}

// RewardsPerEpoch multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) RewardsPerEpoch(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.RewardsPerEpoch(opts, arg0, arg1)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.RewardsPerEpoch not implemented (tag=%s)", merged.currTag)
	return
}

// RewardsToken multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) RewardsToken(opts *bind.CallOpts) (retVal common.Address, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.RewardsToken(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.RewardsToken not implemented (tag=%s)", merged.currTag)
	return
}

// UnclaimedRewards multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) UnclaimedRewards(opts *bind.CallOpts) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}

	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.UnclaimedRewards(opts)

		if methodErr != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	err = import_fmt.Errorf("RewardsDistributorCaller.UnclaimedRewards not implemented (tag=%s)", merged.currTag)
	return
}

// Version multiplexes to different implementations of the method.
func (merged *RewardsDistributorCaller) Version(opts *bind.CallOpts) (retVal string, err error) {
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

	err = import_fmt.Errorf("RewardsDistributorCaller.Version not implemented (tag=%s)", merged.currTag)
	return
}
