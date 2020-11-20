// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package getter

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582021470e149c5fe6b3462b7bf32ca96e71c94c08f0bc445cb6df30554221683c4b64736f6c63430005100032"

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
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
const TellorDisputeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"_result\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reportedMiner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_reportingParty\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"DisputeVoteTallied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"NewDispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newTellor\",\"type\":\"address\"}],\"name\":\"NewTellorAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_position\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_voteWeight\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"}]"

// TellorDisputeFuncSigs maps the 4-byte function signature to its string representation.
var TellorDisputeFuncSigs = map[string]string{
	"ca9a4ea5": "beginDispute(TellorStorage.TellorStorageStruct storage,uint256,uint256,uint256)",
	"694bf49f": "proposeFork(TellorStorage.TellorStorageStruct storage,address)",
	"def6fac7": "tallyVotes(TellorStorage.TellorStorageStruct storage,uint256)",
	"97f5f972": "unlockDisputeFee(TellorStorage.TellorStorageStruct storage,uint256)",
	"d7b651c1": "updateMinDisputeFee(TellorStorage.TellorStorageStruct storage)",
	"22048ecf": "updateTellor(TellorStorage.TellorStorageStruct storage,uint256)",
	"2da0706e": "vote(TellorStorage.TellorStorageStruct storage,uint256,bool)",
}

// TellorDisputeBin is the compiled bytecode used for deploying new contracts.
var TellorDisputeBin = "0x612290610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100875760003560e01c806397f5f9721161006557806397f5f9721461012f578063ca9a4ea51461015f578063d7b651c11461019b578063def6fac7146101c557610087565b806322048ecf1461008c5780632da0706e146100be578063694bf49f146100f6575b600080fd5b81801561009857600080fd5b506100bc600480360360408110156100af57600080fd5b50803590602001356101f5565b005b8180156100ca57600080fd5b506100bc600480360360608110156100e157600080fd5b508035906020810135906040013515156103a1565b81801561010257600080fd5b506100bc6004803603604081101561011957600080fd5b50803590602001356001600160a01b0316610651565b81801561013b57600080fd5b506100bc6004803603604081101561015257600080fd5b5080359060200135610b35565b81801561016b57600080fd5b506100bc6004803603608081101561018257600080fd5b508035906020810135906040810135906060013561133e565b8180156101a757600080fd5b506100bc600480360360208110156101be57600080fd5b5035611dbc565b8180156101d157600080fd5b506100bc600480360360408110156101e857600080fd5b5080359060200135611ead565b6000818152604483016020818152604080842054808552604a870183528185205480865284845282862083516c64697370757465526f756e647360981b8152845190819003600d01812088526005909101808652848820548287015284518083038701815291850185528151918601919091208752845282862054808752949093529320600281015491929160ff6101009091041615156001146102d5576040805162461bcd60e51b8152602060048201526012602482015271766f7465206e6565647320746f207061737360701b604482015290519081900360640190fd5b604080516874616c6c794461746560b81b815281519081900360090190206000908152600583016020522054620151804291909103116103465760405162461bcd60e51b81526004018080602001828103825260338152602001806122096033913960400191505060405180910390fd5b60040154604080516d1d195b1b1bdc90dbdb9d1c9858dd60921b8152815190819003600e0190206000908152603f90970160205290952080546001600160a01b0319166001600160a01b039096169590951790945550505050565b600082815260448085016020908152604080842081516a313637b1b5a73ab6b132b960a91b8152825190819003600b018120865260058201845282862054633f48b1ff60e01b8252600482018a905233602483015294810194909452905190939273__$e6d6eab054cf6cc8ad21596dd9777aec01$__92633f48b1ff92606480840193829003018186803b15801561043857600080fd5b505af415801561044c573d6000803e3d6000fd5b505050506040513d602081101561046257600080fd5b505133600090815260068401602052604090205490915060ff161515600114156104d3576040805162461bcd60e51b815260206004820152601860248201527f53656e6465722068617320616c726561647920766f7465640000000000000000604482015290519081900360640190fd5b80610519576040805162461bcd60e51b81526020600482015260116024820152700557365722062616c616e6365206973203607c1b604482015290519081900360640190fd5b33600090815260478601602052604090205460031415610579576040805162461bcd60e51b81526020600482015260166024820152754d696e657220697320756e646572206469737075746560501b604482015290519081900360640190fd5b3360009081526006830160209081526040808320805460ff1916600190811790915581516c6e756d6265724f66566f74657360981b8152825190819003600d0190208452600586019092529091208054909101905582156105f35760018201546105e9908263ffffffff61216916565b600183015561060e565b6001820154610608908263ffffffff61219a16565b60018301555b60408051841515815290518291339187917f911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e919081900360200190a45050505050565b604080516001600160a01b0383166020808301919091528251808303820181528284018085528151919092012063c7bb46ad60e01b9091526044820185905233606483015230608483015268056bc75e2d6310000060a4830152915173__$e6d6eab054cf6cc8ad21596dd9777aec01$__9163c7bb46ad9160c4808301926000929190829003018186803b1580156106e857600080fd5b505af41580156106fc573d6000803e3d6000fd5b5050604080516b191a5cdc1d5d1950dbdd5b9d60a21b808252825191829003600c90810183206000908152848a016020818152868320805460010190559385528551948590039092019093208352815282822054868352604a8901909152919020549092501590506107b2576000828152604a850160209081526040808320548484526044880183528184208251651bdc9a59d25160d21b815283519081900360060190208552600501909252909120556107c6565b6000828152604a8501602052604090208190555b6000828152604a850160209081526040808320548084526044880180845282852083516c64697370757465526f756e647360981b808252855191829003600d9081018320895260059093018088528689208054600101905585895284885290825285519182900390920181208752818652848720548488529286528086018390528451808203870181529085018552805190860120865290935292208390558282146109d557600082815260448701602081815260408084208151600019870181850152825180820385018152818401808552815191860191909120875260059283018552838720548088529585528387206f6d696e457865637574696f6e4461746560801b90915283519182900360500190912086520190915290912054421015610933576040805162461bcd60e51b81526020600482015260176024820152762234b9b83aba329034b99030b63932b0b23c9037b832b760491b604482015290519081900360640190fd5b600081815260448801602052604090206002015460ff16156109d3576000818152604488016020908152604080832081516874616c6c794461746560b81b81528251908190036009019020845260050190915290205462015180429190910311156109d3576040805162461bcd60e51b815260206004820152601f602482015260008051602061223c833981519152604482015290519081900360640190fd5b505b505060408051610100808201835293815260006020808301828152838501838152606085018481526001608087018181523360a0890181815260c08a019182526001600160a01b039d8e1660e08b019081528c8a526044909f018089528b8a209a518b559651938a019390935593516002890180549451925193518e166301000000026301000000600160b81b0319941515620100000262ff000019941515909e0261ff001993151560ff19909716969096179290921694909417919091169a909a17169890981790975595516003840180549189166001600160a01b031992831617905597516004840180549190981698169790971790955581516a313637b1b5a73ab6b132b960a91b8152825190819003600b018120865260059091018087528286204390559285529285526f6d696e457865637574696f6e4461746560801b835280519283900360100190922083529092522062093a8042019055565b60008181526044830160208181526040808420548452604a860182528084205480855292825280842081516c64697370757465526f756e647360981b8152825190819003600d018120865260059091018084528286205482850152825180830385018152918301835281519184019190912085529091529091205480610bb85750805b6000828152604485016020908152604080832084845281842082516c64697370757465526f756e647360981b8152835190819003600d01902085526005820190935292205480610c06575060015b60408051631c185a5960e21b815281519081900360040190206000908152600585016020529081205415610c74576040805162461bcd60e51b815260206004820152601060248201526f185b1c9958591e481c185a59081bdd5d60821b604482015290519081900360640190fd5b604080516874616c6c794461746560b81b81528151908190036009019020600090815260058501602052205462015180429190910311610ce9576040805162461bcd60e51b815260206004820152601f602482015260008051602061223c833981519152604482015290519081900360640190fd5b600284810154630100000090046001600160a01b0316600090815260478a01602090815260408083208151631c185a5960e21b8152825160049181900391909101902084526005890190925290912060019081905591850154909161010090910460ff1615151415610fd55762015180420642036001820155604080516a1cdd185ad95c90dbdd5b9d60aa1b8152815190819003600b0190206000908152818b01602052208054600019019055610d9f89611dbc565b805460041415610e74576005815560028501546003860154604080516a1cdd185ad9505b5bdd5b9d60aa1b8152815190819003600b0181206000908152828e016020528281205463c7bb46ad60e01b8352600483018f90526001600160a01b0363010000009096048616602484015293909416604482015260648101929092525173__$e6d6eab054cf6cc8ad21596dd9777aec01$__9263c7bb46ad9260848082019391829003018186803b158015610e5757600080fd5b505af4158015610e6b573d6000803e3d6000fd5b50506000835550505b60005b83811015610fcf576040805182860360208083019190915282518083038201815291830183528151918101919091206000908152600589019091522054925082610ebf578792505b60008a6044016000858152602001908152602001600020905073__$e6d6eab054cf6cc8ad21596dd9777aec01$__63c7bb46ad8c308460030160009054906101000a90046001600160a01b031685600501600060405180806266656560e81b815250600301905060405180910390208152602001908152602001600020546040518563ffffffff1660e01b815260040180858152602001846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b0316815260200182815260200194505050505060006040518083038186803b158015610faa57600080fd5b505af4158015610fbe573d6000803e3d6000fd5b505060019093019250610e77915050565b50611292565b6001815560408051681c995c5d595cdd125960ba1b815281519081900360099081018220600090815260058901602081815285832054835260488f018152858320681b5a5b995c94db1bdd60ba1b8652865195869003909401909420825290925291902054600214156110a157604080516476616c756560d81b815281519081900360059081018220600090815290890160208181528483205468074696d657374616d760bc1b8552855194859003600901909420835290815283822054825260068501905291909120555b6040805168074696d657374616d760bc1b815281519081900360090190206000908152600588016020908152828220548252600784019052205460ff1615156001141561112a576040805168074696d657374616d760bc1b81528151908190036009019020600090815260058801602090815282822054825260078401905220805460ff191690555b60005b8481101561128f57604080518287036020808301919091528251808303820181529183018352815191810191909120600090815260058a0190915220549350831561118557600084815260448c016020526040902095505b73__$e6d6eab054cf6cc8ad21596dd9777aec01$__63c7bb46ad8c308960020160039054906101000a90046001600160a01b03168f60440160008a8152602001908152602001600020600501600060405180806266656560e81b815250600301905060405180910390208152602001908152602001600020546040518563ffffffff1660e01b815260040180858152602001846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b0316815260200182815260200194505050505060006040518083038186803b15801561126b57600080fd5b505af415801561127f573d6000803e3d6000fd5b50506001909201915061112d9050565b50505b60408051681b5a5b995c94db1bdd60ba1b815281519081900360090190206000908152600587016020522054600214156113335760408051681c995c5d595cdd125960ba1b81528151908190036009018120600090815260058801602090815283822054825260488d0181528382206b191a5cdc1d5d1950dbdd5b9d60a21b8452845193849003600c01909320825260049092019091522080546000190190555b505050505050505050565b6000838152604885016020908152604080832085845260058101909252909120546113a3576040805162461bcd60e51b815260206004820152601060248201526f04d696e656420626c6f636b20697320360841b604482015290519081900360640190fd5b600582106113ef576040805162461bcd60e51b81526020600482015260146024820152734d696e657220696e6465782069732077726f6e6760601b604482015290519081900360640190fd5b60008381526008820160205260408120836005811061140a57fe5b0154604080516bffffffffffffffffffffffff19606084901b1660208083019190915260348201899052605480830189905283518084039091018152607483018085528151918301919091206b191a5cdc1d5d1950dbdd5b9d60a21b91829052845193849003608001842060009081528c860180855286822054938652865195869003600c01909520815293835284842060019290920191829055808452604a8c0190925292909120546001600160a01b03909316935091801561150657600082815260448a01602090815260408083208151651bdc9a59d25160d21b815282519081900360060190208452600501909152902081905561151c565b506000828152604a890160205260409020819055805b600081815260448a016020818152604080842081516c64697370757465526f756e647360981b808252835191829003600d908101832088526005909301808652848820548989528787529183528451928390039093018220875282855283872060019091019081905587875294845280840185905282518082038501815290830183528051908401208552909152909120839055819083821461172257600082815260448c01602081815260408084208151600019870181850152825180820385018152818401808552815191860191909120875260059283018552838720548088529585528387206f6d696e457865637574696f6e4461746560801b90915283519182900360500190912086520190915290912054421015611680576040805162461bcd60e51b81526020600482015260176024820152762234b9b83aba329034b99030b63932b0b23c9037b832b760491b604482015290519081900360640190fd5b600081815260448d01602052604090206002015460ff161561172057600081815260448d016020908152604080832081516874616c6c794461746560b81b8152825190819003600901902084526005019091529020546201518042919091031115611720576040805162461bcd60e51b815260206004820152601f602482015260008051602061223c833981519152604482015290519081900360640190fd5b505b6000886002141561185e578b60480160008c8152602001908152602001600020600401600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c01905060405180910390208152602001908152602001600020546001018c60480160008d8152602001908152602001600020600401600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c01905060405180910390208152602001908152602001600020819055508b60480160008c8152602001908152602001600020600401600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c01905060405180910390208152602001908152602001600020548c604001600060405180806a1cdd185ad9505b5bdd5b9d60aa1b815250600b019050604051809103902081526020019081526020016000205402905061188e565b5060408051696469737075746546656560b01b8152815190819003600a0190206000908152818d01602052205481025b60405180610100016040528087815260200160008152602001600015158152602001600015158152602001600015158152602001886001600160a01b03168152602001336001600160a01b0316815260200160006001600160a01b03168152508c6044016000878152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548160ff02191690831515021790555060608201518160020160016101000a81548160ff02191690831515021790555060808201518160020160026101000a81548160ff02191690831515021790555060a08201518160020160036101000a8154816001600160a01b0302191690836001600160a01b0316021790555060c08201518160030160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060e08201518160040160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050508a8c604401600087815260200190815260200160002060050160006040518080681c995c5d595cdd125960ba1b81525060090190506040518091039020815260200190815260200160002081905550898c60440160008781526020019081526020016000206005016000604051808068074696d657374616d760bc1b815250600901905060405180910390208152602001908152602001600020819055508760090160008b81526020019081526020016000208960058110611abf57fe5b01548c6044016000878152602001908152602001600020600501600060405180806476616c756560d81b81525060050190506040518091039020815260200190815260200160002081905550816202a3000242018c6044016000878152602001908152602001600020600501600060405180806f6d696e457865637574696f6e4461746560801b81525060100190506040518091039020815260200190815260200160002081905550438c6044016000878152602001908152602001600020600501600060405180806a313637b1b5a73ab6b132b960a91b815250600b0190506040518091039020815260200190815260200160002081905550888c604401600087815260200190815260200160002060050160006040518080681b5a5b995c94db1bdd60ba1b81525060090190506040518091039020815260200190815260200160002081905550808c6044016000878152602001908152602001600020600501600060405180806266656560e81b8152506003019050604051809103902081526020019081526020016000208190555073__$e6d6eab054cf6cc8ad21596dd9777aec01$__63c7bb46ad8d3330856040518563ffffffff1660e01b815260040180858152602001846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b0316815260200182815260200194505050505060006040518083038186803b158015611cdc57600080fd5b505af4158015611cf0573d6000803e3d6000fd5b505050508860021415611d265760008a81526007890160209081526040808320805460ff1916600117905560068b019091528120555b6001600160a01b038716600090815260478d016020526040902054600414611d67576001600160a01b038716600090815260478d0160205260409020600390555b604080518b81526001600160a01b038916602082015281518d9288927feceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64929081900390910190a3505050505050505050505050565b604080516a1cdd185ad9505b5bdd5b9d60aa1b81528151600b918190038201812060009081528484016020818152858320546b7461726765744d696e65727360a01b85528651600c958190039590950185208452828252868420546a1cdd185ad95c90dbdd5b9d60aa1b86528751958690039096019094208352529290922054611e789067d02ab486cedc0000906103e8908490611e5b9082906121c0565b6103e802860281611e6857fe5b0481611e7057fe5b0484036121d8565b60408051696469737075746546656560b01b8152815190819003600a0190206000908152948101602052909320929092555050565b60008181526044830160205260409020600281015460ff1615611f015760405162461bcd60e51b81526004018080602001828103825260218152602001806121e86021913960400191505060405180910390fd5b604080516f6d696e457865637574696f6e4461746560801b815281519081900360100190206000908152600583016020522054421015611f76576040805162461bcd60e51b815260206004820152601f602482015260008051602061223c833981519152604482015290519081900360640190fd5b60038101546001600160a01b0316611fd5576040805162461bcd60e51b815260206004820152601c60248201527f7265706f7274696e672050617274792069732061646472657373203000000000604482015290519081900360640190fd5b60018101546000811315611ff55760028201805461ff0019166101001790555b600282015462010000900460ff16612041576002820154630100000090046001600160a01b03166000908152604785016020526040902080546003141561203b57600481555b506120be565b604080516b746f74616c5f737570706c7960a01b8152815190819003600c01902060009081528186016020522054606490600a020481106120be576004820154604080516001600160a01b039092168252517fc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d9181900360200190a15b604080516874616c6c794461746560b81b8152815190819003600901812060009081526005850160209081529083902042905560028501805460ff19166001179081905560038601548584526001600160a01b0390811692840192909252610100810460ff1615158385015292516301000000909304169185917f21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739919081900360600190a350505050565b60008082131561218657508181018281121561218157fe5b612194565b508181018281131561219457fe5b92915050565b6000808213156121b257508082038281131561218157fe5b508082038281121561219457fe5b60008183106121cf57816121d1565b825b9392505050565b60008183116121cf57816121d156fe4469737075746520686173206265656e20616c726561647920657865637574656454696d6520666f7220766f74696e6720666f72206675727468657220646973707574657320686173206e6f742070617373656454696d6520666f7220766f74696e6720686176656e277420656c617073656400a265627a7a723158201e27429fe9a1cfaf73af74e3708dd424b8b534c564b5f48ce2aa12240c428ec164736f6c63430005100032"

// DeployTellorDispute deploys a new Ethereum contract, binding an instance of TellorDispute to it.
func DeployTellorDispute(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorDispute, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorDisputeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	tellorTransferAddr, _, _, _ := DeployTellorTransfer(auth, backend)
	TellorDisputeBin = strings.Replace(TellorDisputeBin, "__$e6d6eab054cf6cc8ad21596dd9777aec01$__", tellorTransferAddr.String()[2:], -1)

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
func (_TellorDispute *TellorDisputeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_TellorDispute *TellorDisputeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
// Solidity: event DisputeVoteTallied(uint256 indexed _disputeID, int256 _result, address indexed _reportedMiner, address _reportingParty, bool _active)
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
// Solidity: event DisputeVoteTallied(uint256 indexed _disputeID, int256 _result, address indexed _reportedMiner, address _reportingParty, bool _active)
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

// ParseDisputeVoteTallied is a log parse operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: event DisputeVoteTallied(uint256 indexed _disputeID, int256 _result, address indexed _reportedMiner, address _reportingParty, bool _active)
func (_TellorDispute *TellorDisputeFilterer) ParseDisputeVoteTallied(log types.Log) (*TellorDisputeDisputeVoteTallied, error) {
	event := new(TellorDisputeDisputeVoteTallied)
	if err := _TellorDispute.contract.UnpackLog(event, "DisputeVoteTallied", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
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
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
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

// ParseNewDispute is a log parse operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_TellorDispute *TellorDisputeFilterer) ParseNewDispute(log types.Log) (*TellorDisputeNewDispute, error) {
	event := new(TellorDisputeNewDispute)
	if err := _TellorDispute.contract.UnpackLog(event, "NewDispute", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event NewTellorAddress(address _newTellor)
func (_TellorDispute *TellorDisputeFilterer) FilterNewTellorAddress(opts *bind.FilterOpts) (*TellorDisputeNewTellorAddressIterator, error) {

	logs, sub, err := _TellorDispute.contract.FilterLogs(opts, "NewTellorAddress")
	if err != nil {
		return nil, err
	}
	return &TellorDisputeNewTellorAddressIterator{contract: _TellorDispute.contract, event: "NewTellorAddress", logs: logs, sub: sub}, nil
}

// WatchNewTellorAddress is a free log subscription operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: event NewTellorAddress(address _newTellor)
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

// ParseNewTellorAddress is a log parse operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: event NewTellorAddress(address _newTellor)
func (_TellorDispute *TellorDisputeFilterer) ParseNewTellorAddress(log types.Log) (*TellorDisputeNewTellorAddress, error) {
	event := new(TellorDisputeNewTellorAddress)
	if err := _TellorDispute.contract.UnpackLog(event, "NewTellorAddress", log); err != nil {
		return nil, err
	}
	return event, nil
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
	DisputeID  *big.Int
	Position   bool
	Voter      common.Address
	VoteWeight *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter, uint256 indexed _voteWeight)
func (_TellorDispute *TellorDisputeFilterer) FilterVoted(opts *bind.FilterOpts, _disputeID []*big.Int, _voter []common.Address, _voteWeight []*big.Int) (*TellorDisputeVotedIterator, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}
	var _voteWeightRule []interface{}
	for _, _voteWeightItem := range _voteWeight {
		_voteWeightRule = append(_voteWeightRule, _voteWeightItem)
	}

	logs, sub, err := _TellorDispute.contract.FilterLogs(opts, "Voted", _disputeIDRule, _voterRule, _voteWeightRule)
	if err != nil {
		return nil, err
	}
	return &TellorDisputeVotedIterator{contract: _TellorDispute.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter, uint256 indexed _voteWeight)
func (_TellorDispute *TellorDisputeFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *TellorDisputeVoted, _disputeID []*big.Int, _voter []common.Address, _voteWeight []*big.Int) (event.Subscription, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}
	var _voteWeightRule []interface{}
	for _, _voteWeightItem := range _voteWeight {
		_voteWeightRule = append(_voteWeightRule, _voteWeightItem)
	}

	logs, sub, err := _TellorDispute.contract.WatchLogs(opts, "Voted", _disputeIDRule, _voterRule, _voteWeightRule)
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

// ParseVoted is a log parse operation binding the contract event 0x911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter, uint256 indexed _voteWeight)
func (_TellorDispute *TellorDisputeFilterer) ParseVoted(log types.Log) (*TellorDisputeVoted, error) {
	event := new(TellorDisputeVoted)
	if err := _TellorDispute.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TellorGettersABI is the input ABI used to generate the binding from.
const TellorGettersABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[9]\",\"name\":\"\",\"type\":\"uint256[9]\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastNewValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"internalType\":\"address[5]\",\"name\":\"\",\"type\":\"address[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_request\",\"type\":\"bytes32\"}],\"name\":\"getRequestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"\",\"type\":\"uint256[51]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"\",\"type\":\"uint256[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TellorGettersFuncSigs maps the 4-byte function signature to its string representation.
var TellorGettersFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"999cf26c": "allowedToTrade(address,uint256)",
	"70a08231": "balanceOf(address)",
	"4ee2cd7e": "balanceOfAt(address,uint256)",
	"63bb82ad": "didMine(bytes32,address)",
	"a7c438bc": "didVote(uint256,address)",
	"133bee5e": "getAddressVars(bytes32)",
	"af0b1327": "getAllDisputeVars(uint256)",
	"a22e407a": "getCurrentVariables()",
	"da379941": "getDisputeIdByDisputeHash(bytes32)",
	"7f6fd5d9": "getDisputeUintVars(uint256,bytes32)",
	"fc7cf0a0": "getLastNewValue()",
	"3180f8df": "getLastNewValueById(uint256)",
	"c775b542": "getMinedBlockNum(uint256,uint256)",
	"69026d63": "getMinersByRequestIdAndTimestamp(uint256,uint256)",
	"46eee1c4": "getNewValueCountbyRequestId(uint256)",
	"1db842f0": "getRequestIdByQueryHash(bytes32)",
	"6173c0b8": "getRequestIdByRequestQIndex(uint256)",
	"0f0b424d": "getRequestIdByTimestamp(uint256)",
	"b5413029": "getRequestQ()",
	"e0ae93c1": "getRequestUintVars(uint256,bytes32)",
	"e1eee6d6": "getRequestVars(uint256)",
	"733bdef0": "getStakerInfo(address)",
	"11c98512": "getSubmissionsByTimestamp(uint256,uint256)",
	"77fbb663": "getTimestampbyRequestIDandIndex(uint256,uint256)",
	"612c8f7f": "getUintVar(bytes32)",
	"19e8e03b": "getVariablesOnDeck()",
	"3df0777b": "isInDispute(uint256,uint256)",
	"93fa4915": "retrieveData(uint256,uint256)",
	"18160ddd": "totalSupply()",
}

// TellorGettersBin is the compiled bytecode used for deploying new contracts.
var TellorGettersBin = "0x608060405234801561001057600080fd5b506119be806100206000396000f3fe608060405234801561001057600080fd5b50600436106101da5760003560e01c8063733bdef011610104578063af0b1327116100a2578063dd62ed3e11610071578063dd62ed3e1461075f578063e0ae93c11461078d578063e1eee6d6146107b0578063fc7cf0a0146108c7576101da565b8063af0b13271461065d578063b541302914610701578063c775b5421461071f578063da37994114610742576101da565b806393fa4915116100de57806393fa49151461053f578063999cf26c14610562578063a22e407a1461058e578063a7c438bc14610631576101da565b8063733bdef0146104ba57806377fbb663146104f95780637f6fd5d91461051c576101da565b80633df0777b1161017c5780636173c0b81161014b5780636173c0b81461042857806363bb82ad1461044557806369026d631461047157806370a0823114610494576101da565b80633df0777b1461038b57806346eee1c4146103c25780634ee2cd7e146103df578063612c8f7f1461040b576101da565b806318160ddd116101b857806318160ddd146102a257806319e8e03b146102aa5780631db842f0146103385780633180f8df14610355576101da565b80630f0b424d146101df57806311c985121461020e578063133bee5e14610269575b600080fd5b6101fc600480360360208110156101f557600080fd5b50356108cf565b60408051918252519081900360200190f35b6102316004803603604081101561022457600080fd5b50803590602001356108e7565b604051808260a080838360005b8381101561025657818101518382015260200161023e565b5050505090500191505060405180910390f35b6102866004803603602081101561027f57600080fd5b5035610908565b604080516001600160a01b039092168252519081900360200190f35b6101fc61091a565b6102b261092b565b6040518084815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156102fb5781810151838201526020016102e3565b50505050905090810190601f1680156103285780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b6101fc6004803603602081101561034e57600080fd5b5035610945565b6103726004803603602081101561036b57600080fd5b5035610957565b6040805192835290151560208301528051918290030190f35b6103ae600480360360408110156103a157600080fd5b5080359060200135610973565b604080519115158252519081900360200190f35b6101fc600480360360208110156103d857600080fd5b5035610986565b6101fc600480360360408110156103f557600080fd5b506001600160a01b038135169060200135610998565b6101fc6004803603602081101561042157600080fd5b5035610a35565b6101fc6004803603602081101561043e57600080fd5b5035610a47565b6103ae6004803603604081101561045b57600080fd5b50803590602001356001600160a01b0316610a59565b6102316004803603604081101561048757600080fd5b5080359060200135610ac3565b6101fc600480360360208110156104aa57600080fd5b50356001600160a01b0316610add565b6104e0600480360360208110156104d057600080fd5b50356001600160a01b0316610b72565b6040805192835260208301919091528051918290030190f35b6101fc6004803603604081101561050f57600080fd5b5080359060200135610b85565b6101fc6004803603604081101561053257600080fd5b5080359060200135610b98565b6101fc6004803603604081101561055557600080fd5b5080359060200135610bab565b6103ae6004803603604081101561057857600080fd5b506001600160a01b038135169060200135610bbe565b610596610c28565b6040518087815260200186815260200185815260200180602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b838110156105f15781810151838201526020016105d9565b50505050905090810190601f16801561061e5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b6103ae6004803603604081101561064757600080fd5b50803590602001356001600160a01b0316610c4f565b61067a6004803603602081101561067357600080fd5b5035610c62565b604080518a815289151560208201528815159181019190915286151560608201526001600160a01b03808716608083015285811660a0830152841660c082015260e081018361012080838360005b838110156106e05781810151838201526020016106c8565b50505050905001828152602001995050505050505050505060405180910390f35b610709610ca6565b604051815181528082610660808383602061023e565b6101fc6004803603604081101561073557600080fd5b5080359060200135610cb8565b6101fc6004803603602081101561075857600080fd5b5035610ccb565b6101fc6004803603604081101561077557600080fd5b506001600160a01b0381358116916020013516610cdd565b6101fc600480360360408110156107a357600080fd5b5080359060200135610d48565b6107cd600480360360208110156107c657600080fd5b5035610d5b565b604051808060200180602001878152602001868152602001858152602001848152602001838103835289818151815260200191508051906020019080838360005b8381101561082657818101518382015260200161080e565b50505050905090810190601f1680156108535780820380516001836020036101000a031916815260200191505b5083810382528851815288516020918201918a019080838360005b8381101561088657818101518382015260200161086e565b50505050905090810190601f1680156108b35780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b610372610d87565b60006108e1818363ffffffff610d9c16565b92915050565b6108ef61192d565b6109016000848463ffffffff610db216565b9392505050565b60006108e1818363ffffffff610e0c16565b60006109266000610e2b565b905090565b600080606061093a6000610e5e565b925092509250909192565b60006108e1818363ffffffff610f4a16565b60008061096a818463ffffffff610f6016565b91509150915091565b600061090181848463ffffffff610fc616565b60006108e1818363ffffffff610fed16565b60408051633f48b1ff60e01b81526000600482018190526001600160a01b038516602483015260448201849052915173__$e6d6eab054cf6cc8ad21596dd9777aec01$__91633f48b1ff916064808301926020929190829003018186803b158015610a0257600080fd5b505af4158015610a16573d6000803e3d6000fd5b505050506040513d6020811015610a2c57600080fd5b50519392505050565b60006108e1818363ffffffff61100616565b60006108e1818363ffffffff61101816565b6040805163c93299e960e01b8152600060048201819052602482018590526001600160a01b0384166044830152915173__$84fe56546dd3e4270ce7af1b84119b8097$__9163c93299e9916064808301926020929190829003018186803b158015610a0257600080fd5b610acb61192d565b6109016000848463ffffffff61108716565b604080516393b182b360e01b81526000600482018190526001600160a01b0384166024830152915173__$e6d6eab054cf6cc8ad21596dd9777aec01$__916393b182b3916044808301926020929190829003018186803b158015610b4057600080fd5b505af4158015610b54573d6000803e3d6000fd5b505050506040513d6020811015610b6a57600080fd5b505192915050565b60008061096a818463ffffffff6110eb16565b600061090181848463ffffffff61111216565b600061090181848463ffffffff61114516565b600061090181848463ffffffff61116916565b604080516356555cf160e11b81526000600482018190526001600160a01b038516602483015260448201849052915173__$e6d6eab054cf6cc8ad21596dd9777aec01$__9163acaab9e2916064808301926020929190829003018186803b158015610a0257600080fd5b60008060006060600080610c3c600061118d565b949b939a50919850965094509092509050565b600061090181848463ffffffff61134b16565b6000806000806000806000610c7561194b565b6000610c87818b63ffffffff61137d16565b9850985098509850985098509850985098509193959799909294969850565b610cae61196a565b6109266000611592565b600061090181848463ffffffff6115d216565b60006108e1818363ffffffff6115f616565b60408051632fcc801b60e21b81526000600482018190526001600160a01b03808616602484015284166044830152915173__$e6d6eab054cf6cc8ad21596dd9777aec01$__9163bf32006c916064808301926020929190829003018186803b158015610a0257600080fd5b600061090181848463ffffffff61160c16565b6060806000808080610d73818863ffffffff61163016565b949c939b5091995097509550909350915050565b600080610d94600061180b565b915091509091565b6000908152604291909101602052604090205490565b610dba61192d565b6000838152604885016020908152604080832085845260090190915290819020815160a08101928390529160059082845b815481526020019060010190808311610deb57505050505090509392505050565b6000908152603f9190910160205260409020546001600160a01b031690565b604080516b746f74616c5f737570706c7960a01b8152815190819003600c01902060009081528183016020522054919050565b60008060606000610e6e85611881565b600081815260488701602081815260408084208151670746f74616c5469760c41b8152825160089181900391909101812086526004820184528286205495879052938352805460026001821615610100026000190190911604601f810184900484028501840190925281845294955085949291839190830182828015610f355780601f10610f0a57610100808354040283529160200191610f35565b820191906000526020600020905b815481529060010190602001808311610f1857829003601f168201915b50505050509050935093509350509193909250565b6000908152604991909101602052604090205490565b60008181526048830160205260408120600381015482919015610fb657600381018054610faa9187918791906000198101908110610f9a57fe5b9060005260206000200154611169565b60019250925050610fbf565b50600091508190505b9250929050565b60009182526048929092016020908152604080832093835260079093019052205460ff1690565b6000908152604891909101602052604090206003015490565b60009081526040918201602052205490565b60006032821115611070576040805162461bcd60e51b815260206004820152601a60248201527f526571756573745120696e6465782069732061626f7665203530000000000000604482015290519081900360640190fd5b506000908152604391909101602052604090205490565b61108f61192d565b6000838152604885016020908152604080832085845260080190915290819020815160a08101928390529160059082845b81546001600160a01b031681526001909101906020018083116110c057505050505090509392505050565b6001600160a01b031660009081526047919091016020526040902080546001909101549091565b6000828152604884016020526040812060030180548390811061113157fe5b906000526020600020015490509392505050565b60009182526044929092016020908152604080832093835260059093019052205490565b60009182526048929092016020908152604080832093835260069093019052205490565b8054604080516f18dd5c9c995b9d14995c5d595cdd125960821b80825282519182900360109081018320600090815284870160208181528683205469646966666963756c747960b01b8752875196879003600a01872084528282528784205486885288519788900386018820855283835288852054855260488b01808452898620888a528a51998a900388018a2087528585528a87205487528185528a87206a6772616e756c617269747960a81b8b528b519a8b9003600b018b208852600490810186528b882054998b528b519a8b90039098018a2087529484528986205486528352888520670746f74616c5469760c41b8952895198899003600801892086529095018252878420548354601f600260001961010060018516150201909216919091049081018490048402890184019099528888529398899889986060988a988998949795969095949193919291859183018282801561132f5780601f106113045761010080835404028352916020019161132f565b820191906000526020600020905b81548152906001019060200180831161131257829003601f168201915b5050505050925095509550955095509550955091939550919395565b600082815260448401602090815260408083206001600160a01b038516845260060190915290205460ff169392505050565b600080600080600080600061139061194b565b5050506000868152604488016020908152604080832080546002820154600380840154600485015486516101208101808952681c995c5d595cdd125960ba1b905287518082036101290190208a526005808801808b52898c20548352895168074696d657374616d760bc1b81528a519081900360099081019091208d52818c528a8d2054848d01528a516476616c756560d81b81528b51908190039093019092208c52808b52898c2054838b015289516f6d696e457865637574696f6e4461746560801b81528a519081900360100190208c52808b52898c2054606084015289516c6e756d6265724f66566f74657360981b81528a5190819003600d0190208c52808b52898c2054608084015289516a313637b1b5a73ab6b132b960a91b81528a5190819003600b0190208c52808b52898c205460a08401528951681b5a5b995c94db1bdd60ba1b81528a51908190039092019091208b52808a52888b205460c083015288516571756f72756d60d01b815289519081900360060190208b52808a52888b205460e083015288516266656560e81b81528951908190039095019094208a5292909752949096205461010087810191909152600190930154919a5060ff8082169a509281048316985062010000810490921696506001600160a01b0363010000009092048216955091811693921691909295985092959850929598565b61159a61196a565b6040805161066081019182905290600184019060339082845b8154815260200190600101908083116115b35750505050509050919050565b60009182526048929092016020908152604080832093835260059093019052205490565b6000908152604a91909101602052604090205490565b60009182526048929092016020908152604080832093835260049093019052205490565b6000818152604883016020908152604080832060028082015483516a6772616e756c617269747960a81b8152845190819003600b018120875260048401808752858820546f3932b8bab2b9ba28a837b9b4ba34b7b760811b83528651928390036010018320895281885286892054670746f74616c5469760c41b845287519384900360080184208a52918852868920548654601f6000196101006001848116159190910291909101909216979097049687018a90048a0285018a019098528584526060998a9990988998899889989197889793880196929592949091889183018282801561175f5780601f106117345761010080835404028352916020019161175f565b820191906000526020600020905b81548152906001019060200180831161174257829003601f168201915b5050885460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959b508a9450925084019050828280156117ed5780601f106117c2576101008083540402835291602001916117ed565b820191906000526020600020905b8154815290600101906020018083116117d057829003601f168201915b50505050509450965096509650965096509650509295509295509295565b604080517174696d654f664c6173744e657756616c756560701b8082528251601292819003830181206000908152858501602081815286832054835260428801815286832054948452865193849003909501909220815292529181205490918291611877918591611169565b9360019350915050565b60408051610660810191829052600091829182916118c29190600187019060339082845b8154815260200190600101908083116118a55750505050506118dd565b60009081526043909501602052505060409092205492915050565b6020810151600160025b603381101561192757828482603381106118fd57fe5b6020020151111561191f5783816033811061191457fe5b602002015192508091505b6001016118e7565b50915091565b6040518060a001604052806005906020820280388339509192915050565b6040518061012001604052806009906020820280388339509192915050565b604051806106600160405280603390602082028038833950919291505056fea265627a7a72315820039ce6fb28ad6f748a493cec7972f4945b9bacca38552b106594a6c055c07d1164736f6c63430005100032"

// DeployTellorGetters deploys a new Ethereum contract, binding an instance of TellorGetters to it.
func DeployTellorGetters(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorGetters, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorGettersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	tellorGettersLibraryAddr, _, _, _ := DeployTellorGettersLibrary(auth, backend)
	TellorGettersBin = strings.Replace(TellorGettersBin, "__$84fe56546dd3e4270ce7af1b84119b8097$__", tellorGettersLibraryAddr.String()[2:], -1)

	tellorTransferAddr, _, _, _ := DeployTellorTransfer(auth, backend)
	TellorGettersBin = strings.Replace(TellorGettersBin, "__$e6d6eab054cf6cc8ad21596dd9777aec01$__", tellorTransferAddr.String()[2:], -1)

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
func (_TellorGetters *TellorGettersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_TellorGetters *TellorGettersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) Allowance(opts *bind.CallOpts, _user common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "allowance", _user, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_TellorGetters *TellorGettersSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _TellorGetters.Contract.Allowance(&_TellorGetters.CallOpts, _user, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _TellorGetters.Contract.Allowance(&_TellorGetters.CallOpts, _user, _spender)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_TellorGetters *TellorGettersCaller) AllowedToTrade(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "allowedToTrade", _user, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_TellorGetters *TellorGettersSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _TellorGetters.Contract.AllowedToTrade(&_TellorGetters.CallOpts, _user, _amount)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_TellorGetters *TellorGettersCallerSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _TellorGetters.Contract.AllowedToTrade(&_TellorGetters.CallOpts, _user, _amount)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "balanceOf", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_TellorGetters *TellorGettersSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _TellorGetters.Contract.BalanceOf(&_TellorGetters.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _TellorGetters.Contract.BalanceOf(&_TellorGetters.CallOpts, _user)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) BalanceOfAt(opts *bind.CallOpts, _user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "balanceOfAt", _user, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_TellorGetters *TellorGettersSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.BalanceOfAt(&_TellorGetters.CallOpts, _user, _blockNumber)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.BalanceOfAt(&_TellorGetters.CallOpts, _user, _blockNumber)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_TellorGetters *TellorGettersCaller) DidMine(opts *bind.CallOpts, _challenge [32]byte, _miner common.Address) (bool, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "didMine", _challenge, _miner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_TellorGetters *TellorGettersSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _TellorGetters.Contract.DidMine(&_TellorGetters.CallOpts, _challenge, _miner)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_TellorGetters *TellorGettersCallerSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _TellorGetters.Contract.DidMine(&_TellorGetters.CallOpts, _challenge, _miner)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_TellorGetters *TellorGettersCaller) DidVote(opts *bind.CallOpts, _disputeId *big.Int, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "didVote", _disputeId, _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_TellorGetters *TellorGettersSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _TellorGetters.Contract.DidVote(&_TellorGetters.CallOpts, _disputeId, _address)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_TellorGetters *TellorGettersCallerSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _TellorGetters.Contract.DidVote(&_TellorGetters.CallOpts, _disputeId, _address)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_TellorGetters *TellorGettersCaller) GetAddressVars(opts *bind.CallOpts, _data [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getAddressVars", _data)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_TellorGetters *TellorGettersSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _TellorGetters.Contract.GetAddressVars(&_TellorGetters.CallOpts, _data)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_TellorGetters *TellorGettersCallerSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _TellorGetters.Contract.GetAddressVars(&_TellorGetters.CallOpts, _data)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_TellorGetters *TellorGettersCaller) GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getAllDisputeVars", _disputeId)

	if err != nil {
		return *new([32]byte), *new(bool), *new(bool), *new(bool), *new(common.Address), *new(common.Address), *new(common.Address), *new([9]*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)
	out4 := *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	out5 := *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	out6 := *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	out7 := *abi.ConvertType(out[7], new([9]*big.Int)).(*[9]*big.Int)
	out8 := *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, err

}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_TellorGetters *TellorGettersSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetAllDisputeVars(&_TellorGetters.CallOpts, _disputeId)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_TellorGetters *TellorGettersCallerSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetAllDisputeVars(&_TellorGetters.CallOpts, _disputeId)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() view returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_TellorGetters *TellorGettersCaller) GetCurrentVariables(opts *bind.CallOpts) ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getCurrentVariables")

	if err != nil {
		return *new([32]byte), *new(*big.Int), *new(*big.Int), *new(string), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, err

}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() view returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_TellorGetters *TellorGettersSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetCurrentVariables(&_TellorGetters.CallOpts)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() view returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_TellorGetters *TellorGettersCallerSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetCurrentVariables(&_TellorGetters.CallOpts)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetDisputeIdByDisputeHash(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getDisputeIdByDisputeHash", _hash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetDisputeIdByDisputeHash(&_TellorGetters.CallOpts, _hash)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetDisputeIdByDisputeHash(&_TellorGetters.CallOpts, _hash)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetDisputeUintVars(opts *bind.CallOpts, _disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getDisputeUintVars", _disputeId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetDisputeUintVars(&_TellorGetters.CallOpts, _disputeId, _data)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetDisputeUintVars(&_TellorGetters.CallOpts, _disputeId, _data)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_TellorGetters *TellorGettersCaller) GetLastNewValue(opts *bind.CallOpts) (*big.Int, bool, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getLastNewValue")

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_TellorGetters *TellorGettersSession) GetLastNewValue() (*big.Int, bool, error) {
	return _TellorGetters.Contract.GetLastNewValue(&_TellorGetters.CallOpts)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_TellorGetters *TellorGettersCallerSession) GetLastNewValue() (*big.Int, bool, error) {
	return _TellorGetters.Contract.GetLastNewValue(&_TellorGetters.CallOpts)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_TellorGetters *TellorGettersCaller) GetLastNewValueById(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, bool, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getLastNewValueById", _requestId)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_TellorGetters *TellorGettersSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _TellorGetters.Contract.GetLastNewValueById(&_TellorGetters.CallOpts, _requestId)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_TellorGetters *TellorGettersCallerSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _TellorGetters.Contract.GetLastNewValueById(&_TellorGetters.CallOpts, _requestId)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetMinedBlockNum(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getMinedBlockNum", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetMinedBlockNum(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetMinedBlockNum(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_TellorGetters *TellorGettersCaller) GetMinersByRequestIdAndTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getMinersByRequestIdAndTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([5]common.Address)).(*[5]common.Address)

	return out0, err

}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_TellorGetters *TellorGettersSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _TellorGetters.Contract.GetMinersByRequestIdAndTimestamp(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_TellorGetters *TellorGettersCallerSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _TellorGetters.Contract.GetMinersByRequestIdAndTimestamp(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getNewValueCountbyRequestId", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetNewValueCountbyRequestId(&_TellorGetters.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetNewValueCountbyRequestId(&_TellorGetters.CallOpts, _requestId)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(bytes32 _request) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetRequestIdByQueryHash(opts *bind.CallOpts, _request [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getRequestIdByQueryHash", _request)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(bytes32 _request) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestIdByQueryHash(&_TellorGetters.CallOpts, _request)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(bytes32 _request) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestIdByQueryHash(&_TellorGetters.CallOpts, _request)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetRequestIdByRequestQIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getRequestIdByRequestQIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestIdByRequestQIndex(&_TellorGetters.CallOpts, _index)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestIdByRequestQIndex(&_TellorGetters.CallOpts, _index)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetRequestIdByTimestamp(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getRequestIdByTimestamp", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestIdByTimestamp(&_TellorGetters.CallOpts, _timestamp)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestIdByTimestamp(&_TellorGetters.CallOpts, _timestamp)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_TellorGetters *TellorGettersCaller) GetRequestQ(opts *bind.CallOpts) ([51]*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getRequestQ")

	if err != nil {
		return *new([51]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([51]*big.Int)).(*[51]*big.Int)

	return out0, err

}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_TellorGetters *TellorGettersSession) GetRequestQ() ([51]*big.Int, error) {
	return _TellorGetters.Contract.GetRequestQ(&_TellorGetters.CallOpts)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_TellorGetters *TellorGettersCallerSession) GetRequestQ() ([51]*big.Int, error) {
	return _TellorGetters.Contract.GetRequestQ(&_TellorGetters.CallOpts)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetRequestUintVars(opts *bind.CallOpts, _requestId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getRequestUintVars", _requestId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestUintVars(&_TellorGetters.CallOpts, _requestId, _data)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestUintVars(&_TellorGetters.CallOpts, _requestId, _data)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(string, string, bytes32, uint256, uint256, uint256)
func (_TellorGetters *TellorGettersCaller) GetRequestVars(opts *bind.CallOpts, _requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getRequestVars", _requestId)

	if err != nil {
		return *new(string), *new(string), *new([32]byte), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, err

}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(string, string, bytes32, uint256, uint256, uint256)
func (_TellorGetters *TellorGettersSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetRequestVars(&_TellorGetters.CallOpts, _requestId)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(string, string, bytes32, uint256, uint256, uint256)
func (_TellorGetters *TellorGettersCallerSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetRequestVars(&_TellorGetters.CallOpts, _requestId)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_TellorGetters *TellorGettersCaller) GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getStakerInfo", _staker)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_TellorGetters *TellorGettersSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetStakerInfo(&_TellorGetters.CallOpts, _staker)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_TellorGetters *TellorGettersCallerSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetStakerInfo(&_TellorGetters.CallOpts, _staker)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_TellorGetters *TellorGettersCaller) GetSubmissionsByTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getSubmissionsByTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_TellorGetters *TellorGettersSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _TellorGetters.Contract.GetSubmissionsByTimestamp(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_TellorGetters *TellorGettersCallerSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _TellorGetters.Contract.GetSubmissionsByTimestamp(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getTimestampbyRequestIDandIndex", _requestID, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetTimestampbyRequestIDandIndex(&_TellorGetters.CallOpts, _requestID, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetTimestampbyRequestIDandIndex(&_TellorGetters.CallOpts, _requestID, _index)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getUintVar", _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetUintVar(&_TellorGetters.CallOpts, _data)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetUintVar(&_TellorGetters.CallOpts, _data)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() view returns(uint256, uint256, string)
func (_TellorGetters *TellorGettersCaller) GetVariablesOnDeck(opts *bind.CallOpts) (*big.Int, *big.Int, string, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getVariablesOnDeck")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)

	return out0, out1, out2, err

}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() view returns(uint256, uint256, string)
func (_TellorGetters *TellorGettersSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _TellorGetters.Contract.GetVariablesOnDeck(&_TellorGetters.CallOpts)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() view returns(uint256, uint256, string)
func (_TellorGetters *TellorGettersCallerSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _TellorGetters.Contract.GetVariablesOnDeck(&_TellorGetters.CallOpts)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_TellorGetters *TellorGettersCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "isInDispute", _requestId, _timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_TellorGetters *TellorGettersSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _TellorGetters.Contract.IsInDispute(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_TellorGetters *TellorGettersCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _TellorGetters.Contract.IsInDispute(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "retrieveData", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.RetrieveData(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.RetrieveData(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TellorGetters *TellorGettersCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TellorGetters *TellorGettersSession) TotalSupply() (*big.Int, error) {
	return _TellorGetters.Contract.TotalSupply(&_TellorGetters.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) TotalSupply() (*big.Int, error) {
	return _TellorGetters.Contract.TotalSupply(&_TellorGetters.CallOpts)
}

// TellorGettersLibraryABI is the input ABI used to generate the binding from.
const TellorGettersLibraryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newTellor\",\"type\":\"address\"}],\"name\":\"NewTellorAddress\",\"type\":\"event\"}]"

// TellorGettersLibraryFuncSigs maps the 4-byte function signature to its string representation.
var TellorGettersLibraryFuncSigs = map[string]string{
	"c93299e9": "didMine(TellorStorage.TellorStorageStruct storage,bytes32,address)",
}

// TellorGettersLibraryBin is the compiled bytecode used for deploying new contracts.
var TellorGettersLibraryBin = "0x60dd610025600b82828239805160001a60731461001857fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063c93299e9146038575b600080fd5b606760048036036060811015604c57600080fd5b50803590602081013590604001356001600160a01b0316607b565b604080519115158252519081900360200190f35b6000918252604192909201602090815260408083206001600160a01b039094168352929052205460ff169056fea265627a7a72315820bc5d92dc00ab63c01f27b2d22b2367ba48b126ae97468f94393dd55a651bc4f064736f6c63430005100032"

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
func (_TellorGettersLibrary *TellorGettersLibraryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_TellorGettersLibrary *TellorGettersLibraryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
// Solidity: event NewTellorAddress(address _newTellor)
func (_TellorGettersLibrary *TellorGettersLibraryFilterer) FilterNewTellorAddress(opts *bind.FilterOpts) (*TellorGettersLibraryNewTellorAddressIterator, error) {

	logs, sub, err := _TellorGettersLibrary.contract.FilterLogs(opts, "NewTellorAddress")
	if err != nil {
		return nil, err
	}
	return &TellorGettersLibraryNewTellorAddressIterator{contract: _TellorGettersLibrary.contract, event: "NewTellorAddress", logs: logs, sub: sub}, nil
}

// WatchNewTellorAddress is a free log subscription operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: event NewTellorAddress(address _newTellor)
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

// ParseNewTellorAddress is a log parse operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: event NewTellorAddress(address _newTellor)
func (_TellorGettersLibrary *TellorGettersLibraryFilterer) ParseNewTellorAddress(log types.Log) (*TellorGettersLibraryNewTellorAddress, error) {
	event := new(TellorGettersLibraryNewTellorAddress)
	if err := _TellorGettersLibrary.contract.UnpackLog(event, "NewTellorAddress", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TellorStakeABI is the input ABI used to generate the binding from.
const TellorStakeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"NewStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"}]"

// TellorStakeFuncSigs maps the 4-byte function signature to its string representation.
var TellorStakeFuncSigs = map[string]string{
	"820a2d66": "depositStake(TellorStorage.TellorStorageStruct storage)",
	"4601f1cd": "init(TellorStorage.TellorStorageStruct storage)",
	"c9cf5e4c": "requestStakingWithdraw(TellorStorage.TellorStorageStruct storage)",
	"44bacc4b": "withdrawStake(TellorStorage.TellorStorageStruct storage)",
}

// TellorStakeBin is the compiled bytecode used for deploying new contracts.
var TellorStakeBin = "0x610a6b610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c806344bacc4b1461005b5780634601f1cd14610087578063820a2d66146100b1578063c9cf5e4c146100db575b600080fd5b81801561006757600080fd5b506100856004803603602081101561007e57600080fd5b5035610105565b005b81801561009357600080fd5b50610085600480360360208110156100aa57600080fd5b50356101e2565b8180156100bd57600080fd5b50610085600480360360208110156100d457600080fd5b5035610604565b8180156100e757600080fd5b50610085600480360360208110156100fe57600080fd5b5035610678565b3360009081526047820160205260409020600181015462093a80906201518042064203031015610171576040805162461bcd60e51b8152602060048201526012602482015271372064617973206469646e2774207061737360701b604482015290519081900360640190fd5b80546002146101b15760405162461bcd60e51b81526004018080602001828103825260238152602001806109f26023913960400191505060405180910390fd5b600080825560405133917f4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec91a25050565b6040805167646563696d616c7360c01b815281519081900360080190206000908152818301602052205415610252576040805162461bcd60e51b8152602060048201526011602482015270546f6f206d616e7920646563696d616c7360781b604482015290519081900360640190fd5b3060009081526045820160205260408082208151631d6f7b8160e31b8152600481019190915269014542ba12a337c00000196024820152905173__$e6d6eab054cf6cc8ad21596dd9777aec01$__9263eb7bdc089260448082019391829003018186803b1580156102c257600080fd5b505af41580156102d6573d6000803e3d6000fd5b505050506102e26109d3565b506040805160c08101825273e037ec8ec9ec423826750853899394de7f024fee815273cdd8fa31af8475574b8909f135d510579a8087d3602082015273b9dd5afd86547df817da2d0fb89334a6f8edd8919181019190915273230570cd052f40e14c14a81038c6f3aa685d712b6060820152733233afa02644ccd048587f8ba6e99b3c00a34dcc608082015273e010ac6e0248790e08f42d5f697160dedf97e02460a082015260005b60068110156104645773__$e6d6eab054cf6cc8ad21596dd9777aec01$__63eb7bdc088460450160008585600681106103c057fe5b60200201516001600160a01b03166001600160a01b03168152602001908152602001600020683635c9adc5dea000006040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561042a57600080fd5b505af415801561043e573d6000803e3d6000fd5b5050505061045c8383836006811061045257fe5b60200201516107ad565b60010161038b565b50604080516b746f74616c5f737570706c7960a01b8152815190819003600c908101822060009081528386016020818152858320805469014542ba12a337c0000001905567646563696d616c7360c01b855285519485900360080185208352818152858320601290556b7461726765744d696e65727360a01b85528551948590039093018420825280835284822060c890556a1cdd185ad9505b5bdd5b9d60aa1b8452845193849003600b0184208252808352848220683635c9adc5dea000009055696469737075746546656560b01b8452845193849003600a908101852083528184528583206834957444b840e800009055691d1a5b5955185c99d95d60b21b808652865195869003820186208452828552868420610258905585528551948590030190932081529190522054428161059a57fe5b604080517174696d654f664c6173744e657756616c756560701b815281519081900360120181206000908152958201602081815283882095909406420390945569646966666963756c747960b01b8152815190819003600a01902085529190529091206001905550565b61060e81336107ad565b73__$6eea42fadb496cf5107a5a057a700a86ef$__63d7b651c1826040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561065d57600080fd5b505af4158015610671573d6000803e3d6000fd5b5050505050565b336000908152604782016020526040902080546001146106d5576040805162461bcd60e51b8152602060048201526013602482015272135a5b995c881a5cc81b9bdd081cdd185ad959606a1b604482015290519081900360640190fd5b6002815562015180420642036001820155604080516a1cdd185ad95c90dbdd5b9d60aa1b8152815190819003600b0181206000908152828501602052828120805460001901905563d7b651c160e01b825260048201859052915173__$6eea42fadb496cf5107a5a057a700a86ef$__9263d7b651c19260248082019391829003018186803b15801561076657600080fd5b505af415801561077a573d6000803e3d6000fd5b50506040513392507f453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf9150600090a25050565b604080516a1cdd185ad9505b5bdd5b9d60aa1b8152815190819003600b0181206000908152828501602090815290839020546393b182b360e01b8352600483018690526001600160a01b0385166024840152925173__$e6d6eab054cf6cc8ad21596dd9777aec01$__926393b182b3926044808301939192829003018186803b15801561083957600080fd5b505af415801561084d573d6000803e3d6000fd5b505050506040513d602081101561086357600080fd5b505110156108a25760405162461bcd60e51b8152600401808060200182810382526022815260200180610a156022913960400191505060405180910390fd5b6001600160a01b038116600090815260478301602052604090205415806108e357506001600160a01b03811660009081526047830160205260409020546002145b610934576040805162461bcd60e51b815260206004820152601b60248201527f4d696e657220697320696e207468652077726f6e672073746174650000000000604482015290519081900360640190fd5b604080516a1cdd185ad95c90dbdd5b9d60aa1b8152815190819003600b01812060009081528483016020908152838220805460019081019091558385018552808452620151804290810690038285019081526001600160a01b038716808552604789019093528584209451855551930192909255915190917f46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e291a25050565b6040518060c00160405280600690602082028038833950919291505056fe4d696e657220776173206e6f74206c6f636b656420666f72207769746864726177616c42616c616e6365206973206c6f776572207468616e207374616b6520616d6f756e74a265627a7a72315820cb21fb06c722dce9f960d6113fc879c81eef076a7059c9775cb329165fbf371164736f6c63430005100032"

// DeployTellorStake deploys a new Ethereum contract, binding an instance of TellorStake to it.
func DeployTellorStake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorStake, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorStakeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	tellorDisputeAddr, _, _, _ := DeployTellorDispute(auth, backend)
	TellorStakeBin = strings.Replace(TellorStakeBin, "__$6eea42fadb496cf5107a5a057a700a86ef$__", tellorDisputeAddr.String()[2:], -1)

	tellorTransferAddr, _, _, _ := DeployTellorTransfer(auth, backend)
	TellorStakeBin = strings.Replace(TellorStakeBin, "__$e6d6eab054cf6cc8ad21596dd9777aec01$__", tellorTransferAddr.String()[2:], -1)

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
func (_TellorStake *TellorStakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_TellorStake *TellorStakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
// Solidity: event NewStake(address indexed _sender)
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
// Solidity: event NewStake(address indexed _sender)
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

// ParseNewStake is a log parse operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: event NewStake(address indexed _sender)
func (_TellorStake *TellorStakeFilterer) ParseNewStake(log types.Log) (*TellorStakeNewStake, error) {
	event := new(TellorStakeNewStake)
	if err := _TellorStake.contract.UnpackLog(event, "NewStake", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event StakeWithdrawRequested(address indexed _sender)
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
// Solidity: event StakeWithdrawRequested(address indexed _sender)
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

// ParseStakeWithdrawRequested is a log parse operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: event StakeWithdrawRequested(address indexed _sender)
func (_TellorStake *TellorStakeFilterer) ParseStakeWithdrawRequested(log types.Log) (*TellorStakeStakeWithdrawRequested, error) {
	event := new(TellorStakeStakeWithdrawRequested)
	if err := _TellorStake.contract.UnpackLog(event, "StakeWithdrawRequested", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event StakeWithdrawn(address indexed _sender)
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
// Solidity: event StakeWithdrawn(address indexed _sender)
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: event StakeWithdrawn(address indexed _sender)
func (_TellorStake *TellorStakeFilterer) ParseStakeWithdrawn(log types.Log) (*TellorStakeStakeWithdrawn, error) {
	event := new(TellorStakeStakeWithdrawn)
	if err := _TellorStake.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TellorStorageABI is the input ABI used to generate the binding from.
const TellorStorageABI = "[]"

// TellorStorageBin is the compiled bytecode used for deploying new contracts.
var TellorStorageBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820066a05a6a7ec5411b6a3ac2af7ab173da38b7122fa780557ee54dbada7f687ac64736f6c63430005100032"

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
func (_TellorStorage *TellorStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_TellorStorage *TellorStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
const TellorTransferABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeAmount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TellorTransferFuncSigs maps the 4-byte function signature to its string representation.
var TellorTransferFuncSigs = map[string]string{
	"bf32006c": "allowance(TellorStorage.TellorStorageStruct storage,address,address)",
	"acaab9e2": "allowedToTrade(TellorStorage.TellorStorageStruct storage,address,uint256)",
	"850dcc32": "approve(TellorStorage.TellorStorageStruct storage,address,uint256)",
	"93b182b3": "balanceOf(TellorStorage.TellorStorageStruct storage,address)",
	"3f48b1ff": "balanceOfAt(TellorStorage.TellorStorageStruct storage,address,uint256)",
	"c7bb46ad": "doTransfer(TellorStorage.TellorStorageStruct storage,address,address,uint256)",
	"60c7dc47": "stakeAmount()",
	"c84b96f5": "transfer(TellorStorage.TellorStorageStruct storage,address,uint256)",
	"ca501899": "transferFrom(TellorStorage.TellorStorageStruct storage,address,address,uint256)",
	"eb7bdc08": "updateBalanceAtNow(TellorStorage.Checkpoint[] storage,uint256)",
}

// TellorTransferBin is the compiled bytecode used for deploying new contracts.
var TellorTransferBin = "0x610adf610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100a85760003560e01c8063bf32006c11610070578063bf32006c146101aa578063c7bb46ad146101de578063c84b96f514610229578063ca50189914610268578063eb7bdc08146102b1576100a8565b80633f48b1ff146100ad57806360c7dc47146100f1578063850dcc32146100f957806393b182b31461014c578063acaab9e214610178575b600080fd5b6100df600480360360608110156100c357600080fd5b508035906001600160a01b0360208201351690604001356102e1565b60408051918252519081900360200190f35b6100df61048e565b81801561010557600080fd5b506101386004803603606081101561011c57600080fd5b508035906001600160a01b0360208201351690604001356104b2565b604080519115158252519081900360200190f35b6100df6004803603604081101561016257600080fd5b50803590602001356001600160a01b03166105f1565b6101386004803603606081101561018e57600080fd5b508035906001600160a01b0360208201351690604001356105fe565b6100df600480360360608110156101c057600080fd5b508035906001600160a01b03602082013581169160400135166106ae565b8180156101ea57600080fd5b506102276004803603608081101561020157600080fd5b508035906001600160a01b036020820135811691604081013590911690606001356106db565b005b81801561023557600080fd5b506101386004803603606081101561024c57600080fd5b508035906001600160a01b0360208201351690604001356108b4565b81801561027457600080fd5b506101386004803603608081101561028b57600080fd5b508035906001600160a01b036020820135811691604081013590911690606001356108cc565b8180156102bd57600080fd5b50610227600480360360408110156102d457600080fd5b508035906020013561097e565b6001600160a01b03821660009081526045840160205260408120805415806103295750828160008154811061031257fe5b6000918252602090912001546001600160801b0316115b15610338576000915050610487565b80548190600019810190811061034a57fe5b6000918252602090912001546001600160801b0316831061039c5780548190600019810190811061037757fe5b600091825260209091200154600160801b90046001600160801b031691506104879050565b8054600090600119015b818111156104545760006002600183850101049050858482815481106103c857fe5b6000918252602090912001546001600160801b03161415610417578381815481106103ef57fe5b600091825260209091200154600160801b90046001600160801b031694506104879350505050565b8584828154811061042457fe5b6000918252602090912001546001600160801b031610156104475780925061044e565b6001810391505b506103a6565b82828154811061046057fe5b600091825260209091200154600160801b90046001600160801b0316935061048792505050565b9392505050565b7f7be108969d31a3f0b261465c71f2b0ba9301cd914d55d9091c3b36a49d4d41b281565b60006001600160a01b038316610506576040805162461bcd60e51b81526020600482015260146024820152735370656e64657220697320302d6164647265737360601b604482015290519081900360640190fd5b33600090815260468501602090815260408083206001600160a01b03871684529091529020541580610536575081155b610587576040805162461bcd60e51b815260206004820152601b60248201527f5370656e64657220697320616c726561647920617070726f7665640000000000604482015290519081900360640190fd5b33600081815260468601602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b60006104878383436102e1565b6001600160a01b03821660009081526047840160205260408120541580159061064157506001600160a01b03831660009081526047850160205260409020546005115b15610699577f7be108969d31a3f0b261465c71f2b0ba9301cd914d55d9091c3b36a49d4d41b260009081526040808601602052902054829061068386866105f1565b031061069157506001610487565b506000610487565b816106a485856105f1565b1015949350505050565b6001600160a01b039182166000908152604693909301602090815260408085209290931684525290205490565b806107175760405162461bcd60e51b8152600401808060200182810382526021815260200180610a8a6021913960400191505060405180910390fd5b6001600160a01b03821661076a576040805162461bcd60e51b815260206004820152601560248201527452656365697665722069732030206164647265737360581b604482015290519081900360640190fd5b6107758484836105fe565b6107b05760405162461bcd60e51b8152600401808060200182810382526027815260200180610a636027913960400191505060405180910390fd5b60006107bc85856105f1565b6001600160a01b038516600090815260458701602052604090209091506107e59083830361097e565b6107ef85846105f1565b905080828201101561083c576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b6001600160a01b038316600090815260458601602052604090206108629082840161097e565b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a35050505050565b60006108c2843385856106db565b5060019392505050565b6001600160a01b0383166000908152604685016020908152604080832033845290915281205482111561093b576040805162461bcd60e51b8152602060048201526012602482015271416c6c6f77616e63652069732077726f6e6760701b604482015290519081900360640190fd5b6001600160a01b03841660009081526046860160209081526040808320338452909152902080548390039055610973858585856106db565b506001949350505050565b815415806109b35750815443908390600019810190811061099b57fe5b6000918252602090912001546001600160801b031614155b15610a245760408051808201909152436001600160801b0390811682528281166020808401918252855460018101875560008781529190912093519301805491516fffffffffffffffffffffffffffffffff19909216938316939093178216600160801b9190921602179055610a5e565b815460009083906000198101908110610a3957fe5b600091825260209091200180546001600160801b03808516600160801b029116179055505b505056fe53686f756c6420686176652073756666696369656e742062616c616e636520746f207472616465547269656420746f2073656e64206e6f6e2d706f73697469766520616d6f756e74a265627a7a72315820663e7dc1ed6de9e71b4a9b09ec96b19f3ac21d0d5dee35a8c04a2e73577ba8f064736f6c63430005100032"

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
func (_TellorTransfer *TellorTransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_TellorTransfer *TellorTransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) StakeAmount(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "stakeAmount")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) StakeAmount() ([32]byte, error) {
	return _TellorTransfer.Contract.StakeAmount(&_TellorTransfer.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) StakeAmount() ([32]byte, error) {
	return _TellorTransfer.Contract.StakeAmount(&_TellorTransfer.CallOpts)
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
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
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
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_TellorTransfer *TellorTransferFilterer) ParseApproval(log types.Log) (*TellorTransferApproval, error) {
	event := new(TellorTransferApproval)
	if err := _TellorTransfer.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_TellorTransfer *TellorTransferFilterer) ParseTransfer(log types.Log) (*TellorTransferTransfer, error) {
	event := new(TellorTransferTransfer)
	if err := _TellorTransfer.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UtilitiesABI is the input ABI used to generate the binding from.
const UtilitiesABI = "[]"

// UtilitiesBin is the compiled bytecode used for deploying new contracts.
var UtilitiesBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582055003b332f9e3c6f8fcf8cb45869ae48a6ff815b35e81f7c7932f9831e665a2f64736f6c63430005100032"

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
func (_Utilities *UtilitiesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_Utilities *UtilitiesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
