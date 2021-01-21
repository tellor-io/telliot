// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balancer

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

// BBronzeABI is the input ABI used to generate the binding from.
const BBronzeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BBronzeFuncSigs maps the 4-byte function signature to its string representation.
var BBronzeFuncSigs = map[string]string{
	"9a86139b": "getColor()",
}

// BBronzeBin is the compiled bytecode used for deploying new contracts.
var BBronzeBin = "0x6080604052348015600f57600080fd5b5060878061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80639a86139b14602d575b600080fd5b60336045565b60408051918252519081900360200190f35b6542524f4e5a4560d01b9056fea265627a7a723158206ae941f700e575b4f48f4780eea8e9989fe1e22bf52c493ea81d74a6e973a6a564736f6c634300050c0032"

// DeployBBronze deploys a new Ethereum contract, binding an instance of BBronze to it.
func DeployBBronze(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BBronze, error) {
	parsed, err := abi.JSON(strings.NewReader(BBronzeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BBronzeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BBronze{BBronzeCaller: BBronzeCaller{contract: contract}, BBronzeTransactor: BBronzeTransactor{contract: contract}, BBronzeFilterer: BBronzeFilterer{contract: contract}}, nil
}

// BBronze is an auto generated Go binding around an Ethereum contract.
type BBronze struct {
	BBronzeCaller     // Read-only binding to the contract
	BBronzeTransactor // Write-only binding to the contract
	BBronzeFilterer   // Log filterer for contract events
}

// BBronzeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BBronzeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BBronzeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BBronzeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BBronzeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BBronzeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BBronzeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BBronzeSession struct {
	Contract     *BBronze          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BBronzeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BBronzeCallerSession struct {
	Contract *BBronzeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BBronzeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BBronzeTransactorSession struct {
	Contract     *BBronzeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BBronzeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BBronzeRaw struct {
	Contract *BBronze // Generic contract binding to access the raw methods on
}

// BBronzeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BBronzeCallerRaw struct {
	Contract *BBronzeCaller // Generic read-only contract binding to access the raw methods on
}

// BBronzeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BBronzeTransactorRaw struct {
	Contract *BBronzeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBBronze creates a new instance of BBronze, bound to a specific deployed contract.
func NewBBronze(address common.Address, backend bind.ContractBackend) (*BBronze, error) {
	contract, err := bindBBronze(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BBronze{BBronzeCaller: BBronzeCaller{contract: contract}, BBronzeTransactor: BBronzeTransactor{contract: contract}, BBronzeFilterer: BBronzeFilterer{contract: contract}}, nil
}

// NewBBronzeCaller creates a new read-only instance of BBronze, bound to a specific deployed contract.
func NewBBronzeCaller(address common.Address, caller bind.ContractCaller) (*BBronzeCaller, error) {
	contract, err := bindBBronze(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BBronzeCaller{contract: contract}, nil
}

// NewBBronzeTransactor creates a new write-only instance of BBronze, bound to a specific deployed contract.
func NewBBronzeTransactor(address common.Address, transactor bind.ContractTransactor) (*BBronzeTransactor, error) {
	contract, err := bindBBronze(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BBronzeTransactor{contract: contract}, nil
}

// NewBBronzeFilterer creates a new log filterer instance of BBronze, bound to a specific deployed contract.
func NewBBronzeFilterer(address common.Address, filterer bind.ContractFilterer) (*BBronzeFilterer, error) {
	contract, err := bindBBronze(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BBronzeFilterer{contract: contract}, nil
}

// bindBBronze binds a generic wrapper to an already deployed contract.
func bindBBronze(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BBronzeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BBronze *BBronzeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BBronze.Contract.BBronzeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BBronze *BBronzeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BBronze.Contract.BBronzeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BBronze *BBronzeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BBronze.Contract.BBronzeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BBronze *BBronzeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BBronze.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BBronze *BBronzeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BBronze.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BBronze *BBronzeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BBronze.Contract.contract.Transact(opts, method, params...)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BBronze *BBronzeCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BBronze.contract.Call(opts, &out, "getColor")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BBronze *BBronzeSession) GetColor() ([32]byte, error) {
	return _BBronze.Contract.GetColor(&_BBronze.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BBronze *BBronzeCallerSession) GetColor() ([32]byte, error) {
	return _BBronze.Contract.GetColor(&_BBronze.CallOpts)
}

// BColorABI is the input ABI used to generate the binding from.
const BColorABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BColorFuncSigs maps the 4-byte function signature to its string representation.
var BColorFuncSigs = map[string]string{
	"9a86139b": "getColor()",
}

// BColor is an auto generated Go binding around an Ethereum contract.
type BColor struct {
	BColorCaller     // Read-only binding to the contract
	BColorTransactor // Write-only binding to the contract
	BColorFilterer   // Log filterer for contract events
}

// BColorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BColorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BColorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BColorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BColorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BColorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BColorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BColorSession struct {
	Contract     *BColor           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BColorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BColorCallerSession struct {
	Contract *BColorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BColorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BColorTransactorSession struct {
	Contract     *BColorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BColorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BColorRaw struct {
	Contract *BColor // Generic contract binding to access the raw methods on
}

// BColorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BColorCallerRaw struct {
	Contract *BColorCaller // Generic read-only contract binding to access the raw methods on
}

// BColorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BColorTransactorRaw struct {
	Contract *BColorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBColor creates a new instance of BColor, bound to a specific deployed contract.
func NewBColor(address common.Address, backend bind.ContractBackend) (*BColor, error) {
	contract, err := bindBColor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BColor{BColorCaller: BColorCaller{contract: contract}, BColorTransactor: BColorTransactor{contract: contract}, BColorFilterer: BColorFilterer{contract: contract}}, nil
}

// NewBColorCaller creates a new read-only instance of BColor, bound to a specific deployed contract.
func NewBColorCaller(address common.Address, caller bind.ContractCaller) (*BColorCaller, error) {
	contract, err := bindBColor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BColorCaller{contract: contract}, nil
}

// NewBColorTransactor creates a new write-only instance of BColor, bound to a specific deployed contract.
func NewBColorTransactor(address common.Address, transactor bind.ContractTransactor) (*BColorTransactor, error) {
	contract, err := bindBColor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BColorTransactor{contract: contract}, nil
}

// NewBColorFilterer creates a new log filterer instance of BColor, bound to a specific deployed contract.
func NewBColorFilterer(address common.Address, filterer bind.ContractFilterer) (*BColorFilterer, error) {
	contract, err := bindBColor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BColorFilterer{contract: contract}, nil
}

// bindBColor binds a generic wrapper to an already deployed contract.
func bindBColor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BColorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BColor *BColorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BColor.Contract.BColorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BColor *BColorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BColor.Contract.BColorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BColor *BColorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BColor.Contract.BColorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BColor *BColorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BColor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BColor *BColorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BColor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BColor *BColorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BColor.Contract.contract.Transact(opts, method, params...)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BColor *BColorCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BColor.contract.Call(opts, &out, "getColor")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BColor *BColorSession) GetColor() ([32]byte, error) {
	return _BColor.Contract.GetColor(&_BColor.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BColor *BColorCallerSession) GetColor() ([32]byte, error) {
	return _BColor.Contract.GetColor(&_BColor.CallOpts)
}

// BConstABI is the input ABI used to generate the binding from.
const BConstABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"BONE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BPOW_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EXIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INIT_POOL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_IN_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OUT_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_TOTAL_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BConstFuncSigs maps the 4-byte function signature to its string representation.
var BConstFuncSigs = map[string]string{
	"c36596a6": "BONE()",
	"189d00ca": "BPOW_PRECISION()",
	"c6580d12": "EXIT_FEE()",
	"9381cd2b": "INIT_POOL_SUPPLY()",
	"b0e0d136": "MAX_BOUND_TOKENS()",
	"bc694ea2": "MAX_BPOW_BASE()",
	"bc063e1a": "MAX_FEE()",
	"ec093021": "MAX_IN_RATIO()",
	"992e2a92": "MAX_OUT_RATIO()",
	"09a3bbe4": "MAX_TOTAL_WEIGHT()",
	"e4a28a52": "MAX_WEIGHT()",
	"867378c5": "MIN_BALANCE()",
	"b7b800a4": "MIN_BOUND_TOKENS()",
	"ba019dab": "MIN_BPOW_BASE()",
	"76c7a3c7": "MIN_FEE()",
	"218b5382": "MIN_WEIGHT()",
	"9a86139b": "getColor()",
}

// BConstBin is the compiled bytecode used for deploying new contracts.
var BConstBin = "0x608060405234801561001057600080fd5b50610288806100206000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c8063b0e0d136116100a2578063bc694ea211610071578063bc694ea214610182578063c36596a614610132578063c6580d121461018a578063e4a28a5214610110578063ec093021146101925761010b565b8063b0e0d13614610162578063b7b800a41461016a578063ba019dab14610172578063bc063e1a1461017a5761010b565b8063867378c5116100de578063867378c5146101425780639381cd2b1461014a578063992e2a92146101525780639a86139b1461015a5761010b565b806309a3bbe414610110578063189d00ca1461012a578063218b53821461013257806376c7a3c71461013a575b600080fd5b61011861019a565b60408051918252519081900360200190f35b6101186101a7565b6101186101bb565b6101186101c7565b6101186101d9565b6101186101ed565b6101186101fa565b610118610206565b610118610213565b610118610218565b61011861021d565b610118610222565b610118610232565b61011861023e565b610118610243565b6802b5e3af16b188000081565b6402540be400670de0b6b3a76400005b0481565b670de0b6b3a764000081565b620f4240670de0b6b3a76400006101b7565b64e8d4a51000670de0b6b3a76400006101b7565b68056bc75e2d6310000081565b6704a03ce68d21555681565b6542524f4e5a4560d01b90565b600881565b600281565b600181565b600a670de0b6b3a76400006101b7565b671bc16d674ec7ffff81565b600081565b6002670de0b6b3a76400006101b756fea265627a7a723158204011a9bb547cf731e85788e3d3db6a03e86a0285e559699f1ecf9b9e9433d78c64736f6c634300050c0032"

// DeployBConst deploys a new Ethereum contract, binding an instance of BConst to it.
func DeployBConst(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BConst, error) {
	parsed, err := abi.JSON(strings.NewReader(BConstABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BConstBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BConst{BConstCaller: BConstCaller{contract: contract}, BConstTransactor: BConstTransactor{contract: contract}, BConstFilterer: BConstFilterer{contract: contract}}, nil
}

// BConst is an auto generated Go binding around an Ethereum contract.
type BConst struct {
	BConstCaller     // Read-only binding to the contract
	BConstTransactor // Write-only binding to the contract
	BConstFilterer   // Log filterer for contract events
}

// BConstCaller is an auto generated read-only Go binding around an Ethereum contract.
type BConstCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BConstTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BConstTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BConstFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BConstFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BConstSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BConstSession struct {
	Contract     *BConst           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BConstCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BConstCallerSession struct {
	Contract *BConstCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BConstTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BConstTransactorSession struct {
	Contract     *BConstTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BConstRaw is an auto generated low-level Go binding around an Ethereum contract.
type BConstRaw struct {
	Contract *BConst // Generic contract binding to access the raw methods on
}

// BConstCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BConstCallerRaw struct {
	Contract *BConstCaller // Generic read-only contract binding to access the raw methods on
}

// BConstTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BConstTransactorRaw struct {
	Contract *BConstTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBConst creates a new instance of BConst, bound to a specific deployed contract.
func NewBConst(address common.Address, backend bind.ContractBackend) (*BConst, error) {
	contract, err := bindBConst(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BConst{BConstCaller: BConstCaller{contract: contract}, BConstTransactor: BConstTransactor{contract: contract}, BConstFilterer: BConstFilterer{contract: contract}}, nil
}

// NewBConstCaller creates a new read-only instance of BConst, bound to a specific deployed contract.
func NewBConstCaller(address common.Address, caller bind.ContractCaller) (*BConstCaller, error) {
	contract, err := bindBConst(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BConstCaller{contract: contract}, nil
}

// NewBConstTransactor creates a new write-only instance of BConst, bound to a specific deployed contract.
func NewBConstTransactor(address common.Address, transactor bind.ContractTransactor) (*BConstTransactor, error) {
	contract, err := bindBConst(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BConstTransactor{contract: contract}, nil
}

// NewBConstFilterer creates a new log filterer instance of BConst, bound to a specific deployed contract.
func NewBConstFilterer(address common.Address, filterer bind.ContractFilterer) (*BConstFilterer, error) {
	contract, err := bindBConst(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BConstFilterer{contract: contract}, nil
}

// bindBConst binds a generic wrapper to an already deployed contract.
func bindBConst(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BConstABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BConst *BConstRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BConst.Contract.BConstCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BConst *BConstRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BConst.Contract.BConstTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BConst *BConstRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BConst.Contract.BConstTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BConst *BConstCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BConst.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BConst *BConstTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BConst.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BConst *BConstTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BConst.Contract.contract.Transact(opts, method, params...)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BConst *BConstCaller) BONE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "BONE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BConst *BConstSession) BONE() (*big.Int, error) {
	return _BConst.Contract.BONE(&_BConst.CallOpts)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BConst *BConstCallerSession) BONE() (*big.Int, error) {
	return _BConst.Contract.BONE(&_BConst.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BConst *BConstCaller) BPOWPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "BPOW_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BConst *BConstSession) BPOWPRECISION() (*big.Int, error) {
	return _BConst.Contract.BPOWPRECISION(&_BConst.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BConst *BConstCallerSession) BPOWPRECISION() (*big.Int, error) {
	return _BConst.Contract.BPOWPRECISION(&_BConst.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BConst *BConstCaller) EXITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "EXIT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BConst *BConstSession) EXITFEE() (*big.Int, error) {
	return _BConst.Contract.EXITFEE(&_BConst.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BConst *BConstCallerSession) EXITFEE() (*big.Int, error) {
	return _BConst.Contract.EXITFEE(&_BConst.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BConst *BConstCaller) INITPOOLSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "INIT_POOL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BConst *BConstSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _BConst.Contract.INITPOOLSUPPLY(&_BConst.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BConst *BConstCallerSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _BConst.Contract.INITPOOLSUPPLY(&_BConst.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BConst *BConstCaller) MAXBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MAX_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BConst *BConstSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _BConst.Contract.MAXBOUNDTOKENS(&_BConst.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BConst *BConstCallerSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _BConst.Contract.MAXBOUNDTOKENS(&_BConst.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BConst *BConstCaller) MAXBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MAX_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BConst *BConstSession) MAXBPOWBASE() (*big.Int, error) {
	return _BConst.Contract.MAXBPOWBASE(&_BConst.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BConst *BConstCallerSession) MAXBPOWBASE() (*big.Int, error) {
	return _BConst.Contract.MAXBPOWBASE(&_BConst.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BConst *BConstCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BConst *BConstSession) MAXFEE() (*big.Int, error) {
	return _BConst.Contract.MAXFEE(&_BConst.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BConst *BConstCallerSession) MAXFEE() (*big.Int, error) {
	return _BConst.Contract.MAXFEE(&_BConst.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BConst *BConstCaller) MAXINRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MAX_IN_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BConst *BConstSession) MAXINRATIO() (*big.Int, error) {
	return _BConst.Contract.MAXINRATIO(&_BConst.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BConst *BConstCallerSession) MAXINRATIO() (*big.Int, error) {
	return _BConst.Contract.MAXINRATIO(&_BConst.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BConst *BConstCaller) MAXOUTRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MAX_OUT_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BConst *BConstSession) MAXOUTRATIO() (*big.Int, error) {
	return _BConst.Contract.MAXOUTRATIO(&_BConst.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BConst *BConstCallerSession) MAXOUTRATIO() (*big.Int, error) {
	return _BConst.Contract.MAXOUTRATIO(&_BConst.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BConst *BConstCaller) MAXTOTALWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MAX_TOTAL_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BConst *BConstSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _BConst.Contract.MAXTOTALWEIGHT(&_BConst.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BConst *BConstCallerSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _BConst.Contract.MAXTOTALWEIGHT(&_BConst.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BConst *BConstCaller) MAXWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MAX_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BConst *BConstSession) MAXWEIGHT() (*big.Int, error) {
	return _BConst.Contract.MAXWEIGHT(&_BConst.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BConst *BConstCallerSession) MAXWEIGHT() (*big.Int, error) {
	return _BConst.Contract.MAXWEIGHT(&_BConst.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BConst *BConstCaller) MINBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MIN_BALANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BConst *BConstSession) MINBALANCE() (*big.Int, error) {
	return _BConst.Contract.MINBALANCE(&_BConst.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BConst *BConstCallerSession) MINBALANCE() (*big.Int, error) {
	return _BConst.Contract.MINBALANCE(&_BConst.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BConst *BConstCaller) MINBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MIN_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BConst *BConstSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _BConst.Contract.MINBOUNDTOKENS(&_BConst.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BConst *BConstCallerSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _BConst.Contract.MINBOUNDTOKENS(&_BConst.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BConst *BConstCaller) MINBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MIN_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BConst *BConstSession) MINBPOWBASE() (*big.Int, error) {
	return _BConst.Contract.MINBPOWBASE(&_BConst.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BConst *BConstCallerSession) MINBPOWBASE() (*big.Int, error) {
	return _BConst.Contract.MINBPOWBASE(&_BConst.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BConst *BConstCaller) MINFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MIN_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BConst *BConstSession) MINFEE() (*big.Int, error) {
	return _BConst.Contract.MINFEE(&_BConst.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BConst *BConstCallerSession) MINFEE() (*big.Int, error) {
	return _BConst.Contract.MINFEE(&_BConst.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BConst *BConstCaller) MINWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MIN_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BConst *BConstSession) MINWEIGHT() (*big.Int, error) {
	return _BConst.Contract.MINWEIGHT(&_BConst.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BConst *BConstCallerSession) MINWEIGHT() (*big.Int, error) {
	return _BConst.Contract.MINWEIGHT(&_BConst.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BConst *BConstCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "getColor")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BConst *BConstSession) GetColor() ([32]byte, error) {
	return _BConst.Contract.GetColor(&_BConst.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BConst *BConstCallerSession) GetColor() ([32]byte, error) {
	return _BConst.Contract.GetColor(&_BConst.CallOpts)
}

// BFactoryABI is the input ABI used to generate the binding from.
const BFactoryABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"blabs\",\"type\":\"address\"}],\"name\":\"LOG_BLABS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"LOG_NEW_POOL\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractBPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"collect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBLabs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"isBPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"newBPool\",\"outputs\":[{\"internalType\":\"contractBPool\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"setBLabs\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BFactoryFuncSigs maps the 4-byte function signature to its string representation.
var BFactoryFuncSigs = map[string]string{
	"06ec16f8": "collect(address)",
	"36ffb167": "getBLabs()",
	"9a86139b": "getColor()",
	"c2bb6dc2": "isBPool(address)",
	"d556c5dc": "newBPool()",
	"c6ce34fb": "setBLabs(address)",
}

// BFactoryBin is the compiled bytecode used for deploying new contracts.
var BFactoryBin = "0x608060405234801561001057600080fd5b50600180546001600160a01b03191633179055615bdf806100326000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806306ec16f81461006757806336ffb1671461008f5780639a86139b146100b3578063c2bb6dc2146100cd578063c6ce34fb14610107578063d556c5dc1461012d575b600080fd5b61008d6004803603602081101561007d57600080fd5b50356001600160a01b0316610135565b005b6100976102cd565b604080516001600160a01b039092168252519081900360200190f35b6100bb6102dc565b60408051918252519081900360200190f35b6100f3600480360360208110156100e357600080fd5b50356001600160a01b03166102e9565b604080519115158252519081900360200190f35b61008d6004803603602081101561011d57600080fd5b50356001600160a01b0316610307565b6100976103ae565b6001546001600160a01b03163314610184576040805162461bcd60e51b815260206004820152600d60248201526c4552525f4e4f545f424c41425360981b604482015290519081900360640190fd5b604080516370a0823160e01b815230600482015290516000916001600160a01b038416916370a0823191602480820192602092909190829003018186803b1580156101ce57600080fd5b505afa1580156101e2573d6000803e3d6000fd5b505050506040513d60208110156101f857600080fd5b50516001546040805163a9059cbb60e01b81526001600160a01b0392831660048201526024810184905290519293506000929185169163a9059cbb9160448082019260209290919082900301818787803b15801561025557600080fd5b505af1158015610269573d6000803e3d6000fd5b505050506040513d602081101561027f57600080fd5b50519050806102c8576040805162461bcd60e51b815260206004820152601060248201526f11549497d15490cc8c17d1905253115160821b604482015290519081900360640190fd5b505050565b6001546001600160a01b031690565b6542524f4e5a4560d01b90565b6001600160a01b031660009081526020819052604090205460ff1690565b6001546001600160a01b03163314610356576040805162461bcd60e51b815260206004820152600d60248201526c4552525f4e4f545f424c41425360981b604482015290519081900360640190fd5b6040516001600160a01b0382169033907ff586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d90600090a3600180546001600160a01b0319166001600160a01b0392909216919091179055565b6000806040516103bd9061048e565b604051809103906000f0801580156103d9573d6000803e3d6000fd5b506001600160a01b038116600081815260208190526040808220805460ff1916600117905551929350909133917f8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f94591a3604080516392eefe9b60e01b815233600482015290516001600160a01b038316916392eefe9b91602480830192600092919082900301818387803b15801561047057600080fd5b505af1158015610484573d6000803e3d6000fd5b5092935050505090565b61570f8061049c8339019056fe60c0604052601360808190527f42616c616e63657220506f6f6c20546f6b656e0000000000000000000000000060a0908152620000409160039190620000f4565b506040805180820190915260038082527f425054000000000000000000000000000000000000000000000000000000000060209092019182526200008791600491620000f4565b506005805460ff19166012179055348015620000a257600080fd5b50600680546005805462010000600160b01b031916336201000081029190911790915564e8d4a510006007556001600160a01b03199091161760ff60a01b191690556008805460ff1916905562000199565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200013757805160ff191683800117855562000167565b8280016001018555821562000167579182015b82811115620001675782518255916020019190600101906200014a565b506200017592915062000179565b5090565b6200019691905b8082111562000175576000815560010162000180565b90565b61556680620001a96000396000f3fe608060405234801561001057600080fd5b50600436106103db5760003560e01c80638d4e40831161020a578063bc694ea211610125578063d73dd623116100b8578063ec09302111610087578063ec09302114610c6e578063f1b8a9b714610c76578063f8b2cb4f14610c9c578063f8d6aed414610cc2578063fde924f714610cfd576103db565b8063d73dd62314610be2578063dd62ed3e14610c0e578063e4a28a52146104e1578063e4e1e53814610c3c576103db565b8063cc77828d116100f4578063cc77828d14610ba4578063cd2ed8fb14610bac578063cf5e7bd314610bb4578063d4cadf6814610bda576103db565b8063bc694ea214610b3c578063be3bbd2e14610b44578063c36596a614610555578063c6580d1214610b9c576103db565b8063a221ee491161019d578063b7b800a41161016c578063b7b800a414610ae9578063ba019dab14610af1578063ba9530a614610af9578063bc063e1a14610b34576103db565b8063a221ee4914610a09578063a9059cbb14610a3e578063b02f0b7314610a6a578063b0e0d13614610ae1576103db565b8063948d8ce6116101d9578063948d8ce6146109cb57806395d89b41146109f1578063992e2a92146109f95780639a86139b14610a01576103db565b80638d4e40831461098d57806392eefe9b14610995578063936c3477146109bb5780639381cd2b146109c3576103db565b806349b59552116102fa57806376c7a3c71161028d5780638656b6531161025c5780638656b653146108e9578063867378c514610924578063892980121461092c5780638c28cbe814610967576103db565b806376c7a3c71461080d5780637c5e9ea4146108155780638201aa3f1461086e57806382f652ad146108ae576103db565b80635db34277116102c95780635db342771461075757806366188463146107895780636d06dfa0146107b557806370a08231146107e7576103db565b806349b595521461067e5780634bb278f31461069d5780634f69c0d4146106a55780635c1bbaf71461071c576103db565b8063218b538211610372578063313ce56711610341578063313ce567146105dd57806334e19907146105fb5780633fdddaa21461061a57806346ab38f11461064c576103db565b8063218b53821461055557806323b872dd1461055d5780632f37b624146105935780633018205f146105b9576103db565b80631446a7ff116103ae5780631446a7ff146104e957806315e84af91461051757806318160ddd14610545578063189d00ca1461054d576103db565b806302c96748146103e057806306fdde0314610424578063095ea7b3146104a157806309a3bbe4146104e1575b600080fd5b610412600480360360608110156103f657600080fd5b506001600160a01b038135169060208101359060400135610d05565b60408051918252519081900360200190f35b61042c611065565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561046657818101518382015260200161044e565b50505050905090810190601f1680156104935780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6104cd600480360360408110156104b757600080fd5b506001600160a01b0381351690602001356110fb565b604080519115158252519081900360200190f35b610412611150565b610412600480360360408110156104ff57600080fd5b506001600160a01b038135811691602001351661115d565b6104126004803603604081101561052d57600080fd5b506001600160a01b03813581169160200135166112b2565b6104126113fe565b610412611404565b610412611418565b6104cd6004803603606081101561057357600080fd5b506001600160a01b03813581169160208101359091169060400135611424565b6104cd600480360360208110156105a957600080fd5b50356001600160a01b031661157e565b6105c161159c565b604080516001600160a01b039092168252519081900360200190f35b6105e56115fa565b6040805160ff9092168252519081900360200190f35b6106186004803603602081101561061157600080fd5b5035611603565b005b6106186004803603606081101561063057600080fd5b506001600160a01b038135169060208101359060400135611800565b6104126004803603606081101561066257600080fd5b506001600160a01b038135169060208101359060400135611c0d565b6106186004803603602081101561069457600080fd5b50351515611f0c565b61061861208f565b610618600480360360408110156106bb57600080fd5b813591908101906040810160208201356401000000008111156106dd57600080fd5b8201836020820111156106ef57600080fd5b8035906020019184602083028401116401000000008311171561071157600080fd5b509092509050612286565b610412600480360360c081101561073257600080fd5b5080359060208101359060408101359060608101359060808101359060a0013561257f565b6104126004803603606081101561076d57600080fd5b506001600160a01b038135169060208101359060400135612637565b6104cd6004803603604081101561079f57600080fd5b506001600160a01b03813516906020013561291a565b610412600480360360608110156107cb57600080fd5b506001600160a01b0381351690602081013590604001356129f2565b610412600480360360208110156107fd57600080fd5b50356001600160a01b0316612d03565b610412612d1e565b610855600480360360a081101561082b57600080fd5b506001600160a01b0381358116916020810135916040820135169060608101359060800135612d30565b6040805192835260208301919091528051918290030190f35b610855600480360360a081101561088457600080fd5b506001600160a01b03813581169160208101359160408201351690606081013590608001356131f3565b610412600480360360c08110156108c457600080fd5b5080359060208101359060408101359060608101359060808101359060a0013561369d565b610412600480360360c08110156108ff57600080fd5b5080359060208101359060408101359060608101359060808101359060a0013561375c565b6104126137fd565b610412600480360360c081101561094257600080fd5b5080359060208101359060408101359060608101359060808101359060a00135613811565b6106186004803603602081101561097d57600080fd5b50356001600160a01b03166138c1565b6104cd613a75565b610618600480360360208110156109ab57600080fd5b50356001600160a01b0316613a7e565b610412613bbc565b610412613c11565b610412600480360360208110156109e157600080fd5b50356001600160a01b0316613c1e565b61042c613ce8565b610412613d49565b610412613d55565b610412600480360360a0811015610a1f57600080fd5b5080359060208101359060408101359060608101359060800135613d62565b6104cd60048036036040811015610a5457600080fd5b506001600160a01b038135169060200135613dc7565b61061860048036036040811015610a8057600080fd5b81359190810190604081016020820135640100000000811115610aa257600080fd5b820183602082011115610ab457600080fd5b80359060200191846020830284011164010000000083111715610ad657600080fd5b509092509050613ddd565b610412614124565b610412614129565b61041261412e565b610412600480360360c0811015610b0f57600080fd5b5080359060208101359060408101359060608101359060808101359060a00135614133565b6104126141b4565b6104126141c4565b610b4c6141d0565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015610b88578181015183820152602001610b70565b505050509050019250505060405180910390f35b6104126142c8565b610b4c6142cd565b61041261431b565b61061860048036036020811015610bca57600080fd5b50356001600160a01b0316614321565b6104126146a3565b6104cd60048036036040811015610bf857600080fd5b506001600160a01b0381351690602001356146f8565b61041260048036036040811015610c2457600080fd5b506001600160a01b0381358116916020013516614779565b61061860048036036060811015610c5257600080fd5b506001600160a01b0381351690602081013590604001356147a4565b6104126149fb565b61041260048036036020811015610c8c57600080fd5b50356001600160a01b0316614a0b565b61041260048036036020811015610cb257600080fd5b50356001600160a01b0316614ae7565b610412600480360360c0811015610cd857600080fd5b5080359060208101359060408101359060608101359060808101359060a00135614bb1565b6104cd614c34565b6000336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff1615610db3576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff00191661010017905560085460ff16610e0d576040805162461bcd60e51b815260206004820152601160248201527011549497d393d517d19253905312569151607a1b604482015290519081900360640190fd5b6001600160a01b0384166000908152600a602052604090205460ff16610e6a576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b0384166000908152600a60205260409020600390810154610e9f91670de0b6b3a76400005b04600101614c44565b831115610ee7576040805162461bcd60e51b81526020600482015260116024820152704552525f4d41585f4f55545f524154494f60781b604482015290519081900360640190fd5b6001600160a01b0384166000908152600a6020526040902060038101546002808301549054600b54600754610f219493929190899061369d565b915081610f67576040805162461bcd60e51b815260206004820152600f60248201526e08aa4a4be9a82a890be82a0a0a49eb608b1b604482015290519081900360640190fd5b82821115610fab576040805162461bcd60e51b815260206004820152600c60248201526b22a9292fa624a6a4aa2fa4a760a11b604482015290519081900360640190fd5b610fb9816003015485614d0d565b60038201556000610fca8382614c44565b6040805187815290519192506001600160a01b0388169133917fe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed919081900360200190a36110183384614d6f565b61102a6110258483614d0d565b614d7d565b600554611046906201000090046001600160a01b031682614d89565b611051863387614d93565b50506005805461ff00191690559392505050565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156110f15780601f106110c6576101008083540402835291602001916110f1565b820191906000526020600020905b8154815290600101906020018083116110d457829003601f168201915b5050505050905090565b3360008181526001602090815260408083206001600160a01b03871680855290835281842086905581518681529151939490939092600080516020615512833981519152928290030190a35060015b92915050565b6802b5e3af16b188000081565b600554600090610100900460ff16156111ab576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6001600160a01b0383166000908152600a602052604090205460ff16611208576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b0382166000908152600a602052604090205460ff16611265576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b038084166000908152600a602052604080822092851682528120600380840154600280860154928401549084015493946112a99492939290613d62565b95945050505050565b600554600090610100900460ff1615611300576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6001600160a01b0383166000908152600a602052604090205460ff1661135d576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b0382166000908152600a602052604090205460ff166113ba576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b038084166000908152600a60205260408082209285168252902060038083015460028085015492840154908401546007546112a994929190613d62565b60025490565b6402540be400670de0b6b3a76400005b0481565b670de0b6b3a764000081565b6000336001600160a01b038516148061146057506001600160a01b03841660009081526001602090815260408083203384529091529020548211155b6114a9576040805162461bcd60e51b815260206004820152601560248201527422a9292fa12a27a5a2a72fa120a22fa1a0a62622a960591b604482015290519081900360640190fd5b6114b4848484614e5e565b336001600160a01b038516148015906114f257506001600160a01b038416600090815260016020908152604080832033845290915290205460001914155b15611574576001600160a01b03841660009081526001602090815260408083203384529091529020546115259083614d0d565b6001600160a01b03858116600090815260016020908152604080832033808552908352928190208590558051948552519287169391926000805160206155128339815191529281900390910190a35b5060019392505050565b6001600160a01b03166000908152600a602052604090205460ff1690565b600554600090610100900460ff16156115ea576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b506006546001600160a01b031690565b60055460ff1690565b336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff16156116af576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff00191661010017905560085460ff1615611709576040805162461bcd60e51b815260206004820152601060248201526f11549497d254d7d1925390531256915160821b604482015290519081900360640190fd5b6006546001600160a01b0316331461175d576040805162461bcd60e51b815260206004820152601260248201527122a9292fa727aa2fa1a7a72a2927a62622a960711b604482015290519081900360640190fd5b64e8d4a510008110156117a5576040805162461bcd60e51b815260206004820152600b60248201526a4552525f4d494e5f46454560a81b604482015290519081900360640190fd5b67016345785d8a00008111156117f0576040805162461bcd60e51b815260206004820152600b60248201526a4552525f4d41585f46454560a81b604482015290519081900360640190fd5b6007556005805461ff0019169055565b336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff16156118ac576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff001916610100179055600654336001600160a01b0390911614611911576040805162461bcd60e51b815260206004820152601260248201527122a9292fa727aa2fa1a7a72a2927a62622a960711b604482015290519081900360640190fd5b6001600160a01b0383166000908152600a602052604090205460ff1661196e576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b60085460ff16156119b9576040805162461bcd60e51b815260206004820152601060248201526f11549497d254d7d1925390531256915160821b604482015290519081900360640190fd5b670de0b6b3a7640000811015611a07576040805162461bcd60e51b815260206004820152600e60248201526d11549497d3525397d5d15251d21560921b604482015290519081900360640190fd5b6802b5e3af16b1880000811115611a56576040805162461bcd60e51b815260206004820152600e60248201526d11549497d3505617d5d15251d21560921b604482015290519081900360640190fd5b620f4240821015611aa0576040805162461bcd60e51b815260206004820152600f60248201526e4552525f4d494e5f42414c414e434560881b604482015290519081900360640190fd5b6001600160a01b0383166000908152600a602052604090206002015480821115611b3757611ad9600b54611ad48484614d0d565b614f6e565b600b8190556802b5e3af16b18800001015611b32576040805162461bcd60e51b815260206004820152601460248201527311549497d3505617d513d5105317d5d15251d21560621b604482015290519081900360640190fd5b611b58565b80821015611b5857611b54600b54611b4f8385614d0d565b614d0d565b600b555b6001600160a01b0384166000908152600a602052604090206002810183905560030180549084905580841115611ba157611b9c8533611b978785614d0d565b614fbb565b611bfb565b80841015611bfb576000611bb58286614d0d565b90506000611bc4826000614c44565b9050611bda8733611bd58585614d0d565b614d93565b600554611bf89088906201000090046001600160a01b031683614d93565b50505b50506005805461ff0019169055505050565b6000336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff1615611cbb576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff00191661010017905560085460ff16611d15576040805162461bcd60e51b815260206004820152601160248201527011549497d393d517d19253905312569151607a1b604482015290519081900360640190fd5b6001600160a01b0384166000908152600a602052604090205460ff16611d72576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b0384166000908152600a6020526040902060038101546002808301549054600b54600754611dac94939291908990613811565b915082821015611df3576040805162461bcd60e51b815260206004820152600d60248201526c11549497d31253525517d3d555609a1b604482015290519081900360640190fd5b6001600160a01b0385166000908152600a60205260409020600390810154611e2391670de0b6b3a7640000610e96565b821115611e6b576040805162461bcd60e51b81526020600482015260116024820152704552525f4d41585f4f55545f524154494f60781b604482015290519081900360640190fd5b611e79816003015483614d0d565b60038201556000611e8a8582614c44565b6040805185815290519192506001600160a01b0388169133917fe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed919081900360200190a3611ed83386614d6f565b611ee56110258683614d0d565b600554611f01906201000090046001600160a01b031682614d89565b611051863385614d93565b336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff1615611fb8576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff00191661010017905560085460ff1615612012576040805162461bcd60e51b815260206004820152601060248201526f11549497d254d7d1925390531256915160821b604482015290519081900360640190fd5b6006546001600160a01b03163314612066576040805162461bcd60e51b815260206004820152601260248201527122a9292fa727aa2fa1a7a72a2927a62622a960711b604482015290519081900360640190fd5b60068054911515600160a01b0260ff60a01b199092169190911790556005805461ff0019169055565b336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff161561213b576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff001916610100179055600654336001600160a01b03909116146121a0576040805162461bcd60e51b815260206004820152601260248201527122a9292fa727aa2fa1a7a72a2927a62622a960711b604482015290519081900360640190fd5b60085460ff16156121eb576040805162461bcd60e51b815260206004820152601060248201526f11549497d254d7d1925390531256915160821b604482015290519081900360640190fd5b60095460021115612234576040805162461bcd60e51b815260206004820152600e60248201526d4552525f4d494e5f544f4b454e5360901b604482015290519081900360640190fd5b6008805460ff191660011790556006805460ff60a01b1916600160a01b17905561226668056bc75e2d63100000615014565b6122793368056bc75e2d63100000614d89565b6005805461ff0019169055565b336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff1615612332576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff00191661010017905560085460ff1661238c576040805162461bcd60e51b815260206004820152601160248201527011549497d393d517d19253905312569151607a1b604482015290519081900360640190fd5b60006123966113fe565b905060006123a4858361501d565b9050806123ea576040805162461bcd60e51b815260206004820152600f60248201526e08aa4a4be9a82a890be82a0a0a49eb608b1b604482015290519081900360640190fd5b60005b60095481101561256b5760006009828154811061240657fe5b60009182526020808320909101546001600160a01b0316808352600a90915260408220600301549092509061243b8583614c44565b905080612481576040805162461bcd60e51b815260206004820152600f60248201526e08aa4a4be9a82a890be82a0a0a49eb608b1b604482015290519081900360640190fd5b87878581811061248d57fe5b905060200201358111156124d7576040805162461bcd60e51b815260206004820152600c60248201526b22a9292fa624a6a4aa2fa4a760a11b604482015290519081900360640190fd5b6001600160a01b0383166000908152600a60205260409020600301546124fd9082614f6e565b6001600160a01b0384166000818152600a60209081526040918290206003019390935580518481529051919233927f63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a9281900390910190a3612560833383614fbb565b5050506001016123ed565b5061257585615014565b611bfb3386614d89565b60008061258c878661501d565b9050600061259a8786614f6e565b905060006125a8828961501d565b905060006125be670de0b6b3a76400008561501d565b905060006125cc8383615125565b905060006125da828e614c44565b905060006125e8828f614d0d565b90506000612607612601670de0b6b3a76400008a614d0d565b8b614c44565b90506126248261261f670de0b6b3a764000084614d0d565b61501d565b9f9e505050505050505050505050505050565b6000336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff16156126e5576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff00191661010017905560085460ff1661273f576040805162461bcd60e51b815260206004820152601160248201527011549497d393d517d19253905312569151607a1b604482015290519081900360640190fd5b6001600160a01b0384166000908152600a602052604090205460ff1661279c576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b0384166000908152600a60205260409020600301546127ce906002670de0b6b3a76400005b04614c44565b831115612815576040805162461bcd60e51b815260206004820152601060248201526f4552525f4d41585f494e5f524154494f60801b604482015290519081900360640190fd5b6001600160a01b0384166000908152600a6020526040902060038101546002808301549054600b5460075461284f9493929190899061375c565b915082821015612896576040805162461bcd60e51b815260206004820152600d60248201526c11549497d31253525517d3d555609a1b604482015290519081900360640190fd5b6128a4816003015485614f6e565b60038201556040805185815290516001600160a01b0387169133917f63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a9181900360200190a36128f282615014565b6128fc3383614d89565b612907853386614fbb565b506005805461ff00191690559392505050565b3360009081526001602090815260408083206001600160a01b03861684529091528120548083111561296f573360009081526001602090815260408083206001600160a01b038816845290915281205561299e565b6129798184614d0d565b3360009081526001602090815260408083206001600160a01b03891684529091529020555b3360008181526001602090815260408083206001600160a01b038916808552908352928190205481519081529051929392600080516020615512833981519152929181900390910190a35060019392505050565b6000336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff1615612aa0576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff00191661010017905560085460ff16612afa576040805162461bcd60e51b815260206004820152601160248201527011549497d393d517d19253905312569151607a1b604482015290519081900360640190fd5b6001600160a01b0384166000908152600a602052604090205460ff16612b57576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b0384166000908152600a6020526040902060038101546002808301549054600b54600754612b919493929190899061257f565b915081612bd7576040805162461bcd60e51b815260206004820152600f60248201526e08aa4a4be9a82a890be82a0a0a49eb608b1b604482015290519081900360640190fd5b82821115612c1b576040805162461bcd60e51b815260206004820152600c60248201526b22a9292fa624a6a4aa2fa4a760a11b604482015290519081900360640190fd5b6001600160a01b0385166000908152600a6020526040902060030154612c4b906002670de0b6b3a76400006127c8565b821115612c92576040805162461bcd60e51b815260206004820152601060248201526f4552525f4d41585f494e5f524154494f60801b604482015290519081900360640190fd5b612ca0816003015483614f6e565b60038201556040805183815290516001600160a01b0387169133917f63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a9181900360200190a3612cee84615014565b612cf83385614d89565b612907853384614fbb565b6001600160a01b031660009081526020819052604090205490565b620f4240670de0b6b3a7640000611414565b60408051602080825236908201819052600092839233926001600160e01b03198535169285929081908101848480828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff1615612dcd576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff0019166101001790556001600160a01b0387166000908152600a602052604090205460ff16612e39576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b0385166000908152600a602052604090205460ff16612e96576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b600654600160a01b900460ff16612eea576040805162461bcd60e51b81526020600482015260136024820152724552525f535741505f4e4f545f5055424c494360681b604482015290519081900360640190fd5b6001600160a01b038088166000908152600a602052604080822092881682529020600380820154612f2391670de0b6b3a7640000610e96565b861115612f6b576040805162461bcd60e51b81526020600482015260116024820152704552525f4d41585f4f55545f524154494f60781b604482015290519081900360640190fd5b6000612f8c8360030154846002015484600301548560020154600754613d62565b905085811115612fd9576040805162461bcd60e51b81526020600482015260136024820152724552525f4241445f4c494d49545f505249434560681b604482015290519081900360640190fd5b612ff983600301548460020154846003015485600201548b600754614bb1565b94508885111561303f576040805162461bcd60e51b815260206004820152600c60248201526b22a9292fa624a6a4aa2fa4a760a11b604482015290519081900360640190fd5b61304d836003015486614f6e565b8360030181905550613063826003015488614d0d565b600380840182905584015460028086015490850154600754613086949190613d62565b9350808410156130cf576040805162461bcd60e51b815260206004820152600f60248201526e08aa4a4be9a82a890be82a0a0a49eb608b1b604482015290519081900360640190fd5b85841115613116576040805162461bcd60e51b815260206004820152600f60248201526e4552525f4c494d49545f505249434560881b604482015290519081900360640190fd5b613120858861501d565b811115613166576040805162461bcd60e51b815260206004820152600f60248201526e08aa4a4be9a82a890be82a0a0a49eb608b1b604482015290519081900360640190fd5b876001600160a01b03168a6001600160a01b0316336001600160a01b03167f908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378888b604051808381526020018281526020019250505060405180910390a46131ce8a3387614fbb565b6131d9883389614d93565b5050506005805461ff001916905590969095509350505050565b60408051602080825236908201819052600092839233926001600160e01b03198535169285929081908101848480828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff1615613290576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff0019166101001790556001600160a01b0387166000908152600a602052604090205460ff166132fc576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b0385166000908152600a602052604090205460ff16613359576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b600654600160a01b900460ff166133ad576040805162461bcd60e51b81526020600482015260136024820152724552525f535741505f4e4f545f5055424c494360681b604482015290519081900360640190fd5b6001600160a01b038088166000908152600a60205260408082209288168252902060038201546133e7906002670de0b6b3a76400006127c8565b88111561342e576040805162461bcd60e51b815260206004820152601060248201526f4552525f4d41585f494e5f524154494f60801b604482015290519081900360640190fd5b600061344f8360030154846002015484600301548560020154600754613d62565b90508581111561349c576040805162461bcd60e51b81526020600482015260136024820152724552525f4241445f4c494d49545f505249434560681b604482015290519081900360640190fd5b6134bc83600301548460020154846003015485600201548d600754614133565b945086851015613503576040805162461bcd60e51b815260206004820152600d60248201526c11549497d31253525517d3d555609a1b604482015290519081900360640190fd5b61351183600301548a614f6e565b8360030181905550613527826003015486614d0d565b60038084018290558401546002808601549085015460075461354a949190613d62565b935080841015613593576040805162461bcd60e51b815260206004820152600f60248201526e08aa4a4be9a82a890be82a0a0a49eb608b1b604482015290519081900360640190fd5b858411156135da576040805162461bcd60e51b815260206004820152600f60248201526e4552525f4c494d49545f505249434560881b604482015290519081900360640190fd5b6135e4898661501d565b81111561362a576040805162461bcd60e51b815260206004820152600f60248201526e08aa4a4be9a82a890be82a0a0a49eb608b1b604482015290519081900360640190fd5b876001600160a01b03168a6001600160a01b0316336001600160a01b03167f908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d433788c89604051808381526020018281526020019250505060405180910390a46136928a338b614fbb565b6131d9883387614d93565b6000806136aa878661501d565b905060006136c0670de0b6b3a764000083614d0d565b905060006136ce8286614c44565b905060006136e88761261f670de0b6b3a764000085614d0d565b905060006136f68c83614d0d565b90506000613704828e61501d565b905060006137128288615125565b90506000613720828e614c44565b9050600061372e8e83614d0d565b90506137478161261f670de0b6b3a76400006000614d0d565b99505050505050505050509695505050505050565b600080613769878661501d565b90506000613788613782670de0b6b3a764000084614d0d565b85614c44565b905060006137a7866137a2670de0b6b3a764000085614d0d565b614c44565b905060006137b58b83614f6e565b905060006137c3828d61501d565b905060006137d18287615125565b905060006137df828d614c44565b90506137eb818d614d0d565b9e9d5050505050505050505050505050565b64e8d4a51000670de0b6b3a7640000611414565b60008061381e878661501d565b90506000613839856137a2670de0b6b3a76400006000614d0d565b905060006138478883614d0d565b90506000613855828a61501d565b905060006138748261386f670de0b6b3a76400008861501d565b615125565b90506000613882828e614c44565b905060006138908e83614d0d565b905060006138a9612601670de0b6b3a76400008a614d0d565b9050612624826137a2670de0b6b3a764000084614d0d565b336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff161561396d576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff0019166101001790556001600160a01b0381166000908152600a602052604090205460ff166139d9576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b604080516370a0823160e01b815230600482015290516001600160a01b038316916370a08231916024808301926020929190829003018186803b158015613a1f57600080fd5b505afa158015613a33573d6000803e3d6000fd5b505050506040513d6020811015613a4957600080fd5b50516001600160a01b039091166000908152600a60205260409020600301556005805461ff0019169055565b60085460ff1690565b336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff1615613b2a576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff001916610100179055600654336001600160a01b0390911614613b8f576040805162461bcd60e51b815260206004820152601260248201527122a9292fa727aa2fa1a7a72a2927a62622a960711b604482015290519081900360640190fd5b600680546001600160a01b0319166001600160a01b03929092169190911790556005805461ff0019169055565b600554600090610100900460ff1615613c0a576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b50600b5490565b68056bc75e2d6310000081565b600554600090610100900460ff1615613c6c576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6001600160a01b0382166000908152600a602052604090205460ff16613cc9576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b506001600160a01b03166000908152600a602052604090206002015490565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156110f15780601f106110c6576101008083540402835291602001916110f1565b6704a03ce68d21555681565b6542524f4e5a4560d01b90565b600080613d6f878761501d565b90506000613d7d868661501d565b90506000613d8b838361501d565b90506000613dad670de0b6b3a764000061261f670de0b6b3a764000089614d0d565b9050613db98282614c44565b9a9950505050505050505050565b6000613dd4338484614e5e565b50600192915050565b336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff1615613e89576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff00191661010017905560085460ff16613ee3576040805162461bcd60e51b815260206004820152601160248201527011549497d393d517d19253905312569151607a1b604482015290519081900360640190fd5b6000613eed6113fe565b90506000613efc856000614c44565b90506000613f0a8683614d0d565b90506000613f18828561501d565b905080613f5e576040805162461bcd60e51b815260206004820152600f60248201526e08aa4a4be9a82a890be82a0a0a49eb608b1b604482015290519081900360640190fd5b613f683388614d6f565b600554613f84906201000090046001600160a01b031684614d89565b613f8d82614d7d565b60005b60095481101561410f57600060098281548110613fa957fe5b60009182526020808320909101546001600160a01b0316808352600a909152604082206003015490925090613fde8583614c44565b905080614024576040805162461bcd60e51b815260206004820152600f60248201526e08aa4a4be9a82a890be82a0a0a49eb608b1b604482015290519081900360640190fd5b89898581811061403057fe5b9050602002013581101561407b576040805162461bcd60e51b815260206004820152600d60248201526c11549497d31253525517d3d555609a1b604482015290519081900360640190fd5b6001600160a01b0383166000908152600a60205260409020600301546140a19082614d0d565b6001600160a01b0384166000818152600a60209081526040918290206003019390935580518481529051919233927fe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed9281900390910190a3614104833383614d93565b505050600101613f90565b50506005805461ff0019169055505050505050565b600881565b600281565b600181565b600080614140878661501d565b90506000614156670de0b6b3a764000085614d0d565b90506141628582614c44565b905060006141748a61261f8c85614f6e565b905060006141828285615125565b90506000614198670de0b6b3a764000083614d0d565b90506141a48a82614c44565b9c9b505050505050505050505050565b600a670de0b6b3a7640000611414565b671bc16d674ec7ffff81565b600554606090610100900460ff161561421e576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b60085460ff16614269576040805162461bcd60e51b815260206004820152601160248201527011549497d393d517d19253905312569151607a1b604482015290519081900360640190fd5b60098054806020026020016040519081016040528092919081815260200182805480156110f157602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116142a1575050505050905090565b600081565b600554606090610100900460ff1615614269576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b60095490565b336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff16156143cd576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff001916610100179055600654336001600160a01b0390911614614432576040805162461bcd60e51b815260206004820152601260248201527122a9292fa727aa2fa1a7a72a2927a62622a960711b604482015290519081900360640190fd5b6001600160a01b0381166000908152600a602052604090205460ff1661448f576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b60085460ff16156144da576040805162461bcd60e51b815260206004820152601060248201526f11549497d254d7d1925390531256915160821b604482015290519081900360640190fd5b6001600160a01b0381166000908152600a6020526040812060030154906145018282614c44565b600b546001600160a01b0385166000908152600a602052604090206002015491925061452c91614d0d565b600b556001600160a01b0383166000908152600a602052604090206001015460098054600019810191908290811061456057fe5b600091825260209091200154600980546001600160a01b03909216918490811061458657fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081600a6000600985815481106145c657fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190206001015560098054806145f957fe5b60008281526020808220600019908401810180546001600160a01b031916905590920190925560408051608081018252838152808301848152818301858152606083018681526001600160a01b038c168752600a909552929094209051815460ff191690151517815592516001840155516002830155516003909101556146858533611bd58787614d0d565b600554611bfb9086906201000090046001600160a01b031685614d93565b600554600090610100900460ff16156146f1576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b5060075490565b3360009081526001602090815260408083206001600160a01b03861684529091528120546147269083614f6e565b3360008181526001602090815260408083206001600160a01b038916808552908352928190208590558051948552519193600080516020615512833981519152929081900390910190a350600192915050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a26006546001600160a01b03163314614859576040805162461bcd60e51b815260206004820152601260248201527122a9292fa727aa2fa1a7a72a2927a62622a960711b604482015290519081900360640190fd5b6001600160a01b0383166000908152600a602052604090205460ff16156148b6576040805162461bcd60e51b815260206004820152600c60248201526b11549497d254d7d093d5539160a21b604482015290519081900360640190fd5b60085460ff1615614901576040805162461bcd60e51b815260206004820152601060248201526f11549497d254d7d1925390531256915160821b604482015290519081900360640190fd5b600954600811614949576040805162461bcd60e51b815260206004820152600e60248201526d4552525f4d41585f544f4b454e5360901b604482015290519081900360640190fd5b6040805160808101825260018082526009805460208085019182526000858701818152606087018281526001600160a01b038c16808452600a9094529782209651875460ff1916901515178755925186860155915160028601559451600390940193909355805491820181559091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0180546001600160a01b03191690911790556149f6838383611800565b505050565b6002670de0b6b3a7640000611414565b600554600090610100900460ff1615614a59576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6001600160a01b0382166000908152600a602052604090205460ff16614ab6576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b0382166000908152600a6020526040902060020154600b54614ae090829061501d565b9392505050565b600554600090610100900460ff1615614b35576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6001600160a01b0382166000908152600a602052604090205460ff16614b92576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b506001600160a01b03166000908152600a602052604090206003015490565b600080614bbe858861501d565b90506000614bcc8786614d0d565b90506000614bda888361501d565b90506000614be88285615125565b9050614bfc81670de0b6b3a7640000614d0d565b9050614c10670de0b6b3a764000087614d0d565b9450614c25614c1f8c83614c44565b8661501d565b9b9a5050505050505050505050565b600654600160a01b900460ff1690565b6000828202831580614c5e575082848281614c5b57fe5b04145b614ca2576040805162461bcd60e51b815260206004820152601060248201526f4552525f4d554c5f4f564552464c4f5760801b604482015290519081900360640190fd5b6706f05b59d3b20000810181811015614cf5576040805162461bcd60e51b815260206004820152601060248201526f4552525f4d554c5f4f564552464c4f5760801b604482015290519081900360640190fd5b6000670de0b6b3a7640000825b049695505050505050565b6000806000614d1c8585615233565b915091508015614d67576040805162461bcd60e51b81526020600482015260116024820152704552525f5355425f554e444552464c4f5760781b604482015290519081900360640190fd5b509392505050565b614d798282615258565b5050565b614d8681615263565b50565b614d798282615333565b6040805163a9059cbb60e01b81526001600160a01b03848116600483015260248201849052915160009286169163a9059cbb91604480830192602092919082900301818787803b158015614de657600080fd5b505af1158015614dfa573d6000803e3d6000fd5b505050506040513d6020811015614e1057600080fd5b5051905080614e58576040805162461bcd60e51b815260206004820152600f60248201526e4552525f45524332305f46414c534560881b604482015290519081900360640190fd5b50505050565b6001600160a01b038316600090815260208190526040902054811115614ec2576040805162461bcd60e51b815260206004820152601460248201527311549497d25394d551919250d251539517d0905360621b604482015290519081900360640190fd5b6001600160a01b038316600090815260208190526040902054614ee59082614d0d565b6001600160a01b038085166000908152602081905260408082209390935590841681522054614f149082614f6e565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600082820183811015614ae0576040805162461bcd60e51b815260206004820152601060248201526f4552525f4144445f4f564552464c4f5760801b604482015290519081900360640190fd5b604080516323b872dd60e01b81526001600160a01b0384811660048301523060248301526044820184905291516000928616916323b872dd91606480830192602092919082900301818787803b158015614de657600080fd5b614d868161533e565b600081615060576040805162461bcd60e51b815260206004820152600c60248201526b4552525f4449565f5a45524f60a01b604482015290519081900360640190fd5b670de0b6b3a764000083028315806150885750670de0b6b3a764000084828161508557fe5b04145b6150cc576040805162461bcd60e51b815260206004820152601060248201526f11549497d1125597d25395115493905360821b604482015290519081900360640190fd5b6002830481018181101561511a576040805162461bcd60e51b815260206004820152601060248201526f11549497d1125597d25395115493905360821b604482015290519081900360640190fd5b6000848281614d0257fe5b60006001831015615175576040805162461bcd60e51b81526020600482015260156024820152744552525f42504f575f424153455f544f4f5f4c4f5760581b604482015290519081900360640190fd5b671bc16d674ec7ffff8311156151cb576040805162461bcd60e51b815260206004820152601660248201527508aa4a4be84a09eaebe8482a68abea89e9ebe90928e960531b604482015290519081900360640190fd5b60006151d6836153b3565b905060006151e48483614d0d565b905060006151fa866151f5856153ce565b6153dc565b90508161520b57925061114a915050565b600061521c87846305f5e100615433565b90506152288282614c44565b979650505050505050565b6000808284106152495750508082036000615251565b505081810360015b9250929050565b614d79823083614e5e565b306000908152602081905260409020548111156152be576040805162461bcd60e51b815260206004820152601460248201527311549497d25394d551919250d251539517d0905360621b604482015290519081900360640190fd5b306000908152602081905260409020546152d89082614d0d565b306000908152602081905260409020556002546152f59082614d0d565b60025560408051828152905160009130917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a350565b614d79308383614e5e565b306000908152602081905260409020546153589082614f6e565b306000908152602081905260409020556002546153759082614f6e565b60025560408051828152905130916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a350565b6000670de0b6b3a76400006153c7836153ce565b0292915050565b670de0b6b3a7640000900490565b600080600283066153f557670de0b6b3a76400006153f7565b835b90506002830492505b8215614ae0576154108485614c44565b93506002830615615428576154258185614c44565b90505b600283049250615400565b600082818061544a87670de0b6b3a7640000615233565b9092509050670de0b6b3a764000080600060015b888410615502576000670de0b6b3a7640000820290506000806154928a61548d85670de0b6b3a7640000614d0d565b615233565b915091506154a4876137a2848c614c44565b96506154b0878461501d565b9650866154bf57505050615502565b87156154c9579315935b80156154d3579315935b84156154ea576154e38688614d0d565b95506154f7565b6154f48688614f6e565b95505b50505060010161545e565b5090999850505050505050505056fe8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925a265627a7a72315820a7cc59bf2e6a7f7e1b4b74a8cd9ed2d0ffd628b1ea9f42f7d1fac030c01dad6764736f6c634300050c0032a265627a7a7231582005315e7c86c7d033bccb795209cffa72b15d7ad406c56c8c8bf9443c514155da64736f6c634300050c0032"

// DeployBFactory deploys a new Ethereum contract, binding an instance of BFactory to it.
func DeployBFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(BFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BFactory{BFactoryCaller: BFactoryCaller{contract: contract}, BFactoryTransactor: BFactoryTransactor{contract: contract}, BFactoryFilterer: BFactoryFilterer{contract: contract}}, nil
}

// BFactory is an auto generated Go binding around an Ethereum contract.
type BFactory struct {
	BFactoryCaller     // Read-only binding to the contract
	BFactoryTransactor // Write-only binding to the contract
	BFactoryFilterer   // Log filterer for contract events
}

// BFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BFactorySession struct {
	Contract     *BFactory         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BFactoryCallerSession struct {
	Contract *BFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BFactoryTransactorSession struct {
	Contract     *BFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BFactoryRaw struct {
	Contract *BFactory // Generic contract binding to access the raw methods on
}

// BFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BFactoryCallerRaw struct {
	Contract *BFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// BFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BFactoryTransactorRaw struct {
	Contract *BFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBFactory creates a new instance of BFactory, bound to a specific deployed contract.
func NewBFactory(address common.Address, backend bind.ContractBackend) (*BFactory, error) {
	contract, err := bindBFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BFactory{BFactoryCaller: BFactoryCaller{contract: contract}, BFactoryTransactor: BFactoryTransactor{contract: contract}, BFactoryFilterer: BFactoryFilterer{contract: contract}}, nil
}

// NewBFactoryCaller creates a new read-only instance of BFactory, bound to a specific deployed contract.
func NewBFactoryCaller(address common.Address, caller bind.ContractCaller) (*BFactoryCaller, error) {
	contract, err := bindBFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BFactoryCaller{contract: contract}, nil
}

// NewBFactoryTransactor creates a new write-only instance of BFactory, bound to a specific deployed contract.
func NewBFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*BFactoryTransactor, error) {
	contract, err := bindBFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BFactoryTransactor{contract: contract}, nil
}

// NewBFactoryFilterer creates a new log filterer instance of BFactory, bound to a specific deployed contract.
func NewBFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*BFactoryFilterer, error) {
	contract, err := bindBFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BFactoryFilterer{contract: contract}, nil
}

// bindBFactory binds a generic wrapper to an already deployed contract.
func bindBFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BFactory *BFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BFactory.Contract.BFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BFactory *BFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BFactory.Contract.BFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BFactory *BFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BFactory.Contract.BFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BFactory *BFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BFactory *BFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BFactory *BFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BFactory.Contract.contract.Transact(opts, method, params...)
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_BFactory *BFactoryCaller) GetBLabs(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BFactory.contract.Call(opts, &out, "getBLabs")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_BFactory *BFactorySession) GetBLabs() (common.Address, error) {
	return _BFactory.Contract.GetBLabs(&_BFactory.CallOpts)
}

// GetBLabs is a free data retrieval call binding the contract method 0x36ffb167.
//
// Solidity: function getBLabs() view returns(address)
func (_BFactory *BFactoryCallerSession) GetBLabs() (common.Address, error) {
	return _BFactory.Contract.GetBLabs(&_BFactory.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BFactory *BFactoryCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BFactory.contract.Call(opts, &out, "getColor")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BFactory *BFactorySession) GetColor() ([32]byte, error) {
	return _BFactory.Contract.GetColor(&_BFactory.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BFactory *BFactoryCallerSession) GetColor() ([32]byte, error) {
	return _BFactory.Contract.GetColor(&_BFactory.CallOpts)
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_BFactory *BFactoryCaller) IsBPool(opts *bind.CallOpts, b common.Address) (bool, error) {
	var out []interface{}
	err := _BFactory.contract.Call(opts, &out, "isBPool", b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_BFactory *BFactorySession) IsBPool(b common.Address) (bool, error) {
	return _BFactory.Contract.IsBPool(&_BFactory.CallOpts, b)
}

// IsBPool is a free data retrieval call binding the contract method 0xc2bb6dc2.
//
// Solidity: function isBPool(address b) view returns(bool)
func (_BFactory *BFactoryCallerSession) IsBPool(b common.Address) (bool, error) {
	return _BFactory.Contract.IsBPool(&_BFactory.CallOpts, b)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_BFactory *BFactoryTransactor) Collect(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _BFactory.contract.Transact(opts, "collect", pool)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_BFactory *BFactorySession) Collect(pool common.Address) (*types.Transaction, error) {
	return _BFactory.Contract.Collect(&_BFactory.TransactOpts, pool)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address pool) returns()
func (_BFactory *BFactoryTransactorSession) Collect(pool common.Address) (*types.Transaction, error) {
	return _BFactory.Contract.Collect(&_BFactory.TransactOpts, pool)
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_BFactory *BFactoryTransactor) NewBPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BFactory.contract.Transact(opts, "newBPool")
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_BFactory *BFactorySession) NewBPool() (*types.Transaction, error) {
	return _BFactory.Contract.NewBPool(&_BFactory.TransactOpts)
}

// NewBPool is a paid mutator transaction binding the contract method 0xd556c5dc.
//
// Solidity: function newBPool() returns(address)
func (_BFactory *BFactoryTransactorSession) NewBPool() (*types.Transaction, error) {
	return _BFactory.Contract.NewBPool(&_BFactory.TransactOpts)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_BFactory *BFactoryTransactor) SetBLabs(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _BFactory.contract.Transact(opts, "setBLabs", b)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_BFactory *BFactorySession) SetBLabs(b common.Address) (*types.Transaction, error) {
	return _BFactory.Contract.SetBLabs(&_BFactory.TransactOpts, b)
}

// SetBLabs is a paid mutator transaction binding the contract method 0xc6ce34fb.
//
// Solidity: function setBLabs(address b) returns()
func (_BFactory *BFactoryTransactorSession) SetBLabs(b common.Address) (*types.Transaction, error) {
	return _BFactory.Contract.SetBLabs(&_BFactory.TransactOpts, b)
}

// BFactoryLOGBLABSIterator is returned from FilterLOGBLABS and is used to iterate over the raw logs and unpacked data for LOGBLABS events raised by the BFactory contract.
type BFactoryLOGBLABSIterator struct {
	Event *BFactoryLOGBLABS // Event containing the contract specifics and raw log

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
func (it *BFactoryLOGBLABSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BFactoryLOGBLABS)
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
		it.Event = new(BFactoryLOGBLABS)
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
func (it *BFactoryLOGBLABSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BFactoryLOGBLABSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BFactoryLOGBLABS represents a LOGBLABS event raised by the BFactory contract.
type BFactoryLOGBLABS struct {
	Caller common.Address
	Blabs  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLOGBLABS is a free log retrieval operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_BFactory *BFactoryFilterer) FilterLOGBLABS(opts *bind.FilterOpts, caller []common.Address, blabs []common.Address) (*BFactoryLOGBLABSIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var blabsRule []interface{}
	for _, blabsItem := range blabs {
		blabsRule = append(blabsRule, blabsItem)
	}

	logs, sub, err := _BFactory.contract.FilterLogs(opts, "LOG_BLABS", callerRule, blabsRule)
	if err != nil {
		return nil, err
	}
	return &BFactoryLOGBLABSIterator{contract: _BFactory.contract, event: "LOG_BLABS", logs: logs, sub: sub}, nil
}

// WatchLOGBLABS is a free log subscription operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_BFactory *BFactoryFilterer) WatchLOGBLABS(opts *bind.WatchOpts, sink chan<- *BFactoryLOGBLABS, caller []common.Address, blabs []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var blabsRule []interface{}
	for _, blabsItem := range blabs {
		blabsRule = append(blabsRule, blabsItem)
	}

	logs, sub, err := _BFactory.contract.WatchLogs(opts, "LOG_BLABS", callerRule, blabsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BFactoryLOGBLABS)
				if err := _BFactory.contract.UnpackLog(event, "LOG_BLABS", log); err != nil {
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

// ParseLOGBLABS is a log parse operation binding the contract event 0xf586fa6ee1fc42f5b727f3b214ccbd0b6d7e698c45d49ba32f224fbb8670155d.
//
// Solidity: event LOG_BLABS(address indexed caller, address indexed blabs)
func (_BFactory *BFactoryFilterer) ParseLOGBLABS(log types.Log) (*BFactoryLOGBLABS, error) {
	event := new(BFactoryLOGBLABS)
	if err := _BFactory.contract.UnpackLog(event, "LOG_BLABS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BFactoryLOGNEWPOOLIterator is returned from FilterLOGNEWPOOL and is used to iterate over the raw logs and unpacked data for LOGNEWPOOL events raised by the BFactory contract.
type BFactoryLOGNEWPOOLIterator struct {
	Event *BFactoryLOGNEWPOOL // Event containing the contract specifics and raw log

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
func (it *BFactoryLOGNEWPOOLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BFactoryLOGNEWPOOL)
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
		it.Event = new(BFactoryLOGNEWPOOL)
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
func (it *BFactoryLOGNEWPOOLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BFactoryLOGNEWPOOLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BFactoryLOGNEWPOOL represents a LOGNEWPOOL event raised by the BFactory contract.
type BFactoryLOGNEWPOOL struct {
	Caller common.Address
	Pool   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLOGNEWPOOL is a free log retrieval operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_BFactory *BFactoryFilterer) FilterLOGNEWPOOL(opts *bind.FilterOpts, caller []common.Address, pool []common.Address) (*BFactoryLOGNEWPOOLIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _BFactory.contract.FilterLogs(opts, "LOG_NEW_POOL", callerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &BFactoryLOGNEWPOOLIterator{contract: _BFactory.contract, event: "LOG_NEW_POOL", logs: logs, sub: sub}, nil
}

// WatchLOGNEWPOOL is a free log subscription operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_BFactory *BFactoryFilterer) WatchLOGNEWPOOL(opts *bind.WatchOpts, sink chan<- *BFactoryLOGNEWPOOL, caller []common.Address, pool []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _BFactory.contract.WatchLogs(opts, "LOG_NEW_POOL", callerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BFactoryLOGNEWPOOL)
				if err := _BFactory.contract.UnpackLog(event, "LOG_NEW_POOL", log); err != nil {
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

// ParseLOGNEWPOOL is a log parse operation binding the contract event 0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945.
//
// Solidity: event LOG_NEW_POOL(address indexed caller, address indexed pool)
func (_BFactory *BFactoryFilterer) ParseLOGNEWPOOL(log types.Log) (*BFactoryLOGNEWPOOL, error) {
	event := new(BFactoryLOGNEWPOOL)
	if err := _BFactory.contract.UnpackLog(event, "LOG_NEW_POOL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BMathABI is the input ABI used to generate the binding from.
const BMathABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"BONE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BPOW_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EXIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INIT_POOL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_IN_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OUT_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_TOTAL_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcInGivenOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcOutGivenIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcPoolInGivenSingleOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcPoolOutGivenSingleIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcSingleInGivenPoolOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcSingleOutGivenPoolIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcSpotPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BMathFuncSigs maps the 4-byte function signature to its string representation.
var BMathFuncSigs = map[string]string{
	"c36596a6": "BONE()",
	"189d00ca": "BPOW_PRECISION()",
	"c6580d12": "EXIT_FEE()",
	"9381cd2b": "INIT_POOL_SUPPLY()",
	"b0e0d136": "MAX_BOUND_TOKENS()",
	"bc694ea2": "MAX_BPOW_BASE()",
	"bc063e1a": "MAX_FEE()",
	"ec093021": "MAX_IN_RATIO()",
	"992e2a92": "MAX_OUT_RATIO()",
	"09a3bbe4": "MAX_TOTAL_WEIGHT()",
	"e4a28a52": "MAX_WEIGHT()",
	"867378c5": "MIN_BALANCE()",
	"b7b800a4": "MIN_BOUND_TOKENS()",
	"ba019dab": "MIN_BPOW_BASE()",
	"76c7a3c7": "MIN_FEE()",
	"218b5382": "MIN_WEIGHT()",
	"f8d6aed4": "calcInGivenOut(uint256,uint256,uint256,uint256,uint256,uint256)",
	"ba9530a6": "calcOutGivenIn(uint256,uint256,uint256,uint256,uint256,uint256)",
	"82f652ad": "calcPoolInGivenSingleOut(uint256,uint256,uint256,uint256,uint256,uint256)",
	"8656b653": "calcPoolOutGivenSingleIn(uint256,uint256,uint256,uint256,uint256,uint256)",
	"5c1bbaf7": "calcSingleInGivenPoolOut(uint256,uint256,uint256,uint256,uint256,uint256)",
	"89298012": "calcSingleOutGivenPoolIn(uint256,uint256,uint256,uint256,uint256,uint256)",
	"a221ee49": "calcSpotPrice(uint256,uint256,uint256,uint256,uint256)",
	"9a86139b": "getColor()",
}

// BMathBin is the compiled bytecode used for deploying new contracts.
var BMathBin = "0x608060405234801561001057600080fd5b50610dbb806100206000396000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c8063a221ee49116100c3578063bc694ea21161007c578063bc694ea21461032b578063c36596a61461017f578063c6580d1214610333578063e4a28a521461015d578063ec0930211461033b578063f8d6aed41461034357610158565b8063a221ee491461029b578063b0e0d136146102d0578063b7b800a4146102d8578063ba019dab146102e0578063ba9530a6146102e8578063bc063e1a1461032357610158565b80638656b653116101155780638656b65314610205578063867378c51461024057806389298012146102485780639381cd2b14610283578063992e2a921461028b5780639a86139b1461029357610158565b806309a3bbe41461015d578063189d00ca14610177578063218b53821461017f5780635c1bbaf71461018757806376c7a3c7146101c257806382f652ad146101ca575b600080fd5b61016561037e565b60408051918252519081900360200190f35b61016561038b565b61016561039f565b610165600480360360c081101561019d57600080fd5b5080359060208101359060408101359060608101359060808101359060a001356103ab565b610165610463565b610165600480360360c08110156101e057600080fd5b5080359060208101359060408101359060608101359060808101359060a00135610475565b610165600480360360c081101561021b57600080fd5b5080359060208101359060408101359060608101359060808101359060a00135610534565b6101656105d5565b610165600480360360c081101561025e57600080fd5b5080359060208101359060408101359060608101359060808101359060a001356105e9565b610165610699565b6101656106a6565b6101656106b2565b610165600480360360a08110156102b157600080fd5b50803590602081013590604081013590606081013590608001356106bf565b610165610724565b610165610729565b61016561072e565b610165600480360360c08110156102fe57600080fd5b5080359060208101359060408101359060608101359060808101359060a00135610733565b6101656107b4565b6101656107c4565b6101656107d0565b6101656107d5565b610165600480360360c081101561035957600080fd5b5080359060208101359060408101359060608101359060808101359060a001356107e5565b6802b5e3af16b188000081565b6402540be400670de0b6b3a76400005b0481565b670de0b6b3a764000081565b6000806103b88786610868565b905060006103c6878661097d565b905060006103d48289610868565b905060006103ea670de0b6b3a764000085610868565b905060006103f883836109d1565b90506000610406828e610adf565b90506000610414828f610ba1565b9050600061043361042d670de0b6b3a76400008a610ba1565b8b610adf565b90506104508261044b670de0b6b3a764000084610ba1565b610868565b9f9e505050505050505050505050505050565b620f4240670de0b6b3a764000061039b565b6000806104828786610868565b90506000610498670de0b6b3a764000083610ba1565b905060006104a68286610adf565b905060006104c08761044b670de0b6b3a764000085610ba1565b905060006104ce8c83610ba1565b905060006104dc828e610868565b905060006104ea82886109d1565b905060006104f8828e610adf565b905060006105068e83610ba1565b905061051f8161044b670de0b6b3a76400006000610ba1565b99505050505050505050509695505050505050565b6000806105418786610868565b9050600061056061055a670de0b6b3a764000084610ba1565b85610adf565b9050600061057f8661057a670de0b6b3a764000085610ba1565b610adf565b9050600061058d8b8361097d565b9050600061059b828d610868565b905060006105a982876109d1565b905060006105b7828d610adf565b90506105c3818d610ba1565b9e9d5050505050505050505050505050565b64e8d4a51000670de0b6b3a764000061039b565b6000806105f68786610868565b905060006106118561057a670de0b6b3a76400006000610ba1565b9050600061061f8883610ba1565b9050600061062d828a610868565b9050600061064c82610647670de0b6b3a764000088610868565b6109d1565b9050600061065a828e610adf565b905060006106688e83610ba1565b9050600061068161042d670de0b6b3a76400008a610ba1565b90506104508261057a670de0b6b3a764000084610ba1565b68056bc75e2d6310000081565b6704a03ce68d21555681565b6542524f4e5a4560d01b90565b6000806106cc8787610868565b905060006106da8686610868565b905060006106e88383610868565b9050600061070a670de0b6b3a764000061044b670de0b6b3a764000089610ba1565b90506107168282610adf565b9a9950505050505050505050565b600881565b600281565b600181565b6000806107408786610868565b90506000610756670de0b6b3a764000085610ba1565b90506107628582610adf565b905060006107748a61044b8c8561097d565b9050600061078282856109d1565b90506000610798670de0b6b3a764000083610ba1565b90506107a48a82610adf565b9c9b505050505050505050505050565b600a670de0b6b3a764000061039b565b671bc16d674ec7ffff81565b600081565b6002670de0b6b3a764000061039b565b6000806107f28588610868565b905060006108008786610ba1565b9050600061080e8883610868565b9050600061081c82856109d1565b905061083081670de0b6b3a7640000610ba1565b9050610844670de0b6b3a764000087610ba1565b94506108596108538c83610adf565b86610868565b9b9a5050505050505050505050565b6000816108ab576040805162461bcd60e51b815260206004820152600c60248201526b4552525f4449565f5a45524f60a01b604482015290519081900360640190fd5b670de0b6b3a764000083028315806108d35750670de0b6b3a76400008482816108d057fe5b04145b610917576040805162461bcd60e51b815260206004820152601060248201526f11549497d1125597d25395115493905360821b604482015290519081900360640190fd5b60028304810181811015610965576040805162461bcd60e51b815260206004820152601060248201526f11549497d1125597d25395115493905360821b604482015290519081900360640190fd5b600084828161097057fe5b0493505050505b92915050565b6000828201838110156109ca576040805162461bcd60e51b815260206004820152601060248201526f4552525f4144445f4f564552464c4f5760801b604482015290519081900360640190fd5b9392505050565b60006001831015610a21576040805162461bcd60e51b81526020600482015260156024820152744552525f42504f575f424153455f544f4f5f4c4f5760581b604482015290519081900360640190fd5b671bc16d674ec7ffff831115610a77576040805162461bcd60e51b815260206004820152601660248201527508aa4a4be84a09eaebe8482a68abea89e9ebe90928e960531b604482015290519081900360640190fd5b6000610a8283610c03565b90506000610a908483610ba1565b90506000610aa686610aa185610c1e565b610c2c565b905081610ab7579250610977915050565b6000610ac887846305f5e100610c83565b9050610ad48282610adf565b979650505050505050565b6000828202831580610af9575082848281610af657fe5b04145b610b3d576040805162461bcd60e51b815260206004820152601060248201526f4552525f4d554c5f4f564552464c4f5760801b604482015290519081900360640190fd5b6706f05b59d3b20000810181811015610b90576040805162461bcd60e51b815260206004820152601060248201526f4552525f4d554c5f4f564552464c4f5760801b604482015290519081900360640190fd5b6000670de0b6b3a764000082610970565b6000806000610bb08585610d61565b915091508015610bfb576040805162461bcd60e51b81526020600482015260116024820152704552525f5355425f554e444552464c4f5760781b604482015290519081900360640190fd5b509392505050565b6000670de0b6b3a7640000610c1783610c1e565b0292915050565b670de0b6b3a7640000900490565b60008060028306610c4557670de0b6b3a7640000610c47565b835b90506002830492505b82156109ca57610c608485610adf565b93506002830615610c7857610c758185610adf565b90505b600283049250610c50565b6000828180610c9a87670de0b6b3a7640000610d61565b9092509050670de0b6b3a764000080600060015b888410610d52576000670de0b6b3a764000082029050600080610ce28a610cdd85670de0b6b3a7640000610ba1565b610d61565b91509150610cf48761057a848c610adf565b9650610d008784610868565b965086610d0f57505050610d52565b8715610d19579315935b8015610d23579315935b8415610d3a57610d338688610ba1565b9550610d47565b610d44868861097d565b95505b505050600101610cae565b50909998505050505050505050565b600080828410610d775750508082036000610d7f565b505081810360015b925092905056fea265627a7a72315820bbf4ec9c6e437f50fa556b0724926195682cb574d4ddb55271b6363fea587e2364736f6c634300050c0032"

// DeployBMath deploys a new Ethereum contract, binding an instance of BMath to it.
func DeployBMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BMath, error) {
	parsed, err := abi.JSON(strings.NewReader(BMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BMath{BMathCaller: BMathCaller{contract: contract}, BMathTransactor: BMathTransactor{contract: contract}, BMathFilterer: BMathFilterer{contract: contract}}, nil
}

// BMath is an auto generated Go binding around an Ethereum contract.
type BMath struct {
	BMathCaller     // Read-only binding to the contract
	BMathTransactor // Write-only binding to the contract
	BMathFilterer   // Log filterer for contract events
}

// BMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type BMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BMathSession struct {
	Contract     *BMath            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BMathCallerSession struct {
	Contract *BMathCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BMathTransactorSession struct {
	Contract     *BMathTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type BMathRaw struct {
	Contract *BMath // Generic contract binding to access the raw methods on
}

// BMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BMathCallerRaw struct {
	Contract *BMathCaller // Generic read-only contract binding to access the raw methods on
}

// BMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BMathTransactorRaw struct {
	Contract *BMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBMath creates a new instance of BMath, bound to a specific deployed contract.
func NewBMath(address common.Address, backend bind.ContractBackend) (*BMath, error) {
	contract, err := bindBMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BMath{BMathCaller: BMathCaller{contract: contract}, BMathTransactor: BMathTransactor{contract: contract}, BMathFilterer: BMathFilterer{contract: contract}}, nil
}

// NewBMathCaller creates a new read-only instance of BMath, bound to a specific deployed contract.
func NewBMathCaller(address common.Address, caller bind.ContractCaller) (*BMathCaller, error) {
	contract, err := bindBMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BMathCaller{contract: contract}, nil
}

// NewBMathTransactor creates a new write-only instance of BMath, bound to a specific deployed contract.
func NewBMathTransactor(address common.Address, transactor bind.ContractTransactor) (*BMathTransactor, error) {
	contract, err := bindBMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BMathTransactor{contract: contract}, nil
}

// NewBMathFilterer creates a new log filterer instance of BMath, bound to a specific deployed contract.
func NewBMathFilterer(address common.Address, filterer bind.ContractFilterer) (*BMathFilterer, error) {
	contract, err := bindBMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BMathFilterer{contract: contract}, nil
}

// bindBMath binds a generic wrapper to an already deployed contract.
func bindBMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BMath *BMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BMath.Contract.BMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BMath *BMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BMath.Contract.BMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BMath *BMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BMath.Contract.BMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BMath *BMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BMath *BMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BMath *BMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BMath.Contract.contract.Transact(opts, method, params...)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BMath *BMathCaller) BONE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "BONE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BMath *BMathSession) BONE() (*big.Int, error) {
	return _BMath.Contract.BONE(&_BMath.CallOpts)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BMath *BMathCallerSession) BONE() (*big.Int, error) {
	return _BMath.Contract.BONE(&_BMath.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BMath *BMathCaller) BPOWPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "BPOW_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BMath *BMathSession) BPOWPRECISION() (*big.Int, error) {
	return _BMath.Contract.BPOWPRECISION(&_BMath.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BMath *BMathCallerSession) BPOWPRECISION() (*big.Int, error) {
	return _BMath.Contract.BPOWPRECISION(&_BMath.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BMath *BMathCaller) EXITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "EXIT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BMath *BMathSession) EXITFEE() (*big.Int, error) {
	return _BMath.Contract.EXITFEE(&_BMath.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BMath *BMathCallerSession) EXITFEE() (*big.Int, error) {
	return _BMath.Contract.EXITFEE(&_BMath.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BMath *BMathCaller) INITPOOLSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "INIT_POOL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BMath *BMathSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _BMath.Contract.INITPOOLSUPPLY(&_BMath.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BMath *BMathCallerSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _BMath.Contract.INITPOOLSUPPLY(&_BMath.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BMath *BMathCaller) MAXBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "MAX_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BMath *BMathSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _BMath.Contract.MAXBOUNDTOKENS(&_BMath.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BMath *BMathCallerSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _BMath.Contract.MAXBOUNDTOKENS(&_BMath.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BMath *BMathCaller) MAXBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "MAX_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BMath *BMathSession) MAXBPOWBASE() (*big.Int, error) {
	return _BMath.Contract.MAXBPOWBASE(&_BMath.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BMath *BMathCallerSession) MAXBPOWBASE() (*big.Int, error) {
	return _BMath.Contract.MAXBPOWBASE(&_BMath.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BMath *BMathCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BMath *BMathSession) MAXFEE() (*big.Int, error) {
	return _BMath.Contract.MAXFEE(&_BMath.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BMath *BMathCallerSession) MAXFEE() (*big.Int, error) {
	return _BMath.Contract.MAXFEE(&_BMath.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BMath *BMathCaller) MAXINRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "MAX_IN_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BMath *BMathSession) MAXINRATIO() (*big.Int, error) {
	return _BMath.Contract.MAXINRATIO(&_BMath.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BMath *BMathCallerSession) MAXINRATIO() (*big.Int, error) {
	return _BMath.Contract.MAXINRATIO(&_BMath.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BMath *BMathCaller) MAXOUTRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "MAX_OUT_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BMath *BMathSession) MAXOUTRATIO() (*big.Int, error) {
	return _BMath.Contract.MAXOUTRATIO(&_BMath.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BMath *BMathCallerSession) MAXOUTRATIO() (*big.Int, error) {
	return _BMath.Contract.MAXOUTRATIO(&_BMath.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BMath *BMathCaller) MAXTOTALWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "MAX_TOTAL_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BMath *BMathSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _BMath.Contract.MAXTOTALWEIGHT(&_BMath.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BMath *BMathCallerSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _BMath.Contract.MAXTOTALWEIGHT(&_BMath.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BMath *BMathCaller) MAXWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "MAX_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BMath *BMathSession) MAXWEIGHT() (*big.Int, error) {
	return _BMath.Contract.MAXWEIGHT(&_BMath.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BMath *BMathCallerSession) MAXWEIGHT() (*big.Int, error) {
	return _BMath.Contract.MAXWEIGHT(&_BMath.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BMath *BMathCaller) MINBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "MIN_BALANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BMath *BMathSession) MINBALANCE() (*big.Int, error) {
	return _BMath.Contract.MINBALANCE(&_BMath.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BMath *BMathCallerSession) MINBALANCE() (*big.Int, error) {
	return _BMath.Contract.MINBALANCE(&_BMath.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BMath *BMathCaller) MINBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "MIN_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BMath *BMathSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _BMath.Contract.MINBOUNDTOKENS(&_BMath.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BMath *BMathCallerSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _BMath.Contract.MINBOUNDTOKENS(&_BMath.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BMath *BMathCaller) MINBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "MIN_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BMath *BMathSession) MINBPOWBASE() (*big.Int, error) {
	return _BMath.Contract.MINBPOWBASE(&_BMath.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BMath *BMathCallerSession) MINBPOWBASE() (*big.Int, error) {
	return _BMath.Contract.MINBPOWBASE(&_BMath.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BMath *BMathCaller) MINFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "MIN_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BMath *BMathSession) MINFEE() (*big.Int, error) {
	return _BMath.Contract.MINFEE(&_BMath.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BMath *BMathCallerSession) MINFEE() (*big.Int, error) {
	return _BMath.Contract.MINFEE(&_BMath.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BMath *BMathCaller) MINWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "MIN_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BMath *BMathSession) MINWEIGHT() (*big.Int, error) {
	return _BMath.Contract.MINWEIGHT(&_BMath.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BMath *BMathCallerSession) MINWEIGHT() (*big.Int, error) {
	return _BMath.Contract.MINWEIGHT(&_BMath.CallOpts)
}

// CalcInGivenOut is a free data retrieval call binding the contract method 0xf8d6aed4.
//
// Solidity: function calcInGivenOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BMath *BMathCaller) CalcInGivenOut(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "calcInGivenOut", tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountOut, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcInGivenOut is a free data retrieval call binding the contract method 0xf8d6aed4.
//
// Solidity: function calcInGivenOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BMath *BMathSession) CalcInGivenOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BMath.Contract.CalcInGivenOut(&_BMath.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountOut, swapFee)
}

// CalcInGivenOut is a free data retrieval call binding the contract method 0xf8d6aed4.
//
// Solidity: function calcInGivenOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BMath *BMathCallerSession) CalcInGivenOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BMath.Contract.CalcInGivenOut(&_BMath.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountOut, swapFee)
}

// CalcOutGivenIn is a free data retrieval call binding the contract method 0xba9530a6.
//
// Solidity: function calcOutGivenIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BMath *BMathCaller) CalcOutGivenIn(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "calcOutGivenIn", tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountIn, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcOutGivenIn is a free data retrieval call binding the contract method 0xba9530a6.
//
// Solidity: function calcOutGivenIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BMath *BMathSession) CalcOutGivenIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BMath.Contract.CalcOutGivenIn(&_BMath.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountIn, swapFee)
}

// CalcOutGivenIn is a free data retrieval call binding the contract method 0xba9530a6.
//
// Solidity: function calcOutGivenIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BMath *BMathCallerSession) CalcOutGivenIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BMath.Contract.CalcOutGivenIn(&_BMath.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountIn, swapFee)
}

// CalcPoolInGivenSingleOut is a free data retrieval call binding the contract method 0x82f652ad.
//
// Solidity: function calcPoolInGivenSingleOut(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 poolAmountIn)
func (_BMath *BMathCaller) CalcPoolInGivenSingleOut(opts *bind.CallOpts, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "calcPoolInGivenSingleOut", tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, tokenAmountOut, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcPoolInGivenSingleOut is a free data retrieval call binding the contract method 0x82f652ad.
//
// Solidity: function calcPoolInGivenSingleOut(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 poolAmountIn)
func (_BMath *BMathSession) CalcPoolInGivenSingleOut(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BMath.Contract.CalcPoolInGivenSingleOut(&_BMath.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, tokenAmountOut, swapFee)
}

// CalcPoolInGivenSingleOut is a free data retrieval call binding the contract method 0x82f652ad.
//
// Solidity: function calcPoolInGivenSingleOut(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 poolAmountIn)
func (_BMath *BMathCallerSession) CalcPoolInGivenSingleOut(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BMath.Contract.CalcPoolInGivenSingleOut(&_BMath.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, tokenAmountOut, swapFee)
}

// CalcPoolOutGivenSingleIn is a free data retrieval call binding the contract method 0x8656b653.
//
// Solidity: function calcPoolOutGivenSingleIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 poolAmountOut)
func (_BMath *BMathCaller) CalcPoolOutGivenSingleIn(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "calcPoolOutGivenSingleIn", tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, tokenAmountIn, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcPoolOutGivenSingleIn is a free data retrieval call binding the contract method 0x8656b653.
//
// Solidity: function calcPoolOutGivenSingleIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 poolAmountOut)
func (_BMath *BMathSession) CalcPoolOutGivenSingleIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BMath.Contract.CalcPoolOutGivenSingleIn(&_BMath.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, tokenAmountIn, swapFee)
}

// CalcPoolOutGivenSingleIn is a free data retrieval call binding the contract method 0x8656b653.
//
// Solidity: function calcPoolOutGivenSingleIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 poolAmountOut)
func (_BMath *BMathCallerSession) CalcPoolOutGivenSingleIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BMath.Contract.CalcPoolOutGivenSingleIn(&_BMath.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, tokenAmountIn, swapFee)
}

// CalcSingleInGivenPoolOut is a free data retrieval call binding the contract method 0x5c1bbaf7.
//
// Solidity: function calcSingleInGivenPoolOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BMath *BMathCaller) CalcSingleInGivenPoolOut(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "calcSingleInGivenPoolOut", tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, poolAmountOut, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcSingleInGivenPoolOut is a free data retrieval call binding the contract method 0x5c1bbaf7.
//
// Solidity: function calcSingleInGivenPoolOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BMath *BMathSession) CalcSingleInGivenPoolOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BMath.Contract.CalcSingleInGivenPoolOut(&_BMath.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, poolAmountOut, swapFee)
}

// CalcSingleInGivenPoolOut is a free data retrieval call binding the contract method 0x5c1bbaf7.
//
// Solidity: function calcSingleInGivenPoolOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BMath *BMathCallerSession) CalcSingleInGivenPoolOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BMath.Contract.CalcSingleInGivenPoolOut(&_BMath.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, poolAmountOut, swapFee)
}

// CalcSingleOutGivenPoolIn is a free data retrieval call binding the contract method 0x89298012.
//
// Solidity: function calcSingleOutGivenPoolIn(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BMath *BMathCaller) CalcSingleOutGivenPoolIn(opts *bind.CallOpts, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "calcSingleOutGivenPoolIn", tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, poolAmountIn, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcSingleOutGivenPoolIn is a free data retrieval call binding the contract method 0x89298012.
//
// Solidity: function calcSingleOutGivenPoolIn(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BMath *BMathSession) CalcSingleOutGivenPoolIn(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BMath.Contract.CalcSingleOutGivenPoolIn(&_BMath.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, poolAmountIn, swapFee)
}

// CalcSingleOutGivenPoolIn is a free data retrieval call binding the contract method 0x89298012.
//
// Solidity: function calcSingleOutGivenPoolIn(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BMath *BMathCallerSession) CalcSingleOutGivenPoolIn(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BMath.Contract.CalcSingleOutGivenPoolIn(&_BMath.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, poolAmountIn, swapFee)
}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_BMath *BMathCaller) CalcSpotPrice(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "calcSpotPrice", tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_BMath *BMathSession) CalcSpotPrice(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BMath.Contract.CalcSpotPrice(&_BMath.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)
}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_BMath *BMathCallerSession) CalcSpotPrice(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BMath.Contract.CalcSpotPrice(&_BMath.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BMath *BMathCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BMath.contract.Call(opts, &out, "getColor")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BMath *BMathSession) GetColor() ([32]byte, error) {
	return _BMath.Contract.GetColor(&_BMath.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BMath *BMathCallerSession) GetColor() ([32]byte, error) {
	return _BMath.Contract.GetColor(&_BMath.CallOpts)
}

// BNumABI is the input ABI used to generate the binding from.
const BNumABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"BONE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BPOW_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EXIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INIT_POOL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_IN_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OUT_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_TOTAL_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BNumFuncSigs maps the 4-byte function signature to its string representation.
var BNumFuncSigs = map[string]string{
	"c36596a6": "BONE()",
	"189d00ca": "BPOW_PRECISION()",
	"c6580d12": "EXIT_FEE()",
	"9381cd2b": "INIT_POOL_SUPPLY()",
	"b0e0d136": "MAX_BOUND_TOKENS()",
	"bc694ea2": "MAX_BPOW_BASE()",
	"bc063e1a": "MAX_FEE()",
	"ec093021": "MAX_IN_RATIO()",
	"992e2a92": "MAX_OUT_RATIO()",
	"09a3bbe4": "MAX_TOTAL_WEIGHT()",
	"e4a28a52": "MAX_WEIGHT()",
	"867378c5": "MIN_BALANCE()",
	"b7b800a4": "MIN_BOUND_TOKENS()",
	"ba019dab": "MIN_BPOW_BASE()",
	"76c7a3c7": "MIN_FEE()",
	"218b5382": "MIN_WEIGHT()",
	"9a86139b": "getColor()",
}

// BNumBin is the compiled bytecode used for deploying new contracts.
var BNumBin = "0x608060405234801561001057600080fd5b50610288806100206000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c8063b0e0d136116100a2578063bc694ea211610071578063bc694ea214610182578063c36596a614610132578063c6580d121461018a578063e4a28a5214610110578063ec093021146101925761010b565b8063b0e0d13614610162578063b7b800a41461016a578063ba019dab14610172578063bc063e1a1461017a5761010b565b8063867378c5116100de578063867378c5146101425780639381cd2b1461014a578063992e2a92146101525780639a86139b1461015a5761010b565b806309a3bbe414610110578063189d00ca1461012a578063218b53821461013257806376c7a3c71461013a575b600080fd5b61011861019a565b60408051918252519081900360200190f35b6101186101a7565b6101186101bb565b6101186101c7565b6101186101d9565b6101186101ed565b6101186101fa565b610118610206565b610118610213565b610118610218565b61011861021d565b610118610222565b610118610232565b61011861023e565b610118610243565b6802b5e3af16b188000081565b6402540be400670de0b6b3a76400005b0481565b670de0b6b3a764000081565b620f4240670de0b6b3a76400006101b7565b64e8d4a51000670de0b6b3a76400006101b7565b68056bc75e2d6310000081565b6704a03ce68d21555681565b6542524f4e5a4560d01b90565b600881565b600281565b600181565b600a670de0b6b3a76400006101b7565b671bc16d674ec7ffff81565b600081565b6002670de0b6b3a76400006101b756fea265627a7a72315820a123677dc0dae08e1bece2d4aa7f6370a4d69204a4f6bd75e326f3eac61876c264736f6c634300050c0032"

// DeployBNum deploys a new Ethereum contract, binding an instance of BNum to it.
func DeployBNum(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BNum, error) {
	parsed, err := abi.JSON(strings.NewReader(BNumABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BNumBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BNum{BNumCaller: BNumCaller{contract: contract}, BNumTransactor: BNumTransactor{contract: contract}, BNumFilterer: BNumFilterer{contract: contract}}, nil
}

// BNum is an auto generated Go binding around an Ethereum contract.
type BNum struct {
	BNumCaller     // Read-only binding to the contract
	BNumTransactor // Write-only binding to the contract
	BNumFilterer   // Log filterer for contract events
}

// BNumCaller is an auto generated read-only Go binding around an Ethereum contract.
type BNumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BNumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BNumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BNumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BNumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BNumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BNumSession struct {
	Contract     *BNum             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BNumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BNumCallerSession struct {
	Contract *BNumCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BNumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BNumTransactorSession struct {
	Contract     *BNumTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BNumRaw is an auto generated low-level Go binding around an Ethereum contract.
type BNumRaw struct {
	Contract *BNum // Generic contract binding to access the raw methods on
}

// BNumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BNumCallerRaw struct {
	Contract *BNumCaller // Generic read-only contract binding to access the raw methods on
}

// BNumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BNumTransactorRaw struct {
	Contract *BNumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBNum creates a new instance of BNum, bound to a specific deployed contract.
func NewBNum(address common.Address, backend bind.ContractBackend) (*BNum, error) {
	contract, err := bindBNum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BNum{BNumCaller: BNumCaller{contract: contract}, BNumTransactor: BNumTransactor{contract: contract}, BNumFilterer: BNumFilterer{contract: contract}}, nil
}

// NewBNumCaller creates a new read-only instance of BNum, bound to a specific deployed contract.
func NewBNumCaller(address common.Address, caller bind.ContractCaller) (*BNumCaller, error) {
	contract, err := bindBNum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BNumCaller{contract: contract}, nil
}

// NewBNumTransactor creates a new write-only instance of BNum, bound to a specific deployed contract.
func NewBNumTransactor(address common.Address, transactor bind.ContractTransactor) (*BNumTransactor, error) {
	contract, err := bindBNum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BNumTransactor{contract: contract}, nil
}

// NewBNumFilterer creates a new log filterer instance of BNum, bound to a specific deployed contract.
func NewBNumFilterer(address common.Address, filterer bind.ContractFilterer) (*BNumFilterer, error) {
	contract, err := bindBNum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BNumFilterer{contract: contract}, nil
}

// bindBNum binds a generic wrapper to an already deployed contract.
func bindBNum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BNumABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BNum *BNumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BNum.Contract.BNumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BNum *BNumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BNum.Contract.BNumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BNum *BNumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BNum.Contract.BNumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BNum *BNumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BNum.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BNum *BNumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BNum.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BNum *BNumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BNum.Contract.contract.Transact(opts, method, params...)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BNum *BNumCaller) BONE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "BONE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BNum *BNumSession) BONE() (*big.Int, error) {
	return _BNum.Contract.BONE(&_BNum.CallOpts)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BNum *BNumCallerSession) BONE() (*big.Int, error) {
	return _BNum.Contract.BONE(&_BNum.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BNum *BNumCaller) BPOWPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "BPOW_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BNum *BNumSession) BPOWPRECISION() (*big.Int, error) {
	return _BNum.Contract.BPOWPRECISION(&_BNum.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BNum *BNumCallerSession) BPOWPRECISION() (*big.Int, error) {
	return _BNum.Contract.BPOWPRECISION(&_BNum.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BNum *BNumCaller) EXITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "EXIT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BNum *BNumSession) EXITFEE() (*big.Int, error) {
	return _BNum.Contract.EXITFEE(&_BNum.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BNum *BNumCallerSession) EXITFEE() (*big.Int, error) {
	return _BNum.Contract.EXITFEE(&_BNum.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BNum *BNumCaller) INITPOOLSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "INIT_POOL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BNum *BNumSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _BNum.Contract.INITPOOLSUPPLY(&_BNum.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BNum *BNumCallerSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _BNum.Contract.INITPOOLSUPPLY(&_BNum.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BNum *BNumCaller) MAXBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MAX_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BNum *BNumSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _BNum.Contract.MAXBOUNDTOKENS(&_BNum.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BNum *BNumCallerSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _BNum.Contract.MAXBOUNDTOKENS(&_BNum.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BNum *BNumCaller) MAXBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MAX_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BNum *BNumSession) MAXBPOWBASE() (*big.Int, error) {
	return _BNum.Contract.MAXBPOWBASE(&_BNum.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BNum *BNumCallerSession) MAXBPOWBASE() (*big.Int, error) {
	return _BNum.Contract.MAXBPOWBASE(&_BNum.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BNum *BNumCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BNum *BNumSession) MAXFEE() (*big.Int, error) {
	return _BNum.Contract.MAXFEE(&_BNum.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BNum *BNumCallerSession) MAXFEE() (*big.Int, error) {
	return _BNum.Contract.MAXFEE(&_BNum.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BNum *BNumCaller) MAXINRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MAX_IN_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BNum *BNumSession) MAXINRATIO() (*big.Int, error) {
	return _BNum.Contract.MAXINRATIO(&_BNum.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BNum *BNumCallerSession) MAXINRATIO() (*big.Int, error) {
	return _BNum.Contract.MAXINRATIO(&_BNum.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BNum *BNumCaller) MAXOUTRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MAX_OUT_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BNum *BNumSession) MAXOUTRATIO() (*big.Int, error) {
	return _BNum.Contract.MAXOUTRATIO(&_BNum.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BNum *BNumCallerSession) MAXOUTRATIO() (*big.Int, error) {
	return _BNum.Contract.MAXOUTRATIO(&_BNum.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BNum *BNumCaller) MAXTOTALWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MAX_TOTAL_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BNum *BNumSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _BNum.Contract.MAXTOTALWEIGHT(&_BNum.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BNum *BNumCallerSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _BNum.Contract.MAXTOTALWEIGHT(&_BNum.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BNum *BNumCaller) MAXWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MAX_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BNum *BNumSession) MAXWEIGHT() (*big.Int, error) {
	return _BNum.Contract.MAXWEIGHT(&_BNum.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BNum *BNumCallerSession) MAXWEIGHT() (*big.Int, error) {
	return _BNum.Contract.MAXWEIGHT(&_BNum.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BNum *BNumCaller) MINBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MIN_BALANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BNum *BNumSession) MINBALANCE() (*big.Int, error) {
	return _BNum.Contract.MINBALANCE(&_BNum.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BNum *BNumCallerSession) MINBALANCE() (*big.Int, error) {
	return _BNum.Contract.MINBALANCE(&_BNum.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BNum *BNumCaller) MINBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MIN_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BNum *BNumSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _BNum.Contract.MINBOUNDTOKENS(&_BNum.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BNum *BNumCallerSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _BNum.Contract.MINBOUNDTOKENS(&_BNum.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BNum *BNumCaller) MINBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MIN_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BNum *BNumSession) MINBPOWBASE() (*big.Int, error) {
	return _BNum.Contract.MINBPOWBASE(&_BNum.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BNum *BNumCallerSession) MINBPOWBASE() (*big.Int, error) {
	return _BNum.Contract.MINBPOWBASE(&_BNum.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BNum *BNumCaller) MINFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MIN_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BNum *BNumSession) MINFEE() (*big.Int, error) {
	return _BNum.Contract.MINFEE(&_BNum.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BNum *BNumCallerSession) MINFEE() (*big.Int, error) {
	return _BNum.Contract.MINFEE(&_BNum.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BNum *BNumCaller) MINWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MIN_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BNum *BNumSession) MINWEIGHT() (*big.Int, error) {
	return _BNum.Contract.MINWEIGHT(&_BNum.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BNum *BNumCallerSession) MINWEIGHT() (*big.Int, error) {
	return _BNum.Contract.MINWEIGHT(&_BNum.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BNum *BNumCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "getColor")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BNum *BNumSession) GetColor() ([32]byte, error) {
	return _BNum.Contract.GetColor(&_BNum.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BNum *BNumCallerSession) GetColor() ([32]byte, error) {
	return _BNum.Contract.GetColor(&_BNum.CallOpts)
}

// BPoolABI is the input ABI used to generate the binding from.
const BPoolABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LOG_CALL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"name\":\"LOG_EXIT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"name\":\"LOG_JOIN\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"name\":\"LOG_SWAP\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"BONE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BPOW_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EXIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INIT_POOL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_IN_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OUT_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_TOTAL_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denorm\",\"type\":\"uint256\"}],\"name\":\"bind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcInGivenOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcOutGivenIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcPoolInGivenSingleOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcPoolOutGivenSingleIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcSingleInGivenPoolOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcSingleOutGivenPoolIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcSpotPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmountsOut\",\"type\":\"uint256[]\"}],\"name\":\"exitPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPoolAmountIn\",\"type\":\"uint256\"}],\"name\":\"exitswapExternAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"exitswapPoolAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDenormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFinalTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getNormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getSpotPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getSpotPriceSansFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSwapFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalDenormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"gulp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"isBound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPublicSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"maxAmountsIn\",\"type\":\"uint256[]\"}],\"name\":\"joinPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPoolAmountOut\",\"type\":\"uint256\"}],\"name\":\"joinswapExternAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"}],\"name\":\"joinswapPoolAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denorm\",\"type\":\"uint256\"}],\"name\":\"rebind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"public_\",\"type\":\"bool\"}],\"name\":\"setPublicSwap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"setSwapFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"swapExactAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceAfter\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"swapExactAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceAfter\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unbind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BPoolFuncSigs maps the 4-byte function signature to its string representation.
var BPoolFuncSigs = map[string]string{
	"c36596a6": "BONE()",
	"189d00ca": "BPOW_PRECISION()",
	"c6580d12": "EXIT_FEE()",
	"9381cd2b": "INIT_POOL_SUPPLY()",
	"b0e0d136": "MAX_BOUND_TOKENS()",
	"bc694ea2": "MAX_BPOW_BASE()",
	"bc063e1a": "MAX_FEE()",
	"ec093021": "MAX_IN_RATIO()",
	"992e2a92": "MAX_OUT_RATIO()",
	"09a3bbe4": "MAX_TOTAL_WEIGHT()",
	"e4a28a52": "MAX_WEIGHT()",
	"867378c5": "MIN_BALANCE()",
	"b7b800a4": "MIN_BOUND_TOKENS()",
	"ba019dab": "MIN_BPOW_BASE()",
	"76c7a3c7": "MIN_FEE()",
	"218b5382": "MIN_WEIGHT()",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"e4e1e538": "bind(address,uint256,uint256)",
	"f8d6aed4": "calcInGivenOut(uint256,uint256,uint256,uint256,uint256,uint256)",
	"ba9530a6": "calcOutGivenIn(uint256,uint256,uint256,uint256,uint256,uint256)",
	"82f652ad": "calcPoolInGivenSingleOut(uint256,uint256,uint256,uint256,uint256,uint256)",
	"8656b653": "calcPoolOutGivenSingleIn(uint256,uint256,uint256,uint256,uint256,uint256)",
	"5c1bbaf7": "calcSingleInGivenPoolOut(uint256,uint256,uint256,uint256,uint256,uint256)",
	"89298012": "calcSingleOutGivenPoolIn(uint256,uint256,uint256,uint256,uint256,uint256)",
	"a221ee49": "calcSpotPrice(uint256,uint256,uint256,uint256,uint256)",
	"313ce567": "decimals()",
	"66188463": "decreaseApproval(address,uint256)",
	"b02f0b73": "exitPool(uint256,uint256[])",
	"02c96748": "exitswapExternAmountOut(address,uint256,uint256)",
	"46ab38f1": "exitswapPoolAmountIn(address,uint256,uint256)",
	"4bb278f3": "finalize()",
	"f8b2cb4f": "getBalance(address)",
	"9a86139b": "getColor()",
	"3018205f": "getController()",
	"cc77828d": "getCurrentTokens()",
	"948d8ce6": "getDenormalizedWeight(address)",
	"be3bbd2e": "getFinalTokens()",
	"f1b8a9b7": "getNormalizedWeight(address)",
	"cd2ed8fb": "getNumTokens()",
	"15e84af9": "getSpotPrice(address,address)",
	"1446a7ff": "getSpotPriceSansFee(address,address)",
	"d4cadf68": "getSwapFee()",
	"936c3477": "getTotalDenormalizedWeight()",
	"8c28cbe8": "gulp(address)",
	"d73dd623": "increaseApproval(address,uint256)",
	"2f37b624": "isBound(address)",
	"8d4e4083": "isFinalized()",
	"fde924f7": "isPublicSwap()",
	"4f69c0d4": "joinPool(uint256,uint256[])",
	"5db34277": "joinswapExternAmountIn(address,uint256,uint256)",
	"6d06dfa0": "joinswapPoolAmountOut(address,uint256,uint256)",
	"06fdde03": "name()",
	"3fdddaa2": "rebind(address,uint256,uint256)",
	"92eefe9b": "setController(address)",
	"49b59552": "setPublicSwap(bool)",
	"34e19907": "setSwapFee(uint256)",
	"8201aa3f": "swapExactAmountIn(address,uint256,address,uint256,uint256)",
	"7c5e9ea4": "swapExactAmountOut(address,uint256,address,uint256,uint256)",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"cf5e7bd3": "unbind(address)",
}

// BPoolBin is the compiled bytecode used for deploying new contracts.
var BPoolBin = "0x60c0604052601360808190527f42616c616e63657220506f6f6c20546f6b656e0000000000000000000000000060a0908152620000409160039190620000f4565b506040805180820190915260038082527f425054000000000000000000000000000000000000000000000000000000000060209092019182526200008791600491620000f4565b506005805460ff19166012179055348015620000a257600080fd5b50600680546005805462010000600160b01b031916336201000081029190911790915564e8d4a510006007556001600160a01b03199091161760ff60a01b191690556008805460ff1916905562000199565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200013757805160ff191683800117855562000167565b8280016001018555821562000167579182015b82811115620001675782518255916020019190600101906200014a565b506200017592915062000179565b5090565b6200019691905b8082111562000175576000815560010162000180565b90565b61556680620001a96000396000f3fe608060405234801561001057600080fd5b50600436106103db5760003560e01c80638d4e40831161020a578063bc694ea211610125578063d73dd623116100b8578063ec09302111610087578063ec09302114610c6e578063f1b8a9b714610c76578063f8b2cb4f14610c9c578063f8d6aed414610cc2578063fde924f714610cfd576103db565b8063d73dd62314610be2578063dd62ed3e14610c0e578063e4a28a52146104e1578063e4e1e53814610c3c576103db565b8063cc77828d116100f4578063cc77828d14610ba4578063cd2ed8fb14610bac578063cf5e7bd314610bb4578063d4cadf6814610bda576103db565b8063bc694ea214610b3c578063be3bbd2e14610b44578063c36596a614610555578063c6580d1214610b9c576103db565b8063a221ee491161019d578063b7b800a41161016c578063b7b800a414610ae9578063ba019dab14610af1578063ba9530a614610af9578063bc063e1a14610b34576103db565b8063a221ee4914610a09578063a9059cbb14610a3e578063b02f0b7314610a6a578063b0e0d13614610ae1576103db565b8063948d8ce6116101d9578063948d8ce6146109cb57806395d89b41146109f1578063992e2a92146109f95780639a86139b14610a01576103db565b80638d4e40831461098d57806392eefe9b14610995578063936c3477146109bb5780639381cd2b146109c3576103db565b806349b59552116102fa57806376c7a3c71161028d5780638656b6531161025c5780638656b653146108e9578063867378c514610924578063892980121461092c5780638c28cbe814610967576103db565b806376c7a3c71461080d5780637c5e9ea4146108155780638201aa3f1461086e57806382f652ad146108ae576103db565b80635db34277116102c95780635db342771461075757806366188463146107895780636d06dfa0146107b557806370a08231146107e7576103db565b806349b595521461067e5780634bb278f31461069d5780634f69c0d4146106a55780635c1bbaf71461071c576103db565b8063218b538211610372578063313ce56711610341578063313ce567146105dd57806334e19907146105fb5780633fdddaa21461061a57806346ab38f11461064c576103db565b8063218b53821461055557806323b872dd1461055d5780632f37b624146105935780633018205f146105b9576103db565b80631446a7ff116103ae5780631446a7ff146104e957806315e84af91461051757806318160ddd14610545578063189d00ca1461054d576103db565b806302c96748146103e057806306fdde0314610424578063095ea7b3146104a157806309a3bbe4146104e1575b600080fd5b610412600480360360608110156103f657600080fd5b506001600160a01b038135169060208101359060400135610d05565b60408051918252519081900360200190f35b61042c611065565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561046657818101518382015260200161044e565b50505050905090810190601f1680156104935780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6104cd600480360360408110156104b757600080fd5b506001600160a01b0381351690602001356110fb565b604080519115158252519081900360200190f35b610412611150565b610412600480360360408110156104ff57600080fd5b506001600160a01b038135811691602001351661115d565b6104126004803603604081101561052d57600080fd5b506001600160a01b03813581169160200135166112b2565b6104126113fe565b610412611404565b610412611418565b6104cd6004803603606081101561057357600080fd5b506001600160a01b03813581169160208101359091169060400135611424565b6104cd600480360360208110156105a957600080fd5b50356001600160a01b031661157e565b6105c161159c565b604080516001600160a01b039092168252519081900360200190f35b6105e56115fa565b6040805160ff9092168252519081900360200190f35b6106186004803603602081101561061157600080fd5b5035611603565b005b6106186004803603606081101561063057600080fd5b506001600160a01b038135169060208101359060400135611800565b6104126004803603606081101561066257600080fd5b506001600160a01b038135169060208101359060400135611c0d565b6106186004803603602081101561069457600080fd5b50351515611f0c565b61061861208f565b610618600480360360408110156106bb57600080fd5b813591908101906040810160208201356401000000008111156106dd57600080fd5b8201836020820111156106ef57600080fd5b8035906020019184602083028401116401000000008311171561071157600080fd5b509092509050612286565b610412600480360360c081101561073257600080fd5b5080359060208101359060408101359060608101359060808101359060a0013561257f565b6104126004803603606081101561076d57600080fd5b506001600160a01b038135169060208101359060400135612637565b6104cd6004803603604081101561079f57600080fd5b506001600160a01b03813516906020013561291a565b610412600480360360608110156107cb57600080fd5b506001600160a01b0381351690602081013590604001356129f2565b610412600480360360208110156107fd57600080fd5b50356001600160a01b0316612d03565b610412612d1e565b610855600480360360a081101561082b57600080fd5b506001600160a01b0381358116916020810135916040820135169060608101359060800135612d30565b6040805192835260208301919091528051918290030190f35b610855600480360360a081101561088457600080fd5b506001600160a01b03813581169160208101359160408201351690606081013590608001356131f3565b610412600480360360c08110156108c457600080fd5b5080359060208101359060408101359060608101359060808101359060a0013561369d565b610412600480360360c08110156108ff57600080fd5b5080359060208101359060408101359060608101359060808101359060a0013561375c565b6104126137fd565b610412600480360360c081101561094257600080fd5b5080359060208101359060408101359060608101359060808101359060a00135613811565b6106186004803603602081101561097d57600080fd5b50356001600160a01b03166138c1565b6104cd613a75565b610618600480360360208110156109ab57600080fd5b50356001600160a01b0316613a7e565b610412613bbc565b610412613c11565b610412600480360360208110156109e157600080fd5b50356001600160a01b0316613c1e565b61042c613ce8565b610412613d49565b610412613d55565b610412600480360360a0811015610a1f57600080fd5b5080359060208101359060408101359060608101359060800135613d62565b6104cd60048036036040811015610a5457600080fd5b506001600160a01b038135169060200135613dc7565b61061860048036036040811015610a8057600080fd5b81359190810190604081016020820135640100000000811115610aa257600080fd5b820183602082011115610ab457600080fd5b80359060200191846020830284011164010000000083111715610ad657600080fd5b509092509050613ddd565b610412614124565b610412614129565b61041261412e565b610412600480360360c0811015610b0f57600080fd5b5080359060208101359060408101359060608101359060808101359060a00135614133565b6104126141b4565b6104126141c4565b610b4c6141d0565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015610b88578181015183820152602001610b70565b505050509050019250505060405180910390f35b6104126142c8565b610b4c6142cd565b61041261431b565b61061860048036036020811015610bca57600080fd5b50356001600160a01b0316614321565b6104126146a3565b6104cd60048036036040811015610bf857600080fd5b506001600160a01b0381351690602001356146f8565b61041260048036036040811015610c2457600080fd5b506001600160a01b0381358116916020013516614779565b61061860048036036060811015610c5257600080fd5b506001600160a01b0381351690602081013590604001356147a4565b6104126149fb565b61041260048036036020811015610c8c57600080fd5b50356001600160a01b0316614a0b565b61041260048036036020811015610cb257600080fd5b50356001600160a01b0316614ae7565b610412600480360360c0811015610cd857600080fd5b5080359060208101359060408101359060608101359060808101359060a00135614bb1565b6104cd614c34565b6000336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff1615610db3576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff00191661010017905560085460ff16610e0d576040805162461bcd60e51b815260206004820152601160248201527011549497d393d517d19253905312569151607a1b604482015290519081900360640190fd5b6001600160a01b0384166000908152600a602052604090205460ff16610e6a576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b0384166000908152600a60205260409020600390810154610e9f91670de0b6b3a76400005b04600101614c44565b831115610ee7576040805162461bcd60e51b81526020600482015260116024820152704552525f4d41585f4f55545f524154494f60781b604482015290519081900360640190fd5b6001600160a01b0384166000908152600a6020526040902060038101546002808301549054600b54600754610f219493929190899061369d565b915081610f67576040805162461bcd60e51b815260206004820152600f60248201526e08aa4a4be9a82a890be82a0a0a49eb608b1b604482015290519081900360640190fd5b82821115610fab576040805162461bcd60e51b815260206004820152600c60248201526b22a9292fa624a6a4aa2fa4a760a11b604482015290519081900360640190fd5b610fb9816003015485614d0d565b60038201556000610fca8382614c44565b6040805187815290519192506001600160a01b0388169133917fe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed919081900360200190a36110183384614d6f565b61102a6110258483614d0d565b614d7d565b600554611046906201000090046001600160a01b031682614d89565b611051863387614d93565b50506005805461ff00191690559392505050565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156110f15780601f106110c6576101008083540402835291602001916110f1565b820191906000526020600020905b8154815290600101906020018083116110d457829003601f168201915b5050505050905090565b3360008181526001602090815260408083206001600160a01b03871680855290835281842086905581518681529151939490939092600080516020615512833981519152928290030190a35060015b92915050565b6802b5e3af16b188000081565b600554600090610100900460ff16156111ab576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6001600160a01b0383166000908152600a602052604090205460ff16611208576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b0382166000908152600a602052604090205460ff16611265576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b038084166000908152600a602052604080822092851682528120600380840154600280860154928401549084015493946112a99492939290613d62565b95945050505050565b600554600090610100900460ff1615611300576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6001600160a01b0383166000908152600a602052604090205460ff1661135d576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b0382166000908152600a602052604090205460ff166113ba576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b038084166000908152600a60205260408082209285168252902060038083015460028085015492840154908401546007546112a994929190613d62565b60025490565b6402540be400670de0b6b3a76400005b0481565b670de0b6b3a764000081565b6000336001600160a01b038516148061146057506001600160a01b03841660009081526001602090815260408083203384529091529020548211155b6114a9576040805162461bcd60e51b815260206004820152601560248201527422a9292fa12a27a5a2a72fa120a22fa1a0a62622a960591b604482015290519081900360640190fd5b6114b4848484614e5e565b336001600160a01b038516148015906114f257506001600160a01b038416600090815260016020908152604080832033845290915290205460001914155b15611574576001600160a01b03841660009081526001602090815260408083203384529091529020546115259083614d0d565b6001600160a01b03858116600090815260016020908152604080832033808552908352928190208590558051948552519287169391926000805160206155128339815191529281900390910190a35b5060019392505050565b6001600160a01b03166000908152600a602052604090205460ff1690565b600554600090610100900460ff16156115ea576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b506006546001600160a01b031690565b60055460ff1690565b336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff16156116af576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff00191661010017905560085460ff1615611709576040805162461bcd60e51b815260206004820152601060248201526f11549497d254d7d1925390531256915160821b604482015290519081900360640190fd5b6006546001600160a01b0316331461175d576040805162461bcd60e51b815260206004820152601260248201527122a9292fa727aa2fa1a7a72a2927a62622a960711b604482015290519081900360640190fd5b64e8d4a510008110156117a5576040805162461bcd60e51b815260206004820152600b60248201526a4552525f4d494e5f46454560a81b604482015290519081900360640190fd5b67016345785d8a00008111156117f0576040805162461bcd60e51b815260206004820152600b60248201526a4552525f4d41585f46454560a81b604482015290519081900360640190fd5b6007556005805461ff0019169055565b336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff16156118ac576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff001916610100179055600654336001600160a01b0390911614611911576040805162461bcd60e51b815260206004820152601260248201527122a9292fa727aa2fa1a7a72a2927a62622a960711b604482015290519081900360640190fd5b6001600160a01b0383166000908152600a602052604090205460ff1661196e576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b60085460ff16156119b9576040805162461bcd60e51b815260206004820152601060248201526f11549497d254d7d1925390531256915160821b604482015290519081900360640190fd5b670de0b6b3a7640000811015611a07576040805162461bcd60e51b815260206004820152600e60248201526d11549497d3525397d5d15251d21560921b604482015290519081900360640190fd5b6802b5e3af16b1880000811115611a56576040805162461bcd60e51b815260206004820152600e60248201526d11549497d3505617d5d15251d21560921b604482015290519081900360640190fd5b620f4240821015611aa0576040805162461bcd60e51b815260206004820152600f60248201526e4552525f4d494e5f42414c414e434560881b604482015290519081900360640190fd5b6001600160a01b0383166000908152600a602052604090206002015480821115611b3757611ad9600b54611ad48484614d0d565b614f6e565b600b8190556802b5e3af16b18800001015611b32576040805162461bcd60e51b815260206004820152601460248201527311549497d3505617d513d5105317d5d15251d21560621b604482015290519081900360640190fd5b611b58565b80821015611b5857611b54600b54611b4f8385614d0d565b614d0d565b600b555b6001600160a01b0384166000908152600a602052604090206002810183905560030180549084905580841115611ba157611b9c8533611b978785614d0d565b614fbb565b611bfb565b80841015611bfb576000611bb58286614d0d565b90506000611bc4826000614c44565b9050611bda8733611bd58585614d0d565b614d93565b600554611bf89088906201000090046001600160a01b031683614d93565b50505b50506005805461ff0019169055505050565b6000336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff1615611cbb576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff00191661010017905560085460ff16611d15576040805162461bcd60e51b815260206004820152601160248201527011549497d393d517d19253905312569151607a1b604482015290519081900360640190fd5b6001600160a01b0384166000908152600a602052604090205460ff16611d72576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b0384166000908152600a6020526040902060038101546002808301549054600b54600754611dac94939291908990613811565b915082821015611df3576040805162461bcd60e51b815260206004820152600d60248201526c11549497d31253525517d3d555609a1b604482015290519081900360640190fd5b6001600160a01b0385166000908152600a60205260409020600390810154611e2391670de0b6b3a7640000610e96565b821115611e6b576040805162461bcd60e51b81526020600482015260116024820152704552525f4d41585f4f55545f524154494f60781b604482015290519081900360640190fd5b611e79816003015483614d0d565b60038201556000611e8a8582614c44565b6040805185815290519192506001600160a01b0388169133917fe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed919081900360200190a3611ed83386614d6f565b611ee56110258683614d0d565b600554611f01906201000090046001600160a01b031682614d89565b611051863385614d93565b336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff1615611fb8576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff00191661010017905560085460ff1615612012576040805162461bcd60e51b815260206004820152601060248201526f11549497d254d7d1925390531256915160821b604482015290519081900360640190fd5b6006546001600160a01b03163314612066576040805162461bcd60e51b815260206004820152601260248201527122a9292fa727aa2fa1a7a72a2927a62622a960711b604482015290519081900360640190fd5b60068054911515600160a01b0260ff60a01b199092169190911790556005805461ff0019169055565b336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff161561213b576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff001916610100179055600654336001600160a01b03909116146121a0576040805162461bcd60e51b815260206004820152601260248201527122a9292fa727aa2fa1a7a72a2927a62622a960711b604482015290519081900360640190fd5b60085460ff16156121eb576040805162461bcd60e51b815260206004820152601060248201526f11549497d254d7d1925390531256915160821b604482015290519081900360640190fd5b60095460021115612234576040805162461bcd60e51b815260206004820152600e60248201526d4552525f4d494e5f544f4b454e5360901b604482015290519081900360640190fd5b6008805460ff191660011790556006805460ff60a01b1916600160a01b17905561226668056bc75e2d63100000615014565b6122793368056bc75e2d63100000614d89565b6005805461ff0019169055565b336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff1615612332576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff00191661010017905560085460ff1661238c576040805162461bcd60e51b815260206004820152601160248201527011549497d393d517d19253905312569151607a1b604482015290519081900360640190fd5b60006123966113fe565b905060006123a4858361501d565b9050806123ea576040805162461bcd60e51b815260206004820152600f60248201526e08aa4a4be9a82a890be82a0a0a49eb608b1b604482015290519081900360640190fd5b60005b60095481101561256b5760006009828154811061240657fe5b60009182526020808320909101546001600160a01b0316808352600a90915260408220600301549092509061243b8583614c44565b905080612481576040805162461bcd60e51b815260206004820152600f60248201526e08aa4a4be9a82a890be82a0a0a49eb608b1b604482015290519081900360640190fd5b87878581811061248d57fe5b905060200201358111156124d7576040805162461bcd60e51b815260206004820152600c60248201526b22a9292fa624a6a4aa2fa4a760a11b604482015290519081900360640190fd5b6001600160a01b0383166000908152600a60205260409020600301546124fd9082614f6e565b6001600160a01b0384166000818152600a60209081526040918290206003019390935580518481529051919233927f63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a9281900390910190a3612560833383614fbb565b5050506001016123ed565b5061257585615014565b611bfb3386614d89565b60008061258c878661501d565b9050600061259a8786614f6e565b905060006125a8828961501d565b905060006125be670de0b6b3a76400008561501d565b905060006125cc8383615125565b905060006125da828e614c44565b905060006125e8828f614d0d565b90506000612607612601670de0b6b3a76400008a614d0d565b8b614c44565b90506126248261261f670de0b6b3a764000084614d0d565b61501d565b9f9e505050505050505050505050505050565b6000336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff16156126e5576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff00191661010017905560085460ff1661273f576040805162461bcd60e51b815260206004820152601160248201527011549497d393d517d19253905312569151607a1b604482015290519081900360640190fd5b6001600160a01b0384166000908152600a602052604090205460ff1661279c576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b0384166000908152600a60205260409020600301546127ce906002670de0b6b3a76400005b04614c44565b831115612815576040805162461bcd60e51b815260206004820152601060248201526f4552525f4d41585f494e5f524154494f60801b604482015290519081900360640190fd5b6001600160a01b0384166000908152600a6020526040902060038101546002808301549054600b5460075461284f9493929190899061375c565b915082821015612896576040805162461bcd60e51b815260206004820152600d60248201526c11549497d31253525517d3d555609a1b604482015290519081900360640190fd5b6128a4816003015485614f6e565b60038201556040805185815290516001600160a01b0387169133917f63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a9181900360200190a36128f282615014565b6128fc3383614d89565b612907853386614fbb565b506005805461ff00191690559392505050565b3360009081526001602090815260408083206001600160a01b03861684529091528120548083111561296f573360009081526001602090815260408083206001600160a01b038816845290915281205561299e565b6129798184614d0d565b3360009081526001602090815260408083206001600160a01b03891684529091529020555b3360008181526001602090815260408083206001600160a01b038916808552908352928190205481519081529051929392600080516020615512833981519152929181900390910190a35060019392505050565b6000336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff1615612aa0576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff00191661010017905560085460ff16612afa576040805162461bcd60e51b815260206004820152601160248201527011549497d393d517d19253905312569151607a1b604482015290519081900360640190fd5b6001600160a01b0384166000908152600a602052604090205460ff16612b57576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b0384166000908152600a6020526040902060038101546002808301549054600b54600754612b919493929190899061257f565b915081612bd7576040805162461bcd60e51b815260206004820152600f60248201526e08aa4a4be9a82a890be82a0a0a49eb608b1b604482015290519081900360640190fd5b82821115612c1b576040805162461bcd60e51b815260206004820152600c60248201526b22a9292fa624a6a4aa2fa4a760a11b604482015290519081900360640190fd5b6001600160a01b0385166000908152600a6020526040902060030154612c4b906002670de0b6b3a76400006127c8565b821115612c92576040805162461bcd60e51b815260206004820152601060248201526f4552525f4d41585f494e5f524154494f60801b604482015290519081900360640190fd5b612ca0816003015483614f6e565b60038201556040805183815290516001600160a01b0387169133917f63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a9181900360200190a3612cee84615014565b612cf83385614d89565b612907853384614fbb565b6001600160a01b031660009081526020819052604090205490565b620f4240670de0b6b3a7640000611414565b60408051602080825236908201819052600092839233926001600160e01b03198535169285929081908101848480828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff1615612dcd576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff0019166101001790556001600160a01b0387166000908152600a602052604090205460ff16612e39576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b0385166000908152600a602052604090205460ff16612e96576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b600654600160a01b900460ff16612eea576040805162461bcd60e51b81526020600482015260136024820152724552525f535741505f4e4f545f5055424c494360681b604482015290519081900360640190fd5b6001600160a01b038088166000908152600a602052604080822092881682529020600380820154612f2391670de0b6b3a7640000610e96565b861115612f6b576040805162461bcd60e51b81526020600482015260116024820152704552525f4d41585f4f55545f524154494f60781b604482015290519081900360640190fd5b6000612f8c8360030154846002015484600301548560020154600754613d62565b905085811115612fd9576040805162461bcd60e51b81526020600482015260136024820152724552525f4241445f4c494d49545f505249434560681b604482015290519081900360640190fd5b612ff983600301548460020154846003015485600201548b600754614bb1565b94508885111561303f576040805162461bcd60e51b815260206004820152600c60248201526b22a9292fa624a6a4aa2fa4a760a11b604482015290519081900360640190fd5b61304d836003015486614f6e565b8360030181905550613063826003015488614d0d565b600380840182905584015460028086015490850154600754613086949190613d62565b9350808410156130cf576040805162461bcd60e51b815260206004820152600f60248201526e08aa4a4be9a82a890be82a0a0a49eb608b1b604482015290519081900360640190fd5b85841115613116576040805162461bcd60e51b815260206004820152600f60248201526e4552525f4c494d49545f505249434560881b604482015290519081900360640190fd5b613120858861501d565b811115613166576040805162461bcd60e51b815260206004820152600f60248201526e08aa4a4be9a82a890be82a0a0a49eb608b1b604482015290519081900360640190fd5b876001600160a01b03168a6001600160a01b0316336001600160a01b03167f908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378888b604051808381526020018281526020019250505060405180910390a46131ce8a3387614fbb565b6131d9883389614d93565b5050506005805461ff001916905590969095509350505050565b60408051602080825236908201819052600092839233926001600160e01b03198535169285929081908101848480828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff1615613290576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff0019166101001790556001600160a01b0387166000908152600a602052604090205460ff166132fc576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b0385166000908152600a602052604090205460ff16613359576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b600654600160a01b900460ff166133ad576040805162461bcd60e51b81526020600482015260136024820152724552525f535741505f4e4f545f5055424c494360681b604482015290519081900360640190fd5b6001600160a01b038088166000908152600a60205260408082209288168252902060038201546133e7906002670de0b6b3a76400006127c8565b88111561342e576040805162461bcd60e51b815260206004820152601060248201526f4552525f4d41585f494e5f524154494f60801b604482015290519081900360640190fd5b600061344f8360030154846002015484600301548560020154600754613d62565b90508581111561349c576040805162461bcd60e51b81526020600482015260136024820152724552525f4241445f4c494d49545f505249434560681b604482015290519081900360640190fd5b6134bc83600301548460020154846003015485600201548d600754614133565b945086851015613503576040805162461bcd60e51b815260206004820152600d60248201526c11549497d31253525517d3d555609a1b604482015290519081900360640190fd5b61351183600301548a614f6e565b8360030181905550613527826003015486614d0d565b60038084018290558401546002808601549085015460075461354a949190613d62565b935080841015613593576040805162461bcd60e51b815260206004820152600f60248201526e08aa4a4be9a82a890be82a0a0a49eb608b1b604482015290519081900360640190fd5b858411156135da576040805162461bcd60e51b815260206004820152600f60248201526e4552525f4c494d49545f505249434560881b604482015290519081900360640190fd5b6135e4898661501d565b81111561362a576040805162461bcd60e51b815260206004820152600f60248201526e08aa4a4be9a82a890be82a0a0a49eb608b1b604482015290519081900360640190fd5b876001600160a01b03168a6001600160a01b0316336001600160a01b03167f908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d433788c89604051808381526020018281526020019250505060405180910390a46136928a338b614fbb565b6131d9883387614d93565b6000806136aa878661501d565b905060006136c0670de0b6b3a764000083614d0d565b905060006136ce8286614c44565b905060006136e88761261f670de0b6b3a764000085614d0d565b905060006136f68c83614d0d565b90506000613704828e61501d565b905060006137128288615125565b90506000613720828e614c44565b9050600061372e8e83614d0d565b90506137478161261f670de0b6b3a76400006000614d0d565b99505050505050505050509695505050505050565b600080613769878661501d565b90506000613788613782670de0b6b3a764000084614d0d565b85614c44565b905060006137a7866137a2670de0b6b3a764000085614d0d565b614c44565b905060006137b58b83614f6e565b905060006137c3828d61501d565b905060006137d18287615125565b905060006137df828d614c44565b90506137eb818d614d0d565b9e9d5050505050505050505050505050565b64e8d4a51000670de0b6b3a7640000611414565b60008061381e878661501d565b90506000613839856137a2670de0b6b3a76400006000614d0d565b905060006138478883614d0d565b90506000613855828a61501d565b905060006138748261386f670de0b6b3a76400008861501d565b615125565b90506000613882828e614c44565b905060006138908e83614d0d565b905060006138a9612601670de0b6b3a76400008a614d0d565b9050612624826137a2670de0b6b3a764000084614d0d565b336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff161561396d576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff0019166101001790556001600160a01b0381166000908152600a602052604090205460ff166139d9576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b604080516370a0823160e01b815230600482015290516001600160a01b038316916370a08231916024808301926020929190829003018186803b158015613a1f57600080fd5b505afa158015613a33573d6000803e3d6000fd5b505050506040513d6020811015613a4957600080fd5b50516001600160a01b039091166000908152600a60205260409020600301556005805461ff0019169055565b60085460ff1690565b336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff1615613b2a576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff001916610100179055600654336001600160a01b0390911614613b8f576040805162461bcd60e51b815260206004820152601260248201527122a9292fa727aa2fa1a7a72a2927a62622a960711b604482015290519081900360640190fd5b600680546001600160a01b0319166001600160a01b03929092169190911790556005805461ff0019169055565b600554600090610100900460ff1615613c0a576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b50600b5490565b68056bc75e2d6310000081565b600554600090610100900460ff1615613c6c576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6001600160a01b0382166000908152600a602052604090205460ff16613cc9576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b506001600160a01b03166000908152600a602052604090206002015490565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156110f15780601f106110c6576101008083540402835291602001916110f1565b6704a03ce68d21555681565b6542524f4e5a4560d01b90565b600080613d6f878761501d565b90506000613d7d868661501d565b90506000613d8b838361501d565b90506000613dad670de0b6b3a764000061261f670de0b6b3a764000089614d0d565b9050613db98282614c44565b9a9950505050505050505050565b6000613dd4338484614e5e565b50600192915050565b336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff1615613e89576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff00191661010017905560085460ff16613ee3576040805162461bcd60e51b815260206004820152601160248201527011549497d393d517d19253905312569151607a1b604482015290519081900360640190fd5b6000613eed6113fe565b90506000613efc856000614c44565b90506000613f0a8683614d0d565b90506000613f18828561501d565b905080613f5e576040805162461bcd60e51b815260206004820152600f60248201526e08aa4a4be9a82a890be82a0a0a49eb608b1b604482015290519081900360640190fd5b613f683388614d6f565b600554613f84906201000090046001600160a01b031684614d89565b613f8d82614d7d565b60005b60095481101561410f57600060098281548110613fa957fe5b60009182526020808320909101546001600160a01b0316808352600a909152604082206003015490925090613fde8583614c44565b905080614024576040805162461bcd60e51b815260206004820152600f60248201526e08aa4a4be9a82a890be82a0a0a49eb608b1b604482015290519081900360640190fd5b89898581811061403057fe5b9050602002013581101561407b576040805162461bcd60e51b815260206004820152600d60248201526c11549497d31253525517d3d555609a1b604482015290519081900360640190fd5b6001600160a01b0383166000908152600a60205260409020600301546140a19082614d0d565b6001600160a01b0384166000818152600a60209081526040918290206003019390935580518481529051919233927fe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed9281900390910190a3614104833383614d93565b505050600101613f90565b50506005805461ff0019169055505050505050565b600881565b600281565b600181565b600080614140878661501d565b90506000614156670de0b6b3a764000085614d0d565b90506141628582614c44565b905060006141748a61261f8c85614f6e565b905060006141828285615125565b90506000614198670de0b6b3a764000083614d0d565b90506141a48a82614c44565b9c9b505050505050505050505050565b600a670de0b6b3a7640000611414565b671bc16d674ec7ffff81565b600554606090610100900460ff161561421e576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b60085460ff16614269576040805162461bcd60e51b815260206004820152601160248201527011549497d393d517d19253905312569151607a1b604482015290519081900360640190fd5b60098054806020026020016040519081016040528092919081815260200182805480156110f157602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116142a1575050505050905090565b600081565b600554606090610100900460ff1615614269576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b60095490565b336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2600554610100900460ff16156143cd576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6005805461ff001916610100179055600654336001600160a01b0390911614614432576040805162461bcd60e51b815260206004820152601260248201527122a9292fa727aa2fa1a7a72a2927a62622a960711b604482015290519081900360640190fd5b6001600160a01b0381166000908152600a602052604090205460ff1661448f576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b60085460ff16156144da576040805162461bcd60e51b815260206004820152601060248201526f11549497d254d7d1925390531256915160821b604482015290519081900360640190fd5b6001600160a01b0381166000908152600a6020526040812060030154906145018282614c44565b600b546001600160a01b0385166000908152600a602052604090206002015491925061452c91614d0d565b600b556001600160a01b0383166000908152600a602052604090206001015460098054600019810191908290811061456057fe5b600091825260209091200154600980546001600160a01b03909216918490811061458657fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081600a6000600985815481106145c657fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190206001015560098054806145f957fe5b60008281526020808220600019908401810180546001600160a01b031916905590920190925560408051608081018252838152808301848152818301858152606083018681526001600160a01b038c168752600a909552929094209051815460ff191690151517815592516001840155516002830155516003909101556146858533611bd58787614d0d565b600554611bfb9086906201000090046001600160a01b031685614d93565b600554600090610100900460ff16156146f1576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b5060075490565b3360009081526001602090815260408083206001600160a01b03861684529091528120546147269083614f6e565b3360008181526001602090815260408083206001600160a01b038916808552908352928190208590558051948552519193600080516020615512833981519152929081900390910190a350600192915050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b336001600160a01b03166000356001600160e01b0319166001600160e01b03191660003660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a26006546001600160a01b03163314614859576040805162461bcd60e51b815260206004820152601260248201527122a9292fa727aa2fa1a7a72a2927a62622a960711b604482015290519081900360640190fd5b6001600160a01b0383166000908152600a602052604090205460ff16156148b6576040805162461bcd60e51b815260206004820152600c60248201526b11549497d254d7d093d5539160a21b604482015290519081900360640190fd5b60085460ff1615614901576040805162461bcd60e51b815260206004820152601060248201526f11549497d254d7d1925390531256915160821b604482015290519081900360640190fd5b600954600811614949576040805162461bcd60e51b815260206004820152600e60248201526d4552525f4d41585f544f4b454e5360901b604482015290519081900360640190fd5b6040805160808101825260018082526009805460208085019182526000858701818152606087018281526001600160a01b038c16808452600a9094529782209651875460ff1916901515178755925186860155915160028601559451600390940193909355805491820181559091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0180546001600160a01b03191690911790556149f6838383611800565b505050565b6002670de0b6b3a7640000611414565b600554600090610100900460ff1615614a59576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6001600160a01b0382166000908152600a602052604090205460ff16614ab6576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b6001600160a01b0382166000908152600a6020526040902060020154600b54614ae090829061501d565b9392505050565b600554600090610100900460ff1615614b35576040805162461bcd60e51b815260206004820152600b60248201526a4552525f5245454e54525960a81b604482015290519081900360640190fd5b6001600160a01b0382166000908152600a602052604090205460ff16614b92576040805162461bcd60e51b815260206004820152600d60248201526c11549497d393d517d093d55391609a1b604482015290519081900360640190fd5b506001600160a01b03166000908152600a602052604090206003015490565b600080614bbe858861501d565b90506000614bcc8786614d0d565b90506000614bda888361501d565b90506000614be88285615125565b9050614bfc81670de0b6b3a7640000614d0d565b9050614c10670de0b6b3a764000087614d0d565b9450614c25614c1f8c83614c44565b8661501d565b9b9a5050505050505050505050565b600654600160a01b900460ff1690565b6000828202831580614c5e575082848281614c5b57fe5b04145b614ca2576040805162461bcd60e51b815260206004820152601060248201526f4552525f4d554c5f4f564552464c4f5760801b604482015290519081900360640190fd5b6706f05b59d3b20000810181811015614cf5576040805162461bcd60e51b815260206004820152601060248201526f4552525f4d554c5f4f564552464c4f5760801b604482015290519081900360640190fd5b6000670de0b6b3a7640000825b049695505050505050565b6000806000614d1c8585615233565b915091508015614d67576040805162461bcd60e51b81526020600482015260116024820152704552525f5355425f554e444552464c4f5760781b604482015290519081900360640190fd5b509392505050565b614d798282615258565b5050565b614d8681615263565b50565b614d798282615333565b6040805163a9059cbb60e01b81526001600160a01b03848116600483015260248201849052915160009286169163a9059cbb91604480830192602092919082900301818787803b158015614de657600080fd5b505af1158015614dfa573d6000803e3d6000fd5b505050506040513d6020811015614e1057600080fd5b5051905080614e58576040805162461bcd60e51b815260206004820152600f60248201526e4552525f45524332305f46414c534560881b604482015290519081900360640190fd5b50505050565b6001600160a01b038316600090815260208190526040902054811115614ec2576040805162461bcd60e51b815260206004820152601460248201527311549497d25394d551919250d251539517d0905360621b604482015290519081900360640190fd5b6001600160a01b038316600090815260208190526040902054614ee59082614d0d565b6001600160a01b038085166000908152602081905260408082209390935590841681522054614f149082614f6e565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600082820183811015614ae0576040805162461bcd60e51b815260206004820152601060248201526f4552525f4144445f4f564552464c4f5760801b604482015290519081900360640190fd5b604080516323b872dd60e01b81526001600160a01b0384811660048301523060248301526044820184905291516000928616916323b872dd91606480830192602092919082900301818787803b158015614de657600080fd5b614d868161533e565b600081615060576040805162461bcd60e51b815260206004820152600c60248201526b4552525f4449565f5a45524f60a01b604482015290519081900360640190fd5b670de0b6b3a764000083028315806150885750670de0b6b3a764000084828161508557fe5b04145b6150cc576040805162461bcd60e51b815260206004820152601060248201526f11549497d1125597d25395115493905360821b604482015290519081900360640190fd5b6002830481018181101561511a576040805162461bcd60e51b815260206004820152601060248201526f11549497d1125597d25395115493905360821b604482015290519081900360640190fd5b6000848281614d0257fe5b60006001831015615175576040805162461bcd60e51b81526020600482015260156024820152744552525f42504f575f424153455f544f4f5f4c4f5760581b604482015290519081900360640190fd5b671bc16d674ec7ffff8311156151cb576040805162461bcd60e51b815260206004820152601660248201527508aa4a4be84a09eaebe8482a68abea89e9ebe90928e960531b604482015290519081900360640190fd5b60006151d6836153b3565b905060006151e48483614d0d565b905060006151fa866151f5856153ce565b6153dc565b90508161520b57925061114a915050565b600061521c87846305f5e100615433565b90506152288282614c44565b979650505050505050565b6000808284106152495750508082036000615251565b505081810360015b9250929050565b614d79823083614e5e565b306000908152602081905260409020548111156152be576040805162461bcd60e51b815260206004820152601460248201527311549497d25394d551919250d251539517d0905360621b604482015290519081900360640190fd5b306000908152602081905260409020546152d89082614d0d565b306000908152602081905260409020556002546152f59082614d0d565b60025560408051828152905160009130917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a350565b614d79308383614e5e565b306000908152602081905260409020546153589082614f6e565b306000908152602081905260409020556002546153759082614f6e565b60025560408051828152905130916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a350565b6000670de0b6b3a76400006153c7836153ce565b0292915050565b670de0b6b3a7640000900490565b600080600283066153f557670de0b6b3a76400006153f7565b835b90506002830492505b8215614ae0576154108485614c44565b93506002830615615428576154258185614c44565b90505b600283049250615400565b600082818061544a87670de0b6b3a7640000615233565b9092509050670de0b6b3a764000080600060015b888410615502576000670de0b6b3a7640000820290506000806154928a61548d85670de0b6b3a7640000614d0d565b615233565b915091506154a4876137a2848c614c44565b96506154b0878461501d565b9650866154bf57505050615502565b87156154c9579315935b80156154d3579315935b84156154ea576154e38688614d0d565b95506154f7565b6154f48688614f6e565b95505b50505060010161545e565b5090999850505050505050505056fe8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925a265627a7a72315820a7cc59bf2e6a7f7e1b4b74a8cd9ed2d0ffd628b1ea9f42f7d1fac030c01dad6764736f6c634300050c0032"

// DeployBPool deploys a new Ethereum contract, binding an instance of BPool to it.
func DeployBPool(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BPool, error) {
	parsed, err := abi.JSON(strings.NewReader(BPoolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BPoolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BPool{BPoolCaller: BPoolCaller{contract: contract}, BPoolTransactor: BPoolTransactor{contract: contract}, BPoolFilterer: BPoolFilterer{contract: contract}}, nil
}

// BPool is an auto generated Go binding around an Ethereum contract.
type BPool struct {
	BPoolCaller     // Read-only binding to the contract
	BPoolTransactor // Write-only binding to the contract
	BPoolFilterer   // Log filterer for contract events
}

// BPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type BPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BPoolSession struct {
	Contract     *BPool            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BPoolCallerSession struct {
	Contract *BPoolCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BPoolTransactorSession struct {
	Contract     *BPoolTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type BPoolRaw struct {
	Contract *BPool // Generic contract binding to access the raw methods on
}

// BPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BPoolCallerRaw struct {
	Contract *BPoolCaller // Generic read-only contract binding to access the raw methods on
}

// BPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BPoolTransactorRaw struct {
	Contract *BPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBPool creates a new instance of BPool, bound to a specific deployed contract.
func NewBPool(address common.Address, backend bind.ContractBackend) (*BPool, error) {
	contract, err := bindBPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BPool{BPoolCaller: BPoolCaller{contract: contract}, BPoolTransactor: BPoolTransactor{contract: contract}, BPoolFilterer: BPoolFilterer{contract: contract}}, nil
}

// NewBPoolCaller creates a new read-only instance of BPool, bound to a specific deployed contract.
func NewBPoolCaller(address common.Address, caller bind.ContractCaller) (*BPoolCaller, error) {
	contract, err := bindBPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BPoolCaller{contract: contract}, nil
}

// NewBPoolTransactor creates a new write-only instance of BPool, bound to a specific deployed contract.
func NewBPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*BPoolTransactor, error) {
	contract, err := bindBPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BPoolTransactor{contract: contract}, nil
}

// NewBPoolFilterer creates a new log filterer instance of BPool, bound to a specific deployed contract.
func NewBPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*BPoolFilterer, error) {
	contract, err := bindBPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BPoolFilterer{contract: contract}, nil
}

// bindBPool binds a generic wrapper to an already deployed contract.
func bindBPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BPool *BPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BPool.Contract.BPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BPool *BPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BPool.Contract.BPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BPool *BPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BPool.Contract.BPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BPool *BPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BPool *BPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BPool *BPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BPool.Contract.contract.Transact(opts, method, params...)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BPool *BPoolCaller) BONE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "BONE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BPool *BPoolSession) BONE() (*big.Int, error) {
	return _BPool.Contract.BONE(&_BPool.CallOpts)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BPool *BPoolCallerSession) BONE() (*big.Int, error) {
	return _BPool.Contract.BONE(&_BPool.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BPool *BPoolCaller) BPOWPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "BPOW_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BPool *BPoolSession) BPOWPRECISION() (*big.Int, error) {
	return _BPool.Contract.BPOWPRECISION(&_BPool.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BPool *BPoolCallerSession) BPOWPRECISION() (*big.Int, error) {
	return _BPool.Contract.BPOWPRECISION(&_BPool.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BPool *BPoolCaller) EXITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "EXIT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BPool *BPoolSession) EXITFEE() (*big.Int, error) {
	return _BPool.Contract.EXITFEE(&_BPool.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BPool *BPoolCallerSession) EXITFEE() (*big.Int, error) {
	return _BPool.Contract.EXITFEE(&_BPool.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BPool *BPoolCaller) INITPOOLSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "INIT_POOL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BPool *BPoolSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _BPool.Contract.INITPOOLSUPPLY(&_BPool.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BPool *BPoolCallerSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _BPool.Contract.INITPOOLSUPPLY(&_BPool.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BPool *BPoolCaller) MAXBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "MAX_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BPool *BPoolSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _BPool.Contract.MAXBOUNDTOKENS(&_BPool.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BPool *BPoolCallerSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _BPool.Contract.MAXBOUNDTOKENS(&_BPool.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BPool *BPoolCaller) MAXBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "MAX_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BPool *BPoolSession) MAXBPOWBASE() (*big.Int, error) {
	return _BPool.Contract.MAXBPOWBASE(&_BPool.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BPool *BPoolCallerSession) MAXBPOWBASE() (*big.Int, error) {
	return _BPool.Contract.MAXBPOWBASE(&_BPool.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BPool *BPoolCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BPool *BPoolSession) MAXFEE() (*big.Int, error) {
	return _BPool.Contract.MAXFEE(&_BPool.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BPool *BPoolCallerSession) MAXFEE() (*big.Int, error) {
	return _BPool.Contract.MAXFEE(&_BPool.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BPool *BPoolCaller) MAXINRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "MAX_IN_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BPool *BPoolSession) MAXINRATIO() (*big.Int, error) {
	return _BPool.Contract.MAXINRATIO(&_BPool.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BPool *BPoolCallerSession) MAXINRATIO() (*big.Int, error) {
	return _BPool.Contract.MAXINRATIO(&_BPool.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BPool *BPoolCaller) MAXOUTRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "MAX_OUT_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BPool *BPoolSession) MAXOUTRATIO() (*big.Int, error) {
	return _BPool.Contract.MAXOUTRATIO(&_BPool.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BPool *BPoolCallerSession) MAXOUTRATIO() (*big.Int, error) {
	return _BPool.Contract.MAXOUTRATIO(&_BPool.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BPool *BPoolCaller) MAXTOTALWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "MAX_TOTAL_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BPool *BPoolSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _BPool.Contract.MAXTOTALWEIGHT(&_BPool.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BPool *BPoolCallerSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _BPool.Contract.MAXTOTALWEIGHT(&_BPool.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BPool *BPoolCaller) MAXWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "MAX_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BPool *BPoolSession) MAXWEIGHT() (*big.Int, error) {
	return _BPool.Contract.MAXWEIGHT(&_BPool.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BPool *BPoolCallerSession) MAXWEIGHT() (*big.Int, error) {
	return _BPool.Contract.MAXWEIGHT(&_BPool.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BPool *BPoolCaller) MINBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "MIN_BALANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BPool *BPoolSession) MINBALANCE() (*big.Int, error) {
	return _BPool.Contract.MINBALANCE(&_BPool.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BPool *BPoolCallerSession) MINBALANCE() (*big.Int, error) {
	return _BPool.Contract.MINBALANCE(&_BPool.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BPool *BPoolCaller) MINBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "MIN_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BPool *BPoolSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _BPool.Contract.MINBOUNDTOKENS(&_BPool.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BPool *BPoolCallerSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _BPool.Contract.MINBOUNDTOKENS(&_BPool.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BPool *BPoolCaller) MINBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "MIN_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BPool *BPoolSession) MINBPOWBASE() (*big.Int, error) {
	return _BPool.Contract.MINBPOWBASE(&_BPool.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BPool *BPoolCallerSession) MINBPOWBASE() (*big.Int, error) {
	return _BPool.Contract.MINBPOWBASE(&_BPool.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BPool *BPoolCaller) MINFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "MIN_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BPool *BPoolSession) MINFEE() (*big.Int, error) {
	return _BPool.Contract.MINFEE(&_BPool.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BPool *BPoolCallerSession) MINFEE() (*big.Int, error) {
	return _BPool.Contract.MINFEE(&_BPool.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BPool *BPoolCaller) MINWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "MIN_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BPool *BPoolSession) MINWEIGHT() (*big.Int, error) {
	return _BPool.Contract.MINWEIGHT(&_BPool.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BPool *BPoolCallerSession) MINWEIGHT() (*big.Int, error) {
	return _BPool.Contract.MINWEIGHT(&_BPool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_BPool *BPoolCaller) Allowance(opts *bind.CallOpts, src common.Address, dst common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "allowance", src, dst)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_BPool *BPoolSession) Allowance(src common.Address, dst common.Address) (*big.Int, error) {
	return _BPool.Contract.Allowance(&_BPool.CallOpts, src, dst)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_BPool *BPoolCallerSession) Allowance(src common.Address, dst common.Address) (*big.Int, error) {
	return _BPool.Contract.Allowance(&_BPool.CallOpts, src, dst)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_BPool *BPoolCaller) BalanceOf(opts *bind.CallOpts, whom common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "balanceOf", whom)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_BPool *BPoolSession) BalanceOf(whom common.Address) (*big.Int, error) {
	return _BPool.Contract.BalanceOf(&_BPool.CallOpts, whom)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_BPool *BPoolCallerSession) BalanceOf(whom common.Address) (*big.Int, error) {
	return _BPool.Contract.BalanceOf(&_BPool.CallOpts, whom)
}

// CalcInGivenOut is a free data retrieval call binding the contract method 0xf8d6aed4.
//
// Solidity: function calcInGivenOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BPool *BPoolCaller) CalcInGivenOut(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "calcInGivenOut", tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountOut, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcInGivenOut is a free data retrieval call binding the contract method 0xf8d6aed4.
//
// Solidity: function calcInGivenOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BPool *BPoolSession) CalcInGivenOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BPool.Contract.CalcInGivenOut(&_BPool.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountOut, swapFee)
}

// CalcInGivenOut is a free data retrieval call binding the contract method 0xf8d6aed4.
//
// Solidity: function calcInGivenOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BPool *BPoolCallerSession) CalcInGivenOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BPool.Contract.CalcInGivenOut(&_BPool.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountOut, swapFee)
}

// CalcOutGivenIn is a free data retrieval call binding the contract method 0xba9530a6.
//
// Solidity: function calcOutGivenIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BPool *BPoolCaller) CalcOutGivenIn(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "calcOutGivenIn", tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountIn, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcOutGivenIn is a free data retrieval call binding the contract method 0xba9530a6.
//
// Solidity: function calcOutGivenIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BPool *BPoolSession) CalcOutGivenIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BPool.Contract.CalcOutGivenIn(&_BPool.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountIn, swapFee)
}

// CalcOutGivenIn is a free data retrieval call binding the contract method 0xba9530a6.
//
// Solidity: function calcOutGivenIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BPool *BPoolCallerSession) CalcOutGivenIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BPool.Contract.CalcOutGivenIn(&_BPool.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountIn, swapFee)
}

// CalcPoolInGivenSingleOut is a free data retrieval call binding the contract method 0x82f652ad.
//
// Solidity: function calcPoolInGivenSingleOut(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 poolAmountIn)
func (_BPool *BPoolCaller) CalcPoolInGivenSingleOut(opts *bind.CallOpts, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "calcPoolInGivenSingleOut", tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, tokenAmountOut, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcPoolInGivenSingleOut is a free data retrieval call binding the contract method 0x82f652ad.
//
// Solidity: function calcPoolInGivenSingleOut(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 poolAmountIn)
func (_BPool *BPoolSession) CalcPoolInGivenSingleOut(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BPool.Contract.CalcPoolInGivenSingleOut(&_BPool.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, tokenAmountOut, swapFee)
}

// CalcPoolInGivenSingleOut is a free data retrieval call binding the contract method 0x82f652ad.
//
// Solidity: function calcPoolInGivenSingleOut(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 poolAmountIn)
func (_BPool *BPoolCallerSession) CalcPoolInGivenSingleOut(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BPool.Contract.CalcPoolInGivenSingleOut(&_BPool.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, tokenAmountOut, swapFee)
}

// CalcPoolOutGivenSingleIn is a free data retrieval call binding the contract method 0x8656b653.
//
// Solidity: function calcPoolOutGivenSingleIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 poolAmountOut)
func (_BPool *BPoolCaller) CalcPoolOutGivenSingleIn(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "calcPoolOutGivenSingleIn", tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, tokenAmountIn, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcPoolOutGivenSingleIn is a free data retrieval call binding the contract method 0x8656b653.
//
// Solidity: function calcPoolOutGivenSingleIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 poolAmountOut)
func (_BPool *BPoolSession) CalcPoolOutGivenSingleIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BPool.Contract.CalcPoolOutGivenSingleIn(&_BPool.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, tokenAmountIn, swapFee)
}

// CalcPoolOutGivenSingleIn is a free data retrieval call binding the contract method 0x8656b653.
//
// Solidity: function calcPoolOutGivenSingleIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 poolAmountOut)
func (_BPool *BPoolCallerSession) CalcPoolOutGivenSingleIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BPool.Contract.CalcPoolOutGivenSingleIn(&_BPool.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, tokenAmountIn, swapFee)
}

// CalcSingleInGivenPoolOut is a free data retrieval call binding the contract method 0x5c1bbaf7.
//
// Solidity: function calcSingleInGivenPoolOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BPool *BPoolCaller) CalcSingleInGivenPoolOut(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "calcSingleInGivenPoolOut", tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, poolAmountOut, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcSingleInGivenPoolOut is a free data retrieval call binding the contract method 0x5c1bbaf7.
//
// Solidity: function calcSingleInGivenPoolOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BPool *BPoolSession) CalcSingleInGivenPoolOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BPool.Contract.CalcSingleInGivenPoolOut(&_BPool.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, poolAmountOut, swapFee)
}

// CalcSingleInGivenPoolOut is a free data retrieval call binding the contract method 0x5c1bbaf7.
//
// Solidity: function calcSingleInGivenPoolOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BPool *BPoolCallerSession) CalcSingleInGivenPoolOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BPool.Contract.CalcSingleInGivenPoolOut(&_BPool.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, poolAmountOut, swapFee)
}

// CalcSingleOutGivenPoolIn is a free data retrieval call binding the contract method 0x89298012.
//
// Solidity: function calcSingleOutGivenPoolIn(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BPool *BPoolCaller) CalcSingleOutGivenPoolIn(opts *bind.CallOpts, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "calcSingleOutGivenPoolIn", tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, poolAmountIn, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcSingleOutGivenPoolIn is a free data retrieval call binding the contract method 0x89298012.
//
// Solidity: function calcSingleOutGivenPoolIn(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BPool *BPoolSession) CalcSingleOutGivenPoolIn(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BPool.Contract.CalcSingleOutGivenPoolIn(&_BPool.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, poolAmountIn, swapFee)
}

// CalcSingleOutGivenPoolIn is a free data retrieval call binding the contract method 0x89298012.
//
// Solidity: function calcSingleOutGivenPoolIn(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BPool *BPoolCallerSession) CalcSingleOutGivenPoolIn(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BPool.Contract.CalcSingleOutGivenPoolIn(&_BPool.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, poolAmountIn, swapFee)
}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_BPool *BPoolCaller) CalcSpotPrice(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "calcSpotPrice", tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_BPool *BPoolSession) CalcSpotPrice(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BPool.Contract.CalcSpotPrice(&_BPool.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)
}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_BPool *BPoolCallerSession) CalcSpotPrice(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BPool.Contract.CalcSpotPrice(&_BPool.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BPool *BPoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BPool *BPoolSession) Decimals() (uint8, error) {
	return _BPool.Contract.Decimals(&_BPool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BPool *BPoolCallerSession) Decimals() (uint8, error) {
	return _BPool.Contract.Decimals(&_BPool.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_BPool *BPoolCaller) GetBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "getBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_BPool *BPoolSession) GetBalance(token common.Address) (*big.Int, error) {
	return _BPool.Contract.GetBalance(&_BPool.CallOpts, token)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_BPool *BPoolCallerSession) GetBalance(token common.Address) (*big.Int, error) {
	return _BPool.Contract.GetBalance(&_BPool.CallOpts, token)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BPool *BPoolCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "getColor")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BPool *BPoolSession) GetColor() ([32]byte, error) {
	return _BPool.Contract.GetColor(&_BPool.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BPool *BPoolCallerSession) GetColor() ([32]byte, error) {
	return _BPool.Contract.GetColor(&_BPool.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_BPool *BPoolCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "getController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_BPool *BPoolSession) GetController() (common.Address, error) {
	return _BPool.Contract.GetController(&_BPool.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_BPool *BPoolCallerSession) GetController() (common.Address, error) {
	return _BPool.Contract.GetController(&_BPool.CallOpts)
}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_BPool *BPoolCaller) GetCurrentTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "getCurrentTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_BPool *BPoolSession) GetCurrentTokens() ([]common.Address, error) {
	return _BPool.Contract.GetCurrentTokens(&_BPool.CallOpts)
}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_BPool *BPoolCallerSession) GetCurrentTokens() ([]common.Address, error) {
	return _BPool.Contract.GetCurrentTokens(&_BPool.CallOpts)
}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_BPool *BPoolCaller) GetDenormalizedWeight(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "getDenormalizedWeight", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_BPool *BPoolSession) GetDenormalizedWeight(token common.Address) (*big.Int, error) {
	return _BPool.Contract.GetDenormalizedWeight(&_BPool.CallOpts, token)
}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_BPool *BPoolCallerSession) GetDenormalizedWeight(token common.Address) (*big.Int, error) {
	return _BPool.Contract.GetDenormalizedWeight(&_BPool.CallOpts, token)
}

// GetFinalTokens is a free data retrieval call binding the contract method 0xbe3bbd2e.
//
// Solidity: function getFinalTokens() view returns(address[] tokens)
func (_BPool *BPoolCaller) GetFinalTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "getFinalTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFinalTokens is a free data retrieval call binding the contract method 0xbe3bbd2e.
//
// Solidity: function getFinalTokens() view returns(address[] tokens)
func (_BPool *BPoolSession) GetFinalTokens() ([]common.Address, error) {
	return _BPool.Contract.GetFinalTokens(&_BPool.CallOpts)
}

// GetFinalTokens is a free data retrieval call binding the contract method 0xbe3bbd2e.
//
// Solidity: function getFinalTokens() view returns(address[] tokens)
func (_BPool *BPoolCallerSession) GetFinalTokens() ([]common.Address, error) {
	return _BPool.Contract.GetFinalTokens(&_BPool.CallOpts)
}

// GetNormalizedWeight is a free data retrieval call binding the contract method 0xf1b8a9b7.
//
// Solidity: function getNormalizedWeight(address token) view returns(uint256)
func (_BPool *BPoolCaller) GetNormalizedWeight(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "getNormalizedWeight", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNormalizedWeight is a free data retrieval call binding the contract method 0xf1b8a9b7.
//
// Solidity: function getNormalizedWeight(address token) view returns(uint256)
func (_BPool *BPoolSession) GetNormalizedWeight(token common.Address) (*big.Int, error) {
	return _BPool.Contract.GetNormalizedWeight(&_BPool.CallOpts, token)
}

// GetNormalizedWeight is a free data retrieval call binding the contract method 0xf1b8a9b7.
//
// Solidity: function getNormalizedWeight(address token) view returns(uint256)
func (_BPool *BPoolCallerSession) GetNormalizedWeight(token common.Address) (*big.Int, error) {
	return _BPool.Contract.GetNormalizedWeight(&_BPool.CallOpts, token)
}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_BPool *BPoolCaller) GetNumTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "getNumTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_BPool *BPoolSession) GetNumTokens() (*big.Int, error) {
	return _BPool.Contract.GetNumTokens(&_BPool.CallOpts)
}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_BPool *BPoolCallerSession) GetNumTokens() (*big.Int, error) {
	return _BPool.Contract.GetNumTokens(&_BPool.CallOpts)
}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_BPool *BPoolCaller) GetSpotPrice(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "getSpotPrice", tokenIn, tokenOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_BPool *BPoolSession) GetSpotPrice(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _BPool.Contract.GetSpotPrice(&_BPool.CallOpts, tokenIn, tokenOut)
}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_BPool *BPoolCallerSession) GetSpotPrice(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _BPool.Contract.GetSpotPrice(&_BPool.CallOpts, tokenIn, tokenOut)
}

// GetSpotPriceSansFee is a free data retrieval call binding the contract method 0x1446a7ff.
//
// Solidity: function getSpotPriceSansFee(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_BPool *BPoolCaller) GetSpotPriceSansFee(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "getSpotPriceSansFee", tokenIn, tokenOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSpotPriceSansFee is a free data retrieval call binding the contract method 0x1446a7ff.
//
// Solidity: function getSpotPriceSansFee(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_BPool *BPoolSession) GetSpotPriceSansFee(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _BPool.Contract.GetSpotPriceSansFee(&_BPool.CallOpts, tokenIn, tokenOut)
}

// GetSpotPriceSansFee is a free data retrieval call binding the contract method 0x1446a7ff.
//
// Solidity: function getSpotPriceSansFee(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_BPool *BPoolCallerSession) GetSpotPriceSansFee(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _BPool.Contract.GetSpotPriceSansFee(&_BPool.CallOpts, tokenIn, tokenOut)
}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_BPool *BPoolCaller) GetSwapFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "getSwapFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_BPool *BPoolSession) GetSwapFee() (*big.Int, error) {
	return _BPool.Contract.GetSwapFee(&_BPool.CallOpts)
}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_BPool *BPoolCallerSession) GetSwapFee() (*big.Int, error) {
	return _BPool.Contract.GetSwapFee(&_BPool.CallOpts)
}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_BPool *BPoolCaller) GetTotalDenormalizedWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "getTotalDenormalizedWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_BPool *BPoolSession) GetTotalDenormalizedWeight() (*big.Int, error) {
	return _BPool.Contract.GetTotalDenormalizedWeight(&_BPool.CallOpts)
}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_BPool *BPoolCallerSession) GetTotalDenormalizedWeight() (*big.Int, error) {
	return _BPool.Contract.GetTotalDenormalizedWeight(&_BPool.CallOpts)
}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_BPool *BPoolCaller) IsBound(opts *bind.CallOpts, t common.Address) (bool, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "isBound", t)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_BPool *BPoolSession) IsBound(t common.Address) (bool, error) {
	return _BPool.Contract.IsBound(&_BPool.CallOpts, t)
}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_BPool *BPoolCallerSession) IsBound(t common.Address) (bool, error) {
	return _BPool.Contract.IsBound(&_BPool.CallOpts, t)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_BPool *BPoolCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "isFinalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_BPool *BPoolSession) IsFinalized() (bool, error) {
	return _BPool.Contract.IsFinalized(&_BPool.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_BPool *BPoolCallerSession) IsFinalized() (bool, error) {
	return _BPool.Contract.IsFinalized(&_BPool.CallOpts)
}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_BPool *BPoolCaller) IsPublicSwap(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "isPublicSwap")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_BPool *BPoolSession) IsPublicSwap() (bool, error) {
	return _BPool.Contract.IsPublicSwap(&_BPool.CallOpts)
}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_BPool *BPoolCallerSession) IsPublicSwap() (bool, error) {
	return _BPool.Contract.IsPublicSwap(&_BPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BPool *BPoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BPool *BPoolSession) Name() (string, error) {
	return _BPool.Contract.Name(&_BPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BPool *BPoolCallerSession) Name() (string, error) {
	return _BPool.Contract.Name(&_BPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BPool *BPoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BPool *BPoolSession) Symbol() (string, error) {
	return _BPool.Contract.Symbol(&_BPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BPool *BPoolCallerSession) Symbol() (string, error) {
	return _BPool.Contract.Symbol(&_BPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BPool *BPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BPool *BPoolSession) TotalSupply() (*big.Int, error) {
	return _BPool.Contract.TotalSupply(&_BPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BPool *BPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _BPool.Contract.TotalSupply(&_BPool.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_BPool *BPoolTransactor) Approve(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "approve", dst, amt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_BPool *BPoolSession) Approve(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.Approve(&_BPool.TransactOpts, dst, amt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_BPool *BPoolTransactorSession) Approve(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.Approve(&_BPool.TransactOpts, dst, amt)
}

// Bind is a paid mutator transaction binding the contract method 0xe4e1e538.
//
// Solidity: function bind(address token, uint256 balance, uint256 denorm) returns()
func (_BPool *BPoolTransactor) Bind(opts *bind.TransactOpts, token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "bind", token, balance, denorm)
}

// Bind is a paid mutator transaction binding the contract method 0xe4e1e538.
//
// Solidity: function bind(address token, uint256 balance, uint256 denorm) returns()
func (_BPool *BPoolSession) Bind(token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.Bind(&_BPool.TransactOpts, token, balance, denorm)
}

// Bind is a paid mutator transaction binding the contract method 0xe4e1e538.
//
// Solidity: function bind(address token, uint256 balance, uint256 denorm) returns()
func (_BPool *BPoolTransactorSession) Bind(token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.Bind(&_BPool.TransactOpts, token, balance, denorm)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_BPool *BPoolTransactor) DecreaseApproval(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "decreaseApproval", dst, amt)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_BPool *BPoolSession) DecreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.DecreaseApproval(&_BPool.TransactOpts, dst, amt)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_BPool *BPoolTransactorSession) DecreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.DecreaseApproval(&_BPool.TransactOpts, dst, amt)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_BPool *BPoolTransactor) ExitPool(opts *bind.TransactOpts, poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "exitPool", poolAmountIn, minAmountsOut)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_BPool *BPoolSession) ExitPool(poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _BPool.Contract.ExitPool(&_BPool.TransactOpts, poolAmountIn, minAmountsOut)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_BPool *BPoolTransactorSession) ExitPool(poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _BPool.Contract.ExitPool(&_BPool.TransactOpts, poolAmountIn, minAmountsOut)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256 poolAmountIn)
func (_BPool *BPoolTransactor) ExitswapExternAmountOut(opts *bind.TransactOpts, tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "exitswapExternAmountOut", tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256 poolAmountIn)
func (_BPool *BPoolSession) ExitswapExternAmountOut(tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.ExitswapExternAmountOut(&_BPool.TransactOpts, tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256 poolAmountIn)
func (_BPool *BPoolTransactorSession) ExitswapExternAmountOut(tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.ExitswapExternAmountOut(&_BPool.TransactOpts, tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256 tokenAmountOut)
func (_BPool *BPoolTransactor) ExitswapPoolAmountIn(opts *bind.TransactOpts, tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "exitswapPoolAmountIn", tokenOut, poolAmountIn, minAmountOut)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256 tokenAmountOut)
func (_BPool *BPoolSession) ExitswapPoolAmountIn(tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.ExitswapPoolAmountIn(&_BPool.TransactOpts, tokenOut, poolAmountIn, minAmountOut)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256 tokenAmountOut)
func (_BPool *BPoolTransactorSession) ExitswapPoolAmountIn(tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.ExitswapPoolAmountIn(&_BPool.TransactOpts, tokenOut, poolAmountIn, minAmountOut)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_BPool *BPoolTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_BPool *BPoolSession) Finalize() (*types.Transaction, error) {
	return _BPool.Contract.Finalize(&_BPool.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_BPool *BPoolTransactorSession) Finalize() (*types.Transaction, error) {
	return _BPool.Contract.Finalize(&_BPool.TransactOpts)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_BPool *BPoolTransactor) Gulp(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "gulp", token)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_BPool *BPoolSession) Gulp(token common.Address) (*types.Transaction, error) {
	return _BPool.Contract.Gulp(&_BPool.TransactOpts, token)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_BPool *BPoolTransactorSession) Gulp(token common.Address) (*types.Transaction, error) {
	return _BPool.Contract.Gulp(&_BPool.TransactOpts, token)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_BPool *BPoolTransactor) IncreaseApproval(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "increaseApproval", dst, amt)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_BPool *BPoolSession) IncreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.IncreaseApproval(&_BPool.TransactOpts, dst, amt)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_BPool *BPoolTransactorSession) IncreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.IncreaseApproval(&_BPool.TransactOpts, dst, amt)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_BPool *BPoolTransactor) JoinPool(opts *bind.TransactOpts, poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "joinPool", poolAmountOut, maxAmountsIn)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_BPool *BPoolSession) JoinPool(poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _BPool.Contract.JoinPool(&_BPool.TransactOpts, poolAmountOut, maxAmountsIn)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_BPool *BPoolTransactorSession) JoinPool(poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _BPool.Contract.JoinPool(&_BPool.TransactOpts, poolAmountOut, maxAmountsIn)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256 poolAmountOut)
func (_BPool *BPoolTransactor) JoinswapExternAmountIn(opts *bind.TransactOpts, tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "joinswapExternAmountIn", tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256 poolAmountOut)
func (_BPool *BPoolSession) JoinswapExternAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.JoinswapExternAmountIn(&_BPool.TransactOpts, tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256 poolAmountOut)
func (_BPool *BPoolTransactorSession) JoinswapExternAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.JoinswapExternAmountIn(&_BPool.TransactOpts, tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256 tokenAmountIn)
func (_BPool *BPoolTransactor) JoinswapPoolAmountOut(opts *bind.TransactOpts, tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "joinswapPoolAmountOut", tokenIn, poolAmountOut, maxAmountIn)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256 tokenAmountIn)
func (_BPool *BPoolSession) JoinswapPoolAmountOut(tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.JoinswapPoolAmountOut(&_BPool.TransactOpts, tokenIn, poolAmountOut, maxAmountIn)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256 tokenAmountIn)
func (_BPool *BPoolTransactorSession) JoinswapPoolAmountOut(tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.JoinswapPoolAmountOut(&_BPool.TransactOpts, tokenIn, poolAmountOut, maxAmountIn)
}

// Rebind is a paid mutator transaction binding the contract method 0x3fdddaa2.
//
// Solidity: function rebind(address token, uint256 balance, uint256 denorm) returns()
func (_BPool *BPoolTransactor) Rebind(opts *bind.TransactOpts, token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "rebind", token, balance, denorm)
}

// Rebind is a paid mutator transaction binding the contract method 0x3fdddaa2.
//
// Solidity: function rebind(address token, uint256 balance, uint256 denorm) returns()
func (_BPool *BPoolSession) Rebind(token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.Rebind(&_BPool.TransactOpts, token, balance, denorm)
}

// Rebind is a paid mutator transaction binding the contract method 0x3fdddaa2.
//
// Solidity: function rebind(address token, uint256 balance, uint256 denorm) returns()
func (_BPool *BPoolTransactorSession) Rebind(token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.Rebind(&_BPool.TransactOpts, token, balance, denorm)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address manager) returns()
func (_BPool *BPoolTransactor) SetController(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "setController", manager)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address manager) returns()
func (_BPool *BPoolSession) SetController(manager common.Address) (*types.Transaction, error) {
	return _BPool.Contract.SetController(&_BPool.TransactOpts, manager)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address manager) returns()
func (_BPool *BPoolTransactorSession) SetController(manager common.Address) (*types.Transaction, error) {
	return _BPool.Contract.SetController(&_BPool.TransactOpts, manager)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool public_) returns()
func (_BPool *BPoolTransactor) SetPublicSwap(opts *bind.TransactOpts, public_ bool) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "setPublicSwap", public_)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool public_) returns()
func (_BPool *BPoolSession) SetPublicSwap(public_ bool) (*types.Transaction, error) {
	return _BPool.Contract.SetPublicSwap(&_BPool.TransactOpts, public_)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool public_) returns()
func (_BPool *BPoolTransactorSession) SetPublicSwap(public_ bool) (*types.Transaction, error) {
	return _BPool.Contract.SetPublicSwap(&_BPool.TransactOpts, public_)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_BPool *BPoolTransactor) SetSwapFee(opts *bind.TransactOpts, swapFee *big.Int) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "setSwapFee", swapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_BPool *BPoolSession) SetSwapFee(swapFee *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.SetSwapFee(&_BPool.TransactOpts, swapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_BPool *BPoolTransactorSession) SetSwapFee(swapFee *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.SetSwapFee(&_BPool.TransactOpts, swapFee)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256 tokenAmountOut, uint256 spotPriceAfter)
func (_BPool *BPoolTransactor) SwapExactAmountIn(opts *bind.TransactOpts, tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "swapExactAmountIn", tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256 tokenAmountOut, uint256 spotPriceAfter)
func (_BPool *BPoolSession) SwapExactAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.SwapExactAmountIn(&_BPool.TransactOpts, tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256 tokenAmountOut, uint256 spotPriceAfter)
func (_BPool *BPoolTransactorSession) SwapExactAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.SwapExactAmountIn(&_BPool.TransactOpts, tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256 tokenAmountIn, uint256 spotPriceAfter)
func (_BPool *BPoolTransactor) SwapExactAmountOut(opts *bind.TransactOpts, tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "swapExactAmountOut", tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256 tokenAmountIn, uint256 spotPriceAfter)
func (_BPool *BPoolSession) SwapExactAmountOut(tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.SwapExactAmountOut(&_BPool.TransactOpts, tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256 tokenAmountIn, uint256 spotPriceAfter)
func (_BPool *BPoolTransactorSession) SwapExactAmountOut(tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.SwapExactAmountOut(&_BPool.TransactOpts, tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_BPool *BPoolTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "transfer", dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_BPool *BPoolSession) Transfer(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.Transfer(&_BPool.TransactOpts, dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_BPool *BPoolTransactorSession) Transfer(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.Transfer(&_BPool.TransactOpts, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_BPool *BPoolTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "transferFrom", src, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_BPool *BPoolSession) TransferFrom(src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.TransferFrom(&_BPool.TransactOpts, src, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_BPool *BPoolTransactorSession) TransferFrom(src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BPool.Contract.TransferFrom(&_BPool.TransactOpts, src, dst, amt)
}

// Unbind is a paid mutator transaction binding the contract method 0xcf5e7bd3.
//
// Solidity: function unbind(address token) returns()
func (_BPool *BPoolTransactor) Unbind(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BPool.contract.Transact(opts, "unbind", token)
}

// Unbind is a paid mutator transaction binding the contract method 0xcf5e7bd3.
//
// Solidity: function unbind(address token) returns()
func (_BPool *BPoolSession) Unbind(token common.Address) (*types.Transaction, error) {
	return _BPool.Contract.Unbind(&_BPool.TransactOpts, token)
}

// Unbind is a paid mutator transaction binding the contract method 0xcf5e7bd3.
//
// Solidity: function unbind(address token) returns()
func (_BPool *BPoolTransactorSession) Unbind(token common.Address) (*types.Transaction, error) {
	return _BPool.Contract.Unbind(&_BPool.TransactOpts, token)
}

// BPoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BPool contract.
type BPoolApprovalIterator struct {
	Event *BPoolApproval // Event containing the contract specifics and raw log

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
func (it *BPoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BPoolApproval)
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
		it.Event = new(BPoolApproval)
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
func (it *BPoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BPoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BPoolApproval represents a Approval event raised by the BPool contract.
type BPoolApproval struct {
	Src common.Address
	Dst common.Address
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_BPool *BPoolFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*BPoolApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BPool.contract.FilterLogs(opts, "Approval", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &BPoolApprovalIterator{contract: _BPool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_BPool *BPoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BPoolApproval, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BPool.contract.WatchLogs(opts, "Approval", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BPoolApproval)
				if err := _BPool.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_BPool *BPoolFilterer) ParseApproval(log types.Log) (*BPoolApproval, error) {
	event := new(BPoolApproval)
	if err := _BPool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BPoolLOGEXITIterator is returned from FilterLOGEXIT and is used to iterate over the raw logs and unpacked data for LOGEXIT events raised by the BPool contract.
type BPoolLOGEXITIterator struct {
	Event *BPoolLOGEXIT // Event containing the contract specifics and raw log

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
func (it *BPoolLOGEXITIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BPoolLOGEXIT)
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
		it.Event = new(BPoolLOGEXIT)
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
func (it *BPoolLOGEXITIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BPoolLOGEXITIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BPoolLOGEXIT represents a LOGEXIT event raised by the BPool contract.
type BPoolLOGEXIT struct {
	Caller         common.Address
	TokenOut       common.Address
	TokenAmountOut *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLOGEXIT is a free log retrieval operation binding the contract event 0xe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed.
//
// Solidity: event LOG_EXIT(address indexed caller, address indexed tokenOut, uint256 tokenAmountOut)
func (_BPool *BPoolFilterer) FilterLOGEXIT(opts *bind.FilterOpts, caller []common.Address, tokenOut []common.Address) (*BPoolLOGEXITIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _BPool.contract.FilterLogs(opts, "LOG_EXIT", callerRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &BPoolLOGEXITIterator{contract: _BPool.contract, event: "LOG_EXIT", logs: logs, sub: sub}, nil
}

// WatchLOGEXIT is a free log subscription operation binding the contract event 0xe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed.
//
// Solidity: event LOG_EXIT(address indexed caller, address indexed tokenOut, uint256 tokenAmountOut)
func (_BPool *BPoolFilterer) WatchLOGEXIT(opts *bind.WatchOpts, sink chan<- *BPoolLOGEXIT, caller []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _BPool.contract.WatchLogs(opts, "LOG_EXIT", callerRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BPoolLOGEXIT)
				if err := _BPool.contract.UnpackLog(event, "LOG_EXIT", log); err != nil {
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

// ParseLOGEXIT is a log parse operation binding the contract event 0xe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed.
//
// Solidity: event LOG_EXIT(address indexed caller, address indexed tokenOut, uint256 tokenAmountOut)
func (_BPool *BPoolFilterer) ParseLOGEXIT(log types.Log) (*BPoolLOGEXIT, error) {
	event := new(BPoolLOGEXIT)
	if err := _BPool.contract.UnpackLog(event, "LOG_EXIT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BPoolLOGJOINIterator is returned from FilterLOGJOIN and is used to iterate over the raw logs and unpacked data for LOGJOIN events raised by the BPool contract.
type BPoolLOGJOINIterator struct {
	Event *BPoolLOGJOIN // Event containing the contract specifics and raw log

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
func (it *BPoolLOGJOINIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BPoolLOGJOIN)
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
		it.Event = new(BPoolLOGJOIN)
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
func (it *BPoolLOGJOINIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BPoolLOGJOINIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BPoolLOGJOIN represents a LOGJOIN event raised by the BPool contract.
type BPoolLOGJOIN struct {
	Caller        common.Address
	TokenIn       common.Address
	TokenAmountIn *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLOGJOIN is a free log retrieval operation binding the contract event 0x63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a.
//
// Solidity: event LOG_JOIN(address indexed caller, address indexed tokenIn, uint256 tokenAmountIn)
func (_BPool *BPoolFilterer) FilterLOGJOIN(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address) (*BPoolLOGJOINIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _BPool.contract.FilterLogs(opts, "LOG_JOIN", callerRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return &BPoolLOGJOINIterator{contract: _BPool.contract, event: "LOG_JOIN", logs: logs, sub: sub}, nil
}

// WatchLOGJOIN is a free log subscription operation binding the contract event 0x63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a.
//
// Solidity: event LOG_JOIN(address indexed caller, address indexed tokenIn, uint256 tokenAmountIn)
func (_BPool *BPoolFilterer) WatchLOGJOIN(opts *bind.WatchOpts, sink chan<- *BPoolLOGJOIN, caller []common.Address, tokenIn []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _BPool.contract.WatchLogs(opts, "LOG_JOIN", callerRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BPoolLOGJOIN)
				if err := _BPool.contract.UnpackLog(event, "LOG_JOIN", log); err != nil {
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

// ParseLOGJOIN is a log parse operation binding the contract event 0x63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a.
//
// Solidity: event LOG_JOIN(address indexed caller, address indexed tokenIn, uint256 tokenAmountIn)
func (_BPool *BPoolFilterer) ParseLOGJOIN(log types.Log) (*BPoolLOGJOIN, error) {
	event := new(BPoolLOGJOIN)
	if err := _BPool.contract.UnpackLog(event, "LOG_JOIN", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BPoolLOGSWAPIterator is returned from FilterLOGSWAP and is used to iterate over the raw logs and unpacked data for LOGSWAP events raised by the BPool contract.
type BPoolLOGSWAPIterator struct {
	Event *BPoolLOGSWAP // Event containing the contract specifics and raw log

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
func (it *BPoolLOGSWAPIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BPoolLOGSWAP)
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
		it.Event = new(BPoolLOGSWAP)
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
func (it *BPoolLOGSWAPIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BPoolLOGSWAPIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BPoolLOGSWAP represents a LOGSWAP event raised by the BPool contract.
type BPoolLOGSWAP struct {
	Caller         common.Address
	TokenIn        common.Address
	TokenOut       common.Address
	TokenAmountIn  *big.Int
	TokenAmountOut *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLOGSWAP is a free log retrieval operation binding the contract event 0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378.
//
// Solidity: event LOG_SWAP(address indexed caller, address indexed tokenIn, address indexed tokenOut, uint256 tokenAmountIn, uint256 tokenAmountOut)
func (_BPool *BPoolFilterer) FilterLOGSWAP(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address, tokenOut []common.Address) (*BPoolLOGSWAPIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _BPool.contract.FilterLogs(opts, "LOG_SWAP", callerRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &BPoolLOGSWAPIterator{contract: _BPool.contract, event: "LOG_SWAP", logs: logs, sub: sub}, nil
}

// WatchLOGSWAP is a free log subscription operation binding the contract event 0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378.
//
// Solidity: event LOG_SWAP(address indexed caller, address indexed tokenIn, address indexed tokenOut, uint256 tokenAmountIn, uint256 tokenAmountOut)
func (_BPool *BPoolFilterer) WatchLOGSWAP(opts *bind.WatchOpts, sink chan<- *BPoolLOGSWAP, caller []common.Address, tokenIn []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _BPool.contract.WatchLogs(opts, "LOG_SWAP", callerRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BPoolLOGSWAP)
				if err := _BPool.contract.UnpackLog(event, "LOG_SWAP", log); err != nil {
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

// ParseLOGSWAP is a log parse operation binding the contract event 0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378.
//
// Solidity: event LOG_SWAP(address indexed caller, address indexed tokenIn, address indexed tokenOut, uint256 tokenAmountIn, uint256 tokenAmountOut)
func (_BPool *BPoolFilterer) ParseLOGSWAP(log types.Log) (*BPoolLOGSWAP, error) {
	event := new(BPoolLOGSWAP)
	if err := _BPool.contract.UnpackLog(event, "LOG_SWAP", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BPoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BPool contract.
type BPoolTransferIterator struct {
	Event *BPoolTransfer // Event containing the contract specifics and raw log

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
func (it *BPoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BPoolTransfer)
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
		it.Event = new(BPoolTransfer)
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
func (it *BPoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BPoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BPoolTransfer represents a Transfer event raised by the BPool contract.
type BPoolTransfer struct {
	Src common.Address
	Dst common.Address
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_BPool *BPoolFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*BPoolTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BPool.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &BPoolTransferIterator{contract: _BPool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_BPool *BPoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BPoolTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BPool.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BPoolTransfer)
				if err := _BPool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_BPool *BPoolFilterer) ParseTransfer(log types.Log) (*BPoolTransfer, error) {
	event := new(BPoolTransfer)
	if err := _BPool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BTokenABI is the input ABI used to generate the binding from.
const BTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"BONE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BPOW_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EXIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INIT_POOL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_IN_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OUT_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_TOTAL_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BTokenFuncSigs maps the 4-byte function signature to its string representation.
var BTokenFuncSigs = map[string]string{
	"c36596a6": "BONE()",
	"189d00ca": "BPOW_PRECISION()",
	"c6580d12": "EXIT_FEE()",
	"9381cd2b": "INIT_POOL_SUPPLY()",
	"b0e0d136": "MAX_BOUND_TOKENS()",
	"bc694ea2": "MAX_BPOW_BASE()",
	"bc063e1a": "MAX_FEE()",
	"ec093021": "MAX_IN_RATIO()",
	"992e2a92": "MAX_OUT_RATIO()",
	"09a3bbe4": "MAX_TOTAL_WEIGHT()",
	"e4a28a52": "MAX_WEIGHT()",
	"867378c5": "MIN_BALANCE()",
	"b7b800a4": "MIN_BOUND_TOKENS()",
	"ba019dab": "MIN_BPOW_BASE()",
	"76c7a3c7": "MIN_FEE()",
	"218b5382": "MIN_WEIGHT()",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"66188463": "decreaseApproval(address,uint256)",
	"9a86139b": "getColor()",
	"d73dd623": "increaseApproval(address,uint256)",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// BTokenBin is the compiled bytecode used for deploying new contracts.
var BTokenBin = "0x60c0604052601360808190527f42616c616e63657220506f6f6c20546f6b656e0000000000000000000000000060a090815261003e91600391906100a3565b506040805180820190915260038082527f42505400000000000000000000000000000000000000000000000000000000006020909201918252610083916004916100a3565b506005805460ff1916601217905534801561009d57600080fd5b5061013e565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100e457805160ff1916838001178555610111565b82800160010185558215610111579182015b828111156101115782518255916020019190600101906100f6565b5061011d929150610121565b5090565b61013b91905b8082111561011d5760008155600101610127565b90565b610bae8061014d6000396000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c8063992e2a92116100f9578063bc694ea211610097578063d73dd62311610071578063d73dd623146103ea578063dd62ed3e14610416578063e4a28a5214610286578063ec09302114610444576101c4565b8063bc694ea2146103da578063c36596a6146102b0578063c6580d12146103e2576101c4565b8063b0e0d136116100d3578063b0e0d136146103ba578063b7b800a4146103c2578063ba019dab146103ca578063bc063e1a146103d2576101c4565b8063992e2a921461037e5780639a86139b14610386578063a9059cbb1461038e576101c4565b8063313ce5671161016657806376c7a3c71161014057806376c7a3c71461035e578063867378c5146103665780639381cd2b1461036e57806395d89b4114610376576101c4565b8063313ce567146102ee578063661884631461030c57806370a0823114610338576101c4565b806318160ddd116101a257806318160ddd146102a0578063189d00ca146102a8578063218b5382146102b057806323b872dd146102b8576101c4565b806306fdde03146101c9578063095ea7b31461024657806309a3bbe414610286575b600080fd5b6101d161044c565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561020b5781810151838201526020016101f3565b50505050905090810190601f1680156102385780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102726004803603604081101561025c57600080fd5b506001600160a01b0381351690602001356104e2565b604080519115158252519081900360200190f35b61028e610536565b60408051918252519081900360200190f35b61028e610543565b61028e610549565b61028e61055d565b610272600480360360608110156102ce57600080fd5b506001600160a01b03813581169160208101359091169060400135610569565b6102f66106c3565b6040805160ff9092168252519081900360200190f35b6102726004803603604081101561032257600080fd5b506001600160a01b0381351690602001356106cc565b61028e6004803603602081101561034e57600080fd5b50356001600160a01b03166107a4565b61028e6107bf565b61028e6107d1565b61028e6107e5565b6101d16107f2565b61028e610853565b61028e61085f565b610272600480360360408110156103a457600080fd5b506001600160a01b03813516906020013561086c565b61028e610882565b61028e610887565b61028e61088c565b61028e610891565b61028e6108a1565b61028e6108ad565b6102726004803603604081101561040057600080fd5b506001600160a01b0381351690602001356108b2565b61028e6004803603604081101561042c57600080fd5b506001600160a01b0381358116916020013516610933565b61028e61095e565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104d85780601f106104ad576101008083540402835291602001916104d8565b820191906000526020600020905b8154815290600101906020018083116104bb57829003601f168201915b5050505050905090565b3360008181526001602090815260408083206001600160a01b03871680855290835281842086905581518681529151939490939092600080516020610b5a833981519152928290030190a350600192915050565b6802b5e3af16b188000081565b60025490565b6402540be400670de0b6b3a76400005b0481565b670de0b6b3a764000081565b6000336001600160a01b03851614806105a557506001600160a01b03841660009081526001602090815260408083203384529091529020548211155b6105ee576040805162461bcd60e51b815260206004820152601560248201527422a9292fa12a27a5a2a72fa120a22fa1a0a62622a960591b604482015290519081900360640190fd5b6105f984848461096e565b336001600160a01b0385161480159061063757506001600160a01b038416600090815260016020908152604080832033845290915290205460001914155b156106b9576001600160a01b038416600090815260016020908152604080832033845290915290205461066a9083610a7e565b6001600160a01b0385811660009081526001602090815260408083203380855290835292819020859055805194855251928716939192600080516020610b5a8339815191529281900390910190a35b5060019392505050565b60055460ff1690565b3360009081526001602090815260408083206001600160a01b038616845290915281205480831115610721573360009081526001602090815260408083206001600160a01b0388168452909152812055610750565b61072b8184610a7e565b3360009081526001602090815260408083206001600160a01b03891684529091529020555b3360008181526001602090815260408083206001600160a01b038916808552908352928190205481519081529051929392600080516020610b5a833981519152929181900390910190a35060019392505050565b6001600160a01b031660009081526020819052604090205490565b620f4240670de0b6b3a7640000610559565b64e8d4a51000670de0b6b3a7640000610559565b68056bc75e2d6310000081565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104d85780601f106104ad576101008083540402835291602001916104d8565b6704a03ce68d21555681565b6542524f4e5a4560d01b90565b600061087933848461096e565b50600192915050565b600881565b600281565b600181565b600a670de0b6b3a7640000610559565b671bc16d674ec7ffff81565b600081565b3360009081526001602090815260408083206001600160a01b03861684529091528120546108e09083610ae0565b3360008181526001602090815260408083206001600160a01b038916808552908352928190208590558051948552519193600080516020610b5a833981519152929081900390910190a350600192915050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6002670de0b6b3a7640000610559565b6001600160a01b0383166000908152602081905260409020548111156109d2576040805162461bcd60e51b815260206004820152601460248201527311549497d25394d551919250d251539517d0905360621b604482015290519081900360640190fd5b6001600160a01b0383166000908152602081905260409020546109f59082610a7e565b6001600160a01b038085166000908152602081905260408082209390935590841681522054610a249082610ae0565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000806000610a8d8585610b34565b915091508015610ad8576040805162461bcd60e51b81526020600482015260116024820152704552525f5355425f554e444552464c4f5760781b604482015290519081900360640190fd5b509392505050565b600082820183811015610b2d576040805162461bcd60e51b815260206004820152601060248201526f4552525f4144445f4f564552464c4f5760801b604482015290519081900360640190fd5b9392505050565b600080828410610b4a5750508082036000610b52565b505081810360015b925092905056fe8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925a265627a7a7231582051049149f1bfb9db00d688ad23a4d7be90411fa3eb410a40c9dc19d34100480964736f6c634300050c0032"

// DeployBToken deploys a new Ethereum contract, binding an instance of BToken to it.
func DeployBToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BToken{BTokenCaller: BTokenCaller{contract: contract}, BTokenTransactor: BTokenTransactor{contract: contract}, BTokenFilterer: BTokenFilterer{contract: contract}}, nil
}

// BToken is an auto generated Go binding around an Ethereum contract.
type BToken struct {
	BTokenCaller     // Read-only binding to the contract
	BTokenTransactor // Write-only binding to the contract
	BTokenFilterer   // Log filterer for contract events
}

// BTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BTokenSession struct {
	Contract     *BToken           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BTokenCallerSession struct {
	Contract *BTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BTokenTransactorSession struct {
	Contract     *BTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BTokenRaw struct {
	Contract *BToken // Generic contract binding to access the raw methods on
}

// BTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BTokenCallerRaw struct {
	Contract *BTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BTokenTransactorRaw struct {
	Contract *BTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBToken creates a new instance of BToken, bound to a specific deployed contract.
func NewBToken(address common.Address, backend bind.ContractBackend) (*BToken, error) {
	contract, err := bindBToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BToken{BTokenCaller: BTokenCaller{contract: contract}, BTokenTransactor: BTokenTransactor{contract: contract}, BTokenFilterer: BTokenFilterer{contract: contract}}, nil
}

// NewBTokenCaller creates a new read-only instance of BToken, bound to a specific deployed contract.
func NewBTokenCaller(address common.Address, caller bind.ContractCaller) (*BTokenCaller, error) {
	contract, err := bindBToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BTokenCaller{contract: contract}, nil
}

// NewBTokenTransactor creates a new write-only instance of BToken, bound to a specific deployed contract.
func NewBTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BTokenTransactor, error) {
	contract, err := bindBToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BTokenTransactor{contract: contract}, nil
}

// NewBTokenFilterer creates a new log filterer instance of BToken, bound to a specific deployed contract.
func NewBTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BTokenFilterer, error) {
	contract, err := bindBToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BTokenFilterer{contract: contract}, nil
}

// bindBToken binds a generic wrapper to an already deployed contract.
func bindBToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BToken *BTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BToken.Contract.BTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BToken *BTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BToken.Contract.BTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BToken *BTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BToken.Contract.BTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BToken *BTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BToken *BTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BToken *BTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BToken.Contract.contract.Transact(opts, method, params...)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BToken *BTokenCaller) BONE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "BONE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BToken *BTokenSession) BONE() (*big.Int, error) {
	return _BToken.Contract.BONE(&_BToken.CallOpts)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BToken *BTokenCallerSession) BONE() (*big.Int, error) {
	return _BToken.Contract.BONE(&_BToken.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BToken *BTokenCaller) BPOWPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "BPOW_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BToken *BTokenSession) BPOWPRECISION() (*big.Int, error) {
	return _BToken.Contract.BPOWPRECISION(&_BToken.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BToken *BTokenCallerSession) BPOWPRECISION() (*big.Int, error) {
	return _BToken.Contract.BPOWPRECISION(&_BToken.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BToken *BTokenCaller) EXITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "EXIT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BToken *BTokenSession) EXITFEE() (*big.Int, error) {
	return _BToken.Contract.EXITFEE(&_BToken.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BToken *BTokenCallerSession) EXITFEE() (*big.Int, error) {
	return _BToken.Contract.EXITFEE(&_BToken.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BToken *BTokenCaller) INITPOOLSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "INIT_POOL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BToken *BTokenSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _BToken.Contract.INITPOOLSUPPLY(&_BToken.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BToken *BTokenCallerSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _BToken.Contract.INITPOOLSUPPLY(&_BToken.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BToken *BTokenCaller) MAXBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "MAX_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BToken *BTokenSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _BToken.Contract.MAXBOUNDTOKENS(&_BToken.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BToken *BTokenCallerSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _BToken.Contract.MAXBOUNDTOKENS(&_BToken.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BToken *BTokenCaller) MAXBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "MAX_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BToken *BTokenSession) MAXBPOWBASE() (*big.Int, error) {
	return _BToken.Contract.MAXBPOWBASE(&_BToken.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BToken *BTokenCallerSession) MAXBPOWBASE() (*big.Int, error) {
	return _BToken.Contract.MAXBPOWBASE(&_BToken.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BToken *BTokenCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BToken *BTokenSession) MAXFEE() (*big.Int, error) {
	return _BToken.Contract.MAXFEE(&_BToken.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BToken *BTokenCallerSession) MAXFEE() (*big.Int, error) {
	return _BToken.Contract.MAXFEE(&_BToken.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BToken *BTokenCaller) MAXINRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "MAX_IN_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BToken *BTokenSession) MAXINRATIO() (*big.Int, error) {
	return _BToken.Contract.MAXINRATIO(&_BToken.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BToken *BTokenCallerSession) MAXINRATIO() (*big.Int, error) {
	return _BToken.Contract.MAXINRATIO(&_BToken.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BToken *BTokenCaller) MAXOUTRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "MAX_OUT_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BToken *BTokenSession) MAXOUTRATIO() (*big.Int, error) {
	return _BToken.Contract.MAXOUTRATIO(&_BToken.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BToken *BTokenCallerSession) MAXOUTRATIO() (*big.Int, error) {
	return _BToken.Contract.MAXOUTRATIO(&_BToken.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BToken *BTokenCaller) MAXTOTALWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "MAX_TOTAL_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BToken *BTokenSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _BToken.Contract.MAXTOTALWEIGHT(&_BToken.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BToken *BTokenCallerSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _BToken.Contract.MAXTOTALWEIGHT(&_BToken.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BToken *BTokenCaller) MAXWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "MAX_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BToken *BTokenSession) MAXWEIGHT() (*big.Int, error) {
	return _BToken.Contract.MAXWEIGHT(&_BToken.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BToken *BTokenCallerSession) MAXWEIGHT() (*big.Int, error) {
	return _BToken.Contract.MAXWEIGHT(&_BToken.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BToken *BTokenCaller) MINBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "MIN_BALANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BToken *BTokenSession) MINBALANCE() (*big.Int, error) {
	return _BToken.Contract.MINBALANCE(&_BToken.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BToken *BTokenCallerSession) MINBALANCE() (*big.Int, error) {
	return _BToken.Contract.MINBALANCE(&_BToken.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BToken *BTokenCaller) MINBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "MIN_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BToken *BTokenSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _BToken.Contract.MINBOUNDTOKENS(&_BToken.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BToken *BTokenCallerSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _BToken.Contract.MINBOUNDTOKENS(&_BToken.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BToken *BTokenCaller) MINBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "MIN_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BToken *BTokenSession) MINBPOWBASE() (*big.Int, error) {
	return _BToken.Contract.MINBPOWBASE(&_BToken.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BToken *BTokenCallerSession) MINBPOWBASE() (*big.Int, error) {
	return _BToken.Contract.MINBPOWBASE(&_BToken.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BToken *BTokenCaller) MINFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "MIN_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BToken *BTokenSession) MINFEE() (*big.Int, error) {
	return _BToken.Contract.MINFEE(&_BToken.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BToken *BTokenCallerSession) MINFEE() (*big.Int, error) {
	return _BToken.Contract.MINFEE(&_BToken.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BToken *BTokenCaller) MINWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "MIN_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BToken *BTokenSession) MINWEIGHT() (*big.Int, error) {
	return _BToken.Contract.MINWEIGHT(&_BToken.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BToken *BTokenCallerSession) MINWEIGHT() (*big.Int, error) {
	return _BToken.Contract.MINWEIGHT(&_BToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_BToken *BTokenCaller) Allowance(opts *bind.CallOpts, src common.Address, dst common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "allowance", src, dst)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_BToken *BTokenSession) Allowance(src common.Address, dst common.Address) (*big.Int, error) {
	return _BToken.Contract.Allowance(&_BToken.CallOpts, src, dst)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_BToken *BTokenCallerSession) Allowance(src common.Address, dst common.Address) (*big.Int, error) {
	return _BToken.Contract.Allowance(&_BToken.CallOpts, src, dst)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_BToken *BTokenCaller) BalanceOf(opts *bind.CallOpts, whom common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "balanceOf", whom)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_BToken *BTokenSession) BalanceOf(whom common.Address) (*big.Int, error) {
	return _BToken.Contract.BalanceOf(&_BToken.CallOpts, whom)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_BToken *BTokenCallerSession) BalanceOf(whom common.Address) (*big.Int, error) {
	return _BToken.Contract.BalanceOf(&_BToken.CallOpts, whom)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BToken *BTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BToken *BTokenSession) Decimals() (uint8, error) {
	return _BToken.Contract.Decimals(&_BToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BToken *BTokenCallerSession) Decimals() (uint8, error) {
	return _BToken.Contract.Decimals(&_BToken.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BToken *BTokenCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "getColor")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BToken *BTokenSession) GetColor() ([32]byte, error) {
	return _BToken.Contract.GetColor(&_BToken.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BToken *BTokenCallerSession) GetColor() ([32]byte, error) {
	return _BToken.Contract.GetColor(&_BToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BToken *BTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BToken *BTokenSession) Name() (string, error) {
	return _BToken.Contract.Name(&_BToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BToken *BTokenCallerSession) Name() (string, error) {
	return _BToken.Contract.Name(&_BToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BToken *BTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BToken *BTokenSession) Symbol() (string, error) {
	return _BToken.Contract.Symbol(&_BToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BToken *BTokenCallerSession) Symbol() (string, error) {
	return _BToken.Contract.Symbol(&_BToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BToken *BTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BToken *BTokenSession) TotalSupply() (*big.Int, error) {
	return _BToken.Contract.TotalSupply(&_BToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BToken *BTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BToken.Contract.TotalSupply(&_BToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_BToken *BTokenTransactor) Approve(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BToken.contract.Transact(opts, "approve", dst, amt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_BToken *BTokenSession) Approve(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BToken.Contract.Approve(&_BToken.TransactOpts, dst, amt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_BToken *BTokenTransactorSession) Approve(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BToken.Contract.Approve(&_BToken.TransactOpts, dst, amt)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_BToken *BTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BToken.contract.Transact(opts, "decreaseApproval", dst, amt)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_BToken *BTokenSession) DecreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BToken.Contract.DecreaseApproval(&_BToken.TransactOpts, dst, amt)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_BToken *BTokenTransactorSession) DecreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BToken.Contract.DecreaseApproval(&_BToken.TransactOpts, dst, amt)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_BToken *BTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BToken.contract.Transact(opts, "increaseApproval", dst, amt)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_BToken *BTokenSession) IncreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BToken.Contract.IncreaseApproval(&_BToken.TransactOpts, dst, amt)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_BToken *BTokenTransactorSession) IncreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BToken.Contract.IncreaseApproval(&_BToken.TransactOpts, dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_BToken *BTokenTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BToken.contract.Transact(opts, "transfer", dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_BToken *BTokenSession) Transfer(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BToken.Contract.Transfer(&_BToken.TransactOpts, dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_BToken *BTokenTransactorSession) Transfer(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BToken.Contract.Transfer(&_BToken.TransactOpts, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_BToken *BTokenTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BToken.contract.Transact(opts, "transferFrom", src, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_BToken *BTokenSession) TransferFrom(src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BToken.Contract.TransferFrom(&_BToken.TransactOpts, src, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_BToken *BTokenTransactorSession) TransferFrom(src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BToken.Contract.TransferFrom(&_BToken.TransactOpts, src, dst, amt)
}

// BTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BToken contract.
type BTokenApprovalIterator struct {
	Event *BTokenApproval // Event containing the contract specifics and raw log

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
func (it *BTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BTokenApproval)
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
		it.Event = new(BTokenApproval)
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
func (it *BTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BTokenApproval represents a Approval event raised by the BToken contract.
type BTokenApproval struct {
	Src common.Address
	Dst common.Address
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_BToken *BTokenFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*BTokenApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BToken.contract.FilterLogs(opts, "Approval", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &BTokenApprovalIterator{contract: _BToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_BToken *BTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BTokenApproval, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BToken.contract.WatchLogs(opts, "Approval", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BTokenApproval)
				if err := _BToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_BToken *BTokenFilterer) ParseApproval(log types.Log) (*BTokenApproval, error) {
	event := new(BTokenApproval)
	if err := _BToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BToken contract.
type BTokenTransferIterator struct {
	Event *BTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BTokenTransfer)
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
		it.Event = new(BTokenTransfer)
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
func (it *BTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BTokenTransfer represents a Transfer event raised by the BToken contract.
type BTokenTransfer struct {
	Src common.Address
	Dst common.Address
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_BToken *BTokenFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*BTokenTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BToken.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &BTokenTransferIterator{contract: _BToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_BToken *BTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BTokenTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BToken.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BTokenTransfer)
				if err := _BToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_BToken *BTokenFilterer) ParseTransfer(log types.Log) (*BTokenTransfer, error) {
	event := new(BTokenTransfer)
	if err := _BToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BTokenBaseABI is the input ABI used to generate the binding from.
const BTokenBaseABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"BONE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BPOW_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EXIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INIT_POOL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_IN_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OUT_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_TOTAL_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BTokenBaseFuncSigs maps the 4-byte function signature to its string representation.
var BTokenBaseFuncSigs = map[string]string{
	"c36596a6": "BONE()",
	"189d00ca": "BPOW_PRECISION()",
	"c6580d12": "EXIT_FEE()",
	"9381cd2b": "INIT_POOL_SUPPLY()",
	"b0e0d136": "MAX_BOUND_TOKENS()",
	"bc694ea2": "MAX_BPOW_BASE()",
	"bc063e1a": "MAX_FEE()",
	"ec093021": "MAX_IN_RATIO()",
	"992e2a92": "MAX_OUT_RATIO()",
	"09a3bbe4": "MAX_TOTAL_WEIGHT()",
	"e4a28a52": "MAX_WEIGHT()",
	"867378c5": "MIN_BALANCE()",
	"b7b800a4": "MIN_BOUND_TOKENS()",
	"ba019dab": "MIN_BPOW_BASE()",
	"76c7a3c7": "MIN_FEE()",
	"218b5382": "MIN_WEIGHT()",
	"9a86139b": "getColor()",
}

// BTokenBaseBin is the compiled bytecode used for deploying new contracts.
var BTokenBaseBin = "0x608060405234801561001057600080fd5b50610288806100206000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c8063b0e0d136116100a2578063bc694ea211610071578063bc694ea214610182578063c36596a614610132578063c6580d121461018a578063e4a28a5214610110578063ec093021146101925761010b565b8063b0e0d13614610162578063b7b800a41461016a578063ba019dab14610172578063bc063e1a1461017a5761010b565b8063867378c5116100de578063867378c5146101425780639381cd2b1461014a578063992e2a92146101525780639a86139b1461015a5761010b565b806309a3bbe414610110578063189d00ca1461012a578063218b53821461013257806376c7a3c71461013a575b600080fd5b61011861019a565b60408051918252519081900360200190f35b6101186101a7565b6101186101bb565b6101186101c7565b6101186101d9565b6101186101ed565b6101186101fa565b610118610206565b610118610213565b610118610218565b61011861021d565b610118610222565b610118610232565b61011861023e565b610118610243565b6802b5e3af16b188000081565b6402540be400670de0b6b3a76400005b0481565b670de0b6b3a764000081565b620f4240670de0b6b3a76400006101b7565b64e8d4a51000670de0b6b3a76400006101b7565b68056bc75e2d6310000081565b6704a03ce68d21555681565b6542524f4e5a4560d01b90565b600881565b600281565b600181565b600a670de0b6b3a76400006101b7565b671bc16d674ec7ffff81565b600081565b6002670de0b6b3a76400006101b756fea265627a7a72315820f794ca1319f6016c87117cec19fab1bec743659d9a2fc909982edb8fddd7d17764736f6c634300050c0032"

// DeployBTokenBase deploys a new Ethereum contract, binding an instance of BTokenBase to it.
func DeployBTokenBase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BTokenBase, error) {
	parsed, err := abi.JSON(strings.NewReader(BTokenBaseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BTokenBaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BTokenBase{BTokenBaseCaller: BTokenBaseCaller{contract: contract}, BTokenBaseTransactor: BTokenBaseTransactor{contract: contract}, BTokenBaseFilterer: BTokenBaseFilterer{contract: contract}}, nil
}

// BTokenBase is an auto generated Go binding around an Ethereum contract.
type BTokenBase struct {
	BTokenBaseCaller     // Read-only binding to the contract
	BTokenBaseTransactor // Write-only binding to the contract
	BTokenBaseFilterer   // Log filterer for contract events
}

// BTokenBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type BTokenBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTokenBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BTokenBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTokenBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BTokenBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTokenBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BTokenBaseSession struct {
	Contract     *BTokenBase       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BTokenBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BTokenBaseCallerSession struct {
	Contract *BTokenBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BTokenBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BTokenBaseTransactorSession struct {
	Contract     *BTokenBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BTokenBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type BTokenBaseRaw struct {
	Contract *BTokenBase // Generic contract binding to access the raw methods on
}

// BTokenBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BTokenBaseCallerRaw struct {
	Contract *BTokenBaseCaller // Generic read-only contract binding to access the raw methods on
}

// BTokenBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BTokenBaseTransactorRaw struct {
	Contract *BTokenBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBTokenBase creates a new instance of BTokenBase, bound to a specific deployed contract.
func NewBTokenBase(address common.Address, backend bind.ContractBackend) (*BTokenBase, error) {
	contract, err := bindBTokenBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BTokenBase{BTokenBaseCaller: BTokenBaseCaller{contract: contract}, BTokenBaseTransactor: BTokenBaseTransactor{contract: contract}, BTokenBaseFilterer: BTokenBaseFilterer{contract: contract}}, nil
}

// NewBTokenBaseCaller creates a new read-only instance of BTokenBase, bound to a specific deployed contract.
func NewBTokenBaseCaller(address common.Address, caller bind.ContractCaller) (*BTokenBaseCaller, error) {
	contract, err := bindBTokenBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BTokenBaseCaller{contract: contract}, nil
}

// NewBTokenBaseTransactor creates a new write-only instance of BTokenBase, bound to a specific deployed contract.
func NewBTokenBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*BTokenBaseTransactor, error) {
	contract, err := bindBTokenBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BTokenBaseTransactor{contract: contract}, nil
}

// NewBTokenBaseFilterer creates a new log filterer instance of BTokenBase, bound to a specific deployed contract.
func NewBTokenBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*BTokenBaseFilterer, error) {
	contract, err := bindBTokenBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BTokenBaseFilterer{contract: contract}, nil
}

// bindBTokenBase binds a generic wrapper to an already deployed contract.
func bindBTokenBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BTokenBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BTokenBase *BTokenBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BTokenBase.Contract.BTokenBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BTokenBase *BTokenBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTokenBase.Contract.BTokenBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BTokenBase *BTokenBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BTokenBase.Contract.BTokenBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BTokenBase *BTokenBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BTokenBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BTokenBase *BTokenBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTokenBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BTokenBase *BTokenBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BTokenBase.Contract.contract.Transact(opts, method, params...)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BTokenBase *BTokenBaseCaller) BONE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTokenBase.contract.Call(opts, &out, "BONE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BTokenBase *BTokenBaseSession) BONE() (*big.Int, error) {
	return _BTokenBase.Contract.BONE(&_BTokenBase.CallOpts)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BTokenBase *BTokenBaseCallerSession) BONE() (*big.Int, error) {
	return _BTokenBase.Contract.BONE(&_BTokenBase.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BTokenBase *BTokenBaseCaller) BPOWPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTokenBase.contract.Call(opts, &out, "BPOW_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BTokenBase *BTokenBaseSession) BPOWPRECISION() (*big.Int, error) {
	return _BTokenBase.Contract.BPOWPRECISION(&_BTokenBase.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BTokenBase *BTokenBaseCallerSession) BPOWPRECISION() (*big.Int, error) {
	return _BTokenBase.Contract.BPOWPRECISION(&_BTokenBase.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BTokenBase *BTokenBaseCaller) EXITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTokenBase.contract.Call(opts, &out, "EXIT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BTokenBase *BTokenBaseSession) EXITFEE() (*big.Int, error) {
	return _BTokenBase.Contract.EXITFEE(&_BTokenBase.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BTokenBase *BTokenBaseCallerSession) EXITFEE() (*big.Int, error) {
	return _BTokenBase.Contract.EXITFEE(&_BTokenBase.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BTokenBase *BTokenBaseCaller) INITPOOLSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTokenBase.contract.Call(opts, &out, "INIT_POOL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BTokenBase *BTokenBaseSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _BTokenBase.Contract.INITPOOLSUPPLY(&_BTokenBase.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BTokenBase *BTokenBaseCallerSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _BTokenBase.Contract.INITPOOLSUPPLY(&_BTokenBase.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BTokenBase *BTokenBaseCaller) MAXBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTokenBase.contract.Call(opts, &out, "MAX_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BTokenBase *BTokenBaseSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _BTokenBase.Contract.MAXBOUNDTOKENS(&_BTokenBase.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BTokenBase *BTokenBaseCallerSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _BTokenBase.Contract.MAXBOUNDTOKENS(&_BTokenBase.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BTokenBase *BTokenBaseCaller) MAXBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTokenBase.contract.Call(opts, &out, "MAX_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BTokenBase *BTokenBaseSession) MAXBPOWBASE() (*big.Int, error) {
	return _BTokenBase.Contract.MAXBPOWBASE(&_BTokenBase.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BTokenBase *BTokenBaseCallerSession) MAXBPOWBASE() (*big.Int, error) {
	return _BTokenBase.Contract.MAXBPOWBASE(&_BTokenBase.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BTokenBase *BTokenBaseCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTokenBase.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BTokenBase *BTokenBaseSession) MAXFEE() (*big.Int, error) {
	return _BTokenBase.Contract.MAXFEE(&_BTokenBase.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BTokenBase *BTokenBaseCallerSession) MAXFEE() (*big.Int, error) {
	return _BTokenBase.Contract.MAXFEE(&_BTokenBase.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BTokenBase *BTokenBaseCaller) MAXINRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTokenBase.contract.Call(opts, &out, "MAX_IN_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BTokenBase *BTokenBaseSession) MAXINRATIO() (*big.Int, error) {
	return _BTokenBase.Contract.MAXINRATIO(&_BTokenBase.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BTokenBase *BTokenBaseCallerSession) MAXINRATIO() (*big.Int, error) {
	return _BTokenBase.Contract.MAXINRATIO(&_BTokenBase.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BTokenBase *BTokenBaseCaller) MAXOUTRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTokenBase.contract.Call(opts, &out, "MAX_OUT_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BTokenBase *BTokenBaseSession) MAXOUTRATIO() (*big.Int, error) {
	return _BTokenBase.Contract.MAXOUTRATIO(&_BTokenBase.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BTokenBase *BTokenBaseCallerSession) MAXOUTRATIO() (*big.Int, error) {
	return _BTokenBase.Contract.MAXOUTRATIO(&_BTokenBase.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BTokenBase *BTokenBaseCaller) MAXTOTALWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTokenBase.contract.Call(opts, &out, "MAX_TOTAL_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BTokenBase *BTokenBaseSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _BTokenBase.Contract.MAXTOTALWEIGHT(&_BTokenBase.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BTokenBase *BTokenBaseCallerSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _BTokenBase.Contract.MAXTOTALWEIGHT(&_BTokenBase.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BTokenBase *BTokenBaseCaller) MAXWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTokenBase.contract.Call(opts, &out, "MAX_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BTokenBase *BTokenBaseSession) MAXWEIGHT() (*big.Int, error) {
	return _BTokenBase.Contract.MAXWEIGHT(&_BTokenBase.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BTokenBase *BTokenBaseCallerSession) MAXWEIGHT() (*big.Int, error) {
	return _BTokenBase.Contract.MAXWEIGHT(&_BTokenBase.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BTokenBase *BTokenBaseCaller) MINBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTokenBase.contract.Call(opts, &out, "MIN_BALANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BTokenBase *BTokenBaseSession) MINBALANCE() (*big.Int, error) {
	return _BTokenBase.Contract.MINBALANCE(&_BTokenBase.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BTokenBase *BTokenBaseCallerSession) MINBALANCE() (*big.Int, error) {
	return _BTokenBase.Contract.MINBALANCE(&_BTokenBase.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BTokenBase *BTokenBaseCaller) MINBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTokenBase.contract.Call(opts, &out, "MIN_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BTokenBase *BTokenBaseSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _BTokenBase.Contract.MINBOUNDTOKENS(&_BTokenBase.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BTokenBase *BTokenBaseCallerSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _BTokenBase.Contract.MINBOUNDTOKENS(&_BTokenBase.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BTokenBase *BTokenBaseCaller) MINBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTokenBase.contract.Call(opts, &out, "MIN_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BTokenBase *BTokenBaseSession) MINBPOWBASE() (*big.Int, error) {
	return _BTokenBase.Contract.MINBPOWBASE(&_BTokenBase.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BTokenBase *BTokenBaseCallerSession) MINBPOWBASE() (*big.Int, error) {
	return _BTokenBase.Contract.MINBPOWBASE(&_BTokenBase.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BTokenBase *BTokenBaseCaller) MINFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTokenBase.contract.Call(opts, &out, "MIN_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BTokenBase *BTokenBaseSession) MINFEE() (*big.Int, error) {
	return _BTokenBase.Contract.MINFEE(&_BTokenBase.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BTokenBase *BTokenBaseCallerSession) MINFEE() (*big.Int, error) {
	return _BTokenBase.Contract.MINFEE(&_BTokenBase.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BTokenBase *BTokenBaseCaller) MINWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTokenBase.contract.Call(opts, &out, "MIN_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BTokenBase *BTokenBaseSession) MINWEIGHT() (*big.Int, error) {
	return _BTokenBase.Contract.MINWEIGHT(&_BTokenBase.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BTokenBase *BTokenBaseCallerSession) MINWEIGHT() (*big.Int, error) {
	return _BTokenBase.Contract.MINWEIGHT(&_BTokenBase.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BTokenBase *BTokenBaseCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BTokenBase.contract.Call(opts, &out, "getColor")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BTokenBase *BTokenBaseSession) GetColor() ([32]byte, error) {
	return _BTokenBase.Contract.GetColor(&_BTokenBase.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BTokenBase *BTokenBaseCallerSession) GetColor() ([32]byte, error) {
	return _BTokenBase.Contract.GetColor(&_BTokenBase.CallOpts)
}

// BTokenBaseApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BTokenBase contract.
type BTokenBaseApprovalIterator struct {
	Event *BTokenBaseApproval // Event containing the contract specifics and raw log

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
func (it *BTokenBaseApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BTokenBaseApproval)
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
		it.Event = new(BTokenBaseApproval)
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
func (it *BTokenBaseApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BTokenBaseApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BTokenBaseApproval represents a Approval event raised by the BTokenBase contract.
type BTokenBaseApproval struct {
	Src common.Address
	Dst common.Address
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_BTokenBase *BTokenBaseFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*BTokenBaseApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BTokenBase.contract.FilterLogs(opts, "Approval", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &BTokenBaseApprovalIterator{contract: _BTokenBase.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_BTokenBase *BTokenBaseFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BTokenBaseApproval, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BTokenBase.contract.WatchLogs(opts, "Approval", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BTokenBaseApproval)
				if err := _BTokenBase.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_BTokenBase *BTokenBaseFilterer) ParseApproval(log types.Log) (*BTokenBaseApproval, error) {
	event := new(BTokenBaseApproval)
	if err := _BTokenBase.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BTokenBaseTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BTokenBase contract.
type BTokenBaseTransferIterator struct {
	Event *BTokenBaseTransfer // Event containing the contract specifics and raw log

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
func (it *BTokenBaseTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BTokenBaseTransfer)
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
		it.Event = new(BTokenBaseTransfer)
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
func (it *BTokenBaseTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BTokenBaseTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BTokenBaseTransfer represents a Transfer event raised by the BTokenBase contract.
type BTokenBaseTransfer struct {
	Src common.Address
	Dst common.Address
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_BTokenBase *BTokenBaseFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*BTokenBaseTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BTokenBase.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &BTokenBaseTransferIterator{contract: _BTokenBase.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_BTokenBase *BTokenBaseFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BTokenBaseTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BTokenBase.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BTokenBaseTransfer)
				if err := _BTokenBase.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_BTokenBase *BTokenBaseFilterer) ParseTransfer(log types.Log) (*BTokenBaseTransfer, error) {
	event := new(BTokenBaseTransfer)
	if err := _BTokenBase.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, src common.Address, dst common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", src, dst)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(src common.Address, dst common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, src, dst)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(src common.Address, dst common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, src, dst)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, whom common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", whom)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(whom common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, whom)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(whom common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, whom)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", dst, amt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_IERC20 *IERC20Session) Approve(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, dst, amt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_IERC20 *IERC20Session) Transfer(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", src, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, src, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, src, dst, amt)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Src common.Address
	Dst common.Address
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*IERC20ApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	Src common.Address
	Dst common.Address
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*IERC20TransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
