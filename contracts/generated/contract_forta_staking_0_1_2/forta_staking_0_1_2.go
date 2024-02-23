// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_forta_staking_0_1_2

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// FortaStakingMetaData contains all meta data concerning the FortaStaking contract.
var FortaStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"AccessManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newWithdrawalDelay\",\"type\":\"uint256\"}],\"name\":\"DelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isFrozen\",\"type\":\"bool\"}],\"name\":\"Froze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"MaxStakeReached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"SlashDelegatorsPercentSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SlashedShareSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"subjectGateway\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allocator\",\"type\":\"address\"}],\"name\":\"StakeHelpersConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensSwept\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"TreasurySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WithdrawalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_SLASHABLE_PERCENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WITHDRAWAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_WITHDRAWAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"activeSharesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"activeSharesToStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"activeStakeFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocator\",\"outputs\":[{\"internalType\":\"contractIStakeAllocator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStakeSubjectGateway\",\"name\":\"_subjectGateway\",\"type\":\"address\"},{\"internalType\":\"contractIStakeAllocator\",\"name\":\"_allocator\",\"type\":\"address\"}],\"name\":\"configureStakeHelpers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeValue\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"freeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"}],\"name\":\"getDelegatedSubjectType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"}],\"name\":\"getDelegatorSubjectType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"}],\"name\":\"getSubjectTypeAgency\",\"outputs\":[{\"internalType\":\"enumSubjectTypeValidator.SubjectStakeAgency\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"inactiveSharesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"inactiveSharesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"inactiveSharesToStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"inactiveStakeFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__manager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"__stakedToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"__withdrawalDelay\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"__treasury\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sharesValue\",\"type\":\"uint256\"}],\"name\":\"initiateWithdrawal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"isFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"oldSubjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"oldSubject\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"newSubjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"newSubject\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"openProposals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"relayPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"setAccessManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newDelay\",\"type\":\"uint64\"}],\"name\":\"setDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ensRegistry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ensName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setReentrancyGuard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"setSlashDelegatorsPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"sharesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"proposerPercent\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashDelegatorsPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"activeSharesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stakeToActiveShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"inactiveSharesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stakeToInactiveShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subjectGateway\",\"outputs\":[{\"internalType\":\"contractIStakeSubjectGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweep\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalActiveStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"totalInactiveShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInactiveStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523060a0523480156200001557600080fd5b5060405162005d4538038062005d45833981016040819052620000389162000180565b6001600160a01b038116608052600054610100900460ff1615808015620000665750600054600160ff909116105b8062000096575062000083306200017160201b62002b3e1760201c565b15801562000096575060005460ff166001145b620000fe5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b6000805460ff19166001179055801562000122576000805461ff0019166101001790555b801562000169576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050620001b2565b6001600160a01b03163b151590565b6000602082840312156200019357600080fd5b81516001600160a01b0381168114620001ab57600080fd5b9392505050565b60805160a051615b54620001f1600039600081816115a1015281816115e101528181611ca301528181611ce30152611d72015260005050615b546000f3fe6080604052600436106103805760003560e01c806362772b14116101d1578063c07975ad11610102578063dc4653ef116100a0578063f08d35ee1161006f578063f08d35ee14610ae8578063f0f4426014610b08578063f242432a14610b28578063f8058b0614610b4857600080fd5b8063dc4653ef14610a31578063e985e9c514610a51578063edac6db814610a9b578063edea0bac14610ab057600080fd5b8063c9580804116100dc578063c9580804146109bb578063cc7a262e146109db578063d858a7e5146109fc578063da5bfb9414610a1157600080fd5b8063c07975ad1461095b578063c10733021461097b578063c133a5621461099b57600080fd5b8063a238f9df1161016f578063ac9650d811610149578063ac9650d8146108c0578063b4d60fb3146108ed578063b8dc491b1461090d578063bd85b0391461092d57600080fd5b8063a238f9df14610868578063a290bf381461087f578063aa5dcecc1461089f57600080fd5b8063762fa7b7116101ab578063762fa7b7146107ef5780638acdea4d1461081c57806390f1ccba14610831578063a22cb4651461084857600080fd5b806362772b141461078f57806364a0f901146107af57806375e130ad146107cf57600080fd5b80633121db1c116102b65780634a5f2b5d116102545780634f558e79116102235780634f558e79146106f357806352d1902d1461072357806354fd4d501461073857806361d027b31461076957600080fd5b80634a5f2b5d1461067c5780634e04275f146106935780634e1273f4146106b35780634f1ef286146106e057600080fd5b8063371f422611610290578063371f4226146105f85780633c3e04dd1461060e5780633f4899141461062e57806346e897871461064e57600080fd5b80633121db1c14610598578063321ebc54146105b85780633659cfe6146105d857600080fd5b80631a82ef4e11610323578063226cc300116102fd578063226cc3001461052257806328f73148146105425780632cb31144146105585780632eb2c2d61461057857600080fd5b80631a82ef4e146104975780631daa0445146104b75780631fd5a956146104e957600080fd5b8063093af6a91161035f578063093af6a91461040a5780630e10103f1461042a5780630e89341c1461044a578063115a90ee1461047757600080fd5b8062fdd58e1461038557806301ffc9a7146103b857806302fe5305146103e8575b600080fd5b34801561039157600080fd5b506103a56103a0366004614a64565b610b68565b6040519081526020015b60405180910390f35b3480156103c457600080fd5b506103d86103d3366004614aa6565b610c04565b60405190151581526020016103af565b3480156103f457600080fd5b50610408610403366004614b62565b610c54565b005b34801561041657600080fd5b50610408610425366004614bc0565b610c9f565b34801561043657600080fd5b506103a5610445366004614c1b565b610ef3565b34801561045657600080fd5b5061046a610465366004614c37565b610f09565b6040516103af9190614ca8565b34801561048357600080fd5b506103a5610492366004614cbb565b610f9e565b3480156104a357600080fd5b506103a56104b2366004614cbb565b610fe0565b3480156104c357600080fd5b506104d76104d2366004614cdd565b61101c565b60405160ff90911681526020016103af565b3480156104f557600080fd5b506101ce5461050a906001600160a01b031681565b6040516001600160a01b0390911681526020016103af565b34801561052e57600080fd5b5061040861053d366004614d06565b61103a565b34801561054e57600080fd5b506101c5546103a5565b34801561056457600080fd5b506103a5610573366004614d46565b6111af565b34801561058457600080fd5b50610408610593366004614e2d565b6114f4565b3480156105a457600080fd5b506104086105b3366004614eda565b611540565b3480156105c457600080fd5b506103a56105d3366004614c1b565b611587565b3480156105e457600080fd5b506104086105f3366004614f5e565b611596565b34801561060457600080fd5b506101c7546103a5565b34801561061a57600080fd5b50610408610629366004614f7b565b611676565b34801561063a57600080fd5b506103a5610649366004614c1b565b61177b565b34801561065a57600080fd5b506103a5610669366004614c37565b6101d16020526000908152604090205481565b34801561068857600080fd5b506103a56201518081565b34801561069f57600080fd5b506104086106ae366004614fcb565b611948565b3480156106bf57600080fd5b506106d36106ce366004615025565b611b6f565b6040516103af919061512c565b6104086106ee36600461513f565b611c98565b3480156106ff57600080fd5b506103d861070e366004614c37565b60009081526101916020526040902054151590565b34801561072f57600080fd5b506103a5611d65565b34801561074457600080fd5b5061046a60405180604001604052806005815260200164181718971960d91b81525081565b34801561077557600080fd5b506101cd54600160401b90046001600160a01b031661050a565b34801561079b57600080fd5b506103a56107aa366004614cbb565b611e19565b3480156107bb57600080fd5b506103a56107ca366004614cbb565b611e30565b3480156107db57600080fd5b506103d86107ea366004614c1b565b611e5f565b3480156107fb57600080fd5b5061080f61080a366004614cdd565b611ea1565b6040516103af91906151bc565b34801561082857600080fd5b506103a5605a81565b34801561083d57600080fd5b506103a56101cf5481565b34801561085457600080fd5b506104086108633660046151ca565b611ef7565b34801561087457600080fd5b506103a56276a70081565b34801561088b57600080fd5b506103a561089a366004614c1b565b611f02565b3480156108ab57600080fd5b506101d05461050a906001600160a01b031681565b3480156108cc57600080fd5b506108e06108db3660046151f8565b611f24565b6040516103af919061526c565b3480156108f957600080fd5b50610408610908366004614c37565b612011565b34801561091957600080fd5b506103a5610928366004614f7b565b612085565b34801561093957600080fd5b506103a5610948366004614c37565b6000908152610191602052604090205490565b34801561096757600080fd5b506104086109763660046152ce565b6121d1565b34801561098757600080fd5b50610408610996366004615315565b61226f565b3480156109a757600080fd5b506104d76109b6366004614cdd565b612357565b3480156109c757600080fd5b506104086109d6366004614f5e565b61236d565b3480156109e757600080fd5b506101c35461050a906001600160a01b031681565b348015610a0857600080fd5b50610408612427565b348015610a1d57600080fd5b506103a5610a2c366004615330565b6124b2565b348015610a3d57600080fd5b506103a5610a4c366004614c1b565b6124c2565b348015610a5d57600080fd5b506103d8610a6c366004614f7b565b6001600160a01b0391821660009081526101606020908152604080832093909416825291909152205460ff1690565b348015610aa757600080fd5b506104086124e4565b348015610abc57600080fd5b50610ad0610acb366004614d46565b612570565b6040516001600160401b0390911681526020016103af565b348015610af457600080fd5b506103a5610b03366004615330565b612831565b348015610b1457600080fd5b50610408610b23366004614f5e565b612841565b348015610b3457600080fd5b50610408610b43366004615365565b6128fd565b348015610b5457600080fd5b506103a5610b633660046153cd565b612942565b60006001600160a01b038316610bd85760405162461bcd60e51b815260206004820152602a60248201527f455243313135353a2061646472657373207a65726f206973206e6f742061207660448201526930b634b21037bbb732b960b11b60648201526084015b60405180910390fd5b50600081815261015f602090815260408083206001600160a01b03861684529091529020545b92915050565b60006001600160e01b03198216636cdb3d1360e11b1480610c3557506001600160e01b031982166303a24d0760e21b145b80610bfe57506301ffc9a760e01b6001600160e01b0319831614610bfe565b6000610c608133612b4d565b610c925780335b6040516301d4003760e61b815260048101929092526001600160a01b03166024820152604401610bcf565b610c9b82612bd2565b5050565b7fcbe0462e67cb804f0011a6c0b71e9ff49be70d0f907ffdf4f3c74499282ab0b1610cca8133612b4d565b610cd5578033610c67565b6002610ce16101d25490565b1415610cff5760405162461bcd60e51b8152600401610bcf9061541d565b610d0a60026101d255565b60ff861615610d315760405163c2628c0b60e01b815260ff87166004820152602401610bcf565b60ff8416600214610d5a5760405163c2628c0b60e01b815260ff85166004820152602401610bcf565b610d648686611e5f565b15610d8257604051632799ebef60e01b815260040160405180910390fd5b6000610d8e8787612be6565b90506000610d9c8483610b68565b90506000610daa8383611e30565b90506000610db88888612be6565b90506000610dc68284610fe0565b9050610dd56101c48685612c32565b610de26101c48385612c75565b610ded878686612cae565b610e2687838360005b6040519080825280601f01601f191660200182016040528015610e20576020820181803683370190505b50612e49565b866001600160a01b0316888a60ff167f3e5900eae50f376487f0d2272f8f857884c7adde3b5be2cd8c7b1e7bdfc47e1486604051610e6691815260200190565b60405180910390a46101d054604051636fda184760e11b81526001600160a01b039091169063dfb4308e90610ea99085908d908d908d908a908990600401615467565b600060405180830381600087803b158015610ec357600080fd5b505af1158015610ed7573d6000803e3d6000fd5b505050505050505050610eeb60016101d255565b505050505050565b6000610f026109488484612be6565b9392505050565b60606101618054610f199061549b565b80601f0160208091040260200160405190810160405280929190818152602001828054610f459061549b565b8015610f925780601f10610f6757610100808354040283529160200191610f92565b820191906000526020600020905b815481529060010190602001808311610f7557829003601f168201915b50505050509050919050565b600082815261019160205260408120548015610fd55760008481526101c66020526040902054610fd0905b8483612f66565b610fd8565b60005b949350505050565b60008281526101c4602052604081205481905b905080156110145760008481526101916020526040902054610fd090610fc9565b509092915050565b600060ff82166003141561103257506002919050565b5060ff919050565b7f12b42e8a160f6064dc959c6f251e3af0750ad213dbecf573b4710d67d6c28e396110658133612b4d565b611070578033610c67565b8360ff811615801590611087575060ff8116600114155b8015611097575060ff8116600214155b80156110a7575060ff8116600314155b156110ca5760405163c2628c0b60e01b815260ff82166004820152602401610bcf565b60006110d68686612be6565b90506110e181613016565b831561110c5760008181526101d160205260408120805491611102836154ec565b9190505550611158565b60008181526101d160205260409020546001111561112b576000611147565b60008181526101d1602052604090205461114790600190615507565b60008281526101d160205260409020555b60008181526101d1602090815260409182902054915191151582523391879160ff8a16917fd520b4ee12f45aacea6f06fd4831c1410e0d3f874a80f64424715327d7f705ca910160405180910390a4505050505050565b60008360ff8116158015906111c8575060ff8116600114155b80156111d8575060ff8116600214155b80156111e8575060ff8116600314155b1561120b5760405163c2628c0b60e01b815260ff82166004820152602401610bcf565b8460048061121883611ea1565b600481111561122957611229615184565b1415611256578161123983611ea1565b826040516365f3939760e01b8152600401610bcf9392919061551e565b60026112626101d25490565b14156112805760405162461bcd60e51b8152600401610bcf9061541d565b61128b60026101d255565b6101ce546001600160a01b03166112d65760405163eac0d38960e01b815260206004820152600e60248201526d7375626a6563744761746577617960901b6044820152606401610bcf565b6101ce5460405163882b291760e01b815260ff89166004820152602481018890526001600160a01b039091169063882b29179060440160206040518083038186803b15801561132457600080fd5b505afa158015611338573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061135c9190615542565b61137957604051632b3dcd7960e21b815260040160405180910390fd5b3360006113868989612be6565b905060006113958a8a8a613064565b909850905080156113d157604051899060ff8c16907f8db486b45abb5d250b92d9eeea290a901f784921d6d28705cfaf2e9fa025d56390600090a35b60006113dd838a610fe0565b6101c3549091506113f9906001600160a01b031685308c613157565b6114066101c4848b612c75565b6114138484836000610df6565b836001600160a01b03168a8c60ff167f3e5900eae50f376487f0d2272f8f857884c7adde3b5be2cd8c7b1e7bdfc47e148c60405161145391815260200190565b60405180910390a46101d060009054906101000a90046001600160a01b03166001600160a01b031663dfb4308e848d8d888e876040518763ffffffff1660e01b81526004016114a796959493929190615467565b600060405180830381600087803b1580156114c157600080fd5b505af11580156114d5573d6000803e3d6000fd5b509299505050505050506114ea60016101d255565b5050509392505050565b6001600160a01b03851633148061151057506115108533610a6c565b61152c5760405162461bcd60e51b8152600401610bcf9061555f565b61153985858585856131c2565b5050505050565b7f664245c7af190fec316596e8231f724e8171b1966cfcd124347ac5a66c34f64a61156b8133612b4d565b611576578033610c67565b6115818484846133aa565b50505050565b6000610f0261094884846134c0565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156115df5760405162461bcd60e51b8152600401610bcf906155ae565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316611628600080516020615ab8833981519152546001600160a01b031690565b6001600160a01b03161461164e5760405162461bcd60e51b8152600401610bcf906155fa565b61165781613508565b604080516000808252602082019092526116739183919061353e565b50565b60006116828133612b4d565b61168d578033610c67565b6001600160a01b0383166116d65760405163eac0d38960e01b815260206004820152600f60248201526e5f7375626a6563744761746577617960881b6044820152606401610bcf565b6001600160a01b03821661171a5760405163eac0d38960e01b815260206004820152600a6024820152692fb0b63637b1b0ba37b960b11b6044820152606401610bcf565b6101ce80546001600160a01b03199081166001600160a01b038681169182179093556101d080549092169285169283179091556040517fe727095b96a07a7ea4675d80358e4c6d0a7b287b367e4df2e1ee384e7d59ecc190600090a3505050565b60008260ff811615801590611794575060ff8116600114155b80156117a4575060ff8116600214155b80156117b4575060ff8116600314155b156117d75760405163c2628c0b60e01b815260ff82166004820152602401610bcf565b3360006117e486866134c0565b90506117f08282610b68565b61180d576040516365efc92d60e01b815260040160405180910390fd5b610100811760009081526101d160205260409020541561184057604051632799ebef60e01b815260040160405180910390fd5b610100811760009081526101c8602090815260408083206001600160a01b0386168452825291829020825191820190925281546001600160401b03168152611887906136bd565b6118a457604051630f2ca6e760e01b815260040160405180910390fd5b805467ffffffffffffffff19168155826001600160a01b0316868860ff167f07e9e8a51e2cf8929e95153e486868bda03206cdbbf30ee732b17e62b43c282760405160405180910390a460006118fa8484610b68565b905060006119088483610f9e565b90506119176101c68583612c32565b611922858584612cae565b6101c35461193a906001600160a01b031686836136ec565b9550505050505b5092915050565b600054610100900460ff16158080156119685750600054600160ff909116105b806119825750303b158015611982575060005460ff166001145b61199e5760405162461bcd60e51b8152600401610bcf90615646565b6000805460ff1916600117905580156119c1576000805461ff0019166101001790555b6001600160a01b038216611a055760405163eac0d38960e01b815260206004820152600a6024820152695f5f747265617375727960b01b6044820152606401610bcf565b6001600160a01b038416611a4c5760405163eac0d38960e01b815260206004820152600d60248201526c2fafb9ba30b5b2b22a37b5b2b760991b6044820152606401610bcf565b611a558561371c565b611a6d604051806020016040528060008152506137d6565b611a75613806565b6101cd80546001600160401b0385166001600160e01b03199091168117600160401b6001600160a01b0386811691909102919091179092556101c380546001600160a01b031916928716929092179091556040519081527f63e09f16584208fba1fc7ff64c62b00f07bec177c0d97ca6689891b1e77a35c79060200160405180910390a16040516001600160a01b03831681527f3c864541ef71378c6229510ed90f376565ee42d9c5e0904a984a9e863e6db44f9060200160405180910390a18015611539576000805461ff001916905560405160018152600080516020615ad88339815191529060200160405180910390a15050505050565b60608151835114611bd45760405162461bcd60e51b815260206004820152602960248201527f455243313135353a206163636f756e747320616e6420696473206c656e677468604482015268040dad2e6dac2e8c6d60bb1b6064820152608401610bcf565b600083516001600160401b03811115611bef57611bef614ac3565b604051908082528060200260200182016040528015611c18578160200160208202803683370190505b50905060005b8451811015611c9057611c63858281518110611c3c57611c3c615694565b6020026020010151858381518110611c5657611c56615694565b6020026020010151610b68565b828281518110611c7557611c75615694565b6020908102919091010152611c89816154ec565b9050611c1e565b509392505050565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161415611ce15760405162461bcd60e51b8152600401610bcf906155ae565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316611d2a600080516020615ab8833981519152546001600160a01b031690565b6001600160a01b031614611d505760405162461bcd60e51b8152600401610bcf906155fa565b611d5982613508565b610c9b8282600161353e565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611e055760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610bcf565b50600080516020615ab88339815191525b90565b60008281526101c660205260408120548190610ff3565b600082815261019160205260408120548015610fd55760008481526101c46020526040902054610fd090610fc9565b600080611e6c8484612be6565b60008181526101d16020526040902054909150151580610fd8575060009081526101cc602052604090205460ff169392505050565b600060ff821660011415611eb757506001919050565b60ff821660021415611ecb57506002919050565b60ff821660031415611edf57506003919050565b60ff8216611eef57506004919050565b506000919050565b610c9b33838361382f565b6000610f02611f118484612be6565b60009081526101c4602052604090205490565b6060816001600160401b03811115611f3e57611f3e614ac3565b604051908082528060200260200182016040528015611f7157816020015b6060815260200190600190039081611f5c5790505b50905060005b8281101561194157611fe130858584818110611f9557611f95615694565b9050602002810190611fa791906156aa565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061391192505050565b828281518110611ff357611ff3615694565b60200260200101819052508080612009906154ec565b915050611f77565b7f24791c44c040514a5d2580696fc45e7d3cb6c9fa65bf3db2e4755362d6c155b561203c8133612b4d565b612047578033610c67565b6101cf8290556040518281527f76096f684901589fef72a8b0a0321bf061b8597a791db667ab58e7eb922c639c906020015b60405180910390a15050565b60007f8aef0597c0be1e090afba1f387ee99f604b5d975ccbed6215cdf146ffd5c49fc6120b28133612b4d565b6120bd578033610c67565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a082319060240160206040518083038186803b1580156120ff57600080fd5b505afa158015612113573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061213791906156f7565b6101c3549091506001600160a01b0386811691161415612178576101c55461215f9082615507565b905061216b6101c75490565b6121759082615507565b90505b6121838585836136ec565b604080516001600160a01b038681168252602082018490528716917fd092d7fceb5ea5a962639fcc27a7bb315e7637e699e3b108cd570c38c7584300910160405180910390a2949350505050565b6101c3546001600160a01b031663d505accf336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152306024820152604481018890526064810187905260ff8616608482015260a4810185905260c4810184905260e401600060405180830381600087803b15801561225057600080fd5b505af1158015612264573d6000803e3d6000fd5b505050505050505050565b600061227b8133612b4d565b612286578033610c67565b62015180826001600160401b031610156122c657604051625a5b2760e31b81526001600160401b0383166004820152620151806024820152604401610bcf565b6276a700826001600160401b0316111561230757604051637f16270360e01b81526001600160401b03831660048201526276a7006024820152604401610bcf565b6101cd805467ffffffffffffffff19166001600160401b0384169081179091556040519081527f63e09f16584208fba1fc7ff64c62b00f07bec177c0d97ca6689891b1e77a35c790602001612079565b600060ff82166002141561103257506003919050565b60006123798133612b4d565b612384578033610c67565b61239e6001600160a01b038316637965db0b60e01b613936565b6123dc576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b6044820152606401610bcf565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a25050565b6065546001600160a01b03166124755760405163eac0d38960e01b81526020600482015260126024820152712fb232b83932b1b0ba32b22fb937baba32b960711b6044820152606401610bcf565b606580546001600160a01b03191690556040516000907f7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc80908290a2565b6000610fd8826103a086866134c0565b6000610f026124d184846134c0565b60009081526101c6602052604090205490565b600054600290610100900460ff16158015612506575060005460ff8083169116105b6125225760405162461bcd60e51b8152600401610bcf90615646565b6000805461ffff191660ff83161761010017905561253e613952565b6000805461ff001916905560405160ff82168152600080516020615ad88339815191529060200160405180910390a150565b60008360ff811615801590612589575060ff8116600114155b8015612599575060ff8116600214155b80156125a9575060ff8116600314155b156125cc5760405163c2628c0b60e01b815260ff82166004820152602401610bcf565b60026125d86101d25490565b14156125f65760405162461bcd60e51b8152600401610bcf9061541d565b61260160026101d255565b33600061260e8787612be6565b905061261a8282610b68565b6126375760405163d7d0b56f60e01b815260040160405180910390fd5b6101cd546000906001600160401b031661265042613984565b61265a9190615710565b60008381526101c8602090815260408083206001600160a01b03881684529091529020805467ffffffffffffffff19166001600160401b038316179055905060006126ae876126a98686610b68565b6139f0565b905060006126bc8483611e30565b905060006126cf61010019861683611e19565b905060006126dc8c611ea1565b90506126eb6101c48785612c32565b6126fd6101c661010019881685612c75565b612708878786612cae565b61271a87610100198816846000610df6565b600281600481111561272e5761272e615184565b148061274b5750600381600481111561274957612749615184565b145b156127cf576101d060009054906101000a90046001600160a01b03166001600160a01b031663d0d87ac8878e8e8b888a6040518763ffffffff1660e01b815260040161279c96959493929190615467565b600060405180830381600087803b1580156127b657600080fd5b505af11580156127ca573d6000803e3d6000fd5b505050505b6040516001600160401b03861681526001600160a01b038816908c9060ff8f16907f8b14266b7a7bfb46d73cda0a8ca3a635f38e8fb4b275b68aff68cb1b5a2aea3a9060200160405180910390a4509296505050505050611c9060016101d255565b6000610fd8826103a08686612be6565b600061284d8133612b4d565b612858578033610c67565b6001600160a01b03821661289d5760405163eac0d38960e01b815260206004820152600b60248201526a6e6577547265617375727960a81b6044820152606401610bcf565b6101cd805468010000000000000000600160e01b031916600160401b6001600160a01b038516908102919091179091556040519081527f3c864541ef71378c6229510ed90f376565ee42d9c5e0904a984a9e863e6db44f90602001612079565b6001600160a01b03851633148061291957506129198533610a6c565b6129355760405162461bcd60e51b8152600401610bcf9061555f565b6115398585858585613a06565b60007f12b42e8a160f6064dc959c6f251e3af0750ad213dbecf573b4710d67d6c28e3961296f8133612b4d565b61297a578033610c67565b8660038061298783611ea1565b600481111561299857612998615184565b14156129a8578161123983611ea1565b60006129b48a8a612be6565b905060026129c18b611ea1565b60048111156129d2576129d2615184565b1415612a3c5760006129e9896101cf546064612f66565b905060006129f7828b615507565b9050612a05838d8d84613b39565b8115612a35576000612a168d612357565b90506000612a24828e612be6565b9050612a3281838f87613b39565b50505b5050612a48565b612a48818b8b8b613b39565b6000612a5689886064612f66565b90508015612ab8576001600160a01b038816612aa05760405163eac0d38960e01b8152602060048201526008602482015267383937b837b9b2b960c11b6044820152606401610bcf565b6101c354612ab8906001600160a01b031689836136ec565b6101c3546101cd54612ae7916001600160a01b0390811691600160401b900416612ae2848d615507565b6136ec565b876001600160a01b03168a8c60ff167fbad5bf3ab3814a2220a4737f205e42060a9b3b66422e774059048cec2f6ed0ac84604051612b2791815260200190565b60405180910390a450969998505050505050505050565b6001600160a01b03163b151590565b603354604051632474521560e21b8152600481018490526001600160a01b03838116602483015260009216906391d148549060440160206040518083038186803b158015612b9a57600080fd5b505afa158015612bae573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f029190615542565b8051610c9b906101619060208401906149bf565b6040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff909116176101001790565b60008281526020849052604081208054839290612c50908490615507565b9250508190555080836001016000828254612c6b9190615507565b9091555050505050565b60008281526020849052604081208054839290612c9390849061573b565b9250508190555080836001016000828254612c6b919061573b565b6001600160a01b038316612d105760405162461bcd60e51b815260206004820152602360248201527f455243313135353a206275726e2066726f6d20746865207a65726f206164647260448201526265737360e81b6064820152608401610bcf565b336000612d1c84613ce0565b90506000612d2984613ce0565b9050612d4983876000858560405180602001604052806000815250613d2b565b600085815261015f602090815260408083206001600160a01b038a16845290915290205484811015612dc95760405162461bcd60e51b8152602060048201526024808201527f455243313135353a206275726e20616d6f756e7420657863656564732062616c604482015263616e636560e01b6064820152608401610bcf565b600086815261015f602090815260408083206001600160a01b038b81168086529184528285208a8703905582518b81529384018a90529092908816917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a46040805160208101909152600090525b50505050505050565b6001600160a01b038416612ea95760405162461bcd60e51b815260206004820152602160248201527f455243313135353a206d696e7420746f20746865207a65726f206164647265736044820152607360f81b6064820152608401610bcf565b336000612eb585613ce0565b90506000612ec285613ce0565b9050612ed383600089858589613d2b565b600086815261015f602090815260408083206001600160a01b038b16845290915281208054879290612f0690849061573b565b909155505060408051878152602081018790526001600160a01b03808a1692600092918716917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4612e4083600089898989613ee5565b600080806000198587098587029250828110838203039150508060001415612fa157838281612f9757612f97615753565b0492505050610f02565b808411612fad57600080fd5b60008486880960026001871981018816978890046003810283188082028403028082028403028082028403028082028403028082028403029081029092039091026000889003889004909101858311909403939093029303949094049190911702949350505050565b60008181526101cc602052604090205460ff16156116735760008181526101cc60209081526040808320805460ff191690556101d1909152812080549161305c836154ec565b919050555050565b6101ce546040516276841960e61b815260ff8516600482015260248101849052600091829182916001600160a01b031690631da106409060440160206040518083038186803b1580156130b657600080fd5b505afa1580156130ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130ee91906156f7565b9050806130fb8787611f02565b1061310e5760006001925092505061314f565b600061311a8787611f02565b6131249083615507565b905061313085826139f0565b828661313c8a8a611f02565b613146919061573b565b10159350935050505b935093915050565b6040516001600160a01b03808516602483015283166044820152606481018290526115819085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152614050565b81518351146132245760405162461bcd60e51b815260206004820152602860248201527f455243313135353a2069647320616e6420616d6f756e7473206c656e677468206044820152670dad2e6dac2e8c6d60c31b6064820152608401610bcf565b6001600160a01b03841661324a5760405162461bcd60e51b8152600401610bcf90615769565b33613259818787878787613d2b565b60005b845181101561334457600085828151811061327957613279615694565b60200260200101519050600085838151811061329757613297615694565b602090810291909101810151600084815261015f835260408082206001600160a01b038e1683529093529190912054909150818110156132e95760405162461bcd60e51b8152600401610bcf906157ae565b600083815261015f602090815260408083206001600160a01b038e8116855292528083208585039055908b1682528120805484929061332990849061573b565b925050819055505050508061333d906154ec565b905061325c565b50846001600160a01b0316866001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb87876040516133949291906157f8565b60405180910390a4610eeb818787878787614122565b6040516302571be360e01b81527f91d1777781884d03a6757a803996e38de2a42967fb37eeaca72729271025a9e260048201526001600160a01b038416906302571be39060240160206040518083038186803b15801561340957600080fd5b505afa15801561341d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613441919061581d565b6001600160a01b031663c47f002783836040518363ffffffff1660e01b815260040161346e92919061583a565b602060405180830381600087803b15801561348857600080fd5b505af115801561349c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061158191906156f7565b6040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff9091161790565b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e36135338133612b4d565b610c9b578033610c67565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff161561357657613571836141ec565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156135af57600080fd5b505afa9250505080156135df575060408051601f3d908101601f191682019092526135dc918101906156f7565b60015b6136425760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610bcf565b600080516020615ab883398151915281146136b15760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610bcf565b50613571838383614288565b60006136d282516001600160401b0316151590565b8015610bfe57505051426001600160401b03909116111590565b6040516001600160a01b03831660248201526044810182905261357190849063a9059cbb60e01b9060640161318b565b600054610100900460ff161580801561373c5750600054600160ff909116105b806137565750303b158015613756575060005460ff166001145b6137725760405162461bcd60e51b8152600401610bcf90615646565b6000805460ff191660011790558015613795576000805461ff0019166101001790555b61379e826142ad565b6137a6613806565b8015610c9b576000805461ff001916905560405160018152600080516020615ad883398151915290602001612079565b600054610100900460ff166137fd5760405162461bcd60e51b8152600401610bcf90615869565b611673816143f5565b600054610100900460ff1661382d5760405162461bcd60e51b8152600401610bcf90615869565b565b816001600160a01b0316836001600160a01b031614156138a35760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2073657474696e6720617070726f76616c20737461747573604482015268103337b91039b2b63360b91b6064820152608401610bcf565b6001600160a01b0383811660008181526101606020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6060610f028383604051806060016040528060278152602001615af860279139614425565b6000613941836144c3565b8015610f025750610f0283836144f6565b600054610100900460ff166139795760405162461bcd60e51b8152600401610bcf90615869565b61382d60016101d255565b60006001600160401b038211156139ec5760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203660448201526534206269747360d01b6064820152608401610bcf565b5090565b60008183106139ff5781610f02565b5090919050565b6001600160a01b038416613a2c5760405162461bcd60e51b8152600401610bcf90615769565b336000613a3885613ce0565b90506000613a4585613ce0565b9050613a55838989858589613d2b565b600086815261015f602090815260408083206001600160a01b038c16845290915290205485811015613a995760405162461bcd60e51b8152600401610bcf906157ae565b600087815261015f602090815260408083206001600160a01b038d8116855292528083208985039055908a16825281208054889290613ad990849061573b565b909155505060408051888152602081018890526001600160a01b03808b16928c821692918816917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4612264848a8a8a8a8a613ee5565b60008481526101c4602052604081205490613b586101001987166124d1565b90506000613b72613b69838561573b565b605a6064612f66565b905080841115613b9557604051631c67aae360e11b815260040160405180910390fd5b6000613bab8486613ba6868361573b565b612f66565b90506000613bb98287615507565b9050613bc86101c48a84612c32565b613bda6101c6610100198b1683612c32565b6000613be589611ea1565b90506002816004811115613bfb57613bfb615184565b1480613c1857506003816004811115613c1657613c16615184565b145b15613c8c576101d054604051631a1b0f5960e31b81526001600160a01b039091169063d0d87ac890613c59908d908d908d906000908a908290600401615467565b600060405180830381600087803b158015613c7357600080fd5b505af1158015613c87573d6000803e3d6000fd5b505050505b336001600160a01b0316888a60ff167f59d631535a5e48c1231a728e536ce39dba8d6c7f49905ec570e3db296430e02e8a604051613ccc91815260200190565b60405180910390a450505050505050505050565b60408051600180825281830190925260609160009190602080830190803683370190505090508281600081518110613d1a57613d1a615694565b602090810291909101015292915050565b60005b8351811015613ed657613d5e848281518110613d4c57613d4c615694565b60200260200101516101009081161490565b15613e8a576000613d85858381518110613d7a57613d7a615694565b602002602001015190565b905060ff81166003148015613da257506001600160a01b03861615155b8015613db657506001600160a01b03871615155b15613e84576101d05485516001600160a01b039091169063e2c11b8d90879085908110613de557613de5615694565b6020026020010151838a8a898881518110613e0257613e02615694565b60209081029190910101516040516001600160e01b031960e088901b168152600481019590955260ff90931660248501526001600160a01b039182166044850152166064830152608482015260a401600060405180830381600087803b158015613e6b57600080fd5b505af1158015613e7f573d6000803e3d6000fd5b505050505b50613ec4565b6001600160a01b0386161580613ea757506001600160a01b038516155b613ec45760405163068331f960e01b815260040160405180910390fd5b80613ece816154ec565b915050613d2e565b50610eeb8686868686866145d5565b6001600160a01b0384163b15610eeb5760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e6190613f2990899089908890889088906004016158b4565b602060405180830381600087803b158015613f4357600080fd5b505af1925050508015613f73575060408051601f3d908101601f19168201909252613f70918101906158ee565b60015b61402057613f7f61590b565b806308c379a01415613fb95750613f94615926565b80613f9f5750613fbb565b8060405162461bcd60e51b8152600401610bcf9190614ca8565b505b60405162461bcd60e51b815260206004820152603460248201527f455243313135353a207472616e7366657220746f206e6f6e20455243313135356044820152732932b1b2b4bb32b91034b6b83632b6b2b73a32b960611b6064820152608401610bcf565b6001600160e01b0319811663f23a6e6160e01b14612e405760405162461bcd60e51b8152600401610bcf906159af565b60006140a5826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166147519092919063ffffffff16565b80519091501561357157808060200190518101906140c39190615542565b6135715760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610bcf565b6001600160a01b0384163b15610eeb5760405163bc197c8160e01b81526001600160a01b0385169063bc197c819061416690899089908890889088906004016159f7565b602060405180830381600087803b15801561418057600080fd5b505af19250505080156141b0575060408051601f3d908101601f191682019092526141ad918101906158ee565b60015b6141bc57613f7f61590b565b6001600160e01b0319811663bc197c8160e01b14612e405760405162461bcd60e51b8152600401610bcf906159af565b6001600160a01b0381163b6142595760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610bcf565b600080516020615ab883398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b61429183614760565b60008251118061429e5750805b156135715761158183836147a0565b600054610100900460ff16158080156142cd5750600054600160ff909116105b806142e75750303b1580156142e7575060005460ff166001145b6143035760405162461bcd60e51b8152600401610bcf90615646565b6000805460ff191660011790558015614326576000805461ff0019166101001790555b6143406001600160a01b038316637965db0b60e01b613936565b61437e576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b6044820152606401610bcf565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a28015610c9b576000805461ff001916905560405160018152600080516020615ad883398151915290602001612079565b600054610100900460ff1661441c5760405162461bcd60e51b8152600401610bcf90615869565b61167381612bd2565b60606001600160a01b0384163b61444e5760405162461bcd60e51b8152600401610bcf90615a55565b600080856001600160a01b0316856040516144699190615a9b565b600060405180830381855af49150503d80600081146144a4576040519150601f19603f3d011682016040523d82523d6000602084013e6144a9565b606091505b50915091506144b9828286614855565b9695505050505050565b60006144d6826301ffc9a760e01b6144f6565b8015610bfe57506144ef826001600160e01b03196144f6565b1592915050565b604080516001600160e01b0319831660248083019190915282518083039091018152604490910182526020810180516001600160e01b03166301ffc9a760e01b179052905160009190829081906001600160a01b038716906175309061455d908690615a9b565b6000604051808303818686fa925050503d8060008114614599576040519150601f19603f3d011682016040523d82523d6000602084013e61459e565b606091505b50915091506020815110156145b95760009350505050610bfe565b8180156144b95750808060200190518101906144b99190615542565b6001600160a01b03851661465d5760005b835181101561465b5782818151811061460157614601615694565b6020026020010151610191600086848151811061462057614620615694565b602002602001015181526020019081526020016000206000828254614645919061573b565b909155506146549050816154ec565b90506145e6565b505b6001600160a01b038416610eeb5760005b8351811015612e4057600084828151811061468b5761468b615694565b6020026020010151905060008483815181106146a9576146a9615694565b60200260200101519050600061019160008481526020019081526020016000205490508181101561472d5760405162461bcd60e51b815260206004820152602860248201527f455243313135353a206275726e20616d6f756e74206578636565647320746f74604482015267616c537570706c7960c01b6064820152608401610bcf565b600092835261019160205260409092209103905561474a816154ec565b905061466e565b6060610fd8848460008561488e565b614769816141ec565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606001600160a01b0383163b6147c95760405162461bcd60e51b8152600401610bcf90615a55565b600080846001600160a01b0316846040516147e49190615a9b565b600060405180830381855af49150503d806000811461481f576040519150601f19603f3d011682016040523d82523d6000602084013e614824565b606091505b509150915061484c8282604051806060016040528060278152602001615af860279139614855565b95945050505050565b60608315614864575081610f02565b8251156148745782518084602001fd5b8160405162461bcd60e51b8152600401610bcf9190614ca8565b6060824710156148ef5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610bcf565b6001600160a01b0385163b6149465760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610bcf565b600080866001600160a01b031685876040516149629190615a9b565b60006040518083038185875af1925050503d806000811461499f576040519150601f19603f3d011682016040523d82523d6000602084013e6149a4565b606091505b50915091506149b4828286614855565b979650505050505050565b8280546149cb9061549b565b90600052602060002090601f0160209004810192826149ed5760008555614a33565b82601f10614a0657805160ff1916838001178555614a33565b82800160010185558215614a33579182015b82811115614a33578251825591602001919060010190614a18565b506139ec9291505b808211156139ec5760008155600101614a3b565b6001600160a01b038116811461167357600080fd5b60008060408385031215614a7757600080fd5b8235614a8281614a4f565b946020939093013593505050565b6001600160e01b03198116811461167357600080fd5b600060208284031215614ab857600080fd5b8135610f0281614a90565b634e487b7160e01b600052604160045260246000fd5b601f8201601f191681016001600160401b0381118282101715614afe57614afe614ac3565b6040525050565b60006001600160401b03831115614b1e57614b1e614ac3565b604051614b35601f8501601f191660200182614ad9565b809150838152848484011115614b4a57600080fd5b83836020830137600060208583010152509392505050565b600060208284031215614b7457600080fd5b81356001600160401b03811115614b8a57600080fd5b8201601f81018413614b9b57600080fd5b610fd884823560208401614b05565b803560ff81168114614bbb57600080fd5b919050565b600080600080600060a08688031215614bd857600080fd5b614be186614baa565b945060208601359350614bf660408701614baa565b9250606086013591506080860135614c0d81614a4f565b809150509295509295909350565b60008060408385031215614c2e57600080fd5b614a8283614baa565b600060208284031215614c4957600080fd5b5035919050565b60005b83811015614c6b578181015183820152602001614c53565b838111156115815750506000910152565b60008151808452614c94816020860160208601614c50565b601f01601f19169290920160200192915050565b602081526000610f026020830184614c7c565b60008060408385031215614cce57600080fd5b50508035926020909101359150565b600060208284031215614cef57600080fd5b610f0282614baa565b801515811461167357600080fd5b600080600060608486031215614d1b57600080fd5b614d2484614baa565b9250602084013591506040840135614d3b81614cf8565b809150509250925092565b600080600060608486031215614d5b57600080fd5b614d6484614baa565b95602085013595506040909401359392505050565b60006001600160401b03821115614d9257614d92614ac3565b5060051b60200190565b600082601f830112614dad57600080fd5b81356020614dba82614d79565b604051614dc78282614ad9565b83815260059390931b8501820192828101915086841115614de757600080fd5b8286015b84811015614e025780358352918301918301614deb565b509695505050505050565b600082601f830112614e1e57600080fd5b610f0283833560208501614b05565b600080600080600060a08688031215614e4557600080fd5b8535614e5081614a4f565b94506020860135614e6081614a4f565b935060408601356001600160401b0380821115614e7c57600080fd5b614e8889838a01614d9c565b94506060880135915080821115614e9e57600080fd5b614eaa89838a01614d9c565b93506080880135915080821115614ec057600080fd5b50614ecd88828901614e0d565b9150509295509295909350565b600080600060408486031215614eef57600080fd5b8335614efa81614a4f565b925060208401356001600160401b0380821115614f1657600080fd5b818601915086601f830112614f2a57600080fd5b813581811115614f3957600080fd5b876020828501011115614f4b57600080fd5b6020830194508093505050509250925092565b600060208284031215614f7057600080fd5b8135610f0281614a4f565b60008060408385031215614f8e57600080fd5b8235614f9981614a4f565b91506020830135614fa981614a4f565b809150509250929050565b80356001600160401b0381168114614bbb57600080fd5b60008060008060808587031215614fe157600080fd5b8435614fec81614a4f565b93506020850135614ffc81614a4f565b925061500a60408601614fb4565b9150606085013561501a81614a4f565b939692955090935050565b6000806040838503121561503857600080fd5b82356001600160401b038082111561504f57600080fd5b818501915085601f83011261506357600080fd5b8135602061507082614d79565b60405161507d8282614ad9565b83815260059390931b850182019282810191508984111561509d57600080fd5b948201945b838610156150c45785356150b581614a4f565b825294820194908201906150a2565b965050860135925050808211156150da57600080fd5b506150e785828601614d9c565b9150509250929050565b600081518084526020808501945080840160005b8381101561512157815187529582019590820190600101615105565b509495945050505050565b602081526000610f0260208301846150f1565b6000806040838503121561515257600080fd5b823561515d81614a4f565b915060208301356001600160401b0381111561517857600080fd5b6150e785828601614e0d565b634e487b7160e01b600052602160045260246000fd5b600581106151b857634e487b7160e01b600052602160045260246000fd5b9052565b60208101610bfe828461519a565b600080604083850312156151dd57600080fd5b82356151e881614a4f565b91506020830135614fa981614cf8565b6000806020838503121561520b57600080fd5b82356001600160401b038082111561522257600080fd5b818501915085601f83011261523657600080fd5b81358181111561524557600080fd5b8660208260051b850101111561525a57600080fd5b60209290920196919550909350505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b828110156152c157603f198886030184526152af858351614c7c565b94509285019290850190600101615293565b5092979650505050505050565b600080600080600060a086880312156152e657600080fd5b85359450602086013593506152fd60408701614baa565b94979396509394606081013594506080013592915050565b60006020828403121561532757600080fd5b610f0282614fb4565b60008060006060848603121561534557600080fd5b61534e84614baa565b9250602084013591506040840135614d3b81614a4f565b600080600080600060a0868803121561537d57600080fd5b853561538881614a4f565b9450602086013561539881614a4f565b9350604086013592506060860135915060808601356001600160401b038111156153c157600080fd5b614ecd88828901614e0d565b600080600080600060a086880312156153e557600080fd5b6153ee86614baa565b94506020860135935060408601359250606086013561540c81614a4f565b949793965091946080013592915050565b6020808252602a908201527f5265656e7472616e637947756172645570677261646561626c653a207265656e6040820152691d1c985b9d0818d85b1b60b21b606082015260800190565b95865260ff94909416602086015260408501929092526001600160a01b03166060840152608083015260a082015260c00190565b600181811c908216806154af57607f821691505b602082108114156154d057634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b6000600019821415615500576155006154d6565b5060010190565b600082821015615519576155196154d6565b500390565b60ff8416815260608101615535602083018561519a565b610fd8604083018461519a565b60006020828403121561555457600080fd5b8151610f0281614cf8565b6020808252602f908201527f455243313135353a2063616c6c6572206973206e6f7420746f6b656e206f776e60408201526e195c881b9bdc88185c1c1c9bdd9959608a1b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b6000808335601e198436030181126156c157600080fd5b8301803591506001600160401b038211156156db57600080fd5b6020019150368190038213156156f057600080fd5b9250929050565b60006020828403121561570957600080fd5b5051919050565b60006001600160401b03808316818516808303821115615732576157326154d6565b01949350505050565b6000821982111561574e5761574e6154d6565b500190565b634e487b7160e01b600052601260045260246000fd5b60208082526025908201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604082015264647265737360d81b606082015260800190565b6020808252602a908201527f455243313135353a20696e73756666696369656e742062616c616e636520666f60408201526939103a3930b739b332b960b11b606082015260800190565b60408152600061580b60408301856150f1565b828103602084015261484c81856150f1565b60006020828403121561582f57600080fd5b8151610f0281614a4f565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b6001600160a01b03868116825285166020820152604081018490526060810183905260a0608082018190526000906149b490830184614c7c565b60006020828403121561590057600080fd5b8151610f0281614a90565b600060033d1115611e165760046000803e5060005160e01c90565b600060443d10156159345790565b6040516003193d81016004833e81513d6001600160401b03816024840111818411171561596357505050505090565b828501915081518181111561597b5750505050505090565b843d87010160208285010111156159955750505050505090565b6159a460208286010187614ad9565b509095945050505050565b60208082526028908201527f455243313135353a204552433131353552656365697665722072656a656374656040820152676420746f6b656e7360c01b606082015260800190565b6001600160a01b0386811682528516602082015260a060408201819052600090615a23908301866150f1565b8281036060840152615a3581866150f1565b90508281036080840152615a498185614c7c565b98975050505050505050565b60208082526026908201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6040820152651b9d1c9858dd60d21b606082015260800190565b60008251615aad818460208701614c50565b919091019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220212aeb2c69b51c511fb6a697ca426e065a345009cbc101709671ddaca944028a64736f6c63430008090033",
}

// FortaStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use FortaStakingMetaData.ABI instead.
var FortaStakingABI = FortaStakingMetaData.ABI

// FortaStakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FortaStakingMetaData.Bin instead.
var FortaStakingBin = FortaStakingMetaData.Bin

// DeployFortaStaking deploys a new Ethereum contract, binding an instance of FortaStaking to it.
func DeployFortaStaking(auth *bind.TransactOpts, backend bind.ContractBackend, _forwarder common.Address) (common.Address, *types.Transaction, *FortaStaking, error) {
	parsed, err := FortaStakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FortaStakingBin), backend, _forwarder)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FortaStaking{FortaStakingCaller: FortaStakingCaller{contract: contract}, FortaStakingTransactor: FortaStakingTransactor{contract: contract}, FortaStakingFilterer: FortaStakingFilterer{contract: contract}}, nil
}

// FortaStaking is an auto generated Go binding around an Ethereum contract.
type FortaStaking struct {
	FortaStakingCaller     // Read-only binding to the contract
	FortaStakingTransactor // Write-only binding to the contract
	FortaStakingFilterer   // Log filterer for contract events
}

// FortaStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type FortaStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FortaStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FortaStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FortaStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FortaStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FortaStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FortaStakingSession struct {
	Contract     *FortaStaking     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FortaStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FortaStakingCallerSession struct {
	Contract *FortaStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// FortaStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FortaStakingTransactorSession struct {
	Contract     *FortaStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FortaStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type FortaStakingRaw struct {
	Contract *FortaStaking // Generic contract binding to access the raw methods on
}

// FortaStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FortaStakingCallerRaw struct {
	Contract *FortaStakingCaller // Generic read-only contract binding to access the raw methods on
}

// FortaStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FortaStakingTransactorRaw struct {
	Contract *FortaStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFortaStaking creates a new instance of FortaStaking, bound to a specific deployed contract.
func NewFortaStaking(address common.Address, backend bind.ContractBackend) (*FortaStaking, error) {
	contract, err := bindFortaStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FortaStaking{FortaStakingCaller: FortaStakingCaller{contract: contract}, FortaStakingTransactor: FortaStakingTransactor{contract: contract}, FortaStakingFilterer: FortaStakingFilterer{contract: contract}}, nil
}

// NewFortaStakingCaller creates a new read-only instance of FortaStaking, bound to a specific deployed contract.
func NewFortaStakingCaller(address common.Address, caller bind.ContractCaller) (*FortaStakingCaller, error) {
	contract, err := bindFortaStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FortaStakingCaller{contract: contract}, nil
}

// NewFortaStakingTransactor creates a new write-only instance of FortaStaking, bound to a specific deployed contract.
func NewFortaStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*FortaStakingTransactor, error) {
	contract, err := bindFortaStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FortaStakingTransactor{contract: contract}, nil
}

// NewFortaStakingFilterer creates a new log filterer instance of FortaStaking, bound to a specific deployed contract.
func NewFortaStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*FortaStakingFilterer, error) {
	contract, err := bindFortaStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FortaStakingFilterer{contract: contract}, nil
}

// bindFortaStaking binds a generic wrapper to an already deployed contract.
func bindFortaStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FortaStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FortaStaking *FortaStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FortaStaking.Contract.FortaStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FortaStaking *FortaStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FortaStaking.Contract.FortaStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FortaStaking *FortaStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FortaStaking.Contract.FortaStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FortaStaking *FortaStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FortaStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FortaStaking *FortaStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FortaStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FortaStaking *FortaStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FortaStaking.Contract.contract.Transact(opts, method, params...)
}

// MAXSLASHABLEPERCENT is a free data retrieval call binding the contract method 0x8acdea4d.
//
// Solidity: function MAX_SLASHABLE_PERCENT() view returns(uint256)
func (_FortaStaking *FortaStakingCaller) MAXSLASHABLEPERCENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "MAX_SLASHABLE_PERCENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSLASHABLEPERCENT is a free data retrieval call binding the contract method 0x8acdea4d.
//
// Solidity: function MAX_SLASHABLE_PERCENT() view returns(uint256)
func (_FortaStaking *FortaStakingSession) MAXSLASHABLEPERCENT() (*big.Int, error) {
	return _FortaStaking.Contract.MAXSLASHABLEPERCENT(&_FortaStaking.CallOpts)
}

// MAXSLASHABLEPERCENT is a free data retrieval call binding the contract method 0x8acdea4d.
//
// Solidity: function MAX_SLASHABLE_PERCENT() view returns(uint256)
func (_FortaStaking *FortaStakingCallerSession) MAXSLASHABLEPERCENT() (*big.Int, error) {
	return _FortaStaking.Contract.MAXSLASHABLEPERCENT(&_FortaStaking.CallOpts)
}

// MAXWITHDRAWALDELAY is a free data retrieval call binding the contract method 0xa238f9df.
//
// Solidity: function MAX_WITHDRAWAL_DELAY() view returns(uint256)
func (_FortaStaking *FortaStakingCaller) MAXWITHDRAWALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "MAX_WITHDRAWAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWITHDRAWALDELAY is a free data retrieval call binding the contract method 0xa238f9df.
//
// Solidity: function MAX_WITHDRAWAL_DELAY() view returns(uint256)
func (_FortaStaking *FortaStakingSession) MAXWITHDRAWALDELAY() (*big.Int, error) {
	return _FortaStaking.Contract.MAXWITHDRAWALDELAY(&_FortaStaking.CallOpts)
}

// MAXWITHDRAWALDELAY is a free data retrieval call binding the contract method 0xa238f9df.
//
// Solidity: function MAX_WITHDRAWAL_DELAY() view returns(uint256)
func (_FortaStaking *FortaStakingCallerSession) MAXWITHDRAWALDELAY() (*big.Int, error) {
	return _FortaStaking.Contract.MAXWITHDRAWALDELAY(&_FortaStaking.CallOpts)
}

// MINWITHDRAWALDELAY is a free data retrieval call binding the contract method 0x4a5f2b5d.
//
// Solidity: function MIN_WITHDRAWAL_DELAY() view returns(uint256)
func (_FortaStaking *FortaStakingCaller) MINWITHDRAWALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "MIN_WITHDRAWAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWITHDRAWALDELAY is a free data retrieval call binding the contract method 0x4a5f2b5d.
//
// Solidity: function MIN_WITHDRAWAL_DELAY() view returns(uint256)
func (_FortaStaking *FortaStakingSession) MINWITHDRAWALDELAY() (*big.Int, error) {
	return _FortaStaking.Contract.MINWITHDRAWALDELAY(&_FortaStaking.CallOpts)
}

// MINWITHDRAWALDELAY is a free data retrieval call binding the contract method 0x4a5f2b5d.
//
// Solidity: function MIN_WITHDRAWAL_DELAY() view returns(uint256)
func (_FortaStaking *FortaStakingCallerSession) MINWITHDRAWALDELAY() (*big.Int, error) {
	return _FortaStaking.Contract.MINWITHDRAWALDELAY(&_FortaStaking.CallOpts)
}

// ActiveSharesToStake is a free data retrieval call binding the contract method 0x64a0f901.
//
// Solidity: function activeSharesToStake(uint256 activeSharesId, uint256 amount) view returns(uint256)
func (_FortaStaking *FortaStakingCaller) ActiveSharesToStake(opts *bind.CallOpts, activeSharesId *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "activeSharesToStake", activeSharesId, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveSharesToStake is a free data retrieval call binding the contract method 0x64a0f901.
//
// Solidity: function activeSharesToStake(uint256 activeSharesId, uint256 amount) view returns(uint256)
func (_FortaStaking *FortaStakingSession) ActiveSharesToStake(activeSharesId *big.Int, amount *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.ActiveSharesToStake(&_FortaStaking.CallOpts, activeSharesId, amount)
}

// ActiveSharesToStake is a free data retrieval call binding the contract method 0x64a0f901.
//
// Solidity: function activeSharesToStake(uint256 activeSharesId, uint256 amount) view returns(uint256)
func (_FortaStaking *FortaStakingCallerSession) ActiveSharesToStake(activeSharesId *big.Int, amount *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.ActiveSharesToStake(&_FortaStaking.CallOpts, activeSharesId, amount)
}

// ActiveStakeFor is a free data retrieval call binding the contract method 0xa290bf38.
//
// Solidity: function activeStakeFor(uint8 subjectType, uint256 subject) view returns(uint256)
func (_FortaStaking *FortaStakingCaller) ActiveStakeFor(opts *bind.CallOpts, subjectType uint8, subject *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "activeStakeFor", subjectType, subject)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveStakeFor is a free data retrieval call binding the contract method 0xa290bf38.
//
// Solidity: function activeStakeFor(uint8 subjectType, uint256 subject) view returns(uint256)
func (_FortaStaking *FortaStakingSession) ActiveStakeFor(subjectType uint8, subject *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.ActiveStakeFor(&_FortaStaking.CallOpts, subjectType, subject)
}

// ActiveStakeFor is a free data retrieval call binding the contract method 0xa290bf38.
//
// Solidity: function activeStakeFor(uint8 subjectType, uint256 subject) view returns(uint256)
func (_FortaStaking *FortaStakingCallerSession) ActiveStakeFor(subjectType uint8, subject *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.ActiveStakeFor(&_FortaStaking.CallOpts, subjectType, subject)
}

// Allocator is a free data retrieval call binding the contract method 0xaa5dcecc.
//
// Solidity: function allocator() view returns(address)
func (_FortaStaking *FortaStakingCaller) Allocator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "allocator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Allocator is a free data retrieval call binding the contract method 0xaa5dcecc.
//
// Solidity: function allocator() view returns(address)
func (_FortaStaking *FortaStakingSession) Allocator() (common.Address, error) {
	return _FortaStaking.Contract.Allocator(&_FortaStaking.CallOpts)
}

// Allocator is a free data retrieval call binding the contract method 0xaa5dcecc.
//
// Solidity: function allocator() view returns(address)
func (_FortaStaking *FortaStakingCallerSession) Allocator() (common.Address, error) {
	return _FortaStaking.Contract.Allocator(&_FortaStaking.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_FortaStaking *FortaStakingCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_FortaStaking *FortaStakingSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.BalanceOf(&_FortaStaking.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_FortaStaking *FortaStakingCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.BalanceOf(&_FortaStaking.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_FortaStaking *FortaStakingCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_FortaStaking *FortaStakingSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _FortaStaking.Contract.BalanceOfBatch(&_FortaStaking.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_FortaStaking *FortaStakingCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _FortaStaking.Contract.BalanceOfBatch(&_FortaStaking.CallOpts, accounts, ids)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_FortaStaking *FortaStakingCaller) Exists(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "exists", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_FortaStaking *FortaStakingSession) Exists(id *big.Int) (bool, error) {
	return _FortaStaking.Contract.Exists(&_FortaStaking.CallOpts, id)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_FortaStaking *FortaStakingCallerSession) Exists(id *big.Int) (bool, error) {
	return _FortaStaking.Contract.Exists(&_FortaStaking.CallOpts, id)
}

// GetDelegatedSubjectType is a free data retrieval call binding the contract method 0x1daa0445.
//
// Solidity: function getDelegatedSubjectType(uint8 subjectType) pure returns(uint8)
func (_FortaStaking *FortaStakingCaller) GetDelegatedSubjectType(opts *bind.CallOpts, subjectType uint8) (uint8, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "getDelegatedSubjectType", subjectType)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetDelegatedSubjectType is a free data retrieval call binding the contract method 0x1daa0445.
//
// Solidity: function getDelegatedSubjectType(uint8 subjectType) pure returns(uint8)
func (_FortaStaking *FortaStakingSession) GetDelegatedSubjectType(subjectType uint8) (uint8, error) {
	return _FortaStaking.Contract.GetDelegatedSubjectType(&_FortaStaking.CallOpts, subjectType)
}

// GetDelegatedSubjectType is a free data retrieval call binding the contract method 0x1daa0445.
//
// Solidity: function getDelegatedSubjectType(uint8 subjectType) pure returns(uint8)
func (_FortaStaking *FortaStakingCallerSession) GetDelegatedSubjectType(subjectType uint8) (uint8, error) {
	return _FortaStaking.Contract.GetDelegatedSubjectType(&_FortaStaking.CallOpts, subjectType)
}

// GetDelegatorSubjectType is a free data retrieval call binding the contract method 0xc133a562.
//
// Solidity: function getDelegatorSubjectType(uint8 subjectType) pure returns(uint8)
func (_FortaStaking *FortaStakingCaller) GetDelegatorSubjectType(opts *bind.CallOpts, subjectType uint8) (uint8, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "getDelegatorSubjectType", subjectType)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetDelegatorSubjectType is a free data retrieval call binding the contract method 0xc133a562.
//
// Solidity: function getDelegatorSubjectType(uint8 subjectType) pure returns(uint8)
func (_FortaStaking *FortaStakingSession) GetDelegatorSubjectType(subjectType uint8) (uint8, error) {
	return _FortaStaking.Contract.GetDelegatorSubjectType(&_FortaStaking.CallOpts, subjectType)
}

// GetDelegatorSubjectType is a free data retrieval call binding the contract method 0xc133a562.
//
// Solidity: function getDelegatorSubjectType(uint8 subjectType) pure returns(uint8)
func (_FortaStaking *FortaStakingCallerSession) GetDelegatorSubjectType(subjectType uint8) (uint8, error) {
	return _FortaStaking.Contract.GetDelegatorSubjectType(&_FortaStaking.CallOpts, subjectType)
}

// GetSubjectTypeAgency is a free data retrieval call binding the contract method 0x762fa7b7.
//
// Solidity: function getSubjectTypeAgency(uint8 subjectType) pure returns(uint8)
func (_FortaStaking *FortaStakingCaller) GetSubjectTypeAgency(opts *bind.CallOpts, subjectType uint8) (uint8, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "getSubjectTypeAgency", subjectType)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetSubjectTypeAgency is a free data retrieval call binding the contract method 0x762fa7b7.
//
// Solidity: function getSubjectTypeAgency(uint8 subjectType) pure returns(uint8)
func (_FortaStaking *FortaStakingSession) GetSubjectTypeAgency(subjectType uint8) (uint8, error) {
	return _FortaStaking.Contract.GetSubjectTypeAgency(&_FortaStaking.CallOpts, subjectType)
}

// GetSubjectTypeAgency is a free data retrieval call binding the contract method 0x762fa7b7.
//
// Solidity: function getSubjectTypeAgency(uint8 subjectType) pure returns(uint8)
func (_FortaStaking *FortaStakingCallerSession) GetSubjectTypeAgency(subjectType uint8) (uint8, error) {
	return _FortaStaking.Contract.GetSubjectTypeAgency(&_FortaStaking.CallOpts, subjectType)
}

// InactiveSharesOf is a free data retrieval call binding the contract method 0xda5bfb94.
//
// Solidity: function inactiveSharesOf(uint8 subjectType, uint256 subject, address account) view returns(uint256)
func (_FortaStaking *FortaStakingCaller) InactiveSharesOf(opts *bind.CallOpts, subjectType uint8, subject *big.Int, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "inactiveSharesOf", subjectType, subject, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InactiveSharesOf is a free data retrieval call binding the contract method 0xda5bfb94.
//
// Solidity: function inactiveSharesOf(uint8 subjectType, uint256 subject, address account) view returns(uint256)
func (_FortaStaking *FortaStakingSession) InactiveSharesOf(subjectType uint8, subject *big.Int, account common.Address) (*big.Int, error) {
	return _FortaStaking.Contract.InactiveSharesOf(&_FortaStaking.CallOpts, subjectType, subject, account)
}

// InactiveSharesOf is a free data retrieval call binding the contract method 0xda5bfb94.
//
// Solidity: function inactiveSharesOf(uint8 subjectType, uint256 subject, address account) view returns(uint256)
func (_FortaStaking *FortaStakingCallerSession) InactiveSharesOf(subjectType uint8, subject *big.Int, account common.Address) (*big.Int, error) {
	return _FortaStaking.Contract.InactiveSharesOf(&_FortaStaking.CallOpts, subjectType, subject, account)
}

// InactiveSharesToStake is a free data retrieval call binding the contract method 0x115a90ee.
//
// Solidity: function inactiveSharesToStake(uint256 inactiveSharesId, uint256 amount) view returns(uint256)
func (_FortaStaking *FortaStakingCaller) InactiveSharesToStake(opts *bind.CallOpts, inactiveSharesId *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "inactiveSharesToStake", inactiveSharesId, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InactiveSharesToStake is a free data retrieval call binding the contract method 0x115a90ee.
//
// Solidity: function inactiveSharesToStake(uint256 inactiveSharesId, uint256 amount) view returns(uint256)
func (_FortaStaking *FortaStakingSession) InactiveSharesToStake(inactiveSharesId *big.Int, amount *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.InactiveSharesToStake(&_FortaStaking.CallOpts, inactiveSharesId, amount)
}

// InactiveSharesToStake is a free data retrieval call binding the contract method 0x115a90ee.
//
// Solidity: function inactiveSharesToStake(uint256 inactiveSharesId, uint256 amount) view returns(uint256)
func (_FortaStaking *FortaStakingCallerSession) InactiveSharesToStake(inactiveSharesId *big.Int, amount *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.InactiveSharesToStake(&_FortaStaking.CallOpts, inactiveSharesId, amount)
}

// InactiveStakeFor is a free data retrieval call binding the contract method 0xdc4653ef.
//
// Solidity: function inactiveStakeFor(uint8 subjectType, uint256 subject) view returns(uint256)
func (_FortaStaking *FortaStakingCaller) InactiveStakeFor(opts *bind.CallOpts, subjectType uint8, subject *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "inactiveStakeFor", subjectType, subject)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InactiveStakeFor is a free data retrieval call binding the contract method 0xdc4653ef.
//
// Solidity: function inactiveStakeFor(uint8 subjectType, uint256 subject) view returns(uint256)
func (_FortaStaking *FortaStakingSession) InactiveStakeFor(subjectType uint8, subject *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.InactiveStakeFor(&_FortaStaking.CallOpts, subjectType, subject)
}

// InactiveStakeFor is a free data retrieval call binding the contract method 0xdc4653ef.
//
// Solidity: function inactiveStakeFor(uint8 subjectType, uint256 subject) view returns(uint256)
func (_FortaStaking *FortaStakingCallerSession) InactiveStakeFor(subjectType uint8, subject *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.InactiveStakeFor(&_FortaStaking.CallOpts, subjectType, subject)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_FortaStaking *FortaStakingCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_FortaStaking *FortaStakingSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _FortaStaking.Contract.IsApprovedForAll(&_FortaStaking.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_FortaStaking *FortaStakingCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _FortaStaking.Contract.IsApprovedForAll(&_FortaStaking.CallOpts, account, operator)
}

// IsFrozen is a free data retrieval call binding the contract method 0x75e130ad.
//
// Solidity: function isFrozen(uint8 subjectType, uint256 subject) view returns(bool)
func (_FortaStaking *FortaStakingCaller) IsFrozen(opts *bind.CallOpts, subjectType uint8, subject *big.Int) (bool, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "isFrozen", subjectType, subject)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFrozen is a free data retrieval call binding the contract method 0x75e130ad.
//
// Solidity: function isFrozen(uint8 subjectType, uint256 subject) view returns(bool)
func (_FortaStaking *FortaStakingSession) IsFrozen(subjectType uint8, subject *big.Int) (bool, error) {
	return _FortaStaking.Contract.IsFrozen(&_FortaStaking.CallOpts, subjectType, subject)
}

// IsFrozen is a free data retrieval call binding the contract method 0x75e130ad.
//
// Solidity: function isFrozen(uint8 subjectType, uint256 subject) view returns(bool)
func (_FortaStaking *FortaStakingCallerSession) IsFrozen(subjectType uint8, subject *big.Int) (bool, error) {
	return _FortaStaking.Contract.IsFrozen(&_FortaStaking.CallOpts, subjectType, subject)
}

// OpenProposals is a free data retrieval call binding the contract method 0x46e89787.
//
// Solidity: function openProposals(uint256 ) view returns(uint256)
func (_FortaStaking *FortaStakingCaller) OpenProposals(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "openProposals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpenProposals is a free data retrieval call binding the contract method 0x46e89787.
//
// Solidity: function openProposals(uint256 ) view returns(uint256)
func (_FortaStaking *FortaStakingSession) OpenProposals(arg0 *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.OpenProposals(&_FortaStaking.CallOpts, arg0)
}

// OpenProposals is a free data retrieval call binding the contract method 0x46e89787.
//
// Solidity: function openProposals(uint256 ) view returns(uint256)
func (_FortaStaking *FortaStakingCallerSession) OpenProposals(arg0 *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.OpenProposals(&_FortaStaking.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FortaStaking *FortaStakingCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FortaStaking *FortaStakingSession) ProxiableUUID() ([32]byte, error) {
	return _FortaStaking.Contract.ProxiableUUID(&_FortaStaking.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FortaStaking *FortaStakingCallerSession) ProxiableUUID() ([32]byte, error) {
	return _FortaStaking.Contract.ProxiableUUID(&_FortaStaking.CallOpts)
}

// SharesOf is a free data retrieval call binding the contract method 0xf08d35ee.
//
// Solidity: function sharesOf(uint8 subjectType, uint256 subject, address account) view returns(uint256)
func (_FortaStaking *FortaStakingCaller) SharesOf(opts *bind.CallOpts, subjectType uint8, subject *big.Int, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "sharesOf", subjectType, subject, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesOf is a free data retrieval call binding the contract method 0xf08d35ee.
//
// Solidity: function sharesOf(uint8 subjectType, uint256 subject, address account) view returns(uint256)
func (_FortaStaking *FortaStakingSession) SharesOf(subjectType uint8, subject *big.Int, account common.Address) (*big.Int, error) {
	return _FortaStaking.Contract.SharesOf(&_FortaStaking.CallOpts, subjectType, subject, account)
}

// SharesOf is a free data retrieval call binding the contract method 0xf08d35ee.
//
// Solidity: function sharesOf(uint8 subjectType, uint256 subject, address account) view returns(uint256)
func (_FortaStaking *FortaStakingCallerSession) SharesOf(subjectType uint8, subject *big.Int, account common.Address) (*big.Int, error) {
	return _FortaStaking.Contract.SharesOf(&_FortaStaking.CallOpts, subjectType, subject, account)
}

// SlashDelegatorsPercent is a free data retrieval call binding the contract method 0x90f1ccba.
//
// Solidity: function slashDelegatorsPercent() view returns(uint256)
func (_FortaStaking *FortaStakingCaller) SlashDelegatorsPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "slashDelegatorsPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashDelegatorsPercent is a free data retrieval call binding the contract method 0x90f1ccba.
//
// Solidity: function slashDelegatorsPercent() view returns(uint256)
func (_FortaStaking *FortaStakingSession) SlashDelegatorsPercent() (*big.Int, error) {
	return _FortaStaking.Contract.SlashDelegatorsPercent(&_FortaStaking.CallOpts)
}

// SlashDelegatorsPercent is a free data retrieval call binding the contract method 0x90f1ccba.
//
// Solidity: function slashDelegatorsPercent() view returns(uint256)
func (_FortaStaking *FortaStakingCallerSession) SlashDelegatorsPercent() (*big.Int, error) {
	return _FortaStaking.Contract.SlashDelegatorsPercent(&_FortaStaking.CallOpts)
}

// StakeToActiveShares is a free data retrieval call binding the contract method 0x1a82ef4e.
//
// Solidity: function stakeToActiveShares(uint256 activeSharesId, uint256 amount) view returns(uint256)
func (_FortaStaking *FortaStakingCaller) StakeToActiveShares(opts *bind.CallOpts, activeSharesId *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "stakeToActiveShares", activeSharesId, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeToActiveShares is a free data retrieval call binding the contract method 0x1a82ef4e.
//
// Solidity: function stakeToActiveShares(uint256 activeSharesId, uint256 amount) view returns(uint256)
func (_FortaStaking *FortaStakingSession) StakeToActiveShares(activeSharesId *big.Int, amount *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.StakeToActiveShares(&_FortaStaking.CallOpts, activeSharesId, amount)
}

// StakeToActiveShares is a free data retrieval call binding the contract method 0x1a82ef4e.
//
// Solidity: function stakeToActiveShares(uint256 activeSharesId, uint256 amount) view returns(uint256)
func (_FortaStaking *FortaStakingCallerSession) StakeToActiveShares(activeSharesId *big.Int, amount *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.StakeToActiveShares(&_FortaStaking.CallOpts, activeSharesId, amount)
}

// StakeToInactiveShares is a free data retrieval call binding the contract method 0x62772b14.
//
// Solidity: function stakeToInactiveShares(uint256 inactiveSharesId, uint256 amount) view returns(uint256)
func (_FortaStaking *FortaStakingCaller) StakeToInactiveShares(opts *bind.CallOpts, inactiveSharesId *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "stakeToInactiveShares", inactiveSharesId, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeToInactiveShares is a free data retrieval call binding the contract method 0x62772b14.
//
// Solidity: function stakeToInactiveShares(uint256 inactiveSharesId, uint256 amount) view returns(uint256)
func (_FortaStaking *FortaStakingSession) StakeToInactiveShares(inactiveSharesId *big.Int, amount *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.StakeToInactiveShares(&_FortaStaking.CallOpts, inactiveSharesId, amount)
}

// StakeToInactiveShares is a free data retrieval call binding the contract method 0x62772b14.
//
// Solidity: function stakeToInactiveShares(uint256 inactiveSharesId, uint256 amount) view returns(uint256)
func (_FortaStaking *FortaStakingCallerSession) StakeToInactiveShares(inactiveSharesId *big.Int, amount *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.StakeToInactiveShares(&_FortaStaking.CallOpts, inactiveSharesId, amount)
}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_FortaStaking *FortaStakingCaller) StakedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "stakedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_FortaStaking *FortaStakingSession) StakedToken() (common.Address, error) {
	return _FortaStaking.Contract.StakedToken(&_FortaStaking.CallOpts)
}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_FortaStaking *FortaStakingCallerSession) StakedToken() (common.Address, error) {
	return _FortaStaking.Contract.StakedToken(&_FortaStaking.CallOpts)
}

// SubjectGateway is a free data retrieval call binding the contract method 0x1fd5a956.
//
// Solidity: function subjectGateway() view returns(address)
func (_FortaStaking *FortaStakingCaller) SubjectGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "subjectGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SubjectGateway is a free data retrieval call binding the contract method 0x1fd5a956.
//
// Solidity: function subjectGateway() view returns(address)
func (_FortaStaking *FortaStakingSession) SubjectGateway() (common.Address, error) {
	return _FortaStaking.Contract.SubjectGateway(&_FortaStaking.CallOpts)
}

// SubjectGateway is a free data retrieval call binding the contract method 0x1fd5a956.
//
// Solidity: function subjectGateway() view returns(address)
func (_FortaStaking *FortaStakingCallerSession) SubjectGateway() (common.Address, error) {
	return _FortaStaking.Contract.SubjectGateway(&_FortaStaking.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FortaStaking *FortaStakingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FortaStaking *FortaStakingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FortaStaking.Contract.SupportsInterface(&_FortaStaking.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FortaStaking *FortaStakingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FortaStaking.Contract.SupportsInterface(&_FortaStaking.CallOpts, interfaceId)
}

// TotalActiveStake is a free data retrieval call binding the contract method 0x28f73148.
//
// Solidity: function totalActiveStake() view returns(uint256)
func (_FortaStaking *FortaStakingCaller) TotalActiveStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "totalActiveStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalActiveStake is a free data retrieval call binding the contract method 0x28f73148.
//
// Solidity: function totalActiveStake() view returns(uint256)
func (_FortaStaking *FortaStakingSession) TotalActiveStake() (*big.Int, error) {
	return _FortaStaking.Contract.TotalActiveStake(&_FortaStaking.CallOpts)
}

// TotalActiveStake is a free data retrieval call binding the contract method 0x28f73148.
//
// Solidity: function totalActiveStake() view returns(uint256)
func (_FortaStaking *FortaStakingCallerSession) TotalActiveStake() (*big.Int, error) {
	return _FortaStaking.Contract.TotalActiveStake(&_FortaStaking.CallOpts)
}

// TotalInactiveShares is a free data retrieval call binding the contract method 0x321ebc54.
//
// Solidity: function totalInactiveShares(uint8 subjectType, uint256 subject) view returns(uint256)
func (_FortaStaking *FortaStakingCaller) TotalInactiveShares(opts *bind.CallOpts, subjectType uint8, subject *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "totalInactiveShares", subjectType, subject)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalInactiveShares is a free data retrieval call binding the contract method 0x321ebc54.
//
// Solidity: function totalInactiveShares(uint8 subjectType, uint256 subject) view returns(uint256)
func (_FortaStaking *FortaStakingSession) TotalInactiveShares(subjectType uint8, subject *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.TotalInactiveShares(&_FortaStaking.CallOpts, subjectType, subject)
}

// TotalInactiveShares is a free data retrieval call binding the contract method 0x321ebc54.
//
// Solidity: function totalInactiveShares(uint8 subjectType, uint256 subject) view returns(uint256)
func (_FortaStaking *FortaStakingCallerSession) TotalInactiveShares(subjectType uint8, subject *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.TotalInactiveShares(&_FortaStaking.CallOpts, subjectType, subject)
}

// TotalInactiveStake is a free data retrieval call binding the contract method 0x371f4226.
//
// Solidity: function totalInactiveStake() view returns(uint256)
func (_FortaStaking *FortaStakingCaller) TotalInactiveStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "totalInactiveStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalInactiveStake is a free data retrieval call binding the contract method 0x371f4226.
//
// Solidity: function totalInactiveStake() view returns(uint256)
func (_FortaStaking *FortaStakingSession) TotalInactiveStake() (*big.Int, error) {
	return _FortaStaking.Contract.TotalInactiveStake(&_FortaStaking.CallOpts)
}

// TotalInactiveStake is a free data retrieval call binding the contract method 0x371f4226.
//
// Solidity: function totalInactiveStake() view returns(uint256)
func (_FortaStaking *FortaStakingCallerSession) TotalInactiveStake() (*big.Int, error) {
	return _FortaStaking.Contract.TotalInactiveStake(&_FortaStaking.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x0e10103f.
//
// Solidity: function totalShares(uint8 subjectType, uint256 subject) view returns(uint256)
func (_FortaStaking *FortaStakingCaller) TotalShares(opts *bind.CallOpts, subjectType uint8, subject *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "totalShares", subjectType, subject)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x0e10103f.
//
// Solidity: function totalShares(uint8 subjectType, uint256 subject) view returns(uint256)
func (_FortaStaking *FortaStakingSession) TotalShares(subjectType uint8, subject *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.TotalShares(&_FortaStaking.CallOpts, subjectType, subject)
}

// TotalShares is a free data retrieval call binding the contract method 0x0e10103f.
//
// Solidity: function totalShares(uint8 subjectType, uint256 subject) view returns(uint256)
func (_FortaStaking *FortaStakingCallerSession) TotalShares(subjectType uint8, subject *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.TotalShares(&_FortaStaking.CallOpts, subjectType, subject)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_FortaStaking *FortaStakingCaller) TotalSupply(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "totalSupply", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_FortaStaking *FortaStakingSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.TotalSupply(&_FortaStaking.CallOpts, id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_FortaStaking *FortaStakingCallerSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _FortaStaking.Contract.TotalSupply(&_FortaStaking.CallOpts, id)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_FortaStaking *FortaStakingCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_FortaStaking *FortaStakingSession) Treasury() (common.Address, error) {
	return _FortaStaking.Contract.Treasury(&_FortaStaking.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_FortaStaking *FortaStakingCallerSession) Treasury() (common.Address, error) {
	return _FortaStaking.Contract.Treasury(&_FortaStaking.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_FortaStaking *FortaStakingCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_FortaStaking *FortaStakingSession) Uri(arg0 *big.Int) (string, error) {
	return _FortaStaking.Contract.Uri(&_FortaStaking.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_FortaStaking *FortaStakingCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _FortaStaking.Contract.Uri(&_FortaStaking.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_FortaStaking *FortaStakingCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_FortaStaking *FortaStakingSession) Version() (string, error) {
	return _FortaStaking.Contract.Version(&_FortaStaking.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_FortaStaking *FortaStakingCallerSession) Version() (string, error) {
	return _FortaStaking.Contract.Version(&_FortaStaking.CallOpts)
}

// ConfigureStakeHelpers is a paid mutator transaction binding the contract method 0x3c3e04dd.
//
// Solidity: function configureStakeHelpers(address _subjectGateway, address _allocator) returns()
func (_FortaStaking *FortaStakingTransactor) ConfigureStakeHelpers(opts *bind.TransactOpts, _subjectGateway common.Address, _allocator common.Address) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "configureStakeHelpers", _subjectGateway, _allocator)
}

// ConfigureStakeHelpers is a paid mutator transaction binding the contract method 0x3c3e04dd.
//
// Solidity: function configureStakeHelpers(address _subjectGateway, address _allocator) returns()
func (_FortaStaking *FortaStakingSession) ConfigureStakeHelpers(_subjectGateway common.Address, _allocator common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.ConfigureStakeHelpers(&_FortaStaking.TransactOpts, _subjectGateway, _allocator)
}

// ConfigureStakeHelpers is a paid mutator transaction binding the contract method 0x3c3e04dd.
//
// Solidity: function configureStakeHelpers(address _subjectGateway, address _allocator) returns()
func (_FortaStaking *FortaStakingTransactorSession) ConfigureStakeHelpers(_subjectGateway common.Address, _allocator common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.ConfigureStakeHelpers(&_FortaStaking.TransactOpts, _subjectGateway, _allocator)
}

// Deposit is a paid mutator transaction binding the contract method 0x2cb31144.
//
// Solidity: function deposit(uint8 subjectType, uint256 subject, uint256 stakeValue) returns(uint256)
func (_FortaStaking *FortaStakingTransactor) Deposit(opts *bind.TransactOpts, subjectType uint8, subject *big.Int, stakeValue *big.Int) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "deposit", subjectType, subject, stakeValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x2cb31144.
//
// Solidity: function deposit(uint8 subjectType, uint256 subject, uint256 stakeValue) returns(uint256)
func (_FortaStaking *FortaStakingSession) Deposit(subjectType uint8, subject *big.Int, stakeValue *big.Int) (*types.Transaction, error) {
	return _FortaStaking.Contract.Deposit(&_FortaStaking.TransactOpts, subjectType, subject, stakeValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x2cb31144.
//
// Solidity: function deposit(uint8 subjectType, uint256 subject, uint256 stakeValue) returns(uint256)
func (_FortaStaking *FortaStakingTransactorSession) Deposit(subjectType uint8, subject *big.Int, stakeValue *big.Int) (*types.Transaction, error) {
	return _FortaStaking.Contract.Deposit(&_FortaStaking.TransactOpts, subjectType, subject, stakeValue)
}

// DisableRouter is a paid mutator transaction binding the contract method 0xd858a7e5.
//
// Solidity: function disableRouter() returns()
func (_FortaStaking *FortaStakingTransactor) DisableRouter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "disableRouter")
}

// DisableRouter is a paid mutator transaction binding the contract method 0xd858a7e5.
//
// Solidity: function disableRouter() returns()
func (_FortaStaking *FortaStakingSession) DisableRouter() (*types.Transaction, error) {
	return _FortaStaking.Contract.DisableRouter(&_FortaStaking.TransactOpts)
}

// DisableRouter is a paid mutator transaction binding the contract method 0xd858a7e5.
//
// Solidity: function disableRouter() returns()
func (_FortaStaking *FortaStakingTransactorSession) DisableRouter() (*types.Transaction, error) {
	return _FortaStaking.Contract.DisableRouter(&_FortaStaking.TransactOpts)
}

// Freeze is a paid mutator transaction binding the contract method 0x226cc300.
//
// Solidity: function freeze(uint8 subjectType, uint256 subject, bool frozen) returns()
func (_FortaStaking *FortaStakingTransactor) Freeze(opts *bind.TransactOpts, subjectType uint8, subject *big.Int, frozen bool) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "freeze", subjectType, subject, frozen)
}

// Freeze is a paid mutator transaction binding the contract method 0x226cc300.
//
// Solidity: function freeze(uint8 subjectType, uint256 subject, bool frozen) returns()
func (_FortaStaking *FortaStakingSession) Freeze(subjectType uint8, subject *big.Int, frozen bool) (*types.Transaction, error) {
	return _FortaStaking.Contract.Freeze(&_FortaStaking.TransactOpts, subjectType, subject, frozen)
}

// Freeze is a paid mutator transaction binding the contract method 0x226cc300.
//
// Solidity: function freeze(uint8 subjectType, uint256 subject, bool frozen) returns()
func (_FortaStaking *FortaStakingTransactorSession) Freeze(subjectType uint8, subject *big.Int, frozen bool) (*types.Transaction, error) {
	return _FortaStaking.Contract.Freeze(&_FortaStaking.TransactOpts, subjectType, subject, frozen)
}

// Initialize is a paid mutator transaction binding the contract method 0x4e04275f.
//
// Solidity: function initialize(address __manager, address __stakedToken, uint64 __withdrawalDelay, address __treasury) returns()
func (_FortaStaking *FortaStakingTransactor) Initialize(opts *bind.TransactOpts, __manager common.Address, __stakedToken common.Address, __withdrawalDelay uint64, __treasury common.Address) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "initialize", __manager, __stakedToken, __withdrawalDelay, __treasury)
}

// Initialize is a paid mutator transaction binding the contract method 0x4e04275f.
//
// Solidity: function initialize(address __manager, address __stakedToken, uint64 __withdrawalDelay, address __treasury) returns()
func (_FortaStaking *FortaStakingSession) Initialize(__manager common.Address, __stakedToken common.Address, __withdrawalDelay uint64, __treasury common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.Initialize(&_FortaStaking.TransactOpts, __manager, __stakedToken, __withdrawalDelay, __treasury)
}

// Initialize is a paid mutator transaction binding the contract method 0x4e04275f.
//
// Solidity: function initialize(address __manager, address __stakedToken, uint64 __withdrawalDelay, address __treasury) returns()
func (_FortaStaking *FortaStakingTransactorSession) Initialize(__manager common.Address, __stakedToken common.Address, __withdrawalDelay uint64, __treasury common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.Initialize(&_FortaStaking.TransactOpts, __manager, __stakedToken, __withdrawalDelay, __treasury)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xedea0bac.
//
// Solidity: function initiateWithdrawal(uint8 subjectType, uint256 subject, uint256 sharesValue) returns(uint64)
func (_FortaStaking *FortaStakingTransactor) InitiateWithdrawal(opts *bind.TransactOpts, subjectType uint8, subject *big.Int, sharesValue *big.Int) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "initiateWithdrawal", subjectType, subject, sharesValue)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xedea0bac.
//
// Solidity: function initiateWithdrawal(uint8 subjectType, uint256 subject, uint256 sharesValue) returns(uint64)
func (_FortaStaking *FortaStakingSession) InitiateWithdrawal(subjectType uint8, subject *big.Int, sharesValue *big.Int) (*types.Transaction, error) {
	return _FortaStaking.Contract.InitiateWithdrawal(&_FortaStaking.TransactOpts, subjectType, subject, sharesValue)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xedea0bac.
//
// Solidity: function initiateWithdrawal(uint8 subjectType, uint256 subject, uint256 sharesValue) returns(uint64)
func (_FortaStaking *FortaStakingTransactorSession) InitiateWithdrawal(subjectType uint8, subject *big.Int, sharesValue *big.Int) (*types.Transaction, error) {
	return _FortaStaking.Contract.InitiateWithdrawal(&_FortaStaking.TransactOpts, subjectType, subject, sharesValue)
}

// Migrate is a paid mutator transaction binding the contract method 0x093af6a9.
//
// Solidity: function migrate(uint8 oldSubjectType, uint256 oldSubject, uint8 newSubjectType, uint256 newSubject, address staker) returns()
func (_FortaStaking *FortaStakingTransactor) Migrate(opts *bind.TransactOpts, oldSubjectType uint8, oldSubject *big.Int, newSubjectType uint8, newSubject *big.Int, staker common.Address) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "migrate", oldSubjectType, oldSubject, newSubjectType, newSubject, staker)
}

// Migrate is a paid mutator transaction binding the contract method 0x093af6a9.
//
// Solidity: function migrate(uint8 oldSubjectType, uint256 oldSubject, uint8 newSubjectType, uint256 newSubject, address staker) returns()
func (_FortaStaking *FortaStakingSession) Migrate(oldSubjectType uint8, oldSubject *big.Int, newSubjectType uint8, newSubject *big.Int, staker common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.Migrate(&_FortaStaking.TransactOpts, oldSubjectType, oldSubject, newSubjectType, newSubject, staker)
}

// Migrate is a paid mutator transaction binding the contract method 0x093af6a9.
//
// Solidity: function migrate(uint8 oldSubjectType, uint256 oldSubject, uint8 newSubjectType, uint256 newSubject, address staker) returns()
func (_FortaStaking *FortaStakingTransactorSession) Migrate(oldSubjectType uint8, oldSubject *big.Int, newSubjectType uint8, newSubject *big.Int, staker common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.Migrate(&_FortaStaking.TransactOpts, oldSubjectType, oldSubject, newSubjectType, newSubject, staker)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_FortaStaking *FortaStakingTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_FortaStaking *FortaStakingSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _FortaStaking.Contract.Multicall(&_FortaStaking.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_FortaStaking *FortaStakingTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _FortaStaking.Contract.Multicall(&_FortaStaking.TransactOpts, data)
}

// RelayPermit is a paid mutator transaction binding the contract method 0xc07975ad.
//
// Solidity: function relayPermit(uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_FortaStaking *FortaStakingTransactor) RelayPermit(opts *bind.TransactOpts, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "relayPermit", value, deadline, v, r, s)
}

// RelayPermit is a paid mutator transaction binding the contract method 0xc07975ad.
//
// Solidity: function relayPermit(uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_FortaStaking *FortaStakingSession) RelayPermit(value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _FortaStaking.Contract.RelayPermit(&_FortaStaking.TransactOpts, value, deadline, v, r, s)
}

// RelayPermit is a paid mutator transaction binding the contract method 0xc07975ad.
//
// Solidity: function relayPermit(uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_FortaStaking *FortaStakingTransactorSession) RelayPermit(value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _FortaStaking.Contract.RelayPermit(&_FortaStaking.TransactOpts, value, deadline, v, r, s)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_FortaStaking *FortaStakingTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_FortaStaking *FortaStakingSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _FortaStaking.Contract.SafeBatchTransferFrom(&_FortaStaking.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_FortaStaking *FortaStakingTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _FortaStaking.Contract.SafeBatchTransferFrom(&_FortaStaking.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_FortaStaking *FortaStakingTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_FortaStaking *FortaStakingSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _FortaStaking.Contract.SafeTransferFrom(&_FortaStaking.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_FortaStaking *FortaStakingTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _FortaStaking.Contract.SafeTransferFrom(&_FortaStaking.TransactOpts, from, to, id, amount, data)
}

// SetAccessManager is a paid mutator transaction binding the contract method 0xc9580804.
//
// Solidity: function setAccessManager(address newManager) returns()
func (_FortaStaking *FortaStakingTransactor) SetAccessManager(opts *bind.TransactOpts, newManager common.Address) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "setAccessManager", newManager)
}

// SetAccessManager is a paid mutator transaction binding the contract method 0xc9580804.
//
// Solidity: function setAccessManager(address newManager) returns()
func (_FortaStaking *FortaStakingSession) SetAccessManager(newManager common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.SetAccessManager(&_FortaStaking.TransactOpts, newManager)
}

// SetAccessManager is a paid mutator transaction binding the contract method 0xc9580804.
//
// Solidity: function setAccessManager(address newManager) returns()
func (_FortaStaking *FortaStakingTransactorSession) SetAccessManager(newManager common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.SetAccessManager(&_FortaStaking.TransactOpts, newManager)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FortaStaking *FortaStakingTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FortaStaking *FortaStakingSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _FortaStaking.Contract.SetApprovalForAll(&_FortaStaking.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FortaStaking *FortaStakingTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _FortaStaking.Contract.SetApprovalForAll(&_FortaStaking.TransactOpts, operator, approved)
}

// SetDelay is a paid mutator transaction binding the contract method 0xc1073302.
//
// Solidity: function setDelay(uint64 newDelay) returns()
func (_FortaStaking *FortaStakingTransactor) SetDelay(opts *bind.TransactOpts, newDelay uint64) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "setDelay", newDelay)
}

// SetDelay is a paid mutator transaction binding the contract method 0xc1073302.
//
// Solidity: function setDelay(uint64 newDelay) returns()
func (_FortaStaking *FortaStakingSession) SetDelay(newDelay uint64) (*types.Transaction, error) {
	return _FortaStaking.Contract.SetDelay(&_FortaStaking.TransactOpts, newDelay)
}

// SetDelay is a paid mutator transaction binding the contract method 0xc1073302.
//
// Solidity: function setDelay(uint64 newDelay) returns()
func (_FortaStaking *FortaStakingTransactorSession) SetDelay(newDelay uint64) (*types.Transaction, error) {
	return _FortaStaking.Contract.SetDelay(&_FortaStaking.TransactOpts, newDelay)
}

// SetName is a paid mutator transaction binding the contract method 0x3121db1c.
//
// Solidity: function setName(address ensRegistry, string ensName) returns()
func (_FortaStaking *FortaStakingTransactor) SetName(opts *bind.TransactOpts, ensRegistry common.Address, ensName string) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "setName", ensRegistry, ensName)
}

// SetName is a paid mutator transaction binding the contract method 0x3121db1c.
//
// Solidity: function setName(address ensRegistry, string ensName) returns()
func (_FortaStaking *FortaStakingSession) SetName(ensRegistry common.Address, ensName string) (*types.Transaction, error) {
	return _FortaStaking.Contract.SetName(&_FortaStaking.TransactOpts, ensRegistry, ensName)
}

// SetName is a paid mutator transaction binding the contract method 0x3121db1c.
//
// Solidity: function setName(address ensRegistry, string ensName) returns()
func (_FortaStaking *FortaStakingTransactorSession) SetName(ensRegistry common.Address, ensName string) (*types.Transaction, error) {
	return _FortaStaking.Contract.SetName(&_FortaStaking.TransactOpts, ensRegistry, ensName)
}

// SetReentrancyGuard is a paid mutator transaction binding the contract method 0xedac6db8.
//
// Solidity: function setReentrancyGuard() returns()
func (_FortaStaking *FortaStakingTransactor) SetReentrancyGuard(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "setReentrancyGuard")
}

// SetReentrancyGuard is a paid mutator transaction binding the contract method 0xedac6db8.
//
// Solidity: function setReentrancyGuard() returns()
func (_FortaStaking *FortaStakingSession) SetReentrancyGuard() (*types.Transaction, error) {
	return _FortaStaking.Contract.SetReentrancyGuard(&_FortaStaking.TransactOpts)
}

// SetReentrancyGuard is a paid mutator transaction binding the contract method 0xedac6db8.
//
// Solidity: function setReentrancyGuard() returns()
func (_FortaStaking *FortaStakingTransactorSession) SetReentrancyGuard() (*types.Transaction, error) {
	return _FortaStaking.Contract.SetReentrancyGuard(&_FortaStaking.TransactOpts)
}

// SetSlashDelegatorsPercent is a paid mutator transaction binding the contract method 0xb4d60fb3.
//
// Solidity: function setSlashDelegatorsPercent(uint256 percent) returns()
func (_FortaStaking *FortaStakingTransactor) SetSlashDelegatorsPercent(opts *bind.TransactOpts, percent *big.Int) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "setSlashDelegatorsPercent", percent)
}

// SetSlashDelegatorsPercent is a paid mutator transaction binding the contract method 0xb4d60fb3.
//
// Solidity: function setSlashDelegatorsPercent(uint256 percent) returns()
func (_FortaStaking *FortaStakingSession) SetSlashDelegatorsPercent(percent *big.Int) (*types.Transaction, error) {
	return _FortaStaking.Contract.SetSlashDelegatorsPercent(&_FortaStaking.TransactOpts, percent)
}

// SetSlashDelegatorsPercent is a paid mutator transaction binding the contract method 0xb4d60fb3.
//
// Solidity: function setSlashDelegatorsPercent(uint256 percent) returns()
func (_FortaStaking *FortaStakingTransactorSession) SetSlashDelegatorsPercent(percent *big.Int) (*types.Transaction, error) {
	return _FortaStaking.Contract.SetSlashDelegatorsPercent(&_FortaStaking.TransactOpts, percent)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_FortaStaking *FortaStakingTransactor) SetTreasury(opts *bind.TransactOpts, newTreasury common.Address) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "setTreasury", newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_FortaStaking *FortaStakingSession) SetTreasury(newTreasury common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.SetTreasury(&_FortaStaking.TransactOpts, newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_FortaStaking *FortaStakingTransactorSession) SetTreasury(newTreasury common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.SetTreasury(&_FortaStaking.TransactOpts, newTreasury)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newUri) returns()
func (_FortaStaking *FortaStakingTransactor) SetURI(opts *bind.TransactOpts, newUri string) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "setURI", newUri)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newUri) returns()
func (_FortaStaking *FortaStakingSession) SetURI(newUri string) (*types.Transaction, error) {
	return _FortaStaking.Contract.SetURI(&_FortaStaking.TransactOpts, newUri)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newUri) returns()
func (_FortaStaking *FortaStakingTransactorSession) SetURI(newUri string) (*types.Transaction, error) {
	return _FortaStaking.Contract.SetURI(&_FortaStaking.TransactOpts, newUri)
}

// Slash is a paid mutator transaction binding the contract method 0xf8058b06.
//
// Solidity: function slash(uint8 subjectType, uint256 subject, uint256 stakeValue, address proposer, uint256 proposerPercent) returns(uint256)
func (_FortaStaking *FortaStakingTransactor) Slash(opts *bind.TransactOpts, subjectType uint8, subject *big.Int, stakeValue *big.Int, proposer common.Address, proposerPercent *big.Int) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "slash", subjectType, subject, stakeValue, proposer, proposerPercent)
}

// Slash is a paid mutator transaction binding the contract method 0xf8058b06.
//
// Solidity: function slash(uint8 subjectType, uint256 subject, uint256 stakeValue, address proposer, uint256 proposerPercent) returns(uint256)
func (_FortaStaking *FortaStakingSession) Slash(subjectType uint8, subject *big.Int, stakeValue *big.Int, proposer common.Address, proposerPercent *big.Int) (*types.Transaction, error) {
	return _FortaStaking.Contract.Slash(&_FortaStaking.TransactOpts, subjectType, subject, stakeValue, proposer, proposerPercent)
}

// Slash is a paid mutator transaction binding the contract method 0xf8058b06.
//
// Solidity: function slash(uint8 subjectType, uint256 subject, uint256 stakeValue, address proposer, uint256 proposerPercent) returns(uint256)
func (_FortaStaking *FortaStakingTransactorSession) Slash(subjectType uint8, subject *big.Int, stakeValue *big.Int, proposer common.Address, proposerPercent *big.Int) (*types.Transaction, error) {
	return _FortaStaking.Contract.Slash(&_FortaStaking.TransactOpts, subjectType, subject, stakeValue, proposer, proposerPercent)
}

// Sweep is a paid mutator transaction binding the contract method 0xb8dc491b.
//
// Solidity: function sweep(address token, address recipient) returns(uint256)
func (_FortaStaking *FortaStakingTransactor) Sweep(opts *bind.TransactOpts, token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "sweep", token, recipient)
}

// Sweep is a paid mutator transaction binding the contract method 0xb8dc491b.
//
// Solidity: function sweep(address token, address recipient) returns(uint256)
func (_FortaStaking *FortaStakingSession) Sweep(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.Sweep(&_FortaStaking.TransactOpts, token, recipient)
}

// Sweep is a paid mutator transaction binding the contract method 0xb8dc491b.
//
// Solidity: function sweep(address token, address recipient) returns(uint256)
func (_FortaStaking *FortaStakingTransactorSession) Sweep(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.Sweep(&_FortaStaking.TransactOpts, token, recipient)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_FortaStaking *FortaStakingTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_FortaStaking *FortaStakingSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.UpgradeTo(&_FortaStaking.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_FortaStaking *FortaStakingTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.UpgradeTo(&_FortaStaking.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FortaStaking *FortaStakingTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FortaStaking *FortaStakingSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FortaStaking.Contract.UpgradeToAndCall(&_FortaStaking.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FortaStaking *FortaStakingTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FortaStaking.Contract.UpgradeToAndCall(&_FortaStaking.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3f489914.
//
// Solidity: function withdraw(uint8 subjectType, uint256 subject) returns(uint256)
func (_FortaStaking *FortaStakingTransactor) Withdraw(opts *bind.TransactOpts, subjectType uint8, subject *big.Int) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "withdraw", subjectType, subject)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3f489914.
//
// Solidity: function withdraw(uint8 subjectType, uint256 subject) returns(uint256)
func (_FortaStaking *FortaStakingSession) Withdraw(subjectType uint8, subject *big.Int) (*types.Transaction, error) {
	return _FortaStaking.Contract.Withdraw(&_FortaStaking.TransactOpts, subjectType, subject)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3f489914.
//
// Solidity: function withdraw(uint8 subjectType, uint256 subject) returns(uint256)
func (_FortaStaking *FortaStakingTransactorSession) Withdraw(subjectType uint8, subject *big.Int) (*types.Transaction, error) {
	return _FortaStaking.Contract.Withdraw(&_FortaStaking.TransactOpts, subjectType, subject)
}

// FortaStakingAccessManagerUpdatedIterator is returned from FilterAccessManagerUpdated and is used to iterate over the raw logs and unpacked data for AccessManagerUpdated events raised by the FortaStaking contract.
type FortaStakingAccessManagerUpdatedIterator struct {
	Event *FortaStakingAccessManagerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingAccessManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingAccessManagerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingAccessManagerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingAccessManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingAccessManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingAccessManagerUpdated represents a AccessManagerUpdated event raised by the FortaStaking contract.
type FortaStakingAccessManagerUpdated struct {
	NewAddressManager common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAccessManagerUpdated is a free log retrieval operation binding the contract event 0xa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c.
//
// Solidity: event AccessManagerUpdated(address indexed newAddressManager)
func (_FortaStaking *FortaStakingFilterer) FilterAccessManagerUpdated(opts *bind.FilterOpts, newAddressManager []common.Address) (*FortaStakingAccessManagerUpdatedIterator, error) {

	var newAddressManagerRule []interface{}
	for _, newAddressManagerItem := range newAddressManager {
		newAddressManagerRule = append(newAddressManagerRule, newAddressManagerItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "AccessManagerUpdated", newAddressManagerRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingAccessManagerUpdatedIterator{contract: _FortaStaking.contract, event: "AccessManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchAccessManagerUpdated is a free log subscription operation binding the contract event 0xa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c.
//
// Solidity: event AccessManagerUpdated(address indexed newAddressManager)
func (_FortaStaking *FortaStakingFilterer) WatchAccessManagerUpdated(opts *bind.WatchOpts, sink chan<- *FortaStakingAccessManagerUpdated, newAddressManager []common.Address) (event.Subscription, error) {

	var newAddressManagerRule []interface{}
	for _, newAddressManagerItem := range newAddressManager {
		newAddressManagerRule = append(newAddressManagerRule, newAddressManagerItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "AccessManagerUpdated", newAddressManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingAccessManagerUpdated)
				if err := _FortaStaking.contract.UnpackLog(event, "AccessManagerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccessManagerUpdated is a log parse operation binding the contract event 0xa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c.
//
// Solidity: event AccessManagerUpdated(address indexed newAddressManager)
func (_FortaStaking *FortaStakingFilterer) ParseAccessManagerUpdated(log types.Log) (*FortaStakingAccessManagerUpdated, error) {
	event := new(FortaStakingAccessManagerUpdated)
	if err := _FortaStaking.contract.UnpackLog(event, "AccessManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the FortaStaking contract.
type FortaStakingAdminChangedIterator struct {
	Event *FortaStakingAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingAdminChanged represents a AdminChanged event raised by the FortaStaking contract.
type FortaStakingAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_FortaStaking *FortaStakingFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*FortaStakingAdminChangedIterator, error) {

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &FortaStakingAdminChangedIterator{contract: _FortaStaking.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_FortaStaking *FortaStakingFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *FortaStakingAdminChanged) (event.Subscription, error) {

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingAdminChanged)
				if err := _FortaStaking.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_FortaStaking *FortaStakingFilterer) ParseAdminChanged(log types.Log) (*FortaStakingAdminChanged, error) {
	event := new(FortaStakingAdminChanged)
	if err := _FortaStaking.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the FortaStaking contract.
type FortaStakingApprovalForAllIterator struct {
	Event *FortaStakingApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingApprovalForAll represents a ApprovalForAll event raised by the FortaStaking contract.
type FortaStakingApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_FortaStaking *FortaStakingFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*FortaStakingApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingApprovalForAllIterator{contract: _FortaStaking.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_FortaStaking *FortaStakingFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *FortaStakingApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingApprovalForAll)
				if err := _FortaStaking.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_FortaStaking *FortaStakingFilterer) ParseApprovalForAll(log types.Log) (*FortaStakingApprovalForAll, error) {
	event := new(FortaStakingApprovalForAll)
	if err := _FortaStaking.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the FortaStaking contract.
type FortaStakingBeaconUpgradedIterator struct {
	Event *FortaStakingBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingBeaconUpgraded represents a BeaconUpgraded event raised by the FortaStaking contract.
type FortaStakingBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_FortaStaking *FortaStakingFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*FortaStakingBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingBeaconUpgradedIterator{contract: _FortaStaking.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_FortaStaking *FortaStakingFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *FortaStakingBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingBeaconUpgraded)
				if err := _FortaStaking.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_FortaStaking *FortaStakingFilterer) ParseBeaconUpgraded(log types.Log) (*FortaStakingBeaconUpgraded, error) {
	event := new(FortaStakingBeaconUpgraded)
	if err := _FortaStaking.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingDelaySetIterator is returned from FilterDelaySet and is used to iterate over the raw logs and unpacked data for DelaySet events raised by the FortaStaking contract.
type FortaStakingDelaySetIterator struct {
	Event *FortaStakingDelaySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingDelaySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingDelaySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingDelaySet represents a DelaySet event raised by the FortaStaking contract.
type FortaStakingDelaySet struct {
	NewWithdrawalDelay *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDelaySet is a free log retrieval operation binding the contract event 0x63e09f16584208fba1fc7ff64c62b00f07bec177c0d97ca6689891b1e77a35c7.
//
// Solidity: event DelaySet(uint256 newWithdrawalDelay)
func (_FortaStaking *FortaStakingFilterer) FilterDelaySet(opts *bind.FilterOpts) (*FortaStakingDelaySetIterator, error) {

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "DelaySet")
	if err != nil {
		return nil, err
	}
	return &FortaStakingDelaySetIterator{contract: _FortaStaking.contract, event: "DelaySet", logs: logs, sub: sub}, nil
}

// WatchDelaySet is a free log subscription operation binding the contract event 0x63e09f16584208fba1fc7ff64c62b00f07bec177c0d97ca6689891b1e77a35c7.
//
// Solidity: event DelaySet(uint256 newWithdrawalDelay)
func (_FortaStaking *FortaStakingFilterer) WatchDelaySet(opts *bind.WatchOpts, sink chan<- *FortaStakingDelaySet) (event.Subscription, error) {

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "DelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingDelaySet)
				if err := _FortaStaking.contract.UnpackLog(event, "DelaySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelaySet is a log parse operation binding the contract event 0x63e09f16584208fba1fc7ff64c62b00f07bec177c0d97ca6689891b1e77a35c7.
//
// Solidity: event DelaySet(uint256 newWithdrawalDelay)
func (_FortaStaking *FortaStakingFilterer) ParseDelaySet(log types.Log) (*FortaStakingDelaySet, error) {
	event := new(FortaStakingDelaySet)
	if err := _FortaStaking.contract.UnpackLog(event, "DelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingFrozeIterator is returned from FilterFroze and is used to iterate over the raw logs and unpacked data for Froze events raised by the FortaStaking contract.
type FortaStakingFrozeIterator struct {
	Event *FortaStakingFroze // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingFrozeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingFroze)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingFroze)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingFrozeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingFrozeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingFroze represents a Froze event raised by the FortaStaking contract.
type FortaStakingFroze struct {
	SubjectType uint8
	Subject     *big.Int
	By          common.Address
	IsFrozen    bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFroze is a free log retrieval operation binding the contract event 0xd520b4ee12f45aacea6f06fd4831c1410e0d3f874a80f64424715327d7f705ca.
//
// Solidity: event Froze(uint8 indexed subjectType, uint256 indexed subject, address indexed by, bool isFrozen)
func (_FortaStaking *FortaStakingFilterer) FilterFroze(opts *bind.FilterOpts, subjectType []uint8, subject []*big.Int, by []common.Address) (*FortaStakingFrozeIterator, error) {

	var subjectTypeRule []interface{}
	for _, subjectTypeItem := range subjectType {
		subjectTypeRule = append(subjectTypeRule, subjectTypeItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "Froze", subjectTypeRule, subjectRule, byRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingFrozeIterator{contract: _FortaStaking.contract, event: "Froze", logs: logs, sub: sub}, nil
}

// WatchFroze is a free log subscription operation binding the contract event 0xd520b4ee12f45aacea6f06fd4831c1410e0d3f874a80f64424715327d7f705ca.
//
// Solidity: event Froze(uint8 indexed subjectType, uint256 indexed subject, address indexed by, bool isFrozen)
func (_FortaStaking *FortaStakingFilterer) WatchFroze(opts *bind.WatchOpts, sink chan<- *FortaStakingFroze, subjectType []uint8, subject []*big.Int, by []common.Address) (event.Subscription, error) {

	var subjectTypeRule []interface{}
	for _, subjectTypeItem := range subjectType {
		subjectTypeRule = append(subjectTypeRule, subjectTypeItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "Froze", subjectTypeRule, subjectRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingFroze)
				if err := _FortaStaking.contract.UnpackLog(event, "Froze", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFroze is a log parse operation binding the contract event 0xd520b4ee12f45aacea6f06fd4831c1410e0d3f874a80f64424715327d7f705ca.
//
// Solidity: event Froze(uint8 indexed subjectType, uint256 indexed subject, address indexed by, bool isFrozen)
func (_FortaStaking *FortaStakingFilterer) ParseFroze(log types.Log) (*FortaStakingFroze, error) {
	event := new(FortaStakingFroze)
	if err := _FortaStaking.contract.UnpackLog(event, "Froze", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the FortaStaking contract.
type FortaStakingInitializedIterator struct {
	Event *FortaStakingInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingInitialized represents a Initialized event raised by the FortaStaking contract.
type FortaStakingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_FortaStaking *FortaStakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*FortaStakingInitializedIterator, error) {

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FortaStakingInitializedIterator{contract: _FortaStaking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_FortaStaking *FortaStakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FortaStakingInitialized) (event.Subscription, error) {

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingInitialized)
				if err := _FortaStaking.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_FortaStaking *FortaStakingFilterer) ParseInitialized(log types.Log) (*FortaStakingInitialized, error) {
	event := new(FortaStakingInitialized)
	if err := _FortaStaking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingMaxStakeReachedIterator is returned from FilterMaxStakeReached and is used to iterate over the raw logs and unpacked data for MaxStakeReached events raised by the FortaStaking contract.
type FortaStakingMaxStakeReachedIterator struct {
	Event *FortaStakingMaxStakeReached // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingMaxStakeReachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingMaxStakeReached)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingMaxStakeReached)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingMaxStakeReachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingMaxStakeReachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingMaxStakeReached represents a MaxStakeReached event raised by the FortaStaking contract.
type FortaStakingMaxStakeReached struct {
	SubjectType uint8
	Subject     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMaxStakeReached is a free log retrieval operation binding the contract event 0x8db486b45abb5d250b92d9eeea290a901f784921d6d28705cfaf2e9fa025d563.
//
// Solidity: event MaxStakeReached(uint8 indexed subjectType, uint256 indexed subject)
func (_FortaStaking *FortaStakingFilterer) FilterMaxStakeReached(opts *bind.FilterOpts, subjectType []uint8, subject []*big.Int) (*FortaStakingMaxStakeReachedIterator, error) {

	var subjectTypeRule []interface{}
	for _, subjectTypeItem := range subjectType {
		subjectTypeRule = append(subjectTypeRule, subjectTypeItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "MaxStakeReached", subjectTypeRule, subjectRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingMaxStakeReachedIterator{contract: _FortaStaking.contract, event: "MaxStakeReached", logs: logs, sub: sub}, nil
}

// WatchMaxStakeReached is a free log subscription operation binding the contract event 0x8db486b45abb5d250b92d9eeea290a901f784921d6d28705cfaf2e9fa025d563.
//
// Solidity: event MaxStakeReached(uint8 indexed subjectType, uint256 indexed subject)
func (_FortaStaking *FortaStakingFilterer) WatchMaxStakeReached(opts *bind.WatchOpts, sink chan<- *FortaStakingMaxStakeReached, subjectType []uint8, subject []*big.Int) (event.Subscription, error) {

	var subjectTypeRule []interface{}
	for _, subjectTypeItem := range subjectType {
		subjectTypeRule = append(subjectTypeRule, subjectTypeItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "MaxStakeReached", subjectTypeRule, subjectRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingMaxStakeReached)
				if err := _FortaStaking.contract.UnpackLog(event, "MaxStakeReached", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMaxStakeReached is a log parse operation binding the contract event 0x8db486b45abb5d250b92d9eeea290a901f784921d6d28705cfaf2e9fa025d563.
//
// Solidity: event MaxStakeReached(uint8 indexed subjectType, uint256 indexed subject)
func (_FortaStaking *FortaStakingFilterer) ParseMaxStakeReached(log types.Log) (*FortaStakingMaxStakeReached, error) {
	event := new(FortaStakingMaxStakeReached)
	if err := _FortaStaking.contract.UnpackLog(event, "MaxStakeReached", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingRouterUpdatedIterator is returned from FilterRouterUpdated and is used to iterate over the raw logs and unpacked data for RouterUpdated events raised by the FortaStaking contract.
type FortaStakingRouterUpdatedIterator struct {
	Event *FortaStakingRouterUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingRouterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingRouterUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingRouterUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingRouterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingRouterUpdated represents a RouterUpdated event raised by the FortaStaking contract.
type FortaStakingRouterUpdated struct {
	Router common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRouterUpdated is a free log retrieval operation binding the contract event 0x7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc80.
//
// Solidity: event RouterUpdated(address indexed router)
func (_FortaStaking *FortaStakingFilterer) FilterRouterUpdated(opts *bind.FilterOpts, router []common.Address) (*FortaStakingRouterUpdatedIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "RouterUpdated", routerRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingRouterUpdatedIterator{contract: _FortaStaking.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

// WatchRouterUpdated is a free log subscription operation binding the contract event 0x7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc80.
//
// Solidity: event RouterUpdated(address indexed router)
func (_FortaStaking *FortaStakingFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *FortaStakingRouterUpdated, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "RouterUpdated", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingRouterUpdated)
				if err := _FortaStaking.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRouterUpdated is a log parse operation binding the contract event 0x7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc80.
//
// Solidity: event RouterUpdated(address indexed router)
func (_FortaStaking *FortaStakingFilterer) ParseRouterUpdated(log types.Log) (*FortaStakingRouterUpdated, error) {
	event := new(FortaStakingRouterUpdated)
	if err := _FortaStaking.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingSlashDelegatorsPercentSetIterator is returned from FilterSlashDelegatorsPercentSet and is used to iterate over the raw logs and unpacked data for SlashDelegatorsPercentSet events raised by the FortaStaking contract.
type FortaStakingSlashDelegatorsPercentSetIterator struct {
	Event *FortaStakingSlashDelegatorsPercentSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingSlashDelegatorsPercentSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingSlashDelegatorsPercentSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingSlashDelegatorsPercentSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingSlashDelegatorsPercentSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingSlashDelegatorsPercentSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingSlashDelegatorsPercentSet represents a SlashDelegatorsPercentSet event raised by the FortaStaking contract.
type FortaStakingSlashDelegatorsPercentSet struct {
	Percent *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSlashDelegatorsPercentSet is a free log retrieval operation binding the contract event 0x76096f684901589fef72a8b0a0321bf061b8597a791db667ab58e7eb922c639c.
//
// Solidity: event SlashDelegatorsPercentSet(uint256 percent)
func (_FortaStaking *FortaStakingFilterer) FilterSlashDelegatorsPercentSet(opts *bind.FilterOpts) (*FortaStakingSlashDelegatorsPercentSetIterator, error) {

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "SlashDelegatorsPercentSet")
	if err != nil {
		return nil, err
	}
	return &FortaStakingSlashDelegatorsPercentSetIterator{contract: _FortaStaking.contract, event: "SlashDelegatorsPercentSet", logs: logs, sub: sub}, nil
}

// WatchSlashDelegatorsPercentSet is a free log subscription operation binding the contract event 0x76096f684901589fef72a8b0a0321bf061b8597a791db667ab58e7eb922c639c.
//
// Solidity: event SlashDelegatorsPercentSet(uint256 percent)
func (_FortaStaking *FortaStakingFilterer) WatchSlashDelegatorsPercentSet(opts *bind.WatchOpts, sink chan<- *FortaStakingSlashDelegatorsPercentSet) (event.Subscription, error) {

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "SlashDelegatorsPercentSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingSlashDelegatorsPercentSet)
				if err := _FortaStaking.contract.UnpackLog(event, "SlashDelegatorsPercentSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlashDelegatorsPercentSet is a log parse operation binding the contract event 0x76096f684901589fef72a8b0a0321bf061b8597a791db667ab58e7eb922c639c.
//
// Solidity: event SlashDelegatorsPercentSet(uint256 percent)
func (_FortaStaking *FortaStakingFilterer) ParseSlashDelegatorsPercentSet(log types.Log) (*FortaStakingSlashDelegatorsPercentSet, error) {
	event := new(FortaStakingSlashDelegatorsPercentSet)
	if err := _FortaStaking.contract.UnpackLog(event, "SlashDelegatorsPercentSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the FortaStaking contract.
type FortaStakingSlashedIterator struct {
	Event *FortaStakingSlashed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingSlashed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingSlashed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingSlashed represents a Slashed event raised by the FortaStaking contract.
type FortaStakingSlashed struct {
	SubjectType uint8
	Subject     *big.Int
	By          common.Address
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0x59d631535a5e48c1231a728e536ce39dba8d6c7f49905ec570e3db296430e02e.
//
// Solidity: event Slashed(uint8 indexed subjectType, uint256 indexed subject, address indexed by, uint256 value)
func (_FortaStaking *FortaStakingFilterer) FilterSlashed(opts *bind.FilterOpts, subjectType []uint8, subject []*big.Int, by []common.Address) (*FortaStakingSlashedIterator, error) {

	var subjectTypeRule []interface{}
	for _, subjectTypeItem := range subjectType {
		subjectTypeRule = append(subjectTypeRule, subjectTypeItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "Slashed", subjectTypeRule, subjectRule, byRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingSlashedIterator{contract: _FortaStaking.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0x59d631535a5e48c1231a728e536ce39dba8d6c7f49905ec570e3db296430e02e.
//
// Solidity: event Slashed(uint8 indexed subjectType, uint256 indexed subject, address indexed by, uint256 value)
func (_FortaStaking *FortaStakingFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *FortaStakingSlashed, subjectType []uint8, subject []*big.Int, by []common.Address) (event.Subscription, error) {

	var subjectTypeRule []interface{}
	for _, subjectTypeItem := range subjectType {
		subjectTypeRule = append(subjectTypeRule, subjectTypeItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "Slashed", subjectTypeRule, subjectRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingSlashed)
				if err := _FortaStaking.contract.UnpackLog(event, "Slashed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlashed is a log parse operation binding the contract event 0x59d631535a5e48c1231a728e536ce39dba8d6c7f49905ec570e3db296430e02e.
//
// Solidity: event Slashed(uint8 indexed subjectType, uint256 indexed subject, address indexed by, uint256 value)
func (_FortaStaking *FortaStakingFilterer) ParseSlashed(log types.Log) (*FortaStakingSlashed, error) {
	event := new(FortaStakingSlashed)
	if err := _FortaStaking.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingSlashedShareSentIterator is returned from FilterSlashedShareSent and is used to iterate over the raw logs and unpacked data for SlashedShareSent events raised by the FortaStaking contract.
type FortaStakingSlashedShareSentIterator struct {
	Event *FortaStakingSlashedShareSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingSlashedShareSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingSlashedShareSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingSlashedShareSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingSlashedShareSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingSlashedShareSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingSlashedShareSent represents a SlashedShareSent event raised by the FortaStaking contract.
type FortaStakingSlashedShareSent struct {
	SubjectType uint8
	Subject     *big.Int
	By          common.Address
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSlashedShareSent is a free log retrieval operation binding the contract event 0xbad5bf3ab3814a2220a4737f205e42060a9b3b66422e774059048cec2f6ed0ac.
//
// Solidity: event SlashedShareSent(uint8 indexed subjectType, uint256 indexed subject, address indexed by, uint256 value)
func (_FortaStaking *FortaStakingFilterer) FilterSlashedShareSent(opts *bind.FilterOpts, subjectType []uint8, subject []*big.Int, by []common.Address) (*FortaStakingSlashedShareSentIterator, error) {

	var subjectTypeRule []interface{}
	for _, subjectTypeItem := range subjectType {
		subjectTypeRule = append(subjectTypeRule, subjectTypeItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "SlashedShareSent", subjectTypeRule, subjectRule, byRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingSlashedShareSentIterator{contract: _FortaStaking.contract, event: "SlashedShareSent", logs: logs, sub: sub}, nil
}

// WatchSlashedShareSent is a free log subscription operation binding the contract event 0xbad5bf3ab3814a2220a4737f205e42060a9b3b66422e774059048cec2f6ed0ac.
//
// Solidity: event SlashedShareSent(uint8 indexed subjectType, uint256 indexed subject, address indexed by, uint256 value)
func (_FortaStaking *FortaStakingFilterer) WatchSlashedShareSent(opts *bind.WatchOpts, sink chan<- *FortaStakingSlashedShareSent, subjectType []uint8, subject []*big.Int, by []common.Address) (event.Subscription, error) {

	var subjectTypeRule []interface{}
	for _, subjectTypeItem := range subjectType {
		subjectTypeRule = append(subjectTypeRule, subjectTypeItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "SlashedShareSent", subjectTypeRule, subjectRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingSlashedShareSent)
				if err := _FortaStaking.contract.UnpackLog(event, "SlashedShareSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlashedShareSent is a log parse operation binding the contract event 0xbad5bf3ab3814a2220a4737f205e42060a9b3b66422e774059048cec2f6ed0ac.
//
// Solidity: event SlashedShareSent(uint8 indexed subjectType, uint256 indexed subject, address indexed by, uint256 value)
func (_FortaStaking *FortaStakingFilterer) ParseSlashedShareSent(log types.Log) (*FortaStakingSlashedShareSent, error) {
	event := new(FortaStakingSlashedShareSent)
	if err := _FortaStaking.contract.UnpackLog(event, "SlashedShareSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingStakeDepositedIterator is returned from FilterStakeDeposited and is used to iterate over the raw logs and unpacked data for StakeDeposited events raised by the FortaStaking contract.
type FortaStakingStakeDepositedIterator struct {
	Event *FortaStakingStakeDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingStakeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingStakeDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingStakeDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingStakeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingStakeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingStakeDeposited represents a StakeDeposited event raised by the FortaStaking contract.
type FortaStakingStakeDeposited struct {
	SubjectType uint8
	Subject     *big.Int
	Account     common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakeDeposited is a free log retrieval operation binding the contract event 0x3e5900eae50f376487f0d2272f8f857884c7adde3b5be2cd8c7b1e7bdfc47e14.
//
// Solidity: event StakeDeposited(uint8 indexed subjectType, uint256 indexed subject, address indexed account, uint256 amount)
func (_FortaStaking *FortaStakingFilterer) FilterStakeDeposited(opts *bind.FilterOpts, subjectType []uint8, subject []*big.Int, account []common.Address) (*FortaStakingStakeDepositedIterator, error) {

	var subjectTypeRule []interface{}
	for _, subjectTypeItem := range subjectType {
		subjectTypeRule = append(subjectTypeRule, subjectTypeItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "StakeDeposited", subjectTypeRule, subjectRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingStakeDepositedIterator{contract: _FortaStaking.contract, event: "StakeDeposited", logs: logs, sub: sub}, nil
}

// WatchStakeDeposited is a free log subscription operation binding the contract event 0x3e5900eae50f376487f0d2272f8f857884c7adde3b5be2cd8c7b1e7bdfc47e14.
//
// Solidity: event StakeDeposited(uint8 indexed subjectType, uint256 indexed subject, address indexed account, uint256 amount)
func (_FortaStaking *FortaStakingFilterer) WatchStakeDeposited(opts *bind.WatchOpts, sink chan<- *FortaStakingStakeDeposited, subjectType []uint8, subject []*big.Int, account []common.Address) (event.Subscription, error) {

	var subjectTypeRule []interface{}
	for _, subjectTypeItem := range subjectType {
		subjectTypeRule = append(subjectTypeRule, subjectTypeItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "StakeDeposited", subjectTypeRule, subjectRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingStakeDeposited)
				if err := _FortaStaking.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeDeposited is a log parse operation binding the contract event 0x3e5900eae50f376487f0d2272f8f857884c7adde3b5be2cd8c7b1e7bdfc47e14.
//
// Solidity: event StakeDeposited(uint8 indexed subjectType, uint256 indexed subject, address indexed account, uint256 amount)
func (_FortaStaking *FortaStakingFilterer) ParseStakeDeposited(log types.Log) (*FortaStakingStakeDeposited, error) {
	event := new(FortaStakingStakeDeposited)
	if err := _FortaStaking.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingStakeHelpersConfiguredIterator is returned from FilterStakeHelpersConfigured and is used to iterate over the raw logs and unpacked data for StakeHelpersConfigured events raised by the FortaStaking contract.
type FortaStakingStakeHelpersConfiguredIterator struct {
	Event *FortaStakingStakeHelpersConfigured // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingStakeHelpersConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingStakeHelpersConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingStakeHelpersConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingStakeHelpersConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingStakeHelpersConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingStakeHelpersConfigured represents a StakeHelpersConfigured event raised by the FortaStaking contract.
type FortaStakingStakeHelpersConfigured struct {
	SubjectGateway common.Address
	Allocator      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStakeHelpersConfigured is a free log retrieval operation binding the contract event 0xe727095b96a07a7ea4675d80358e4c6d0a7b287b367e4df2e1ee384e7d59ecc1.
//
// Solidity: event StakeHelpersConfigured(address indexed subjectGateway, address indexed allocator)
func (_FortaStaking *FortaStakingFilterer) FilterStakeHelpersConfigured(opts *bind.FilterOpts, subjectGateway []common.Address, allocator []common.Address) (*FortaStakingStakeHelpersConfiguredIterator, error) {

	var subjectGatewayRule []interface{}
	for _, subjectGatewayItem := range subjectGateway {
		subjectGatewayRule = append(subjectGatewayRule, subjectGatewayItem)
	}
	var allocatorRule []interface{}
	for _, allocatorItem := range allocator {
		allocatorRule = append(allocatorRule, allocatorItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "StakeHelpersConfigured", subjectGatewayRule, allocatorRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingStakeHelpersConfiguredIterator{contract: _FortaStaking.contract, event: "StakeHelpersConfigured", logs: logs, sub: sub}, nil
}

// WatchStakeHelpersConfigured is a free log subscription operation binding the contract event 0xe727095b96a07a7ea4675d80358e4c6d0a7b287b367e4df2e1ee384e7d59ecc1.
//
// Solidity: event StakeHelpersConfigured(address indexed subjectGateway, address indexed allocator)
func (_FortaStaking *FortaStakingFilterer) WatchStakeHelpersConfigured(opts *bind.WatchOpts, sink chan<- *FortaStakingStakeHelpersConfigured, subjectGateway []common.Address, allocator []common.Address) (event.Subscription, error) {

	var subjectGatewayRule []interface{}
	for _, subjectGatewayItem := range subjectGateway {
		subjectGatewayRule = append(subjectGatewayRule, subjectGatewayItem)
	}
	var allocatorRule []interface{}
	for _, allocatorItem := range allocator {
		allocatorRule = append(allocatorRule, allocatorItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "StakeHelpersConfigured", subjectGatewayRule, allocatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingStakeHelpersConfigured)
				if err := _FortaStaking.contract.UnpackLog(event, "StakeHelpersConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeHelpersConfigured is a log parse operation binding the contract event 0xe727095b96a07a7ea4675d80358e4c6d0a7b287b367e4df2e1ee384e7d59ecc1.
//
// Solidity: event StakeHelpersConfigured(address indexed subjectGateway, address indexed allocator)
func (_FortaStaking *FortaStakingFilterer) ParseStakeHelpersConfigured(log types.Log) (*FortaStakingStakeHelpersConfigured, error) {
	event := new(FortaStakingStakeHelpersConfigured)
	if err := _FortaStaking.contract.UnpackLog(event, "StakeHelpersConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingTokensSweptIterator is returned from FilterTokensSwept and is used to iterate over the raw logs and unpacked data for TokensSwept events raised by the FortaStaking contract.
type FortaStakingTokensSweptIterator struct {
	Event *FortaStakingTokensSwept // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingTokensSweptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingTokensSwept)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingTokensSwept)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingTokensSweptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingTokensSweptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingTokensSwept represents a TokensSwept event raised by the FortaStaking contract.
type FortaStakingTokensSwept struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensSwept is a free log retrieval operation binding the contract event 0xd092d7fceb5ea5a962639fcc27a7bb315e7637e699e3b108cd570c38c7584300.
//
// Solidity: event TokensSwept(address indexed token, address to, uint256 amount)
func (_FortaStaking *FortaStakingFilterer) FilterTokensSwept(opts *bind.FilterOpts, token []common.Address) (*FortaStakingTokensSweptIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "TokensSwept", tokenRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingTokensSweptIterator{contract: _FortaStaking.contract, event: "TokensSwept", logs: logs, sub: sub}, nil
}

// WatchTokensSwept is a free log subscription operation binding the contract event 0xd092d7fceb5ea5a962639fcc27a7bb315e7637e699e3b108cd570c38c7584300.
//
// Solidity: event TokensSwept(address indexed token, address to, uint256 amount)
func (_FortaStaking *FortaStakingFilterer) WatchTokensSwept(opts *bind.WatchOpts, sink chan<- *FortaStakingTokensSwept, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "TokensSwept", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingTokensSwept)
				if err := _FortaStaking.contract.UnpackLog(event, "TokensSwept", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensSwept is a log parse operation binding the contract event 0xd092d7fceb5ea5a962639fcc27a7bb315e7637e699e3b108cd570c38c7584300.
//
// Solidity: event TokensSwept(address indexed token, address to, uint256 amount)
func (_FortaStaking *FortaStakingFilterer) ParseTokensSwept(log types.Log) (*FortaStakingTokensSwept, error) {
	event := new(FortaStakingTokensSwept)
	if err := _FortaStaking.contract.UnpackLog(event, "TokensSwept", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the FortaStaking contract.
type FortaStakingTransferBatchIterator struct {
	Event *FortaStakingTransferBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingTransferBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingTransferBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingTransferBatch represents a TransferBatch event raised by the FortaStaking contract.
type FortaStakingTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_FortaStaking *FortaStakingFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*FortaStakingTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingTransferBatchIterator{contract: _FortaStaking.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_FortaStaking *FortaStakingFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *FortaStakingTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingTransferBatch)
				if err := _FortaStaking.contract.UnpackLog(event, "TransferBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_FortaStaking *FortaStakingFilterer) ParseTransferBatch(log types.Log) (*FortaStakingTransferBatch, error) {
	event := new(FortaStakingTransferBatch)
	if err := _FortaStaking.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the FortaStaking contract.
type FortaStakingTransferSingleIterator struct {
	Event *FortaStakingTransferSingle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingTransferSingle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingTransferSingle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingTransferSingle represents a TransferSingle event raised by the FortaStaking contract.
type FortaStakingTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_FortaStaking *FortaStakingFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*FortaStakingTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingTransferSingleIterator{contract: _FortaStaking.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_FortaStaking *FortaStakingFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *FortaStakingTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingTransferSingle)
				if err := _FortaStaking.contract.UnpackLog(event, "TransferSingle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_FortaStaking *FortaStakingFilterer) ParseTransferSingle(log types.Log) (*FortaStakingTransferSingle, error) {
	event := new(FortaStakingTransferSingle)
	if err := _FortaStaking.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingTreasurySetIterator is returned from FilterTreasurySet and is used to iterate over the raw logs and unpacked data for TreasurySet events raised by the FortaStaking contract.
type FortaStakingTreasurySetIterator struct {
	Event *FortaStakingTreasurySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingTreasurySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingTreasurySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingTreasurySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingTreasurySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingTreasurySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingTreasurySet represents a TreasurySet event raised by the FortaStaking contract.
type FortaStakingTreasurySet struct {
	NewTreasury common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTreasurySet is a free log retrieval operation binding the contract event 0x3c864541ef71378c6229510ed90f376565ee42d9c5e0904a984a9e863e6db44f.
//
// Solidity: event TreasurySet(address newTreasury)
func (_FortaStaking *FortaStakingFilterer) FilterTreasurySet(opts *bind.FilterOpts) (*FortaStakingTreasurySetIterator, error) {

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "TreasurySet")
	if err != nil {
		return nil, err
	}
	return &FortaStakingTreasurySetIterator{contract: _FortaStaking.contract, event: "TreasurySet", logs: logs, sub: sub}, nil
}

// WatchTreasurySet is a free log subscription operation binding the contract event 0x3c864541ef71378c6229510ed90f376565ee42d9c5e0904a984a9e863e6db44f.
//
// Solidity: event TreasurySet(address newTreasury)
func (_FortaStaking *FortaStakingFilterer) WatchTreasurySet(opts *bind.WatchOpts, sink chan<- *FortaStakingTreasurySet) (event.Subscription, error) {

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "TreasurySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingTreasurySet)
				if err := _FortaStaking.contract.UnpackLog(event, "TreasurySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTreasurySet is a log parse operation binding the contract event 0x3c864541ef71378c6229510ed90f376565ee42d9c5e0904a984a9e863e6db44f.
//
// Solidity: event TreasurySet(address newTreasury)
func (_FortaStaking *FortaStakingFilterer) ParseTreasurySet(log types.Log) (*FortaStakingTreasurySet, error) {
	event := new(FortaStakingTreasurySet)
	if err := _FortaStaking.contract.UnpackLog(event, "TreasurySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the FortaStaking contract.
type FortaStakingURIIterator struct {
	Event *FortaStakingURI // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingURI)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingURI)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingURI represents a URI event raised by the FortaStaking contract.
type FortaStakingURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_FortaStaking *FortaStakingFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*FortaStakingURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingURIIterator{contract: _FortaStaking.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_FortaStaking *FortaStakingFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *FortaStakingURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingURI)
				if err := _FortaStaking.contract.UnpackLog(event, "URI", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_FortaStaking *FortaStakingFilterer) ParseURI(log types.Log) (*FortaStakingURI, error) {
	event := new(FortaStakingURI)
	if err := _FortaStaking.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the FortaStaking contract.
type FortaStakingUpgradedIterator struct {
	Event *FortaStakingUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingUpgraded represents a Upgraded event raised by the FortaStaking contract.
type FortaStakingUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FortaStaking *FortaStakingFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*FortaStakingUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingUpgradedIterator{contract: _FortaStaking.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FortaStaking *FortaStakingFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *FortaStakingUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingUpgraded)
				if err := _FortaStaking.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FortaStaking *FortaStakingFilterer) ParseUpgraded(log types.Log) (*FortaStakingUpgraded, error) {
	event := new(FortaStakingUpgraded)
	if err := _FortaStaking.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingWithdrawalExecutedIterator is returned from FilterWithdrawalExecuted and is used to iterate over the raw logs and unpacked data for WithdrawalExecuted events raised by the FortaStaking contract.
type FortaStakingWithdrawalExecutedIterator struct {
	Event *FortaStakingWithdrawalExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingWithdrawalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingWithdrawalExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingWithdrawalExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingWithdrawalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingWithdrawalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingWithdrawalExecuted represents a WithdrawalExecuted event raised by the FortaStaking contract.
type FortaStakingWithdrawalExecuted struct {
	SubjectType uint8
	Subject     *big.Int
	Account     common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalExecuted is a free log retrieval operation binding the contract event 0x07e9e8a51e2cf8929e95153e486868bda03206cdbbf30ee732b17e62b43c2827.
//
// Solidity: event WithdrawalExecuted(uint8 indexed subjectType, uint256 indexed subject, address indexed account)
func (_FortaStaking *FortaStakingFilterer) FilterWithdrawalExecuted(opts *bind.FilterOpts, subjectType []uint8, subject []*big.Int, account []common.Address) (*FortaStakingWithdrawalExecutedIterator, error) {

	var subjectTypeRule []interface{}
	for _, subjectTypeItem := range subjectType {
		subjectTypeRule = append(subjectTypeRule, subjectTypeItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "WithdrawalExecuted", subjectTypeRule, subjectRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingWithdrawalExecutedIterator{contract: _FortaStaking.contract, event: "WithdrawalExecuted", logs: logs, sub: sub}, nil
}

// WatchWithdrawalExecuted is a free log subscription operation binding the contract event 0x07e9e8a51e2cf8929e95153e486868bda03206cdbbf30ee732b17e62b43c2827.
//
// Solidity: event WithdrawalExecuted(uint8 indexed subjectType, uint256 indexed subject, address indexed account)
func (_FortaStaking *FortaStakingFilterer) WatchWithdrawalExecuted(opts *bind.WatchOpts, sink chan<- *FortaStakingWithdrawalExecuted, subjectType []uint8, subject []*big.Int, account []common.Address) (event.Subscription, error) {

	var subjectTypeRule []interface{}
	for _, subjectTypeItem := range subjectType {
		subjectTypeRule = append(subjectTypeRule, subjectTypeItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "WithdrawalExecuted", subjectTypeRule, subjectRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingWithdrawalExecuted)
				if err := _FortaStaking.contract.UnpackLog(event, "WithdrawalExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalExecuted is a log parse operation binding the contract event 0x07e9e8a51e2cf8929e95153e486868bda03206cdbbf30ee732b17e62b43c2827.
//
// Solidity: event WithdrawalExecuted(uint8 indexed subjectType, uint256 indexed subject, address indexed account)
func (_FortaStaking *FortaStakingFilterer) ParseWithdrawalExecuted(log types.Log) (*FortaStakingWithdrawalExecuted, error) {
	event := new(FortaStakingWithdrawalExecuted)
	if err := _FortaStaking.contract.UnpackLog(event, "WithdrawalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingWithdrawalInitiatedIterator is returned from FilterWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for WithdrawalInitiated events raised by the FortaStaking contract.
type FortaStakingWithdrawalInitiatedIterator struct {
	Event *FortaStakingWithdrawalInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FortaStakingWithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingWithdrawalInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FortaStakingWithdrawalInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FortaStakingWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingWithdrawalInitiated represents a WithdrawalInitiated event raised by the FortaStaking contract.
type FortaStakingWithdrawalInitiated struct {
	SubjectType uint8
	Subject     *big.Int
	Account     common.Address
	Deadline    uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalInitiated is a free log retrieval operation binding the contract event 0x8b14266b7a7bfb46d73cda0a8ca3a635f38e8fb4b275b68aff68cb1b5a2aea3a.
//
// Solidity: event WithdrawalInitiated(uint8 indexed subjectType, uint256 indexed subject, address indexed account, uint64 deadline)
func (_FortaStaking *FortaStakingFilterer) FilterWithdrawalInitiated(opts *bind.FilterOpts, subjectType []uint8, subject []*big.Int, account []common.Address) (*FortaStakingWithdrawalInitiatedIterator, error) {

	var subjectTypeRule []interface{}
	for _, subjectTypeItem := range subjectType {
		subjectTypeRule = append(subjectTypeRule, subjectTypeItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "WithdrawalInitiated", subjectTypeRule, subjectRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingWithdrawalInitiatedIterator{contract: _FortaStaking.contract, event: "WithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalInitiated is a free log subscription operation binding the contract event 0x8b14266b7a7bfb46d73cda0a8ca3a635f38e8fb4b275b68aff68cb1b5a2aea3a.
//
// Solidity: event WithdrawalInitiated(uint8 indexed subjectType, uint256 indexed subject, address indexed account, uint64 deadline)
func (_FortaStaking *FortaStakingFilterer) WatchWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *FortaStakingWithdrawalInitiated, subjectType []uint8, subject []*big.Int, account []common.Address) (event.Subscription, error) {

	var subjectTypeRule []interface{}
	for _, subjectTypeItem := range subjectType {
		subjectTypeRule = append(subjectTypeRule, subjectTypeItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "WithdrawalInitiated", subjectTypeRule, subjectRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingWithdrawalInitiated)
				if err := _FortaStaking.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalInitiated is a log parse operation binding the contract event 0x8b14266b7a7bfb46d73cda0a8ca3a635f38e8fb4b275b68aff68cb1b5a2aea3a.
//
// Solidity: event WithdrawalInitiated(uint8 indexed subjectType, uint256 indexed subject, address indexed account, uint64 deadline)
func (_FortaStaking *FortaStakingFilterer) ParseWithdrawalInitiated(log types.Log) (*FortaStakingWithdrawalInitiated, error) {
	event := new(FortaStakingWithdrawalInitiated)
	if err := _FortaStaking.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
