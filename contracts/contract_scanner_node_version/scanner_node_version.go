// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_scanner_node_version

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

// ScannerNodeVersionMetaData contains all meta data concerning the ScannerNodeVersion contract.
var ScannerNodeVersionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"AccessManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newVersion\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldVersion\",\"type\":\"string\"}],\"name\":\"ScannerNodeBetaVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newVersion\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldVersion\",\"type\":\"string\"}],\"name\":\"ScannerNodeVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__router\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scannerNodeBetaVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scannerNodeVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"setAccessManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ensRegistry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ensName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"setScannerNodeBetaVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"setScannerNodeVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040523060a0523480156200001557600080fd5b50604051620021ee380380620021ee833981016040819052620000389162000105565b6001600160a01b038116608052600054610100900460ff16806200005f575060005460ff16155b620000c75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b600054610100900460ff16158015620000ea576000805461ffff19166101011790555b8015620000fd576000805461ff00191690555b505062000137565b6000602082840312156200011857600080fd5b81516001600160a01b03811681146200013057600080fd5b9392505050565b60805160a0516120756200017960003960008181610420015281816104a50152818161066b01526106f00152600081816101f201526113b301526120756000f3fe6080604052600436106100d25760003560e01c806354fd4d501161007f578063ac9650d811610059578063ac9650d814610252578063c0d786551461027f578063c4d2b6dd1461029f578063c9580804146102bf57600080fd5b806354fd4d501461018c578063572b6c05146101d557806394bb55a21461023257600080fd5b80633659cfe6116100b05780633659cfe614610139578063485cc955146101595780634f1ef2861461017957600080fd5b80633121db1c146100d757806331d66531146100f9578063345db3e114610124575b600080fd5b3480156100e357600080fd5b506100f76100f23660046119fe565b6102df565b005b34801561010557600080fd5b5061010e610378565b60405161011b9190611aab565b60405180910390f35b34801561013057600080fd5b5061010e610407565b34801561014557600080fd5b506100f7610154366004611abe565b610415565b34801561016557600080fd5b506100f7610174366004611adb565b610591565b6100f7610187366004611b2a565b610660565b34801561019857600080fd5b5061010e6040518060400160405280600581526020017f302e312e3100000000000000000000000000000000000000000000000000000081525081565b3480156101e157600080fd5b506102226101f0366004611abe565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0390811691161490565b604051901515815260200161011b565b34801561023e57600080fd5b506100f761024d366004611bee565b6107cd565b34801561025e57600080fd5b5061027261026d366004611c30565b6108c2565b60405161011b9190611ca5565b34801561028b57600080fd5b506100f761029a366004611abe565b6109b8565b3480156102ab57600080fd5b506100f76102ba366004611bee565b610a75565b3480156102cb57600080fd5b506100f76102da366004611abe565b610b6a565b7f664245c7af190fec316596e8231f724e8171b1966cfcd124347ac5a66c34f64a6103118161030c610c28565b610c37565b610367578061031e610c28565b6040517f75000dc000000000000000000000000000000000000000000000000000000000815260048101929092526001600160a01b031660248201526044015b60405180910390fd5b610372848484610cdc565b50505050565b61012e805461038690611d07565b80601f01602080910402602001604051908101604052809291908181526020018280546103b290611d07565b80156103ff5780601f106103d4576101008083540402835291602001916103ff565b820191906000526020600020905b8154815290600101906020018083116103e257829003601f168201915b505050505081565b61012d805461038690611d07565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156104a35760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b606482015260840161035e565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166104fe7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b0316146105695760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b6163746976652070726f787960a01b606482015260840161035e565b61057281610e0b565b6040805160008082526020820190925261058e91839190610e45565b50565b600054610100900460ff16806105aa575060005460ff16155b61060d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161035e565b600054610100900460ff1615801561062f576000805461ffff19166101011790555b61063883611009565b6106418261115b565b6106496112ac565b801561065b576000805461ff00191690555b505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156106ee5760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b606482015260840161035e565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166107497f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b0316146107b45760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b6163746976652070726f787960a01b606482015260840161035e565b6107bd82610e0b565b6107c982826001610e45565b5050565b7f93bf097503ee765c5402c05999a7d54bac82299bf183ba1f5f2681ab2c6bd70c6107fa8161030c610c28565b610807578061031e610c28565b828260405160200161081a929190611d42565b6040516020818303038152906040528051906020012061012d6040516020016108439190611d52565b6040516020818303038152906040528051906020012014156108785760405163ae01a30160e01b815260040160405180910390fd5b7faa8f23ad4857a5d22df0189ebc6150d51b3152d8c621c3e8d24c66387897819b838361012d6040516108ad93929190611ded565b60405180910390a161037261012d8484611907565b60608167ffffffffffffffff8111156108dd576108dd611b14565b60405190808252806020026020018201604052801561091057816020015b60608152602001906001900390816108fb5790505b50905060005b828110156109b0576109803085858481811061093457610934611e84565b90506020028101906109469190611e9a565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061136e92505050565b82828151811061099257610992611e84565b602002602001018190525080806109a890611ef7565b915050610916565b505b92915050565b60006109c68161030c610c28565b6109d3578061031e610c28565b6001600160a01b038216610a2a5760405163eac0d38960e01b815260206004820152600960248201527f6e6577526f757465720000000000000000000000000000000000000000000000604482015260640161035e565b606580546001600160a01b0319166001600160a01b0384169081179091556040517f7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc8090600090a25050565b7fcd2c75e410f265ca5b9c8bed70c109422b498bfc98f7df6ef360bcb407b1bedc610aa28161030c610c28565b610aaf578061031e610c28565b8282604051602001610ac2929190611d42565b6040516020818303038152906040528051906020012061012e604051602001610aeb9190611d52565b604051602081830303815290604052805190602001201415610b205760405163ae01a30160e01b815260040160405180910390fd5b7fe5ac27c9fd9eefb9d9fd2b0878923cf7d44e611ead66cf80b7ccf184ac2927c0838361012e604051610b5593929190611ded565b60405180910390a161037261012e8484611907565b6000610b788161030c610c28565b610b85578061031e610c28565b610b9f6001600160a01b038316637965db0b60e01b611393565b610bdd576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b604482015260640161035e565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a25050565b6000610c326113af565b905090565b6033546040517f91d14854000000000000000000000000000000000000000000000000000000008152600481018490526001600160a01b03838116602483015260009216906391d148549060440160206040518083038186803b158015610c9d57600080fd5b505afa158015610cb1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cd59190611f12565b9392505050565b6040517f02571be30000000000000000000000000000000000000000000000000000000081527f91d1777781884d03a6757a803996e38de2a42967fb37eeaca72729271025a9e260048201526001600160a01b038416906302571be39060240160206040518083038186803b158015610d5457600080fd5b505afa158015610d68573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d8c9190611f34565b6001600160a01b031663c47f002783836040518363ffffffff1660e01b8152600401610db9929190611f51565b602060405180830381600087803b158015610dd357600080fd5b505af1158015610de7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103729190611f6d565b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3610e388161030c610c28565b6107c9578061031e610c28565b6000610e787f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b9050610e8384611412565b600083511180610e905750815b15610ea157610e9f84846114c7565b505b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd9143805460ff1661100257805460ff191660011781556040516001600160a01b0383166024820152610f4e90869060440160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f3659cfe6000000000000000000000000000000000000000000000000000000001790526114c7565b50805460ff191681557f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b03838116911614610ff95760405162461bcd60e51b815260206004820152602f60248201527f45524331393637557067726164653a207570677261646520627265616b73206660448201527f7572746865722075706772616465730000000000000000000000000000000000606482015260840161035e565b611002856115b2565b5050505050565b600054610100900460ff1680611022575060005460ff16155b6110855760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161035e565b600054610100900460ff161580156110a7576000805461ffff19166101011790555b6110c16001600160a01b038316637965db0b60e01b611393565b6110ff576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b604482015260640161035e565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a280156107c9576000805461ff00191690555050565b600054610100900460ff1680611174575060005460ff16155b6111d75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161035e565b600054610100900460ff161580156111f9576000805461ffff19166101011790555b6001600160a01b0382166112505760405163eac0d38960e01b815260206004820152600660248201527f726f757465720000000000000000000000000000000000000000000000000000604482015260640161035e565b606580546001600160a01b0319166001600160a01b0384169081179091556040517f7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc8090600090a280156107c9576000805461ff00191690555050565b600054610100900460ff16806112c5575060005460ff16155b6113285760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161035e565b600054610100900460ff1615801561134a576000805461ffff19166101011790555b6113526115f2565b61135a6115f2565b801561058e576000805461ff001916905550565b6060610cd58383604051806060016040528060278152602001612019602791396116a3565b600061139e83611777565b8015610cd55750610cd583836117c2565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633141561140d576000366113f0601482611f86565b6113fc92369290611f9d565b61140591611fc7565b60601c905090565b503390565b803b6114865760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e747261637400000000000000000000000000000000000000606482015260840161035e565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b6060823b6115265760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b606482015260840161035e565b600080846001600160a01b0316846040516115419190611ffc565b600060405180830381855af49150503d806000811461157c576040519150601f19603f3d011682016040523d82523d6000602084013e611581565b606091505b50915091506115a98282604051806060016040528060278152602001612019602791396118ce565b95945050505050565b6115bb81611412565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b600054610100900460ff168061160b575060005460ff16155b61166e5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161035e565b600054610100900460ff1615801561135a576000805461ffff1916610101179055801561058e576000805461ff001916905550565b6060833b6117025760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b606482015260840161035e565b600080856001600160a01b03168560405161171d9190611ffc565b600060405180830381855af49150503d8060008114611758576040519150601f19603f3d011682016040523d82523d6000602084013e61175d565b606091505b509150915061176d8282866118ce565b9695505050505050565b600061178a826301ffc9a760e01b6117c2565b80156109b257506117bb827fffffffff000000000000000000000000000000000000000000000000000000006117c2565b1592915050565b604080517fffffffff00000000000000000000000000000000000000000000000000000000831660248083019190915282518083039091018152604490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166301ffc9a760e01b179052905160009190829081906001600160a01b0387169061753090611856908690611ffc565b6000604051808303818686fa925050503d8060008114611892576040519150601f19603f3d011682016040523d82523d6000602084013e611897565b606091505b50915091506020815110156118b257600093505050506109b2565b81801561176d57508080602001905181019061176d9190611f12565b606083156118dd575081610cd5565b8251156118ed5782518084602001fd5b8160405162461bcd60e51b815260040161035e9190611aab565b82805461191390611d07565b90600052602060002090601f016020900481019282611935576000855561197b565b82601f1061194e5782800160ff1982351617855561197b565b8280016001018555821561197b579182015b8281111561197b578235825591602001919060010190611960565b5061198792915061198b565b5090565b5b80821115611987576000815560010161198c565b6001600160a01b038116811461058e57600080fd5b60008083601f8401126119c757600080fd5b50813567ffffffffffffffff8111156119df57600080fd5b6020830191508360208285010111156119f757600080fd5b9250929050565b600080600060408486031215611a1357600080fd5b8335611a1e816119a0565b9250602084013567ffffffffffffffff811115611a3a57600080fd5b611a46868287016119b5565b9497909650939450505050565b60005b83811015611a6e578181015183820152602001611a56565b838111156103725750506000910152565b60008151808452611a97816020860160208601611a53565b601f01601f19169290920160200192915050565b602081526000610cd56020830184611a7f565b600060208284031215611ad057600080fd5b8135610cd5816119a0565b60008060408385031215611aee57600080fd5b8235611af9816119a0565b91506020830135611b09816119a0565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b60008060408385031215611b3d57600080fd5b8235611b48816119a0565b9150602083013567ffffffffffffffff80821115611b6557600080fd5b818501915085601f830112611b7957600080fd5b813581811115611b8b57611b8b611b14565b604051601f8201601f19908116603f01168101908382118183101715611bb357611bb3611b14565b81604052828152886020848701011115611bcc57600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b60008060208385031215611c0157600080fd5b823567ffffffffffffffff811115611c1857600080fd5b611c24858286016119b5565b90969095509350505050565b60008060208385031215611c4357600080fd5b823567ffffffffffffffff80821115611c5b57600080fd5b818501915085601f830112611c6f57600080fd5b813581811115611c7e57600080fd5b8660208260051b8501011115611c9357600080fd5b60209290920196919550909350505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015611cfa57603f19888603018452611ce8858351611a7f565b94509285019290850190600101611ccc565b5092979650505050505050565b600181811c90821680611d1b57607f821691505b60208210811415611d3c57634e487b7160e01b600052602260045260246000fd5b50919050565b8183823760009101908152919050565b6000808354611d6081611d07565b60018281168015611d785760018114611d8957611db8565b60ff19841687528287019450611db8565b8760005260208060002060005b85811015611daf5781548a820152908401908201611d96565b50505082870194505b50929695505050505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b604081526000611e01604083018587611dc4565b60208382038185015260008554611e1781611d07565b80855260018281168015611e325760018114611e4657611e74565b60ff19841687870152604087019450611e74565b896000528560002060005b84811015611e6c578154898201890152908301908701611e51565b880187019550505b50929a9950505050505050505050565b634e487b7160e01b600052603260045260246000fd5b6000808335601e19843603018112611eb157600080fd5b83018035915067ffffffffffffffff821115611ecc57600080fd5b6020019150368190038213156119f757600080fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415611f0b57611f0b611ee1565b5060010190565b600060208284031215611f2457600080fd5b81518015158114610cd557600080fd5b600060208284031215611f4657600080fd5b8151610cd5816119a0565b602081526000611f65602083018486611dc4565b949350505050565b600060208284031215611f7f57600080fd5b5051919050565b600082821015611f9857611f98611ee1565b500390565b60008085851115611fad57600080fd5b83861115611fba57600080fd5b5050820193919092039150565b6bffffffffffffffffffffffff198135818116916014851015611ff45780818660140360031b1b83161692505b505092915050565b6000825161200e818460208701611a53565b919091019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220f5b4e55643d7984446b1f7588e62740e6ebdd068e85d7dc1e04344536f42f73864736f6c63430008090033",
}

// ScannerNodeVersionABI is the input ABI used to generate the binding from.
// Deprecated: Use ScannerNodeVersionMetaData.ABI instead.
var ScannerNodeVersionABI = ScannerNodeVersionMetaData.ABI

// ScannerNodeVersionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ScannerNodeVersionMetaData.Bin instead.
var ScannerNodeVersionBin = ScannerNodeVersionMetaData.Bin

// DeployScannerNodeVersion deploys a new Ethereum contract, binding an instance of ScannerNodeVersion to it.
func DeployScannerNodeVersion(auth *bind.TransactOpts, backend bind.ContractBackend, forwarder common.Address) (common.Address, *types.Transaction, *ScannerNodeVersion, error) {
	parsed, err := ScannerNodeVersionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ScannerNodeVersionBin), backend, forwarder)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ScannerNodeVersion{ScannerNodeVersionCaller: ScannerNodeVersionCaller{contract: contract}, ScannerNodeVersionTransactor: ScannerNodeVersionTransactor{contract: contract}, ScannerNodeVersionFilterer: ScannerNodeVersionFilterer{contract: contract}}, nil
}

// ScannerNodeVersion is an auto generated Go binding around an Ethereum contract.
type ScannerNodeVersion struct {
	ScannerNodeVersionCaller     // Read-only binding to the contract
	ScannerNodeVersionTransactor // Write-only binding to the contract
	ScannerNodeVersionFilterer   // Log filterer for contract events
}

// ScannerNodeVersionCaller is an auto generated read-only Go binding around an Ethereum contract.
type ScannerNodeVersionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScannerNodeVersionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScannerNodeVersionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScannerNodeVersionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScannerNodeVersionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScannerNodeVersionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScannerNodeVersionSession struct {
	Contract     *ScannerNodeVersion // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ScannerNodeVersionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScannerNodeVersionCallerSession struct {
	Contract *ScannerNodeVersionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ScannerNodeVersionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScannerNodeVersionTransactorSession struct {
	Contract     *ScannerNodeVersionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ScannerNodeVersionRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScannerNodeVersionRaw struct {
	Contract *ScannerNodeVersion // Generic contract binding to access the raw methods on
}

// ScannerNodeVersionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScannerNodeVersionCallerRaw struct {
	Contract *ScannerNodeVersionCaller // Generic read-only contract binding to access the raw methods on
}

// ScannerNodeVersionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScannerNodeVersionTransactorRaw struct {
	Contract *ScannerNodeVersionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScannerNodeVersion creates a new instance of ScannerNodeVersion, bound to a specific deployed contract.
func NewScannerNodeVersion(address common.Address, backend bind.ContractBackend) (*ScannerNodeVersion, error) {
	contract, err := bindScannerNodeVersion(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ScannerNodeVersion{ScannerNodeVersionCaller: ScannerNodeVersionCaller{contract: contract}, ScannerNodeVersionTransactor: ScannerNodeVersionTransactor{contract: contract}, ScannerNodeVersionFilterer: ScannerNodeVersionFilterer{contract: contract}}, nil
}

// NewScannerNodeVersionCaller creates a new read-only instance of ScannerNodeVersion, bound to a specific deployed contract.
func NewScannerNodeVersionCaller(address common.Address, caller bind.ContractCaller) (*ScannerNodeVersionCaller, error) {
	contract, err := bindScannerNodeVersion(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScannerNodeVersionCaller{contract: contract}, nil
}

// NewScannerNodeVersionTransactor creates a new write-only instance of ScannerNodeVersion, bound to a specific deployed contract.
func NewScannerNodeVersionTransactor(address common.Address, transactor bind.ContractTransactor) (*ScannerNodeVersionTransactor, error) {
	contract, err := bindScannerNodeVersion(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScannerNodeVersionTransactor{contract: contract}, nil
}

// NewScannerNodeVersionFilterer creates a new log filterer instance of ScannerNodeVersion, bound to a specific deployed contract.
func NewScannerNodeVersionFilterer(address common.Address, filterer bind.ContractFilterer) (*ScannerNodeVersionFilterer, error) {
	contract, err := bindScannerNodeVersion(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScannerNodeVersionFilterer{contract: contract}, nil
}

// bindScannerNodeVersion binds a generic wrapper to an already deployed contract.
func bindScannerNodeVersion(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ScannerNodeVersionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScannerNodeVersion *ScannerNodeVersionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScannerNodeVersion.Contract.ScannerNodeVersionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScannerNodeVersion *ScannerNodeVersionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.ScannerNodeVersionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScannerNodeVersion *ScannerNodeVersionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.ScannerNodeVersionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScannerNodeVersion *ScannerNodeVersionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScannerNodeVersion.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScannerNodeVersion *ScannerNodeVersionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScannerNodeVersion *ScannerNodeVersionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.contract.Transact(opts, method, params...)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_ScannerNodeVersion *ScannerNodeVersionCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _ScannerNodeVersion.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_ScannerNodeVersion *ScannerNodeVersionSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _ScannerNodeVersion.Contract.IsTrustedForwarder(&_ScannerNodeVersion.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_ScannerNodeVersion *ScannerNodeVersionCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _ScannerNodeVersion.Contract.IsTrustedForwarder(&_ScannerNodeVersion.CallOpts, forwarder)
}

// ScannerNodeBetaVersion is a free data retrieval call binding the contract method 0x31d66531.
//
// Solidity: function scannerNodeBetaVersion() view returns(string)
func (_ScannerNodeVersion *ScannerNodeVersionCaller) ScannerNodeBetaVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ScannerNodeVersion.contract.Call(opts, &out, "scannerNodeBetaVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ScannerNodeBetaVersion is a free data retrieval call binding the contract method 0x31d66531.
//
// Solidity: function scannerNodeBetaVersion() view returns(string)
func (_ScannerNodeVersion *ScannerNodeVersionSession) ScannerNodeBetaVersion() (string, error) {
	return _ScannerNodeVersion.Contract.ScannerNodeBetaVersion(&_ScannerNodeVersion.CallOpts)
}

// ScannerNodeBetaVersion is a free data retrieval call binding the contract method 0x31d66531.
//
// Solidity: function scannerNodeBetaVersion() view returns(string)
func (_ScannerNodeVersion *ScannerNodeVersionCallerSession) ScannerNodeBetaVersion() (string, error) {
	return _ScannerNodeVersion.Contract.ScannerNodeBetaVersion(&_ScannerNodeVersion.CallOpts)
}

// ScannerNodeVersion is a free data retrieval call binding the contract method 0x345db3e1.
//
// Solidity: function scannerNodeVersion() view returns(string)
func (_ScannerNodeVersion *ScannerNodeVersionCaller) ScannerNodeVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ScannerNodeVersion.contract.Call(opts, &out, "scannerNodeVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ScannerNodeVersion is a free data retrieval call binding the contract method 0x345db3e1.
//
// Solidity: function scannerNodeVersion() view returns(string)
func (_ScannerNodeVersion *ScannerNodeVersionSession) ScannerNodeVersion() (string, error) {
	return _ScannerNodeVersion.Contract.ScannerNodeVersion(&_ScannerNodeVersion.CallOpts)
}

// ScannerNodeVersion is a free data retrieval call binding the contract method 0x345db3e1.
//
// Solidity: function scannerNodeVersion() view returns(string)
func (_ScannerNodeVersion *ScannerNodeVersionCallerSession) ScannerNodeVersion() (string, error) {
	return _ScannerNodeVersion.Contract.ScannerNodeVersion(&_ScannerNodeVersion.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ScannerNodeVersion *ScannerNodeVersionCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ScannerNodeVersion.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ScannerNodeVersion *ScannerNodeVersionSession) Version() (string, error) {
	return _ScannerNodeVersion.Contract.Version(&_ScannerNodeVersion.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ScannerNodeVersion *ScannerNodeVersionCallerSession) Version() (string, error) {
	return _ScannerNodeVersion.Contract.Version(&_ScannerNodeVersion.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address __manager, address __router) returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactor) Initialize(opts *bind.TransactOpts, __manager common.Address, __router common.Address) (*types.Transaction, error) {
	return _ScannerNodeVersion.contract.Transact(opts, "initialize", __manager, __router)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address __manager, address __router) returns()
func (_ScannerNodeVersion *ScannerNodeVersionSession) Initialize(__manager common.Address, __router common.Address) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.Initialize(&_ScannerNodeVersion.TransactOpts, __manager, __router)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address __manager, address __router) returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactorSession) Initialize(__manager common.Address, __router common.Address) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.Initialize(&_ScannerNodeVersion.TransactOpts, __manager, __router)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ScannerNodeVersion *ScannerNodeVersionTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _ScannerNodeVersion.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ScannerNodeVersion *ScannerNodeVersionSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.Multicall(&_ScannerNodeVersion.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ScannerNodeVersion *ScannerNodeVersionTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.Multicall(&_ScannerNodeVersion.TransactOpts, data)
}

// SetAccessManager is a paid mutator transaction binding the contract method 0xc9580804.
//
// Solidity: function setAccessManager(address newManager) returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactor) SetAccessManager(opts *bind.TransactOpts, newManager common.Address) (*types.Transaction, error) {
	return _ScannerNodeVersion.contract.Transact(opts, "setAccessManager", newManager)
}

// SetAccessManager is a paid mutator transaction binding the contract method 0xc9580804.
//
// Solidity: function setAccessManager(address newManager) returns()
func (_ScannerNodeVersion *ScannerNodeVersionSession) SetAccessManager(newManager common.Address) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.SetAccessManager(&_ScannerNodeVersion.TransactOpts, newManager)
}

// SetAccessManager is a paid mutator transaction binding the contract method 0xc9580804.
//
// Solidity: function setAccessManager(address newManager) returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactorSession) SetAccessManager(newManager common.Address) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.SetAccessManager(&_ScannerNodeVersion.TransactOpts, newManager)
}

// SetName is a paid mutator transaction binding the contract method 0x3121db1c.
//
// Solidity: function setName(address ensRegistry, string ensName) returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactor) SetName(opts *bind.TransactOpts, ensRegistry common.Address, ensName string) (*types.Transaction, error) {
	return _ScannerNodeVersion.contract.Transact(opts, "setName", ensRegistry, ensName)
}

// SetName is a paid mutator transaction binding the contract method 0x3121db1c.
//
// Solidity: function setName(address ensRegistry, string ensName) returns()
func (_ScannerNodeVersion *ScannerNodeVersionSession) SetName(ensRegistry common.Address, ensName string) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.SetName(&_ScannerNodeVersion.TransactOpts, ensRegistry, ensName)
}

// SetName is a paid mutator transaction binding the contract method 0x3121db1c.
//
// Solidity: function setName(address ensRegistry, string ensName) returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactorSession) SetName(ensRegistry common.Address, ensName string) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.SetName(&_ScannerNodeVersion.TransactOpts, ensRegistry, ensName)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address newRouter) returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _ScannerNodeVersion.contract.Transact(opts, "setRouter", newRouter)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address newRouter) returns()
func (_ScannerNodeVersion *ScannerNodeVersionSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.SetRouter(&_ScannerNodeVersion.TransactOpts, newRouter)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address newRouter) returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.SetRouter(&_ScannerNodeVersion.TransactOpts, newRouter)
}

// SetScannerNodeBetaVersion is a paid mutator transaction binding the contract method 0xc4d2b6dd.
//
// Solidity: function setScannerNodeBetaVersion(string _version) returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactor) SetScannerNodeBetaVersion(opts *bind.TransactOpts, _version string) (*types.Transaction, error) {
	return _ScannerNodeVersion.contract.Transact(opts, "setScannerNodeBetaVersion", _version)
}

// SetScannerNodeBetaVersion is a paid mutator transaction binding the contract method 0xc4d2b6dd.
//
// Solidity: function setScannerNodeBetaVersion(string _version) returns()
func (_ScannerNodeVersion *ScannerNodeVersionSession) SetScannerNodeBetaVersion(_version string) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.SetScannerNodeBetaVersion(&_ScannerNodeVersion.TransactOpts, _version)
}

// SetScannerNodeBetaVersion is a paid mutator transaction binding the contract method 0xc4d2b6dd.
//
// Solidity: function setScannerNodeBetaVersion(string _version) returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactorSession) SetScannerNodeBetaVersion(_version string) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.SetScannerNodeBetaVersion(&_ScannerNodeVersion.TransactOpts, _version)
}

// SetScannerNodeVersion is a paid mutator transaction binding the contract method 0x94bb55a2.
//
// Solidity: function setScannerNodeVersion(string _version) returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactor) SetScannerNodeVersion(opts *bind.TransactOpts, _version string) (*types.Transaction, error) {
	return _ScannerNodeVersion.contract.Transact(opts, "setScannerNodeVersion", _version)
}

// SetScannerNodeVersion is a paid mutator transaction binding the contract method 0x94bb55a2.
//
// Solidity: function setScannerNodeVersion(string _version) returns()
func (_ScannerNodeVersion *ScannerNodeVersionSession) SetScannerNodeVersion(_version string) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.SetScannerNodeVersion(&_ScannerNodeVersion.TransactOpts, _version)
}

// SetScannerNodeVersion is a paid mutator transaction binding the contract method 0x94bb55a2.
//
// Solidity: function setScannerNodeVersion(string _version) returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactorSession) SetScannerNodeVersion(_version string) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.SetScannerNodeVersion(&_ScannerNodeVersion.TransactOpts, _version)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _ScannerNodeVersion.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ScannerNodeVersion *ScannerNodeVersionSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.UpgradeTo(&_ScannerNodeVersion.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.UpgradeTo(&_ScannerNodeVersion.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ScannerNodeVersion.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ScannerNodeVersion *ScannerNodeVersionSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.UpgradeToAndCall(&_ScannerNodeVersion.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.UpgradeToAndCall(&_ScannerNodeVersion.TransactOpts, newImplementation, data)
}

// ScannerNodeVersionAccessManagerUpdatedIterator is returned from FilterAccessManagerUpdated and is used to iterate over the raw logs and unpacked data for AccessManagerUpdated events raised by the ScannerNodeVersion contract.
type ScannerNodeVersionAccessManagerUpdatedIterator struct {
	Event *ScannerNodeVersionAccessManagerUpdated // Event containing the contract specifics and raw log

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
func (it *ScannerNodeVersionAccessManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerNodeVersionAccessManagerUpdated)
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
		it.Event = new(ScannerNodeVersionAccessManagerUpdated)
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
func (it *ScannerNodeVersionAccessManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerNodeVersionAccessManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerNodeVersionAccessManagerUpdated represents a AccessManagerUpdated event raised by the ScannerNodeVersion contract.
type ScannerNodeVersionAccessManagerUpdated struct {
	NewAddressManager common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAccessManagerUpdated is a free log retrieval operation binding the contract event 0xa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c.
//
// Solidity: event AccessManagerUpdated(address indexed newAddressManager)
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) FilterAccessManagerUpdated(opts *bind.FilterOpts, newAddressManager []common.Address) (*ScannerNodeVersionAccessManagerUpdatedIterator, error) {

	var newAddressManagerRule []interface{}
	for _, newAddressManagerItem := range newAddressManager {
		newAddressManagerRule = append(newAddressManagerRule, newAddressManagerItem)
	}

	logs, sub, err := _ScannerNodeVersion.contract.FilterLogs(opts, "AccessManagerUpdated", newAddressManagerRule)
	if err != nil {
		return nil, err
	}
	return &ScannerNodeVersionAccessManagerUpdatedIterator{contract: _ScannerNodeVersion.contract, event: "AccessManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchAccessManagerUpdated is a free log subscription operation binding the contract event 0xa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c.
//
// Solidity: event AccessManagerUpdated(address indexed newAddressManager)
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) WatchAccessManagerUpdated(opts *bind.WatchOpts, sink chan<- *ScannerNodeVersionAccessManagerUpdated, newAddressManager []common.Address) (event.Subscription, error) {

	var newAddressManagerRule []interface{}
	for _, newAddressManagerItem := range newAddressManager {
		newAddressManagerRule = append(newAddressManagerRule, newAddressManagerItem)
	}

	logs, sub, err := _ScannerNodeVersion.contract.WatchLogs(opts, "AccessManagerUpdated", newAddressManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerNodeVersionAccessManagerUpdated)
				if err := _ScannerNodeVersion.contract.UnpackLog(event, "AccessManagerUpdated", log); err != nil {
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
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) ParseAccessManagerUpdated(log types.Log) (*ScannerNodeVersionAccessManagerUpdated, error) {
	event := new(ScannerNodeVersionAccessManagerUpdated)
	if err := _ScannerNodeVersion.contract.UnpackLog(event, "AccessManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerNodeVersionAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ScannerNodeVersion contract.
type ScannerNodeVersionAdminChangedIterator struct {
	Event *ScannerNodeVersionAdminChanged // Event containing the contract specifics and raw log

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
func (it *ScannerNodeVersionAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerNodeVersionAdminChanged)
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
		it.Event = new(ScannerNodeVersionAdminChanged)
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
func (it *ScannerNodeVersionAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerNodeVersionAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerNodeVersionAdminChanged represents a AdminChanged event raised by the ScannerNodeVersion contract.
type ScannerNodeVersionAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ScannerNodeVersionAdminChangedIterator, error) {

	logs, sub, err := _ScannerNodeVersion.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ScannerNodeVersionAdminChangedIterator{contract: _ScannerNodeVersion.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ScannerNodeVersionAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ScannerNodeVersion.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerNodeVersionAdminChanged)
				if err := _ScannerNodeVersion.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) ParseAdminChanged(log types.Log) (*ScannerNodeVersionAdminChanged, error) {
	event := new(ScannerNodeVersionAdminChanged)
	if err := _ScannerNodeVersion.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerNodeVersionBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ScannerNodeVersion contract.
type ScannerNodeVersionBeaconUpgradedIterator struct {
	Event *ScannerNodeVersionBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ScannerNodeVersionBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerNodeVersionBeaconUpgraded)
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
		it.Event = new(ScannerNodeVersionBeaconUpgraded)
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
func (it *ScannerNodeVersionBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerNodeVersionBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerNodeVersionBeaconUpgraded represents a BeaconUpgraded event raised by the ScannerNodeVersion contract.
type ScannerNodeVersionBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ScannerNodeVersionBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ScannerNodeVersion.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ScannerNodeVersionBeaconUpgradedIterator{contract: _ScannerNodeVersion.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ScannerNodeVersionBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ScannerNodeVersion.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerNodeVersionBeaconUpgraded)
				if err := _ScannerNodeVersion.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) ParseBeaconUpgraded(log types.Log) (*ScannerNodeVersionBeaconUpgraded, error) {
	event := new(ScannerNodeVersionBeaconUpgraded)
	if err := _ScannerNodeVersion.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerNodeVersionRouterUpdatedIterator is returned from FilterRouterUpdated and is used to iterate over the raw logs and unpacked data for RouterUpdated events raised by the ScannerNodeVersion contract.
type ScannerNodeVersionRouterUpdatedIterator struct {
	Event *ScannerNodeVersionRouterUpdated // Event containing the contract specifics and raw log

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
func (it *ScannerNodeVersionRouterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerNodeVersionRouterUpdated)
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
		it.Event = new(ScannerNodeVersionRouterUpdated)
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
func (it *ScannerNodeVersionRouterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerNodeVersionRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerNodeVersionRouterUpdated represents a RouterUpdated event raised by the ScannerNodeVersion contract.
type ScannerNodeVersionRouterUpdated struct {
	Router common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRouterUpdated is a free log retrieval operation binding the contract event 0x7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc80.
//
// Solidity: event RouterUpdated(address indexed router)
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) FilterRouterUpdated(opts *bind.FilterOpts, router []common.Address) (*ScannerNodeVersionRouterUpdatedIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _ScannerNodeVersion.contract.FilterLogs(opts, "RouterUpdated", routerRule)
	if err != nil {
		return nil, err
	}
	return &ScannerNodeVersionRouterUpdatedIterator{contract: _ScannerNodeVersion.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

// WatchRouterUpdated is a free log subscription operation binding the contract event 0x7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc80.
//
// Solidity: event RouterUpdated(address indexed router)
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *ScannerNodeVersionRouterUpdated, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _ScannerNodeVersion.contract.WatchLogs(opts, "RouterUpdated", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerNodeVersionRouterUpdated)
				if err := _ScannerNodeVersion.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
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
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) ParseRouterUpdated(log types.Log) (*ScannerNodeVersionRouterUpdated, error) {
	event := new(ScannerNodeVersionRouterUpdated)
	if err := _ScannerNodeVersion.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerNodeVersionScannerNodeBetaVersionUpdatedIterator is returned from FilterScannerNodeBetaVersionUpdated and is used to iterate over the raw logs and unpacked data for ScannerNodeBetaVersionUpdated events raised by the ScannerNodeVersion contract.
type ScannerNodeVersionScannerNodeBetaVersionUpdatedIterator struct {
	Event *ScannerNodeVersionScannerNodeBetaVersionUpdated // Event containing the contract specifics and raw log

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
func (it *ScannerNodeVersionScannerNodeBetaVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerNodeVersionScannerNodeBetaVersionUpdated)
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
		it.Event = new(ScannerNodeVersionScannerNodeBetaVersionUpdated)
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
func (it *ScannerNodeVersionScannerNodeBetaVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerNodeVersionScannerNodeBetaVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerNodeVersionScannerNodeBetaVersionUpdated represents a ScannerNodeBetaVersionUpdated event raised by the ScannerNodeVersion contract.
type ScannerNodeVersionScannerNodeBetaVersionUpdated struct {
	NewVersion string
	OldVersion string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterScannerNodeBetaVersionUpdated is a free log retrieval operation binding the contract event 0xe5ac27c9fd9eefb9d9fd2b0878923cf7d44e611ead66cf80b7ccf184ac2927c0.
//
// Solidity: event ScannerNodeBetaVersionUpdated(string newVersion, string oldVersion)
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) FilterScannerNodeBetaVersionUpdated(opts *bind.FilterOpts) (*ScannerNodeVersionScannerNodeBetaVersionUpdatedIterator, error) {

	logs, sub, err := _ScannerNodeVersion.contract.FilterLogs(opts, "ScannerNodeBetaVersionUpdated")
	if err != nil {
		return nil, err
	}
	return &ScannerNodeVersionScannerNodeBetaVersionUpdatedIterator{contract: _ScannerNodeVersion.contract, event: "ScannerNodeBetaVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchScannerNodeBetaVersionUpdated is a free log subscription operation binding the contract event 0xe5ac27c9fd9eefb9d9fd2b0878923cf7d44e611ead66cf80b7ccf184ac2927c0.
//
// Solidity: event ScannerNodeBetaVersionUpdated(string newVersion, string oldVersion)
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) WatchScannerNodeBetaVersionUpdated(opts *bind.WatchOpts, sink chan<- *ScannerNodeVersionScannerNodeBetaVersionUpdated) (event.Subscription, error) {

	logs, sub, err := _ScannerNodeVersion.contract.WatchLogs(opts, "ScannerNodeBetaVersionUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerNodeVersionScannerNodeBetaVersionUpdated)
				if err := _ScannerNodeVersion.contract.UnpackLog(event, "ScannerNodeBetaVersionUpdated", log); err != nil {
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

// ParseScannerNodeBetaVersionUpdated is a log parse operation binding the contract event 0xe5ac27c9fd9eefb9d9fd2b0878923cf7d44e611ead66cf80b7ccf184ac2927c0.
//
// Solidity: event ScannerNodeBetaVersionUpdated(string newVersion, string oldVersion)
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) ParseScannerNodeBetaVersionUpdated(log types.Log) (*ScannerNodeVersionScannerNodeBetaVersionUpdated, error) {
	event := new(ScannerNodeVersionScannerNodeBetaVersionUpdated)
	if err := _ScannerNodeVersion.contract.UnpackLog(event, "ScannerNodeBetaVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerNodeVersionScannerNodeVersionUpdatedIterator is returned from FilterScannerNodeVersionUpdated and is used to iterate over the raw logs and unpacked data for ScannerNodeVersionUpdated events raised by the ScannerNodeVersion contract.
type ScannerNodeVersionScannerNodeVersionUpdatedIterator struct {
	Event *ScannerNodeVersionScannerNodeVersionUpdated // Event containing the contract specifics and raw log

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
func (it *ScannerNodeVersionScannerNodeVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerNodeVersionScannerNodeVersionUpdated)
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
		it.Event = new(ScannerNodeVersionScannerNodeVersionUpdated)
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
func (it *ScannerNodeVersionScannerNodeVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerNodeVersionScannerNodeVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerNodeVersionScannerNodeVersionUpdated represents a ScannerNodeVersionUpdated event raised by the ScannerNodeVersion contract.
type ScannerNodeVersionScannerNodeVersionUpdated struct {
	NewVersion string
	OldVersion string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterScannerNodeVersionUpdated is a free log retrieval operation binding the contract event 0xaa8f23ad4857a5d22df0189ebc6150d51b3152d8c621c3e8d24c66387897819b.
//
// Solidity: event ScannerNodeVersionUpdated(string newVersion, string oldVersion)
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) FilterScannerNodeVersionUpdated(opts *bind.FilterOpts) (*ScannerNodeVersionScannerNodeVersionUpdatedIterator, error) {

	logs, sub, err := _ScannerNodeVersion.contract.FilterLogs(opts, "ScannerNodeVersionUpdated")
	if err != nil {
		return nil, err
	}
	return &ScannerNodeVersionScannerNodeVersionUpdatedIterator{contract: _ScannerNodeVersion.contract, event: "ScannerNodeVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchScannerNodeVersionUpdated is a free log subscription operation binding the contract event 0xaa8f23ad4857a5d22df0189ebc6150d51b3152d8c621c3e8d24c66387897819b.
//
// Solidity: event ScannerNodeVersionUpdated(string newVersion, string oldVersion)
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) WatchScannerNodeVersionUpdated(opts *bind.WatchOpts, sink chan<- *ScannerNodeVersionScannerNodeVersionUpdated) (event.Subscription, error) {

	logs, sub, err := _ScannerNodeVersion.contract.WatchLogs(opts, "ScannerNodeVersionUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerNodeVersionScannerNodeVersionUpdated)
				if err := _ScannerNodeVersion.contract.UnpackLog(event, "ScannerNodeVersionUpdated", log); err != nil {
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

// ParseScannerNodeVersionUpdated is a log parse operation binding the contract event 0xaa8f23ad4857a5d22df0189ebc6150d51b3152d8c621c3e8d24c66387897819b.
//
// Solidity: event ScannerNodeVersionUpdated(string newVersion, string oldVersion)
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) ParseScannerNodeVersionUpdated(log types.Log) (*ScannerNodeVersionScannerNodeVersionUpdated, error) {
	event := new(ScannerNodeVersionScannerNodeVersionUpdated)
	if err := _ScannerNodeVersion.contract.UnpackLog(event, "ScannerNodeVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScannerNodeVersionUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ScannerNodeVersion contract.
type ScannerNodeVersionUpgradedIterator struct {
	Event *ScannerNodeVersionUpgraded // Event containing the contract specifics and raw log

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
func (it *ScannerNodeVersionUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerNodeVersionUpgraded)
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
		it.Event = new(ScannerNodeVersionUpgraded)
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
func (it *ScannerNodeVersionUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerNodeVersionUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerNodeVersionUpgraded represents a Upgraded event raised by the ScannerNodeVersion contract.
type ScannerNodeVersionUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ScannerNodeVersionUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ScannerNodeVersion.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ScannerNodeVersionUpgradedIterator{contract: _ScannerNodeVersion.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ScannerNodeVersionUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ScannerNodeVersion.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerNodeVersionUpgraded)
				if err := _ScannerNodeVersion.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) ParseUpgraded(log types.Log) (*ScannerNodeVersionUpgraded, error) {
	event := new(ScannerNodeVersionUpgraded)
	if err := _ScannerNodeVersion.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
