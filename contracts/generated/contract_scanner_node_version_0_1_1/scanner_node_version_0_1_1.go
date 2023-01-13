// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_scanner_node_version_0_1_1

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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"AccessManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newVersion\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldVersion\",\"type\":\"string\"}],\"name\":\"ScannerNodeBetaVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newVersion\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldVersion\",\"type\":\"string\"}],\"name\":\"ScannerNodeVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"disableRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__manager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scannerNodeBetaVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scannerNodeVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"setAccessManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ensRegistry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ensName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"setScannerNodeBetaVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"setScannerNodeVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040523060a0523480156200001557600080fd5b5060405162001fe938038062001fe9833981016040819052620000389162000180565b6001600160a01b038116608052600054610100900460ff1615808015620000665750600054600160ff909116105b8062000096575062000083306200017160201b62000b601760201c565b15801562000096575060005460ff166001145b620000fe5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b6000805460ff19166001179055801562000122576000805461ff0019166101001790555b801562000169576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050620001b2565b6001600160a01b03163b151590565b6000602082840312156200019357600080fd5b81516001600160a01b0381168114620001ab57600080fd5b9392505050565b60805160a051611dee620001fb6000396000818161041201528181610452015281816104f20152818161053201526105c50152600081816101e80152610fa80152611dee6000f3fe6080604052600436106100dd5760003560e01c8063572b6c051161007f578063c4d2b6dd11610059578063c4d2b6dd14610275578063c4d66de814610295578063c9580804146102b5578063d858a7e5146102d557600080fd5b8063572b6c05146101cb57806394bb55a214610228578063ac9650d81461024857600080fd5b80633659cfe6116100bb5780633659cfe6146101445780634f1ef2861461016457806352d1902d1461017757806354fd4d501461019a57600080fd5b80633121db1c146100e257806331d6653114610104578063345db3e11461012f575b600080fd5b3480156100ee57600080fd5b506101026100fd366004611664565b6102ea565b005b34801561011057600080fd5b5061011961036a565b6040516101269190611711565b60405180910390f35b34801561013b57600080fd5b506101196103f9565b34801561015057600080fd5b5061010261015f366004611724565b610407565b610102610172366004611757565b6104e7565b34801561018357600080fd5b5061018c6105b8565b604051908152602001610126565b3480156101a657600080fd5b5061011960405180604001604052806005815260200164302e312e3160d81b81525081565b3480156101d757600080fd5b506102186101e6366004611724565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0390811691161490565b6040519015158152602001610126565b34801561023457600080fd5b5061010261024336600461181b565b61066b565b34801561025457600080fd5b5061026861026336600461185d565b610760565b60405161012691906118d2565b34801561028157600080fd5b5061010261029036600461181b565b610856565b3480156102a157600080fd5b506101026102b0366004611724565b61094b565b3480156102c157600080fd5b506101026102d0366004611724565b610a17565b3480156102e157600080fd5b50610102610ad5565b7f664245c7af190fec316596e8231f724e8171b1966cfcd124347ac5a66c34f64a61031c81610317610b6f565b610b7e565b6103595780610329610b6f565b6040516301d4003760e61b815260048101929092526001600160a01b031660248201526044015b60405180910390fd5b610364848484610c0a565b50505050565b61012e805461037890611934565b80601f01602080910402602001604051908101604052809291908181526020018280546103a490611934565b80156103f15780601f106103c6576101008083540402835291602001916103f1565b820191906000526020600020905b8154815290600101906020018083116103d457829003601f168201915b505050505081565b61012d805461037890611934565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156104505760405162461bcd60e51b81526004016103509061196f565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610499600080516020611d72833981519152546001600160a01b031690565b6001600160a01b0316146104bf5760405162461bcd60e51b8152600401610350906119bb565b6104c881610d20565b604080516000808252602082019092526104e491839190610d5a565b50565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156105305760405162461bcd60e51b81526004016103509061196f565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610579600080516020611d72833981519152546001600160a01b031690565b6001600160a01b03161461059f5760405162461bcd60e51b8152600401610350906119bb565b6105a882610d20565b6105b482826001610d5a565b5050565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146106585760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610350565b50600080516020611d7283398151915290565b7f93bf097503ee765c5402c05999a7d54bac82299bf183ba1f5f2681ab2c6bd70c61069881610317610b6f565b6106a55780610329610b6f565b82826040516020016106b8929190611a07565b6040516020818303038152906040528051906020012061012d6040516020016106e19190611a17565b6040516020818303038152906040528051906020012014156107165760405163ae01a30160e01b815260040160405180910390fd5b7faa8f23ad4857a5d22df0189ebc6150d51b3152d8c621c3e8d24c66387897819b838361012d60405161074b93929190611ab2565b60405180910390a161036461012d848461156d565b60608167ffffffffffffffff81111561077b5761077b611741565b6040519080825280602002602001820160405280156107ae57816020015b60608152602001906001900390816107995790505b50905060005b8281101561084e5761081e308585848181106107d2576107d2611b49565b90506020028101906107e49190611b5f565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610ed992505050565b82828151811061083057610830611b49565b6020026020010181905250808061084690611bbc565b9150506107b4565b505b92915050565b7fcd2c75e410f265ca5b9c8bed70c109422b498bfc98f7df6ef360bcb407b1bedc61088381610317610b6f565b6108905780610329610b6f565b82826040516020016108a3929190611a07565b6040516020818303038152906040528051906020012061012e6040516020016108cc9190611a17565b6040516020818303038152906040528051906020012014156109015760405163ae01a30160e01b815260040160405180910390fd5b7fe5ac27c9fd9eefb9d9fd2b0878923cf7d44e611ead66cf80b7ccf184ac2927c0838361012e60405161093693929190611ab2565b60405180910390a161036461012e848461156d565b600054610100900460ff161580801561096b5750600054600160ff909116105b806109855750303b158015610985575060005460ff166001145b6109a15760405162461bcd60e51b815260040161035090611bd7565b6000805460ff1916600117905580156109c4576000805461ff0019166101001790555b6109cd82610efe565b80156105b4576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15050565b6000610a2581610317610b6f565b610a325780610329610b6f565b610a4c6001600160a01b038316637965db0b60e01b610f88565b610a8a576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b6044820152606401610350565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a25050565b6065546001600160a01b0316610b235760405163eac0d38960e01b81526020600482015260126024820152712fb232b83932b1b0ba32b22fb937baba32b960711b6044820152606401610350565b606580546001600160a01b03191690556040516000907f7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc80908290a2565b6001600160a01b03163b151590565b6000610b79610fa4565b905090565b603354604051632474521560e21b8152600481018490526001600160a01b03838116602483015260009216906391d148549060440160206040518083038186803b158015610bcb57600080fd5b505afa158015610bdf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c039190611c25565b9392505050565b6040516302571be360e01b81527f91d1777781884d03a6757a803996e38de2a42967fb37eeaca72729271025a9e260048201526001600160a01b038416906302571be39060240160206040518083038186803b158015610c6957600080fd5b505afa158015610c7d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ca19190611c47565b6001600160a01b031663c47f002783836040518363ffffffff1660e01b8152600401610cce929190611c64565b602060405180830381600087803b158015610ce857600080fd5b505af1158015610cfc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103649190611c80565b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3610d4d81610317610b6f565b6105b45780610329610b6f565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610d9257610d8d83611007565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610dcb57600080fd5b505afa925050508015610dfb575060408051601f3d908101601f19168201909252610df891810190611c80565b60015b610e5e5760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610350565b600080516020611d728339815191528114610ecd5760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610350565b50610d8d8383836110a3565b6060610c038383604051806060016040528060278152602001611d92602791396110c8565b600054610100900460ff1615808015610f1e5750600054600160ff909116105b80610f385750303b158015610f38575060005460ff166001145b610f545760405162461bcd60e51b815260040161035090611bd7565b6000805460ff191660011790558015610f77576000805461ff0019166101001790555b610f8082611166565b6109cd6112c0565b6000610f938361132d565b8015610c035750610c038383611360565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633141561100257600036610fe5601482611c99565b610ff192369290611cb0565b610ffa91611cda565b60601c905090565b503390565b6001600160a01b0381163b6110745760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610350565b600080516020611d7283398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6110ac8361143f565b6000825111806110b95750805b15610d8d57610364838361147f565b60606001600160a01b0384163b6110f15760405162461bcd60e51b815260040161035090611d0f565b600080856001600160a01b03168560405161110c9190611d55565b600060405180830381855af49150503d8060008114611147576040519150601f19603f3d011682016040523d82523d6000602084013e61114c565b606091505b509150915061115c828286611534565b9695505050505050565b600054610100900460ff16158080156111865750600054600160ff909116105b806111a05750303b1580156111a0575060005460ff166001145b6111bc5760405162461bcd60e51b815260040161035090611bd7565b6000805460ff1916600117905580156111df576000805461ff0019166101001790555b6111f96001600160a01b038316637965db0b60e01b610f88565b611237576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b6044820152606401610350565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a280156105b4576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001610a0b565b600054610100900460ff1661132b5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401610350565b565b6000611340826301ffc9a760e01b611360565b80156108505750611359826001600160e01b0319611360565b1592915050565b604080516001600160e01b0319831660248083019190915282518083039091018152604490910182526020810180516001600160e01b03166301ffc9a760e01b179052905160009190829081906001600160a01b03871690617530906113c7908690611d55565b6000604051808303818686fa925050503d8060008114611403576040519150601f19603f3d011682016040523d82523d6000602084013e611408565b606091505b50915091506020815110156114235760009350505050610850565b81801561115c57508080602001905181019061115c9190611c25565b61144881611007565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606001600160a01b0383163b6114a85760405162461bcd60e51b815260040161035090611d0f565b600080846001600160a01b0316846040516114c39190611d55565b600060405180830381855af49150503d80600081146114fe576040519150601f19603f3d011682016040523d82523d6000602084013e611503565b606091505b509150915061152b8282604051806060016040528060278152602001611d9260279139611534565b95945050505050565b60608315611543575081610c03565b8251156115535782518084602001fd5b8160405162461bcd60e51b81526004016103509190611711565b82805461157990611934565b90600052602060002090601f01602090048101928261159b57600085556115e1565b82601f106115b45782800160ff198235161785556115e1565b828001600101855582156115e1579182015b828111156115e15782358255916020019190600101906115c6565b506115ed9291506115f1565b5090565b5b808211156115ed57600081556001016115f2565b6001600160a01b03811681146104e457600080fd5b60008083601f84011261162d57600080fd5b50813567ffffffffffffffff81111561164557600080fd5b60208301915083602082850101111561165d57600080fd5b9250929050565b60008060006040848603121561167957600080fd5b833561168481611606565b9250602084013567ffffffffffffffff8111156116a057600080fd5b6116ac8682870161161b565b9497909650939450505050565b60005b838110156116d45781810151838201526020016116bc565b838111156103645750506000910152565b600081518084526116fd8160208601602086016116b9565b601f01601f19169290920160200192915050565b602081526000610c0360208301846116e5565b60006020828403121561173657600080fd5b8135610c0381611606565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561176a57600080fd5b823561177581611606565b9150602083013567ffffffffffffffff8082111561179257600080fd5b818501915085601f8301126117a657600080fd5b8135818111156117b8576117b8611741565b604051601f8201601f19908116603f011681019083821181831017156117e0576117e0611741565b816040528281528860208487010111156117f957600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b6000806020838503121561182e57600080fd5b823567ffffffffffffffff81111561184557600080fd5b6118518582860161161b565b90969095509350505050565b6000806020838503121561187057600080fd5b823567ffffffffffffffff8082111561188857600080fd5b818501915085601f83011261189c57600080fd5b8135818111156118ab57600080fd5b8660208260051b85010111156118c057600080fd5b60209290920196919550909350505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561192757603f198886030184526119158583516116e5565b945092850192908501906001016118f9565b5092979650505050505050565b600181811c9082168061194857607f821691505b6020821081141561196957634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b8183823760009101908152919050565b6000808354611a2581611934565b60018281168015611a3d5760018114611a4e57611a7d565b60ff19841687528287019450611a7d565b8760005260208060002060005b85811015611a745781548a820152908401908201611a5b565b50505082870194505b50929695505050505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b604081526000611ac6604083018587611a89565b60208382038185015260008554611adc81611934565b80855260018281168015611af75760018114611b0b57611b39565b60ff19841687870152604087019450611b39565b896000528560002060005b84811015611b31578154898201890152908301908701611b16565b880187019550505b50929a9950505050505050505050565b634e487b7160e01b600052603260045260246000fd5b6000808335601e19843603018112611b7657600080fd5b83018035915067ffffffffffffffff821115611b9157600080fd5b60200191503681900382131561165d57600080fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415611bd057611bd0611ba6565b5060010190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b600060208284031215611c3757600080fd5b81518015158114610c0357600080fd5b600060208284031215611c5957600080fd5b8151610c0381611606565b602081526000611c78602083018486611a89565b949350505050565b600060208284031215611c9257600080fd5b5051919050565b600082821015611cab57611cab611ba6565b500390565b60008085851115611cc057600080fd5b83861115611ccd57600080fd5b5050820193919092039150565b6bffffffffffffffffffffffff198135818116916014851015611d075780818660140360031b1b83161692505b505092915050565b60208082526026908201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6040820152651b9d1c9858dd60d21b606082015260800190565b60008251611d678184602087016116b9565b919091019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220026035f969c8f1e92efd36d45f8e0c91edab83fc3d39d1d1e522f455260692fe64736f6c63430008090033",
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

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ScannerNodeVersion *ScannerNodeVersionCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ScannerNodeVersion.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ScannerNodeVersion *ScannerNodeVersionSession) ProxiableUUID() ([32]byte, error) {
	return _ScannerNodeVersion.Contract.ProxiableUUID(&_ScannerNodeVersion.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ScannerNodeVersion *ScannerNodeVersionCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ScannerNodeVersion.Contract.ProxiableUUID(&_ScannerNodeVersion.CallOpts)
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

// DisableRouter is a paid mutator transaction binding the contract method 0xd858a7e5.
//
// Solidity: function disableRouter() returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactor) DisableRouter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScannerNodeVersion.contract.Transact(opts, "disableRouter")
}

// DisableRouter is a paid mutator transaction binding the contract method 0xd858a7e5.
//
// Solidity: function disableRouter() returns()
func (_ScannerNodeVersion *ScannerNodeVersionSession) DisableRouter() (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.DisableRouter(&_ScannerNodeVersion.TransactOpts)
}

// DisableRouter is a paid mutator transaction binding the contract method 0xd858a7e5.
//
// Solidity: function disableRouter() returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactorSession) DisableRouter() (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.DisableRouter(&_ScannerNodeVersion.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address __manager) returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactor) Initialize(opts *bind.TransactOpts, __manager common.Address) (*types.Transaction, error) {
	return _ScannerNodeVersion.contract.Transact(opts, "initialize", __manager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address __manager) returns()
func (_ScannerNodeVersion *ScannerNodeVersionSession) Initialize(__manager common.Address) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.Initialize(&_ScannerNodeVersion.TransactOpts, __manager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address __manager) returns()
func (_ScannerNodeVersion *ScannerNodeVersionTransactorSession) Initialize(__manager common.Address) (*types.Transaction, error) {
	return _ScannerNodeVersion.Contract.Initialize(&_ScannerNodeVersion.TransactOpts, __manager)
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

// ScannerNodeVersionInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ScannerNodeVersion contract.
type ScannerNodeVersionInitializedIterator struct {
	Event *ScannerNodeVersionInitialized // Event containing the contract specifics and raw log

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
func (it *ScannerNodeVersionInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScannerNodeVersionInitialized)
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
		it.Event = new(ScannerNodeVersionInitialized)
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
func (it *ScannerNodeVersionInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScannerNodeVersionInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScannerNodeVersionInitialized represents a Initialized event raised by the ScannerNodeVersion contract.
type ScannerNodeVersionInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) FilterInitialized(opts *bind.FilterOpts) (*ScannerNodeVersionInitializedIterator, error) {

	logs, sub, err := _ScannerNodeVersion.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ScannerNodeVersionInitializedIterator{contract: _ScannerNodeVersion.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ScannerNodeVersionInitialized) (event.Subscription, error) {

	logs, sub, err := _ScannerNodeVersion.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScannerNodeVersionInitialized)
				if err := _ScannerNodeVersion.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ScannerNodeVersion *ScannerNodeVersionFilterer) ParseInitialized(log types.Log) (*ScannerNodeVersionInitialized, error) {
	event := new(ScannerNodeVersionInitialized)
	if err := _ScannerNodeVersion.contract.UnpackLog(event, "Initialized", log); err != nil {
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
