// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_router

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

// RouterMetaData contains all meta data concerning the Router contract.
var RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"AccessManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"revertsOnFail\",\"type\":\"bool\"}],\"name\":\"RoutingUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"hookHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__manager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"setAccessManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ensRegistry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ensName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"revertsOnFail\",\"type\":\"bool\"}],\"name\":\"setRoutingTable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040523060a0523480156200001557600080fd5b5060405162001ffa38038062001ffa833981016040819052620000389162000105565b6001600160a01b038116608052600054610100900460ff16806200005f575060005460ff16155b620000c75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b600054610100900460ff16158015620000ea576000805461ffff19166101011790555b8015620000fd576000805461ff00191690555b505062000137565b6000602082840312156200011857600080fd5b81516001600160a01b03811681146200013057600080fd5b9392505050565b60805160a051611e8162000179600039600081816103180152818161039d01528181610785015261080a0152600081816101c701526111ec0152611e816000f3fe6080604052600436106100b15760003560e01c806354fd4d5011610069578063ac9650d81161004e578063ac9650d814610207578063c4d66de814610234578063c95808041461025457600080fd5b806354fd4d501461014b578063572b6c05146101aa57600080fd5b80633fa222c41161009a5780633fa222c4146100f85780634d807af9146101185780634f1ef2861461013857600080fd5b80633121db1c146100b65780633659cfe6146100d8575b600080fd5b3480156100c257600080fd5b506100d66100d13660046118e7565b610274565b005b3480156100e457600080fd5b506100d66100f336600461193c565b61030d565b34801561010457600080fd5b506100d6610113366004611959565b610489565b34801561012457600080fd5b506100d66101333660046119a9565b6105e2565b6100d6610146366004611a28565b61077a565b34801561015757600080fd5b506101946040518060400160405280600581526020017f302e312e3000000000000000000000000000000000000000000000000000000081525081565b6040516101a19190611b44565b60405180910390f35b3480156101b657600080fd5b506101f76101c536600461193c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0390811691161490565b60405190151581526020016101a1565b34801561021357600080fd5b50610227610222366004611b57565b6108e7565b6040516101a19190611bcc565b34801561024057600080fd5b506100d661024f36600461193c565b6109dd565b34801561026057600080fd5b506100d661026f36600461193c565b610aa1565b7f664245c7af190fec316596e8231f724e8171b1966cfcd124347ac5a66c34f64a6102a6816102a1610b6c565b610b7b565b6102fc57806102b3610b6c565b6040517f75000dc000000000000000000000000000000000000000000000000000000000815260048101929092526001600160a01b031660248201526044015b60405180910390fd5b610307848484610c20565b50505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016141561039b5760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b60648201526084016102f3565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166103f67f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b0316146104615760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b6163746976652070726f787960a01b60648201526084016102f3565b61046a81610d4f565b6040805160008082526020820190925261048691839190610d89565b50565b60006104986004828486611c2e565b6104a191611c58565b6001600160e01b03198116600090815260c960205260408120919250906104c790610f46565b905060005b818110156105db576001600160e01b03198316600090815260c96020526040812081906104f99084610f50565b6001600160a01b03168787604051610512929190611c88565b6000604051808303816000865af19150503d806000811461054f576040519150601f19603f3d011682016040523d82523d6000602084013e610554565b606091505b506001600160e01b03198716600090815260ca6020526040902054919350915060ff16156105c657816105c6576040517fbde758bc0000000000000000000000000000000000000000000000000000000081526001600160e01b031986166004820152602481018490526044016102f3565b505080806105d390611cae565b9150506104cc565b5050505050565b7f56623da34b1ec5cb86498f15a28504a6323a0eedfb150423fe6f418d952826ee61060f816102a1610b6c565b61061c57806102b3610b6c565b82156106a5576001600160e01b03198516600090815260c9602052604090206106459085610f5c565b61067b576040517f91a1f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160e01b03198516600090815260ca60205260409020805460ff1916831515179055610720565b6001600160e01b03198516600090815260c9602052604090206106c89085610f71565b6106fe576040517fd314330900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160e01b03198516600090815260ca60205260409020805460ff191690555b60408051841515815283151560208201526001600160a01b038616916001600160e01b03198816917fd7732b31e65a6e584dd5881f392ef25ab0b38bdd0b2fb55641e4be90d4dd4f9a910160405180910390a35050505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156108085760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b60648201526084016102f3565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166108637f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b0316146108ce5760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b6163746976652070726f787960a01b60648201526084016102f3565b6108d782610d4f565b6108e382826001610d89565b5050565b60608167ffffffffffffffff81111561090257610902611a12565b60405190808252806020026020018201604052801561093557816020015b60608152602001906001900390816109205790505b50905060005b828110156109d5576109a53085858481811061095957610959611cc9565b905060200281019061096b9190611cdf565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610f8692505050565b8282815181106109b7576109b7611cc9565b602002602001018190525080806109cd90611cae565b91505061093b565b505b92915050565b600054610100900460ff16806109f6575060005460ff16155b610a595760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102f3565b600054610100900460ff16158015610a7b576000805461ffff19166101011790555b610a8482610fab565b610a8c61110a565b80156108e3576000805461ff00191690555050565b6000610aaf816102a1610b6c565b610abc57806102b3610b6c565b610ad66001600160a01b038316637965db0b60e01b6111cc565b610b14576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b60448201526064016102f3565b6033805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a25050565b6000610b766111e8565b905090565b6033546040517f91d14854000000000000000000000000000000000000000000000000000000008152600481018490526001600160a01b03838116602483015260009216906391d148549060440160206040518083038186803b158015610be157600080fd5b505afa158015610bf5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c199190611d26565b9392505050565b6040517f02571be30000000000000000000000000000000000000000000000000000000081527f91d1777781884d03a6757a803996e38de2a42967fb37eeaca72729271025a9e260048201526001600160a01b038416906302571be39060240160206040518083038186803b158015610c9857600080fd5b505afa158015610cac573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cd09190611d43565b6001600160a01b031663c47f002783836040518363ffffffff1660e01b8152600401610cfd929190611d60565b602060405180830381600087803b158015610d1757600080fd5b505af1158015610d2b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103079190611d8f565b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3610d7c816102a1610b6c565b6108e357806102b3610b6c565b6000610dbc7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b9050610dc78461124b565b600083511180610dd45750815b15610de557610de3848461130d565b505b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd9143805460ff166105db57805460ff191660011781556040516001600160a01b0383166024820152610e9290869060440160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f3659cfe60000000000000000000000000000000000000000000000000000000017905261130d565b50805460ff191681557f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b03838116911614610f3d5760405162461bcd60e51b815260206004820152602f60248201527f45524331393637557067726164653a207570677261646520627265616b73206660448201527f757274686572207570677261646573000000000000000000000000000000000060648201526084016102f3565b6105db856113f8565b60006109d7825490565b6000610c198383611438565b6000610c19836001600160a01b038416611462565b6000610c19836001600160a01b0384166114b1565b6060610c198383604051806060016040528060278152602001611e25602791396115a4565b600054610100900460ff1680610fc4575060005460ff16155b6110275760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102f3565b600054610100900460ff16158015611049576000805461ffff19166101011790555b6110636001600160a01b038316637965db0b60e01b6111cc565b6110a1576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b60448201526064016102f3565b6033805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a280156108e3576000805461ff00191690555050565b600054610100900460ff1680611123575060005460ff16155b6111865760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102f3565b600054610100900460ff161580156111a8576000805461ffff19166101011790555b6111b0611678565b6111b8611678565b8015610486576000805461ff001916905550565b60006111d783611729565b8015610c195750610c19838361175c565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633141561124657600036611229601482611da8565b61123592369290611c2e565b61123e91611dbf565b60601c905090565b503390565b803b6112bf5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e74726163740000000000000000000000000000000000000060648201526084016102f3565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b6060823b61136c5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084016102f3565b600080846001600160a01b0316846040516113879190611df2565b600060405180830381855af49150503d80600081146113c2576040519150601f19603f3d011682016040523d82523d6000602084013e6113c7565b606091505b50915091506113ef8282604051806060016040528060278152602001611e2560279139611850565b95945050505050565b6114018161124b565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b600082600001828154811061144f5761144f611cc9565b9060005260206000200154905092915050565b60008181526001830160205260408120546114a9575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556109d7565b5060006109d7565b6000818152600183016020526040812054801561159a5760006114d5600183611da8565b85549091506000906114e990600190611da8565b905081811461154e57600086600001828154811061150957611509611cc9565b906000526020600020015490508087600001848154811061152c5761152c611cc9565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061155f5761155f611e0e565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506109d7565b60009150506109d7565b6060833b6116035760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084016102f3565b600080856001600160a01b03168560405161161e9190611df2565b600060405180830381855af49150503d8060008114611659576040519150601f19603f3d011682016040523d82523d6000602084013e61165e565b606091505b509150915061166e828286611850565b9695505050505050565b600054610100900460ff1680611691575060005460ff16155b6116f45760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102f3565b600054610100900460ff161580156111b8576000805461ffff19166101011790558015610486576000805461ff001916905550565b600061173c826301ffc9a760e01b61175c565b80156109d75750611755826001600160e01b031961175c565b1592915050565b604080516001600160e01b0319831660248083019190915282518083039091018152604490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166301ffc9a760e01b179052905160009190829081906001600160a01b03871690617530906117d8908690611df2565b6000604051808303818686fa925050503d8060008114611814576040519150601f19603f3d011682016040523d82523d6000602084013e611819565b606091505b509150915060208151101561183457600093505050506109d7565b81801561166e57508080602001905181019061166e9190611d26565b6060831561185f575081610c19565b82511561186f5782518084602001fd5b8160405162461bcd60e51b81526004016102f39190611b44565b6001600160a01b038116811461048657600080fd5b60008083601f8401126118b057600080fd5b50813567ffffffffffffffff8111156118c857600080fd5b6020830191508360208285010111156118e057600080fd5b9250929050565b6000806000604084860312156118fc57600080fd5b833561190781611889565b9250602084013567ffffffffffffffff81111561192357600080fd5b61192f8682870161189e565b9497909650939450505050565b60006020828403121561194e57600080fd5b8135610c1981611889565b6000806020838503121561196c57600080fd5b823567ffffffffffffffff81111561198357600080fd5b61198f8582860161189e565b90969095509350505050565b801515811461048657600080fd5b600080600080608085870312156119bf57600080fd5b84356001600160e01b0319811681146119d757600080fd5b935060208501356119e781611889565b925060408501356119f78161199b565b91506060850135611a078161199b565b939692955090935050565b634e487b7160e01b600052604160045260246000fd5b60008060408385031215611a3b57600080fd5b8235611a4681611889565b9150602083013567ffffffffffffffff80821115611a6357600080fd5b818501915085601f830112611a7757600080fd5b813581811115611a8957611a89611a12565b604051601f8201601f19908116603f01168101908382118183101715611ab157611ab1611a12565b81604052828152886020848701011115611aca57600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b60005b83811015611b07578181015183820152602001611aef565b838111156103075750506000910152565b60008151808452611b30816020860160208601611aec565b601f01601f19169290920160200192915050565b602081526000610c196020830184611b18565b60008060208385031215611b6a57600080fd5b823567ffffffffffffffff80821115611b8257600080fd5b818501915085601f830112611b9657600080fd5b813581811115611ba557600080fd5b8660208260051b8501011115611bba57600080fd5b60209290920196919550909350505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015611c2157603f19888603018452611c0f858351611b18565b94509285019290850190600101611bf3565b5092979650505050505050565b60008085851115611c3e57600080fd5b83861115611c4b57600080fd5b5050820193919092039150565b6001600160e01b03198135818116916004851015611c805780818660040360031b1b83161692505b505092915050565b8183823760009101908152919050565b634e487b7160e01b600052601160045260246000fd5b6000600019821415611cc257611cc2611c98565b5060010190565b634e487b7160e01b600052603260045260246000fd5b6000808335601e19843603018112611cf657600080fd5b83018035915067ffffffffffffffff821115611d1157600080fd5b6020019150368190038213156118e057600080fd5b600060208284031215611d3857600080fd5b8151610c198161199b565b600060208284031215611d5557600080fd5b8151610c1981611889565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b600060208284031215611da157600080fd5b5051919050565b600082821015611dba57611dba611c98565b500390565b6bffffffffffffffffffffffff198135818116916014851015611c805760149490940360031b84901b1690921692915050565b60008251611e04818460208701611aec565b9190910192915050565b634e487b7160e01b600052603160045260246000fdfe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212201b9462bfb10899ca4bbfe1f35d46c52e35417ca1cf6d2efeee3402db52d7a45464736f6c63430008090033",
}

// RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterMetaData.ABI instead.
var RouterABI = RouterMetaData.ABI

// RouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RouterMetaData.Bin instead.
var RouterBin = RouterMetaData.Bin

// DeployRouter deploys a new Ethereum contract, binding an instance of Router to it.
func DeployRouter(auth *bind.TransactOpts, backend bind.ContractBackend, forwarder common.Address) (common.Address, *types.Transaction, *Router, error) {
	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RouterBin), backend, forwarder)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Router *RouterCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Router *RouterSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Router.Contract.IsTrustedForwarder(&_Router.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Router *RouterCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Router.Contract.IsTrustedForwarder(&_Router.CallOpts, forwarder)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Router *RouterCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Router *RouterSession) Version() (string, error) {
	return _Router.Contract.Version(&_Router.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Router *RouterCallerSession) Version() (string, error) {
	return _Router.Contract.Version(&_Router.CallOpts)
}

// HookHandler is a paid mutator transaction binding the contract method 0x3fa222c4.
//
// Solidity: function hookHandler(bytes payload) returns()
func (_Router *RouterTransactor) HookHandler(opts *bind.TransactOpts, payload []byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "hookHandler", payload)
}

// HookHandler is a paid mutator transaction binding the contract method 0x3fa222c4.
//
// Solidity: function hookHandler(bytes payload) returns()
func (_Router *RouterSession) HookHandler(payload []byte) (*types.Transaction, error) {
	return _Router.Contract.HookHandler(&_Router.TransactOpts, payload)
}

// HookHandler is a paid mutator transaction binding the contract method 0x3fa222c4.
//
// Solidity: function hookHandler(bytes payload) returns()
func (_Router *RouterTransactorSession) HookHandler(payload []byte) (*types.Transaction, error) {
	return _Router.Contract.HookHandler(&_Router.TransactOpts, payload)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address __manager) returns()
func (_Router *RouterTransactor) Initialize(opts *bind.TransactOpts, __manager common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "initialize", __manager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address __manager) returns()
func (_Router *RouterSession) Initialize(__manager common.Address) (*types.Transaction, error) {
	return _Router.Contract.Initialize(&_Router.TransactOpts, __manager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address __manager) returns()
func (_Router *RouterTransactorSession) Initialize(__manager common.Address) (*types.Transaction, error) {
	return _Router.Contract.Initialize(&_Router.TransactOpts, __manager)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Router *RouterTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Router *RouterSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Router.Contract.Multicall(&_Router.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Router *RouterTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Router.Contract.Multicall(&_Router.TransactOpts, data)
}

// SetAccessManager is a paid mutator transaction binding the contract method 0xc9580804.
//
// Solidity: function setAccessManager(address newManager) returns()
func (_Router *RouterTransactor) SetAccessManager(opts *bind.TransactOpts, newManager common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setAccessManager", newManager)
}

// SetAccessManager is a paid mutator transaction binding the contract method 0xc9580804.
//
// Solidity: function setAccessManager(address newManager) returns()
func (_Router *RouterSession) SetAccessManager(newManager common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetAccessManager(&_Router.TransactOpts, newManager)
}

// SetAccessManager is a paid mutator transaction binding the contract method 0xc9580804.
//
// Solidity: function setAccessManager(address newManager) returns()
func (_Router *RouterTransactorSession) SetAccessManager(newManager common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetAccessManager(&_Router.TransactOpts, newManager)
}

// SetName is a paid mutator transaction binding the contract method 0x3121db1c.
//
// Solidity: function setName(address ensRegistry, string ensName) returns()
func (_Router *RouterTransactor) SetName(opts *bind.TransactOpts, ensRegistry common.Address, ensName string) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setName", ensRegistry, ensName)
}

// SetName is a paid mutator transaction binding the contract method 0x3121db1c.
//
// Solidity: function setName(address ensRegistry, string ensName) returns()
func (_Router *RouterSession) SetName(ensRegistry common.Address, ensName string) (*types.Transaction, error) {
	return _Router.Contract.SetName(&_Router.TransactOpts, ensRegistry, ensName)
}

// SetName is a paid mutator transaction binding the contract method 0x3121db1c.
//
// Solidity: function setName(address ensRegistry, string ensName) returns()
func (_Router *RouterTransactorSession) SetName(ensRegistry common.Address, ensName string) (*types.Transaction, error) {
	return _Router.Contract.SetName(&_Router.TransactOpts, ensRegistry, ensName)
}

// SetRoutingTable is a paid mutator transaction binding the contract method 0x4d807af9.
//
// Solidity: function setRoutingTable(bytes4 sig, address target, bool enable, bool revertsOnFail) returns()
func (_Router *RouterTransactor) SetRoutingTable(opts *bind.TransactOpts, sig [4]byte, target common.Address, enable bool, revertsOnFail bool) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setRoutingTable", sig, target, enable, revertsOnFail)
}

// SetRoutingTable is a paid mutator transaction binding the contract method 0x4d807af9.
//
// Solidity: function setRoutingTable(bytes4 sig, address target, bool enable, bool revertsOnFail) returns()
func (_Router *RouterSession) SetRoutingTable(sig [4]byte, target common.Address, enable bool, revertsOnFail bool) (*types.Transaction, error) {
	return _Router.Contract.SetRoutingTable(&_Router.TransactOpts, sig, target, enable, revertsOnFail)
}

// SetRoutingTable is a paid mutator transaction binding the contract method 0x4d807af9.
//
// Solidity: function setRoutingTable(bytes4 sig, address target, bool enable, bool revertsOnFail) returns()
func (_Router *RouterTransactorSession) SetRoutingTable(sig [4]byte, target common.Address, enable bool, revertsOnFail bool) (*types.Transaction, error) {
	return _Router.Contract.SetRoutingTable(&_Router.TransactOpts, sig, target, enable, revertsOnFail)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Router *RouterTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Router *RouterSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Router.Contract.UpgradeTo(&_Router.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Router *RouterTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Router.Contract.UpgradeTo(&_Router.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Router *RouterTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Router *RouterSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Router.Contract.UpgradeToAndCall(&_Router.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Router *RouterTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Router.Contract.UpgradeToAndCall(&_Router.TransactOpts, newImplementation, data)
}

// RouterAccessManagerUpdatedIterator is returned from FilterAccessManagerUpdated and is used to iterate over the raw logs and unpacked data for AccessManagerUpdated events raised by the Router contract.
type RouterAccessManagerUpdatedIterator struct {
	Event *RouterAccessManagerUpdated // Event containing the contract specifics and raw log

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
func (it *RouterAccessManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterAccessManagerUpdated)
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
		it.Event = new(RouterAccessManagerUpdated)
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
func (it *RouterAccessManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterAccessManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterAccessManagerUpdated represents a AccessManagerUpdated event raised by the Router contract.
type RouterAccessManagerUpdated struct {
	NewAddressManager common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAccessManagerUpdated is a free log retrieval operation binding the contract event 0xa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c.
//
// Solidity: event AccessManagerUpdated(address indexed newAddressManager)
func (_Router *RouterFilterer) FilterAccessManagerUpdated(opts *bind.FilterOpts, newAddressManager []common.Address) (*RouterAccessManagerUpdatedIterator, error) {

	var newAddressManagerRule []interface{}
	for _, newAddressManagerItem := range newAddressManager {
		newAddressManagerRule = append(newAddressManagerRule, newAddressManagerItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "AccessManagerUpdated", newAddressManagerRule)
	if err != nil {
		return nil, err
	}
	return &RouterAccessManagerUpdatedIterator{contract: _Router.contract, event: "AccessManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchAccessManagerUpdated is a free log subscription operation binding the contract event 0xa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c.
//
// Solidity: event AccessManagerUpdated(address indexed newAddressManager)
func (_Router *RouterFilterer) WatchAccessManagerUpdated(opts *bind.WatchOpts, sink chan<- *RouterAccessManagerUpdated, newAddressManager []common.Address) (event.Subscription, error) {

	var newAddressManagerRule []interface{}
	for _, newAddressManagerItem := range newAddressManager {
		newAddressManagerRule = append(newAddressManagerRule, newAddressManagerItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "AccessManagerUpdated", newAddressManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterAccessManagerUpdated)
				if err := _Router.contract.UnpackLog(event, "AccessManagerUpdated", log); err != nil {
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
func (_Router *RouterFilterer) ParseAccessManagerUpdated(log types.Log) (*RouterAccessManagerUpdated, error) {
	event := new(RouterAccessManagerUpdated)
	if err := _Router.contract.UnpackLog(event, "AccessManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Router contract.
type RouterAdminChangedIterator struct {
	Event *RouterAdminChanged // Event containing the contract specifics and raw log

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
func (it *RouterAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterAdminChanged)
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
		it.Event = new(RouterAdminChanged)
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
func (it *RouterAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterAdminChanged represents a AdminChanged event raised by the Router contract.
type RouterAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Router *RouterFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*RouterAdminChangedIterator, error) {

	logs, sub, err := _Router.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &RouterAdminChangedIterator{contract: _Router.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Router *RouterFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *RouterAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Router.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterAdminChanged)
				if err := _Router.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Router *RouterFilterer) ParseAdminChanged(log types.Log) (*RouterAdminChanged, error) {
	event := new(RouterAdminChanged)
	if err := _Router.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Router contract.
type RouterBeaconUpgradedIterator struct {
	Event *RouterBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *RouterBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterBeaconUpgraded)
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
		it.Event = new(RouterBeaconUpgraded)
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
func (it *RouterBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterBeaconUpgraded represents a BeaconUpgraded event raised by the Router contract.
type RouterBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Router *RouterFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*RouterBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &RouterBeaconUpgradedIterator{contract: _Router.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Router *RouterFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *RouterBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterBeaconUpgraded)
				if err := _Router.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Router *RouterFilterer) ParseBeaconUpgraded(log types.Log) (*RouterBeaconUpgraded, error) {
	event := new(RouterBeaconUpgraded)
	if err := _Router.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterRoutingUpdatedIterator is returned from FilterRoutingUpdated and is used to iterate over the raw logs and unpacked data for RoutingUpdated events raised by the Router contract.
type RouterRoutingUpdatedIterator struct {
	Event *RouterRoutingUpdated // Event containing the contract specifics and raw log

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
func (it *RouterRoutingUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterRoutingUpdated)
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
		it.Event = new(RouterRoutingUpdated)
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
func (it *RouterRoutingUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterRoutingUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterRoutingUpdated represents a RoutingUpdated event raised by the Router contract.
type RouterRoutingUpdated struct {
	Sig           [4]byte
	Target        common.Address
	Enable        bool
	RevertsOnFail bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRoutingUpdated is a free log retrieval operation binding the contract event 0xd7732b31e65a6e584dd5881f392ef25ab0b38bdd0b2fb55641e4be90d4dd4f9a.
//
// Solidity: event RoutingUpdated(bytes4 indexed sig, address indexed target, bool enable, bool revertsOnFail)
func (_Router *RouterFilterer) FilterRoutingUpdated(opts *bind.FilterOpts, sig [][4]byte, target []common.Address) (*RouterRoutingUpdatedIterator, error) {

	var sigRule []interface{}
	for _, sigItem := range sig {
		sigRule = append(sigRule, sigItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "RoutingUpdated", sigRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &RouterRoutingUpdatedIterator{contract: _Router.contract, event: "RoutingUpdated", logs: logs, sub: sub}, nil
}

// WatchRoutingUpdated is a free log subscription operation binding the contract event 0xd7732b31e65a6e584dd5881f392ef25ab0b38bdd0b2fb55641e4be90d4dd4f9a.
//
// Solidity: event RoutingUpdated(bytes4 indexed sig, address indexed target, bool enable, bool revertsOnFail)
func (_Router *RouterFilterer) WatchRoutingUpdated(opts *bind.WatchOpts, sink chan<- *RouterRoutingUpdated, sig [][4]byte, target []common.Address) (event.Subscription, error) {

	var sigRule []interface{}
	for _, sigItem := range sig {
		sigRule = append(sigRule, sigItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "RoutingUpdated", sigRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterRoutingUpdated)
				if err := _Router.contract.UnpackLog(event, "RoutingUpdated", log); err != nil {
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

// ParseRoutingUpdated is a log parse operation binding the contract event 0xd7732b31e65a6e584dd5881f392ef25ab0b38bdd0b2fb55641e4be90d4dd4f9a.
//
// Solidity: event RoutingUpdated(bytes4 indexed sig, address indexed target, bool enable, bool revertsOnFail)
func (_Router *RouterFilterer) ParseRoutingUpdated(log types.Log) (*RouterRoutingUpdated, error) {
	event := new(RouterRoutingUpdated)
	if err := _Router.contract.UnpackLog(event, "RoutingUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Router contract.
type RouterUpgradedIterator struct {
	Event *RouterUpgraded // Event containing the contract specifics and raw log

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
func (it *RouterUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterUpgraded)
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
		it.Event = new(RouterUpgraded)
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
func (it *RouterUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterUpgraded represents a Upgraded event raised by the Router contract.
type RouterUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Router *RouterFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*RouterUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &RouterUpgradedIterator{contract: _Router.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Router *RouterFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *RouterUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterUpgraded)
				if err := _Router.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Router *RouterFilterer) ParseUpgraded(log types.Log) (*RouterUpgraded, error) {
	event := new(RouterUpgraded)
	if err := _Router.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
