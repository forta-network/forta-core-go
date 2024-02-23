// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_scanner_pool_registry_0_1_0

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

// IStakeSubjectStakeThreshold is an auto generated low-level Go binding around an user-defined struct.
type IStakeSubjectStakeThreshold struct {
	Min       *big.Int
	Max       *big.Int
	Activated bool
}

// ScannerPoolRegistryCoreScannerNode is an auto generated low-level Go binding around an user-defined struct.
type ScannerPoolRegistryCoreScannerNode struct {
	Registered    bool
	Disabled      bool
	ScannerPoolId *big.Int
	ChainId       *big.Int
	Metadata      string
}

// ScannerPoolRegistryCoreScannerNodeRegistration is an auto generated low-level Go binding around an user-defined struct.
type ScannerPoolRegistryCoreScannerNodeRegistration struct {
	Scanner       common.Address
	ScannerPoolId *big.Int
	ChainId       *big.Int
	Metadata      string
	Timestamp     *big.Int
}

// ScannerPoolRegistryMetaData contains all meta data concerning the ScannerPoolRegistry contract.
var ScannerPoolRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakeAllocator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"AccessManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"enabledScanners\",\"type\":\"uint256\"}],\"name\":\"EnabledScannersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"activated\",\"type\":\"bool\"}],\"name\":\"ManagedStakeThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"ManagerEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"RegistrationDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"scannerId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"disableFlag\",\"type\":\"bool\"}],\"name\":\"ScannerEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"ScannerPoolRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"scannerId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"scannerPool\",\"type\":\"uint256\"}],\"name\":\"ScannerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newHandler\",\"type\":\"address\"}],\"name\":\"SubjectHandlerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"}],\"name\":\"disableScanner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"}],\"name\":\"enableScanner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"managedId\",\"type\":\"uint256\"}],\"name\":\"getManagedStakeThreshold\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"activated\",\"type\":\"bool\"}],\"internalType\":\"structIStakeSubject.StakeThreshold\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getManagerAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"}],\"name\":\"getManagerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"}],\"name\":\"getScanner\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structScannerPoolRegistryCore.ScannerNode\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"}],\"name\":\"getScannerState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"operational\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSubjectHandler\",\"outputs\":[{\"internalType\":\"contractIStakeSubjectGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"getTotalManagedSubjects\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__manager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"__name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"__symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"__stakeSubjectGateway\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"__registrationDelay\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"}],\"name\":\"isScannerDisabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"}],\"name\":\"isScannerOperational\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"}],\"name\":\"isScannerRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"}],\"name\":\"isScannerRegisteredTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"}],\"name\":\"monitoredChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structScannerPoolRegistryCore.ScannerNodeRegistration\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"name\":\"registerMigratedScannerNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scannerPoolAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"registerMigratedScannerPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structScannerPoolRegistryCore.ScannerNodeRegistration\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"registerScannerNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"registerScannerPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"registeredScannerAddressAtIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"registeredScannerAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structScannerPoolRegistryCore.ScannerNode\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registrationDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"}],\"name\":\"scannerAddressToId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerId\",\"type\":\"uint256\"}],\"name\":\"scannerIdToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"setAccessManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"activated\",\"type\":\"bool\"}],\"internalType\":\"structIStakeSubject.StakeThreshold\",\"name\":\"newStakeThreshold\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"setManagedStakeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ensRegistry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ensName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setRegistrationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subjectGateway\",\"type\":\"address\"}],\"name\":\"setSubjectHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"}],\"name\":\"totalScannersRegistered\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"updateEnabledScanners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"updateScannerMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"}],\"name\":\"willNewScannerShutdownPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e06040523060a0523480156200001557600080fd5b5060405162006053380380620060538339810160408190526200003891620001f2565b6001600160a01b0380831660805281908116620000905760405163eac0d38960e01b815260206004820152601060248201526f2fafb9ba30b5b2a0b63637b1b0ba37b960811b60448201526064015b60405180910390fd5b6001600160a01b031660c052600054610100900460ff1615808015620000bd5750600054600160ff909116105b80620000ed5750620000da30620001c660201b620025101760201c565b158015620000ed575060005460ff166001145b620001525760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840162000087565b6000805460ff19166001179055801562000176576000805461ff0019166101001790555b8015620001bd576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050506200022a565b6001600160a01b03163b151590565b80516001600160a01b0381168114620001ed57600080fd5b919050565b600080604083850312156200020657600080fd5b6200021183620001d5565b91506200022160208401620001d5565b90509250929050565b60805160a05160c051615d93620002c0600039600081816121870152818161222c01528181612dc501528181612e7201528181612f5001528181612fe70152818161307e015281816138640152818161394c01528181613a4a01528181613b130152613c54015260008181611070015281816110b0015281816112ff0152818161133f0152611463015260005050615d936000f3fe6080604052600436106103815760003560e01c8063760d0d39116101d1578063b7fb70f611610102578063d18cc647116100a0578063e873c9dd1161006f578063e873c9dd14610af9578063e985e9c514610b27578063f04dc4b714610b71578063f65cfc9614610b8857600080fd5b8063d18cc64714610a84578063d5c35af414610aa4578063d858a7e514610ac4578063e11cf71e14610ad957600080fd5b8063c80678ef116100dc578063c80678ef14610a04578063c87b56dd14610a24578063c958080414610a44578063ca00928314610a6457600080fd5b8063b7fb70f6146109a6578063b88d4fde146109c6578063c72d33d8146109e657600080fd5b806390b717e21161016f5780639f79b63a116101495780639f79b63a1461091a578063a22cb46514610939578063a9048b5714610959578063ac9650d81461097957600080fd5b806390b717e2146108c5578063911b7b30146108e557806395d89b411461090557600080fd5b806382fe1bcc116101ab57806382fe1bcc1461081a578063841280ac146108415780638e79a369146108855780638fb270e0146108a557600080fd5b8063760d0d39146107a0578063773ed13c146107da5780637dfe7c42146107fa57600080fd5b8063408e35f0116102b657806354fd4d50116102545780636877063a116102235780636877063a1461070e57806370a082311461072e5780637434d8e71461074e57806375b30be61461078057600080fd5b806354fd4d501461067d578063579a6988146106ae5780635a74fc29146106ce5780636352211e146106ee57600080fd5b8063474dd82111610290578063474dd821146106155780634f1ef286146106355780634f6ccce71461064857806352d1902d1461066857600080fd5b8063408e35f0146105b557806342842e0e146105d557806344014b6b146105f557600080fd5b80631d734bf2116103235780632f745c59116102fd5780632f745c59146105165780633121db1c14610536578063331b2c8a146105565780633659cfe61461059557600080fd5b80631d734bf2146104a95780631e93b8e0146104c957806323b872dd146104f657600080fd5b8063095ea7b31161035f578063095ea7b314610415578063113f65c81461043757806318160ddd146104655780631c91a9a41461047b57600080fd5b806301ffc9a71461038657806306fdde03146103bb578063081812fc146103dd575b600080fd5b34801561039257600080fd5b506103a66103a1366004615072565b610ba8565b60405190151581526020015b60405180910390f35b3480156103c757600080fd5b506103d0610bb9565b6040516103b291906150e7565b3480156103e957600080fd5b506103fd6103f83660046150fa565b610c4c565b6040516001600160a01b0390911681526020016103b2565b34801561042157600080fd5b50610435610430366004615128565b610c74565b005b34801561044357600080fd5b506104576104523660046150fa565b610d8f565b6040519081526020016103b2565b34801561047157600080fd5b5061019354610457565b34801561048757600080fd5b506104576104963660046150fa565b60009081526101ff602052604090205490565b3480156104b557600080fd5b506104576104c4366004615128565b610d9b565b3480156104d557600080fd5b506104e96104e4366004615154565b610e0e565b6040516103b29190615176565b34801561050257600080fd5b506104356105113660046151bd565b610f56565b34801561052257600080fd5b50610457610531366004615128565b610f87565b34801561054257600080fd5b50610435610551366004615240565b61101e565b34801561056257600080fd5b506103a6610571366004615295565b6001600160a01b031660009081526101fd6020526040902054610100900460ff1690565b3480156105a157600080fd5b506104356105b0366004615295565b611065565b3480156105c157600080fd5b506104356105d0366004615295565b611145565b3480156105e157600080fd5b506104356105f03660046151bd565b61125d565b34801561060157600080fd5b50610435610610366004615295565b611278565b34801561062157600080fd5b506104356106303660046152d8565b611298565b6104356106433660046153cd565b6112f4565b34801561065457600080fd5b506104576106633660046150fa565b6113c1565b34801561067457600080fd5b50610457611456565b34801561068957600080fd5b506103d0604051806040016040528060058152602001640302e312e360dc1b81525081565b3480156106ba57600080fd5b506103a66106c93660046150fa565b611509565b3480156106da57600080fd5b506104356106e936600461541d565b611529565b3480156106fa57600080fd5b506103fd6107093660046150fa565b611605565b34801561071a57600080fd5b506104e9610729366004615295565b611610565b34801561073a57600080fd5b50610457610749366004615295565b61173d565b34801561075a57600080fd5b5061076e610769366004615295565b6117c4565b6040516103b29695949392919061545f565b34801561078c57600080fd5b5061043561079b3660046154a8565b611824565b3480156107ac57600080fd5b506103a66107bb366004615295565b6001600160a01b031660009081526101fd602052604090205460ff1690565b3480156107e657600080fd5b506103a66107f5366004615548565b6118f2565b34801561080657600080fd5b5061043561081536600461556d565b611912565b34801561082657600080fd5b50610457610835366004615295565b6001600160a01b031690565b34801561084d57600080fd5b5061086161085c3660046150fa565b611ac0565b604080518251815260208084015190820152918101511515908201526060016103b2565b34801561089157600080fd5b506103fd6108a0366004615154565b611b30565b3480156108b157600080fd5b506104576108c03660046150fa565b611b49565b3480156108d157600080fd5b506104356108e03660046150fa565b611b61565b3480156108f157600080fd5b506103a6610900366004615295565b611ba0565b34801561091157600080fd5b506103d0611c32565b34801561092657600080fd5b506101c3546001600160a01b03166103fd565b34801561094557600080fd5b506104356109543660046155c9565b611c42565b34801561096557600080fd5b50610435610974366004615295565b611c4d565b34801561098557600080fd5b506109996109943660046155f7565b611d60565b6040516103b2919061566c565b3480156109b257600080fd5b506104356109c13660046156ce565b611e4e565b3480156109d257600080fd5b506104356109e13660046156ff565b611f5c565b3480156109f257600080fd5b506103fd610a013660046150fa565b90565b348015610a1057600080fd5b506103a6610a1f366004615128565b611f8e565b348015610a3057600080fd5b506103d0610a3f3660046150fa565b611fa7565b348015610a5057600080fd5b50610435610a5f366004615295565b61201a565b348015610a7057600080fd5b50610435610a7f366004615154565b6120d4565b348015610a9057600080fd5b506103a6610a9f3660046150fa565b61215d565b348015610ab057600080fd5b50610435610abf366004615240565b612300565b348015610ad057600080fd5b50610435612454565b348015610ae557600080fd5b50610457610af43660046150fa565b6124df565b348015610b0557600080fd5b50610457610b143660046150fa565b6000908152610201602052604090205490565b348015610b3357600080fd5b506103a6610b4236600461576b565b6001600160a01b0391821660009081526101646020908152604080832093909416825291909152205460ff1690565b348015610b7d57600080fd5b506104576102025481565b348015610b9457600080fd5b506103fd610ba3366004615154565b6124f7565b6000610bb38261251f565b92915050565b606061015f8054610bc990615799565b80601f0160208091040260200160405190810160405280929190818152602001828054610bf590615799565b8015610c425780601f10610c1757610100808354040283529160200191610c42565b820191906000526020600020905b815481529060010190602001808311610c2557829003601f168201915b5050505050905090565b6000610c5782612544565b50600090815261016360205260409020546001600160a01b031690565b6000610c7f826125a4565b9050806001600160a01b0316836001600160a01b03161415610cf25760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b0382161480610d0e5750610d0e8133610b42565b610d805760405162461bcd60e51b815260206004820152603e60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206e6f7220617070726f76656420666f7220616c6c00006064820152608401610ce9565b610d8a838361260a565b505050565b6000610bb33383612679565b60007fcbe0462e67cb804f0011a6c0b71e9ff49be70d0f907ffdf4f3c74499282ab0b1610dc8813361275f565b610dfa5780335b6040516301d4003760e61b815260048101929092526001600160a01b03166024820152604401610ce9565b610e048484612679565b91505b5092915050565b610e446040518060a001604052806000151581526020016000151581526020016000815260200160008152602001606081525090565b60008381526101fe602052604081206101fd9190610e6290856127e4565b6001600160a01b031681526020808201929092526040908101600020815160a081018352815460ff808216151583526101009091041615159381019390935260018101549183019190915260028101546060830152600381018054608084019190610ecc90615799565b80601f0160208091040260200160405190810160405280929190818152602001828054610ef890615799565b8015610f455780601f10610f1a57610100808354040283529160200191610f45565b820191906000526020600020905b815481529060010190602001808311610f2857829003601f168201915b505050505081525050905092915050565b610f6033826127f0565b610f7c5760405162461bcd60e51b8152600401610ce9906157ce565b610d8a83838361286f565b6000610f928361173d565b8210610ff45760405162461bcd60e51b815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201526a74206f6620626f756e647360a81b6064820152608401610ce9565b506001600160a01b0391909116600090815261019160209081526040808320938352929052205490565b7f664245c7af190fec316596e8231f724e8171b1966cfcd124347ac5a66c34f64a611049813361275f565b611054578033610dcf565b61105f848484612a19565b50505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156110ae5760405162461bcd60e51b8152600401610ce99061581c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166110f7600080516020615cf7833981519152546001600160a01b031690565b6001600160a01b03161461111d5760405162461bcd60e51b8152600401610ce990615868565b61112681612b2f565b6040805160008082526020820190925261114291839190612b65565b50565b80611169816001600160a01b031660009081526101fd602052604090205460ff1690565b61119157604051631cc39a1f60e01b81526001600160a01b0382166004820152602401610ce9565b61119a82612cdf565b6111b75760405163a46f5f1160e01b815260040160405180910390fd5b6001600160a01b03821660009081526101fd6020526040902054610100900460ff161561120257604051635fe2c72560e11b81526001600160a01b0383166004820152602401610ce9565b6001600160a01b03821660009081526101fd602052604090206001015461122890612d1a565b6001600160a01b03821660009081526101fd602052604090206001015461124e90612d8e565b6112598260016130b5565b5050565b610d8a83838360405180602001604052806000815250611f5c565b6000611284813361275f565b61128f578033610dcf565b6112598261313c565b7fcbe0462e67cb804f0011a6c0b71e9ff49be70d0f907ffdf4f3c74499282ab0b16112c3813361275f565b6112ce578033610dcf565b6112d7836131cf565b8115610d8a57610d8a6112ed6020850185615295565b60016130b5565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016141561133d5760405162461bcd60e51b8152600401610ce99061581c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316611386600080516020615cf7833981519152546001600160a01b031690565b6001600160a01b0316146113ac5760405162461bcd60e51b8152600401610ce990615868565b6113b582612b2f565b61125982826001612b65565b60006113cd6101935490565b82106114305760405162461bcd60e51b815260206004820152602c60248201527f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60448201526b7574206f6620626f756e647360a01b6064820152608401610ce9565b6101938281548110611444576114446158b4565b90600052602060002001549050919050565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146114f65760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610ce9565b50600080516020615cf783398151915290565b600081815261016160205260408120546001600160a01b03161515610bb3565b8261153381611605565b6001600160a01b0316336001600160a01b03161461157957335b60405163694c4e1560e01b81526001600160a01b03909116600482015260248101829052604401610ce9565b811561159e57600084815261022960205260409020611598908461340b565b506115b9565b6000848152610229602052604090206115b79084613420565b505b826001600160a01b0316847f538b6537a6fe8f0deae9f3b86ad1924d5e5b3d5a683055276b2824f918be043e846040516115f7911515815260200190565b60405180910390a350505050565b6000610bb3826125a4565b6116466040518060a001604052806000151581526020016000151581526020016000815260200160008152602001606081525090565b6001600160a01b03821660009081526101fd6020908152604091829020825160a081018452815460ff8082161515835261010090910416151592810192909252600181015492820192909252600282015460608201526003820180549192916080840191906116b490615799565b80601f01602080910402602001604051908101604052809291908181526020018280546116e090615799565b801561172d5780601f106117025761010080835404028352916020019161172d565b820191906000526020600020905b81548152906001019060200180831161171057829003601f168201915b5050505050815250509050919050565b60006001600160a01b0382166117a75760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b6064820152608401610ce9565b506001600160a01b03166000908152610162602052604090205490565b6000806000606060008060006117d988611610565b8051909150806117ea5760006117f7565b6117f78260400151611605565b8260600151836080015161180a8c611ba0565b602090950151939c929b5090995097509195509350915050565b600054610100900460ff16158080156118445750600054600160ff909116105b8061185e5750303b15801561185e575060005460ff166001145b61187a5760405162461bcd60e51b8152600401610ce9906158ca565b6000805460ff19166001179055801561189d576000805461ff0019166101001790555b6118a688613435565b6118b48787878787876134f7565b80156118e8576000805461ff001916905560405160018152600080516020615d178339815191529060200160405180910390a15b5050505050505050565b60008281526102296020526040812061190b9083613688565b9392505050565b826020013561192081611605565b6001600160a01b0316336001600160a01b03161461193e573361154d565b42610202548560800135611952919061592e565b10156119715760405163e794805160e01b815260040160405180910390fd5b611a8d6119816020860186615295565b611a517f7cb62875b2afadb4cb4ed8910346f6f929b3380857391a94d729cbd47d406af16119b26020890189615295565b602089013560408a01356119c960608c018c615946565b6040516020016119da92919061598d565b604051602081830303815290604052805190602001208b60800135604051602001611a36969594939291909586526001600160a01b0394909416602086015260408501929092526060840152608083015260a082015260c00190565b604051602081830303815290604052805190602001206136aa565b85858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506136f892505050565b611aaa57604051632b84c44f60e11b815260040160405180910390fd5b611ab3846131cf565b61105f8460200135613844565b611ae6604051806060016040528060008152602001600081526020016000151581525090565b506000908152610201602090815260408083205483526102008252918290208251606081018452815481526001820154928101929092526002015460ff1615159181019190915290565b60008281526102296020526040812061190b90836127e4565b60008181526101fe60205260408120610bb390613b80565b7f5f1b4fe6fc447c4d93b2d928a24950315e60be9356ac67fdbfe6aa3d8a985b4d611b8c813361275f565b611b97578033610dcf565b61125982613b8a565b6001600160a01b03811660009081526101fd60209081526040808320600281015484526102009092528220815460ff168015611be357508154610100900460ff16155b8015611c025750600281015460ff161580611c025750611c0284613bfc565b8015611c2a57506001820154600090815261016160205260409020546001600160a01b031615155b949350505050565b60606101608054610bc990615799565b611259338383613cda565b80611c71816001600160a01b031660009081526101fd602052604090205460ff1690565b611c9957604051631cc39a1f60e01b81526001600160a01b0382166004820152602401610ce9565b611ca282612cdf565b611cbf5760405163a46f5f1160e01b815260040160405180910390fd5b6001600160a01b03821660009081526101fd6020526040902054610100900460ff16611d09576040516314d02cb160e31b81526001600160a01b0383166004820152602401610ce9565b6001600160a01b03821660009081526101fd6020526040902060010154611d2f90613da2565b6001600160a01b03821660009081526101fd6020526040902060010154611d5590613844565b6112598260006130b5565b60608167ffffffffffffffff811115611d7b57611d7b61532a565b604051908082528060200260200182016040528015611dae57816020015b6060815260200190600190039081611d995790505b50905060005b82811015610e0757611e1e30858584818110611dd257611dd26158b4565b9050602002810190611de49190615946565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613dc292505050565b828281518110611e3057611e306158b4565b60200260200101819052508080611e469061599d565b915050611db4565b7f5f1b4fe6fc447c4d93b2d928a24950315e60be9356ac67fdbfe6aa3d8a985b4d611e79813361275f565b611e84578033610dcf565b81611ebc576040516303b3e63560e41b815260206004820152600760248201526618da185a5b925960ca1b6044820152606401610ce9565b8235602084013511611ee157604051632ca637fd60e21b815260040160405180910390fd5b817fc89d7e1caf32a74dc9fc2c0197a92afe1936d38c73dd808b0f92fc50677b91ff84356020860135611f1a60608801604089016159b8565b60408051938452602084019290925215159082015260600160405180910390a26000828152610200602052604090208390611f5582826159d5565b5050505050565b611f6633836127f0565b611f825760405162461bcd60e51b8152600401610ce9906157ce565b61105f84848484613de7565b60008181526101fe6020526040812061190b9084613688565b6060611fb282612544565b6000611fc960408051602081019091526000815290565b90506000815111611fe9576040518060200160405280600081525061190b565b80611ff384613e1a565b604051602001612004929190615a0b565b6040516020818303038152906040529392505050565b6000612026813361275f565b612031578033610dcf565b61204b6001600160a01b038316637965db0b60e01b613f18565b612089576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b6044820152606401610ce9565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a25050565b7f5f1b4fe6fc447c4d93b2d928a24950315e60be9356ac67fdbfe6aa3d8a985b4d6120ff813361275f565b61210a578033610dcf565b60008381526101ff6020526040908190208390555183907fe98aca89d957a1075e807d80a8f0f7d280340888edefe3da76c278ca302922c8906121509085815260200190565b60405180910390a2505050565b604051630de12d7d60e31b8152600260048201526024810182905260009081906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690636f096be89060440160206040518083038186803b1580156121c957600080fd5b505afa1580156121dd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122019190615a3a565b604051631837f2f960e31b815260026004820152602481018590529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063c1bf97c89060440160206040518083038186803b15801561226e57600080fd5b505afa158015612282573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122a69190615a3a565b6000858152610201602090815260408083205483526102008252808320548884526101ff909252909120549192509081906122e290600161592e565b6122ec858561592e565b6122f69190615a69565b1095945050505050565b6001600160a01b03831660009081526101fd602052604090205460ff1661234557604051631cc39a1f60e01b81526001600160a01b0384166004820152602401610ce9565b6001600160a01b03831660009081526101fd602052604090206001015461236b90611605565b6001600160a01b0316336001600160a01b0316146123c4576001600160a01b03831660009081526101fd60205260409081902060010154905163694c4e1560e01b81523360048201526024810191909152604401610ce9565b6001600160a01b03831660009081526101fd602052604090206123eb906003018383614f4f565b506001600160a01b03831660008181526101fd602052604090819020600281015460019091015491519092917ffc43d422261aa008210f185c9a82dfeb4faa37fcdac42e357e80bf0a3446e3e891612447918791879190615aa6565b60405180910390a3505050565b6065546001600160a01b03166124a25760405163eac0d38960e01b81526020600482015260126024820152712fb232b83932b1b0ba32b22fb937baba32b960711b6044820152606401610ce9565b606580546001600160a01b03191690556040516000907f7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc80908290a2565b600081815261022960205260408120610bb390613b80565b60008281526101fe6020526040812061190b90836127e4565b6001600160a01b03163b151590565b60006001600160e01b0319821663780e9d6360e01b1480610bb35750610bb382613f34565b600081815261016160205260409020546001600160a01b03166111425760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b6044820152606401610ce9565b600081815261016160205260408120546001600160a01b031680610bb35760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b6044820152606401610ce9565b905090565b60008181526101636020526040902080546001600160a01b0319166001600160a01b0384169081179091558190612640826125a4565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b60006001600160a01b0383166126c75760405163eac0d38960e01b81526020600482015260126024820152717363616e6e6572506f6f6c4164647265737360701b6044820152606401610ce9565b816126ff576040516303b3e63560e41b815260206004820152600760248201526618da185a5b925960ca1b6044820152606401610ce9565b61270e6101fc80546001019055565b506101fc5461271d8382613f84565b6000818152610201602052604080822084905551839183917f7e191d1226c4da936653bac3b9ae5c76dd880cf5221a80bcaa224448ccc670489190a392915050565b603354604051632474521560e21b8152600481018490526001600160a01b03838116602483015260009216906391d148549060440160206040518083038186803b1580156127ac57600080fd5b505afa1580156127c0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061190b9190615aca565b600061190b8383613f9e565b6000806127fc836125a4565b9050806001600160a01b0316846001600160a01b0316148061284457506001600160a01b038082166000908152610164602090815260408083209388168352929052205460ff165b80610e045750836001600160a01b031661285d84610c4c565b6001600160a01b031614949350505050565b826001600160a01b0316612882826125a4565b6001600160a01b0316146128e65760405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201526437bbb732b960d91b6064820152608401610ce9565b6001600160a01b0382166129485760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610ce9565b612953838383613fc8565b61295e60008261260a565b6001600160a01b038316600090815261016260205260408120805460019290612988908490615ae7565b90915550506001600160a01b0382166000908152610162602052604081208054600192906129b790849061592e565b90915550506000818152610161602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6040516302571be360e01b81527f91d1777781884d03a6757a803996e38de2a42967fb37eeaca72729271025a9e260048201526001600160a01b038416906302571be39060240160206040518083038186803b158015612a7857600080fd5b505afa158015612a8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ab09190615afe565b6001600160a01b031663c47f002783836040518363ffffffff1660e01b8152600401612add929190615b1b565b602060405180830381600087803b158015612af757600080fd5b505af1158015612b0b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061105f9190615a3a565b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3612b5a813361275f565b611259578033610dcf565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615612b9857610d8a83613fd3565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b158015612bd157600080fd5b505afa925050508015612c01575060408051601f3d908101601f19168201909252612bfe91810190615a3a565b60015b612c645760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610ce9565b600080516020615cf78339815191528114612cd35760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610ce9565b50610d8a83838361406f565b6000612cea82614094565b80610bb35750610bb37fcbe0462e67cb804f0011a6c0b71e9ff49be70d0f907ffdf4f3c74499282ab0b13361275f565b60008181526101ff60205260408120805460019290612d3a908490615ae7565b909155505060008181526101ff60205260409081902054905182917fe98aca89d957a1075e807d80a8f0f7d280340888edefe3da76c278ca302922c891612d8391815260200190565b60405180910390a250565b60008181526101ff6020526040902054612da55750565b604051631837f2f960e31b815260026004820152602481018290526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063c1bf97c89060440160206040518083038186803b158015612e0f57600080fd5b505afa158015612e23573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e479190615a3a565b604051631837f2f960e31b815260036004820152602481018490529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063c1bf97c89060440160206040518083038186803b158015612eb457600080fd5b505afa158015612ec8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612eec9190615a3a565b90506000612efa828461592e565b90506000612f07856140cc565b9050808211612f17575050505050565b6000612f238284615ae7565b9050808410612fbd576040516328535e1f60e21b81526002600482015260248101879052604481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063a14d787c906064015b600060405180830381600087803b158015612f9d57600080fd5b505af1158015612fb1573d6000803e3d6000fd5b50505050505050505050565b831561305a576040516328535e1f60e21b81526002600482015260248101879052604481018590527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063a14d787c90606401600060405180830381600087803b15801561303357600080fd5b505af1158015613047573d6000803e3d6000fd5b5050505083816130579190615ae7565b90505b6040516313b743c960e21b81526002600482015260248101879052604481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690634edd0f2490606401612f83565b6001600160a01b03821660009081526101fd60205260409020805461ff001916610100831515021790556130e882611ba0565b15156001600160a01b0383167f800131ba226bbb1e9c537607e0f4d73e692ae496e7d1fb9b2aea059369d37fa233604080516001600160a01b03909216825285151560208301520160405180910390a35050565b6001600160a01b0381166131845760405163eac0d38960e01b815260206004820152600e60248201526d7375626a6563744761746577617960901b6044820152606401610ce9565b6101c380546001600160a01b0319166001600160a01b0383169081179091556040517f16d72e484786227d7b6dd05e079be9ff9904a81a0a9ec36fc346b20f8c47aff090600090a250565b6131df6107bb6020830183615295565b15613216576131f16020820182615295565b6040516371cf021b60e01b81526001600160a01b039091166004820152602401610ce9565b80604001356102016000836020013581526020019081526020016000205414613276576020808201356000908152610201909152604090819020548151630432cec160e31b8152610ce99284013590600401918252602082015260400190565b6040518060a0016040528060011515815260200160001515815260200182602001358152602001826040013581526020018280606001906132b79190615946565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525093909452506101fd925061330090506020850185615295565b6001600160a01b031681526020808201929092526040908101600020835181548585015115156101000261ff00199215159290921661ffff19909116171781559083015160018201556060830151600282015560808301518051919261336e92600385019290910190614fd3565b5061339e915061338390506020830183615295565b60208084013560009081526101fe909152604090209061340b565b5060408101356133b46108356020840184615295565b7ffc43d422261aa008210f185c9a82dfeb4faa37fcdac42e357e80bf0a3446e3e86133e26060850185615946565b85602001356040516133f693929190615aa6565b60405180910390a36111428160200135613da2565b600061190b836001600160a01b038416614100565b600061190b836001600160a01b03841661414f565b600054610100900460ff16158080156134555750600054600160ff909116105b8061346f5750303b15801561346f575060005460ff166001145b61348b5760405162461bcd60e51b8152600401610ce9906158ca565b6000805460ff1916600117905580156134ae576000805461ff0019166101001790555b6134b782614242565b6134bf61438a565b8015611259576000805461ff001916905560405160018152600080516020615d17833981519152906020015b60405180910390a15050565b600054610100900460ff16158080156135175750600054600160ff909116105b806135315750303b158015613531575060005460ff166001145b61354d5760405162461bcd60e51b8152600401610ce9906158ca565b6000805460ff191660011790558015613570576000805461ff0019166101001790555b6135e387878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8b0181900481028201810190925289815292508991508890819084018382808284376000920191909152506143b392505050565b6135eb61438a565b613639604051806040016040528060138152602001725363616e6e6572506f6f6c526567697374727960681b815250604051806040016040528060018152602001603160f81b8152506143e4565b61364283614415565b61364b82613b8a565b801561367f576000805461ff001916905560405160018152600080516020615d178339815191529060200160405180910390a15b50505050505050565b6001600160a01b0381166000908152600183016020526040812054151561190b565b6000610bb36136b7614497565b8360405161190160f01b6020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b60008060006137078585614514565b9092509050600081600481111561372057613720615b2f565b14801561373e5750856001600160a01b0316826001600160a01b0316145b1561374e5760019250505061190b565b600080876001600160a01b0316631626ba7e60e01b8888604051602401613776929190615b45565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516137b49190615b5e565b600060405180830381855afa9150503d80600081146137ef576040519150601f19603f3d011682016040523d82523d6000602084013e6137f4565b606091505b5091509150818015613807575080516020145b801561383857508051630b135d3f60e11b9061382c9083016020908101908401615b7a565b6001600160e01b031916145b98975050505050505050565b604051631837f2f960e31b815260026004820152602481018290526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063c1bf97c89060440160206040518083038186803b1580156138ae57600080fd5b505afa1580156138c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138e69190615a3a565b6000838152610201602090815260408083205483526102008252808320548684526101ff909252909120549192509081906139219084615a69565b111561392c57505050565b604051630de12d7d60e31b815260026004820152602481018490526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690636f096be89060440160206040518083038186803b15801561399657600080fd5b505afa1580156139aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139ce9190615a3a565b60008581526101ff602052604090205490915082906139ed858461592e565b6139f79190615a69565b1015613a1557604051626bcc2960e71b815260040160405180910390fd5b6000613a20856140cc565b604051629e4db160e81b815260026004820152602481018790529091506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639e4db1009060440160206040518083038186803b158015613a8c57600080fd5b505afa158015613aa0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613ac49190615a3a565b905081811115613ad657505050505050565b6000613ae28284615ae7565b905083811115613aef5750825b60405163499572af60e01b81526002600482015260248101889052604481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063499572af90606401600060405180830381600087803b158015613b5f57600080fd5b505af1158015613b73573d6000803e3d6000fd5b5050505050505050505050565b6000610bb3825490565b80613bc0576040516303b3e63560e41b815260206004820152600560248201526464656c617960d81b6044820152606401610ce9565b6102028190556040518181527fbb57183c638abde4f4007bb9e16e4dc9fe60ec4e316cc19e9329d9782cabaeac9060200160405180910390a150565b6001600160a01b0381811660009081526101fd6020908152604080832060028082015485526102009093528184208054600183015493516369ce5e8960e11b81526004810195909552602485019390935293949093927f00000000000000000000000000000000000000000000000000000000000000009091169063d39cbd129060440160206040518083038186803b158015613c9857600080fd5b505afa158015613cac573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613cd09190615a3a565b1015949350505050565b816001600160a01b0316836001600160a01b03161415613d3c5760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610ce9565b6001600160a01b0383811660008181526101646020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c319101612447565b60008181526101ff60205260408120805460019290612d3a90849061592e565b606061190b8383604051806060016040528060278152602001615d3760279139614584565b613df284848461286f565b613dfe84848484614622565b61105f5760405162461bcd60e51b8152600401610ce990615b97565b606081613e3e5750506040805180820190915260018152600360fc1b602082015290565b8160005b8115613e685780613e528161599d565b9150613e619050600a83615a69565b9150613e42565b60008167ffffffffffffffff811115613e8357613e8361532a565b6040519080825280601f01601f191660200182016040528015613ead576020820181803683370190505b5090505b8415611c2a57613ec2600183615ae7565b9150613ecf600a86615be9565b613eda90603061592e565b60f81b818381518110613eef57613eef6158b4565b60200101906001600160f81b031916908160001a905350613f11600a86615a69565b9450613eb1565b6000613f238361472f565b801561190b575061190b8383614762565b60006001600160e01b031982166380ac58cd60e01b1480613f6557506001600160e01b03198216635b5e139f60e01b145b80610bb357506301ffc9a760e01b6001600160e01b0319831614610bb3565b611259828260405180602001604052806000815250614841565b6000826000018281548110613fb557613fb56158b4565b9060005260206000200154905092915050565b610d8a838383614874565b6001600160a01b0381163b6140405760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610ce9565b600080516020615cf783398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6140788361492e565b6000825111806140855750805b15610d8a5761105f838361496e565b600061409f82614a23565b80610bb357506001600160a01b03821660009081526101fd6020526040902060010154610bb390336118f2565b60008181526101ff60209081526040808320546102018352818420548452610200909252822060010154610bb39190615bfd565b600081815260018301602052604081205461414757508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610bb3565b506000610bb3565b60008181526001830160205260408120548015614238576000614173600183615ae7565b855490915060009061418790600190615ae7565b90508181146141ec5760008660000182815481106141a7576141a76158b4565b90600052602060002001549050808760000184815481106141ca576141ca6158b4565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806141fd576141fd615c1c565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610bb3565b6000915050610bb3565b600054610100900460ff16158080156142625750600054600160ff909116105b8061427c5750303b15801561427c575060005460ff166001145b6142985760405162461bcd60e51b8152600401610ce9906158ca565b6000805460ff1916600117905580156142bb576000805461ff0019166101001790555b6142d56001600160a01b038316637965db0b60e01b613f18565b614313576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b6044820152606401610ce9565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a28015611259576000805461ff001916905560405160018152600080516020615d17833981519152906020016134eb565b600054610100900460ff166143b15760405162461bcd60e51b8152600401610ce990615c32565b565b600054610100900460ff166143da5760405162461bcd60e51b8152600401610ce990615c32565b6112598282614a6f565b600054610100900460ff1661440b5760405162461bcd60e51b8152600401610ce990615c32565b6112598282614abf565b600054610100900460ff16158080156144355750600054600160ff909116105b8061444f5750303b15801561444f575060005460ff166001145b61446b5760405162461bcd60e51b8152600401610ce9906158ca565b6000805460ff19166001179055801561448e576000805461ff0019166101001790555b6134bf8261313c565b60006126057f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6144c76101c85490565b6101c9546040805160208101859052908101839052606081018290524660808201523060a082015260009060c0016040516020818303038152906040528051906020012090509392505050565b60008082516041141561454b5760208301516040840151606085015160001a61453f87828585614b02565b9450945050505061457d565b825160401415614575576020830151604084015161456a868383614bef565b93509350505061457d565b506000905060025b9250929050565b60606001600160a01b0384163b6145ad5760405162461bcd60e51b8152600401610ce990615c7d565b600080856001600160a01b0316856040516145c89190615b5e565b600060405180830381855af49150503d8060008114614603576040519150601f19603f3d011682016040523d82523d6000602084013e614608565b606091505b5091509150614618828286614c28565b9695505050505050565b60006001600160a01b0384163b1561472457604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290614666903390899088908890600401615cc3565b602060405180830381600087803b15801561468057600080fd5b505af19250505080156146b0575060408051601f3d908101601f191682019092526146ad91810190615b7a565b60015b61470a573d8080156146de576040519150601f19603f3d011682016040523d82523d6000602084013e6146e3565b606091505b5080516147025760405162461bcd60e51b8152600401610ce990615b97565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050611c2a565b506001949350505050565b6000614742826301ffc9a760e01b614762565b8015610bb3575061475b826001600160e01b0319614762565b1592915050565b604080516001600160e01b0319831660248083019190915282518083039091018152604490910182526020810180516001600160e01b03166301ffc9a760e01b179052905160009190829081906001600160a01b03871690617530906147c9908690615b5e565b6000604051808303818686fa925050503d8060008114614805576040519150601f19603f3d011682016040523d82523d6000602084013e61480a565b606091505b50915091506020815110156148255760009350505050610bb3565b8180156146185750808060200190518101906146189190615aca565b61484b8383614c61565b6148586000848484614622565b610d8a5760405162461bcd60e51b8152600401610ce990615b97565b6001600160a01b0383166148d1576148cc816101938054600083815261019460205260408120829055600182018355919091527ffc8af01f449989052b52093a58fc9f42d0b11f0c6dd5dca0463dab62346ccc680155565b6148f4565b816001600160a01b0316836001600160a01b0316146148f4576148f48382614db2565b6001600160a01b03821661490b57610d8a81614e54565b826001600160a01b0316826001600160a01b031614610d8a57610d8a8282614f09565b61493781613fd3565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606001600160a01b0383163b6149975760405162461bcd60e51b8152600401610ce990615c7d565b600080846001600160a01b0316846040516149b29190615b5e565b600060405180830381855af49150503d80600081146149ed576040519150601f19603f3d011682016040523d82523d6000602084013e6149f2565b606091505b5091509150614a1a8282604051806060016040528060278152602001615d3760279139614c28565b95945050505050565b6000336001600160a01b0383161480610bb357506001600160a01b03821660009081526101fd60205260409020600101543390614a5f90611605565b6001600160a01b03161492915050565b600054610100900460ff16614a965760405162461bcd60e51b8152600401610ce990615c32565b8151614aaa9061015f906020850190614fd3565b508051610d8a90610160906020840190614fd3565b600054610100900460ff16614ae65760405162461bcd60e51b8152600401610ce990615c32565b8151602092830120815191909201206101c8919091556101c955565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115614b395750600090506003614be6565b8460ff16601b14158015614b5157508460ff16601c14155b15614b625750600090506004614be6565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015614bb6573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116614bdf57600060019250925050614be6565b9150600090505b94509492505050565b6000806001600160ff1b03831681614c0c60ff86901c601b61592e565b9050614c1a87828885614b02565b935093505050935093915050565b60608315614c3757508161190b565b825115614c475782518084602001fd5b8160405162461bcd60e51b8152600401610ce991906150e7565b6001600160a01b038216614cb75760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610ce9565b600081815261016160205260409020546001600160a01b031615614d1d5760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610ce9565b614d2960008383613fc8565b6001600160a01b038216600090815261016260205260408120805460019290614d5390849061592e565b90915550506000818152610161602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b60006001614dbf8461173d565b614dc99190615ae7565b60008381526101926020526040902054909150808214614e1f576001600160a01b038416600090815261019160209081526040808320858452825280832054848452818420819055835261019290915290208190555b506000918252610192602090815260408084208490556001600160a01b03909416835261019181528383209183525290812055565b61019354600090614e6790600190615ae7565b600083815261019460205260408120546101938054939450909284908110614e9157614e916158b4565b90600052602060002001549050806101938381548110614eb357614eb36158b4565b60009182526020808320909101929092558281526101949091526040808220849055858252812055610193805480614eed57614eed615c1c565b6001900381819060005260206000200160009055905550505050565b6000614f148361173d565b6001600160a01b0390931660009081526101916020908152604080832086845282528083208590559382526101929052919091209190915550565b828054614f5b90615799565b90600052602060002090601f016020900481019282614f7d5760008555614fc3565b82601f10614f965782800160ff19823516178555614fc3565b82800160010185558215614fc3579182015b82811115614fc3578235825591602001919060010190614fa8565b50614fcf929150615047565b5090565b828054614fdf90615799565b90600052602060002090601f0160209004810192826150015760008555614fc3565b82601f1061501a57805160ff1916838001178555614fc3565b82800160010185558215614fc3579182015b82811115614fc357825182559160200191906001019061502c565b5b80821115614fcf5760008155600101615048565b6001600160e01b03198116811461114257600080fd5b60006020828403121561508457600080fd5b813561190b8161505c565b60005b838110156150aa578181015183820152602001615092565b8381111561105f5750506000910152565b600081518084526150d381602086016020860161508f565b601f01601f19169290920160200192915050565b60208152600061190b60208301846150bb565b60006020828403121561510c57600080fd5b5035919050565b6001600160a01b038116811461114257600080fd5b6000806040838503121561513b57600080fd5b823561514681615113565b946020939093013593505050565b6000806040838503121561516757600080fd5b50508035926020909101359150565b6020815281511515602082015260208201511515604082015260408201516060820152606082015160808201526000608083015160a080840152610e0460c08401826150bb565b6000806000606084860312156151d257600080fd5b83356151dd81615113565b925060208401356151ed81615113565b929592945050506040919091013590565b60008083601f84011261521057600080fd5b50813567ffffffffffffffff81111561522857600080fd5b60208301915083602082850101111561457d57600080fd5b60008060006040848603121561525557600080fd5b833561526081615113565b9250602084013567ffffffffffffffff81111561527c57600080fd5b615288868287016151fe565b9497909650939450505050565b6000602082840312156152a757600080fd5b813561190b81615113565b600060a082840312156152c457600080fd5b50919050565b801515811461114257600080fd5b600080604083850312156152eb57600080fd5b823567ffffffffffffffff81111561530257600080fd5b61530e858286016152b2565b925050602083013561531f816152ca565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261535157600080fd5b813567ffffffffffffffff8082111561536c5761536c61532a565b604051601f8301601f19908116603f011681019082821181831017156153945761539461532a565b816040528381528660208588010111156153ad57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080604083850312156153e057600080fd5b82356153eb81615113565b9150602083013567ffffffffffffffff81111561540757600080fd5b61541385828601615340565b9150509250929050565b60008060006060848603121561543257600080fd5b83359250602084013561544481615113565b91506040840135615454816152ca565b809150509250925092565b861515815260018060a01b038616602082015284604082015260c06060820152600061548e60c08301866150bb565b93151560808301525090151560a090910152949350505050565b600080600080600080600060a0888a0312156154c357600080fd5b87356154ce81615113565b9650602088013567ffffffffffffffff808211156154eb57600080fd5b6154f78b838c016151fe565b909850965060408a013591508082111561551057600080fd5b5061551d8a828b016151fe565b909550935050606088013561553181615113565b809250506080880135905092959891949750929550565b6000806040838503121561555b57600080fd5b82359150602083013561531f81615113565b60008060006040848603121561558257600080fd5b833567ffffffffffffffff8082111561559a57600080fd5b6155a6878388016152b2565b945060208601359150808211156155bc57600080fd5b50615288868287016151fe565b600080604083850312156155dc57600080fd5b82356155e781615113565b9150602083013561531f816152ca565b6000806020838503121561560a57600080fd5b823567ffffffffffffffff8082111561562257600080fd5b818501915085601f83011261563657600080fd5b81358181111561564557600080fd5b8660208260051b850101111561565a57600080fd5b60209290920196919550909350505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b828110156156c157603f198886030184526156af8583516150bb565b94509285019290850190600101615693565b5092979650505050505050565b60008082840360808112156156e257600080fd5b60608112156156f057600080fd5b50919360608501359350915050565b6000806000806080858703121561571557600080fd5b843561572081615113565b9350602085013561573081615113565b925060408501359150606085013567ffffffffffffffff81111561575357600080fd5b61575f87828801615340565b91505092959194509250565b6000806040838503121561577e57600080fd5b823561578981615113565b9150602083013561531f81615113565b600181811c908216806157ad57607f821691505b602082108114156152c457634e487b7160e01b600052602260045260246000fd5b6020808252602e908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526d1c881b9bdc88185c1c1c9bdd995960921b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b6000821982111561594157615941615918565b500190565b6000808335601e1984360301811261595d57600080fd5b83018035915067ffffffffffffffff82111561597857600080fd5b60200191503681900382131561457d57600080fd5b8183823760009101908152919050565b60006000198214156159b1576159b1615918565b5060010190565b6000602082840312156159ca57600080fd5b813561190b816152ca565b81358155602082013560018201556002810160408301356159f5816152ca565b815490151560ff1660ff19919091161790555050565b60008351615a1d81846020880161508f565b835190830190615a3181836020880161508f565b01949350505050565b600060208284031215615a4c57600080fd5b5051919050565b634e487b7160e01b600052601260045260246000fd5b600082615a7857615a78615a53565b500490565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b604081526000615aba604083018587615a7d565b9050826020830152949350505050565b600060208284031215615adc57600080fd5b815161190b816152ca565b600082821015615af957615af9615918565b500390565b600060208284031215615b1057600080fd5b815161190b81615113565b602081526000611c2a602083018486615a7d565b634e487b7160e01b600052602160045260246000fd5b828152604060208201526000611c2a60408301846150bb565b60008251615b7081846020870161508f565b9190910192915050565b600060208284031215615b8c57600080fd5b815161190b8161505c565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b600082615bf857615bf8615a53565b500690565b6000816000190483118215151615615c1757615c17615918565b500290565b634e487b7160e01b600052603160045260246000fd5b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b60208082526026908201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6040820152651b9d1c9858dd60d21b606082015260800190565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090614618908301846150bb56fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a264697066735822122053a7dc85ec8e0aadda0f4bfad3b6f8e30ee343e01bf2d2435119ed49171aab8664736f6c63430008090033",
}

// ScannerPoolRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ScannerPoolRegistryMetaData.ABI instead.
var ScannerPoolRegistryABI = ScannerPoolRegistryMetaData.ABI

// ScannerPoolRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ScannerPoolRegistryMetaData.Bin instead.
var ScannerPoolRegistryBin = ScannerPoolRegistryMetaData.Bin

// DeployScannerPoolRegistry deploys a new Ethereum contract, binding an instance of ScannerPoolRegistry to it.
func DeployScannerPoolRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, forwarder common.Address, stakeAllocator common.Address) (common.Address, *types.Transaction, *ScannerPoolRegistry, error) {
	parsed, err := ScannerPoolRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ScannerPoolRegistryBin), backend, forwarder, stakeAllocator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ScannerPoolRegistry{ScannerPoolRegistryCaller: ScannerPoolRegistryCaller{contract: contract}, ScannerPoolRegistryTransactor: ScannerPoolRegistryTransactor{contract: contract}, ScannerPoolRegistryFilterer: ScannerPoolRegistryFilterer{contract: contract}}, nil
}

// ScannerPoolRegistry is an auto generated Go binding around an Ethereum contract.
type ScannerPoolRegistry struct {
	ScannerPoolRegistryCaller     // Read-only binding to the contract
	ScannerPoolRegistryTransactor // Write-only binding to the contract
	ScannerPoolRegistryFilterer   // Log filterer for contract events
}

// ScannerPoolRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ScannerPoolRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScannerPoolRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScannerPoolRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScannerPoolRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScannerPoolRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScannerPoolRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScannerPoolRegistrySession struct {
	Contract     *ScannerPoolRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ScannerPoolRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScannerPoolRegistryCallerSession struct {
	Contract *ScannerPoolRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ScannerPoolRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScannerPoolRegistryTransactorSession struct {
	Contract     *ScannerPoolRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ScannerPoolRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScannerPoolRegistryRaw struct {
	Contract *ScannerPoolRegistry // Generic contract binding to access the raw methods on
}

// ScannerPoolRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScannerPoolRegistryCallerRaw struct {
	Contract *ScannerPoolRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ScannerPoolRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScannerPoolRegistryTransactorRaw struct {
	Contract *ScannerPoolRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScannerPoolRegistry creates a new instance of ScannerPoolRegistry, bound to a specific deployed contract.
func NewScannerPoolRegistry(address common.Address, backend bind.ContractBackend) (*ScannerPoolRegistry, error) {
	contract, err := bindScannerPoolRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistry{ScannerPoolRegistryCaller: ScannerPoolRegistryCaller{contract: contract}, ScannerPoolRegistryTransactor: ScannerPoolRegistryTransactor{contract: contract}, ScannerPoolRegistryFilterer: ScannerPoolRegistryFilterer{contract: contract}}, nil
}

// NewScannerPoolRegistryCaller creates a new read-only instance of ScannerPoolRegistry, bound to a specific deployed contract.
func NewScannerPoolRegistryCaller(address common.Address, caller bind.ContractCaller) (*ScannerPoolRegistryCaller, error) {
	contract, err := bindScannerPoolRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistryCaller{contract: contract}, nil
}

// NewScannerPoolRegistryTransactor creates a new write-only instance of ScannerPoolRegistry, bound to a specific deployed contract.
func NewScannerPoolRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ScannerPoolRegistryTransactor, error) {
	contract, err := bindScannerPoolRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistryTransactor{contract: contract}, nil
}

// NewScannerPoolRegistryFilterer creates a new log filterer instance of ScannerPoolRegistry, bound to a specific deployed contract.
func NewScannerPoolRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ScannerPoolRegistryFilterer, error) {
	contract, err := bindScannerPoolRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistryFilterer{contract: contract}, nil
}

// bindScannerPoolRegistry binds a generic wrapper to an already deployed contract.
func bindScannerPoolRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ScannerPoolRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScannerPoolRegistry *ScannerPoolRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScannerPoolRegistry.Contract.ScannerPoolRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScannerPoolRegistry *ScannerPoolRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.ScannerPoolRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScannerPoolRegistry *ScannerPoolRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.ScannerPoolRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScannerPoolRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ScannerPoolRegistry.Contract.BalanceOf(&_ScannerPoolRegistry.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ScannerPoolRegistry.Contract.BalanceOf(&_ScannerPoolRegistry.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ScannerPoolRegistry.Contract.GetApproved(&_ScannerPoolRegistry.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ScannerPoolRegistry.Contract.GetApproved(&_ScannerPoolRegistry.CallOpts, tokenId)
}

// GetManagedStakeThreshold is a free data retrieval call binding the contract method 0x841280ac.
//
// Solidity: function getManagedStakeThreshold(uint256 managedId) view returns((uint256,uint256,bool))
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) GetManagedStakeThreshold(opts *bind.CallOpts, managedId *big.Int) (IStakeSubjectStakeThreshold, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "getManagedStakeThreshold", managedId)

	if err != nil {
		return *new(IStakeSubjectStakeThreshold), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeSubjectStakeThreshold)).(*IStakeSubjectStakeThreshold)

	return out0, err

}

// GetManagedStakeThreshold is a free data retrieval call binding the contract method 0x841280ac.
//
// Solidity: function getManagedStakeThreshold(uint256 managedId) view returns((uint256,uint256,bool))
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) GetManagedStakeThreshold(managedId *big.Int) (IStakeSubjectStakeThreshold, error) {
	return _ScannerPoolRegistry.Contract.GetManagedStakeThreshold(&_ScannerPoolRegistry.CallOpts, managedId)
}

// GetManagedStakeThreshold is a free data retrieval call binding the contract method 0x841280ac.
//
// Solidity: function getManagedStakeThreshold(uint256 managedId) view returns((uint256,uint256,bool))
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) GetManagedStakeThreshold(managedId *big.Int) (IStakeSubjectStakeThreshold, error) {
	return _ScannerPoolRegistry.Contract.GetManagedStakeThreshold(&_ScannerPoolRegistry.CallOpts, managedId)
}

// GetManagerAt is a free data retrieval call binding the contract method 0x8e79a369.
//
// Solidity: function getManagerAt(uint256 scannerPoolId, uint256 index) view returns(address)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) GetManagerAt(opts *bind.CallOpts, scannerPoolId *big.Int, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "getManagerAt", scannerPoolId, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetManagerAt is a free data retrieval call binding the contract method 0x8e79a369.
//
// Solidity: function getManagerAt(uint256 scannerPoolId, uint256 index) view returns(address)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) GetManagerAt(scannerPoolId *big.Int, index *big.Int) (common.Address, error) {
	return _ScannerPoolRegistry.Contract.GetManagerAt(&_ScannerPoolRegistry.CallOpts, scannerPoolId, index)
}

// GetManagerAt is a free data retrieval call binding the contract method 0x8e79a369.
//
// Solidity: function getManagerAt(uint256 scannerPoolId, uint256 index) view returns(address)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) GetManagerAt(scannerPoolId *big.Int, index *big.Int) (common.Address, error) {
	return _ScannerPoolRegistry.Contract.GetManagerAt(&_ScannerPoolRegistry.CallOpts, scannerPoolId, index)
}

// GetManagerCount is a free data retrieval call binding the contract method 0xe11cf71e.
//
// Solidity: function getManagerCount(uint256 scannerPoolId) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) GetManagerCount(opts *bind.CallOpts, scannerPoolId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "getManagerCount", scannerPoolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetManagerCount is a free data retrieval call binding the contract method 0xe11cf71e.
//
// Solidity: function getManagerCount(uint256 scannerPoolId) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) GetManagerCount(scannerPoolId *big.Int) (*big.Int, error) {
	return _ScannerPoolRegistry.Contract.GetManagerCount(&_ScannerPoolRegistry.CallOpts, scannerPoolId)
}

// GetManagerCount is a free data retrieval call binding the contract method 0xe11cf71e.
//
// Solidity: function getManagerCount(uint256 scannerPoolId) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) GetManagerCount(scannerPoolId *big.Int) (*big.Int, error) {
	return _ScannerPoolRegistry.Contract.GetManagerCount(&_ScannerPoolRegistry.CallOpts, scannerPoolId)
}

// GetScanner is a free data retrieval call binding the contract method 0x6877063a.
//
// Solidity: function getScanner(address scanner) view returns((bool,bool,uint256,uint256,string))
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) GetScanner(opts *bind.CallOpts, scanner common.Address) (ScannerPoolRegistryCoreScannerNode, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "getScanner", scanner)

	if err != nil {
		return *new(ScannerPoolRegistryCoreScannerNode), err
	}

	out0 := *abi.ConvertType(out[0], new(ScannerPoolRegistryCoreScannerNode)).(*ScannerPoolRegistryCoreScannerNode)

	return out0, err

}

// GetScanner is a free data retrieval call binding the contract method 0x6877063a.
//
// Solidity: function getScanner(address scanner) view returns((bool,bool,uint256,uint256,string))
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) GetScanner(scanner common.Address) (ScannerPoolRegistryCoreScannerNode, error) {
	return _ScannerPoolRegistry.Contract.GetScanner(&_ScannerPoolRegistry.CallOpts, scanner)
}

// GetScanner is a free data retrieval call binding the contract method 0x6877063a.
//
// Solidity: function getScanner(address scanner) view returns((bool,bool,uint256,uint256,string))
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) GetScanner(scanner common.Address) (ScannerPoolRegistryCoreScannerNode, error) {
	return _ScannerPoolRegistry.Contract.GetScanner(&_ScannerPoolRegistry.CallOpts, scanner)
}

// GetScannerState is a free data retrieval call binding the contract method 0x7434d8e7.
//
// Solidity: function getScannerState(address scanner) view returns(bool registered, address owner, uint256 chainId, string metadata, bool operational, bool disabled)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) GetScannerState(opts *bind.CallOpts, scanner common.Address) (struct {
	Registered  bool
	Owner       common.Address
	ChainId     *big.Int
	Metadata    string
	Operational bool
	Disabled    bool
}, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "getScannerState", scanner)

	outstruct := new(struct {
		Registered  bool
		Owner       common.Address
		ChainId     *big.Int
		Metadata    string
		Operational bool
		Disabled    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Registered = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ChainId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Metadata = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Operational = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Disabled = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// GetScannerState is a free data retrieval call binding the contract method 0x7434d8e7.
//
// Solidity: function getScannerState(address scanner) view returns(bool registered, address owner, uint256 chainId, string metadata, bool operational, bool disabled)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) GetScannerState(scanner common.Address) (struct {
	Registered  bool
	Owner       common.Address
	ChainId     *big.Int
	Metadata    string
	Operational bool
	Disabled    bool
}, error) {
	return _ScannerPoolRegistry.Contract.GetScannerState(&_ScannerPoolRegistry.CallOpts, scanner)
}

// GetScannerState is a free data retrieval call binding the contract method 0x7434d8e7.
//
// Solidity: function getScannerState(address scanner) view returns(bool registered, address owner, uint256 chainId, string metadata, bool operational, bool disabled)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) GetScannerState(scanner common.Address) (struct {
	Registered  bool
	Owner       common.Address
	ChainId     *big.Int
	Metadata    string
	Operational bool
	Disabled    bool
}, error) {
	return _ScannerPoolRegistry.Contract.GetScannerState(&_ScannerPoolRegistry.CallOpts, scanner)
}

// GetSubjectHandler is a free data retrieval call binding the contract method 0x9f79b63a.
//
// Solidity: function getSubjectHandler() view returns(address)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) GetSubjectHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "getSubjectHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSubjectHandler is a free data retrieval call binding the contract method 0x9f79b63a.
//
// Solidity: function getSubjectHandler() view returns(address)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) GetSubjectHandler() (common.Address, error) {
	return _ScannerPoolRegistry.Contract.GetSubjectHandler(&_ScannerPoolRegistry.CallOpts)
}

// GetSubjectHandler is a free data retrieval call binding the contract method 0x9f79b63a.
//
// Solidity: function getSubjectHandler() view returns(address)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) GetSubjectHandler() (common.Address, error) {
	return _ScannerPoolRegistry.Contract.GetSubjectHandler(&_ScannerPoolRegistry.CallOpts)
}

// GetTotalManagedSubjects is a free data retrieval call binding the contract method 0x1c91a9a4.
//
// Solidity: function getTotalManagedSubjects(uint256 subject) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) GetTotalManagedSubjects(opts *bind.CallOpts, subject *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "getTotalManagedSubjects", subject)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalManagedSubjects is a free data retrieval call binding the contract method 0x1c91a9a4.
//
// Solidity: function getTotalManagedSubjects(uint256 subject) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) GetTotalManagedSubjects(subject *big.Int) (*big.Int, error) {
	return _ScannerPoolRegistry.Contract.GetTotalManagedSubjects(&_ScannerPoolRegistry.CallOpts, subject)
}

// GetTotalManagedSubjects is a free data retrieval call binding the contract method 0x1c91a9a4.
//
// Solidity: function getTotalManagedSubjects(uint256 subject) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) GetTotalManagedSubjects(subject *big.Int) (*big.Int, error) {
	return _ScannerPoolRegistry.Contract.GetTotalManagedSubjects(&_ScannerPoolRegistry.CallOpts, subject)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ScannerPoolRegistry.Contract.IsApprovedForAll(&_ScannerPoolRegistry.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ScannerPoolRegistry.Contract.IsApprovedForAll(&_ScannerPoolRegistry.CallOpts, owner, operator)
}

// IsManager is a free data retrieval call binding the contract method 0x773ed13c.
//
// Solidity: function isManager(uint256 scannerPoolId, address manager) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) IsManager(opts *bind.CallOpts, scannerPoolId *big.Int, manager common.Address) (bool, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "isManager", scannerPoolId, manager)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0x773ed13c.
//
// Solidity: function isManager(uint256 scannerPoolId, address manager) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) IsManager(scannerPoolId *big.Int, manager common.Address) (bool, error) {
	return _ScannerPoolRegistry.Contract.IsManager(&_ScannerPoolRegistry.CallOpts, scannerPoolId, manager)
}

// IsManager is a free data retrieval call binding the contract method 0x773ed13c.
//
// Solidity: function isManager(uint256 scannerPoolId, address manager) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) IsManager(scannerPoolId *big.Int, manager common.Address) (bool, error) {
	return _ScannerPoolRegistry.Contract.IsManager(&_ScannerPoolRegistry.CallOpts, scannerPoolId, manager)
}

// IsRegistered is a free data retrieval call binding the contract method 0x579a6988.
//
// Solidity: function isRegistered(uint256 scannerPoolId) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) IsRegistered(opts *bind.CallOpts, scannerPoolId *big.Int) (bool, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "isRegistered", scannerPoolId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0x579a6988.
//
// Solidity: function isRegistered(uint256 scannerPoolId) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) IsRegistered(scannerPoolId *big.Int) (bool, error) {
	return _ScannerPoolRegistry.Contract.IsRegistered(&_ScannerPoolRegistry.CallOpts, scannerPoolId)
}

// IsRegistered is a free data retrieval call binding the contract method 0x579a6988.
//
// Solidity: function isRegistered(uint256 scannerPoolId) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) IsRegistered(scannerPoolId *big.Int) (bool, error) {
	return _ScannerPoolRegistry.Contract.IsRegistered(&_ScannerPoolRegistry.CallOpts, scannerPoolId)
}

// IsScannerDisabled is a free data retrieval call binding the contract method 0x331b2c8a.
//
// Solidity: function isScannerDisabled(address scanner) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) IsScannerDisabled(opts *bind.CallOpts, scanner common.Address) (bool, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "isScannerDisabled", scanner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsScannerDisabled is a free data retrieval call binding the contract method 0x331b2c8a.
//
// Solidity: function isScannerDisabled(address scanner) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) IsScannerDisabled(scanner common.Address) (bool, error) {
	return _ScannerPoolRegistry.Contract.IsScannerDisabled(&_ScannerPoolRegistry.CallOpts, scanner)
}

// IsScannerDisabled is a free data retrieval call binding the contract method 0x331b2c8a.
//
// Solidity: function isScannerDisabled(address scanner) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) IsScannerDisabled(scanner common.Address) (bool, error) {
	return _ScannerPoolRegistry.Contract.IsScannerDisabled(&_ScannerPoolRegistry.CallOpts, scanner)
}

// IsScannerOperational is a free data retrieval call binding the contract method 0x911b7b30.
//
// Solidity: function isScannerOperational(address scanner) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) IsScannerOperational(opts *bind.CallOpts, scanner common.Address) (bool, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "isScannerOperational", scanner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsScannerOperational is a free data retrieval call binding the contract method 0x911b7b30.
//
// Solidity: function isScannerOperational(address scanner) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) IsScannerOperational(scanner common.Address) (bool, error) {
	return _ScannerPoolRegistry.Contract.IsScannerOperational(&_ScannerPoolRegistry.CallOpts, scanner)
}

// IsScannerOperational is a free data retrieval call binding the contract method 0x911b7b30.
//
// Solidity: function isScannerOperational(address scanner) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) IsScannerOperational(scanner common.Address) (bool, error) {
	return _ScannerPoolRegistry.Contract.IsScannerOperational(&_ScannerPoolRegistry.CallOpts, scanner)
}

// IsScannerRegistered is a free data retrieval call binding the contract method 0x760d0d39.
//
// Solidity: function isScannerRegistered(address scanner) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) IsScannerRegistered(opts *bind.CallOpts, scanner common.Address) (bool, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "isScannerRegistered", scanner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsScannerRegistered is a free data retrieval call binding the contract method 0x760d0d39.
//
// Solidity: function isScannerRegistered(address scanner) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) IsScannerRegistered(scanner common.Address) (bool, error) {
	return _ScannerPoolRegistry.Contract.IsScannerRegistered(&_ScannerPoolRegistry.CallOpts, scanner)
}

// IsScannerRegistered is a free data retrieval call binding the contract method 0x760d0d39.
//
// Solidity: function isScannerRegistered(address scanner) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) IsScannerRegistered(scanner common.Address) (bool, error) {
	return _ScannerPoolRegistry.Contract.IsScannerRegistered(&_ScannerPoolRegistry.CallOpts, scanner)
}

// IsScannerRegisteredTo is a free data retrieval call binding the contract method 0xc80678ef.
//
// Solidity: function isScannerRegisteredTo(address scanner, uint256 scannerPoolId) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) IsScannerRegisteredTo(opts *bind.CallOpts, scanner common.Address, scannerPoolId *big.Int) (bool, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "isScannerRegisteredTo", scanner, scannerPoolId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsScannerRegisteredTo is a free data retrieval call binding the contract method 0xc80678ef.
//
// Solidity: function isScannerRegisteredTo(address scanner, uint256 scannerPoolId) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) IsScannerRegisteredTo(scanner common.Address, scannerPoolId *big.Int) (bool, error) {
	return _ScannerPoolRegistry.Contract.IsScannerRegisteredTo(&_ScannerPoolRegistry.CallOpts, scanner, scannerPoolId)
}

// IsScannerRegisteredTo is a free data retrieval call binding the contract method 0xc80678ef.
//
// Solidity: function isScannerRegisteredTo(address scanner, uint256 scannerPoolId) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) IsScannerRegisteredTo(scanner common.Address, scannerPoolId *big.Int) (bool, error) {
	return _ScannerPoolRegistry.Contract.IsScannerRegisteredTo(&_ScannerPoolRegistry.CallOpts, scanner, scannerPoolId)
}

// MonitoredChainId is a free data retrieval call binding the contract method 0xe873c9dd.
//
// Solidity: function monitoredChainId(uint256 scannerPoolId) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) MonitoredChainId(opts *bind.CallOpts, scannerPoolId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "monitoredChainId", scannerPoolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MonitoredChainId is a free data retrieval call binding the contract method 0xe873c9dd.
//
// Solidity: function monitoredChainId(uint256 scannerPoolId) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) MonitoredChainId(scannerPoolId *big.Int) (*big.Int, error) {
	return _ScannerPoolRegistry.Contract.MonitoredChainId(&_ScannerPoolRegistry.CallOpts, scannerPoolId)
}

// MonitoredChainId is a free data retrieval call binding the contract method 0xe873c9dd.
//
// Solidity: function monitoredChainId(uint256 scannerPoolId) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) MonitoredChainId(scannerPoolId *big.Int) (*big.Int, error) {
	return _ScannerPoolRegistry.Contract.MonitoredChainId(&_ScannerPoolRegistry.CallOpts, scannerPoolId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) Name() (string, error) {
	return _ScannerPoolRegistry.Contract.Name(&_ScannerPoolRegistry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) Name() (string, error) {
	return _ScannerPoolRegistry.Contract.Name(&_ScannerPoolRegistry.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 subject) view returns(address)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) OwnerOf(opts *bind.CallOpts, subject *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "ownerOf", subject)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 subject) view returns(address)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) OwnerOf(subject *big.Int) (common.Address, error) {
	return _ScannerPoolRegistry.Contract.OwnerOf(&_ScannerPoolRegistry.CallOpts, subject)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 subject) view returns(address)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) OwnerOf(subject *big.Int) (common.Address, error) {
	return _ScannerPoolRegistry.Contract.OwnerOf(&_ScannerPoolRegistry.CallOpts, subject)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _ScannerPoolRegistry.Contract.ProxiableUUID(&_ScannerPoolRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ScannerPoolRegistry.Contract.ProxiableUUID(&_ScannerPoolRegistry.CallOpts)
}

// RegisteredScannerAddressAtIndex is a free data retrieval call binding the contract method 0xf65cfc96.
//
// Solidity: function registeredScannerAddressAtIndex(uint256 scannerPoolId, uint256 index) view returns(address)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) RegisteredScannerAddressAtIndex(opts *bind.CallOpts, scannerPoolId *big.Int, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "registeredScannerAddressAtIndex", scannerPoolId, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegisteredScannerAddressAtIndex is a free data retrieval call binding the contract method 0xf65cfc96.
//
// Solidity: function registeredScannerAddressAtIndex(uint256 scannerPoolId, uint256 index) view returns(address)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) RegisteredScannerAddressAtIndex(scannerPoolId *big.Int, index *big.Int) (common.Address, error) {
	return _ScannerPoolRegistry.Contract.RegisteredScannerAddressAtIndex(&_ScannerPoolRegistry.CallOpts, scannerPoolId, index)
}

// RegisteredScannerAddressAtIndex is a free data retrieval call binding the contract method 0xf65cfc96.
//
// Solidity: function registeredScannerAddressAtIndex(uint256 scannerPoolId, uint256 index) view returns(address)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) RegisteredScannerAddressAtIndex(scannerPoolId *big.Int, index *big.Int) (common.Address, error) {
	return _ScannerPoolRegistry.Contract.RegisteredScannerAddressAtIndex(&_ScannerPoolRegistry.CallOpts, scannerPoolId, index)
}

// RegisteredScannerAtIndex is a free data retrieval call binding the contract method 0x1e93b8e0.
//
// Solidity: function registeredScannerAtIndex(uint256 scannerPoolId, uint256 index) view returns((bool,bool,uint256,uint256,string))
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) RegisteredScannerAtIndex(opts *bind.CallOpts, scannerPoolId *big.Int, index *big.Int) (ScannerPoolRegistryCoreScannerNode, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "registeredScannerAtIndex", scannerPoolId, index)

	if err != nil {
		return *new(ScannerPoolRegistryCoreScannerNode), err
	}

	out0 := *abi.ConvertType(out[0], new(ScannerPoolRegistryCoreScannerNode)).(*ScannerPoolRegistryCoreScannerNode)

	return out0, err

}

// RegisteredScannerAtIndex is a free data retrieval call binding the contract method 0x1e93b8e0.
//
// Solidity: function registeredScannerAtIndex(uint256 scannerPoolId, uint256 index) view returns((bool,bool,uint256,uint256,string))
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) RegisteredScannerAtIndex(scannerPoolId *big.Int, index *big.Int) (ScannerPoolRegistryCoreScannerNode, error) {
	return _ScannerPoolRegistry.Contract.RegisteredScannerAtIndex(&_ScannerPoolRegistry.CallOpts, scannerPoolId, index)
}

// RegisteredScannerAtIndex is a free data retrieval call binding the contract method 0x1e93b8e0.
//
// Solidity: function registeredScannerAtIndex(uint256 scannerPoolId, uint256 index) view returns((bool,bool,uint256,uint256,string))
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) RegisteredScannerAtIndex(scannerPoolId *big.Int, index *big.Int) (ScannerPoolRegistryCoreScannerNode, error) {
	return _ScannerPoolRegistry.Contract.RegisteredScannerAtIndex(&_ScannerPoolRegistry.CallOpts, scannerPoolId, index)
}

// RegistrationDelay is a free data retrieval call binding the contract method 0xf04dc4b7.
//
// Solidity: function registrationDelay() view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) RegistrationDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "registrationDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegistrationDelay is a free data retrieval call binding the contract method 0xf04dc4b7.
//
// Solidity: function registrationDelay() view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) RegistrationDelay() (*big.Int, error) {
	return _ScannerPoolRegistry.Contract.RegistrationDelay(&_ScannerPoolRegistry.CallOpts)
}

// RegistrationDelay is a free data retrieval call binding the contract method 0xf04dc4b7.
//
// Solidity: function registrationDelay() view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) RegistrationDelay() (*big.Int, error) {
	return _ScannerPoolRegistry.Contract.RegistrationDelay(&_ScannerPoolRegistry.CallOpts)
}

// ScannerAddressToId is a free data retrieval call binding the contract method 0x82fe1bcc.
//
// Solidity: function scannerAddressToId(address scanner) pure returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) ScannerAddressToId(opts *bind.CallOpts, scanner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "scannerAddressToId", scanner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScannerAddressToId is a free data retrieval call binding the contract method 0x82fe1bcc.
//
// Solidity: function scannerAddressToId(address scanner) pure returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) ScannerAddressToId(scanner common.Address) (*big.Int, error) {
	return _ScannerPoolRegistry.Contract.ScannerAddressToId(&_ScannerPoolRegistry.CallOpts, scanner)
}

// ScannerAddressToId is a free data retrieval call binding the contract method 0x82fe1bcc.
//
// Solidity: function scannerAddressToId(address scanner) pure returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) ScannerAddressToId(scanner common.Address) (*big.Int, error) {
	return _ScannerPoolRegistry.Contract.ScannerAddressToId(&_ScannerPoolRegistry.CallOpts, scanner)
}

// ScannerIdToAddress is a free data retrieval call binding the contract method 0xc72d33d8.
//
// Solidity: function scannerIdToAddress(uint256 scannerId) pure returns(address)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) ScannerIdToAddress(opts *bind.CallOpts, scannerId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "scannerIdToAddress", scannerId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ScannerIdToAddress is a free data retrieval call binding the contract method 0xc72d33d8.
//
// Solidity: function scannerIdToAddress(uint256 scannerId) pure returns(address)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) ScannerIdToAddress(scannerId *big.Int) (common.Address, error) {
	return _ScannerPoolRegistry.Contract.ScannerIdToAddress(&_ScannerPoolRegistry.CallOpts, scannerId)
}

// ScannerIdToAddress is a free data retrieval call binding the contract method 0xc72d33d8.
//
// Solidity: function scannerIdToAddress(uint256 scannerId) pure returns(address)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) ScannerIdToAddress(scannerId *big.Int) (common.Address, error) {
	return _ScannerPoolRegistry.Contract.ScannerIdToAddress(&_ScannerPoolRegistry.CallOpts, scannerId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ScannerPoolRegistry.Contract.SupportsInterface(&_ScannerPoolRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ScannerPoolRegistry.Contract.SupportsInterface(&_ScannerPoolRegistry.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) Symbol() (string, error) {
	return _ScannerPoolRegistry.Contract.Symbol(&_ScannerPoolRegistry.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) Symbol() (string, error) {
	return _ScannerPoolRegistry.Contract.Symbol(&_ScannerPoolRegistry.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ScannerPoolRegistry.Contract.TokenByIndex(&_ScannerPoolRegistry.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ScannerPoolRegistry.Contract.TokenByIndex(&_ScannerPoolRegistry.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ScannerPoolRegistry.Contract.TokenOfOwnerByIndex(&_ScannerPoolRegistry.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ScannerPoolRegistry.Contract.TokenOfOwnerByIndex(&_ScannerPoolRegistry.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) TokenURI(tokenId *big.Int) (string, error) {
	return _ScannerPoolRegistry.Contract.TokenURI(&_ScannerPoolRegistry.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ScannerPoolRegistry.Contract.TokenURI(&_ScannerPoolRegistry.CallOpts, tokenId)
}

// TotalScannersRegistered is a free data retrieval call binding the contract method 0x8fb270e0.
//
// Solidity: function totalScannersRegistered(uint256 scannerPoolId) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) TotalScannersRegistered(opts *bind.CallOpts, scannerPoolId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "totalScannersRegistered", scannerPoolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalScannersRegistered is a free data retrieval call binding the contract method 0x8fb270e0.
//
// Solidity: function totalScannersRegistered(uint256 scannerPoolId) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) TotalScannersRegistered(scannerPoolId *big.Int) (*big.Int, error) {
	return _ScannerPoolRegistry.Contract.TotalScannersRegistered(&_ScannerPoolRegistry.CallOpts, scannerPoolId)
}

// TotalScannersRegistered is a free data retrieval call binding the contract method 0x8fb270e0.
//
// Solidity: function totalScannersRegistered(uint256 scannerPoolId) view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) TotalScannersRegistered(scannerPoolId *big.Int) (*big.Int, error) {
	return _ScannerPoolRegistry.Contract.TotalScannersRegistered(&_ScannerPoolRegistry.CallOpts, scannerPoolId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) TotalSupply() (*big.Int, error) {
	return _ScannerPoolRegistry.Contract.TotalSupply(&_ScannerPoolRegistry.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) TotalSupply() (*big.Int, error) {
	return _ScannerPoolRegistry.Contract.TotalSupply(&_ScannerPoolRegistry.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) Version() (string, error) {
	return _ScannerPoolRegistry.Contract.Version(&_ScannerPoolRegistry.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) Version() (string, error) {
	return _ScannerPoolRegistry.Contract.Version(&_ScannerPoolRegistry.CallOpts)
}

// WillNewScannerShutdownPool is a free data retrieval call binding the contract method 0xd18cc647.
//
// Solidity: function willNewScannerShutdownPool(uint256 scannerPoolId) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) WillNewScannerShutdownPool(opts *bind.CallOpts, scannerPoolId *big.Int) (bool, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "willNewScannerShutdownPool", scannerPoolId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WillNewScannerShutdownPool is a free data retrieval call binding the contract method 0xd18cc647.
//
// Solidity: function willNewScannerShutdownPool(uint256 scannerPoolId) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) WillNewScannerShutdownPool(scannerPoolId *big.Int) (bool, error) {
	return _ScannerPoolRegistry.Contract.WillNewScannerShutdownPool(&_ScannerPoolRegistry.CallOpts, scannerPoolId)
}

// WillNewScannerShutdownPool is a free data retrieval call binding the contract method 0xd18cc647.
//
// Solidity: function willNewScannerShutdownPool(uint256 scannerPoolId) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) WillNewScannerShutdownPool(scannerPoolId *big.Int) (bool, error) {
	return _ScannerPoolRegistry.Contract.WillNewScannerShutdownPool(&_ScannerPoolRegistry.CallOpts, scannerPoolId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.Approve(&_ScannerPoolRegistry.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.Approve(&_ScannerPoolRegistry.TransactOpts, to, tokenId)
}

// DisableRouter is a paid mutator transaction binding the contract method 0xd858a7e5.
//
// Solidity: function disableRouter() returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) DisableRouter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "disableRouter")
}

// DisableRouter is a paid mutator transaction binding the contract method 0xd858a7e5.
//
// Solidity: function disableRouter() returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) DisableRouter() (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.DisableRouter(&_ScannerPoolRegistry.TransactOpts)
}

// DisableRouter is a paid mutator transaction binding the contract method 0xd858a7e5.
//
// Solidity: function disableRouter() returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) DisableRouter() (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.DisableRouter(&_ScannerPoolRegistry.TransactOpts)
}

// DisableScanner is a paid mutator transaction binding the contract method 0x408e35f0.
//
// Solidity: function disableScanner(address scanner) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) DisableScanner(opts *bind.TransactOpts, scanner common.Address) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "disableScanner", scanner)
}

// DisableScanner is a paid mutator transaction binding the contract method 0x408e35f0.
//
// Solidity: function disableScanner(address scanner) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) DisableScanner(scanner common.Address) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.DisableScanner(&_ScannerPoolRegistry.TransactOpts, scanner)
}

// DisableScanner is a paid mutator transaction binding the contract method 0x408e35f0.
//
// Solidity: function disableScanner(address scanner) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) DisableScanner(scanner common.Address) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.DisableScanner(&_ScannerPoolRegistry.TransactOpts, scanner)
}

// EnableScanner is a paid mutator transaction binding the contract method 0xa9048b57.
//
// Solidity: function enableScanner(address scanner) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) EnableScanner(opts *bind.TransactOpts, scanner common.Address) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "enableScanner", scanner)
}

// EnableScanner is a paid mutator transaction binding the contract method 0xa9048b57.
//
// Solidity: function enableScanner(address scanner) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) EnableScanner(scanner common.Address) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.EnableScanner(&_ScannerPoolRegistry.TransactOpts, scanner)
}

// EnableScanner is a paid mutator transaction binding the contract method 0xa9048b57.
//
// Solidity: function enableScanner(address scanner) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) EnableScanner(scanner common.Address) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.EnableScanner(&_ScannerPoolRegistry.TransactOpts, scanner)
}

// Initialize is a paid mutator transaction binding the contract method 0x75b30be6.
//
// Solidity: function initialize(address __manager, string __name, string __symbol, address __stakeSubjectGateway, uint256 __registrationDelay) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) Initialize(opts *bind.TransactOpts, __manager common.Address, __name string, __symbol string, __stakeSubjectGateway common.Address, __registrationDelay *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "initialize", __manager, __name, __symbol, __stakeSubjectGateway, __registrationDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0x75b30be6.
//
// Solidity: function initialize(address __manager, string __name, string __symbol, address __stakeSubjectGateway, uint256 __registrationDelay) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) Initialize(__manager common.Address, __name string, __symbol string, __stakeSubjectGateway common.Address, __registrationDelay *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.Initialize(&_ScannerPoolRegistry.TransactOpts, __manager, __name, __symbol, __stakeSubjectGateway, __registrationDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0x75b30be6.
//
// Solidity: function initialize(address __manager, string __name, string __symbol, address __stakeSubjectGateway, uint256 __registrationDelay) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) Initialize(__manager common.Address, __name string, __symbol string, __stakeSubjectGateway common.Address, __registrationDelay *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.Initialize(&_ScannerPoolRegistry.TransactOpts, __manager, __name, __symbol, __stakeSubjectGateway, __registrationDelay)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.Multicall(&_ScannerPoolRegistry.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.Multicall(&_ScannerPoolRegistry.TransactOpts, data)
}

// RegisterMigratedScannerNode is a paid mutator transaction binding the contract method 0x474dd821.
//
// Solidity: function registerMigratedScannerNode((address,uint256,uint256,string,uint256) req, bool disabled) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) RegisterMigratedScannerNode(opts *bind.TransactOpts, req ScannerPoolRegistryCoreScannerNodeRegistration, disabled bool) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "registerMigratedScannerNode", req, disabled)
}

// RegisterMigratedScannerNode is a paid mutator transaction binding the contract method 0x474dd821.
//
// Solidity: function registerMigratedScannerNode((address,uint256,uint256,string,uint256) req, bool disabled) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) RegisterMigratedScannerNode(req ScannerPoolRegistryCoreScannerNodeRegistration, disabled bool) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.RegisterMigratedScannerNode(&_ScannerPoolRegistry.TransactOpts, req, disabled)
}

// RegisterMigratedScannerNode is a paid mutator transaction binding the contract method 0x474dd821.
//
// Solidity: function registerMigratedScannerNode((address,uint256,uint256,string,uint256) req, bool disabled) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) RegisterMigratedScannerNode(req ScannerPoolRegistryCoreScannerNodeRegistration, disabled bool) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.RegisterMigratedScannerNode(&_ScannerPoolRegistry.TransactOpts, req, disabled)
}

// RegisterMigratedScannerPool is a paid mutator transaction binding the contract method 0x1d734bf2.
//
// Solidity: function registerMigratedScannerPool(address scannerPoolAddress, uint256 chainId) returns(uint256 scannerPoolId)
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) RegisterMigratedScannerPool(opts *bind.TransactOpts, scannerPoolAddress common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "registerMigratedScannerPool", scannerPoolAddress, chainId)
}

// RegisterMigratedScannerPool is a paid mutator transaction binding the contract method 0x1d734bf2.
//
// Solidity: function registerMigratedScannerPool(address scannerPoolAddress, uint256 chainId) returns(uint256 scannerPoolId)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) RegisterMigratedScannerPool(scannerPoolAddress common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.RegisterMigratedScannerPool(&_ScannerPoolRegistry.TransactOpts, scannerPoolAddress, chainId)
}

// RegisterMigratedScannerPool is a paid mutator transaction binding the contract method 0x1d734bf2.
//
// Solidity: function registerMigratedScannerPool(address scannerPoolAddress, uint256 chainId) returns(uint256 scannerPoolId)
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) RegisterMigratedScannerPool(scannerPoolAddress common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.RegisterMigratedScannerPool(&_ScannerPoolRegistry.TransactOpts, scannerPoolAddress, chainId)
}

// RegisterScannerNode is a paid mutator transaction binding the contract method 0x7dfe7c42.
//
// Solidity: function registerScannerNode((address,uint256,uint256,string,uint256) req, bytes signature) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) RegisterScannerNode(opts *bind.TransactOpts, req ScannerPoolRegistryCoreScannerNodeRegistration, signature []byte) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "registerScannerNode", req, signature)
}

// RegisterScannerNode is a paid mutator transaction binding the contract method 0x7dfe7c42.
//
// Solidity: function registerScannerNode((address,uint256,uint256,string,uint256) req, bytes signature) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) RegisterScannerNode(req ScannerPoolRegistryCoreScannerNodeRegistration, signature []byte) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.RegisterScannerNode(&_ScannerPoolRegistry.TransactOpts, req, signature)
}

// RegisterScannerNode is a paid mutator transaction binding the contract method 0x7dfe7c42.
//
// Solidity: function registerScannerNode((address,uint256,uint256,string,uint256) req, bytes signature) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) RegisterScannerNode(req ScannerPoolRegistryCoreScannerNodeRegistration, signature []byte) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.RegisterScannerNode(&_ScannerPoolRegistry.TransactOpts, req, signature)
}

// RegisterScannerPool is a paid mutator transaction binding the contract method 0x113f65c8.
//
// Solidity: function registerScannerPool(uint256 chainId) returns(uint256 scannerPoolId)
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) RegisterScannerPool(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "registerScannerPool", chainId)
}

// RegisterScannerPool is a paid mutator transaction binding the contract method 0x113f65c8.
//
// Solidity: function registerScannerPool(uint256 chainId) returns(uint256 scannerPoolId)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) RegisterScannerPool(chainId *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.RegisterScannerPool(&_ScannerPoolRegistry.TransactOpts, chainId)
}

// RegisterScannerPool is a paid mutator transaction binding the contract method 0x113f65c8.
//
// Solidity: function registerScannerPool(uint256 chainId) returns(uint256 scannerPoolId)
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) RegisterScannerPool(chainId *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.RegisterScannerPool(&_ScannerPoolRegistry.TransactOpts, chainId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.SafeTransferFrom(&_ScannerPoolRegistry.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.SafeTransferFrom(&_ScannerPoolRegistry.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.SafeTransferFrom0(&_ScannerPoolRegistry.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.SafeTransferFrom0(&_ScannerPoolRegistry.TransactOpts, from, to, tokenId, data)
}

// SetAccessManager is a paid mutator transaction binding the contract method 0xc9580804.
//
// Solidity: function setAccessManager(address newManager) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) SetAccessManager(opts *bind.TransactOpts, newManager common.Address) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "setAccessManager", newManager)
}

// SetAccessManager is a paid mutator transaction binding the contract method 0xc9580804.
//
// Solidity: function setAccessManager(address newManager) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) SetAccessManager(newManager common.Address) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.SetAccessManager(&_ScannerPoolRegistry.TransactOpts, newManager)
}

// SetAccessManager is a paid mutator transaction binding the contract method 0xc9580804.
//
// Solidity: function setAccessManager(address newManager) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) SetAccessManager(newManager common.Address) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.SetAccessManager(&_ScannerPoolRegistry.TransactOpts, newManager)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.SetApprovalForAll(&_ScannerPoolRegistry.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.SetApprovalForAll(&_ScannerPoolRegistry.TransactOpts, operator, approved)
}

// SetManagedStakeThreshold is a paid mutator transaction binding the contract method 0xb7fb70f6.
//
// Solidity: function setManagedStakeThreshold((uint256,uint256,bool) newStakeThreshold, uint256 chainId) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) SetManagedStakeThreshold(opts *bind.TransactOpts, newStakeThreshold IStakeSubjectStakeThreshold, chainId *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "setManagedStakeThreshold", newStakeThreshold, chainId)
}

// SetManagedStakeThreshold is a paid mutator transaction binding the contract method 0xb7fb70f6.
//
// Solidity: function setManagedStakeThreshold((uint256,uint256,bool) newStakeThreshold, uint256 chainId) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) SetManagedStakeThreshold(newStakeThreshold IStakeSubjectStakeThreshold, chainId *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.SetManagedStakeThreshold(&_ScannerPoolRegistry.TransactOpts, newStakeThreshold, chainId)
}

// SetManagedStakeThreshold is a paid mutator transaction binding the contract method 0xb7fb70f6.
//
// Solidity: function setManagedStakeThreshold((uint256,uint256,bool) newStakeThreshold, uint256 chainId) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) SetManagedStakeThreshold(newStakeThreshold IStakeSubjectStakeThreshold, chainId *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.SetManagedStakeThreshold(&_ScannerPoolRegistry.TransactOpts, newStakeThreshold, chainId)
}

// SetManager is a paid mutator transaction binding the contract method 0x5a74fc29.
//
// Solidity: function setManager(uint256 scannerPoolId, address manager, bool enable) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) SetManager(opts *bind.TransactOpts, scannerPoolId *big.Int, manager common.Address, enable bool) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "setManager", scannerPoolId, manager, enable)
}

// SetManager is a paid mutator transaction binding the contract method 0x5a74fc29.
//
// Solidity: function setManager(uint256 scannerPoolId, address manager, bool enable) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) SetManager(scannerPoolId *big.Int, manager common.Address, enable bool) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.SetManager(&_ScannerPoolRegistry.TransactOpts, scannerPoolId, manager, enable)
}

// SetManager is a paid mutator transaction binding the contract method 0x5a74fc29.
//
// Solidity: function setManager(uint256 scannerPoolId, address manager, bool enable) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) SetManager(scannerPoolId *big.Int, manager common.Address, enable bool) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.SetManager(&_ScannerPoolRegistry.TransactOpts, scannerPoolId, manager, enable)
}

// SetName is a paid mutator transaction binding the contract method 0x3121db1c.
//
// Solidity: function setName(address ensRegistry, string ensName) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) SetName(opts *bind.TransactOpts, ensRegistry common.Address, ensName string) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "setName", ensRegistry, ensName)
}

// SetName is a paid mutator transaction binding the contract method 0x3121db1c.
//
// Solidity: function setName(address ensRegistry, string ensName) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) SetName(ensRegistry common.Address, ensName string) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.SetName(&_ScannerPoolRegistry.TransactOpts, ensRegistry, ensName)
}

// SetName is a paid mutator transaction binding the contract method 0x3121db1c.
//
// Solidity: function setName(address ensRegistry, string ensName) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) SetName(ensRegistry common.Address, ensName string) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.SetName(&_ScannerPoolRegistry.TransactOpts, ensRegistry, ensName)
}

// SetRegistrationDelay is a paid mutator transaction binding the contract method 0x90b717e2.
//
// Solidity: function setRegistrationDelay(uint256 delay) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) SetRegistrationDelay(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "setRegistrationDelay", delay)
}

// SetRegistrationDelay is a paid mutator transaction binding the contract method 0x90b717e2.
//
// Solidity: function setRegistrationDelay(uint256 delay) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) SetRegistrationDelay(delay *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.SetRegistrationDelay(&_ScannerPoolRegistry.TransactOpts, delay)
}

// SetRegistrationDelay is a paid mutator transaction binding the contract method 0x90b717e2.
//
// Solidity: function setRegistrationDelay(uint256 delay) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) SetRegistrationDelay(delay *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.SetRegistrationDelay(&_ScannerPoolRegistry.TransactOpts, delay)
}

// SetSubjectHandler is a paid mutator transaction binding the contract method 0x44014b6b.
//
// Solidity: function setSubjectHandler(address subjectGateway) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) SetSubjectHandler(opts *bind.TransactOpts, subjectGateway common.Address) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "setSubjectHandler", subjectGateway)
}

// SetSubjectHandler is a paid mutator transaction binding the contract method 0x44014b6b.
//
// Solidity: function setSubjectHandler(address subjectGateway) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) SetSubjectHandler(subjectGateway common.Address) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.SetSubjectHandler(&_ScannerPoolRegistry.TransactOpts, subjectGateway)
}

// SetSubjectHandler is a paid mutator transaction binding the contract method 0x44014b6b.
//
// Solidity: function setSubjectHandler(address subjectGateway) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) SetSubjectHandler(subjectGateway common.Address) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.SetSubjectHandler(&_ScannerPoolRegistry.TransactOpts, subjectGateway)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.TransferFrom(&_ScannerPoolRegistry.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.TransferFrom(&_ScannerPoolRegistry.TransactOpts, from, to, tokenId)
}

// UpdateEnabledScanners is a paid mutator transaction binding the contract method 0xca009283.
//
// Solidity: function updateEnabledScanners(uint256 scannerPoolId, uint256 count) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) UpdateEnabledScanners(opts *bind.TransactOpts, scannerPoolId *big.Int, count *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "updateEnabledScanners", scannerPoolId, count)
}

// UpdateEnabledScanners is a paid mutator transaction binding the contract method 0xca009283.
//
// Solidity: function updateEnabledScanners(uint256 scannerPoolId, uint256 count) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) UpdateEnabledScanners(scannerPoolId *big.Int, count *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.UpdateEnabledScanners(&_ScannerPoolRegistry.TransactOpts, scannerPoolId, count)
}

// UpdateEnabledScanners is a paid mutator transaction binding the contract method 0xca009283.
//
// Solidity: function updateEnabledScanners(uint256 scannerPoolId, uint256 count) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) UpdateEnabledScanners(scannerPoolId *big.Int, count *big.Int) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.UpdateEnabledScanners(&_ScannerPoolRegistry.TransactOpts, scannerPoolId, count)
}

// UpdateScannerMetadata is a paid mutator transaction binding the contract method 0xd5c35af4.
//
// Solidity: function updateScannerMetadata(address scanner, string metadata) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) UpdateScannerMetadata(opts *bind.TransactOpts, scanner common.Address, metadata string) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "updateScannerMetadata", scanner, metadata)
}

// UpdateScannerMetadata is a paid mutator transaction binding the contract method 0xd5c35af4.
//
// Solidity: function updateScannerMetadata(address scanner, string metadata) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) UpdateScannerMetadata(scanner common.Address, metadata string) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.UpdateScannerMetadata(&_ScannerPoolRegistry.TransactOpts, scanner, metadata)
}

// UpdateScannerMetadata is a paid mutator transaction binding the contract method 0xd5c35af4.
//
// Solidity: function updateScannerMetadata(address scanner, string metadata) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) UpdateScannerMetadata(scanner common.Address, metadata string) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.UpdateScannerMetadata(&_ScannerPoolRegistry.TransactOpts, scanner, metadata)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.UpgradeTo(&_ScannerPoolRegistry.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.UpgradeTo(&_ScannerPoolRegistry.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ScannerPoolRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.UpgradeToAndCall(&_ScannerPoolRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ScannerPoolRegistry *ScannerPoolRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ScannerPoolRegistry.Contract.UpgradeToAndCall(&_ScannerPoolRegistry.TransactOpts, newImplementation, data)
}

// ScannerPoolRegistryAccessManagerUpdatedIterator is returned from FilterAccessManagerUpdated and is used to iterate over the raw logs and unpacked data for AccessManagerUpdated events raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryAccessManagerUpdatedIterator struct {
	Event *ScannerPoolRegistryAccessManagerUpdated // Event containing the contract specifics and raw log

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
func (it *ScannerPoolRegistryAccessManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerPoolRegistryAccessManagerUpdated)
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
		it.Event = new(ScannerPoolRegistryAccessManagerUpdated)
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
func (it *ScannerPoolRegistryAccessManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerPoolRegistryAccessManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerPoolRegistryAccessManagerUpdated represents a AccessManagerUpdated event raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryAccessManagerUpdated struct {
	NewAddressManager common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAccessManagerUpdated is a free log retrieval operation binding the contract event 0xa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c.
//
// Solidity: event AccessManagerUpdated(address indexed newAddressManager)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) FilterAccessManagerUpdated(opts *bind.FilterOpts, newAddressManager []common.Address) (*ScannerPoolRegistryAccessManagerUpdatedIterator, error) {

	var newAddressManagerRule []interface{}
	for _, newAddressManagerItem := range newAddressManager {
		newAddressManagerRule = append(newAddressManagerRule, newAddressManagerItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.FilterLogs(opts, "AccessManagerUpdated", newAddressManagerRule)
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistryAccessManagerUpdatedIterator{contract: _ScannerPoolRegistry.contract, event: "AccessManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchAccessManagerUpdated is a free log subscription operation binding the contract event 0xa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c.
//
// Solidity: event AccessManagerUpdated(address indexed newAddressManager)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) WatchAccessManagerUpdated(opts *bind.WatchOpts, sink chan<- *ScannerPoolRegistryAccessManagerUpdated, newAddressManager []common.Address) (event.Subscription, error) {

	var newAddressManagerRule []interface{}
	for _, newAddressManagerItem := range newAddressManager {
		newAddressManagerRule = append(newAddressManagerRule, newAddressManagerItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.WatchLogs(opts, "AccessManagerUpdated", newAddressManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerPoolRegistryAccessManagerUpdated)
				if err := _ScannerPoolRegistry.contract.UnpackLog(event, "AccessManagerUpdated", log); err != nil {
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
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) ParseAccessManagerUpdated(log types.Log) (*ScannerPoolRegistryAccessManagerUpdated, error) {
	event := new(ScannerPoolRegistryAccessManagerUpdated)
	if err := _ScannerPoolRegistry.contract.UnpackLog(event, "AccessManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerPoolRegistryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryAdminChangedIterator struct {
	Event *ScannerPoolRegistryAdminChanged // Event containing the contract specifics and raw log

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
func (it *ScannerPoolRegistryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerPoolRegistryAdminChanged)
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
		it.Event = new(ScannerPoolRegistryAdminChanged)
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
func (it *ScannerPoolRegistryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerPoolRegistryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerPoolRegistryAdminChanged represents a AdminChanged event raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ScannerPoolRegistryAdminChangedIterator, error) {

	logs, sub, err := _ScannerPoolRegistry.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistryAdminChangedIterator{contract: _ScannerPoolRegistry.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ScannerPoolRegistryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ScannerPoolRegistry.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerPoolRegistryAdminChanged)
				if err := _ScannerPoolRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) ParseAdminChanged(log types.Log) (*ScannerPoolRegistryAdminChanged, error) {
	event := new(ScannerPoolRegistryAdminChanged)
	if err := _ScannerPoolRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerPoolRegistryApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryApprovalIterator struct {
	Event *ScannerPoolRegistryApproval // Event containing the contract specifics and raw log

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
func (it *ScannerPoolRegistryApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerPoolRegistryApproval)
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
		it.Event = new(ScannerPoolRegistryApproval)
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
func (it *ScannerPoolRegistryApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerPoolRegistryApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerPoolRegistryApproval represents a Approval event raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ScannerPoolRegistryApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistryApprovalIterator{contract: _ScannerPoolRegistry.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ScannerPoolRegistryApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerPoolRegistryApproval)
				if err := _ScannerPoolRegistry.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) ParseApproval(log types.Log) (*ScannerPoolRegistryApproval, error) {
	event := new(ScannerPoolRegistryApproval)
	if err := _ScannerPoolRegistry.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerPoolRegistryApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryApprovalForAllIterator struct {
	Event *ScannerPoolRegistryApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ScannerPoolRegistryApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerPoolRegistryApprovalForAll)
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
		it.Event = new(ScannerPoolRegistryApprovalForAll)
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
func (it *ScannerPoolRegistryApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerPoolRegistryApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerPoolRegistryApprovalForAll represents a ApprovalForAll event raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ScannerPoolRegistryApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistryApprovalForAllIterator{contract: _ScannerPoolRegistry.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ScannerPoolRegistryApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerPoolRegistryApprovalForAll)
				if err := _ScannerPoolRegistry.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) ParseApprovalForAll(log types.Log) (*ScannerPoolRegistryApprovalForAll, error) {
	event := new(ScannerPoolRegistryApprovalForAll)
	if err := _ScannerPoolRegistry.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerPoolRegistryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryBeaconUpgradedIterator struct {
	Event *ScannerPoolRegistryBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ScannerPoolRegistryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerPoolRegistryBeaconUpgraded)
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
		it.Event = new(ScannerPoolRegistryBeaconUpgraded)
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
func (it *ScannerPoolRegistryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerPoolRegistryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerPoolRegistryBeaconUpgraded represents a BeaconUpgraded event raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ScannerPoolRegistryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistryBeaconUpgradedIterator{contract: _ScannerPoolRegistry.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ScannerPoolRegistryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerPoolRegistryBeaconUpgraded)
				if err := _ScannerPoolRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) ParseBeaconUpgraded(log types.Log) (*ScannerPoolRegistryBeaconUpgraded, error) {
	event := new(ScannerPoolRegistryBeaconUpgraded)
	if err := _ScannerPoolRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerPoolRegistryEnabledScannersChangedIterator is returned from FilterEnabledScannersChanged and is used to iterate over the raw logs and unpacked data for EnabledScannersChanged events raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryEnabledScannersChangedIterator struct {
	Event *ScannerPoolRegistryEnabledScannersChanged // Event containing the contract specifics and raw log

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
func (it *ScannerPoolRegistryEnabledScannersChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerPoolRegistryEnabledScannersChanged)
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
		it.Event = new(ScannerPoolRegistryEnabledScannersChanged)
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
func (it *ScannerPoolRegistryEnabledScannersChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerPoolRegistryEnabledScannersChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerPoolRegistryEnabledScannersChanged represents a EnabledScannersChanged event raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryEnabledScannersChanged struct {
	ScannerPoolId   *big.Int
	EnabledScanners *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEnabledScannersChanged is a free log retrieval operation binding the contract event 0xe98aca89d957a1075e807d80a8f0f7d280340888edefe3da76c278ca302922c8.
//
// Solidity: event EnabledScannersChanged(uint256 indexed scannerPoolId, uint256 enabledScanners)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) FilterEnabledScannersChanged(opts *bind.FilterOpts, scannerPoolId []*big.Int) (*ScannerPoolRegistryEnabledScannersChangedIterator, error) {

	var scannerPoolIdRule []interface{}
	for _, scannerPoolIdItem := range scannerPoolId {
		scannerPoolIdRule = append(scannerPoolIdRule, scannerPoolIdItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.FilterLogs(opts, "EnabledScannersChanged", scannerPoolIdRule)
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistryEnabledScannersChangedIterator{contract: _ScannerPoolRegistry.contract, event: "EnabledScannersChanged", logs: logs, sub: sub}, nil
}

// WatchEnabledScannersChanged is a free log subscription operation binding the contract event 0xe98aca89d957a1075e807d80a8f0f7d280340888edefe3da76c278ca302922c8.
//
// Solidity: event EnabledScannersChanged(uint256 indexed scannerPoolId, uint256 enabledScanners)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) WatchEnabledScannersChanged(opts *bind.WatchOpts, sink chan<- *ScannerPoolRegistryEnabledScannersChanged, scannerPoolId []*big.Int) (event.Subscription, error) {

	var scannerPoolIdRule []interface{}
	for _, scannerPoolIdItem := range scannerPoolId {
		scannerPoolIdRule = append(scannerPoolIdRule, scannerPoolIdItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.WatchLogs(opts, "EnabledScannersChanged", scannerPoolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerPoolRegistryEnabledScannersChanged)
				if err := _ScannerPoolRegistry.contract.UnpackLog(event, "EnabledScannersChanged", log); err != nil {
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

// ParseEnabledScannersChanged is a log parse operation binding the contract event 0xe98aca89d957a1075e807d80a8f0f7d280340888edefe3da76c278ca302922c8.
//
// Solidity: event EnabledScannersChanged(uint256 indexed scannerPoolId, uint256 enabledScanners)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) ParseEnabledScannersChanged(log types.Log) (*ScannerPoolRegistryEnabledScannersChanged, error) {
	event := new(ScannerPoolRegistryEnabledScannersChanged)
	if err := _ScannerPoolRegistry.contract.UnpackLog(event, "EnabledScannersChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerPoolRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryInitializedIterator struct {
	Event *ScannerPoolRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ScannerPoolRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerPoolRegistryInitialized)
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
		it.Event = new(ScannerPoolRegistryInitialized)
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
func (it *ScannerPoolRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerPoolRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerPoolRegistryInitialized represents a Initialized event raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ScannerPoolRegistryInitializedIterator, error) {

	logs, sub, err := _ScannerPoolRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistryInitializedIterator{contract: _ScannerPoolRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ScannerPoolRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ScannerPoolRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerPoolRegistryInitialized)
				if err := _ScannerPoolRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) ParseInitialized(log types.Log) (*ScannerPoolRegistryInitialized, error) {
	event := new(ScannerPoolRegistryInitialized)
	if err := _ScannerPoolRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerPoolRegistryManagedStakeThresholdChangedIterator is returned from FilterManagedStakeThresholdChanged and is used to iterate over the raw logs and unpacked data for ManagedStakeThresholdChanged events raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryManagedStakeThresholdChangedIterator struct {
	Event *ScannerPoolRegistryManagedStakeThresholdChanged // Event containing the contract specifics and raw log

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
func (it *ScannerPoolRegistryManagedStakeThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerPoolRegistryManagedStakeThresholdChanged)
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
		it.Event = new(ScannerPoolRegistryManagedStakeThresholdChanged)
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
func (it *ScannerPoolRegistryManagedStakeThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerPoolRegistryManagedStakeThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerPoolRegistryManagedStakeThresholdChanged represents a ManagedStakeThresholdChanged event raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryManagedStakeThresholdChanged struct {
	ChainId   *big.Int
	Min       *big.Int
	Max       *big.Int
	Activated bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterManagedStakeThresholdChanged is a free log retrieval operation binding the contract event 0xc89d7e1caf32a74dc9fc2c0197a92afe1936d38c73dd808b0f92fc50677b91ff.
//
// Solidity: event ManagedStakeThresholdChanged(uint256 indexed chainId, uint256 min, uint256 max, bool activated)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) FilterManagedStakeThresholdChanged(opts *bind.FilterOpts, chainId []*big.Int) (*ScannerPoolRegistryManagedStakeThresholdChangedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.FilterLogs(opts, "ManagedStakeThresholdChanged", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistryManagedStakeThresholdChangedIterator{contract: _ScannerPoolRegistry.contract, event: "ManagedStakeThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchManagedStakeThresholdChanged is a free log subscription operation binding the contract event 0xc89d7e1caf32a74dc9fc2c0197a92afe1936d38c73dd808b0f92fc50677b91ff.
//
// Solidity: event ManagedStakeThresholdChanged(uint256 indexed chainId, uint256 min, uint256 max, bool activated)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) WatchManagedStakeThresholdChanged(opts *bind.WatchOpts, sink chan<- *ScannerPoolRegistryManagedStakeThresholdChanged, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.WatchLogs(opts, "ManagedStakeThresholdChanged", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerPoolRegistryManagedStakeThresholdChanged)
				if err := _ScannerPoolRegistry.contract.UnpackLog(event, "ManagedStakeThresholdChanged", log); err != nil {
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

// ParseManagedStakeThresholdChanged is a log parse operation binding the contract event 0xc89d7e1caf32a74dc9fc2c0197a92afe1936d38c73dd808b0f92fc50677b91ff.
//
// Solidity: event ManagedStakeThresholdChanged(uint256 indexed chainId, uint256 min, uint256 max, bool activated)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) ParseManagedStakeThresholdChanged(log types.Log) (*ScannerPoolRegistryManagedStakeThresholdChanged, error) {
	event := new(ScannerPoolRegistryManagedStakeThresholdChanged)
	if err := _ScannerPoolRegistry.contract.UnpackLog(event, "ManagedStakeThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerPoolRegistryManagerEnabledIterator is returned from FilterManagerEnabled and is used to iterate over the raw logs and unpacked data for ManagerEnabled events raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryManagerEnabledIterator struct {
	Event *ScannerPoolRegistryManagerEnabled // Event containing the contract specifics and raw log

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
func (it *ScannerPoolRegistryManagerEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerPoolRegistryManagerEnabled)
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
		it.Event = new(ScannerPoolRegistryManagerEnabled)
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
func (it *ScannerPoolRegistryManagerEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerPoolRegistryManagerEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerPoolRegistryManagerEnabled represents a ManagerEnabled event raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryManagerEnabled struct {
	ScannerPoolId *big.Int
	Manager       common.Address
	Enabled       bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterManagerEnabled is a free log retrieval operation binding the contract event 0x538b6537a6fe8f0deae9f3b86ad1924d5e5b3d5a683055276b2824f918be043e.
//
// Solidity: event ManagerEnabled(uint256 indexed scannerPoolId, address indexed manager, bool enabled)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) FilterManagerEnabled(opts *bind.FilterOpts, scannerPoolId []*big.Int, manager []common.Address) (*ScannerPoolRegistryManagerEnabledIterator, error) {

	var scannerPoolIdRule []interface{}
	for _, scannerPoolIdItem := range scannerPoolId {
		scannerPoolIdRule = append(scannerPoolIdRule, scannerPoolIdItem)
	}
	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.FilterLogs(opts, "ManagerEnabled", scannerPoolIdRule, managerRule)
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistryManagerEnabledIterator{contract: _ScannerPoolRegistry.contract, event: "ManagerEnabled", logs: logs, sub: sub}, nil
}

// WatchManagerEnabled is a free log subscription operation binding the contract event 0x538b6537a6fe8f0deae9f3b86ad1924d5e5b3d5a683055276b2824f918be043e.
//
// Solidity: event ManagerEnabled(uint256 indexed scannerPoolId, address indexed manager, bool enabled)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) WatchManagerEnabled(opts *bind.WatchOpts, sink chan<- *ScannerPoolRegistryManagerEnabled, scannerPoolId []*big.Int, manager []common.Address) (event.Subscription, error) {

	var scannerPoolIdRule []interface{}
	for _, scannerPoolIdItem := range scannerPoolId {
		scannerPoolIdRule = append(scannerPoolIdRule, scannerPoolIdItem)
	}
	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.WatchLogs(opts, "ManagerEnabled", scannerPoolIdRule, managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerPoolRegistryManagerEnabled)
				if err := _ScannerPoolRegistry.contract.UnpackLog(event, "ManagerEnabled", log); err != nil {
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

// ParseManagerEnabled is a log parse operation binding the contract event 0x538b6537a6fe8f0deae9f3b86ad1924d5e5b3d5a683055276b2824f918be043e.
//
// Solidity: event ManagerEnabled(uint256 indexed scannerPoolId, address indexed manager, bool enabled)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) ParseManagerEnabled(log types.Log) (*ScannerPoolRegistryManagerEnabled, error) {
	event := new(ScannerPoolRegistryManagerEnabled)
	if err := _ScannerPoolRegistry.contract.UnpackLog(event, "ManagerEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerPoolRegistryRegistrationDelaySetIterator is returned from FilterRegistrationDelaySet and is used to iterate over the raw logs and unpacked data for RegistrationDelaySet events raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryRegistrationDelaySetIterator struct {
	Event *ScannerPoolRegistryRegistrationDelaySet // Event containing the contract specifics and raw log

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
func (it *ScannerPoolRegistryRegistrationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerPoolRegistryRegistrationDelaySet)
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
		it.Event = new(ScannerPoolRegistryRegistrationDelaySet)
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
func (it *ScannerPoolRegistryRegistrationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerPoolRegistryRegistrationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerPoolRegistryRegistrationDelaySet represents a RegistrationDelaySet event raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryRegistrationDelaySet struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRegistrationDelaySet is a free log retrieval operation binding the contract event 0xbb57183c638abde4f4007bb9e16e4dc9fe60ec4e316cc19e9329d9782cabaeac.
//
// Solidity: event RegistrationDelaySet(uint256 delay)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) FilterRegistrationDelaySet(opts *bind.FilterOpts) (*ScannerPoolRegistryRegistrationDelaySetIterator, error) {

	logs, sub, err := _ScannerPoolRegistry.contract.FilterLogs(opts, "RegistrationDelaySet")
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistryRegistrationDelaySetIterator{contract: _ScannerPoolRegistry.contract, event: "RegistrationDelaySet", logs: logs, sub: sub}, nil
}

// WatchRegistrationDelaySet is a free log subscription operation binding the contract event 0xbb57183c638abde4f4007bb9e16e4dc9fe60ec4e316cc19e9329d9782cabaeac.
//
// Solidity: event RegistrationDelaySet(uint256 delay)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) WatchRegistrationDelaySet(opts *bind.WatchOpts, sink chan<- *ScannerPoolRegistryRegistrationDelaySet) (event.Subscription, error) {

	logs, sub, err := _ScannerPoolRegistry.contract.WatchLogs(opts, "RegistrationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerPoolRegistryRegistrationDelaySet)
				if err := _ScannerPoolRegistry.contract.UnpackLog(event, "RegistrationDelaySet", log); err != nil {
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

// ParseRegistrationDelaySet is a log parse operation binding the contract event 0xbb57183c638abde4f4007bb9e16e4dc9fe60ec4e316cc19e9329d9782cabaeac.
//
// Solidity: event RegistrationDelaySet(uint256 delay)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) ParseRegistrationDelaySet(log types.Log) (*ScannerPoolRegistryRegistrationDelaySet, error) {
	event := new(ScannerPoolRegistryRegistrationDelaySet)
	if err := _ScannerPoolRegistry.contract.UnpackLog(event, "RegistrationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerPoolRegistryRouterUpdatedIterator is returned from FilterRouterUpdated and is used to iterate over the raw logs and unpacked data for RouterUpdated events raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryRouterUpdatedIterator struct {
	Event *ScannerPoolRegistryRouterUpdated // Event containing the contract specifics and raw log

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
func (it *ScannerPoolRegistryRouterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerPoolRegistryRouterUpdated)
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
		it.Event = new(ScannerPoolRegistryRouterUpdated)
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
func (it *ScannerPoolRegistryRouterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerPoolRegistryRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerPoolRegistryRouterUpdated represents a RouterUpdated event raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryRouterUpdated struct {
	Router common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRouterUpdated is a free log retrieval operation binding the contract event 0x7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc80.
//
// Solidity: event RouterUpdated(address indexed router)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) FilterRouterUpdated(opts *bind.FilterOpts, router []common.Address) (*ScannerPoolRegistryRouterUpdatedIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.FilterLogs(opts, "RouterUpdated", routerRule)
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistryRouterUpdatedIterator{contract: _ScannerPoolRegistry.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

// WatchRouterUpdated is a free log subscription operation binding the contract event 0x7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc80.
//
// Solidity: event RouterUpdated(address indexed router)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *ScannerPoolRegistryRouterUpdated, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.WatchLogs(opts, "RouterUpdated", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerPoolRegistryRouterUpdated)
				if err := _ScannerPoolRegistry.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
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
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) ParseRouterUpdated(log types.Log) (*ScannerPoolRegistryRouterUpdated, error) {
	event := new(ScannerPoolRegistryRouterUpdated)
	if err := _ScannerPoolRegistry.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerPoolRegistryScannerEnabledIterator is returned from FilterScannerEnabled and is used to iterate over the raw logs and unpacked data for ScannerEnabled events raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryScannerEnabledIterator struct {
	Event *ScannerPoolRegistryScannerEnabled // Event containing the contract specifics and raw log

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
func (it *ScannerPoolRegistryScannerEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerPoolRegistryScannerEnabled)
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
		it.Event = new(ScannerPoolRegistryScannerEnabled)
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
func (it *ScannerPoolRegistryScannerEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerPoolRegistryScannerEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerPoolRegistryScannerEnabled represents a ScannerEnabled event raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryScannerEnabled struct {
	ScannerId   *big.Int
	Enabled     bool
	Sender      common.Address
	DisableFlag bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterScannerEnabled is a free log retrieval operation binding the contract event 0x800131ba226bbb1e9c537607e0f4d73e692ae496e7d1fb9b2aea059369d37fa2.
//
// Solidity: event ScannerEnabled(uint256 indexed scannerId, bool indexed enabled, address sender, bool disableFlag)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) FilterScannerEnabled(opts *bind.FilterOpts, scannerId []*big.Int, enabled []bool) (*ScannerPoolRegistryScannerEnabledIterator, error) {

	var scannerIdRule []interface{}
	for _, scannerIdItem := range scannerId {
		scannerIdRule = append(scannerIdRule, scannerIdItem)
	}
	var enabledRule []interface{}
	for _, enabledItem := range enabled {
		enabledRule = append(enabledRule, enabledItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.FilterLogs(opts, "ScannerEnabled", scannerIdRule, enabledRule)
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistryScannerEnabledIterator{contract: _ScannerPoolRegistry.contract, event: "ScannerEnabled", logs: logs, sub: sub}, nil
}

// WatchScannerEnabled is a free log subscription operation binding the contract event 0x800131ba226bbb1e9c537607e0f4d73e692ae496e7d1fb9b2aea059369d37fa2.
//
// Solidity: event ScannerEnabled(uint256 indexed scannerId, bool indexed enabled, address sender, bool disableFlag)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) WatchScannerEnabled(opts *bind.WatchOpts, sink chan<- *ScannerPoolRegistryScannerEnabled, scannerId []*big.Int, enabled []bool) (event.Subscription, error) {

	var scannerIdRule []interface{}
	for _, scannerIdItem := range scannerId {
		scannerIdRule = append(scannerIdRule, scannerIdItem)
	}
	var enabledRule []interface{}
	for _, enabledItem := range enabled {
		enabledRule = append(enabledRule, enabledItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.WatchLogs(opts, "ScannerEnabled", scannerIdRule, enabledRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerPoolRegistryScannerEnabled)
				if err := _ScannerPoolRegistry.contract.UnpackLog(event, "ScannerEnabled", log); err != nil {
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

// ParseScannerEnabled is a log parse operation binding the contract event 0x800131ba226bbb1e9c537607e0f4d73e692ae496e7d1fb9b2aea059369d37fa2.
//
// Solidity: event ScannerEnabled(uint256 indexed scannerId, bool indexed enabled, address sender, bool disableFlag)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) ParseScannerEnabled(log types.Log) (*ScannerPoolRegistryScannerEnabled, error) {
	event := new(ScannerPoolRegistryScannerEnabled)
	if err := _ScannerPoolRegistry.contract.UnpackLog(event, "ScannerEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerPoolRegistryScannerPoolRegisteredIterator is returned from FilterScannerPoolRegistered and is used to iterate over the raw logs and unpacked data for ScannerPoolRegistered events raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryScannerPoolRegisteredIterator struct {
	Event *ScannerPoolRegistryScannerPoolRegistered // Event containing the contract specifics and raw log

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
func (it *ScannerPoolRegistryScannerPoolRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerPoolRegistryScannerPoolRegistered)
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
		it.Event = new(ScannerPoolRegistryScannerPoolRegistered)
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
func (it *ScannerPoolRegistryScannerPoolRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerPoolRegistryScannerPoolRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerPoolRegistryScannerPoolRegistered represents a ScannerPoolRegistered event raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryScannerPoolRegistered struct {
	ScannerPoolId *big.Int
	ChainId       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterScannerPoolRegistered is a free log retrieval operation binding the contract event 0x7e191d1226c4da936653bac3b9ae5c76dd880cf5221a80bcaa224448ccc67048.
//
// Solidity: event ScannerPoolRegistered(uint256 indexed scannerPoolId, uint256 indexed chainId)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) FilterScannerPoolRegistered(opts *bind.FilterOpts, scannerPoolId []*big.Int, chainId []*big.Int) (*ScannerPoolRegistryScannerPoolRegisteredIterator, error) {

	var scannerPoolIdRule []interface{}
	for _, scannerPoolIdItem := range scannerPoolId {
		scannerPoolIdRule = append(scannerPoolIdRule, scannerPoolIdItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.FilterLogs(opts, "ScannerPoolRegistered", scannerPoolIdRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistryScannerPoolRegisteredIterator{contract: _ScannerPoolRegistry.contract, event: "ScannerPoolRegistered", logs: logs, sub: sub}, nil
}

// WatchScannerPoolRegistered is a free log subscription operation binding the contract event 0x7e191d1226c4da936653bac3b9ae5c76dd880cf5221a80bcaa224448ccc67048.
//
// Solidity: event ScannerPoolRegistered(uint256 indexed scannerPoolId, uint256 indexed chainId)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) WatchScannerPoolRegistered(opts *bind.WatchOpts, sink chan<- *ScannerPoolRegistryScannerPoolRegistered, scannerPoolId []*big.Int, chainId []*big.Int) (event.Subscription, error) {

	var scannerPoolIdRule []interface{}
	for _, scannerPoolIdItem := range scannerPoolId {
		scannerPoolIdRule = append(scannerPoolIdRule, scannerPoolIdItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.WatchLogs(opts, "ScannerPoolRegistered", scannerPoolIdRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerPoolRegistryScannerPoolRegistered)
				if err := _ScannerPoolRegistry.contract.UnpackLog(event, "ScannerPoolRegistered", log); err != nil {
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

// ParseScannerPoolRegistered is a log parse operation binding the contract event 0x7e191d1226c4da936653bac3b9ae5c76dd880cf5221a80bcaa224448ccc67048.
//
// Solidity: event ScannerPoolRegistered(uint256 indexed scannerPoolId, uint256 indexed chainId)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) ParseScannerPoolRegistered(log types.Log) (*ScannerPoolRegistryScannerPoolRegistered, error) {
	event := new(ScannerPoolRegistryScannerPoolRegistered)
	if err := _ScannerPoolRegistry.contract.UnpackLog(event, "ScannerPoolRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerPoolRegistryScannerUpdatedIterator is returned from FilterScannerUpdated and is used to iterate over the raw logs and unpacked data for ScannerUpdated events raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryScannerUpdatedIterator struct {
	Event *ScannerPoolRegistryScannerUpdated // Event containing the contract specifics and raw log

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
func (it *ScannerPoolRegistryScannerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerPoolRegistryScannerUpdated)
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
		it.Event = new(ScannerPoolRegistryScannerUpdated)
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
func (it *ScannerPoolRegistryScannerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerPoolRegistryScannerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerPoolRegistryScannerUpdated represents a ScannerUpdated event raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryScannerUpdated struct {
	ScannerId   *big.Int
	ChainId     *big.Int
	Metadata    string
	ScannerPool *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterScannerUpdated is a free log retrieval operation binding the contract event 0xfc43d422261aa008210f185c9a82dfeb4faa37fcdac42e357e80bf0a3446e3e8.
//
// Solidity: event ScannerUpdated(uint256 indexed scannerId, uint256 indexed chainId, string metadata, uint256 scannerPool)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) FilterScannerUpdated(opts *bind.FilterOpts, scannerId []*big.Int, chainId []*big.Int) (*ScannerPoolRegistryScannerUpdatedIterator, error) {

	var scannerIdRule []interface{}
	for _, scannerIdItem := range scannerId {
		scannerIdRule = append(scannerIdRule, scannerIdItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.FilterLogs(opts, "ScannerUpdated", scannerIdRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistryScannerUpdatedIterator{contract: _ScannerPoolRegistry.contract, event: "ScannerUpdated", logs: logs, sub: sub}, nil
}

// WatchScannerUpdated is a free log subscription operation binding the contract event 0xfc43d422261aa008210f185c9a82dfeb4faa37fcdac42e357e80bf0a3446e3e8.
//
// Solidity: event ScannerUpdated(uint256 indexed scannerId, uint256 indexed chainId, string metadata, uint256 scannerPool)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) WatchScannerUpdated(opts *bind.WatchOpts, sink chan<- *ScannerPoolRegistryScannerUpdated, scannerId []*big.Int, chainId []*big.Int) (event.Subscription, error) {

	var scannerIdRule []interface{}
	for _, scannerIdItem := range scannerId {
		scannerIdRule = append(scannerIdRule, scannerIdItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.WatchLogs(opts, "ScannerUpdated", scannerIdRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerPoolRegistryScannerUpdated)
				if err := _ScannerPoolRegistry.contract.UnpackLog(event, "ScannerUpdated", log); err != nil {
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

// ParseScannerUpdated is a log parse operation binding the contract event 0xfc43d422261aa008210f185c9a82dfeb4faa37fcdac42e357e80bf0a3446e3e8.
//
// Solidity: event ScannerUpdated(uint256 indexed scannerId, uint256 indexed chainId, string metadata, uint256 scannerPool)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) ParseScannerUpdated(log types.Log) (*ScannerPoolRegistryScannerUpdated, error) {
	event := new(ScannerPoolRegistryScannerUpdated)
	if err := _ScannerPoolRegistry.contract.UnpackLog(event, "ScannerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerPoolRegistrySubjectHandlerUpdatedIterator is returned from FilterSubjectHandlerUpdated and is used to iterate over the raw logs and unpacked data for SubjectHandlerUpdated events raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistrySubjectHandlerUpdatedIterator struct {
	Event *ScannerPoolRegistrySubjectHandlerUpdated // Event containing the contract specifics and raw log

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
func (it *ScannerPoolRegistrySubjectHandlerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerPoolRegistrySubjectHandlerUpdated)
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
		it.Event = new(ScannerPoolRegistrySubjectHandlerUpdated)
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
func (it *ScannerPoolRegistrySubjectHandlerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerPoolRegistrySubjectHandlerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerPoolRegistrySubjectHandlerUpdated represents a SubjectHandlerUpdated event raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistrySubjectHandlerUpdated struct {
	NewHandler common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubjectHandlerUpdated is a free log retrieval operation binding the contract event 0x16d72e484786227d7b6dd05e079be9ff9904a81a0a9ec36fc346b20f8c47aff0.
//
// Solidity: event SubjectHandlerUpdated(address indexed newHandler)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) FilterSubjectHandlerUpdated(opts *bind.FilterOpts, newHandler []common.Address) (*ScannerPoolRegistrySubjectHandlerUpdatedIterator, error) {

	var newHandlerRule []interface{}
	for _, newHandlerItem := range newHandler {
		newHandlerRule = append(newHandlerRule, newHandlerItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.FilterLogs(opts, "SubjectHandlerUpdated", newHandlerRule)
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistrySubjectHandlerUpdatedIterator{contract: _ScannerPoolRegistry.contract, event: "SubjectHandlerUpdated", logs: logs, sub: sub}, nil
}

// WatchSubjectHandlerUpdated is a free log subscription operation binding the contract event 0x16d72e484786227d7b6dd05e079be9ff9904a81a0a9ec36fc346b20f8c47aff0.
//
// Solidity: event SubjectHandlerUpdated(address indexed newHandler)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) WatchSubjectHandlerUpdated(opts *bind.WatchOpts, sink chan<- *ScannerPoolRegistrySubjectHandlerUpdated, newHandler []common.Address) (event.Subscription, error) {

	var newHandlerRule []interface{}
	for _, newHandlerItem := range newHandler {
		newHandlerRule = append(newHandlerRule, newHandlerItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.WatchLogs(opts, "SubjectHandlerUpdated", newHandlerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerPoolRegistrySubjectHandlerUpdated)
				if err := _ScannerPoolRegistry.contract.UnpackLog(event, "SubjectHandlerUpdated", log); err != nil {
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

// ParseSubjectHandlerUpdated is a log parse operation binding the contract event 0x16d72e484786227d7b6dd05e079be9ff9904a81a0a9ec36fc346b20f8c47aff0.
//
// Solidity: event SubjectHandlerUpdated(address indexed newHandler)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) ParseSubjectHandlerUpdated(log types.Log) (*ScannerPoolRegistrySubjectHandlerUpdated, error) {
	event := new(ScannerPoolRegistrySubjectHandlerUpdated)
	if err := _ScannerPoolRegistry.contract.UnpackLog(event, "SubjectHandlerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerPoolRegistryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryTransferIterator struct {
	Event *ScannerPoolRegistryTransfer // Event containing the contract specifics and raw log

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
func (it *ScannerPoolRegistryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerPoolRegistryTransfer)
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
		it.Event = new(ScannerPoolRegistryTransfer)
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
func (it *ScannerPoolRegistryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerPoolRegistryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerPoolRegistryTransfer represents a Transfer event raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ScannerPoolRegistryTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistryTransferIterator{contract: _ScannerPoolRegistry.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ScannerPoolRegistryTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerPoolRegistryTransfer)
				if err := _ScannerPoolRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) ParseTransfer(log types.Log) (*ScannerPoolRegistryTransfer, error) {
	event := new(ScannerPoolRegistryTransfer)
	if err := _ScannerPoolRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerPoolRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryUpgradedIterator struct {
	Event *ScannerPoolRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *ScannerPoolRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerPoolRegistryUpgraded)
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
		it.Event = new(ScannerPoolRegistryUpgraded)
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
func (it *ScannerPoolRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerPoolRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerPoolRegistryUpgraded represents a Upgraded event raised by the ScannerPoolRegistry contract.
type ScannerPoolRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ScannerPoolRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ScannerPoolRegistryUpgradedIterator{contract: _ScannerPoolRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ScannerPoolRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ScannerPoolRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerPoolRegistryUpgraded)
				if err := _ScannerPoolRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ScannerPoolRegistry *ScannerPoolRegistryFilterer) ParseUpgraded(log types.Log) (*ScannerPoolRegistryUpgraded, error) {
	event := new(ScannerPoolRegistryUpgraded)
	if err := _ScannerPoolRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
