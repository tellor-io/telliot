// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72305820b0b25139b75e676b66895a7d1fb0a1f523d8c20a178651d70497b0640c4406c364736f6c634300050a0032`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// TellorDisputeABI is the input ABI used to generate the binding from.
const TellorDisputeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"NewDispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_position\",\"type\":\"bool\"},{\"indexed\":true,\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_result\",\"type\":\"int256\"},{\"indexed\":true,\"name\":\"_reportedMiner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_reportingParty\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"DisputeVoteTallied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_newTellor\",\"type\":\"address\"}],\"name\":\"NewTellorAddress\",\"type\":\"event\"}]"

// TellorDisputeBin is the compiled bytecode used for deploying new contracts.
const TellorDisputeBin = `0x6115f2610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100615760003560e01c80632da0706e14610066578063694bf49f146100a0578063ca9a4ea5146100d9578063def6fac714610115578063e15f6f7014610145575b600080fd5b81801561007257600080fd5b5061009e6004803603606081101561008957600080fd5b5080359060208101359060400135151561016f565b005b8180156100ac57600080fd5b5061009e600480360360408110156100c357600080fd5b50803590602001356001600160a01b0316610362565b8180156100e557600080fd5b5061009e600480360360808110156100fc57600080fd5b508035906020810135906040810135906060013561079d565b81801561012157600080fd5b5061009e6004803603604081101561013857600080fd5b5080359060200135610dda565b81801561015157600080fd5b5061009e6004803603602081101561016857600080fd5b50356113e7565b600082815260448085016020908152604080842081516a313637b1b5a73ab6b132b960a91b8152825190819003600b018120865260058201845282862054633f48b1ff60e01b8252600482018a905233602483015294810194909452905190939273__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__92633f48b1ff92606480840193829003018186803b15801561020657600080fd5b505af415801561021a573d6000803e3d6000fd5b505050506040513d602081101561023057600080fd5b505133600090815260068401602052604090205490915060ff1615156001141561025957600080fd5b6000811161026657600080fd5b3360009081526047860160205260409020546003141561028557600080fd5b336000908152600680840160209081526040808420805460ff1916600190811790915581516c6e756d6265724f66566f74657360981b8152825190819003600d01812086526005880180855283872080549093019092556571756f72756d60d01b81528251908190039094019093208452919052902080548201905582156103165760018201805482019055610322565b60018201805482900390555b6040805184151581529051339186917f86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae09181900360200190a35050505050565b604080516bffffffffffffffffffffffff19606084901b1660208083019190915282518083036014018152603490920183528151918101919091206000818152604a860190925291902054156103b757600080fd5b60408051696469737075746546656560b01b8152815190819003600a01812060009081528286016020528281205463c7bb46ad60e01b8352600483018790523360248401523060448401526064830152915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9263c7bb46ad9260848082019391829003018186803b15801561044057600080fd5b505af4158015610454573d6000803e3d6000fd5b5050505082604001600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c0190506040518091039020815260200190815260200160002060008154809291906001019190505550600083604001600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c019050604051809103902081526020019081526020016000205490508084604a0160008481526020019081526020016000208190555060405180610100016040528083815260200160008152602001600015158152602001600015158152602001600115158152602001336001600160a01b03168152602001336001600160a01b03168152602001846001600160a01b0316815250846044016000838152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548160ff02191690831515021790555060608201518160020160016101000a81548160ff02191690831515021790555060808201518160020160026101000a81548160ff02191690831515021790555060a08201518160020160036101000a8154816001600160a01b0302191690836001600160a01b0316021790555060c08201518160030160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060e08201518160040160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555090505043846044016000838152602001908152602001600020600501600060405180806a313637b1b5a73ab6b132b960a91b815250600b01905060405180910390208152602001908152602001600020819055508360400160006040518080696469737075746546656560b01b815250600a0190506040518091039020815260200190815260200160002054846044016000838152602001908152602001600020600501600060405180806266656560e81b815250600301905060405180910390208152602001908152602001600020819055504262093a8001846044016000838152602001908152602001600020600501600060405180806f6d696e457865637574696f6e4461746560801b8152506010019050604051809103902081526020019081526020016000208190555050505050565b6000838152604885016020908152604080832085845260058101909252909120546090439190910311156107d057600080fd5b60008381526005820160205260409020546107ea57600080fd5b600582106107f757600080fd5b60008381526008820160205260408120836005811061081257fe5b0154604080516bffffffffffffffffffffffff19606084901b1660208083019190915260348201899052605480830189905283518084039091018152607490920183528151918101919091206000818152604a8b01909252919020546001600160a01b039092169250901561088657600080fd5b60408051696469737075746546656560b01b8152815190819003600a0181206000908152828a016020528281205463c7bb46ad60e01b8352600483018b90523360248401523060448401526064830152915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9263c7bb46ad9260848082019391829003018186803b15801561090f57600080fd5b505af4158015610923573d6000803e3d6000fd5b5050505086604001600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c019050604051809103902081526020019081526020016000205460010187604001600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c0190506040518091039020815260200190815260200160002081905550600087604001600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c019050604051809103902081526020019081526020016000205490508088604a0160008481526020019081526020016000208190555060405180610100016040528083815260200160008152602001600015158152602001600015158152602001600015158152602001846001600160a01b03168152602001336001600160a01b0316815260200160006001600160a01b0316815250886044016000838152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548160ff02191690831515021790555060608201518160020160016101000a81548160ff02191690831515021790555060808201518160020160026101000a81548160ff02191690831515021790555060a08201518160020160036101000a8154816001600160a01b0302191690836001600160a01b0316021790555060c08201518160030160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060e08201518160040160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050508688604401600083815260200190815260200160002060050160006040518080681c995c5d595cdd125960ba1b81525060090190506040518091039020815260200190815260200160002081905550858860440160008381526020019081526020016000206005016000604051808068074696d657374616d760bc1b815250600901905060405180910390208152602001908152602001600020819055508360090160008781526020019081526020016000208560058110610c2a57fe5b0154600082815260448a016020818152604080842081516476616c756560d81b8152825190819003600590810182208752909101808452828620969096558685528383526f6d696e457865637574696f6e4461746560801b81528151908190036010018120855285835281852062093a80420190558685528383526a313637b1b5a73ab6b132b960a91b8152815190819003600b0181208552858352818520439055868552838352681b5a5b995c94db1bdd60ba1b8152815190819003600901812085528583528185208b9055696469737075746546656560b01b8152815190819003600a0181208552818e018352818520548786529383526266656560e81b815281519081900360030190208452939052919020556002851415610d7357600087815260488901602090815260408083208984526007019091529020805460ff191660011790555b6001600160a01b038316600081815260478a016020908152604091829020600390558151898152908101929092528051899284927feceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da6492918290030190a35050505050505050565b600081815260448301602090815260408083208151681c995c5d595cdd125960ba1b81528251908190036009019020845260058101835281842054845260488601909252909120600282015460ff1615610e3357600080fd5b604080516f6d696e457865637574696f6e4461746560801b8152815190819003601001902060009081526005840160205220544211610e7157600080fd5b600282015462010000900460ff16611258576002820154630100000090046001600160a01b03166000908152604785016020526040812060018401549091121561110d576000815562015180420642036001820155604080516a1cdd185ad95c90dbdd5b9d60aa1b8152815190819003600b0190206000908152818701602052208054600019019055610f03856113e7565b60028301546003840154604080516a1cdd185ad9505b5bdd5b9d60aa1b8152815190819003600b0181206000908152828a016020528281205463c7bb46ad60e01b8352600483018b90526001600160a01b0363010000009096048616602484015293909416604482015260648101929092525173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9263c7bb46ad9260848082019391829003018186803b158015610fad57600080fd5b505af4158015610fc1573d6000803e3d6000fd5b505050600380850154604080516266656560e81b815281519081900390930183206000908152600588016020528181205463c7bb46ad60e01b8552600485018b90523060248601526001600160a01b03909316604485015260648401929092525173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__935063c7bb46ad926084808201939291829003018186803b15801561105b57600080fd5b505af415801561106f573d6000803e3d6000fd5b50505060028401805461ff001916610100179055506040805168074696d657374616d760bc1b815281519081900360090190206000908152600585016020908152828220548252600785019052205460ff16151560011415611108576040805168074696d657374616d760bc1b815281519081900360090190206000908152600585016020908152828220548252600685019052908120555b611252565b600181556002830154604080516266656560e81b815281519081900360030181206000908152600587016020528281205463c7bb46ad60e01b8352600483018a90523060248401526001600160a01b0363010000009095049490941660448301526064820193909352905173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9263c7bb46ad9260848082019391829003018186803b1580156111af57600080fd5b505af41580156111c3573d6000803e3d6000fd5b50506040805168074696d657374616d760bc1b815281519081900360090190206000908152600587016020908152828220548252600787019052205460ff1615156001141591506112529050576040805168074696d657374616d760bc1b81528151908190036009019020600090815260058501602090815282822054825260078501905220805460ff191690555b50611369565b60008260010154131561136957604080516b746f74616c5f737570706c7960a01b8152815190819003600c01902060009081528186016020522054606490601402604080516571756f72756d60d01b815281519081900360060190206000908152600586016020522054919004106112cf57600080fd5b600482018054604080516d1d195b1b1bdc90dbdb9d1c9858dd60921b8152815190819003600e0181206000908152603f890160209081529083902080546001600160a01b0319166001600160a01b0395861617905560028701805461ff00191661010017905593549092168252517fc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d929181900390910190a15b60028201805460ff19166001908117918290558301546003840154604080519283526001600160a01b039182166020840152610100840460ff16151583820152516301000000909304169185917f21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739919081900360600190a350505050565b604080516b7461726765744d696e65727360a01b8152815190819003600c01812060009081528284016020818152848320546a1cdd185ad95c90dbdd5b9d60aa1b8552855194859003600b0190942083525291909120546103e8919082028161144c57fe5b04101561154657604080516b7461726765744d696e65727360a01b8152815190819003600c01812060009081528284016020818152848320546a1cdd185ad95c90dbdd5b9d60aa1b8552855194859003600b0190942083525291909120546115159167d02ab486cedc0000916103e891611508918302816114c957fe5b604080516a1cdd185ad9505b5bdd5b9d60aa1b8152815190819003600b019020600090815281890160205220549190046103e80363ffffffff61158016565b8161150f57fe5b046115a7565b60408051696469737075746546656560b01b8152815190819003600a0190206000908152818401602052205561157d565b60408051696469737075746546656560b01b8152815190819003600a01902060009081528183016020522067d02ab486cedc000090555b50565b600082820283158061159a57508284828161159757fe5b04145b6115a057fe5b9392505050565b60008183116115b657816115a0565b509091905056fea265627a7a72305820dadb2b8460cb9a7746c8f349d0279e13a2fc595a14a9d398772ad34d573d450064736f6c634300050a0032`

// DeployTellorDispute deploys a new Ethereum contract, binding an instance of TellorDispute to it.
func DeployTellorDispute(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorDispute, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorDisputeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TellorDisputeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TellorDispute{TellorDisputeCaller: TellorDisputeCaller{contract: contract}, TellorDisputeTransactor: TellorDisputeTransactor{contract: contract}, TellorDisputeFilterer: TellorDisputeFilterer{contract: contract}}, nil
}

// TellorDispute is an auto generated Go binding around an Ethereum contract.
type TellorDispute struct {
	TellorDisputeCaller     // Read-only binding to the contract
	TellorDisputeTransactor // Write-only binding to the contract
	TellorDisputeFilterer   // Log filterer for contract events
}

// TellorDisputeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorDisputeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorDisputeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorDisputeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorDisputeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorDisputeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorDisputeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorDisputeSession struct {
	Contract     *TellorDispute    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorDisputeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorDisputeCallerSession struct {
	Contract *TellorDisputeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TellorDisputeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorDisputeTransactorSession struct {
	Contract     *TellorDisputeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TellorDisputeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorDisputeRaw struct {
	Contract *TellorDispute // Generic contract binding to access the raw methods on
}

// TellorDisputeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorDisputeCallerRaw struct {
	Contract *TellorDisputeCaller // Generic read-only contract binding to access the raw methods on
}

// TellorDisputeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorDisputeTransactorRaw struct {
	Contract *TellorDisputeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellorDispute creates a new instance of TellorDispute, bound to a specific deployed contract.
func NewTellorDispute(address common.Address, backend bind.ContractBackend) (*TellorDispute, error) {
	contract, err := bindTellorDispute(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TellorDispute{TellorDisputeCaller: TellorDisputeCaller{contract: contract}, TellorDisputeTransactor: TellorDisputeTransactor{contract: contract}, TellorDisputeFilterer: TellorDisputeFilterer{contract: contract}}, nil
}

// NewTellorDisputeCaller creates a new read-only instance of TellorDispute, bound to a specific deployed contract.
func NewTellorDisputeCaller(address common.Address, caller bind.ContractCaller) (*TellorDisputeCaller, error) {
	contract, err := bindTellorDispute(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorDisputeCaller{contract: contract}, nil
}

// NewTellorDisputeTransactor creates a new write-only instance of TellorDispute, bound to a specific deployed contract.
func NewTellorDisputeTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorDisputeTransactor, error) {
	contract, err := bindTellorDispute(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorDisputeTransactor{contract: contract}, nil
}

// NewTellorDisputeFilterer creates a new log filterer instance of TellorDispute, bound to a specific deployed contract.
func NewTellorDisputeFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorDisputeFilterer, error) {
	contract, err := bindTellorDispute(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorDisputeFilterer{contract: contract}, nil
}

// bindTellorDispute binds a generic wrapper to an already deployed contract.
func bindTellorDispute(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorDisputeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorDispute *TellorDisputeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TellorDispute.Contract.TellorDisputeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorDispute *TellorDisputeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorDispute.Contract.TellorDisputeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorDispute *TellorDisputeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorDispute.Contract.TellorDisputeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorDispute *TellorDisputeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TellorDispute.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorDispute *TellorDisputeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorDispute.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorDispute *TellorDisputeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorDispute.Contract.contract.Transact(opts, method, params...)
}

// TellorDisputeDisputeVoteTalliedIterator is returned from FilterDisputeVoteTallied and is used to iterate over the raw logs and unpacked data for DisputeVoteTallied events raised by the TellorDispute contract.
type TellorDisputeDisputeVoteTalliedIterator struct {
	Event *TellorDisputeDisputeVoteTallied // Event containing the contract specifics and raw log

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
func (it *TellorDisputeDisputeVoteTalliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorDisputeDisputeVoteTallied)
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
		it.Event = new(TellorDisputeDisputeVoteTallied)
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
func (it *TellorDisputeDisputeVoteTalliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorDisputeDisputeVoteTalliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorDisputeDisputeVoteTallied represents a DisputeVoteTallied event raised by the TellorDispute contract.
type TellorDisputeDisputeVoteTallied struct {
	DisputeID      *big.Int
	Result         *big.Int
	ReportedMiner  common.Address
	ReportingParty common.Address
	Active         bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDisputeVoteTallied is a free log retrieval operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: e DisputeVoteTallied(_disputeID indexed uint256, _result int256, _reportedMiner indexed address, _reportingParty address, _active bool)
func (_TellorDispute *TellorDisputeFilterer) FilterDisputeVoteTallied(opts *bind.FilterOpts, _disputeID []*big.Int, _reportedMiner []common.Address) (*TellorDisputeDisputeVoteTalliedIterator, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _reportedMinerRule []interface{}
	for _, _reportedMinerItem := range _reportedMiner {
		_reportedMinerRule = append(_reportedMinerRule, _reportedMinerItem)
	}

	logs, sub, err := _TellorDispute.contract.FilterLogs(opts, "DisputeVoteTallied", _disputeIDRule, _reportedMinerRule)
	if err != nil {
		return nil, err
	}
	return &TellorDisputeDisputeVoteTalliedIterator{contract: _TellorDispute.contract, event: "DisputeVoteTallied", logs: logs, sub: sub}, nil
}

// WatchDisputeVoteTallied is a free log subscription operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: e DisputeVoteTallied(_disputeID indexed uint256, _result int256, _reportedMiner indexed address, _reportingParty address, _active bool)
func (_TellorDispute *TellorDisputeFilterer) WatchDisputeVoteTallied(opts *bind.WatchOpts, sink chan<- *TellorDisputeDisputeVoteTallied, _disputeID []*big.Int, _reportedMiner []common.Address) (event.Subscription, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _reportedMinerRule []interface{}
	for _, _reportedMinerItem := range _reportedMiner {
		_reportedMinerRule = append(_reportedMinerRule, _reportedMinerItem)
	}

	logs, sub, err := _TellorDispute.contract.WatchLogs(opts, "DisputeVoteTallied", _disputeIDRule, _reportedMinerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorDisputeDisputeVoteTallied)
				if err := _TellorDispute.contract.UnpackLog(event, "DisputeVoteTallied", log); err != nil {
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

// TellorDisputeNewDisputeIterator is returned from FilterNewDispute and is used to iterate over the raw logs and unpacked data for NewDispute events raised by the TellorDispute contract.
type TellorDisputeNewDisputeIterator struct {
	Event *TellorDisputeNewDispute // Event containing the contract specifics and raw log

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
func (it *TellorDisputeNewDisputeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorDisputeNewDispute)
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
		it.Event = new(TellorDisputeNewDispute)
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
func (it *TellorDisputeNewDisputeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorDisputeNewDisputeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorDisputeNewDispute represents a NewDispute event raised by the TellorDispute contract.
type TellorDisputeNewDispute struct {
	DisputeId *big.Int
	RequestId *big.Int
	Timestamp *big.Int
	Miner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewDispute is a free log retrieval operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: e NewDispute(_disputeId indexed uint256, _requestId indexed uint256, _timestamp uint256, _miner address)
func (_TellorDispute *TellorDisputeFilterer) FilterNewDispute(opts *bind.FilterOpts, _disputeId []*big.Int, _requestId []*big.Int) (*TellorDisputeNewDisputeIterator, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _TellorDispute.contract.FilterLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &TellorDisputeNewDisputeIterator{contract: _TellorDispute.contract, event: "NewDispute", logs: logs, sub: sub}, nil
}

// WatchNewDispute is a free log subscription operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: e NewDispute(_disputeId indexed uint256, _requestId indexed uint256, _timestamp uint256, _miner address)
func (_TellorDispute *TellorDisputeFilterer) WatchNewDispute(opts *bind.WatchOpts, sink chan<- *TellorDisputeNewDispute, _disputeId []*big.Int, _requestId []*big.Int) (event.Subscription, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _TellorDispute.contract.WatchLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorDisputeNewDispute)
				if err := _TellorDispute.contract.UnpackLog(event, "NewDispute", log); err != nil {
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

// TellorDisputeNewTellorAddressIterator is returned from FilterNewTellorAddress and is used to iterate over the raw logs and unpacked data for NewTellorAddress events raised by the TellorDispute contract.
type TellorDisputeNewTellorAddressIterator struct {
	Event *TellorDisputeNewTellorAddress // Event containing the contract specifics and raw log

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
func (it *TellorDisputeNewTellorAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorDisputeNewTellorAddress)
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
		it.Event = new(TellorDisputeNewTellorAddress)
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
func (it *TellorDisputeNewTellorAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorDisputeNewTellorAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorDisputeNewTellorAddress represents a NewTellorAddress event raised by the TellorDispute contract.
type TellorDisputeNewTellorAddress struct {
	NewTellor common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTellorAddress is a free log retrieval operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: e NewTellorAddress(_newTellor address)
func (_TellorDispute *TellorDisputeFilterer) FilterNewTellorAddress(opts *bind.FilterOpts) (*TellorDisputeNewTellorAddressIterator, error) {

	logs, sub, err := _TellorDispute.contract.FilterLogs(opts, "NewTellorAddress")
	if err != nil {
		return nil, err
	}
	return &TellorDisputeNewTellorAddressIterator{contract: _TellorDispute.contract, event: "NewTellorAddress", logs: logs, sub: sub}, nil
}

// WatchNewTellorAddress is a free log subscription operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: e NewTellorAddress(_newTellor address)
func (_TellorDispute *TellorDisputeFilterer) WatchNewTellorAddress(opts *bind.WatchOpts, sink chan<- *TellorDisputeNewTellorAddress) (event.Subscription, error) {

	logs, sub, err := _TellorDispute.contract.WatchLogs(opts, "NewTellorAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorDisputeNewTellorAddress)
				if err := _TellorDispute.contract.UnpackLog(event, "NewTellorAddress", log); err != nil {
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

// TellorDisputeVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the TellorDispute contract.
type TellorDisputeVotedIterator struct {
	Event *TellorDisputeVoted // Event containing the contract specifics and raw log

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
func (it *TellorDisputeVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorDisputeVoted)
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
		it.Event = new(TellorDisputeVoted)
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
func (it *TellorDisputeVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorDisputeVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorDisputeVoted represents a Voted event raised by the TellorDispute contract.
type TellorDisputeVoted struct {
	DisputeID *big.Int
	Position  bool
	Voter     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: e Voted(_disputeID indexed uint256, _position bool, _voter indexed address)
func (_TellorDispute *TellorDisputeFilterer) FilterVoted(opts *bind.FilterOpts, _disputeID []*big.Int, _voter []common.Address) (*TellorDisputeVotedIterator, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}

	logs, sub, err := _TellorDispute.contract.FilterLogs(opts, "Voted", _disputeIDRule, _voterRule)
	if err != nil {
		return nil, err
	}
	return &TellorDisputeVotedIterator{contract: _TellorDispute.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: e Voted(_disputeID indexed uint256, _position bool, _voter indexed address)
func (_TellorDispute *TellorDisputeFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *TellorDisputeVoted, _disputeID []*big.Int, _voter []common.Address) (event.Subscription, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}

	logs, sub, err := _TellorDispute.contract.WatchLogs(opts, "Voted", _disputeIDRule, _voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorDisputeVoted)
				if err := _TellorDispute.contract.UnpackLog(event, "Voted", log); err != nil {
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

// TellorGettersABI is the input ABI used to generate the binding from.
const TellorGettersABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSymbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVariablesOnDeck\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_request\",\"type\":\"bytes32\"}],\"name\":\"getRequestIdByQueryHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"address[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestID\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentVariables\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256[9]\"},{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[51]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastNewValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TellorGettersBin is the compiled bytecode used for deploying new contracts.
const TellorGettersBin = `0x608060405234801561001057600080fd5b50611a2a806100206000396000f3fe608060405234801561001057600080fd5b50600436106101f05760003560e01c806370a082311161010f578063af0b1327116100a2578063dd62ed3e11610071578063dd62ed3e146107fa578063e0ae93c114610828578063e1eee6d61461084b578063fc7cf0a014610962576101f0565b8063af0b1327146106f8578063b54130291461079c578063c775b542146107ba578063da379941146107dd576101f0565b806393fa4915116100de57806393fa4915146105da578063999cf26c146105fd578063a22e407a14610629578063a7c438bc146106cc576101f0565b806370a082311461052f578063733bdef01461055557806377fbb663146105945780637f6fd5d9146105b7576101f0565b80633180f8df11610187578063612c8f7f11610156578063612c8f7f146104a65780636173c0b8146104c357806363bb82ad146104e057806369026d631461050c576101f0565b80633180f8df146103f05780633df0777b1461042657806346eee1c41461045d5780634ee2cd7e1461047a576101f0565b806317d7de7c116101c357806317d7de7c1461033557806318160ddd1461033d57806319e8e03b146103455780631db842f0146103d3576101f0565b80630f0b424d146101f557806311c9851214610224578063133bee5e1461027f57806315070401146102b8575b600080fd5b6102126004803603602081101561020b57600080fd5b503561096a565b60408051918252519081900360200190f35b6102476004803603604081101561023a57600080fd5b5080359060200135610982565b604051808260a080838360005b8381101561026c578181015183820152602001610254565b5050505090500191505060405180910390f35b61029c6004803603602081101561029557600080fd5b50356109a3565b604080516001600160a01b039092168252519081900360200190f35b6102c06109b5565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102fa5781810151838201526020016102e2565b50505050905090810190601f1680156103275780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102c06109c6565b6102126109d2565b61034d6109de565b6040518084815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561039657818101518382015260200161037e565b50505050905090810190601f1680156103c35780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b610212600480360360208110156103e957600080fd5b50356109f8565b61040d6004803603602081101561040657600080fd5b5035610a0a565b6040805192835290151560208301528051918290030190f35b6104496004803603604081101561043c57600080fd5b5080359060200135610a26565b604080519115158252519081900360200190f35b6102126004803603602081101561047357600080fd5b5035610a39565b6102126004803603604081101561049057600080fd5b506001600160a01b038135169060200135610a4b565b610212600480360360208110156104bc57600080fd5b5035610ae8565b610212600480360360208110156104d957600080fd5b5035610afa565b610449600480360360408110156104f657600080fd5b50803590602001356001600160a01b0316610b0c565b6102476004803603604081101561052257600080fd5b5080359060200135610b1f565b6102126004803603602081101561054557600080fd5b50356001600160a01b0316610b39565b61057b6004803603602081101561056b57600080fd5b50356001600160a01b0316610bce565b6040805192835260208301919091528051918290030190f35b610212600480360360408110156105aa57600080fd5b5080359060200135610be1565b610212600480360360408110156105cd57600080fd5b5080359060200135610bf4565b610212600480360360408110156105f057600080fd5b5080359060200135610c07565b6104496004803603604081101561061357600080fd5b506001600160a01b038135169060200135610c1a565b610631610c84565b6040518087815260200186815260200185815260200180602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b8381101561068c578181015183820152602001610674565b50505050905090810190601f1680156106b95780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b610449600480360360408110156106e257600080fd5b50803590602001356001600160a01b0316610cab565b6107156004803603602081101561070e57600080fd5b5035610cbe565b604080518a815289151560208201528815159181019190915286151560608201526001600160a01b03808716608083015285811660a0830152841660c082015260e081018361012080838360005b8381101561077b578181015183820152602001610763565b50505050905001828152602001995050505050505050505060405180910390f35b6107a4610d02565b6040518151815280826106608083836020610254565b610212600480360360408110156107d057600080fd5b5080359060200135610d14565b610212600480360360208110156107f357600080fd5b5035610d27565b6102126004803603604081101561081057600080fd5b506001600160a01b0381358116916020013516610d39565b6102126004803603604081101561083e57600080fd5b5080359060200135610da4565b6108686004803603602081101561086157600080fd5b5035610db7565b604051808060200180602001878152602001868152602001858152602001848152602001838103835289818151815260200191508051906020019080838360005b838110156108c15781810151838201526020016108a9565b50505050905090810190601f1680156108ee5780820380516001836020036101000a031916815260200191505b5083810382528851815288516020918201918a019080838360005b83811015610921578181015183820152602001610909565b50505050905090810190601f16801561094e5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b61040d610de3565b600061097c818363ffffffff610df816565b92915050565b61098a611999565b61099c6000848463ffffffff610e0e16565b9392505050565b600061097c818363ffffffff610e6816565b60606109c16000610e87565b905090565b60606109c16000610ea4565b60006109c16000610ece565b60008060606109ed6000610f01565b925092509250909192565b600061097c818363ffffffff610fed16565b600080610a1d818463ffffffff61100316565b91509150915091565b600061099c81848463ffffffff61106916565b600061097c818363ffffffff61109016565b60408051633f48b1ff60e01b81526000600482018190526001600160a01b038516602483015260448201849052915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__91633f48b1ff916064808301926020929190829003018186803b158015610ab557600080fd5b505af4158015610ac9573d6000803e3d6000fd5b505050506040513d6020811015610adf57600080fd5b50519392505050565b600061097c818363ffffffff6110a916565b600061097c818363ffffffff6110bb16565b600061099c81848463ffffffff6110e216565b610b27611999565b61099c6000848463ffffffff61110f16565b604080516393b182b360e01b81526000600482018190526001600160a01b0384166024830152915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__916393b182b3916044808301926020929190829003018186803b158015610b9c57600080fd5b505af4158015610bb0573d6000803e3d6000fd5b505050506040513d6020811015610bc657600080fd5b505192915050565b600080610a1d818463ffffffff61117316565b600061099c81848463ffffffff61119a16565b600061099c81848463ffffffff6111cd16565b600061099c81848463ffffffff6111f116565b604080516356555cf160e11b81526000600482018190526001600160a01b038516602483015260448201849052915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163acaab9e2916064808301926020929190829003018186803b158015610ab557600080fd5b60008060006060600080610c986000611215565b949b939a50919850965094509092509050565b600061099c81848463ffffffff6113d316565b6000806000806000806000610cd16119b7565b6000610ce3818b63ffffffff61140516565b9850985098509850985098509850985098509193959799909294969850565b610d0a6119d6565b6109c1600061161a565b600061099c81848463ffffffff61165a16565b600061097c818363ffffffff61167e16565b60408051632fcc801b60e21b81526000600482018190526001600160a01b03808616602484015284166044830152915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163bf32006c916064808301926020929190829003018186803b158015610ab557600080fd5b600061099c81848463ffffffff61169416565b6060806000808080610dcf818863ffffffff6116b816565b949c939b5091995097509550909350915050565b600080610df06000611893565b915091509091565b6000908152604291909101602052604090205490565b610e16611999565b6000838152604885016020908152604080832085845260090190915290819020815160a08101928390529160059082845b815481526020019060010190808311610e4757505050505090509392505050565b6000908152603f9190910160205260409020546001600160a01b031690565b50604080518082019091526002815261151560f21b602082015290565b5060408051808201909152600f81526e54656c6c6f7220547269627574657360881b602082015290565b604080516b746f74616c5f737570706c7960a01b8152815190819003600c01902060009081528183016020522054919050565b60008060606000610f1185611909565b600081815260488701602081815260408084208151670746f74616c5469760c41b8152825160089181900391909101812086526004820184528286205495879052938352805460026001821615610100026000190190911604601f810184900484028501840190925281845294955085949291839190830182828015610fd85780601f10610fad57610100808354040283529160200191610fd8565b820191906000526020600020905b815481529060010190602001808311610fbb57829003601f168201915b50505050509050935093509350509193909250565b6000908152604991909101602052604090205490565b600081815260488301602052604081206003810154829190156110595760038101805461104d918791879190600019810190811061103d57fe5b90600052602060002001546111f1565b60019250925050611062565b50600091508190505b9250929050565b60009182526048929092016020908152604080832093835260079093019052205460ff1690565b6000908152604891909101602052604090206003015490565b60009081526040918201602052205490565b600060328211156110cb57600080fd5b506000908152604391909101602052604090205490565b6000918252604192909201602090815260408083206001600160a01b039094168352929052205460ff1690565b611117611999565b6000838152604885016020908152604080832085845260080190915290819020815160a08101928390529160059082845b81546001600160a01b0316815260019091019060200180831161114857505050505090509392505050565b6001600160a01b031660009081526047919091016020526040902080546001909101549091565b600082815260488401602052604081206003018054839081106111b957fe5b906000526020600020015490509392505050565b60009182526044929092016020908152604080832093835260059093019052205490565b60009182526048929092016020908152604080832093835260069093019052205490565b8054604080516f18dd5c9c995b9d14995c5d595cdd125960821b80825282519182900360109081018320600090815284870160208181528683205469646966666963756c747960b01b8752875196879003600a01872084528282528784205486885288519788900386018820855283835288852054855260488b01808452898620888a528a51998a900388018a2087528585528a87205487528185528a87206a6772616e756c617269747960a81b8b528b519a8b9003600b018b208852600490810186528b882054998b528b519a8b90039098018a2087529484528986205486528352888520670746f74616c5469760c41b8952895198899003600801892086529095018252878420548354601f600260001961010060018516150201909216919091049081018490048402890184019099528888529398899889986060988a98899894979596909594919391929185918301828280156113b75780601f1061138c576101008083540402835291602001916113b7565b820191906000526020600020905b81548152906001019060200180831161139a57829003601f168201915b5050505050925095509550955095509550955091939550919395565b600082815260448401602090815260408083206001600160a01b038516845260060190915290205460ff169392505050565b60008060008060008060006114186119b7565b5050506000868152604488016020908152604080832080546002820154600380840154600485015486516101208101808952681c995c5d595cdd125960ba1b905287518082036101290190208a526005808801808b52898c20548352895168074696d657374616d760bc1b81528a519081900360099081019091208d52818c528a8d2054848d01528a516476616c756560d81b81528b51908190039093019092208c52808b52898c2054838b015289516f6d696e457865637574696f6e4461746560801b81528a519081900360100190208c52808b52898c2054606084015289516c6e756d6265724f66566f74657360981b81528a5190819003600d0190208c52808b52898c2054608084015289516a313637b1b5a73ab6b132b960a91b81528a5190819003600b0190208c52808b52898c205460a08401528951681b5a5b995c94db1bdd60ba1b81528a51908190039092019091208b52808a52888b205460c083015288516571756f72756d60d01b815289519081900360060190208b52808a52888b205460e083015288516266656560e81b81528951908190039095019094208a5292909752949096205461010087810191909152600190930154919a5060ff8082169a509281048316985062010000810490921696506001600160a01b0363010000009092048216955091811693921691909295985092959850929598565b6116226119d6565b6040805161066081019182905290600184019060339082845b81548152602001906001019080831161163b5750505050509050919050565b60009182526048929092016020908152604080832093835260059093019052205490565b6000908152604a91909101602052604090205490565b60009182526048929092016020908152604080832093835260049093019052205490565b6000818152604883016020908152604080832060028082015483516a6772616e756c617269747960a81b8152845190819003600b018120875260048401808752858820546f3932b8bab2b9ba28a837b9b4ba34b7b760811b83528651928390036010018320895281885286892054670746f74616c5469760c41b845287519384900360080184208a52918852868920548654601f6000196101006001848116159190910291909101909216979097049687018a90048a0285018a019098528584526060998a999098899889988998919788979388019692959294909188918301828280156117e75780601f106117bc576101008083540402835291602001916117e7565b820191906000526020600020905b8154815290600101906020018083116117ca57829003601f168201915b5050885460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959b508a9450925084019050828280156118755780601f1061184a57610100808354040283529160200191611875565b820191906000526020600020905b81548152906001019060200180831161185857829003601f168201915b50505050509450965096509650965096509650509295509295509295565b604080517174696d654f664c6173744e657756616c756560701b80825282516012928190038301812060009081528585016020818152868320548352604288018152868320549484528651938490039095019092208152925291812054909182916118ff9185916111f1565b9360019350915050565b604080516106608101918290526000918291829161194a9190600187019060339082845b81548152602001906001019080831161192d575050505050611965565b60009081526043909501602052505060409092205492915050565b60008060005b603381101561199357602081028401518084101561198a578093508192505b5060010161196b565b50915091565b6040518060a001604052806005906020820280388339509192915050565b6040518061012001604052806009906020820280388339509192915050565b604051806106600160405280603390602082028038833950919291505056fea265627a7a72305820cc7cdcaa2fea4ffcef4610ccbe2f219baa0c95dc886d7d996fc8219f373ffd7364736f6c634300050a0032`

// DeployTellorGetters deploys a new Ethereum contract, binding an instance of TellorGetters to it.
func DeployTellorGetters(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorGetters, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorGettersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TellorGettersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TellorGetters{TellorGettersCaller: TellorGettersCaller{contract: contract}, TellorGettersTransactor: TellorGettersTransactor{contract: contract}, TellorGettersFilterer: TellorGettersFilterer{contract: contract}}, nil
}

// TellorGetters is an auto generated Go binding around an Ethereum contract.
type TellorGetters struct {
	TellorGettersCaller     // Read-only binding to the contract
	TellorGettersTransactor // Write-only binding to the contract
	TellorGettersFilterer   // Log filterer for contract events
}

// TellorGettersCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorGettersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorGettersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorGettersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorGettersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorGettersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorGettersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorGettersSession struct {
	Contract     *TellorGetters    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorGettersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorGettersCallerSession struct {
	Contract *TellorGettersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TellorGettersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorGettersTransactorSession struct {
	Contract     *TellorGettersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TellorGettersRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorGettersRaw struct {
	Contract *TellorGetters // Generic contract binding to access the raw methods on
}

// TellorGettersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorGettersCallerRaw struct {
	Contract *TellorGettersCaller // Generic read-only contract binding to access the raw methods on
}

// TellorGettersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorGettersTransactorRaw struct {
	Contract *TellorGettersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellorGetters creates a new instance of TellorGetters, bound to a specific deployed contract.
func NewTellorGetters(address common.Address, backend bind.ContractBackend) (*TellorGetters, error) {
	contract, err := bindTellorGetters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TellorGetters{TellorGettersCaller: TellorGettersCaller{contract: contract}, TellorGettersTransactor: TellorGettersTransactor{contract: contract}, TellorGettersFilterer: TellorGettersFilterer{contract: contract}}, nil
}

// NewTellorGettersCaller creates a new read-only instance of TellorGetters, bound to a specific deployed contract.
func NewTellorGettersCaller(address common.Address, caller bind.ContractCaller) (*TellorGettersCaller, error) {
	contract, err := bindTellorGetters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorGettersCaller{contract: contract}, nil
}

// NewTellorGettersTransactor creates a new write-only instance of TellorGetters, bound to a specific deployed contract.
func NewTellorGettersTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorGettersTransactor, error) {
	contract, err := bindTellorGetters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorGettersTransactor{contract: contract}, nil
}

// NewTellorGettersFilterer creates a new log filterer instance of TellorGetters, bound to a specific deployed contract.
func NewTellorGettersFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorGettersFilterer, error) {
	contract, err := bindTellorGetters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorGettersFilterer{contract: contract}, nil
}

// bindTellorGetters binds a generic wrapper to an already deployed contract.
func bindTellorGetters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorGettersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorGetters *TellorGettersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TellorGetters.Contract.TellorGettersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorGetters *TellorGettersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorGetters.Contract.TellorGettersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorGetters *TellorGettersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorGetters.Contract.TellorGettersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorGetters *TellorGettersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TellorGetters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorGetters *TellorGettersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorGetters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorGetters *TellorGettersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorGetters.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_user address, _spender address) constant returns(uint256)
func (_TellorGetters *TellorGettersCaller) Allowance(opts *bind.CallOpts, _user common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "allowance", _user, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_user address, _spender address) constant returns(uint256)
func (_TellorGetters *TellorGettersSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _TellorGetters.Contract.Allowance(&_TellorGetters.CallOpts, _user, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_user address, _spender address) constant returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _TellorGetters.Contract.Allowance(&_TellorGetters.CallOpts, _user, _spender)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(_user address, _amount uint256) constant returns(bool)
func (_TellorGetters *TellorGettersCaller) AllowedToTrade(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "allowedToTrade", _user, _amount)
	return *ret0, err
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(_user address, _amount uint256) constant returns(bool)
func (_TellorGetters *TellorGettersSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _TellorGetters.Contract.AllowedToTrade(&_TellorGetters.CallOpts, _user, _amount)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(_user address, _amount uint256) constant returns(bool)
func (_TellorGetters *TellorGettersCallerSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _TellorGetters.Contract.AllowedToTrade(&_TellorGetters.CallOpts, _user, _amount)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_user address) constant returns(uint256)
func (_TellorGetters *TellorGettersCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "balanceOf", _user)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_user address) constant returns(uint256)
func (_TellorGetters *TellorGettersSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _TellorGetters.Contract.BalanceOf(&_TellorGetters.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_user address) constant returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _TellorGetters.Contract.BalanceOf(&_TellorGetters.CallOpts, _user)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(_user address, _blockNumber uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersCaller) BalanceOfAt(opts *bind.CallOpts, _user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "balanceOfAt", _user, _blockNumber)
	return *ret0, err
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(_user address, _blockNumber uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.BalanceOfAt(&_TellorGetters.CallOpts, _user, _blockNumber)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(_user address, _blockNumber uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.BalanceOfAt(&_TellorGetters.CallOpts, _user, _blockNumber)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(_challenge bytes32, _miner address) constant returns(bool)
func (_TellorGetters *TellorGettersCaller) DidMine(opts *bind.CallOpts, _challenge [32]byte, _miner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "didMine", _challenge, _miner)
	return *ret0, err
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(_challenge bytes32, _miner address) constant returns(bool)
func (_TellorGetters *TellorGettersSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _TellorGetters.Contract.DidMine(&_TellorGetters.CallOpts, _challenge, _miner)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(_challenge bytes32, _miner address) constant returns(bool)
func (_TellorGetters *TellorGettersCallerSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _TellorGetters.Contract.DidMine(&_TellorGetters.CallOpts, _challenge, _miner)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(_disputeId uint256, _address address) constant returns(bool)
func (_TellorGetters *TellorGettersCaller) DidVote(opts *bind.CallOpts, _disputeId *big.Int, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "didVote", _disputeId, _address)
	return *ret0, err
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(_disputeId uint256, _address address) constant returns(bool)
func (_TellorGetters *TellorGettersSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _TellorGetters.Contract.DidVote(&_TellorGetters.CallOpts, _disputeId, _address)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(_disputeId uint256, _address address) constant returns(bool)
func (_TellorGetters *TellorGettersCallerSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _TellorGetters.Contract.DidVote(&_TellorGetters.CallOpts, _disputeId, _address)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(_data bytes32) constant returns(address)
func (_TellorGetters *TellorGettersCaller) GetAddressVars(opts *bind.CallOpts, _data [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "getAddressVars", _data)
	return *ret0, err
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(_data bytes32) constant returns(address)
func (_TellorGetters *TellorGettersSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _TellorGetters.Contract.GetAddressVars(&_TellorGetters.CallOpts, _data)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(_data bytes32) constant returns(address)
func (_TellorGetters *TellorGettersCallerSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _TellorGetters.Contract.GetAddressVars(&_TellorGetters.CallOpts, _data)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(_disputeId uint256) constant returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_TellorGetters *TellorGettersCaller) GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
		ret2 = new(bool)
		ret3 = new(bool)
		ret4 = new(common.Address)
		ret5 = new(common.Address)
		ret6 = new(common.Address)
		ret7 = new([9]*big.Int)
		ret8 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
		ret7,
		ret8,
	}
	err := _TellorGetters.contract.Call(opts, out, "getAllDisputeVars", _disputeId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, *ret7, *ret8, err
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(_disputeId uint256) constant returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_TellorGetters *TellorGettersSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetAllDisputeVars(&_TellorGetters.CallOpts, _disputeId)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(_disputeId uint256) constant returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_TellorGetters *TellorGettersCallerSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetAllDisputeVars(&_TellorGetters.CallOpts, _disputeId)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() constant returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_TellorGetters *TellorGettersCaller) GetCurrentVariables(opts *bind.CallOpts) ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(string)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _TellorGetters.contract.Call(opts, out, "getCurrentVariables")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() constant returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_TellorGetters *TellorGettersSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetCurrentVariables(&_TellorGetters.CallOpts)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() constant returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_TellorGetters *TellorGettersCallerSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetCurrentVariables(&_TellorGetters.CallOpts)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(_hash bytes32) constant returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetDisputeIdByDisputeHash(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "getDisputeIdByDisputeHash", _hash)
	return *ret0, err
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(_hash bytes32) constant returns(uint256)
func (_TellorGetters *TellorGettersSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetDisputeIdByDisputeHash(&_TellorGetters.CallOpts, _hash)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(_hash bytes32) constant returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetDisputeIdByDisputeHash(&_TellorGetters.CallOpts, _hash)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(_disputeId uint256, _data bytes32) constant returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetDisputeUintVars(opts *bind.CallOpts, _disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "getDisputeUintVars", _disputeId, _data)
	return *ret0, err
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(_disputeId uint256, _data bytes32) constant returns(uint256)
func (_TellorGetters *TellorGettersSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetDisputeUintVars(&_TellorGetters.CallOpts, _disputeId, _data)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(_disputeId uint256, _data bytes32) constant returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetDisputeUintVars(&_TellorGetters.CallOpts, _disputeId, _data)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() constant returns(uint256, bool)
func (_TellorGetters *TellorGettersCaller) GetLastNewValue(opts *bind.CallOpts) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _TellorGetters.contract.Call(opts, out, "getLastNewValue")
	return *ret0, *ret1, err
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() constant returns(uint256, bool)
func (_TellorGetters *TellorGettersSession) GetLastNewValue() (*big.Int, bool, error) {
	return _TellorGetters.Contract.GetLastNewValue(&_TellorGetters.CallOpts)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() constant returns(uint256, bool)
func (_TellorGetters *TellorGettersCallerSession) GetLastNewValue() (*big.Int, bool, error) {
	return _TellorGetters.Contract.GetLastNewValue(&_TellorGetters.CallOpts)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(_requestId uint256) constant returns(uint256, bool)
func (_TellorGetters *TellorGettersCaller) GetLastNewValueById(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _TellorGetters.contract.Call(opts, out, "getLastNewValueById", _requestId)
	return *ret0, *ret1, err
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(_requestId uint256) constant returns(uint256, bool)
func (_TellorGetters *TellorGettersSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _TellorGetters.Contract.GetLastNewValueById(&_TellorGetters.CallOpts, _requestId)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(_requestId uint256) constant returns(uint256, bool)
func (_TellorGetters *TellorGettersCallerSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _TellorGetters.Contract.GetLastNewValueById(&_TellorGetters.CallOpts, _requestId)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetMinedBlockNum(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "getMinedBlockNum", _requestId, _timestamp)
	return *ret0, err
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetMinedBlockNum(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetMinedBlockNum(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(_requestId uint256, _timestamp uint256) constant returns(address[5])
func (_TellorGetters *TellorGettersCaller) GetMinersByRequestIdAndTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	var (
		ret0 = new([5]common.Address)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "getMinersByRequestIdAndTimestamp", _requestId, _timestamp)
	return *ret0, err
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(_requestId uint256, _timestamp uint256) constant returns(address[5])
func (_TellorGetters *TellorGettersSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _TellorGetters.Contract.GetMinersByRequestIdAndTimestamp(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(_requestId uint256, _timestamp uint256) constant returns(address[5])
func (_TellorGetters *TellorGettersCallerSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _TellorGetters.Contract.GetMinersByRequestIdAndTimestamp(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_TellorGetters *TellorGettersCaller) GetName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "getName")
	return *ret0, err
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_TellorGetters *TellorGettersSession) GetName() (string, error) {
	return _TellorGetters.Contract.GetName(&_TellorGetters.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_TellorGetters *TellorGettersCallerSession) GetName() (string, error) {
	return _TellorGetters.Contract.GetName(&_TellorGetters.CallOpts)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(_requestId uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "getNewValueCountbyRequestId", _requestId)
	return *ret0, err
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(_requestId uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetNewValueCountbyRequestId(&_TellorGetters.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(_requestId uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetNewValueCountbyRequestId(&_TellorGetters.CallOpts, _requestId)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(_request bytes32) constant returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetRequestIdByQueryHash(opts *bind.CallOpts, _request [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "getRequestIdByQueryHash", _request)
	return *ret0, err
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(_request bytes32) constant returns(uint256)
func (_TellorGetters *TellorGettersSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestIdByQueryHash(&_TellorGetters.CallOpts, _request)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(_request bytes32) constant returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestIdByQueryHash(&_TellorGetters.CallOpts, _request)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(_index uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetRequestIdByRequestQIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "getRequestIdByRequestQIndex", _index)
	return *ret0, err
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(_index uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestIdByRequestQIndex(&_TellorGetters.CallOpts, _index)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(_index uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestIdByRequestQIndex(&_TellorGetters.CallOpts, _index)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(_timestamp uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetRequestIdByTimestamp(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "getRequestIdByTimestamp", _timestamp)
	return *ret0, err
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(_timestamp uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestIdByTimestamp(&_TellorGetters.CallOpts, _timestamp)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(_timestamp uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestIdByTimestamp(&_TellorGetters.CallOpts, _timestamp)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() constant returns(uint256[51])
func (_TellorGetters *TellorGettersCaller) GetRequestQ(opts *bind.CallOpts) ([51]*big.Int, error) {
	var (
		ret0 = new([51]*big.Int)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "getRequestQ")
	return *ret0, err
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() constant returns(uint256[51])
func (_TellorGetters *TellorGettersSession) GetRequestQ() ([51]*big.Int, error) {
	return _TellorGetters.Contract.GetRequestQ(&_TellorGetters.CallOpts)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() constant returns(uint256[51])
func (_TellorGetters *TellorGettersCallerSession) GetRequestQ() ([51]*big.Int, error) {
	return _TellorGetters.Contract.GetRequestQ(&_TellorGetters.CallOpts)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(_requestId uint256, _data bytes32) constant returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetRequestUintVars(opts *bind.CallOpts, _requestId *big.Int, _data [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "getRequestUintVars", _requestId, _data)
	return *ret0, err
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(_requestId uint256, _data bytes32) constant returns(uint256)
func (_TellorGetters *TellorGettersSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestUintVars(&_TellorGetters.CallOpts, _requestId, _data)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(_requestId uint256, _data bytes32) constant returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestUintVars(&_TellorGetters.CallOpts, _requestId, _data)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(_requestId uint256) constant returns(string, string, bytes32, uint256, uint256, uint256)
func (_TellorGetters *TellorGettersCaller) GetRequestVars(opts *bind.CallOpts, _requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new([32]byte)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _TellorGetters.contract.Call(opts, out, "getRequestVars", _requestId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(_requestId uint256) constant returns(string, string, bytes32, uint256, uint256, uint256)
func (_TellorGetters *TellorGettersSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetRequestVars(&_TellorGetters.CallOpts, _requestId)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(_requestId uint256) constant returns(string, string, bytes32, uint256, uint256, uint256)
func (_TellorGetters *TellorGettersCallerSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetRequestVars(&_TellorGetters.CallOpts, _requestId)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(_staker address) constant returns(uint256, uint256)
func (_TellorGetters *TellorGettersCaller) GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _TellorGetters.contract.Call(opts, out, "getStakerInfo", _staker)
	return *ret0, *ret1, err
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(_staker address) constant returns(uint256, uint256)
func (_TellorGetters *TellorGettersSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetStakerInfo(&_TellorGetters.CallOpts, _staker)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(_staker address) constant returns(uint256, uint256)
func (_TellorGetters *TellorGettersCallerSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetStakerInfo(&_TellorGetters.CallOpts, _staker)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(_requestId uint256, _timestamp uint256) constant returns(uint256[5])
func (_TellorGetters *TellorGettersCaller) GetSubmissionsByTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	var (
		ret0 = new([5]*big.Int)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "getSubmissionsByTimestamp", _requestId, _timestamp)
	return *ret0, err
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(_requestId uint256, _timestamp uint256) constant returns(uint256[5])
func (_TellorGetters *TellorGettersSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _TellorGetters.Contract.GetSubmissionsByTimestamp(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(_requestId uint256, _timestamp uint256) constant returns(uint256[5])
func (_TellorGetters *TellorGettersCallerSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _TellorGetters.Contract.GetSubmissionsByTimestamp(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() constant returns(string)
func (_TellorGetters *TellorGettersCaller) GetSymbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "getSymbol")
	return *ret0, err
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() constant returns(string)
func (_TellorGetters *TellorGettersSession) GetSymbol() (string, error) {
	return _TellorGetters.Contract.GetSymbol(&_TellorGetters.CallOpts)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() constant returns(string)
func (_TellorGetters *TellorGettersCallerSession) GetSymbol() (string, error) {
	return _TellorGetters.Contract.GetSymbol(&_TellorGetters.CallOpts)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(_requestID uint256, _index uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "getTimestampbyRequestIDandIndex", _requestID, _index)
	return *ret0, err
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(_requestID uint256, _index uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetTimestampbyRequestIDandIndex(&_TellorGetters.CallOpts, _requestID, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(_requestID uint256, _index uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetTimestampbyRequestIDandIndex(&_TellorGetters.CallOpts, _requestID, _index)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(_data bytes32) constant returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "getUintVar", _data)
	return *ret0, err
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(_data bytes32) constant returns(uint256)
func (_TellorGetters *TellorGettersSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetUintVar(&_TellorGetters.CallOpts, _data)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(_data bytes32) constant returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetUintVar(&_TellorGetters.CallOpts, _data)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() constant returns(uint256, uint256, string)
func (_TellorGetters *TellorGettersCaller) GetVariablesOnDeck(opts *bind.CallOpts) (*big.Int, *big.Int, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _TellorGetters.contract.Call(opts, out, "getVariablesOnDeck")
	return *ret0, *ret1, *ret2, err
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() constant returns(uint256, uint256, string)
func (_TellorGetters *TellorGettersSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _TellorGetters.Contract.GetVariablesOnDeck(&_TellorGetters.CallOpts)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() constant returns(uint256, uint256, string)
func (_TellorGetters *TellorGettersCallerSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _TellorGetters.Contract.GetVariablesOnDeck(&_TellorGetters.CallOpts)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(_requestId uint256, _timestamp uint256) constant returns(bool)
func (_TellorGetters *TellorGettersCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "isInDispute", _requestId, _timestamp)
	return *ret0, err
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(_requestId uint256, _timestamp uint256) constant returns(bool)
func (_TellorGetters *TellorGettersSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _TellorGetters.Contract.IsInDispute(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(_requestId uint256, _timestamp uint256) constant returns(bool)
func (_TellorGetters *TellorGettersCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _TellorGetters.Contract.IsInDispute(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "retrieveData", _requestId, _timestamp)
	return *ret0, err
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.RetrieveData(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.RetrieveData(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TellorGetters *TellorGettersCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorGetters.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TellorGetters *TellorGettersSession) TotalSupply() (*big.Int, error) {
	return _TellorGetters.Contract.TotalSupply(&_TellorGetters.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) TotalSupply() (*big.Int, error) {
	return _TellorGetters.Contract.TotalSupply(&_TellorGetters.CallOpts)
}

// TellorGettersLibraryABI is the input ABI used to generate the binding from.
const TellorGettersLibraryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_newTellor\",\"type\":\"address\"}],\"name\":\"NewTellorAddress\",\"type\":\"event\"}]"

// TellorGettersLibraryBin is the compiled bytecode used for deploying new contracts.
const TellorGettersLibraryBin = `0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72305820da6ebe9dfbd29ba4e44451f27443534bf73520041279e6b4bb3ce66caf83e1e364736f6c634300050a0032`

// DeployTellorGettersLibrary deploys a new Ethereum contract, binding an instance of TellorGettersLibrary to it.
func DeployTellorGettersLibrary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorGettersLibrary, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorGettersLibraryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TellorGettersLibraryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TellorGettersLibrary{TellorGettersLibraryCaller: TellorGettersLibraryCaller{contract: contract}, TellorGettersLibraryTransactor: TellorGettersLibraryTransactor{contract: contract}, TellorGettersLibraryFilterer: TellorGettersLibraryFilterer{contract: contract}}, nil
}

// TellorGettersLibrary is an auto generated Go binding around an Ethereum contract.
type TellorGettersLibrary struct {
	TellorGettersLibraryCaller     // Read-only binding to the contract
	TellorGettersLibraryTransactor // Write-only binding to the contract
	TellorGettersLibraryFilterer   // Log filterer for contract events
}

// TellorGettersLibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorGettersLibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorGettersLibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorGettersLibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorGettersLibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorGettersLibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorGettersLibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorGettersLibrarySession struct {
	Contract     *TellorGettersLibrary // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TellorGettersLibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorGettersLibraryCallerSession struct {
	Contract *TellorGettersLibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// TellorGettersLibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorGettersLibraryTransactorSession struct {
	Contract     *TellorGettersLibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// TellorGettersLibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorGettersLibraryRaw struct {
	Contract *TellorGettersLibrary // Generic contract binding to access the raw methods on
}

// TellorGettersLibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorGettersLibraryCallerRaw struct {
	Contract *TellorGettersLibraryCaller // Generic read-only contract binding to access the raw methods on
}

// TellorGettersLibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorGettersLibraryTransactorRaw struct {
	Contract *TellorGettersLibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellorGettersLibrary creates a new instance of TellorGettersLibrary, bound to a specific deployed contract.
func NewTellorGettersLibrary(address common.Address, backend bind.ContractBackend) (*TellorGettersLibrary, error) {
	contract, err := bindTellorGettersLibrary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TellorGettersLibrary{TellorGettersLibraryCaller: TellorGettersLibraryCaller{contract: contract}, TellorGettersLibraryTransactor: TellorGettersLibraryTransactor{contract: contract}, TellorGettersLibraryFilterer: TellorGettersLibraryFilterer{contract: contract}}, nil
}

// NewTellorGettersLibraryCaller creates a new read-only instance of TellorGettersLibrary, bound to a specific deployed contract.
func NewTellorGettersLibraryCaller(address common.Address, caller bind.ContractCaller) (*TellorGettersLibraryCaller, error) {
	contract, err := bindTellorGettersLibrary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorGettersLibraryCaller{contract: contract}, nil
}

// NewTellorGettersLibraryTransactor creates a new write-only instance of TellorGettersLibrary, bound to a specific deployed contract.
func NewTellorGettersLibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorGettersLibraryTransactor, error) {
	contract, err := bindTellorGettersLibrary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorGettersLibraryTransactor{contract: contract}, nil
}

// NewTellorGettersLibraryFilterer creates a new log filterer instance of TellorGettersLibrary, bound to a specific deployed contract.
func NewTellorGettersLibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorGettersLibraryFilterer, error) {
	contract, err := bindTellorGettersLibrary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorGettersLibraryFilterer{contract: contract}, nil
}

// bindTellorGettersLibrary binds a generic wrapper to an already deployed contract.
func bindTellorGettersLibrary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorGettersLibraryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorGettersLibrary *TellorGettersLibraryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TellorGettersLibrary.Contract.TellorGettersLibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorGettersLibrary *TellorGettersLibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorGettersLibrary.Contract.TellorGettersLibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorGettersLibrary *TellorGettersLibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorGettersLibrary.Contract.TellorGettersLibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorGettersLibrary *TellorGettersLibraryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TellorGettersLibrary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorGettersLibrary *TellorGettersLibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorGettersLibrary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorGettersLibrary *TellorGettersLibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorGettersLibrary.Contract.contract.Transact(opts, method, params...)
}

// TellorGettersLibraryNewTellorAddressIterator is returned from FilterNewTellorAddress and is used to iterate over the raw logs and unpacked data for NewTellorAddress events raised by the TellorGettersLibrary contract.
type TellorGettersLibraryNewTellorAddressIterator struct {
	Event *TellorGettersLibraryNewTellorAddress // Event containing the contract specifics and raw log

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
func (it *TellorGettersLibraryNewTellorAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorGettersLibraryNewTellorAddress)
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
		it.Event = new(TellorGettersLibraryNewTellorAddress)
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
func (it *TellorGettersLibraryNewTellorAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorGettersLibraryNewTellorAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorGettersLibraryNewTellorAddress represents a NewTellorAddress event raised by the TellorGettersLibrary contract.
type TellorGettersLibraryNewTellorAddress struct {
	NewTellor common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTellorAddress is a free log retrieval operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: e NewTellorAddress(_newTellor address)
func (_TellorGettersLibrary *TellorGettersLibraryFilterer) FilterNewTellorAddress(opts *bind.FilterOpts) (*TellorGettersLibraryNewTellorAddressIterator, error) {

	logs, sub, err := _TellorGettersLibrary.contract.FilterLogs(opts, "NewTellorAddress")
	if err != nil {
		return nil, err
	}
	return &TellorGettersLibraryNewTellorAddressIterator{contract: _TellorGettersLibrary.contract, event: "NewTellorAddress", logs: logs, sub: sub}, nil
}

// WatchNewTellorAddress is a free log subscription operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: e NewTellorAddress(_newTellor address)
func (_TellorGettersLibrary *TellorGettersLibraryFilterer) WatchNewTellorAddress(opts *bind.WatchOpts, sink chan<- *TellorGettersLibraryNewTellorAddress) (event.Subscription, error) {

	logs, sub, err := _TellorGettersLibrary.contract.WatchLogs(opts, "NewTellorAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorGettersLibraryNewTellorAddress)
				if err := _TellorGettersLibrary.contract.UnpackLog(event, "NewTellorAddress", log); err != nil {
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

// TellorMasterABI is the input ABI used to generate the binding from.
const TellorMasterABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSymbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVariablesOnDeck\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_request\",\"type\":\"bytes32\"}],\"name\":\"getRequestIdByQueryHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDeity\",\"type\":\"address\"}],\"name\":\"changeDeity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"address[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestID\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentVariables\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tellorContract\",\"type\":\"address\"}],\"name\":\"changeTellorContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256[9]\"},{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[51]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastNewValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_tellorContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_newTellor\",\"type\":\"address\"}],\"name\":\"NewTellorAddress\",\"type\":\"event\"}]"

// TellorMasterBin is the compiled bytecode used for deploying new contracts.
const TellorMasterBin = `0x608060405234801561001057600080fd5b506040516120123803806120128339818101604052602081101561003357600080fd5b5051604080517f4601f1cd000000000000000000000000000000000000000000000000000000008152600060048201819052915173__$799602413129f49037f52758954cf5aa52$__92634601f1cd9260248082019391829003018186803b15801561009e57600080fd5b505af41580156100b2573d6000803e3d6000fd5b5050604080517f5f6f776e657200000000000000000000000000000000000000000000000000008152815190819003600690810182206000908152603f602081815285832080546001600160a01b0319908116339081179092557f5f646569747900000000000000000000000000000000000000000000000000008752875196879003909501862084528282528684208054861690911790557f74656c6c6f72436f6e74726163740000000000000000000000000000000000008552855194859003600e01852083529081529084902080546001600160a01b03891693168317905590825291517fc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d9450908190039091019150a150611e3c806101d66000396000f3fe6080604052600436106101f95760003560e01c806370a082311161010d578063ae0a8279116100a0578063da3799411161006f578063da37994114610a3e578063dd62ed3e14610a68578063e0ae93c114610aa3578063e1eee6d614610ad3578063fc7cf0a014610bf7576101f9565b8063ae0a8279146108ff578063af0b132714610932578063b5413029146109e3578063c775b54214610a0e576101f9565b806393fa4915116100dc57806393fa4915146107ad578063999cf26c146107dd578063a22e407a14610816578063a7c438bc146108c6576101f9565b806370a08231146106ce578063733bdef01461070157806377fbb6631461074d5780637f6fd5d91461077d576101f9565b80633180f8df116101905780634ee2cd7e1161015f5780634ee2cd7e146105d8578063612c8f7f146106115780636173c0b81461063b57806363bb82ad1461066557806369026d631461069e576101f9565b80633180f8df146104f25780633df0777b1461053557806346eee1c41461057957806347abd7f1146105a3576101f9565b806317d7de7c116101cc57806317d7de7c1461040357806318160ddd1461041857806319e8e03b1461042d5780631db842f0146104c8576101f9565b80630f0b424d1461028f57806311c98512146102cb578063133bee5e146103335780631507040114610379575b604080516d1d195b1b1bdc90dbdb9d1c9858dd60921b8152815190819003600e0181206000908152603f602090815283822054601f369081018390048302850183019095528484526001600160a01b03169360609392918190840183828082843760009201829052508451949550938493509150506020840185600019f43d604051816000823e82801561028b578282f35b8282fd5b34801561029b57600080fd5b506102b9600480360360208110156102b257600080fd5b5035610c0c565b60408051918252519081900360200190f35b3480156102d757600080fd5b506102fb600480360360408110156102ee57600080fd5b5080359060200135610c24565b604051808260a080838360005b83811015610320578181015183820152602001610308565b5050505090500191505060405180910390f35b34801561033f57600080fd5b5061035d6004803603602081101561035657600080fd5b5035610c45565b604080516001600160a01b039092168252519081900360200190f35b34801561038557600080fd5b5061038e610c57565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103c85781810151838201526020016103b0565b50505050905090810190601f1680156103f55780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561040f57600080fd5b5061038e610c68565b34801561042457600080fd5b506102b9610c74565b34801561043957600080fd5b50610442610c80565b6040518084815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561048b578181015183820152602001610473565b50505050905090810190601f1680156104b85780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b3480156104d457600080fd5b506102b9600480360360208110156104eb57600080fd5b5035610c9a565b3480156104fe57600080fd5b5061051c6004803603602081101561051557600080fd5b5035610cac565b6040805192835290151560208301528051918290030190f35b34801561054157600080fd5b506105656004803603604081101561055857600080fd5b5080359060200135610cc8565b604080519115158252519081900360200190f35b34801561058557600080fd5b506102b96004803603602081101561059c57600080fd5b5035610cdb565b3480156105af57600080fd5b506105d6600480360360208110156105c657600080fd5b50356001600160a01b0316610ced565b005b3480156105e457600080fd5b506102b9600480360360408110156105fb57600080fd5b506001600160a01b038135169060200135610d01565b34801561061d57600080fd5b506102b96004803603602081101561063457600080fd5b5035610d9e565b34801561064757600080fd5b506102b96004803603602081101561065e57600080fd5b5035610db0565b34801561067157600080fd5b506105656004803603604081101561068857600080fd5b50803590602001356001600160a01b0316610dc2565b3480156106aa57600080fd5b506102fb600480360360408110156106c157600080fd5b5080359060200135610dd5565b3480156106da57600080fd5b506102b9600480360360208110156106f157600080fd5b50356001600160a01b0316610def565b34801561070d57600080fd5b506107346004803603602081101561072457600080fd5b50356001600160a01b0316610e84565b6040805192835260208301919091528051918290030190f35b34801561075957600080fd5b506102b96004803603604081101561077057600080fd5b5080359060200135610e97565b34801561078957600080fd5b506102b9600480360360408110156107a057600080fd5b5080359060200135610eaa565b3480156107b957600080fd5b506102b9600480360360408110156107d057600080fd5b5080359060200135610ebd565b3480156107e957600080fd5b506105656004803603604081101561080057600080fd5b506001600160a01b038135169060200135610ed0565b34801561082257600080fd5b5061082b610f3a565b6040518087815260200186815260200185815260200180602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b8381101561088657818101518382015260200161086e565b50505050905090810190601f1680156108b35780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b3480156108d257600080fd5b50610565600480360360408110156108e957600080fd5b50803590602001356001600160a01b0316610f61565b34801561090b57600080fd5b506105d66004803603602081101561092257600080fd5b50356001600160a01b0316610f74565b34801561093e57600080fd5b5061095c6004803603602081101561095557600080fd5b5035610f85565b604080518a815289151560208201528815159181019190915286151560608201526001600160a01b03808716608083015285811660a0830152841660c082015260e081018361012080838360005b838110156109c25781810151838201526020016109aa565b50505050905001828152602001995050505050505050505060405180910390f35b3480156109ef57600080fd5b506109f8610fc9565b6040518151815280826106608083836020610308565b348015610a1a57600080fd5b506102b960048036036040811015610a3157600080fd5b5080359060200135610fdb565b348015610a4a57600080fd5b506102b960048036036020811015610a6157600080fd5b5035610fee565b348015610a7457600080fd5b506102b960048036036040811015610a8b57600080fd5b506001600160a01b0381358116916020013516611000565b348015610aaf57600080fd5b506102b960048036036040811015610ac657600080fd5b508035906020013561106b565b348015610adf57600080fd5b50610afd60048036036020811015610af657600080fd5b503561107e565b604051808060200180602001878152602001868152602001858152602001848152602001838103835289818151815260200191508051906020019080838360005b83811015610b56578181015183820152602001610b3e565b50505050905090810190601f168015610b835780820380516001836020036101000a031916815260200191505b5083810382528851815288516020918201918a019080838360005b83811015610bb6578181015183820152602001610b9e565b50505050905090810190601f168015610be35780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b348015610c0357600080fd5b5061051c6110aa565b6000610c1e818363ffffffff6110bf16565b92915050565b610c2c611dab565b610c3e6000848463ffffffff6110d516565b9392505050565b6000610c1e818363ffffffff61112f16565b6060610c63600061114e565b905090565b6060610c63600061116b565b6000610c636000611195565b6000806060610c8f60006111c8565b925092509250909192565b6000610c1e818363ffffffff6112b416565b600080610cbf818463ffffffff6112ca16565b91509150915091565b6000610c3e81848463ffffffff61133016565b6000610c1e818363ffffffff61135716565b610cfe60008263ffffffff61137016565b50565b60408051633f48b1ff60e01b81526000600482018190526001600160a01b038516602483015260448201849052915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__91633f48b1ff916064808301926020929190829003018186803b158015610d6b57600080fd5b505af4158015610d7f573d6000803e3d6000fd5b505050506040513d6020811015610d9557600080fd5b50519392505050565b6000610c1e818363ffffffff6113f916565b6000610c1e818363ffffffff61140b16565b6000610c3e81848463ffffffff61143216565b610ddd611dab565b610c3e6000848463ffffffff61145f16565b604080516393b182b360e01b81526000600482018190526001600160a01b0384166024830152915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__916393b182b3916044808301926020929190829003018186803b158015610e5257600080fd5b505af4158015610e66573d6000803e3d6000fd5b505050506040513d6020811015610e7c57600080fd5b505192915050565b600080610cbf818463ffffffff6114c316565b6000610c3e81848463ffffffff6114ea16565b6000610c3e81848463ffffffff61151d16565b6000610c3e81848463ffffffff61154116565b604080516356555cf160e11b81526000600482018190526001600160a01b038516602483015260448201849052915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163acaab9e2916064808301926020929190829003018186803b158015610d6b57600080fd5b60008060006060600080610f4e6000611565565b949b939a50919850965094509092509050565b6000610c3e81848463ffffffff61172316565b610cfe60008263ffffffff61175516565b6000806000806000806000610f98611dc9565b6000610faa818b63ffffffff61181716565b9850985098509850985098509850985098509193959799909294969850565b610fd1611de8565b610c636000611a2c565b6000610c3e81848463ffffffff611a6c16565b6000610c1e818363ffffffff611a9016565b60408051632fcc801b60e21b81526000600482018190526001600160a01b03808616602484015284166044830152915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163bf32006c916064808301926020929190829003018186803b158015610d6b57600080fd5b6000610c3e81848463ffffffff611aa616565b6060806000808080611096818863ffffffff611aca16565b949c939b5091995097509550909350915050565b6000806110b76000611ca5565b915091509091565b6000908152604291909101602052604090205490565b6110dd611dab565b6000838152604885016020908152604080832085845260090190915290819020815160a08101928390529160059082845b81548152602001906001019080831161110e57505050505090509392505050565b6000908152603f9190910160205260409020546001600160a01b031690565b50604080518082019091526002815261151560f21b602082015290565b5060408051808201909152600f81526e54656c6c6f7220547269627574657360881b602082015290565b604080516b746f74616c5f737570706c7960a01b8152815190819003600c01902060009081528183016020522054919050565b600080606060006111d885611d1b565b600081815260488701602081815260408084208151670746f74616c5469760c41b8152825160089181900391909101812086526004820184528286205495879052938352805460026001821615610100026000190190911604601f81018490048402850184019092528184529495508594929183919083018282801561129f5780601f106112745761010080835404028352916020019161129f565b820191906000526020600020905b81548152906001019060200180831161128257829003601f168201915b50505050509050935093509350509193909250565b6000908152604991909101602052604090205490565b6000818152604883016020526040812060038101548291901561132057600381018054611314918791879190600019810190811061130457fe5b9060005260206000200154611541565b60019250925050611329565b50600091508190505b9250929050565b60009182526048929092016020908152604080832093835260079093019052205460ff1690565b6000908152604891909101602052604090206003015490565b60408051655f646569747960d01b815281519081900360060190206000908152603f840160205220546001600160a01b031633146113ad57600080fd5b60408051655f646569747960d01b815281519081900360060190206000908152603f90930160205290912080546001600160a01b039092166001600160a01b0319909216919091179055565b60009081526040918201602052205490565b6000603282111561141b57600080fd5b506000908152604391909101602052604090205490565b6000918252604192909201602090815260408083206001600160a01b039094168352929052205460ff1690565b611467611dab565b6000838152604885016020908152604080832085845260080190915290819020815160a08101928390529160059082845b81546001600160a01b0316815260019091019060200180831161149857505050505090509392505050565b6001600160a01b031660009081526047919091016020526040902080546001909101549091565b6000828152604884016020526040812060030180548390811061150957fe5b906000526020600020015490509392505050565b60009182526044929092016020908152604080832093835260059093019052205490565b60009182526048929092016020908152604080832093835260069093019052205490565b8054604080516f18dd5c9c995b9d14995c5d595cdd125960821b80825282519182900360109081018320600090815284870160208181528683205469646966666963756c747960b01b8752875196879003600a01872084528282528784205486885288519788900386018820855283835288852054855260488b01808452898620888a528a51998a900388018a2087528585528a87205487528185528a87206a6772616e756c617269747960a81b8b528b519a8b9003600b018b208852600490810186528b882054998b528b519a8b90039098018a2087529484528986205486528352888520670746f74616c5469760c41b8952895198899003600801892086529095018252878420548354601f600260001961010060018516150201909216919091049081018490048402890184019099528888529398899889986060988a98899894979596909594919391929185918301828280156117075780601f106116dc57610100808354040283529160200191611707565b820191906000526020600020905b8154815290600101906020018083116116ea57829003601f168201915b5050505050925095509550955095509550955091939550919395565b600082815260448401602090815260408083206001600160a01b038516845260060190915290205460ff169392505050565b60408051655f646569747960d01b815281519081900360060190206000908152603f840160205220546001600160a01b0316331461179257600080fd5b604080516d1d195b1b1bdc90dbdb9d1c9858dd60921b8152815190819003600e0181206000908152603f850160209081529083902080546001600160a01b0386166001600160a01b03199091168117909155825291517fc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d929181900390910190a15050565b600080600080600080600061182a611dc9565b5050506000868152604488016020908152604080832080546002820154600380840154600485015486516101208101808952681c995c5d595cdd125960ba1b905287518082036101290190208a526005808801808b52898c20548352895168074696d657374616d760bc1b81528a519081900360099081019091208d52818c528a8d2054848d01528a516476616c756560d81b81528b51908190039093019092208c52808b52898c2054838b015289516f6d696e457865637574696f6e4461746560801b81528a519081900360100190208c52808b52898c2054606084015289516c6e756d6265724f66566f74657360981b81528a5190819003600d0190208c52808b52898c2054608084015289516a313637b1b5a73ab6b132b960a91b81528a5190819003600b0190208c52808b52898c205460a08401528951681b5a5b995c94db1bdd60ba1b81528a51908190039092019091208b52808a52888b205460c083015288516571756f72756d60d01b815289519081900360060190208b52808a52888b205460e083015288516266656560e81b81528951908190039095019094208a5292909752949096205461010087810191909152600190930154919a5060ff8082169a509281048316985062010000810490921696506001600160a01b0363010000009092048216955091811693921691909295985092959850929598565b611a34611de8565b6040805161066081019182905290600184019060339082845b815481526020019060010190808311611a4d5750505050509050919050565b60009182526048929092016020908152604080832093835260059093019052205490565b6000908152604a91909101602052604090205490565b60009182526048929092016020908152604080832093835260049093019052205490565b6000818152604883016020908152604080832060028082015483516a6772616e756c617269747960a81b8152845190819003600b018120875260048401808752858820546f3932b8bab2b9ba28a837b9b4ba34b7b760811b83528651928390036010018320895281885286892054670746f74616c5469760c41b845287519384900360080184208a52918852868920548654601f6000196101006001848116159190910291909101909216979097049687018a90048a0285018a019098528584526060998a99909889988998899891978897938801969295929490918891830182828015611bf95780601f10611bce57610100808354040283529160200191611bf9565b820191906000526020600020905b815481529060010190602001808311611bdc57829003601f168201915b5050885460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959b508a945092508401905082828015611c875780601f10611c5c57610100808354040283529160200191611c87565b820191906000526020600020905b815481529060010190602001808311611c6a57829003601f168201915b50505050509450965096509650965096509650509295509295509295565b604080517174696d654f664c6173744e657756616c756560701b8082528251601292819003830181206000908152858501602081815286832054835260428801815286832054948452865193849003909501909220815292529181205490918291611d11918591611541565b9360019350915050565b6040805161066081019182905260009182918291611d5c9190600187019060339082845b815481526020019060010190808311611d3f575050505050611d77565b60009081526043909501602052505060409092205492915050565b60008060005b6033811015611da5576020810284015180841015611d9c578093508192505b50600101611d7d565b50915091565b6040518060a001604052806005906020820280388339509192915050565b6040518061012001604052806009906020820280388339509192915050565b604051806106600160405280603390602082028038833950919291505056fea265627a7a723058208f825b37f40114c9f8689b2b8792c2aa6e313b233bb3f67a270a1197636999e964736f6c634300050a0032`

// DeployTellorMaster deploys a new Ethereum contract, binding an instance of TellorMaster to it.
func DeployTellorMaster(auth *bind.TransactOpts, backend bind.ContractBackend, _tellorContract common.Address) (common.Address, *types.Transaction, *TellorMaster, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorMasterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TellorMasterBin), backend, _tellorContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TellorMaster{TellorMasterCaller: TellorMasterCaller{contract: contract}, TellorMasterTransactor: TellorMasterTransactor{contract: contract}, TellorMasterFilterer: TellorMasterFilterer{contract: contract}}, nil
}

// TellorMaster is an auto generated Go binding around an Ethereum contract.
type TellorMaster struct {
	TellorMasterCaller     // Read-only binding to the contract
	TellorMasterTransactor // Write-only binding to the contract
	TellorMasterFilterer   // Log filterer for contract events
}

// TellorMasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorMasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorMasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorMasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorMasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorMasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorMasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorMasterSession struct {
	Contract     *TellorMaster     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorMasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorMasterCallerSession struct {
	Contract *TellorMasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TellorMasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorMasterTransactorSession struct {
	Contract     *TellorMasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TellorMasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorMasterRaw struct {
	Contract *TellorMaster // Generic contract binding to access the raw methods on
}

// TellorMasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorMasterCallerRaw struct {
	Contract *TellorMasterCaller // Generic read-only contract binding to access the raw methods on
}

// TellorMasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorMasterTransactorRaw struct {
	Contract *TellorMasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellorMaster creates a new instance of TellorMaster, bound to a specific deployed contract.
func NewTellorMaster(address common.Address, backend bind.ContractBackend) (*TellorMaster, error) {
	contract, err := bindTellorMaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TellorMaster{TellorMasterCaller: TellorMasterCaller{contract: contract}, TellorMasterTransactor: TellorMasterTransactor{contract: contract}, TellorMasterFilterer: TellorMasterFilterer{contract: contract}}, nil
}

// NewTellorMasterCaller creates a new read-only instance of TellorMaster, bound to a specific deployed contract.
func NewTellorMasterCaller(address common.Address, caller bind.ContractCaller) (*TellorMasterCaller, error) {
	contract, err := bindTellorMaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorMasterCaller{contract: contract}, nil
}

// NewTellorMasterTransactor creates a new write-only instance of TellorMaster, bound to a specific deployed contract.
func NewTellorMasterTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorMasterTransactor, error) {
	contract, err := bindTellorMaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorMasterTransactor{contract: contract}, nil
}

// NewTellorMasterFilterer creates a new log filterer instance of TellorMaster, bound to a specific deployed contract.
func NewTellorMasterFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorMasterFilterer, error) {
	contract, err := bindTellorMaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorMasterFilterer{contract: contract}, nil
}

// bindTellorMaster binds a generic wrapper to an already deployed contract.
func bindTellorMaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorMasterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorMaster *TellorMasterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TellorMaster.Contract.TellorMasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorMaster *TellorMasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorMaster.Contract.TellorMasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorMaster *TellorMasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorMaster.Contract.TellorMasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorMaster *TellorMasterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TellorMaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorMaster *TellorMasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorMaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorMaster *TellorMasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorMaster.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_user address, _spender address) constant returns(uint256)
func (_TellorMaster *TellorMasterCaller) Allowance(opts *bind.CallOpts, _user common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "allowance", _user, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_user address, _spender address) constant returns(uint256)
func (_TellorMaster *TellorMasterSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _TellorMaster.Contract.Allowance(&_TellorMaster.CallOpts, _user, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_user address, _spender address) constant returns(uint256)
func (_TellorMaster *TellorMasterCallerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _TellorMaster.Contract.Allowance(&_TellorMaster.CallOpts, _user, _spender)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(_user address, _amount uint256) constant returns(bool)
func (_TellorMaster *TellorMasterCaller) AllowedToTrade(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "allowedToTrade", _user, _amount)
	return *ret0, err
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(_user address, _amount uint256) constant returns(bool)
func (_TellorMaster *TellorMasterSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _TellorMaster.Contract.AllowedToTrade(&_TellorMaster.CallOpts, _user, _amount)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(_user address, _amount uint256) constant returns(bool)
func (_TellorMaster *TellorMasterCallerSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _TellorMaster.Contract.AllowedToTrade(&_TellorMaster.CallOpts, _user, _amount)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_user address) constant returns(uint256)
func (_TellorMaster *TellorMasterCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "balanceOf", _user)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_user address) constant returns(uint256)
func (_TellorMaster *TellorMasterSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _TellorMaster.Contract.BalanceOf(&_TellorMaster.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_user address) constant returns(uint256)
func (_TellorMaster *TellorMasterCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _TellorMaster.Contract.BalanceOf(&_TellorMaster.CallOpts, _user)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(_user address, _blockNumber uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterCaller) BalanceOfAt(opts *bind.CallOpts, _user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "balanceOfAt", _user, _blockNumber)
	return *ret0, err
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(_user address, _blockNumber uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _TellorMaster.Contract.BalanceOfAt(&_TellorMaster.CallOpts, _user, _blockNumber)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(_user address, _blockNumber uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterCallerSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _TellorMaster.Contract.BalanceOfAt(&_TellorMaster.CallOpts, _user, _blockNumber)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(_challenge bytes32, _miner address) constant returns(bool)
func (_TellorMaster *TellorMasterCaller) DidMine(opts *bind.CallOpts, _challenge [32]byte, _miner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "didMine", _challenge, _miner)
	return *ret0, err
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(_challenge bytes32, _miner address) constant returns(bool)
func (_TellorMaster *TellorMasterSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _TellorMaster.Contract.DidMine(&_TellorMaster.CallOpts, _challenge, _miner)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(_challenge bytes32, _miner address) constant returns(bool)
func (_TellorMaster *TellorMasterCallerSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _TellorMaster.Contract.DidMine(&_TellorMaster.CallOpts, _challenge, _miner)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(_disputeId uint256, _address address) constant returns(bool)
func (_TellorMaster *TellorMasterCaller) DidVote(opts *bind.CallOpts, _disputeId *big.Int, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "didVote", _disputeId, _address)
	return *ret0, err
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(_disputeId uint256, _address address) constant returns(bool)
func (_TellorMaster *TellorMasterSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _TellorMaster.Contract.DidVote(&_TellorMaster.CallOpts, _disputeId, _address)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(_disputeId uint256, _address address) constant returns(bool)
func (_TellorMaster *TellorMasterCallerSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _TellorMaster.Contract.DidVote(&_TellorMaster.CallOpts, _disputeId, _address)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(_data bytes32) constant returns(address)
func (_TellorMaster *TellorMasterCaller) GetAddressVars(opts *bind.CallOpts, _data [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "getAddressVars", _data)
	return *ret0, err
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(_data bytes32) constant returns(address)
func (_TellorMaster *TellorMasterSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _TellorMaster.Contract.GetAddressVars(&_TellorMaster.CallOpts, _data)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(_data bytes32) constant returns(address)
func (_TellorMaster *TellorMasterCallerSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _TellorMaster.Contract.GetAddressVars(&_TellorMaster.CallOpts, _data)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(_disputeId uint256) constant returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_TellorMaster *TellorMasterCaller) GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
		ret2 = new(bool)
		ret3 = new(bool)
		ret4 = new(common.Address)
		ret5 = new(common.Address)
		ret6 = new(common.Address)
		ret7 = new([9]*big.Int)
		ret8 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
		ret7,
		ret8,
	}
	err := _TellorMaster.contract.Call(opts, out, "getAllDisputeVars", _disputeId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, *ret7, *ret8, err
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(_disputeId uint256) constant returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_TellorMaster *TellorMasterSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _TellorMaster.Contract.GetAllDisputeVars(&_TellorMaster.CallOpts, _disputeId)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(_disputeId uint256) constant returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_TellorMaster *TellorMasterCallerSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _TellorMaster.Contract.GetAllDisputeVars(&_TellorMaster.CallOpts, _disputeId)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() constant returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_TellorMaster *TellorMasterCaller) GetCurrentVariables(opts *bind.CallOpts) ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(string)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _TellorMaster.contract.Call(opts, out, "getCurrentVariables")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() constant returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_TellorMaster *TellorMasterSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _TellorMaster.Contract.GetCurrentVariables(&_TellorMaster.CallOpts)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() constant returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_TellorMaster *TellorMasterCallerSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _TellorMaster.Contract.GetCurrentVariables(&_TellorMaster.CallOpts)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(_hash bytes32) constant returns(uint256)
func (_TellorMaster *TellorMasterCaller) GetDisputeIdByDisputeHash(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "getDisputeIdByDisputeHash", _hash)
	return *ret0, err
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(_hash bytes32) constant returns(uint256)
func (_TellorMaster *TellorMasterSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _TellorMaster.Contract.GetDisputeIdByDisputeHash(&_TellorMaster.CallOpts, _hash)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(_hash bytes32) constant returns(uint256)
func (_TellorMaster *TellorMasterCallerSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _TellorMaster.Contract.GetDisputeIdByDisputeHash(&_TellorMaster.CallOpts, _hash)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(_disputeId uint256, _data bytes32) constant returns(uint256)
func (_TellorMaster *TellorMasterCaller) GetDisputeUintVars(opts *bind.CallOpts, _disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "getDisputeUintVars", _disputeId, _data)
	return *ret0, err
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(_disputeId uint256, _data bytes32) constant returns(uint256)
func (_TellorMaster *TellorMasterSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _TellorMaster.Contract.GetDisputeUintVars(&_TellorMaster.CallOpts, _disputeId, _data)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(_disputeId uint256, _data bytes32) constant returns(uint256)
func (_TellorMaster *TellorMasterCallerSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _TellorMaster.Contract.GetDisputeUintVars(&_TellorMaster.CallOpts, _disputeId, _data)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() constant returns(uint256, bool)
func (_TellorMaster *TellorMasterCaller) GetLastNewValue(opts *bind.CallOpts) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _TellorMaster.contract.Call(opts, out, "getLastNewValue")
	return *ret0, *ret1, err
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() constant returns(uint256, bool)
func (_TellorMaster *TellorMasterSession) GetLastNewValue() (*big.Int, bool, error) {
	return _TellorMaster.Contract.GetLastNewValue(&_TellorMaster.CallOpts)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() constant returns(uint256, bool)
func (_TellorMaster *TellorMasterCallerSession) GetLastNewValue() (*big.Int, bool, error) {
	return _TellorMaster.Contract.GetLastNewValue(&_TellorMaster.CallOpts)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(_requestId uint256) constant returns(uint256, bool)
func (_TellorMaster *TellorMasterCaller) GetLastNewValueById(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _TellorMaster.contract.Call(opts, out, "getLastNewValueById", _requestId)
	return *ret0, *ret1, err
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(_requestId uint256) constant returns(uint256, bool)
func (_TellorMaster *TellorMasterSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _TellorMaster.Contract.GetLastNewValueById(&_TellorMaster.CallOpts, _requestId)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(_requestId uint256) constant returns(uint256, bool)
func (_TellorMaster *TellorMasterCallerSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _TellorMaster.Contract.GetLastNewValueById(&_TellorMaster.CallOpts, _requestId)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterCaller) GetMinedBlockNum(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "getMinedBlockNum", _requestId, _timestamp)
	return *ret0, err
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorMaster.Contract.GetMinedBlockNum(&_TellorMaster.CallOpts, _requestId, _timestamp)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterCallerSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorMaster.Contract.GetMinedBlockNum(&_TellorMaster.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(_requestId uint256, _timestamp uint256) constant returns(address[5])
func (_TellorMaster *TellorMasterCaller) GetMinersByRequestIdAndTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	var (
		ret0 = new([5]common.Address)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "getMinersByRequestIdAndTimestamp", _requestId, _timestamp)
	return *ret0, err
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(_requestId uint256, _timestamp uint256) constant returns(address[5])
func (_TellorMaster *TellorMasterSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _TellorMaster.Contract.GetMinersByRequestIdAndTimestamp(&_TellorMaster.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(_requestId uint256, _timestamp uint256) constant returns(address[5])
func (_TellorMaster *TellorMasterCallerSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _TellorMaster.Contract.GetMinersByRequestIdAndTimestamp(&_TellorMaster.CallOpts, _requestId, _timestamp)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_TellorMaster *TellorMasterCaller) GetName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "getName")
	return *ret0, err
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_TellorMaster *TellorMasterSession) GetName() (string, error) {
	return _TellorMaster.Contract.GetName(&_TellorMaster.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_TellorMaster *TellorMasterCallerSession) GetName() (string, error) {
	return _TellorMaster.Contract.GetName(&_TellorMaster.CallOpts)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(_requestId uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "getNewValueCountbyRequestId", _requestId)
	return *ret0, err
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(_requestId uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _TellorMaster.Contract.GetNewValueCountbyRequestId(&_TellorMaster.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(_requestId uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _TellorMaster.Contract.GetNewValueCountbyRequestId(&_TellorMaster.CallOpts, _requestId)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(_request bytes32) constant returns(uint256)
func (_TellorMaster *TellorMasterCaller) GetRequestIdByQueryHash(opts *bind.CallOpts, _request [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "getRequestIdByQueryHash", _request)
	return *ret0, err
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(_request bytes32) constant returns(uint256)
func (_TellorMaster *TellorMasterSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _TellorMaster.Contract.GetRequestIdByQueryHash(&_TellorMaster.CallOpts, _request)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(_request bytes32) constant returns(uint256)
func (_TellorMaster *TellorMasterCallerSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _TellorMaster.Contract.GetRequestIdByQueryHash(&_TellorMaster.CallOpts, _request)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(_index uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterCaller) GetRequestIdByRequestQIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "getRequestIdByRequestQIndex", _index)
	return *ret0, err
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(_index uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _TellorMaster.Contract.GetRequestIdByRequestQIndex(&_TellorMaster.CallOpts, _index)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(_index uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterCallerSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _TellorMaster.Contract.GetRequestIdByRequestQIndex(&_TellorMaster.CallOpts, _index)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(_timestamp uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterCaller) GetRequestIdByTimestamp(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "getRequestIdByTimestamp", _timestamp)
	return *ret0, err
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(_timestamp uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _TellorMaster.Contract.GetRequestIdByTimestamp(&_TellorMaster.CallOpts, _timestamp)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(_timestamp uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterCallerSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _TellorMaster.Contract.GetRequestIdByTimestamp(&_TellorMaster.CallOpts, _timestamp)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() constant returns(uint256[51])
func (_TellorMaster *TellorMasterCaller) GetRequestQ(opts *bind.CallOpts) ([51]*big.Int, error) {
	var (
		ret0 = new([51]*big.Int)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "getRequestQ")
	return *ret0, err
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() constant returns(uint256[51])
func (_TellorMaster *TellorMasterSession) GetRequestQ() ([51]*big.Int, error) {
	return _TellorMaster.Contract.GetRequestQ(&_TellorMaster.CallOpts)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() constant returns(uint256[51])
func (_TellorMaster *TellorMasterCallerSession) GetRequestQ() ([51]*big.Int, error) {
	return _TellorMaster.Contract.GetRequestQ(&_TellorMaster.CallOpts)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(_requestId uint256, _data bytes32) constant returns(uint256)
func (_TellorMaster *TellorMasterCaller) GetRequestUintVars(opts *bind.CallOpts, _requestId *big.Int, _data [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "getRequestUintVars", _requestId, _data)
	return *ret0, err
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(_requestId uint256, _data bytes32) constant returns(uint256)
func (_TellorMaster *TellorMasterSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _TellorMaster.Contract.GetRequestUintVars(&_TellorMaster.CallOpts, _requestId, _data)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(_requestId uint256, _data bytes32) constant returns(uint256)
func (_TellorMaster *TellorMasterCallerSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _TellorMaster.Contract.GetRequestUintVars(&_TellorMaster.CallOpts, _requestId, _data)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(_requestId uint256) constant returns(string, string, bytes32, uint256, uint256, uint256)
func (_TellorMaster *TellorMasterCaller) GetRequestVars(opts *bind.CallOpts, _requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new([32]byte)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _TellorMaster.contract.Call(opts, out, "getRequestVars", _requestId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(_requestId uint256) constant returns(string, string, bytes32, uint256, uint256, uint256)
func (_TellorMaster *TellorMasterSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _TellorMaster.Contract.GetRequestVars(&_TellorMaster.CallOpts, _requestId)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(_requestId uint256) constant returns(string, string, bytes32, uint256, uint256, uint256)
func (_TellorMaster *TellorMasterCallerSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _TellorMaster.Contract.GetRequestVars(&_TellorMaster.CallOpts, _requestId)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(_staker address) constant returns(uint256, uint256)
func (_TellorMaster *TellorMasterCaller) GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _TellorMaster.contract.Call(opts, out, "getStakerInfo", _staker)
	return *ret0, *ret1, err
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(_staker address) constant returns(uint256, uint256)
func (_TellorMaster *TellorMasterSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _TellorMaster.Contract.GetStakerInfo(&_TellorMaster.CallOpts, _staker)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(_staker address) constant returns(uint256, uint256)
func (_TellorMaster *TellorMasterCallerSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _TellorMaster.Contract.GetStakerInfo(&_TellorMaster.CallOpts, _staker)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(_requestId uint256, _timestamp uint256) constant returns(uint256[5])
func (_TellorMaster *TellorMasterCaller) GetSubmissionsByTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	var (
		ret0 = new([5]*big.Int)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "getSubmissionsByTimestamp", _requestId, _timestamp)
	return *ret0, err
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(_requestId uint256, _timestamp uint256) constant returns(uint256[5])
func (_TellorMaster *TellorMasterSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _TellorMaster.Contract.GetSubmissionsByTimestamp(&_TellorMaster.CallOpts, _requestId, _timestamp)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(_requestId uint256, _timestamp uint256) constant returns(uint256[5])
func (_TellorMaster *TellorMasterCallerSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _TellorMaster.Contract.GetSubmissionsByTimestamp(&_TellorMaster.CallOpts, _requestId, _timestamp)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() constant returns(string)
func (_TellorMaster *TellorMasterCaller) GetSymbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "getSymbol")
	return *ret0, err
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() constant returns(string)
func (_TellorMaster *TellorMasterSession) GetSymbol() (string, error) {
	return _TellorMaster.Contract.GetSymbol(&_TellorMaster.CallOpts)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() constant returns(string)
func (_TellorMaster *TellorMasterCallerSession) GetSymbol() (string, error) {
	return _TellorMaster.Contract.GetSymbol(&_TellorMaster.CallOpts)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(_requestID uint256, _index uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "getTimestampbyRequestIDandIndex", _requestID, _index)
	return *ret0, err
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(_requestID uint256, _index uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _TellorMaster.Contract.GetTimestampbyRequestIDandIndex(&_TellorMaster.CallOpts, _requestID, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(_requestID uint256, _index uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterCallerSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _TellorMaster.Contract.GetTimestampbyRequestIDandIndex(&_TellorMaster.CallOpts, _requestID, _index)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(_data bytes32) constant returns(uint256)
func (_TellorMaster *TellorMasterCaller) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "getUintVar", _data)
	return *ret0, err
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(_data bytes32) constant returns(uint256)
func (_TellorMaster *TellorMasterSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _TellorMaster.Contract.GetUintVar(&_TellorMaster.CallOpts, _data)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(_data bytes32) constant returns(uint256)
func (_TellorMaster *TellorMasterCallerSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _TellorMaster.Contract.GetUintVar(&_TellorMaster.CallOpts, _data)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() constant returns(uint256, uint256, string)
func (_TellorMaster *TellorMasterCaller) GetVariablesOnDeck(opts *bind.CallOpts) (*big.Int, *big.Int, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _TellorMaster.contract.Call(opts, out, "getVariablesOnDeck")
	return *ret0, *ret1, *ret2, err
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() constant returns(uint256, uint256, string)
func (_TellorMaster *TellorMasterSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _TellorMaster.Contract.GetVariablesOnDeck(&_TellorMaster.CallOpts)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() constant returns(uint256, uint256, string)
func (_TellorMaster *TellorMasterCallerSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _TellorMaster.Contract.GetVariablesOnDeck(&_TellorMaster.CallOpts)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(_requestId uint256, _timestamp uint256) constant returns(bool)
func (_TellorMaster *TellorMasterCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "isInDispute", _requestId, _timestamp)
	return *ret0, err
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(_requestId uint256, _timestamp uint256) constant returns(bool)
func (_TellorMaster *TellorMasterSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _TellorMaster.Contract.IsInDispute(&_TellorMaster.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(_requestId uint256, _timestamp uint256) constant returns(bool)
func (_TellorMaster *TellorMasterCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _TellorMaster.Contract.IsInDispute(&_TellorMaster.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "retrieveData", _requestId, _timestamp)
	return *ret0, err
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorMaster.Contract.RetrieveData(&_TellorMaster.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_TellorMaster *TellorMasterCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorMaster.Contract.RetrieveData(&_TellorMaster.CallOpts, _requestId, _timestamp)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TellorMaster *TellorMasterCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TellorMaster.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TellorMaster *TellorMasterSession) TotalSupply() (*big.Int, error) {
	return _TellorMaster.Contract.TotalSupply(&_TellorMaster.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TellorMaster *TellorMasterCallerSession) TotalSupply() (*big.Int, error) {
	return _TellorMaster.Contract.TotalSupply(&_TellorMaster.CallOpts)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(_newDeity address) returns()
func (_TellorMaster *TellorMasterTransactor) ChangeDeity(opts *bind.TransactOpts, _newDeity common.Address) (*types.Transaction, error) {
	return _TellorMaster.contract.Transact(opts, "changeDeity", _newDeity)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(_newDeity address) returns()
func (_TellorMaster *TellorMasterSession) ChangeDeity(_newDeity common.Address) (*types.Transaction, error) {
	return _TellorMaster.Contract.ChangeDeity(&_TellorMaster.TransactOpts, _newDeity)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(_newDeity address) returns()
func (_TellorMaster *TellorMasterTransactorSession) ChangeDeity(_newDeity common.Address) (*types.Transaction, error) {
	return _TellorMaster.Contract.ChangeDeity(&_TellorMaster.TransactOpts, _newDeity)
}

// ChangeTellorContract is a paid mutator transaction binding the contract method 0xae0a8279.
//
// Solidity: function changeTellorContract(_tellorContract address) returns()
func (_TellorMaster *TellorMasterTransactor) ChangeTellorContract(opts *bind.TransactOpts, _tellorContract common.Address) (*types.Transaction, error) {
	return _TellorMaster.contract.Transact(opts, "changeTellorContract", _tellorContract)
}

// ChangeTellorContract is a paid mutator transaction binding the contract method 0xae0a8279.
//
// Solidity: function changeTellorContract(_tellorContract address) returns()
func (_TellorMaster *TellorMasterSession) ChangeTellorContract(_tellorContract common.Address) (*types.Transaction, error) {
	return _TellorMaster.Contract.ChangeTellorContract(&_TellorMaster.TransactOpts, _tellorContract)
}

// ChangeTellorContract is a paid mutator transaction binding the contract method 0xae0a8279.
//
// Solidity: function changeTellorContract(_tellorContract address) returns()
func (_TellorMaster *TellorMasterTransactorSession) ChangeTellorContract(_tellorContract common.Address) (*types.Transaction, error) {
	return _TellorMaster.Contract.ChangeTellorContract(&_TellorMaster.TransactOpts, _tellorContract)
}

// TellorMasterNewTellorAddressIterator is returned from FilterNewTellorAddress and is used to iterate over the raw logs and unpacked data for NewTellorAddress events raised by the TellorMaster contract.
type TellorMasterNewTellorAddressIterator struct {
	Event *TellorMasterNewTellorAddress // Event containing the contract specifics and raw log

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
func (it *TellorMasterNewTellorAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorMasterNewTellorAddress)
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
		it.Event = new(TellorMasterNewTellorAddress)
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
func (it *TellorMasterNewTellorAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorMasterNewTellorAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorMasterNewTellorAddress represents a NewTellorAddress event raised by the TellorMaster contract.
type TellorMasterNewTellorAddress struct {
	NewTellor common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTellorAddress is a free log retrieval operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: e NewTellorAddress(_newTellor address)
func (_TellorMaster *TellorMasterFilterer) FilterNewTellorAddress(opts *bind.FilterOpts) (*TellorMasterNewTellorAddressIterator, error) {

	logs, sub, err := _TellorMaster.contract.FilterLogs(opts, "NewTellorAddress")
	if err != nil {
		return nil, err
	}
	return &TellorMasterNewTellorAddressIterator{contract: _TellorMaster.contract, event: "NewTellorAddress", logs: logs, sub: sub}, nil
}

// WatchNewTellorAddress is a free log subscription operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: e NewTellorAddress(_newTellor address)
func (_TellorMaster *TellorMasterFilterer) WatchNewTellorAddress(opts *bind.WatchOpts, sink chan<- *TellorMasterNewTellorAddress) (event.Subscription, error) {

	logs, sub, err := _TellorMaster.contract.WatchLogs(opts, "NewTellorAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorMasterNewTellorAddress)
				if err := _TellorMaster.contract.UnpackLog(event, "NewTellorAddress", log); err != nil {
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

// TellorStakeABI is the input ABI used to generate the binding from.
const TellorStakeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"NewStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawRequested\",\"type\":\"event\"}]"

// TellorStakeBin is the compiled bytecode used for deploying new contracts.
const TellorStakeBin = `0x6108c3610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c806344bacc4b1461005b5780634601f1cd14610087578063820a2d66146100b1578063c9cf5e4c146100db575b600080fd5b81801561006757600080fd5b506100856004803603602081101561007e57600080fd5b5035610105565b005b81801561009357600080fd5b50610085600480360360208110156100aa57600080fd5b5035610173565b8180156100bd57600080fd5b50610085600480360360208110156100d457600080fd5b5035610559565b8180156100e757600080fd5b50610085600480360360208110156100fe57600080fd5b50356105cd565b3360009081526047820160205260409020600181015462093a8090620151804206420303101561013457600080fd5b805460021461014257600080fd5b600080825560405133917f4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec91a25050565b6040805167646563696d616c7360c01b8152815190819003600801902060009081528183016020522054156101a757600080fd5b3060009081526045820160205260408082208151631d6f7b8160e31b8152600481019190915269014542ba12a337c00000196024820152905173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9263eb7bdc089260448082019391829003018186803b15801561021757600080fd5b505af415801561022b573d6000803e3d6000fd5b50505050610237610870565b506040805160c08101825273e037ec8ec9ec423826750853899394de7f024fee815273cdd8fa31af8475574b8909f135d510579a8087d3602082015273b9dd5afd86547df817da2d0fb89334a6f8edd8919181019190915273230570cd052f40e14c14a81038c6f3aa685d712b6060820152733233afa02644ccd048587f8ba6e99b3c00a34dcc608082015273e010ac6e0248790e08f42d5f697160dedf97e02460a082015260005b60068110156103b95773__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__63eb7bdc0884604501600085856006811061031557fe5b60200201516001600160a01b03166001600160a01b03168152602001908152602001600020683635c9adc5dea000006040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561037f57600080fd5b505af4158015610393573d6000803e3d6000fd5b505050506103b1838383600681106103a757fe5b60200201516106c4565b6001016102e0565b50604080516b746f74616c5f737570706c7960a01b8152815190819003600c908101822060009081528386016020818152858320805469014542ba12a337c0000001905567646563696d616c7360c01b855285519485900360080185208352818152858320601290556b7461726765744d696e65727360a01b85528551948590039093018420825280835284822060c890556a1cdd185ad9505b5bdd5b9d60aa1b8452845193849003600b0184208252808352848220683635c9adc5dea000009055696469737075746546656560b01b8452845193849003600a908101852083528184528583206834957444b840e800009055691d1a5b5955185c99d95d60b21b80865286519586900382018620845282855286842061025890558552855194859003019093208152919052205442816104ef57fe5b604080517174696d654f664c6173744e657756616c756560701b815281519081900360120181206000908152958201602081815283882095909406420390945569646966666963756c747960b01b8152815190819003600a01902085529190529091206001905550565b61056381336106c4565b73__$5c50ea773c3f1223822c80c8edbce14ea7$__63e15f6f70826040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b1580156105b257600080fd5b505af41580156105c6573d6000803e3d6000fd5b5050505050565b336000908152604782016020526040902080546001146105ec57600080fd5b6002815562015180420642036001820155604080516a1cdd185ad95c90dbdd5b9d60aa1b8152815190819003600b01812060009081528285016020528281208054600019019055630e15f6f760e41b825260048201859052915173__$5c50ea773c3f1223822c80c8edbce14ea7$__9263e15f6f709260248082019391829003018186803b15801561067d57600080fd5b505af4158015610691573d6000803e3d6000fd5b50506040513392507f453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf9150600090a25050565b604080516a1cdd185ad9505b5bdd5b9d60aa1b8152815190819003600b0181206000908152828501602090815290839020546393b182b360e01b8352600483018690526001600160a01b0385166024840152925173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__926393b182b3926044808301939192829003018186803b15801561075057600080fd5b505af4158015610764573d6000803e3d6000fd5b505050506040513d602081101561077a57600080fd5b5051101561078757600080fd5b6001600160a01b038116600090815260478301602052604090205415806107c857506001600160a01b03811660009081526047830160205260409020546002145b6107d157600080fd5b604080516a1cdd185ad95c90dbdd5b9d60aa1b8152815190819003600b01812060009081528483016020908152838220805460019081019091558385018552808452620151804290810690038285019081526001600160a01b038716808552604789019093528584209451855551930192909255915190917f46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e291a25050565b6040518060c00160405280600690602082028038833950919291505056fea265627a7a723058203cc3787f4bf5992f9b9f3d798c1bfdb3f004f46517aa8f3cf7a95e43f392268564736f6c634300050a0032`

// DeployTellorStake deploys a new Ethereum contract, binding an instance of TellorStake to it.
func DeployTellorStake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorStake, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorStakeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TellorStakeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TellorStake{TellorStakeCaller: TellorStakeCaller{contract: contract}, TellorStakeTransactor: TellorStakeTransactor{contract: contract}, TellorStakeFilterer: TellorStakeFilterer{contract: contract}}, nil
}

// TellorStake is an auto generated Go binding around an Ethereum contract.
type TellorStake struct {
	TellorStakeCaller     // Read-only binding to the contract
	TellorStakeTransactor // Write-only binding to the contract
	TellorStakeFilterer   // Log filterer for contract events
}

// TellorStakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorStakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorStakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorStakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorStakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorStakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorStakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorStakeSession struct {
	Contract     *TellorStake      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorStakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorStakeCallerSession struct {
	Contract *TellorStakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TellorStakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorStakeTransactorSession struct {
	Contract     *TellorStakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TellorStakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorStakeRaw struct {
	Contract *TellorStake // Generic contract binding to access the raw methods on
}

// TellorStakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorStakeCallerRaw struct {
	Contract *TellorStakeCaller // Generic read-only contract binding to access the raw methods on
}

// TellorStakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorStakeTransactorRaw struct {
	Contract *TellorStakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellorStake creates a new instance of TellorStake, bound to a specific deployed contract.
func NewTellorStake(address common.Address, backend bind.ContractBackend) (*TellorStake, error) {
	contract, err := bindTellorStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TellorStake{TellorStakeCaller: TellorStakeCaller{contract: contract}, TellorStakeTransactor: TellorStakeTransactor{contract: contract}, TellorStakeFilterer: TellorStakeFilterer{contract: contract}}, nil
}

// NewTellorStakeCaller creates a new read-only instance of TellorStake, bound to a specific deployed contract.
func NewTellorStakeCaller(address common.Address, caller bind.ContractCaller) (*TellorStakeCaller, error) {
	contract, err := bindTellorStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorStakeCaller{contract: contract}, nil
}

// NewTellorStakeTransactor creates a new write-only instance of TellorStake, bound to a specific deployed contract.
func NewTellorStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorStakeTransactor, error) {
	contract, err := bindTellorStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorStakeTransactor{contract: contract}, nil
}

// NewTellorStakeFilterer creates a new log filterer instance of TellorStake, bound to a specific deployed contract.
func NewTellorStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorStakeFilterer, error) {
	contract, err := bindTellorStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorStakeFilterer{contract: contract}, nil
}

// bindTellorStake binds a generic wrapper to an already deployed contract.
func bindTellorStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorStakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorStake *TellorStakeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TellorStake.Contract.TellorStakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorStake *TellorStakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorStake.Contract.TellorStakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorStake *TellorStakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorStake.Contract.TellorStakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorStake *TellorStakeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TellorStake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorStake *TellorStakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorStake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorStake *TellorStakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorStake.Contract.contract.Transact(opts, method, params...)
}

// TellorStakeNewStakeIterator is returned from FilterNewStake and is used to iterate over the raw logs and unpacked data for NewStake events raised by the TellorStake contract.
type TellorStakeNewStakeIterator struct {
	Event *TellorStakeNewStake // Event containing the contract specifics and raw log

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
func (it *TellorStakeNewStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorStakeNewStake)
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
		it.Event = new(TellorStakeNewStake)
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
func (it *TellorStakeNewStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorStakeNewStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorStakeNewStake represents a NewStake event raised by the TellorStake contract.
type TellorStakeNewStake struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewStake is a free log retrieval operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: e NewStake(_sender indexed address)
func (_TellorStake *TellorStakeFilterer) FilterNewStake(opts *bind.FilterOpts, _sender []common.Address) (*TellorStakeNewStakeIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _TellorStake.contract.FilterLogs(opts, "NewStake", _senderRule)
	if err != nil {
		return nil, err
	}
	return &TellorStakeNewStakeIterator{contract: _TellorStake.contract, event: "NewStake", logs: logs, sub: sub}, nil
}

// WatchNewStake is a free log subscription operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: e NewStake(_sender indexed address)
func (_TellorStake *TellorStakeFilterer) WatchNewStake(opts *bind.WatchOpts, sink chan<- *TellorStakeNewStake, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _TellorStake.contract.WatchLogs(opts, "NewStake", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorStakeNewStake)
				if err := _TellorStake.contract.UnpackLog(event, "NewStake", log); err != nil {
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

// TellorStakeStakeWithdrawRequestedIterator is returned from FilterStakeWithdrawRequested and is used to iterate over the raw logs and unpacked data for StakeWithdrawRequested events raised by the TellorStake contract.
type TellorStakeStakeWithdrawRequestedIterator struct {
	Event *TellorStakeStakeWithdrawRequested // Event containing the contract specifics and raw log

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
func (it *TellorStakeStakeWithdrawRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorStakeStakeWithdrawRequested)
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
		it.Event = new(TellorStakeStakeWithdrawRequested)
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
func (it *TellorStakeStakeWithdrawRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorStakeStakeWithdrawRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorStakeStakeWithdrawRequested represents a StakeWithdrawRequested event raised by the TellorStake contract.
type TellorStakeStakeWithdrawRequested struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawRequested is a free log retrieval operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: e StakeWithdrawRequested(_sender indexed address)
func (_TellorStake *TellorStakeFilterer) FilterStakeWithdrawRequested(opts *bind.FilterOpts, _sender []common.Address) (*TellorStakeStakeWithdrawRequestedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _TellorStake.contract.FilterLogs(opts, "StakeWithdrawRequested", _senderRule)
	if err != nil {
		return nil, err
	}
	return &TellorStakeStakeWithdrawRequestedIterator{contract: _TellorStake.contract, event: "StakeWithdrawRequested", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawRequested is a free log subscription operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: e StakeWithdrawRequested(_sender indexed address)
func (_TellorStake *TellorStakeFilterer) WatchStakeWithdrawRequested(opts *bind.WatchOpts, sink chan<- *TellorStakeStakeWithdrawRequested, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _TellorStake.contract.WatchLogs(opts, "StakeWithdrawRequested", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorStakeStakeWithdrawRequested)
				if err := _TellorStake.contract.UnpackLog(event, "StakeWithdrawRequested", log); err != nil {
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

// TellorStakeStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the TellorStake contract.
type TellorStakeStakeWithdrawnIterator struct {
	Event *TellorStakeStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *TellorStakeStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorStakeStakeWithdrawn)
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
		it.Event = new(TellorStakeStakeWithdrawn)
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
func (it *TellorStakeStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorStakeStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorStakeStakeWithdrawn represents a StakeWithdrawn event raised by the TellorStake contract.
type TellorStakeStakeWithdrawn struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: e StakeWithdrawn(_sender indexed address)
func (_TellorStake *TellorStakeFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, _sender []common.Address) (*TellorStakeStakeWithdrawnIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _TellorStake.contract.FilterLogs(opts, "StakeWithdrawn", _senderRule)
	if err != nil {
		return nil, err
	}
	return &TellorStakeStakeWithdrawnIterator{contract: _TellorStake.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: e StakeWithdrawn(_sender indexed address)
func (_TellorStake *TellorStakeFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *TellorStakeStakeWithdrawn, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _TellorStake.contract.WatchLogs(opts, "StakeWithdrawn", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorStakeStakeWithdrawn)
				if err := _TellorStake.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// TellorStorageABI is the input ABI used to generate the binding from.
const TellorStorageABI = "[]"

// TellorStorageBin is the compiled bytecode used for deploying new contracts.
const TellorStorageBin = `0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7230582058ec2f862ab3051bed82bfb7dfd8444570e7f6447a286cf14fc693a0a697f7c664736f6c634300050a0032`

// DeployTellorStorage deploys a new Ethereum contract, binding an instance of TellorStorage to it.
func DeployTellorStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TellorStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TellorStorage{TellorStorageCaller: TellorStorageCaller{contract: contract}, TellorStorageTransactor: TellorStorageTransactor{contract: contract}, TellorStorageFilterer: TellorStorageFilterer{contract: contract}}, nil
}

// TellorStorage is an auto generated Go binding around an Ethereum contract.
type TellorStorage struct {
	TellorStorageCaller     // Read-only binding to the contract
	TellorStorageTransactor // Write-only binding to the contract
	TellorStorageFilterer   // Log filterer for contract events
}

// TellorStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorStorageSession struct {
	Contract     *TellorStorage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorStorageCallerSession struct {
	Contract *TellorStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TellorStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorStorageTransactorSession struct {
	Contract     *TellorStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TellorStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorStorageRaw struct {
	Contract *TellorStorage // Generic contract binding to access the raw methods on
}

// TellorStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorStorageCallerRaw struct {
	Contract *TellorStorageCaller // Generic read-only contract binding to access the raw methods on
}

// TellorStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorStorageTransactorRaw struct {
	Contract *TellorStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellorStorage creates a new instance of TellorStorage, bound to a specific deployed contract.
func NewTellorStorage(address common.Address, backend bind.ContractBackend) (*TellorStorage, error) {
	contract, err := bindTellorStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TellorStorage{TellorStorageCaller: TellorStorageCaller{contract: contract}, TellorStorageTransactor: TellorStorageTransactor{contract: contract}, TellorStorageFilterer: TellorStorageFilterer{contract: contract}}, nil
}

// NewTellorStorageCaller creates a new read-only instance of TellorStorage, bound to a specific deployed contract.
func NewTellorStorageCaller(address common.Address, caller bind.ContractCaller) (*TellorStorageCaller, error) {
	contract, err := bindTellorStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorStorageCaller{contract: contract}, nil
}

// NewTellorStorageTransactor creates a new write-only instance of TellorStorage, bound to a specific deployed contract.
func NewTellorStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorStorageTransactor, error) {
	contract, err := bindTellorStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorStorageTransactor{contract: contract}, nil
}

// NewTellorStorageFilterer creates a new log filterer instance of TellorStorage, bound to a specific deployed contract.
func NewTellorStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorStorageFilterer, error) {
	contract, err := bindTellorStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorStorageFilterer{contract: contract}, nil
}

// bindTellorStorage binds a generic wrapper to an already deployed contract.
func bindTellorStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorStorage *TellorStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TellorStorage.Contract.TellorStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorStorage *TellorStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorStorage.Contract.TellorStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorStorage *TellorStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorStorage.Contract.TellorStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorStorage *TellorStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TellorStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorStorage *TellorStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorStorage *TellorStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorStorage.Contract.contract.Transact(opts, method, params...)
}

// TellorTransferABI is the input ABI used to generate the binding from.
const TellorTransferABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// TellorTransferBin is the compiled bytecode used for deploying new contracts.
const TellorTransferBin = `0x610936610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100a85760003560e01c8063bf32006c11610070578063bf32006c146101c5578063c7bb46ad146101f9578063c84b96f514610244578063ca50189914610283578063eb7bdc08146102cc576100a8565b80633f48b1ff146100ad578063850dcc32146100f157806393b182b3146101445780639be5647f14610170578063acaab9e214610193575b600080fd5b6100df600480360360608110156100c357600080fd5b508035906001600160a01b0360208201351690604001356102fc565b60408051918252519081900360200190f35b8180156100fd57600080fd5b506101306004803603606081101561011457600080fd5b508035906001600160a01b036020820135169060400135610395565b604080519115158252519081900360200190f35b6100df6004803603604081101561015a57600080fd5b50803590602001356001600160a01b0316610428565b6100df6004803603604081101561018657600080fd5b508035906020013561043e565b610130600480360360608110156101a957600080fd5b508035906001600160a01b03602082013516906040013561056e565b6100df600480360360608110156101db57600080fd5b508035906001600160a01b0360208201358116916040013516610614565b81801561020557600080fd5b506102426004803603608081101561021c57600080fd5b508035906001600160a01b03602082013581169160408101359091169060600135610641565b005b81801561025057600080fd5b506101306004803603606081101561026757600080fd5b508035906001600160a01b03602082013516906040013561073f565b81801561028f57600080fd5b50610130600480360360808110156102a657600080fd5b508035906001600160a01b03602082013581169160408101359091169060600135610757565b8180156102d857600080fd5b50610242600480360360408110156102ef57600080fd5b50803590602001356107cc565b6001600160a01b0382166000908152604584016020526040812054158061035a57506001600160a01b03831660009081526045850160205260408120805484929061034357fe5b6000918252602090912001546001600160801b0316115b156103675750600061038e565b6001600160a01b0383166000908152604585016020526040902061038b908361043e565b90505b9392505050565b60006103a284338461056e565b6103ab57600080fd5b6001600160a01b0383166103be57600080fd5b33600081815260468601602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b60006104358383436102fc565b90505b92915050565b815460009061044f57506000610438565b82548390600019810190811061046157fe5b6000918252602090912001546001600160801b031682106104b15782548390600019810190811061048e57fe5b600091825260209091200154600160801b90046001600160801b03169050610438565b826000815481106104be57fe5b6000918252602090912001546001600160801b03168210156104e257506000610438565b8254600090600019015b8181111561053d57600060026001838501010490508486828154811061050e57fe5b6000918252602090912001546001600160801b03161161053057809250610537565b6001810391505b506104ec565b84828154811061054957fe5b600091825260209091200154600160801b90046001600160801b031695945050505050565b6001600160a01b0382166000908152604784016020526040812054156105ed57604080516a1cdd185ad9505b5bdd5b9d60aa1b8152815190819003600b0190206000908152818601602052908120546105db9084906105cf90818989610428565b9063ffffffff6108a516565b106105e85750600161038e565b61060a565b60006105fd836105cf8787610428565b1061060a5750600161038e565b5060009392505050565b6001600160a01b039182166000908152604693909301602090815260408085209290931684525290205490565b6000811161064e57600080fd5b6001600160a01b03821661066157600080fd5b61066c84848361056e565b61067557600080fd5b60006106828585436102fc565b6001600160a01b038516600090815260458701602052604090209091506106ab908383036107cc565b6106b68584436102fc565b90508082820110156106c757600080fd5b6001600160a01b038316600090815260458601602052604090206106ed908284016107cc565b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a35050505050565b600061074d84338585610641565b5060019392505050565b6001600160a01b0383166000908152604685016020908152604080832033845290915281205482111561078957600080fd5b6001600160a01b038416600090815260468601602090815260408083203384529091529020805483900390556107c185858585610641565b506001949350505050565b81541580610800575081544390839060001981019081106107e957fe5b6000918252602090912001546001600160801b0316105b15610867578154600090839061081982600183016108b7565b8154811061082357fe5b600091825260209091200180546001600160801b03848116600160801b024382166fffffffffffffffffffffffffffffffff199093169290921716179055506108a1565b81546000908390600019810190811061087c57fe5b600091825260209091200180546001600160801b03808516600160801b029116179055505b5050565b6000828211156108b157fe5b50900390565b8154818355818111156108db576000838152602090206108db9181019083016108e0565b505050565b6108fe91905b808211156108fa57600081556001016108e6565b5090565b9056fea265627a7a72305820782f81105098c84c9eb011f38e6970ea78434680be8cc6b5d842e4b5772271ea64736f6c634300050a0032`

// DeployTellorTransfer deploys a new Ethereum contract, binding an instance of TellorTransfer to it.
func DeployTellorTransfer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorTransfer, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorTransferABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TellorTransferBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TellorTransfer{TellorTransferCaller: TellorTransferCaller{contract: contract}, TellorTransferTransactor: TellorTransferTransactor{contract: contract}, TellorTransferFilterer: TellorTransferFilterer{contract: contract}}, nil
}

// TellorTransfer is an auto generated Go binding around an Ethereum contract.
type TellorTransfer struct {
	TellorTransferCaller     // Read-only binding to the contract
	TellorTransferTransactor // Write-only binding to the contract
	TellorTransferFilterer   // Log filterer for contract events
}

// TellorTransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorTransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorTransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorTransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorTransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorTransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorTransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorTransferSession struct {
	Contract     *TellorTransfer   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorTransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorTransferCallerSession struct {
	Contract *TellorTransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TellorTransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorTransferTransactorSession struct {
	Contract     *TellorTransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TellorTransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorTransferRaw struct {
	Contract *TellorTransfer // Generic contract binding to access the raw methods on
}

// TellorTransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorTransferCallerRaw struct {
	Contract *TellorTransferCaller // Generic read-only contract binding to access the raw methods on
}

// TellorTransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorTransferTransactorRaw struct {
	Contract *TellorTransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellorTransfer creates a new instance of TellorTransfer, bound to a specific deployed contract.
func NewTellorTransfer(address common.Address, backend bind.ContractBackend) (*TellorTransfer, error) {
	contract, err := bindTellorTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TellorTransfer{TellorTransferCaller: TellorTransferCaller{contract: contract}, TellorTransferTransactor: TellorTransferTransactor{contract: contract}, TellorTransferFilterer: TellorTransferFilterer{contract: contract}}, nil
}

// NewTellorTransferCaller creates a new read-only instance of TellorTransfer, bound to a specific deployed contract.
func NewTellorTransferCaller(address common.Address, caller bind.ContractCaller) (*TellorTransferCaller, error) {
	contract, err := bindTellorTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorTransferCaller{contract: contract}, nil
}

// NewTellorTransferTransactor creates a new write-only instance of TellorTransfer, bound to a specific deployed contract.
func NewTellorTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorTransferTransactor, error) {
	contract, err := bindTellorTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorTransferTransactor{contract: contract}, nil
}

// NewTellorTransferFilterer creates a new log filterer instance of TellorTransfer, bound to a specific deployed contract.
func NewTellorTransferFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorTransferFilterer, error) {
	contract, err := bindTellorTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorTransferFilterer{contract: contract}, nil
}

// bindTellorTransfer binds a generic wrapper to an already deployed contract.
func bindTellorTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorTransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorTransfer *TellorTransferRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TellorTransfer.Contract.TellorTransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorTransfer *TellorTransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorTransfer.Contract.TellorTransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorTransfer *TellorTransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorTransfer.Contract.TellorTransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorTransfer *TellorTransferCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TellorTransfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorTransfer *TellorTransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorTransfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorTransfer *TellorTransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorTransfer.Contract.contract.Transact(opts, method, params...)
}

// TellorTransferApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TellorTransfer contract.
type TellorTransferApprovalIterator struct {
	Event *TellorTransferApproval // Event containing the contract specifics and raw log

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
func (it *TellorTransferApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorTransferApproval)
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
		it.Event = new(TellorTransferApproval)
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
func (it *TellorTransferApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorTransferApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorTransferApproval represents a Approval event raised by the TellorTransfer contract.
type TellorTransferApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_TellorTransfer *TellorTransferFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*TellorTransferApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _TellorTransfer.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &TellorTransferApprovalIterator{contract: _TellorTransfer.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_TellorTransfer *TellorTransferFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TellorTransferApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _TellorTransfer.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorTransferApproval)
				if err := _TellorTransfer.contract.UnpackLog(event, "Approval", log); err != nil {
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

// TellorTransferTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TellorTransfer contract.
type TellorTransferTransferIterator struct {
	Event *TellorTransferTransfer // Event containing the contract specifics and raw log

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
func (it *TellorTransferTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorTransferTransfer)
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
		it.Event = new(TellorTransferTransfer)
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
func (it *TellorTransferTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorTransferTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorTransferTransfer represents a Transfer event raised by the TellorTransfer contract.
type TellorTransferTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_TellorTransfer *TellorTransferFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*TellorTransferTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _TellorTransfer.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &TellorTransferTransferIterator{contract: _TellorTransfer.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_TellorTransfer *TellorTransferFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TellorTransferTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _TellorTransfer.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorTransferTransfer)
				if err := _TellorTransfer.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// UtilitiesABI is the input ABI used to generate the binding from.
const UtilitiesABI = "[]"

// UtilitiesBin is the compiled bytecode used for deploying new contracts.
const UtilitiesBin = `0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058203a5f66bda6f7a12c4d2010ffb4ddea6b3e3fea0390b238de6b2ad2897120313b64736f6c634300050a0032`

// DeployUtilities deploys a new Ethereum contract, binding an instance of Utilities to it.
func DeployUtilities(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utilities, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilitiesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilitiesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utilities{UtilitiesCaller: UtilitiesCaller{contract: contract}, UtilitiesTransactor: UtilitiesTransactor{contract: contract}, UtilitiesFilterer: UtilitiesFilterer{contract: contract}}, nil
}

// Utilities is an auto generated Go binding around an Ethereum contract.
type Utilities struct {
	UtilitiesCaller     // Read-only binding to the contract
	UtilitiesTransactor // Write-only binding to the contract
	UtilitiesFilterer   // Log filterer for contract events
}

// UtilitiesCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilitiesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilitiesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilitiesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilitiesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilitiesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilitiesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilitiesSession struct {
	Contract     *Utilities        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilitiesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilitiesCallerSession struct {
	Contract *UtilitiesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// UtilitiesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilitiesTransactorSession struct {
	Contract     *UtilitiesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UtilitiesRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilitiesRaw struct {
	Contract *Utilities // Generic contract binding to access the raw methods on
}

// UtilitiesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilitiesCallerRaw struct {
	Contract *UtilitiesCaller // Generic read-only contract binding to access the raw methods on
}

// UtilitiesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilitiesTransactorRaw struct {
	Contract *UtilitiesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtilities creates a new instance of Utilities, bound to a specific deployed contract.
func NewUtilities(address common.Address, backend bind.ContractBackend) (*Utilities, error) {
	contract, err := bindUtilities(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utilities{UtilitiesCaller: UtilitiesCaller{contract: contract}, UtilitiesTransactor: UtilitiesTransactor{contract: contract}, UtilitiesFilterer: UtilitiesFilterer{contract: contract}}, nil
}

// NewUtilitiesCaller creates a new read-only instance of Utilities, bound to a specific deployed contract.
func NewUtilitiesCaller(address common.Address, caller bind.ContractCaller) (*UtilitiesCaller, error) {
	contract, err := bindUtilities(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilitiesCaller{contract: contract}, nil
}

// NewUtilitiesTransactor creates a new write-only instance of Utilities, bound to a specific deployed contract.
func NewUtilitiesTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilitiesTransactor, error) {
	contract, err := bindUtilities(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilitiesTransactor{contract: contract}, nil
}

// NewUtilitiesFilterer creates a new log filterer instance of Utilities, bound to a specific deployed contract.
func NewUtilitiesFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilitiesFilterer, error) {
	contract, err := bindUtilities(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilitiesFilterer{contract: contract}, nil
}

// bindUtilities binds a generic wrapper to an already deployed contract.
func bindUtilities(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilitiesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utilities *UtilitiesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utilities.Contract.UtilitiesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utilities *UtilitiesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utilities.Contract.UtilitiesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utilities *UtilitiesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utilities.Contract.UtilitiesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utilities *UtilitiesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utilities.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utilities *UtilitiesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utilities.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utilities *UtilitiesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utilities.Contract.contract.Transact(opts, method, params...)
}
