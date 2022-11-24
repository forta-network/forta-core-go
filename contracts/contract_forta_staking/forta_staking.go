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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"AccessManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newWithdrawalDelay\",\"type\":\"uint256\"}],\"name\":\"DelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isFrozen\",\"type\":\"bool\"}],\"name\":\"Froze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"MaxStakeReached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"SlashDelegatorsPercentSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SlashedShareSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"subjectGateway\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allocator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardsDistributor\",\"type\":\"address\"}],\"name\":\"StakeHelpersConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensSwept\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"TreasurySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WithdrawalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_SLASHABLE_PERCENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WITHDRAWAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_WITHDRAWAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"activeSharesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"activeSharesToStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"activeStakeFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStakeSubjectGateway\",\"name\":\"__subjectGateway\",\"type\":\"address\"},{\"internalType\":\"contractIStakeAllocator\",\"name\":\"__allocator\",\"type\":\"address\"},{\"internalType\":\"contractIRewardsDistributor\",\"name\":\"__distributor\",\"type\":\"address\"}],\"name\":\"configureStakeHelpers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeValue\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"freeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"}],\"name\":\"getDelegatedSubjectType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"}],\"name\":\"getDelegatorSubjectType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"}],\"name\":\"getSubjectTypeAgency\",\"outputs\":[{\"internalType\":\"enumSubjectTypeValidator.SubjectStakeAgency\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"inactiveSharesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"inactiveSharesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"inactiveSharesToStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"inactiveStakeFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__manager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"__stakedToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"__withdrawalDelay\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"__treasury\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sharesValue\",\"type\":\"uint256\"}],\"name\":\"initiateWithdrawal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"isFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"oldSubjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"oldSubject\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"newSubjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"newSubject\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"relayPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"setAccessManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newDelay\",\"type\":\"uint64\"}],\"name\":\"setDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ensRegistry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ensName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"setSlashDelegatorsPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"sharesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"proposerPercent\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashDelegatorsPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"activeSharesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stakeToActiveShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"inactiveSharesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stakeToInactiveShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subjectGateway\",\"outputs\":[{\"internalType\":\"contractIStakeSubjectGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweep\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalActiveStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"totalInactiveShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInactiveStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"subjectType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523060a0523480156200001557600080fd5b5060405162005cf338038062005cf3833981016040819052620000389162000180565b6001600160a01b038116608052600054610100900460ff1615808015620000665750600054600160ff909116105b8062000096575062000083306200017160201b62002a531760201c565b15801562000096575060005460ff166001145b620000fe5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b6000805460ff19166001179055801562000122576000805461ff0019166101001790555b801562000169576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050620001b2565b6001600160a01b03163b151590565b6000602082840312156200019357600080fd5b81516001600160a01b0381168114620001ab57600080fd5b9392505050565b60805160a051615af8620001fb600039600081816114d00152818161151001528181611aed01528181611b2d0152611d1f01526000818161074201526144e50152615af86000f3fe60806040526004361061036a5760003560e01c806361d027b3116101c6578063c07975ad116100f7578063dc4653ef11610095578063f08d35ee1161006f578063f08d35ee14610abb578063f0f4426014610adb578063f242432a14610afb578063f8058b0614610b1b57600080fd5b8063dc4653ef14610a19578063e985e9c514610a39578063edea0bac14610a8357600080fd5b8063c9580804116100d1578063c9580804146109a3578063cc7a262e146109c3578063d858a7e5146109e4578063da5bfb94146109f957600080fd5b8063c07975ad14610943578063c107330214610963578063c133a5621461098357600080fd5b8063a22cb46511610164578063ac9650d81161013e578063ac9650d8146108a8578063b4d60fb3146108d5578063b8dc491b146108f5578063bd85b0391461091557600080fd5b8063a22cb46514610851578063a238f9df14610871578063a290bf381461088857600080fd5b806375e130ad116101a057806375e130ad146107d8578063762fa7b7146107f85780638acdea4d1461082557806390f1ccba1461083a57600080fd5b806361d027b31461077257806362772b141461079857806364a0f901146107b857600080fd5b80633121db1c116102a05780634e1273f41161023e5780635128e693116102185780635128e693146106bf57806352d1902d146106df57806354fd4d50146106f4578063572b6c051461072557600080fd5b80634e1273f41461064f5780634f1ef2861461067c5780634f558e791461068f57600080fd5b8063371f42261161027a578063371f4226146105e25780633f489914146105f85780634a5f2b5d146106185780634e04275f1461062f57600080fd5b80633121db1c14610582578063321ebc54146105a25780633659cfe6146105c257600080fd5b80631a82ef4e1161030d578063226cc300116102e7578063226cc3001461050c57806328f731481461052c5780632cb31144146105425780632eb2c2d61461056257600080fd5b80631a82ef4e146104815780631daa0445146104a15780631fd5a956146104d357600080fd5b8063093af6a911610349578063093af6a9146103f45780630e10103f146104145780630e89341c14610434578063115a90ee1461046157600080fd5b8062fdd58e1461036f57806301ffc9a7146103a257806302fe5305146103d2575b600080fd5b34801561037b57600080fd5b5061038f61038a3660046149d3565b610b3b565b6040519081526020015b60405180910390f35b3480156103ae57600080fd5b506103c26103bd366004614a15565b610bd7565b6040519015158152602001610399565b3480156103de57600080fd5b506103f26103ed366004614ad1565b610c27565b005b34801561040057600080fd5b506103f261040f366004614b2f565b610c7f565b34801561042057600080fd5b5061038f61042f366004614b8a565b610e97565b34801561044057600080fd5b5061045461044f366004614ba6565b610ead565b6040516103999190614c17565b34801561046d57600080fd5b5061038f61047c366004614c2a565b610f42565b34801561048d57600080fd5b5061038f61049c366004614c2a565b610f84565b3480156104ad57600080fd5b506104c16104bc366004614c4c565b610fc0565b60405160ff9091168152602001610399565b3480156104df57600080fd5b506101ce546104f4906001600160a01b031681565b6040516001600160a01b039091168152602001610399565b34801561051857600080fd5b506103f2610527366004614c75565b610fde565b34801561053857600080fd5b506101c55461038f565b34801561054e57600080fd5b5061038f61055d366004614cb5565b6110fe565b34801561056e57600080fd5b506103f261057d366004614d9c565b61140d565b34801561058e57600080fd5b506103f261059d366004614e49565b61146b565b3480156105ae57600080fd5b5061038f6105bd366004614b8a565b6114b6565b3480156105ce57600080fd5b506103f26105dd366004614ecd565b6114c5565b3480156105ee57600080fd5b506101c75461038f565b34801561060457600080fd5b5061038f610613366004614b8a565b6115a5565b34801561062457600080fd5b5061038f6201518081565b34801561063b57600080fd5b506103f261064a366004614f01565b611780565b34801561065b57600080fd5b5061066f61066a366004614f5b565b6119b9565b6040516103999190615062565b6103f261068a366004615075565b611ae2565b34801561069b57600080fd5b506103c26106aa366004614ba6565b60009081526101916020526040902054151590565b3480156106cb57600080fd5b506103f26106da3660046150ba565b611baf565b3480156106eb57600080fd5b5061038f611d12565b34801561070057600080fd5b5061045460405180604001604052806005815260200164302e312e3160d81b81525081565b34801561073157600080fd5b506103c2610740366004614ecd565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0390811691161490565b34801561077e57600080fd5b506101cd54600160401b90046001600160a01b03166104f4565b3480156107a457600080fd5b5061038f6107b3366004614c2a565b611dc6565b3480156107c457600080fd5b5061038f6107d3366004614c2a565b611ddd565b3480156107e457600080fd5b506103c26107f3366004614b8a565b611e0c565b34801561080457600080fd5b50610818610813366004614c4c565b611e38565b6040516103999190615132565b34801561083157600080fd5b5061038f605a81565b34801561084657600080fd5b5061038f6101cf5481565b34801561085d57600080fd5b506103f261086c366004615140565b611e8e565b34801561087d57600080fd5b5061038f6276a70081565b34801561089457600080fd5b5061038f6108a3366004614b8a565b611ea0565b3480156108b457600080fd5b506108c86108c3366004615179565b611ec2565b60405161039991906151ed565b3480156108e157600080fd5b506103f26108f0366004614ba6565b611faf565b34801561090157600080fd5b5061038f61091036600461524f565b612027565b34801561092157600080fd5b5061038f610930366004614ba6565b6000908152610191602052604090205490565b34801561094f57600080fd5b506103f261095e36600461527d565b612177565b34801561096f57600080fd5b506103f261097e3660046152c4565b61221c565b34801561098f57600080fd5b506104c161099e366004614c4c565b612308565b3480156109af57600080fd5b506103f26109be366004614ecd565b61231e565b3480156109cf57600080fd5b506101c3546104f4906001600160a01b031681565b3480156109f057600080fd5b506103f26123dc565b348015610a0557600080fd5b5061038f610a143660046152df565b612467565b348015610a2557600080fd5b5061038f610a34366004614b8a565b612477565b348015610a4557600080fd5b506103c2610a5436600461524f565b6001600160a01b0391821660009081526101606020908152604080832093909416825291909152205460ff1690565b348015610a8f57600080fd5b50610aa3610a9e366004614cb5565b612499565b6040516001600160401b039091168152602001610399565b348015610ac757600080fd5b5061038f610ad63660046152df565b61272c565b348015610ae757600080fd5b506103f2610af6366004614ecd565b61273c565b348015610b0757600080fd5b506103f2610b16366004615314565b6127fc565b348015610b2757600080fd5b5061038f610b3636600461537c565b612853565b60006001600160a01b038316610bab5760405162461bcd60e51b815260206004820152602a60248201527f455243313135353a2061646472657373207a65726f206973206e6f742061207660448201526930b634b21037bbb732b960b11b60648201526084015b60405180910390fd5b50600081815261015f602090815260408083206001600160a01b03861684529091529020545b92915050565b60006001600160e01b03198216636cdb3d1360e11b1480610c0857506001600160e01b031982166303a24d0760e21b145b80610bd157506301ffc9a760e01b6001600160e01b0319831614610bd1565b6000610c3a81610c35612a62565b612a6c565b610c725780610c47612a62565b6040516301d4003760e61b815260048101929092526001600160a01b03166024820152604401610ba2565b610c7b82612af1565b5050565b7fcbe0462e67cb804f0011a6c0b71e9ff49be70d0f907ffdf4f3c74499282ab0b1610cac81610c35612a62565b610cb95780610c47612a62565b60ff861615610ce05760405163c2628c0b60e01b815260ff87166004820152602401610ba2565b60ff8416600214610d095760405163c2628c0b60e01b815260ff85166004820152602401610ba2565b610d138686611e0c565b15610d3157604051632799ebef60e01b815260040160405180910390fd5b6000610d3d8787612b05565b90506000610d4b8483610b3b565b90506000610d598383611ddd565b90506000610d678888612b05565b90506000610d758284610f84565b9050610d846101c48685612b51565b610d916101c48385612b94565b610d9c878686612bcd565b610dd587838360005b6040519080825280601f01601f191660200182016040528015610dcf576020820181803683370190505b50612d73565b866001600160a01b0316888a60ff167f3e5900eae50f376487f0d2272f8f857884c7adde3b5be2cd8c7b1e7bdfc47e1486604051610e1591815260200190565b60405180910390a46101d054604051636fda184760e11b81526001600160a01b039091169063dfb4308e90610e589085908d908d908d908a9089906004016153cc565b600060405180830381600087803b158015610e7257600080fd5b505af1158015610e86573d6000803e3d6000fd5b505050505050505050505050505050565b6000610ea66109308484612b05565b9392505050565b60606101618054610ebd90615400565b80601f0160208091040260200160405190810160405280929190818152602001828054610ee990615400565b8015610f365780601f10610f0b57610100808354040283529160200191610f36565b820191906000526020600020905b815481529060010190602001808311610f1957829003601f168201915b50505050509050919050565b600082815261019160205260408120548015610f795760008481526101c66020526040902054610f74905b8483612e9b565b610f7c565b60005b949350505050565b60008281526101c4602052604081205481905b90508015610fb85760008481526101916020526040902054610f7490610f6d565b509092915050565b600060ff821660031415610fd657506002919050565b5060ff919050565b7f12b42e8a160f6064dc959c6f251e3af0750ad213dbecf573b4710d67d6c28e3961100b81610c35612a62565b6110185780610c47612a62565b8360ff81161580159061102f575060ff8116600114155b801561103f575060ff8116600214155b801561104f575060ff8116600314155b156110725760405163c2628c0b60e01b815260ff82166004820152602401610ba2565b826101cc60006110828888612b05565b81526020810191909152604001600020805460ff19169115159190911790556110a9612a62565b6001600160a01b0316848660ff167fd520b4ee12f45aacea6f06fd4831c1410e0d3f874a80f64424715327d7f705ca866040516110ea911515815260200190565b60405180910390a45050505050565b905090565b60008360ff811615801590611117575060ff8116600114155b8015611127575060ff8116600214155b8015611137575060ff8116600314155b1561115a5760405163c2628c0b60e01b815260ff82166004820152602401610ba2565b8460048061116783611e38565b6004811115611178576111786150fa565b14156111a5578161118883611e38565b826040516365f3939760e01b8152600401610ba29392919061543b565b6101ce546001600160a01b03166111f05760405163eac0d38960e01b815260206004820152600e60248201526d7375626a6563744761746577617960901b6044820152606401610ba2565b6101ce5460405163882b291760e01b815260ff89166004820152602481018890526001600160a01b039091169063882b29179060440160206040518083038186803b15801561123e57600080fd5b505afa158015611252573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611276919061545f565b61129357604051632b3dcd7960e21b815260040160405180910390fd5b600061129d612a62565b905060006112ab8989612b05565b905060006112ba8a8a8a612f4b565b909850905080156112f657604051899060ff8c16907f8db486b45abb5d250b92d9eeea290a901f784921d6d28705cfaf2e9fa025d56390600090a35b6000611302838a610f84565b6101c35490915061131e906001600160a01b031685308c61303e565b61132b6101c4848b612b94565b6113388484836000610da5565b836001600160a01b03168a8c60ff167f3e5900eae50f376487f0d2272f8f857884c7adde3b5be2cd8c7b1e7bdfc47e148c60405161137891815260200190565b60405180910390a46101d060009054906101000a90046001600160a01b03166001600160a01b031663dfb4308e848d8d888e876040518763ffffffff1660e01b81526004016113cc969594939291906153cc565b600060405180830381600087803b1580156113e657600080fd5b505af11580156113fa573d6000803e3d6000fd5b50929d9c50505050505050505050505050565b611415612a62565b6001600160a01b0316856001600160a01b0316148061143b575061143b85610a54612a62565b6114575760405162461bcd60e51b8152600401610ba29061547c565b61146485858585856130a9565b5050505050565b7f664245c7af190fec316596e8231f724e8171b1966cfcd124347ac5a66c34f64a61149881610c35612a62565b6114a55780610c47612a62565b6114b08484846132a4565b50505050565b6000610ea661093084846133ba565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016141561150e5760405162461bcd60e51b8152600401610ba2906154cb565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316611557600080516020615a7c833981519152546001600160a01b031690565b6001600160a01b03161461157d5760405162461bcd60e51b8152600401610ba290615517565b61158681613402565b604080516000808252602082019092526115a29183919061343c565b50565b60008260ff8116158015906115be575060ff8116600114155b80156115ce575060ff8116600214155b80156115de575060ff8116600314155b156116015760405163c2628c0b60e01b815260ff82166004820152602401610ba2565b600061160b612a62565b9050600061161986866133ba565b90506116258282610b3b565b611642576040516365efc92d60e01b815260040160405180910390fd5b610100811760009081526101cc602052604090205460ff161561167857604051632799ebef60e01b815260040160405180910390fd5b610100811760009081526101c8602090815260408083206001600160a01b0386168452825291829020825191820190925281546001600160401b031681526116bf906135bb565b6116dc57604051630f2ca6e760e01b815260040160405180910390fd5b805467ffffffffffffffff19168155826001600160a01b0316868860ff167f07e9e8a51e2cf8929e95153e486868bda03206cdbbf30ee732b17e62b43c282760405160405180910390a460006117328484610b3b565b905060006117408483610f42565b905061174f6101c68583612b51565b61175a858584612bcd565b6101c354611772906001600160a01b031686836135ea565b9550505050505b5092915050565b600054610100900460ff16158080156117a05750600054600160ff909116105b806117ba5750303b1580156117ba575060005460ff166001145b6117d65760405162461bcd60e51b8152600401610ba290615563565b6000805460ff1916600117905580156117f9576000805461ff0019166101001790555b6001600160a01b03821661183d5760405163eac0d38960e01b815260206004820152600a6024820152695f5f747265617375727960b01b6044820152606401610ba2565b6001600160a01b0384166118845760405163eac0d38960e01b815260206004820152600d60248201526c2fafb9ba30b5b2b22a37b5b2b760991b6044820152606401610ba2565b61188d8561361a565b6118a5604051806020016040528060008152506136e6565b6118ad613716565b6101cd80546001600160401b0385166001600160e01b03199091168117600160401b6001600160a01b0386811691909102919091179092556101c380546001600160a01b031916928716929092179091556040519081527f63e09f16584208fba1fc7ff64c62b00f07bec177c0d97ca6689891b1e77a35c79060200160405180910390a16040516001600160a01b03831681527f3c864541ef71378c6229510ed90f376565ee42d9c5e0904a984a9e863e6db44f9060200160405180910390a18015611464576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050505050565b60608151835114611a1e5760405162461bcd60e51b815260206004820152602960248201527f455243313135353a206163636f756e747320616e6420696473206c656e677468604482015268040dad2e6dac2e8c6d60bb1b6064820152608401610ba2565b600083516001600160401b03811115611a3957611a39614a32565b604051908082528060200260200182016040528015611a62578160200160208202803683370190505b50905060005b8451811015611ada57611aad858281518110611a8657611a866155b1565b6020026020010151858381518110611aa057611aa06155b1565b6020026020010151610b3b565b828281518110611abf57611abf6155b1565b6020908102919091010152611ad3816155dd565b9050611a68565b509392505050565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161415611b2b5760405162461bcd60e51b8152600401610ba2906154cb565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316611b74600080516020615a7c833981519152546001600160a01b031690565b6001600160a01b031614611b9a5760405162461bcd60e51b8152600401610ba290615517565b611ba382613402565b610c7b8282600161343c565b6000611bbd81610c35612a62565b611bca5780610c47612a62565b6001600160a01b038416611c145760405163eac0d38960e01b815260206004820152601060248201526f5f5f7375626a6563744761746577617960801b6044820152606401610ba2565b6001600160a01b038316611c595760405163eac0d38960e01b815260206004820152600b60248201526a2fafb0b63637b1b0ba37b960a91b6044820152606401610ba2565b6001600160a01b038216611ca05760405163eac0d38960e01b815260206004820152600d60248201526c2fafb234b9ba3934b13aba37b960991b6044820152606401610ba2565b6101ce80546001600160a01b03199081166001600160a01b038781169182179093556101d0805483168785169081179091556101d180549093169386169384179092556040517f99c462401f8e035efc2fc0fc71ec13cfd7399e9dce176e9aef9cb2e4487e44ed90600090a450505050565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611db25760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610ba2565b50600080516020615a7c8339815191525b90565b60008281526101c660205260408120548190610f97565b600082815261019160205260408120548015610f795760008481526101c46020526040902054610f7490610f6d565b60006101cc6000611e1d8585612b05565b815260208101919091526040016000205460ff169392505050565b600060ff821660011415611e4e57506001919050565b60ff821660021415611e6257506002919050565b60ff821660031415611e7657506003919050565b60ff8216611e8657506004919050565b506000919050565b610c7b611e99612a62565b838361373f565b6000610ea6611eaf8484612b05565b60009081526101c4602052604090205490565b6060816001600160401b03811115611edc57611edc614a32565b604051908082528060200260200182016040528015611f0f57816020015b6060815260200190600190039081611efa5790505b50905060005b8281101561177957611f7f30858584818110611f3357611f336155b1565b9050602002810190611f4591906155f8565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061382192505050565b828281518110611f9157611f916155b1565b60200260200101819052508080611fa7906155dd565b915050611f15565b7f24791c44c040514a5d2580696fc45e7d3cb6c9fa65bf3db2e4755362d6c155b5611fdc81610c35612a62565b611fe95780610c47612a62565b6101cf8290556040518281527f76096f684901589fef72a8b0a0321bf061b8597a791db667ab58e7eb922c639c906020015b60405180910390a15050565b60007f8aef0597c0be1e090afba1f387ee99f604b5d975ccbed6215cdf146ffd5c49fc61205681610c35612a62565b6120635780610c47612a62565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a082319060240160206040518083038186803b1580156120a557600080fd5b505afa1580156120b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120dd9190615645565b6101c3549091506001600160a01b038681169116141561211e576101c554612105908261565e565b90506121116101c75490565b61211b908261565e565b90505b6121298585836135ea565b604080516001600160a01b038681168252602082018490528716917fd092d7fceb5ea5a962639fcc27a7bb315e7637e699e3b108cd570c38c7584300910160405180910390a2949350505050565b6101c3546001600160a01b031663d505accf612191612a62565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152306024820152604481018890526064810187905260ff8616608482015260a4810185905260c4810184905260e401600060405180830381600087803b1580156121fd57600080fd5b505af1158015612211573d6000803e3d6000fd5b505050505050505050565b600061222a81610c35612a62565b6122375780610c47612a62565b62015180826001600160401b0316101561227757604051625a5b2760e31b81526001600160401b0383166004820152620151806024820152604401610ba2565b6276a700826001600160401b031611156122b857604051637f16270360e01b81526001600160401b03831660048201526276a7006024820152604401610ba2565b6101cd805467ffffffffffffffff19166001600160401b0384169081179091556040519081527f63e09f16584208fba1fc7ff64c62b00f07bec177c0d97ca6689891b1e77a35c79060200161201b565b600060ff821660021415610fd657506003919050565b600061232c81610c35612a62565b6123395780610c47612a62565b6123536001600160a01b038316637965db0b60e01b613846565b612391576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b6044820152606401610ba2565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a25050565b6065546001600160a01b031661242a5760405163eac0d38960e01b81526020600482015260126024820152712fb232b83932b1b0ba32b22fb937baba32b960711b6044820152606401610ba2565b606580546001600160a01b03191690556040516000907f7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc80908290a2565b6000610f7c8261038a86866133ba565b6000610ea661248684846133ba565b60009081526101c6602052604090205490565b60008360ff8116158015906124b2575060ff8116600114155b80156124c2575060ff8116600214155b80156124d2575060ff8116600314155b156124f55760405163c2628c0b60e01b815260ff82166004820152602401610ba2565b60006124ff612a62565b9050600061250d8787612b05565b90506125198282610b3b565b6125365760405163d7d0b56f60e01b815260040160405180910390fd5b6101cd546000906001600160401b031661254f42613862565b6125599190615675565b60008381526101c8602090815260408083206001600160a01b03881684529091529020805467ffffffffffffffff19166001600160401b038316179055905060006125ad876125a88686610b3b565b6138ce565b905060006125bb8483611ddd565b905060006125ce61010019861683611dc6565b905060006125db8c611e38565b90506125ea6101c48785612b51565b6125fc6101c661010019881685612b94565b612607878786612bcd565b61261987610100198816846000610da5565b600281600481111561262d5761262d6150fa565b148061264a57506003816004811115612648576126486150fa565b145b156126ce576101d060009054906101000a90046001600160a01b03166001600160a01b031663d0d87ac8878e8e8b888a6040518763ffffffff1660e01b815260040161269b969594939291906153cc565b600060405180830381600087803b1580156126b557600080fd5b505af11580156126c9573d6000803e3d6000fd5b505050505b6040516001600160401b03861681526001600160a01b038816908c9060ff8f16907f8b14266b7a7bfb46d73cda0a8ca3a635f38e8fb4b275b68aff68cb1b5a2aea3a9060200160405180910390a450929a9950505050505050505050565b6000610f7c8261038a8686612b05565b600061274a81610c35612a62565b6127575780610c47612a62565b6001600160a01b03821661279c5760405163eac0d38960e01b815260206004820152600b60248201526a6e6577547265617375727960a81b6044820152606401610ba2565b6101cd805468010000000000000000600160e01b031916600160401b6001600160a01b038516908102919091179091556040519081527f3c864541ef71378c6229510ed90f376565ee42d9c5e0904a984a9e863e6db44f9060200161201b565b612804612a62565b6001600160a01b0316856001600160a01b0316148061282a575061282a85610a54612a62565b6128465760405162461bcd60e51b8152600401610ba29061547c565b61146485858585856138e4565b60007f12b42e8a160f6064dc959c6f251e3af0750ad213dbecf573b4710d67d6c28e3961288281610c35612a62565b61288f5780610c47612a62565b8660038061289c83611e38565b60048111156128ad576128ad6150fa565b14156128bd578161118883611e38565b60006128c98a8a612b05565b905060026128d68b611e38565b60048111156128e7576128e76150fa565b14156129515760006128fe896101cf546064612e9b565b9050600061290c828b61565e565b905061291a838d8d84613a22565b811561294a57600061292b8d612308565b90506000612939828e612b05565b905061294781838f87613a22565b50505b505061295d565b61295d818b8b8b613a22565b600061296b89886064612e9b565b905080156129cd576001600160a01b0388166129b55760405163eac0d38960e01b8152602060048201526008602482015267383937b837b9b2b960c11b6044820152606401610ba2565b6101c3546129cd906001600160a01b031689836135ea565b6101c3546101cd546129fc916001600160a01b0390811691600160401b9004166129f7848d61565e565b6135ea565b876001600160a01b03168a8c60ff167fbad5bf3ab3814a2220a4737f205e42060a9b3b66422e774059048cec2f6ed0ac84604051612a3c91815260200190565b60405180910390a450969998505050505050505050565b6001600160a01b03163b151590565b60006110f9613bd0565b603354604051632474521560e21b8152600481018490526001600160a01b03838116602483015260009216906391d148549060440160206040518083038186803b158015612ab957600080fd5b505afa158015612acd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ea6919061545f565b8051610c7b9061016190602084019061492e565b6040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff909116176101001790565b60008281526020849052604081208054839290612b6f90849061565e565b9250508190555080836001016000828254612b8a919061565e565b9091555050505050565b60008281526020849052604081208054839290612bb29084906156a0565b9250508190555080836001016000828254612b8a91906156a0565b6001600160a01b038316612c2f5760405162461bcd60e51b815260206004820152602360248201527f455243313135353a206275726e2066726f6d20746865207a65726f206164647260448201526265737360e81b6064820152608401610ba2565b6000612c39612a62565b90506000612c4684613bda565b90506000612c5384613bda565b9050612c7383876000858560405180602001604052806000815250613c25565b600085815261015f602090815260408083206001600160a01b038a16845290915290205484811015612cf35760405162461bcd60e51b8152602060048201526024808201527f455243313135353a206275726e20616d6f756e7420657863656564732062616c604482015263616e636560e01b6064820152608401610ba2565b600086815261015f602090815260408083206001600160a01b038b81168086529184528285208a8703905582518b81529384018a90529092908816917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a46040805160208101909152600090525b50505050505050565b6001600160a01b038416612dd35760405162461bcd60e51b815260206004820152602160248201527f455243313135353a206d696e7420746f20746865207a65726f206164647265736044820152607360f81b6064820152608401610ba2565b6000612ddd612a62565b90506000612dea85613bda565b90506000612df785613bda565b9050612e0883600089858589613c25565b600086815261015f602090815260408083206001600160a01b038b16845290915281208054879290612e3b9084906156a0565b909155505060408051878152602081018790526001600160a01b03808a1692600092918716917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4612d6a83600089898989613ddf565b600080806000198587098587029250828110838203039150508060001415612ed657838281612ecc57612ecc6156b8565b0492505050610ea6565b808411612ee257600080fd5b60008486880960026001871981018816978890046003810283188082028403028082028403028082028403028082028403028082028403029081029092039091026000889003889004909101858311909403939093029303949094049190911702949350505050565b6101ce546040516276841960e61b815260ff8516600482015260248101849052600091829182916001600160a01b031690631da106409060440160206040518083038186803b158015612f9d57600080fd5b505afa158015612fb1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fd59190615645565b905080612fe28787611ea0565b10612ff557600060019250925050613036565b60006130018787611ea0565b61300b908361565e565b905061301785826138ce565b82866130238a8a611ea0565b61302d91906156a0565b10159350935050505b935093915050565b6040516001600160a01b03808516602483015283166044820152606481018290526114b09085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152613f4a565b815183511461310b5760405162461bcd60e51b815260206004820152602860248201527f455243313135353a2069647320616e6420616d6f756e7473206c656e677468206044820152670dad2e6dac2e8c6d60c31b6064820152608401610ba2565b6001600160a01b0384166131315760405162461bcd60e51b8152600401610ba2906156ce565b600061313b612a62565b905061314b818787878787613c25565b60005b845181101561323657600085828151811061316b5761316b6155b1565b602002602001015190506000858381518110613189576131896155b1565b602090810291909101810151600084815261015f835260408082206001600160a01b038e1683529093529190912054909150818110156131db5760405162461bcd60e51b8152600401610ba290615713565b600083815261015f602090815260408083206001600160a01b038e8116855292528083208585039055908b1682528120805484929061321b9084906156a0565b925050819055505050508061322f906155dd565b905061314e565b50846001600160a01b0316866001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb878760405161328692919061575d565b60405180910390a461329c81878787878761401c565b505050505050565b6040516302571be360e01b81527f91d1777781884d03a6757a803996e38de2a42967fb37eeaca72729271025a9e260048201526001600160a01b038416906302571be39060240160206040518083038186803b15801561330357600080fd5b505afa158015613317573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061333b9190615782565b6001600160a01b031663c47f002783836040518363ffffffff1660e01b815260040161336892919061579f565b602060405180830381600087803b15801561338257600080fd5b505af1158015613396573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114b09190615645565b6040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff9091161790565b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e361342f81610c35612a62565b610c7b5780610c47612a62565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff16156134745761346f836140e6565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156134ad57600080fd5b505afa9250505080156134dd575060408051601f3d908101601f191682019092526134da91810190615645565b60015b6135405760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610ba2565b600080516020615a7c83398151915281146135af5760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610ba2565b5061346f838383614182565b60006135d082516001600160401b0316151590565b8015610bd157505051426001600160401b03909116111590565b6040516001600160a01b03831660248201526044810182905261346f90849063a9059cbb60e01b90606401613072565b600054610100900460ff161580801561363a5750600054600160ff909116105b806136545750303b158015613654575060005460ff166001145b6136705760405162461bcd60e51b8152600401610ba290615563565b6000805460ff191660011790558015613693576000805461ff0019166101001790555b61369c826141a7565b6136a4613716565b8015610c7b576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200161201b565b600054610100900460ff1661370d5760405162461bcd60e51b8152600401610ba2906157ce565b6115a281614301565b600054610100900460ff1661373d5760405162461bcd60e51b8152600401610ba2906157ce565b565b816001600160a01b0316836001600160a01b031614156137b35760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2073657474696e6720617070726f76616c20737461747573604482015268103337b91039b2b63360b91b6064820152608401610ba2565b6001600160a01b0383811660008181526101606020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6060610ea68383604051806060016040528060278152602001615a9c60279139614331565b6000613851836143cf565b8015610ea65750610ea68383614402565b60006001600160401b038211156138ca5760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203660448201526534206269747360d01b6064820152608401610ba2565b5090565b60008183106138dd5781610ea6565b5090919050565b6001600160a01b03841661390a5760405162461bcd60e51b8152600401610ba2906156ce565b6000613914612a62565b9050600061392185613bda565b9050600061392e85613bda565b905061393e838989858589613c25565b600086815261015f602090815260408083206001600160a01b038c168452909152902054858110156139825760405162461bcd60e51b8152600401610ba290615713565b600087815261015f602090815260408083206001600160a01b038d8116855292528083208985039055908a168252812080548892906139c29084906156a0565b909155505060408051888152602081018890526001600160a01b03808b16928c821692918816917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4612211848a8a8a8a8a613ddf565b60008481526101c4602052604081205490613a41610100198716612486565b90506000613a5b613a5283856156a0565b605a6064612e9b565b905080841115613a7e57604051631c67aae360e11b815260040160405180910390fd5b6000613a948486613a8f86836156a0565b612e9b565b90506000613aa2828761565e565b9050613ab16101c48a84612b51565b613ac36101c6610100198b1683612b51565b6000613ace89611e38565b90506002816004811115613ae457613ae46150fa565b1480613b0157506003816004811115613aff57613aff6150fa565b145b15613b75576101d054604051631a1b0f5960e31b81526001600160a01b039091169063d0d87ac890613b42908d908d908d906000908a9082906004016153cc565b600060405180830381600087803b158015613b5c57600080fd5b505af1158015613b70573d6000803e3d6000fd5b505050505b613b7d612a62565b6001600160a01b0316888a60ff167f59d631535a5e48c1231a728e536ce39dba8d6c7f49905ec570e3db296430e02e8a604051613bbc91815260200190565b60405180910390a450505050505050505050565b60006110f96144e1565b60408051600180825281830190925260609160009190602080830190803683370190505090508281600081518110613c1457613c146155b1565b602090810291909101015292915050565b60005b8351811015613dd057613c58848281518110613c4657613c466155b1565b60200260200101516101009081161490565b15613d84576000613c7f858381518110613c7457613c746155b1565b602002602001015190565b905060ff81166003148015613c9c57506001600160a01b03861615155b8015613cb057506001600160a01b03871615155b15613d7e576101d05485516001600160a01b039091169063e2c11b8d90879085908110613cdf57613cdf6155b1565b6020026020010151838a8a898881518110613cfc57613cfc6155b1565b60209081029190910101516040516001600160e01b031960e088901b168152600481019590955260ff90931660248501526001600160a01b039182166044850152166064830152608482015260a401600060405180830381600087803b158015613d6557600080fd5b505af1158015613d79573d6000803e3d6000fd5b505050505b50613dbe565b6001600160a01b0386161580613da157506001600160a01b038516155b613dbe5760405163068331f960e01b815260040160405180910390fd5b80613dc8816155dd565b915050613c28565b5061329c868686868686614544565b6001600160a01b0384163b1561329c5760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e6190613e239089908990889088908890600401615819565b602060405180830381600087803b158015613e3d57600080fd5b505af1925050508015613e6d575060408051601f3d908101601f19168201909252613e6a91810190615853565b60015b613f1a57613e79615870565b806308c379a01415613eb35750613e8e61588b565b80613e995750613eb5565b8060405162461bcd60e51b8152600401610ba29190614c17565b505b60405162461bcd60e51b815260206004820152603460248201527f455243313135353a207472616e7366657220746f206e6f6e20455243313135356044820152732932b1b2b4bb32b91034b6b83632b6b2b73a32b960611b6064820152608401610ba2565b6001600160e01b0319811663f23a6e6160e01b14612d6a5760405162461bcd60e51b8152600401610ba290615914565b6000613f9f826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166146c09092919063ffffffff16565b80519091501561346f5780806020019051810190613fbd919061545f565b61346f5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610ba2565b6001600160a01b0384163b1561329c5760405163bc197c8160e01b81526001600160a01b0385169063bc197c8190614060908990899088908890889060040161595c565b602060405180830381600087803b15801561407a57600080fd5b505af19250505080156140aa575060408051601f3d908101601f191682019092526140a791810190615853565b60015b6140b657613e79615870565b6001600160e01b0319811663bc197c8160e01b14612d6a5760405162461bcd60e51b8152600401610ba290615914565b6001600160a01b0381163b6141535760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610ba2565b600080516020615a7c83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b61418b836146cf565b6000825111806141985750805b1561346f576114b0838361470f565b600054610100900460ff16158080156141c75750600054600160ff909116105b806141e15750303b1580156141e1575060005460ff166001145b6141fd5760405162461bcd60e51b8152600401610ba290615563565b6000805460ff191660011790558015614220576000805461ff0019166101001790555b61423a6001600160a01b038316637965db0b60e01b613846565b614278576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b6044820152606401610ba2565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a28015610c7b576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200161201b565b600054610100900460ff166143285760405162461bcd60e51b8152600401610ba2906157ce565b6115a281612af1565b60606001600160a01b0384163b61435a5760405162461bcd60e51b8152600401610ba2906159ba565b600080856001600160a01b0316856040516143759190615a00565b600060405180830381855af49150503d80600081146143b0576040519150601f19603f3d011682016040523d82523d6000602084013e6143b5565b606091505b50915091506143c58282866147c4565b9695505050505050565b60006143e2826301ffc9a760e01b614402565b8015610bd157506143fb826001600160e01b0319614402565b1592915050565b604080516001600160e01b0319831660248083019190915282518083039091018152604490910182526020810180516001600160e01b03166301ffc9a760e01b179052905160009190829081906001600160a01b0387169061753090614469908690615a00565b6000604051808303818686fa925050503d80600081146144a5576040519150601f19603f3d011682016040523d82523d6000602084013e6144aa565b606091505b50915091506020815110156144c55760009350505050610bd1565b8180156143c55750808060200190518101906143c5919061545f565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633141561453f5760003661452260148261565e565b61452e92369290615a1c565b61453791615a46565b60601c905090565b503390565b6001600160a01b0385166145cc5760005b83518110156145ca57828181518110614570576145706155b1565b6020026020010151610191600086848151811061458f5761458f6155b1565b6020026020010151815260200190815260200160002060008282546145b491906156a0565b909155506145c39050816155dd565b9050614555565b505b6001600160a01b03841661329c5760005b8351811015612d6a5760008482815181106145fa576145fa6155b1565b602002602001015190506000848381518110614618576146186155b1565b60200260200101519050600061019160008481526020019081526020016000205490508181101561469c5760405162461bcd60e51b815260206004820152602860248201527f455243313135353a206275726e20616d6f756e74206578636565647320746f74604482015267616c537570706c7960c01b6064820152608401610ba2565b60009283526101916020526040909220910390556146b9816155dd565b90506145dd565b6060610f7c84846000856147fd565b6146d8816140e6565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606001600160a01b0383163b6147385760405162461bcd60e51b8152600401610ba2906159ba565b600080846001600160a01b0316846040516147539190615a00565b600060405180830381855af49150503d806000811461478e576040519150601f19603f3d011682016040523d82523d6000602084013e614793565b606091505b50915091506147bb8282604051806060016040528060278152602001615a9c602791396147c4565b95945050505050565b606083156147d3575081610ea6565b8251156147e35782518084602001fd5b8160405162461bcd60e51b8152600401610ba29190614c17565b60608247101561485e5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610ba2565b6001600160a01b0385163b6148b55760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610ba2565b600080866001600160a01b031685876040516148d19190615a00565b60006040518083038185875af1925050503d806000811461490e576040519150601f19603f3d011682016040523d82523d6000602084013e614913565b606091505b50915091506149238282866147c4565b979650505050505050565b82805461493a90615400565b90600052602060002090601f01602090048101928261495c57600085556149a2565b82601f1061497557805160ff19168380011785556149a2565b828001600101855582156149a2579182015b828111156149a2578251825591602001919060010190614987565b506138ca9291505b808211156138ca57600081556001016149aa565b6001600160a01b03811681146115a257600080fd5b600080604083850312156149e657600080fd5b82356149f1816149be565b946020939093013593505050565b6001600160e01b0319811681146115a257600080fd5b600060208284031215614a2757600080fd5b8135610ea6816149ff565b634e487b7160e01b600052604160045260246000fd5b601f8201601f191681016001600160401b0381118282101715614a6d57614a6d614a32565b6040525050565b60006001600160401b03831115614a8d57614a8d614a32565b604051614aa4601f8501601f191660200182614a48565b809150838152848484011115614ab957600080fd5b83836020830137600060208583010152509392505050565b600060208284031215614ae357600080fd5b81356001600160401b03811115614af957600080fd5b8201601f81018413614b0a57600080fd5b610f7c84823560208401614a74565b803560ff81168114614b2a57600080fd5b919050565b600080600080600060a08688031215614b4757600080fd5b614b5086614b19565b945060208601359350614b6560408701614b19565b9250606086013591506080860135614b7c816149be565b809150509295509295909350565b60008060408385031215614b9d57600080fd5b6149f183614b19565b600060208284031215614bb857600080fd5b5035919050565b60005b83811015614bda578181015183820152602001614bc2565b838111156114b05750506000910152565b60008151808452614c03816020860160208601614bbf565b601f01601f19169290920160200192915050565b602081526000610ea66020830184614beb565b60008060408385031215614c3d57600080fd5b50508035926020909101359150565b600060208284031215614c5e57600080fd5b610ea682614b19565b80151581146115a257600080fd5b600080600060608486031215614c8a57600080fd5b614c9384614b19565b9250602084013591506040840135614caa81614c67565b809150509250925092565b600080600060608486031215614cca57600080fd5b614cd384614b19565b95602085013595506040909401359392505050565b60006001600160401b03821115614d0157614d01614a32565b5060051b60200190565b600082601f830112614d1c57600080fd5b81356020614d2982614ce8565b604051614d368282614a48565b83815260059390931b8501820192828101915086841115614d5657600080fd5b8286015b84811015614d715780358352918301918301614d5a565b509695505050505050565b600082601f830112614d8d57600080fd5b610ea683833560208501614a74565b600080600080600060a08688031215614db457600080fd5b8535614dbf816149be565b94506020860135614dcf816149be565b935060408601356001600160401b0380821115614deb57600080fd5b614df789838a01614d0b565b94506060880135915080821115614e0d57600080fd5b614e1989838a01614d0b565b93506080880135915080821115614e2f57600080fd5b50614e3c88828901614d7c565b9150509295509295909350565b600080600060408486031215614e5e57600080fd5b8335614e69816149be565b925060208401356001600160401b0380821115614e8557600080fd5b818601915086601f830112614e9957600080fd5b813581811115614ea857600080fd5b876020828501011115614eba57600080fd5b6020830194508093505050509250925092565b600060208284031215614edf57600080fd5b8135610ea6816149be565b80356001600160401b0381168114614b2a57600080fd5b60008060008060808587031215614f1757600080fd5b8435614f22816149be565b93506020850135614f32816149be565b9250614f4060408601614eea565b91506060850135614f50816149be565b939692955090935050565b60008060408385031215614f6e57600080fd5b82356001600160401b0380821115614f8557600080fd5b818501915085601f830112614f9957600080fd5b81356020614fa682614ce8565b604051614fb38282614a48565b83815260059390931b8501820192828101915089841115614fd357600080fd5b948201945b83861015614ffa578535614feb816149be565b82529482019490820190614fd8565b9650508601359250508082111561501057600080fd5b5061501d85828601614d0b565b9150509250929050565b600081518084526020808501945080840160005b838110156150575781518752958201959082019060010161503b565b509495945050505050565b602081526000610ea66020830184615027565b6000806040838503121561508857600080fd5b8235615093816149be565b915060208301356001600160401b038111156150ae57600080fd5b61501d85828601614d7c565b6000806000606084860312156150cf57600080fd5b83356150da816149be565b925060208401356150ea816149be565b91506040840135614caa816149be565b634e487b7160e01b600052602160045260246000fd5b6005811061512e57634e487b7160e01b600052602160045260246000fd5b9052565b60208101610bd18284615110565b6000806040838503121561515357600080fd5b823561515e816149be565b9150602083013561516e81614c67565b809150509250929050565b6000806020838503121561518c57600080fd5b82356001600160401b03808211156151a357600080fd5b818501915085601f8301126151b757600080fd5b8135818111156151c657600080fd5b8660208260051b85010111156151db57600080fd5b60209290920196919550909350505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561524257603f19888603018452615230858351614beb565b94509285019290850190600101615214565b5092979650505050505050565b6000806040838503121561526257600080fd5b823561526d816149be565b9150602083013561516e816149be565b600080600080600060a0868803121561529557600080fd5b85359450602086013593506152ac60408701614b19565b94979396509394606081013594506080013592915050565b6000602082840312156152d657600080fd5b610ea682614eea565b6000806000606084860312156152f457600080fd5b6152fd84614b19565b9250602084013591506040840135614caa816149be565b600080600080600060a0868803121561532c57600080fd5b8535615337816149be565b94506020860135615347816149be565b9350604086013592506060860135915060808601356001600160401b0381111561537057600080fd5b614e3c88828901614d7c565b600080600080600060a0868803121561539457600080fd5b61539d86614b19565b9450602086013593506040860135925060608601356153bb816149be565b949793965091946080013592915050565b95865260ff94909416602086015260408501929092526001600160a01b03166060840152608083015260a082015260c00190565b600181811c9082168061541457607f821691505b6020821081141561543557634e487b7160e01b600052602260045260246000fd5b50919050565b60ff84168152606081016154526020830185615110565b610f7c6040830184615110565b60006020828403121561547157600080fd5b8151610ea681614c67565b6020808252602f908201527f455243313135353a2063616c6c6572206973206e6f7420746f6b656e206f776e60408201526e195c881b9bdc88185c1c1c9bdd9959608a1b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156155f1576155f16155c7565b5060010190565b6000808335601e1984360301811261560f57600080fd5b8301803591506001600160401b0382111561562957600080fd5b60200191503681900382131561563e57600080fd5b9250929050565b60006020828403121561565757600080fd5b5051919050565b600082821015615670576156706155c7565b500390565b60006001600160401b03808316818516808303821115615697576156976155c7565b01949350505050565b600082198211156156b3576156b36155c7565b500190565b634e487b7160e01b600052601260045260246000fd5b60208082526025908201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604082015264647265737360d81b606082015260800190565b6020808252602a908201527f455243313135353a20696e73756666696369656e742062616c616e636520666f60408201526939103a3930b739b332b960b11b606082015260800190565b6040815260006157706040830185615027565b82810360208401526147bb8185615027565b60006020828403121561579457600080fd5b8151610ea6816149be565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b6001600160a01b03868116825285166020820152604081018490526060810183905260a06080820181905260009061492390830184614beb565b60006020828403121561586557600080fd5b8151610ea6816149ff565b600060033d1115611dc35760046000803e5060005160e01c90565b600060443d10156158995790565b6040516003193d81016004833e81513d6001600160401b0381602484011181841117156158c857505050505090565b82850191508151818111156158e05750505050505090565b843d87010160208285010111156158fa5750505050505090565b61590960208286010187614a48565b509095945050505050565b60208082526028908201527f455243313135353a204552433131353552656365697665722072656a656374656040820152676420746f6b656e7360c01b606082015260800190565b6001600160a01b0386811682528516602082015260a06040820181905260009061598890830186615027565b828103606084015261599a8186615027565b905082810360808401526159ae8185614beb565b98975050505050505050565b60208082526026908201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6040820152651b9d1c9858dd60d21b606082015260800190565b60008251615a12818460208701614bbf565b9190910192915050565b60008085851115615a2c57600080fd5b83861115615a3957600080fd5b5050820193919092039150565b6bffffffffffffffffffffffff198135818116916014851015615a735780818660140360031b1b83161692505b50509291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212206992396fc4c87b191f66494e2478323cf026a76b999e207da4638723879c695264736f6c63430008090033",
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

// ConfigureStakeHelpers is a paid mutator transaction binding the contract method 0x5128e693.
//
// Solidity: function configureStakeHelpers(address __subjectGateway, address __allocator, address __distributor) returns()
func (_FortaStaking *FortaStakingTransactor) ConfigureStakeHelpers(opts *bind.TransactOpts, __subjectGateway common.Address, __allocator common.Address, __distributor common.Address) (*types.Transaction, error) {
	return _FortaStaking.contract.Transact(opts, "configureStakeHelpers", __subjectGateway, __allocator, __distributor)
}

// ConfigureStakeHelpers is a paid mutator transaction binding the contract method 0x5128e693.
//
// Solidity: function configureStakeHelpers(address __subjectGateway, address __allocator, address __distributor) returns()
func (_FortaStaking *FortaStakingSession) ConfigureStakeHelpers(__subjectGateway common.Address, __allocator common.Address, __distributor common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.ConfigureStakeHelpers(&_FortaStaking.TransactOpts, __subjectGateway, __allocator, __distributor)
}

// ConfigureStakeHelpers is a paid mutator transaction binding the contract method 0x5128e693.
//
// Solidity: function configureStakeHelpers(address __subjectGateway, address __allocator, address __distributor) returns()
func (_FortaStaking *FortaStakingTransactorSession) ConfigureStakeHelpers(__subjectGateway common.Address, __allocator common.Address, __distributor common.Address) (*types.Transaction, error) {
	return _FortaStaking.Contract.ConfigureStakeHelpers(&_FortaStaking.TransactOpts, __subjectGateway, __allocator, __distributor)
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
	SubjectGateway     common.Address
	Allocator          common.Address
	RewardsDistributor common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterStakeHelpersConfigured is a free log retrieval operation binding the contract event 0x99c462401f8e035efc2fc0fc71ec13cfd7399e9dce176e9aef9cb2e4487e44ed.
//
// Solidity: event StakeHelpersConfigured(address indexed subjectGateway, address indexed allocator, address indexed rewardsDistributor)
func (_FortaStaking *FortaStakingFilterer) FilterStakeHelpersConfigured(opts *bind.FilterOpts, subjectGateway []common.Address, allocator []common.Address, rewardsDistributor []common.Address) (*FortaStakingStakeHelpersConfiguredIterator, error) {

	var subjectGatewayRule []interface{}
	for _, subjectGatewayItem := range subjectGateway {
		subjectGatewayRule = append(subjectGatewayRule, subjectGatewayItem)
	}
	var allocatorRule []interface{}
	for _, allocatorItem := range allocator {
		allocatorRule = append(allocatorRule, allocatorItem)
	}
	var rewardsDistributorRule []interface{}
	for _, rewardsDistributorItem := range rewardsDistributor {
		rewardsDistributorRule = append(rewardsDistributorRule, rewardsDistributorItem)
	}

	logs, sub, err := _FortaStaking.contract.FilterLogs(opts, "StakeHelpersConfigured", subjectGatewayRule, allocatorRule, rewardsDistributorRule)
	if err != nil {
		return nil, err
	}
	return &FortaStakingStakeHelpersConfiguredIterator{contract: _FortaStaking.contract, event: "StakeHelpersConfigured", logs: logs, sub: sub}, nil
}

// WatchStakeHelpersConfigured is a free log subscription operation binding the contract event 0x99c462401f8e035efc2fc0fc71ec13cfd7399e9dce176e9aef9cb2e4487e44ed.
//
// Solidity: event StakeHelpersConfigured(address indexed subjectGateway, address indexed allocator, address indexed rewardsDistributor)
func (_FortaStaking *FortaStakingFilterer) WatchStakeHelpersConfigured(opts *bind.WatchOpts, sink chan<- *FortaStakingStakeHelpersConfigured, subjectGateway []common.Address, allocator []common.Address, rewardsDistributor []common.Address) (event.Subscription, error) {

	var subjectGatewayRule []interface{}
	for _, subjectGatewayItem := range subjectGateway {
		subjectGatewayRule = append(subjectGatewayRule, subjectGatewayItem)
	}
	var allocatorRule []interface{}
	for _, allocatorItem := range allocator {
		allocatorRule = append(allocatorRule, allocatorItem)
	}
	var rewardsDistributorRule []interface{}
	for _, rewardsDistributorItem := range rewardsDistributor {
		rewardsDistributorRule = append(rewardsDistributorRule, rewardsDistributorItem)
	}

	logs, sub, err := _FortaStaking.contract.WatchLogs(opts, "StakeHelpersConfigured", subjectGatewayRule, allocatorRule, rewardsDistributorRule)
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

// ParseStakeHelpersConfigured is a log parse operation binding the contract event 0x99c462401f8e035efc2fc0fc71ec13cfd7399e9dce176e9aef9cb2e4487e44ed.
//
// Solidity: event StakeHelpersConfigured(address indexed subjectGateway, address indexed allocator, address indexed rewardsDistributor)
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
