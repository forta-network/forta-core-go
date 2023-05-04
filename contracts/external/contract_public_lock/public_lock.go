// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_public_lock

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

// PublicLockMetaData contains all meta data concerning the PublicLock contract.
var PublicLockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CANNOT_APPROVE_SELF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CANT_BE_SMALLER_THAN_SUPPLY\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CANT_EXTEND_NON_EXPIRING_KEY\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GAS_REFUND_FAILED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INSUFFICIENT_ERC20_VALUE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INSUFFICIENT_VALUE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_ADDRESS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"hookIndex\",\"type\":\"uint8\"}],\"name\":\"INVALID_HOOK\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_LENGTH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KEY_NOT_VALID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KEY_TRANSFERS_DISABLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LOCK_HAS_CHANGED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LOCK_SOLD_OUT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MAX_KEYS_REACHED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MIGRATION_REQUIRED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NON_COMPLIANT_ERC721_RECEIVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NON_RENEWABLE_LOCK\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NOT_ENOUGH_FUNDS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NOT_ENOUGH_TIME\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NOT_READY_FOR_RENEWAL\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NO_SUCH_KEY\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NULL_VALUE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ONLY_KEY_MANAGER_OR_APPROVED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ONLY_LOCK_MANAGER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ONLY_LOCK_MANAGER_OR_KEY_GRANTER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OUT_OF_RANGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OWNER_CANT_BE_ADDRESS_ZERO\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SCHEMA_VERSION_NOT_CORRECT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TRANSFER_TO_SELF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UNAUTHORIZED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UNAUTHORIZED_KEY_MANAGER_UPDATE\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sendTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"}],\"name\":\"CancelKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newExpiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"timeAdded\",\"type\":\"bool\"}],\"name\":\"ExpirationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ExpireKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"GasRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"name\":\"KeyExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"KeyGranterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"KeyGranterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newManager\",\"type\":\"address\"}],\"name\":\"KeyManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedRecordsCount\",\"type\":\"uint256\"}],\"name\":\"KeysMigrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxNumberOfKeys\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxKeysPerAcccount\",\"type\":\"uint256\"}],\"name\":\"LockConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"LockManagerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"LockManagerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseTokenURI\",\"type\":\"string\"}],\"name\":\"LockMetadata\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldKeyPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"keyPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"PricingChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"freeTrialLength\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundPenaltyBasisPoints\",\"type\":\"uint256\"}],\"name\":\"RefundPenaltyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transferFeeBasisPoints\",\"type\":\"uint256\"}],\"name\":\"TransferFeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lockAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"unlockAddress\",\"type\":\"address\"}],\"name\":\"UnlockCallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KEY_GRANTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LOCK_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addKeyGranter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addLockManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_approved\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keyOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAndRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expirationDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"expireAndRefundFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"extend\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freeTrialLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasRefundValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_refundValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getCancelAndRefundValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keyOwner\",\"type\":\"address\"}],\"name\":\"getHasValidKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"getTransferFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"grantKeyExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_expirationTimestamps\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_keyManagers\",\"type\":\"address[]\"}],\"name\":\"grantKeys\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_lockCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_expirationDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_keyPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxNumberOfKeys\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_lockName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isKeyGranter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isLockManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"isValidKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"keyExpirationTimestampFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"keyManagerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keyPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"lendKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxKeysPerAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxNumberOfKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenIdFrom\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenIdTo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mergeKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfOwners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onKeyCancelHook\",\"outputs\":[{\"internalType\":\"contractILockKeyCancelHook\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onKeyExtendHook\",\"outputs\":[{\"internalType\":\"contractILockKeyExtendHook\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onKeyGrantHook\",\"outputs\":[{\"internalType\":\"contractILockKeyGrantHook\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onKeyPurchaseHook\",\"outputs\":[{\"internalType\":\"contractILockKeyPurchaseHook\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onKeyTransferHook\",\"outputs\":[{\"internalType\":\"contractILockKeyTransferHook\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onTokenURIHook\",\"outputs\":[{\"internalType\":\"contractILockTokenURIHook\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onValidKeyHook\",\"outputs\":[{\"internalType\":\"contractILockValidKeyHook\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicLockVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_recipients\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_referrers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_keyManagers\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_data\",\"type\":\"bytes[]\"}],\"name\":\"purchase\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"purchasePriceFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minKeyPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"referrerFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundPenaltyBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"}],\"name\":\"renewMembershipFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceLockManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_granter\",\"type\":\"address\"}],\"name\":\"revokeKeyGranter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"schemaVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_onKeyPurchaseHook\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_onKeyCancelHook\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_onValidKeyHook\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_onTokenURIHook\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_onKeyTransferHook\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_onKeyExtendHook\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_onKeyGrantHook\",\"type\":\"address\"}],\"name\":\"setEventHooks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_refundValue\",\"type\":\"uint256\"}],\"name\":\"setGasRefundValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_keyManager\",\"type\":\"address\"}],\"name\":\"setKeyManagerOf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_lockName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_lockSymbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_baseTokenURI\",\"type\":\"string\"}],\"name\":\"setLockMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_feeBasisPoint\",\"type\":\"uint256\"}],\"name\":\"setReferrerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenIdFrom\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeShared\",\"type\":\"uint256\"}],\"name\":\"shareKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keyOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keyOwner\",\"type\":\"address\"}],\"name\":\"totalKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_valueBasisPoint\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"unlendKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockProtocol\",\"outputs\":[{\"internalType\":\"contractIUnlock\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_keyPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"updateKeyPricing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newExpirationDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxNumberOfKeys\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxKeysPerAcccount\",\"type\":\"uint256\"}],\"name\":\"updateLockConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_freeTrialLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_refundPenaltyBasisPoints\",\"type\":\"uint256\"}],\"name\":\"updateRefundPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateSchemaVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_transferFeeBasisPoints\",\"type\":\"uint256\"}],\"name\":\"updateTransferFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PublicLockABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicLockMetaData.ABI instead.
var PublicLockABI = PublicLockMetaData.ABI

// PublicLock is an auto generated Go binding around an Ethereum contract.
type PublicLock struct {
	PublicLockCaller     // Read-only binding to the contract
	PublicLockTransactor // Write-only binding to the contract
	PublicLockFilterer   // Log filterer for contract events
}

// PublicLockCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicLockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicLockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicLockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicLockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicLockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicLockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicLockSession struct {
	Contract     *PublicLock       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PublicLockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicLockCallerSession struct {
	Contract *PublicLockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PublicLockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicLockTransactorSession struct {
	Contract     *PublicLockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PublicLockRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicLockRaw struct {
	Contract *PublicLock // Generic contract binding to access the raw methods on
}

// PublicLockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicLockCallerRaw struct {
	Contract *PublicLockCaller // Generic read-only contract binding to access the raw methods on
}

// PublicLockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicLockTransactorRaw struct {
	Contract *PublicLockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicLock creates a new instance of PublicLock, bound to a specific deployed contract.
func NewPublicLock(address common.Address, backend bind.ContractBackend) (*PublicLock, error) {
	contract, err := bindPublicLock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicLock{PublicLockCaller: PublicLockCaller{contract: contract}, PublicLockTransactor: PublicLockTransactor{contract: contract}, PublicLockFilterer: PublicLockFilterer{contract: contract}}, nil
}

// NewPublicLockCaller creates a new read-only instance of PublicLock, bound to a specific deployed contract.
func NewPublicLockCaller(address common.Address, caller bind.ContractCaller) (*PublicLockCaller, error) {
	contract, err := bindPublicLock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicLockCaller{contract: contract}, nil
}

// NewPublicLockTransactor creates a new write-only instance of PublicLock, bound to a specific deployed contract.
func NewPublicLockTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicLockTransactor, error) {
	contract, err := bindPublicLock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicLockTransactor{contract: contract}, nil
}

// NewPublicLockFilterer creates a new log filterer instance of PublicLock, bound to a specific deployed contract.
func NewPublicLockFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicLockFilterer, error) {
	contract, err := bindPublicLock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicLockFilterer{contract: contract}, nil
}

// bindPublicLock binds a generic wrapper to an already deployed contract.
func bindPublicLock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PublicLockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicLock *PublicLockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicLock.Contract.PublicLockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicLock *PublicLockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicLock.Contract.PublicLockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicLock *PublicLockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicLock.Contract.PublicLockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicLock *PublicLockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicLock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicLock *PublicLockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicLock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicLock *PublicLockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicLock.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PublicLock *PublicLockCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PublicLock *PublicLockSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PublicLock.Contract.DEFAULTADMINROLE(&_PublicLock.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PublicLock *PublicLockCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PublicLock.Contract.DEFAULTADMINROLE(&_PublicLock.CallOpts)
}

// KEYGRANTERROLE is a free data retrieval call binding the contract method 0x23100509.
//
// Solidity: function KEY_GRANTER_ROLE() view returns(bytes32)
func (_PublicLock *PublicLockCaller) KEYGRANTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "KEY_GRANTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KEYGRANTERROLE is a free data retrieval call binding the contract method 0x23100509.
//
// Solidity: function KEY_GRANTER_ROLE() view returns(bytes32)
func (_PublicLock *PublicLockSession) KEYGRANTERROLE() ([32]byte, error) {
	return _PublicLock.Contract.KEYGRANTERROLE(&_PublicLock.CallOpts)
}

// KEYGRANTERROLE is a free data retrieval call binding the contract method 0x23100509.
//
// Solidity: function KEY_GRANTER_ROLE() view returns(bytes32)
func (_PublicLock *PublicLockCallerSession) KEYGRANTERROLE() ([32]byte, error) {
	return _PublicLock.Contract.KEYGRANTERROLE(&_PublicLock.CallOpts)
}

// LOCKMANAGERROLE is a free data retrieval call binding the contract method 0x8ca2fbad.
//
// Solidity: function LOCK_MANAGER_ROLE() view returns(bytes32)
func (_PublicLock *PublicLockCaller) LOCKMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "LOCK_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LOCKMANAGERROLE is a free data retrieval call binding the contract method 0x8ca2fbad.
//
// Solidity: function LOCK_MANAGER_ROLE() view returns(bytes32)
func (_PublicLock *PublicLockSession) LOCKMANAGERROLE() ([32]byte, error) {
	return _PublicLock.Contract.LOCKMANAGERROLE(&_PublicLock.CallOpts)
}

// LOCKMANAGERROLE is a free data retrieval call binding the contract method 0x8ca2fbad.
//
// Solidity: function LOCK_MANAGER_ROLE() view returns(bytes32)
func (_PublicLock *PublicLockCallerSession) LOCKMANAGERROLE() ([32]byte, error) {
	return _PublicLock.Contract.LOCKMANAGERROLE(&_PublicLock.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _keyOwner) view returns(uint256 balance)
func (_PublicLock *PublicLockCaller) BalanceOf(opts *bind.CallOpts, _keyOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "balanceOf", _keyOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _keyOwner) view returns(uint256 balance)
func (_PublicLock *PublicLockSession) BalanceOf(_keyOwner common.Address) (*big.Int, error) {
	return _PublicLock.Contract.BalanceOf(&_PublicLock.CallOpts, _keyOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _keyOwner) view returns(uint256 balance)
func (_PublicLock *PublicLockCallerSession) BalanceOf(_keyOwner common.Address) (*big.Int, error) {
	return _PublicLock.Contract.BalanceOf(&_PublicLock.CallOpts, _keyOwner)
}

// ExpirationDuration is a free data retrieval call binding the contract method 0x11a4c03a.
//
// Solidity: function expirationDuration() view returns(uint256)
func (_PublicLock *PublicLockCaller) ExpirationDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "expirationDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpirationDuration is a free data retrieval call binding the contract method 0x11a4c03a.
//
// Solidity: function expirationDuration() view returns(uint256)
func (_PublicLock *PublicLockSession) ExpirationDuration() (*big.Int, error) {
	return _PublicLock.Contract.ExpirationDuration(&_PublicLock.CallOpts)
}

// ExpirationDuration is a free data retrieval call binding the contract method 0x11a4c03a.
//
// Solidity: function expirationDuration() view returns(uint256)
func (_PublicLock *PublicLockCallerSession) ExpirationDuration() (*big.Int, error) {
	return _PublicLock.Contract.ExpirationDuration(&_PublicLock.CallOpts)
}

// FreeTrialLength is a free data retrieval call binding the contract method 0xa375cb05.
//
// Solidity: function freeTrialLength() view returns(uint256)
func (_PublicLock *PublicLockCaller) FreeTrialLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "freeTrialLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreeTrialLength is a free data retrieval call binding the contract method 0xa375cb05.
//
// Solidity: function freeTrialLength() view returns(uint256)
func (_PublicLock *PublicLockSession) FreeTrialLength() (*big.Int, error) {
	return _PublicLock.Contract.FreeTrialLength(&_PublicLock.CallOpts)
}

// FreeTrialLength is a free data retrieval call binding the contract method 0xa375cb05.
//
// Solidity: function freeTrialLength() view returns(uint256)
func (_PublicLock *PublicLockCallerSession) FreeTrialLength() (*big.Int, error) {
	return _PublicLock.Contract.FreeTrialLength(&_PublicLock.CallOpts)
}

// GasRefundValue is a free data retrieval call binding the contract method 0x6207a8da.
//
// Solidity: function gasRefundValue() view returns(uint256 _refundValue)
func (_PublicLock *PublicLockCaller) GasRefundValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "gasRefundValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasRefundValue is a free data retrieval call binding the contract method 0x6207a8da.
//
// Solidity: function gasRefundValue() view returns(uint256 _refundValue)
func (_PublicLock *PublicLockSession) GasRefundValue() (*big.Int, error) {
	return _PublicLock.Contract.GasRefundValue(&_PublicLock.CallOpts)
}

// GasRefundValue is a free data retrieval call binding the contract method 0x6207a8da.
//
// Solidity: function gasRefundValue() view returns(uint256 _refundValue)
func (_PublicLock *PublicLockCallerSession) GasRefundValue() (*big.Int, error) {
	return _PublicLock.Contract.GasRefundValue(&_PublicLock.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address)
func (_PublicLock *PublicLockCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "getApproved", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address)
func (_PublicLock *PublicLockSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _PublicLock.Contract.GetApproved(&_PublicLock.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address)
func (_PublicLock *PublicLockCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _PublicLock.Contract.GetApproved(&_PublicLock.CallOpts, _tokenId)
}

// GetCancelAndRefundValue is a free data retrieval call binding the contract method 0x92ac98a5.
//
// Solidity: function getCancelAndRefundValue(uint256 _tokenId) view returns(uint256 refund)
func (_PublicLock *PublicLockCaller) GetCancelAndRefundValue(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "getCancelAndRefundValue", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCancelAndRefundValue is a free data retrieval call binding the contract method 0x92ac98a5.
//
// Solidity: function getCancelAndRefundValue(uint256 _tokenId) view returns(uint256 refund)
func (_PublicLock *PublicLockSession) GetCancelAndRefundValue(_tokenId *big.Int) (*big.Int, error) {
	return _PublicLock.Contract.GetCancelAndRefundValue(&_PublicLock.CallOpts, _tokenId)
}

// GetCancelAndRefundValue is a free data retrieval call binding the contract method 0x92ac98a5.
//
// Solidity: function getCancelAndRefundValue(uint256 _tokenId) view returns(uint256 refund)
func (_PublicLock *PublicLockCallerSession) GetCancelAndRefundValue(_tokenId *big.Int) (*big.Int, error) {
	return _PublicLock.Contract.GetCancelAndRefundValue(&_PublicLock.CallOpts, _tokenId)
}

// GetHasValidKey is a free data retrieval call binding the contract method 0x6d8ea5b4.
//
// Solidity: function getHasValidKey(address _keyOwner) view returns(bool isValid)
func (_PublicLock *PublicLockCaller) GetHasValidKey(opts *bind.CallOpts, _keyOwner common.Address) (bool, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "getHasValidKey", _keyOwner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetHasValidKey is a free data retrieval call binding the contract method 0x6d8ea5b4.
//
// Solidity: function getHasValidKey(address _keyOwner) view returns(bool isValid)
func (_PublicLock *PublicLockSession) GetHasValidKey(_keyOwner common.Address) (bool, error) {
	return _PublicLock.Contract.GetHasValidKey(&_PublicLock.CallOpts, _keyOwner)
}

// GetHasValidKey is a free data retrieval call binding the contract method 0x6d8ea5b4.
//
// Solidity: function getHasValidKey(address _keyOwner) view returns(bool isValid)
func (_PublicLock *PublicLockCallerSession) GetHasValidKey(_keyOwner common.Address) (bool, error) {
	return _PublicLock.Contract.GetHasValidKey(&_PublicLock.CallOpts, _keyOwner)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PublicLock *PublicLockCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PublicLock *PublicLockSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PublicLock.Contract.GetRoleAdmin(&_PublicLock.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PublicLock *PublicLockCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PublicLock.Contract.GetRoleAdmin(&_PublicLock.CallOpts, role)
}

// GetTransferFee is a free data retrieval call binding the contract method 0xb1a3b25d.
//
// Solidity: function getTransferFee(uint256 _tokenId, uint256 _time) view returns(uint256)
func (_PublicLock *PublicLockCaller) GetTransferFee(opts *bind.CallOpts, _tokenId *big.Int, _time *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "getTransferFee", _tokenId, _time)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransferFee is a free data retrieval call binding the contract method 0xb1a3b25d.
//
// Solidity: function getTransferFee(uint256 _tokenId, uint256 _time) view returns(uint256)
func (_PublicLock *PublicLockSession) GetTransferFee(_tokenId *big.Int, _time *big.Int) (*big.Int, error) {
	return _PublicLock.Contract.GetTransferFee(&_PublicLock.CallOpts, _tokenId, _time)
}

// GetTransferFee is a free data retrieval call binding the contract method 0xb1a3b25d.
//
// Solidity: function getTransferFee(uint256 _tokenId, uint256 _time) view returns(uint256)
func (_PublicLock *PublicLockCallerSession) GetTransferFee(_tokenId *big.Int, _time *big.Int) (*big.Int, error) {
	return _PublicLock.Contract.GetTransferFee(&_PublicLock.CallOpts, _tokenId, _time)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PublicLock *PublicLockCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PublicLock *PublicLockSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PublicLock.Contract.HasRole(&_PublicLock.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PublicLock *PublicLockCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PublicLock.Contract.HasRole(&_PublicLock.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_PublicLock *PublicLockCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "isApprovedForAll", _owner, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_PublicLock *PublicLockSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _PublicLock.Contract.IsApprovedForAll(&_PublicLock.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_PublicLock *PublicLockCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _PublicLock.Contract.IsApprovedForAll(&_PublicLock.CallOpts, _owner, _operator)
}

// IsKeyGranter is a free data retrieval call binding the contract method 0x52b0f638.
//
// Solidity: function isKeyGranter(address account) view returns(bool)
func (_PublicLock *PublicLockCaller) IsKeyGranter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "isKeyGranter", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKeyGranter is a free data retrieval call binding the contract method 0x52b0f638.
//
// Solidity: function isKeyGranter(address account) view returns(bool)
func (_PublicLock *PublicLockSession) IsKeyGranter(account common.Address) (bool, error) {
	return _PublicLock.Contract.IsKeyGranter(&_PublicLock.CallOpts, account)
}

// IsKeyGranter is a free data retrieval call binding the contract method 0x52b0f638.
//
// Solidity: function isKeyGranter(address account) view returns(bool)
func (_PublicLock *PublicLockCallerSession) IsKeyGranter(account common.Address) (bool, error) {
	return _PublicLock.Contract.IsKeyGranter(&_PublicLock.CallOpts, account)
}

// IsLockManager is a free data retrieval call binding the contract method 0xaae4b8f7.
//
// Solidity: function isLockManager(address account) view returns(bool)
func (_PublicLock *PublicLockCaller) IsLockManager(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "isLockManager", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLockManager is a free data retrieval call binding the contract method 0xaae4b8f7.
//
// Solidity: function isLockManager(address account) view returns(bool)
func (_PublicLock *PublicLockSession) IsLockManager(account common.Address) (bool, error) {
	return _PublicLock.Contract.IsLockManager(&_PublicLock.CallOpts, account)
}

// IsLockManager is a free data retrieval call binding the contract method 0xaae4b8f7.
//
// Solidity: function isLockManager(address account) view returns(bool)
func (_PublicLock *PublicLockCallerSession) IsLockManager(account common.Address) (bool, error) {
	return _PublicLock.Contract.IsLockManager(&_PublicLock.CallOpts, account)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_PublicLock *PublicLockCaller) IsOwner(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "isOwner", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_PublicLock *PublicLockSession) IsOwner(account common.Address) (bool, error) {
	return _PublicLock.Contract.IsOwner(&_PublicLock.CallOpts, account)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_PublicLock *PublicLockCallerSession) IsOwner(account common.Address) (bool, error) {
	return _PublicLock.Contract.IsOwner(&_PublicLock.CallOpts, account)
}

// IsValidKey is a free data retrieval call binding the contract method 0xa98d3623.
//
// Solidity: function isValidKey(uint256 _tokenId) view returns(bool)
func (_PublicLock *PublicLockCaller) IsValidKey(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "isValidKey", _tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidKey is a free data retrieval call binding the contract method 0xa98d3623.
//
// Solidity: function isValidKey(uint256 _tokenId) view returns(bool)
func (_PublicLock *PublicLockSession) IsValidKey(_tokenId *big.Int) (bool, error) {
	return _PublicLock.Contract.IsValidKey(&_PublicLock.CallOpts, _tokenId)
}

// IsValidKey is a free data retrieval call binding the contract method 0xa98d3623.
//
// Solidity: function isValidKey(uint256 _tokenId) view returns(bool)
func (_PublicLock *PublicLockCallerSession) IsValidKey(_tokenId *big.Int) (bool, error) {
	return _PublicLock.Contract.IsValidKey(&_PublicLock.CallOpts, _tokenId)
}

// KeyExpirationTimestampFor is a free data retrieval call binding the contract method 0x54b249fb.
//
// Solidity: function keyExpirationTimestampFor(uint256 _tokenId) view returns(uint256)
func (_PublicLock *PublicLockCaller) KeyExpirationTimestampFor(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "keyExpirationTimestampFor", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KeyExpirationTimestampFor is a free data retrieval call binding the contract method 0x54b249fb.
//
// Solidity: function keyExpirationTimestampFor(uint256 _tokenId) view returns(uint256)
func (_PublicLock *PublicLockSession) KeyExpirationTimestampFor(_tokenId *big.Int) (*big.Int, error) {
	return _PublicLock.Contract.KeyExpirationTimestampFor(&_PublicLock.CallOpts, _tokenId)
}

// KeyExpirationTimestampFor is a free data retrieval call binding the contract method 0x54b249fb.
//
// Solidity: function keyExpirationTimestampFor(uint256 _tokenId) view returns(uint256)
func (_PublicLock *PublicLockCallerSession) KeyExpirationTimestampFor(_tokenId *big.Int) (*big.Int, error) {
	return _PublicLock.Contract.KeyExpirationTimestampFor(&_PublicLock.CallOpts, _tokenId)
}

// KeyManagerOf is a free data retrieval call binding the contract method 0x4d025fed.
//
// Solidity: function keyManagerOf(uint256 ) view returns(address)
func (_PublicLock *PublicLockCaller) KeyManagerOf(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "keyManagerOf", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyManagerOf is a free data retrieval call binding the contract method 0x4d025fed.
//
// Solidity: function keyManagerOf(uint256 ) view returns(address)
func (_PublicLock *PublicLockSession) KeyManagerOf(arg0 *big.Int) (common.Address, error) {
	return _PublicLock.Contract.KeyManagerOf(&_PublicLock.CallOpts, arg0)
}

// KeyManagerOf is a free data retrieval call binding the contract method 0x4d025fed.
//
// Solidity: function keyManagerOf(uint256 ) view returns(address)
func (_PublicLock *PublicLockCallerSession) KeyManagerOf(arg0 *big.Int) (common.Address, error) {
	return _PublicLock.Contract.KeyManagerOf(&_PublicLock.CallOpts, arg0)
}

// KeyPrice is a free data retrieval call binding the contract method 0x10e56973.
//
// Solidity: function keyPrice() view returns(uint256)
func (_PublicLock *PublicLockCaller) KeyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "keyPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KeyPrice is a free data retrieval call binding the contract method 0x10e56973.
//
// Solidity: function keyPrice() view returns(uint256)
func (_PublicLock *PublicLockSession) KeyPrice() (*big.Int, error) {
	return _PublicLock.Contract.KeyPrice(&_PublicLock.CallOpts)
}

// KeyPrice is a free data retrieval call binding the contract method 0x10e56973.
//
// Solidity: function keyPrice() view returns(uint256)
func (_PublicLock *PublicLockCallerSession) KeyPrice() (*big.Int, error) {
	return _PublicLock.Contract.KeyPrice(&_PublicLock.CallOpts)
}

// MaxKeysPerAddress is a free data retrieval call binding the contract method 0xd52e4a10.
//
// Solidity: function maxKeysPerAddress() view returns(uint256)
func (_PublicLock *PublicLockCaller) MaxKeysPerAddress(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "maxKeysPerAddress")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxKeysPerAddress is a free data retrieval call binding the contract method 0xd52e4a10.
//
// Solidity: function maxKeysPerAddress() view returns(uint256)
func (_PublicLock *PublicLockSession) MaxKeysPerAddress() (*big.Int, error) {
	return _PublicLock.Contract.MaxKeysPerAddress(&_PublicLock.CallOpts)
}

// MaxKeysPerAddress is a free data retrieval call binding the contract method 0xd52e4a10.
//
// Solidity: function maxKeysPerAddress() view returns(uint256)
func (_PublicLock *PublicLockCallerSession) MaxKeysPerAddress() (*big.Int, error) {
	return _PublicLock.Contract.MaxKeysPerAddress(&_PublicLock.CallOpts)
}

// MaxNumberOfKeys is a free data retrieval call binding the contract method 0x74b6c106.
//
// Solidity: function maxNumberOfKeys() view returns(uint256)
func (_PublicLock *PublicLockCaller) MaxNumberOfKeys(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "maxNumberOfKeys")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxNumberOfKeys is a free data retrieval call binding the contract method 0x74b6c106.
//
// Solidity: function maxNumberOfKeys() view returns(uint256)
func (_PublicLock *PublicLockSession) MaxNumberOfKeys() (*big.Int, error) {
	return _PublicLock.Contract.MaxNumberOfKeys(&_PublicLock.CallOpts)
}

// MaxNumberOfKeys is a free data retrieval call binding the contract method 0x74b6c106.
//
// Solidity: function maxNumberOfKeys() view returns(uint256)
func (_PublicLock *PublicLockCallerSession) MaxNumberOfKeys() (*big.Int, error) {
	return _PublicLock.Contract.MaxNumberOfKeys(&_PublicLock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PublicLock *PublicLockCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PublicLock *PublicLockSession) Name() (string, error) {
	return _PublicLock.Contract.Name(&_PublicLock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PublicLock *PublicLockCallerSession) Name() (string, error) {
	return _PublicLock.Contract.Name(&_PublicLock.CallOpts)
}

// NumberOfOwners is a free data retrieval call binding the contract method 0x93fd1844.
//
// Solidity: function numberOfOwners() view returns(uint256)
func (_PublicLock *PublicLockCaller) NumberOfOwners(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "numberOfOwners")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfOwners is a free data retrieval call binding the contract method 0x93fd1844.
//
// Solidity: function numberOfOwners() view returns(uint256)
func (_PublicLock *PublicLockSession) NumberOfOwners() (*big.Int, error) {
	return _PublicLock.Contract.NumberOfOwners(&_PublicLock.CallOpts)
}

// NumberOfOwners is a free data retrieval call binding the contract method 0x93fd1844.
//
// Solidity: function numberOfOwners() view returns(uint256)
func (_PublicLock *PublicLockCallerSession) NumberOfOwners() (*big.Int, error) {
	return _PublicLock.Contract.NumberOfOwners(&_PublicLock.CallOpts)
}

// OnKeyCancelHook is a free data retrieval call binding the contract method 0x217751bc.
//
// Solidity: function onKeyCancelHook() view returns(address)
func (_PublicLock *PublicLockCaller) OnKeyCancelHook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "onKeyCancelHook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OnKeyCancelHook is a free data retrieval call binding the contract method 0x217751bc.
//
// Solidity: function onKeyCancelHook() view returns(address)
func (_PublicLock *PublicLockSession) OnKeyCancelHook() (common.Address, error) {
	return _PublicLock.Contract.OnKeyCancelHook(&_PublicLock.CallOpts)
}

// OnKeyCancelHook is a free data retrieval call binding the contract method 0x217751bc.
//
// Solidity: function onKeyCancelHook() view returns(address)
func (_PublicLock *PublicLockCallerSession) OnKeyCancelHook() (common.Address, error) {
	return _PublicLock.Contract.OnKeyCancelHook(&_PublicLock.CallOpts)
}

// OnKeyExtendHook is a free data retrieval call binding the contract method 0xc907c3ec.
//
// Solidity: function onKeyExtendHook() view returns(address)
func (_PublicLock *PublicLockCaller) OnKeyExtendHook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "onKeyExtendHook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OnKeyExtendHook is a free data retrieval call binding the contract method 0xc907c3ec.
//
// Solidity: function onKeyExtendHook() view returns(address)
func (_PublicLock *PublicLockSession) OnKeyExtendHook() (common.Address, error) {
	return _PublicLock.Contract.OnKeyExtendHook(&_PublicLock.CallOpts)
}

// OnKeyExtendHook is a free data retrieval call binding the contract method 0xc907c3ec.
//
// Solidity: function onKeyExtendHook() view returns(address)
func (_PublicLock *PublicLockCallerSession) OnKeyExtendHook() (common.Address, error) {
	return _PublicLock.Contract.OnKeyExtendHook(&_PublicLock.CallOpts)
}

// OnKeyGrantHook is a free data retrieval call binding the contract method 0xb129694e.
//
// Solidity: function onKeyGrantHook() view returns(address)
func (_PublicLock *PublicLockCaller) OnKeyGrantHook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "onKeyGrantHook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OnKeyGrantHook is a free data retrieval call binding the contract method 0xb129694e.
//
// Solidity: function onKeyGrantHook() view returns(address)
func (_PublicLock *PublicLockSession) OnKeyGrantHook() (common.Address, error) {
	return _PublicLock.Contract.OnKeyGrantHook(&_PublicLock.CallOpts)
}

// OnKeyGrantHook is a free data retrieval call binding the contract method 0xb129694e.
//
// Solidity: function onKeyGrantHook() view returns(address)
func (_PublicLock *PublicLockCallerSession) OnKeyGrantHook() (common.Address, error) {
	return _PublicLock.Contract.OnKeyGrantHook(&_PublicLock.CallOpts)
}

// OnKeyPurchaseHook is a free data retrieval call binding the contract method 0x2d33dd5b.
//
// Solidity: function onKeyPurchaseHook() view returns(address)
func (_PublicLock *PublicLockCaller) OnKeyPurchaseHook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "onKeyPurchaseHook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OnKeyPurchaseHook is a free data retrieval call binding the contract method 0x2d33dd5b.
//
// Solidity: function onKeyPurchaseHook() view returns(address)
func (_PublicLock *PublicLockSession) OnKeyPurchaseHook() (common.Address, error) {
	return _PublicLock.Contract.OnKeyPurchaseHook(&_PublicLock.CallOpts)
}

// OnKeyPurchaseHook is a free data retrieval call binding the contract method 0x2d33dd5b.
//
// Solidity: function onKeyPurchaseHook() view returns(address)
func (_PublicLock *PublicLockCallerSession) OnKeyPurchaseHook() (common.Address, error) {
	return _PublicLock.Contract.OnKeyPurchaseHook(&_PublicLock.CallOpts)
}

// OnKeyTransferHook is a free data retrieval call binding the contract method 0x389f07e8.
//
// Solidity: function onKeyTransferHook() view returns(address)
func (_PublicLock *PublicLockCaller) OnKeyTransferHook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "onKeyTransferHook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OnKeyTransferHook is a free data retrieval call binding the contract method 0x389f07e8.
//
// Solidity: function onKeyTransferHook() view returns(address)
func (_PublicLock *PublicLockSession) OnKeyTransferHook() (common.Address, error) {
	return _PublicLock.Contract.OnKeyTransferHook(&_PublicLock.CallOpts)
}

// OnKeyTransferHook is a free data retrieval call binding the contract method 0x389f07e8.
//
// Solidity: function onKeyTransferHook() view returns(address)
func (_PublicLock *PublicLockCallerSession) OnKeyTransferHook() (common.Address, error) {
	return _PublicLock.Contract.OnKeyTransferHook(&_PublicLock.CallOpts)
}

// OnTokenURIHook is a free data retrieval call binding the contract method 0x7ec2a724.
//
// Solidity: function onTokenURIHook() view returns(address)
func (_PublicLock *PublicLockCaller) OnTokenURIHook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "onTokenURIHook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OnTokenURIHook is a free data retrieval call binding the contract method 0x7ec2a724.
//
// Solidity: function onTokenURIHook() view returns(address)
func (_PublicLock *PublicLockSession) OnTokenURIHook() (common.Address, error) {
	return _PublicLock.Contract.OnTokenURIHook(&_PublicLock.CallOpts)
}

// OnTokenURIHook is a free data retrieval call binding the contract method 0x7ec2a724.
//
// Solidity: function onTokenURIHook() view returns(address)
func (_PublicLock *PublicLockCallerSession) OnTokenURIHook() (common.Address, error) {
	return _PublicLock.Contract.OnTokenURIHook(&_PublicLock.CallOpts)
}

// OnValidKeyHook is a free data retrieval call binding the contract method 0x26e9ca07.
//
// Solidity: function onValidKeyHook() view returns(address)
func (_PublicLock *PublicLockCaller) OnValidKeyHook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "onValidKeyHook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OnValidKeyHook is a free data retrieval call binding the contract method 0x26e9ca07.
//
// Solidity: function onValidKeyHook() view returns(address)
func (_PublicLock *PublicLockSession) OnValidKeyHook() (common.Address, error) {
	return _PublicLock.Contract.OnValidKeyHook(&_PublicLock.CallOpts)
}

// OnValidKeyHook is a free data retrieval call binding the contract method 0x26e9ca07.
//
// Solidity: function onValidKeyHook() view returns(address)
func (_PublicLock *PublicLockCallerSession) OnValidKeyHook() (common.Address, error) {
	return _PublicLock.Contract.OnValidKeyHook(&_PublicLock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PublicLock *PublicLockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PublicLock *PublicLockSession) Owner() (common.Address, error) {
	return _PublicLock.Contract.Owner(&_PublicLock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PublicLock *PublicLockCallerSession) Owner() (common.Address, error) {
	return _PublicLock.Contract.Owner(&_PublicLock.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address)
func (_PublicLock *PublicLockCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "ownerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address)
func (_PublicLock *PublicLockSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _PublicLock.Contract.OwnerOf(&_PublicLock.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address)
func (_PublicLock *PublicLockCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _PublicLock.Contract.OwnerOf(&_PublicLock.CallOpts, _tokenId)
}

// PublicLockVersion is a free data retrieval call binding the contract method 0xd1bbd49c.
//
// Solidity: function publicLockVersion() pure returns(uint16)
func (_PublicLock *PublicLockCaller) PublicLockVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "publicLockVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// PublicLockVersion is a free data retrieval call binding the contract method 0xd1bbd49c.
//
// Solidity: function publicLockVersion() pure returns(uint16)
func (_PublicLock *PublicLockSession) PublicLockVersion() (uint16, error) {
	return _PublicLock.Contract.PublicLockVersion(&_PublicLock.CallOpts)
}

// PublicLockVersion is a free data retrieval call binding the contract method 0xd1bbd49c.
//
// Solidity: function publicLockVersion() pure returns(uint16)
func (_PublicLock *PublicLockCallerSession) PublicLockVersion() (uint16, error) {
	return _PublicLock.Contract.PublicLockVersion(&_PublicLock.CallOpts)
}

// PurchasePriceFor is a free data retrieval call binding the contract method 0x097ba333.
//
// Solidity: function purchasePriceFor(address _recipient, address _referrer, bytes _data) view returns(uint256 minKeyPrice)
func (_PublicLock *PublicLockCaller) PurchasePriceFor(opts *bind.CallOpts, _recipient common.Address, _referrer common.Address, _data []byte) (*big.Int, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "purchasePriceFor", _recipient, _referrer, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PurchasePriceFor is a free data retrieval call binding the contract method 0x097ba333.
//
// Solidity: function purchasePriceFor(address _recipient, address _referrer, bytes _data) view returns(uint256 minKeyPrice)
func (_PublicLock *PublicLockSession) PurchasePriceFor(_recipient common.Address, _referrer common.Address, _data []byte) (*big.Int, error) {
	return _PublicLock.Contract.PurchasePriceFor(&_PublicLock.CallOpts, _recipient, _referrer, _data)
}

// PurchasePriceFor is a free data retrieval call binding the contract method 0x097ba333.
//
// Solidity: function purchasePriceFor(address _recipient, address _referrer, bytes _data) view returns(uint256 minKeyPrice)
func (_PublicLock *PublicLockCallerSession) PurchasePriceFor(_recipient common.Address, _referrer common.Address, _data []byte) (*big.Int, error) {
	return _PublicLock.Contract.PurchasePriceFor(&_PublicLock.CallOpts, _recipient, _referrer, _data)
}

// ReferrerFees is a free data retrieval call binding the contract method 0xc23135dd.
//
// Solidity: function referrerFees(address ) view returns(uint256)
func (_PublicLock *PublicLockCaller) ReferrerFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "referrerFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReferrerFees is a free data retrieval call binding the contract method 0xc23135dd.
//
// Solidity: function referrerFees(address ) view returns(uint256)
func (_PublicLock *PublicLockSession) ReferrerFees(arg0 common.Address) (*big.Int, error) {
	return _PublicLock.Contract.ReferrerFees(&_PublicLock.CallOpts, arg0)
}

// ReferrerFees is a free data retrieval call binding the contract method 0xc23135dd.
//
// Solidity: function referrerFees(address ) view returns(uint256)
func (_PublicLock *PublicLockCallerSession) ReferrerFees(arg0 common.Address) (*big.Int, error) {
	return _PublicLock.Contract.ReferrerFees(&_PublicLock.CallOpts, arg0)
}

// RefundPenaltyBasisPoints is a free data retrieval call binding the contract method 0x56e0d51f.
//
// Solidity: function refundPenaltyBasisPoints() view returns(uint256)
func (_PublicLock *PublicLockCaller) RefundPenaltyBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "refundPenaltyBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RefundPenaltyBasisPoints is a free data retrieval call binding the contract method 0x56e0d51f.
//
// Solidity: function refundPenaltyBasisPoints() view returns(uint256)
func (_PublicLock *PublicLockSession) RefundPenaltyBasisPoints() (*big.Int, error) {
	return _PublicLock.Contract.RefundPenaltyBasisPoints(&_PublicLock.CallOpts)
}

// RefundPenaltyBasisPoints is a free data retrieval call binding the contract method 0x56e0d51f.
//
// Solidity: function refundPenaltyBasisPoints() view returns(uint256)
func (_PublicLock *PublicLockCallerSession) RefundPenaltyBasisPoints() (*big.Int, error) {
	return _PublicLock.Contract.RefundPenaltyBasisPoints(&_PublicLock.CallOpts)
}

// SchemaVersion is a free data retrieval call binding the contract method 0x4e2ce6d3.
//
// Solidity: function schemaVersion() view returns(uint256)
func (_PublicLock *PublicLockCaller) SchemaVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "schemaVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SchemaVersion is a free data retrieval call binding the contract method 0x4e2ce6d3.
//
// Solidity: function schemaVersion() view returns(uint256)
func (_PublicLock *PublicLockSession) SchemaVersion() (*big.Int, error) {
	return _PublicLock.Contract.SchemaVersion(&_PublicLock.CallOpts)
}

// SchemaVersion is a free data retrieval call binding the contract method 0x4e2ce6d3.
//
// Solidity: function schemaVersion() view returns(uint256)
func (_PublicLock *PublicLockCallerSession) SchemaVersion() (*big.Int, error) {
	return _PublicLock.Contract.SchemaVersion(&_PublicLock.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PublicLock *PublicLockCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PublicLock *PublicLockSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PublicLock.Contract.SupportsInterface(&_PublicLock.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PublicLock *PublicLockCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PublicLock.Contract.SupportsInterface(&_PublicLock.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PublicLock *PublicLockCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PublicLock *PublicLockSession) Symbol() (string, error) {
	return _PublicLock.Contract.Symbol(&_PublicLock.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PublicLock *PublicLockCallerSession) Symbol() (string, error) {
	return _PublicLock.Contract.Symbol(&_PublicLock.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_PublicLock *PublicLockCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "tokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_PublicLock *PublicLockSession) TokenAddress() (common.Address, error) {
	return _PublicLock.Contract.TokenAddress(&_PublicLock.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_PublicLock *PublicLockCallerSession) TokenAddress() (common.Address, error) {
	return _PublicLock.Contract.TokenAddress(&_PublicLock.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) view returns(uint256)
func (_PublicLock *PublicLockCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "tokenByIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) view returns(uint256)
func (_PublicLock *PublicLockSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _PublicLock.Contract.TokenByIndex(&_PublicLock.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) view returns(uint256)
func (_PublicLock *PublicLockCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _PublicLock.Contract.TokenByIndex(&_PublicLock.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _keyOwner, uint256 _index) view returns(uint256)
func (_PublicLock *PublicLockCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _keyOwner common.Address, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "tokenOfOwnerByIndex", _keyOwner, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _keyOwner, uint256 _index) view returns(uint256)
func (_PublicLock *PublicLockSession) TokenOfOwnerByIndex(_keyOwner common.Address, _index *big.Int) (*big.Int, error) {
	return _PublicLock.Contract.TokenOfOwnerByIndex(&_PublicLock.CallOpts, _keyOwner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _keyOwner, uint256 _index) view returns(uint256)
func (_PublicLock *PublicLockCallerSession) TokenOfOwnerByIndex(_keyOwner common.Address, _index *big.Int) (*big.Int, error) {
	return _PublicLock.Contract.TokenOfOwnerByIndex(&_PublicLock.CallOpts, _keyOwner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_PublicLock *PublicLockCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_PublicLock *PublicLockSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _PublicLock.Contract.TokenURI(&_PublicLock.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_PublicLock *PublicLockCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _PublicLock.Contract.TokenURI(&_PublicLock.CallOpts, _tokenId)
}

// TotalKeys is a free data retrieval call binding the contract method 0x812eecd4.
//
// Solidity: function totalKeys(address _keyOwner) view returns(uint256)
func (_PublicLock *PublicLockCaller) TotalKeys(opts *bind.CallOpts, _keyOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "totalKeys", _keyOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalKeys is a free data retrieval call binding the contract method 0x812eecd4.
//
// Solidity: function totalKeys(address _keyOwner) view returns(uint256)
func (_PublicLock *PublicLockSession) TotalKeys(_keyOwner common.Address) (*big.Int, error) {
	return _PublicLock.Contract.TotalKeys(&_PublicLock.CallOpts, _keyOwner)
}

// TotalKeys is a free data retrieval call binding the contract method 0x812eecd4.
//
// Solidity: function totalKeys(address _keyOwner) view returns(uint256)
func (_PublicLock *PublicLockCallerSession) TotalKeys(_keyOwner common.Address) (*big.Int, error) {
	return _PublicLock.Contract.TotalKeys(&_PublicLock.CallOpts, _keyOwner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PublicLock *PublicLockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PublicLock *PublicLockSession) TotalSupply() (*big.Int, error) {
	return _PublicLock.Contract.TotalSupply(&_PublicLock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PublicLock *PublicLockCallerSession) TotalSupply() (*big.Int, error) {
	return _PublicLock.Contract.TotalSupply(&_PublicLock.CallOpts)
}

// TransferFeeBasisPoints is a free data retrieval call binding the contract method 0x183767da.
//
// Solidity: function transferFeeBasisPoints() view returns(uint256)
func (_PublicLock *PublicLockCaller) TransferFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "transferFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferFeeBasisPoints is a free data retrieval call binding the contract method 0x183767da.
//
// Solidity: function transferFeeBasisPoints() view returns(uint256)
func (_PublicLock *PublicLockSession) TransferFeeBasisPoints() (*big.Int, error) {
	return _PublicLock.Contract.TransferFeeBasisPoints(&_PublicLock.CallOpts)
}

// TransferFeeBasisPoints is a free data retrieval call binding the contract method 0x183767da.
//
// Solidity: function transferFeeBasisPoints() view returns(uint256)
func (_PublicLock *PublicLockCallerSession) TransferFeeBasisPoints() (*big.Int, error) {
	return _PublicLock.Contract.TransferFeeBasisPoints(&_PublicLock.CallOpts)
}

// UnlockProtocol is a free data retrieval call binding the contract method 0x0f15023b.
//
// Solidity: function unlockProtocol() view returns(address)
func (_PublicLock *PublicLockCaller) UnlockProtocol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PublicLock.contract.Call(opts, &out, "unlockProtocol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnlockProtocol is a free data retrieval call binding the contract method 0x0f15023b.
//
// Solidity: function unlockProtocol() view returns(address)
func (_PublicLock *PublicLockSession) UnlockProtocol() (common.Address, error) {
	return _PublicLock.Contract.UnlockProtocol(&_PublicLock.CallOpts)
}

// UnlockProtocol is a free data retrieval call binding the contract method 0x0f15023b.
//
// Solidity: function unlockProtocol() view returns(address)
func (_PublicLock *PublicLockCallerSession) UnlockProtocol() (common.Address, error) {
	return _PublicLock.Contract.UnlockProtocol(&_PublicLock.CallOpts)
}

// AddKeyGranter is a paid mutator transaction binding the contract method 0x564aa99d.
//
// Solidity: function addKeyGranter(address account) returns()
func (_PublicLock *PublicLockTransactor) AddKeyGranter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "addKeyGranter", account)
}

// AddKeyGranter is a paid mutator transaction binding the contract method 0x564aa99d.
//
// Solidity: function addKeyGranter(address account) returns()
func (_PublicLock *PublicLockSession) AddKeyGranter(account common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.AddKeyGranter(&_PublicLock.TransactOpts, account)
}

// AddKeyGranter is a paid mutator transaction binding the contract method 0x564aa99d.
//
// Solidity: function addKeyGranter(address account) returns()
func (_PublicLock *PublicLockTransactorSession) AddKeyGranter(account common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.AddKeyGranter(&_PublicLock.TransactOpts, account)
}

// AddLockManager is a paid mutator transaction binding the contract method 0xd2503485.
//
// Solidity: function addLockManager(address account) returns()
func (_PublicLock *PublicLockTransactor) AddLockManager(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "addLockManager", account)
}

// AddLockManager is a paid mutator transaction binding the contract method 0xd2503485.
//
// Solidity: function addLockManager(address account) returns()
func (_PublicLock *PublicLockSession) AddLockManager(account common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.AddLockManager(&_PublicLock.TransactOpts, account)
}

// AddLockManager is a paid mutator transaction binding the contract method 0xd2503485.
//
// Solidity: function addLockManager(address account) returns()
func (_PublicLock *PublicLockTransactorSession) AddLockManager(account common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.AddLockManager(&_PublicLock.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _approved, uint256 _tokenId) returns()
func (_PublicLock *PublicLockTransactor) Approve(opts *bind.TransactOpts, _approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "approve", _approved, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _approved, uint256 _tokenId) returns()
func (_PublicLock *PublicLockSession) Approve(_approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.Approve(&_PublicLock.TransactOpts, _approved, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _approved, uint256 _tokenId) returns()
func (_PublicLock *PublicLockTransactorSession) Approve(_approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.Approve(&_PublicLock.TransactOpts, _approved, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _tokenId) returns()
func (_PublicLock *PublicLockTransactor) Burn(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "burn", _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _tokenId) returns()
func (_PublicLock *PublicLockSession) Burn(_tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.Burn(&_PublicLock.TransactOpts, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _tokenId) returns()
func (_PublicLock *PublicLockTransactorSession) Burn(_tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.Burn(&_PublicLock.TransactOpts, _tokenId)
}

// CancelAndRefund is a paid mutator transaction binding the contract method 0xd32bfb6c.
//
// Solidity: function cancelAndRefund(uint256 _tokenId) returns()
func (_PublicLock *PublicLockTransactor) CancelAndRefund(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "cancelAndRefund", _tokenId)
}

// CancelAndRefund is a paid mutator transaction binding the contract method 0xd32bfb6c.
//
// Solidity: function cancelAndRefund(uint256 _tokenId) returns()
func (_PublicLock *PublicLockSession) CancelAndRefund(_tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.CancelAndRefund(&_PublicLock.TransactOpts, _tokenId)
}

// CancelAndRefund is a paid mutator transaction binding the contract method 0xd32bfb6c.
//
// Solidity: function cancelAndRefund(uint256 _tokenId) returns()
func (_PublicLock *PublicLockTransactorSession) CancelAndRefund(_tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.CancelAndRefund(&_PublicLock.TransactOpts, _tokenId)
}

// ExpireAndRefundFor is a paid mutator transaction binding the contract method 0x558b71e9.
//
// Solidity: function expireAndRefundFor(uint256 _tokenId, uint256 _amount) returns()
func (_PublicLock *PublicLockTransactor) ExpireAndRefundFor(opts *bind.TransactOpts, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "expireAndRefundFor", _tokenId, _amount)
}

// ExpireAndRefundFor is a paid mutator transaction binding the contract method 0x558b71e9.
//
// Solidity: function expireAndRefundFor(uint256 _tokenId, uint256 _amount) returns()
func (_PublicLock *PublicLockSession) ExpireAndRefundFor(_tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.ExpireAndRefundFor(&_PublicLock.TransactOpts, _tokenId, _amount)
}

// ExpireAndRefundFor is a paid mutator transaction binding the contract method 0x558b71e9.
//
// Solidity: function expireAndRefundFor(uint256 _tokenId, uint256 _amount) returns()
func (_PublicLock *PublicLockTransactorSession) ExpireAndRefundFor(_tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.ExpireAndRefundFor(&_PublicLock.TransactOpts, _tokenId, _amount)
}

// Extend is a paid mutator transaction binding the contract method 0xd813cc19.
//
// Solidity: function extend(uint256 _value, uint256 _tokenId, address _referrer, bytes _data) payable returns()
func (_PublicLock *PublicLockTransactor) Extend(opts *bind.TransactOpts, _value *big.Int, _tokenId *big.Int, _referrer common.Address, _data []byte) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "extend", _value, _tokenId, _referrer, _data)
}

// Extend is a paid mutator transaction binding the contract method 0xd813cc19.
//
// Solidity: function extend(uint256 _value, uint256 _tokenId, address _referrer, bytes _data) payable returns()
func (_PublicLock *PublicLockSession) Extend(_value *big.Int, _tokenId *big.Int, _referrer common.Address, _data []byte) (*types.Transaction, error) {
	return _PublicLock.Contract.Extend(&_PublicLock.TransactOpts, _value, _tokenId, _referrer, _data)
}

// Extend is a paid mutator transaction binding the contract method 0xd813cc19.
//
// Solidity: function extend(uint256 _value, uint256 _tokenId, address _referrer, bytes _data) payable returns()
func (_PublicLock *PublicLockTransactorSession) Extend(_value *big.Int, _tokenId *big.Int, _referrer common.Address, _data []byte) (*types.Transaction, error) {
	return _PublicLock.Contract.Extend(&_PublicLock.TransactOpts, _value, _tokenId, _referrer, _data)
}

// GrantKeyExtension is a paid mutator transaction binding the contract method 0x4cd38c1d.
//
// Solidity: function grantKeyExtension(uint256 _tokenId, uint256 _duration) returns()
func (_PublicLock *PublicLockTransactor) GrantKeyExtension(opts *bind.TransactOpts, _tokenId *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "grantKeyExtension", _tokenId, _duration)
}

// GrantKeyExtension is a paid mutator transaction binding the contract method 0x4cd38c1d.
//
// Solidity: function grantKeyExtension(uint256 _tokenId, uint256 _duration) returns()
func (_PublicLock *PublicLockSession) GrantKeyExtension(_tokenId *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.GrantKeyExtension(&_PublicLock.TransactOpts, _tokenId, _duration)
}

// GrantKeyExtension is a paid mutator transaction binding the contract method 0x4cd38c1d.
//
// Solidity: function grantKeyExtension(uint256 _tokenId, uint256 _duration) returns()
func (_PublicLock *PublicLockTransactorSession) GrantKeyExtension(_tokenId *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.GrantKeyExtension(&_PublicLock.TransactOpts, _tokenId, _duration)
}

// GrantKeys is a paid mutator transaction binding the contract method 0x81a3c943.
//
// Solidity: function grantKeys(address[] _recipients, uint256[] _expirationTimestamps, address[] _keyManagers) returns(uint256[])
func (_PublicLock *PublicLockTransactor) GrantKeys(opts *bind.TransactOpts, _recipients []common.Address, _expirationTimestamps []*big.Int, _keyManagers []common.Address) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "grantKeys", _recipients, _expirationTimestamps, _keyManagers)
}

// GrantKeys is a paid mutator transaction binding the contract method 0x81a3c943.
//
// Solidity: function grantKeys(address[] _recipients, uint256[] _expirationTimestamps, address[] _keyManagers) returns(uint256[])
func (_PublicLock *PublicLockSession) GrantKeys(_recipients []common.Address, _expirationTimestamps []*big.Int, _keyManagers []common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.GrantKeys(&_PublicLock.TransactOpts, _recipients, _expirationTimestamps, _keyManagers)
}

// GrantKeys is a paid mutator transaction binding the contract method 0x81a3c943.
//
// Solidity: function grantKeys(address[] _recipients, uint256[] _expirationTimestamps, address[] _keyManagers) returns(uint256[])
func (_PublicLock *PublicLockTransactorSession) GrantKeys(_recipients []common.Address, _expirationTimestamps []*big.Int, _keyManagers []common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.GrantKeys(&_PublicLock.TransactOpts, _recipients, _expirationTimestamps, _keyManagers)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PublicLock *PublicLockTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PublicLock *PublicLockSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.GrantRole(&_PublicLock.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PublicLock *PublicLockTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.GrantRole(&_PublicLock.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x6eadde43.
//
// Solidity: function initialize(address _lockCreator, uint256 _expirationDuration, address _tokenAddress, uint256 _keyPrice, uint256 _maxNumberOfKeys, string _lockName) returns()
func (_PublicLock *PublicLockTransactor) Initialize(opts *bind.TransactOpts, _lockCreator common.Address, _expirationDuration *big.Int, _tokenAddress common.Address, _keyPrice *big.Int, _maxNumberOfKeys *big.Int, _lockName string) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "initialize", _lockCreator, _expirationDuration, _tokenAddress, _keyPrice, _maxNumberOfKeys, _lockName)
}

// Initialize is a paid mutator transaction binding the contract method 0x6eadde43.
//
// Solidity: function initialize(address _lockCreator, uint256 _expirationDuration, address _tokenAddress, uint256 _keyPrice, uint256 _maxNumberOfKeys, string _lockName) returns()
func (_PublicLock *PublicLockSession) Initialize(_lockCreator common.Address, _expirationDuration *big.Int, _tokenAddress common.Address, _keyPrice *big.Int, _maxNumberOfKeys *big.Int, _lockName string) (*types.Transaction, error) {
	return _PublicLock.Contract.Initialize(&_PublicLock.TransactOpts, _lockCreator, _expirationDuration, _tokenAddress, _keyPrice, _maxNumberOfKeys, _lockName)
}

// Initialize is a paid mutator transaction binding the contract method 0x6eadde43.
//
// Solidity: function initialize(address _lockCreator, uint256 _expirationDuration, address _tokenAddress, uint256 _keyPrice, uint256 _maxNumberOfKeys, string _lockName) returns()
func (_PublicLock *PublicLockTransactorSession) Initialize(_lockCreator common.Address, _expirationDuration *big.Int, _tokenAddress common.Address, _keyPrice *big.Int, _maxNumberOfKeys *big.Int, _lockName string) (*types.Transaction, error) {
	return _PublicLock.Contract.Initialize(&_PublicLock.TransactOpts, _lockCreator, _expirationDuration, _tokenAddress, _keyPrice, _maxNumberOfKeys, _lockName)
}

// LendKey is a paid mutator transaction binding the contract method 0x0c2db8d1.
//
// Solidity: function lendKey(address _from, address _recipient, uint256 _tokenId) returns()
func (_PublicLock *PublicLockTransactor) LendKey(opts *bind.TransactOpts, _from common.Address, _recipient common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "lendKey", _from, _recipient, _tokenId)
}

// LendKey is a paid mutator transaction binding the contract method 0x0c2db8d1.
//
// Solidity: function lendKey(address _from, address _recipient, uint256 _tokenId) returns()
func (_PublicLock *PublicLockSession) LendKey(_from common.Address, _recipient common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.LendKey(&_PublicLock.TransactOpts, _from, _recipient, _tokenId)
}

// LendKey is a paid mutator transaction binding the contract method 0x0c2db8d1.
//
// Solidity: function lendKey(address _from, address _recipient, uint256 _tokenId) returns()
func (_PublicLock *PublicLockTransactorSession) LendKey(_from common.Address, _recipient common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.LendKey(&_PublicLock.TransactOpts, _from, _recipient, _tokenId)
}

// MergeKeys is a paid mutator transaction binding the contract method 0x068208cd.
//
// Solidity: function mergeKeys(uint256 _tokenIdFrom, uint256 _tokenIdTo, uint256 _amount) returns()
func (_PublicLock *PublicLockTransactor) MergeKeys(opts *bind.TransactOpts, _tokenIdFrom *big.Int, _tokenIdTo *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "mergeKeys", _tokenIdFrom, _tokenIdTo, _amount)
}

// MergeKeys is a paid mutator transaction binding the contract method 0x068208cd.
//
// Solidity: function mergeKeys(uint256 _tokenIdFrom, uint256 _tokenIdTo, uint256 _amount) returns()
func (_PublicLock *PublicLockSession) MergeKeys(_tokenIdFrom *big.Int, _tokenIdTo *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.MergeKeys(&_PublicLock.TransactOpts, _tokenIdFrom, _tokenIdTo, _amount)
}

// MergeKeys is a paid mutator transaction binding the contract method 0x068208cd.
//
// Solidity: function mergeKeys(uint256 _tokenIdFrom, uint256 _tokenIdTo, uint256 _amount) returns()
func (_PublicLock *PublicLockTransactorSession) MergeKeys(_tokenIdFrom *big.Int, _tokenIdTo *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.MergeKeys(&_PublicLock.TransactOpts, _tokenIdFrom, _tokenIdTo, _amount)
}

// Migrate is a paid mutator transaction binding the contract method 0x8932a90d.
//
// Solidity: function migrate(bytes ) returns()
func (_PublicLock *PublicLockTransactor) Migrate(opts *bind.TransactOpts, arg0 []byte) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "migrate", arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0x8932a90d.
//
// Solidity: function migrate(bytes ) returns()
func (_PublicLock *PublicLockSession) Migrate(arg0 []byte) (*types.Transaction, error) {
	return _PublicLock.Contract.Migrate(&_PublicLock.TransactOpts, arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0x8932a90d.
//
// Solidity: function migrate(bytes ) returns()
func (_PublicLock *PublicLockTransactorSession) Migrate(arg0 []byte) (*types.Transaction, error) {
	return _PublicLock.Contract.Migrate(&_PublicLock.TransactOpts, arg0)
}

// Purchase is a paid mutator transaction binding the contract method 0x33818997.
//
// Solidity: function purchase(uint256[] _values, address[] _recipients, address[] _referrers, address[] _keyManagers, bytes[] _data) payable returns(uint256[])
func (_PublicLock *PublicLockTransactor) Purchase(opts *bind.TransactOpts, _values []*big.Int, _recipients []common.Address, _referrers []common.Address, _keyManagers []common.Address, _data [][]byte) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "purchase", _values, _recipients, _referrers, _keyManagers, _data)
}

// Purchase is a paid mutator transaction binding the contract method 0x33818997.
//
// Solidity: function purchase(uint256[] _values, address[] _recipients, address[] _referrers, address[] _keyManagers, bytes[] _data) payable returns(uint256[])
func (_PublicLock *PublicLockSession) Purchase(_values []*big.Int, _recipients []common.Address, _referrers []common.Address, _keyManagers []common.Address, _data [][]byte) (*types.Transaction, error) {
	return _PublicLock.Contract.Purchase(&_PublicLock.TransactOpts, _values, _recipients, _referrers, _keyManagers, _data)
}

// Purchase is a paid mutator transaction binding the contract method 0x33818997.
//
// Solidity: function purchase(uint256[] _values, address[] _recipients, address[] _referrers, address[] _keyManagers, bytes[] _data) payable returns(uint256[])
func (_PublicLock *PublicLockTransactorSession) Purchase(_values []*big.Int, _recipients []common.Address, _referrers []common.Address, _keyManagers []common.Address, _data [][]byte) (*types.Transaction, error) {
	return _PublicLock.Contract.Purchase(&_PublicLock.TransactOpts, _values, _recipients, _referrers, _keyManagers, _data)
}

// RenewMembershipFor is a paid mutator transaction binding the contract method 0x8505fe95.
//
// Solidity: function renewMembershipFor(uint256 _tokenId, address _referrer) returns()
func (_PublicLock *PublicLockTransactor) RenewMembershipFor(opts *bind.TransactOpts, _tokenId *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "renewMembershipFor", _tokenId, _referrer)
}

// RenewMembershipFor is a paid mutator transaction binding the contract method 0x8505fe95.
//
// Solidity: function renewMembershipFor(uint256 _tokenId, address _referrer) returns()
func (_PublicLock *PublicLockSession) RenewMembershipFor(_tokenId *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.RenewMembershipFor(&_PublicLock.TransactOpts, _tokenId, _referrer)
}

// RenewMembershipFor is a paid mutator transaction binding the contract method 0x8505fe95.
//
// Solidity: function renewMembershipFor(uint256 _tokenId, address _referrer) returns()
func (_PublicLock *PublicLockTransactorSession) RenewMembershipFor(_tokenId *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.RenewMembershipFor(&_PublicLock.TransactOpts, _tokenId, _referrer)
}

// RenounceLockManager is a paid mutator transaction binding the contract method 0xf0ba6040.
//
// Solidity: function renounceLockManager() returns()
func (_PublicLock *PublicLockTransactor) RenounceLockManager(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "renounceLockManager")
}

// RenounceLockManager is a paid mutator transaction binding the contract method 0xf0ba6040.
//
// Solidity: function renounceLockManager() returns()
func (_PublicLock *PublicLockSession) RenounceLockManager() (*types.Transaction, error) {
	return _PublicLock.Contract.RenounceLockManager(&_PublicLock.TransactOpts)
}

// RenounceLockManager is a paid mutator transaction binding the contract method 0xf0ba6040.
//
// Solidity: function renounceLockManager() returns()
func (_PublicLock *PublicLockTransactorSession) RenounceLockManager() (*types.Transaction, error) {
	return _PublicLock.Contract.RenounceLockManager(&_PublicLock.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PublicLock *PublicLockTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PublicLock *PublicLockSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.RenounceRole(&_PublicLock.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PublicLock *PublicLockTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.RenounceRole(&_PublicLock.TransactOpts, role, account)
}

// RevokeKeyGranter is a paid mutator transaction binding the contract method 0x2af9162a.
//
// Solidity: function revokeKeyGranter(address _granter) returns()
func (_PublicLock *PublicLockTransactor) RevokeKeyGranter(opts *bind.TransactOpts, _granter common.Address) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "revokeKeyGranter", _granter)
}

// RevokeKeyGranter is a paid mutator transaction binding the contract method 0x2af9162a.
//
// Solidity: function revokeKeyGranter(address _granter) returns()
func (_PublicLock *PublicLockSession) RevokeKeyGranter(_granter common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.RevokeKeyGranter(&_PublicLock.TransactOpts, _granter)
}

// RevokeKeyGranter is a paid mutator transaction binding the contract method 0x2af9162a.
//
// Solidity: function revokeKeyGranter(address _granter) returns()
func (_PublicLock *PublicLockTransactorSession) RevokeKeyGranter(_granter common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.RevokeKeyGranter(&_PublicLock.TransactOpts, _granter)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PublicLock *PublicLockTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PublicLock *PublicLockSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.RevokeRole(&_PublicLock.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PublicLock *PublicLockTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.RevokeRole(&_PublicLock.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_PublicLock *PublicLockTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_PublicLock *PublicLockSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.SafeTransferFrom(&_PublicLock.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_PublicLock *PublicLockTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.SafeTransferFrom(&_PublicLock.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_PublicLock *PublicLockTransactor) SafeTransferFrom0(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "safeTransferFrom0", _from, _to, _tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_PublicLock *PublicLockSession) SafeTransferFrom0(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _PublicLock.Contract.SafeTransferFrom0(&_PublicLock.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_PublicLock *PublicLockTransactorSession) SafeTransferFrom0(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _PublicLock.Contract.SafeTransferFrom0(&_PublicLock.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _to, bool _approved) returns()
func (_PublicLock *PublicLockTransactor) SetApprovalForAll(opts *bind.TransactOpts, _to common.Address, _approved bool) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "setApprovalForAll", _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _to, bool _approved) returns()
func (_PublicLock *PublicLockSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _PublicLock.Contract.SetApprovalForAll(&_PublicLock.TransactOpts, _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _to, bool _approved) returns()
func (_PublicLock *PublicLockTransactorSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _PublicLock.Contract.SetApprovalForAll(&_PublicLock.TransactOpts, _to, _approved)
}

// SetEventHooks is a paid mutator transaction binding the contract method 0x74cac47d.
//
// Solidity: function setEventHooks(address _onKeyPurchaseHook, address _onKeyCancelHook, address _onValidKeyHook, address _onTokenURIHook, address _onKeyTransferHook, address _onKeyExtendHook, address _onKeyGrantHook) returns()
func (_PublicLock *PublicLockTransactor) SetEventHooks(opts *bind.TransactOpts, _onKeyPurchaseHook common.Address, _onKeyCancelHook common.Address, _onValidKeyHook common.Address, _onTokenURIHook common.Address, _onKeyTransferHook common.Address, _onKeyExtendHook common.Address, _onKeyGrantHook common.Address) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "setEventHooks", _onKeyPurchaseHook, _onKeyCancelHook, _onValidKeyHook, _onTokenURIHook, _onKeyTransferHook, _onKeyExtendHook, _onKeyGrantHook)
}

// SetEventHooks is a paid mutator transaction binding the contract method 0x74cac47d.
//
// Solidity: function setEventHooks(address _onKeyPurchaseHook, address _onKeyCancelHook, address _onValidKeyHook, address _onTokenURIHook, address _onKeyTransferHook, address _onKeyExtendHook, address _onKeyGrantHook) returns()
func (_PublicLock *PublicLockSession) SetEventHooks(_onKeyPurchaseHook common.Address, _onKeyCancelHook common.Address, _onValidKeyHook common.Address, _onTokenURIHook common.Address, _onKeyTransferHook common.Address, _onKeyExtendHook common.Address, _onKeyGrantHook common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.SetEventHooks(&_PublicLock.TransactOpts, _onKeyPurchaseHook, _onKeyCancelHook, _onValidKeyHook, _onTokenURIHook, _onKeyTransferHook, _onKeyExtendHook, _onKeyGrantHook)
}

// SetEventHooks is a paid mutator transaction binding the contract method 0x74cac47d.
//
// Solidity: function setEventHooks(address _onKeyPurchaseHook, address _onKeyCancelHook, address _onValidKeyHook, address _onTokenURIHook, address _onKeyTransferHook, address _onKeyExtendHook, address _onKeyGrantHook) returns()
func (_PublicLock *PublicLockTransactorSession) SetEventHooks(_onKeyPurchaseHook common.Address, _onKeyCancelHook common.Address, _onValidKeyHook common.Address, _onTokenURIHook common.Address, _onKeyTransferHook common.Address, _onKeyExtendHook common.Address, _onKeyGrantHook common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.SetEventHooks(&_PublicLock.TransactOpts, _onKeyPurchaseHook, _onKeyCancelHook, _onValidKeyHook, _onTokenURIHook, _onKeyTransferHook, _onKeyExtendHook, _onKeyGrantHook)
}

// SetGasRefundValue is a paid mutator transaction binding the contract method 0xf5766b39.
//
// Solidity: function setGasRefundValue(uint256 _refundValue) returns()
func (_PublicLock *PublicLockTransactor) SetGasRefundValue(opts *bind.TransactOpts, _refundValue *big.Int) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "setGasRefundValue", _refundValue)
}

// SetGasRefundValue is a paid mutator transaction binding the contract method 0xf5766b39.
//
// Solidity: function setGasRefundValue(uint256 _refundValue) returns()
func (_PublicLock *PublicLockSession) SetGasRefundValue(_refundValue *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.SetGasRefundValue(&_PublicLock.TransactOpts, _refundValue)
}

// SetGasRefundValue is a paid mutator transaction binding the contract method 0xf5766b39.
//
// Solidity: function setGasRefundValue(uint256 _refundValue) returns()
func (_PublicLock *PublicLockTransactorSession) SetGasRefundValue(_refundValue *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.SetGasRefundValue(&_PublicLock.TransactOpts, _refundValue)
}

// SetKeyManagerOf is a paid mutator transaction binding the contract method 0xb11d7ec1.
//
// Solidity: function setKeyManagerOf(uint256 _tokenId, address _keyManager) returns()
func (_PublicLock *PublicLockTransactor) SetKeyManagerOf(opts *bind.TransactOpts, _tokenId *big.Int, _keyManager common.Address) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "setKeyManagerOf", _tokenId, _keyManager)
}

// SetKeyManagerOf is a paid mutator transaction binding the contract method 0xb11d7ec1.
//
// Solidity: function setKeyManagerOf(uint256 _tokenId, address _keyManager) returns()
func (_PublicLock *PublicLockSession) SetKeyManagerOf(_tokenId *big.Int, _keyManager common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.SetKeyManagerOf(&_PublicLock.TransactOpts, _tokenId, _keyManager)
}

// SetKeyManagerOf is a paid mutator transaction binding the contract method 0xb11d7ec1.
//
// Solidity: function setKeyManagerOf(uint256 _tokenId, address _keyManager) returns()
func (_PublicLock *PublicLockTransactorSession) SetKeyManagerOf(_tokenId *big.Int, _keyManager common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.SetKeyManagerOf(&_PublicLock.TransactOpts, _tokenId, _keyManager)
}

// SetLockMetadata is a paid mutator transaction binding the contract method 0xd1b8759b.
//
// Solidity: function setLockMetadata(string _lockName, string _lockSymbol, string _baseTokenURI) returns()
func (_PublicLock *PublicLockTransactor) SetLockMetadata(opts *bind.TransactOpts, _lockName string, _lockSymbol string, _baseTokenURI string) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "setLockMetadata", _lockName, _lockSymbol, _baseTokenURI)
}

// SetLockMetadata is a paid mutator transaction binding the contract method 0xd1b8759b.
//
// Solidity: function setLockMetadata(string _lockName, string _lockSymbol, string _baseTokenURI) returns()
func (_PublicLock *PublicLockSession) SetLockMetadata(_lockName string, _lockSymbol string, _baseTokenURI string) (*types.Transaction, error) {
	return _PublicLock.Contract.SetLockMetadata(&_PublicLock.TransactOpts, _lockName, _lockSymbol, _baseTokenURI)
}

// SetLockMetadata is a paid mutator transaction binding the contract method 0xd1b8759b.
//
// Solidity: function setLockMetadata(string _lockName, string _lockSymbol, string _baseTokenURI) returns()
func (_PublicLock *PublicLockTransactorSession) SetLockMetadata(_lockName string, _lockSymbol string, _baseTokenURI string) (*types.Transaction, error) {
	return _PublicLock.Contract.SetLockMetadata(&_PublicLock.TransactOpts, _lockName, _lockSymbol, _baseTokenURI)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address account) returns()
func (_PublicLock *PublicLockTransactor) SetOwner(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "setOwner", account)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address account) returns()
func (_PublicLock *PublicLockSession) SetOwner(account common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.SetOwner(&_PublicLock.TransactOpts, account)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address account) returns()
func (_PublicLock *PublicLockTransactorSession) SetOwner(account common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.SetOwner(&_PublicLock.TransactOpts, account)
}

// SetReferrerFee is a paid mutator transaction binding the contract method 0xdebe2b0d.
//
// Solidity: function setReferrerFee(address _referrer, uint256 _feeBasisPoint) returns()
func (_PublicLock *PublicLockTransactor) SetReferrerFee(opts *bind.TransactOpts, _referrer common.Address, _feeBasisPoint *big.Int) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "setReferrerFee", _referrer, _feeBasisPoint)
}

// SetReferrerFee is a paid mutator transaction binding the contract method 0xdebe2b0d.
//
// Solidity: function setReferrerFee(address _referrer, uint256 _feeBasisPoint) returns()
func (_PublicLock *PublicLockSession) SetReferrerFee(_referrer common.Address, _feeBasisPoint *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.SetReferrerFee(&_PublicLock.TransactOpts, _referrer, _feeBasisPoint)
}

// SetReferrerFee is a paid mutator transaction binding the contract method 0xdebe2b0d.
//
// Solidity: function setReferrerFee(address _referrer, uint256 _feeBasisPoint) returns()
func (_PublicLock *PublicLockTransactorSession) SetReferrerFee(_referrer common.Address, _feeBasisPoint *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.SetReferrerFee(&_PublicLock.TransactOpts, _referrer, _feeBasisPoint)
}

// ShareKey is a paid mutator transaction binding the contract method 0xf12c6b6e.
//
// Solidity: function shareKey(address _to, uint256 _tokenIdFrom, uint256 _timeShared) returns()
func (_PublicLock *PublicLockTransactor) ShareKey(opts *bind.TransactOpts, _to common.Address, _tokenIdFrom *big.Int, _timeShared *big.Int) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "shareKey", _to, _tokenIdFrom, _timeShared)
}

// ShareKey is a paid mutator transaction binding the contract method 0xf12c6b6e.
//
// Solidity: function shareKey(address _to, uint256 _tokenIdFrom, uint256 _timeShared) returns()
func (_PublicLock *PublicLockSession) ShareKey(_to common.Address, _tokenIdFrom *big.Int, _timeShared *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.ShareKey(&_PublicLock.TransactOpts, _to, _tokenIdFrom, _timeShared)
}

// ShareKey is a paid mutator transaction binding the contract method 0xf12c6b6e.
//
// Solidity: function shareKey(address _to, uint256 _tokenIdFrom, uint256 _timeShared) returns()
func (_PublicLock *PublicLockTransactorSession) ShareKey(_to common.Address, _tokenIdFrom *big.Int, _timeShared *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.ShareKey(&_PublicLock.TransactOpts, _to, _tokenIdFrom, _timeShared)
}

// Transfer is a paid mutator transaction binding the contract method 0xf8548e36.
//
// Solidity: function transfer(uint256 _tokenId, address _to, uint256 _valueBasisPoint) returns(bool success)
func (_PublicLock *PublicLockTransactor) Transfer(opts *bind.TransactOpts, _tokenId *big.Int, _to common.Address, _valueBasisPoint *big.Int) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "transfer", _tokenId, _to, _valueBasisPoint)
}

// Transfer is a paid mutator transaction binding the contract method 0xf8548e36.
//
// Solidity: function transfer(uint256 _tokenId, address _to, uint256 _valueBasisPoint) returns(bool success)
func (_PublicLock *PublicLockSession) Transfer(_tokenId *big.Int, _to common.Address, _valueBasisPoint *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.Transfer(&_PublicLock.TransactOpts, _tokenId, _to, _valueBasisPoint)
}

// Transfer is a paid mutator transaction binding the contract method 0xf8548e36.
//
// Solidity: function transfer(uint256 _tokenId, address _to, uint256 _valueBasisPoint) returns(bool success)
func (_PublicLock *PublicLockTransactorSession) Transfer(_tokenId *big.Int, _to common.Address, _valueBasisPoint *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.Transfer(&_PublicLock.TransactOpts, _tokenId, _to, _valueBasisPoint)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _recipient, uint256 _tokenId) returns()
func (_PublicLock *PublicLockTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _recipient common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "transferFrom", _from, _recipient, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _recipient, uint256 _tokenId) returns()
func (_PublicLock *PublicLockSession) TransferFrom(_from common.Address, _recipient common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.TransferFrom(&_PublicLock.TransactOpts, _from, _recipient, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _recipient, uint256 _tokenId) returns()
func (_PublicLock *PublicLockTransactorSession) TransferFrom(_from common.Address, _recipient common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.TransferFrom(&_PublicLock.TransactOpts, _from, _recipient, _tokenId)
}

// UnlendKey is a paid mutator transaction binding the contract method 0x407dc589.
//
// Solidity: function unlendKey(address _recipient, uint256 _tokenId) returns()
func (_PublicLock *PublicLockTransactor) UnlendKey(opts *bind.TransactOpts, _recipient common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "unlendKey", _recipient, _tokenId)
}

// UnlendKey is a paid mutator transaction binding the contract method 0x407dc589.
//
// Solidity: function unlendKey(address _recipient, uint256 _tokenId) returns()
func (_PublicLock *PublicLockSession) UnlendKey(_recipient common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.UnlendKey(&_PublicLock.TransactOpts, _recipient, _tokenId)
}

// UnlendKey is a paid mutator transaction binding the contract method 0x407dc589.
//
// Solidity: function unlendKey(address _recipient, uint256 _tokenId) returns()
func (_PublicLock *PublicLockTransactorSession) UnlendKey(_recipient common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.UnlendKey(&_PublicLock.TransactOpts, _recipient, _tokenId)
}

// UpdateKeyPricing is a paid mutator transaction binding the contract method 0xa2e4cd2e.
//
// Solidity: function updateKeyPricing(uint256 _keyPrice, address _tokenAddress) returns()
func (_PublicLock *PublicLockTransactor) UpdateKeyPricing(opts *bind.TransactOpts, _keyPrice *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "updateKeyPricing", _keyPrice, _tokenAddress)
}

// UpdateKeyPricing is a paid mutator transaction binding the contract method 0xa2e4cd2e.
//
// Solidity: function updateKeyPricing(uint256 _keyPrice, address _tokenAddress) returns()
func (_PublicLock *PublicLockSession) UpdateKeyPricing(_keyPrice *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.UpdateKeyPricing(&_PublicLock.TransactOpts, _keyPrice, _tokenAddress)
}

// UpdateKeyPricing is a paid mutator transaction binding the contract method 0xa2e4cd2e.
//
// Solidity: function updateKeyPricing(uint256 _keyPrice, address _tokenAddress) returns()
func (_PublicLock *PublicLockTransactorSession) UpdateKeyPricing(_keyPrice *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _PublicLock.Contract.UpdateKeyPricing(&_PublicLock.TransactOpts, _keyPrice, _tokenAddress)
}

// UpdateLockConfig is a paid mutator transaction binding the contract method 0x282478df.
//
// Solidity: function updateLockConfig(uint256 _newExpirationDuration, uint256 _maxNumberOfKeys, uint256 _maxKeysPerAcccount) returns()
func (_PublicLock *PublicLockTransactor) UpdateLockConfig(opts *bind.TransactOpts, _newExpirationDuration *big.Int, _maxNumberOfKeys *big.Int, _maxKeysPerAcccount *big.Int) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "updateLockConfig", _newExpirationDuration, _maxNumberOfKeys, _maxKeysPerAcccount)
}

// UpdateLockConfig is a paid mutator transaction binding the contract method 0x282478df.
//
// Solidity: function updateLockConfig(uint256 _newExpirationDuration, uint256 _maxNumberOfKeys, uint256 _maxKeysPerAcccount) returns()
func (_PublicLock *PublicLockSession) UpdateLockConfig(_newExpirationDuration *big.Int, _maxNumberOfKeys *big.Int, _maxKeysPerAcccount *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.UpdateLockConfig(&_PublicLock.TransactOpts, _newExpirationDuration, _maxNumberOfKeys, _maxKeysPerAcccount)
}

// UpdateLockConfig is a paid mutator transaction binding the contract method 0x282478df.
//
// Solidity: function updateLockConfig(uint256 _newExpirationDuration, uint256 _maxNumberOfKeys, uint256 _maxKeysPerAcccount) returns()
func (_PublicLock *PublicLockTransactorSession) UpdateLockConfig(_newExpirationDuration *big.Int, _maxNumberOfKeys *big.Int, _maxKeysPerAcccount *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.UpdateLockConfig(&_PublicLock.TransactOpts, _newExpirationDuration, _maxNumberOfKeys, _maxKeysPerAcccount)
}

// UpdateRefundPenalty is a paid mutator transaction binding the contract method 0x39f46986.
//
// Solidity: function updateRefundPenalty(uint256 _freeTrialLength, uint256 _refundPenaltyBasisPoints) returns()
func (_PublicLock *PublicLockTransactor) UpdateRefundPenalty(opts *bind.TransactOpts, _freeTrialLength *big.Int, _refundPenaltyBasisPoints *big.Int) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "updateRefundPenalty", _freeTrialLength, _refundPenaltyBasisPoints)
}

// UpdateRefundPenalty is a paid mutator transaction binding the contract method 0x39f46986.
//
// Solidity: function updateRefundPenalty(uint256 _freeTrialLength, uint256 _refundPenaltyBasisPoints) returns()
func (_PublicLock *PublicLockSession) UpdateRefundPenalty(_freeTrialLength *big.Int, _refundPenaltyBasisPoints *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.UpdateRefundPenalty(&_PublicLock.TransactOpts, _freeTrialLength, _refundPenaltyBasisPoints)
}

// UpdateRefundPenalty is a paid mutator transaction binding the contract method 0x39f46986.
//
// Solidity: function updateRefundPenalty(uint256 _freeTrialLength, uint256 _refundPenaltyBasisPoints) returns()
func (_PublicLock *PublicLockTransactorSession) UpdateRefundPenalty(_freeTrialLength *big.Int, _refundPenaltyBasisPoints *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.UpdateRefundPenalty(&_PublicLock.TransactOpts, _freeTrialLength, _refundPenaltyBasisPoints)
}

// UpdateSchemaVersion is a paid mutator transaction binding the contract method 0xf32e8b24.
//
// Solidity: function updateSchemaVersion() returns()
func (_PublicLock *PublicLockTransactor) UpdateSchemaVersion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "updateSchemaVersion")
}

// UpdateSchemaVersion is a paid mutator transaction binding the contract method 0xf32e8b24.
//
// Solidity: function updateSchemaVersion() returns()
func (_PublicLock *PublicLockSession) UpdateSchemaVersion() (*types.Transaction, error) {
	return _PublicLock.Contract.UpdateSchemaVersion(&_PublicLock.TransactOpts)
}

// UpdateSchemaVersion is a paid mutator transaction binding the contract method 0xf32e8b24.
//
// Solidity: function updateSchemaVersion() returns()
func (_PublicLock *PublicLockTransactorSession) UpdateSchemaVersion() (*types.Transaction, error) {
	return _PublicLock.Contract.UpdateSchemaVersion(&_PublicLock.TransactOpts)
}

// UpdateTransferFee is a paid mutator transaction binding the contract method 0x8577a6d5.
//
// Solidity: function updateTransferFee(uint256 _transferFeeBasisPoints) returns()
func (_PublicLock *PublicLockTransactor) UpdateTransferFee(opts *bind.TransactOpts, _transferFeeBasisPoints *big.Int) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "updateTransferFee", _transferFeeBasisPoints)
}

// UpdateTransferFee is a paid mutator transaction binding the contract method 0x8577a6d5.
//
// Solidity: function updateTransferFee(uint256 _transferFeeBasisPoints) returns()
func (_PublicLock *PublicLockSession) UpdateTransferFee(_transferFeeBasisPoints *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.UpdateTransferFee(&_PublicLock.TransactOpts, _transferFeeBasisPoints)
}

// UpdateTransferFee is a paid mutator transaction binding the contract method 0x8577a6d5.
//
// Solidity: function updateTransferFee(uint256 _transferFeeBasisPoints) returns()
func (_PublicLock *PublicLockTransactorSession) UpdateTransferFee(_transferFeeBasisPoints *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.UpdateTransferFee(&_PublicLock.TransactOpts, _transferFeeBasisPoints)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _tokenAddress, address _recipient, uint256 _amount) returns()
func (_PublicLock *PublicLockTransactor) Withdraw(opts *bind.TransactOpts, _tokenAddress common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PublicLock.contract.Transact(opts, "withdraw", _tokenAddress, _recipient, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _tokenAddress, address _recipient, uint256 _amount) returns()
func (_PublicLock *PublicLockSession) Withdraw(_tokenAddress common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.Withdraw(&_PublicLock.TransactOpts, _tokenAddress, _recipient, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _tokenAddress, address _recipient, uint256 _amount) returns()
func (_PublicLock *PublicLockTransactorSession) Withdraw(_tokenAddress common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PublicLock.Contract.Withdraw(&_PublicLock.TransactOpts, _tokenAddress, _recipient, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PublicLock *PublicLockTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicLock.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PublicLock *PublicLockSession) Receive() (*types.Transaction, error) {
	return _PublicLock.Contract.Receive(&_PublicLock.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PublicLock *PublicLockTransactorSession) Receive() (*types.Transaction, error) {
	return _PublicLock.Contract.Receive(&_PublicLock.TransactOpts)
}

// PublicLockApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PublicLock contract.
type PublicLockApprovalIterator struct {
	Event *PublicLockApproval // Event containing the contract specifics and raw log

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
func (it *PublicLockApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockApproval)
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
		it.Event = new(PublicLockApproval)
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
func (it *PublicLockApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockApproval represents a Approval event raised by the PublicLock contract.
type PublicLockApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_PublicLock *PublicLockFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*PublicLockApprovalIterator, error) {

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

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PublicLockApprovalIterator{contract: _PublicLock.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_PublicLock *PublicLockFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PublicLockApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockApproval)
				if err := _PublicLock.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_PublicLock *PublicLockFilterer) ParseApproval(log types.Log) (*PublicLockApproval, error) {
	event := new(PublicLockApproval)
	if err := _PublicLock.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the PublicLock contract.
type PublicLockApprovalForAllIterator struct {
	Event *PublicLockApprovalForAll // Event containing the contract specifics and raw log

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
func (it *PublicLockApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockApprovalForAll)
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
		it.Event = new(PublicLockApprovalForAll)
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
func (it *PublicLockApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockApprovalForAll represents a ApprovalForAll event raised by the PublicLock contract.
type PublicLockApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_PublicLock *PublicLockFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*PublicLockApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PublicLockApprovalForAllIterator{contract: _PublicLock.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_PublicLock *PublicLockFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *PublicLockApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockApprovalForAll)
				if err := _PublicLock.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_PublicLock *PublicLockFilterer) ParseApprovalForAll(log types.Log) (*PublicLockApprovalForAll, error) {
	event := new(PublicLockApprovalForAll)
	if err := _PublicLock.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockCancelKeyIterator is returned from FilterCancelKey and is used to iterate over the raw logs and unpacked data for CancelKey events raised by the PublicLock contract.
type PublicLockCancelKeyIterator struct {
	Event *PublicLockCancelKey // Event containing the contract specifics and raw log

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
func (it *PublicLockCancelKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockCancelKey)
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
		it.Event = new(PublicLockCancelKey)
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
func (it *PublicLockCancelKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockCancelKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockCancelKey represents a CancelKey event raised by the PublicLock contract.
type PublicLockCancelKey struct {
	TokenId *big.Int
	Owner   common.Address
	SendTo  common.Address
	Refund  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCancelKey is a free log retrieval operation binding the contract event 0x0a7068a9989857441c039a14a42b67ed71dd1fcfe5a9b17cc87b252e47bce528.
//
// Solidity: event CancelKey(uint256 indexed tokenId, address indexed owner, address indexed sendTo, uint256 refund)
func (_PublicLock *PublicLockFilterer) FilterCancelKey(opts *bind.FilterOpts, tokenId []*big.Int, owner []common.Address, sendTo []common.Address) (*PublicLockCancelKeyIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var sendToRule []interface{}
	for _, sendToItem := range sendTo {
		sendToRule = append(sendToRule, sendToItem)
	}

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "CancelKey", tokenIdRule, ownerRule, sendToRule)
	if err != nil {
		return nil, err
	}
	return &PublicLockCancelKeyIterator{contract: _PublicLock.contract, event: "CancelKey", logs: logs, sub: sub}, nil
}

// WatchCancelKey is a free log subscription operation binding the contract event 0x0a7068a9989857441c039a14a42b67ed71dd1fcfe5a9b17cc87b252e47bce528.
//
// Solidity: event CancelKey(uint256 indexed tokenId, address indexed owner, address indexed sendTo, uint256 refund)
func (_PublicLock *PublicLockFilterer) WatchCancelKey(opts *bind.WatchOpts, sink chan<- *PublicLockCancelKey, tokenId []*big.Int, owner []common.Address, sendTo []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var sendToRule []interface{}
	for _, sendToItem := range sendTo {
		sendToRule = append(sendToRule, sendToItem)
	}

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "CancelKey", tokenIdRule, ownerRule, sendToRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockCancelKey)
				if err := _PublicLock.contract.UnpackLog(event, "CancelKey", log); err != nil {
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

// ParseCancelKey is a log parse operation binding the contract event 0x0a7068a9989857441c039a14a42b67ed71dd1fcfe5a9b17cc87b252e47bce528.
//
// Solidity: event CancelKey(uint256 indexed tokenId, address indexed owner, address indexed sendTo, uint256 refund)
func (_PublicLock *PublicLockFilterer) ParseCancelKey(log types.Log) (*PublicLockCancelKey, error) {
	event := new(PublicLockCancelKey)
	if err := _PublicLock.contract.UnpackLog(event, "CancelKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockExpirationChangedIterator is returned from FilterExpirationChanged and is used to iterate over the raw logs and unpacked data for ExpirationChanged events raised by the PublicLock contract.
type PublicLockExpirationChangedIterator struct {
	Event *PublicLockExpirationChanged // Event containing the contract specifics and raw log

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
func (it *PublicLockExpirationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockExpirationChanged)
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
		it.Event = new(PublicLockExpirationChanged)
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
func (it *PublicLockExpirationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockExpirationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockExpirationChanged represents a ExpirationChanged event raised by the PublicLock contract.
type PublicLockExpirationChanged struct {
	TokenId       *big.Int
	NewExpiration *big.Int
	Amount        *big.Int
	TimeAdded     bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExpirationChanged is a free log retrieval operation binding the contract event 0x3c907806849e9204e0e26bb095dfe4b3071576c4323f766735c548211556d052.
//
// Solidity: event ExpirationChanged(uint256 indexed tokenId, uint256 newExpiration, uint256 amount, bool timeAdded)
func (_PublicLock *PublicLockFilterer) FilterExpirationChanged(opts *bind.FilterOpts, tokenId []*big.Int) (*PublicLockExpirationChangedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "ExpirationChanged", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PublicLockExpirationChangedIterator{contract: _PublicLock.contract, event: "ExpirationChanged", logs: logs, sub: sub}, nil
}

// WatchExpirationChanged is a free log subscription operation binding the contract event 0x3c907806849e9204e0e26bb095dfe4b3071576c4323f766735c548211556d052.
//
// Solidity: event ExpirationChanged(uint256 indexed tokenId, uint256 newExpiration, uint256 amount, bool timeAdded)
func (_PublicLock *PublicLockFilterer) WatchExpirationChanged(opts *bind.WatchOpts, sink chan<- *PublicLockExpirationChanged, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "ExpirationChanged", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockExpirationChanged)
				if err := _PublicLock.contract.UnpackLog(event, "ExpirationChanged", log); err != nil {
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

// ParseExpirationChanged is a log parse operation binding the contract event 0x3c907806849e9204e0e26bb095dfe4b3071576c4323f766735c548211556d052.
//
// Solidity: event ExpirationChanged(uint256 indexed tokenId, uint256 newExpiration, uint256 amount, bool timeAdded)
func (_PublicLock *PublicLockFilterer) ParseExpirationChanged(log types.Log) (*PublicLockExpirationChanged, error) {
	event := new(PublicLockExpirationChanged)
	if err := _PublicLock.contract.UnpackLog(event, "ExpirationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockExpireKeyIterator is returned from FilterExpireKey and is used to iterate over the raw logs and unpacked data for ExpireKey events raised by the PublicLock contract.
type PublicLockExpireKeyIterator struct {
	Event *PublicLockExpireKey // Event containing the contract specifics and raw log

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
func (it *PublicLockExpireKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockExpireKey)
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
		it.Event = new(PublicLockExpireKey)
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
func (it *PublicLockExpireKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockExpireKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockExpireKey represents a ExpireKey event raised by the PublicLock contract.
type PublicLockExpireKey struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExpireKey is a free log retrieval operation binding the contract event 0x59f2fe866dd27a1c2d34115520888c3150365cbc931aab97fa88c4b9ab40b795.
//
// Solidity: event ExpireKey(uint256 indexed tokenId)
func (_PublicLock *PublicLockFilterer) FilterExpireKey(opts *bind.FilterOpts, tokenId []*big.Int) (*PublicLockExpireKeyIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "ExpireKey", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PublicLockExpireKeyIterator{contract: _PublicLock.contract, event: "ExpireKey", logs: logs, sub: sub}, nil
}

// WatchExpireKey is a free log subscription operation binding the contract event 0x59f2fe866dd27a1c2d34115520888c3150365cbc931aab97fa88c4b9ab40b795.
//
// Solidity: event ExpireKey(uint256 indexed tokenId)
func (_PublicLock *PublicLockFilterer) WatchExpireKey(opts *bind.WatchOpts, sink chan<- *PublicLockExpireKey, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "ExpireKey", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockExpireKey)
				if err := _PublicLock.contract.UnpackLog(event, "ExpireKey", log); err != nil {
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

// ParseExpireKey is a log parse operation binding the contract event 0x59f2fe866dd27a1c2d34115520888c3150365cbc931aab97fa88c4b9ab40b795.
//
// Solidity: event ExpireKey(uint256 indexed tokenId)
func (_PublicLock *PublicLockFilterer) ParseExpireKey(log types.Log) (*PublicLockExpireKey, error) {
	event := new(PublicLockExpireKey)
	if err := _PublicLock.contract.UnpackLog(event, "ExpireKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockGasRefundedIterator is returned from FilterGasRefunded and is used to iterate over the raw logs and unpacked data for GasRefunded events raised by the PublicLock contract.
type PublicLockGasRefundedIterator struct {
	Event *PublicLockGasRefunded // Event containing the contract specifics and raw log

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
func (it *PublicLockGasRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockGasRefunded)
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
		it.Event = new(PublicLockGasRefunded)
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
func (it *PublicLockGasRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockGasRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockGasRefunded represents a GasRefunded event raised by the PublicLock contract.
type PublicLockGasRefunded struct {
	Receiver       common.Address
	RefundedAmount *big.Int
	TokenAddress   common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGasRefunded is a free log retrieval operation binding the contract event 0x522a883b471164223f18b50f326da8671372b64b4792eac0e63d447e714c3e3b.
//
// Solidity: event GasRefunded(address indexed receiver, uint256 refundedAmount, address tokenAddress)
func (_PublicLock *PublicLockFilterer) FilterGasRefunded(opts *bind.FilterOpts, receiver []common.Address) (*PublicLockGasRefundedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "GasRefunded", receiverRule)
	if err != nil {
		return nil, err
	}
	return &PublicLockGasRefundedIterator{contract: _PublicLock.contract, event: "GasRefunded", logs: logs, sub: sub}, nil
}

// WatchGasRefunded is a free log subscription operation binding the contract event 0x522a883b471164223f18b50f326da8671372b64b4792eac0e63d447e714c3e3b.
//
// Solidity: event GasRefunded(address indexed receiver, uint256 refundedAmount, address tokenAddress)
func (_PublicLock *PublicLockFilterer) WatchGasRefunded(opts *bind.WatchOpts, sink chan<- *PublicLockGasRefunded, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "GasRefunded", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockGasRefunded)
				if err := _PublicLock.contract.UnpackLog(event, "GasRefunded", log); err != nil {
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

// ParseGasRefunded is a log parse operation binding the contract event 0x522a883b471164223f18b50f326da8671372b64b4792eac0e63d447e714c3e3b.
//
// Solidity: event GasRefunded(address indexed receiver, uint256 refundedAmount, address tokenAddress)
func (_PublicLock *PublicLockFilterer) ParseGasRefunded(log types.Log) (*PublicLockGasRefunded, error) {
	event := new(PublicLockGasRefunded)
	if err := _PublicLock.contract.UnpackLog(event, "GasRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PublicLock contract.
type PublicLockInitializedIterator struct {
	Event *PublicLockInitialized // Event containing the contract specifics and raw log

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
func (it *PublicLockInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockInitialized)
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
		it.Event = new(PublicLockInitialized)
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
func (it *PublicLockInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockInitialized represents a Initialized event raised by the PublicLock contract.
type PublicLockInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PublicLock *PublicLockFilterer) FilterInitialized(opts *bind.FilterOpts) (*PublicLockInitializedIterator, error) {

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PublicLockInitializedIterator{contract: _PublicLock.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PublicLock *PublicLockFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PublicLockInitialized) (event.Subscription, error) {

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockInitialized)
				if err := _PublicLock.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PublicLock *PublicLockFilterer) ParseInitialized(log types.Log) (*PublicLockInitialized, error) {
	event := new(PublicLockInitialized)
	if err := _PublicLock.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockKeyExtendedIterator is returned from FilterKeyExtended and is used to iterate over the raw logs and unpacked data for KeyExtended events raised by the PublicLock contract.
type PublicLockKeyExtendedIterator struct {
	Event *PublicLockKeyExtended // Event containing the contract specifics and raw log

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
func (it *PublicLockKeyExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockKeyExtended)
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
		it.Event = new(PublicLockKeyExtended)
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
func (it *PublicLockKeyExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockKeyExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockKeyExtended represents a KeyExtended event raised by the PublicLock contract.
type PublicLockKeyExtended struct {
	TokenId      *big.Int
	NewTimestamp *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterKeyExtended is a free log retrieval operation binding the contract event 0x3ca112768ff7861e008ace1c11570c52e404c043e585545b5957a1e20961dde3.
//
// Solidity: event KeyExtended(uint256 indexed tokenId, uint256 newTimestamp)
func (_PublicLock *PublicLockFilterer) FilterKeyExtended(opts *bind.FilterOpts, tokenId []*big.Int) (*PublicLockKeyExtendedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "KeyExtended", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PublicLockKeyExtendedIterator{contract: _PublicLock.contract, event: "KeyExtended", logs: logs, sub: sub}, nil
}

// WatchKeyExtended is a free log subscription operation binding the contract event 0x3ca112768ff7861e008ace1c11570c52e404c043e585545b5957a1e20961dde3.
//
// Solidity: event KeyExtended(uint256 indexed tokenId, uint256 newTimestamp)
func (_PublicLock *PublicLockFilterer) WatchKeyExtended(opts *bind.WatchOpts, sink chan<- *PublicLockKeyExtended, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "KeyExtended", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockKeyExtended)
				if err := _PublicLock.contract.UnpackLog(event, "KeyExtended", log); err != nil {
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

// ParseKeyExtended is a log parse operation binding the contract event 0x3ca112768ff7861e008ace1c11570c52e404c043e585545b5957a1e20961dde3.
//
// Solidity: event KeyExtended(uint256 indexed tokenId, uint256 newTimestamp)
func (_PublicLock *PublicLockFilterer) ParseKeyExtended(log types.Log) (*PublicLockKeyExtended, error) {
	event := new(PublicLockKeyExtended)
	if err := _PublicLock.contract.UnpackLog(event, "KeyExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockKeyGranterAddedIterator is returned from FilterKeyGranterAdded and is used to iterate over the raw logs and unpacked data for KeyGranterAdded events raised by the PublicLock contract.
type PublicLockKeyGranterAddedIterator struct {
	Event *PublicLockKeyGranterAdded // Event containing the contract specifics and raw log

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
func (it *PublicLockKeyGranterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockKeyGranterAdded)
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
		it.Event = new(PublicLockKeyGranterAdded)
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
func (it *PublicLockKeyGranterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockKeyGranterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockKeyGranterAdded represents a KeyGranterAdded event raised by the PublicLock contract.
type PublicLockKeyGranterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKeyGranterAdded is a free log retrieval operation binding the contract event 0x684f8a71407db0c334454ebe9c9b288549317893a20b10dc779ec5c118dcd121.
//
// Solidity: event KeyGranterAdded(address indexed account)
func (_PublicLock *PublicLockFilterer) FilterKeyGranterAdded(opts *bind.FilterOpts, account []common.Address) (*PublicLockKeyGranterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "KeyGranterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &PublicLockKeyGranterAddedIterator{contract: _PublicLock.contract, event: "KeyGranterAdded", logs: logs, sub: sub}, nil
}

// WatchKeyGranterAdded is a free log subscription operation binding the contract event 0x684f8a71407db0c334454ebe9c9b288549317893a20b10dc779ec5c118dcd121.
//
// Solidity: event KeyGranterAdded(address indexed account)
func (_PublicLock *PublicLockFilterer) WatchKeyGranterAdded(opts *bind.WatchOpts, sink chan<- *PublicLockKeyGranterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "KeyGranterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockKeyGranterAdded)
				if err := _PublicLock.contract.UnpackLog(event, "KeyGranterAdded", log); err != nil {
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

// ParseKeyGranterAdded is a log parse operation binding the contract event 0x684f8a71407db0c334454ebe9c9b288549317893a20b10dc779ec5c118dcd121.
//
// Solidity: event KeyGranterAdded(address indexed account)
func (_PublicLock *PublicLockFilterer) ParseKeyGranterAdded(log types.Log) (*PublicLockKeyGranterAdded, error) {
	event := new(PublicLockKeyGranterAdded)
	if err := _PublicLock.contract.UnpackLog(event, "KeyGranterAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockKeyGranterRemovedIterator is returned from FilterKeyGranterRemoved and is used to iterate over the raw logs and unpacked data for KeyGranterRemoved events raised by the PublicLock contract.
type PublicLockKeyGranterRemovedIterator struct {
	Event *PublicLockKeyGranterRemoved // Event containing the contract specifics and raw log

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
func (it *PublicLockKeyGranterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockKeyGranterRemoved)
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
		it.Event = new(PublicLockKeyGranterRemoved)
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
func (it *PublicLockKeyGranterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockKeyGranterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockKeyGranterRemoved represents a KeyGranterRemoved event raised by the PublicLock contract.
type PublicLockKeyGranterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKeyGranterRemoved is a free log retrieval operation binding the contract event 0x766f6199fea59554b9ff688e413302b9200f85d74811c053c12d945ac6d8dd7a.
//
// Solidity: event KeyGranterRemoved(address indexed account)
func (_PublicLock *PublicLockFilterer) FilterKeyGranterRemoved(opts *bind.FilterOpts, account []common.Address) (*PublicLockKeyGranterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "KeyGranterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &PublicLockKeyGranterRemovedIterator{contract: _PublicLock.contract, event: "KeyGranterRemoved", logs: logs, sub: sub}, nil
}

// WatchKeyGranterRemoved is a free log subscription operation binding the contract event 0x766f6199fea59554b9ff688e413302b9200f85d74811c053c12d945ac6d8dd7a.
//
// Solidity: event KeyGranterRemoved(address indexed account)
func (_PublicLock *PublicLockFilterer) WatchKeyGranterRemoved(opts *bind.WatchOpts, sink chan<- *PublicLockKeyGranterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "KeyGranterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockKeyGranterRemoved)
				if err := _PublicLock.contract.UnpackLog(event, "KeyGranterRemoved", log); err != nil {
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

// ParseKeyGranterRemoved is a log parse operation binding the contract event 0x766f6199fea59554b9ff688e413302b9200f85d74811c053c12d945ac6d8dd7a.
//
// Solidity: event KeyGranterRemoved(address indexed account)
func (_PublicLock *PublicLockFilterer) ParseKeyGranterRemoved(log types.Log) (*PublicLockKeyGranterRemoved, error) {
	event := new(PublicLockKeyGranterRemoved)
	if err := _PublicLock.contract.UnpackLog(event, "KeyGranterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockKeyManagerChangedIterator is returned from FilterKeyManagerChanged and is used to iterate over the raw logs and unpacked data for KeyManagerChanged events raised by the PublicLock contract.
type PublicLockKeyManagerChangedIterator struct {
	Event *PublicLockKeyManagerChanged // Event containing the contract specifics and raw log

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
func (it *PublicLockKeyManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockKeyManagerChanged)
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
		it.Event = new(PublicLockKeyManagerChanged)
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
func (it *PublicLockKeyManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockKeyManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockKeyManagerChanged represents a KeyManagerChanged event raised by the PublicLock contract.
type PublicLockKeyManagerChanged struct {
	TokenId    *big.Int
	NewManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterKeyManagerChanged is a free log retrieval operation binding the contract event 0x9d2895c45a420624de863a2f437b022d879f457bf7a829044055a10c5a6fd5e3.
//
// Solidity: event KeyManagerChanged(uint256 indexed _tokenId, address indexed _newManager)
func (_PublicLock *PublicLockFilterer) FilterKeyManagerChanged(opts *bind.FilterOpts, _tokenId []*big.Int, _newManager []common.Address) (*PublicLockKeyManagerChangedIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _newManagerRule []interface{}
	for _, _newManagerItem := range _newManager {
		_newManagerRule = append(_newManagerRule, _newManagerItem)
	}

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "KeyManagerChanged", _tokenIdRule, _newManagerRule)
	if err != nil {
		return nil, err
	}
	return &PublicLockKeyManagerChangedIterator{contract: _PublicLock.contract, event: "KeyManagerChanged", logs: logs, sub: sub}, nil
}

// WatchKeyManagerChanged is a free log subscription operation binding the contract event 0x9d2895c45a420624de863a2f437b022d879f457bf7a829044055a10c5a6fd5e3.
//
// Solidity: event KeyManagerChanged(uint256 indexed _tokenId, address indexed _newManager)
func (_PublicLock *PublicLockFilterer) WatchKeyManagerChanged(opts *bind.WatchOpts, sink chan<- *PublicLockKeyManagerChanged, _tokenId []*big.Int, _newManager []common.Address) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _newManagerRule []interface{}
	for _, _newManagerItem := range _newManager {
		_newManagerRule = append(_newManagerRule, _newManagerItem)
	}

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "KeyManagerChanged", _tokenIdRule, _newManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockKeyManagerChanged)
				if err := _PublicLock.contract.UnpackLog(event, "KeyManagerChanged", log); err != nil {
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

// ParseKeyManagerChanged is a log parse operation binding the contract event 0x9d2895c45a420624de863a2f437b022d879f457bf7a829044055a10c5a6fd5e3.
//
// Solidity: event KeyManagerChanged(uint256 indexed _tokenId, address indexed _newManager)
func (_PublicLock *PublicLockFilterer) ParseKeyManagerChanged(log types.Log) (*PublicLockKeyManagerChanged, error) {
	event := new(PublicLockKeyManagerChanged)
	if err := _PublicLock.contract.UnpackLog(event, "KeyManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockKeysMigratedIterator is returned from FilterKeysMigrated and is used to iterate over the raw logs and unpacked data for KeysMigrated events raised by the PublicLock contract.
type PublicLockKeysMigratedIterator struct {
	Event *PublicLockKeysMigrated // Event containing the contract specifics and raw log

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
func (it *PublicLockKeysMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockKeysMigrated)
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
		it.Event = new(PublicLockKeysMigrated)
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
func (it *PublicLockKeysMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockKeysMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockKeysMigrated represents a KeysMigrated event raised by the PublicLock contract.
type PublicLockKeysMigrated struct {
	UpdatedRecordsCount *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterKeysMigrated is a free log retrieval operation binding the contract event 0x4e075b7dcd463fe1b504dad3ccd1c4f87011a70c1e50f1805183df5d559834cd.
//
// Solidity: event KeysMigrated(uint256 updatedRecordsCount)
func (_PublicLock *PublicLockFilterer) FilterKeysMigrated(opts *bind.FilterOpts) (*PublicLockKeysMigratedIterator, error) {

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "KeysMigrated")
	if err != nil {
		return nil, err
	}
	return &PublicLockKeysMigratedIterator{contract: _PublicLock.contract, event: "KeysMigrated", logs: logs, sub: sub}, nil
}

// WatchKeysMigrated is a free log subscription operation binding the contract event 0x4e075b7dcd463fe1b504dad3ccd1c4f87011a70c1e50f1805183df5d559834cd.
//
// Solidity: event KeysMigrated(uint256 updatedRecordsCount)
func (_PublicLock *PublicLockFilterer) WatchKeysMigrated(opts *bind.WatchOpts, sink chan<- *PublicLockKeysMigrated) (event.Subscription, error) {

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "KeysMigrated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockKeysMigrated)
				if err := _PublicLock.contract.UnpackLog(event, "KeysMigrated", log); err != nil {
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

// ParseKeysMigrated is a log parse operation binding the contract event 0x4e075b7dcd463fe1b504dad3ccd1c4f87011a70c1e50f1805183df5d559834cd.
//
// Solidity: event KeysMigrated(uint256 updatedRecordsCount)
func (_PublicLock *PublicLockFilterer) ParseKeysMigrated(log types.Log) (*PublicLockKeysMigrated, error) {
	event := new(PublicLockKeysMigrated)
	if err := _PublicLock.contract.UnpackLog(event, "KeysMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockLockConfigIterator is returned from FilterLockConfig and is used to iterate over the raw logs and unpacked data for LockConfig events raised by the PublicLock contract.
type PublicLockLockConfigIterator struct {
	Event *PublicLockLockConfig // Event containing the contract specifics and raw log

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
func (it *PublicLockLockConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockLockConfig)
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
		it.Event = new(PublicLockLockConfig)
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
func (it *PublicLockLockConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockLockConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockLockConfig represents a LockConfig event raised by the PublicLock contract.
type PublicLockLockConfig struct {
	ExpirationDuration *big.Int
	MaxNumberOfKeys    *big.Int
	MaxKeysPerAcccount *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLockConfig is a free log retrieval operation binding the contract event 0x9a09448a3f24d3a01ccc67103c7cddbeea820176a18182cc83d0bce585f26a5b.
//
// Solidity: event LockConfig(uint256 expirationDuration, uint256 maxNumberOfKeys, uint256 maxKeysPerAcccount)
func (_PublicLock *PublicLockFilterer) FilterLockConfig(opts *bind.FilterOpts) (*PublicLockLockConfigIterator, error) {

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "LockConfig")
	if err != nil {
		return nil, err
	}
	return &PublicLockLockConfigIterator{contract: _PublicLock.contract, event: "LockConfig", logs: logs, sub: sub}, nil
}

// WatchLockConfig is a free log subscription operation binding the contract event 0x9a09448a3f24d3a01ccc67103c7cddbeea820176a18182cc83d0bce585f26a5b.
//
// Solidity: event LockConfig(uint256 expirationDuration, uint256 maxNumberOfKeys, uint256 maxKeysPerAcccount)
func (_PublicLock *PublicLockFilterer) WatchLockConfig(opts *bind.WatchOpts, sink chan<- *PublicLockLockConfig) (event.Subscription, error) {

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "LockConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockLockConfig)
				if err := _PublicLock.contract.UnpackLog(event, "LockConfig", log); err != nil {
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

// ParseLockConfig is a log parse operation binding the contract event 0x9a09448a3f24d3a01ccc67103c7cddbeea820176a18182cc83d0bce585f26a5b.
//
// Solidity: event LockConfig(uint256 expirationDuration, uint256 maxNumberOfKeys, uint256 maxKeysPerAcccount)
func (_PublicLock *PublicLockFilterer) ParseLockConfig(log types.Log) (*PublicLockLockConfig, error) {
	event := new(PublicLockLockConfig)
	if err := _PublicLock.contract.UnpackLog(event, "LockConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockLockManagerAddedIterator is returned from FilterLockManagerAdded and is used to iterate over the raw logs and unpacked data for LockManagerAdded events raised by the PublicLock contract.
type PublicLockLockManagerAddedIterator struct {
	Event *PublicLockLockManagerAdded // Event containing the contract specifics and raw log

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
func (it *PublicLockLockManagerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockLockManagerAdded)
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
		it.Event = new(PublicLockLockManagerAdded)
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
func (it *PublicLockLockManagerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockLockManagerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockLockManagerAdded represents a LockManagerAdded event raised by the PublicLock contract.
type PublicLockLockManagerAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLockManagerAdded is a free log retrieval operation binding the contract event 0x91d5c045d5bd98bf59a379b259ebca05b93bf79af1845fdf87e3172385d4c7f7.
//
// Solidity: event LockManagerAdded(address indexed account)
func (_PublicLock *PublicLockFilterer) FilterLockManagerAdded(opts *bind.FilterOpts, account []common.Address) (*PublicLockLockManagerAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "LockManagerAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &PublicLockLockManagerAddedIterator{contract: _PublicLock.contract, event: "LockManagerAdded", logs: logs, sub: sub}, nil
}

// WatchLockManagerAdded is a free log subscription operation binding the contract event 0x91d5c045d5bd98bf59a379b259ebca05b93bf79af1845fdf87e3172385d4c7f7.
//
// Solidity: event LockManagerAdded(address indexed account)
func (_PublicLock *PublicLockFilterer) WatchLockManagerAdded(opts *bind.WatchOpts, sink chan<- *PublicLockLockManagerAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "LockManagerAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockLockManagerAdded)
				if err := _PublicLock.contract.UnpackLog(event, "LockManagerAdded", log); err != nil {
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

// ParseLockManagerAdded is a log parse operation binding the contract event 0x91d5c045d5bd98bf59a379b259ebca05b93bf79af1845fdf87e3172385d4c7f7.
//
// Solidity: event LockManagerAdded(address indexed account)
func (_PublicLock *PublicLockFilterer) ParseLockManagerAdded(log types.Log) (*PublicLockLockManagerAdded, error) {
	event := new(PublicLockLockManagerAdded)
	if err := _PublicLock.contract.UnpackLog(event, "LockManagerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockLockManagerRemovedIterator is returned from FilterLockManagerRemoved and is used to iterate over the raw logs and unpacked data for LockManagerRemoved events raised by the PublicLock contract.
type PublicLockLockManagerRemovedIterator struct {
	Event *PublicLockLockManagerRemoved // Event containing the contract specifics and raw log

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
func (it *PublicLockLockManagerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockLockManagerRemoved)
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
		it.Event = new(PublicLockLockManagerRemoved)
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
func (it *PublicLockLockManagerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockLockManagerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockLockManagerRemoved represents a LockManagerRemoved event raised by the PublicLock contract.
type PublicLockLockManagerRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLockManagerRemoved is a free log retrieval operation binding the contract event 0x42885193b8178d25fca25a38e6fcc93918501e91be06d85e0c8afb3bad952380.
//
// Solidity: event LockManagerRemoved(address indexed account)
func (_PublicLock *PublicLockFilterer) FilterLockManagerRemoved(opts *bind.FilterOpts, account []common.Address) (*PublicLockLockManagerRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "LockManagerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &PublicLockLockManagerRemovedIterator{contract: _PublicLock.contract, event: "LockManagerRemoved", logs: logs, sub: sub}, nil
}

// WatchLockManagerRemoved is a free log subscription operation binding the contract event 0x42885193b8178d25fca25a38e6fcc93918501e91be06d85e0c8afb3bad952380.
//
// Solidity: event LockManagerRemoved(address indexed account)
func (_PublicLock *PublicLockFilterer) WatchLockManagerRemoved(opts *bind.WatchOpts, sink chan<- *PublicLockLockManagerRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "LockManagerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockLockManagerRemoved)
				if err := _PublicLock.contract.UnpackLog(event, "LockManagerRemoved", log); err != nil {
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

// ParseLockManagerRemoved is a log parse operation binding the contract event 0x42885193b8178d25fca25a38e6fcc93918501e91be06d85e0c8afb3bad952380.
//
// Solidity: event LockManagerRemoved(address indexed account)
func (_PublicLock *PublicLockFilterer) ParseLockManagerRemoved(log types.Log) (*PublicLockLockManagerRemoved, error) {
	event := new(PublicLockLockManagerRemoved)
	if err := _PublicLock.contract.UnpackLog(event, "LockManagerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockLockMetadataIterator is returned from FilterLockMetadata and is used to iterate over the raw logs and unpacked data for LockMetadata events raised by the PublicLock contract.
type PublicLockLockMetadataIterator struct {
	Event *PublicLockLockMetadata // Event containing the contract specifics and raw log

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
func (it *PublicLockLockMetadataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockLockMetadata)
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
		it.Event = new(PublicLockLockMetadata)
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
func (it *PublicLockLockMetadataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockLockMetadataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockLockMetadata represents a LockMetadata event raised by the PublicLock contract.
type PublicLockLockMetadata struct {
	Name         string
	Symbol       string
	BaseTokenURI string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLockMetadata is a free log retrieval operation binding the contract event 0x1e6d6a19e45ae156dcf4155bc83cf8f59e98d536000998f0e95f4cd330ecfb3e.
//
// Solidity: event LockMetadata(string name, string symbol, string baseTokenURI)
func (_PublicLock *PublicLockFilterer) FilterLockMetadata(opts *bind.FilterOpts) (*PublicLockLockMetadataIterator, error) {

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "LockMetadata")
	if err != nil {
		return nil, err
	}
	return &PublicLockLockMetadataIterator{contract: _PublicLock.contract, event: "LockMetadata", logs: logs, sub: sub}, nil
}

// WatchLockMetadata is a free log subscription operation binding the contract event 0x1e6d6a19e45ae156dcf4155bc83cf8f59e98d536000998f0e95f4cd330ecfb3e.
//
// Solidity: event LockMetadata(string name, string symbol, string baseTokenURI)
func (_PublicLock *PublicLockFilterer) WatchLockMetadata(opts *bind.WatchOpts, sink chan<- *PublicLockLockMetadata) (event.Subscription, error) {

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "LockMetadata")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockLockMetadata)
				if err := _PublicLock.contract.UnpackLog(event, "LockMetadata", log); err != nil {
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

// ParseLockMetadata is a log parse operation binding the contract event 0x1e6d6a19e45ae156dcf4155bc83cf8f59e98d536000998f0e95f4cd330ecfb3e.
//
// Solidity: event LockMetadata(string name, string symbol, string baseTokenURI)
func (_PublicLock *PublicLockFilterer) ParseLockMetadata(log types.Log) (*PublicLockLockMetadata, error) {
	event := new(PublicLockLockMetadata)
	if err := _PublicLock.contract.UnpackLog(event, "LockMetadata", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PublicLock contract.
type PublicLockOwnershipTransferredIterator struct {
	Event *PublicLockOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PublicLockOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockOwnershipTransferred)
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
		it.Event = new(PublicLockOwnershipTransferred)
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
func (it *PublicLockOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockOwnershipTransferred represents a OwnershipTransferred event raised by the PublicLock contract.
type PublicLockOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address previousOwner, address newOwner)
func (_PublicLock *PublicLockFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts) (*PublicLockOwnershipTransferredIterator, error) {

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return &PublicLockOwnershipTransferredIterator{contract: _PublicLock.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address previousOwner, address newOwner)
func (_PublicLock *PublicLockFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PublicLockOwnershipTransferred) (event.Subscription, error) {

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockOwnershipTransferred)
				if err := _PublicLock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address previousOwner, address newOwner)
func (_PublicLock *PublicLockFilterer) ParseOwnershipTransferred(log types.Log) (*PublicLockOwnershipTransferred, error) {
	event := new(PublicLockOwnershipTransferred)
	if err := _PublicLock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockPricingChangedIterator is returned from FilterPricingChanged and is used to iterate over the raw logs and unpacked data for PricingChanged events raised by the PublicLock contract.
type PublicLockPricingChangedIterator struct {
	Event *PublicLockPricingChanged // Event containing the contract specifics and raw log

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
func (it *PublicLockPricingChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockPricingChanged)
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
		it.Event = new(PublicLockPricingChanged)
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
func (it *PublicLockPricingChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockPricingChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockPricingChanged represents a PricingChanged event raised by the PublicLock contract.
type PublicLockPricingChanged struct {
	OldKeyPrice     *big.Int
	KeyPrice        *big.Int
	OldTokenAddress common.Address
	TokenAddress    common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPricingChanged is a free log retrieval operation binding the contract event 0x3615065ccf48367ac483ac86701248e2e5ff55bdd9be845007d34a3b68d719d4.
//
// Solidity: event PricingChanged(uint256 oldKeyPrice, uint256 keyPrice, address oldTokenAddress, address tokenAddress)
func (_PublicLock *PublicLockFilterer) FilterPricingChanged(opts *bind.FilterOpts) (*PublicLockPricingChangedIterator, error) {

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "PricingChanged")
	if err != nil {
		return nil, err
	}
	return &PublicLockPricingChangedIterator{contract: _PublicLock.contract, event: "PricingChanged", logs: logs, sub: sub}, nil
}

// WatchPricingChanged is a free log subscription operation binding the contract event 0x3615065ccf48367ac483ac86701248e2e5ff55bdd9be845007d34a3b68d719d4.
//
// Solidity: event PricingChanged(uint256 oldKeyPrice, uint256 keyPrice, address oldTokenAddress, address tokenAddress)
func (_PublicLock *PublicLockFilterer) WatchPricingChanged(opts *bind.WatchOpts, sink chan<- *PublicLockPricingChanged) (event.Subscription, error) {

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "PricingChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockPricingChanged)
				if err := _PublicLock.contract.UnpackLog(event, "PricingChanged", log); err != nil {
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

// ParsePricingChanged is a log parse operation binding the contract event 0x3615065ccf48367ac483ac86701248e2e5ff55bdd9be845007d34a3b68d719d4.
//
// Solidity: event PricingChanged(uint256 oldKeyPrice, uint256 keyPrice, address oldTokenAddress, address tokenAddress)
func (_PublicLock *PublicLockFilterer) ParsePricingChanged(log types.Log) (*PublicLockPricingChanged, error) {
	event := new(PublicLockPricingChanged)
	if err := _PublicLock.contract.UnpackLog(event, "PricingChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockRefundPenaltyChangedIterator is returned from FilterRefundPenaltyChanged and is used to iterate over the raw logs and unpacked data for RefundPenaltyChanged events raised by the PublicLock contract.
type PublicLockRefundPenaltyChangedIterator struct {
	Event *PublicLockRefundPenaltyChanged // Event containing the contract specifics and raw log

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
func (it *PublicLockRefundPenaltyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockRefundPenaltyChanged)
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
		it.Event = new(PublicLockRefundPenaltyChanged)
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
func (it *PublicLockRefundPenaltyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockRefundPenaltyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockRefundPenaltyChanged represents a RefundPenaltyChanged event raised by the PublicLock contract.
type PublicLockRefundPenaltyChanged struct {
	FreeTrialLength          *big.Int
	RefundPenaltyBasisPoints *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterRefundPenaltyChanged is a free log retrieval operation binding the contract event 0xd6867bc538320e67d7bdc35860c27c08486eb490b4fd9b820fff18fb28381d3c.
//
// Solidity: event RefundPenaltyChanged(uint256 freeTrialLength, uint256 refundPenaltyBasisPoints)
func (_PublicLock *PublicLockFilterer) FilterRefundPenaltyChanged(opts *bind.FilterOpts) (*PublicLockRefundPenaltyChangedIterator, error) {

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "RefundPenaltyChanged")
	if err != nil {
		return nil, err
	}
	return &PublicLockRefundPenaltyChangedIterator{contract: _PublicLock.contract, event: "RefundPenaltyChanged", logs: logs, sub: sub}, nil
}

// WatchRefundPenaltyChanged is a free log subscription operation binding the contract event 0xd6867bc538320e67d7bdc35860c27c08486eb490b4fd9b820fff18fb28381d3c.
//
// Solidity: event RefundPenaltyChanged(uint256 freeTrialLength, uint256 refundPenaltyBasisPoints)
func (_PublicLock *PublicLockFilterer) WatchRefundPenaltyChanged(opts *bind.WatchOpts, sink chan<- *PublicLockRefundPenaltyChanged) (event.Subscription, error) {

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "RefundPenaltyChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockRefundPenaltyChanged)
				if err := _PublicLock.contract.UnpackLog(event, "RefundPenaltyChanged", log); err != nil {
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

// ParseRefundPenaltyChanged is a log parse operation binding the contract event 0xd6867bc538320e67d7bdc35860c27c08486eb490b4fd9b820fff18fb28381d3c.
//
// Solidity: event RefundPenaltyChanged(uint256 freeTrialLength, uint256 refundPenaltyBasisPoints)
func (_PublicLock *PublicLockFilterer) ParseRefundPenaltyChanged(log types.Log) (*PublicLockRefundPenaltyChanged, error) {
	event := new(PublicLockRefundPenaltyChanged)
	if err := _PublicLock.contract.UnpackLog(event, "RefundPenaltyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PublicLock contract.
type PublicLockRoleAdminChangedIterator struct {
	Event *PublicLockRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PublicLockRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockRoleAdminChanged)
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
		it.Event = new(PublicLockRoleAdminChanged)
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
func (it *PublicLockRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockRoleAdminChanged represents a RoleAdminChanged event raised by the PublicLock contract.
type PublicLockRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PublicLock *PublicLockFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PublicLockRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PublicLockRoleAdminChangedIterator{contract: _PublicLock.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PublicLock *PublicLockFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PublicLockRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockRoleAdminChanged)
				if err := _PublicLock.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PublicLock *PublicLockFilterer) ParseRoleAdminChanged(log types.Log) (*PublicLockRoleAdminChanged, error) {
	event := new(PublicLockRoleAdminChanged)
	if err := _PublicLock.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PublicLock contract.
type PublicLockRoleGrantedIterator struct {
	Event *PublicLockRoleGranted // Event containing the contract specifics and raw log

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
func (it *PublicLockRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockRoleGranted)
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
		it.Event = new(PublicLockRoleGranted)
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
func (it *PublicLockRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockRoleGranted represents a RoleGranted event raised by the PublicLock contract.
type PublicLockRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PublicLock *PublicLockFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PublicLockRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PublicLockRoleGrantedIterator{contract: _PublicLock.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PublicLock *PublicLockFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PublicLockRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockRoleGranted)
				if err := _PublicLock.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PublicLock *PublicLockFilterer) ParseRoleGranted(log types.Log) (*PublicLockRoleGranted, error) {
	event := new(PublicLockRoleGranted)
	if err := _PublicLock.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PublicLock contract.
type PublicLockRoleRevokedIterator struct {
	Event *PublicLockRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PublicLockRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockRoleRevoked)
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
		it.Event = new(PublicLockRoleRevoked)
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
func (it *PublicLockRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockRoleRevoked represents a RoleRevoked event raised by the PublicLock contract.
type PublicLockRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PublicLock *PublicLockFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PublicLockRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PublicLockRoleRevokedIterator{contract: _PublicLock.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PublicLock *PublicLockFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PublicLockRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockRoleRevoked)
				if err := _PublicLock.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PublicLock *PublicLockFilterer) ParseRoleRevoked(log types.Log) (*PublicLockRoleRevoked, error) {
	event := new(PublicLockRoleRevoked)
	if err := _PublicLock.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PublicLock contract.
type PublicLockTransferIterator struct {
	Event *PublicLockTransfer // Event containing the contract specifics and raw log

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
func (it *PublicLockTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockTransfer)
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
		it.Event = new(PublicLockTransfer)
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
func (it *PublicLockTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockTransfer represents a Transfer event raised by the PublicLock contract.
type PublicLockTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_PublicLock *PublicLockFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*PublicLockTransferIterator, error) {

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

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PublicLockTransferIterator{contract: _PublicLock.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_PublicLock *PublicLockFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PublicLockTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockTransfer)
				if err := _PublicLock.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_PublicLock *PublicLockFilterer) ParseTransfer(log types.Log) (*PublicLockTransfer, error) {
	event := new(PublicLockTransfer)
	if err := _PublicLock.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockTransferFeeChangedIterator is returned from FilterTransferFeeChanged and is used to iterate over the raw logs and unpacked data for TransferFeeChanged events raised by the PublicLock contract.
type PublicLockTransferFeeChangedIterator struct {
	Event *PublicLockTransferFeeChanged // Event containing the contract specifics and raw log

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
func (it *PublicLockTransferFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockTransferFeeChanged)
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
		it.Event = new(PublicLockTransferFeeChanged)
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
func (it *PublicLockTransferFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockTransferFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockTransferFeeChanged represents a TransferFeeChanged event raised by the PublicLock contract.
type PublicLockTransferFeeChanged struct {
	TransferFeeBasisPoints *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterTransferFeeChanged is a free log retrieval operation binding the contract event 0x0496ed1e61eb69727f9659a8e859288db4758ffb1f744d1c1424634f90a257f4.
//
// Solidity: event TransferFeeChanged(uint256 transferFeeBasisPoints)
func (_PublicLock *PublicLockFilterer) FilterTransferFeeChanged(opts *bind.FilterOpts) (*PublicLockTransferFeeChangedIterator, error) {

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "TransferFeeChanged")
	if err != nil {
		return nil, err
	}
	return &PublicLockTransferFeeChangedIterator{contract: _PublicLock.contract, event: "TransferFeeChanged", logs: logs, sub: sub}, nil
}

// WatchTransferFeeChanged is a free log subscription operation binding the contract event 0x0496ed1e61eb69727f9659a8e859288db4758ffb1f744d1c1424634f90a257f4.
//
// Solidity: event TransferFeeChanged(uint256 transferFeeBasisPoints)
func (_PublicLock *PublicLockFilterer) WatchTransferFeeChanged(opts *bind.WatchOpts, sink chan<- *PublicLockTransferFeeChanged) (event.Subscription, error) {

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "TransferFeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockTransferFeeChanged)
				if err := _PublicLock.contract.UnpackLog(event, "TransferFeeChanged", log); err != nil {
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

// ParseTransferFeeChanged is a log parse operation binding the contract event 0x0496ed1e61eb69727f9659a8e859288db4758ffb1f744d1c1424634f90a257f4.
//
// Solidity: event TransferFeeChanged(uint256 transferFeeBasisPoints)
func (_PublicLock *PublicLockFilterer) ParseTransferFeeChanged(log types.Log) (*PublicLockTransferFeeChanged, error) {
	event := new(PublicLockTransferFeeChanged)
	if err := _PublicLock.contract.UnpackLog(event, "TransferFeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockUnlockCallFailedIterator is returned from FilterUnlockCallFailed and is used to iterate over the raw logs and unpacked data for UnlockCallFailed events raised by the PublicLock contract.
type PublicLockUnlockCallFailedIterator struct {
	Event *PublicLockUnlockCallFailed // Event containing the contract specifics and raw log

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
func (it *PublicLockUnlockCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockUnlockCallFailed)
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
		it.Event = new(PublicLockUnlockCallFailed)
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
func (it *PublicLockUnlockCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockUnlockCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockUnlockCallFailed represents a UnlockCallFailed event raised by the PublicLock contract.
type PublicLockUnlockCallFailed struct {
	LockAddress   common.Address
	UnlockAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnlockCallFailed is a free log retrieval operation binding the contract event 0x6b18946261693dfd6c760d986b28ad2238b5b0267f8e5b6bc40a2f998e2f20ac.
//
// Solidity: event UnlockCallFailed(address indexed lockAddress, address unlockAddress)
func (_PublicLock *PublicLockFilterer) FilterUnlockCallFailed(opts *bind.FilterOpts, lockAddress []common.Address) (*PublicLockUnlockCallFailedIterator, error) {

	var lockAddressRule []interface{}
	for _, lockAddressItem := range lockAddress {
		lockAddressRule = append(lockAddressRule, lockAddressItem)
	}

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "UnlockCallFailed", lockAddressRule)
	if err != nil {
		return nil, err
	}
	return &PublicLockUnlockCallFailedIterator{contract: _PublicLock.contract, event: "UnlockCallFailed", logs: logs, sub: sub}, nil
}

// WatchUnlockCallFailed is a free log subscription operation binding the contract event 0x6b18946261693dfd6c760d986b28ad2238b5b0267f8e5b6bc40a2f998e2f20ac.
//
// Solidity: event UnlockCallFailed(address indexed lockAddress, address unlockAddress)
func (_PublicLock *PublicLockFilterer) WatchUnlockCallFailed(opts *bind.WatchOpts, sink chan<- *PublicLockUnlockCallFailed, lockAddress []common.Address) (event.Subscription, error) {

	var lockAddressRule []interface{}
	for _, lockAddressItem := range lockAddress {
		lockAddressRule = append(lockAddressRule, lockAddressItem)
	}

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "UnlockCallFailed", lockAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockUnlockCallFailed)
				if err := _PublicLock.contract.UnpackLog(event, "UnlockCallFailed", log); err != nil {
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

// ParseUnlockCallFailed is a log parse operation binding the contract event 0x6b18946261693dfd6c760d986b28ad2238b5b0267f8e5b6bc40a2f998e2f20ac.
//
// Solidity: event UnlockCallFailed(address indexed lockAddress, address unlockAddress)
func (_PublicLock *PublicLockFilterer) ParseUnlockCallFailed(log types.Log) (*PublicLockUnlockCallFailed, error) {
	event := new(PublicLockUnlockCallFailed)
	if err := _PublicLock.contract.UnpackLog(event, "UnlockCallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicLockWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the PublicLock contract.
type PublicLockWithdrawalIterator struct {
	Event *PublicLockWithdrawal // Event containing the contract specifics and raw log

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
func (it *PublicLockWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicLockWithdrawal)
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
		it.Event = new(PublicLockWithdrawal)
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
func (it *PublicLockWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicLockWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicLockWithdrawal represents a Withdrawal event raised by the PublicLock contract.
type PublicLockWithdrawal struct {
	Sender       common.Address
	TokenAddress common.Address
	Recipient    common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x342e7ff505a8a0364cd0dc2ff195c315e43bce86b204846ecd36913e117b109e.
//
// Solidity: event Withdrawal(address indexed sender, address indexed tokenAddress, address indexed recipient, uint256 amount)
func (_PublicLock *PublicLockFilterer) FilterWithdrawal(opts *bind.FilterOpts, sender []common.Address, tokenAddress []common.Address, recipient []common.Address) (*PublicLockWithdrawalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _PublicLock.contract.FilterLogs(opts, "Withdrawal", senderRule, tokenAddressRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &PublicLockWithdrawalIterator{contract: _PublicLock.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x342e7ff505a8a0364cd0dc2ff195c315e43bce86b204846ecd36913e117b109e.
//
// Solidity: event Withdrawal(address indexed sender, address indexed tokenAddress, address indexed recipient, uint256 amount)
func (_PublicLock *PublicLockFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *PublicLockWithdrawal, sender []common.Address, tokenAddress []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _PublicLock.contract.WatchLogs(opts, "Withdrawal", senderRule, tokenAddressRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicLockWithdrawal)
				if err := _PublicLock.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x342e7ff505a8a0364cd0dc2ff195c315e43bce86b204846ecd36913e117b109e.
//
// Solidity: event Withdrawal(address indexed sender, address indexed tokenAddress, address indexed recipient, uint256 amount)
func (_PublicLock *PublicLockFilterer) ParseWithdrawal(log types.Log) (*PublicLockWithdrawal, error) {
	event := new(PublicLockWithdrawal)
	if err := _PublicLock.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
