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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"AccessManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"revertsOnFail\",\"type\":\"bool\"}],\"name\":\"RoutingUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"hookHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__manager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"setAccessManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ensRegistry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ensName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"revertsOnFail\",\"type\":\"bool\"}],\"name\":\"setRoutingTable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040523060a0523480156200001557600080fd5b5060405162001ea538038062001ea5833981016040819052620000389162000180565b6001600160a01b038116608052600054610100900460ff1615808015620000665750600054600160ff909116105b8062000096575062000083306200017160201b62000a761760201c565b15801562000096575060005460ff166001145b620000fe5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b6000805460ff19166001179055801562000122576000805461ff0019166101001790555b801562000169576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050620001b2565b6001600160a01b03163b151590565b6000602082840312156200019357600080fd5b81516001600160a01b0381168114620001ab57600080fd5b9392505050565b60805160a051611caa620001fb600039600081816102f00152818161033001528181610675015281816106b501526107480152600081816101b90152610fff0152611caa6000f3fe60806040526004361061009c5760003560e01c806352d1902d1161006457806352d1902d1461013657806354fd4d501461015e578063572b6c051461019c578063ac9650d8146101f9578063c4d66de814610226578063c95808041461024657600080fd5b80633121db1c146100a15780633659cfe6146100c35780633fa222c4146100e35780634d807af9146101035780634f1ef28614610123575b600080fd5b3480156100ad57600080fd5b506100c16100bc3660046115c6565b610266565b005b3480156100cf57600080fd5b506100c16100de36600461161b565b6102e6565b3480156100ef57600080fd5b506100c16100fe366004611638565b6103c5565b34801561010f57600080fd5b506100c161011e366004611688565b610505565b6100c1610131366004611707565b61066b565b34801561014257600080fd5b5061014b61073b565b6040519081526020015b60405180910390f35b34801561016a57600080fd5b5061018f604051806040016040528060058152602001640302e312e360dc1b81525081565b6040516101559190611823565b3480156101a857600080fd5b506101e96101b736600461161b565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0390811691161490565b6040519015158152602001610155565b34801561020557600080fd5b50610219610214366004611836565b6107ee565b60405161015591906118ab565b34801561023257600080fd5b506100c161024136600461161b565b6108e4565b34801561025257600080fd5b506100c161026136600461161b565b6109b8565b7f664245c7af190fec316596e8231f724e8171b1966cfcd124347ac5a66c34f64a61029881610293610a85565b610a94565b6102d557806102a5610a85565b6040516301d4003760e61b815260048101929092526001600160a01b031660248201526044015b60405180910390fd5b6102e0848484610b11565b50505050565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016300361032e5760405162461bcd60e51b81526004016102cc9061190d565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610377600080516020611c2e833981519152546001600160a01b031690565b6001600160a01b03161461039d5760405162461bcd60e51b81526004016102cc90611959565b6103a681610c09565b604080516000808252602082019092526103c291839190610c43565b50565b60006103d460048284866119a5565b6103dd916119cf565b6001600160e01b03198116600090815260c9602052604081209192509061040390610db3565b905060005b818110156104fe576001600160e01b03198316600090815260c96020526040812081906104359084610dbd565b6001600160a01b0316878760405161044e9291906119ff565b6000604051808303816000865af19150503d806000811461048b576040519150601f19603f3d011682016040523d82523d6000602084013e610490565b606091505b506001600160e01b03198716600090815260ca6020526040902054919350915060ff16156104e957816104e957604051632f79d62f60e21b81526001600160e01b031986166004820152602481018490526044016102cc565b505080806104f690611a25565b915050610408565b5050505050565b7f56623da34b1ec5cb86498f15a28504a6323a0eedfb150423fe6f418d952826ee61053281610293610a85565b61053f57806102a5610a85565b82156105af576001600160e01b03198516600090815260c9602052604090206105689085610dc9565b610585576040516348d0fc2160e11b815260040160405180910390fd5b6001600160e01b03198516600090815260ca60205260409020805460ff1916831515179055610611565b6001600160e01b03198516600090815260c9602052604090206105d29085610dde565b6105ef5760405163d314330960e01b815260040160405180910390fd5b6001600160e01b03198516600090815260ca60205260409020805460ff191690555b60408051841515815283151560208201526001600160a01b038616916001600160e01b03198816917fd7732b31e65a6e584dd5881f392ef25ab0b38bdd0b2fb55641e4be90d4dd4f9a910160405180910390a35050505050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036106b35760405162461bcd60e51b81526004016102cc9061190d565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166106fc600080516020611c2e833981519152546001600160a01b031690565b6001600160a01b0316146107225760405162461bcd60e51b81526004016102cc90611959565b61072b82610c09565b61073782826001610c43565b5050565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107db5760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c000000000000000060648201526084016102cc565b50600080516020611c2e83398151915290565b60608167ffffffffffffffff811115610809576108096116f1565b60405190808252806020026020018201604052801561083c57816020015b60608152602001906001900390816108275790505b50905060005b828110156108dc576108ac3085858481811061086057610860611a3e565b90506020028101906108729190611a54565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610df392505050565b8282815181106108be576108be611a3e565b602002602001018190525080806108d490611a25565b915050610842565b505b92915050565b600054610100900460ff16158080156109045750600054600160ff909116105b8061091e5750303b15801561091e575060005460ff166001145b61093a5760405162461bcd60e51b81526004016102cc90611a9b565b6000805460ff19166001179055801561095d576000805461ff0019166101001790555b61096682610e18565b61096e610f72565b8015610737576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15050565b60006109c681610293610a85565b6109d357806102a5610a85565b6109ed6001600160a01b038316637965db0b60e01b610fdf565b610a2b576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b60448201526064016102cc565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a25050565b6001600160a01b03163b151590565b6000610a8f610ffb565b905090565b603354604051632474521560e21b8152600481018490526001600160a01b03838116602483015260009216906391d1485490604401602060405180830381865afa158015610ae6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b0a9190611ae9565b9392505050565b6040516302571be360e01b81527f91d1777781884d03a6757a803996e38de2a42967fb37eeaca72729271025a9e260048201526001600160a01b038416906302571be390602401602060405180830381865afa158015610b75573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b999190611b06565b6001600160a01b031663c47f002783836040518363ffffffff1660e01b8152600401610bc6929190611b23565b6020604051808303816000875af1158015610be5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102e09190611b52565b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3610c3681610293610a85565b61073757806102a5610a85565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610c7b57610c768361105d565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610cd5575060408051601f3d908101601f19168201909252610cd291810190611b52565b60015b610d385760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b60648201526084016102cc565b600080516020611c2e8339815191528114610da75760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b60648201526084016102cc565b50610c768383836110f9565b60006108de825490565b6000610b0a838361111e565b6000610b0a836001600160a01b038416611148565b6000610b0a836001600160a01b038416611197565b6060610b0a8383604051806060016040528060278152602001611c4e6027913961128a565b600054610100900460ff1615808015610e385750600054600160ff909116105b80610e525750303b158015610e52575060005460ff166001145b610e6e5760405162461bcd60e51b81526004016102cc90611a9b565b6000805460ff191660011790558015610e91576000805461ff0019166101001790555b610eab6001600160a01b038316637965db0b60e01b610fdf565b610ee9576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b60448201526064016102cc565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a28015610737576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020016109ac565b600054610100900460ff16610fdd5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b60648201526084016102cc565b565b6000610fea83611328565b8015610b0a5750610b0a838361135b565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633036110585760003661103b601482611b6b565b611047923692906119a5565b61105091611b82565b60601c905090565b503390565b6001600160a01b0381163b6110ca5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084016102cc565b600080516020611c2e83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6111028361143a565b60008251118061110f5750805b15610c76576102e0838361147a565b600082600001828154811061113557611135611a3e565b9060005260206000200154905092915050565b600081815260018301602052604081205461118f575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556108de565b5060006108de565b600081815260018301602052604081205480156112805760006111bb600183611b6b565b85549091506000906111cf90600190611b6b565b90508181146112345760008660000182815481106111ef576111ef611a3e565b906000526020600020015490508087600001848154811061121257611212611a3e565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061124557611245611bb5565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506108de565b60009150506108de565b60606001600160a01b0384163b6112b35760405162461bcd60e51b81526004016102cc90611bcb565b600080856001600160a01b0316856040516112ce9190611c11565b600060405180830381855af49150503d8060008114611309576040519150601f19603f3d011682016040523d82523d6000602084013e61130e565b606091505b509150915061131e82828661152f565b9695505050505050565b600061133b826301ffc9a760e01b61135b565b80156108de5750611354826001600160e01b031961135b565b1592915050565b604080516001600160e01b0319831660248083019190915282518083039091018152604490910182526020810180516001600160e01b03166301ffc9a760e01b179052905160009190829081906001600160a01b03871690617530906113c2908690611c11565b6000604051808303818686fa925050503d80600081146113fe576040519150601f19603f3d011682016040523d82523d6000602084013e611403565b606091505b509150915060208151101561141e57600093505050506108de565b81801561131e57508080602001905181019061131e9190611ae9565b6114438161105d565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606001600160a01b0383163b6114a35760405162461bcd60e51b81526004016102cc90611bcb565b600080846001600160a01b0316846040516114be9190611c11565b600060405180830381855af49150503d80600081146114f9576040519150601f19603f3d011682016040523d82523d6000602084013e6114fe565b606091505b50915091506115268282604051806060016040528060278152602001611c4e6027913961152f565b95945050505050565b6060831561153e575081610b0a565b82511561154e5782518084602001fd5b8160405162461bcd60e51b81526004016102cc9190611823565b6001600160a01b03811681146103c257600080fd5b60008083601f84011261158f57600080fd5b50813567ffffffffffffffff8111156115a757600080fd5b6020830191508360208285010111156115bf57600080fd5b9250929050565b6000806000604084860312156115db57600080fd5b83356115e681611568565b9250602084013567ffffffffffffffff81111561160257600080fd5b61160e8682870161157d565b9497909650939450505050565b60006020828403121561162d57600080fd5b8135610b0a81611568565b6000806020838503121561164b57600080fd5b823567ffffffffffffffff81111561166257600080fd5b61166e8582860161157d565b90969095509350505050565b80151581146103c257600080fd5b6000806000806080858703121561169e57600080fd5b84356001600160e01b0319811681146116b657600080fd5b935060208501356116c681611568565b925060408501356116d68161167a565b915060608501356116e68161167a565b939692955090935050565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561171a57600080fd5b823561172581611568565b9150602083013567ffffffffffffffff8082111561174257600080fd5b818501915085601f83011261175657600080fd5b813581811115611768576117686116f1565b604051601f8201601f19908116603f01168101908382118183101715611790576117906116f1565b816040528281528860208487010111156117a957600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b60005b838110156117e65781810151838201526020016117ce565b838111156102e05750506000910152565b6000815180845261180f8160208601602086016117cb565b601f01601f19169290920160200192915050565b602081526000610b0a60208301846117f7565b6000806020838503121561184957600080fd5b823567ffffffffffffffff8082111561186157600080fd5b818501915085601f83011261187557600080fd5b81358181111561188457600080fd5b8660208260051b850101111561189957600080fd5b60209290920196919550909350505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561190057603f198886030184526118ee8583516117f7565b945092850192908501906001016118d2565b5092979650505050505050565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b600080858511156119b557600080fd5b838611156119c257600080fd5b5050820193919092039150565b6001600160e01b031981358181169160048510156119f75780818660040360031b1b83161692505b505092915050565b8183823760009101908152919050565b634e487b7160e01b600052601160045260246000fd5b600060018201611a3757611a37611a0f565b5060010190565b634e487b7160e01b600052603260045260246000fd5b6000808335601e19843603018112611a6b57600080fd5b83018035915067ffffffffffffffff821115611a8657600080fd5b6020019150368190038213156115bf57600080fd5b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b600060208284031215611afb57600080fd5b8151610b0a8161167a565b600060208284031215611b1857600080fd5b8151610b0a81611568565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b600060208284031215611b6457600080fd5b5051919050565b600082821015611b7d57611b7d611a0f565b500390565b6bffffffffffffffffffffffff1981358181169160148510156119f75760149490940360031b84901b1690921692915050565b634e487b7160e01b600052603160045260246000fd5b60208082526026908201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6040820152651b9d1c9858dd60d21b606082015260800190565b60008251611c238184602087016117cb565b919091019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a264697066735822122004c865a50159536d9cb48c39494f090c7fa000a05fe26c73af1188eb25e6e71c64736f6c634300080f0033",
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

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Router *RouterCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Router *RouterSession) ProxiableUUID() ([32]byte, error) {
	return _Router.Contract.ProxiableUUID(&_Router.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Router *RouterCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Router.Contract.ProxiableUUID(&_Router.CallOpts)
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

// RouterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Router contract.
type RouterInitializedIterator struct {
	Event *RouterInitialized // Event containing the contract specifics and raw log

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
func (it *RouterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterInitialized)
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
		it.Event = new(RouterInitialized)
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
func (it *RouterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterInitialized represents a Initialized event raised by the Router contract.
type RouterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Router *RouterFilterer) FilterInitialized(opts *bind.FilterOpts) (*RouterInitializedIterator, error) {

	logs, sub, err := _Router.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RouterInitializedIterator{contract: _Router.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Router *RouterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RouterInitialized) (event.Subscription, error) {

	logs, sub, err := _Router.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterInitialized)
				if err := _Router.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Router *RouterFilterer) ParseInitialized(log types.Log) (*RouterInitialized, error) {
	event := new(RouterInitialized)
	if err := _Router.contract.UnpackLog(event, "Initialized", log); err != nil {
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
