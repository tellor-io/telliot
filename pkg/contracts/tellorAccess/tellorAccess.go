// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tellorAccess

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
var AddressBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205cc54955cd051d0128d67616b06c3a63905ba90fd2f23b2a06b6bb659e9c28c064736f6c63430007000033"

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
var EnumerableSetBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220dea7b5e18a17ee7345e8494209fb6f96319ea03e2576f99472ec40490a97000864736f6c63430007000033"

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

// IOracleABI is the input ABI used to generate the binding from.
const IOracleABI = "[{\"inputs\":[],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMarketClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isTerminated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceTWAPLong\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"newPrice\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceTWAPShort\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"newPrice\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingAsset\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IOracleFuncSigs maps the 4-byte function signature to its string representation.
var IOracleFuncSigs = map[string]string{
	"d8dfeb45": "collateral()",
	"b7e86c1f": "isMarketClosed()",
	"d1cc9976": "isTerminated()",
	"0e222f9b": "priceTWAPLong()",
	"ccbdbee2": "priceTWAPShort()",
	"7158da7c": "underlyingAsset()",
}

// IOracle is an auto generated Go binding around an Ethereum contract.
type IOracle struct {
	IOracleCaller     // Read-only binding to the contract
	IOracleTransactor // Write-only binding to the contract
	IOracleFilterer   // Log filterer for contract events
}

// IOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOracleSession struct {
	Contract     *IOracle          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOracleCallerSession struct {
	Contract *IOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOracleTransactorSession struct {
	Contract     *IOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOracleRaw struct {
	Contract *IOracle // Generic contract binding to access the raw methods on
}

// IOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOracleCallerRaw struct {
	Contract *IOracleCaller // Generic read-only contract binding to access the raw methods on
}

// IOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOracleTransactorRaw struct {
	Contract *IOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOracle creates a new instance of IOracle, bound to a specific deployed contract.
func NewIOracle(address common.Address, backend bind.ContractBackend) (*IOracle, error) {
	contract, err := bindIOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOracle{IOracleCaller: IOracleCaller{contract: contract}, IOracleTransactor: IOracleTransactor{contract: contract}, IOracleFilterer: IOracleFilterer{contract: contract}}, nil
}

// NewIOracleCaller creates a new read-only instance of IOracle, bound to a specific deployed contract.
func NewIOracleCaller(address common.Address, caller bind.ContractCaller) (*IOracleCaller, error) {
	contract, err := bindIOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOracleCaller{contract: contract}, nil
}

// NewIOracleTransactor creates a new write-only instance of IOracle, bound to a specific deployed contract.
func NewIOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*IOracleTransactor, error) {
	contract, err := bindIOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOracleTransactor{contract: contract}, nil
}

// NewIOracleFilterer creates a new log filterer instance of IOracle, bound to a specific deployed contract.
func NewIOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*IOracleFilterer, error) {
	contract, err := bindIOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOracleFilterer{contract: contract}, nil
}

// bindIOracle binds a generic wrapper to an already deployed contract.
func bindIOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOracle *IOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOracle.Contract.IOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOracle *IOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOracle.Contract.IOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOracle *IOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOracle.Contract.IOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOracle *IOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOracle *IOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOracle *IOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOracle.Contract.contract.Transact(opts, method, params...)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(string)
func (_IOracle *IOracleCaller) Collateral(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IOracle.contract.Call(opts, &out, "collateral")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(string)
func (_IOracle *IOracleSession) Collateral() (string, error) {
	return _IOracle.Contract.Collateral(&_IOracle.CallOpts)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(string)
func (_IOracle *IOracleCallerSession) Collateral() (string, error) {
	return _IOracle.Contract.Collateral(&_IOracle.CallOpts)
}

// UnderlyingAsset is a free data retrieval call binding the contract method 0x7158da7c.
//
// Solidity: function underlyingAsset() view returns(string)
func (_IOracle *IOracleCaller) UnderlyingAsset(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IOracle.contract.Call(opts, &out, "underlyingAsset")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UnderlyingAsset is a free data retrieval call binding the contract method 0x7158da7c.
//
// Solidity: function underlyingAsset() view returns(string)
func (_IOracle *IOracleSession) UnderlyingAsset() (string, error) {
	return _IOracle.Contract.UnderlyingAsset(&_IOracle.CallOpts)
}

// UnderlyingAsset is a free data retrieval call binding the contract method 0x7158da7c.
//
// Solidity: function underlyingAsset() view returns(string)
func (_IOracle *IOracleCallerSession) UnderlyingAsset() (string, error) {
	return _IOracle.Contract.UnderlyingAsset(&_IOracle.CallOpts)
}

// IsMarketClosed is a paid mutator transaction binding the contract method 0xb7e86c1f.
//
// Solidity: function isMarketClosed() returns(bool)
func (_IOracle *IOracleTransactor) IsMarketClosed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOracle.contract.Transact(opts, "isMarketClosed")
}

// IsMarketClosed is a paid mutator transaction binding the contract method 0xb7e86c1f.
//
// Solidity: function isMarketClosed() returns(bool)
func (_IOracle *IOracleSession) IsMarketClosed() (*types.Transaction, error) {
	return _IOracle.Contract.IsMarketClosed(&_IOracle.TransactOpts)
}

// IsMarketClosed is a paid mutator transaction binding the contract method 0xb7e86c1f.
//
// Solidity: function isMarketClosed() returns(bool)
func (_IOracle *IOracleTransactorSession) IsMarketClosed() (*types.Transaction, error) {
	return _IOracle.Contract.IsMarketClosed(&_IOracle.TransactOpts)
}

// IsTerminated is a paid mutator transaction binding the contract method 0xd1cc9976.
//
// Solidity: function isTerminated() returns(bool)
func (_IOracle *IOracleTransactor) IsTerminated(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOracle.contract.Transact(opts, "isTerminated")
}

// IsTerminated is a paid mutator transaction binding the contract method 0xd1cc9976.
//
// Solidity: function isTerminated() returns(bool)
func (_IOracle *IOracleSession) IsTerminated() (*types.Transaction, error) {
	return _IOracle.Contract.IsTerminated(&_IOracle.TransactOpts)
}

// IsTerminated is a paid mutator transaction binding the contract method 0xd1cc9976.
//
// Solidity: function isTerminated() returns(bool)
func (_IOracle *IOracleTransactorSession) IsTerminated() (*types.Transaction, error) {
	return _IOracle.Contract.IsTerminated(&_IOracle.TransactOpts)
}

// PriceTWAPLong is a paid mutator transaction binding the contract method 0x0e222f9b.
//
// Solidity: function priceTWAPLong() returns(int256 newPrice, uint256 newTimestamp)
func (_IOracle *IOracleTransactor) PriceTWAPLong(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOracle.contract.Transact(opts, "priceTWAPLong")
}

// PriceTWAPLong is a paid mutator transaction binding the contract method 0x0e222f9b.
//
// Solidity: function priceTWAPLong() returns(int256 newPrice, uint256 newTimestamp)
func (_IOracle *IOracleSession) PriceTWAPLong() (*types.Transaction, error) {
	return _IOracle.Contract.PriceTWAPLong(&_IOracle.TransactOpts)
}

// PriceTWAPLong is a paid mutator transaction binding the contract method 0x0e222f9b.
//
// Solidity: function priceTWAPLong() returns(int256 newPrice, uint256 newTimestamp)
func (_IOracle *IOracleTransactorSession) PriceTWAPLong() (*types.Transaction, error) {
	return _IOracle.Contract.PriceTWAPLong(&_IOracle.TransactOpts)
}

// PriceTWAPShort is a paid mutator transaction binding the contract method 0xccbdbee2.
//
// Solidity: function priceTWAPShort() returns(int256 newPrice, uint256 newTimestamp)
func (_IOracle *IOracleTransactor) PriceTWAPShort(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOracle.contract.Transact(opts, "priceTWAPShort")
}

// PriceTWAPShort is a paid mutator transaction binding the contract method 0xccbdbee2.
//
// Solidity: function priceTWAPShort() returns(int256 newPrice, uint256 newTimestamp)
func (_IOracle *IOracleSession) PriceTWAPShort() (*types.Transaction, error) {
	return _IOracle.Contract.PriceTWAPShort(&_IOracle.TransactOpts)
}

// PriceTWAPShort is a paid mutator transaction binding the contract method 0xccbdbee2.
//
// Solidity: function priceTWAPShort() returns(int256 newPrice, uint256 newTimestamp)
func (_IOracle *IOracleTransactorSession) PriceTWAPShort() (*types.Transaction, error) {
	return _IOracle.Contract.PriceTWAPShort(&_IOracle.TransactOpts)
}

// OracleABI is the input ABI used to generate the binding from.
const OracleABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"collateral_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"underlyingAsset_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_markRequestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_indexRequestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"NewValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPORTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin_address\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reporter_address\",\"type\":\"address\"}],\"name\":\"addReporter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getCurrentValue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getDataBefore\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"indexRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin_address\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMarketClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reporter_address\",\"type\":\"address\"}],\"name\":\"isReporter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isTerminated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"markRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceTWAPLong\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"newPrice\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceTWAPShort\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"newPrice\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reporter_address\",\"type\":\"address\"}],\"name\":\"removeReporter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"submitValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"terminate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"timestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingAsset\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"values\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OracleFuncSigs maps the 4-byte function signature to its string representation.
var OracleFuncSigs = map[string]string{
	"a217fddf": "DEFAULT_ADMIN_ROLE()",
	"3f60d799": "REPORTER_ROLE()",
	"70480275": "addAdmin(address)",
	"dd8755f2": "addReporter(address)",
	"d8dfeb45": "collateral()",
	"3fcad964": "getCurrentValue(uint256)",
	"66b44611": "getDataBefore(uint256,uint256)",
	"46eee1c4": "getNewValueCountbyRequestId(uint256)",
	"248a9ca3": "getRoleAdmin(bytes32)",
	"9010d07c": "getRoleMember(bytes32,uint256)",
	"ca15c873": "getRoleMemberCount(bytes32)",
	"77fbb663": "getTimestampbyRequestIDandIndex(uint256,uint256)",
	"2f2ff15d": "grantRole(bytes32,address)",
	"91d14854": "hasRole(bytes32,address)",
	"38af876d": "indexRequestId()",
	"24d7806c": "isAdmin(address)",
	"b7e86c1f": "isMarketClosed()",
	"044ad7be": "isReporter(address)",
	"d1cc9976": "isTerminated()",
	"da47c199": "markRequestId()",
	"0e222f9b": "priceTWAPLong()",
	"ccbdbee2": "priceTWAPShort()",
	"5de5c212": "removeReporter(address)",
	"8bad0c0a": "renounceAdmin()",
	"36568abe": "renounceRole(bytes32,address)",
	"93fa4915": "retrieveData(uint256,uint256)",
	"d547741f": "revokeRole(bytes32,address)",
	"62f55112": "submitValue(uint256,uint256)",
	"0c08bf88": "terminate()",
	"fb0ceb04": "timestamps(uint256,uint256)",
	"7158da7c": "underlyingAsset()",
	"a3183701": "values(uint256,uint256)",
}

// OracleBin is the compiled bytecode used for deploying new contracts.
var OracleBin = "0x60806040523480156200001157600080fd5b50604051620016cd380380620016cd833981810160405260808110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010a57600080fd5b9083019060208201858111156200012057600080fd5b82516401000000008111828201881017156200013b57600080fd5b82525081516020918201929091019080838360005b838110156200016a57818101518382015260200162000150565b50505050905090810190601f168015620001985780820380516001836020036101000a031916815260200191505b50604090815260208201519101519092509050620001b860003362000223565b620001e57f176c2b761bfeb5dab89f614b6c08152e31d9230394b3605eabf32249ea1c89a6600062000233565b8351620001fa90600390602087019062000389565b5082516200021090600490602086019062000389565b5060069190915560075550620004259050565b6200022f828262000285565b5050565b600082815260208190526040808220600201549051839285917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a460009182526020829052604090912060020155565b600082815260208181526040909120620002aa91839062000d9c620002fe821b17901c565b156200022f57620002ba6200031e565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b600062000315836001600160a01b03841662000322565b90505b92915050565b3390565b600062000330838362000371565b620003685750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000318565b50600062000318565b60009081526001919091016020526040902054151590565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003cc57805160ff1916838001178555620003fc565b82800160010185558215620003fc579182015b82811115620003fc578251825591602001919060010190620003df565b506200040a9291506200040e565b5090565b5b808211156200040a57600081556001016200040f565b61129880620004356000396000f3fe608060405234801561001057600080fd5b50600436106101f05760003560e01c806377fbb6631161010f578063ca15c873116100a2578063d8dfeb4511610071578063d8dfeb45146105c5578063da47c199146105cd578063dd8755f2146105d5578063fb0ceb04146105fb576101f0565b8063ca15c8731461056c578063ccbdbee214610589578063d1cc997614610591578063d547741f14610599576101f0565b806393fa4915116100de57806393fa491514610516578063a217fddf14610539578063a318370114610541578063b7e86c1f14610564576101f0565b806377fbb663146104805780638bad0c0a146104a35780639010d07c146104ab57806391d14854146104ea576101f0565b80633f60d7991161018757806362f551121161015657806362f551121461039757806366b44611146103ba57806370480275146103dd5780637158da7c14610403576101f0565b80633f60d7991461030f5780633fcad9641461031757806346eee1c4146103545780635de5c21214610371576101f0565b806324d7806c116101c357806324d7806c146102895780632f2ff15d146102af57806336568abe146102db57806338af876d14610307576101f0565b8063044ad7be146101f55780630c08bf881461022f5780630e222f9b14610239578063248a9ca31461025a575b600080fd5b61021b6004803603602081101561020b57600080fd5b50356001600160a01b031661061e565b604080519115158252519081900360200190f35b61023761063e565b005b61024161069f565b6040805192835260208301919091528051918290030190f35b6102776004803603602081101561027057600080fd5b50356106f2565b60408051918252519081900360200190f35b61021b6004803603602081101561029f57600080fd5b50356001600160a01b0316610707565b610237600480360360408110156102c557600080fd5b50803590602001356001600160a01b0316610713565b610237600480360360408110156102f157600080fd5b50803590602001356001600160a01b031661077f565b6102776107e0565b6102776107e6565b6103346004803603602081101561032d57600080fd5b50356107f8565b604080519315158452602084019290925282820152519081900360600190f35b6102776004803603602081101561036a57600080fd5b5035610850565b6102376004803603602081101561038757600080fd5b50356001600160a01b0316610862565b610237600480360360408110156103ad57600080fd5b50803590602001356108cf565b610334600480360360408110156103d057600080fd5b508035906020013561099e565b610237600480360360208110156103f357600080fd5b50356001600160a01b0316610a14565b61040b610a71565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561044557818101518382015260200161042d565b50505050905090810190601f1680156104725780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102776004803603604081101561049657600080fd5b5080359060200135610b07565b610237610b60565b6104ce600480360360408110156104c157600080fd5b5080359060200135610b6d565b604080516001600160a01b039092168252519081900360200190f35b61021b6004803603604081101561050057600080fd5b50803590602001356001600160a01b0316610b8c565b6102776004803603604081101561052c57600080fd5b5080359060200135610ba4565b610277610bc1565b6102776004803603604081101561055757600080fd5b5080359060200135610bc6565b61021b610be3565b6102776004803603602081101561058257600080fd5b5035610bf1565b610241610c08565b61021b610c3b565b610237600480360360408110156105af57600080fd5b50803590602001356001600160a01b0316610c44565b61040b610c9d565b610277610cfe565b610237600480360360208110156105eb57600080fd5b50356001600160a01b0316610d04565b6102776004803603604081101561061157600080fd5b5080359060200135610d6e565b60006106386000805160206111b283398151915283610b8c565b92915050565b61064733610707565b610690576040805162461bcd60e51b81526020600482015260156024820152742932b9ba3934b1ba32b2103a379030b236b4b7399760591b604482015290519081900360640190fd5b6005805460ff19166001179055565b60008060006106af600654610850565b905060006106c260065460018403610b07565b905060006106d260065483610ba4565b905080156106e557935091506106ee9050565b50600093509150505b9091565b60009081526020819052604090206002015490565b60006106388183610b8c565b60008281526020819052604090206002015461073690610731610db1565b610b8c565b6107715760405162461bcd60e51b815260040180806020018281038252602f815260200180611183602f913960400191505060405180910390fd5b61077b8282610db5565b5050565b610787610db1565b6001600160a01b0316816001600160a01b0316146107d65760405162461bcd60e51b815260040180806020018281038252602f815260200180611234602f913960400191505060405180910390fd5b61077b8282610e1e565b60075481565b6000805160206111b283398151915281565b60008060008061080785610850565b905060006108188660018403610b07565b905060006108268783610ba4565b9050801561083d5760019550935091506108499050565b50600094508493509150505b9193909250565b60009081526002602052604090205490565b61086b33610707565b6108b4576040805162461bcd60e51b81526020600482015260156024820152742932b9ba3934b1ba32b2103a379030b236b4b7399760591b604482015290519081900360640190fd5b6108cc6000805160206111b283398151915282610c44565b50565b6108d83361061e565b806108e757506108e733610707565b6109225760405162461bcd60e51b81526004018080602001828103825260328152602001806111d26032913960400191505060405180910390fd5b6000828152600160208181526040808420428086529083528185208690558685526002835281852080549485018155855293829020909201839055815185815290810192909252818101839052517fba11e319aee26e7bbac889432515ba301ec8f6d27bf6b94829c21a65c5f6ff259181900360600190a15050565b60008060008060006109b08787610e87565b91509150816109cb5760008060009450945094505050610a0d565b60006109d78883610b07565b905060006109e58983610ba4565b905080156109fd576001965094509250610a0d915050565b6000806000965096509650505050505b9250925092565b610a1d33610707565b610a66576040805162461bcd60e51b81526020600482015260156024820152742932b9ba3934b1ba32b2103a379030b236b4b7399760591b604482015290519081900360640190fd5b6108cc600082610713565b60048054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610afd5780601f10610ad257610100808354040283529160200191610afd565b820191906000526020600020905b815481529060010190602001808311610ae057829003601f168201915b5050505050905090565b600082815260026020526040812054801580610b235750828111155b15610b32576000915050610638565b6000848152600260205260409020805484908110610b4c57fe5b906000526020600020015491505092915050565b610b6b60003361077f565b565b6000828152602081905260408120610b859083610f8f565b9392505050565b6000828152602081905260408120610b859083610f9b565b600091825260016020908152604080842092845291905290205490565b600081565b600160209081526000928352604080842090915290825290205481565b600554610100900460ff1690565b600081815260208190526040812061063890610fb0565b6000806000610c18600754610850565b90506000610c2b60075460018403610b07565b905060006106d260075483610ba4565b60055460ff1690565b600082815260208190526040902060020154610c6290610731610db1565b6107d65760405162461bcd60e51b81526004018080602001828103825260308152602001806112046030913960400191505060405180910390fd5b60038054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610afd5780601f10610ad257610100808354040283529160200191610afd565b60065481565b610d0d33610707565b610d56576040805162461bcd60e51b81526020600482015260156024820152742932b9ba3934b1ba32b2103a379030b236b4b7399760591b604482015290519081900360640190fd5b6108cc6000805160206111b283398151915282610713565b60026020528160005260406000208181548110610d8757fe5b90600052602060002001600091509150505481565b6000610b85836001600160a01b038416610fbb565b3390565b6000828152602081905260409020610dcd9082610d9c565b1561077b57610dda610db1565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6000828152602081905260409020610e369082611005565b1561077b57610e43610db1565b6001600160a01b0316816001600160a01b0316837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45050565b6000806000610e9585610850565b90508015610f7f57600080600019830181610eb08982610b07565b9050878110610eca57600080965096505050505050610f88565b610ed48983610b07565b905087811015610eee5750600195509350610f8892505050565b82600281840304600101019350610f058985610b07565b905087811015610f45576000610f1e8a86600101610b07565b9050888110610f395760018597509750505050505050610f88565b84600101935050610f7a565b6000610f548a60018703610b07565b905088811015610f7257600180860397509750505050505050610f88565b600185039250505b610eee565b60008092509250505b9250929050565b6000610b85838361101a565b6000610b85836001600160a01b03841661107e565b600061063882611096565b6000610fc7838361107e565b610ffd57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610638565b506000610638565b6000610b85836001600160a01b03841661109a565b8154600090821061105c5760405162461bcd60e51b81526004018080602001828103825260228152602001806111616022913960400191505060405180910390fd5b82600001828154811061106b57fe5b9060005260206000200154905092915050565b60009081526001919091016020526040902054151590565b5490565b6000818152600183016020526040812054801561115657835460001980830191908101906000908790839081106110cd57fe5b90600052602060002001549050808760000184815481106110ea57fe5b60009182526020808320909101929092558281526001898101909252604090209084019055865487908061111a57fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050610638565b600091505061063856fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e6473416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f206772616e74176c2b761bfeb5dab89f614b6c08152e31d9230394b3605eabf32249ea1c89a653656e646572206d75737420626520616e2041646d696e206f72205265706f7274657220746f207375626d697456616c7565416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f207265766f6b65416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636520726f6c657320666f722073656c66a26469706673582212200bb95b2fbe807b011a48e1d48afcf4ab3c9291abf033c760a9b5b58c067a8d3364736f6c63430007000033"

// DeployOracle deploys a new Ethereum contract, binding an instance of Oracle to it.
func DeployOracle(auth *bind.TransactOpts, backend bind.ContractBackend, collateral_ string, underlyingAsset_ string, _markRequestId *big.Int, _indexRequestId *big.Int) (common.Address, *types.Transaction, *Oracle, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OracleBin), backend, collateral_, underlyingAsset_, _markRequestId, _indexRequestId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Oracle *OracleCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Oracle *OracleSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Oracle.Contract.DEFAULTADMINROLE(&_Oracle.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Oracle *OracleCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Oracle.Contract.DEFAULTADMINROLE(&_Oracle.CallOpts)
}

// REPORTERROLE is a free data retrieval call binding the contract method 0x3f60d799.
//
// Solidity: function REPORTER_ROLE() view returns(bytes32)
func (_Oracle *OracleCaller) REPORTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "REPORTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REPORTERROLE is a free data retrieval call binding the contract method 0x3f60d799.
//
// Solidity: function REPORTER_ROLE() view returns(bytes32)
func (_Oracle *OracleSession) REPORTERROLE() ([32]byte, error) {
	return _Oracle.Contract.REPORTERROLE(&_Oracle.CallOpts)
}

// REPORTERROLE is a free data retrieval call binding the contract method 0x3f60d799.
//
// Solidity: function REPORTER_ROLE() view returns(bytes32)
func (_Oracle *OracleCallerSession) REPORTERROLE() ([32]byte, error) {
	return _Oracle.Contract.REPORTERROLE(&_Oracle.CallOpts)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(string)
func (_Oracle *OracleCaller) Collateral(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "collateral")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(string)
func (_Oracle *OracleSession) Collateral() (string, error) {
	return _Oracle.Contract.Collateral(&_Oracle.CallOpts)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(string)
func (_Oracle *OracleCallerSession) Collateral() (string, error) {
	return _Oracle.Contract.Collateral(&_Oracle.CallOpts)
}

// GetCurrentValue is a free data retrieval call binding the contract method 0x3fcad964.
//
// Solidity: function getCurrentValue(uint256 _requestId) view returns(bool, uint256, uint256)
func (_Oracle *OracleCaller) GetCurrentValue(opts *bind.CallOpts, _requestId *big.Int) (bool, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getCurrentValue", _requestId)

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
func (_Oracle *OracleSession) GetCurrentValue(_requestId *big.Int) (bool, *big.Int, *big.Int, error) {
	return _Oracle.Contract.GetCurrentValue(&_Oracle.CallOpts, _requestId)
}

// GetCurrentValue is a free data retrieval call binding the contract method 0x3fcad964.
//
// Solidity: function getCurrentValue(uint256 _requestId) view returns(bool, uint256, uint256)
func (_Oracle *OracleCallerSession) GetCurrentValue(_requestId *big.Int) (bool, *big.Int, *big.Int, error) {
	return _Oracle.Contract.GetCurrentValue(&_Oracle.CallOpts, _requestId)
}

// GetDataBefore is a free data retrieval call binding the contract method 0x66b44611.
//
// Solidity: function getDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool, uint256, uint256)
func (_Oracle *OracleCaller) GetDataBefore(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getDataBefore", _requestId, _timestamp)

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
func (_Oracle *OracleSession) GetDataBefore(_requestId *big.Int, _timestamp *big.Int) (bool, *big.Int, *big.Int, error) {
	return _Oracle.Contract.GetDataBefore(&_Oracle.CallOpts, _requestId, _timestamp)
}

// GetDataBefore is a free data retrieval call binding the contract method 0x66b44611.
//
// Solidity: function getDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool, uint256, uint256)
func (_Oracle *OracleCallerSession) GetDataBefore(_requestId *big.Int, _timestamp *big.Int) (bool, *big.Int, *big.Int, error) {
	return _Oracle.Contract.GetDataBefore(&_Oracle.CallOpts, _requestId, _timestamp)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_Oracle *OracleCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getNewValueCountbyRequestId", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_Oracle *OracleSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _Oracle.Contract.GetNewValueCountbyRequestId(&_Oracle.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_Oracle *OracleCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _Oracle.Contract.GetNewValueCountbyRequestId(&_Oracle.CallOpts, _requestId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Oracle *OracleCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Oracle *OracleSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Oracle.Contract.GetRoleAdmin(&_Oracle.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Oracle *OracleCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Oracle.Contract.GetRoleAdmin(&_Oracle.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Oracle *OracleCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Oracle *OracleSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Oracle.Contract.GetRoleMember(&_Oracle.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Oracle *OracleCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Oracle.Contract.GetRoleMember(&_Oracle.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Oracle *OracleCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Oracle *OracleSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Oracle.Contract.GetRoleMemberCount(&_Oracle.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Oracle *OracleCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Oracle.Contract.GetRoleMemberCount(&_Oracle.CallOpts, role)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 _index) view returns(uint256)
func (_Oracle *OracleCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestId *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getTimestampbyRequestIDandIndex", _requestId, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 _index) view returns(uint256)
func (_Oracle *OracleSession) GetTimestampbyRequestIDandIndex(_requestId *big.Int, _index *big.Int) (*big.Int, error) {
	return _Oracle.Contract.GetTimestampbyRequestIDandIndex(&_Oracle.CallOpts, _requestId, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 _index) view returns(uint256)
func (_Oracle *OracleCallerSession) GetTimestampbyRequestIDandIndex(_requestId *big.Int, _index *big.Int) (*big.Int, error) {
	return _Oracle.Contract.GetTimestampbyRequestIDandIndex(&_Oracle.CallOpts, _requestId, _index)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Oracle *OracleCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Oracle *OracleSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Oracle.Contract.HasRole(&_Oracle.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Oracle *OracleCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Oracle.Contract.HasRole(&_Oracle.CallOpts, role, account)
}

// IndexRequestId is a free data retrieval call binding the contract method 0x38af876d.
//
// Solidity: function indexRequestId() view returns(uint256)
func (_Oracle *OracleCaller) IndexRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "indexRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IndexRequestId is a free data retrieval call binding the contract method 0x38af876d.
//
// Solidity: function indexRequestId() view returns(uint256)
func (_Oracle *OracleSession) IndexRequestId() (*big.Int, error) {
	return _Oracle.Contract.IndexRequestId(&_Oracle.CallOpts)
}

// IndexRequestId is a free data retrieval call binding the contract method 0x38af876d.
//
// Solidity: function indexRequestId() view returns(uint256)
func (_Oracle *OracleCallerSession) IndexRequestId() (*big.Int, error) {
	return _Oracle.Contract.IndexRequestId(&_Oracle.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _admin_address) view returns(bool)
func (_Oracle *OracleCaller) IsAdmin(opts *bind.CallOpts, _admin_address common.Address) (bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "isAdmin", _admin_address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _admin_address) view returns(bool)
func (_Oracle *OracleSession) IsAdmin(_admin_address common.Address) (bool, error) {
	return _Oracle.Contract.IsAdmin(&_Oracle.CallOpts, _admin_address)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _admin_address) view returns(bool)
func (_Oracle *OracleCallerSession) IsAdmin(_admin_address common.Address) (bool, error) {
	return _Oracle.Contract.IsAdmin(&_Oracle.CallOpts, _admin_address)
}

// IsMarketClosed is a free data retrieval call binding the contract method 0xb7e86c1f.
//
// Solidity: function isMarketClosed() view returns(bool)
func (_Oracle *OracleCaller) IsMarketClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "isMarketClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMarketClosed is a free data retrieval call binding the contract method 0xb7e86c1f.
//
// Solidity: function isMarketClosed() view returns(bool)
func (_Oracle *OracleSession) IsMarketClosed() (bool, error) {
	return _Oracle.Contract.IsMarketClosed(&_Oracle.CallOpts)
}

// IsMarketClosed is a free data retrieval call binding the contract method 0xb7e86c1f.
//
// Solidity: function isMarketClosed() view returns(bool)
func (_Oracle *OracleCallerSession) IsMarketClosed() (bool, error) {
	return _Oracle.Contract.IsMarketClosed(&_Oracle.CallOpts)
}

// IsReporter is a free data retrieval call binding the contract method 0x044ad7be.
//
// Solidity: function isReporter(address _reporter_address) view returns(bool)
func (_Oracle *OracleCaller) IsReporter(opts *bind.CallOpts, _reporter_address common.Address) (bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "isReporter", _reporter_address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReporter is a free data retrieval call binding the contract method 0x044ad7be.
//
// Solidity: function isReporter(address _reporter_address) view returns(bool)
func (_Oracle *OracleSession) IsReporter(_reporter_address common.Address) (bool, error) {
	return _Oracle.Contract.IsReporter(&_Oracle.CallOpts, _reporter_address)
}

// IsReporter is a free data retrieval call binding the contract method 0x044ad7be.
//
// Solidity: function isReporter(address _reporter_address) view returns(bool)
func (_Oracle *OracleCallerSession) IsReporter(_reporter_address common.Address) (bool, error) {
	return _Oracle.Contract.IsReporter(&_Oracle.CallOpts, _reporter_address)
}

// IsTerminated is a free data retrieval call binding the contract method 0xd1cc9976.
//
// Solidity: function isTerminated() view returns(bool)
func (_Oracle *OracleCaller) IsTerminated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "isTerminated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTerminated is a free data retrieval call binding the contract method 0xd1cc9976.
//
// Solidity: function isTerminated() view returns(bool)
func (_Oracle *OracleSession) IsTerminated() (bool, error) {
	return _Oracle.Contract.IsTerminated(&_Oracle.CallOpts)
}

// IsTerminated is a free data retrieval call binding the contract method 0xd1cc9976.
//
// Solidity: function isTerminated() view returns(bool)
func (_Oracle *OracleCallerSession) IsTerminated() (bool, error) {
	return _Oracle.Contract.IsTerminated(&_Oracle.CallOpts)
}

// MarkRequestId is a free data retrieval call binding the contract method 0xda47c199.
//
// Solidity: function markRequestId() view returns(uint256)
func (_Oracle *OracleCaller) MarkRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "markRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarkRequestId is a free data retrieval call binding the contract method 0xda47c199.
//
// Solidity: function markRequestId() view returns(uint256)
func (_Oracle *OracleSession) MarkRequestId() (*big.Int, error) {
	return _Oracle.Contract.MarkRequestId(&_Oracle.CallOpts)
}

// MarkRequestId is a free data retrieval call binding the contract method 0xda47c199.
//
// Solidity: function markRequestId() view returns(uint256)
func (_Oracle *OracleCallerSession) MarkRequestId() (*big.Int, error) {
	return _Oracle.Contract.MarkRequestId(&_Oracle.CallOpts)
}

// PriceTWAPLong is a free data retrieval call binding the contract method 0x0e222f9b.
//
// Solidity: function priceTWAPLong() view returns(int256 newPrice, uint256 newTimestamp)
func (_Oracle *OracleCaller) PriceTWAPLong(opts *bind.CallOpts) (struct {
	NewPrice     *big.Int
	NewTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "priceTWAPLong")

	outstruct := new(struct {
		NewPrice     *big.Int
		NewTimestamp *big.Int
	})

	outstruct.NewPrice = out[0].(*big.Int)
	outstruct.NewTimestamp = out[1].(*big.Int)

	return *outstruct, err

}

// PriceTWAPLong is a free data retrieval call binding the contract method 0x0e222f9b.
//
// Solidity: function priceTWAPLong() view returns(int256 newPrice, uint256 newTimestamp)
func (_Oracle *OracleSession) PriceTWAPLong() (struct {
	NewPrice     *big.Int
	NewTimestamp *big.Int
}, error) {
	return _Oracle.Contract.PriceTWAPLong(&_Oracle.CallOpts)
}

// PriceTWAPLong is a free data retrieval call binding the contract method 0x0e222f9b.
//
// Solidity: function priceTWAPLong() view returns(int256 newPrice, uint256 newTimestamp)
func (_Oracle *OracleCallerSession) PriceTWAPLong() (struct {
	NewPrice     *big.Int
	NewTimestamp *big.Int
}, error) {
	return _Oracle.Contract.PriceTWAPLong(&_Oracle.CallOpts)
}

// PriceTWAPShort is a free data retrieval call binding the contract method 0xccbdbee2.
//
// Solidity: function priceTWAPShort() view returns(int256 newPrice, uint256 newTimestamp)
func (_Oracle *OracleCaller) PriceTWAPShort(opts *bind.CallOpts) (struct {
	NewPrice     *big.Int
	NewTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "priceTWAPShort")

	outstruct := new(struct {
		NewPrice     *big.Int
		NewTimestamp *big.Int
	})

	outstruct.NewPrice = out[0].(*big.Int)
	outstruct.NewTimestamp = out[1].(*big.Int)

	return *outstruct, err

}

// PriceTWAPShort is a free data retrieval call binding the contract method 0xccbdbee2.
//
// Solidity: function priceTWAPShort() view returns(int256 newPrice, uint256 newTimestamp)
func (_Oracle *OracleSession) PriceTWAPShort() (struct {
	NewPrice     *big.Int
	NewTimestamp *big.Int
}, error) {
	return _Oracle.Contract.PriceTWAPShort(&_Oracle.CallOpts)
}

// PriceTWAPShort is a free data retrieval call binding the contract method 0xccbdbee2.
//
// Solidity: function priceTWAPShort() view returns(int256 newPrice, uint256 newTimestamp)
func (_Oracle *OracleCallerSession) PriceTWAPShort() (struct {
	NewPrice     *big.Int
	NewTimestamp *big.Int
}, error) {
	return _Oracle.Contract.PriceTWAPShort(&_Oracle.CallOpts)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_Oracle *OracleCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "retrieveData", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_Oracle *OracleSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _Oracle.Contract.RetrieveData(&_Oracle.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_Oracle *OracleCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _Oracle.Contract.RetrieveData(&_Oracle.CallOpts, _requestId, _timestamp)
}

// Timestamps is a free data retrieval call binding the contract method 0xfb0ceb04.
//
// Solidity: function timestamps(uint256 , uint256 ) view returns(uint256)
func (_Oracle *OracleCaller) Timestamps(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "timestamps", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timestamps is a free data retrieval call binding the contract method 0xfb0ceb04.
//
// Solidity: function timestamps(uint256 , uint256 ) view returns(uint256)
func (_Oracle *OracleSession) Timestamps(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Oracle.Contract.Timestamps(&_Oracle.CallOpts, arg0, arg1)
}

// Timestamps is a free data retrieval call binding the contract method 0xfb0ceb04.
//
// Solidity: function timestamps(uint256 , uint256 ) view returns(uint256)
func (_Oracle *OracleCallerSession) Timestamps(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Oracle.Contract.Timestamps(&_Oracle.CallOpts, arg0, arg1)
}

// UnderlyingAsset is a free data retrieval call binding the contract method 0x7158da7c.
//
// Solidity: function underlyingAsset() view returns(string)
func (_Oracle *OracleCaller) UnderlyingAsset(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "underlyingAsset")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UnderlyingAsset is a free data retrieval call binding the contract method 0x7158da7c.
//
// Solidity: function underlyingAsset() view returns(string)
func (_Oracle *OracleSession) UnderlyingAsset() (string, error) {
	return _Oracle.Contract.UnderlyingAsset(&_Oracle.CallOpts)
}

// UnderlyingAsset is a free data retrieval call binding the contract method 0x7158da7c.
//
// Solidity: function underlyingAsset() view returns(string)
func (_Oracle *OracleCallerSession) UnderlyingAsset() (string, error) {
	return _Oracle.Contract.UnderlyingAsset(&_Oracle.CallOpts)
}

// Values is a free data retrieval call binding the contract method 0xa3183701.
//
// Solidity: function values(uint256 , uint256 ) view returns(uint256)
func (_Oracle *OracleCaller) Values(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "values", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Values is a free data retrieval call binding the contract method 0xa3183701.
//
// Solidity: function values(uint256 , uint256 ) view returns(uint256)
func (_Oracle *OracleSession) Values(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Oracle.Contract.Values(&_Oracle.CallOpts, arg0, arg1)
}

// Values is a free data retrieval call binding the contract method 0xa3183701.
//
// Solidity: function values(uint256 , uint256 ) view returns(uint256)
func (_Oracle *OracleCallerSession) Values(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Oracle.Contract.Values(&_Oracle.CallOpts, arg0, arg1)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin_address) returns()
func (_Oracle *OracleTransactor) AddAdmin(opts *bind.TransactOpts, _admin_address common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "addAdmin", _admin_address)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin_address) returns()
func (_Oracle *OracleSession) AddAdmin(_admin_address common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.AddAdmin(&_Oracle.TransactOpts, _admin_address)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin_address) returns()
func (_Oracle *OracleTransactorSession) AddAdmin(_admin_address common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.AddAdmin(&_Oracle.TransactOpts, _admin_address)
}

// AddReporter is a paid mutator transaction binding the contract method 0xdd8755f2.
//
// Solidity: function addReporter(address _reporter_address) returns()
func (_Oracle *OracleTransactor) AddReporter(opts *bind.TransactOpts, _reporter_address common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "addReporter", _reporter_address)
}

// AddReporter is a paid mutator transaction binding the contract method 0xdd8755f2.
//
// Solidity: function addReporter(address _reporter_address) returns()
func (_Oracle *OracleSession) AddReporter(_reporter_address common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.AddReporter(&_Oracle.TransactOpts, _reporter_address)
}

// AddReporter is a paid mutator transaction binding the contract method 0xdd8755f2.
//
// Solidity: function addReporter(address _reporter_address) returns()
func (_Oracle *OracleTransactorSession) AddReporter(_reporter_address common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.AddReporter(&_Oracle.TransactOpts, _reporter_address)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Oracle *OracleTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Oracle *OracleSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.GrantRole(&_Oracle.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Oracle *OracleTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.GrantRole(&_Oracle.TransactOpts, role, account)
}

// RemoveReporter is a paid mutator transaction binding the contract method 0x5de5c212.
//
// Solidity: function removeReporter(address _reporter_address) returns()
func (_Oracle *OracleTransactor) RemoveReporter(opts *bind.TransactOpts, _reporter_address common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "removeReporter", _reporter_address)
}

// RemoveReporter is a paid mutator transaction binding the contract method 0x5de5c212.
//
// Solidity: function removeReporter(address _reporter_address) returns()
func (_Oracle *OracleSession) RemoveReporter(_reporter_address common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.RemoveReporter(&_Oracle.TransactOpts, _reporter_address)
}

// RemoveReporter is a paid mutator transaction binding the contract method 0x5de5c212.
//
// Solidity: function removeReporter(address _reporter_address) returns()
func (_Oracle *OracleTransactorSession) RemoveReporter(_reporter_address common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.RemoveReporter(&_Oracle.TransactOpts, _reporter_address)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_Oracle *OracleTransactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_Oracle *OracleSession) RenounceAdmin() (*types.Transaction, error) {
	return _Oracle.Contract.RenounceAdmin(&_Oracle.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_Oracle *OracleTransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _Oracle.Contract.RenounceAdmin(&_Oracle.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Oracle *OracleTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Oracle *OracleSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.RenounceRole(&_Oracle.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Oracle *OracleTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.RenounceRole(&_Oracle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Oracle *OracleTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Oracle *OracleSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.RevokeRole(&_Oracle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Oracle *OracleTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.RevokeRole(&_Oracle.TransactOpts, role, account)
}

// SubmitValue is a paid mutator transaction binding the contract method 0x62f55112.
//
// Solidity: function submitValue(uint256 _requestId, uint256 _value) returns()
func (_Oracle *OracleTransactor) SubmitValue(opts *bind.TransactOpts, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "submitValue", _requestId, _value)
}

// SubmitValue is a paid mutator transaction binding the contract method 0x62f55112.
//
// Solidity: function submitValue(uint256 _requestId, uint256 _value) returns()
func (_Oracle *OracleSession) SubmitValue(_requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SubmitValue(&_Oracle.TransactOpts, _requestId, _value)
}

// SubmitValue is a paid mutator transaction binding the contract method 0x62f55112.
//
// Solidity: function submitValue(uint256 _requestId, uint256 _value) returns()
func (_Oracle *OracleTransactorSession) SubmitValue(_requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SubmitValue(&_Oracle.TransactOpts, _requestId, _value)
}

// Terminate is a paid mutator transaction binding the contract method 0x0c08bf88.
//
// Solidity: function terminate() returns()
func (_Oracle *OracleTransactor) Terminate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "terminate")
}

// Terminate is a paid mutator transaction binding the contract method 0x0c08bf88.
//
// Solidity: function terminate() returns()
func (_Oracle *OracleSession) Terminate() (*types.Transaction, error) {
	return _Oracle.Contract.Terminate(&_Oracle.TransactOpts)
}

// Terminate is a paid mutator transaction binding the contract method 0x0c08bf88.
//
// Solidity: function terminate() returns()
func (_Oracle *OracleTransactorSession) Terminate() (*types.Transaction, error) {
	return _Oracle.Contract.Terminate(&_Oracle.TransactOpts)
}

// OracleNewValueIterator is returned from FilterNewValue and is used to iterate over the raw logs and unpacked data for NewValue events raised by the Oracle contract.
type OracleNewValueIterator struct {
	Event *OracleNewValue // Event containing the contract specifics and raw log

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
func (it *OracleNewValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleNewValue)
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
		it.Event = new(OracleNewValue)
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
func (it *OracleNewValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleNewValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleNewValue represents a NewValue event raised by the Oracle contract.
type OracleNewValue struct {
	RequestId *big.Int
	Time      *big.Int
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewValue is a free log retrieval operation binding the contract event 0xba11e319aee26e7bbac889432515ba301ec8f6d27bf6b94829c21a65c5f6ff25.
//
// Solidity: event NewValue(uint256 _requestId, uint256 _time, uint256 _value)
func (_Oracle *OracleFilterer) FilterNewValue(opts *bind.FilterOpts) (*OracleNewValueIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "NewValue")
	if err != nil {
		return nil, err
	}
	return &OracleNewValueIterator{contract: _Oracle.contract, event: "NewValue", logs: logs, sub: sub}, nil
}

// WatchNewValue is a free log subscription operation binding the contract event 0xba11e319aee26e7bbac889432515ba301ec8f6d27bf6b94829c21a65c5f6ff25.
//
// Solidity: event NewValue(uint256 _requestId, uint256 _time, uint256 _value)
func (_Oracle *OracleFilterer) WatchNewValue(opts *bind.WatchOpts, sink chan<- *OracleNewValue) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "NewValue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleNewValue)
				if err := _Oracle.contract.UnpackLog(event, "NewValue", log); err != nil {
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

// ParseNewValue is a log parse operation binding the contract event 0xba11e319aee26e7bbac889432515ba301ec8f6d27bf6b94829c21a65c5f6ff25.
//
// Solidity: event NewValue(uint256 _requestId, uint256 _time, uint256 _value)
func (_Oracle *OracleFilterer) ParseNewValue(log types.Log) (*OracleNewValue, error) {
	event := new(OracleNewValue)
	if err := _Oracle.contract.UnpackLog(event, "NewValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Oracle contract.
type OracleRoleAdminChangedIterator struct {
	Event *OracleRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *OracleRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRoleAdminChanged)
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
		it.Event = new(OracleRoleAdminChanged)
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
func (it *OracleRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRoleAdminChanged represents a RoleAdminChanged event raised by the Oracle contract.
type OracleRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Oracle *OracleFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*OracleRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &OracleRoleAdminChangedIterator{contract: _Oracle.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Oracle *OracleFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *OracleRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRoleAdminChanged)
				if err := _Oracle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseRoleAdminChanged(log types.Log) (*OracleRoleAdminChanged, error) {
	event := new(OracleRoleAdminChanged)
	if err := _Oracle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Oracle contract.
type OracleRoleGrantedIterator struct {
	Event *OracleRoleGranted // Event containing the contract specifics and raw log

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
func (it *OracleRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRoleGranted)
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
		it.Event = new(OracleRoleGranted)
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
func (it *OracleRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRoleGranted represents a RoleGranted event raised by the Oracle contract.
type OracleRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Oracle *OracleFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OracleRoleGrantedIterator, error) {

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

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OracleRoleGrantedIterator{contract: _Oracle.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Oracle *OracleFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *OracleRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRoleGranted)
				if err := _Oracle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseRoleGranted(log types.Log) (*OracleRoleGranted, error) {
	event := new(OracleRoleGranted)
	if err := _Oracle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Oracle contract.
type OracleRoleRevokedIterator struct {
	Event *OracleRoleRevoked // Event containing the contract specifics and raw log

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
func (it *OracleRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRoleRevoked)
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
		it.Event = new(OracleRoleRevoked)
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
func (it *OracleRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRoleRevoked represents a RoleRevoked event raised by the Oracle contract.
type OracleRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Oracle *OracleFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OracleRoleRevokedIterator, error) {

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

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OracleRoleRevokedIterator{contract: _Oracle.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Oracle *OracleFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *OracleRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRoleRevoked)
				if err := _Oracle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseRoleRevoked(log types.Log) (*OracleRoleRevoked, error) {
	event := new(OracleRoleRevoked)
	if err := _Oracle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a03f702109ebb4b2937b923b4095e183c36936ae0260c723b97c5786d4652f0e64736f6c63430007000033"

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

// TellorAccessABI is the input ABI used to generate the binding from.
const TellorAccessABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"NewValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPORTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin_address\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reporter_address\",\"type\":\"address\"}],\"name\":\"addReporter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getCurrentValue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getDataBefore\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin_address\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reporter_address\",\"type\":\"address\"}],\"name\":\"isReporter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reporter_address\",\"type\":\"address\"}],\"name\":\"removeReporter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"submitValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"timestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"values\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TellorAccessFuncSigs maps the 4-byte function signature to its string representation.
var TellorAccessFuncSigs = map[string]string{
	"a217fddf": "DEFAULT_ADMIN_ROLE()",
	"3f60d799": "REPORTER_ROLE()",
	"70480275": "addAdmin(address)",
	"dd8755f2": "addReporter(address)",
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
	"5de5c212": "removeReporter(address)",
	"8bad0c0a": "renounceAdmin()",
	"36568abe": "renounceRole(bytes32,address)",
	"93fa4915": "retrieveData(uint256,uint256)",
	"d547741f": "revokeRole(bytes32,address)",
	"62f55112": "submitValue(uint256,uint256)",
	"fb0ceb04": "timestamps(uint256,uint256)",
	"a3183701": "values(uint256,uint256)",
}

// TellorAccessBin is the compiled bytecode used for deploying new contracts.
var TellorAccessBin = "0x60806040523480156200001157600080fd5b506200001f60003362000052565b6200004c7f176c2b761bfeb5dab89f614b6c08152e31d9230394b3605eabf32249ea1c89a6600062000062565b620001b8565b6200005e8282620000b4565b5050565b600082815260208190526040808220600201549051839285917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a460009182526020829052604090912060020155565b600082815260208181526040909120620000d991839062000a226200012d821b17901c565b156200005e57620000e96200014d565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b600062000144836001600160a01b03841662000151565b90505b92915050565b3390565b60006200015f8383620001a0565b620001975750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000147565b50600062000147565b60009081526001919091016020526040902054151590565b610f1e80620001c86000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c806370480275116100c3578063a217fddf1161007c578063a217fddf146103e8578063a3183701146103f0578063ca15c87314610413578063d547741f14610430578063dd8755f21461045c578063fb0ceb04146104825761014d565b8063704802751461030957806377fbb6631461032f5780638bad0c0a146103525780639010d07c1461035a57806391d148541461039957806393fa4915146103c55761014d565b80633f60d799116101155780633f60d7991461023b5780633fcad9641461024357806346eee1c4146102805780635de5c2121461029d57806362f55112146102c357806366b44611146102e65761014d565b8063044ad7be14610152578063248a9ca31461018c57806324d7806c146101bb5780632f2ff15d146101e157806336568abe1461020f575b600080fd5b6101786004803603602081101561016857600080fd5b50356001600160a01b03166104a5565b604080519115158252519081900360200190f35b6101a9600480360360208110156101a257600080fd5b50356104c5565b60408051918252519081900360200190f35b610178600480360360208110156101d157600080fd5b50356001600160a01b03166104da565b61020d600480360360408110156101f757600080fd5b50803590602001356001600160a01b03166104e6565b005b61020d6004803603604081101561022557600080fd5b50803590602001356001600160a01b0316610552565b6101a96105b3565b6102606004803603602081101561025957600080fd5b50356105c5565b604080519315158452602084019290925282820152519081900360600190f35b6101a96004803603602081101561029657600080fd5b503561061d565b61020d600480360360208110156102b357600080fd5b50356001600160a01b031661062f565b61020d600480360360408110156102d957600080fd5b508035906020013561069c565b610260600480360360408110156102fc57600080fd5b508035906020013561076b565b61020d6004803603602081101561031f57600080fd5b50356001600160a01b03166107e1565b6101a96004803603604081101561034557600080fd5b508035906020013561083e565b61020d610897565b61037d6004803603604081101561037057600080fd5b50803590602001356108a4565b604080516001600160a01b039092168252519081900360200190f35b610178600480360360408110156103af57600080fd5b50803590602001356001600160a01b03166108c3565b6101a9600480360360408110156103db57600080fd5b50803590602001356108db565b6101a96108f8565b6101a96004803603604081101561040657600080fd5b50803590602001356108fd565b6101a96004803603602081101561042957600080fd5b503561091a565b61020d6004803603604081101561044657600080fd5b50803590602001356001600160a01b0316610931565b61020d6004803603602081101561047257600080fd5b50356001600160a01b031661098a565b6101a96004803603604081101561049857600080fd5b50803590602001356109f4565b60006104bf600080516020610e38833981519152836108c3565b92915050565b60009081526020819052604090206002015490565b60006104bf81836108c3565b60008281526020819052604090206002015461050990610504610a37565b6108c3565b6105445760405162461bcd60e51b815260040180806020018281038252602f815260200180610e09602f913960400191505060405180910390fd5b61054e8282610a3b565b5050565b61055a610a37565b6001600160a01b0316816001600160a01b0316146105a95760405162461bcd60e51b815260040180806020018281038252602f815260200180610eba602f913960400191505060405180910390fd5b61054e8282610aa4565b600080516020610e3883398151915281565b6000806000806105d48561061d565b905060006105e5866001840361083e565b905060006105f387836108db565b9050801561060a5760019550935091506106169050565b50600094508493509150505b9193909250565b60009081526002602052604090205490565b610638336104da565b610681576040805162461bcd60e51b81526020600482015260156024820152742932b9ba3934b1ba32b2103a379030b236b4b7399760591b604482015290519081900360640190fd5b610699600080516020610e3883398151915282610931565b50565b6106a5336104a5565b806106b457506106b4336104da565b6106ef5760405162461bcd60e51b8152600401808060200182810382526032815260200180610e586032913960400191505060405180910390fd5b6000828152600160208181526040808420428086529083528185208690558685526002835281852080549485018155855293829020909201839055815185815290810192909252818101839052517fba11e319aee26e7bbac889432515ba301ec8f6d27bf6b94829c21a65c5f6ff259181900360600190a15050565b600080600080600061077d8787610b0d565b915091508161079857600080600094509450945050506107da565b60006107a4888361083e565b905060006107b289836108db565b905080156107ca5760019650945092506107da915050565b6000806000965096509650505050505b9250925092565b6107ea336104da565b610833576040805162461bcd60e51b81526020600482015260156024820152742932b9ba3934b1ba32b2103a379030b236b4b7399760591b604482015290519081900360640190fd5b6106996000826104e6565b60008281526002602052604081205480158061085a5750828111155b156108695760009150506104bf565b600084815260026020526040902080548490811061088357fe5b906000526020600020015491505092915050565b6108a2600033610552565b565b60008281526020819052604081206108bc9083610c15565b9392505050565b60008281526020819052604081206108bc9083610c21565b600091825260016020908152604080842092845291905290205490565b600081565b600160209081526000928352604080842090915290825290205481565b60008181526020819052604081206104bf90610c36565b60008281526020819052604090206002015461094f90610504610a37565b6105a95760405162461bcd60e51b8152600401808060200182810382526030815260200180610e8a6030913960400191505060405180910390fd5b610993336104da565b6109dc576040805162461bcd60e51b81526020600482015260156024820152742932b9ba3934b1ba32b2103a379030b236b4b7399760591b604482015290519081900360640190fd5b610699600080516020610e38833981519152826104e6565b60026020528160005260406000208181548110610a0d57fe5b90600052602060002001600091509150505481565b60006108bc836001600160a01b038416610c41565b3390565b6000828152602081905260409020610a539082610a22565b1561054e57610a60610a37565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6000828152602081905260409020610abc9082610c8b565b1561054e57610ac9610a37565b6001600160a01b0316816001600160a01b0316837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45050565b6000806000610b1b8561061d565b90508015610c0557600080600019830181610b36898261083e565b9050878110610b5057600080965096505050505050610c0e565b610b5a898361083e565b905087811015610b745750600195509350610c0e92505050565b82600281840304600101019350610b8b898561083e565b905087811015610bcb576000610ba48a8660010161083e565b9050888110610bbf5760018597509750505050505050610c0e565b84600101935050610c00565b6000610bda8a6001870361083e565b905088811015610bf857600180860397509750505050505050610c0e565b600185039250505b610b74565b60008092509250505b9250929050565b60006108bc8383610ca0565b60006108bc836001600160a01b038416610d04565b60006104bf82610d1c565b6000610c4d8383610d04565b610c83575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556104bf565b5060006104bf565b60006108bc836001600160a01b038416610d20565b81546000908210610ce25760405162461bcd60e51b8152600401808060200182810382526022815260200180610de76022913960400191505060405180910390fd5b826000018281548110610cf157fe5b9060005260206000200154905092915050565b60009081526001919091016020526040902054151590565b5490565b60008181526001830160205260408120548015610ddc5783546000198083019190810190600090879083908110610d5357fe5b9060005260206000200154905080876000018481548110610d7057fe5b600091825260208083209091019290925582815260018981019092526040902090840190558654879080610da057fe5b600190038181906000526020600020016000905590558660010160008781526020019081526020016000206000905560019450505050506104bf565b60009150506104bf56fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e6473416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f206772616e74176c2b761bfeb5dab89f614b6c08152e31d9230394b3605eabf32249ea1c89a653656e646572206d75737420626520616e2041646d696e206f72205265706f7274657220746f207375626d697456616c7565416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f207265766f6b65416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636520726f6c657320666f722073656c66a2646970667358221220cb3f20dcc96b8272dd153ce430a7de99e9471fbd72aed1d1d8512f97b24e19d764736f6c63430007000033"

// DeployTellorAccess deploys a new Ethereum contract, binding an instance of TellorAccess to it.
func DeployTellorAccess(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorAccess, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorAccessABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TellorAccessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TellorAccess{TellorAccessCaller: TellorAccessCaller{contract: contract}, TellorAccessTransactor: TellorAccessTransactor{contract: contract}, TellorAccessFilterer: TellorAccessFilterer{contract: contract}}, nil
}

// TellorAccess is an auto generated Go binding around an Ethereum contract.
type TellorAccess struct {
	TellorAccessCaller     // Read-only binding to the contract
	TellorAccessTransactor // Write-only binding to the contract
	TellorAccessFilterer   // Log filterer for contract events
}

// TellorAccessCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorAccessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorAccessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorAccessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorAccessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorAccessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorAccessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorAccessSession struct {
	Contract     *TellorAccess     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorAccessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorAccessCallerSession struct {
	Contract *TellorAccessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TellorAccessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorAccessTransactorSession struct {
	Contract     *TellorAccessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TellorAccessRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorAccessRaw struct {
	Contract *TellorAccess // Generic contract binding to access the raw methods on
}

// TellorAccessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorAccessCallerRaw struct {
	Contract *TellorAccessCaller // Generic read-only contract binding to access the raw methods on
}

// TellorAccessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorAccessTransactorRaw struct {
	Contract *TellorAccessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellorAccess creates a new instance of TellorAccess, bound to a specific deployed contract.
func NewTellorAccess(address common.Address, backend bind.ContractBackend) (*TellorAccess, error) {
	contract, err := bindTellorAccess(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TellorAccess{TellorAccessCaller: TellorAccessCaller{contract: contract}, TellorAccessTransactor: TellorAccessTransactor{contract: contract}, TellorAccessFilterer: TellorAccessFilterer{contract: contract}}, nil
}

// NewTellorAccessCaller creates a new read-only instance of TellorAccess, bound to a specific deployed contract.
func NewTellorAccessCaller(address common.Address, caller bind.ContractCaller) (*TellorAccessCaller, error) {
	contract, err := bindTellorAccess(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorAccessCaller{contract: contract}, nil
}

// NewTellorAccessTransactor creates a new write-only instance of TellorAccess, bound to a specific deployed contract.
func NewTellorAccessTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorAccessTransactor, error) {
	contract, err := bindTellorAccess(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorAccessTransactor{contract: contract}, nil
}

// NewTellorAccessFilterer creates a new log filterer instance of TellorAccess, bound to a specific deployed contract.
func NewTellorAccessFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorAccessFilterer, error) {
	contract, err := bindTellorAccess(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorAccessFilterer{contract: contract}, nil
}

// bindTellorAccess binds a generic wrapper to an already deployed contract.
func bindTellorAccess(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorAccessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorAccess *TellorAccessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorAccess.Contract.TellorAccessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorAccess *TellorAccessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorAccess.Contract.TellorAccessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorAccess *TellorAccessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorAccess.Contract.TellorAccessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorAccess *TellorAccessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorAccess.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorAccess *TellorAccessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorAccess.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorAccess *TellorAccessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorAccess.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TellorAccess *TellorAccessCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorAccess.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TellorAccess *TellorAccessSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TellorAccess.Contract.DEFAULTADMINROLE(&_TellorAccess.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TellorAccess *TellorAccessCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TellorAccess.Contract.DEFAULTADMINROLE(&_TellorAccess.CallOpts)
}

// REPORTERROLE is a free data retrieval call binding the contract method 0x3f60d799.
//
// Solidity: function REPORTER_ROLE() view returns(bytes32)
func (_TellorAccess *TellorAccessCaller) REPORTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorAccess.contract.Call(opts, &out, "REPORTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REPORTERROLE is a free data retrieval call binding the contract method 0x3f60d799.
//
// Solidity: function REPORTER_ROLE() view returns(bytes32)
func (_TellorAccess *TellorAccessSession) REPORTERROLE() ([32]byte, error) {
	return _TellorAccess.Contract.REPORTERROLE(&_TellorAccess.CallOpts)
}

// REPORTERROLE is a free data retrieval call binding the contract method 0x3f60d799.
//
// Solidity: function REPORTER_ROLE() view returns(bytes32)
func (_TellorAccess *TellorAccessCallerSession) REPORTERROLE() ([32]byte, error) {
	return _TellorAccess.Contract.REPORTERROLE(&_TellorAccess.CallOpts)
}

// GetCurrentValue is a free data retrieval call binding the contract method 0x3fcad964.
//
// Solidity: function getCurrentValue(uint256 _requestId) view returns(bool, uint256, uint256)
func (_TellorAccess *TellorAccessCaller) GetCurrentValue(opts *bind.CallOpts, _requestId *big.Int) (bool, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _TellorAccess.contract.Call(opts, &out, "getCurrentValue", _requestId)

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
func (_TellorAccess *TellorAccessSession) GetCurrentValue(_requestId *big.Int) (bool, *big.Int, *big.Int, error) {
	return _TellorAccess.Contract.GetCurrentValue(&_TellorAccess.CallOpts, _requestId)
}

// GetCurrentValue is a free data retrieval call binding the contract method 0x3fcad964.
//
// Solidity: function getCurrentValue(uint256 _requestId) view returns(bool, uint256, uint256)
func (_TellorAccess *TellorAccessCallerSession) GetCurrentValue(_requestId *big.Int) (bool, *big.Int, *big.Int, error) {
	return _TellorAccess.Contract.GetCurrentValue(&_TellorAccess.CallOpts, _requestId)
}

// GetDataBefore is a free data retrieval call binding the contract method 0x66b44611.
//
// Solidity: function getDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool, uint256, uint256)
func (_TellorAccess *TellorAccessCaller) GetDataBefore(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _TellorAccess.contract.Call(opts, &out, "getDataBefore", _requestId, _timestamp)

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
func (_TellorAccess *TellorAccessSession) GetDataBefore(_requestId *big.Int, _timestamp *big.Int) (bool, *big.Int, *big.Int, error) {
	return _TellorAccess.Contract.GetDataBefore(&_TellorAccess.CallOpts, _requestId, _timestamp)
}

// GetDataBefore is a free data retrieval call binding the contract method 0x66b44611.
//
// Solidity: function getDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool, uint256, uint256)
func (_TellorAccess *TellorAccessCallerSession) GetDataBefore(_requestId *big.Int, _timestamp *big.Int) (bool, *big.Int, *big.Int, error) {
	return _TellorAccess.Contract.GetDataBefore(&_TellorAccess.CallOpts, _requestId, _timestamp)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_TellorAccess *TellorAccessCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorAccess.contract.Call(opts, &out, "getNewValueCountbyRequestId", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_TellorAccess *TellorAccessSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _TellorAccess.Contract.GetNewValueCountbyRequestId(&_TellorAccess.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_TellorAccess *TellorAccessCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _TellorAccess.Contract.GetNewValueCountbyRequestId(&_TellorAccess.CallOpts, _requestId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TellorAccess *TellorAccessCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TellorAccess.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TellorAccess *TellorAccessSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TellorAccess.Contract.GetRoleAdmin(&_TellorAccess.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TellorAccess *TellorAccessCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TellorAccess.Contract.GetRoleAdmin(&_TellorAccess.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TellorAccess *TellorAccessCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TellorAccess.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TellorAccess *TellorAccessSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _TellorAccess.Contract.GetRoleMember(&_TellorAccess.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TellorAccess *TellorAccessCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _TellorAccess.Contract.GetRoleMember(&_TellorAccess.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TellorAccess *TellorAccessCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorAccess.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TellorAccess *TellorAccessSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _TellorAccess.Contract.GetRoleMemberCount(&_TellorAccess.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TellorAccess *TellorAccessCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _TellorAccess.Contract.GetRoleMemberCount(&_TellorAccess.CallOpts, role)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 _index) view returns(uint256)
func (_TellorAccess *TellorAccessCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestId *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorAccess.contract.Call(opts, &out, "getTimestampbyRequestIDandIndex", _requestId, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 _index) view returns(uint256)
func (_TellorAccess *TellorAccessSession) GetTimestampbyRequestIDandIndex(_requestId *big.Int, _index *big.Int) (*big.Int, error) {
	return _TellorAccess.Contract.GetTimestampbyRequestIDandIndex(&_TellorAccess.CallOpts, _requestId, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 _index) view returns(uint256)
func (_TellorAccess *TellorAccessCallerSession) GetTimestampbyRequestIDandIndex(_requestId *big.Int, _index *big.Int) (*big.Int, error) {
	return _TellorAccess.Contract.GetTimestampbyRequestIDandIndex(&_TellorAccess.CallOpts, _requestId, _index)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TellorAccess *TellorAccessCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TellorAccess.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TellorAccess *TellorAccessSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TellorAccess.Contract.HasRole(&_TellorAccess.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TellorAccess *TellorAccessCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TellorAccess.Contract.HasRole(&_TellorAccess.CallOpts, role, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _admin_address) view returns(bool)
func (_TellorAccess *TellorAccessCaller) IsAdmin(opts *bind.CallOpts, _admin_address common.Address) (bool, error) {
	var out []interface{}
	err := _TellorAccess.contract.Call(opts, &out, "isAdmin", _admin_address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _admin_address) view returns(bool)
func (_TellorAccess *TellorAccessSession) IsAdmin(_admin_address common.Address) (bool, error) {
	return _TellorAccess.Contract.IsAdmin(&_TellorAccess.CallOpts, _admin_address)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _admin_address) view returns(bool)
func (_TellorAccess *TellorAccessCallerSession) IsAdmin(_admin_address common.Address) (bool, error) {
	return _TellorAccess.Contract.IsAdmin(&_TellorAccess.CallOpts, _admin_address)
}

// IsReporter is a free data retrieval call binding the contract method 0x044ad7be.
//
// Solidity: function isReporter(address _reporter_address) view returns(bool)
func (_TellorAccess *TellorAccessCaller) IsReporter(opts *bind.CallOpts, _reporter_address common.Address) (bool, error) {
	var out []interface{}
	err := _TellorAccess.contract.Call(opts, &out, "isReporter", _reporter_address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReporter is a free data retrieval call binding the contract method 0x044ad7be.
//
// Solidity: function isReporter(address _reporter_address) view returns(bool)
func (_TellorAccess *TellorAccessSession) IsReporter(_reporter_address common.Address) (bool, error) {
	return _TellorAccess.Contract.IsReporter(&_TellorAccess.CallOpts, _reporter_address)
}

// IsReporter is a free data retrieval call binding the contract method 0x044ad7be.
//
// Solidity: function isReporter(address _reporter_address) view returns(bool)
func (_TellorAccess *TellorAccessCallerSession) IsReporter(_reporter_address common.Address) (bool, error) {
	return _TellorAccess.Contract.IsReporter(&_TellorAccess.CallOpts, _reporter_address)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorAccess *TellorAccessCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorAccess.contract.Call(opts, &out, "retrieveData", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorAccess *TellorAccessSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorAccess.Contract.RetrieveData(&_TellorAccess.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorAccess *TellorAccessCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorAccess.Contract.RetrieveData(&_TellorAccess.CallOpts, _requestId, _timestamp)
}

// Timestamps is a free data retrieval call binding the contract method 0xfb0ceb04.
//
// Solidity: function timestamps(uint256 , uint256 ) view returns(uint256)
func (_TellorAccess *TellorAccessCaller) Timestamps(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorAccess.contract.Call(opts, &out, "timestamps", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timestamps is a free data retrieval call binding the contract method 0xfb0ceb04.
//
// Solidity: function timestamps(uint256 , uint256 ) view returns(uint256)
func (_TellorAccess *TellorAccessSession) Timestamps(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TellorAccess.Contract.Timestamps(&_TellorAccess.CallOpts, arg0, arg1)
}

// Timestamps is a free data retrieval call binding the contract method 0xfb0ceb04.
//
// Solidity: function timestamps(uint256 , uint256 ) view returns(uint256)
func (_TellorAccess *TellorAccessCallerSession) Timestamps(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TellorAccess.Contract.Timestamps(&_TellorAccess.CallOpts, arg0, arg1)
}

// Values is a free data retrieval call binding the contract method 0xa3183701.
//
// Solidity: function values(uint256 , uint256 ) view returns(uint256)
func (_TellorAccess *TellorAccessCaller) Values(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorAccess.contract.Call(opts, &out, "values", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Values is a free data retrieval call binding the contract method 0xa3183701.
//
// Solidity: function values(uint256 , uint256 ) view returns(uint256)
func (_TellorAccess *TellorAccessSession) Values(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TellorAccess.Contract.Values(&_TellorAccess.CallOpts, arg0, arg1)
}

// Values is a free data retrieval call binding the contract method 0xa3183701.
//
// Solidity: function values(uint256 , uint256 ) view returns(uint256)
func (_TellorAccess *TellorAccessCallerSession) Values(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TellorAccess.Contract.Values(&_TellorAccess.CallOpts, arg0, arg1)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin_address) returns()
func (_TellorAccess *TellorAccessTransactor) AddAdmin(opts *bind.TransactOpts, _admin_address common.Address) (*types.Transaction, error) {
	return _TellorAccess.contract.Transact(opts, "addAdmin", _admin_address)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin_address) returns()
func (_TellorAccess *TellorAccessSession) AddAdmin(_admin_address common.Address) (*types.Transaction, error) {
	return _TellorAccess.Contract.AddAdmin(&_TellorAccess.TransactOpts, _admin_address)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin_address) returns()
func (_TellorAccess *TellorAccessTransactorSession) AddAdmin(_admin_address common.Address) (*types.Transaction, error) {
	return _TellorAccess.Contract.AddAdmin(&_TellorAccess.TransactOpts, _admin_address)
}

// AddReporter is a paid mutator transaction binding the contract method 0xdd8755f2.
//
// Solidity: function addReporter(address _reporter_address) returns()
func (_TellorAccess *TellorAccessTransactor) AddReporter(opts *bind.TransactOpts, _reporter_address common.Address) (*types.Transaction, error) {
	return _TellorAccess.contract.Transact(opts, "addReporter", _reporter_address)
}

// AddReporter is a paid mutator transaction binding the contract method 0xdd8755f2.
//
// Solidity: function addReporter(address _reporter_address) returns()
func (_TellorAccess *TellorAccessSession) AddReporter(_reporter_address common.Address) (*types.Transaction, error) {
	return _TellorAccess.Contract.AddReporter(&_TellorAccess.TransactOpts, _reporter_address)
}

// AddReporter is a paid mutator transaction binding the contract method 0xdd8755f2.
//
// Solidity: function addReporter(address _reporter_address) returns()
func (_TellorAccess *TellorAccessTransactorSession) AddReporter(_reporter_address common.Address) (*types.Transaction, error) {
	return _TellorAccess.Contract.AddReporter(&_TellorAccess.TransactOpts, _reporter_address)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TellorAccess *TellorAccessTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TellorAccess.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TellorAccess *TellorAccessSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TellorAccess.Contract.GrantRole(&_TellorAccess.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TellorAccess *TellorAccessTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TellorAccess.Contract.GrantRole(&_TellorAccess.TransactOpts, role, account)
}

// RemoveReporter is a paid mutator transaction binding the contract method 0x5de5c212.
//
// Solidity: function removeReporter(address _reporter_address) returns()
func (_TellorAccess *TellorAccessTransactor) RemoveReporter(opts *bind.TransactOpts, _reporter_address common.Address) (*types.Transaction, error) {
	return _TellorAccess.contract.Transact(opts, "removeReporter", _reporter_address)
}

// RemoveReporter is a paid mutator transaction binding the contract method 0x5de5c212.
//
// Solidity: function removeReporter(address _reporter_address) returns()
func (_TellorAccess *TellorAccessSession) RemoveReporter(_reporter_address common.Address) (*types.Transaction, error) {
	return _TellorAccess.Contract.RemoveReporter(&_TellorAccess.TransactOpts, _reporter_address)
}

// RemoveReporter is a paid mutator transaction binding the contract method 0x5de5c212.
//
// Solidity: function removeReporter(address _reporter_address) returns()
func (_TellorAccess *TellorAccessTransactorSession) RemoveReporter(_reporter_address common.Address) (*types.Transaction, error) {
	return _TellorAccess.Contract.RemoveReporter(&_TellorAccess.TransactOpts, _reporter_address)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_TellorAccess *TellorAccessTransactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorAccess.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_TellorAccess *TellorAccessSession) RenounceAdmin() (*types.Transaction, error) {
	return _TellorAccess.Contract.RenounceAdmin(&_TellorAccess.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_TellorAccess *TellorAccessTransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _TellorAccess.Contract.RenounceAdmin(&_TellorAccess.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TellorAccess *TellorAccessTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TellorAccess.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TellorAccess *TellorAccessSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TellorAccess.Contract.RenounceRole(&_TellorAccess.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TellorAccess *TellorAccessTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TellorAccess.Contract.RenounceRole(&_TellorAccess.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TellorAccess *TellorAccessTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TellorAccess.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TellorAccess *TellorAccessSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TellorAccess.Contract.RevokeRole(&_TellorAccess.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TellorAccess *TellorAccessTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TellorAccess.Contract.RevokeRole(&_TellorAccess.TransactOpts, role, account)
}

// SubmitValue is a paid mutator transaction binding the contract method 0x62f55112.
//
// Solidity: function submitValue(uint256 _requestId, uint256 _value) returns()
func (_TellorAccess *TellorAccessTransactor) SubmitValue(opts *bind.TransactOpts, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _TellorAccess.contract.Transact(opts, "submitValue", _requestId, _value)
}

// SubmitValue is a paid mutator transaction binding the contract method 0x62f55112.
//
// Solidity: function submitValue(uint256 _requestId, uint256 _value) returns()
func (_TellorAccess *TellorAccessSession) SubmitValue(_requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _TellorAccess.Contract.SubmitValue(&_TellorAccess.TransactOpts, _requestId, _value)
}

// SubmitValue is a paid mutator transaction binding the contract method 0x62f55112.
//
// Solidity: function submitValue(uint256 _requestId, uint256 _value) returns()
func (_TellorAccess *TellorAccessTransactorSession) SubmitValue(_requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _TellorAccess.Contract.SubmitValue(&_TellorAccess.TransactOpts, _requestId, _value)
}

// TellorAccessNewValueIterator is returned from FilterNewValue and is used to iterate over the raw logs and unpacked data for NewValue events raised by the TellorAccess contract.
type TellorAccessNewValueIterator struct {
	Event *TellorAccessNewValue // Event containing the contract specifics and raw log

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
func (it *TellorAccessNewValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorAccessNewValue)
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
		it.Event = new(TellorAccessNewValue)
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
func (it *TellorAccessNewValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorAccessNewValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorAccessNewValue represents a NewValue event raised by the TellorAccess contract.
type TellorAccessNewValue struct {
	RequestId *big.Int
	Time      *big.Int
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewValue is a free log retrieval operation binding the contract event 0xba11e319aee26e7bbac889432515ba301ec8f6d27bf6b94829c21a65c5f6ff25.
//
// Solidity: event NewValue(uint256 _requestId, uint256 _time, uint256 _value)
func (_TellorAccess *TellorAccessFilterer) FilterNewValue(opts *bind.FilterOpts) (*TellorAccessNewValueIterator, error) {

	logs, sub, err := _TellorAccess.contract.FilterLogs(opts, "NewValue")
	if err != nil {
		return nil, err
	}
	return &TellorAccessNewValueIterator{contract: _TellorAccess.contract, event: "NewValue", logs: logs, sub: sub}, nil
}

// WatchNewValue is a free log subscription operation binding the contract event 0xba11e319aee26e7bbac889432515ba301ec8f6d27bf6b94829c21a65c5f6ff25.
//
// Solidity: event NewValue(uint256 _requestId, uint256 _time, uint256 _value)
func (_TellorAccess *TellorAccessFilterer) WatchNewValue(opts *bind.WatchOpts, sink chan<- *TellorAccessNewValue) (event.Subscription, error) {

	logs, sub, err := _TellorAccess.contract.WatchLogs(opts, "NewValue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorAccessNewValue)
				if err := _TellorAccess.contract.UnpackLog(event, "NewValue", log); err != nil {
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

// ParseNewValue is a log parse operation binding the contract event 0xba11e319aee26e7bbac889432515ba301ec8f6d27bf6b94829c21a65c5f6ff25.
//
// Solidity: event NewValue(uint256 _requestId, uint256 _time, uint256 _value)
func (_TellorAccess *TellorAccessFilterer) ParseNewValue(log types.Log) (*TellorAccessNewValue, error) {
	event := new(TellorAccessNewValue)
	if err := _TellorAccess.contract.UnpackLog(event, "NewValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorAccessRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TellorAccess contract.
type TellorAccessRoleAdminChangedIterator struct {
	Event *TellorAccessRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TellorAccessRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorAccessRoleAdminChanged)
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
		it.Event = new(TellorAccessRoleAdminChanged)
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
func (it *TellorAccessRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorAccessRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorAccessRoleAdminChanged represents a RoleAdminChanged event raised by the TellorAccess contract.
type TellorAccessRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TellorAccess *TellorAccessFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TellorAccessRoleAdminChangedIterator, error) {

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

	logs, sub, err := _TellorAccess.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TellorAccessRoleAdminChangedIterator{contract: _TellorAccess.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TellorAccess *TellorAccessFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TellorAccessRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _TellorAccess.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorAccessRoleAdminChanged)
				if err := _TellorAccess.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_TellorAccess *TellorAccessFilterer) ParseRoleAdminChanged(log types.Log) (*TellorAccessRoleAdminChanged, error) {
	event := new(TellorAccessRoleAdminChanged)
	if err := _TellorAccess.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorAccessRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TellorAccess contract.
type TellorAccessRoleGrantedIterator struct {
	Event *TellorAccessRoleGranted // Event containing the contract specifics and raw log

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
func (it *TellorAccessRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorAccessRoleGranted)
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
		it.Event = new(TellorAccessRoleGranted)
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
func (it *TellorAccessRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorAccessRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorAccessRoleGranted represents a RoleGranted event raised by the TellorAccess contract.
type TellorAccessRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TellorAccess *TellorAccessFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TellorAccessRoleGrantedIterator, error) {

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

	logs, sub, err := _TellorAccess.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TellorAccessRoleGrantedIterator{contract: _TellorAccess.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TellorAccess *TellorAccessFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TellorAccessRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TellorAccess.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorAccessRoleGranted)
				if err := _TellorAccess.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_TellorAccess *TellorAccessFilterer) ParseRoleGranted(log types.Log) (*TellorAccessRoleGranted, error) {
	event := new(TellorAccessRoleGranted)
	if err := _TellorAccess.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorAccessRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TellorAccess contract.
type TellorAccessRoleRevokedIterator struct {
	Event *TellorAccessRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TellorAccessRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorAccessRoleRevoked)
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
		it.Event = new(TellorAccessRoleRevoked)
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
func (it *TellorAccessRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorAccessRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorAccessRoleRevoked represents a RoleRevoked event raised by the TellorAccess contract.
type TellorAccessRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TellorAccess *TellorAccessFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TellorAccessRoleRevokedIterator, error) {

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

	logs, sub, err := _TellorAccess.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TellorAccessRoleRevokedIterator{contract: _TellorAccess.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TellorAccess *TellorAccessFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TellorAccessRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TellorAccess.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorAccessRoleRevoked)
				if err := _TellorAccess.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_TellorAccess *TellorAccessFilterer) ParseRoleRevoked(log types.Log) (*TellorAccessRoleRevoked, error) {
	event := new(TellorAccessRoleRevoked)
	if err := _TellorAccess.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
