// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_forta_staking

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
)

// FortaStakingMetaData contains all meta data concerning the FortaStaking contract.
var FortaStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"AccessManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newWithdrawalDelay\",\"type\":\"uint256\"}],\"name\":\"DelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isFrozen\",\"type\":\"bool\"}],\"name\":\"Froze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"MaxStakeReached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Rewarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"StakeParamsManagerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensSwept\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"TreasurySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WithdrawalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"activeStakeFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"availableReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeValue\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"freeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"inactiveSharesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"inactiveStakeFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__router\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"__stakedToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"__withdrawalDelay\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"__treasury\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sharesValue\",\"type\":\"uint256\"}],\"name\":\"initiateWithdrawal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"isFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"relayPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"releaseReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"setAccessManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newDelay\",\"type\":\"uint64\"}],\"name\":\"setDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ensRegistry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ensName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStakeController\",\"name\":\"newStakingParameters\",\"type\":\"address\"}],\"name\":\"setStakingParametersManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"sharesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeValue\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweep\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalActiveStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"totalInactiveShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInactiveStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FortaStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use FortaStakingMetaData.ABI instead.
var FortaStakingABI = FortaStakingMetaData.ABI

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
	parsed, err := abi.JSON(strings.NewReader(FortaStakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// AvailableReward is a free data retrieval call binding the contract method 0x8c5588ac.
//
// Solidity: function availableReward(uint8 subjectType, uint256 subject, address account) view returns(uint256)
func (_FortaStaking *FortaStakingCaller) AvailableReward(opts *bind.CallOpts, subjectType uint8, subject *big.Int, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "availableReward", subjectType, subject, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableReward is a free data retrieval call binding the contract method 0x8c5588ac.
//
// Solidity: function availableReward(uint8 subjectType, uint256 subject, address account) view returns(uint256)
func (_FortaStaking *FortaStakingSession) AvailableReward(subjectType uint8, subject *big.Int, account common.Address) (*big.Int, error) {
	return _FortaStaking.Contract.AvailableReward(&_FortaStaking.CallOpts, subjectType, subject, account)
}

// AvailableReward is a free data retrieval call binding the contract method 0x8c5588ac.
//
// Solidity: function availableReward(uint8 subjectType, uint256 subject, address account) view returns(uint256)
func (_FortaStaking *FortaStakingCallerSession) AvailableReward(subjectType uint8, subject *big.Int, account common.Address) (*big.Int, error) {
	return _FortaStaking.Contract.AvailableReward(&_FortaStaking.CallOpts, subjectType, subject, account)
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

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_FortaStaking *FortaStakingCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _FortaStaking.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_FortaStaking *FortaStakingSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _FortaStaking.Contract.IsTrustedForwarder(&_FortaStaking.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_FortaStaking *FortaStakingCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _FortaStaking.Contract.IsTrustedForwarder(&_FortaStaking.CallOpts, forwarder)
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

// Initialize is a paid mutator transaction binding the contract method 0x8432d6b6.
//
// Solidity: function initialize(address __manager, address __router, address __stakedToken, uint64 __withdrawalDelay, address __treasury) returns()
func (_FortaStaking *FortaStakingTransactor) Initialize(opts *bind.TransactOpts, __manager common.Address, __router common.Address, __stakedToken common.Address, __withdrawalDelay uint64, __treasury common.Address) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "initialize", __manager, __router, __stakedToken, __withdrawalDelay, __treasury)
}

// Initialize is a paid mutator transaction binding the contract method 0x8432d6b6.
//
// Solidity: function initialize(address __manager, address __router, address __stakedToken, uint64 __withdrawalDelay, address __treasury) returns()
func (_FortaStaking *FortaStakingSession) Initialize(__manager common.Address, __router common.Address, __stakedToken common.Address, __withdrawalDelay uint64, __treasury common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.Initialize(&_FortaStaking.TransactOpts, __manager, __router, __stakedToken, __withdrawalDelay, __treasury)
}

// Initialize is a paid mutator transaction binding the contract method 0x8432d6b6.
//
// Solidity: function initialize(address __manager, address __router, address __stakedToken, uint64 __withdrawalDelay, address __treasury) returns()
func (_FortaStaking *FortaStakingTransactorSession) Initialize(__manager common.Address, __router common.Address, __stakedToken common.Address, __withdrawalDelay uint64, __treasury common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.Initialize(&_FortaStaking.TransactOpts, __manager, __router, __stakedToken, __withdrawalDelay, __treasury)
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

// ReleaseReward is a paid mutator transaction binding the contract method 0xd1e59d1c.
//
// Solidity: function releaseReward(uint8 subjectType, uint256 subject, address account) returns(uint256)
func (_FortaStaking *FortaStakingTransactor) ReleaseReward(opts *bind.TransactOpts, subjectType uint8, subject *big.Int, account common.Address) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "releaseReward", subjectType, subject, account)
}

// ReleaseReward is a paid mutator transaction binding the contract method 0xd1e59d1c.
//
// Solidity: function releaseReward(uint8 subjectType, uint256 subject, address account) returns(uint256)
func (_FortaStaking *FortaStakingSession) ReleaseReward(subjectType uint8, subject *big.Int, account common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.ReleaseReward(&_FortaStaking.TransactOpts, subjectType, subject, account)
}

// ReleaseReward is a paid mutator transaction binding the contract method 0xd1e59d1c.
//
// Solidity: function releaseReward(uint8 subjectType, uint256 subject, address account) returns(uint256)
func (_FortaStaking *FortaStakingTransactorSession) ReleaseReward(subjectType uint8, subject *big.Int, account common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.ReleaseReward(&_FortaStaking.TransactOpts, subjectType, subject, account)
}

// Reward is a paid mutator transaction binding the contract method 0x75eb81d6.
//
// Solidity: function reward(uint8 subjectType, uint256 subject, uint256 value) returns()
func (_FortaStaking *FortaStakingTransactor) Reward(opts *bind.TransactOpts, subjectType uint8, subject *big.Int, value *big.Int) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "reward", subjectType, subject, value)
}

// Reward is a paid mutator transaction binding the contract method 0x75eb81d6.
//
// Solidity: function reward(uint8 subjectType, uint256 subject, uint256 value) returns()
func (_FortaStaking *FortaStakingSession) Reward(subjectType uint8, subject *big.Int, value *big.Int) (*types.Transaction, error) {
	return _FortaStaking.Contract.Reward(&_FortaStaking.TransactOpts, subjectType, subject, value)
}

// Reward is a paid mutator transaction binding the contract method 0x75eb81d6.
//
// Solidity: function reward(uint8 subjectType, uint256 subject, uint256 value) returns()
func (_FortaStaking *FortaStakingTransactorSession) Reward(subjectType uint8, subject *big.Int, value *big.Int) (*types.Transaction, error) {
	return _FortaStaking.Contract.Reward(&_FortaStaking.TransactOpts, subjectType, subject, value)
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

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address newRouter) returns()
func (_FortaStaking *FortaStakingTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "setRouter", newRouter)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address newRouter) returns()
func (_FortaStaking *FortaStakingSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.SetRouter(&_FortaStaking.TransactOpts, newRouter)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address newRouter) returns()
func (_FortaStaking *FortaStakingTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.SetRouter(&_FortaStaking.TransactOpts, newRouter)
}

// SetStakingParametersManager is a paid mutator transaction binding the contract method 0x0a615413.
//
// Solidity: function setStakingParametersManager(address newStakingParameters) returns()
func (_FortaStaking *FortaStakingTransactor) SetStakingParametersManager(opts *bind.TransactOpts, newStakingParameters common.Address) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "setStakingParametersManager", newStakingParameters)
}

// SetStakingParametersManager is a paid mutator transaction binding the contract method 0x0a615413.
//
// Solidity: function setStakingParametersManager(address newStakingParameters) returns()
func (_FortaStaking *FortaStakingSession) SetStakingParametersManager(newStakingParameters common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.SetStakingParametersManager(&_FortaStaking.TransactOpts, newStakingParameters)
}

// SetStakingParametersManager is a paid mutator transaction binding the contract method 0x0a615413.
//
// Solidity: function setStakingParametersManager(address newStakingParameters) returns()
func (_FortaStaking *FortaStakingTransactorSession) SetStakingParametersManager(newStakingParameters common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.SetStakingParametersManager(&_FortaStaking.TransactOpts, newStakingParameters)
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

// Slash is a paid mutator transaction binding the contract method 0xfe54ec80.
//
// Solidity: function slash(uint8 subjectType, uint256 subject, uint256 stakeValue) returns(uint256)
func (_FortaStaking *FortaStakingTransactor) Slash(opts *bind.TransactOpts, subjectType uint8, subject *big.Int, stakeValue *big.Int) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "slash", subjectType, subject, stakeValue)
}

// Slash is a paid mutator transaction binding the contract method 0xfe54ec80.
//
// Solidity: function slash(uint8 subjectType, uint256 subject, uint256 stakeValue) returns(uint256)
func (_FortaStaking *FortaStakingSession) Slash(subjectType uint8, subject *big.Int, stakeValue *big.Int) (*types.Transaction, error) {
	return _FortaStaking.Contract.Slash(&_FortaStaking.TransactOpts, subjectType, subject, stakeValue)
}

// Slash is a paid mutator transaction binding the contract method 0xfe54ec80.
//
// Solidity: function slash(uint8 subjectType, uint256 subject, uint256 stakeValue) returns(uint256)
func (_FortaStaking *FortaStakingTransactorSession) Slash(subjectType uint8, subject *big.Int, stakeValue *big.Int) (*types.Transaction, error) {
	return _FortaStaking.Contract.Slash(&_FortaStaking.TransactOpts, subjectType, subject, stakeValue)
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

// FortaStakingReleasedIterator is returned from FilterReleased and is used to iterate over the raw logs and unpacked data for Released events raised by the FortaStaking contract.
type FortaStakingReleasedIterator struct {
	Event *FortaStakingReleased // Event containing the contract specifics and raw log

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
func (it *FortaStakingReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingReleased)
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
		it.Event = new(FortaStakingReleased)
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
func (it *FortaStakingReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingReleased represents a Released event raised by the FortaStaking contract.
type FortaStakingReleased struct {
	SubjectType uint8
	Subject     *big.Int
	To          common.Address
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReleased is a free log retrieval operation binding the contract event 0x8f404fb0169bf58e5bb9d571055da9ac0fc4eb973f2e6c7d1b108024381f0fc6.
//
// Solidity: event Released(uint8 indexed subjectType, uint256 indexed subject, address indexed to, uint256 value)
func (_FortaStaking *FortaStakingFilterer) FilterReleased(opts *bind.FilterOpts, subjectType []uint8, subject []*big.Int, to []common.Address) (*FortaStakingReleasedIterator, error) {

	var subjectTypeRule []interface{}
	for _, subjectTypeItem := range subjectType {
		subjectTypeRule = append(subjectTypeRule, subjectTypeItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "Released", subjectTypeRule, subjectRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingReleasedIterator{contract: _FortaStaking.contract, event: "Released", logs: logs, sub: sub}, nil
}

// WatchReleased is a free log subscription operation binding the contract event 0x8f404fb0169bf58e5bb9d571055da9ac0fc4eb973f2e6c7d1b108024381f0fc6.
//
// Solidity: event Released(uint8 indexed subjectType, uint256 indexed subject, address indexed to, uint256 value)
func (_FortaStaking *FortaStakingFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *FortaStakingReleased, subjectType []uint8, subject []*big.Int, to []common.Address) (event.Subscription, error) {

	var subjectTypeRule []interface{}
	for _, subjectTypeItem := range subjectType {
		subjectTypeRule = append(subjectTypeRule, subjectTypeItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "Released", subjectTypeRule, subjectRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingReleased)
				if err := _FortaStaking.contract.UnpackLog(event, "Released", log); err != nil {
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

// ParseReleased is a log parse operation binding the contract event 0x8f404fb0169bf58e5bb9d571055da9ac0fc4eb973f2e6c7d1b108024381f0fc6.
//
// Solidity: event Released(uint8 indexed subjectType, uint256 indexed subject, address indexed to, uint256 value)
func (_FortaStaking *FortaStakingFilterer) ParseReleased(log types.Log) (*FortaStakingReleased, error) {
	event := new(FortaStakingReleased)
	if err := _FortaStaking.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FortaStakingRewardedIterator is returned from FilterRewarded and is used to iterate over the raw logs and unpacked data for Rewarded events raised by the FortaStaking contract.
type FortaStakingRewardedIterator struct {
	Event *FortaStakingRewarded // Event containing the contract specifics and raw log

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
func (it *FortaStakingRewardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingRewarded)
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
		it.Event = new(FortaStakingRewarded)
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
func (it *FortaStakingRewardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingRewardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingRewarded represents a Rewarded event raised by the FortaStaking contract.
type FortaStakingRewarded struct {
	SubjectType uint8
	Subject     *big.Int
	From        common.Address
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewarded is a free log retrieval operation binding the contract event 0x4bae81d7c78d9effd9b0ffe353b35b51c2695820985aa889b3d8916a017f60a0.
//
// Solidity: event Rewarded(uint8 indexed subjectType, uint256 indexed subject, address indexed from, uint256 value)
func (_FortaStaking *FortaStakingFilterer) FilterRewarded(opts *bind.FilterOpts, subjectType []uint8, subject []*big.Int, from []common.Address) (*FortaStakingRewardedIterator, error) {

	var subjectTypeRule []interface{}
	for _, subjectTypeItem := range subjectType {
		subjectTypeRule = append(subjectTypeRule, subjectTypeItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "Rewarded", subjectTypeRule, subjectRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingRewardedIterator{contract: _FortaStaking.contract, event: "Rewarded", logs: logs, sub: sub}, nil
}

// WatchRewarded is a free log subscription operation binding the contract event 0x4bae81d7c78d9effd9b0ffe353b35b51c2695820985aa889b3d8916a017f60a0.
//
// Solidity: event Rewarded(uint8 indexed subjectType, uint256 indexed subject, address indexed from, uint256 value)
func (_FortaStaking *FortaStakingFilterer) WatchRewarded(opts *bind.WatchOpts, sink chan<- *FortaStakingRewarded, subjectType []uint8, subject []*big.Int, from []common.Address) (event.Subscription, error) {

	var subjectTypeRule []interface{}
	for _, subjectTypeItem := range subjectType {
		subjectTypeRule = append(subjectTypeRule, subjectTypeItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "Rewarded", subjectTypeRule, subjectRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingRewarded)
				if err := _FortaStaking.contract.UnpackLog(event, "Rewarded", log); err != nil {
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

// ParseRewarded is a log parse operation binding the contract event 0x4bae81d7c78d9effd9b0ffe353b35b51c2695820985aa889b3d8916a017f60a0.
//
// Solidity: event Rewarded(uint8 indexed subjectType, uint256 indexed subject, address indexed from, uint256 value)
func (_FortaStaking *FortaStakingFilterer) ParseRewarded(log types.Log) (*FortaStakingRewarded, error) {
	event := new(FortaStakingRewarded)
	if err := _FortaStaking.contract.UnpackLog(event, "Rewarded", log); err != nil {
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

// FortaStakingStakeParamsManagerSetIterator is returned from FilterStakeParamsManagerSet and is used to iterate over the raw logs and unpacked data for StakeParamsManagerSet events raised by the FortaStaking contract.
type FortaStakingStakeParamsManagerSetIterator struct {
	Event *FortaStakingStakeParamsManagerSet // Event containing the contract specifics and raw log

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
func (it *FortaStakingStakeParamsManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FortaStakingStakeParamsManagerSet)
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
		it.Event = new(FortaStakingStakeParamsManagerSet)
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
func (it *FortaStakingStakeParamsManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FortaStakingStakeParamsManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FortaStakingStakeParamsManagerSet represents a StakeParamsManagerSet event raised by the FortaStaking contract.
type FortaStakingStakeParamsManagerSet struct {
	NewManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStakeParamsManagerSet is a free log retrieval operation binding the contract event 0x06b2874e4c6a9fd4be614bef4bf7205f773309fd0d578a6f9b08d7b1f95347fb.
//
// Solidity: event StakeParamsManagerSet(address indexed newManager)
func (_FortaStaking *FortaStakingFilterer) FilterStakeParamsManagerSet(opts *bind.FilterOpts, newManager []common.Address) (*FortaStakingStakeParamsManagerSetIterator, error) {

	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "StakeParamsManagerSet", newManagerRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingStakeParamsManagerSetIterator{contract: _FortaStaking.contract, event: "StakeParamsManagerSet", logs: logs, sub: sub}, nil
}

// WatchStakeParamsManagerSet is a free log subscription operation binding the contract event 0x06b2874e4c6a9fd4be614bef4bf7205f773309fd0d578a6f9b08d7b1f95347fb.
//
// Solidity: event StakeParamsManagerSet(address indexed newManager)
func (_FortaStaking *FortaStakingFilterer) WatchStakeParamsManagerSet(opts *bind.WatchOpts, sink chan<- *FortaStakingStakeParamsManagerSet, newManager []common.Address) (event.Subscription, error) {

	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "StakeParamsManagerSet", newManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FortaStakingStakeParamsManagerSet)
				if err := _FortaStaking.contract.UnpackLog(event, "StakeParamsManagerSet", log); err != nil {
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

// ParseStakeParamsManagerSet is a log parse operation binding the contract event 0x06b2874e4c6a9fd4be614bef4bf7205f773309fd0d578a6f9b08d7b1f95347fb.
//
// Solidity: event StakeParamsManagerSet(address indexed newManager)
func (_FortaStaking *FortaStakingFilterer) ParseStakeParamsManagerSet(log types.Log) (*FortaStakingStakeParamsManagerSet, error) {
	event := new(FortaStakingStakeParamsManagerSet)
	if err := _FortaStaking.contract.UnpackLog(event, "StakeParamsManagerSet", log); err != nil {
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
