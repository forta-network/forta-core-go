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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"AccessManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newVersion\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldVersion\",\"type\":\"string\"}],\"name\":\"ScannerNodeVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__router\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scannerNodeVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"setAccessManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ensRegistry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ensName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"setScannerNodeVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040523060a0523480156200001557600080fd5b50604051620020b9380380620020b9833981016040819052620000389162000105565b6001600160a01b038116608052600054610100900460ff16806200005f575060005460ff16155b620000c75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b600054610100900460ff16158015620000ea576000805461ffff19166101011790555b8015620000fd576000805461ff00191690555b505062000137565b6000602082840312156200011857600080fd5b81516001600160a01b03811681146200013057600080fd5b9392505050565b60805160a051611f4062000179600039600081816103c70152818161044c0152818161061201526106970152600081816101c7015261127e0152611f406000f3fe6080604052600436106100bc5760003560e01c806354fd4d5011610074578063ac9650d81161004e578063ac9650d814610227578063c0d7865514610254578063c95808041461027457600080fd5b806354fd4d5014610161578063572b6c05146101aa57806394bb55a21461020757600080fd5b80633659cfe6116100a55780633659cfe61461010e578063485cc9551461012e5780634f1ef2861461014e57600080fd5b80633121db1c146100c1578063345db3e1146100e3575b600080fd5b3480156100cd57600080fd5b506100e16100dc3660046118c9565b610294565b005b3480156100ef57600080fd5b506100f861032d565b6040516101059190611976565b60405180910390f35b34801561011a57600080fd5b506100e1610129366004611989565b6103bc565b34801561013a57600080fd5b506100e16101493660046119a6565b610538565b6100e161015c3660046119f5565b610607565b34801561016d57600080fd5b506100f86040518060400160405280600581526020017f302e312e3000000000000000000000000000000000000000000000000000000081525081565b3480156101b657600080fd5b506101f76101c5366004611989565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0390811691161490565b6040519015158152602001610105565b34801561021357600080fd5b506100e1610222366004611ab9565b610774565b34801561023357600080fd5b50610247610242366004611afb565b610882565b6040516101059190611b70565b34801561026057600080fd5b506100e161026f366004611989565b610978565b34801561028057600080fd5b506100e161028f366004611989565b610a35565b7f664245c7af190fec316596e8231f724e8171b1966cfcd124347ac5a66c34f64a6102c6816102c1610af3565b610b02565b61031c57806102d3610af3565b6040517f75000dc000000000000000000000000000000000000000000000000000000000815260048101929092526001600160a01b031660248201526044015b60405180910390fd5b610327848484610ba7565b50505050565b61012d805461033b90611bd2565b80601f016020809104026020016040519081016040528092919081815260200182805461036790611bd2565b80156103b45780601f10610389576101008083540402835291602001916103b4565b820191906000526020600020905b81548152906001019060200180831161039757829003601f168201915b505050505081565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016141561044a5760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b6064820152608401610313565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166104a57f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b0316146105105760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b6163746976652070726f787960a01b6064820152608401610313565b61051981610cd6565b6040805160008082526020820190925261053591839190610d10565b50565b600054610100900460ff1680610551575060005460ff16155b6105b45760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610313565b600054610100900460ff161580156105d6576000805461ffff19166101011790555b6105df83610ed4565b6105e882611026565b6105f0611177565b8015610602576000805461ff00191690555b505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156106955760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b6064820152608401610313565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166106f07f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b03161461075b5760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b6163746976652070726f787960a01b6064820152608401610313565b61076482610cd6565b61077082826001610d10565b5050565b7f93bf097503ee765c5402c05999a7d54bac82299bf183ba1f5f2681ab2c6bd70c6107a1816102c1610af3565b6107ae57806102d3610af3565b82826040516020016107c1929190611c0d565b6040516020818303038152906040528051906020012061012d6040516020016107ea9190611c1d565b604051602081830303815290604052805190602001201415610838576040517fae01a30100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7faa8f23ad4857a5d22df0189ebc6150d51b3152d8c621c3e8d24c66387897819b838361012d60405161086d93929190611cb8565b60405180910390a161032761012d84846117d2565b60608167ffffffffffffffff81111561089d5761089d6119df565b6040519080825280602002602001820160405280156108d057816020015b60608152602001906001900390816108bb5790505b50905060005b8281101561097057610940308585848181106108f4576108f4611d4f565b90506020028101906109069190611d65565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061123992505050565b82828151811061095257610952611d4f565b6020026020010181905250808061096890611dc2565b9150506108d6565b505b92915050565b6000610986816102c1610af3565b61099357806102d3610af3565b6001600160a01b0382166109ea5760405163eac0d38960e01b815260206004820152600960248201527f6e6577526f7574657200000000000000000000000000000000000000000000006044820152606401610313565b606580546001600160a01b0319166001600160a01b0384169081179091556040517f7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc8090600090a25050565b6000610a43816102c1610af3565b610a5057806102d3610af3565b610a6a6001600160a01b038316637965db0b60e01b61125e565b610aa8576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b6044820152606401610313565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a25050565b6000610afd61127a565b905090565b6033546040517f91d14854000000000000000000000000000000000000000000000000000000008152600481018490526001600160a01b03838116602483015260009216906391d148549060440160206040518083038186803b158015610b6857600080fd5b505afa158015610b7c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ba09190611ddd565b9392505050565b6040517f02571be30000000000000000000000000000000000000000000000000000000081527f91d1777781884d03a6757a803996e38de2a42967fb37eeaca72729271025a9e260048201526001600160a01b038416906302571be39060240160206040518083038186803b158015610c1f57600080fd5b505afa158015610c33573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c579190611dff565b6001600160a01b031663c47f002783836040518363ffffffff1660e01b8152600401610c84929190611e1c565b602060405180830381600087803b158015610c9e57600080fd5b505af1158015610cb2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103279190611e38565b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3610d03816102c1610af3565b61077057806102d3610af3565b6000610d437f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b9050610d4e846112dd565b600083511180610d5b5750815b15610d6c57610d6a8484611392565b505b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd9143805460ff16610ecd57805460ff191660011781556040516001600160a01b0383166024820152610e1990869060440160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f3659cfe600000000000000000000000000000000000000000000000000000000179052611392565b50805460ff191681557f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b03838116911614610ec45760405162461bcd60e51b815260206004820152602f60248201527f45524331393637557067726164653a207570677261646520627265616b73206660448201527f75727468657220757067726164657300000000000000000000000000000000006064820152608401610313565b610ecd8561147d565b5050505050565b600054610100900460ff1680610eed575060005460ff16155b610f505760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610313565b600054610100900460ff16158015610f72576000805461ffff19166101011790555b610f8c6001600160a01b038316637965db0b60e01b61125e565b610fca576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b6044820152606401610313565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a28015610770576000805461ff00191690555050565b600054610100900460ff168061103f575060005460ff16155b6110a25760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610313565b600054610100900460ff161580156110c4576000805461ffff19166101011790555b6001600160a01b03821661111b5760405163eac0d38960e01b815260206004820152600660248201527f726f7574657200000000000000000000000000000000000000000000000000006044820152606401610313565b606580546001600160a01b0319166001600160a01b0384169081179091556040517f7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc8090600090a28015610770576000805461ff00191690555050565b600054610100900460ff1680611190575060005460ff16155b6111f35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610313565b600054610100900460ff16158015611215576000805461ffff19166101011790555b61121d6114bd565b6112256114bd565b8015610535576000805461ff001916905550565b6060610ba08383604051806060016040528060278152602001611ee46027913961156e565b600061126983611642565b8015610ba05750610ba0838361168d565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163314156112d8576000366112bb601482611e51565b6112c792369290611e68565b6112d091611e92565b60601c905090565b503390565b803b6113515760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e7472616374000000000000000000000000000000000000006064820152608401610313565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b6060823b6113f15760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608401610313565b600080846001600160a01b03168460405161140c9190611ec7565b600060405180830381855af49150503d8060008114611447576040519150601f19603f3d011682016040523d82523d6000602084013e61144c565b606091505b50915091506114748282604051806060016040528060278152602001611ee460279139611799565b95945050505050565b611486816112dd565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b600054610100900460ff16806114d6575060005460ff16155b6115395760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610313565b600054610100900460ff16158015611225576000805461ffff19166101011790558015610535576000805461ff001916905550565b6060833b6115cd5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608401610313565b600080856001600160a01b0316856040516115e89190611ec7565b600060405180830381855af49150503d8060008114611623576040519150601f19603f3d011682016040523d82523d6000602084013e611628565b606091505b5091509150611638828286611799565b9695505050505050565b6000611655826301ffc9a760e01b61168d565b80156109725750611686827fffffffff0000000000000000000000000000000000000000000000000000000061168d565b1592915050565b604080517fffffffff00000000000000000000000000000000000000000000000000000000831660248083019190915282518083039091018152604490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166301ffc9a760e01b179052905160009190829081906001600160a01b0387169061753090611721908690611ec7565b6000604051808303818686fa925050503d806000811461175d576040519150601f19603f3d011682016040523d82523d6000602084013e611762565b606091505b509150915060208151101561177d5760009350505050610972565b8180156116385750808060200190518101906116389190611ddd565b606083156117a8575081610ba0565b8251156117b85782518084602001fd5b8160405162461bcd60e51b81526004016103139190611976565b8280546117de90611bd2565b90600052602060002090601f0160209004810192826118005760008555611846565b82601f106118195782800160ff19823516178555611846565b82800160010185558215611846579182015b8281111561184657823582559160200191906001019061182b565b50611852929150611856565b5090565b5b808211156118525760008155600101611857565b6001600160a01b038116811461053557600080fd5b60008083601f84011261189257600080fd5b50813567ffffffffffffffff8111156118aa57600080fd5b6020830191508360208285010111156118c257600080fd5b9250929050565b6000806000604084860312156118de57600080fd5b83356118e98161186b565b9250602084013567ffffffffffffffff81111561190557600080fd5b61191186828701611880565b9497909650939450505050565b60005b83811015611939578181015183820152602001611921565b838111156103275750506000910152565b6000815180845261196281602086016020860161191e565b601f01601f19169290920160200192915050565b602081526000610ba0602083018461194a565b60006020828403121561199b57600080fd5b8135610ba08161186b565b600080604083850312156119b957600080fd5b82356119c48161186b565b915060208301356119d48161186b565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b60008060408385031215611a0857600080fd5b8235611a138161186b565b9150602083013567ffffffffffffffff80821115611a3057600080fd5b818501915085601f830112611a4457600080fd5b813581811115611a5657611a566119df565b604051601f8201601f19908116603f01168101908382118183101715611a7e57611a7e6119df565b81604052828152886020848701011115611a9757600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b60008060208385031215611acc57600080fd5b823567ffffffffffffffff811115611ae357600080fd5b611aef85828601611880565b90969095509350505050565b60008060208385031215611b0e57600080fd5b823567ffffffffffffffff80821115611b2657600080fd5b818501915085601f830112611b3a57600080fd5b813581811115611b4957600080fd5b8660208260051b8501011115611b5e57600080fd5b60209290920196919550909350505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015611bc557603f19888603018452611bb385835161194a565b94509285019290850190600101611b97565b5092979650505050505050565b600181811c90821680611be657607f821691505b60208210811415611c0757634e487b7160e01b600052602260045260246000fd5b50919050565b8183823760009101908152919050565b6000808354611c2b81611bd2565b60018281168015611c435760018114611c5457611c83565b60ff19841687528287019450611c83565b8760005260208060002060005b85811015611c7a5781548a820152908401908201611c61565b50505082870194505b50929695505050505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b604081526000611ccc604083018587611c8f565b60208382038185015260008554611ce281611bd2565b80855260018281168015611cfd5760018114611d1157611d3f565b60ff19841687870152604087019450611d3f565b896000528560002060005b84811015611d37578154898201890152908301908701611d1c565b880187019550505b50929a9950505050505050505050565b634e487b7160e01b600052603260045260246000fd5b6000808335601e19843603018112611d7c57600080fd5b83018035915067ffffffffffffffff821115611d9757600080fd5b6020019150368190038213156118c257600080fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415611dd657611dd6611dac565b5060010190565b600060208284031215611def57600080fd5b81518015158114610ba057600080fd5b600060208284031215611e1157600080fd5b8151610ba08161186b565b602081526000611e30602083018486611c8f565b949350505050565b600060208284031215611e4a57600080fd5b5051919050565b600082821015611e6357611e63611dac565b500390565b60008085851115611e7857600080fd5b83861115611e8557600080fd5b5050820193919092039150565b6bffffffffffffffffffffffff198135818116916014851015611ebf5780818660140360031b1b83161692505b505092915050565b60008251611ed981846020870161191e565b919091019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212201fd8076f00cce0f17726099b21892b42cff21c4fd9ce08fcf1d2abb1594a83d564736f6c63430008090033",
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

// SetScannerNodeVersion is a paid mutator transaction binding the contract method 0x94bb55a2.
//
// Solidity: function setScannerNodeVersion(string version) returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactor) SetScannerNodeVersion(opts *bind.TransactOpts, version string) (*types.Transaction, error) {
	return _ScannerNodeVersion.contract.Transact(opts, "setScannerNodeVersion", version)
}

// SetScannerNodeVersion is a paid mutator transaction binding the contract method 0x94bb55a2.
//
// Solidity: function setScannerNodeVersion(string version) returns()
func (_ScannerNodeVersion *ScannerNodeVersionSession) SetScannerNodeVersion(version string) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.SetScannerNodeVersion(&_ScannerNodeVersion.TransactOpts, version)
}

// SetScannerNodeVersion is a paid mutator transaction binding the contract method 0x94bb55a2.
//
// Solidity: function setScannerNodeVersion(string version) returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactorSession) SetScannerNodeVersion(version string) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.SetScannerNodeVersion(&_ScannerNodeVersion.TransactOpts, version)
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
