// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_scanner_pool_registry

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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakeAllocator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"AccessManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"enabledScanners\",\"type\":\"uint256\"}],\"name\":\"EnabledScannersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"activated\",\"type\":\"bool\"}],\"name\":\"ManagedStakeThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"ManagerEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"RegistrationDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"scannerId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"disableFlag\",\"type\":\"bool\"}],\"name\":\"ScannerEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"ScannerPoolRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"scannerId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"scannerPool\",\"type\":\"uint256\"}],\"name\":\"ScannerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newHandler\",\"type\":\"address\"}],\"name\":\"SubjectHandlerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"}],\"name\":\"disableScanner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"}],\"name\":\"enableScanner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"managedId\",\"type\":\"uint256\"}],\"name\":\"getManagedStakeThreshold\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"activated\",\"type\":\"bool\"}],\"internalType\":\"structIStakeSubject.StakeThreshold\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getManagerAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"}],\"name\":\"getManagerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"}],\"name\":\"getScanner\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structScannerPoolRegistryCore.ScannerNode\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"}],\"name\":\"getScannerState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"operational\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSubjectHandler\",\"outputs\":[{\"internalType\":\"contractIStakeSubjectGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"getTotalManagedSubjects\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__manager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"__name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"__symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"__stakeSubjectGateway\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"__registrationDelay\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"}],\"name\":\"isScannerDisabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"}],\"name\":\"isScannerOperational\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"}],\"name\":\"isScannerRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"}],\"name\":\"isScannerRegisteredTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"}],\"name\":\"monitoredChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structScannerPoolRegistryCore.ScannerNodeRegistration\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"name\":\"registerMigratedScannerNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scannerPoolAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"registerMigratedScannerPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structScannerPoolRegistryCore.ScannerNodeRegistration\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"registerScannerNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"registerScannerPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"registeredScannerAddressAtIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"registeredScannerAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structScannerPoolRegistryCore.ScannerNode\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registrationDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"}],\"name\":\"scannerAddressToId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerId\",\"type\":\"uint256\"}],\"name\":\"scannerIdToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"setAccessManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"activated\",\"type\":\"bool\"}],\"internalType\":\"structIStakeSubject.StakeThreshold\",\"name\":\"newStakeThreshold\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"setManagedStakeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ensRegistry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ensName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setRegistrationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subjectGateway\",\"type\":\"address\"}],\"name\":\"setSubjectHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scannerPoolId\",\"type\":\"uint256\"}],\"name\":\"totalScannersRegistered\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scanner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"updateScannerMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e06040523060a0523480156200001557600080fd5b50604051620057c0380380620057c08339810160408190526200003891620001f2565b6001600160a01b0380831660805281908116620000905760405163eac0d38960e01b815260206004820152601060248201526f2fafb9ba30b5b2a0b63637b1b0ba37b960811b60448201526064015b60405180910390fd5b6001600160a01b031660c052600054610100900460ff1615808015620000bd5750600054600160ff909116105b80620000ed5750620000da30620001c660201b620022661760201c565b158015620000ed575060005460ff166001145b620001525760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840162000087565b6000805460ff19166001179055801562000176576000805461ff0019166101001790555b8015620001bd576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050506200022a565b6001600160a01b03163b151590565b80516001600160a01b0381168114620001ed57600080fd5b919050565b600080604083850312156200020657600080fd5b6200021183620001d5565b91506200022160208401620001d5565b90509250929050565b60805160a05160c0516155436200027d600039600061335e0152600081816110a3015281816110e3015281816112c901528181611309015261142d0152600081816106c0015261437201526155436000f3fe6080604052600436106103765760003560e01c80637434d8e7116101d1578063a9048b5711610102578063c9580804116100a0578063e873c9dd1161006f578063e873c9dd14610afb578063e985e9c514610b29578063f04dc4b714610b73578063f65cfc9614610b8a57600080fd5b8063c958080414610a86578063d5c35af414610aa6578063d858a7e514610ac6578063e11cf71e14610adb57600080fd5b8063b88d4fde116100dc578063b88d4fde14610a08578063c72d33d814610a28578063c80678ef14610a46578063c87b56dd14610a6657600080fd5b8063a9048b571461099b578063ac9650d8146109bb578063b7fb70f6146109e857600080fd5b80638e79a3691161016f578063911b7b3011610149578063911b7b301461092757806395d89b41146109475780639f79b63a1461095c578063a22cb4651461097b57600080fd5b80638e79a369146108c75780638fb270e0146108e757806390b717e21461090757600080fd5b8063773ed13c116101ab578063773ed13c1461081c5780637dfe7c421461083c57806382fe1bcc1461085c578063841280ac1461088357600080fd5b80637434d8e71461079057806375b30be6146107c2578063760d0d39146107e257600080fd5b8063408e35f0116102ab57806354fd4d50116102495780635a74fc29116102235780635a74fc29146107105780636352211e146107305780636877063a1461075057806370a082311461077057600080fd5b806354fd4d5014610672578063572b6c05146106a3578063579a6988146106f057600080fd5b8063474dd82111610285578063474dd8211461060a5780634f1ef2861461062a5780634f6ccce71461063d57806352d1902d1461065d57600080fd5b8063408e35f0146105aa57806342842e0e146105ca57806344014b6b146105ea57600080fd5b80631d734bf2116103185780632f745c59116102f25780632f745c591461050b5780633121db1c1461052b578063331b2c8a1461054b5780633659cfe61461058a57600080fd5b80631d734bf21461049e5780631e93b8e0146104be57806323b872dd146104eb57600080fd5b8063095ea7b311610354578063095ea7b31461040a578063113f65c81461042c57806318160ddd1461045a5780631c91a9a41461047057600080fd5b806301ffc9a71461037b57806306fdde03146103b0578063081812fc146103d2575b600080fd5b34801561038757600080fd5b5061039b6103963660046147e2565b610baa565b60405190151581526020015b60405180910390f35b3480156103bc57600080fd5b506103c5610bbb565b6040516103a79190614857565b3480156103de57600080fd5b506103f26103ed36600461486a565b610c4e565b6040516001600160a01b0390911681526020016103a7565b34801561041657600080fd5b5061042a610425366004614898565b610c76565b005b34801561043857600080fd5b5061044c61044736600461486a565b610da3565b6040519081526020016103a7565b34801561046657600080fd5b506101935461044c565b34801561047c57600080fd5b5061044c61048b36600461486a565b60009081526101ff602052604090205490565b3480156104aa57600080fd5b5061044c6104b9366004614898565b610db6565b3480156104ca57600080fd5b506104de6104d93660046148c4565b610e36565b6040516103a791906148e6565b3480156104f757600080fd5b5061042a61050636600461492d565b610f7e565b34801561051757600080fd5b5061044c610526366004614898565b610fb6565b34801561053757600080fd5b5061042a6105463660046149b0565b61104d565b34801561055757600080fd5b5061039b610566366004614a05565b6001600160a01b031660009081526101fd6020526040902054610100900460ff1690565b34801561059657600080fd5b5061042a6105a5366004614a05565b611098565b3480156105b657600080fd5b5061042a6105c5366004614a05565b611178565b3480156105d657600080fd5b5061042a6105e536600461492d565b61121f565b3480156105f657600080fd5b5061042a610605366004614a05565b61123a565b34801561061657600080fd5b5061042a610625366004614a48565b61125e565b61042a610638366004614b3d565b6112be565b34801561064957600080fd5b5061044c61065836600461486a565b61138b565b34801561066957600080fd5b5061044c611420565b34801561067e57600080fd5b506103c5604051806040016040528060058152602001640302e312e360dc1b81525081565b3480156106af57600080fd5b5061039b6106be366004614a05565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0390811691161490565b3480156106fc57600080fd5b5061039b61070b36600461486a565b6114d3565b34801561071c57600080fd5b5061042a61072b366004614b8d565b6114f3565b34801561073c57600080fd5b506103f261074b36600461486a565b6115dc565b34801561075c57600080fd5b506104de61076b366004614a05565b6115e7565b34801561077c57600080fd5b5061044c61078b366004614a05565b611714565b34801561079c57600080fd5b506107b06107ab366004614a05565b61179b565b6040516103a796959493929190614bcf565b3480156107ce57600080fd5b5061042a6107dd366004614c18565b6117fb565b3480156107ee57600080fd5b5061039b6107fd366004614a05565b6001600160a01b031660009081526101fd602052604090205460ff1690565b34801561082857600080fd5b5061039b610837366004614cb8565b6118c9565b34801561084857600080fd5b5061042a610857366004614cdd565b6118e9565b34801561086857600080fd5b5061044c610877366004614a05565b6001600160a01b031690565b34801561088f57600080fd5b506108a361089e36600461486a565b611a93565b604080518251815260208084015190820152918101511515908201526060016103a7565b3480156108d357600080fd5b506103f26108e23660046148c4565b611af8565b3480156108f357600080fd5b5061044c61090236600461486a565b611b11565b34801561091357600080fd5b5061042a61092236600461486a565b611b29565b34801561093357600080fd5b5061039b610942366004614a05565b611b6c565b34801561095357600080fd5b506103c5611bfe565b34801561096857600080fd5b506101c3546001600160a01b03166103f2565b34801561098757600080fd5b5061042a610996366004614d39565b611c0e565b3480156109a757600080fd5b5061042a6109b6366004614a05565b611c20565b3480156109c757600080fd5b506109db6109d6366004614d67565b611cc3565b6040516103a79190614ddc565b3480156109f457600080fd5b5061042a610a03366004614e3e565b611db1565b348015610a1457600080fd5b5061042a610a23366004614e6f565b611ec3565b348015610a3457600080fd5b506103f2610a4336600461486a565b90565b348015610a5257600080fd5b5061039b610a61366004614898565b611efc565b348015610a7257600080fd5b506103c5610a8136600461486a565b611f15565b348015610a9257600080fd5b5061042a610aa1366004614a05565b611f88565b348015610ab257600080fd5b5061042a610ac13660046149b0565b612046565b348015610ad257600080fd5b5061042a6121aa565b348015610ae757600080fd5b5061044c610af636600461486a565b612235565b348015610b0757600080fd5b5061044c610b1636600461486a565b6000908152610201602052604090205490565b348015610b3557600080fd5b5061039b610b44366004614edb565b6001600160a01b0391821660009081526101646020908152604080832093909416825291909152205460ff1690565b348015610b7f57600080fd5b5061044c6102025481565b348015610b9657600080fd5b506103f2610ba53660046148c4565b61224d565b6000610bb582612275565b92915050565b606061015f8054610bcb90614f09565b80601f0160208091040260200160405190810160405280929190818152602001828054610bf790614f09565b8015610c445780601f10610c1957610100808354040283529160200191610c44565b820191906000526020600020905b815481529060010190602001808311610c2757829003601f168201915b5050505050905090565b6000610c598261229a565b50600090815261016360205260409020546001600160a01b031690565b6000610c81826122fa565b9050806001600160a01b0316836001600160a01b03161415610cf45760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b806001600160a01b0316610d0661235b565b6001600160a01b03161480610d225750610d2281610b4461235b565b610d945760405162461bcd60e51b815260206004820152603e60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206e6f7220617070726f76656420666f7220616c6c00006064820152608401610ceb565b610d9e838361236a565b505050565b6000610bb5610db061235b565b836123d9565b60007fcbe0462e67cb804f0011a6c0b71e9ff49be70d0f907ffdf4f3c74499282ab0b1610dea81610de561235b565b6124bf565b610e225780610df761235b565b6040516301d4003760e61b815260048101929092526001600160a01b03166024820152604401610ceb565b610e2c84846123d9565b91505b5092915050565b610e6c6040518060a001604052806000151581526020016000151581526020016000815260200160008152602001606081525090565b60008381526101fe602052604081206101fd9190610e8a9085612544565b6001600160a01b031681526020808201929092526040908101600020815160a081018352815460ff808216151583526101009091041615159381019390935260018101549183019190915260028101546060830152600381018054608084019190610ef490614f09565b80601f0160208091040260200160405190810160405280929190818152602001828054610f2090614f09565b8015610f6d5780601f10610f4257610100808354040283529160200191610f6d565b820191906000526020600020905b815481529060010190602001808311610f5057829003601f168201915b505050505081525050905092915050565b610f8f610f8961235b565b82612550565b610fab5760405162461bcd60e51b8152600401610ceb90614f3e565b610d9e8383836125cf565b6000610fc183611714565b82106110235760405162461bcd60e51b815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201526a74206f6620626f756e647360a81b6064820152608401610ceb565b506001600160a01b0391909116600090815261019160209081526040808320938352929052205490565b7f664245c7af190fec316596e8231f724e8171b1966cfcd124347ac5a66c34f64a61107a81610de561235b565b6110875780610df761235b565b611092848484612779565b50505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156110e15760405162461bcd60e51b8152600401610ceb90614f8c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661112a6000805160206154a7833981519152546001600160a01b031690565b6001600160a01b0316146111505760405162461bcd60e51b8152600401610ceb90614fd8565b6111598161288f565b60408051600080825260208201909252611175918391906128c9565b50565b8061119c816001600160a01b031660009081526101fd602052604090205460ff1690565b6111c457604051631cc39a1f60e01b81526001600160a01b0382166004820152602401610ceb565b6111cd82612a43565b6111ea5760405163a46f5f1160e01b815260040160405180910390fd5b6111f5826001612a80565b6001600160a01b03821660009081526101fd602052604090206001015461121b90612b0e565b5050565b610d9e83838360405180602001604052806000815250611ec3565b600061124881610de561235b565b6112555780610df761235b565b61121b82612b82565b7fcbe0462e67cb804f0011a6c0b71e9ff49be70d0f907ffdf4f3c74499282ab0b161128b81610de561235b565b6112985780610df761235b565b6112a183612c15565b8115610d9e57610d9e6112b76020850185614a05565b6001612a80565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156113075760405162461bcd60e51b8152600401610ceb90614f8c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166113506000805160206154a7833981519152546001600160a01b031690565b6001600160a01b0316146113765760405162461bcd60e51b8152600401610ceb90614fd8565b61137f8261288f565b61121b828260016128c9565b60006113976101935490565b82106113fa5760405162461bcd60e51b815260206004820152602c60248201527f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60448201526b7574206f6620626f756e647360a01b6064820152608401610ceb565b610193828154811061140e5761140e615024565b90600052602060002001549050919050565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146114c05760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610ceb565b506000805160206154a783398151915290565b600081815261016160205260408120546001600160a01b03161515610bb5565b826114fd816115dc565b6001600160a01b031661150e61235b565b6001600160a01b0316146115505761152461235b565b60405163694c4e1560e01b81526001600160a01b03909116600482015260248101829052604401610ceb565b811561157557600084815261022f6020526040902061156f9084612e51565b50611590565b600084815261022f6020526040902061158e9084612e66565b505b826001600160a01b0316847f538b6537a6fe8f0deae9f3b86ad1924d5e5b3d5a683055276b2824f918be043e846040516115ce911515815260200190565b60405180910390a350505050565b6000610bb5826122fa565b61161d6040518060a001604052806000151581526020016000151581526020016000815260200160008152602001606081525090565b6001600160a01b03821660009081526101fd6020908152604091829020825160a081018452815460ff80821615158352610100909104161515928101929092526001810154928201929092526002820154606082015260038201805491929160808401919061168b90614f09565b80601f01602080910402602001604051908101604052809291908181526020018280546116b790614f09565b80156117045780601f106116d957610100808354040283529160200191611704565b820191906000526020600020905b8154815290600101906020018083116116e757829003601f168201915b5050505050815250509050919050565b60006001600160a01b03821661177e5760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b6064820152608401610ceb565b506001600160a01b03166000908152610162602052604090205490565b6000806000606060008060006117b0886115e7565b8051909150806117c15760006117ce565b6117ce82604001516115dc565b826060015183608001516117e18c611b6c565b602090950151939c929b5090995097509195509350915050565b600054610100900460ff161580801561181b5750600054600160ff909116105b806118355750303b158015611835575060005460ff166001145b6118515760405162461bcd60e51b8152600401610ceb9061503a565b6000805460ff191660011790558015611874576000805461ff0019166101001790555b61187d88612e7b565b61188b878787878787612f3d565b80156118bf576000805461ff0019169055604051600181526000805160206154c78339815191529060200160405180910390a15b5050505050505050565b600082815261022f602052604081206118e290836130ce565b9392505050565b82602001356118f7816115dc565b6001600160a01b031661190861235b565b6001600160a01b03161461191e5761152461235b565b42610202548560800135611932919061509e565b10156119515760405163e794805160e01b815260040160405180910390fd5b611a6d6119616020860186614a05565b611a317f7cb62875b2afadb4cb4ed8910346f6f929b3380857391a94d729cbd47d406af16119926020890189614a05565b602089013560408a01356119a960608c018c6150b6565b6040516020016119ba9291906150fd565b604051602081830303815290604052805190602001208b60800135604051602001611a16969594939291909586526001600160a01b0394909416602086015260408501929092526060840152608083015260a082015260c00190565b604051602081830303815290604052805190602001206130f0565b85858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061313e92505050565b611a8a57604051632b84c44f60e11b815260040160405180910390fd5b61109284612c15565b611ab9604051806060016040528060008152602001600081526020016000151581525090565b50600090815261020060209081526040918290208251606081018452815481526001820154928101929092526002015460ff1615159181019190915290565b600082815261022f602052604081206118e29083612544565b60008181526101fe60205260408120610bb59061328a565b7f5f1b4fe6fc447c4d93b2d928a24950315e60be9356ac67fdbfe6aa3d8a985b4d611b5681610de561235b565b611b635780610df761235b565b61121b82613294565b6001600160a01b03811660009081526101fd60209081526040808320600281015484526102009092528220815460ff168015611baf57508154610100900460ff16155b8015611bce5750600281015460ff161580611bce5750611bce84613306565b8015611bf657506001820154600090815261016160205260409020546001600160a01b031615155b949350505050565b60606101608054610bcb90614f09565b61121b611c1961235b565b83836133e4565b80611c44816001600160a01b031660009081526101fd602052604090205460ff1690565b611c6c57604051631cc39a1f60e01b81526001600160a01b0382166004820152602401610ceb565b611c7582612a43565b611c925760405163a46f5f1160e01b815260040160405180910390fd5b611c9d826000612a80565b6001600160a01b03821660009081526101fd602052604090206001015461121b906134ac565b60608167ffffffffffffffff811115611cde57611cde614a9a565b604051908082528060200260200182016040528015611d1157816020015b6060815260200190600190039081611cfc5790505b50905060005b82811015610e2f57611d8130858584818110611d3557611d35615024565b9050602002810190611d4791906150b6565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506134cc92505050565b828281518110611d9357611d93615024565b60200260200101819052508080611da99061510d565b915050611d17565b7f5f1b4fe6fc447c4d93b2d928a24950315e60be9356ac67fdbfe6aa3d8a985b4d611dde81610de561235b565b611deb5780610df761235b565b81611e23576040516303b3e63560e41b815260206004820152600760248201526618da185a5b925960ca1b6044820152606401610ceb565b8235602084013511611e4857604051632ca637fd60e21b815260040160405180910390fd5b817fc89d7e1caf32a74dc9fc2c0197a92afe1936d38c73dd808b0f92fc50677b91ff84356020860135611e816060880160408901615128565b60408051938452602084019290925215159082015260600160405180910390a26000828152610200602052604090208390611ebc8282615145565b5050505050565b611ed4611ece61235b565b83612550565b611ef05760405162461bcd60e51b8152600401610ceb90614f3e565b611092848484846134f1565b60008181526101fe602052604081206118e290846130ce565b6060611f208261229a565b6000611f3760408051602081019091526000815290565b90506000815111611f5757604051806020016040528060008152506118e2565b80611f6184613524565b604051602001611f7292919061517b565b6040516020818303038152906040529392505050565b6000611f9681610de561235b565b611fa35780610df761235b565b611fbd6001600160a01b038316637965db0b60e01b613622565b611ffb576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b6044820152606401610ceb565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a25050565b6001600160a01b03831660009081526101fd602052604090205460ff1661208b57604051631cc39a1f60e01b81526001600160a01b0384166004820152602401610ceb565b6001600160a01b03831660009081526101fd60205260409020600101546120b1906115dc565b6001600160a01b03166120c261235b565b6001600160a01b03161461211a576120d861235b565b6001600160a01b0384811660009081526101fd60205260409081902060010154905163694c4e1560e01b81529290911660048301526024820152604401610ceb565b6001600160a01b03831660009081526101fd602052604090206121419060030183836146bf565b506001600160a01b03831660008181526101fd602052604090819020600281015460019091015491519092917ffc43d422261aa008210f185c9a82dfeb4faa37fcdac42e357e80bf0a3446e3e89161219d9187918791906151d3565b60405180910390a3505050565b6065546001600160a01b03166121f85760405163eac0d38960e01b81526020600482015260126024820152712fb232b83932b1b0ba32b22fb937baba32b960711b6044820152606401610ceb565b606580546001600160a01b03191690556040516000907f7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc80908290a2565b600081815261022f60205260408120610bb59061328a565b60008281526101fe602052604081206118e29083612544565b6001600160a01b03163b151590565b60006001600160e01b0319821663780e9d6360e01b1480610bb55750610bb58261363e565b600081815261016160205260409020546001600160a01b03166111755760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b6044820152606401610ceb565b600081815261016160205260408120546001600160a01b031680610bb55760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b6044820152606401610ceb565b600061236561368e565b905090565b60008181526101636020526040902080546001600160a01b0319166001600160a01b03841690811790915581906123a0826122fa565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b60006001600160a01b0383166124275760405163eac0d38960e01b81526020600482015260126024820152717363616e6e6572506f6f6c4164647265737360701b6044820152606401610ceb565b8161245f576040516303b3e63560e41b815260206004820152600760248201526618da185a5b925960ca1b6044820152606401610ceb565b61246e6101fc80546001019055565b506101fc5461247d8382613698565b6000818152610201602052604080822084905551839183917f7e191d1226c4da936653bac3b9ae5c76dd880cf5221a80bcaa224448ccc670489190a392915050565b603354604051632474521560e21b8152600481018490526001600160a01b03838116602483015260009216906391d148549060440160206040518083038186803b15801561250c57600080fd5b505afa158015612520573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118e291906151f7565b60006118e283836136b2565b60008061255c836122fa565b9050806001600160a01b0316846001600160a01b031614806125a457506001600160a01b038082166000908152610164602090815260408083209388168352929052205460ff165b80610e2c5750836001600160a01b03166125bd84610c4e565b6001600160a01b031614949350505050565b826001600160a01b03166125e2826122fa565b6001600160a01b0316146126465760405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201526437bbb732b960d91b6064820152608401610ceb565b6001600160a01b0382166126a85760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610ceb565b6126b38383836136dc565b6126be60008261236a565b6001600160a01b0383166000908152610162602052604081208054600192906126e8908490615214565b90915550506001600160a01b03821660009081526101626020526040812080546001929061271790849061509e565b90915550506000818152610161602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6040516302571be360e01b81527f91d1777781884d03a6757a803996e38de2a42967fb37eeaca72729271025a9e260048201526001600160a01b038416906302571be39060240160206040518083038186803b1580156127d857600080fd5b505afa1580156127ec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612810919061522b565b6001600160a01b031663c47f002783836040518363ffffffff1660e01b815260040161283d929190615248565b602060405180830381600087803b15801561285757600080fd5b505af115801561286b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611092919061525c565b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e36128bc81610de561235b565b61121b5780610df761235b565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff16156128fc57610d9e836136e7565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561293557600080fd5b505afa925050508015612965575060408051601f3d908101601f191682019092526129629181019061525c565b60015b6129c85760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610ceb565b6000805160206154a78339815191528114612a375760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610ceb565b50610d9e838383613783565b6000612a4e826137a8565b80610bb55750610bb57fcbe0462e67cb804f0011a6c0b71e9ff49be70d0f907ffdf4f3c74499282ab0b1610de561235b565b6001600160a01b03821660009081526101fd60205260409020805461ff00191661010083151502179055612ab382611b6c565b15156001600160a01b0383167f800131ba226bbb1e9c537607e0f4d73e692ae496e7d1fb9b2aea059369d37fa2612ae861235b565b604080516001600160a01b03909216825285151560208301520160405180910390a35050565b60008181526101ff60205260408120805460019290612b2e908490615214565b909155505060008181526101ff60205260409081902054905182917fe98aca89d957a1075e807d80a8f0f7d280340888edefe3da76c278ca302922c891612b7791815260200190565b60405180910390a250565b6001600160a01b038116612bca5760405163eac0d38960e01b815260206004820152600e60248201526d7375626a6563744761746577617960901b6044820152606401610ceb565b6101c380546001600160a01b0319166001600160a01b0383169081179091556040517f16d72e484786227d7b6dd05e079be9ff9904a81a0a9ec36fc346b20f8c47aff090600090a250565b612c256107fd6020830183614a05565b15612c5c57612c376020820182614a05565b6040516371cf021b60e01b81526001600160a01b039091166004820152602401610ceb565b80604001356102016000836020013581526020019081526020016000205414612cbc576020808201356000908152610201909152604090819020548151630432cec160e31b8152610ceb9284013590600401918252602082015260400190565b6040518060a001604052806001151581526020016000151581526020018260200135815260200182604001358152602001828060600190612cfd91906150b6565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525093909452506101fd9250612d4690506020850185614a05565b6001600160a01b031681526020808201929092526040908101600020835181548585015115156101000261ff00199215159290921661ffff199091161717815590830151600182015560608301516002820155608083015180519192612db492600385019290910190614743565b50612de49150612dc990506020830183614a05565b60208084013560009081526101fe9091526040902090612e51565b506040810135612dfa6108776020840184614a05565b7ffc43d422261aa008210f185c9a82dfeb4faa37fcdac42e357e80bf0a3446e3e8612e2860608501856150b6565b8560200135604051612e3c939291906151d3565b60405180910390a361117581602001356134ac565b60006118e2836001600160a01b0384166137e2565b60006118e2836001600160a01b038416613831565b600054610100900460ff1615808015612e9b5750600054600160ff909116105b80612eb55750303b158015612eb5575060005460ff166001145b612ed15760405162461bcd60e51b8152600401610ceb9061503a565b6000805460ff191660011790558015612ef4576000805461ff0019166101001790555b612efd82613924565b612f05613a6c565b801561121b576000805461ff0019169055604051600181526000805160206154c7833981519152906020015b60405180910390a15050565b600054610100900460ff1615808015612f5d5750600054600160ff909116105b80612f775750303b158015612f77575060005460ff166001145b612f935760405162461bcd60e51b8152600401610ceb9061503a565b6000805460ff191660011790558015612fb6576000805461ff0019166101001790555b61302987878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8b018190048102820181019092528981529250899150889081908401838280828437600092019190915250613a9592505050565b613031613a6c565b61307f604051806040016040528060138152602001725363616e6e6572506f6f6c526567697374727960681b815250604051806040016040528060018152602001603160f81b815250613ac6565b61308883613af7565b61309182613294565b80156130c5576000805461ff0019169055604051600181526000805160206154c78339815191529060200160405180910390a15b50505050505050565b6001600160a01b038116600090815260018301602052604081205415156118e2565b6000610bb56130fd613b79565b8360405161190160f01b6020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b600080600061314d8585613bf6565b9092509050600081600481111561316657613166615275565b1480156131845750856001600160a01b0316826001600160a01b0316145b15613194576001925050506118e2565b600080876001600160a01b0316631626ba7e60e01b88886040516024016131bc92919061528b565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516131fa91906152a4565b600060405180830381855afa9150503d8060008114613235576040519150601f19603f3d011682016040523d82523d6000602084013e61323a565b606091505b509150915081801561324d575080516020145b801561327e57508051630b135d3f60e11b9061327290830160209081019084016152c0565b6001600160e01b031916145b98975050505050505050565b6000610bb5825490565b806132ca576040516303b3e63560e41b815260206004820152600560248201526464656c617960d81b6044820152606401610ceb565b6102028190556040518181527fbb57183c638abde4f4007bb9e16e4dc9fe60ec4e316cc19e9329d9782cabaeac9060200160405180910390a150565b6001600160a01b0381811660009081526101fd6020908152604080832060028082015485526102009093528184208054600183015493516369ce5e8960e11b81526004810195909552602485019390935293949093927f00000000000000000000000000000000000000000000000000000000000000009091169063d39cbd129060440160206040518083038186803b1580156133a257600080fd5b505afa1580156133b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133da919061525c565b1015949350505050565b816001600160a01b0316836001600160a01b031614156134465760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610ceb565b6001600160a01b0383811660008181526101646020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910161219d565b60008181526101ff60205260408120805460019290612b2e90849061509e565b60606118e283836040518060600160405280602781526020016154e760279139613c66565b6134fc8484846125cf565b61350884848484613d04565b6110925760405162461bcd60e51b8152600401610ceb906152dd565b6060816135485750506040805180820190915260018152600360fc1b602082015290565b8160005b8115613572578061355c8161510d565b915061356b9050600a83615345565b915061354c565b60008167ffffffffffffffff81111561358d5761358d614a9a565b6040519080825280601f01601f1916602001820160405280156135b7576020820181803683370190505b5090505b8415611bf6576135cc600183615214565b91506135d9600a86615359565b6135e490603061509e565b60f81b8183815181106135f9576135f9615024565b60200101906001600160f81b031916908160001a90535061361b600a86615345565b94506135bb565b600061362d83613e18565b80156118e257506118e28383613e4b565b60006001600160e01b031982166380ac58cd60e01b148061366f57506001600160e01b03198216635b5e139f60e01b145b80610bb557506301ffc9a760e01b6001600160e01b0319831614610bb5565b6000612365613f2a565b61121b828260405180602001604052806000815250613f34565b60008260000182815481106136c9576136c9615024565b9060005260206000200154905092915050565b610d9e838383613f67565b6001600160a01b0381163b6137545760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610ceb565b6000805160206154a783398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b61378c83614021565b6000825111806137995750805b15610d9e576110928383614061565b60006137b382614116565b80610bb557506001600160a01b03821660009081526101fd6020526040902060010154610bb59061083761235b565b600081815260018301602052604081205461382957508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610bb5565b506000610bb5565b6000818152600183016020526040812054801561391a576000613855600183615214565b855490915060009061386990600190615214565b90508181146138ce57600086600001828154811061388957613889615024565b90600052602060002001549050808760000184815481106138ac576138ac615024565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806138df576138df61536d565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610bb5565b6000915050610bb5565b600054610100900460ff16158080156139445750600054600160ff909116105b8061395e5750303b15801561395e575060005460ff166001145b61397a5760405162461bcd60e51b8152600401610ceb9061503a565b6000805460ff19166001179055801561399d576000805461ff0019166101001790555b6139b76001600160a01b038316637965db0b60e01b613622565b6139f5576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b6044820152606401610ceb565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a2801561121b576000805461ff0019169055604051600181526000805160206154c783398151915290602001612f31565b600054610100900460ff16613a935760405162461bcd60e51b8152600401610ceb90615383565b565b600054610100900460ff16613abc5760405162461bcd60e51b8152600401610ceb90615383565b61121b828261417c565b600054610100900460ff16613aed5760405162461bcd60e51b8152600401610ceb90615383565b61121b82826141cc565b600054610100900460ff1615808015613b175750600054600160ff909116105b80613b315750303b158015613b31575060005460ff166001145b613b4d5760405162461bcd60e51b8152600401610ceb9061503a565b6000805460ff191660011790558015613b70576000805461ff0019166101001790555b612f0582612b82565b60006123657f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f613ba96101c85490565b6101c9546040805160208101859052908101839052606081018290524660808201523060a082015260009060c0016040516020818303038152906040528051906020012090509392505050565b600080825160411415613c2d5760208301516040840151606085015160001a613c218782858561420f565b94509450505050613c5f565b825160401415613c575760208301516040840151613c4c8683836142fc565b935093505050613c5f565b506000905060025b9250929050565b60606001600160a01b0384163b613c8f5760405162461bcd60e51b8152600401610ceb906153ce565b600080856001600160a01b031685604051613caa91906152a4565b600060405180830381855af49150503d8060008114613ce5576040519150601f19603f3d011682016040523d82523d6000602084013e613cea565b606091505b5091509150613cfa828286614335565b9695505050505050565b60006001600160a01b0384163b15613e0d57836001600160a01b031663150b7a02613d2d61235b565b8786866040518563ffffffff1660e01b8152600401613d4f9493929190615414565b602060405180830381600087803b158015613d6957600080fd5b505af1925050508015613d99575060408051601f3d908101601f19168201909252613d96918101906152c0565b60015b613df3573d808015613dc7576040519150601f19603f3d011682016040523d82523d6000602084013e613dcc565b606091505b508051613deb5760405162461bcd60e51b8152600401610ceb906152dd565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050611bf6565b506001949350505050565b6000613e2b826301ffc9a760e01b613e4b565b8015610bb55750613e44826001600160e01b0319613e4b565b1592915050565b604080516001600160e01b0319831660248083019190915282518083039091018152604490910182526020810180516001600160e01b03166301ffc9a760e01b179052905160009190829081906001600160a01b0387169061753090613eb29086906152a4565b6000604051808303818686fa925050503d8060008114613eee576040519150601f19603f3d011682016040523d82523d6000602084013e613ef3565b606091505b5091509150602081511015613f0e5760009350505050610bb5565b818015613cfa575080806020019051810190613cfa91906151f7565b600061236561436e565b613f3e83836143d1565b613f4b6000848484613d04565b610d9e5760405162461bcd60e51b8152600401610ceb906152dd565b6001600160a01b038316613fc457613fbf816101938054600083815261019460205260408120829055600182018355919091527ffc8af01f449989052b52093a58fc9f42d0b11f0c6dd5dca0463dab62346ccc680155565b613fe7565b816001600160a01b0316836001600160a01b031614613fe757613fe78382614522565b6001600160a01b038216613ffe57610d9e816145c4565b826001600160a01b0316826001600160a01b031614610d9e57610d9e8282614679565b61402a816136e7565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606001600160a01b0383163b61408a5760405162461bcd60e51b8152600401610ceb906153ce565b600080846001600160a01b0316846040516140a591906152a4565b600060405180830381855af49150503d80600081146140e0576040519150601f19603f3d011682016040523d82523d6000602084013e6140e5565b606091505b509150915061410d82826040518060600160405280602781526020016154e760279139614335565b95945050505050565b6000816001600160a01b031661412a61235b565b6001600160a01b03161480610bb5575061414261235b565b6001600160a01b0383811660009081526101fd602052604090206001015491169061416c906115dc565b6001600160a01b03161492915050565b600054610100900460ff166141a35760405162461bcd60e51b8152600401610ceb90615383565b81516141b79061015f906020850190614743565b508051610d9e90610160906020840190614743565b600054610100900460ff166141f35760405162461bcd60e51b8152600401610ceb90615383565b8151602092830120815191909201206101c8919091556101c955565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561424657506000905060036142f3565b8460ff16601b1415801561425e57508460ff16601c14155b1561426f57506000905060046142f3565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156142c3573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166142ec576000600192509250506142f3565b9150600090505b94509492505050565b6000806001600160ff1b0383168161431960ff86901c601b61509e565b90506143278782888561420f565b935093505050935093915050565b606083156143445750816118e2565b8251156143545782518084602001fd5b8160405162461bcd60e51b8152600401610ceb9190614857565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163314156143cc576000366143af601482615214565b6143bb92369290615447565b6143c491615471565b60601c905090565b503390565b6001600160a01b0382166144275760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610ceb565b600081815261016160205260409020546001600160a01b03161561448d5760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610ceb565b614499600083836136dc565b6001600160a01b0382166000908152610162602052604081208054600192906144c390849061509e565b90915550506000818152610161602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b6000600161452f84611714565b6145399190615214565b6000838152610192602052604090205490915080821461458f576001600160a01b038416600090815261019160209081526040808320858452825280832054848452818420819055835261019290915290208190555b506000918252610192602090815260408084208490556001600160a01b03909416835261019181528383209183525290812055565b610193546000906145d790600190615214565b60008381526101946020526040812054610193805493945090928490811061460157614601615024565b9060005260206000200154905080610193838154811061462357614623615024565b6000918252602080832090910192909255828152610194909152604080822084905585825281205561019380548061465d5761465d61536d565b6001900381819060005260206000200160009055905550505050565b600061468483611714565b6001600160a01b0390931660009081526101916020908152604080832086845282528083208590559382526101929052919091209190915550565b8280546146cb90614f09565b90600052602060002090601f0160209004810192826146ed5760008555614733565b82601f106147065782800160ff19823516178555614733565b82800160010185558215614733579182015b82811115614733578235825591602001919060010190614718565b5061473f9291506147b7565b5090565b82805461474f90614f09565b90600052602060002090601f0160209004810192826147715760008555614733565b82601f1061478a57805160ff1916838001178555614733565b82800160010185558215614733579182015b8281111561473357825182559160200191906001019061479c565b5b8082111561473f57600081556001016147b8565b6001600160e01b03198116811461117557600080fd5b6000602082840312156147f457600080fd5b81356118e2816147cc565b60005b8381101561481a578181015183820152602001614802565b838111156110925750506000910152565b600081518084526148438160208601602086016147ff565b601f01601f19169290920160200192915050565b6020815260006118e2602083018461482b565b60006020828403121561487c57600080fd5b5035919050565b6001600160a01b038116811461117557600080fd5b600080604083850312156148ab57600080fd5b82356148b681614883565b946020939093013593505050565b600080604083850312156148d757600080fd5b50508035926020909101359150565b6020815281511515602082015260208201511515604082015260408201516060820152606082015160808201526000608083015160a080840152610e2c60c084018261482b565b60008060006060848603121561494257600080fd5b833561494d81614883565b9250602084013561495d81614883565b929592945050506040919091013590565b60008083601f84011261498057600080fd5b50813567ffffffffffffffff81111561499857600080fd5b602083019150836020828501011115613c5f57600080fd5b6000806000604084860312156149c557600080fd5b83356149d081614883565b9250602084013567ffffffffffffffff8111156149ec57600080fd5b6149f88682870161496e565b9497909650939450505050565b600060208284031215614a1757600080fd5b81356118e281614883565b600060a08284031215614a3457600080fd5b50919050565b801515811461117557600080fd5b60008060408385031215614a5b57600080fd5b823567ffffffffffffffff811115614a7257600080fd5b614a7e85828601614a22565b9250506020830135614a8f81614a3a565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112614ac157600080fd5b813567ffffffffffffffff80821115614adc57614adc614a9a565b604051601f8301601f19908116603f01168101908282118183101715614b0457614b04614a9a565b81604052838152866020858801011115614b1d57600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060408385031215614b5057600080fd5b8235614b5b81614883565b9150602083013567ffffffffffffffff811115614b7757600080fd5b614b8385828601614ab0565b9150509250929050565b600080600060608486031215614ba257600080fd5b833592506020840135614bb481614883565b91506040840135614bc481614a3a565b809150509250925092565b861515815260018060a01b038616602082015284604082015260c060608201526000614bfe60c083018661482b565b93151560808301525090151560a090910152949350505050565b600080600080600080600060a0888a031215614c3357600080fd5b8735614c3e81614883565b9650602088013567ffffffffffffffff80821115614c5b57600080fd5b614c678b838c0161496e565b909850965060408a0135915080821115614c8057600080fd5b50614c8d8a828b0161496e565b9095509350506060880135614ca181614883565b809250506080880135905092959891949750929550565b60008060408385031215614ccb57600080fd5b823591506020830135614a8f81614883565b600080600060408486031215614cf257600080fd5b833567ffffffffffffffff80821115614d0a57600080fd5b614d1687838801614a22565b94506020860135915080821115614d2c57600080fd5b506149f88682870161496e565b60008060408385031215614d4c57600080fd5b8235614d5781614883565b91506020830135614a8f81614a3a565b60008060208385031215614d7a57600080fd5b823567ffffffffffffffff80821115614d9257600080fd5b818501915085601f830112614da657600080fd5b813581811115614db557600080fd5b8660208260051b8501011115614dca57600080fd5b60209290920196919550909350505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015614e3157603f19888603018452614e1f85835161482b565b94509285019290850190600101614e03565b5092979650505050505050565b6000808284036080811215614e5257600080fd5b6060811215614e6057600080fd5b50919360608501359350915050565b60008060008060808587031215614e8557600080fd5b8435614e9081614883565b93506020850135614ea081614883565b925060408501359150606085013567ffffffffffffffff811115614ec357600080fd5b614ecf87828801614ab0565b91505092959194509250565b60008060408385031215614eee57600080fd5b8235614ef981614883565b91506020830135614a8f81614883565b600181811c90821680614f1d57607f821691505b60208210811415614a3457634e487b7160e01b600052602260045260246000fd5b6020808252602e908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526d1c881b9bdc88185c1c1c9bdd995960921b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b600082198211156150b1576150b1615088565b500190565b6000808335601e198436030181126150cd57600080fd5b83018035915067ffffffffffffffff8211156150e857600080fd5b602001915036819003821315613c5f57600080fd5b8183823760009101908152919050565b600060001982141561512157615121615088565b5060010190565b60006020828403121561513a57600080fd5b81356118e281614a3a565b813581556020820135600182015560028101604083013561516581614a3a565b815490151560ff1660ff19919091161790555050565b6000835161518d8184602088016147ff565b8351908301906151a18183602088016147ff565b01949350505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6040815260006151e76040830185876151aa565b9050826020830152949350505050565b60006020828403121561520957600080fd5b81516118e281614a3a565b60008282101561522657615226615088565b500390565b60006020828403121561523d57600080fd5b81516118e281614883565b602081526000611bf66020830184866151aa565b60006020828403121561526e57600080fd5b5051919050565b634e487b7160e01b600052602160045260246000fd5b828152604060208201526000611bf6604083018461482b565b600082516152b68184602087016147ff565b9190910192915050565b6000602082840312156152d257600080fd5b81516118e2816147cc565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b634e487b7160e01b600052601260045260246000fd5b6000826153545761535461532f565b500490565b6000826153685761536861532f565b500690565b634e487b7160e01b600052603160045260246000fd5b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b60208082526026908201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6040820152651b9d1c9858dd60d21b606082015260800190565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090613cfa9083018461482b565b6000808585111561545757600080fd5b8386111561546457600080fd5b5050820193919092039150565b6bffffffffffffffffffffffff19813581811691601485101561549e5780818660140360031b1b83161692505b50509291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212207c31ee26ed9db4c1ea3e3bb0d44f710e2dc38414354cd8a3c8dbb8c58c1b3c5e64736f6c63430008090033",
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
	parsed, err := abi.JSON(strings.NewReader(ScannerPoolRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistryCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _ScannerPoolRegistry.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistrySession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _ScannerPoolRegistry.Contract.IsTrustedForwarder(&_ScannerPoolRegistry.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_ScannerPoolRegistry *ScannerPoolRegistryCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _ScannerPoolRegistry.Contract.IsTrustedForwarder(&_ScannerPoolRegistry.CallOpts, forwarder)
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
