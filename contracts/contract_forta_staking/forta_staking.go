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
	Bin: "0x60c06040523060a0523480156200001557600080fd5b506040516200615c3803806200615c833981016040819052620000389162000105565b6001600160a01b038116608052600054610100900460ff16806200005f575060005460ff16155b620000c75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b600054610100900460ff16158015620000ea576000805461ffff19166101011790555b8015620000fd576000805461ff00191690555b505062000137565b6000602082840312156200011857600080fd5b81516001600160a01b03811681146200013057600080fd5b9392505050565b60805160a051615fe362000179600039600081816112130152818161129801528181611717015261179c0152600081816105b40152614d360152615fe36000f3fe6080604052600436106102db5760003560e01c806375eb81d611610184578063c9580804116100d6578063e985e9c51161008a578063f0f4426011610064578063f0f44260146108db578063f242432a146108fb578063fe54ec801461091b57600080fd5b8063e985e9c514610838578063edea0bac14610882578063f08d35ee146108bb57600080fd5b8063d1e59d1c116100bb578063d1e59d1c146107d8578063da5bfb94146107f8578063dc4653ef1461081857600080fd5b8063c95808041461077f578063cc7a262e1461079f57600080fd5b8063ac9650d811610138578063c07975ad11610112578063c07975ad1461071f578063c0d786551461073f578063c10733021461075f57600080fd5b8063ac9650d8146106a4578063b8dc491b146106d1578063bd85b039146106f157600080fd5b80638c5588ac116101695780638c5588ac14610644578063a22cb46514610664578063a290bf381461068457600080fd5b806375eb81d6146106045780638432d6b61461062457600080fd5b80633121db1c1161023d5780634e1273f4116101f157806354fd4d50116101cb57806354fd4d501461054e578063572b6c051461059757806375e130ad146105e457600080fd5b80634e1273f4146104de5780634f1ef2861461050b5780634f558e791461051e57600080fd5b80633659cfe6116102225780633659cfe614610488578063371f4226146104a85780633f489914146104be57600080fd5b80633121db1c14610448578063321ebc541461046857600080fd5b80630e89341c1161029457806328f731481161027957806328f73148146103f25780632cb31144146104085780632eb2c2d61461042857600080fd5b80630e89341c146103a5578063226cc300146103d257600080fd5b806302fe5305116102c557806302fe5305146103435780630a615413146103655780630e10103f1461038557600080fd5b8062fdd58e146102e057806301ffc9a714610313575b600080fd5b3480156102ec57600080fd5b506103006102fb3660046152b8565b61093b565b6040519081526020015b60405180910390f35b34801561031f57600080fd5b5061033361032e3660046152fa565b6109ea565b604051901515815260200161030a565b34801561034f57600080fd5b5061036361035e3660046153b8565b610a6c565b005b34801561037157600080fd5b50610363610380366004615401565b610add565b34801561039157600080fd5b506103006103a0366004615434565b610ba7565b3480156103b157600080fd5b506103c56103c0366004615450565b610c04565b60405161030a91906154c1565b3480156103de57600080fd5b506103636103ed3660046154e2565b610c99565b3480156103fe57600080fd5b506101c554610300565b34801561041457600080fd5b50610300610423366004615522565b610de0565b34801561043457600080fd5b5061036361044336600461560a565b6110b7565b34801561045457600080fd5b506103636104633660046156b8565b61116b565b34801561047457600080fd5b50610300610483366004615434565b6111b6565b34801561049457600080fd5b506103636104a3366004615401565b611208565b3480156104b457600080fd5b506101c754610300565b3480156104ca57600080fd5b506103006104d9366004615434565b611384565b3480156104ea57600080fd5b506104fe6104f936600461573d565b6115ce565b60405161030a9190615845565b610363610519366004615858565b61170c565b34801561052a57600080fd5b50610333610539366004615450565b60009081526101916020526040902054151590565b34801561055a57600080fd5b506103c56040518060400160405280600581526020017f302e312e3000000000000000000000000000000000000000000000000000000081525081565b3480156105a357600080fd5b506103336105b2366004615401565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0390811691161490565b3480156105f057600080fd5b506103336105ff366004615434565b611875565b34801561061057600080fd5b5061036361061f366004615522565b6118e8565b34801561063057600080fd5b5061036361063f3660046158b6565b6119f5565b34801561065057600080fd5b5061030061065f366004615925565b611c03565b34801561067057600080fd5b5061036361067f36600461595a565b611c6c565b34801561069057600080fd5b5061030061069f366004615434565b611c7e565b3480156106b057600080fd5b506106c46106bf366004615993565b611ce7565b60405161030a9190615a08565b3480156106dd57600080fd5b506103006106ec366004615a6a565b611dd5565b3480156106fd57600080fd5b5061030061070c366004615450565b6000908152610191602052604090205490565b34801561072b57600080fd5b5061036361073a366004615a98565b611f4f565b34801561074b57600080fd5b5061036361075a366004615401565b611ff4565b34801561076b57600080fd5b5061036361077a366004615adf565b6120b1565b34801561078b57600080fd5b5061036361079a366004615401565b612125565b3480156107ab57600080fd5b506101c3546107c0906001600160a01b031681565b6040516001600160a01b03909116815260200161030a565b3480156107e457600080fd5b506103006107f3366004615925565b6121e3565b34801561080457600080fd5b50610300610813366004615925565b6123b1565b34801561082457600080fd5b50610300610833366004615434565b61240c565b34801561084457600080fd5b50610333610853366004615a6a565b6001600160a01b0391821660009081526101606020908152604080832093909416825291909152205460ff1690565b34801561088e57600080fd5b506108a261089d366004615522565b612471565b60405167ffffffffffffffff909116815260200161030a565b3480156108c757600080fd5b506103006108d6366004615925565b612684565b3480156108e757600080fd5b506103636108f6366004615401565b6126db565b34801561090757600080fd5b50610363610916366004615afa565b6127c2565b34801561092757600080fd5b50610300610936366004615522565b61286f565b60006001600160a01b0383166109be5760405162461bcd60e51b815260206004820152602b60248201527f455243313135353a2062616c616e636520717565727920666f7220746865207a60448201527f65726f206164647265737300000000000000000000000000000000000000000060648201526084015b60405180910390fd5b50600081815261015f602090815260408083206001600160a01b03861684529091529020545b92915050565b60006001600160e01b031982167fd9b67a26000000000000000000000000000000000000000000000000000000001480610a4d57506001600160e01b031982167f0e89341c00000000000000000000000000000000000000000000000000000000145b806109e457506301ffc9a760e01b6001600160e01b03198316146109e4565b6000610a7f81610a7a612a8a565b612a94565b610ad05780610a8c612a8a565b6040517f75000dc000000000000000000000000000000000000000000000000000000000815260048101929092526001600160a01b031660248201526044016109b5565b610ad982612b32565b5050565b6000610aeb81610a7a612a8a565b610af85780610a8c612a8a565b6001600160a01b038216610b4f5760405163eac0d38960e01b815260206004820152601460248201527f6e65775374616b696e67506172616d657465727300000000000000000000000060448201526064016109b5565b6040516001600160a01b038316907f06b2874e4c6a9fd4be614bef4bf7205f773309fd0d578a6f9b08d7b1f95347fb90600090a2506101ce80546001600160a01b0319166001600160a01b0392909216919091179055565b6000610bfd61070c84846040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff909116176101001790565b9392505050565b60606101618054610c1490615b63565b80601f0160208091040260200160405190810160405280929190818152602001828054610c4090615b63565b8015610c8d5780601f10610c6257610100808354040283529160200191610c8d565b820191906000526020600020905b815481529060010190602001808311610c7057829003601f168201915b50505050509050919050565b7f12b42e8a160f6064dc959c6f251e3af0750ad213dbecf573b4710d67d6c28e39610cc681610a7a612a8a565b610cd35780610a8c612a8a565b8360ff811615801590610cea575060ff8116600114155b15610d0d5760405163c2628c0b60e01b815260ff821660048201526024016109b5565b826101cc6000610d6488886040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff909116176101001790565b81526020810191909152604001600020805460ff1916911515919091179055610d8b612a8a565b6001600160a01b0316848660ff167fd520b4ee12f45aacea6f06fd4831c1410e0d3f874a80f64424715327d7f705ca86604051610dcc911515815260200190565b60405180910390a45050505050565b905090565b60008360ff811615801590610df9575060ff8116600114155b15610e1c5760405163c2628c0b60e01b815260ff821660048201526024016109b5565b6101ce546001600160a01b0316610e765760405163eac0d38960e01b815260206004820152601260248201527f5f7374616b696e67506172616d6574657273000000000000000000000000000060448201526064016109b5565b6101ce546040517f882b291700000000000000000000000000000000000000000000000000000000815260ff87166004820152602481018690526001600160a01b039091169063882b29179060440160206040518083038186803b158015610edd57600080fd5b505afa158015610ef1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f159190615b9e565b610f4b576040517facf735e400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610f55612a8a565b90506000610faa87876040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff909116176101001790565b90506000610fb9888888612b46565b90965090508015610ff557604051879060ff8a16907f8db486b45abb5d250b92d9eeea290a901f784921d6d28705cfaf2e9fa025d56390600090a35b60006110018388612c53565b6101c35490915061101d906001600160a01b031685308a612c97565b61102a6101c48489612d1b565b61106384848360005b6040519080825280601f01601f19166020018201604052801561105d576020820181803683370190505b50612d5e565b836001600160a01b0316888a60ff167f3e5900eae50f376487f0d2272f8f857884c7adde3b5be2cd8c7b1e7bdfc47e148a6040516110a391815260200190565b60405180910390a498975050505050505050565b6110bf612a8a565b6001600160a01b0316856001600160a01b031614806110e557506110e585610853612a8a565b6111575760405162461bcd60e51b815260206004820152603260248201527f455243313135353a207472616e736665722063616c6c6572206973206e6f742060448201527f6f776e6572206e6f7220617070726f766564000000000000000000000000000060648201526084016109b5565b6111648585858585612e98565b5050505050565b7f664245c7af190fec316596e8231f724e8171b1966cfcd124347ac5a66c34f64a61119881610a7a612a8a565b6111a55780610a8c612a8a565b6111b0848484613121565b50505050565b6000610bfd61070c84846040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff9091161790565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156112965760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b60648201526084016109b5565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166112f17f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b03161461135c5760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b6163746976652070726f787960a01b60648201526084016109b5565b61136581613250565b604080516000808252602082019092526113819183919061328a565b50565b60008260ff81161580159061139d575060ff8116600114155b156113c05760405163c2628c0b60e01b815260ff821660048201526024016109b5565b60006113ca612a8a565b9050600061141b86866040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff9091161790565b9050611427828261093b565b61145d576040517f65efc92d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610100811760009081526101cc602052604090205460ff16156114ac576040517f2799ebef00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610100811760009081526101c8602090815260408083206001600160a01b03861684528252918290208251918201909252815467ffffffffffffffff1681526114f490613432565b61152a576040517f0f2ca6e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805467ffffffffffffffff19168155826001600160a01b0316868860ff167f07e9e8a51e2cf8929e95153e486868bda03206cdbbf30ee732b17e62b43c282760405160405180910390a46000611580848461093b565b9050600061158e8483613463565b905061159d6101c685836134a0565b6115a88585846134d9565b6101c3546115c0906001600160a01b0316868361368c565b9550505050505b5092915050565b606081518351146116475760405162461bcd60e51b815260206004820152602960248201527f455243313135353a206163636f756e747320616e6420696473206c656e67746860448201527f206d69736d61746368000000000000000000000000000000000000000000000060648201526084016109b5565b6000835167ffffffffffffffff81111561166357611663615317565b60405190808252806020026020018201604052801561168c578160200160208202803683370190505b50905060005b8451811015611704576116d78582815181106116b0576116b0615bbb565b60200260200101518583815181106116ca576116ca615bbb565b602002602001015161093b565b8282815181106116e9576116e9615bbb565b60209081029190910101526116fd81615be7565b9050611692565b509392505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016141561179a5760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b60648201526084016109b5565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166117f57f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b0316146118605760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b6163746976652070726f787960a01b60648201526084016109b5565b61186982613250565b610ad98282600161328a565b60006101cc60006118cd85856040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff909116176101001790565b815260208101919091526040016000205460ff169392505050565b8260ff8116158015906118ff575060ff8116600114155b156119225760405163c2628c0b60e01b815260ff821660048201526024016109b5565b6101c354611942906001600160a01b031661193b612a8a565b3085612c97565b6119a061199685856040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff909116176101001790565b6101c99084612d1b565b6119a8612a8a565b6001600160a01b0316838560ff167f4bae81d7c78d9effd9b0ffe353b35b51c2695820985aa889b3d8916a017f60a0856040516119e791815260200190565b60405180910390a450505050565b600054610100900460ff1680611a0e575060005460ff16155b611a715760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016109b5565b600054610100900460ff16158015611a93576000805461ffff19166101011790555b6001600160a01b038216611aea5760405163eac0d38960e01b815260206004820152600a60248201527f5f5f74726561737572790000000000000000000000000000000000000000000060448201526064016109b5565b611af3866136da565b611afc8561382c565b611b0461397d565b611b1c60405180602001604052806000815250613a3f565b611b24613b0b565b6101c380546001600160a01b0319166001600160a01b03868116919091179091556101cd805467ffffffffffffffff86166001600160e01b0319909116811768010000000000000000938616939093029290921790556040519081527f63e09f16584208fba1fc7ff64c62b00f07bec177c0d97ca6689891b1e77a35c79060200160405180910390a16040516001600160a01b03831681527f3c864541ef71378c6229510ed90f376565ee42d9c5e0904a984a9e863e6db44f9060200160405180910390a18015611bfb576000805461ff00191690555b505050505050565b600080611c5785856040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff909116176101001790565b9050611c638184613bb1565b95945050505050565b610ad9611c77612a8a565b8383613c03565b6000610bfd611cd484846040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff909116176101001790565b60009081526101c4602052604090205490565b60608167ffffffffffffffff811115611d0257611d02615317565b604051908082528060200260200182016040528015611d3557816020015b6060815260200190600190039081611d205790505b50905060005b828110156115c757611da530858584818110611d5957611d59615bbb565b9050602002810190611d6b9190615c02565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613cf992505050565b828281518110611db757611db7615bbb565b60200260200101819052508080611dcd90615be7565b915050611d3b565b60007f8aef0597c0be1e090afba1f387ee99f604b5d975ccbed6215cdf146ffd5c49fc611e0481610a7a612a8a565b611e115780610a8c612a8a565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000906001600160a01b038616906370a082319060240160206040518083038186803b158015611e6c57600080fd5b505afa158015611e80573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ea49190615c50565b6101c3549091506001600160a01b0386811691161415611ef6576101c554611ecc9082615c69565b9050611ed86101c75490565b611ee29082615c69565b6101ca54909150611ef39082615c69565b90505b611f0185858361368c565b604080516001600160a01b038681168252602082018490528716917fd092d7fceb5ea5a962639fcc27a7bb315e7637e699e3b108cd570c38c7584300910160405180910390a2949350505050565b6101c3546001600160a01b031663d505accf611f69612a8a565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152306024820152604481018890526064810187905260ff8616608482015260a4810185905260c4810184905260e401600060405180830381600087803b158015611fd557600080fd5b505af1158015611fe9573d6000803e3d6000fd5b505050505050505050565b600061200281610a7a612a8a565b61200f5780610a8c612a8a565b6001600160a01b0382166120665760405163eac0d38960e01b815260206004820152600960248201527f6e6577526f75746572000000000000000000000000000000000000000000000060448201526064016109b5565b606580546001600160a01b0319166001600160a01b0384169081179091556040517f7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc8090600090a25050565b60006120bf81610a7a612a8a565b6120cc5780610a8c612a8a565b6101cd805467ffffffffffffffff191667ffffffffffffffff84169081179091556040519081527f63e09f16584208fba1fc7ff64c62b00f07bec177c0d97ca6689891b1e77a35c7906020015b60405180910390a15050565b600061213381610a7a612a8a565b6121405780610a8c612a8a565b61215a6001600160a01b038316637965db0b60e01b613d1e565b612198576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b60448201526064016109b5565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a25050565b60008360ff8116158015906121fc575060ff8116600114155b1561221f5760405163c2628c0b60e01b815260ff821660048201526024016109b5565b600061227286866040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff909116176101001790565b905060006122808286613bb1565b905061228f6101c983836134a0565b6122b28561229c83613d3a565b60008581526101cb602052604090209190613dbd565b6101c3546122ca906001600160a01b0316868361368c565b846001600160a01b0316868860ff167f8f404fb0169bf58e5bb9d571055da9ac0fc4eb973f2e6c7d1b108024381f0fc68460405161230a91815260200190565b60405180910390a4843b1515801561233757506123376001600160a01b038616631f0ae67f60e21b613d1e565b156123a757604051631f0ae67f60e21b815260ff8816600482015260248101879052604481018290526001600160a01b03861690637c2b99fc90606401600060405180830381600087803b15801561238e57600080fd5b505af11580156123a2573d6000803e3d6000fd5b505050505b9695505050505050565b6000612404826102fb86866040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff9091161790565b949350505050565b6000610bfd61245e84846040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff9091161790565b60009081526101c6602052604090205490565b60008360ff81161580159061248a575060ff8116600114155b156124ad5760405163c2628c0b60e01b815260ff821660048201526024016109b5565b60006124b7612a8a565b9050600061250c87876040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff909116176101001790565b9050612518828261093b565b61254e576040517fd7d0b56f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6101cd5460009067ffffffffffffffff1661256842613e59565b6125729190615c80565b60008381526101c8602090815260408083206001600160a01b03881684529091529020805467ffffffffffffffff191667ffffffffffffffff8316179055905060006125c7876125c2868661093b565b613ed9565b905060006125d58483613eef565b905060006125e861010019861683613f22565b90506125f76101c486846134a0565b6126096101c661010019871684612d1b565b6126148686856134d9565b61262686610100198716836000611033565b60405167ffffffffffffffff851681526001600160a01b038716908b9060ff8e16907f8b14266b7a7bfb46d73cda0a8ca3a635f38e8fb4b275b68aff68cb1b5a2aea3a9060200160405180910390a450919998505050505050505050565b6000612404826102fb86866040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff909116176101001790565b60006126e981610a7a612a8a565b6126f65780610a8c612a8a565b6001600160a01b03821661274d5760405163eac0d38960e01b815260206004820152600b60248201527f6e6577547265617375727900000000000000000000000000000000000000000060448201526064016109b5565b6101cd80547fffffffff0000000000000000000000000000000000000000ffffffffffffffff16680100000000000000006001600160a01b038516908102919091179091556040519081527f3c864541ef71378c6229510ed90f376565ee42d9c5e0904a984a9e863e6db44f90602001612119565b6127ca612a8a565b6001600160a01b0316856001600160a01b031614806127f057506127f085610853612a8a565b6128625760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2063616c6c6572206973206e6f74206f776e6572206e6f7260448201527f20617070726f766564000000000000000000000000000000000000000000000060648201526084016109b5565b6111648585858585613f39565b60007f12b42e8a160f6064dc959c6f251e3af0750ad213dbecf573b4710d67d6c28e3961289e81610a7a612a8a565b6128ab5780610a8c612a8a565b8460ff8116158015906128c2575060ff8116600114155b156128e55760405163c2628c0b60e01b815260ff821660048201526024016109b5565b600061293887876040805160f884901b6001600160f81b031916602080830191909152602180830194909452825180830390940184526041909101909152815191012060091b60ff909116176101001790565b60008181526101c4602052604081205491925061295961010019841661245e565b905060006129736009600a61296e8587615cac565b6140e8565b9050808811156129af576040517f38cf55c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006129c5846129bf8582615cac565b8b6140e8565b905060006129d3828b615c69565b90506129df8183615cac565b99506129ee6101c487846134a0565b612a006101c6610100198816836134a0565b6101c3546101cd54612a2b916001600160a01b0390811691680100000000000000009004168c61368c565b612a33612a8a565b6001600160a01b03168b8d60ff167f59d631535a5e48c1231a728e536ce39dba8d6c7f49905ec570e3db296430e02e8d604051612a7291815260200190565b60405180910390a450979a9950505050505050505050565b6000610ddb6141c7565b6033546040517f91d14854000000000000000000000000000000000000000000000000000000008152600481018490526001600160a01b03838116602483015260009216906391d148549060440160206040518083038186803b158015612afa57600080fd5b505afa158015612b0e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bfd9190615b9e565b8051610ad990610161906020840190615213565b6101ce546040517f1da1064000000000000000000000000000000000000000000000000000000000815260ff8516600482015260248101849052600091829182916001600160a01b031690631da106409060440160206040518083038186803b158015612bb257600080fd5b505afa158015612bc6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bea9190615c50565b905080612bf78787611c7e565b10612c0a57600060019250925050612c4b565b6000612c168787611c7e565b612c209083615c69565b9050612c2c8582613ed9565b8286612c388a8a611c7e565b612c429190615cac565b10159350935050505b935093915050565b60008281526101c4602052604081205481905b90508015612c8f57612c8a838261296e876000908152610191602052604090205490565b612404565b509092915050565b6040516001600160a01b03808516602483015283166044820152606481018290526111b09085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526141d1565b60008281526020849052604081208054839290612d39908490615cac565b9250508190555080836001016000828254612d549190615cac565b9091555050505050565b6001600160a01b038416612dda5760405162461bcd60e51b815260206004820152602160248201527f455243313135353a206d696e7420746f20746865207a65726f2061646472657360448201527f730000000000000000000000000000000000000000000000000000000000000060648201526084016109b5565b6000612de4612a8a565b9050612e0581600087612df6886142b6565b612dff886142b6565b87614301565b600084815261015f602090815260408083206001600160a01b038916845290915281208054859290612e38908490615cac565b909155505060408051858152602081018590526001600160a01b0380881692600092918516917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4611164816000878787876144dc565b8151835114612f0f5760405162461bcd60e51b815260206004820152602860248201527f455243313135353a2069647320616e6420616d6f756e7473206c656e6774682060448201527f6d69736d6174636800000000000000000000000000000000000000000000000060648201526084016109b5565b6001600160a01b038416612f735760405162461bcd60e51b815260206004820152602560248201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604482015264647265737360d81b60648201526084016109b5565b6000612f7d612a8a565b9050612f8d818787878787614301565b60005b84518110156130bb576000858281518110612fad57612fad615bbb565b602002602001015190506000858381518110612fcb57612fcb615bbb565b602090810291909101810151600084815261015f835260408082206001600160a01b038e1683529093529190912054909150818110156130605760405162461bcd60e51b815260206004820152602a60248201527f455243313135353a20696e73756666696369656e742062616c616e636520666f60448201526939103a3930b739b332b960b11b60648201526084016109b5565b600083815261015f602090815260408083206001600160a01b038e8116855292528083208585039055908b168252812080548492906130a0908490615cac565b92505081905550505050806130b490615be7565b9050612f90565b50846001600160a01b0316866001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb878760405161310b929190615cc4565b60405180910390a4611bfb818787878787614691565b6040517f02571be30000000000000000000000000000000000000000000000000000000081527f91d1777781884d03a6757a803996e38de2a42967fb37eeaca72729271025a9e260048201526001600160a01b038416906302571be39060240160206040518083038186803b15801561319957600080fd5b505afa1580156131ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131d19190615ce9565b6001600160a01b031663c47f002783836040518363ffffffff1660e01b81526004016131fe929190615d06565b602060405180830381600087803b15801561321857600080fd5b505af115801561322c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111b09190615c50565b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e361327d81610a7a612a8a565b610ad95780610a8c612a8a565b60006132bd7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b90506132c88461479c565b6000835111806132d55750815b156132e6576132e48484614851565b505b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd9143805460ff1661116457805460ff191660011781556040516001600160a01b038316602482015261337e90869060440160408051601f198184030181529190526020810180516001600160e01b03167f3659cfe600000000000000000000000000000000000000000000000000000000179052614851565b50805460ff191681557f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b038381169116146134295760405162461bcd60e51b815260206004820152602f60248201527f45524331393637557067726164653a207570677261646520627265616b73206660448201527f757274686572207570677261646573000000000000000000000000000000000060648201526084016109b5565b61116485614933565b6000613448825167ffffffffffffffff16151590565b80156109e4575050514267ffffffffffffffff909116111590565b6000828152610191602052604081205480156134965760008481526101c66020526040902054612c8a90849083906140e8565b5060009392505050565b600082815260208490526040812080548392906134be908490615c69565b9250508190555080836001016000828254612d549190615c69565b6001600160a01b0383166135555760405162461bcd60e51b815260206004820152602360248201527f455243313135353a206275726e2066726f6d20746865207a65726f206164647260448201527f657373000000000000000000000000000000000000000000000000000000000060648201526084016109b5565b600061355f612a8a565b905061358f81856000613571876142b6565b61357a876142b6565b60405180602001604052806000815250614301565b600083815261015f602090815260408083206001600160a01b0388168452909152902054828110156136285760405162461bcd60e51b8152602060048201526024808201527f455243313135353a206275726e20616d6f756e7420657863656564732062616c60448201527f616e63650000000000000000000000000000000000000000000000000000000060648201526084016109b5565b600084815261015f602090815260408083206001600160a01b03898116808652918452828520888703905582518981529384018890529092908616917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f629101610dcc565b6040516001600160a01b0383166024820152604481018290526136d59084907fa9059cbb0000000000000000000000000000000000000000000000000000000090606401612ce4565b505050565b600054610100900460ff16806136f3575060005460ff16155b6137565760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016109b5565b600054610100900460ff16158015613778576000805461ffff19166101011790555b6137926001600160a01b038316637965db0b60e01b613d1e565b6137d0576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b60448201526064016109b5565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a28015610ad9576000805461ff00191690555050565b600054610100900460ff1680613845575060005460ff16155b6138a85760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016109b5565b600054610100900460ff161580156138ca576000805461ffff19166101011790555b6001600160a01b0382166139215760405163eac0d38960e01b815260206004820152600660248201527f726f75746572000000000000000000000000000000000000000000000000000060448201526064016109b5565b606580546001600160a01b0319166001600160a01b0384169081179091556040517f7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc8090600090a28015610ad9576000805461ff00191690555050565b600054610100900460ff1680613996575060005460ff16155b6139f95760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016109b5565b600054610100900460ff16158015613a1b576000805461ffff19166101011790555b613a23614973565b613a2b614973565b8015611381576000805461ff001916905550565b600054610100900460ff1680613a58575060005460ff16155b613abb5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016109b5565b600054610100900460ff16158015613add576000805461ffff19166101011790555b613ae5614973565b613aed614973565b613af682614a24565b8015610ad9576000805461ff00191690555050565b600054610100900460ff1680613b24575060005460ff16155b613b875760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016109b5565b600054610100900460ff16158015613ba9576000805461ffff19166101011790555b613a1b614973565b60008281526101cb602090815260408083206001600160a01b0385168452909152812054610bfd90613bf4613bef86613bea878961093b565b614acb565b613d3a565b613bfe9190615d35565b614b04565b816001600160a01b0316836001600160a01b03161415613c8b5760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2073657474696e6720617070726f76616c2073746174757360448201527f20666f722073656c66000000000000000000000000000000000000000000000060648201526084016109b5565b6001600160a01b0383811660008181526101606020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6060610bfd8383604051806060016040528060278152602001615f8760279139614b56565b6000613d2983614c20565b8015610bfd5750610bfd8383614c53565b60006001600160ff1b03821115613db95760405162461bcd60e51b815260206004820152602860248201527f53616665436173743a2076616c756520646f65736e27742066697420696e206160448201527f6e20696e7432353600000000000000000000000000000000000000000000000060648201526084016109b5565b5090565b6001600160a01b038216613e165760405163eac0d38960e01b81526004016109b59060208082526004908201527f6d696e7400000000000000000000000000000000000000000000000000000000604082015260600190565b6001600160a01b03821660009081526020849052604081208054839290613e3e908490615d74565b9250508190555080836001016000828254612d549190615d74565b600067ffffffffffffffff821115613db95760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203660448201527f342062697473000000000000000000000000000000000000000000000000000060648201526084016109b5565b6000818310613ee85781610bfd565b5090919050565b6000828152610191602052604081205480156134965760008481526101c46020526040902054612c8a90849083906140e8565b60008281526101c660205260408120548190612c66565b6001600160a01b038416613f9d5760405162461bcd60e51b815260206004820152602560248201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604482015264647265737360d81b60648201526084016109b5565b6000613fa7612a8a565b9050613fb8818787612df6886142b6565b600084815261015f602090815260408083206001600160a01b038a1684529091529020548381101561403f5760405162461bcd60e51b815260206004820152602a60248201527f455243313135353a20696e73756666696369656e742062616c616e636520666f60448201526939103a3930b739b332b960b11b60648201526084016109b5565b600085815261015f602090815260408083206001600160a01b038b811685529252808320878503905590881682528120805486929061407f908490615cac565b909155505060408051868152602081018690526001600160a01b03808916928a821692918616917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a46140df8288888888886144dc565b50505050505050565b60008080600019868509868502925082811083820303915050806000141561411557508390049050610bfd565b80851161414e576040517fc8fea11100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8486850960008681038716968790049687800280890260039081026002908103909202808b028202830302808b028202830302808b028202830302808b028202830302808b0282029092039091029282900382900460010186851190950394909402929094039390930417919091020290509392505050565b6000610ddb614d32565b6000614226826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316614d989092919063ffffffff16565b8051909150156136d557808060200190518101906142449190615b9e565b6136d55760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016109b5565b604080516001808252818301909252606091600091906020808301908036833701905050905082816000815181106142f0576142f0615bbb565b602090810291909101015292915050565b60005b83518110156144cd5761433484828151811061432257614322615bbb565b60200260200101516101009081161490565b1561446857600061437a613bef86848151811061435357614353615bbb565b602002602001015186858151811061436d5761436d615bbb565b6020026020010151614acb565b90506001600160a01b0387166143d0576143cb86826101cb60008987815181106143a6576143a6615bbb565b60200260200101518152602001908152602001600020613dbd9092919063ffffffff16565b614462565b6001600160a01b03861661441f576143cb87826101cb60008987815181106143fa576143fa615bbb565b60200260200101518152602001908152602001600020614da79092919063ffffffff16565b6144628787836101cb60008a888151811061443c5761443c615bbb565b60200260200101518152602001908152602001600020614e43909392919063ffffffff16565b506144bb565b6001600160a01b038616158061448557506001600160a01b038516155b6144bb576040517f068331f900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806144c581615be7565b915050614304565b50611bfb868686868686614f53565b6001600160a01b0384163b15611bfb5760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e61906145209089908990889088908890600401615db3565b602060405180830381600087803b15801561453a57600080fd5b505af192505050801561456a575060408051601f3d908101601f1916820190925261456791810190615deb565b60015b61462057614576615e08565b806308c379a014156145b0575061458b615e23565b8061459657506145b2565b8060405162461bcd60e51b81526004016109b591906154c1565b505b60405162461bcd60e51b815260206004820152603460248201527f455243313135353a207472616e7366657220746f206e6f6e204552433131353560448201527f526563656976657220696d706c656d656e74657200000000000000000000000060648201526084016109b5565b6001600160e01b0319811663f23a6e6160e01b146140df5760405162461bcd60e51b815260206004820152602860248201527f455243313135353a204552433131353552656365697665722072656a656374656044820152676420746f6b656e7360c01b60648201526084016109b5565b6001600160a01b0384163b15611bfb5760405163bc197c8160e01b81526001600160a01b0385169063bc197c81906146d59089908990889088908890600401615ead565b602060405180830381600087803b1580156146ef57600080fd5b505af192505050801561471f575060408051601f3d908101601f1916820190925261471c91810190615deb565b60015b61472b57614576615e08565b6001600160e01b0319811663bc197c8160e01b146140df5760405162461bcd60e51b815260206004820152602860248201527f455243313135353a204552433131353552656365697665722072656a656374656044820152676420746f6b656e7360c01b60648201526084016109b5565b803b6148105760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e74726163740000000000000000000000000000000000000060648201526084016109b5565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b6060823b6148b05760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084016109b5565b600080846001600160a01b0316846040516148cb9190615f0b565b600060405180830381855af49150503d8060008114614906576040519150601f19603f3d011682016040523d82523d6000602084013e61490b565b606091505b5091509150611c638282604051806060016040528060278152602001615f8760279139615061565b61493c8161479c565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b600054610100900460ff168061498c575060005460ff16155b6149ef5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016109b5565b600054610100900460ff16158015613a2b576000805461ffff19166101011790558015611381576000805461ff001916905550565b600054610100900460ff1680614a3d575060005460ff16155b614aa05760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016109b5565b600054610100900460ff16158015614ac2576000805461ffff19166101011790555b613af682612b32565b60008281526101916020526040812054600083118015614aeb5750600081115b614af6576000612404565b612404838261296e8761509a565b600080821215613db95760405162461bcd60e51b815260206004820181905260248201527f53616665436173743a2076616c7565206d75737420626520706f73697469766560448201526064016109b5565b6060833b614bb55760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084016109b5565b600080856001600160a01b031685604051614bd09190615f0b565b600060405180830381855af49150503d8060008114614c0b576040519150601f19603f3d011682016040523d82523d6000602084013e614c10565b606091505b50915091506123a7828286615061565b6000614c33826301ffc9a760e01b614c53565b80156109e45750614c4c826001600160e01b0319614c53565b1592915050565b604080516001600160e01b0319831660248083019190915282518083039091018152604490910182526020810180516001600160e01b03166301ffc9a760e01b179052905160009190829081906001600160a01b0387169061753090614cba908690615f0b565b6000604051808303818686fa925050503d8060008114614cf6576040519150601f19603f3d011682016040523d82523d6000602084013e614cfb565b606091505b5091509150602081511015614d1657600093505050506109e4565b8180156123a75750808060200190518101906123a79190615b9e565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316331415614d9057600036614d73601482615c69565b614d7f92369290615f27565b614d8891615f51565b60601c905090565b503390565b90565b606061240484846000856150d4565b6001600160a01b038216614e005760405163eac0d38960e01b81526004016109b59060208082526004908201527f6275726e00000000000000000000000000000000000000000000000000000000604082015260600190565b6001600160a01b03821660009081526020849052604081208054839290614e28908490615d35565b9250508190555080836001016000828254612d549190615d35565b6001600160a01b038316614e9c5760405163eac0d38960e01b81526004016109b59060208082526004908201527f66726f6d00000000000000000000000000000000000000000000000000000000604082015260600190565b6001600160a01b038216614ef35760405163eac0d38960e01b815260206004820152600260248201527f746f00000000000000000000000000000000000000000000000000000000000060448201526064016109b5565b6001600160a01b03831660009081526020859052604081208054839290614f1b908490615d35565b90915550506001600160a01b03821660009081526020859052604081208054839290614f48908490615d74565b909155505050505050565b6001600160a01b038516614fdb5760005b8351811015614fd957828181518110614f7f57614f7f615bbb565b60200260200101516101916000868481518110614f9e57614f9e615bbb565b602002602001015181526020019081526020016000206000828254614fc39190615cac565b90915550614fd2905081615be7565b9050614f64565b505b6001600160a01b038416611bfb5760005b83518110156140df5782818151811061500757615007615bbb565b6020026020010151610191600086848151811061502657615026615bbb565b60200260200101518152602001908152602001600020600082825461504b9190615c69565b9091555061505a905081615be7565b9050614fec565b60608315615070575081610bfd565b8251156150805782518084602001fd5b8160405162461bcd60e51b81526004016109b591906154c1565b60008181526101cb60205260408120600101546109e49060008481526101c960205260409020546150ca90613d3a565b613bfe9190615d74565b60608247101561514c5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016109b5565b843b61519a5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016109b5565b600080866001600160a01b031685876040516151b69190615f0b565b60006040518083038185875af1925050503d80600081146151f3576040519150601f19603f3d011682016040523d82523d6000602084013e6151f8565b606091505b5091509150615208828286615061565b979650505050505050565b82805461521f90615b63565b90600052602060002090601f0160209004810192826152415760008555615287565b82601f1061525a57805160ff1916838001178555615287565b82800160010185558215615287579182015b8281111561528757825182559160200191906001019061526c565b50613db99291505b80821115613db9576000815560010161528f565b6001600160a01b038116811461138157600080fd5b600080604083850312156152cb57600080fd5b82356152d6816152a3565b946020939093013593505050565b6001600160e01b03198116811461138157600080fd5b60006020828403121561530c57600080fd5b8135610bfd816152e4565b634e487b7160e01b600052604160045260246000fd5b601f8201601f1916810167ffffffffffffffff8111828210171561535357615353615317565b6040525050565b600067ffffffffffffffff83111561537457615374615317565b60405161538b601f8501601f19166020018261532d565b8091508381528484840111156153a057600080fd5b83836020830137600060208583010152509392505050565b6000602082840312156153ca57600080fd5b813567ffffffffffffffff8111156153e157600080fd5b8201601f810184136153f257600080fd5b6124048482356020840161535a565b60006020828403121561541357600080fd5b8135610bfd816152a3565b803560ff8116811461542f57600080fd5b919050565b6000806040838503121561544757600080fd5b6152d68361541e565b60006020828403121561546257600080fd5b5035919050565b60005b8381101561548457818101518382015260200161546c565b838111156111b05750506000910152565b600081518084526154ad816020860160208601615469565b601f01601f19169290920160200192915050565b602081526000610bfd6020830184615495565b801515811461138157600080fd5b6000806000606084860312156154f757600080fd5b6155008461541e565b9250602084013591506040840135615517816154d4565b809150509250925092565b60008060006060848603121561553757600080fd5b6155408461541e565b95602085013595506040909401359392505050565b600067ffffffffffffffff82111561556f5761556f615317565b5060051b60200190565b600082601f83011261558a57600080fd5b8135602061559782615555565b6040516155a4828261532d565b83815260059390931b85018201928281019150868411156155c457600080fd5b8286015b848110156155df57803583529183019183016155c8565b509695505050505050565b600082601f8301126155fb57600080fd5b610bfd8383356020850161535a565b600080600080600060a0868803121561562257600080fd5b853561562d816152a3565b9450602086013561563d816152a3565b9350604086013567ffffffffffffffff8082111561565a57600080fd5b61566689838a01615579565b9450606088013591508082111561567c57600080fd5b61568889838a01615579565b9350608088013591508082111561569e57600080fd5b506156ab888289016155ea565b9150509295509295909350565b6000806000604084860312156156cd57600080fd5b83356156d8816152a3565b9250602084013567ffffffffffffffff808211156156f557600080fd5b818601915086601f83011261570957600080fd5b81358181111561571857600080fd5b87602082850101111561572a57600080fd5b6020830194508093505050509250925092565b6000806040838503121561575057600080fd5b823567ffffffffffffffff8082111561576857600080fd5b818501915085601f83011261577c57600080fd5b8135602061578982615555565b604051615796828261532d565b83815260059390931b85018201928281019150898411156157b657600080fd5b948201945b838610156157dd5785356157ce816152a3565b825294820194908201906157bb565b965050860135925050808211156157f357600080fd5b5061580085828601615579565b9150509250929050565b600081518084526020808501945080840160005b8381101561583a5781518752958201959082019060010161581e565b509495945050505050565b602081526000610bfd602083018461580a565b6000806040838503121561586b57600080fd5b8235615876816152a3565b9150602083013567ffffffffffffffff81111561589257600080fd5b615800858286016155ea565b803567ffffffffffffffff8116811461542f57600080fd5b600080600080600060a086880312156158ce57600080fd5b85356158d9816152a3565b945060208601356158e9816152a3565b935060408601356158f9816152a3565b92506159076060870161589e565b91506080860135615917816152a3565b809150509295509295909350565b60008060006060848603121561593a57600080fd5b6159438461541e565b9250602084013591506040840135615517816152a3565b6000806040838503121561596d57600080fd5b8235615978816152a3565b91506020830135615988816154d4565b809150509250929050565b600080602083850312156159a657600080fd5b823567ffffffffffffffff808211156159be57600080fd5b818501915085601f8301126159d257600080fd5b8135818111156159e157600080fd5b8660208260051b85010111156159f657600080fd5b60209290920196919550909350505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015615a5d57603f19888603018452615a4b858351615495565b94509285019290850190600101615a2f565b5092979650505050505050565b60008060408385031215615a7d57600080fd5b8235615a88816152a3565b91506020830135615988816152a3565b600080600080600060a08688031215615ab057600080fd5b8535945060208601359350615ac76040870161541e565b94979396509394606081013594506080013592915050565b600060208284031215615af157600080fd5b610bfd8261589e565b600080600080600060a08688031215615b1257600080fd5b8535615b1d816152a3565b94506020860135615b2d816152a3565b93506040860135925060608601359150608086013567ffffffffffffffff811115615b5757600080fd5b6156ab888289016155ea565b600181811c90821680615b7757607f821691505b60208210811415615b9857634e487b7160e01b600052602260045260246000fd5b50919050565b600060208284031215615bb057600080fd5b8151610bfd816154d4565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415615bfb57615bfb615bd1565b5060010190565b6000808335601e19843603018112615c1957600080fd5b83018035915067ffffffffffffffff821115615c3457600080fd5b602001915036819003821315615c4957600080fd5b9250929050565b600060208284031215615c6257600080fd5b5051919050565b600082821015615c7b57615c7b615bd1565b500390565b600067ffffffffffffffff808316818516808303821115615ca357615ca3615bd1565b01949350505050565b60008219821115615cbf57615cbf615bd1565b500190565b604081526000615cd7604083018561580a565b8281036020840152611c63818561580a565b600060208284031215615cfb57600080fd5b8151610bfd816152a3565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b600080831283600160ff1b01831281151615615d5357615d53615bd1565b836001600160ff1b03018313811615615d6e57615d6e615bd1565b50500390565b6000808212826001600160ff1b0303841381151615615d9557615d95615bd1565b82600160ff1b038412811615615dad57615dad615bd1565b50500190565b60006001600160a01b03808816835280871660208401525084604083015283606083015260a0608083015261520860a0830184615495565b600060208284031215615dfd57600080fd5b8151610bfd816152e4565b600060033d1115614d955760046000803e5060005160e01c90565b600060443d1015615e315790565b6040516003193d81016004833e81513d67ffffffffffffffff8160248401118184111715615e6157505050505090565b8285019150815181811115615e795750505050505090565b843d8701016020828501011115615e935750505050505090565b615ea26020828601018761532d565b509095945050505050565b60006001600160a01b03808816835280871660208401525060a06040830152615ed960a083018661580a565b8281036060840152615eeb818661580a565b90508281036080840152615eff8185615495565b98975050505050505050565b60008251615f1d818460208701615469565b9190910192915050565b60008085851115615f3757600080fd5b83861115615f4457600080fd5b5050820193919092039150565b6bffffffffffffffffffffffff198135818116916014851015615f7e5780818660140360031b1b83161692505b50509291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220b8d3f371886704a0bca757b18db99d2e890d25e799a95cdb43882a35c89b927464736f6c63430008090033",
}

// FortaStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use FortaStakingMetaData.ABI instead.
var FortaStakingABI = FortaStakingMetaData.ABI

// FortaStakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FortaStakingMetaData.Bin instead.
var FortaStakingBin = FortaStakingMetaData.Bin

// DeployFortaStaking deploys a new Ethereum contract, binding an instance of FortaStaking to it.
func DeployFortaStaking(auth *bind.TransactOpts, backend bind.ContractBackend, forwarder common.Address) (common.Address, *types.Transaction, *FortaStaking, error) {
	parsed, err := FortaStakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FortaStakingBin), backend, forwarder)
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
