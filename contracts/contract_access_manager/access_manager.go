// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_access_manager

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

// AccessManagerMetaData contains all meta data concerning the AccessManager contract.
var AccessManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ensRegistry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ensName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"admin\",\"type\":\"bytes32\"}],\"name\":\"setNewRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040523060a0523480156200001557600080fd5b5060405162001f0638038062001f06833981016040819052620000389162000105565b6001600160a01b038116608052600054610100900460ff16806200005f575060005460ff16155b620000c75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b600054610100900460ff16158015620000ea576000805461ffff19166101011790555b8015620000fd576000805461ff00191690555b505062000137565b6000602082840312156200011857600080fd5b81516001600160a01b03811681146200013057600080fd5b9392505050565b60805160a051611d8d6200017960003960008181610535015281816105ba015281816106b10152610736015260008181610268015261115f0152611d8d6000f3fe6080604052600436106100e85760003560e01c806354fd4d501161008a578063a73494ea11610059578063a73494ea146102f3578063ac9650d814610313578063c4d66de814610340578063d547741f1461036057600080fd5b806354fd4d50146101f5578063572b6c051461024b57806391d1485414610298578063a217fddf146102de57600080fd5b80633121db1c116100c65780633121db1c1461018257806336568abe146101a25780633659cfe6146101c25780634f1ef286146101e257600080fd5b806301ffc9a7146100ed578063248a9ca3146101225780632f2ff15d14610160575b600080fd5b3480156100f957600080fd5b5061010d610108366004611755565b610380565b60405190151581526020015b60405180910390f35b34801561012e57600080fd5b5061015261013d366004611797565b60009081526065602052604090206001015490565b604051908152602001610119565b34801561016c57600080fd5b5061018061017b3660046117c5565b610419565b005b34801561018e57600080fd5b5061018061019d3660046117f5565b61044b565b3480156101ae57600080fd5b506101806101bd3660046117c5565b610489565b3480156101ce57600080fd5b506101806101dd36600461187a565b61052a565b6101806101f03660046118ad565b6106a6565b34801561020157600080fd5b5061023e6040518060400160405280600581526020017f302e312e3000000000000000000000000000000000000000000000000000000081525081565b60405161011991906119c9565b34801561025757600080fd5b5061010d61026636600461187a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0390811691161490565b3480156102a457600080fd5b5061010d6102b33660046117c5565b60009182526065602090815260408084206001600160a01b0393909316845291905290205460ff1690565b3480156102ea57600080fd5b50610152600081565b3480156102ff57600080fd5b5061018061030e3660046119dc565b61080f565b34801561031f57600080fd5b5061033361032e3660046119fe565b610827565b6040516101199190611a73565b34801561034c57600080fd5b5061018061035b36600461187a565b61091c565b34801561036c57600080fd5b5061018061037b3660046117c5565b610a5a565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061041357507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b60008281526065602052604090206001015461043c81610437610a82565b610a91565b6104468383610b11565b505050565b7f664245c7af190fec316596e8231f724e8171b1966cfcd124347ac5a66c34f64a61047881610437610a82565b610483848484610bb4565b50505050565b610491610a82565b6001600160a01b0316816001600160a01b03161461051c5760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201527f20726f6c657320666f722073656c66000000000000000000000000000000000060648201526084015b60405180910390fd5b6105268282610ce3565b5050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156105b85760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b6064820152608401610513565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166106137f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b03161461067e5760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b6163746976652070726f787960a01b6064820152608401610513565b61068781610d84565b604080516000808252602082019092526106a391839190610db1565b50565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156107345760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b6064820152608401610513565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661078f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b0316146107fa5760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b6163746976652070726f787960a01b6064820152608401610513565b61080382610d84565b61052682826001610db1565b600061081d81610437610a82565b6104468383610f75565b60608167ffffffffffffffff81111561084257610842611897565b60405190808252806020026020018201604052801561087557816020015b60608152602001906001900390816108605790505b50905060005b82811015610915576108e53085858481811061089957610899611ad5565b90506020028101906108ab9190611aeb565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610fc092505050565b8282815181106108f7576108f7611ad5565b6020026020010181905250808061090d90611b4f565b91505061087b565b5092915050565b600054610100900460ff1680610935575060005460ff16155b6109985760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610513565b600054610100900460ff161580156109ba576000805461ffff19166101011790555b6001600160a01b038216610a2a576040517feac0d38900000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5f5f61646d696e000000000000000000000000000000000000000000000000006044820152606401610513565b610a32610fec565b610a3a6110b6565b610a45600083610b11565b8015610526576000805461ff00191690555050565b600082815260656020526040902060010154610a7881610437610a82565b6104468383610ce3565b6000610a8c61115b565b905090565b60008281526065602090815260408083206001600160a01b038516845290915290205460ff1661052657610acf816001600160a01b031660146111be565b610ada8360206111be565b604051602001610aeb929190611b6a565b60408051601f198184030181529082905262461bcd60e51b8252610513916004016119c9565b60008281526065602090815260408083206001600160a01b038516845290915290205460ff166105265760008281526065602090815260408083206001600160a01b03851684529091529020805460ff19166001179055610b70610a82565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6040517f02571be30000000000000000000000000000000000000000000000000000000081527f91d1777781884d03a6757a803996e38de2a42967fb37eeaca72729271025a9e260048201526001600160a01b038416906302571be39060240160206040518083038186803b158015610c2c57600080fd5b505afa158015610c40573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c649190611beb565b6001600160a01b031663c47f002783836040518363ffffffff1660e01b8152600401610c91929190611c08565b602060405180830381600087803b158015610cab57600080fd5b505af1158015610cbf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104839190611c37565b60008281526065602090815260408083206001600160a01b038516845290915290205460ff16156105265760008281526065602090815260408083206001600160a01b03851684529091529020805460ff19169055610d40610a82565b6001600160a01b0316816001600160a01b0316837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45050565b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e361052681610437610a82565b6000610de47f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b9050610def8461139f565b600083511180610dfc5750815b15610e0d57610e0b848461146c565b505b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd9143805460ff16610f6e57805460ff191660011781556040516001600160a01b0383166024820152610eba90869060440160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f3659cfe60000000000000000000000000000000000000000000000000000000017905261146c565b50805460ff191681557f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b03838116911614610f655760405162461bcd60e51b815260206004820152602f60248201527f45524331393637557067726164653a207570677261646520627265616b73206660448201527f75727468657220757067726164657300000000000000000000000000000000006064820152608401610513565b610f6e85611557565b5050505050565b600082815260656020526040808220600101805490849055905190918391839186917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a4505050565b6060610fe58383604051806060016040528060278152602001611d3160279139611597565b9392505050565b600054610100900460ff1680611005575060005460ff16155b6110685760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610513565b600054610100900460ff1615801561108a576000805461ffff19166101011790555b61109261166b565b61109a61166b565b6110a261166b565b80156106a3576000805461ff001916905550565b600054610100900460ff16806110cf575060005460ff16155b6111325760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610513565b600054610100900460ff16158015611092576000805461ffff191661010117905561109a61166b565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163314156111b95760003661119c601482611c50565b6111a892369290611c67565b6111b191611c91565b60601c905090565b503390565b606060006111cd836002611cc6565b6111d8906002611ce5565b67ffffffffffffffff8111156111f0576111f0611897565b6040519080825280601f01601f19166020018201604052801561121a576020820181803683370190505b5090507f30000000000000000000000000000000000000000000000000000000000000008160008151811061125157611251611ad5565b60200101906001600160f81b031916908160001a9053507f78000000000000000000000000000000000000000000000000000000000000008160018151811061129c5761129c611ad5565b60200101906001600160f81b031916908160001a90535060006112c0846002611cc6565b6112cb906001611ce5565b90505b6001811115611350577f303132333435363738396162636465660000000000000000000000000000000085600f166010811061130c5761130c611ad5565b1a60f81b82828151811061132257611322611ad5565b60200101906001600160f81b031916908160001a90535060049490941c9361134981611cfd565b90506112ce565b508315610fe55760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610513565b803b6114135760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e7472616374000000000000000000000000000000000000006064820152608401610513565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060823b6114cb5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608401610513565b600080846001600160a01b0316846040516114e69190611d14565b600060405180830381855af49150503d8060008114611521576040519150601f19603f3d011682016040523d82523d6000602084013e611526565b606091505b509150915061154e8282604051806060016040528060278152602001611d316027913961171c565b95945050505050565b6115608161139f565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060833b6115f65760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608401610513565b600080856001600160a01b0316856040516116119190611d14565b600060405180830381855af49150503d806000811461164c576040519150601f19603f3d011682016040523d82523d6000602084013e611651565b606091505b509150915061166182828661171c565b9695505050505050565b600054610100900460ff1680611684575060005460ff16155b6116e75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610513565b600054610100900460ff161580156110a2576000805461ffff191661010117905580156106a3576000805461ff001916905550565b6060831561172b575081610fe5565b82511561173b5782518084602001fd5b8160405162461bcd60e51b815260040161051391906119c9565b60006020828403121561176757600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610fe557600080fd5b6000602082840312156117a957600080fd5b5035919050565b6001600160a01b03811681146106a357600080fd5b600080604083850312156117d857600080fd5b8235915060208301356117ea816117b0565b809150509250929050565b60008060006040848603121561180a57600080fd5b8335611815816117b0565b9250602084013567ffffffffffffffff8082111561183257600080fd5b818601915086601f83011261184657600080fd5b81358181111561185557600080fd5b87602082850101111561186757600080fd5b6020830194508093505050509250925092565b60006020828403121561188c57600080fd5b8135610fe5816117b0565b634e487b7160e01b600052604160045260246000fd5b600080604083850312156118c057600080fd5b82356118cb816117b0565b9150602083013567ffffffffffffffff808211156118e857600080fd5b818501915085601f8301126118fc57600080fd5b81358181111561190e5761190e611897565b604051601f8201601f19908116603f0116810190838211818310171561193657611936611897565b8160405282815288602084870101111561194f57600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b60005b8381101561198c578181015183820152602001611974565b838111156104835750506000910152565b600081518084526119b5816020860160208601611971565b601f01601f19169290920160200192915050565b602081526000610fe5602083018461199d565b600080604083850312156119ef57600080fd5b50508035926020909101359150565b60008060208385031215611a1157600080fd5b823567ffffffffffffffff80821115611a2957600080fd5b818501915085601f830112611a3d57600080fd5b813581811115611a4c57600080fd5b8660208260051b8501011115611a6157600080fd5b60209290920196919550909350505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015611ac857603f19888603018452611ab685835161199d565b94509285019290850190600101611a9a565b5092979650505050505050565b634e487b7160e01b600052603260045260246000fd5b6000808335601e19843603018112611b0257600080fd5b83018035915067ffffffffffffffff821115611b1d57600080fd5b602001915036819003821315611b3257600080fd5b9250929050565b634e487b7160e01b600052601160045260246000fd5b6000600019821415611b6357611b63611b39565b5060010190565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351611ba2816017850160208801611971565b7f206973206d697373696e6720726f6c65200000000000000000000000000000006017918401918201528351611bdf816028840160208801611971565b01602801949350505050565b600060208284031215611bfd57600080fd5b8151610fe5816117b0565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b600060208284031215611c4957600080fd5b5051919050565b600082821015611c6257611c62611b39565b500390565b60008085851115611c7757600080fd5b83861115611c8457600080fd5b5050820193919092039150565b6bffffffffffffffffffffffff198135818116916014851015611cbe5780818660140360031b1b83161692505b505092915050565b6000816000190483118215151615611ce057611ce0611b39565b500290565b60008219821115611cf857611cf8611b39565b500190565b600081611d0c57611d0c611b39565b506000190190565b60008251611d26818460208701611971565b919091019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220c97dce1ef1afd17a5065467c811432359ed2d9d1536e7046d261985765c1442e64736f6c63430008090033",
}

// AccessManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessManagerMetaData.ABI instead.
var AccessManagerABI = AccessManagerMetaData.ABI

// AccessManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccessManagerMetaData.Bin instead.
var AccessManagerBin = AccessManagerMetaData.Bin

// DeployAccessManager deploys a new Ethereum contract, binding an instance of AccessManager to it.
func DeployAccessManager(auth *bind.TransactOpts, backend bind.ContractBackend, forwarder common.Address) (common.Address, *types.Transaction, *AccessManager, error) {
	parsed, err := AccessManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccessManagerBin), backend, forwarder)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccessManager{AccessManagerCaller: AccessManagerCaller{contract: contract}, AccessManagerTransactor: AccessManagerTransactor{contract: contract}, AccessManagerFilterer: AccessManagerFilterer{contract: contract}}, nil
}

// AccessManager is an auto generated Go binding around an Ethereum contract.
type AccessManager struct {
	AccessManagerCaller     // Read-only binding to the contract
	AccessManagerTransactor // Write-only binding to the contract
	AccessManagerFilterer   // Log filterer for contract events
}

// AccessManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessManagerSession struct {
	Contract     *AccessManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessManagerCallerSession struct {
	Contract *AccessManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AccessManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessManagerTransactorSession struct {
	Contract     *AccessManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccessManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessManagerRaw struct {
	Contract *AccessManager // Generic contract binding to access the raw methods on
}

// AccessManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessManagerCallerRaw struct {
	Contract *AccessManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AccessManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessManagerTransactorRaw struct {
	Contract *AccessManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessManager creates a new instance of AccessManager, bound to a specific deployed contract.
func NewAccessManager(address common.Address, backend bind.ContractBackend) (*AccessManager, error) {
	contract, err := bindAccessManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessManager{AccessManagerCaller: AccessManagerCaller{contract: contract}, AccessManagerTransactor: AccessManagerTransactor{contract: contract}, AccessManagerFilterer: AccessManagerFilterer{contract: contract}}, nil
}

// NewAccessManagerCaller creates a new read-only instance of AccessManager, bound to a specific deployed contract.
func NewAccessManagerCaller(address common.Address, caller bind.ContractCaller) (*AccessManagerCaller, error) {
	contract, err := bindAccessManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessManagerCaller{contract: contract}, nil
}

// NewAccessManagerTransactor creates a new write-only instance of AccessManager, bound to a specific deployed contract.
func NewAccessManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessManagerTransactor, error) {
	contract, err := bindAccessManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessManagerTransactor{contract: contract}, nil
}

// NewAccessManagerFilterer creates a new log filterer instance of AccessManager, bound to a specific deployed contract.
func NewAccessManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessManagerFilterer, error) {
	contract, err := bindAccessManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessManagerFilterer{contract: contract}, nil
}

// bindAccessManager binds a generic wrapper to an already deployed contract.
func bindAccessManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessManager *AccessManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessManager.Contract.AccessManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessManager *AccessManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessManager.Contract.AccessManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessManager *AccessManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessManager.Contract.AccessManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessManager *AccessManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessManager *AccessManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessManager *AccessManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessManager *AccessManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessManager *AccessManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessManager.Contract.DEFAULTADMINROLE(&_AccessManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessManager *AccessManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessManager.Contract.DEFAULTADMINROLE(&_AccessManager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessManager *AccessManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AccessManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessManager *AccessManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessManager.Contract.GetRoleAdmin(&_AccessManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessManager *AccessManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessManager.Contract.GetRoleAdmin(&_AccessManager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessManager *AccessManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessManager *AccessManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessManager.Contract.HasRole(&_AccessManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessManager *AccessManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessManager.Contract.HasRole(&_AccessManager.CallOpts, role, account)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_AccessManager *AccessManagerCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _AccessManager.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_AccessManager *AccessManagerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _AccessManager.Contract.IsTrustedForwarder(&_AccessManager.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_AccessManager *AccessManagerCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _AccessManager.Contract.IsTrustedForwarder(&_AccessManager.CallOpts, forwarder)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessManager *AccessManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AccessManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessManager *AccessManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessManager.Contract.SupportsInterface(&_AccessManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessManager *AccessManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessManager.Contract.SupportsInterface(&_AccessManager.CallOpts, interfaceId)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_AccessManager *AccessManagerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AccessManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_AccessManager *AccessManagerSession) Version() (string, error) {
	return _AccessManager.Contract.Version(&_AccessManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_AccessManager *AccessManagerCallerSession) Version() (string, error) {
	return _AccessManager.Contract.Version(&_AccessManager.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessManager *AccessManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessManager *AccessManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessManager.Contract.GrantRole(&_AccessManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessManager *AccessManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessManager.Contract.GrantRole(&_AccessManager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address __admin) returns()
func (_AccessManager *AccessManagerTransactor) Initialize(opts *bind.TransactOpts, __admin common.Address) (*types.Transaction, error) {
	return _AccessManager.contract.Transact(opts, "initialize", __admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address __admin) returns()
func (_AccessManager *AccessManagerSession) Initialize(__admin common.Address) (*types.Transaction, error) {
	return _AccessManager.Contract.Initialize(&_AccessManager.TransactOpts, __admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address __admin) returns()
func (_AccessManager *AccessManagerTransactorSession) Initialize(__admin common.Address) (*types.Transaction, error) {
	return _AccessManager.Contract.Initialize(&_AccessManager.TransactOpts, __admin)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_AccessManager *AccessManagerTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _AccessManager.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_AccessManager *AccessManagerSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _AccessManager.Contract.Multicall(&_AccessManager.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_AccessManager *AccessManagerTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _AccessManager.Contract.Multicall(&_AccessManager.TransactOpts, data)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessManager *AccessManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessManager.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessManager *AccessManagerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessManager.Contract.RenounceRole(&_AccessManager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessManager *AccessManagerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessManager.Contract.RenounceRole(&_AccessManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessManager *AccessManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessManager *AccessManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessManager.Contract.RevokeRole(&_AccessManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessManager *AccessManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessManager.Contract.RevokeRole(&_AccessManager.TransactOpts, role, account)
}

// SetName is a paid mutator transaction binding the contract method 0x3121db1c.
//
// Solidity: function setName(address ensRegistry, string ensName) returns()
func (_AccessManager *AccessManagerTransactor) SetName(opts *bind.TransactOpts, ensRegistry common.Address, ensName string) (*types.Transaction, error) {
	return _AccessManager.contract.Transact(opts, "setName", ensRegistry, ensName)
}

// SetName is a paid mutator transaction binding the contract method 0x3121db1c.
//
// Solidity: function setName(address ensRegistry, string ensName) returns()
func (_AccessManager *AccessManagerSession) SetName(ensRegistry common.Address, ensName string) (*types.Transaction, error) {
	return _AccessManager.Contract.SetName(&_AccessManager.TransactOpts, ensRegistry, ensName)
}

// SetName is a paid mutator transaction binding the contract method 0x3121db1c.
//
// Solidity: function setName(address ensRegistry, string ensName) returns()
func (_AccessManager *AccessManagerTransactorSession) SetName(ensRegistry common.Address, ensName string) (*types.Transaction, error) {
	return _AccessManager.Contract.SetName(&_AccessManager.TransactOpts, ensRegistry, ensName)
}

// SetNewRole is a paid mutator transaction binding the contract method 0xa73494ea.
//
// Solidity: function setNewRole(bytes32 role, bytes32 admin) returns()
func (_AccessManager *AccessManagerTransactor) SetNewRole(opts *bind.TransactOpts, role [32]byte, admin [32]byte) (*types.Transaction, error) {
	return _AccessManager.contract.Transact(opts, "setNewRole", role, admin)
}

// SetNewRole is a paid mutator transaction binding the contract method 0xa73494ea.
//
// Solidity: function setNewRole(bytes32 role, bytes32 admin) returns()
func (_AccessManager *AccessManagerSession) SetNewRole(role [32]byte, admin [32]byte) (*types.Transaction, error) {
	return _AccessManager.Contract.SetNewRole(&_AccessManager.TransactOpts, role, admin)
}

// SetNewRole is a paid mutator transaction binding the contract method 0xa73494ea.
//
// Solidity: function setNewRole(bytes32 role, bytes32 admin) returns()
func (_AccessManager *AccessManagerTransactorSession) SetNewRole(role [32]byte, admin [32]byte) (*types.Transaction, error) {
	return _AccessManager.Contract.SetNewRole(&_AccessManager.TransactOpts, role, admin)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AccessManager *AccessManagerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _AccessManager.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AccessManager *AccessManagerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AccessManager.Contract.UpgradeTo(&_AccessManager.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AccessManager *AccessManagerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AccessManager.Contract.UpgradeTo(&_AccessManager.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AccessManager *AccessManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AccessManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AccessManager *AccessManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AccessManager.Contract.UpgradeToAndCall(&_AccessManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AccessManager *AccessManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AccessManager.Contract.UpgradeToAndCall(&_AccessManager.TransactOpts, newImplementation, data)
}

// AccessManagerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the AccessManager contract.
type AccessManagerAdminChangedIterator struct {
	Event *AccessManagerAdminChanged // Event containing the contract specifics and raw log

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
func (it *AccessManagerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessManagerAdminChanged)
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
		it.Event = new(AccessManagerAdminChanged)
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
func (it *AccessManagerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessManagerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessManagerAdminChanged represents a AdminChanged event raised by the AccessManager contract.
type AccessManagerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AccessManager *AccessManagerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AccessManagerAdminChangedIterator, error) {

	logs, sub, err := _AccessManager.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AccessManagerAdminChangedIterator{contract: _AccessManager.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AccessManager *AccessManagerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AccessManagerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _AccessManager.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessManagerAdminChanged)
				if err := _AccessManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_AccessManager *AccessManagerFilterer) ParseAdminChanged(log types.Log) (*AccessManagerAdminChanged, error) {
	event := new(AccessManagerAdminChanged)
	if err := _AccessManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessManagerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the AccessManager contract.
type AccessManagerBeaconUpgradedIterator struct {
	Event *AccessManagerBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *AccessManagerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessManagerBeaconUpgraded)
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
		it.Event = new(AccessManagerBeaconUpgraded)
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
func (it *AccessManagerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessManagerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessManagerBeaconUpgraded represents a BeaconUpgraded event raised by the AccessManager contract.
type AccessManagerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_AccessManager *AccessManagerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*AccessManagerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _AccessManager.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &AccessManagerBeaconUpgradedIterator{contract: _AccessManager.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_AccessManager *AccessManagerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *AccessManagerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _AccessManager.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessManagerBeaconUpgraded)
				if err := _AccessManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_AccessManager *AccessManagerFilterer) ParseBeaconUpgraded(log types.Log) (*AccessManagerBeaconUpgraded, error) {
	event := new(AccessManagerBeaconUpgraded)
	if err := _AccessManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AccessManager contract.
type AccessManagerRoleAdminChangedIterator struct {
	Event *AccessManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AccessManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessManagerRoleAdminChanged)
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
		it.Event = new(AccessManagerRoleAdminChanged)
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
func (it *AccessManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessManagerRoleAdminChanged represents a RoleAdminChanged event raised by the AccessManager contract.
type AccessManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessManager *AccessManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AccessManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _AccessManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AccessManagerRoleAdminChangedIterator{contract: _AccessManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessManager *AccessManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AccessManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AccessManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessManagerRoleAdminChanged)
				if err := _AccessManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AccessManager *AccessManagerFilterer) ParseRoleAdminChanged(log types.Log) (*AccessManagerRoleAdminChanged, error) {
	event := new(AccessManagerRoleAdminChanged)
	if err := _AccessManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AccessManager contract.
type AccessManagerRoleGrantedIterator struct {
	Event *AccessManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *AccessManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessManagerRoleGranted)
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
		it.Event = new(AccessManagerRoleGranted)
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
func (it *AccessManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessManagerRoleGranted represents a RoleGranted event raised by the AccessManager contract.
type AccessManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessManager *AccessManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _AccessManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessManagerRoleGrantedIterator{contract: _AccessManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessManager *AccessManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AccessManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccessManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessManagerRoleGranted)
				if err := _AccessManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AccessManager *AccessManagerFilterer) ParseRoleGranted(log types.Log) (*AccessManagerRoleGranted, error) {
	event := new(AccessManagerRoleGranted)
	if err := _AccessManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AccessManager contract.
type AccessManagerRoleRevokedIterator struct {
	Event *AccessManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AccessManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessManagerRoleRevoked)
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
		it.Event = new(AccessManagerRoleRevoked)
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
func (it *AccessManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessManagerRoleRevoked represents a RoleRevoked event raised by the AccessManager contract.
type AccessManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessManager *AccessManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _AccessManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessManagerRoleRevokedIterator{contract: _AccessManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessManager *AccessManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AccessManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccessManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessManagerRoleRevoked)
				if err := _AccessManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AccessManager *AccessManagerFilterer) ParseRoleRevoked(log types.Log) (*AccessManagerRoleRevoked, error) {
	event := new(AccessManagerRoleRevoked)
	if err := _AccessManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AccessManager contract.
type AccessManagerUpgradedIterator struct {
	Event *AccessManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *AccessManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessManagerUpgraded)
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
		it.Event = new(AccessManagerUpgraded)
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
func (it *AccessManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessManagerUpgraded represents a Upgraded event raised by the AccessManager contract.
type AccessManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AccessManager *AccessManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AccessManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AccessManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AccessManagerUpgradedIterator{contract: _AccessManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AccessManager *AccessManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AccessManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AccessManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessManagerUpgraded)
				if err := _AccessManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_AccessManager *AccessManagerFilterer) ParseUpgraded(log types.Log) (*AccessManagerUpgraded, error) {
	event := new(AccessManagerUpgraded)
	if err := _AccessManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
