// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oldTellor

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
var SafeMathBin = "0x604c6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820ae7a6e542dbb565b320d4108e45cc18a4f720ee1597df1f211d7ccce701559680029"

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
const TellorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_propNewTellorAddress\",\"type\":\"address\"}],\"name\":\"proposeFork\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"requestStakingWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_c_sapi\",\"type\":\"string\"},{\"name\":\"_c_symbol\",\"type\":\"string\"},{\"name\":\"_granularity\",\"type\":\"uint256\"},{\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"requestData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"tallyVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nonce\",\"type\":\"string\"},{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"submitMiningSolution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pendingOwner\",\"type\":\"address\"}],\"name\":\"proposeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_minerIndex\",\"type\":\"uint256\"}],\"name\":\"beginDispute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"name\":\"_supportsDispute\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TellorFuncSigs maps the 4-byte function signature to its string representation.
var TellorFuncSigs = map[string]string{
	"752d49a1": "addTip(uint256,uint256)",
	"095ea7b3": "approve(address,uint256)",
	"8581af19": "beginDispute(uint256,uint256,uint256)",
	"4e71e0c8": "claimOwnership()",
	"0d2d76a2": "depositStake()",
	"26b7d9f6": "proposeFork(address)",
	"710bf322": "proposeOwnership(address)",
	"3fff2816": "requestData(string,string,uint256,uint256)",
	"28449c3a": "requestStakingWithdraw()",
	"68c180d5": "submitMiningSolution(string,uint256,uint256)",
	"4d318b0e": "tallyVotes(uint256)",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"c9d27afe": "vote(uint256,bool)",
	"bed9d861": "withdrawStake()",
}

// TellorBin is the compiled bytecode used for deploying new contracts.
var TellorBin = "0x608060405234801561001057600080fd5b50610d01806100206000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80634e71e0c8116100975780638581af19116100665780638581af1914610354578063a9059cbb1461037d578063bed9d861146103a9578063c9d27afe146103b1576100f5565b80634e71e0c81461028d57806368c180d514610295578063710bf3221461030b578063752d49a114610331576100f5565b806326b7d9f6116100d357806326b7d9f61461017a57806328449c3a146101a05780633fff2816146101a85780634d318b0e14610270576100f5565b8063095ea7b3146100fa5780630d2d76a21461013a57806323b872dd14610144575b600080fd5b6101266004803603604081101561011057600080fd5b506001600160a01b0381351690602001356103d6565b604080519115158252519081900360200190f35b610142610476565b005b6101266004803603606081101561015a57600080fd5b506001600160a01b038135811691602081013590911690604001356104e3565b6101426004803603602081101561019057600080fd5b50356001600160a01b031661058c565b610142610609565b610142600480360360808110156101be57600080fd5b8101906020810181356401000000008111156101d957600080fd5b8201836020820111156101eb57600080fd5b8035906020019184600183028401116401000000008311171561020d57600080fd5b91939092909160208101903564010000000081111561022b57600080fd5b82018360208201111561023d57600080fd5b8035906020019184600183028401116401000000008311171561025f57600080fd5b91935091508035906020013561065c565b6101426004803603602081101561028657600080fd5b503561073e565b610142610798565b610142600480360360608110156102ab57600080fd5b8101906020810181356401000000008111156102c657600080fd5b8201836020820111156102d857600080fd5b803590602001918460018302840111640100000000831117156102fa57600080fd5b9193509150803590602001356107a4565b6101426004803603602081101561032157600080fd5b50356001600160a01b0316610858565b6101426004803603604081101561034757600080fd5b508035906020013561086c565b6101426004803603606081101561036a57600080fd5b50803590602081013590604001356108e9565b6101266004803603604081101561039357600080fd5b506001600160a01b03813516906020013561096e565b6101426109db565b610142600480360360408110156103c757600080fd5b50803590602001351515610a2e565b60408051600160e11b634286e6190281526000600482018190526001600160a01b038516602483015260448201849052915173__$a7780dac84603eae88ac24b827b34ffc9e$__9163850dcc32916064808301926020929190829003018186803b15801561044357600080fd5b505af4158015610457573d6000803e3d6000fd5b505050506040513d602081101561046d57600080fd5b50519392505050565b60408051600160e11b63410516b3028152600060048201819052915173__$95e2a20849a26b4eb147d4480c0d0f303a$__9263820a2d669260248082019391829003018186803b1580156104c957600080fd5b505af41580156104dd573d6000803e3d6000fd5b50505050565b60408051600160e01b63ca5018990281526000600482018190526001600160a01b0380871660248401528516604483015260648201849052915173__$a7780dac84603eae88ac24b827b34ffc9e$__9163ca501899916084808301926020929190829003018186803b15801561055857600080fd5b505af415801561056c573d6000803e3d6000fd5b505050506040513d602081101561058257600080fd5b5051949350505050565b60408051600160e01b63694bf49f0281526000600482018190526001600160a01b0384166024830152915173__$01102ac036bcb1aea51b873d34f7cf54e9$__9263694bf49f9260448082019391829003018186803b1580156105ee57600080fd5b505af4158015610602573d6000803e3d6000fd5b5050505050565b60408051600160e21b633273d793028152600060048201819052915173__$95e2a20849a26b4eb147d4480c0d0f303a$__9263c9cf5e4c9260248082019391829003018186803b1580156104c957600080fd5b600073__$93445d4249bacb01c598edb6abcc89e8c4$__6385997bf690918888888888886040518863ffffffff1660e01b81526004018088815260200180602001806020018581526020018481526020018381038352898982818152602001925080828437600083820152601f01601f191690910184810383528781526020019050878780828437600081840152601f19601f820116905080830192505050995050505050505050505060006040518083038186803b15801561071e57600080fd5b505af4158015610732573d6000803e3d6000fd5b50505050505050505050565b60408051600160e01b63def6fac702815260006004820181905260248201849052915173__$01102ac036bcb1aea51b873d34f7cf54e9$__9263def6fac79260448082019391829003018186803b1580156105ee57600080fd5b6107a26000610a90565b565b600073__$93445d4249bacb01c598edb6abcc89e8c4$__63a098b5b49091868686866040518663ffffffff1660e01b815260040180868152602001806020018481526020018381526020018281038252868682818152602001925080828437600081840152601f19601f820116905080830192505050965050505050505060006040518083038186803b15801561083a57600080fd5b505af415801561084e573d6000803e3d6000fd5b5050505050505050565b61086960008263ffffffff610bdc16565b50565b60408051600160e01b6302e8f21b0281526000600482018190526024820185905260448201849052915173__$93445d4249bacb01c598edb6abcc89e8c4$__926302e8f21b9260648082019391829003018186803b1580156108cd57600080fd5b505af41580156108e1573d6000803e3d6000fd5b505050505050565b60408051600160e01b63ca9a4ea5028152600060048201819052602482018690526044820185905260648201849052915173__$01102ac036bcb1aea51b873d34f7cf54e9$__9263ca9a4ea59260848082019391829003018186803b15801561095157600080fd5b505af4158015610965573d6000803e3d6000fd5b50505050505050565b60408051600160e01b63c84b96f50281526000600482018190526001600160a01b038516602483015260448201849052915173__$a7780dac84603eae88ac24b827b34ffc9e$__9163c84b96f5916064808301926020929190829003018186803b15801561044357600080fd5b60408051600160e01b6344bacc4b028152600060048201819052915173__$95e2a20849a26b4eb147d4480c0d0f303a$__926344bacc4b9260248082019391829003018186803b1580156104c957600080fd5b60408051600160e11b6316d03837028152600060048201819052602482018590528315156044830152915173__$01102ac036bcb1aea51b873d34f7cf54e9$__92632da0706e9260648082019391829003018186803b1580156108cd57600080fd5b60408051600160991b6c3832b73234b733afb7bbb732b9028152815190819003600d0190206000908152603f83016020522054336001600160a01b0390911614610ad957600080fd5b60408051600160991b6c3832b73234b733afb7bbb732b9028152815190819003600d0181206000908152603f8401602081815284832054600160d11b652fb7bbb732b902855285519485900360060185208452919052928120546001600160a01b039384169316917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a360408051600160991b6c3832b73234b733afb7bbb732b9028152815190819003600d0181206000908152603f909301602081815283852054600160d11b652fb7bbb732b9028452845193849003600601909320855252912080546001600160a01b0319166001600160a01b03909216919091179055565b60408051600160d11b652fb7bbb732b902815281519081900360060190206000908152603f84016020522054336001600160a01b0390911614610c1e57600080fd5b60408051600160d11b652fb7bbb732b902815281519081900360060181206000908152603f8501602052918220546001600160a01b03808516939116917fb51454ce8c7f26becd312a46c4815553887f2ec876a0b8dc813b87f62edf6f809190a360408051600160991b6c3832b73234b733afb7bbb732b9028152815190819003600d0190206000908152603f90930160205290912080546001600160a01b039092166001600160a01b031990921691909117905556fea165627a7a72305820630d68a2d26db30416b2bf4562dcac945523c6632ee76a76e2021610dffbb3390029"

// DeployTellor deploys a new Ethereum contract, binding an instance of Tellor to it.
func DeployTellor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tellor, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	tellorDisputeAddr, _, _, _ := DeployTellorDispute(auth, backend)
	TellorBin = strings.Replace(TellorBin, "__$01102ac036bcb1aea51b873d34f7cf54e9$__", tellorDisputeAddr.String()[2:], -1)

	tellorLibraryAddr, _, _, _ := DeployTellorLibrary(auth, backend)
	TellorBin = strings.Replace(TellorBin, "__$93445d4249bacb01c598edb6abcc89e8c4$__", tellorLibraryAddr.String()[2:], -1)

	tellorStakeAddr, _, _, _ := DeployTellorStake(auth, backend)
	TellorBin = strings.Replace(TellorBin, "__$95e2a20849a26b4eb147d4480c0d0f303a$__", tellorStakeAddr.String()[2:], -1)

	tellorTransferAddr, _, _, _ := DeployTellorTransfer(auth, backend)
	TellorBin = strings.Replace(TellorBin, "__$a7780dac84603eae88ac24b827b34ffc9e$__", tellorTransferAddr.String()[2:], -1)

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

// RequestData is a paid mutator transaction binding the contract method 0x3fff2816.
//
// Solidity: function requestData(string _c_sapi, string _c_symbol, uint256 _granularity, uint256 _tip) returns()
func (_Tellor *TellorTransactor) RequestData(opts *bind.TransactOpts, _c_sapi string, _c_symbol string, _granularity *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "requestData", _c_sapi, _c_symbol, _granularity, _tip)
}

// RequestData is a paid mutator transaction binding the contract method 0x3fff2816.
//
// Solidity: function requestData(string _c_sapi, string _c_symbol, uint256 _granularity, uint256 _tip) returns()
func (_Tellor *TellorSession) RequestData(_c_sapi string, _c_symbol string, _granularity *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.RequestData(&_Tellor.TransactOpts, _c_sapi, _c_symbol, _granularity, _tip)
}

// RequestData is a paid mutator transaction binding the contract method 0x3fff2816.
//
// Solidity: function requestData(string _c_sapi, string _c_symbol, uint256 _granularity, uint256 _tip) returns()
func (_Tellor *TellorTransactorSession) RequestData(_c_sapi string, _c_symbol string, _granularity *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.RequestData(&_Tellor.TransactOpts, _c_sapi, _c_symbol, _granularity, _tip)
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

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_Tellor *TellorTransactor) SubmitMiningSolution(opts *bind.TransactOpts, _nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "submitMiningSolution", _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_Tellor *TellorSession) SubmitMiningSolution(_nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.SubmitMiningSolution(&_Tellor.TransactOpts, _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_Tellor *TellorTransactorSession) SubmitMiningSolution(_nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
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
const TellorDisputeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"NewDispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_position\",\"type\":\"bool\"},{\"indexed\":true,\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_result\",\"type\":\"int256\"},{\"indexed\":true,\"name\":\"_reportedMiner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_reportingParty\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"DisputeVoteTallied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_newTellor\",\"type\":\"address\"}],\"name\":\"NewTellorAddress\",\"type\":\"event\"}]"

// TellorDisputeFuncSigs maps the 4-byte function signature to its string representation.
var TellorDisputeFuncSigs = map[string]string{
	"ca9a4ea5": "beginDispute(TellorStorage.TellorStorageStruct storage,uint256,uint256,uint256)",
	"694bf49f": "proposeFork(TellorStorage.TellorStorageStruct storage,address)",
	"def6fac7": "tallyVotes(TellorStorage.TellorStorageStruct storage,uint256)",
	"e15f6f70": "updateDisputeFee(TellorStorage.TellorStorageStruct storage)",
	"2da0706e": "vote(TellorStorage.TellorStorageStruct storage,uint256,bool)",
}

// TellorDisputeBin is the compiled bytecode used for deploying new contracts.
var TellorDisputeBin = "0x611702610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100615760003560e01c80632da0706e14610066578063694bf49f146100a0578063ca9a4ea5146100d9578063def6fac714610115578063e15f6f7014610145575b600080fd5b81801561007257600080fd5b5061009e6004803603606081101561008957600080fd5b5080359060208101359060400135151561016f565b005b8180156100ac57600080fd5b5061009e600480360360408110156100c357600080fd5b50803590602001356001600160a01b031661039a565b8180156100e557600080fd5b5061009e600480360360808110156100fc57600080fd5b50803590602081013590604081013590606001356107e6565b81801561012157600080fd5b5061009e6004803603604081101561013857600080fd5b5080359060200135610e3d565b81801561015157600080fd5b5061009e6004803603602081101561016857600080fd5b5035611494565b60008281526044808501602090815260408084208151600160a91b6a313637b1b5a73ab6b132b9028152825190819003600b018120865260058201845282862054600160e01b633f48b1ff028252600482018a905233602483015294810194909452905190939273__$a7780dac84603eae88ac24b827b34ffc9e$__92633f48b1ff92606480840193829003018186803b15801561020c57600080fd5b505af4158015610220573d6000803e3d6000fd5b505050506040513d602081101561023657600080fd5b505133600090815260068401602052604090205490915060ff1615156001141561025f57600080fd5b6000811161026c57600080fd5b3360009081526047860160205260409020546003141561028b57600080fd5b336000908152600680840160209081526040808420805460ff1916600190811790915581517f6e756d6265724f66566f746573000000000000000000000000000000000000008152825190819003600d0181208652600588018085528387208054909301909255600160d01b6571756f72756d02815282519081900390940190932084529190529020805482019055821561033f576001820154610335908263ffffffff61164216565b600183015561035a565b6001820154610354908263ffffffff61167316565b60018301555b6040805184151581529051339186917f86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae09181900360200190a35050505050565b604080516001600160a01b03831660601b60208083019190915282518083036014018152603490920183528151918101919091206000818152604a860190925291902054156103e857600080fd5b60408051600160b01b6964697370757465466565028152815190819003600a018120600090815282860160205282812054600160e01b63c7bb46ad028352600483018790523360248401523060448401526064830152915173__$a7780dac84603eae88ac24b827b34ffc9e$__9263c7bb46ad9260848082019391829003018186803b15801561047757600080fd5b505af415801561048b573d6000803e3d6000fd5b505050508260400160006040518080600160a21b6b191a5cdc1d5d1950dbdd5b9d02815250600c019050604051809103902081526020019081526020016000206000815480929190600101919050555060008360400160006040518080600160a21b6b191a5cdc1d5d1950dbdd5b9d02815250600c019050604051809103902081526020019081526020016000205490508084604a0160008481526020019081526020016000208190555060405180610100016040528083815260200160008152602001600015158152602001600015158152602001600115158152602001336001600160a01b03168152602001336001600160a01b03168152602001846001600160a01b0316815250846044016000838152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548160ff02191690831515021790555060608201518160020160016101000a81548160ff02191690831515021790555060808201518160020160026101000a81548160ff02191690831515021790555060a08201518160020160036101000a8154816001600160a01b0302191690836001600160a01b0316021790555060c08201518160030160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060e08201518160040160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050504384604401600083815260200190815260200160002060050160006040518080600160a91b6a313637b1b5a73ab6b132b902815250600b01905060405180910390208152602001908152602001600020819055508360400160006040518080600160b01b696469737075746546656502815250600a019050604051809103902081526020019081526020016000205484604401600083815260200190815260200160002060050160006040518080600160e81b6266656502815250600301905060405180910390208152602001908152602001600020819055504262093a800184604401600083815260200190815260200160002060050160006040518080600160801b6f6d696e457865637574696f6e44617465028152506010019050604051809103902081526020019081526020016000208190555050505050565b60008381526048850160209081526040808320858452600581019092529091205460904391909103111561081957600080fd5b600083815260058201602052604090205461083357600080fd5b6005821061084057600080fd5b60008381526008820160205260408120836005811061085b57fe5b0154604080516001600160a01b03909216606081901b60208085019190915260348401899052605480850189905283518086039091018152607490940183528351938101939093206000818152604a8b01909452919092205491925090156108c257600080fd5b60408051600160b01b6964697370757465466565028152815190819003600a0181206000908152828a0160205282812054600160e01b63c7bb46ad028352600483018b90523360248401523060448401526064830152915173__$a7780dac84603eae88ac24b827b34ffc9e$__9263c7bb46ad9260848082019391829003018186803b15801561095157600080fd5b505af4158015610965573d6000803e3d6000fd5b505050508660400160006040518080600160a21b6b191a5cdc1d5d1950dbdd5b9d02815250600c01905060405180910390208152602001908152602001600020546001018760400160006040518080600160a21b6b191a5cdc1d5d1950dbdd5b9d02815250600c019050604051809103902081526020019081526020016000208190555060008760400160006040518080600160a21b6b191a5cdc1d5d1950dbdd5b9d02815250600c019050604051809103902081526020019081526020016000205490508088604a0160008481526020019081526020016000208190555060405180610100016040528083815260200160008152602001600015158152602001600015158152602001600015158152602001846001600160a01b03168152602001336001600160a01b0316815260200160006001600160a01b0316815250886044016000838152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548160ff02191690831515021790555060608201518160020160016101000a81548160ff02191690831515021790555060808201518160020160026101000a81548160ff02191690831515021790555060a08201518160020160036101000a8154816001600160a01b0302191690836001600160a01b0316021790555060c08201518160030160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060e08201518160040160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050508688604401600083815260200190815260200160002060050160006040518080600160ba1b681c995c5d595cdd125902815250600901905060405180910390208152602001908152602001600020819055508588604401600083815260200190815260200160002060050160006040518080600160bc1b68074696d657374616d702815250600901905060405180910390208152602001908152602001600020819055508360090160008781526020019081526020016000208560058110610c7b57fe5b0154600082815260448a01602081815260408084208151600160d81b6476616c756502815282519081900360059081018220875290910180845282862096909655868552838352600160801b6f6d696e457865637574696f6e446174650281528151908190036010018120855285835281852062093a8042019055868552838352600160a91b6a313637b1b5a73ab6b132b9028152815190819003600b0181208552858352818520439055868552838352600160ba1b681b5a5b995c94db1bdd028152815190819003600901812085528583528185208b9055600160b01b6964697370757465466565028152815190819003600a0181208552818e01835281852054878652938352600160e81b6266656502815281519081900360030190208452939052919020556002851415610dd657600087815260488901602090815260408083208984526007019091529020805460ff191660011790555b6001600160a01b038316600081815260478a016020908152604091829020600390558151898152908101929092528051899284927feceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da6492918290030190a35050505050505050565b600081815260448301602090815260408083208151600160ba1b681c995c5d595cdd12590281528251908190036009019020845260058101835281842054845260488601909252909120600282015460ff1615610e9957600080fd5b60408051600160801b6f6d696e457865637574696f6e44617465028152815190819003601001902060009081526005840160205220544211610eda57600080fd5b600282015462010000900460ff166112e2576002820154630100000090046001600160a01b03166000908152604785016020526040812060018401549091121561118b57600081556201518042064203600182015560408051600160aa1b6a1cdd185ad95c90dbdd5b9d028152815190819003600b0190206000908152818701602052208054600019019055610f6f85611494565b6002830154600384015460408051600160aa1b6a1cdd185ad9505b5bdd5b9d028152815190819003600b0181206000908152828a0160205282812054600160e01b63c7bb46ad028352600483018b90526001600160a01b0363010000009096048616602484015293909416604482015260648101929092525173__$a7780dac84603eae88ac24b827b34ffc9e$__9263c7bb46ad9260848082019391829003018186803b15801561101f57600080fd5b505af4158015611033573d6000803e3d6000fd5b50505060038085015460408051600160e81b62666565028152815190819003909301832060009081526005880160205281812054600160e01b63c7bb46ad028552600485018b90523060248601526001600160a01b03909316604485015260648401929092525173__$a7780dac84603eae88ac24b827b34ffc9e$__935063c7bb46ad926084808201939291829003018186803b1580156110d357600080fd5b505af41580156110e7573d6000803e3d6000fd5b50505060028401805461ff0019166101001790555060408051600160bc1b68074696d657374616d702815281519081900360090190206000908152600585016020908152828220548252600785019052205460ff161515600114156111865760408051600160bc1b68074696d657374616d702815281519081900360090190206000908152600585016020908152828220548252600685019052908120555b6112dc565b60018155600283015460408051600160e81b62666565028152815190819003600301812060009081526005870160205282812054600160e01b63c7bb46ad028352600483018a90523060248401526001600160a01b0363010000009095049490941660448301526064820193909352905173__$a7780dac84603eae88ac24b827b34ffc9e$__9263c7bb46ad9260848082019391829003018186803b15801561123357600080fd5b505af4158015611247573d6000803e3d6000fd5b505060408051600160bc1b68074696d657374616d702815281519081900360090190206000908152600587016020908152828220548252600787019052205460ff1615156001141591506112dc90505760408051600160bc1b68074696d657374616d70281528151908190036009019020600090815260058501602090815282822054825260078501905220805460ff191690555b50611416565b60008260010154131561141657604080517f746f74616c5f737570706c7900000000000000000000000000000000000000008152815190819003600c0190206000908152818601602052205460649060140260408051600160d01b6571756f72756d028152815190819003600601902060009081526005860160205220549190041061136d57600080fd5b600482018054604080517f74656c6c6f72436f6e74726163740000000000000000000000000000000000008152815190819003600e0181206000908152603f890160209081529083902080546001600160a01b0319166001600160a01b0395861617905560028701805461ff00191661010017905593549092168252517fc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d929181900390910190a15b60028201805460ff19166001908117918290558301546003840154604080519283526001600160a01b039182166020840152610100840460ff16151583820152516301000000909304169185917f21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739919081900360600190a350505050565b60408051600160a01b6b7461726765744d696e657273028152815190819003600c0181206000908152828401602081815284832054600160aa1b6a1cdd185ad95c90dbdd5b9d028552855194859003600b0190942083525291909120546103e891908202816114ff57fe5b0410156116055760408051600160a01b6b7461726765744d696e657273028152815190819003600c0181206000908152828401602081815284832054600160aa1b6a1cdd185ad95c90dbdd5b9d028552855194859003600b0190942083525291909120546115d19167d02ab486cedc0000916103e8916115c49183028161158257fe5b60408051600160aa1b6a1cdd185ad9505b5bdd5b9d028152815190819003600b019020600090815281890160205220549190046103e80363ffffffff61169916565b816115cb57fe5b046116c0565b60408051600160b01b6964697370757465466565028152815190819003600a0190206000908152818401602052205561163f565b60408051600160b01b6964697370757465466565028152815190819003600a01902060009081528183016020522067d02ab486cedc000090555b50565b60008082131561165f57508181018281121561165a57fe5b61166d565b508181018281131561166d57fe5b92915050565b60008082131561168b57508082038281131561165a57fe5b508082038281121561166d57fe5b60008282028315806116b35750828482816116b057fe5b04145b6116b957fe5b9392505050565b60008183116116cf57816116b9565b509091905056fea165627a7a72305820257b4e0c609c5b73a9f4c7b59485b352f2d058b812d5f68b22df0038b179196d0029"

// DeployTellorDispute deploys a new Ethereum contract, binding an instance of TellorDispute to it.
func DeployTellorDispute(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorDispute, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorDisputeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	tellorTransferAddr, _, _, _ := DeployTellorTransfer(auth, backend)
	TellorDisputeBin = strings.Replace(TellorDisputeBin, "__$a7780dac84603eae88ac24b827b34ffc9e$__", tellorTransferAddr.String()[2:], -1)

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
	DisputeID *big.Int
	Position  bool
	Voter     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter)
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
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter)
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

// ParseVoted is a log parse operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter)
func (_TellorDispute *TellorDisputeFilterer) ParseVoted(log types.Log) (*TellorDisputeVoted, error) {
	event := new(TellorDisputeVoted)
	if err := _TellorDispute.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorGettersLibraryABI is the input ABI used to generate the binding from.
const TellorGettersLibraryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_newTellor\",\"type\":\"address\"}],\"name\":\"NewTellorAddress\",\"type\":\"event\"}]"

// TellorGettersLibraryBin is the compiled bytecode used for deploying new contracts.
var TellorGettersLibraryBin = "0x604c6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820ef0ba1686e05d690debf6bee65b3cf53f14e3e1a518408d8bdbb0a94ea7031ac0029"

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
const TellorLibraryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_tip\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"TipAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_querySymbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_granularity\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"DataRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_currentChallenge\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_currentRequestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_multiplier\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"NewChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_onDeckQueryHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_onDeckTotalTips\",\"type\":\"uint256\"}],\"name\":\"NewRequestOnDeck\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_time\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_totalTips\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NewValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_miner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_nonce\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NonceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipProposed\",\"type\":\"event\"}]"

// TellorLibraryFuncSigs maps the 4-byte function signature to its string representation.
var TellorLibraryFuncSigs = map[string]string{
	"02e8f21b": "addTip(TellorStorage.TellorStorageStruct storage,uint256,uint256)",
	"85997bf6": "requestData(TellorStorage.TellorStorageStruct storage,string,string,uint256,uint256)",
	"a098b5b4": "submitMiningSolution(TellorStorage.TellorStorageStruct storage,string,uint256,uint256)",
}

// TellorLibraryBin is the compiled bytecode used for deploying new contracts.
var TellorLibraryBin = "0x6123f1610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c806302e8f21b1461005057806385997bf614610088578063a098b5b4146101ce575b600080fd5b81801561005c57600080fd5b506100866004803603606081101561007357600080fd5b508035906020810135906040013561028d565b005b81801561009457600080fd5b50610086600480360360a08110156100ab57600080fd5b813591908101906040810160208201356401000000008111156100cd57600080fd5b8201836020820111156100df57600080fd5b8035906020019184600183028401116401000000008311171561010157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561015457600080fd5b82018360208201111561016657600080fd5b8035906020019184600183028401116401000000008311171561018857600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602001356103a9565b8180156101da57600080fd5b50610086600480360360808110156101f157600080fd5b8135919081019060408101602082013564010000000081111561021357600080fd5b82018360208201111561022557600080fd5b8035906020019184600183028401116401000000008311171561024757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020013561081a565b6000821161029a57600080fd5b80156103215760408051600160e01b63c7bb46ad0281526004810185905233602482015230604482015260648101839052905173__$a7780dac84603eae88ac24b827b34ffc9e$__9163c7bb46ad916084808301926000929190829003018186803b15801561030857600080fd5b505af415801561031c573d6000803e3d6000fd5b505050505b61032c838383610ce5565b600082815260488401602090815260408083208151600160c41b670746f74616c5469702815282519081900360080181208552600490910183529281902054848452918301919091528051849233927fd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e882092918290030190a3505050565b600082116103b657600080fd5b670de0b6b3a76400008211156103cb57600080fd5b8351849084906103da57600080fd5b60408151106103e857600080fd5b600082856040516020018083805190602001908083835b6020831061041e5780518252601f1990920191602091820191016103ff565b51815160209384036101000a60001901801990921691161790529201938452506040805180850381529382018152835193820193909320600081815260498e019092529290205491935050151590506107f357604080517f72657175657374436f756e7400000000000000000000000000000000000000008082528251600c928190038301812060009081528c8501602081815286832080546001019055938352855192839003909401822081529282528383205460808201855287825281830187905281850186905284518481528084018652606083015280845260488d01835293909220825180519192610519928492909101906121a6565b50602082810151805161053292600185019201906121a6565b506040820151600282015560608201518051610558916003840191602090910190612224565b505050600081815260488a01602081815260408084208151600160a81b6a6772616e756c6172697479028152825190819003600b018120865260049091018084528286208c9055868652848452600080516020612366833981519152825282519182900360100182208652808452828620869055868652938352600160c41b670746f74616c546970281528151908190036008019020845291815281832083905584835260498c0190529020819055841561068e5760408051600160e01b63c7bb46ad028152600481018b905233602482015230604482015260648101879052905173__$a7780dac84603eae88ac24b827b34ffc9e$__9163c7bb46ad916084808301926000929190829003018186803b15801561067557600080fd5b505af4158015610689573d6000803e3d6000fd5b505050505b610699898287610ce5565b600081815260488a016020908152604091829020825192830189905260608301889052608080845281546002600180831615610100026000190190921604918501829052859433947f6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc99493928401928d928d929091829182019060a0830190889080156107675780601f1061073c57610100808354040283529160200191610767565b820191906000526020600020905b81548152906001019060200180831161074a57829003601f168201915b50508381038252865460026000196101006001841615020190911604808252602090910190879080156107db5780601f106107b0576101008083540402835291602001916107db565b820191906000526020600020905b8154815290600101906020018083116107be57829003601f168201915b5050965050505050505060405180910390a350610810565b60008181526049890160205260409020546108109089908661028d565b5050505050505050565b33600090815260478501602052604090205460011461083857600080fd5b604080516000805160206123868339815191528152815190819003601001902060009081528186016020522054821461087057600080fd5b8360400160006040518080600160b01b69646966666963756c747902815250600a0190506040518091039020815260200190815260200160002054600260038660000154338760405160200180848152602001836001600160a01b03166001600160a01b031660601b815260140182805190602001908083835b602083106109095780518252601f1990920191602091820191016108ea565b6001836020036101000a038019825116818451168082178552505050505050905001935050505060405160208183030381529060405280519060200120604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b602083106109945780518252601f199092019160209182019101610975565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa1580156109d3573d6000803e3d6000fd5b5050506040515160601b60405160200180826bffffffffffffffffffffffff19166bffffffffffffffffffffffff191681526014019150506040516020818303038152906040526040518082805190602001908083835b60208310610a495780518252601f199092019160209182019101610a2a565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610a88573d6000803e3d6000fd5b5050506040513d6020811015610a9d57600080fd5b505181610aa657fe5b0615610ab157600080fd5b83546000908152604185016020908152604080832033845290915290205460ff1615610adc57600080fd5b60408051600160a01b6b736c6f7450726f6772657373028152815190819003600c019020600090815281860160205220548190603586019060058110610b1e57fe5b600202015560408051600160a01b6b736c6f7450726f6772657373028152815190819003600c019020600090815281860160205220543390603586019060058110610b6557fe5b60020201600190810180546001600160a01b0319166001600160a01b03939093169290921790915560408051600160a01b6b736c6f7450726f67726573730281528151600c91819003919091018120600090815287830160209081528382208054860190558854825260418901815283822033808452908252848320805460ff19169096179095558854818401879052938301849052606080845288519084015287518795947fe6d63a2aee0aaed2ab49702313ce54114f2145af219d7db30d6624af9e6dffc4948a9489949293919283926080840192918801918190849084905b83811015610c5f578181015183820152602001610c47565b50505050905090810190601f168015610c8c5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a360408051600160a01b6b736c6f7450726f6772657373028152815190819003600c0190206000908152818601602052205460051415610cdf57610cdf8484846113b3565b50505050565b6000828152604884016020526040812090610cff8561209c565b90508215610d745760408051600160c41b670746f74616c5469702815281519081900360080190206000908152600484016020522054610d45908463ffffffff6120f816565b60408051600160c41b670746f74616c54697028152815190819003600801902060009081526004850160205220555b60408051600160c41b670746f74616c54697028152815190819003600801812060009081526004850160209081528382205460008051602061238683398151915284528451938490036010019093208252888401905291909120546111355760008360040160006040518080600160c41b670746f74616c546970281525060080190506040518091039020815260200190815260200160002081905550848660400160006040518080600080516020612386833981519152815250601001905060405180910390208152602001908152602001600020819055508086604001600060405180806000805160206123a68339815191528152506010019050604051809103902081526020019081526020016000208190555080866000015460014303406040516020018084815260200183815260200182815260200193505050506040516020818303038152906040528051906020012086600001819055508560400160006040518080600080516020612386833981519152815250601001905060405180910390208152602001908152602001600020547f9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f4287600001548860400160006040518080600160b01b69646966666963756c747902815250600a01905060405180910390208152602001908152602001600020548960480160008b6040016000604051808060008051602061238683398151915281525060100190506040518091039020815260200190815260200160002054815260200190815260200160002060040160006040518080600160a81b6a6772616e756c617269747902815250600b01905060405180910390208152602001908152602001600020548a60480160008c604001600060405180806000805160206123868339815191528152506010019050604051809103902081526020019081526020016000205481526020019081526020016000206000018b604001600060405180806000805160206123a683398151915281525060100190506040518091039020815260200190815260200160002054604051808681526020018581526020018481526020018060200183815260200182810382528481815460018160011615610100020316600290048152602001915080546001816001161561010002031660029004801561111e5780601f106110f35761010080835404028352916020019161111e565b820191906000526020600020905b81548152906001019060200180831161110157829003601f168201915b5050965050505050505060405180910390a26113ab565b600082815260488701602090815260408083208151600160c41b670746f74616c546970281528251908190036008019020845260040190915290205481118061117c575081155b156112495760028084015460408051602081018390529081018490526060808252865460001961010060018316150201169390930492810183905287927f5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c928792909186919081906080820190869080156112385780601f1061120d57610100808354040283529160200191611238565b820191906000526020600020905b81548152906001019060200180831161121b57829003601f168201915b505094505050505060405180910390a25b6040805160008051602061236683398151915281528151908190036010019020600090815260048501602052205461135d576040805161066081019182905260009182916112b99160018b019060339082845b81548152602001906001019080831161129c57505050505061210e565b9092509050818311806112ca575081155b1561135657828860010182603381106112df57fe5b0155600081815260438901602081815260408084208054855260488d018352818520825160008051602061236683398151915280825284519182900360109081018320895260049384018752858920899055898952968652928e905591825282519182900390940190208452918801905290208190555b50506113ab565b83156113ab576040805160008051602061236683398151915281528151908190036010019020600090815260048501602052205484906001880190603381106113a257fe5b01805490910190555b505050505050565b600081815260488401602090815260408083208151600160701b7174696d654f664c6173744e657756616c75650281528251601291819003919091018120855282880180855283862054600160b21b691d1a5b5955185c99d95d0283528451600a9381900384018120885282875285882054600160b01b69646966666963756c747902808352875192839003860183208a52848952878a20549083528751928390039095019091208852919095529285205491946064429590950390930302929092059091019081136114b75760408051600160b01b69646966666963756c7479028152815190819003600a019020600090815281870160205220600190556114e9565b60408051600160b01b69646966666963756c7479028152815190819003600a0190206000908152818701602052208190555b6000603c4260408051600160701b7174696d654f664c6173744e657756616c756502815281519081900360120190206000908152818a0160205220919006420390819055905061153761225e565b6040805160a081019091526035880160056000835b8282101561158d5760408051808201909152600283028501805482526001908101546001600160a01b0316602080840191909152918352909201910161154c565b509293506001925050505b60058110156116b75760008282600581106115af57fe5b602002015151905060008383600581106115c557fe5b602002015160200151905060008390505b6000811180156115f957508460018203600581106115f057fe5b60200201515183105b1561166b5784600182036005811061160d57fe5b60200201515185826005811061161f57fe5b6020020151528460001982016005811061163557fe5b60200201516020015185826005811061164a57fe5b602090810291909101516001600160a01b03909216910152600019016115d6565b838110156116ac578285826005811061168057fe5b6020020151528185826005811061169357fe5b602090810291909101516001600160a01b039092169101525b505050600101611598565b5060005b60058110156117d35773__$a7780dac84603eae88ac24b827b34ffc9e$__63c7bb46ad89308585600581106116ec57fe5b60200201516020015160058d604001600060405180806000805160206123a6833981519152815250601001905060405180910390208152602001908152602001600020548161173757fe5b04674563918244f40000016040518563ffffffff1660e01b815260040180858152602001846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b0316815260200182815260200194505050505060006040518083038186803b1580156117af57600080fd5b505af41580156117c3573d6000803e3d6000fd5b5050600190920191506116bb9050565b6040808301515181516000805160206123a683398151915280825283516010928190038301812060009081528d8601602081815287832054948452875193849003909501832082528452859020548d54898352938201949094526005909106909203828401526060820152905187917fc509d04e0782579ee59912139246ccbdf6c36c43abd90591d912017f3467dd16919081900360800190a2604080517f746f74616c5f737570706c7900000000000000000000000000000000000000008152815190819003600c0181206000908152828b016020908152838220805468017da3a04c7b3e0000019055600160d11b652fb7bbb732b902835283519283900360060183208252603f8c01905282812054600160e01b63c7bb46ad028352600483018c90523060248401526001600160a01b031660448301526722b1c8c1227a00006064830152915173__$a7780dac84603eae88ac24b827b34ffc9e$__9263c7bb46ad9260848082019391829003018186803b15801561195357600080fd5b505af4158015611967573d6000803e3d6000fd5b505050508160026005811061197857fe5b6020908102919091015151600085815260068801835260408082209290925560038801805460018101825590825283822001869055815160a08101835285518401516001600160a01b0390811682528685015185015181168286015286840151850151811682850152606080880151860151821690830152608080880151860151909116908201528682526008890190935220611a1691600561228c565b506040805160a081018252835151815260208085015151818301528483015151828401526060808601515190830152608080860151519083015260008681526009890190915291909120611a6b9160056122e0565b506000838152600586016020908152604080832043905560428b01825280832089905560348b018054600181018255908452828420018690558051600160a01b6b736c6f7450726f6772657373028152815190819003600c01902083528a81019091528120819055611adc8961209c565b60408051600080516020612386833981519152815281519081900360100190206000908152818c01602052208190559050801561205d5788604801600082815260200190815260200160002060040160006040518080600160c41b670746f74616c54697028152506008019050604051809103902081526020019081526020016000205489604001600060405180806000805160206123a6833981519152815250601001905060405180910390208152602001908152602001600020819055506000896001018a6048016000848152602001908152602001600020600401600060405180806000805160206123668339815191528152506010019050604051809103902081526020019081526020016000205460338110611bf957fe5b018190555060008960430160008b6048016000858152602001908152602001600020600401600060405180806000805160206123668339815191528152506010019050604051809103902081526020019081526020016000205481526020019081526020016000208190555060008960480160008381526020019081526020016000206004016000604051808060008051602061236683398151915281525060100190506040518091039020815260200190815260200160002081905550600089604801600083815260200190815260200160002060040160006040518080600160c41b670746f74616c5469702815250600801905060405180910390208152602001908152602001600020819055506000611d148a61209c565b9050888a6000015460014303406040516020018084805190602001908083835b60208310611d535780518252601f199092019160209182019101611d34565b6001836020036101000a0380198251168184511680821785525050505050509050018381526020018281526020019350505050604051602081830303815290604052805190602001208a60000181905550817f9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f428b600001548c60400160006040518080600160b01b69646966666963756c747902815250600a01905060405180910390208152602001908152602001600020548d604801600087815260200190815260200160002060040160006040518080600160a81b6a6772616e756c617269747902815250600b01905060405180910390208152602001908152602001600020548e60480160008881526020019081526020016000206000018f604001600060405180806000805160206123a6833981519152815250601001905060405180910390208152602001908152602001600020546040518086815260200185815260200184815260200180602001838152602001828103825284818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015611f485780601f10611f1d57610100808354040283529160200191611f48565b820191906000526020600020905b815481529060010190602001808311611f2b57829003601f168201915b5050965050505050505060405180910390a2600081815260488b01602090815260408083206002808201548351600160c41b670746f74616c546970281528451908190036008018120875260048401865295849020549486018190529285018490526060808652825460001961010060018316150201169190910490850181905285947f5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c949293929181906080820190869080156120475780601f1061201c57610100808354040283529160200191612047565b820191906000526020600020905b81548152906001019060200180831161202a57829003601f168201915b505094505050505060405180910390a250612091565b604080516000805160206123a6833981519152815281519081900360100190206000908152818b0160205290812081905589555b505050505050505050565b60408051610660810191829052600091829182916120dd9190600187019060339082845b8154815260200190600101908083116120c057505050505061215c565b60009081526043909501602052505060409092205492915050565b60008282018381101561210757fe5b9392505050565b6106408101516032805b8015612156578284826033811061212b57fe5b6020020151101561214d5783816033811061214257fe5b602002015192508091505b60001901612118565b50915091565b6020810151600060015b6033811015612156578284826033811061217c57fe5b6020020151111561219e5783816033811061219357fe5b602002015192508091505b600101612166565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106121e757805160ff1916838001178555612214565b82800160010185558215612214579182015b828111156122145782518255916020019190600101906121f9565b5061222092915061230d565b5090565b82805482825590600052602060002090810192821561221457916020028201828111156122145782518255916020019190600101906121f9565b6040518061014001604052806005905b61227661232a565b81526020019060019003908161226e5790505090565b82600581019282156122d4579160200282015b828111156122d457825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019061229f565b50612220929150612341565b826005810192821561221457916020028201828111156122145782518255916020019190600101906121f9565b61232791905b808211156122205760008155600101612313565b90565b604080518082019091526000808252602082015290565b61232791905b808211156122205780546001600160a01b031916815560010161234756fe7265717565737451506f736974696f6e0000000000000000000000000000000063757272656e745265717565737449640000000000000000000000000000000063757272656e74546f74616c5469707300000000000000000000000000000000a165627a7a723058206f99470bba95c8ba562fc8cfb952d5a478e131198ceb4607d9950aab91445aaf0029"

// DeployTellorLibrary deploys a new Ethereum contract, binding an instance of TellorLibrary to it.
func DeployTellorLibrary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorLibrary, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorLibraryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	tellorTransferAddr, _, _, _ := DeployTellorTransfer(auth, backend)
	TellorLibraryBin = strings.Replace(TellorLibraryBin, "__$a7780dac84603eae88ac24b827b34ffc9e$__", tellorTransferAddr.String()[2:], -1)

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

// TellorLibraryDataRequestedIterator is returned from FilterDataRequested and is used to iterate over the raw logs and unpacked data for DataRequested events raised by the TellorLibrary contract.
type TellorLibraryDataRequestedIterator struct {
	Event *TellorLibraryDataRequested // Event containing the contract specifics and raw log

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
func (it *TellorLibraryDataRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorLibraryDataRequested)
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
		it.Event = new(TellorLibraryDataRequested)
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
func (it *TellorLibraryDataRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorLibraryDataRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorLibraryDataRequested represents a DataRequested event raised by the TellorLibrary contract.
type TellorLibraryDataRequested struct {
	Sender      common.Address
	Query       string
	QuerySymbol string
	Granularity *big.Int
	RequestId   *big.Int
	TotalTips   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDataRequested is a free log retrieval operation binding the contract event 0x6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc9.
//
// Solidity: event DataRequested(address indexed _sender, string _query, string _querySymbol, uint256 _granularity, uint256 indexed _requestId, uint256 _totalTips)
func (_TellorLibrary *TellorLibraryFilterer) FilterDataRequested(opts *bind.FilterOpts, _sender []common.Address, _requestId []*big.Int) (*TellorLibraryDataRequestedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _TellorLibrary.contract.FilterLogs(opts, "DataRequested", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &TellorLibraryDataRequestedIterator{contract: _TellorLibrary.contract, event: "DataRequested", logs: logs, sub: sub}, nil
}

// WatchDataRequested is a free log subscription operation binding the contract event 0x6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc9.
//
// Solidity: event DataRequested(address indexed _sender, string _query, string _querySymbol, uint256 _granularity, uint256 indexed _requestId, uint256 _totalTips)
func (_TellorLibrary *TellorLibraryFilterer) WatchDataRequested(opts *bind.WatchOpts, sink chan<- *TellorLibraryDataRequested, _sender []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _TellorLibrary.contract.WatchLogs(opts, "DataRequested", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorLibraryDataRequested)
				if err := _TellorLibrary.contract.UnpackLog(event, "DataRequested", log); err != nil {
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

// ParseDataRequested is a log parse operation binding the contract event 0x6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc9.
//
// Solidity: event DataRequested(address indexed _sender, string _query, string _querySymbol, uint256 _granularity, uint256 indexed _requestId, uint256 _totalTips)
func (_TellorLibrary *TellorLibraryFilterer) ParseDataRequested(log types.Log) (*TellorLibraryDataRequested, error) {
	event := new(TellorLibraryDataRequested)
	if err := _TellorLibrary.contract.UnpackLog(event, "DataRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	CurrentRequestId *big.Int
	Difficulty       *big.Int
	Multiplier       *big.Int
	Query            string
	TotalTips        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewChallenge is a free log retrieval operation binding the contract event 0x9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42.
//
// Solidity: event NewChallenge(bytes32 _currentChallenge, uint256 indexed _currentRequestId, uint256 _difficulty, uint256 _multiplier, string _query, uint256 _totalTips)
func (_TellorLibrary *TellorLibraryFilterer) FilterNewChallenge(opts *bind.FilterOpts, _currentRequestId []*big.Int) (*TellorLibraryNewChallengeIterator, error) {

	var _currentRequestIdRule []interface{}
	for _, _currentRequestIdItem := range _currentRequestId {
		_currentRequestIdRule = append(_currentRequestIdRule, _currentRequestIdItem)
	}

	logs, sub, err := _TellorLibrary.contract.FilterLogs(opts, "NewChallenge", _currentRequestIdRule)
	if err != nil {
		return nil, err
	}
	return &TellorLibraryNewChallengeIterator{contract: _TellorLibrary.contract, event: "NewChallenge", logs: logs, sub: sub}, nil
}

// WatchNewChallenge is a free log subscription operation binding the contract event 0x9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42.
//
// Solidity: event NewChallenge(bytes32 _currentChallenge, uint256 indexed _currentRequestId, uint256 _difficulty, uint256 _multiplier, string _query, uint256 _totalTips)
func (_TellorLibrary *TellorLibraryFilterer) WatchNewChallenge(opts *bind.WatchOpts, sink chan<- *TellorLibraryNewChallenge, _currentRequestId []*big.Int) (event.Subscription, error) {

	var _currentRequestIdRule []interface{}
	for _, _currentRequestIdItem := range _currentRequestId {
		_currentRequestIdRule = append(_currentRequestIdRule, _currentRequestIdItem)
	}

	logs, sub, err := _TellorLibrary.contract.WatchLogs(opts, "NewChallenge", _currentRequestIdRule)
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

// ParseNewChallenge is a log parse operation binding the contract event 0x9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42.
//
// Solidity: event NewChallenge(bytes32 _currentChallenge, uint256 indexed _currentRequestId, uint256 _difficulty, uint256 _multiplier, string _query, uint256 _totalTips)
func (_TellorLibrary *TellorLibraryFilterer) ParseNewChallenge(log types.Log) (*TellorLibraryNewChallenge, error) {
	event := new(TellorLibraryNewChallenge)
	if err := _TellorLibrary.contract.UnpackLog(event, "NewChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorLibraryNewRequestOnDeckIterator is returned from FilterNewRequestOnDeck and is used to iterate over the raw logs and unpacked data for NewRequestOnDeck events raised by the TellorLibrary contract.
type TellorLibraryNewRequestOnDeckIterator struct {
	Event *TellorLibraryNewRequestOnDeck // Event containing the contract specifics and raw log

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
func (it *TellorLibraryNewRequestOnDeckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorLibraryNewRequestOnDeck)
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
		it.Event = new(TellorLibraryNewRequestOnDeck)
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
func (it *TellorLibraryNewRequestOnDeckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorLibraryNewRequestOnDeckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorLibraryNewRequestOnDeck represents a NewRequestOnDeck event raised by the TellorLibrary contract.
type TellorLibraryNewRequestOnDeck struct {
	RequestId       *big.Int
	Query           string
	OnDeckQueryHash [32]byte
	OnDeckTotalTips *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewRequestOnDeck is a free log retrieval operation binding the contract event 0x5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c.
//
// Solidity: event NewRequestOnDeck(uint256 indexed _requestId, string _query, bytes32 _onDeckQueryHash, uint256 _onDeckTotalTips)
func (_TellorLibrary *TellorLibraryFilterer) FilterNewRequestOnDeck(opts *bind.FilterOpts, _requestId []*big.Int) (*TellorLibraryNewRequestOnDeckIterator, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _TellorLibrary.contract.FilterLogs(opts, "NewRequestOnDeck", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &TellorLibraryNewRequestOnDeckIterator{contract: _TellorLibrary.contract, event: "NewRequestOnDeck", logs: logs, sub: sub}, nil
}

// WatchNewRequestOnDeck is a free log subscription operation binding the contract event 0x5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c.
//
// Solidity: event NewRequestOnDeck(uint256 indexed _requestId, string _query, bytes32 _onDeckQueryHash, uint256 _onDeckTotalTips)
func (_TellorLibrary *TellorLibraryFilterer) WatchNewRequestOnDeck(opts *bind.WatchOpts, sink chan<- *TellorLibraryNewRequestOnDeck, _requestId []*big.Int) (event.Subscription, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _TellorLibrary.contract.WatchLogs(opts, "NewRequestOnDeck", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorLibraryNewRequestOnDeck)
				if err := _TellorLibrary.contract.UnpackLog(event, "NewRequestOnDeck", log); err != nil {
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

// ParseNewRequestOnDeck is a log parse operation binding the contract event 0x5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c.
//
// Solidity: event NewRequestOnDeck(uint256 indexed _requestId, string _query, bytes32 _onDeckQueryHash, uint256 _onDeckTotalTips)
func (_TellorLibrary *TellorLibraryFilterer) ParseNewRequestOnDeck(log types.Log) (*TellorLibraryNewRequestOnDeck, error) {
	event := new(TellorLibraryNewRequestOnDeck)
	if err := _TellorLibrary.contract.UnpackLog(event, "NewRequestOnDeck", log); err != nil {
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
	RequestId        *big.Int
	Time             *big.Int
	Value            *big.Int
	TotalTips        *big.Int
	CurrentChallenge [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewValue is a free log retrieval operation binding the contract event 0xc509d04e0782579ee59912139246ccbdf6c36c43abd90591d912017f3467dd16.
//
// Solidity: event NewValue(uint256 indexed _requestId, uint256 _time, uint256 _value, uint256 _totalTips, bytes32 _currentChallenge)
func (_TellorLibrary *TellorLibraryFilterer) FilterNewValue(opts *bind.FilterOpts, _requestId []*big.Int) (*TellorLibraryNewValueIterator, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _TellorLibrary.contract.FilterLogs(opts, "NewValue", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &TellorLibraryNewValueIterator{contract: _TellorLibrary.contract, event: "NewValue", logs: logs, sub: sub}, nil
}

// WatchNewValue is a free log subscription operation binding the contract event 0xc509d04e0782579ee59912139246ccbdf6c36c43abd90591d912017f3467dd16.
//
// Solidity: event NewValue(uint256 indexed _requestId, uint256 _time, uint256 _value, uint256 _totalTips, bytes32 _currentChallenge)
func (_TellorLibrary *TellorLibraryFilterer) WatchNewValue(opts *bind.WatchOpts, sink chan<- *TellorLibraryNewValue, _requestId []*big.Int) (event.Subscription, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _TellorLibrary.contract.WatchLogs(opts, "NewValue", _requestIdRule)
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

// ParseNewValue is a log parse operation binding the contract event 0xc509d04e0782579ee59912139246ccbdf6c36c43abd90591d912017f3467dd16.
//
// Solidity: event NewValue(uint256 indexed _requestId, uint256 _time, uint256 _value, uint256 _totalTips, bytes32 _currentChallenge)
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
	RequestId        *big.Int
	Value            *big.Int
	CurrentChallenge [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNonceSubmitted is a free log retrieval operation binding the contract event 0xe6d63a2aee0aaed2ab49702313ce54114f2145af219d7db30d6624af9e6dffc4.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256 indexed _requestId, uint256 _value, bytes32 _currentChallenge)
func (_TellorLibrary *TellorLibraryFilterer) FilterNonceSubmitted(opts *bind.FilterOpts, _miner []common.Address, _requestId []*big.Int) (*TellorLibraryNonceSubmittedIterator, error) {

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _TellorLibrary.contract.FilterLogs(opts, "NonceSubmitted", _minerRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &TellorLibraryNonceSubmittedIterator{contract: _TellorLibrary.contract, event: "NonceSubmitted", logs: logs, sub: sub}, nil
}

// WatchNonceSubmitted is a free log subscription operation binding the contract event 0xe6d63a2aee0aaed2ab49702313ce54114f2145af219d7db30d6624af9e6dffc4.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256 indexed _requestId, uint256 _value, bytes32 _currentChallenge)
func (_TellorLibrary *TellorLibraryFilterer) WatchNonceSubmitted(opts *bind.WatchOpts, sink chan<- *TellorLibraryNonceSubmitted, _miner []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _TellorLibrary.contract.WatchLogs(opts, "NonceSubmitted", _minerRule, _requestIdRule)
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

// ParseNonceSubmitted is a log parse operation binding the contract event 0xe6d63a2aee0aaed2ab49702313ce54114f2145af219d7db30d6624af9e6dffc4.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256 indexed _requestId, uint256 _value, bytes32 _currentChallenge)
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
const TellorStakeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"NewStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawRequested\",\"type\":\"event\"}]"

// TellorStakeFuncSigs maps the 4-byte function signature to its string representation.
var TellorStakeFuncSigs = map[string]string{
	"820a2d66": "depositStake(TellorStorage.TellorStorageStruct storage)",
	"4601f1cd": "init(TellorStorage.TellorStorageStruct storage)",
	"c9cf5e4c": "requestStakingWithdraw(TellorStorage.TellorStorageStruct storage)",
	"44bacc4b": "withdrawStake(TellorStorage.TellorStorageStruct storage)",
}

// TellorStakeBin is the compiled bytecode used for deploying new contracts.
var TellorStakeBin = "0x61090b610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c806344bacc4b1461005b5780634601f1cd14610087578063820a2d66146100b1578063c9cf5e4c146100db575b600080fd5b81801561006757600080fd5b506100856004803603602081101561007e57600080fd5b5035610105565b005b81801561009357600080fd5b50610085600480360360208110156100aa57600080fd5b5035610173565b8180156100bd57600080fd5b50610085600480360360208110156100d457600080fd5b503561059b565b8180156100e757600080fd5b50610085600480360360208110156100fe57600080fd5b503561060f565b3360009081526047820160205260409020600181015462093a8090620151804206420303101561013457600080fd5b805460021461014257600080fd5b600080825560405133917f4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec91a25050565b60408051600160c01b67646563696d616c73028152815190819003600801902060009081528183016020522054156101aa57600080fd5b3060009081526045820160205260408082208151600160e31b631d6f7b81028152600481019190915269014542ba12a337c00000196024820152905173__$a7780dac84603eae88ac24b827b34ffc9e$__9263eb7bdc089260448082019391829003018186803b15801561021d57600080fd5b505af4158015610231573d6000803e3d6000fd5b5050505061023d6108c1565b506040805160c08101825273e037ec8ec9ec423826750853899394de7f024fee815273cdd8fa31af8475574b8909f135d510579a8087d3602082015273b9dd5afd86547df817da2d0fb89334a6f8edd8919181019190915273230570cd052f40e14c14a81038c6f3aa685d712b6060820152733233afa02644ccd048587f8ba6e99b3c00a34dcc608082015273e010ac6e0248790e08f42d5f697160dedf97e02460a082015260005b60068110156103bf5773__$a7780dac84603eae88ac24b827b34ffc9e$__63eb7bdc0884604501600085856006811061031b57fe5b60200201516001600160a01b03166001600160a01b03168152602001908152602001600020683635c9adc5dea000006040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561038557600080fd5b505af4158015610399573d6000803e3d6000fd5b505050506103b7838383600681106103ad57fe5b602002015161070c565b6001016102e6565b50604080517f746f74616c5f737570706c7900000000000000000000000000000000000000008152815190819003600c908101822060009081528386016020818152858320805469014542ba12a337c00000019055600160c01b67646563696d616c7302855285519485900360080185208352818152858320601290557f7461726765744d696e657273000000000000000000000000000000000000000085528551948590039093018420825280835284822060c89055600160aa1b6a1cdd185ad9505b5bdd5b9d028452845193849003600b0184208252808352848220683635c9adc5dea000009055600160b01b6964697370757465466565028452845193849003600a908101852083528184528583206834957444b840e800009055600160b21b691d1a5b5955185c99d95d02808652865195869003820186208452828552868420610258905585528551948590030190932081529190522054428161052357fe5b604080517f74696d654f664c6173744e657756616c756500000000000000000000000000008152815190819003601201812060009081529582016020818152838820959094064203909455600160b01b69646966666963756c7479028152815190819003600a01902085529190529091206001905550565b6105a5813361070c565b73__$01102ac036bcb1aea51b873d34f7cf54e9$__63e15f6f70826040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b1580156105f457600080fd5b505af4158015610608573d6000803e3d6000fd5b5050505050565b3360009081526047820160205260409020805460011461062e57600080fd5b600281556201518042064203600182015560408051600160aa1b6a1cdd185ad95c90dbdd5b9d028152815190819003600b01812060009081528285016020528281208054600019019055600160e41b630e15f6f702825260048201859052915173__$01102ac036bcb1aea51b873d34f7cf54e9$__9263e15f6f709260248082019391829003018186803b1580156106c557600080fd5b505af41580156106d9573d6000803e3d6000fd5b50506040513392507f453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf9150600090a25050565b60408051600160aa1b6a1cdd185ad9505b5bdd5b9d028152815190819003600b018120600090815282850160209081529083902054600160e01b6393b182b3028352600483018690526001600160a01b0385166024840152925173__$a7780dac84603eae88ac24b827b34ffc9e$__926393b182b3926044808301939192829003018186803b15801561079e57600080fd5b505af41580156107b2573d6000803e3d6000fd5b505050506040513d60208110156107c857600080fd5b505110156107d557600080fd5b6001600160a01b0381166000908152604783016020526040902054158061081657506001600160a01b03811660009081526047830160205260409020546002145b61081f57600080fd5b60408051600160aa1b6a1cdd185ad95c90dbdd5b9d028152815190819003600b01812060009081528483016020908152838220805460019081019091558385018552808452620151804290810690038285019081526001600160a01b038716808552604789019093528584209451855551930192909255915190917f46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e291a25050565b6040518060c00160405280600690602082028038833950919291505056fea165627a7a72305820e4d74c42596b73ed976d06629f73d2d4b0a661a7ebcee80d81e01b938712bf1b0029"

// DeployTellorStake deploys a new Ethereum contract, binding an instance of TellorStake to it.
func DeployTellorStake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorStake, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorStakeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	tellorDisputeAddr, _, _, _ := DeployTellorDispute(auth, backend)
	TellorStakeBin = strings.Replace(TellorStakeBin, "__$01102ac036bcb1aea51b873d34f7cf54e9$__", tellorDisputeAddr.String()[2:], -1)

	tellorTransferAddr, _, _, _ := DeployTellorTransfer(auth, backend)
	TellorStakeBin = strings.Replace(TellorStakeBin, "__$a7780dac84603eae88ac24b827b34ffc9e$__", tellorTransferAddr.String()[2:], -1)

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
var TellorStorageBin = "0x604c6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820515486cb4b39fbc1385b8760532b75d5c47ac362c649b7f9014d66fd04861a9c0029"

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
const TellorTransferABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// TellorTransferFuncSigs maps the 4-byte function signature to its string representation.
var TellorTransferFuncSigs = map[string]string{
	"bf32006c": "allowance(TellorStorage.TellorStorageStruct storage,address,address)",
	"acaab9e2": "allowedToTrade(TellorStorage.TellorStorageStruct storage,address,uint256)",
	"850dcc32": "approve(TellorStorage.TellorStorageStruct storage,address,uint256)",
	"93b182b3": "balanceOf(TellorStorage.TellorStorageStruct storage,address)",
	"3f48b1ff": "balanceOfAt(TellorStorage.TellorStorageStruct storage,address,uint256)",
	"c7bb46ad": "doTransfer(TellorStorage.TellorStorageStruct storage,address,address,uint256)",
	"9be5647f": "getBalanceAt(TellorStorage.Checkpoint[] storage,uint256)",
	"c84b96f5": "transfer(TellorStorage.TellorStorageStruct storage,address,uint256)",
	"ca501899": "transferFrom(TellorStorage.TellorStorageStruct storage,address,address,uint256)",
	"eb7bdc08": "updateBalanceAtNow(TellorStorage.Checkpoint[] storage,uint256)",
}

// TellorTransferBin is the compiled bytecode used for deploying new contracts.
var TellorTransferBin = "0x610930610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100a85760003560e01c8063bf32006c11610070578063bf32006c146101c5578063c7bb46ad146101f9578063c84b96f514610244578063ca50189914610283578063eb7bdc08146102cc576100a8565b80633f48b1ff146100ad578063850dcc32146100f157806393b182b3146101445780639be5647f14610170578063acaab9e214610193575b600080fd5b6100df600480360360608110156100c357600080fd5b508035906001600160a01b0360208201351690604001356102fc565b60408051918252519081900360200190f35b8180156100fd57600080fd5b506101306004803603606081101561011457600080fd5b508035906001600160a01b036020820135169060400135610395565b604080519115158252519081900360200190f35b6100df6004803603604081101561015a57600080fd5b50803590602001356001600160a01b0316610428565b6100df6004803603604081101561018657600080fd5b508035906020013561043e565b610130600480360360608110156101a957600080fd5b508035906001600160a01b03602082013516906040013561056e565b6100df600480360360608110156101db57600080fd5b508035906001600160a01b0360208201358116916040013516610617565b81801561020557600080fd5b506102426004803603608081101561021c57600080fd5b508035906001600160a01b03602082013581169160408101359091169060600135610644565b005b81801561025057600080fd5b506101306004803603606081101561026757600080fd5b508035906001600160a01b036020820135169060400135610742565b81801561028f57600080fd5b50610130600480360360808110156102a657600080fd5b508035906001600160a01b0360208201358116916040810135909116906060013561075a565b8180156102d857600080fd5b50610242600480360360408110156102ef57600080fd5b50803590602001356107cf565b6001600160a01b0382166000908152604584016020526040812054158061035a57506001600160a01b03831660009081526045850160205260408120805484929061034357fe5b6000918252602090912001546001600160801b0316115b156103675750600061038e565b6001600160a01b0383166000908152604585016020526040902061038b908361043e565b90505b9392505050565b60006103a284338461056e565b6103ab57600080fd5b6001600160a01b0383166103be57600080fd5b33600081815260468601602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b60006104358383436102fc565b90505b92915050565b815460009061044f57506000610438565b82548390600019810190811061046157fe5b6000918252602090912001546001600160801b031682106104b15782548390600019810190811061048e57fe5b600091825260209091200154600160801b90046001600160801b03169050610438565b826000815481106104be57fe5b6000918252602090912001546001600160801b03168210156104e257506000610438565b8254600090600019015b8181111561053d57600060026001838501010490508486828154811061050e57fe5b6000918252602090912001546001600160801b03161161053057809250610537565b6001810391505b506104ec565b84828154811061054957fe5b600091825260209091200154600160801b90046001600160801b031695945050505050565b6001600160a01b0382166000908152604784016020526040812054156105f05760408051600160aa1b6a1cdd185ad9505b5bdd5b9d028152815190819003600b0190206000908152818601602052908120546105de9084906105d290818989610428565b9063ffffffff6108a816565b106105eb5750600161038e565b61060d565b6000610600836105d28787610428565b1061060d5750600161038e565b5060009392505050565b6001600160a01b039182166000908152604693909301602090815260408085209290931684525290205490565b6000811161065157600080fd5b6001600160a01b03821661066457600080fd5b61066f84848361056e565b61067857600080fd5b60006106858585436102fc565b6001600160a01b038516600090815260458701602052604090209091506106ae908383036107cf565b6106b98584436102fc565b90508082820110156106ca57600080fd5b6001600160a01b038316600090815260458601602052604090206106f0908284016107cf565b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a35050505050565b600061075084338585610644565b5060019392505050565b6001600160a01b0383166000908152604685016020908152604080832033845290915281205482111561078c57600080fd5b6001600160a01b038416600090815260468601602090815260408083203384529091529020805483900390556107c485858585610644565b506001949350505050565b81541580610803575081544390839060001981019081106107ec57fe5b6000918252602090912001546001600160801b0316105b1561086a578154600090839061081c82600183016108ba565b8154811061082657fe5b600091825260209091200180546001600160801b03848116600160801b024382166fffffffffffffffffffffffffffffffff199093169290921716179055506108a4565b81546000908390600019810190811061087f57fe5b600091825260209091200180546001600160801b03808516600160801b029116179055505b5050565b6000828211156108b457fe5b50900390565b8154818355818111156108de576000838152602090206108de9181019083016108e3565b505050565b61090191905b808211156108fd57600081556001016108e9565b5090565b9056fea165627a7a72305820c6c295baf81bf26c7b082884cc9c10ad50f641862df6f0474f13212da204a9fe0029"

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
var UtilitiesBin = "0x604c6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a723058206c63a2ce0342c420ef12cd2eeb9c403ca61321488bc0ae195b58425d8ee4f84f0029"

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
