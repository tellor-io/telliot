// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tellorCurrent

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
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158205339e8c5b6a23aba8502939e43402a6ccf9231a5062796981969a011bfe3062d64736f6c63430005100032"

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

// TellorABI is the input ABI used to generate the binding from.
const TellorABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minerIndex\",\"type\":\"uint256\"}],\"name\":\"beginDispute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNewCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256\",\"name\":\"_difficutly\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNewVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"idsOnDeck\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"tipsOnDeck\",\"type\":\"uint256[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTopRequestIDs\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_propNewTellorAddress\",\"type\":\"address\"}],\"name\":\"proposeFork\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_pendingOwner\",\"type\":\"address\"}],\"name\":\"proposeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"requestStakingWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"}],\"name\":\"submitMiningSolution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"tallyVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"unlockDisputeFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"updateTellor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_supportsDispute\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TellorFuncSigs maps the 4-byte function signature to its string representation.
var TellorFuncSigs = map[string]string{
	"752d49a1": "addTip(uint256,uint256)",
	"095ea7b3": "approve(address,uint256)",
	"8581af19": "beginDispute(uint256,uint256,uint256)",
	"4e71e0c8": "claimOwnership()",
	"313ce567": "decimals()",
	"0d2d76a2": "depositStake()",
	"4049f198": "getNewCurrentVariables()",
	"9a7077ab": "getNewVariablesOnDeck()",
	"fe1cd15d": "getTopRequestIDs()",
	"06fdde03": "name()",
	"26b7d9f6": "proposeFork(address)",
	"710bf322": "proposeOwnership(address)",
	"28449c3a": "requestStakingWithdraw()",
	"4350283e": "submitMiningSolution(string,uint256[5],uint256[5])",
	"95d89b41": "symbol()",
	"4d318b0e": "tallyVotes(uint256)",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"9a01ca13": "unlockDisputeFee(uint256)",
	"f458ab98": "updateTellor(uint256)",
	"c9d27afe": "vote(uint256,bool)",
	"bed9d861": "withdrawStake()",
}

// TellorBin is the compiled bytecode used for deploying new contracts.
var TellorBin = "0x608060405234801561001057600080fd5b506110dd806100206000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c8063710bf322116100b85780639a7077ab1161007c5780639a7077ab14610415578063a9059cbb14610481578063bed9d861146104ad578063c9d27afe146104b5578063f458ab98146104da578063fe1cd15d146104f757610142565b8063710bf3221461037e578063752d49a1146103a45780638581af19146103c757806395d89b41146103f05780639a01ca13146103f857610142565b806328449c3a1161010a57806328449c3a1461026a578063313ce567146102725780634049f198146102905780634350283e146102e55780634d318b0e146103595780634e71e0c81461037657610142565b806306fdde0314610147578063095ea7b3146101c45780630d2d76a21461020457806323b872dd1461020e57806326b7d9f614610244575b600080fd5b61014f610537565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610189578181015183820152602001610171565b50505050905090810190601f1680156101b65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101f0600480360360408110156101da57600080fd5b506001600160a01b038135169060200135610560565b604080519115158252519081900360200190f35b61020c6105fd565b005b6101f06004803603606081101561022457600080fd5b506001600160a01b03813581169160208101359091169060400135610667565b61020c6004803603602081101561025a57600080fd5b50356001600160a01b031661070d565b61020c610787565b61027a6107d7565b6040805160ff9092168252519081900360200190f35b6102986107dc565b604051848152602081018460a080838360005b838110156102c35781810151838201526020016102ab565b5050505090500183815260200182815260200194505050505060405180910390f35b61020c60048036036101608110156102fc57600080fd5b81019060208101813564010000000081111561031757600080fd5b82018360208201111561032957600080fd5b8035906020019184600183028401116401000000008311171561034b57600080fd5b919350915060a08101610801565b61020c6004803603602081101561036f57600080fd5b50356108db565b61020c610932565b61020c6004803603602081101561039457600080fd5b50356001600160a01b0316610982565b61020c600480360360408110156103ba57600080fd5b50803590602001356109e1565b61020c600480360360608110156103dd57600080fd5b5080359060208101359060400135610a5b565b61014f610add565b61020c6004803603602081101561040e57600080fd5b5035610afa565b61041d610b51565b604051808360a080838360005b8381101561044257818101518382015260200161042a565b5050505090500182600560200280838360005b8381101561046d578181015183820152602001610455565b505050509050019250505060405180910390f35b6101f06004803603604081101561049757600080fd5b506001600160a01b038135169060200135610b73565b61020c610bdd565b61020c600480360360408110156104cb57600080fd5b50803590602001351515610c2d565b61020c600480360360208110156104f057600080fd5b5035610c8c565b6104ff610ce3565b604051808260a080838360005b8381101561052457818101518382015260200161050c565b5050505090500191505060405180910390f35b60408051808201909152600f81526e54656c6c6f7220547269627574657360881b602082015290565b60408051634286e61960e11b81526000600482018190526001600160a01b038516602483015260448201849052915173__$0a7adcfe26a87eb516915da32ab69f1b09$__9163850dcc32916064808301926020929190829003018186803b1580156105ca57600080fd5b505af41580156105de573d6000803e3d6000fd5b505050506040513d60208110156105f457600080fd5b50519392505050565b6040805163410516b360e11b8152600060048201819052915173__$a099bb662982b258a7aafa50db8f7174f3$__9263820a2d669260248082019391829003018186803b15801561064d57600080fd5b505af4158015610661573d6000803e3d6000fd5b50505050565b6040805163ca50189960e01b81526000600482018190526001600160a01b0380871660248401528516604483015260648201849052915173__$0a7adcfe26a87eb516915da32ab69f1b09$__9163ca501899916084808301926020929190829003018186803b1580156106d957600080fd5b505af41580156106ed573d6000803e3d6000fd5b505050506040513d602081101561070357600080fd5b5051949350505050565b6040805163694bf49f60e01b81526000600482018190526001600160a01b0384166024830152915173__$cd7a2037986a81685ce8a6561121d7deb1$__9263694bf49f9260448082019391829003018186803b15801561076c57600080fd5b505af4158015610780573d6000803e3d6000fd5b5050505050565b60408051633273d79360e21b8152600060048201819052915173__$a099bb662982b258a7aafa50db8f7174f3$__9263c9cf5e4c9260248082019391829003018186803b15801561064d57600080fd5b601290565b60006107e661108a565b6000806107f36000610cfa565b935093509350935090919293565b600073__$4bd309c742d0217566e7eb7764debb4a33$__63a4bc40679091868686866040518663ffffffff1660e01b8152600401808681526020018060200184600560200280828437600083820152601f01601f191690910190508360a080828437600083820152601f01601f191690910183810383528681526020019050868680828437600081840152601f19601f820116905080830192505050965050505050505060006040518083038186803b1580156108bd57600080fd5b505af41580156108d1573d6000803e3d6000fd5b5050505050505050565b6040805163def6fac760e01b815260006004820181905260248201849052915173__$cd7a2037986a81685ce8a6561121d7deb1$__9263def6fac79260448082019391829003018186803b15801561076c57600080fd5b6040805163314691ff60e01b8152600060048201819052915173__$4bd309c742d0217566e7eb7764debb4a33$__9263314691ff9260248082019391829003018186803b15801561064d57600080fd5b6040805163291f8b7360e01b81526000600482018190526001600160a01b0384166024830152915173__$4bd309c742d0217566e7eb7764debb4a33$__9263291f8b739260448082019391829003018186803b15801561076c57600080fd5b604080516302e8f21b60e01b81526000600482018190526024820185905260448201849052915173__$4bd309c742d0217566e7eb7764debb4a33$__926302e8f21b9260648082019391829003018186803b158015610a3f57600080fd5b505af4158015610a53573d6000803e3d6000fd5b505050505050565b6040805163ca9a4ea560e01b8152600060048201819052602482018690526044820185905260648201849052915173__$cd7a2037986a81685ce8a6561121d7deb1$__9263ca9a4ea59260848082019391829003018186803b158015610ac057600080fd5b505af4158015610ad4573d6000803e3d6000fd5b50505050505050565b6040805180820190915260038152622a292160e91b602082015290565b60408051634bfafcb960e11b815260006004820181905260248201849052915173__$cd7a2037986a81685ce8a6561121d7deb1$__926397f5f9729260448082019391829003018186803b15801561076c57600080fd5b610b5961108a565b610b6161108a565b610b6b6000610da9565b915091509091565b6040805163c84b96f560e01b81526000600482018190526001600160a01b038516602483015260448201849052915173__$0a7adcfe26a87eb516915da32ab69f1b09$__9163c84b96f5916064808301926020929190829003018186803b1580156105ca57600080fd5b604080516344bacc4b60e01b8152600060048201819052915173__$a099bb662982b258a7aafa50db8f7174f3$__926344bacc4b9260248082019391829003018186803b15801561064d57600080fd5b604080516316d0383760e11b8152600060048201819052602482018590528315156044830152915173__$cd7a2037986a81685ce8a6561121d7deb1$__92632da0706e9260648082019391829003018186803b158015610a3f57600080fd5b604080516322048ecf60e01b815260006004820181905260248201849052915173__$cd7a2037986a81685ce8a6561121d7deb1$__926322048ecf9260448082019391829003018186803b15801561076c57600080fd5b610ceb61108a565b610cf56000610e49565b905090565b6000610d0461108a565b600080805b6005811015610d3f57856035018160058110610d2157fe5b6002020154848260058110610d3257fe5b6020020152600101610d09565b505083546040805169646966666963756c747960b01b8152815190819003600a01812060009081528288016020818152848320546f63757272656e74546f74616c5469707360801b8552855194859003601001909420835252919091205491945091509193509193565b610db161108a565b610db961108a565b610dc283610e49565b915060005b6005811015610e4357836048016000848360058110610de257fe5b6020020151815260200190815260200160002060040160006040518080670746f74616c5469760c41b81525060080190506040518091039020815260200190815260200160002054828260058110610e3657fe5b6020020152600101610dc7565b50915091565b610e5161108a565b610e5961108a565b610e6161108a565b60408051610660810191829052610e9a91600187019060339082845b815481526020019060010190808311610e7d575050505050610f36565b909250905060005b6005811015610f2e57828160058110610eb757fe5b602002015115610efd57846043016000838360058110610ed357fe5b6020020151815260200190815260200160002054848260058110610ef357fe5b6020020152610f26565b846035018160040360058110610f0f57fe5b6002020154848260058110610f2057fe5b60200201525b600101610ea2565b505050919050565b610f3e61108a565b610f4661108a565b60208301516000805b6005811015610fc957858160010160338110610f6757fe5b6020020151858260058110610f7857fe5b602002015260018101848260058110610f8d57fe5b602002015282858260058110610f9f57fe5b60200201511015610fc157848160058110610fb657fe5b602002015192508091505b600101610f4f565b5060065b60338110156110825782868260338110610fe357fe5b6020020151111561107a57858160338110610ffa57fe5b602002015185836005811061100b57fe5b60200201528084836005811061101d57fe5b602002015285816033811061102e57fe5b6020020151925060005b6005811015611078578386826005811061104e57fe5b602002015110156110705785816005811061106557fe5b602002015193508092505b600101611038565b505b600101610fcd565b505050915091565b6040518060a00160405280600590602082028038833950919291505056fea265627a7a7231582012641903dc272b06302c8eca3e27a4e04b6dcbeed7996685f3c8d8c917df1fb264736f6c63430005100032"

// DeployTellor deploys a new Ethereum contract, binding an instance of Tellor to it.
func DeployTellor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tellor, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	tellorTransferAddr, _, _, _ := DeployTellorTransfer(auth, backend)
	TellorBin = strings.Replace(TellorBin, "__$0a7adcfe26a87eb516915da32ab69f1b09$__", tellorTransferAddr.String()[2:], -1)

	tellorLibraryAddr, _, _, _ := DeployTellorLibrary(auth, backend)
	TellorBin = strings.Replace(TellorBin, "__$4bd309c742d0217566e7eb7764debb4a33$__", tellorLibraryAddr.String()[2:], -1)

	tellorStakeAddr, _, _, _ := DeployTellorStake(auth, backend)
	TellorBin = strings.Replace(TellorBin, "__$a099bb662982b258a7aafa50db8f7174f3$__", tellorStakeAddr.String()[2:], -1)

	tellorDisputeAddr, _, _, _ := DeployTellorDispute(auth, backend)
	TellorBin = strings.Replace(TellorBin, "__$cd7a2037986a81685ce8a6561121d7deb1$__", tellorDisputeAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TellorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tellor{TellorCaller: TellorCaller{contract: contract}, TellorTransactor: TellorTransactor{contract: contract}, TellorFilterer: TellorFilterer{contract: contract}}, nil
}

// Tellor is an auto generated Go binding around an Ethereum contract.
type Tellor struct {
	TellorCaller     // Read-only binding to the contract
	TellorTransactor // Write-only binding to the contract
	TellorFilterer   // Log filterer for contract events
}

// TellorCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorSession struct {
	Contract     *Tellor           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorCallerSession struct {
	Contract *TellorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TellorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorTransactorSession struct {
	Contract     *TellorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorRaw struct {
	Contract *Tellor // Generic contract binding to access the raw methods on
}

// TellorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorCallerRaw struct {
	Contract *TellorCaller // Generic read-only contract binding to access the raw methods on
}

// TellorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorTransactorRaw struct {
	Contract *TellorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellor creates a new instance of Tellor, bound to a specific deployed contract.
func NewTellor(address common.Address, backend bind.ContractBackend) (*Tellor, error) {
	contract, err := bindTellor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tellor{TellorCaller: TellorCaller{contract: contract}, TellorTransactor: TellorTransactor{contract: contract}, TellorFilterer: TellorFilterer{contract: contract}}, nil
}

// NewTellorCaller creates a new read-only instance of Tellor, bound to a specific deployed contract.
func NewTellorCaller(address common.Address, caller bind.ContractCaller) (*TellorCaller, error) {
	contract, err := bindTellor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorCaller{contract: contract}, nil
}

// NewTellorTransactor creates a new write-only instance of Tellor, bound to a specific deployed contract.
func NewTellorTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorTransactor, error) {
	contract, err := bindTellor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorTransactor{contract: contract}, nil
}

// NewTellorFilterer creates a new log filterer instance of Tellor, bound to a specific deployed contract.
func NewTellorFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorFilterer, error) {
	contract, err := bindTellor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorFilterer{contract: contract}, nil
}

// bindTellor binds a generic wrapper to an already deployed contract.
func bindTellor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tellor *TellorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tellor.Contract.TellorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tellor *TellorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tellor.Contract.TellorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tellor *TellorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tellor.Contract.TellorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tellor *TellorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tellor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tellor *TellorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tellor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tellor *TellorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tellor.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Tellor *TellorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Tellor *TellorSession) Decimals() (uint8, error) {
	return _Tellor.Contract.Decimals(&_Tellor.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Tellor *TellorCallerSession) Decimals() (uint8, error) {
	return _Tellor.Contract.Decimals(&_Tellor.CallOpts)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficutly, uint256 _tip)
func (_Tellor *TellorCaller) GetNewCurrentVariables(opts *bind.CallOpts) (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "getNewCurrentVariables")

	outstruct := new(struct {
		Challenge  [32]byte
		RequestIds [5]*big.Int
		Difficutly *big.Int
		Tip        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Challenge = out[0].([32]byte)
	outstruct.RequestIds = out[1].([5]*big.Int)
	outstruct.Difficutly = out[2].(*big.Int)
	outstruct.Tip = out[3].(*big.Int)

	return *outstruct, err

}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficutly, uint256 _tip)
func (_Tellor *TellorSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	return _Tellor.Contract.GetNewCurrentVariables(&_Tellor.CallOpts)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficutly, uint256 _tip)
func (_Tellor *TellorCallerSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	return _Tellor.Contract.GetNewCurrentVariables(&_Tellor.CallOpts)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_Tellor *TellorCaller) GetNewVariablesOnDeck(opts *bind.CallOpts) (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "getNewVariablesOnDeck")

	outstruct := new(struct {
		IdsOnDeck  [5]*big.Int
		TipsOnDeck [5]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IdsOnDeck = out[0].([5]*big.Int)
	outstruct.TipsOnDeck = out[1].([5]*big.Int)

	return *outstruct, err

}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_Tellor *TellorSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _Tellor.Contract.GetNewVariablesOnDeck(&_Tellor.CallOpts)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_Tellor *TellorCallerSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _Tellor.Contract.GetNewVariablesOnDeck(&_Tellor.CallOpts)
}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_Tellor *TellorCaller) GetTopRequestIDs(opts *bind.CallOpts) ([5]*big.Int, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "getTopRequestIDs")

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_Tellor *TellorSession) GetTopRequestIDs() ([5]*big.Int, error) {
	return _Tellor.Contract.GetTopRequestIDs(&_Tellor.CallOpts)
}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_Tellor *TellorCallerSession) GetTopRequestIDs() ([5]*big.Int, error) {
	return _Tellor.Contract.GetTopRequestIDs(&_Tellor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Tellor *TellorCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Tellor *TellorSession) Name() (string, error) {
	return _Tellor.Contract.Name(&_Tellor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Tellor *TellorCallerSession) Name() (string, error) {
	return _Tellor.Contract.Name(&_Tellor.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Tellor *TellorCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Tellor *TellorSession) Symbol() (string, error) {
	return _Tellor.Contract.Symbol(&_Tellor.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Tellor *TellorCallerSession) Symbol() (string, error) {
	return _Tellor.Contract.Symbol(&_Tellor.CallOpts)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_Tellor *TellorTransactor) AddTip(opts *bind.TransactOpts, _requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "addTip", _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_Tellor *TellorSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.AddTip(&_Tellor.TransactOpts, _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_Tellor *TellorTransactorSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.AddTip(&_Tellor.TransactOpts, _requestId, _tip)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Tellor *TellorTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Tellor *TellorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.Approve(&_Tellor.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Tellor *TellorTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.Approve(&_Tellor.TransactOpts, _spender, _amount)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_Tellor *TellorTransactor) BeginDispute(opts *bind.TransactOpts, _requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "beginDispute", _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_Tellor *TellorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.BeginDispute(&_Tellor.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_Tellor *TellorTransactorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.BeginDispute(&_Tellor.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Tellor *TellorTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Tellor *TellorSession) ClaimOwnership() (*types.Transaction, error) {
	return _Tellor.Contract.ClaimOwnership(&_Tellor.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Tellor *TellorTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _Tellor.Contract.ClaimOwnership(&_Tellor.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Tellor *TellorTransactor) DepositStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "depositStake")
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Tellor *TellorSession) DepositStake() (*types.Transaction, error) {
	return _Tellor.Contract.DepositStake(&_Tellor.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Tellor *TellorTransactorSession) DepositStake() (*types.Transaction, error) {
	return _Tellor.Contract.DepositStake(&_Tellor.TransactOpts)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_Tellor *TellorTransactor) ProposeFork(opts *bind.TransactOpts, _propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "proposeFork", _propNewTellorAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_Tellor *TellorSession) ProposeFork(_propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _Tellor.Contract.ProposeFork(&_Tellor.TransactOpts, _propNewTellorAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_Tellor *TellorTransactorSession) ProposeFork(_propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _Tellor.Contract.ProposeFork(&_Tellor.TransactOpts, _propNewTellorAddress)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _pendingOwner) returns()
func (_Tellor *TellorTransactor) ProposeOwnership(opts *bind.TransactOpts, _pendingOwner common.Address) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "proposeOwnership", _pendingOwner)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _pendingOwner) returns()
func (_Tellor *TellorSession) ProposeOwnership(_pendingOwner common.Address) (*types.Transaction, error) {
	return _Tellor.Contract.ProposeOwnership(&_Tellor.TransactOpts, _pendingOwner)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _pendingOwner) returns()
func (_Tellor *TellorTransactorSession) ProposeOwnership(_pendingOwner common.Address) (*types.Transaction, error) {
	return _Tellor.Contract.ProposeOwnership(&_Tellor.TransactOpts, _pendingOwner)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Tellor *TellorTransactor) RequestStakingWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "requestStakingWithdraw")
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Tellor *TellorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _Tellor.Contract.RequestStakingWithdraw(&_Tellor.TransactOpts)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Tellor *TellorTransactorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _Tellor.Contract.RequestStakingWithdraw(&_Tellor.TransactOpts)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_Tellor *TellorTransactor) SubmitMiningSolution(opts *bind.TransactOpts, _nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "submitMiningSolution", _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_Tellor *TellorSession) SubmitMiningSolution(_nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.SubmitMiningSolution(&_Tellor.TransactOpts, _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_Tellor *TellorTransactorSession) SubmitMiningSolution(_nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.SubmitMiningSolution(&_Tellor.TransactOpts, _nonce, _requestId, _value)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_Tellor *TellorTransactor) TallyVotes(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "tallyVotes", _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_Tellor *TellorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.TallyVotes(&_Tellor.TransactOpts, _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_Tellor *TellorTransactorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.TallyVotes(&_Tellor.TransactOpts, _disputeId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Tellor *TellorTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Tellor *TellorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.Transfer(&_Tellor.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Tellor *TellorTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.Transfer(&_Tellor.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Tellor *TellorTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Tellor *TellorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.TransferFrom(&_Tellor.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Tellor *TellorTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.TransferFrom(&_Tellor.TransactOpts, _from, _to, _amount)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_Tellor *TellorTransactor) UnlockDisputeFee(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "unlockDisputeFee", _disputeId)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_Tellor *TellorSession) UnlockDisputeFee(_disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.UnlockDisputeFee(&_Tellor.TransactOpts, _disputeId)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_Tellor *TellorTransactorSession) UnlockDisputeFee(_disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.UnlockDisputeFee(&_Tellor.TransactOpts, _disputeId)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_Tellor *TellorTransactor) UpdateTellor(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "updateTellor", _disputeId)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_Tellor *TellorSession) UpdateTellor(_disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.UpdateTellor(&_Tellor.TransactOpts, _disputeId)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_Tellor *TellorTransactorSession) UpdateTellor(_disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.UpdateTellor(&_Tellor.TransactOpts, _disputeId)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_Tellor *TellorTransactor) Vote(opts *bind.TransactOpts, _disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "vote", _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_Tellor *TellorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Tellor.Contract.Vote(&_Tellor.TransactOpts, _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_Tellor *TellorTransactorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Tellor.Contract.Vote(&_Tellor.TransactOpts, _disputeId, _supportsDispute)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Tellor *TellorTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Tellor *TellorSession) WithdrawStake() (*types.Transaction, error) {
	return _Tellor.Contract.WithdrawStake(&_Tellor.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Tellor *TellorTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _Tellor.Contract.WithdrawStake(&_Tellor.TransactOpts)
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
var TellorDisputeBin = "0x612265610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100875760003560e01c806397f5f9721161006557806397f5f9721461012f578063ca9a4ea51461015f578063d7b651c11461019b578063def6fac7146101c557610087565b806322048ecf1461008c5780632da0706e146100be578063694bf49f146100f6575b600080fd5b81801561009857600080fd5b506100bc600480360360408110156100af57600080fd5b50803590602001356101f5565b005b8180156100ca57600080fd5b506100bc600480360360608110156100e157600080fd5b508035906020810135906040013515156103a1565b81801561010257600080fd5b506100bc6004803603604081101561011957600080fd5b50803590602001356001600160a01b0316610651565b81801561013b57600080fd5b506100bc6004803603604081101561015257600080fd5b5080359060200135610b35565b81801561016b57600080fd5b506100bc6004803603608081101561018257600080fd5b508035906020810135906040810135906060013561133e565b8180156101a757600080fd5b506100bc600480360360208110156101be57600080fd5b5035611d91565b8180156101d157600080fd5b506100bc600480360360408110156101e857600080fd5b5080359060200135611e82565b6000818152604483016020818152604080842054808552604a870183528185205480865284845282862083516c64697370757465526f756e647360981b8152845190819003600d01812088526005909101808652848820548287015284518083038701815291850185528151918601919091208752845282862054808752949093529320600281015491929160ff6101009091041615156001146102d5576040805162461bcd60e51b8152602060048201526012602482015271766f7465206e6565647320746f207061737360701b604482015290519081900360640190fd5b604080516874616c6c794461746560b81b815281519081900360090190206000908152600583016020522054620151804291909103116103465760405162461bcd60e51b81526004018080602001828103825260338152602001806121de6033913960400191505060405180910390fd5b60040154604080516d1d195b1b1bdc90dbdb9d1c9858dd60921b8152815190819003600e0190206000908152603f90970160205290952080546001600160a01b0319166001600160a01b039096169590951790945550505050565b600082815260448085016020908152604080842081516a313637b1b5a73ab6b132b960a91b8152825190819003600b018120865260058201845282862054633f48b1ff60e01b8252600482018a905233602483015294810194909452905190939273__$0a7adcfe26a87eb516915da32ab69f1b09$__92633f48b1ff92606480840193829003018186803b15801561043857600080fd5b505af415801561044c573d6000803e3d6000fd5b505050506040513d602081101561046257600080fd5b505133600090815260068401602052604090205490915060ff161515600114156104d3576040805162461bcd60e51b815260206004820152601860248201527f53656e6465722068617320616c726561647920766f7465640000000000000000604482015290519081900360640190fd5b80610519576040805162461bcd60e51b81526020600482015260116024820152700557365722062616c616e6365206973203607c1b604482015290519081900360640190fd5b33600090815260478601602052604090205460031415610579576040805162461bcd60e51b81526020600482015260166024820152754d696e657220697320756e646572206469737075746560501b604482015290519081900360640190fd5b3360009081526006830160209081526040808320805460ff1916600190811790915581516c6e756d6265724f66566f74657360981b8152825190819003600d0190208452600586019092529091208054909101905582156105f35760018201546105e9908263ffffffff61213e16565b600183015561060e565b6001820154610608908263ffffffff61216f16565b60018301555b60408051841515815290518291339187917f911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e919081900360200190a45050505050565b604080516001600160a01b0383166020808301919091528251808303820181528284018085528151919092012063c7bb46ad60e01b9091526044820185905233606483015230608483015268056bc75e2d6310000060a4830152915173__$0a7adcfe26a87eb516915da32ab69f1b09$__9163c7bb46ad9160c4808301926000929190829003018186803b1580156106e857600080fd5b505af41580156106fc573d6000803e3d6000fd5b5050604080516b191a5cdc1d5d1950dbdd5b9d60a21b808252825191829003600c90810183206000908152848a016020818152868320805460010190559385528551948590039092019093208352815282822054868352604a8901909152919020549092501590506107b2576000828152604a850160209081526040808320548484526044880183528184208251651bdc9a59d25160d21b815283519081900360060190208552600501909252909120556107c6565b6000828152604a8501602052604090208190555b6000828152604a850160209081526040808320548084526044880180845282852083516c64697370757465526f756e647360981b808252855191829003600d9081018320895260059093018088528689208054600101905585895284885290825285519182900390920181208752818652848720548488529286528086018390528451808203870181529085018552805190860120865290935292208390558282146109d557600082815260448701602081815260408084208151600019870181850152825180820385018152818401808552815191860191909120875260059283018552838720548088529585528387206f6d696e457865637574696f6e4461746560801b90915283519182900360500190912086520190915290912054421015610933576040805162461bcd60e51b81526020600482015260176024820152762234b9b83aba329034b99030b63932b0b23c9037b832b760491b604482015290519081900360640190fd5b600081815260448801602052604090206002015460ff16156109d3576000818152604488016020908152604080832081516874616c6c794461746560b81b81528251908190036009019020845260050190915290205462015180429190910311156109d3576040805162461bcd60e51b815260206004820152601f6024820152600080516020612211833981519152604482015290519081900360640190fd5b505b505060408051610100808201835293815260006020808301828152838501838152606085018481526001608087018181523360a0890181815260c08a019182526001600160a01b039d8e1660e08b019081528c8a526044909f018089528b8a209a518b559651938a019390935593516002890180549451925193518e166301000000026301000000600160b81b0319941515620100000262ff000019941515909e0261ff001993151560ff19909716969096179290921694909417919091169a909a17169890981790975595516003840180549189166001600160a01b031992831617905597516004840180549190981698169790971790955581516a313637b1b5a73ab6b132b960a91b8152825190819003600b018120865260059091018087528286204390559285529285526f6d696e457865637574696f6e4461746560801b835280519283900360100190922083529092522062093a8042019055565b60008181526044830160208181526040808420548452604a860182528084205480855292825280842081516c64697370757465526f756e647360981b8152825190819003600d018120865260059091018084528286205482850152825180830385018152918301835281519184019190912085529091529091205480610bb85750805b6000828152604485016020908152604080832084845281842082516c64697370757465526f756e647360981b8152835190819003600d01902085526005820190935292205480610c06575060015b60408051631c185a5960e21b815281519081900360040190206000908152600585016020529081205415610c74576040805162461bcd60e51b815260206004820152601060248201526f185b1c9958591e481c185a59081bdd5d60821b604482015290519081900360640190fd5b604080516874616c6c794461746560b81b81528151908190036009019020600090815260058501602052205462015180429190910311610ce9576040805162461bcd60e51b815260206004820152601f6024820152600080516020612211833981519152604482015290519081900360640190fd5b600284810154630100000090046001600160a01b0316600090815260478a01602090815260408083208151631c185a5960e21b8152825160049181900391909101902084526005890190925290912060019081905591850154909161010090910460ff1615151415610fd55762015180420642036001820155604080516a1cdd185ad95c90dbdd5b9d60aa1b8152815190819003600b0190206000908152818b01602052208054600019019055610d9f89611d91565b805460041415610e74576005815560028501546003860154604080516a1cdd185ad9505b5bdd5b9d60aa1b8152815190819003600b0181206000908152828e016020528281205463c7bb46ad60e01b8352600483018f90526001600160a01b0363010000009096048616602484015293909416604482015260648101929092525173__$0a7adcfe26a87eb516915da32ab69f1b09$__9263c7bb46ad9260848082019391829003018186803b158015610e5757600080fd5b505af4158015610e6b573d6000803e3d6000fd5b50506000835550505b60005b83811015610fcf576040805182860360208083019190915282518083038201815291830183528151918101919091206000908152600589019091522054925082610ebf578792505b60008a6044016000858152602001908152602001600020905073__$0a7adcfe26a87eb516915da32ab69f1b09$__63c7bb46ad8c308460030160009054906101000a90046001600160a01b031685600501600060405180806266656560e81b815250600301905060405180910390208152602001908152602001600020546040518563ffffffff1660e01b815260040180858152602001846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b0316815260200182815260200194505050505060006040518083038186803b158015610faa57600080fd5b505af4158015610fbe573d6000803e3d6000fd5b505060019093019250610e77915050565b50611292565b6001815560408051681c995c5d595cdd125960ba1b815281519081900360099081018220600090815260058901602081815285832054835260488f018152858320681b5a5b995c94db1bdd60ba1b8652865195869003909401909420825290925291902054600214156110a157604080516476616c756560d81b815281519081900360059081018220600090815290890160208181528483205468074696d657374616d760bc1b8552855194859003600901909420835290815283822054825260068501905291909120555b6040805168074696d657374616d760bc1b815281519081900360090190206000908152600588016020908152828220548252600784019052205460ff1615156001141561112a576040805168074696d657374616d760bc1b81528151908190036009019020600090815260058801602090815282822054825260078401905220805460ff191690555b60005b8481101561128f57604080518287036020808301919091528251808303820181529183018352815191810191909120600090815260058a0190915220549350831561118557600084815260448c016020526040902095505b73__$0a7adcfe26a87eb516915da32ab69f1b09$__63c7bb46ad8c308960020160039054906101000a90046001600160a01b03168f60440160008a8152602001908152602001600020600501600060405180806266656560e81b815250600301905060405180910390208152602001908152602001600020546040518563ffffffff1660e01b815260040180858152602001846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b0316815260200182815260200194505050505060006040518083038186803b15801561126b57600080fd5b505af415801561127f573d6000803e3d6000fd5b50506001909201915061112d9050565b50505b60408051681b5a5b995c94db1bdd60ba1b815281519081900360090190206000908152600587016020522054600214156113335760408051681c995c5d595cdd125960ba1b81528151908190036009018120600090815260058801602090815283822054825260488d0181528382206b191a5cdc1d5d1950dbdd5b9d60a21b8452845193849003600c01909320825260049092019091522080546000190190555b505050505050505050565b6000838152604885016020908152604080832085845260058101909252909120546113a3576040805162461bcd60e51b815260206004820152601060248201526f04d696e656420626c6f636b20697320360841b604482015290519081900360640190fd5b600582106113ef576040805162461bcd60e51b81526020600482015260146024820152734d696e657220696e6465782069732077726f6e6760601b604482015290519081900360640190fd5b60008381526008820160205260408120836005811061140a57fe5b0154604080516bffffffffffffffffffffffff19606084901b1660208083019190915260348201899052605480830189905283518084039091018152607483018085528151918301919091206b191a5cdc1d5d1950dbdd5b9d60a21b91829052845193849003608001842060009081528c860180855286822054938652865195869003600c01909520815293835284842060019290920191829055808452604a8c0190925292909120546001600160a01b03909316935091801561150657600082815260448a01602090815260408083208151651bdc9a59d25160d21b815282519081900360060190208452600501909152902081905561151c565b506000828152604a890160205260409020819055805b600081815260448a016020818152604080842081516c64697370757465526f756e647360981b808252835191829003600d908101832088526005909301808652848820548989528787529183528451928390039093018220875282855283872060019091019081905587875294845280840185905282518082038501815290830183528051908401208552909152909120839055819083821461172257600082815260448c01602081815260408084208151600019870181850152825180820385018152818401808552815191860191909120875260059283018552838720548088529585528387206f6d696e457865637574696f6e4461746560801b90915283519182900360500190912086520190915290912054421015611680576040805162461bcd60e51b81526020600482015260176024820152762234b9b83aba329034b99030b63932b0b23c9037b832b760491b604482015290519081900360640190fd5b600081815260448d01602052604090206002015460ff161561172057600081815260448d016020908152604080832081516874616c6c794461746560b81b8152825190819003600901902084526005019091529020546201518042919091031115611720576040805162461bcd60e51b815260206004820152601f6024820152600080516020612211833981519152604482015290519081900360640190fd5b505b6000886002141561185e578b60480160008c8152602001908152602001600020600401600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c01905060405180910390208152602001908152602001600020546001018c60480160008d8152602001908152602001600020600401600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c01905060405180910390208152602001908152602001600020819055508b60480160008c8152602001908152602001600020600401600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c01905060405180910390208152602001908152602001600020548c604001600060405180806a1cdd185ad9505b5bdd5b9d60aa1b815250600b019050604051809103902081526020019081526020016000205402905061188e565b5060408051696469737075746546656560b01b8152815190819003600a0190206000908152818d01602052205481025b60405180610100016040528087815260200160008152602001600015158152602001600015158152602001600015158152602001886001600160a01b03168152602001336001600160a01b0316815260200160006001600160a01b03168152508c6044016000878152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548160ff02191690831515021790555060608201518160020160016101000a81548160ff02191690831515021790555060808201518160020160026101000a81548160ff02191690831515021790555060a08201518160020160036101000a8154816001600160a01b0302191690836001600160a01b0316021790555060c08201518160030160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060e08201518160040160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050508a8c604401600087815260200190815260200160002060050160006040518080681c995c5d595cdd125960ba1b81525060090190506040518091039020815260200190815260200160002081905550898c60440160008781526020019081526020016000206005016000604051808068074696d657374616d760bc1b815250600901905060405180910390208152602001908152602001600020819055508760090160008b81526020019081526020016000208960058110611abf57fe5b01548c6044016000878152602001908152602001600020600501600060405180806476616c756560d81b81525060050190506040518091039020815260200190815260200160002081905550816202a3000242018c6044016000878152602001908152602001600020600501600060405180806f6d696e457865637574696f6e4461746560801b81525060100190506040518091039020815260200190815260200160002081905550438c6044016000878152602001908152602001600020600501600060405180806a313637b1b5a73ab6b132b960a91b815250600b0190506040518091039020815260200190815260200160002081905550888c604401600087815260200190815260200160002060050160006040518080681b5a5b995c94db1bdd60ba1b81525060090190506040518091039020815260200190815260200160002081905550808c6044016000878152602001908152602001600020600501600060405180806266656560e81b8152506003019050604051809103902081526020019081526020016000208190555073__$0a7adcfe26a87eb516915da32ab69f1b09$__63c7bb46ad8d3330856040518563ffffffff1660e01b815260040180858152602001846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b0316815260200182815260200194505050505060006040518083038186803b158015611cdc57600080fd5b505af4158015611cf0573d6000803e3d6000fd5b505050508860021415611d265760008a81526007890160209081526040808320805460ff1916600117905560068b019091528120555b6001600160a01b038716600081815260478e0160209081526040918290206003905581518d81529081019290925280518d9288927feceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da6492918290030190a3505050505050505050505050565b604080516a1cdd185ad9505b5bdd5b9d60aa1b81528151600b918190038201812060009081528484016020818152858320546b7461726765744d696e65727360a01b85528651600c958190039590950185208452828252868420546a1cdd185ad95c90dbdd5b9d60aa1b86528751958690039096019094208352529290922054611e4d9067d02ab486cedc0000906103e8908490611e30908290612195565b6103e802860281611e3d57fe5b0481611e4557fe5b0484036121ad565b60408051696469737075746546656560b01b8152815190819003600a0190206000908152948101602052909320929092555050565b60008181526044830160205260409020600281015460ff1615611ed65760405162461bcd60e51b81526004018080602001828103825260218152602001806121bd6021913960400191505060405180910390fd5b604080516f6d696e457865637574696f6e4461746560801b815281519081900360100190206000908152600583016020522054421015611f4b576040805162461bcd60e51b815260206004820152601f6024820152600080516020612211833981519152604482015290519081900360640190fd5b60038101546001600160a01b0316611faa576040805162461bcd60e51b815260206004820152601c60248201527f7265706f7274696e672050617274792069732061646472657373203000000000604482015290519081900360640190fd5b60018101546000811315611fca5760028201805461ff0019166101001790555b600282015462010000900460ff16612016576002820154630100000090046001600160a01b03166000908152604785016020526040902080546003141561201057600481555b50612093565b604080516b746f74616c5f737570706c7960a01b8152815190819003600c01902060009081528186016020522054606490600a02048110612093576004820154604080516001600160a01b039092168252517fc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d9181900360200190a15b604080516874616c6c794461746560b81b8152815190819003600901812060009081526005850160209081529083902042905560028501805460ff19166001179081905560038601548584526001600160a01b0390811692840192909252610100810460ff1615158385015292516301000000909304169185917f21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739919081900360600190a350505050565b60008082131561215b57508181018281121561215657fe5b612169565b508181018281131561216957fe5b92915050565b60008082131561218757508082038281131561215657fe5b508082038281121561216957fe5b60008183106121a457816121a6565b825b9392505050565b60008183116121a457816121a656fe4469737075746520686173206265656e20616c726561647920657865637574656454696d6520666f7220766f74696e6720666f72206675727468657220646973707574657320686173206e6f742070617373656454696d6520666f7220766f74696e6720686176656e277420656c617073656400a265627a7a72315820dd5032ea3d0ede30810349a8e9f261767b8d2ca448b00951b8906cfd297fc36064736f6c63430005100032"

// DeployTellorDispute deploys a new Ethereum contract, binding an instance of TellorDispute to it.
func DeployTellorDispute(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorDispute, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorDisputeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	tellorTransferAddr, _, _, _ := DeployTellorTransfer(auth, backend)
	TellorDisputeBin = strings.Replace(TellorDisputeBin, "__$0a7adcfe26a87eb516915da32ab69f1b09$__", tellorTransferAddr.String()[2:], -1)

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
	event.Raw = log
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
	event.Raw = log
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
	event.Raw = log
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
	event.Raw = log
	return event, nil
}

// TellorGettersLibraryABI is the input ABI used to generate the binding from.
const TellorGettersLibraryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newTellor\",\"type\":\"address\"}],\"name\":\"NewTellorAddress\",\"type\":\"event\"}]"

// TellorGettersLibraryFuncSigs maps the 4-byte function signature to its string representation.
var TellorGettersLibraryFuncSigs = map[string]string{
	"c93299e9": "didMine(TellorStorage.TellorStorageStruct storage,bytes32,address)",
}

// TellorGettersLibraryBin is the compiled bytecode used for deploying new contracts.
var TellorGettersLibraryBin = "0x60dd610025600b82828239805160001a60731461001857fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063c93299e9146038575b600080fd5b606760048036036060811015604c57600080fd5b50803590602081013590604001356001600160a01b0316607b565b604080519115158252519081900360200190f35b6000918252604192909201602090815260408083206001600160a01b039094168352929052205460ff169056fea265627a7a72315820f1fcae754d246d24d77122fd67302f7e74b6db4ec0755902fec9f84eeedb81b864736f6c63430005100032"

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
	event.Raw = log
	return event, nil
}

// TellorLibraryABI is the input ABI used to generate the binding from.
const TellorLibraryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_currentRequestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"NewChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NewValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NonceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"TipAdded\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_tBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentRequestId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentReward\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentTotalTips\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"devShare\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"difficulty\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pending_owner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requestCount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requestQPosition\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"runningTips\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slotProgress\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeOfLastNewValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeTarget\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalTip\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"total_supply\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TellorLibraryFuncSigs maps the 4-byte function signature to its string representation.
var TellorLibraryFuncSigs = map[string]string{
	"b2bdfa7b": "_owner()",
	"6e3cf885": "_tBlock()",
	"02e8f21b": "addTip(TellorStorage.TellorStorageStruct storage,uint256,uint256)",
	"314691ff": "claimOwnership(TellorStorage.TellorStorageStruct storage)",
	"5ae2bfdb": "currentRequestId()",
	"07621eca": "currentReward()",
	"75ad1a2a": "currentTotalTips()",
	"aed04fae": "devShare()",
	"19cae462": "difficulty()",
	"2dbfd604": "newBlock(TellorStorage.TellorStorageStruct storage,string,uint256[5])",
	"7f4ec4c3": "pending_owner()",
	"291f8b73": "proposeOwnership(TellorStorage.TellorStorageStruct storage,address)",
	"5badbe4c": "requestCount()",
	"2bf07e9e": "requestQPosition()",
	"b0dc7c20": "runningTips()",
	"03b3160f": "slotProgress()",
	"a4bc4067": "submitMiningSolution(TellorStorage.TellorStorageStruct storage,string,uint256[5],uint256[5])",
	"6fd4f229": "timeOfLastNewValue()",
	"6fc37811": "timeTarget()",
	"561cb04a": "totalTip()",
	"3940e9ee": "total_supply()",
	"ef84b45f": "updateOnDeck(TellorStorage.TellorStorageStruct storage,uint256,uint256)",
}

// TellorLibraryBin is the compiled bytecode used for deploying new contracts.
var TellorLibraryBin = "0x61290e610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061014c5760003560e01c80635badbe4c116100c25780637f4ec4c3116100865780637f4ec4c314610344578063a4bc40671461034c578063aed04fae146103d4578063b0dc7c20146103dc578063b2bdfa7b146103e4578063ef84b45f146103ec5761014c565b80635badbe4c1461031c5780636e3cf885146103245780636fc378111461032c5780636fd4f2291461033457806375ad1a2a1461033c5761014c565b80632bf07e9e116101145780632bf07e9e146101ec5780632dbfd604146101f4578063314691ff146102da5780633940e9ee14610304578063561cb04a1461030c5780635ae2bfdb146103145761014c565b806302e8f21b1461015157806303b3160f1461018957806307621eca146101a357806319cae462146101ab578063291f8b73146101b3575b600080fd5b81801561015d57600080fd5b506101876004803603606081101561017457600080fd5b5080359060208101359060400135610422565b005b610191610675565b60408051918252519081900360200190f35b610191610687565b6101916106ab565b8180156101bf57600080fd5b50610187600480360360408110156101d657600080fd5b50803590602001356001600160a01b03166106bd565b6101916107c5565b81801561020057600080fd5b50610187600480360360e081101561021757600080fd5b8135919081019060408101602082013564010000000081111561023957600080fd5b82018360208201111561024b57600080fd5b8035906020019184600183028401116401000000008311171561026d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040805160a081810190925293969594818101949350915060059083908390808284376000920191909152509194506107d79350505050565b8180156102e657600080fd5b50610187600480360360208110156102fd57600080fd5b503561103f565b61019161117b565b61019161119f565b6101916111b1565b6101916111d5565b6101916111f9565b61019161120b565b61019161122f565b610191611241565b610191611253565b81801561035857600080fd5b50610187600480360361018081101561037057600080fd5b8135919081019060408101602082013564010000000081111561039257600080fd5b8201836020820111156103a457600080fd5b803590602001918460018302840111640100000000831117156103c657600080fd5b919350915060a08101611265565b610191611336565b61019161135a565b61019161137e565b8180156103f857600080fd5b506101876004803603606081101561040f57600080fd5b5080359060208101359060400135611390565b81610465576040805162461bcd60e51b815260206004820152600e60248201526d052657175657374496420697320360941b604482015290519081900360640190fd5b806104b7576040805162461bcd60e51b815260206004820152601c60248201527f5469702073686f756c642062652067726561746572207468616e203000000000604482015290519081900360640190fd5b7f05de9147d05477c0a5dc675aeea733157f5092f82add148cf39d579cafe3dc98600090815260408085016020529020546001018281141561052b577f05de9147d05477c0a5dc675aeea733157f5092f82add148cf39d579cafe3dc9860009081526040808601602052902081905561057f565b80831061057f576040805162461bcd60e51b815260206004820181905260248201527f526571756573744964206973206e6f74206c657373207468616e20636f756e74604482015290519081900360640190fd5b6040805163c7bb46ad60e01b81526004810186905233602482015230604482015260648101849052905173__$0a7adcfe26a87eb516915da32ab69f1b09$__9163c7bb46ad916084808301926000929190829003018186803b1580156105e457600080fd5b505af41580156105f8573d6000803e3d6000fd5b50505050610607848484611390565b600083815260488501602090815260408083206000805160206127d48339815191528452600401825291829020548251858152918201528151859233927fd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820929081900390910190a350505050565b6000805160206127f483398151915281565b7f9b6853911475b07474368644a0d922ee13bc76a15cd3e97d3e334326424a47d481565b6000805160206127b483398151915281565b6000805160206127948339815191526000908152603f830160205260409020546001600160a01b0316331461072f576040805162461bcd60e51b815260206004820152601360248201527229b2b73232b91034b9903737ba1037bbb732b960691b604482015290519081900360640190fd5b6000805160206127948339815191526000908152603f830160205260408082205490516001600160a01b03808516939216917fb51454ce8c7f26becd312a46c4815553887f2ec876a0b8dc813b87f62edf6f8091a36000805160206128148339815191526000908152603f92909201602052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b60008051602061285483398151915281565b60008051602061283483398151915260009081526040808501602081815282842054845260488701815282842060008051602061277483398151915285529190529082205490914291909103906108306104b0836115b8565b6000805160206127b483398151915260009081526040888101602052808220547fad16221efc80aaf1b7e69bd3ecb61ba5ffa539adf129c3b4ffff769c9b5bbc338352912054610fa092900381029190910591508161088e57600191505b61089b82820160016115d0565b6000805160206127b4833981519152600090815260408981016020528082209290925588546000805160206127748339815191528252919020603c429081069003908190556108e8612655565b60005b6005811015610c165760015b6005811015610a9c57600082815260098a0160205260408120826005811061091b57fe5b0154600084815260088c016020526040812091925090836005811061093c57fe5b01546001600160a01b03169050825b60008111801561097a5750600085815260098d016020526040902060001982016005811061097557fe5b015483105b15610a2d57600085815260098d016020526040902060001982016005811061099e57fe5b0154600086815260098e016020526040902082600581106109bb57fe5b0155600085815260088d01602052604090206000198201600581106109dc57fe5b0154600086815260088e01602052604090206001600160a01b03909116908260058110610a0557fe5b0180546001600160a01b0319166001600160a01b03929092169190911790556000190161094b565b83811015610a9157600085815260098d016020526040902083908260058110610a5257fe5b0155600085815260088d016020526040902082908260058110610a7157fe5b0180546001600160a01b0319166001600160a01b03929092169190911790555b5050506001016108f7565b5060008b60480160008b8460058110610ab157fe5b602002015181526020019081526020016000209050886009016000838152602001908152602001600020600580602002604051908101604052809291908260058015610b12576020028201915b815481526020019060010190808311610afe575b5050505050925082600260058110610b2657fe5b6020908102919091015160008681526006840183526040808220929092558481526008808d01845282822088835290850190935220610b66916005612673565b5060008281526009808b0160209081526040808420888552928501909152909120610b929160056126ae565b50600082815260088a0160205260408120610bac916126e5565b600082815260098a0160205260408120610bc5916126e5565b600381018054600181810183556000928352602080842090920187905586835260058401825260408084204390556000805160206127d48339815191528452600490940190915291812055016108eb565b50827fbeb3b9f46c8d7bb00c873fca10d307538df350866d25f891ffb395147ddbdc458984848e60400160007fdb21f0c4accc4f2f5f1045353763a9ffe7091ceaf0fcceb5831858d96cf8463160001b8152602001908152602001600020546040518085600560200280838360005b83811015610c9d578181015183820152602001610c85565b5050505090500184815260200183600560200280838360005b83811015610cce578181015183820152602001610cb6565b5050505090500182815260200194505050505060405180910390a28751600083815260428c016020908152604082209290925560348c01805460018101825590825291902001829055610d1f612655565b8851600090815260488c016020908152604080832086845260080190915290819020815160a08101928390529160059082845b81546001600160a01b03168152600190910190602001808311610d525750505050509050610d818b88836115df565b60008051602061283483398151915260009081526040808d01602052902080546001019055610dae612655565b610db78c611a4f565b905060005b6005811015610ed157818160058110610dd157fe5b60200201518d6035018260058110610de557fe5b6002020155600060018e0160488f0182858560058110610e0157fe5b60200201518152602001908152602001600020600401600060008051602061285483398151915260001b81526020019081526020016000205460338110610e4457fe5b015560488d016000838360058110610e5857fe5b6020020151815260200190815260200160002060040160006000805160206127d483398151915260001b8152602001908152602001600020548d60400160006000805160206128ba83398151915260001b8152602001908152602001600020600082825401925050819055508080600101915050610dbc565b508a8560014303406040516020018080602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b83811015610f25578181015183820152602001610f0d565b50505050905090810190601f168015610f525780820380516001836020036101000a031916815260200191505b50945050505050604051602081830303815290604052805190602001209450848c60000181905550847f1d85ce10456e29b67de37887496d3f1fcf1b64c79c4d07484038703a9f5c1408828e60400160006000805160206127b483398151915260001b8152602001908152602001600020548f60400160006000805160206128ba83398151915260001b8152602001908152602001600020546040518084600560200280838360005b83811015611013578181015183820152602001610ffb565b5050505091909101938452505060208201526040805191829003019150a2505050505050505050505050565b6000805160206128148339815191526000908152603f820160205260409020546001600160a01b031633146110bb576040805162461bcd60e51b815260206004820152601b60248201527f53656e646572206973206e6f742070656e64696e67206f776e65720000000000604482015290519081900360640190fd5b6000805160206128148339815191526000908152603f820160205260408082205460008051602061279483398151915283528183205491516001600160a01b039182169392909116917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805160206128148339815191526000908152603f9091016020526040808220546000805160206127948339815191528352912080546001600160a01b0319166001600160a01b03909216919091179055565b7fb1557182e4359a1f0c6301278e8f5b35a776ab58d39892581e357578fb28783681565b6000805160206127d483398151915281565b7f7584d7d8701714da9c117f5bf30af73b0b88aca5338a84a21eb28de2fe0d93b881565b7f05de9147d05477c0a5dc675aeea733157f5092f82add148cf39d579cafe3dc9881565b60008051602061283483398151915281565b7fad16221efc80aaf1b7e69bd3ecb61ba5ffa539adf129c3b4ffff769c9b5bbc3381565b60008051602061277483398151915281565b6000805160206128ba83398151915281565b60008051602061281483398151915281565b6112a58585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611b3c92505050565b61132f8585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040805160a081810190925292508791506005908390839080828437600092019190915250506040805160a0818101909252915086906005908390839080828437600092019190915250611dd2915050565b5050505050565b7f8fe9ded8d7c08f720cf0340699024f83522ea66b2bbfb8f557851cb9ee63b54c81565b7fdb21f0c4accc4f2f5f1045353763a9ffe7091ceaf0fcceb5831858d96cf8463181565b60008051602061279483398151915281565b600082815260488401602090815260408083206000805160206127d4833981519152845260048101909252909120546113cf908363ffffffff6124a316565b6000805160206127d4833981519152600090815260048301602052604090205560358401548314806114045750603784015483145b806114125750603984015483145b806114205750603b84015483145b8061142e5750603d84015483145b1561145c576000805160206128ba8339815191526000908152604080860160205290208054830190556115b2565b6000805160206128548339815191526000908152600482016020526040902054611578576040805161066081019182905260009182916114be91600189019060339082845b8154815260200190600101908083116114a15750505050506124b2565b6000805160206127d4833981519152600090815260048601602052604090205491935091508210806114ee575081155b15611571576000805160206127d4833981519152600090815260048401602052604090205460018701826033811061152257fe5b0155600081815260438701602090815260408083208054845260488a01835281842060008051602061285483398151915285526004908101845282852085905590899055860190915290208190555b50506115b2565b600080516020612854833981519152600090815260048201602052604090205482906001860190603381106115a957fe5b01805490910190555b50505050565b60008183106115c757816115c9565b825b9392505050565b60008183136115c757816115c9565b6000805160206128ba83398151915260009081526040848101602052808220548351825163c7bb46ad60e01b8152600481018890523060248201526001600160a01b039091166044820152600a90910461012c670de0b6b3a7640000878102919091048281016064850152935190949192600285049273__$0a7adcfe26a87eb516915da32ab69f1b09$__9263c7bb46ad926084808201939291829003018186803b15801561168d57600080fd5b505af41580156116a1573d6000803e3d6000fd5b5050505073__$0a7adcfe26a87eb516915da32ab69f1b09$__63c7bb46ad8830886001600581106116ce57fe5b6020020151604080516001600160e01b031960e087901b16815260048101949094526001600160a01b039283166024850152911660448301528686016064830152516084808301926000929190829003018186803b15801561172f57600080fd5b505af4158015611743573d6000803e3d6000fd5b5050505073__$0a7adcfe26a87eb516915da32ab69f1b09$__63c7bb46ad88308860026005811061177057fe5b6020020151604080516001600160e01b031960e087901b16815260048101949094526001600160a01b039283166024850152911660448301528686016064830152516084808301926000929190829003018186803b1580156117d157600080fd5b505af41580156117e5573d6000803e3d6000fd5b5050505073__$0a7adcfe26a87eb516915da32ab69f1b09$__63c7bb46ad88308860036005811061181257fe5b6020020151604080516001600160e01b031960e087901b16815260048101949094526001600160a01b039283166024850152911660448301528686016064830152516084808301926000929190829003018186803b15801561187357600080fd5b505af4158015611887573d6000803e3d6000fd5b5050505073__$0a7adcfe26a87eb516915da32ab69f1b09$__63c7bb46ad8830886004600581106118b457fe5b6020020151604080516001600160e01b031960e087901b16815260048101949094526001600160a01b039283166024850152911660448301528686016064830152516084808301926000929190829003018186803b15801561191557600080fd5b505af4158015611929573d6000803e3d6000fd5b50506000805160206128ba833981519152600090815260408a81016020908152818320547fb1557182e4359a1f0c6301278e8f5b35a776ab58d39892581e357578fb28783684528284208054600290920460058b0289010390910190556000805160206127948339815191528352603f8c01905280822054815163c7bb46ad60e01b8152600481018d90523060248201526001600160a01b03909116604482015260648101869052905173__$0a7adcfe26a87eb516915da32ab69f1b09$__945063c7bb46ad935060848083019392829003018186803b158015611a0c57600080fd5b505af4158015611a20573d6000803e3d6000fd5b50506000805160206128ba83398151915260009081526040998a01602052988920989098555050505050505050565b611a57612655565b611a5f612655565b611a67612655565b60408051610660810191829052611aa091600187019060339082845b815481526020019060010190808311611a83575050505050612501565b909250905060005b6005811015611b3457828160058110611abd57fe5b602002015115611b0357846043016000838360058110611ad957fe5b6020020151815260200190815260200160002054848260058110611af957fe5b6020020152611b2c565b846035018160040360058110611b1557fe5b6002020154848260058110611b2657fe5b60200201525b600101611aa8565b505050919050565b8160400160006000805160206127b483398151915260001b815260200190815260200160002054600260038460000154338560405160200180848152602001836001600160a01b03166001600160a01b031660601b815260140182805190602001908083835b60208310611bc15780518252601f199092019160209182019101611ba2565b6001836020036101000a038019825116818451168082178552505050505050905001935050505060405160208183030381529060405280519060200120604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b60208310611c4c5780518252601f199092019160209182019101611c2d565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015611c8b573d6000803e3d6000fd5b5050506040515160601b60405160200180826bffffffffffffffffffffffff19166bffffffffffffffffffffffff191681526014019150506040516020818303038152906040526040518082805190602001908083835b60208310611d015780518252601f199092019160209182019101611ce2565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015611d40573d6000803e3d6000fd5b5050506040513d6020811015611d5557600080fd5b505181611d5e57fe5b061580611d9357506000805160206127748339815191526000908152604080840160205290205461038490603c420642030310155b611dce5760405162461bcd60e51b81526004018080602001828103825260258152602001806128956025913960400191505060405180910390fd5b5050565b6040805133602080830182905283518084038201815292840184528251928101929092206000918252604788019092529190912054600114611e5b576040805162461bcd60e51b815260206004820152601a60248201527f4d696e657220737461747573206973206e6f74207374616b6572000000000000604482015290519081900360640190fd5b60008181526040808701602052902054610384429190910311611eaf5760405162461bcd60e51b815260040180806020018281038252602a81526020018061274a602a913960400191505060405180910390fd5b6035850154835114611efe576040805162461bcd60e51b8152602060048201526013602482015272526571756573742049442069732077726f6e6760681b604482015290519081900360640190fd5b6037850154602084015114611f50576040805162461bcd60e51b8152602060048201526013602482015272526571756573742049442069732077726f6e6760681b604482015290519081900360640190fd5b6039850154604084015114611fa2576040805162461bcd60e51b8152602060048201526013602482015272526571756573742049442069732077726f6e6760681b604482015290519081900360640190fd5b603b850154606084015114611ff4576040805162461bcd60e51b8152602060048201526013602482015272526571756573742049442069732077726f6e6760681b604482015290519081900360640190fd5b603d850154608084015114612046576040805162461bcd60e51b8152602060048201526013602482015272526571756573742049442069732077726f6e6760681b604482015290519081900360640190fd5b60008181526040808701602090815281832042905587546000805160206127f483398151915284528284205481855260418a018352838520338652909252919092205490919060ff16156120cb5760405162461bcd60e51b81526004018080602001828103825260218152602001806128746021913960400191505060405180910390fd5b6000828152604188016020908152604080832033808552908352818420805460ff191660019081179091556000805160206128348339815191528552828c01845282852054855260488c018452828520908552600881019093529220909190836005811061213557fe5b0180546001600160a01b0319166001600160a01b0392909216919091179055845160008080526009830160205260409020836005811061217157fe5b015560208086015160016000908152600984019092526040909120836005811061219757fe5b015560408086015160026000908152600984016020529190912083600581106121bc57fe5b0155606085015160036000908152600983016020526040902083600581106121e057fe5b01556080850151600460009081526009830160205260409020836005811061220457fe5b0155600080805260088201602052604090203390836005811061222357fe5b0180546001600160a01b0319166001600160a01b03929092169190911790556001600090815260088201602052604090203390836005811061226157fe5b0180546001600160a01b0319166001600160a01b03929092169190911790556002600090815260088201602052604090203390836005811061229f57fe5b0180546001600160a01b0319166001600160a01b0392909216919091179055600360009081526008820160205260409020339083600581106122dd57fe5b0180546001600160a01b0319166001600160a01b03929092169190911790556004600090815260088201602052604090203390836005811061231b57fe5b0180546001600160a01b0319166001600160a01b03929092169190911790556001820160051415612375576123518888886107d7565b6000805160206127f483398151915260009081526040808a0160205281205561239b565b6000805160206127f483398151915260009081526040808a016020529020805460010190555b82336001600160a01b03167f0e4e65dc389613b6884b7f8c615e54fd3b894fbbbc534c990037744eea942000898989604051808060200184600560200280838360005b838110156123f65781810151838201526020016123de565b5050505090500183600560200280838360005b83811015612421578181015183820152602001612409565b50505050905001828103825285818151815260200191508051906020019080838360005b8381101561245d578181015183820152602001612445565b50505050905090810190601f16801561248a5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a35050505050505050565b6000828201838110156115c957fe5b610640810151603260315b80156124fb57828482603381106124d057fe5b602002015110156124f2578381603381106124e757fe5b602002015192508091505b600019016124bd565b50915091565b612509612655565b612511612655565b60208301516000805b60058110156125945785816001016033811061253257fe5b602002015185826005811061254357fe5b60200201526001810184826005811061255857fe5b60200201528285826005811061256a57fe5b6020020151101561258c5784816005811061258157fe5b602002015192508091505b60010161251a565b5060065b603381101561264d57828682603381106125ae57fe5b60200201511115612645578581603381106125c557fe5b60200201518583600581106125d657fe5b6020020152808483600581106125e857fe5b60200201528581603381106125f957fe5b6020020151925060005b6005811015612643578386826005811061261957fe5b6020020151101561263b5785816005811061263057fe5b602002015193508092505b600101612603565b505b600101612598565b505050915091565b6040518060a001604052806005906020820280388339509192915050565b826005810192821561269e579182015b8281111561269e578254825591600101919060010190612683565b506126aa929150612708565b5090565b82600581019282156126d9579182015b828111156126d95782548255916001019190600101906126be565b506126aa92915061272f565b506000815560010160008155600101600081556001016000815560010160009055565b61272c91905b808211156126aa5780546001600160a01b031916815560010161270e565b90565b61272c91905b808211156126aa576000815560010161273556fe4d696e65722063616e206f6e6c792077696e2072657761726473206f6e636520706572203135206d696e97e6eb29f6a85471f7cc9b57f9e4c3deaf398cfc9798673160d7798baf0b13a49dbc393ddc18fd27b1d9b1b129059925688d2f2d5818a5ec3ebb750b7c286ea6b12aff7664b16cb99339be399b863feecd64d14817be7e1f042f97e3f358e64e2a9e355a92978430eca9c1aa3a9ba590094bac282594bccf82de16b83046e2c36c505cb2db6644f57b42d87bd9407b0f66788b07d0617a2bc1356a0e69e66f9a44b2657a0f8a90ed8e62f4c4cceca06eacaa9b4b25751ae1ebca9280a70abd68969ea04b74d02bb4d9e6e8e57236e1b9ca31627139ae9f0e465249932e8245021e344bd070f05f1c5b3f0b1266f4f20d837a0a8190a3a2da8b0375eac2ba86ea4d696e657220616c7265616479207375626d6974746564207468652076616c7565496e636f7272656374206e6f6e636520666f722063757272656e74206368616c6c656e6765d26d9834adf5a73309c4974bf654850bb699df8505e70d4cfde365c417b19dfca265627a7a72315820bbf3dc5f9b33468779423e8cbd34e37df5a5ce6a425444cc19f97be21c6fb6df64736f6c63430005100032"

// DeployTellorLibrary deploys a new Ethereum contract, binding an instance of TellorLibrary to it.
func DeployTellorLibrary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorLibrary, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorLibraryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	tellorTransferAddr, _, _, _ := DeployTellorTransfer(auth, backend)
	TellorLibraryBin = strings.Replace(TellorLibraryBin, "__$0a7adcfe26a87eb516915da32ab69f1b09$__", tellorTransferAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TellorLibraryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TellorLibrary{TellorLibraryCaller: TellorLibraryCaller{contract: contract}, TellorLibraryTransactor: TellorLibraryTransactor{contract: contract}, TellorLibraryFilterer: TellorLibraryFilterer{contract: contract}}, nil
}

// TellorLibrary is an auto generated Go binding around an Ethereum contract.
type TellorLibrary struct {
	TellorLibraryCaller     // Read-only binding to the contract
	TellorLibraryTransactor // Write-only binding to the contract
	TellorLibraryFilterer   // Log filterer for contract events
}

// TellorLibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorLibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorLibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorLibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorLibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorLibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorLibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorLibrarySession struct {
	Contract     *TellorLibrary    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorLibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorLibraryCallerSession struct {
	Contract *TellorLibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TellorLibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorLibraryTransactorSession struct {
	Contract     *TellorLibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TellorLibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorLibraryRaw struct {
	Contract *TellorLibrary // Generic contract binding to access the raw methods on
}

// TellorLibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorLibraryCallerRaw struct {
	Contract *TellorLibraryCaller // Generic read-only contract binding to access the raw methods on
}

// TellorLibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorLibraryTransactorRaw struct {
	Contract *TellorLibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellorLibrary creates a new instance of TellorLibrary, bound to a specific deployed contract.
func NewTellorLibrary(address common.Address, backend bind.ContractBackend) (*TellorLibrary, error) {
	contract, err := bindTellorLibrary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TellorLibrary{TellorLibraryCaller: TellorLibraryCaller{contract: contract}, TellorLibraryTransactor: TellorLibraryTransactor{contract: contract}, TellorLibraryFilterer: TellorLibraryFilterer{contract: contract}}, nil
}

// NewTellorLibraryCaller creates a new read-only instance of TellorLibrary, bound to a specific deployed contract.
func NewTellorLibraryCaller(address common.Address, caller bind.ContractCaller) (*TellorLibraryCaller, error) {
	contract, err := bindTellorLibrary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorLibraryCaller{contract: contract}, nil
}

// NewTellorLibraryTransactor creates a new write-only instance of TellorLibrary, bound to a specific deployed contract.
func NewTellorLibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorLibraryTransactor, error) {
	contract, err := bindTellorLibrary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorLibraryTransactor{contract: contract}, nil
}

// NewTellorLibraryFilterer creates a new log filterer instance of TellorLibrary, bound to a specific deployed contract.
func NewTellorLibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorLibraryFilterer, error) {
	contract, err := bindTellorLibrary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorLibraryFilterer{contract: contract}, nil
}

// bindTellorLibrary binds a generic wrapper to an already deployed contract.
func bindTellorLibrary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorLibraryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorLibrary *TellorLibraryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorLibrary.Contract.TellorLibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorLibrary *TellorLibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorLibrary.Contract.TellorLibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorLibrary *TellorLibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorLibrary.Contract.TellorLibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorLibrary *TellorLibraryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorLibrary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorLibrary *TellorLibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorLibrary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorLibrary *TellorLibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorLibrary.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCaller) Owner(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorLibrary.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(bytes32)
func (_TellorLibrary *TellorLibrarySession) Owner() ([32]byte, error) {
	return _TellorLibrary.Contract.Owner(&_TellorLibrary.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCallerSession) Owner() ([32]byte, error) {
	return _TellorLibrary.Contract.Owner(&_TellorLibrary.CallOpts)
}

// TBlock is a free data retrieval call binding the contract method 0x6e3cf885.
//
// Solidity: function _tBlock() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCaller) TBlock(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorLibrary.contract.Call(opts, &out, "_tBlock")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TBlock is a free data retrieval call binding the contract method 0x6e3cf885.
//
// Solidity: function _tBlock() view returns(bytes32)
func (_TellorLibrary *TellorLibrarySession) TBlock() ([32]byte, error) {
	return _TellorLibrary.Contract.TBlock(&_TellorLibrary.CallOpts)
}

// TBlock is a free data retrieval call binding the contract method 0x6e3cf885.
//
// Solidity: function _tBlock() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCallerSession) TBlock() ([32]byte, error) {
	return _TellorLibrary.Contract.TBlock(&_TellorLibrary.CallOpts)
}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5ae2bfdb.
//
// Solidity: function currentRequestId() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCaller) CurrentRequestId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorLibrary.contract.Call(opts, &out, "currentRequestId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5ae2bfdb.
//
// Solidity: function currentRequestId() view returns(bytes32)
func (_TellorLibrary *TellorLibrarySession) CurrentRequestId() ([32]byte, error) {
	return _TellorLibrary.Contract.CurrentRequestId(&_TellorLibrary.CallOpts)
}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5ae2bfdb.
//
// Solidity: function currentRequestId() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCallerSession) CurrentRequestId() ([32]byte, error) {
	return _TellorLibrary.Contract.CurrentRequestId(&_TellorLibrary.CallOpts)
}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCaller) CurrentReward(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorLibrary.contract.Call(opts, &out, "currentReward")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(bytes32)
func (_TellorLibrary *TellorLibrarySession) CurrentReward() ([32]byte, error) {
	return _TellorLibrary.Contract.CurrentReward(&_TellorLibrary.CallOpts)
}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCallerSession) CurrentReward() ([32]byte, error) {
	return _TellorLibrary.Contract.CurrentReward(&_TellorLibrary.CallOpts)
}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCaller) CurrentTotalTips(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorLibrary.contract.Call(opts, &out, "currentTotalTips")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(bytes32)
func (_TellorLibrary *TellorLibrarySession) CurrentTotalTips() ([32]byte, error) {
	return _TellorLibrary.Contract.CurrentTotalTips(&_TellorLibrary.CallOpts)
}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCallerSession) CurrentTotalTips() ([32]byte, error) {
	return _TellorLibrary.Contract.CurrentTotalTips(&_TellorLibrary.CallOpts)
}

// DevShare is a free data retrieval call binding the contract method 0xaed04fae.
//
// Solidity: function devShare() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCaller) DevShare(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorLibrary.contract.Call(opts, &out, "devShare")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DevShare is a free data retrieval call binding the contract method 0xaed04fae.
//
// Solidity: function devShare() view returns(bytes32)
func (_TellorLibrary *TellorLibrarySession) DevShare() ([32]byte, error) {
	return _TellorLibrary.Contract.DevShare(&_TellorLibrary.CallOpts)
}

// DevShare is a free data retrieval call binding the contract method 0xaed04fae.
//
// Solidity: function devShare() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCallerSession) DevShare() ([32]byte, error) {
	return _TellorLibrary.Contract.DevShare(&_TellorLibrary.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCaller) Difficulty(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorLibrary.contract.Call(opts, &out, "difficulty")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(bytes32)
func (_TellorLibrary *TellorLibrarySession) Difficulty() ([32]byte, error) {
	return _TellorLibrary.Contract.Difficulty(&_TellorLibrary.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCallerSession) Difficulty() ([32]byte, error) {
	return _TellorLibrary.Contract.Difficulty(&_TellorLibrary.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCaller) PendingOwner(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorLibrary.contract.Call(opts, &out, "pending_owner")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(bytes32)
func (_TellorLibrary *TellorLibrarySession) PendingOwner() ([32]byte, error) {
	return _TellorLibrary.Contract.PendingOwner(&_TellorLibrary.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCallerSession) PendingOwner() ([32]byte, error) {
	return _TellorLibrary.Contract.PendingOwner(&_TellorLibrary.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCaller) RequestCount(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorLibrary.contract.Call(opts, &out, "requestCount")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(bytes32)
func (_TellorLibrary *TellorLibrarySession) RequestCount() ([32]byte, error) {
	return _TellorLibrary.Contract.RequestCount(&_TellorLibrary.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCallerSession) RequestCount() ([32]byte, error) {
	return _TellorLibrary.Contract.RequestCount(&_TellorLibrary.CallOpts)
}

// RequestQPosition is a free data retrieval call binding the contract method 0x2bf07e9e.
//
// Solidity: function requestQPosition() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCaller) RequestQPosition(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorLibrary.contract.Call(opts, &out, "requestQPosition")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RequestQPosition is a free data retrieval call binding the contract method 0x2bf07e9e.
//
// Solidity: function requestQPosition() view returns(bytes32)
func (_TellorLibrary *TellorLibrarySession) RequestQPosition() ([32]byte, error) {
	return _TellorLibrary.Contract.RequestQPosition(&_TellorLibrary.CallOpts)
}

// RequestQPosition is a free data retrieval call binding the contract method 0x2bf07e9e.
//
// Solidity: function requestQPosition() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCallerSession) RequestQPosition() ([32]byte, error) {
	return _TellorLibrary.Contract.RequestQPosition(&_TellorLibrary.CallOpts)
}

// RunningTips is a free data retrieval call binding the contract method 0xb0dc7c20.
//
// Solidity: function runningTips() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCaller) RunningTips(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorLibrary.contract.Call(opts, &out, "runningTips")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RunningTips is a free data retrieval call binding the contract method 0xb0dc7c20.
//
// Solidity: function runningTips() view returns(bytes32)
func (_TellorLibrary *TellorLibrarySession) RunningTips() ([32]byte, error) {
	return _TellorLibrary.Contract.RunningTips(&_TellorLibrary.CallOpts)
}

// RunningTips is a free data retrieval call binding the contract method 0xb0dc7c20.
//
// Solidity: function runningTips() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCallerSession) RunningTips() ([32]byte, error) {
	return _TellorLibrary.Contract.RunningTips(&_TellorLibrary.CallOpts)
}

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCaller) SlotProgress(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorLibrary.contract.Call(opts, &out, "slotProgress")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(bytes32)
func (_TellorLibrary *TellorLibrarySession) SlotProgress() ([32]byte, error) {
	return _TellorLibrary.Contract.SlotProgress(&_TellorLibrary.CallOpts)
}

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCallerSession) SlotProgress() ([32]byte, error) {
	return _TellorLibrary.Contract.SlotProgress(&_TellorLibrary.CallOpts)
}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCaller) TimeOfLastNewValue(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorLibrary.contract.Call(opts, &out, "timeOfLastNewValue")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(bytes32)
func (_TellorLibrary *TellorLibrarySession) TimeOfLastNewValue() ([32]byte, error) {
	return _TellorLibrary.Contract.TimeOfLastNewValue(&_TellorLibrary.CallOpts)
}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCallerSession) TimeOfLastNewValue() ([32]byte, error) {
	return _TellorLibrary.Contract.TimeOfLastNewValue(&_TellorLibrary.CallOpts)
}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCaller) TimeTarget(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorLibrary.contract.Call(opts, &out, "timeTarget")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(bytes32)
func (_TellorLibrary *TellorLibrarySession) TimeTarget() ([32]byte, error) {
	return _TellorLibrary.Contract.TimeTarget(&_TellorLibrary.CallOpts)
}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCallerSession) TimeTarget() ([32]byte, error) {
	return _TellorLibrary.Contract.TimeTarget(&_TellorLibrary.CallOpts)
}

// TotalTip is a free data retrieval call binding the contract method 0x561cb04a.
//
// Solidity: function totalTip() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCaller) TotalTip(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorLibrary.contract.Call(opts, &out, "totalTip")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TotalTip is a free data retrieval call binding the contract method 0x561cb04a.
//
// Solidity: function totalTip() view returns(bytes32)
func (_TellorLibrary *TellorLibrarySession) TotalTip() ([32]byte, error) {
	return _TellorLibrary.Contract.TotalTip(&_TellorLibrary.CallOpts)
}

// TotalTip is a free data retrieval call binding the contract method 0x561cb04a.
//
// Solidity: function totalTip() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCallerSession) TotalTip() ([32]byte, error) {
	return _TellorLibrary.Contract.TotalTip(&_TellorLibrary.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x3940e9ee.
//
// Solidity: function total_supply() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCaller) TotalSupply(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorLibrary.contract.Call(opts, &out, "total_supply")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x3940e9ee.
//
// Solidity: function total_supply() view returns(bytes32)
func (_TellorLibrary *TellorLibrarySession) TotalSupply() ([32]byte, error) {
	return _TellorLibrary.Contract.TotalSupply(&_TellorLibrary.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x3940e9ee.
//
// Solidity: function total_supply() view returns(bytes32)
func (_TellorLibrary *TellorLibraryCallerSession) TotalSupply() ([32]byte, error) {
	return _TellorLibrary.Contract.TotalSupply(&_TellorLibrary.CallOpts)
}

// TellorLibraryNewChallengeIterator is returned from FilterNewChallenge and is used to iterate over the raw logs and unpacked data for NewChallenge events raised by the TellorLibrary contract.
type TellorLibraryNewChallengeIterator struct {
	Event *TellorLibraryNewChallenge // Event containing the contract specifics and raw log

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
func (it *TellorLibraryNewChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorLibraryNewChallenge)
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
		it.Event = new(TellorLibraryNewChallenge)
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
func (it *TellorLibraryNewChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorLibraryNewChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorLibraryNewChallenge represents a NewChallenge event raised by the TellorLibrary contract.
type TellorLibraryNewChallenge struct {
	CurrentChallenge [32]byte
	CurrentRequestId [5]*big.Int
	Difficulty       *big.Int
	TotalTips        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewChallenge is a free log retrieval operation binding the contract event 0x1d85ce10456e29b67de37887496d3f1fcf1b64c79c4d07484038703a9f5c1408.
//
// Solidity: event NewChallenge(bytes32 indexed _currentChallenge, uint256[5] _currentRequestId, uint256 _difficulty, uint256 _totalTips)
func (_TellorLibrary *TellorLibraryFilterer) FilterNewChallenge(opts *bind.FilterOpts, _currentChallenge [][32]byte) (*TellorLibraryNewChallengeIterator, error) {

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _TellorLibrary.contract.FilterLogs(opts, "NewChallenge", _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return &TellorLibraryNewChallengeIterator{contract: _TellorLibrary.contract, event: "NewChallenge", logs: logs, sub: sub}, nil
}

// WatchNewChallenge is a free log subscription operation binding the contract event 0x1d85ce10456e29b67de37887496d3f1fcf1b64c79c4d07484038703a9f5c1408.
//
// Solidity: event NewChallenge(bytes32 indexed _currentChallenge, uint256[5] _currentRequestId, uint256 _difficulty, uint256 _totalTips)
func (_TellorLibrary *TellorLibraryFilterer) WatchNewChallenge(opts *bind.WatchOpts, sink chan<- *TellorLibraryNewChallenge, _currentChallenge [][32]byte) (event.Subscription, error) {

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _TellorLibrary.contract.WatchLogs(opts, "NewChallenge", _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorLibraryNewChallenge)
				if err := _TellorLibrary.contract.UnpackLog(event, "NewChallenge", log); err != nil {
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

// ParseNewChallenge is a log parse operation binding the contract event 0x1d85ce10456e29b67de37887496d3f1fcf1b64c79c4d07484038703a9f5c1408.
//
// Solidity: event NewChallenge(bytes32 indexed _currentChallenge, uint256[5] _currentRequestId, uint256 _difficulty, uint256 _totalTips)
func (_TellorLibrary *TellorLibraryFilterer) ParseNewChallenge(log types.Log) (*TellorLibraryNewChallenge, error) {
	event := new(TellorLibraryNewChallenge)
	if err := _TellorLibrary.contract.UnpackLog(event, "NewChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorLibraryNewValueIterator is returned from FilterNewValue and is used to iterate over the raw logs and unpacked data for NewValue events raised by the TellorLibrary contract.
type TellorLibraryNewValueIterator struct {
	Event *TellorLibraryNewValue // Event containing the contract specifics and raw log

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
func (it *TellorLibraryNewValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorLibraryNewValue)
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
		it.Event = new(TellorLibraryNewValue)
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
func (it *TellorLibraryNewValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorLibraryNewValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorLibraryNewValue represents a NewValue event raised by the TellorLibrary contract.
type TellorLibraryNewValue struct {
	RequestId        [5]*big.Int
	Time             *big.Int
	Value            [5]*big.Int
	TotalTips        *big.Int
	CurrentChallenge [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewValue is a free log retrieval operation binding the contract event 0xbeb3b9f46c8d7bb00c873fca10d307538df350866d25f891ffb395147ddbdc45.
//
// Solidity: event NewValue(uint256[5] _requestId, uint256 _time, uint256[5] _value, uint256 _totalTips, bytes32 indexed _currentChallenge)
func (_TellorLibrary *TellorLibraryFilterer) FilterNewValue(opts *bind.FilterOpts, _currentChallenge [][32]byte) (*TellorLibraryNewValueIterator, error) {

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _TellorLibrary.contract.FilterLogs(opts, "NewValue", _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return &TellorLibraryNewValueIterator{contract: _TellorLibrary.contract, event: "NewValue", logs: logs, sub: sub}, nil
}

// WatchNewValue is a free log subscription operation binding the contract event 0xbeb3b9f46c8d7bb00c873fca10d307538df350866d25f891ffb395147ddbdc45.
//
// Solidity: event NewValue(uint256[5] _requestId, uint256 _time, uint256[5] _value, uint256 _totalTips, bytes32 indexed _currentChallenge)
func (_TellorLibrary *TellorLibraryFilterer) WatchNewValue(opts *bind.WatchOpts, sink chan<- *TellorLibraryNewValue, _currentChallenge [][32]byte) (event.Subscription, error) {

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _TellorLibrary.contract.WatchLogs(opts, "NewValue", _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorLibraryNewValue)
				if err := _TellorLibrary.contract.UnpackLog(event, "NewValue", log); err != nil {
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

// ParseNewValue is a log parse operation binding the contract event 0xbeb3b9f46c8d7bb00c873fca10d307538df350866d25f891ffb395147ddbdc45.
//
// Solidity: event NewValue(uint256[5] _requestId, uint256 _time, uint256[5] _value, uint256 _totalTips, bytes32 indexed _currentChallenge)
func (_TellorLibrary *TellorLibraryFilterer) ParseNewValue(log types.Log) (*TellorLibraryNewValue, error) {
	event := new(TellorLibraryNewValue)
	if err := _TellorLibrary.contract.UnpackLog(event, "NewValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorLibraryNonceSubmittedIterator is returned from FilterNonceSubmitted and is used to iterate over the raw logs and unpacked data for NonceSubmitted events raised by the TellorLibrary contract.
type TellorLibraryNonceSubmittedIterator struct {
	Event *TellorLibraryNonceSubmitted // Event containing the contract specifics and raw log

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
func (it *TellorLibraryNonceSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorLibraryNonceSubmitted)
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
		it.Event = new(TellorLibraryNonceSubmitted)
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
func (it *TellorLibraryNonceSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorLibraryNonceSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorLibraryNonceSubmitted represents a NonceSubmitted event raised by the TellorLibrary contract.
type TellorLibraryNonceSubmitted struct {
	Miner            common.Address
	Nonce            string
	RequestId        [5]*big.Int
	Value            [5]*big.Int
	CurrentChallenge [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNonceSubmitted is a free log retrieval operation binding the contract event 0x0e4e65dc389613b6884b7f8c615e54fd3b894fbbbc534c990037744eea942000.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256[5] _requestId, uint256[5] _value, bytes32 indexed _currentChallenge)
func (_TellorLibrary *TellorLibraryFilterer) FilterNonceSubmitted(opts *bind.FilterOpts, _miner []common.Address, _currentChallenge [][32]byte) (*TellorLibraryNonceSubmittedIterator, error) {

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _TellorLibrary.contract.FilterLogs(opts, "NonceSubmitted", _minerRule, _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return &TellorLibraryNonceSubmittedIterator{contract: _TellorLibrary.contract, event: "NonceSubmitted", logs: logs, sub: sub}, nil
}

// WatchNonceSubmitted is a free log subscription operation binding the contract event 0x0e4e65dc389613b6884b7f8c615e54fd3b894fbbbc534c990037744eea942000.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256[5] _requestId, uint256[5] _value, bytes32 indexed _currentChallenge)
func (_TellorLibrary *TellorLibraryFilterer) WatchNonceSubmitted(opts *bind.WatchOpts, sink chan<- *TellorLibraryNonceSubmitted, _miner []common.Address, _currentChallenge [][32]byte) (event.Subscription, error) {

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _TellorLibrary.contract.WatchLogs(opts, "NonceSubmitted", _minerRule, _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorLibraryNonceSubmitted)
				if err := _TellorLibrary.contract.UnpackLog(event, "NonceSubmitted", log); err != nil {
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

// ParseNonceSubmitted is a log parse operation binding the contract event 0x0e4e65dc389613b6884b7f8c615e54fd3b894fbbbc534c990037744eea942000.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256[5] _requestId, uint256[5] _value, bytes32 indexed _currentChallenge)
func (_TellorLibrary *TellorLibraryFilterer) ParseNonceSubmitted(log types.Log) (*TellorLibraryNonceSubmitted, error) {
	event := new(TellorLibraryNonceSubmitted)
	if err := _TellorLibrary.contract.UnpackLog(event, "NonceSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorLibraryOwnershipProposedIterator is returned from FilterOwnershipProposed and is used to iterate over the raw logs and unpacked data for OwnershipProposed events raised by the TellorLibrary contract.
type TellorLibraryOwnershipProposedIterator struct {
	Event *TellorLibraryOwnershipProposed // Event containing the contract specifics and raw log

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
func (it *TellorLibraryOwnershipProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorLibraryOwnershipProposed)
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
		it.Event = new(TellorLibraryOwnershipProposed)
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
func (it *TellorLibraryOwnershipProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorLibraryOwnershipProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorLibraryOwnershipProposed represents a OwnershipProposed event raised by the TellorLibrary contract.
type TellorLibraryOwnershipProposed struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipProposed is a free log retrieval operation binding the contract event 0xb51454ce8c7f26becd312a46c4815553887f2ec876a0b8dc813b87f62edf6f80.
//
// Solidity: event OwnershipProposed(address indexed _previousOwner, address indexed _newOwner)
func (_TellorLibrary *TellorLibraryFilterer) FilterOwnershipProposed(opts *bind.FilterOpts, _previousOwner []common.Address, _newOwner []common.Address) (*TellorLibraryOwnershipProposedIterator, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _TellorLibrary.contract.FilterLogs(opts, "OwnershipProposed", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TellorLibraryOwnershipProposedIterator{contract: _TellorLibrary.contract, event: "OwnershipProposed", logs: logs, sub: sub}, nil
}

// WatchOwnershipProposed is a free log subscription operation binding the contract event 0xb51454ce8c7f26becd312a46c4815553887f2ec876a0b8dc813b87f62edf6f80.
//
// Solidity: event OwnershipProposed(address indexed _previousOwner, address indexed _newOwner)
func (_TellorLibrary *TellorLibraryFilterer) WatchOwnershipProposed(opts *bind.WatchOpts, sink chan<- *TellorLibraryOwnershipProposed, _previousOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _TellorLibrary.contract.WatchLogs(opts, "OwnershipProposed", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorLibraryOwnershipProposed)
				if err := _TellorLibrary.contract.UnpackLog(event, "OwnershipProposed", log); err != nil {
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

// ParseOwnershipProposed is a log parse operation binding the contract event 0xb51454ce8c7f26becd312a46c4815553887f2ec876a0b8dc813b87f62edf6f80.
//
// Solidity: event OwnershipProposed(address indexed _previousOwner, address indexed _newOwner)
func (_TellorLibrary *TellorLibraryFilterer) ParseOwnershipProposed(log types.Log) (*TellorLibraryOwnershipProposed, error) {
	event := new(TellorLibraryOwnershipProposed)
	if err := _TellorLibrary.contract.UnpackLog(event, "OwnershipProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorLibraryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TellorLibrary contract.
type TellorLibraryOwnershipTransferredIterator struct {
	Event *TellorLibraryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TellorLibraryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorLibraryOwnershipTransferred)
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
		it.Event = new(TellorLibraryOwnershipTransferred)
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
func (it *TellorLibraryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorLibraryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorLibraryOwnershipTransferred represents a OwnershipTransferred event raised by the TellorLibrary contract.
type TellorLibraryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner)
func (_TellorLibrary *TellorLibraryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _previousOwner []common.Address, _newOwner []common.Address) (*TellorLibraryOwnershipTransferredIterator, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _TellorLibrary.contract.FilterLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TellorLibraryOwnershipTransferredIterator{contract: _TellorLibrary.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner)
func (_TellorLibrary *TellorLibraryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TellorLibraryOwnershipTransferred, _previousOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _TellorLibrary.contract.WatchLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorLibraryOwnershipTransferred)
				if err := _TellorLibrary.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner)
func (_TellorLibrary *TellorLibraryFilterer) ParseOwnershipTransferred(log types.Log) (*TellorLibraryOwnershipTransferred, error) {
	event := new(TellorLibraryOwnershipTransferred)
	if err := _TellorLibrary.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorLibraryTipAddedIterator is returned from FilterTipAdded and is used to iterate over the raw logs and unpacked data for TipAdded events raised by the TellorLibrary contract.
type TellorLibraryTipAddedIterator struct {
	Event *TellorLibraryTipAdded // Event containing the contract specifics and raw log

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
func (it *TellorLibraryTipAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorLibraryTipAdded)
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
		it.Event = new(TellorLibraryTipAdded)
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
func (it *TellorLibraryTipAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorLibraryTipAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorLibraryTipAdded represents a TipAdded event raised by the TellorLibrary contract.
type TellorLibraryTipAdded struct {
	Sender    common.Address
	RequestId *big.Int
	Tip       *big.Int
	TotalTips *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTipAdded is a free log retrieval operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_TellorLibrary *TellorLibraryFilterer) FilterTipAdded(opts *bind.FilterOpts, _sender []common.Address, _requestId []*big.Int) (*TellorLibraryTipAddedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _TellorLibrary.contract.FilterLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &TellorLibraryTipAddedIterator{contract: _TellorLibrary.contract, event: "TipAdded", logs: logs, sub: sub}, nil
}

// WatchTipAdded is a free log subscription operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_TellorLibrary *TellorLibraryFilterer) WatchTipAdded(opts *bind.WatchOpts, sink chan<- *TellorLibraryTipAdded, _sender []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _TellorLibrary.contract.WatchLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorLibraryTipAdded)
				if err := _TellorLibrary.contract.UnpackLog(event, "TipAdded", log); err != nil {
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

// ParseTipAdded is a log parse operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_TellorLibrary *TellorLibraryFilterer) ParseTipAdded(log types.Log) (*TellorLibraryTipAdded, error) {
	event := new(TellorLibraryTipAdded)
	if err := _TellorLibrary.contract.UnpackLog(event, "TipAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorStakeABI is the input ABI used to generate the binding from.
const TellorStakeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"NewStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"}]"

// TellorStakeFuncSigs maps the 4-byte function signature to its string representation.
var TellorStakeFuncSigs = map[string]string{
	"820a2d66": "depositStake(TellorStorage.TellorStorageStruct storage)",
	"c9cf5e4c": "requestStakingWithdraw(TellorStorage.TellorStorageStruct storage)",
	"44bacc4b": "withdrawStake(TellorStorage.TellorStorageStruct storage)",
}

// TellorStakeBin is the compiled bytecode used for deploying new contracts.
var TellorStakeBin = "0x6105f6610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c806344bacc4b14610050578063820a2d661461007c578063c9cf5e4c146100a6575b600080fd5b81801561005c57600080fd5b5061007a6004803603602081101561007357600080fd5b50356100d0565b005b81801561008857600080fd5b5061007a6004803603602081101561009f57600080fd5b50356101ad565b8180156100b257600080fd5b5061007a600480360360208110156100c957600080fd5b5035610221565b3360009081526047820160205260409020600181015462093a8090620151804206420303101561013c576040805162461bcd60e51b8152602060048201526012602482015271372064617973206469646e2774207061737360701b604482015290519081900360640190fd5b805460021461017c5760405162461bcd60e51b815260040180806020018281038252602381526020018061057d6023913960400191505060405180910390fd5b600080825560405133917f4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec91a25050565b6101b78133610356565b73__$cd7a2037986a81685ce8a6561121d7deb1$__63d7b651c1826040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561020657600080fd5b505af415801561021a573d6000803e3d6000fd5b5050505050565b3360009081526047820160205260409020805460011461027e576040805162461bcd60e51b8152602060048201526013602482015272135a5b995c881a5cc81b9bdd081cdd185ad959606a1b604482015290519081900360640190fd5b6002815562015180420642036001820155604080516a1cdd185ad95c90dbdd5b9d60aa1b8152815190819003600b0181206000908152828501602052828120805460001901905563d7b651c160e01b825260048201859052915173__$cd7a2037986a81685ce8a6561121d7deb1$__9263d7b651c19260248082019391829003018186803b15801561030f57600080fd5b505af4158015610323573d6000803e3d6000fd5b50506040513392507f453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf9150600090a25050565b604080516a1cdd185ad9505b5bdd5b9d60aa1b8152815190819003600b0181206000908152828501602090815290839020546393b182b360e01b8352600483018690526001600160a01b0385166024840152925173__$0a7adcfe26a87eb516915da32ab69f1b09$__926393b182b3926044808301939192829003018186803b1580156103e257600080fd5b505af41580156103f6573d6000803e3d6000fd5b505050506040513d602081101561040c57600080fd5b5051101561044b5760405162461bcd60e51b81526004018080602001828103825260228152602001806105a06022913960400191505060405180910390fd5b6001600160a01b0381166000908152604783016020526040902054158061048c57506001600160a01b03811660009081526047830160205260409020546002145b6104dd576040805162461bcd60e51b815260206004820152601b60248201527f4d696e657220697320696e207468652077726f6e672073746174650000000000604482015290519081900360640190fd5b604080516a1cdd185ad95c90dbdd5b9d60aa1b8152815190819003600b01812060009081528483016020908152838220805460019081019091558385018552808452620151804290810690038285019081526001600160a01b038716808552604789019093528584209451855551930192909255915190917f46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e291a2505056fe4d696e657220776173206e6f74206c6f636b656420666f72207769746864726177616c42616c616e6365206973206c6f776572207468616e207374616b6520616d6f756e74a265627a7a723158201cd6ad57ae58d84b9068dc1a464252716cbdb91094a23c410c05f1843545273464736f6c63430005100032"

// DeployTellorStake deploys a new Ethereum contract, binding an instance of TellorStake to it.
func DeployTellorStake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorStake, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorStakeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	tellorTransferAddr, _, _, _ := DeployTellorTransfer(auth, backend)
	TellorStakeBin = strings.Replace(TellorStakeBin, "__$0a7adcfe26a87eb516915da32ab69f1b09$__", tellorTransferAddr.String()[2:], -1)

	tellorDisputeAddr, _, _, _ := DeployTellorDispute(auth, backend)
	TellorStakeBin = strings.Replace(TellorStakeBin, "__$cd7a2037986a81685ce8a6561121d7deb1$__", tellorDisputeAddr.String()[2:], -1)

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
	event.Raw = log
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
	event.Raw = log
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
	event.Raw = log
	return event, nil
}

// TellorStorageABI is the input ABI used to generate the binding from.
const TellorStorageABI = "[]"

// TellorStorageBin is the compiled bytecode used for deploying new contracts.
var TellorStorageBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158208700b8abb16aac9c988f0e00897b788bcb4a66443edbcea42f316880531cea9d64736f6c63430005100032"

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
var TellorTransferBin = "0x610adf610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100a85760003560e01c8063bf32006c11610070578063bf32006c146101aa578063c7bb46ad146101de578063c84b96f514610229578063ca50189914610268578063eb7bdc08146102b1576100a8565b80633f48b1ff146100ad57806360c7dc47146100f1578063850dcc32146100f957806393b182b31461014c578063acaab9e214610178575b600080fd5b6100df600480360360608110156100c357600080fd5b508035906001600160a01b0360208201351690604001356102e1565b60408051918252519081900360200190f35b6100df61048e565b81801561010557600080fd5b506101386004803603606081101561011c57600080fd5b508035906001600160a01b0360208201351690604001356104b2565b604080519115158252519081900360200190f35b6100df6004803603604081101561016257600080fd5b50803590602001356001600160a01b03166105f1565b6101386004803603606081101561018e57600080fd5b508035906001600160a01b0360208201351690604001356105fe565b6100df600480360360608110156101c057600080fd5b508035906001600160a01b03602082013581169160400135166106ae565b8180156101ea57600080fd5b506102276004803603608081101561020157600080fd5b508035906001600160a01b036020820135811691604081013590911690606001356106db565b005b81801561023557600080fd5b506101386004803603606081101561024c57600080fd5b508035906001600160a01b0360208201351690604001356108b4565b81801561027457600080fd5b506101386004803603608081101561028b57600080fd5b508035906001600160a01b036020820135811691604081013590911690606001356108cc565b8180156102bd57600080fd5b50610227600480360360408110156102d457600080fd5b508035906020013561097e565b6001600160a01b03821660009081526045840160205260408120805415806103295750828160008154811061031257fe5b6000918252602090912001546001600160801b0316115b15610338576000915050610487565b80548190600019810190811061034a57fe5b6000918252602090912001546001600160801b0316831061039c5780548190600019810190811061037757fe5b600091825260209091200154600160801b90046001600160801b031691506104879050565b8054600090600119015b818111156104545760006002600183850101049050858482815481106103c857fe5b6000918252602090912001546001600160801b03161415610417578381815481106103ef57fe5b600091825260209091200154600160801b90046001600160801b031694506104879350505050565b8584828154811061042457fe5b6000918252602090912001546001600160801b031610156104475780925061044e565b6001810391505b506103a6565b82828154811061046057fe5b600091825260209091200154600160801b90046001600160801b0316935061048792505050565b9392505050565b7f7be108969d31a3f0b261465c71f2b0ba9301cd914d55d9091c3b36a49d4d41b281565b60006001600160a01b038316610506576040805162461bcd60e51b81526020600482015260146024820152735370656e64657220697320302d6164647265737360601b604482015290519081900360640190fd5b33600090815260468501602090815260408083206001600160a01b03871684529091529020541580610536575081155b610587576040805162461bcd60e51b815260206004820152601b60248201527f5370656e64657220697320616c726561647920617070726f7665640000000000604482015290519081900360640190fd5b33600081815260468601602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b60006104878383436102e1565b6001600160a01b03821660009081526047840160205260408120541580159061064157506001600160a01b03831660009081526047850160205260409020546005115b15610699577f7be108969d31a3f0b261465c71f2b0ba9301cd914d55d9091c3b36a49d4d41b260009081526040808601602052902054829061068386866105f1565b031061069157506001610487565b506000610487565b816106a485856105f1565b1015949350505050565b6001600160a01b039182166000908152604693909301602090815260408085209290931684525290205490565b806107175760405162461bcd60e51b8152600401808060200182810382526021815260200180610a8a6021913960400191505060405180910390fd5b6001600160a01b03821661076a576040805162461bcd60e51b815260206004820152601560248201527452656365697665722069732030206164647265737360581b604482015290519081900360640190fd5b6107758484836105fe565b6107b05760405162461bcd60e51b8152600401808060200182810382526027815260200180610a636027913960400191505060405180910390fd5b60006107bc85856105f1565b6001600160a01b038516600090815260458701602052604090209091506107e59083830361097e565b6107ef85846105f1565b905080828201101561083c576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b6001600160a01b038316600090815260458601602052604090206108629082840161097e565b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a35050505050565b60006108c2843385856106db565b5060019392505050565b6001600160a01b0383166000908152604685016020908152604080832033845290915281205482111561093b576040805162461bcd60e51b8152602060048201526012602482015271416c6c6f77616e63652069732077726f6e6760701b604482015290519081900360640190fd5b6001600160a01b03841660009081526046860160209081526040808320338452909152902080548390039055610973858585856106db565b506001949350505050565b815415806109b35750815443908390600019810190811061099b57fe5b6000918252602090912001546001600160801b031614155b15610a245760408051808201909152436001600160801b0390811682528281166020808401918252855460018101875560008781529190912093519301805491516fffffffffffffffffffffffffffffffff19909216938316939093178216600160801b9190921602179055610a5e565b815460009083906000198101908110610a3957fe5b600091825260209091200180546001600160801b03808516600160801b029116179055505b505056fe53686f756c6420686176652073756666696369656e742062616c616e636520746f207472616465547269656420746f2073656e64206e6f6e2d706f73697469766520616d6f756e74a265627a7a7231582003fc87d84f4ff4d1258e42d12e114ed052499862f2607d7a6563119d040da2f664736f6c63430005100032"

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
	event.Raw = log
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
	event.Raw = log
	return event, nil
}

// UtilitiesABI is the input ABI used to generate the binding from.
const UtilitiesABI = "[]"

// UtilitiesBin is the compiled bytecode used for deploying new contracts.
var UtilitiesBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582044ffbd8bc57eafebac8397c18ffa14cc7a334cef4deb5aca1c730c36cae9bd5564736f6c63430005100032"

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
