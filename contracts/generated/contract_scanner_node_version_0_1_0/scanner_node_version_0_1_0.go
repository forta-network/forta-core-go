// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_scanner_node_version_0_1_0

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
	Bin: "0x60c06040523060a0523480156200001557600080fd5b5060405162001d3738038062001d37833981016040819052620000389162000105565b6001600160a01b038116608052600054610100900460ff16806200005f575060005460ff16155b620000c75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b600054610100900460ff16158015620000ea576000805461ffff19166101011790555b8015620000fd576000805461ff00191690555b505062000137565b6000602082840312156200011857600080fd5b81516001600160a01b03811681146200013057600080fd5b9392505050565b60805160a051611bbe6200017960003960008181610376015281816103b6015281816104c7015261050701526000818161018f0152610eea0152611bbe6000f3fe60806040526004361061009c5760003560e01c806354fd4d501161006457806354fd4d5014610141578063572b6c051461017257806394bb55a2146101cf578063ac9650d8146101ef578063c0d786551461021c578063c95808041461023c57600080fd5b80633121db1c146100a1578063345db3e1146100c35780633659cfe6146100ee578063485cc9551461010e5780634f1ef2861461012e575b600080fd5b3480156100ad57600080fd5b506100c16100bc36600461141b565b61025c565b005b3480156100cf57600080fd5b506100d86102dc565b6040516100e591906114c8565b60405180910390f35b3480156100fa57600080fd5b506100c16101093660046114db565b61036b565b34801561011a57600080fd5b506100c16101293660046114f8565b610434565b6100c161013c366004611547565b6104bc565b34801561014d57600080fd5b506100d8604051806040016040528060058152602001640302e312e360dc1b81525081565b34801561017e57600080fd5b506101bf61018d3660046114db565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0390811691161490565b60405190151581526020016100e5565b3480156101db57600080fd5b506100c16101ea36600461160b565b610576565b3480156101fb57600080fd5b5061020f61020a36600461164d565b61066b565b6040516100e591906116c2565b34801561022857600080fd5b506100c16102373660046114db565b610761565b34801561024857600080fd5b506100c16102573660046114db565b61080a565b7f664245c7af190fec316596e8231f724e8171b1966cfcd124347ac5a66c34f64a61028e816102896108c8565b6108d7565b6102cb578061029b6108c8565b6040516301d4003760e61b815260048101929092526001600160a01b031660248201526044015b60405180910390fd5b6102d6848484610963565b50505050565b61012d80546102ea90611724565b80601f016020809104026020016040519081016040528092919081815260200182805461031690611724565b80156103635780601f1061033857610100808354040283529160200191610363565b820191906000526020600020905b81548152906001019060200180831161034657829003601f168201915b505050505081565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156103b45760405162461bcd60e51b81526004016102c29061175f565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166103e6610a79565b6001600160a01b03161461040c5760405162461bcd60e51b81526004016102c2906117ab565b61041581610aa7565b6040805160008082526020820190925261043191839190610ae1565b50565b600054610100900460ff168061044d575060005460ff16155b6104695760405162461bcd60e51b81526004016102c2906117f7565b600054610100900460ff1615801561048b576000805461ffff19166101011790555b61049483610c2c565b61049d82610d37565b6104a5610e2a565b80156104b7576000805461ff00191690555b505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156105055760405162461bcd60e51b81526004016102c29061175f565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610537610a79565b6001600160a01b03161461055d5760405162461bcd60e51b81526004016102c2906117ab565b61056682610aa7565b61057282826001610ae1565b5050565b7f93bf097503ee765c5402c05999a7d54bac82299bf183ba1f5f2681ab2c6bd70c6105a3816102896108c8565b6105b0578061029b6108c8565b82826040516020016105c3929190611845565b6040516020818303038152906040528051906020012061012d6040516020016105ec9190611855565b6040516020818303038152906040528051906020012014156106215760405163ae01a30160e01b815260040160405180910390fd5b7faa8f23ad4857a5d22df0189ebc6150d51b3152d8c621c3e8d24c66387897819b838361012d604051610656939291906118f0565b60405180910390a16102d661012d8484611324565b60608167ffffffffffffffff81111561068657610686611531565b6040519080825280602002602001820160405280156106b957816020015b60608152602001906001900390816106a45790505b50905060005b8281101561075957610729308585848181106106dd576106dd611987565b90506020028101906106ef919061199d565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610ea592505050565b82828151811061073b5761073b611987565b60200260200101819052508080610751906119fa565b9150506106bf565b505b92915050565b600061076f816102896108c8565b61077c578061029b6108c8565b6001600160a01b0382166107bf5760405163eac0d38960e01b81526020600482015260096024820152683732bba937baba32b960b91b60448201526064016102c2565b606580546001600160a01b0319166001600160a01b0384169081179091556040517f7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc8090600090a25050565b6000610818816102896108c8565b610825578061029b6108c8565b61083f6001600160a01b038316637965db0b60e01b610eca565b61087d576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b60448201526064016102c2565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a25050565b60006108d2610ee6565b905090565b603354604051632474521560e21b8152600481018490526001600160a01b03838116602483015260009216906391d148549060440160206040518083038186803b15801561092457600080fd5b505afa158015610938573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095c9190611a15565b9392505050565b6040516302571be360e01b81527f91d1777781884d03a6757a803996e38de2a42967fb37eeaca72729271025a9e260048201526001600160a01b038416906302571be39060240160206040518083038186803b1580156109c257600080fd5b505afa1580156109d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109fa9190611a37565b6001600160a01b031663c47f002783836040518363ffffffff1660e01b8152600401610a27929190611a54565b602060405180830381600087803b158015610a4157600080fd5b505af1158015610a55573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102d69190611a70565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3610ad4816102896108c8565b610572578061029b6108c8565b6000610aeb610a79565b9050610af684610f49565b600083511180610b035750815b15610b1457610b128484610fee565b505b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd9143805460ff16610c2557805460ff191660011781556040516001600160a01b0383166024820152610b9390869060440160408051601f198184030181529190526020810180516001600160e01b0316631b2ce7f360e11b179052610fee565b50805460ff19168155610ba4610a79565b6001600160a01b0316826001600160a01b031614610c1c5760405162461bcd60e51b815260206004820152602f60248201527f45524331393637557067726164653a207570677261646520627265616b73206660448201526e75727468657220757067726164657360881b60648201526084016102c2565b610c258561109a565b5050505050565b600054610100900460ff1680610c45575060005460ff16155b610c615760405162461bcd60e51b81526004016102c2906117f7565b600054610100900460ff16158015610c83576000805461ffff19166101011790555b610c9d6001600160a01b038316637965db0b60e01b610eca565b610cdb576040516301a1fdbb60e41b815260206004820152600e60248201526d125058d8d95cdcd0dbdb9d1c9bdb60921b60448201526064016102c2565b603380546001600160a01b0319166001600160a01b0384169081179091556040517fa5bc17e575e3b53b23d0e93e121a5a66d1de4d5edb4dfde6027b14d79b7f2b9c90600090a28015610572576000805461ff00191690555050565b600054610100900460ff1680610d50575060005460ff16155b610d6c5760405162461bcd60e51b81526004016102c2906117f7565b600054610100900460ff16158015610d8e576000805461ffff19166101011790555b6001600160a01b038216610dce5760405163eac0d38960e01b81526020600482015260066024820152653937baba32b960d11b60448201526064016102c2565b606580546001600160a01b0319166001600160a01b0384169081179091556040517f7aed1d3e8155a07ccf395e44ea3109a0e2d6c9b29bbbe9f142d9790596f4dc8090600090a28015610572576000805461ff00191690555050565b600054610100900460ff1680610e43575060005460ff16155b610e5f5760405162461bcd60e51b81526004016102c2906117f7565b600054610100900460ff16158015610e81576000805461ffff19166101011790555b610e896110da565b610e916110da565b8015610431576000805461ff001916905550565b606061095c8383604051806060016040528060278152602001611b6260279139611144565b6000610ed5836111d9565b801561095c575061095c838361120c565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316331415610f4457600036610f27601482611a89565b610f3392369290611aa0565b610f3c91611aca565b60601c905090565b503390565b803b610fad5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084016102c2565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b6060823b61100e5760405162461bcd60e51b81526004016102c290611aff565b600080846001600160a01b0316846040516110299190611b45565b600060405180830381855af49150503d8060008114611064576040519150601f19603f3d011682016040523d82523d6000602084013e611069565b606091505b50915091506110918282604051806060016040528060278152602001611b62602791396112eb565b95945050505050565b6110a381610f49565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b600054610100900460ff16806110f3575060005460ff16155b61110f5760405162461bcd60e51b81526004016102c2906117f7565b600054610100900460ff16158015610e91576000805461ffff19166101011790558015610431576000805461ff001916905550565b6060833b6111645760405162461bcd60e51b81526004016102c290611aff565b600080856001600160a01b03168560405161117f9190611b45565b600060405180830381855af49150503d80600081146111ba576040519150601f19603f3d011682016040523d82523d6000602084013e6111bf565b606091505b50915091506111cf8282866112eb565b9695505050505050565b60006111ec826301ffc9a760e01b61120c565b801561075b5750611205826001600160e01b031961120c565b1592915050565b604080516001600160e01b0319831660248083019190915282518083039091018152604490910182526020810180516001600160e01b03166301ffc9a760e01b179052905160009190829081906001600160a01b0387169061753090611273908690611b45565b6000604051808303818686fa925050503d80600081146112af576040519150601f19603f3d011682016040523d82523d6000602084013e6112b4565b606091505b50915091506020815110156112cf576000935050505061075b565b8180156111cf5750808060200190518101906111cf9190611a15565b606083156112fa57508161095c565b82511561130a5782518084602001fd5b8160405162461bcd60e51b81526004016102c291906114c8565b82805461133090611724565b90600052602060002090601f0160209004810192826113525760008555611398565b82601f1061136b5782800160ff19823516178555611398565b82800160010185558215611398579182015b8281111561139857823582559160200191906001019061137d565b506113a49291506113a8565b5090565b5b808211156113a457600081556001016113a9565b6001600160a01b038116811461043157600080fd5b60008083601f8401126113e457600080fd5b50813567ffffffffffffffff8111156113fc57600080fd5b60208301915083602082850101111561141457600080fd5b9250929050565b60008060006040848603121561143057600080fd5b833561143b816113bd565b9250602084013567ffffffffffffffff81111561145757600080fd5b611463868287016113d2565b9497909650939450505050565b60005b8381101561148b578181015183820152602001611473565b838111156102d65750506000910152565b600081518084526114b4816020860160208601611470565b601f01601f19169290920160200192915050565b60208152600061095c602083018461149c565b6000602082840312156114ed57600080fd5b813561095c816113bd565b6000806040838503121561150b57600080fd5b8235611516816113bd565b91506020830135611526816113bd565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561155a57600080fd5b8235611565816113bd565b9150602083013567ffffffffffffffff8082111561158257600080fd5b818501915085601f83011261159657600080fd5b8135818111156115a8576115a8611531565b604051601f8201601f19908116603f011681019083821181831017156115d0576115d0611531565b816040528281528860208487010111156115e957600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b6000806020838503121561161e57600080fd5b823567ffffffffffffffff81111561163557600080fd5b611641858286016113d2565b90969095509350505050565b6000806020838503121561166057600080fd5b823567ffffffffffffffff8082111561167857600080fd5b818501915085601f83011261168c57600080fd5b81358181111561169b57600080fd5b8660208260051b85010111156116b057600080fd5b60209290920196919550909350505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561171757603f1988860301845261170585835161149c565b945092850192908501906001016116e9565b5092979650505050505050565b600181811c9082168061173857607f821691505b6020821081141561175957634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b8183823760009101908152919050565b600080835461186381611724565b6001828116801561187b576001811461188c576118bb565b60ff198416875282870194506118bb565b8760005260208060002060005b858110156118b25781548a820152908401908201611899565b50505082870194505b50929695505050505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6040815260006119046040830185876118c7565b6020838203818501526000855461191a81611724565b80855260018281168015611935576001811461194957611977565b60ff19841687870152604087019450611977565b896000528560002060005b8481101561196f578154898201890152908301908701611954565b880187019550505b50929a9950505050505050505050565b634e487b7160e01b600052603260045260246000fd5b6000808335601e198436030181126119b457600080fd5b83018035915067ffffffffffffffff8211156119cf57600080fd5b60200191503681900382131561141457600080fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415611a0e57611a0e6119e4565b5060010190565b600060208284031215611a2757600080fd5b8151801515811461095c57600080fd5b600060208284031215611a4957600080fd5b815161095c816113bd565b602081526000611a686020830184866118c7565b949350505050565b600060208284031215611a8257600080fd5b5051919050565b600082821015611a9b57611a9b6119e4565b500390565b60008085851115611ab057600080fd5b83861115611abd57600080fd5b5050820193919092039150565b6bffffffffffffffffffffffff198135818116916014851015611af75780818660140360031b1b83161692505b505092915050565b60208082526026908201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6040820152651b9d1c9858dd60d21b606082015260800190565b60008251611b57818460208701611470565b919091019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220eccf7c95276842d185d05858934a58ac3728ad0347a29746671014c20eeed3f264736f6c63430008090033",
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
