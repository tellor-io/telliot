// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OldTellorGettersLibraries

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

// OldTellorGettersLibrariesABI is the input ABI used to generate the binding from.
const OldTellorGettersLibrariesABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newTellor\",\"type\":\"address\"}],\"name\":\"OldNewTellorAddress\",\"type\":\"event\"}]"

// OldTellorGettersLibrariesBin is the compiled bytecode used for deploying new contracts.
var OldTellorGettersLibrariesBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158200e6b8a3af2c23b7a7d806881bf3b8ef2f8dd1a12eb0e4e66d57686d14306fc6664736f6c63430005100032"

// DeployOldTellorGettersLibraries deploys a new Ethereum contract, binding an instance of OldTellorGettersLibraries to it.
func DeployOldTellorGettersLibraries(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OldTellorGettersLibraries, error) {
	parsed, err := abi.JSON(strings.NewReader(OldTellorGettersLibrariesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OldTellorGettersLibrariesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OldTellorGettersLibraries{OldTellorGettersLibrariesCaller: OldTellorGettersLibrariesCaller{contract: contract}, OldTellorGettersLibrariesTransactor: OldTellorGettersLibrariesTransactor{contract: contract}, OldTellorGettersLibrariesFilterer: OldTellorGettersLibrariesFilterer{contract: contract}}, nil
}

// OldTellorGettersLibraries is an auto generated Go binding around an Ethereum contract.
type OldTellorGettersLibraries struct {
	OldTellorGettersLibrariesCaller     // Read-only binding to the contract
	OldTellorGettersLibrariesTransactor // Write-only binding to the contract
	OldTellorGettersLibrariesFilterer   // Log filterer for contract events
}

// OldTellorGettersLibrariesCaller is an auto generated read-only Go binding around an Ethereum contract.
type OldTellorGettersLibrariesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldTellorGettersLibrariesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OldTellorGettersLibrariesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldTellorGettersLibrariesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OldTellorGettersLibrariesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldTellorGettersLibrariesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OldTellorGettersLibrariesSession struct {
	Contract     *OldTellorGettersLibraries // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// OldTellorGettersLibrariesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OldTellorGettersLibrariesCallerSession struct {
	Contract *OldTellorGettersLibrariesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// OldTellorGettersLibrariesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OldTellorGettersLibrariesTransactorSession struct {
	Contract     *OldTellorGettersLibrariesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// OldTellorGettersLibrariesRaw is an auto generated low-level Go binding around an Ethereum contract.
type OldTellorGettersLibrariesRaw struct {
	Contract *OldTellorGettersLibraries // Generic contract binding to access the raw methods on
}

// OldTellorGettersLibrariesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OldTellorGettersLibrariesCallerRaw struct {
	Contract *OldTellorGettersLibrariesCaller // Generic read-only contract binding to access the raw methods on
}

// OldTellorGettersLibrariesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OldTellorGettersLibrariesTransactorRaw struct {
	Contract *OldTellorGettersLibrariesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOldTellorGettersLibraries creates a new instance of OldTellorGettersLibraries, bound to a specific deployed contract.
func NewOldTellorGettersLibraries(address common.Address, backend bind.ContractBackend) (*OldTellorGettersLibraries, error) {
	contract, err := bindOldTellorGettersLibraries(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OldTellorGettersLibraries{OldTellorGettersLibrariesCaller: OldTellorGettersLibrariesCaller{contract: contract}, OldTellorGettersLibrariesTransactor: OldTellorGettersLibrariesTransactor{contract: contract}, OldTellorGettersLibrariesFilterer: OldTellorGettersLibrariesFilterer{contract: contract}}, nil
}

// NewOldTellorGettersLibrariesCaller creates a new read-only instance of OldTellorGettersLibraries, bound to a specific deployed contract.
func NewOldTellorGettersLibrariesCaller(address common.Address, caller bind.ContractCaller) (*OldTellorGettersLibrariesCaller, error) {
	contract, err := bindOldTellorGettersLibraries(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OldTellorGettersLibrariesCaller{contract: contract}, nil
}

// NewOldTellorGettersLibrariesTransactor creates a new write-only instance of OldTellorGettersLibraries, bound to a specific deployed contract.
func NewOldTellorGettersLibrariesTransactor(address common.Address, transactor bind.ContractTransactor) (*OldTellorGettersLibrariesTransactor, error) {
	contract, err := bindOldTellorGettersLibraries(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OldTellorGettersLibrariesTransactor{contract: contract}, nil
}

// NewOldTellorGettersLibrariesFilterer creates a new log filterer instance of OldTellorGettersLibraries, bound to a specific deployed contract.
func NewOldTellorGettersLibrariesFilterer(address common.Address, filterer bind.ContractFilterer) (*OldTellorGettersLibrariesFilterer, error) {
	contract, err := bindOldTellorGettersLibraries(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OldTellorGettersLibrariesFilterer{contract: contract}, nil
}

// bindOldTellorGettersLibraries binds a generic wrapper to an already deployed contract.
func bindOldTellorGettersLibraries(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OldTellorGettersLibrariesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OldTellorGettersLibraries *OldTellorGettersLibrariesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OldTellorGettersLibraries.Contract.OldTellorGettersLibrariesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OldTellorGettersLibraries *OldTellorGettersLibrariesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OldTellorGettersLibraries.Contract.OldTellorGettersLibrariesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OldTellorGettersLibraries *OldTellorGettersLibrariesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OldTellorGettersLibraries.Contract.OldTellorGettersLibrariesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OldTellorGettersLibraries *OldTellorGettersLibrariesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OldTellorGettersLibraries.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OldTellorGettersLibraries *OldTellorGettersLibrariesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OldTellorGettersLibraries.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OldTellorGettersLibraries *OldTellorGettersLibrariesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OldTellorGettersLibraries.Contract.contract.Transact(opts, method, params...)
}

// OldTellorGettersLibrariesOldNewTellorAddressIterator is returned from FilterOldNewTellorAddress and is used to iterate over the raw logs and unpacked data for OldNewTellorAddress events raised by the OldTellorGettersLibraries contract.
type OldTellorGettersLibrariesOldNewTellorAddressIterator struct {
	Event *OldTellorGettersLibrariesOldNewTellorAddress // Event containing the contract specifics and raw log

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
func (it *OldTellorGettersLibrariesOldNewTellorAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OldTellorGettersLibrariesOldNewTellorAddress)
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
		it.Event = new(OldTellorGettersLibrariesOldNewTellorAddress)
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
func (it *OldTellorGettersLibrariesOldNewTellorAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OldTellorGettersLibrariesOldNewTellorAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OldTellorGettersLibrariesOldNewTellorAddress represents a OldNewTellorAddress event raised by the OldTellorGettersLibraries contract.
type OldTellorGettersLibrariesOldNewTellorAddress struct {
	NewTellor common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOldNewTellorAddress is a free log retrieval operation binding the contract event 0x4239d23661972c85791dbf8c5d8e537960414a4be5b728323ea1636b40d6ad3a.
//
// Solidity: event OldNewTellorAddress(address _newTellor)
func (_OldTellorGettersLibraries *OldTellorGettersLibrariesFilterer) FilterOldNewTellorAddress(opts *bind.FilterOpts) (*OldTellorGettersLibrariesOldNewTellorAddressIterator, error) {

	logs, sub, err := _OldTellorGettersLibraries.contract.FilterLogs(opts, "OldNewTellorAddress")
	if err != nil {
		return nil, err
	}
	return &OldTellorGettersLibrariesOldNewTellorAddressIterator{contract: _OldTellorGettersLibraries.contract, event: "OldNewTellorAddress", logs: logs, sub: sub}, nil
}

// WatchOldNewTellorAddress is a free log subscription operation binding the contract event 0x4239d23661972c85791dbf8c5d8e537960414a4be5b728323ea1636b40d6ad3a.
//
// Solidity: event OldNewTellorAddress(address _newTellor)
func (_OldTellorGettersLibraries *OldTellorGettersLibrariesFilterer) WatchOldNewTellorAddress(opts *bind.WatchOpts, sink chan<- *OldTellorGettersLibrariesOldNewTellorAddress) (event.Subscription, error) {

	logs, sub, err := _OldTellorGettersLibraries.contract.WatchLogs(opts, "OldNewTellorAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OldTellorGettersLibrariesOldNewTellorAddress)
				if err := _OldTellorGettersLibraries.contract.UnpackLog(event, "OldNewTellorAddress", log); err != nil {
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

// ParseOldNewTellorAddress is a log parse operation binding the contract event 0x4239d23661972c85791dbf8c5d8e537960414a4be5b728323ea1636b40d6ad3a.
//
// Solidity: event OldNewTellorAddress(address _newTellor)
func (_OldTellorGettersLibraries *OldTellorGettersLibrariesFilterer) ParseOldNewTellorAddress(log types.Log) (*OldTellorGettersLibrariesOldNewTellorAddress, error) {
	event := new(OldTellorGettersLibrariesOldNewTellorAddress)
	if err := _OldTellorGettersLibraries.contract.UnpackLog(event, "OldNewTellorAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
