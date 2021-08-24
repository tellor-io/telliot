// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tellorMesosphere

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

// AccessControlABI is the input ABI used to generate the binding from.
const AccessControlABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AccessControlFuncSigs maps the 4-byte function signature to its string representation.
var AccessControlFuncSigs = map[string]string{
	"a217fddf": "DEFAULT_ADMIN_ROLE()",
	"248a9ca3": "getRoleAdmin(bytes32)",
	"9010d07c": "getRoleMember(bytes32,uint256)",
	"ca15c873": "getRoleMemberCount(bytes32)",
	"2f2ff15d": "grantRole(bytes32,address)",
	"91d14854": "hasRole(bytes32,address)",
	"36568abe": "renounceRole(bytes32,address)",
	"d547741f": "revokeRole(bytes32,address)",
}

// AccessControl is an auto generated Go binding around an Ethereum contract.
type AccessControl struct {
	AccessControlCaller     // Read-only binding to the contract
	AccessControlTransactor // Write-only binding to the contract
	AccessControlFilterer   // Log filterer for contract events
}

// AccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlSession struct {
	Contract     *AccessControl    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlCallerSession struct {
	Contract *AccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlTransactorSession struct {
	Contract     *AccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlRaw struct {
	Contract *AccessControl // Generic contract binding to access the raw methods on
}

// AccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlCallerRaw struct {
	Contract *AccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlTransactorRaw struct {
	Contract *AccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControl creates a new instance of AccessControl, bound to a specific deployed contract.
func NewAccessControl(address common.Address, backend bind.ContractBackend) (*AccessControl, error) {
	contract, err := bindAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControl{AccessControlCaller: AccessControlCaller{contract: contract}, AccessControlTransactor: AccessControlTransactor{contract: contract}, AccessControlFilterer: AccessControlFilterer{contract: contract}}, nil
}

// NewAccessControlCaller creates a new read-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlCaller(address common.Address, caller bind.ContractCaller) (*AccessControlCaller, error) {
	contract, err := bindAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlCaller{contract: contract}, nil
}

// NewAccessControlTransactor creates a new write-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlTransactor, error) {
	contract, err := bindAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlTransactor{contract: contract}, nil
}

// NewAccessControlFilterer creates a new log filterer instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlFilterer, error) {
	contract, err := bindAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlFilterer{contract: contract}, nil
}

// bindAccessControl binds a generic wrapper to an already deployed contract.
func bindAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.AccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControl.Contract.DEFAULTADMINROLE(&_AccessControl.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControl.Contract.DEFAULTADMINROLE(&_AccessControl.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControl.Contract.GetRoleAdmin(&_AccessControl.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControl.Contract.GetRoleAdmin(&_AccessControl.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControl *AccessControlCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControl *AccessControlSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AccessControl.Contract.GetRoleMember(&_AccessControl.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControl *AccessControlCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AccessControl.Contract.GetRoleMember(&_AccessControl.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControl *AccessControlCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControl *AccessControlSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AccessControl.Contract.GetRoleMemberCount(&_AccessControl.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControl *AccessControlCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AccessControl.Contract.GetRoleMemberCount(&_AccessControl.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRole(&_AccessControl.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRole(&_AccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// AccessControlRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AccessControl contract.
type AccessControlRoleAdminChangedIterator struct {
	Event *AccessControlRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleAdminChanged)
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
		it.Event = new(AccessControlRoleAdminChanged)
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
func (it *AccessControlRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleAdminChanged represents a RoleAdminChanged event raised by the AccessControl contract.
type AccessControlRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControl *AccessControlFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AccessControlRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleAdminChangedIterator{contract: _AccessControl.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControl *AccessControlFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AccessControlRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleAdminChanged)
				if err := _AccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControl *AccessControlFilterer) ParseRoleAdminChanged(log types.Log) (*AccessControlRoleAdminChanged, error) {
	event := new(AccessControlRoleAdminChanged)
	if err := _AccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AccessControl contract.
type AccessControlRoleGrantedIterator struct {
	Event *AccessControlRoleGranted // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleGranted)
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
		it.Event = new(AccessControlRoleGranted)
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
func (it *AccessControlRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleGranted represents a RoleGranted event raised by the AccessControl contract.
type AccessControlRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleGrantedIterator{contract: _AccessControl.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AccessControlRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleGranted)
				if err := _AccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) ParseRoleGranted(log types.Log) (*AccessControlRoleGranted, error) {
	event := new(AccessControlRoleGranted)
	if err := _AccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AccessControl contract.
type AccessControlRoleRevokedIterator struct {
	Event *AccessControlRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleRevoked)
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
		it.Event = new(AccessControlRoleRevoked)
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
func (it *AccessControlRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleRevoked represents a RoleRevoked event raised by the AccessControl contract.
type AccessControlRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleRevokedIterator{contract: _AccessControl.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AccessControlRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleRevoked)
				if err := _AccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) ParseRoleRevoked(log types.Log) (*AccessControlRoleRevoked, error) {
	event := new(AccessControlRoleRevoked)
	if err := _AccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBin is the compiled bytecode used for deploying new contracts.
var AddressBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220241e7016c54a59e3fc86fa209a39a30f76c1980ccb900d3cbd7098aa5569771364736f6c63430007000033"

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// EnumerableSetABI is the input ABI used to generate the binding from.
const EnumerableSetABI = "[]"

// EnumerableSetBin is the compiled bytecode used for deploying new contracts.
var EnumerableSetBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c07c81376977e4bc108f9a534d956d3597aa968e0a6431f08fb69f97ccffd83c64736f6c63430007000033"

// DeployEnumerableSet deploys a new Ethereum contract, binding an instance of EnumerableSet to it.
func DeployEnumerableSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnumerableSet, error) {
	parsed, err := abi.JSON(strings.NewReader(EnumerableSetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EnumerableSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// EnumerableSet is an auto generated Go binding around an Ethereum contract.
type EnumerableSet struct {
	EnumerableSetCaller     // Read-only binding to the contract
	EnumerableSetTransactor // Write-only binding to the contract
	EnumerableSetFilterer   // Log filterer for contract events
}

// EnumerableSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnumerableSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnumerableSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnumerableSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnumerableSetSession struct {
	Contract     *EnumerableSet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnumerableSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnumerableSetCallerSession struct {
	Contract *EnumerableSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EnumerableSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnumerableSetTransactorSession struct {
	Contract     *EnumerableSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnumerableSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnumerableSetRaw struct {
	Contract *EnumerableSet // Generic contract binding to access the raw methods on
}

// EnumerableSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnumerableSetCallerRaw struct {
	Contract *EnumerableSetCaller // Generic read-only contract binding to access the raw methods on
}

// EnumerableSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnumerableSetTransactorRaw struct {
	Contract *EnumerableSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnumerableSet creates a new instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSet(address common.Address, backend bind.ContractBackend) (*EnumerableSet, error) {
	contract, err := bindEnumerableSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// NewEnumerableSetCaller creates a new read-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetCaller(address common.Address, caller bind.ContractCaller) (*EnumerableSetCaller, error) {
	contract, err := bindEnumerableSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetCaller{contract: contract}, nil
}

// NewEnumerableSetTransactor creates a new write-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumerableSetTransactor, error) {
	contract, err := bindEnumerableSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetTransactor{contract: contract}, nil
}

// NewEnumerableSetFilterer creates a new log filterer instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumerableSetFilterer, error) {
	contract, err := bindEnumerableSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetFilterer{contract: contract}, nil
}

// bindEnumerableSet binds a generic wrapper to an already deployed contract.
func bindEnumerableSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnumerableSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.EnumerableSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220abe1d1077184a78b4160b3f926e0d2af2798afe7e4404ed5ebbf194004b6c7c964736f6c63430007000033"

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

// TellorMesosphereABI is the input ABI used to generate the binding from.
const TellorMesosphereABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_quorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maximumDeviation\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPORTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin_address\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reporter_address\",\"type\":\"address\"}],\"name\":\"addReporter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"availableReporterIndices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getCurrentValue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getDataBefore\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin_address\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reporter_address\",\"type\":\"address\"}],\"name\":\"isReporter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"latestTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"latestValues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestValuesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_b\",\"type\":\"uint256\"}],\"name\":\"max\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maximumDeviation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_b\",\"type\":\"uint256\"}],\"name\":\"min\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfReporters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"numberReportersFromLatestBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"oldestTimestampFromLatestBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reporter_address\",\"type\":\"address\"}],\"name\":\"removeReporter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"reporterIndices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reporters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"submitValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"timestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maximumDeviation\",\"type\":\"uint256\"}],\"name\":\"updateMaximumDeviation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_quorum\",\"type\":\"uint256\"}],\"name\":\"updateQuorum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"values\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TellorMesosphereFuncSigs maps the 4-byte function signature to its string representation.
var TellorMesosphereFuncSigs = map[string]string{
	"a217fddf": "DEFAULT_ADMIN_ROLE()",
	"3f60d799": "REPORTER_ROLE()",
	"70480275": "addAdmin(address)",
	"dd8755f2": "addReporter(address)",
	"85524a4f": "availableReporterIndices(uint256)",
	"3fcad964": "getCurrentValue(uint256)",
	"66b44611": "getDataBefore(uint256,uint256)",
	"46eee1c4": "getNewValueCountbyRequestId(uint256)",
	"248a9ca3": "getRoleAdmin(bytes32)",
	"9010d07c": "getRoleMember(bytes32,uint256)",
	"ca15c873": "getRoleMemberCount(bytes32)",
	"77fbb663": "getTimestampbyRequestIDandIndex(uint256,uint256)",
	"2f2ff15d": "grantRole(bytes32,address)",
	"91d14854": "hasRole(bytes32,address)",
	"24d7806c": "isAdmin(address)",
	"044ad7be": "isReporter(address)",
	"72d725e6": "latestTimestamps(uint256,uint256)",
	"8ee15314": "latestValues(uint256,uint256)",
	"550d76be": "latestValuesLength()",
	"6d5433e6": "max(uint256,uint256)",
	"9152c29d": "maximumDeviation()",
	"7ae2b5c7": "min(uint256,uint256)",
	"8d1f7daf": "numberOfReporters()",
	"9ee70ff4": "numberReportersFromLatestBlock(uint256)",
	"e8207c4e": "oldestTimestampFromLatestBlock(uint256)",
	"1703a018": "quorum()",
	"5de5c212": "removeReporter(address)",
	"8bad0c0a": "renounceAdmin()",
	"36568abe": "renounceRole(bytes32,address)",
	"6231bc11": "reporterIndices(address)",
	"5e02cb7d": "reporters(uint256)",
	"93fa4915": "retrieveData(uint256,uint256)",
	"d547741f": "revokeRole(bytes32,address)",
	"62f55112": "submitValue(uint256,uint256)",
	"c08d1fe5": "timeLimit()",
	"fb0ceb04": "timestamps(uint256,uint256)",
	"2c8b71d4": "updateMaximumDeviation(uint256)",
	"35680dc2": "updateQuorum(uint256)",
	"a3183701": "values(uint256,uint256)",
}

// TellorMesosphereBin is the compiled bytecode used for deploying new contracts.
var TellorMesosphereBin = "0x60806040523480156200001157600080fd5b5060405162001d0138038062001d01833981810160405260608110156200003757600080fd5b5080516020820151604090920151600e829055600c839055600d819055909190620000646000336200009a565b620000917f176c2b761bfeb5dab89f614b6c08152e31d9230394b3605eabf32249ea1c89a66000620000aa565b50505062000200565b620000a68282620000fc565b5050565b600082815260208190526040808220600201549051839285917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a460009182526020829052604090912060020155565b60008281526020818152604090912062000121918390620010d162000175821b17901c565b15620000a6576200013162000195565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60006200018c836001600160a01b03841662000199565b90505b92915050565b3390565b6000620001a78383620001e8565b620001df575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556200018f565b5060006200018f565b60009081526001919091016020526040902054151590565b611af180620002106000396000f3fe608060405234801561001057600080fd5b506004361061023d5760003560e01c806372d725e61161013b57806393fa4915116100b8578063ca15c8731161007c578063ca15c8731461066e578063d547741f1461068b578063dd8755f2146106b7578063e8207c4e146106dd578063fb0ceb04146106fa5761023d565b806393fa4915146105fb5780639ee70ff41461061e578063a217fddf1461063b578063a318370114610643578063c08d1fe5146106665761023d565b80638d1f7daf116100ff5780638d1f7daf146105795780638ee15314146105815780639010d07c146105a45780639152c29d146105c757806391d14854146105cf5761023d565b806372d725e6146104eb57806377fbb6631461050e5780637ae2b5c71461053157806385524a4f146105545780638bad0c0a146105715761023d565b80633fcad964116101c95780636231bc111161018d5780636231bc111461043657806362f551121461045c57806366b446111461047f5780636d5433e6146104a257806370480275146104c55761023d565b80633fcad9641461037557806346eee1c4146103b2578063550d76be146103cf5780635de5c212146103d75780635e02cb7d146103fd5761023d565b80632c8b71d4116102105780632c8b71d4146102d95780632f2ff15d146102f857806335680dc21461032457806336568abe146103415780633f60d7991461036d5761023d565b8063044ad7be146102425780631703a0181461027c578063248a9ca31461029657806324d7806c146102b3575b600080fd5b6102686004803603602081101561025857600080fd5b50356001600160a01b031661071d565b604080519115158252519081900360200190f35b61028461073d565b60408051918252519081900360200190f35b610284600480360360208110156102ac57600080fd5b5035610743565b610268600480360360208110156102c957600080fd5b50356001600160a01b0316610758565b6102f6600480360360208110156102ef57600080fd5b5035610764565b005b6102f66004803603604081101561030e57600080fd5b50803590602001356001600160a01b03166107bb565b6102f66004803603602081101561033a57600080fd5b5035610827565b6102f66004803603604081101561035757600080fd5b50803590602001356001600160a01b031661087e565b6102846108df565b6103926004803603602081101561038b57600080fd5b50356108f1565b604080519315158452602084019290925282820152519081900360600190f35b610284600480360360208110156103c857600080fd5b5035610992565b6102846109a4565b6102f6600480360360208110156103ed57600080fd5b50356001600160a01b03166109aa565b61041a6004803603602081101561041357600080fd5b5035610a8e565b604080516001600160a01b039092168252519081900360200190f35b6102846004803603602081101561044c57600080fd5b50356001600160a01b0316610aa9565b6102f66004803603604081101561047257600080fd5b5080359060200135610abb565b6103926004803603604081101561049557600080fd5b5080359060200135610c78565b610284600480360360408110156104b857600080fd5b5080359060200135610cee565b6102f6600480360360208110156104db57600080fd5b50356001600160a01b0316610d05565b6102846004803603604081101561050157600080fd5b5080359060200135610d65565b6102846004803603604081101561052457600080fd5b5080359060200135610d82565b6102846004803603604081101561054757600080fd5b5080359060200135610ddb565b6102846004803603602081101561056a57600080fd5b5035610dec565b6102f6610e0a565b610284610e17565b6102846004803603604081101561059757600080fd5b5080359060200135610e1d565b61041a600480360360408110156105ba57600080fd5b5080359060200135610e3a565b610284610e59565b610268600480360360408110156105e557600080fd5b50803590602001356001600160a01b0316610e5f565b6102846004803603604081101561061157600080fd5b5080359060200135610e77565b6102846004803603602081101561063457600080fd5b5035610e94565b610284610ea6565b6102846004803603604081101561065957600080fd5b5080359060200135610eab565b610284610ec8565b6102846004803603602081101561068457600080fd5b5035610ece565b6102f6600480360360408110156106a157600080fd5b50803590602001356001600160a01b0316610ee5565b6102f6600480360360208110156106cd57600080fd5b50356001600160a01b0316610f3e565b610284600480360360208110156106f357600080fd5b5035611091565b6102846004803603604081101561071057600080fd5b50803590602001356110a3565b6000610737600080516020611a1c83398151915283610e5f565b92915050565b600e5481565b60009081526020819052604090206002015490565b60006107378183610e5f565b61076d33610758565b6107b6576040805162461bcd60e51b81526020600482015260156024820152742932b9ba3934b1ba32b2103a379030b236b4b7399760591b604482015290519081900360640190fd5b600d55565b6000828152602081905260409020600201546107de906107d96110e6565b610e5f565b6108195760405162461bcd60e51b815260040180806020018281038252602f8152602001806119ed602f913960400191505060405180910390fd5b61082382826110ea565b5050565b61083033610758565b610879576040805162461bcd60e51b81526020600482015260156024820152742932b9ba3934b1ba32b2103a379030b236b4b7399760591b604482015290519081900360640190fd5b600e55565b6108866110e6565b6001600160a01b0316816001600160a01b0316146108d55760405162461bcd60e51b815260040180806020018281038252602f815260200180611a8d602f913960400191505060405180910390fd5b6108238282611153565b600080516020611a1c83398151915281565b60008060008061090085610992565b905061090f600b546005610ddb565b6000868152600860205260409020541080156109415750600c5442036007600087815260200190815260200160002054115b1561094b57600019015b600061095a8660018403610d82565b905060006109688783610e77565b9050801561097f57600195509350915061098b9050565b50600094508493509150505b9193909250565b60009081526002602052604090205490565b600a5481565b6109b333610758565b6109fc576040805162461bcd60e51b81526020600482015260156024820152742932b9ba3934b1ba32b2103a379030b236b4b7399760591b604482015290519081900360640190fd5b6001600160a01b038116600090815260046020818152604080842080548086526003845291852080546001600160a01b031916905592909152908290556009805460018101825592527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af909101819055600b8054600019019055610823600080516020611a1c83398151915283610ee5565b6003602052600090815260409020546001600160a01b031681565b60046020526000908152604090205481565b610ac43361071d565b610b15576040805162461bcd60e51b815260206004820152601860248201527f5265737472696374656420746f207265706f72746572732e0000000000000000604482015290519081900360640190fd5b6000828152600560209081526040808320338452600483528184208054855290835281842085905585845260068352818420905484529091528120429055808080610b5f866111bc565b93509350935093508315610c705760008681526007602052604090205482148015610b995750600b54600087815260086020526040902054105b15610c155760008681526002602052604081208054600019810192919083908110610bc057fe5b60009182526020808320909101548a835260018252604080842082855283528084208490558b8452600290925291208054919250429184908110610c0057fe5b90600052602060002001819055505050610c46565b6000868152600260209081526040808320805460018101825590845282842042910155888352600790915290208290555b60008681526001602090815260408083204284528252808320869055888352600890915290208190555b505050505050565b6000806000806000610c8a8787611697565b9150915081610ca55760008060009450945094505050610ce7565b6000610cb18883610d82565b90506000610cbf8983610e77565b90508015610cd7576001965094509250610ce7915050565b6000806000965096509650505050505b9250925092565b600081831115610cff575081610737565b50919050565b610d0e33610758565b610d57576040805162461bcd60e51b81526020600482015260156024820152742932b9ba3934b1ba32b2103a379030b236b4b7399760591b604482015290519081900360640190fd5b610d626000826107bb565b50565b600660209081526000928352604080842090915290825290205481565b600082815260026020526040812054801580610d9e5750828111155b15610dad576000915050610737565b6000848152600260205260409020805484908110610dc757fe5b906000526020600020015491505092915050565b600081831015610cff575081610737565b60098181548110610df957fe5b600091825260209091200154905081565b610e1560003361087e565b565b600b5481565b600560209081526000928352604080842090915290825290205481565b6000828152602081905260408120610e52908361179f565b9392505050565b600d5481565b6000828152602081905260408120610e5290836117ab565b600091825260016020908152604080842092845291905290205490565b60086020526000908152604090205481565b600081565b600160209081526000928352604080842090915290825290205481565b600c5481565b6000818152602081905260408120610737906117c0565b600082815260208190526040902060020154610f03906107d96110e6565b6108d55760405162461bcd60e51b8152600401808060200182810382526030815260200180611a3c6030913960400191505060405180910390fd5b610f4733610758565b610f90576040805162461bcd60e51b81526020600482015260156024820152742932b9ba3934b1ba32b2103a379030b236b4b7399760591b604482015290519081900360640190fd5b610f998161071d565b15610fd55760405162461bcd60e51b8152600401808060200182810382526021815260200180611a6c6021913960400191505060405180910390fd5b6009546000901561102557600980546000198101908110610ff257fe5b90600052602060002001549050600980548061100a57fe5b60019003818190600052602060002001600090559055611037565b50600b54600a80546001908101909155015b600081815260036020908152604080832080546001600160a01b0319166001600160a01b038716908117909155835260049091529020819055600b80546001019055610823600080516020611a1c833981519152836107bb565b60076020526000908152604090205481565b600260205281600052604060002081815481106110bc57fe5b90600052602060002001600091509150505481565b6000610e52836001600160a01b0384166117cb565b3390565b600082815260208190526040902061110290826110d1565b156108235761110f6110e6565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b600082815260208190526040902061116b9082611815565b15610823576111786110e6565b6001600160a01b0316816001600160a01b0316837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45050565b6000806000806060600b5467ffffffffffffffff811180156111dd57600080fd5b50604051908082528060200260200182016040528015611207578160200160208202803683370190505b5090506060600b5467ffffffffffffffff8111801561122557600080fd5b5060405190808252806020026020018201604052801561124f578160200160208202803683370190505b50905060004260015b600a54811161132157600c5460008b8152600660209081526040808320858452909152902054429190910310156113195760008a815260056020908152604080832084845290915290205485518690859081106112b157fe5b602002602001018181525050808484806001019550815181106112d057fe5b60209081029190910181019190915260008b815260068252604080822084835290925220548211156113195760008a815260066020908152604080832084845290915290205491505b600101611258565b50600e5482101561134357600080600080975097509750975050505050611690565b60015b8281101561146657805b600081118015611388575085818151811061136757fe5b602002602001015186600183038151811061137e57fe5b6020026020010151115b1561145d57600086828151811061139b57fe5b6020026020010151905060008683815181106113b357fe5b602002602001015190508760018403815181106113cc57fe5b60200260200101518884815181106113e057fe5b6020026020010181815250508660018403815181106113fb57fe5b602002602001015187848151811061140f57fe5b6020026020010181815250508188600185038151811061142b57fe5b6020026020010181815250508087600185038151811061144757fe5b6020908102919091010152505060001901611350565b50600101611346565b5060006114728a6108f1565b5091505080156115fa57600d54818660008151811061148d57fe5b60200260200101518760018703815181106114a457fe5b60200260200101510361271002816114b857fe5b0411156115fa57600e5460018403106115e2576114e9856000815181106114db57fe5b602002602001015182610ddb565b611507866000815181106114f957fe5b602002602001015183610cee565b0361152886600186038151811061151a57fe5b602002602001015183610ddb565b61154887600187038151811061153a57fe5b602002602001015184610cee565b0311156115905760008a8152600660205260408120855182908790600019880190811061157157fe5b60200260200101518152602001908152602001600020819055506115c7565b60008a815260066020526040812085518290879082906115ac57fe5b60200260200101518152602001908152602001600020819055505b6115d08a6111bc565b98509850985098505050505050611690565b60008060008098509850985098505050505050611690565b600060028406600114156116265785600285048151811061161757fe5b6020026020010151905061167b565b6002611670876001838804038151811061163c57fe5b6020026020010151886002888161164f57fe5b048151811061165a57fe5b602002602001015161182a90919063ffffffff16565b8161167757fe5b0490505b60019950975090955090935061169092505050565b9193509193565b60008060006116a585610992565b9050801561178f576000806000198301816116c08982610d82565b90508781106116da57600080965096505050505050611798565b6116e48983610d82565b9050878110156116fe575060019550935061179892505050565b826002818403046001010193506117158985610d82565b90508781101561175557600061172e8a86600101610d82565b90508881106117495760018597509750505050505050611798565b8460010193505061178a565b60006117648a60018703610d82565b90508881101561178257600180860397509750505050505050611798565b600185039250505b6116fe565b60008092509250505b9250929050565b6000610e528383611884565b6000610e52836001600160a01b0384166118e8565b600061073782611900565b60006117d783836118e8565b61180d57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610737565b506000610737565b6000610e52836001600160a01b038416611904565b600082820183811015610e52576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b815460009082106118c65760405162461bcd60e51b81526004018080602001828103825260228152602001806119cb6022913960400191505060405180910390fd5b8260000182815481106118d557fe5b9060005260206000200154905092915050565b60009081526001919091016020526040902054151590565b5490565b600081815260018301602052604081205480156119c0578354600019808301919081019060009087908390811061193757fe5b906000526020600020015490508087600001848154811061195457fe5b60009182526020808320909101929092558281526001898101909252604090209084019055865487908061198457fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050610737565b600091505061073756fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e6473416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f206772616e74176c2b761bfeb5dab89f614b6c08152e31d9230394b3605eabf32249ea1c89a6416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f207265766f6b654164647265737320616c726561647920686173207265706f7274657220726f6c65416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636520726f6c657320666f722073656c66a2646970667358221220f158d645172f7df80985c025544941996090fe2236dcc388af75e6851e06ec7464736f6c63430007000033"

// DeployTellorMesosphere deploys a new Ethereum contract, binding an instance of TellorMesosphere to it.
func DeployTellorMesosphere(auth *bind.TransactOpts, backend bind.ContractBackend, _quorum *big.Int, _timeLimit *big.Int, _maximumDeviation *big.Int) (common.Address, *types.Transaction, *TellorMesosphere, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorMesosphereABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TellorMesosphereBin), backend, _quorum, _timeLimit, _maximumDeviation)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TellorMesosphere{TellorMesosphereCaller: TellorMesosphereCaller{contract: contract}, TellorMesosphereTransactor: TellorMesosphereTransactor{contract: contract}, TellorMesosphereFilterer: TellorMesosphereFilterer{contract: contract}}, nil
}

// TellorMesosphere is an auto generated Go binding around an Ethereum contract.
type TellorMesosphere struct {
	TellorMesosphereCaller     // Read-only binding to the contract
	TellorMesosphereTransactor // Write-only binding to the contract
	TellorMesosphereFilterer   // Log filterer for contract events
}

// TellorMesosphereCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorMesosphereCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorMesosphereTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorMesosphereTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorMesosphereFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorMesosphereFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorMesosphereSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorMesosphereSession struct {
	Contract     *TellorMesosphere // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorMesosphereCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorMesosphereCallerSession struct {
	Contract *TellorMesosphereCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TellorMesosphereTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorMesosphereTransactorSession struct {
	Contract     *TellorMesosphereTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TellorMesosphereRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorMesosphereRaw struct {
	Contract *TellorMesosphere // Generic contract binding to access the raw methods on
}

// TellorMesosphereCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorMesosphereCallerRaw struct {
	Contract *TellorMesosphereCaller // Generic read-only contract binding to access the raw methods on
}

// TellorMesosphereTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorMesosphereTransactorRaw struct {
	Contract *TellorMesosphereTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellorMesosphere creates a new instance of TellorMesosphere, bound to a specific deployed contract.
func NewTellorMesosphere(address common.Address, backend bind.ContractBackend) (*TellorMesosphere, error) {
	contract, err := bindTellorMesosphere(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TellorMesosphere{TellorMesosphereCaller: TellorMesosphereCaller{contract: contract}, TellorMesosphereTransactor: TellorMesosphereTransactor{contract: contract}, TellorMesosphereFilterer: TellorMesosphereFilterer{contract: contract}}, nil
}

// NewTellorMesosphereCaller creates a new read-only instance of TellorMesosphere, bound to a specific deployed contract.
func NewTellorMesosphereCaller(address common.Address, caller bind.ContractCaller) (*TellorMesosphereCaller, error) {
	contract, err := bindTellorMesosphere(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorMesosphereCaller{contract: contract}, nil
}

// NewTellorMesosphereTransactor creates a new write-only instance of TellorMesosphere, bound to a specific deployed contract.
func NewTellorMesosphereTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorMesosphereTransactor, error) {
	contract, err := bindTellorMesosphere(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorMesosphereTransactor{contract: contract}, nil
}

// NewTellorMesosphereFilterer creates a new log filterer instance of TellorMesosphere, bound to a specific deployed contract.
func NewTellorMesosphereFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorMesosphereFilterer, error) {
	contract, err := bindTellorMesosphere(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorMesosphereFilterer{contract: contract}, nil
}

// bindTellorMesosphere binds a generic wrapper to an already deployed contract.
func bindTellorMesosphere(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorMesosphereABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorMesosphere *TellorMesosphereRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorMesosphere.Contract.TellorMesosphereCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorMesosphere *TellorMesosphereRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.TellorMesosphereTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorMesosphere *TellorMesosphereRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.TellorMesosphereTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorMesosphere *TellorMesosphereCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorMesosphere.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorMesosphere *TellorMesosphereTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorMesosphere *TellorMesosphereTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TellorMesosphere *TellorMesosphereCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TellorMesosphere *TellorMesosphereSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TellorMesosphere.Contract.DEFAULTADMINROLE(&_TellorMesosphere.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TellorMesosphere *TellorMesosphereCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TellorMesosphere.Contract.DEFAULTADMINROLE(&_TellorMesosphere.CallOpts)
}

// REPORTERROLE is a free data retrieval call binding the contract method 0x3f60d799.
//
// Solidity: function REPORTER_ROLE() view returns(bytes32)
func (_TellorMesosphere *TellorMesosphereCaller) REPORTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "REPORTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REPORTERROLE is a free data retrieval call binding the contract method 0x3f60d799.
//
// Solidity: function REPORTER_ROLE() view returns(bytes32)
func (_TellorMesosphere *TellorMesosphereSession) REPORTERROLE() ([32]byte, error) {
	return _TellorMesosphere.Contract.REPORTERROLE(&_TellorMesosphere.CallOpts)
}

// REPORTERROLE is a free data retrieval call binding the contract method 0x3f60d799.
//
// Solidity: function REPORTER_ROLE() view returns(bytes32)
func (_TellorMesosphere *TellorMesosphereCallerSession) REPORTERROLE() ([32]byte, error) {
	return _TellorMesosphere.Contract.REPORTERROLE(&_TellorMesosphere.CallOpts)
}

// AvailableReporterIndices is a free data retrieval call binding the contract method 0x85524a4f.
//
// Solidity: function availableReporterIndices(uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCaller) AvailableReporterIndices(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "availableReporterIndices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableReporterIndices is a free data retrieval call binding the contract method 0x85524a4f.
//
// Solidity: function availableReporterIndices(uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereSession) AvailableReporterIndices(arg0 *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.AvailableReporterIndices(&_TellorMesosphere.CallOpts, arg0)
}

// AvailableReporterIndices is a free data retrieval call binding the contract method 0x85524a4f.
//
// Solidity: function availableReporterIndices(uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) AvailableReporterIndices(arg0 *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.AvailableReporterIndices(&_TellorMesosphere.CallOpts, arg0)
}

// GetCurrentValue is a free data retrieval call binding the contract method 0x3fcad964.
//
// Solidity: function getCurrentValue(uint256 _requestId) view returns(bool, uint256, uint256)
func (_TellorMesosphere *TellorMesosphereCaller) GetCurrentValue(opts *bind.CallOpts, _requestId *big.Int) (bool, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "getCurrentValue", _requestId)

	if err != nil {
		return *new(bool), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetCurrentValue is a free data retrieval call binding the contract method 0x3fcad964.
//
// Solidity: function getCurrentValue(uint256 _requestId) view returns(bool, uint256, uint256)
func (_TellorMesosphere *TellorMesosphereSession) GetCurrentValue(_requestId *big.Int) (bool, *big.Int, *big.Int, error) {
	return _TellorMesosphere.Contract.GetCurrentValue(&_TellorMesosphere.CallOpts, _requestId)
}

// GetCurrentValue is a free data retrieval call binding the contract method 0x3fcad964.
//
// Solidity: function getCurrentValue(uint256 _requestId) view returns(bool, uint256, uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) GetCurrentValue(_requestId *big.Int) (bool, *big.Int, *big.Int, error) {
	return _TellorMesosphere.Contract.GetCurrentValue(&_TellorMesosphere.CallOpts, _requestId)
}

// GetDataBefore is a free data retrieval call binding the contract method 0x66b44611.
//
// Solidity: function getDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool, uint256, uint256)
func (_TellorMesosphere *TellorMesosphereCaller) GetDataBefore(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "getDataBefore", _requestId, _timestamp)

	if err != nil {
		return *new(bool), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetDataBefore is a free data retrieval call binding the contract method 0x66b44611.
//
// Solidity: function getDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool, uint256, uint256)
func (_TellorMesosphere *TellorMesosphereSession) GetDataBefore(_requestId *big.Int, _timestamp *big.Int) (bool, *big.Int, *big.Int, error) {
	return _TellorMesosphere.Contract.GetDataBefore(&_TellorMesosphere.CallOpts, _requestId, _timestamp)
}

// GetDataBefore is a free data retrieval call binding the contract method 0x66b44611.
//
// Solidity: function getDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool, uint256, uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) GetDataBefore(_requestId *big.Int, _timestamp *big.Int) (bool, *big.Int, *big.Int, error) {
	return _TellorMesosphere.Contract.GetDataBefore(&_TellorMesosphere.CallOpts, _requestId, _timestamp)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "getNewValueCountbyRequestId", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.GetNewValueCountbyRequestId(&_TellorMesosphere.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.GetNewValueCountbyRequestId(&_TellorMesosphere.CallOpts, _requestId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TellorMesosphere *TellorMesosphereCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TellorMesosphere *TellorMesosphereSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TellorMesosphere.Contract.GetRoleAdmin(&_TellorMesosphere.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TellorMesosphere *TellorMesosphereCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TellorMesosphere.Contract.GetRoleAdmin(&_TellorMesosphere.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TellorMesosphere *TellorMesosphereCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TellorMesosphere *TellorMesosphereSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _TellorMesosphere.Contract.GetRoleMember(&_TellorMesosphere.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TellorMesosphere *TellorMesosphereCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _TellorMesosphere.Contract.GetRoleMember(&_TellorMesosphere.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _TellorMesosphere.Contract.GetRoleMemberCount(&_TellorMesosphere.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _TellorMesosphere.Contract.GetRoleMemberCount(&_TellorMesosphere.CallOpts, role)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 _index) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestId *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "getTimestampbyRequestIDandIndex", _requestId, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 _index) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereSession) GetTimestampbyRequestIDandIndex(_requestId *big.Int, _index *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.GetTimestampbyRequestIDandIndex(&_TellorMesosphere.CallOpts, _requestId, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 _index) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) GetTimestampbyRequestIDandIndex(_requestId *big.Int, _index *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.GetTimestampbyRequestIDandIndex(&_TellorMesosphere.CallOpts, _requestId, _index)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TellorMesosphere *TellorMesosphereCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TellorMesosphere *TellorMesosphereSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TellorMesosphere.Contract.HasRole(&_TellorMesosphere.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TellorMesosphere *TellorMesosphereCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TellorMesosphere.Contract.HasRole(&_TellorMesosphere.CallOpts, role, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _admin_address) view returns(bool)
func (_TellorMesosphere *TellorMesosphereCaller) IsAdmin(opts *bind.CallOpts, _admin_address common.Address) (bool, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "isAdmin", _admin_address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _admin_address) view returns(bool)
func (_TellorMesosphere *TellorMesosphereSession) IsAdmin(_admin_address common.Address) (bool, error) {
	return _TellorMesosphere.Contract.IsAdmin(&_TellorMesosphere.CallOpts, _admin_address)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _admin_address) view returns(bool)
func (_TellorMesosphere *TellorMesosphereCallerSession) IsAdmin(_admin_address common.Address) (bool, error) {
	return _TellorMesosphere.Contract.IsAdmin(&_TellorMesosphere.CallOpts, _admin_address)
}

// IsReporter is a free data retrieval call binding the contract method 0x044ad7be.
//
// Solidity: function isReporter(address _reporter_address) view returns(bool)
func (_TellorMesosphere *TellorMesosphereCaller) IsReporter(opts *bind.CallOpts, _reporter_address common.Address) (bool, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "isReporter", _reporter_address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReporter is a free data retrieval call binding the contract method 0x044ad7be.
//
// Solidity: function isReporter(address _reporter_address) view returns(bool)
func (_TellorMesosphere *TellorMesosphereSession) IsReporter(_reporter_address common.Address) (bool, error) {
	return _TellorMesosphere.Contract.IsReporter(&_TellorMesosphere.CallOpts, _reporter_address)
}

// IsReporter is a free data retrieval call binding the contract method 0x044ad7be.
//
// Solidity: function isReporter(address _reporter_address) view returns(bool)
func (_TellorMesosphere *TellorMesosphereCallerSession) IsReporter(_reporter_address common.Address) (bool, error) {
	return _TellorMesosphere.Contract.IsReporter(&_TellorMesosphere.CallOpts, _reporter_address)
}

// LatestTimestamps is a free data retrieval call binding the contract method 0x72d725e6.
//
// Solidity: function latestTimestamps(uint256 , uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCaller) LatestTimestamps(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "latestTimestamps", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestTimestamps is a free data retrieval call binding the contract method 0x72d725e6.
//
// Solidity: function latestTimestamps(uint256 , uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereSession) LatestTimestamps(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.LatestTimestamps(&_TellorMesosphere.CallOpts, arg0, arg1)
}

// LatestTimestamps is a free data retrieval call binding the contract method 0x72d725e6.
//
// Solidity: function latestTimestamps(uint256 , uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) LatestTimestamps(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.LatestTimestamps(&_TellorMesosphere.CallOpts, arg0, arg1)
}

// LatestValues is a free data retrieval call binding the contract method 0x8ee15314.
//
// Solidity: function latestValues(uint256 , uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCaller) LatestValues(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "latestValues", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestValues is a free data retrieval call binding the contract method 0x8ee15314.
//
// Solidity: function latestValues(uint256 , uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereSession) LatestValues(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.LatestValues(&_TellorMesosphere.CallOpts, arg0, arg1)
}

// LatestValues is a free data retrieval call binding the contract method 0x8ee15314.
//
// Solidity: function latestValues(uint256 , uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) LatestValues(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.LatestValues(&_TellorMesosphere.CallOpts, arg0, arg1)
}

// LatestValuesLength is a free data retrieval call binding the contract method 0x550d76be.
//
// Solidity: function latestValuesLength() view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCaller) LatestValuesLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "latestValuesLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestValuesLength is a free data retrieval call binding the contract method 0x550d76be.
//
// Solidity: function latestValuesLength() view returns(uint256)
func (_TellorMesosphere *TellorMesosphereSession) LatestValuesLength() (*big.Int, error) {
	return _TellorMesosphere.Contract.LatestValuesLength(&_TellorMesosphere.CallOpts)
}

// LatestValuesLength is a free data retrieval call binding the contract method 0x550d76be.
//
// Solidity: function latestValuesLength() view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) LatestValuesLength() (*big.Int, error) {
	return _TellorMesosphere.Contract.LatestValuesLength(&_TellorMesosphere.CallOpts)
}

// Max is a free data retrieval call binding the contract method 0x6d5433e6.
//
// Solidity: function max(uint256 _a, uint256 _b) pure returns(uint256)
func (_TellorMesosphere *TellorMesosphereCaller) Max(opts *bind.CallOpts, _a *big.Int, _b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "max", _a, _b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Max is a free data retrieval call binding the contract method 0x6d5433e6.
//
// Solidity: function max(uint256 _a, uint256 _b) pure returns(uint256)
func (_TellorMesosphere *TellorMesosphereSession) Max(_a *big.Int, _b *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.Max(&_TellorMesosphere.CallOpts, _a, _b)
}

// Max is a free data retrieval call binding the contract method 0x6d5433e6.
//
// Solidity: function max(uint256 _a, uint256 _b) pure returns(uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) Max(_a *big.Int, _b *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.Max(&_TellorMesosphere.CallOpts, _a, _b)
}

// MaximumDeviation is a free data retrieval call binding the contract method 0x9152c29d.
//
// Solidity: function maximumDeviation() view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCaller) MaximumDeviation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "maximumDeviation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumDeviation is a free data retrieval call binding the contract method 0x9152c29d.
//
// Solidity: function maximumDeviation() view returns(uint256)
func (_TellorMesosphere *TellorMesosphereSession) MaximumDeviation() (*big.Int, error) {
	return _TellorMesosphere.Contract.MaximumDeviation(&_TellorMesosphere.CallOpts)
}

// MaximumDeviation is a free data retrieval call binding the contract method 0x9152c29d.
//
// Solidity: function maximumDeviation() view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) MaximumDeviation() (*big.Int, error) {
	return _TellorMesosphere.Contract.MaximumDeviation(&_TellorMesosphere.CallOpts)
}

// Min is a free data retrieval call binding the contract method 0x7ae2b5c7.
//
// Solidity: function min(uint256 _a, uint256 _b) pure returns(uint256)
func (_TellorMesosphere *TellorMesosphereCaller) Min(opts *bind.CallOpts, _a *big.Int, _b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "min", _a, _b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Min is a free data retrieval call binding the contract method 0x7ae2b5c7.
//
// Solidity: function min(uint256 _a, uint256 _b) pure returns(uint256)
func (_TellorMesosphere *TellorMesosphereSession) Min(_a *big.Int, _b *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.Min(&_TellorMesosphere.CallOpts, _a, _b)
}

// Min is a free data retrieval call binding the contract method 0x7ae2b5c7.
//
// Solidity: function min(uint256 _a, uint256 _b) pure returns(uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) Min(_a *big.Int, _b *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.Min(&_TellorMesosphere.CallOpts, _a, _b)
}

// NumberOfReporters is a free data retrieval call binding the contract method 0x8d1f7daf.
//
// Solidity: function numberOfReporters() view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCaller) NumberOfReporters(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "numberOfReporters")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfReporters is a free data retrieval call binding the contract method 0x8d1f7daf.
//
// Solidity: function numberOfReporters() view returns(uint256)
func (_TellorMesosphere *TellorMesosphereSession) NumberOfReporters() (*big.Int, error) {
	return _TellorMesosphere.Contract.NumberOfReporters(&_TellorMesosphere.CallOpts)
}

// NumberOfReporters is a free data retrieval call binding the contract method 0x8d1f7daf.
//
// Solidity: function numberOfReporters() view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) NumberOfReporters() (*big.Int, error) {
	return _TellorMesosphere.Contract.NumberOfReporters(&_TellorMesosphere.CallOpts)
}

// NumberReportersFromLatestBlock is a free data retrieval call binding the contract method 0x9ee70ff4.
//
// Solidity: function numberReportersFromLatestBlock(uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCaller) NumberReportersFromLatestBlock(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "numberReportersFromLatestBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberReportersFromLatestBlock is a free data retrieval call binding the contract method 0x9ee70ff4.
//
// Solidity: function numberReportersFromLatestBlock(uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereSession) NumberReportersFromLatestBlock(arg0 *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.NumberReportersFromLatestBlock(&_TellorMesosphere.CallOpts, arg0)
}

// NumberReportersFromLatestBlock is a free data retrieval call binding the contract method 0x9ee70ff4.
//
// Solidity: function numberReportersFromLatestBlock(uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) NumberReportersFromLatestBlock(arg0 *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.NumberReportersFromLatestBlock(&_TellorMesosphere.CallOpts, arg0)
}

// OldestTimestampFromLatestBlock is a free data retrieval call binding the contract method 0xe8207c4e.
//
// Solidity: function oldestTimestampFromLatestBlock(uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCaller) OldestTimestampFromLatestBlock(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "oldestTimestampFromLatestBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OldestTimestampFromLatestBlock is a free data retrieval call binding the contract method 0xe8207c4e.
//
// Solidity: function oldestTimestampFromLatestBlock(uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereSession) OldestTimestampFromLatestBlock(arg0 *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.OldestTimestampFromLatestBlock(&_TellorMesosphere.CallOpts, arg0)
}

// OldestTimestampFromLatestBlock is a free data retrieval call binding the contract method 0xe8207c4e.
//
// Solidity: function oldestTimestampFromLatestBlock(uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) OldestTimestampFromLatestBlock(arg0 *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.OldestTimestampFromLatestBlock(&_TellorMesosphere.CallOpts, arg0)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCaller) Quorum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "quorum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint256)
func (_TellorMesosphere *TellorMesosphereSession) Quorum() (*big.Int, error) {
	return _TellorMesosphere.Contract.Quorum(&_TellorMesosphere.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) Quorum() (*big.Int, error) {
	return _TellorMesosphere.Contract.Quorum(&_TellorMesosphere.CallOpts)
}

// ReporterIndices is a free data retrieval call binding the contract method 0x6231bc11.
//
// Solidity: function reporterIndices(address ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCaller) ReporterIndices(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "reporterIndices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReporterIndices is a free data retrieval call binding the contract method 0x6231bc11.
//
// Solidity: function reporterIndices(address ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereSession) ReporterIndices(arg0 common.Address) (*big.Int, error) {
	return _TellorMesosphere.Contract.ReporterIndices(&_TellorMesosphere.CallOpts, arg0)
}

// ReporterIndices is a free data retrieval call binding the contract method 0x6231bc11.
//
// Solidity: function reporterIndices(address ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) ReporterIndices(arg0 common.Address) (*big.Int, error) {
	return _TellorMesosphere.Contract.ReporterIndices(&_TellorMesosphere.CallOpts, arg0)
}

// Reporters is a free data retrieval call binding the contract method 0x5e02cb7d.
//
// Solidity: function reporters(uint256 ) view returns(address)
func (_TellorMesosphere *TellorMesosphereCaller) Reporters(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "reporters", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Reporters is a free data retrieval call binding the contract method 0x5e02cb7d.
//
// Solidity: function reporters(uint256 ) view returns(address)
func (_TellorMesosphere *TellorMesosphereSession) Reporters(arg0 *big.Int) (common.Address, error) {
	return _TellorMesosphere.Contract.Reporters(&_TellorMesosphere.CallOpts, arg0)
}

// Reporters is a free data retrieval call binding the contract method 0x5e02cb7d.
//
// Solidity: function reporters(uint256 ) view returns(address)
func (_TellorMesosphere *TellorMesosphereCallerSession) Reporters(arg0 *big.Int) (common.Address, error) {
	return _TellorMesosphere.Contract.Reporters(&_TellorMesosphere.CallOpts, arg0)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "retrieveData", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.RetrieveData(&_TellorMesosphere.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.RetrieveData(&_TellorMesosphere.CallOpts, _requestId, _timestamp)
}

// TimeLimit is a free data retrieval call binding the contract method 0xc08d1fe5.
//
// Solidity: function timeLimit() view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCaller) TimeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "timeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeLimit is a free data retrieval call binding the contract method 0xc08d1fe5.
//
// Solidity: function timeLimit() view returns(uint256)
func (_TellorMesosphere *TellorMesosphereSession) TimeLimit() (*big.Int, error) {
	return _TellorMesosphere.Contract.TimeLimit(&_TellorMesosphere.CallOpts)
}

// TimeLimit is a free data retrieval call binding the contract method 0xc08d1fe5.
//
// Solidity: function timeLimit() view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) TimeLimit() (*big.Int, error) {
	return _TellorMesosphere.Contract.TimeLimit(&_TellorMesosphere.CallOpts)
}

// Timestamps is a free data retrieval call binding the contract method 0xfb0ceb04.
//
// Solidity: function timestamps(uint256 , uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCaller) Timestamps(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "timestamps", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timestamps is a free data retrieval call binding the contract method 0xfb0ceb04.
//
// Solidity: function timestamps(uint256 , uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereSession) Timestamps(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.Timestamps(&_TellorMesosphere.CallOpts, arg0, arg1)
}

// Timestamps is a free data retrieval call binding the contract method 0xfb0ceb04.
//
// Solidity: function timestamps(uint256 , uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) Timestamps(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.Timestamps(&_TellorMesosphere.CallOpts, arg0, arg1)
}

// Values is a free data retrieval call binding the contract method 0xa3183701.
//
// Solidity: function values(uint256 , uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCaller) Values(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorMesosphere.contract.Call(opts, &out, "values", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Values is a free data retrieval call binding the contract method 0xa3183701.
//
// Solidity: function values(uint256 , uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereSession) Values(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.Values(&_TellorMesosphere.CallOpts, arg0, arg1)
}

// Values is a free data retrieval call binding the contract method 0xa3183701.
//
// Solidity: function values(uint256 , uint256 ) view returns(uint256)
func (_TellorMesosphere *TellorMesosphereCallerSession) Values(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TellorMesosphere.Contract.Values(&_TellorMesosphere.CallOpts, arg0, arg1)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin_address) returns()
func (_TellorMesosphere *TellorMesosphereTransactor) AddAdmin(opts *bind.TransactOpts, _admin_address common.Address) (*types.Transaction, error) {
	return _TellorMesosphere.contract.Transact(opts, "addAdmin", _admin_address)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin_address) returns()
func (_TellorMesosphere *TellorMesosphereSession) AddAdmin(_admin_address common.Address) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.AddAdmin(&_TellorMesosphere.TransactOpts, _admin_address)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin_address) returns()
func (_TellorMesosphere *TellorMesosphereTransactorSession) AddAdmin(_admin_address common.Address) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.AddAdmin(&_TellorMesosphere.TransactOpts, _admin_address)
}

// AddReporter is a paid mutator transaction binding the contract method 0xdd8755f2.
//
// Solidity: function addReporter(address _reporter_address) returns()
func (_TellorMesosphere *TellorMesosphereTransactor) AddReporter(opts *bind.TransactOpts, _reporter_address common.Address) (*types.Transaction, error) {
	return _TellorMesosphere.contract.Transact(opts, "addReporter", _reporter_address)
}

// AddReporter is a paid mutator transaction binding the contract method 0xdd8755f2.
//
// Solidity: function addReporter(address _reporter_address) returns()
func (_TellorMesosphere *TellorMesosphereSession) AddReporter(_reporter_address common.Address) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.AddReporter(&_TellorMesosphere.TransactOpts, _reporter_address)
}

// AddReporter is a paid mutator transaction binding the contract method 0xdd8755f2.
//
// Solidity: function addReporter(address _reporter_address) returns()
func (_TellorMesosphere *TellorMesosphereTransactorSession) AddReporter(_reporter_address common.Address) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.AddReporter(&_TellorMesosphere.TransactOpts, _reporter_address)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TellorMesosphere *TellorMesosphereTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TellorMesosphere.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TellorMesosphere *TellorMesosphereSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.GrantRole(&_TellorMesosphere.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TellorMesosphere *TellorMesosphereTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.GrantRole(&_TellorMesosphere.TransactOpts, role, account)
}

// RemoveReporter is a paid mutator transaction binding the contract method 0x5de5c212.
//
// Solidity: function removeReporter(address _reporter_address) returns()
func (_TellorMesosphere *TellorMesosphereTransactor) RemoveReporter(opts *bind.TransactOpts, _reporter_address common.Address) (*types.Transaction, error) {
	return _TellorMesosphere.contract.Transact(opts, "removeReporter", _reporter_address)
}

// RemoveReporter is a paid mutator transaction binding the contract method 0x5de5c212.
//
// Solidity: function removeReporter(address _reporter_address) returns()
func (_TellorMesosphere *TellorMesosphereSession) RemoveReporter(_reporter_address common.Address) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.RemoveReporter(&_TellorMesosphere.TransactOpts, _reporter_address)
}

// RemoveReporter is a paid mutator transaction binding the contract method 0x5de5c212.
//
// Solidity: function removeReporter(address _reporter_address) returns()
func (_TellorMesosphere *TellorMesosphereTransactorSession) RemoveReporter(_reporter_address common.Address) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.RemoveReporter(&_TellorMesosphere.TransactOpts, _reporter_address)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_TellorMesosphere *TellorMesosphereTransactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorMesosphere.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_TellorMesosphere *TellorMesosphereSession) RenounceAdmin() (*types.Transaction, error) {
	return _TellorMesosphere.Contract.RenounceAdmin(&_TellorMesosphere.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_TellorMesosphere *TellorMesosphereTransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _TellorMesosphere.Contract.RenounceAdmin(&_TellorMesosphere.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TellorMesosphere *TellorMesosphereTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TellorMesosphere.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TellorMesosphere *TellorMesosphereSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.RenounceRole(&_TellorMesosphere.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TellorMesosphere *TellorMesosphereTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.RenounceRole(&_TellorMesosphere.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TellorMesosphere *TellorMesosphereTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TellorMesosphere.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TellorMesosphere *TellorMesosphereSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.RevokeRole(&_TellorMesosphere.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TellorMesosphere *TellorMesosphereTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.RevokeRole(&_TellorMesosphere.TransactOpts, role, account)
}

// SubmitValue is a paid mutator transaction binding the contract method 0x62f55112.
//
// Solidity: function submitValue(uint256 _requestId, uint256 _value) returns()
func (_TellorMesosphere *TellorMesosphereTransactor) SubmitValue(opts *bind.TransactOpts, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _TellorMesosphere.contract.Transact(opts, "submitValue", _requestId, _value)
}

// SubmitValue is a paid mutator transaction binding the contract method 0x62f55112.
//
// Solidity: function submitValue(uint256 _requestId, uint256 _value) returns()
func (_TellorMesosphere *TellorMesosphereSession) SubmitValue(_requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.SubmitValue(&_TellorMesosphere.TransactOpts, _requestId, _value)
}

// SubmitValue is a paid mutator transaction binding the contract method 0x62f55112.
//
// Solidity: function submitValue(uint256 _requestId, uint256 _value) returns()
func (_TellorMesosphere *TellorMesosphereTransactorSession) SubmitValue(_requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.SubmitValue(&_TellorMesosphere.TransactOpts, _requestId, _value)
}

// UpdateMaximumDeviation is a paid mutator transaction binding the contract method 0x2c8b71d4.
//
// Solidity: function updateMaximumDeviation(uint256 _maximumDeviation) returns()
func (_TellorMesosphere *TellorMesosphereTransactor) UpdateMaximumDeviation(opts *bind.TransactOpts, _maximumDeviation *big.Int) (*types.Transaction, error) {
	return _TellorMesosphere.contract.Transact(opts, "updateMaximumDeviation", _maximumDeviation)
}

// UpdateMaximumDeviation is a paid mutator transaction binding the contract method 0x2c8b71d4.
//
// Solidity: function updateMaximumDeviation(uint256 _maximumDeviation) returns()
func (_TellorMesosphere *TellorMesosphereSession) UpdateMaximumDeviation(_maximumDeviation *big.Int) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.UpdateMaximumDeviation(&_TellorMesosphere.TransactOpts, _maximumDeviation)
}

// UpdateMaximumDeviation is a paid mutator transaction binding the contract method 0x2c8b71d4.
//
// Solidity: function updateMaximumDeviation(uint256 _maximumDeviation) returns()
func (_TellorMesosphere *TellorMesosphereTransactorSession) UpdateMaximumDeviation(_maximumDeviation *big.Int) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.UpdateMaximumDeviation(&_TellorMesosphere.TransactOpts, _maximumDeviation)
}

// UpdateQuorum is a paid mutator transaction binding the contract method 0x35680dc2.
//
// Solidity: function updateQuorum(uint256 _quorum) returns()
func (_TellorMesosphere *TellorMesosphereTransactor) UpdateQuorum(opts *bind.TransactOpts, _quorum *big.Int) (*types.Transaction, error) {
	return _TellorMesosphere.contract.Transact(opts, "updateQuorum", _quorum)
}

// UpdateQuorum is a paid mutator transaction binding the contract method 0x35680dc2.
//
// Solidity: function updateQuorum(uint256 _quorum) returns()
func (_TellorMesosphere *TellorMesosphereSession) UpdateQuorum(_quorum *big.Int) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.UpdateQuorum(&_TellorMesosphere.TransactOpts, _quorum)
}

// UpdateQuorum is a paid mutator transaction binding the contract method 0x35680dc2.
//
// Solidity: function updateQuorum(uint256 _quorum) returns()
func (_TellorMesosphere *TellorMesosphereTransactorSession) UpdateQuorum(_quorum *big.Int) (*types.Transaction, error) {
	return _TellorMesosphere.Contract.UpdateQuorum(&_TellorMesosphere.TransactOpts, _quorum)
}

// TellorMesosphereRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TellorMesosphere contract.
type TellorMesosphereRoleAdminChangedIterator struct {
	Event *TellorMesosphereRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TellorMesosphereRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorMesosphereRoleAdminChanged)
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
		it.Event = new(TellorMesosphereRoleAdminChanged)
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
func (it *TellorMesosphereRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorMesosphereRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorMesosphereRoleAdminChanged represents a RoleAdminChanged event raised by the TellorMesosphere contract.
type TellorMesosphereRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TellorMesosphere *TellorMesosphereFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TellorMesosphereRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TellorMesosphere.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TellorMesosphereRoleAdminChangedIterator{contract: _TellorMesosphere.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TellorMesosphere *TellorMesosphereFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TellorMesosphereRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TellorMesosphere.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorMesosphereRoleAdminChanged)
				if err := _TellorMesosphere.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TellorMesosphere *TellorMesosphereFilterer) ParseRoleAdminChanged(log types.Log) (*TellorMesosphereRoleAdminChanged, error) {
	event := new(TellorMesosphereRoleAdminChanged)
	if err := _TellorMesosphere.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorMesosphereRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TellorMesosphere contract.
type TellorMesosphereRoleGrantedIterator struct {
	Event *TellorMesosphereRoleGranted // Event containing the contract specifics and raw log

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
func (it *TellorMesosphereRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorMesosphereRoleGranted)
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
		it.Event = new(TellorMesosphereRoleGranted)
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
func (it *TellorMesosphereRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorMesosphereRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorMesosphereRoleGranted represents a RoleGranted event raised by the TellorMesosphere contract.
type TellorMesosphereRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TellorMesosphere *TellorMesosphereFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TellorMesosphereRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TellorMesosphere.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TellorMesosphereRoleGrantedIterator{contract: _TellorMesosphere.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TellorMesosphere *TellorMesosphereFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TellorMesosphereRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TellorMesosphere.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorMesosphereRoleGranted)
				if err := _TellorMesosphere.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TellorMesosphere *TellorMesosphereFilterer) ParseRoleGranted(log types.Log) (*TellorMesosphereRoleGranted, error) {
	event := new(TellorMesosphereRoleGranted)
	if err := _TellorMesosphere.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorMesosphereRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TellorMesosphere contract.
type TellorMesosphereRoleRevokedIterator struct {
	Event *TellorMesosphereRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TellorMesosphereRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorMesosphereRoleRevoked)
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
		it.Event = new(TellorMesosphereRoleRevoked)
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
func (it *TellorMesosphereRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorMesosphereRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorMesosphereRoleRevoked represents a RoleRevoked event raised by the TellorMesosphere contract.
type TellorMesosphereRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TellorMesosphere *TellorMesosphereFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TellorMesosphereRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TellorMesosphere.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TellorMesosphereRoleRevokedIterator{contract: _TellorMesosphere.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TellorMesosphere *TellorMesosphereFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TellorMesosphereRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TellorMesosphere.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorMesosphereRoleRevoked)
				if err := _TellorMesosphere.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TellorMesosphere *TellorMesosphereFilterer) ParseRoleRevoked(log types.Log) (*TellorMesosphereRoleRevoked, error) {
	event := new(TellorMesosphereRoleRevoked)
	if err := _TellorMesosphere.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConsoleABI is the input ABI used to generate the binding from.
const ConsoleABI = "[]"

// ConsoleBin is the compiled bytecode used for deploying new contracts.
var ConsoleBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ce9a69e5e4f289d28528ed0c0490f4e7bc58575a5f96b7ad33167a5c49b7308164736f6c63430007000033"

// DeployConsole deploys a new Ethereum contract, binding an instance of Console to it.
func DeployConsole(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Console, error) {
	parsed, err := abi.JSON(strings.NewReader(ConsoleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ConsoleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Console{ConsoleCaller: ConsoleCaller{contract: contract}, ConsoleTransactor: ConsoleTransactor{contract: contract}, ConsoleFilterer: ConsoleFilterer{contract: contract}}, nil
}

// Console is an auto generated Go binding around an Ethereum contract.
type Console struct {
	ConsoleCaller     // Read-only binding to the contract
	ConsoleTransactor // Write-only binding to the contract
	ConsoleFilterer   // Log filterer for contract events
}

// ConsoleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConsoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsoleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConsoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsoleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConsoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsoleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConsoleSession struct {
	Contract     *Console          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConsoleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConsoleCallerSession struct {
	Contract *ConsoleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ConsoleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConsoleTransactorSession struct {
	Contract     *ConsoleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ConsoleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConsoleRaw struct {
	Contract *Console // Generic contract binding to access the raw methods on
}

// ConsoleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConsoleCallerRaw struct {
	Contract *ConsoleCaller // Generic read-only contract binding to access the raw methods on
}

// ConsoleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConsoleTransactorRaw struct {
	Contract *ConsoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConsole creates a new instance of Console, bound to a specific deployed contract.
func NewConsole(address common.Address, backend bind.ContractBackend) (*Console, error) {
	contract, err := bindConsole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Console{ConsoleCaller: ConsoleCaller{contract: contract}, ConsoleTransactor: ConsoleTransactor{contract: contract}, ConsoleFilterer: ConsoleFilterer{contract: contract}}, nil
}

// NewConsoleCaller creates a new read-only instance of Console, bound to a specific deployed contract.
func NewConsoleCaller(address common.Address, caller bind.ContractCaller) (*ConsoleCaller, error) {
	contract, err := bindConsole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConsoleCaller{contract: contract}, nil
}

// NewConsoleTransactor creates a new write-only instance of Console, bound to a specific deployed contract.
func NewConsoleTransactor(address common.Address, transactor bind.ContractTransactor) (*ConsoleTransactor, error) {
	contract, err := bindConsole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConsoleTransactor{contract: contract}, nil
}

// NewConsoleFilterer creates a new log filterer instance of Console, bound to a specific deployed contract.
func NewConsoleFilterer(address common.Address, filterer bind.ContractFilterer) (*ConsoleFilterer, error) {
	contract, err := bindConsole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConsoleFilterer{contract: contract}, nil
}

// bindConsole binds a generic wrapper to an already deployed contract.
func bindConsole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConsoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Console *ConsoleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Console.Contract.ConsoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Console *ConsoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Console.Contract.ConsoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Console *ConsoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Console.Contract.ConsoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Console *ConsoleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Console.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Console *ConsoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Console.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Console *ConsoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Console.Contract.contract.Transact(opts, method, params...)
}
