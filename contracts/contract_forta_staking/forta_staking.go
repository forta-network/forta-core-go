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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"AccessManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newWithdrawalDelay\",\"type\":\"uint256\"}],\"name\":\"DelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isFrozen\",\"type\":\"bool\"}],\"name\":\"Froze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"MaxStakeReached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Rewarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SlashedShareSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"StakeParamsManagerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensSwept\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"TreasurySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WithdrawalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_WITHDRAWAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_WITHDRAWAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"activeSharesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"activeSharesToStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"activeStakeFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"availableReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeValue\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"freeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"inactiveSharesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"inactiveSharesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"inactiveSharesToStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"inactiveStakeFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"__withdrawalDelay\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"__treasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__stakedToken\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sharesValue\",\"type\":\"uint256\"}],\"name\":\"initiateWithdrawal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"isFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"relayPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"releaseReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"setAccessManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newDelay\",\"type\":\"uint64\"}],\"name\":\"setDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ensRegistry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ensName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStakeController\",\"name\":\"newStakingParameters\",\"type\":\"address\"}],\"name\":\"setStakingParametersManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"sharesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"proposerPercent\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"activeSharesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stakeToActiveShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"inactiveSharesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stakeToInactiveShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweep\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalActiveStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"totalInactiveShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInactiveStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523060a0523480156200001557600080fd5b5060405162005b9438038062005b94833981016040819052620000389162000180565b6001600160a01b038116608052600054610100900460ff1615808015620000665750600054600160ff909116105b8062000096575062000083306200017160201b620027281760201c565b15801562000096575060005460ff166001145b620000fe5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b6000805460ff19166001179055801562000122576000805461ff0019166101001790555b801562000169576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050620001b2565b6001600160a01b03163b151590565b6000602082840312156200019357600080fd5b81516001600160a01b0381168114620001ab57600080fd5b9392505050565b60805160a051615999620001fb6000396000818161114301528181611183015281816115090152818161154901526115d801526000818161061001526141dd01526159996000f3fe6080604052600436106102e35760003560e01c806362772b1411610190578063c0d78655116100dc578063dc4653ef11610095578063f08d35ee1161006f578063f08d35ee1461098f578063f0f44260146109af578063f242432a146109cf578063f8058b06146109ef57600080fd5b8063dc4653ef146108ed578063e985e9c51461090d578063edea0bac1461095757600080fd5b8063c0d786551461082c578063c10733021461084c578063c95808041461086c578063cc7a262e1461088c578063d1e59d1c146108ad578063da5bfb94146108cd57600080fd5b8063a22cb46511610149578063ac9650d811610123578063ac9650d814610791578063b8dc491b146107be578063bd85b039146107de578063c07975ad1461080c57600080fd5b8063a22cb4651461073a578063a238f9df1461075a578063a290bf381461077157600080fd5b806362772b141461067a57806364a0f9011461069a57806375e130ad146106ba57806375eb81d6146106da57806389a0ca6b146106fa5780638c5588ac1461071a57600080fd5b80633121db1c1161024f5780634e1273f41161020857806352d1902d116101e257806352d1902d146105ad57806354fd4d50146105c2578063572b6c05146105f357806361d027b31461064057600080fd5b80634e1273f41461053d5780634f1ef2861461056a5780634f558e791461057d57600080fd5b80633121db1c14610490578063321ebc54146104b05780633659cfe6146104d0578063371f4226146104f05780633f489914146105065780634a5f2b5d1461052657600080fd5b8063115a90ee116102a1578063115a90ee146103da5780631a82ef4e146103fa578063226cc3001461041a57806328f731481461043a5780632cb31144146104505780632eb2c2d61461047057600080fd5b8062fdd58e146102e857806301ffc9a71461031b57806302fe53051461034b5780630a6154131461036d5780630e10103f1461038d5780630e89341c146103ad575b600080fd5b3480156102f457600080fd5b5061030861030336600461481b565b610a0f565b6040519081526020015b60405180910390f35b34801561032757600080fd5b5061033b61033636600461485d565b610aab565b6040519015158152602001610312565b34801561035757600080fd5b5061036b610366366004614919565b610afb565b005b34801561037957600080fd5b5061036b610388366004614961565b610b53565b34801561039957600080fd5b506103086103a8366004614994565b610c14565b3480156103b957600080fd5b506103cd6103c83660046149b0565b610c2a565b6040516103129190614a21565b3480156103e657600080fd5b506103086103f5366004614a34565b610cbf565b34801561040657600080fd5b50610308610415366004614a34565b610d01565b34801561042657600080fd5b5061036b610435366004614a64565b610d3d565b34801561044657600080fd5b506101c554610308565b34801561045c57600080fd5b5061030861046b366004614aa4565b610e3d565b34801561047c57600080fd5b5061036b61048b366004614b8b565b611081565b34801561049c57600080fd5b5061036b6104ab366004614c38565b6110df565b3480156104bc57600080fd5b506103086104cb366004614994565b61112a565b3480156104dc57600080fd5b5061036b6104eb366004614961565b611139565b3480156104fc57600080fd5b506101c754610308565b34801561051257600080fd5b50610308610521366004614994565b611218565b34801561053257600080fd5b506103086201518081565b34801561054957600080fd5b5061055d610558366004614cbc565b6113d6565b6040516103129190614dc3565b61036b610578366004614dd6565b6114ff565b34801561058957600080fd5b5061033b6105983660046149b0565b60009081526101916020526040902054151590565b3480156105b957600080fd5b506103086115cb565b3480156105ce57600080fd5b506103cd60405180604001604052806005815260200164302e312e3160d81b81525081565b3480156105ff57600080fd5b5061033b61060e366004614961565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0390811691161490565b34801561064c57600080fd5b506101cd54600160401b90046001600160a01b03165b6040516001600160a01b039091168152602001610312565b34801561068657600080fd5b50610308610695366004614a34565b61167f565b3480156106a657600080fd5b506103086106b5366004614a34565b611696565b3480156106c657600080fd5b5061033b6106d5366004614994565b6116c5565b3480156106e657600080fd5b5061036b6106f5366004614aa4565b6116f1565b34801561070657600080fd5b5061036b610715366004614e32565b6117b7565b34801561072657600080fd5b50610308610735366004614ea1565b611a03565b34801561074657600080fd5b5061036b610755366004614ed6565b611a25565b34801561076657600080fd5b506103086276a70081565b34801561077d57600080fd5b5061030861078c366004614994565b611a37565b34801561079d57600080fd5b506107b16107ac366004614f0f565b611a59565b6040516103129190614f83565b3480156107ca57600080fd5b506103086107d9366004614fe5565b611b46565b3480156107ea57600080fd5b506103086107f93660046149b0565b6000908152610191602052604090205490565b34801561081857600080fd5b5061036b610827366004615013565b611c98565b34801561083857600080fd5b5061036b610847366004614961565b611d3d565b34801561085857600080fd5b5061036b61086736600461505a565b611de6565b34801561087857600080fd5b5061036b610887366004614961565b611eda565b34801561089857600080fd5b506101c354610662906001600160a01b031681565b3480156108b957600080fd5b506103086108c8366004614ea1565b611f98565b3480156108d957600080fd5b506103086108e8366004614ea1565b612128565b3480156108f957600080fd5b50610308610908366004614994565b612138565b34801561091957600080fd5b5061033b610928366004614fe5565b6001600160a01b0391821660009081526101606020908152604080832093909416825291909152205460ff1690565b34801561096357600080fd5b50610977610972366004614aa4565b61215a565b6040516001600160401b039091168152602001610312565b34801561099b57600080fd5b506103086109aa366004614ea1565b61230d565b3480156109bb57600080fd5b5061036b6109ca366004614961565b61231d565b3480156109db57600080fd5b5061036b6109ea366004615075565b6123dd565b3480156109fb57600080fd5b50610308610a0a3660046150dd565b612434565b60006001600160a01b038316610a7f5760405162461bcd60e51b815260206004820152602a60248201527f455243313135353a2061646472657373207a65726f206973206e6f742061207660448201526930b634b21037bbb732b960b11b60648201526084015b60405180910390fd5b50600081815261015f602090815260408083206001600160a01b03861684529091529020545b92915050565b60006001600160e01b03198216636cdb3d1360e11b1480610adc57506001600160e01b031982166303a24d0760e21b145b80610aa557506301ffc9a760e01b6001600160e01b0319831614610aa5565b6000610b0e81610b09612737565b612741565b610b465780610b1b612737565b6040516301d4003760e61b815260048101929092526001600160a01b03166024820152604401610a76565b610b4f826127b7565b5050565b6000610b6181610b09612737565b610b6e5780610b1b612737565b6001600160a01b038216610bbc5760405163eac0d38960e01b81526020600482015260146024820152736e65775374616b696e67506172616d657465727360601b6044820152606401610a76565b6040516001600160a01b038316907f06b2874e4c6a9fd4be614bef4bf7205f773309fd0d578a6f9b08d7b1f95347fb90600090a2506101ce80546001600160a01b0319166001600160a01b0392909216919091179055565b6000610c236107f984846127c4565b9392505050565b60606101618054610c3a9061512d565b80601f0160208091040260200160405190810160405280929190818152602001828054610c669061512d565b8015610cb35780601f10610c8857610100808354040283529160200191610cb3565b820191906000526020600020905b815481529060010190602001808311610c9657829003601f168201915b50505050509050919050565b600082815261019160205260408120548015610cf65760008481526101c66020526040902054610cf1905b8483612810565b610cf9565b60005b949350505050565b60008281526101c4602052604081205481905b90508015610d355760008481526101916020526040902054610cf190610cea565b509092915050565b7f12b42e8a160f6064dc959c6f251e3af0750ad213dbecf573b4710d67d6c28e39610d6a81610b09612737565b610d775780610b1b612737565b8360ff811615801590610d8e575060ff8116600114155b15610db15760405163c2628c0b60e01b815260ff82166004820152602401610a76565b826101cc6000610dc188886127c4565b81526020810191909152604001600020805460ff1916911515919091179055610de8612737565b6001600160a01b0316848660ff167fd520b4ee12f45aacea6f06fd4831c1410e0d3f874a80f64424715327d7f705ca86604051610e29911515815260200190565b60405180910390a45050505050565b905090565b60008360ff811615801590610e56575060ff8116600114155b15610e795760405163c2628c0b60e01b815260ff82166004820152602401610a76565b6101ce546001600160a01b0316610ec85760405163eac0d38960e01b81526020600482015260126024820152715f7374616b696e67506172616d657465727360701b6044820152606401610a76565b6101ce5460405163882b291760e01b815260ff87166004820152602481018690526001600160a01b039091169063882b291790604401602060405180830381865afa158015610f1b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f3f9190615167565b610f5c57604051632b3dcd7960e21b815260040160405180910390fd5b6000610f66612737565b90506000610f7487876127c4565b90506000610f838888886128bf565b90965090508015610fbf57604051879060ff8a16907f8db486b45abb5d250b92d9eeea290a901f784921d6d28705cfaf2e9fa025d56390600090a35b6000610fcb8388610d01565b6101c354909150610fe7906001600160a01b031685308a6129a3565b610ff46101c48489612a0e565b61102d84848360005b6040519080825280601f01601f191660200182016040528015611027576020820181803683370190505b50612a51565b836001600160a01b0316888a60ff167f3e5900eae50f376487f0d2272f8f857884c7adde3b5be2cd8c7b1e7bdfc47e148a60405161106d91815260200190565b60405180910390a498975050505050505050565b611089612737565b6001600160a01b0316856001600160a01b031614806110af57506110af85610928612737565b6110cb5760405162461bcd60e51b8152600401610a7690615184565b6110d88585858585612b82565b5050505050565b7f664245c7af190fec316596e8231f724e8171b1966cfcd124347ac5a66c34f64a61110c81610b09612737565b6111195780610b1b612737565b611124848484612d75565b50505050565b6000610c236107f98484612e6d565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036111815760405162461bcd60e51b8152600401610a76906151d3565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166111ca60008051602061591d833981519152546001600160a01b031690565b6001600160a01b0316146111f05760405162461bcd60e51b8152600401610a769061521f565b6111f981612eb5565b6040805160008082526020820190925261121591839190612eef565b50565b60008260ff811615801590611231575060ff8116600114155b156112545760405163c2628c0b60e01b815260ff82166004820152602401610a76565b600061125e612737565b9050600061126c8686612e6d565b90506112788282610a0f565b600003611298576040516365efc92d60e01b815260040160405180910390fd5b610100811760009081526101cc602052604090205460ff16156112ce57604051632799ebef60e01b815260040160405180910390fd5b610100811760009081526101c8602090815260408083206001600160a01b0386168452825291829020825191820190925281546001600160401b031681526113159061305f565b61133257604051630f2ca6e760e01b815260040160405180910390fd5b805467ffffffffffffffff19168155826001600160a01b0316868860ff167f07e9e8a51e2cf8929e95153e486868bda03206cdbbf30ee732b17e62b43c282760405160405180910390a460006113888484610a0f565b905060006113968483610cbf565b90506113a56101c6858361308e565b6113b08585846130c7565b6101c3546113c8906001600160a01b03168683613268565b9550505050505b5092915050565b6060815183511461143b5760405162461bcd60e51b815260206004820152602960248201527f455243313135353a206163636f756e747320616e6420696473206c656e677468604482015268040dad2e6dac2e8c6d60bb1b6064820152608401610a76565b600083516001600160401b038111156114565761145661487a565b60405190808252806020026020018201604052801561147f578160200160208202803683370190505b50905060005b84518110156114f7576114ca8582815181106114a3576114a361526b565b60200260200101518583815181106114bd576114bd61526b565b6020026020010151610a0f565b8282815181106114dc576114dc61526b565b60209081029190910101526114f081615297565b9050611485565b509392505050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036115475760405162461bcd60e51b8152600401610a76906151d3565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661159060008051602061591d833981519152546001600160a01b031690565b6001600160a01b0316146115b65760405162461bcd60e51b8152600401610a769061521f565b6115bf82612eb5565b610b4f82826001612eef565b6000306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461166b5760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610a76565b5060008051602061591d8339815191525b90565b60008281526101c660205260408120548190610d14565b600082815261019160205260408120548015610cf65760008481526101c46020526040902054610cf190610cea565b60006101cc60006116d685856127c4565b815260208101919091526040016000205460ff169392505050565b8260ff811615801590611708575060ff8116600114155b1561172b5760405163c2628c0b60e01b815260ff82166004820152602401610a76565b6101c35461174b906001600160a01b0316611744612737565b30856129a3565b61176261175885856127c4565b6101c99084612a0e565b61176a612737565b6001600160a01b0316838560ff167f4bae81d7c78d9effd9b0ffe353b35b51c2695820985aa889b3d8916a017f60a0856040516117a991815260200190565b60405180910390a450505050565b600054610100900460ff16158080156117d75750600054600160ff909116105b806117f15750303b1580156117f1575060005460ff166001145b61180d5760405162461bcd60e51b8152600401610a76906152b0565b6000805460ff191660011790558015611830576000805461ff0019166101001790555b6001600160a01b0383166118745760405163eac0d38960e01b815260206004820152600a6024820152695f5f747265617375727960b01b6044820152606401610a76565b6001600160a01b0382166118bb5760405163eac0d38960e01b815260206004820152600d60248201526c2fafb9ba30b5b2b22a37b5b2b760991b6044820152606401610a76565b6118c486613298565b6118cd856133f2565b6118d5613534565b6118ed6040518060200160405280600081525061355d565b6118f5613534565b6101cd80546001600160401b0386166001600160e01b03199091168117600160401b6001600160a01b0387811691909102919091179092556101c380546001600160a01b031916928516929092179091556040519081527f63e09f16584208fba1fc7ff64c62b00f07bec177c0d97ca6689891b1e77a35c79060200160405180910390a16040516001600160a01b03841681527f3c864541ef71378c6229510ed90f376565ee42d9c5e0904a984a9e863e6db44f9060200160405180910390a180156119fb576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b600080611a1085856127c4565b9050611a1c818461358d565b95945050505050565b610b4f611a30612737565b83836135e1565b6000610c23611a4684846127c4565b60009081526101c4602052604090205490565b6060816001600160401b03811115611a7357611a7361487a565b604051908082528060200260200182016040528015611aa657816020015b6060815260200190600190039081611a915790505b50905060005b828110156113cf57611b1630858584818110611aca57611aca61526b565b9050602002810190611adc91906152fe565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506136c292505050565b828281518110611b2857611b2861526b565b60200260200101819052508080611b3e90615297565b915050611aac565b60007f8aef0597c0be1e090afba1f387ee99f604b5d975ccbed6215cdf146ffd5c49fc611b7581610b09612737565b611b825780610b1b612737565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa158015611bc9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bed919061534b565b6101c3549091506001600160a01b0390811690861603611c3f576101c554611c159082615364565b9050611c216101c75490565b611c2b9082615364565b6101ca54909150611c3c9082615364565b90505b611c4a858583613268565b604080516001600160a01b038681168252602082018490528716917fd092d7fceb5ea5a962639fcc27a7bb315e7637e699e3b108cd570c38c7584300910160405180910390a2949350505050565b6101c3546001600160a01b031663d505accf611cb2612737565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152306024820152604481018890526064810187905260ff8616608482015260a4810185905260c4810184905260e401600060405180830381600087803b158015611d1e57600080fd5b505af1158015611d32573d6000803e3d6000fd5b505050505050505050565b6000611d4b81610b09612737565b611d585780610b1b612737565b6001600160a01b038216611d9b5760405163eac0d38960e01b81526020600482015260096024820152683732bba937baba32b960b91b6044820152606401610a76565b606580546001600160a01b0319166001600160a01b0384169081179091556040517f7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc8090600090a25050565b6000611df481610b09612737565b611e015780610b1b612737565b62015180826001600160401b03161015611e4157604051625a5b2760e31b81526001600160401b0383166004820152620151806024820152604401610a76565b6276a700826001600160401b03161115611e8257604051637f16270360e01b81526001600160401b03831660048201526276a7006024820152604401610a76565b6101cd805467ffffffffffffffff19166001600160401b0384169081179091556040519081527f63e09f16584208fba1fc7ff64c62b00f07bec177c0d97ca6689891b1e77a35c7906020015b60405180910390a15050565b6000611ee881610b09612737565b611ef55780610b1b612737565b611f0f6001600160a01b038316637965db0b60e01b6136e7565b611f4d576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b6044820152606401610a76565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a25050565b60008360ff811615801590611fb1575060ff8116600114155b15611fd45760405163c2628c0b60e01b815260ff82166004820152602401610a76565b6000611fe086866127c4565b90506000611fee828661358d565b9050611ffd6101c9838361308e565b6120208561200a83613703565b60008581526101cb602052604090209190613771565b6101c354612038906001600160a01b03168683613268565b846001600160a01b0316868860ff167f8f404fb0169bf58e5bb9d571055da9ac0fc4eb973f2e6c7d1b108024381f0fc68460405161207891815260200190565b60405180910390a46001600160a01b0385163b151580156120ae57506120ae6001600160a01b038616631f0ae67f60e21b6136e7565b1561211e57604051631f0ae67f60e21b815260ff8816600482015260248101879052604481018290526001600160a01b03861690637c2b99fc90606401600060405180830381600087803b15801561210557600080fd5b505af1158015612119573d6000803e3d6000fd5b505050505b9695505050505050565b6000610cf9826103038686612e6d565b6000610c236121478484612e6d565b60009081526101c6602052604090205490565b60008360ff811615801590612173575060ff8116600114155b156121965760405163c2628c0b60e01b815260ff82166004820152602401610a76565b60006121a0612737565b905060006121ae87876127c4565b90506121ba8282610a0f565b6000036121da5760405163d7d0b56f60e01b815260040160405180910390fd5b6101cd546000906001600160401b03166121f3426137f4565b6121fd919061537b565b60008381526101c8602090815260408083206001600160a01b03881684529091529020805467ffffffffffffffff19166001600160401b038316179055905060006122518761224c8686610a0f565b61385c565b9050600061225f8483611696565b905060006122726101001986168361167f565b90506122816101c4868461308e565b6122936101c661010019871684612a0e565b61229e8686856130c7565b6122b086610100198716836000610ffd565b6040516001600160401b03851681526001600160a01b038716908b9060ff8e16907f8b14266b7a7bfb46d73cda0a8ca3a635f38e8fb4b275b68aff68cb1b5a2aea3a9060200160405180910390a450919998505050505050505050565b6000610cf98261030386866127c4565b600061232b81610b09612737565b6123385780610b1b612737565b6001600160a01b03821661237d5760405163eac0d38960e01b815260206004820152600b60248201526a6e6577547265617375727960a81b6044820152606401610a76565b6101cd805468010000000000000000600160e01b031916600160401b6001600160a01b038516908102919091179091556040519081527f3c864541ef71378c6229510ed90f376565ee42d9c5e0904a984a9e863e6db44f90602001611ece565b6123e5612737565b6001600160a01b0316856001600160a01b0316148061240b575061240b85610928612737565b6124275760405162461bcd60e51b8152600401610a7690615184565b6110d88585858585613872565b60007f12b42e8a160f6064dc959c6f251e3af0750ad213dbecf573b4710d67d6c28e3961246381610b09612737565b6124705780610b1b612737565b600061247c88886127c4565b60008181526101c4602052604081205491925061249d610100198416612147565b9050600061252d6124ae83856153a6565b6101ce60009054906101000a90046001600160a01b03166001600160a01b0316635ba971e56040518163ffffffff1660e01b8152600401602060405180830381865afa158015612502573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612526919061534b565b6064612810565b90508089111561255057604051631c67aae360e11b815260040160405180910390fd5b6000612566848b61256186836153a6565b612810565b90506000612574828c615364565b905061258081836153a6565b9a5061258f6101c4878461308e565b6125a16101c66101001988168361308e565b88156126a2576001600160a01b038a166125e95760405163eac0d38960e01b8152602060048201526008602482015267383937b837b9b2b960c11b6044820152606401610a76565b60006125f78c8b6064612810565b6101c354909150612612906001600160a01b03168c83613268565b8a6001600160a01b03168d8f60ff167fbad5bf3ab3814a2220a4737f205e42060a9b3b66422e774059048cec2f6ed0ac8460405161265291815260200190565b60405180910390a461269c6101c360009054906101000a90046001600160a01b03166101cd60089054906101000a90046001600160a01b0316838f6126979190615364565b613268565b506126c8565b6101c3546101cd546126c8916001600160a01b0390811691600160401b9004168d613268565b6126d0612737565b6001600160a01b03168c8e60ff167f59d631535a5e48c1231a728e536ce39dba8d6c7f49905ec570e3db296430e02e8e60405161270f91815260200190565b60405180910390a450989b9a5050505050505050505050565b6001600160a01b03163b151590565b6000610e386139b0565b603354604051632474521560e21b8152600481018490526001600160a01b03838116602483015260009216906391d1485490604401602060405180830381865afa158015612793573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c239190615167565b610161610b4f8282615404565b6040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff909116176101001790565b600080806000198587098587029250828110838203039150508060000361284a57838281612840576128406154c3565b0492505050610c23565b80841161285657600080fd5b60008486880960026001871981018816978890046003810283188082028403028082028403028082028403028082028403028082028403029081029092039091026000889003889004909101858311909403939093029303949094049190911702949350505050565b6101ce546040516276841960e61b815260ff8516600482015260248101849052600091829182916001600160a01b031690631da1064090604401602060405180830381865afa158015612916573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061293a919061534b565b9050806129478787611a37565b1061295a5760006001925092505061299b565b60006129668787611a37565b6129709083615364565b905061297c858261385c565b82866129888a8a611a37565b61299291906153a6565b10159350935050505b935093915050565b6040516001600160a01b03808516602483015283166044820152606481018290526111249085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526139ba565b60008281526020849052604081208054839290612a2c9084906153a6565b9250508190555080836001016000828254612a4791906153a6565b9091555050505050565b6001600160a01b038416612ab15760405162461bcd60e51b815260206004820152602160248201527f455243313135353a206d696e7420746f20746865207a65726f206164647265736044820152607360f81b6064820152608401610a76565b6000612abb612737565b90506000612ac885613a8c565b90506000612ad585613a8c565b9050612ae683600089858589613ad7565b600086815261015f602090815260408083206001600160a01b038b16845290915281208054879290612b199084906153a6565b909155505060408051878152602081018790526001600160a01b03808a1692600092918716917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4612b7983600089898989613c9b565b50505050505050565b8151835114612be45760405162461bcd60e51b815260206004820152602860248201527f455243313135353a2069647320616e6420616d6f756e7473206c656e677468206044820152670dad2e6dac2e8c6d60c31b6064820152608401610a76565b6001600160a01b038416612c0a5760405162461bcd60e51b8152600401610a76906154d9565b6000612c14612737565b9050612c24818787878787613ad7565b60005b8451811015612d0f576000858281518110612c4457612c4461526b565b602002602001015190506000858381518110612c6257612c6261526b565b602090810291909101810151600084815261015f835260408082206001600160a01b038e168352909352919091205490915081811015612cb45760405162461bcd60e51b8152600401610a769061551e565b600083815261015f602090815260408083206001600160a01b038e8116855292528083208585039055908b16825281208054849290612cf49084906153a6565b9250508190555050505080612d0890615297565b9050612c27565b50846001600160a01b0316866001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051612d5f929190615568565b60405180910390a46119fb818787878787613df6565b6040516302571be360e01b81527f91d1777781884d03a6757a803996e38de2a42967fb37eeaca72729271025a9e260048201526001600160a01b038416906302571be390602401602060405180830381865afa158015612dd9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612dfd919061558d565b6001600160a01b031663c47f002783836040518363ffffffff1660e01b8152600401612e2a9291906155aa565b6020604051808303816000875af1158015612e49573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611124919061534b565b6040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff9091161790565b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3612ee281610b09612737565b610b4f5780610b1b612737565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615612f2757612f2283613eb1565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612f81575060408051601f3d908101601f19168201909252612f7e9181019061534b565b60015b612fe45760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610a76565b60008051602061591d83398151915281146130535760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610a76565b50612f22838383613f4d565b600061307482516001600160401b0316151590565b8015610aa557505051426001600160401b03909116111590565b600082815260208490526040812080548392906130ac908490615364565b9250508190555080836001016000828254612a479190615364565b6001600160a01b0383166131295760405162461bcd60e51b815260206004820152602360248201527f455243313135353a206275726e2066726f6d20746865207a65726f206164647260448201526265737360e81b6064820152608401610a76565b6000613133612737565b9050600061314084613a8c565b9050600061314d84613a8c565b905061316d83876000858560405180602001604052806000815250613ad7565b600085815261015f602090815260408083206001600160a01b038a168452909152902054848110156131ed5760405162461bcd60e51b8152602060048201526024808201527f455243313135353a206275726e20616d6f756e7420657863656564732062616c604482015263616e636560e01b6064820152608401610a76565b600086815261015f602090815260408083206001600160a01b038b81168086529184528285208a8703905582518b81529384018a90529092908816917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4604080516020810190915260009052612b79565b6040516001600160a01b038316602482015260448101829052612f2290849063a9059cbb60e01b906064016129d7565b600054610100900460ff16158080156132b85750600054600160ff909116105b806132d25750303b1580156132d2575060005460ff166001145b6132ee5760405162461bcd60e51b8152600401610a76906152b0565b6000805460ff191660011790558015613311576000805461ff0019166101001790555b61332b6001600160a01b038316637965db0b60e01b6136e7565b613369576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b6044820152606401610a76565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a28015610b4f576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001611ece565b600054610100900460ff16158080156134125750600054600160ff909116105b8061342c5750303b15801561342c575060005460ff166001145b6134485760405162461bcd60e51b8152600401610a76906152b0565b6000805460ff19166001179055801561346b576000805461ff0019166101001790555b6001600160a01b0382166134ab5760405163eac0d38960e01b81526020600482015260066024820152653937baba32b960d11b6044820152606401610a76565b606580546001600160a01b0319166001600160a01b0384169081179091556040517f7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc8090600090a28015610b4f576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001611ece565b600054610100900460ff1661355b5760405162461bcd60e51b8152600401610a76906155d9565b565b600054610100900460ff166135845760405162461bcd60e51b8152600401610a76906155d9565b61121581613f72565b60008281526101cb602090815260408083206001600160a01b0385168452909152812054610c23906135d26135cd866135c68789610a0f565b6000613fa2565b613703565b6135dc9190615624565b613fe1565b816001600160a01b0316836001600160a01b0316036136545760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2073657474696e6720617070726f76616c20737461747573604482015268103337b91039b2b63360b91b6064820152608401610a76565b6001600160a01b0383811660008181526101606020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6060610c23838360405180606001604052806027815260200161593d60279139614033565b60006136f2836140c7565b8015610c235750610c2383836140fa565b60006001600160ff1b0382111561376d5760405162461bcd60e51b815260206004820152602860248201527f53616665436173743a2076616c756520646f65736e27742066697420696e2061604482015267371034b73a191a9b60c11b6064820152608401610a76565b5090565b6001600160a01b0382166137b15760405163eac0d38960e01b8152600401610a76906020808252600490820152631b5a5b9d60e21b604082015260600190565b6001600160a01b038216600090815260208490526040812080548392906137d9908490615663565b9250508190555080836001016000828254612a479190615663565b60006001600160401b0382111561376d5760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203660448201526534206269747360d01b6064820152608401610a76565b600081831061386b5781610c23565b5090919050565b6001600160a01b0384166138985760405162461bcd60e51b8152600401610a76906154d9565b60006138a2612737565b905060006138af85613a8c565b905060006138bc85613a8c565b90506138cc838989858589613ad7565b600086815261015f602090815260408083206001600160a01b038c168452909152902054858110156139105760405162461bcd60e51b8152600401610a769061551e565b600087815261015f602090815260408083206001600160a01b038d8116855292528083208985039055908a168252812080548892906139509084906153a6565b909155505060408051888152602081018890526001600160a01b03808b16928c821692918816917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4611d32848a8a8a8a8a613c9b565b6000610e386141d9565b6000613a0f826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661423b9092919063ffffffff16565b805190915015612f225780806020019051810190613a2d9190615167565b612f225760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610a76565b60408051600180825281830190925260609160009190602080830190803683370190505090508281600081518110613ac657613ac661526b565b602090810291909101015292915050565b60005b8351811015613c8c57613b0a848281518110613af857613af861526b565b60200260200101516101009081161490565b15613c40576000613b526135cd868481518110613b2957613b2961526b565b6020026020010151868581518110613b4357613b4361526b565b60200260200101516001613fa2565b90506001600160a01b038716613ba857613ba386826101cb6000898781518110613b7e57613b7e61526b565b602002602001015181526020019081526020016000206137719092919063ffffffff16565b613c3a565b6001600160a01b038616613bf757613ba387826101cb6000898781518110613bd257613bd261526b565b6020026020010151815260200190815260200160002061424a9092919063ffffffff16565b613c3a8787836101cb60008a8881518110613c1457613c1461526b565b602002602001015181526020019081526020016000206142cd909392919063ffffffff16565b50613c7a565b6001600160a01b0386161580613c5d57506001600160a01b038516155b613c7a5760405163068331f960e01b815260040160405180910390fd5b80613c8481615297565b915050613ada565b506119fb8686868686866143a9565b6001600160a01b0384163b156119fb5760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e6190613cdf90899089908890889088906004016156a4565b6020604051808303816000875af1925050508015613d1a575060408051601f3d908101601f19168201909252613d17918101906156de565b60015b613dc657613d266156fb565b806308c379a003613d5f5750613d3a615716565b80613d455750613d61565b8060405162461bcd60e51b8152600401610a769190614a21565b505b60405162461bcd60e51b815260206004820152603460248201527f455243313135353a207472616e7366657220746f206e6f6e20455243313135356044820152732932b1b2b4bb32b91034b6b83632b6b2b73a32b960611b6064820152608401610a76565b6001600160e01b0319811663f23a6e6160e01b14612b795760405162461bcd60e51b8152600401610a769061579f565b6001600160a01b0384163b156119fb5760405163bc197c8160e01b81526001600160a01b0385169063bc197c8190613e3a90899089908890889088906004016157e7565b6020604051808303816000875af1925050508015613e75575060408051601f3d908101601f19168201909252613e72918101906156de565b60015b613e8157613d266156fb565b6001600160e01b0319811663bc197c8160e01b14612b795760405162461bcd60e51b8152600401610a769061579f565b6001600160a01b0381163b613f1e5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610a76565b60008051602061591d83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b613f5683614525565b600082511180613f635750805b15612f22576111248383614565565b600054610100900460ff16613f995760405162461bcd60e51b8152600401610a76906155d9565b611215816127b7565b60008381526101916020526040812054600084118015613fc25750600081115b613fcd576000611a1c565b611a1c613fd986614611565b85838661464b565b60008082121561376d5760405162461bcd60e51b815260206004820181905260248201527f53616665436173743a2076616c7565206d75737420626520706f7369746976656044820152606401610a76565b60606001600160a01b0384163b61405c5760405162461bcd60e51b8152600401610a7690615845565b600080856001600160a01b031685604051614077919061588b565b600060405180830381855af49150503d80600081146140b2576040519150601f19603f3d011682016040523d82523d6000602084013e6140b7565b606091505b509150915061211e82828661469c565b60006140da826301ffc9a760e01b6140fa565b8015610aa557506140f3826001600160e01b03196140fa565b1592915050565b604080516001600160e01b0319831660248083019190915282518083039091018152604490910182526020810180516001600160e01b03166301ffc9a760e01b179052905160009190829081906001600160a01b038716906175309061416190869061588b565b6000604051808303818686fa925050503d806000811461419d576040519150601f19603f3d011682016040523d82523d6000602084013e6141a2565b606091505b50915091506020815110156141bd5760009350505050610aa5565b81801561211e57508080602001905181019061211e9190615167565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361423657600036614219601482615364565b614225923692906158a7565b61422e916158d1565b60601c905090565b503390565b6060610cf984846000856146d5565b6001600160a01b03821661428a5760405163eac0d38960e01b8152600401610a7690602080825260049082015263313ab93760e11b604082015260600190565b6001600160a01b038216600090815260208490526040812080548392906142b2908490615624565b9250508190555080836001016000828254612a479190615624565b6001600160a01b03831661430d5760405163eac0d38960e01b8152600401610a769060208082526004908201526366726f6d60e01b604082015260600190565b6001600160a01b0382166143495760405163eac0d38960e01b8152602060048201526002602482015261746f60f01b6044820152606401610a76565b6001600160a01b03831660009081526020859052604081208054839290614371908490615624565b90915550506001600160a01b0382166000908152602085905260408120805483929061439e908490615663565b909155505050505050565b6001600160a01b0385166144315760005b835181101561442f578281815181106143d5576143d561526b565b602002602001015161019160008684815181106143f4576143f461526b565b60200260200101518152602001908152602001600020600082825461441991906153a6565b90915550614428905081615297565b90506143ba565b505b6001600160a01b0384166119fb5760005b8351811015612b7957600084828151811061445f5761445f61526b565b60200260200101519050600084838151811061447d5761447d61526b565b6020026020010151905060006101916000848152602001908152602001600020549050818110156145015760405162461bcd60e51b815260206004820152602860248201527f455243313135353a206275726e20616d6f756e74206578636565647320746f74604482015267616c537570706c7960c01b6064820152608401610a76565b600092835261019160205260409092209103905561451e81615297565b9050614442565b61452e81613eb1565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606001600160a01b0383163b61458e5760405162461bcd60e51b8152600401610a7690615845565b600080846001600160a01b0316846040516145a9919061588b565b600060405180830381855af49150503d80600081146145e4576040519150601f19603f3d011682016040523d82523d6000602084013e6145e9565b606091505b5091509150611a1c828260405180606001604052806027815260200161593d6027913961469c565b60008181526101cb6020526040812060010154610aa59060008481526101c9602052604090205461464190613703565b6135dc9190615663565b600080614659868686612810565b9050600183600281111561466f5761466f615906565b14801561468c575060008480614687576146876154c3565b868809115b15611a1c5761211e6001826153a6565b606083156146ab575081610c23565b8251156146bb5782518084602001fd5b8160405162461bcd60e51b8152600401610a769190614a21565b6060824710156147365760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610a76565b6001600160a01b0385163b61478d5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610a76565b600080866001600160a01b031685876040516147a9919061588b565b60006040518083038185875af1925050503d80600081146147e6576040519150601f19603f3d011682016040523d82523d6000602084013e6147eb565b606091505b50915091506147fb82828661469c565b979650505050505050565b6001600160a01b038116811461121557600080fd5b6000806040838503121561482e57600080fd5b823561483981614806565b946020939093013593505050565b6001600160e01b03198116811461121557600080fd5b60006020828403121561486f57600080fd5b8135610c2381614847565b634e487b7160e01b600052604160045260246000fd5b601f8201601f191681016001600160401b03811182821017156148b5576148b561487a565b6040525050565b60006001600160401b038311156148d5576148d561487a565b6040516148ec601f8501601f191660200182614890565b80915083815284848401111561490157600080fd5b83836020830137600060208583010152509392505050565b60006020828403121561492b57600080fd5b81356001600160401b0381111561494157600080fd5b8201601f8101841361495257600080fd5b610cf9848235602084016148bc565b60006020828403121561497357600080fd5b8135610c2381614806565b803560ff8116811461498f57600080fd5b919050565b600080604083850312156149a757600080fd5b6148398361497e565b6000602082840312156149c257600080fd5b5035919050565b60005b838110156149e45781810151838201526020016149cc565b838111156111245750506000910152565b60008151808452614a0d8160208601602086016149c9565b601f01601f19169290920160200192915050565b602081526000610c2360208301846149f5565b60008060408385031215614a4757600080fd5b50508035926020909101359150565b801515811461121557600080fd5b600080600060608486031215614a7957600080fd5b614a828461497e565b9250602084013591506040840135614a9981614a56565b809150509250925092565b600080600060608486031215614ab957600080fd5b614ac28461497e565b95602085013595506040909401359392505050565b60006001600160401b03821115614af057614af061487a565b5060051b60200190565b600082601f830112614b0b57600080fd5b81356020614b1882614ad7565b604051614b258282614890565b83815260059390931b8501820192828101915086841115614b4557600080fd5b8286015b84811015614b605780358352918301918301614b49565b509695505050505050565b600082601f830112614b7c57600080fd5b610c23838335602085016148bc565b600080600080600060a08688031215614ba357600080fd5b8535614bae81614806565b94506020860135614bbe81614806565b935060408601356001600160401b0380821115614bda57600080fd5b614be689838a01614afa565b94506060880135915080821115614bfc57600080fd5b614c0889838a01614afa565b93506080880135915080821115614c1e57600080fd5b50614c2b88828901614b6b565b9150509295509295909350565b600080600060408486031215614c4d57600080fd5b8335614c5881614806565b925060208401356001600160401b0380821115614c7457600080fd5b818601915086601f830112614c8857600080fd5b813581811115614c9757600080fd5b876020828501011115614ca957600080fd5b6020830194508093505050509250925092565b60008060408385031215614ccf57600080fd5b82356001600160401b0380821115614ce657600080fd5b818501915085601f830112614cfa57600080fd5b81356020614d0782614ad7565b604051614d148282614890565b83815260059390931b8501820192828101915089841115614d3457600080fd5b948201945b83861015614d5b578535614d4c81614806565b82529482019490820190614d39565b96505086013592505080821115614d7157600080fd5b50614d7e85828601614afa565b9150509250929050565b600081518084526020808501945080840160005b83811015614db857815187529582019590820190600101614d9c565b509495945050505050565b602081526000610c236020830184614d88565b60008060408385031215614de957600080fd5b8235614df481614806565b915060208301356001600160401b03811115614e0f57600080fd5b614d7e85828601614b6b565b80356001600160401b038116811461498f57600080fd5b600080600080600060a08688031215614e4a57600080fd5b8535614e5581614806565b94506020860135614e6581614806565b9350614e7360408701614e1b565b92506060860135614e8381614806565b91506080860135614e9381614806565b809150509295509295909350565b600080600060608486031215614eb657600080fd5b614ebf8461497e565b9250602084013591506040840135614a9981614806565b60008060408385031215614ee957600080fd5b8235614ef481614806565b91506020830135614f0481614a56565b809150509250929050565b60008060208385031215614f2257600080fd5b82356001600160401b0380821115614f3957600080fd5b818501915085601f830112614f4d57600080fd5b813581811115614f5c57600080fd5b8660208260051b8501011115614f7157600080fd5b60209290920196919550909350505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015614fd857603f19888603018452614fc68583516149f5565b94509285019290850190600101614faa565b5092979650505050505050565b60008060408385031215614ff857600080fd5b823561500381614806565b91506020830135614f0481614806565b600080600080600060a0868803121561502b57600080fd5b85359450602086013593506150426040870161497e565b94979396509394606081013594506080013592915050565b60006020828403121561506c57600080fd5b610c2382614e1b565b600080600080600060a0868803121561508d57600080fd5b853561509881614806565b945060208601356150a881614806565b9350604086013592506060860135915060808601356001600160401b038111156150d157600080fd5b614c2b88828901614b6b565b600080600080600060a086880312156150f557600080fd5b6150fe8661497e565b94506020860135935060408601359250606086013561511c81614806565b949793965091946080013592915050565b600181811c9082168061514157607f821691505b60208210810361516157634e487b7160e01b600052602260045260246000fd5b50919050565b60006020828403121561517957600080fd5b8151610c2381614a56565b6020808252602f908201527f455243313135353a2063616c6c6572206973206e6f7420746f6b656e206f776e60408201526e195c881b9bdc88185c1c1c9bdd9959608a1b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600182016152a9576152a9615281565b5060010190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b6000808335601e1984360301811261531557600080fd5b8301803591506001600160401b0382111561532f57600080fd5b60200191503681900382131561534457600080fd5b9250929050565b60006020828403121561535d57600080fd5b5051919050565b60008282101561537657615376615281565b500390565b60006001600160401b0380831681851680830382111561539d5761539d615281565b01949350505050565b600082198211156153b9576153b9615281565b500190565b601f821115612f2257600081815260208120601f850160051c810160208610156153e55750805b601f850160051c820191505b818110156119fb578281556001016153f1565b81516001600160401b0381111561541d5761541d61487a565b6154318161542b845461512d565b846153be565b602080601f831160018114615466576000841561544e5750858301515b600019600386901b1c1916600185901b1785556119fb565b600085815260208120601f198616915b8281101561549557888601518255948401946001909101908401615476565b50858210156154b35787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052601260045260246000fd5b60208082526025908201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604082015264647265737360d81b606082015260800190565b6020808252602a908201527f455243313135353a20696e73756666696369656e742062616c616e636520666f60408201526939103a3930b739b332b960b11b606082015260800190565b60408152600061557b6040830185614d88565b8281036020840152611a1c8185614d88565b60006020828403121561559f57600080fd5b8151610c2381614806565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b60008083128015600160ff1b85018412161561564257615642615281565b6001600160ff1b038401831381161561565d5761565d615281565b50500390565b600080821280156001600160ff1b038490038513161561568557615685615281565b600160ff1b839003841281161561569e5761569e615281565b50500190565b6001600160a01b03868116825285166020820152604081018490526060810183905260a0608082018190526000906147fb908301846149f5565b6000602082840312156156f057600080fd5b8151610c2381614847565b600060033d111561167c5760046000803e5060005160e01c90565b600060443d10156157245790565b6040516003193d81016004833e81513d6001600160401b03816024840111818411171561575357505050505090565b828501915081518181111561576b5750505050505090565b843d87010160208285010111156157855750505050505090565b61579460208286010187614890565b509095945050505050565b60208082526028908201527f455243313135353a204552433131353552656365697665722072656a656374656040820152676420746f6b656e7360c01b606082015260800190565b6001600160a01b0386811682528516602082015260a06040820181905260009061581390830186614d88565b82810360608401526158258186614d88565b9050828103608084015261583981856149f5565b98975050505050505050565b60208082526026908201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6040820152651b9d1c9858dd60d21b606082015260800190565b6000825161589d8184602087016149c9565b9190910192915050565b600080858511156158b757600080fd5b838611156158c457600080fd5b5050820193919092039150565b6bffffffffffffffffffffffff1981358181169160148510156158fe5780818660140360031b1b83161692505b505092915050565b634e487b7160e01b600052602160045260246000fdfe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220c5b546099bc6b6cf58d1c2622c5bc754e5ea2de44ef353e0bf05aa15e53ca85864736f6c634300080f0033",
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

// Initialize is a paid mutator transaction binding the contract method 0x89a0ca6b.
//
// Solidity: function initialize(address __manager, address __router, uint64 __withdrawalDelay, address __treasury, address __stakedToken) returns()
func (_FortaStaking *FortaStakingTransactor) Initialize(opts *bind.TransactOpts, __manager common.Address, __router common.Address, __withdrawalDelay uint64, __treasury common.Address, __stakedToken common.Address) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "initialize", __manager, __router, __withdrawalDelay, __treasury, __stakedToken)
}

// Initialize is a paid mutator transaction binding the contract method 0x89a0ca6b.
//
// Solidity: function initialize(address __manager, address __router, uint64 __withdrawalDelay, address __treasury, address __stakedToken) returns()
func (_FortaStaking *FortaStakingSession) Initialize(__manager common.Address, __router common.Address, __withdrawalDelay uint64, __treasury common.Address, __stakedToken common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.Initialize(&_FortaStaking.TransactOpts, __manager, __router, __withdrawalDelay, __treasury, __stakedToken)
}

// Initialize is a paid mutator transaction binding the contract method 0x89a0ca6b.
//
// Solidity: function initialize(address __manager, address __router, uint64 __withdrawalDelay, address __treasury, address __stakedToken) returns()
func (_FortaStaking *FortaStakingTransactorSession) Initialize(__manager common.Address, __router common.Address, __withdrawalDelay uint64, __treasury common.Address, __stakedToken common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.Initialize(&_FortaStaking.TransactOpts, __manager, __router, __withdrawalDelay, __treasury, __stakedToken)
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
