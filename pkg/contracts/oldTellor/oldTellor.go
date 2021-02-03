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

// OldTellorABI is the input ABI used to generate the binding from.
const OldTellorABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"theLazyCoon\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minerIndex\",\"type\":\"uint256\"}],\"name\":\"beginDispute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_supportsDispute\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"tallyVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_propNewTellorAddress\",\"type\":\"address\"}],\"name\":\"proposeFork\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_c_sapi\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_c_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_granularity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"requestData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"submitMiningSolution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"testSubmitMiningSolution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_pendingOwner\",\"type\":\"address\"}],\"name\":\"proposeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"requestStakingWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OldTellorBin is the compiled bytecode used for deploying new contracts.
var OldTellorBin = "0x608060405234801561001057600080fd5b506114d7806100206000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c806368c180d5116100a2578063a9059cbb11610071578063a9059cbb146104b9578063b079f64a1461051f578063bed9d8611461056d578063c0a8b65014610577578063c9d27afe146106045761010b565b806368c180d51461036e578063710bf322146103fb578063752d49a11461043f5780638581af19146104775761010b565b806328449c3a116100de57806328449c3a1461024a5780633fff2816146102545780634d318b0e146103365780634e71e0c8146103645761010b565b8063095ea7b3146101105780630d2d76a21461017657806323b872dd1461018057806326b7d9f614610206575b600080fd5b61015c6004803603604081101561012657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061063e565b604051808215151515815260200191505060405180910390f35b61017e610710565b005b6101ec6004803603606081101561019657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061077c565b604051808215151515815260200191505060405180910390f35b6102486004803603602081101561021c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610883565b005b610252610924565b005b6103346004803603608081101561026a57600080fd5b810190808035906020019064010000000081111561028757600080fd5b82018360208201111561029957600080fd5b803590602001918460018302840111640100000000831117156102bb57600080fd5b9091929391929390803590602001906401000000008111156102dc57600080fd5b8201836020820111156102ee57600080fd5b8035906020019184600183028401116401000000008311171561031057600080fd5b90919293919293908035906020019092919080359060200190929190505050610990565b005b6103626004803603602081101561034c57600080fd5b8101908080359060200190929190505050610a7a565b005b61036c610aef565b005b6103f96004803603606081101561038457600080fd5b81019080803590602001906401000000008111156103a157600080fd5b8201836020820111156103b357600080fd5b803590602001918460018302840111640100000000831117156103d557600080fd5b90919293919293908035906020019092919080359060200190929190505050610afb565b005b61043d6004803603602081101561041157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610baf565b005b6104756004803603604081101561045557600080fd5b810190808035906020019092919080359060200190929190505050610bc6565b005b6104b76004803603606081101561048d57600080fd5b81019080803590602001909291908035906020019092919080359060200190929190505050610c44565b005b610505600480360360408110156104cf57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610ccb565b604051808215151515815260200191505060405180910390f35b61056b6004803603604081101561053557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d9d565b005b610575610e47565b005b6106026004803603606081101561058d57600080fd5b81019080803590602001906401000000008111156105aa57600080fd5b8201836020820111156105bc57600080fd5b803590602001918460018302840111640100000000831117156105de57600080fd5b90919293919293908035906020019092919080359060200190929190505050610eb3565b005b61063c6004803603604081101561061a57600080fd5b8101908080359060200190929190803515159060200190929190505050610f67565b005b60008073__OldTellorTransfer_____________________63cb6f7c66909185856040518463ffffffff1660e01b8152600401808481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060206040518083038186803b1580156106cd57600080fd5b505af41580156106e1573d6000803e3d6000fd5b505050506040513d60208110156106f757600080fd5b8101908080519060200190929190505050905092915050565b600073__OldTellorStake________________________636d09e1b490916040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561076257600080fd5b505af4158015610776573d6000803e3d6000fd5b50505050565b60008073__OldTellorTransfer_____________________6335ff14bd90918686866040518563ffffffff1660e01b8152600401808581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200194505050505060206040518083038186803b15801561083f57600080fd5b505af4158015610853573d6000803e3d6000fd5b505050506040513d602081101561086957600080fd5b810190808051906020019092919050505090509392505050565b600073__OldTellorDispute______________________63eaeddad99091836040518363ffffffff1660e01b8152600401808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060006040518083038186803b15801561090957600080fd5b505af415801561091d573d6000803e3d6000fd5b5050505050565b600073__OldTellorStake________________________633c06d13990916040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561097657600080fd5b505af415801561098a573d6000803e3d6000fd5b50505050565b600073__OldTellorLibrary______________________63421ab07290918888888888886040518863ffffffff1660e01b81526004018088815260200180602001806020018581526020018481526020018381038352898982818152602001925080828437600081840152601f19601f8201169050808301925050508381038252878782818152602001925080828437600081840152601f19601f820116905080830192505050995050505050505050505060006040518083038186803b158015610a5a57600080fd5b505af4158015610a6e573d6000803e3d6000fd5b50505050505050505050565b600073__OldTellorDispute______________________63a513e1299091836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b158015610ad457600080fd5b505af4158015610ae8573d6000803e3d6000fd5b5050505050565b610af96000610fe9565b565b600073__OldTellorLibrary______________________63c38a5c8a9091868686866040518663ffffffff1660e01b815260040180868152602001806020018481526020018381526020018281038252868682818152602001925080828437600081840152601f19601f820116905080830192505050965050505050505060006040518083038186803b158015610b9157600080fd5b505af4158015610ba5573d6000803e3d6000fd5b5050505050505050565b610bc38160006112af90919063ffffffff16565b50565b600073__OldTellorLibrary______________________6363687386909184846040518463ffffffff1660e01b815260040180848152602001838152602001828152602001935050505060006040518083038186803b158015610c2857600080fd5b505af4158015610c3c573d6000803e3d6000fd5b505050505050565b600073__OldTellorDispute______________________630edf76e490918585856040518563ffffffff1660e01b81526004018085815260200184815260200183815260200182815260200194505050505060006040518083038186803b158015610cae57600080fd5b505af4158015610cc2573d6000803e3d6000fd5b50505050505050565b60008073__OldTellorTransfer_____________________631e405eda909185856040518463ffffffff1660e01b8152600401808481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060206040518083038186803b158015610d5a57600080fd5b505af4158015610d6e573d6000803e3d6000fd5b505050506040513d6020811015610d8457600080fd5b8101908080519060200190929190505050905092915050565b600073__OldTellorLibrary______________________63fa0837b1909184846040518463ffffffff1660e01b8152600401808481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060006040518083038186803b158015610e2b57600080fd5b505af4158015610e3f573d6000803e3d6000fd5b505050505050565b600073__OldTellorStake________________________63e62c164590916040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b158015610e9957600080fd5b505af4158015610ead573d6000803e3d6000fd5b50505050565b600073__OldTellorLibrary______________________63af026f0d9091868686866040518663ffffffff1660e01b815260040180868152602001806020018481526020018381526020018281038252868682818152602001925080828437600081840152601f19601f820116905080830192505050965050505050505060006040518083038186803b158015610f4957600080fd5b505af4158015610f5d573d6000803e3d6000fd5b5050505050505050565b600073__OldTellorDispute______________________632e8dcb7e909184846040518463ffffffff1660e01b81526004018084815260200183815260200182151515158152602001935050505060006040518083038186803b158015610fcd57600080fd5b505af4158015610fe1573d6000803e3d6000fd5b505050505050565b80603f01600060405180807f70656e64696e675f6f776e657200000000000000000000000000000000000000815250600d0190506040518091039020815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461108b57600080fd5b80603f01600060405180807f70656e64696e675f6f776e657200000000000000000000000000000000000000815250600d0190506040518091039020815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681603f01600060405180807f5f6f776e6572000000000000000000000000000000000000000000000000000081525060060190506040518091039020815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fd02dfe11dd36892231b9ad3e9a9d48153f493acd9af79de7817ac4dac150423660405160405180910390a380603f01600060405180807f70656e64696e675f6f776e657200000000000000000000000000000000000000815250600d0190506040518091039020815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681603f01600060405180807f5f6f776e6572000000000000000000000000000000000000000000000000000081525060060190506040518091039020815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b81603f01600060405180807f5f6f776e6572000000000000000000000000000000000000000000000000000081525060060190506040518091039020815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461135157600080fd5b8073ffffffffffffffffffffffffffffffffffffffff1682603f01600060405180807f5f6f776e6572000000000000000000000000000000000000000000000000000081525060060190506040518091039020815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f496503062c30a72cf552a695690d2219ad04e8af1524dd5c10861634ed11f90b60405160405180910390a38082603f01600060405180807f70656e64696e675f6f776e657200000000000000000000000000000000000000815250600d0190506040518091039020815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505056fea265627a7a723158202625c5d282ae282fc2d531de164831086f3e084838977e328ed7db911c4f238564736f6c63430005100032"

// DeployOldTellor deploys a new Ethereum contract, binding an instance of OldTellor to it.
func DeployOldTellor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OldTellor, error) {
	parsed, err := abi.JSON(strings.NewReader(OldTellorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OldTellorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OldTellor{OldTellorCaller: OldTellorCaller{contract: contract}, OldTellorTransactor: OldTellorTransactor{contract: contract}, OldTellorFilterer: OldTellorFilterer{contract: contract}}, nil
}

// OldTellor is an auto generated Go binding around an Ethereum contract.
type OldTellor struct {
	OldTellorCaller     // Read-only binding to the contract
	OldTellorTransactor // Write-only binding to the contract
	OldTellorFilterer   // Log filterer for contract events
}

// OldTellorCaller is an auto generated read-only Go binding around an Ethereum contract.
type OldTellorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldTellorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OldTellorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldTellorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OldTellorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldTellorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OldTellorSession struct {
	Contract     *OldTellor        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OldTellorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OldTellorCallerSession struct {
	Contract *OldTellorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OldTellorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OldTellorTransactorSession struct {
	Contract     *OldTellorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OldTellorRaw is an auto generated low-level Go binding around an Ethereum contract.
type OldTellorRaw struct {
	Contract *OldTellor // Generic contract binding to access the raw methods on
}

// OldTellorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OldTellorCallerRaw struct {
	Contract *OldTellorCaller // Generic read-only contract binding to access the raw methods on
}

// OldTellorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OldTellorTransactorRaw struct {
	Contract *OldTellorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOldTellor creates a new instance of OldTellor, bound to a specific deployed contract.
func NewOldTellor(address common.Address, backend bind.ContractBackend) (*OldTellor, error) {
	contract, err := bindOldTellor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OldTellor{OldTellorCaller: OldTellorCaller{contract: contract}, OldTellorTransactor: OldTellorTransactor{contract: contract}, OldTellorFilterer: OldTellorFilterer{contract: contract}}, nil
}

// NewOldTellorCaller creates a new read-only instance of OldTellor, bound to a specific deployed contract.
func NewOldTellorCaller(address common.Address, caller bind.ContractCaller) (*OldTellorCaller, error) {
	contract, err := bindOldTellor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OldTellorCaller{contract: contract}, nil
}

// NewOldTellorTransactor creates a new write-only instance of OldTellor, bound to a specific deployed contract.
func NewOldTellorTransactor(address common.Address, transactor bind.ContractTransactor) (*OldTellorTransactor, error) {
	contract, err := bindOldTellor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OldTellorTransactor{contract: contract}, nil
}

// NewOldTellorFilterer creates a new log filterer instance of OldTellor, bound to a specific deployed contract.
func NewOldTellorFilterer(address common.Address, filterer bind.ContractFilterer) (*OldTellorFilterer, error) {
	contract, err := bindOldTellor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OldTellorFilterer{contract: contract}, nil
}

// bindOldTellor binds a generic wrapper to an already deployed contract.
func bindOldTellor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OldTellorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OldTellor *OldTellorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OldTellor.Contract.OldTellorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OldTellor *OldTellorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OldTellor.Contract.OldTellorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OldTellor *OldTellorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OldTellor.Contract.OldTellorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OldTellor *OldTellorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OldTellor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OldTellor *OldTellorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OldTellor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OldTellor *OldTellorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OldTellor.Contract.contract.Transact(opts, method, params...)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_OldTellor *OldTellorTransactor) AddTip(opts *bind.TransactOpts, _requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _OldTellor.contract.Transact(opts, "addTip", _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_OldTellor *OldTellorSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _OldTellor.Contract.AddTip(&_OldTellor.TransactOpts, _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_OldTellor *OldTellorTransactorSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _OldTellor.Contract.AddTip(&_OldTellor.TransactOpts, _requestId, _tip)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_OldTellor *OldTellorTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OldTellor.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_OldTellor *OldTellorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OldTellor.Contract.Approve(&_OldTellor.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_OldTellor *OldTellorTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OldTellor.Contract.Approve(&_OldTellor.TransactOpts, _spender, _amount)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_OldTellor *OldTellorTransactor) BeginDispute(opts *bind.TransactOpts, _requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _OldTellor.contract.Transact(opts, "beginDispute", _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_OldTellor *OldTellorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _OldTellor.Contract.BeginDispute(&_OldTellor.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_OldTellor *OldTellorTransactorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _OldTellor.Contract.BeginDispute(&_OldTellor.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_OldTellor *OldTellorTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OldTellor.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_OldTellor *OldTellorSession) ClaimOwnership() (*types.Transaction, error) {
	return _OldTellor.Contract.ClaimOwnership(&_OldTellor.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_OldTellor *OldTellorTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _OldTellor.Contract.ClaimOwnership(&_OldTellor.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_OldTellor *OldTellorTransactor) DepositStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OldTellor.contract.Transact(opts, "depositStake")
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_OldTellor *OldTellorSession) DepositStake() (*types.Transaction, error) {
	return _OldTellor.Contract.DepositStake(&_OldTellor.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_OldTellor *OldTellorTransactorSession) DepositStake() (*types.Transaction, error) {
	return _OldTellor.Contract.DepositStake(&_OldTellor.TransactOpts)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_OldTellor *OldTellorTransactor) ProposeFork(opts *bind.TransactOpts, _propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _OldTellor.contract.Transact(opts, "proposeFork", _propNewTellorAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_OldTellor *OldTellorSession) ProposeFork(_propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _OldTellor.Contract.ProposeFork(&_OldTellor.TransactOpts, _propNewTellorAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_OldTellor *OldTellorTransactorSession) ProposeFork(_propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _OldTellor.Contract.ProposeFork(&_OldTellor.TransactOpts, _propNewTellorAddress)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _pendingOwner) returns()
func (_OldTellor *OldTellorTransactor) ProposeOwnership(opts *bind.TransactOpts, _pendingOwner common.Address) (*types.Transaction, error) {
	return _OldTellor.contract.Transact(opts, "proposeOwnership", _pendingOwner)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _pendingOwner) returns()
func (_OldTellor *OldTellorSession) ProposeOwnership(_pendingOwner common.Address) (*types.Transaction, error) {
	return _OldTellor.Contract.ProposeOwnership(&_OldTellor.TransactOpts, _pendingOwner)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _pendingOwner) returns()
func (_OldTellor *OldTellorTransactorSession) ProposeOwnership(_pendingOwner common.Address) (*types.Transaction, error) {
	return _OldTellor.Contract.ProposeOwnership(&_OldTellor.TransactOpts, _pendingOwner)
}

// RequestData is a paid mutator transaction binding the contract method 0x3fff2816.
//
// Solidity: function requestData(string _c_sapi, string _c_symbol, uint256 _granularity, uint256 _tip) returns()
func (_OldTellor *OldTellorTransactor) RequestData(opts *bind.TransactOpts, _c_sapi string, _c_symbol string, _granularity *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _OldTellor.contract.Transact(opts, "requestData", _c_sapi, _c_symbol, _granularity, _tip)
}

// RequestData is a paid mutator transaction binding the contract method 0x3fff2816.
//
// Solidity: function requestData(string _c_sapi, string _c_symbol, uint256 _granularity, uint256 _tip) returns()
func (_OldTellor *OldTellorSession) RequestData(_c_sapi string, _c_symbol string, _granularity *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _OldTellor.Contract.RequestData(&_OldTellor.TransactOpts, _c_sapi, _c_symbol, _granularity, _tip)
}

// RequestData is a paid mutator transaction binding the contract method 0x3fff2816.
//
// Solidity: function requestData(string _c_sapi, string _c_symbol, uint256 _granularity, uint256 _tip) returns()
func (_OldTellor *OldTellorTransactorSession) RequestData(_c_sapi string, _c_symbol string, _granularity *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _OldTellor.Contract.RequestData(&_OldTellor.TransactOpts, _c_sapi, _c_symbol, _granularity, _tip)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_OldTellor *OldTellorTransactor) RequestStakingWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OldTellor.contract.Transact(opts, "requestStakingWithdraw")
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_OldTellor *OldTellorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _OldTellor.Contract.RequestStakingWithdraw(&_OldTellor.TransactOpts)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_OldTellor *OldTellorTransactorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _OldTellor.Contract.RequestStakingWithdraw(&_OldTellor.TransactOpts)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_OldTellor *OldTellorTransactor) SubmitMiningSolution(opts *bind.TransactOpts, _nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _OldTellor.contract.Transact(opts, "submitMiningSolution", _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_OldTellor *OldTellorSession) SubmitMiningSolution(_nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _OldTellor.Contract.SubmitMiningSolution(&_OldTellor.TransactOpts, _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_OldTellor *OldTellorTransactorSession) SubmitMiningSolution(_nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _OldTellor.Contract.SubmitMiningSolution(&_OldTellor.TransactOpts, _nonce, _requestId, _value)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_OldTellor *OldTellorTransactor) TallyVotes(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _OldTellor.contract.Transact(opts, "tallyVotes", _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_OldTellor *OldTellorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _OldTellor.Contract.TallyVotes(&_OldTellor.TransactOpts, _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_OldTellor *OldTellorTransactorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _OldTellor.Contract.TallyVotes(&_OldTellor.TransactOpts, _disputeId)
}

// TestSubmitMiningSolution is a paid mutator transaction binding the contract method 0xc0a8b650.
//
// Solidity: function testSubmitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_OldTellor *OldTellorTransactor) TestSubmitMiningSolution(opts *bind.TransactOpts, _nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _OldTellor.contract.Transact(opts, "testSubmitMiningSolution", _nonce, _requestId, _value)
}

// TestSubmitMiningSolution is a paid mutator transaction binding the contract method 0xc0a8b650.
//
// Solidity: function testSubmitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_OldTellor *OldTellorSession) TestSubmitMiningSolution(_nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _OldTellor.Contract.TestSubmitMiningSolution(&_OldTellor.TransactOpts, _nonce, _requestId, _value)
}

// TestSubmitMiningSolution is a paid mutator transaction binding the contract method 0xc0a8b650.
//
// Solidity: function testSubmitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_OldTellor *OldTellorTransactorSession) TestSubmitMiningSolution(_nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _OldTellor.Contract.TestSubmitMiningSolution(&_OldTellor.TransactOpts, _nonce, _requestId, _value)
}

// TheLazyCoon is a paid mutator transaction binding the contract method 0xb079f64a.
//
// Solidity: function theLazyCoon(address _address, uint256 _amount) returns()
func (_OldTellor *OldTellorTransactor) TheLazyCoon(opts *bind.TransactOpts, _address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OldTellor.contract.Transact(opts, "theLazyCoon", _address, _amount)
}

// TheLazyCoon is a paid mutator transaction binding the contract method 0xb079f64a.
//
// Solidity: function theLazyCoon(address _address, uint256 _amount) returns()
func (_OldTellor *OldTellorSession) TheLazyCoon(_address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OldTellor.Contract.TheLazyCoon(&_OldTellor.TransactOpts, _address, _amount)
}

// TheLazyCoon is a paid mutator transaction binding the contract method 0xb079f64a.
//
// Solidity: function theLazyCoon(address _address, uint256 _amount) returns()
func (_OldTellor *OldTellorTransactorSession) TheLazyCoon(_address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OldTellor.Contract.TheLazyCoon(&_OldTellor.TransactOpts, _address, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_OldTellor *OldTellorTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OldTellor.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_OldTellor *OldTellorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OldTellor.Contract.Transfer(&_OldTellor.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_OldTellor *OldTellorTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OldTellor.Contract.Transfer(&_OldTellor.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_OldTellor *OldTellorTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OldTellor.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_OldTellor *OldTellorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OldTellor.Contract.TransferFrom(&_OldTellor.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_OldTellor *OldTellorTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OldTellor.Contract.TransferFrom(&_OldTellor.TransactOpts, _from, _to, _amount)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_OldTellor *OldTellorTransactor) Vote(opts *bind.TransactOpts, _disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _OldTellor.contract.Transact(opts, "vote", _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_OldTellor *OldTellorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _OldTellor.Contract.Vote(&_OldTellor.TransactOpts, _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_OldTellor *OldTellorTransactorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _OldTellor.Contract.Vote(&_OldTellor.TransactOpts, _disputeId, _supportsDispute)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_OldTellor *OldTellorTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OldTellor.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_OldTellor *OldTellorSession) WithdrawStake() (*types.Transaction, error) {
	return _OldTellor.Contract.WithdrawStake(&_OldTellor.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_OldTellor *OldTellorTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _OldTellor.Contract.WithdrawStake(&_OldTellor.TransactOpts)
}
