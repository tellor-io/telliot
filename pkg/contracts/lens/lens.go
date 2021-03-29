// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lens

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

// MainDataID is an auto generated low-level Go binding around an user-defined struct.
type MainDataID struct {
	Id          *big.Int
	Name        string
	Granularity *big.Int
}

// MainValue is an auto generated low-level Go binding around an user-defined struct.
type MainValue struct {
	Meta      MainDataID
	Timestamp *big.Int
	Value     *big.Int
	Tip       *big.Int
}

// ITellorABI is the input ABI used to generate the binding from.
const ITellorABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minerIndex\",\"type\":\"uint256\"}],\"name\":\"beginDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[9]\",\"name\":\"\",\"type\":\"uint256[9]\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastNewValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"internalType\":\"address[5]\",\"name\":\"\",\"type\":\"address[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256\",\"name\":\"_difficutly\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"idsOnDeck\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"tipsOnDeck\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_request\",\"type\":\"bytes32\"}],\"name\":\"getRequestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"\",\"type\":\"uint256[51]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTopRequestIDs\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_propNewTellorAddress\",\"type\":\"address\"}],\"name\":\"proposeFork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_pendingOwner\",\"type\":\"address\"}],\"name\":\"proposeOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestStakingWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"}],\"name\":\"submitMiningSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"submitMiningSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"tallyVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"unlockDisputeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"updateTellor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_supportsDispute\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ITellorFuncSigs maps the 4-byte function signature to its string representation.
var ITellorFuncSigs = map[string]string{
	"752d49a1": "addTip(uint256,uint256)",
	"dd62ed3e": "allowance(address,address)",
	"999cf26c": "allowedToTrade(address,uint256)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"4ee2cd7e": "balanceOfAt(address,uint256)",
	"8581af19": "beginDispute(uint256,uint256,uint256)",
	"4e71e0c8": "claimOwnership()",
	"313ce567": "decimals()",
	"0d2d76a2": "depositStake()",
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
	"4049f198": "getNewCurrentVariables()",
	"46eee1c4": "getNewValueCountbyRequestId(uint256)",
	"9a7077ab": "getNewVariablesOnDeck()",
	"1db842f0": "getRequestIdByQueryHash(bytes32)",
	"6173c0b8": "getRequestIdByRequestQIndex(uint256)",
	"0f0b424d": "getRequestIdByTimestamp(uint256)",
	"b5413029": "getRequestQ()",
	"e0ae93c1": "getRequestUintVars(uint256,bytes32)",
	"e1eee6d6": "getRequestVars(uint256)",
	"733bdef0": "getStakerInfo(address)",
	"11c98512": "getSubmissionsByTimestamp(uint256,uint256)",
	"77fbb663": "getTimestampbyRequestIDandIndex(uint256,uint256)",
	"fe1cd15d": "getTopRequestIDs()",
	"612c8f7f": "getUintVar(bytes32)",
	"19e8e03b": "getVariablesOnDeck()",
	"3df0777b": "isInDispute(uint256,uint256)",
	"06fdde03": "name()",
	"26b7d9f6": "proposeFork(address)",
	"710bf322": "proposeOwnership(address)",
	"28449c3a": "requestStakingWithdraw()",
	"93fa4915": "retrieveData(uint256,uint256)",
	"68c180d5": "submitMiningSolution(string,uint256,uint256)",
	"4350283e": "submitMiningSolution(string,uint256[5],uint256[5])",
	"95d89b41": "symbol()",
	"4d318b0e": "tallyVotes(uint256)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"9a01ca13": "unlockDisputeFee(uint256)",
	"f458ab98": "updateTellor(uint256)",
	"c9d27afe": "vote(uint256,bool)",
	"bed9d861": "withdrawStake()",
}

// ITellor is an auto generated Go binding around an Ethereum contract.
type ITellor struct {
	ITellorCaller     // Read-only binding to the contract
	ITellorTransactor // Write-only binding to the contract
	ITellorFilterer   // Log filterer for contract events
}

// ITellorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITellorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITellorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITellorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITellorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITellorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITellorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITellorSession struct {
	Contract     *ITellor          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITellorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITellorCallerSession struct {
	Contract *ITellorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ITellorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITellorTransactorSession struct {
	Contract     *ITellorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ITellorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITellorRaw struct {
	Contract *ITellor // Generic contract binding to access the raw methods on
}

// ITellorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITellorCallerRaw struct {
	Contract *ITellorCaller // Generic read-only contract binding to access the raw methods on
}

// ITellorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITellorTransactorRaw struct {
	Contract *ITellorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITellor creates a new instance of ITellor, bound to a specific deployed contract.
func NewITellor(address common.Address, backend bind.ContractBackend) (*ITellor, error) {
	contract, err := bindITellor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITellor{ITellorCaller: ITellorCaller{contract: contract}, ITellorTransactor: ITellorTransactor{contract: contract}, ITellorFilterer: ITellorFilterer{contract: contract}}, nil
}

// NewITellorCaller creates a new read-only instance of ITellor, bound to a specific deployed contract.
func NewITellorCaller(address common.Address, caller bind.ContractCaller) (*ITellorCaller, error) {
	contract, err := bindITellor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITellorCaller{contract: contract}, nil
}

// NewITellorTransactor creates a new write-only instance of ITellor, bound to a specific deployed contract.
func NewITellorTransactor(address common.Address, transactor bind.ContractTransactor) (*ITellorTransactor, error) {
	contract, err := bindITellor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITellorTransactor{contract: contract}, nil
}

// NewITellorFilterer creates a new log filterer instance of ITellor, bound to a specific deployed contract.
func NewITellorFilterer(address common.Address, filterer bind.ContractFilterer) (*ITellorFilterer, error) {
	contract, err := bindITellor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITellorFilterer{contract: contract}, nil
}

// bindITellor binds a generic wrapper to an already deployed contract.
func bindITellor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITellorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITellor *ITellorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITellor.Contract.ITellorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITellor *ITellorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITellor.Contract.ITellorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITellor *ITellorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITellor.Contract.ITellorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITellor *ITellorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITellor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITellor *ITellorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITellor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITellor *ITellorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITellor.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_ITellor *ITellorCaller) Allowance(opts *bind.CallOpts, _user common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "allowance", _user, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_ITellor *ITellorSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _ITellor.Contract.Allowance(&_ITellor.CallOpts, _user, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_ITellor *ITellorCallerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _ITellor.Contract.Allowance(&_ITellor.CallOpts, _user, _spender)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_ITellor *ITellorCaller) AllowedToTrade(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "allowedToTrade", _user, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_ITellor *ITellorSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _ITellor.Contract.AllowedToTrade(&_ITellor.CallOpts, _user, _amount)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_ITellor *ITellorCallerSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _ITellor.Contract.AllowedToTrade(&_ITellor.CallOpts, _user, _amount)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ITellor *ITellorCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "balanceOf", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ITellor *ITellorSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _ITellor.Contract.BalanceOf(&_ITellor.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ITellor *ITellorCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _ITellor.Contract.BalanceOf(&_ITellor.CallOpts, _user)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_ITellor *ITellorCaller) BalanceOfAt(opts *bind.CallOpts, _user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "balanceOfAt", _user, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_ITellor *ITellorSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _ITellor.Contract.BalanceOfAt(&_ITellor.CallOpts, _user, _blockNumber)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_ITellor *ITellorCallerSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _ITellor.Contract.BalanceOfAt(&_ITellor.CallOpts, _user, _blockNumber)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_ITellor *ITellorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_ITellor *ITellorSession) Decimals() (uint8, error) {
	return _ITellor.Contract.Decimals(&_ITellor.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_ITellor *ITellorCallerSession) Decimals() (uint8, error) {
	return _ITellor.Contract.Decimals(&_ITellor.CallOpts)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_ITellor *ITellorCaller) DidMine(opts *bind.CallOpts, _challenge [32]byte, _miner common.Address) (bool, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "didMine", _challenge, _miner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_ITellor *ITellorSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _ITellor.Contract.DidMine(&_ITellor.CallOpts, _challenge, _miner)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_ITellor *ITellorCallerSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _ITellor.Contract.DidMine(&_ITellor.CallOpts, _challenge, _miner)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_ITellor *ITellorCaller) DidVote(opts *bind.CallOpts, _disputeId *big.Int, _address common.Address) (bool, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "didVote", _disputeId, _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_ITellor *ITellorSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _ITellor.Contract.DidVote(&_ITellor.CallOpts, _disputeId, _address)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_ITellor *ITellorCallerSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _ITellor.Contract.DidVote(&_ITellor.CallOpts, _disputeId, _address)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_ITellor *ITellorCaller) GetAddressVars(opts *bind.CallOpts, _data [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getAddressVars", _data)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_ITellor *ITellorSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _ITellor.Contract.GetAddressVars(&_ITellor.CallOpts, _data)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_ITellor *ITellorCallerSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _ITellor.Contract.GetAddressVars(&_ITellor.CallOpts, _data)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ITellor *ITellorCaller) GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getAllDisputeVars", _disputeId)

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
func (_ITellor *ITellorSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _ITellor.Contract.GetAllDisputeVars(&_ITellor.CallOpts, _disputeId)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ITellor *ITellorCallerSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _ITellor.Contract.GetAllDisputeVars(&_ITellor.CallOpts, _disputeId)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() view returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ITellor *ITellorCaller) GetCurrentVariables(opts *bind.CallOpts) ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getCurrentVariables")

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
func (_ITellor *ITellorSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _ITellor.Contract.GetCurrentVariables(&_ITellor.CallOpts)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() view returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ITellor *ITellorCallerSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _ITellor.Contract.GetCurrentVariables(&_ITellor.CallOpts)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_ITellor *ITellorCaller) GetDisputeIdByDisputeHash(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getDisputeIdByDisputeHash", _hash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_ITellor *ITellorSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _ITellor.Contract.GetDisputeIdByDisputeHash(&_ITellor.CallOpts, _hash)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_ITellor *ITellorCallerSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _ITellor.Contract.GetDisputeIdByDisputeHash(&_ITellor.CallOpts, _hash)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_ITellor *ITellorCaller) GetDisputeUintVars(opts *bind.CallOpts, _disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getDisputeUintVars", _disputeId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_ITellor *ITellorSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ITellor.Contract.GetDisputeUintVars(&_ITellor.CallOpts, _disputeId, _data)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_ITellor *ITellorCallerSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ITellor.Contract.GetDisputeUintVars(&_ITellor.CallOpts, _disputeId, _data)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_ITellor *ITellorCaller) GetLastNewValue(opts *bind.CallOpts) (*big.Int, bool, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getLastNewValue")

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
func (_ITellor *ITellorSession) GetLastNewValue() (*big.Int, bool, error) {
	return _ITellor.Contract.GetLastNewValue(&_ITellor.CallOpts)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_ITellor *ITellorCallerSession) GetLastNewValue() (*big.Int, bool, error) {
	return _ITellor.Contract.GetLastNewValue(&_ITellor.CallOpts)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_ITellor *ITellorCaller) GetLastNewValueById(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, bool, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getLastNewValueById", _requestId)

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
func (_ITellor *ITellorSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _ITellor.Contract.GetLastNewValueById(&_ITellor.CallOpts, _requestId)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_ITellor *ITellorCallerSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _ITellor.Contract.GetLastNewValueById(&_ITellor.CallOpts, _requestId)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ITellor *ITellorCaller) GetMinedBlockNum(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getMinedBlockNum", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ITellor *ITellorSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ITellor.Contract.GetMinedBlockNum(&_ITellor.CallOpts, _requestId, _timestamp)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ITellor *ITellorCallerSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ITellor.Contract.GetMinedBlockNum(&_ITellor.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_ITellor *ITellorCaller) GetMinersByRequestIdAndTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getMinersByRequestIdAndTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([5]common.Address)).(*[5]common.Address)

	return out0, err

}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_ITellor *ITellorSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _ITellor.Contract.GetMinersByRequestIdAndTimestamp(&_ITellor.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_ITellor *ITellorCallerSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _ITellor.Contract.GetMinersByRequestIdAndTimestamp(&_ITellor.CallOpts, _requestId, _timestamp)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficutly, uint256 _tip)
func (_ITellor *ITellorCaller) GetNewCurrentVariables(opts *bind.CallOpts) (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getNewCurrentVariables")

	outstruct := new(struct {
		Challenge  [32]byte
		RequestIds [5]*big.Int
		Difficutly *big.Int
		Tip        *big.Int
	})

	outstruct.Challenge = out[0].([32]byte)
	outstruct.RequestIds = out[1].([5]*big.Int)
	outstruct.Difficutly = out[2].(*big.Int)
	outstruct.Tip = out[3].(*big.Int)

	return *outstruct, err

}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficutly, uint256 _tip)
func (_ITellor *ITellorSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	return _ITellor.Contract.GetNewCurrentVariables(&_ITellor.CallOpts)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficutly, uint256 _tip)
func (_ITellor *ITellorCallerSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	return _ITellor.Contract.GetNewCurrentVariables(&_ITellor.CallOpts)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_ITellor *ITellorCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getNewValueCountbyRequestId", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_ITellor *ITellorSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _ITellor.Contract.GetNewValueCountbyRequestId(&_ITellor.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_ITellor *ITellorCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _ITellor.Contract.GetNewValueCountbyRequestId(&_ITellor.CallOpts, _requestId)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_ITellor *ITellorCaller) GetNewVariablesOnDeck(opts *bind.CallOpts) (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getNewVariablesOnDeck")

	outstruct := new(struct {
		IdsOnDeck  [5]*big.Int
		TipsOnDeck [5]*big.Int
	})

	outstruct.IdsOnDeck = out[0].([5]*big.Int)
	outstruct.TipsOnDeck = out[1].([5]*big.Int)

	return *outstruct, err

}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_ITellor *ITellorSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _ITellor.Contract.GetNewVariablesOnDeck(&_ITellor.CallOpts)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_ITellor *ITellorCallerSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _ITellor.Contract.GetNewVariablesOnDeck(&_ITellor.CallOpts)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(bytes32 _request) view returns(uint256)
func (_ITellor *ITellorCaller) GetRequestIdByQueryHash(opts *bind.CallOpts, _request [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getRequestIdByQueryHash", _request)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(bytes32 _request) view returns(uint256)
func (_ITellor *ITellorSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _ITellor.Contract.GetRequestIdByQueryHash(&_ITellor.CallOpts, _request)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(bytes32 _request) view returns(uint256)
func (_ITellor *ITellorCallerSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _ITellor.Contract.GetRequestIdByQueryHash(&_ITellor.CallOpts, _request)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_ITellor *ITellorCaller) GetRequestIdByRequestQIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getRequestIdByRequestQIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_ITellor *ITellorSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _ITellor.Contract.GetRequestIdByRequestQIndex(&_ITellor.CallOpts, _index)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_ITellor *ITellorCallerSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _ITellor.Contract.GetRequestIdByRequestQIndex(&_ITellor.CallOpts, _index)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_ITellor *ITellorCaller) GetRequestIdByTimestamp(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getRequestIdByTimestamp", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_ITellor *ITellorSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _ITellor.Contract.GetRequestIdByTimestamp(&_ITellor.CallOpts, _timestamp)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_ITellor *ITellorCallerSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _ITellor.Contract.GetRequestIdByTimestamp(&_ITellor.CallOpts, _timestamp)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_ITellor *ITellorCaller) GetRequestQ(opts *bind.CallOpts) ([51]*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getRequestQ")

	if err != nil {
		return *new([51]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([51]*big.Int)).(*[51]*big.Int)

	return out0, err

}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_ITellor *ITellorSession) GetRequestQ() ([51]*big.Int, error) {
	return _ITellor.Contract.GetRequestQ(&_ITellor.CallOpts)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_ITellor *ITellorCallerSession) GetRequestQ() ([51]*big.Int, error) {
	return _ITellor.Contract.GetRequestQ(&_ITellor.CallOpts)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_ITellor *ITellorCaller) GetRequestUintVars(opts *bind.CallOpts, _requestId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getRequestUintVars", _requestId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_ITellor *ITellorSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ITellor.Contract.GetRequestUintVars(&_ITellor.CallOpts, _requestId, _data)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_ITellor *ITellorCallerSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ITellor.Contract.GetRequestUintVars(&_ITellor.CallOpts, _requestId, _data)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(string, string, bytes32, uint256, uint256, uint256)
func (_ITellor *ITellorCaller) GetRequestVars(opts *bind.CallOpts, _requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getRequestVars", _requestId)

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
func (_ITellor *ITellorSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _ITellor.Contract.GetRequestVars(&_ITellor.CallOpts, _requestId)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(string, string, bytes32, uint256, uint256, uint256)
func (_ITellor *ITellorCallerSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _ITellor.Contract.GetRequestVars(&_ITellor.CallOpts, _requestId)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_ITellor *ITellorCaller) GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getStakerInfo", _staker)

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
func (_ITellor *ITellorSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _ITellor.Contract.GetStakerInfo(&_ITellor.CallOpts, _staker)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_ITellor *ITellorCallerSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _ITellor.Contract.GetStakerInfo(&_ITellor.CallOpts, _staker)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_ITellor *ITellorCaller) GetSubmissionsByTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getSubmissionsByTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_ITellor *ITellorSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _ITellor.Contract.GetSubmissionsByTimestamp(&_ITellor.CallOpts, _requestId, _timestamp)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_ITellor *ITellorCallerSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _ITellor.Contract.GetSubmissionsByTimestamp(&_ITellor.CallOpts, _requestId, _timestamp)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_ITellor *ITellorCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getTimestampbyRequestIDandIndex", _requestID, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_ITellor *ITellorSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _ITellor.Contract.GetTimestampbyRequestIDandIndex(&_ITellor.CallOpts, _requestID, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_ITellor *ITellorCallerSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _ITellor.Contract.GetTimestampbyRequestIDandIndex(&_ITellor.CallOpts, _requestID, _index)
}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_ITellor *ITellorCaller) GetTopRequestIDs(opts *bind.CallOpts) ([5]*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getTopRequestIDs")

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_ITellor *ITellorSession) GetTopRequestIDs() ([5]*big.Int, error) {
	return _ITellor.Contract.GetTopRequestIDs(&_ITellor.CallOpts)
}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_ITellor *ITellorCallerSession) GetTopRequestIDs() ([5]*big.Int, error) {
	return _ITellor.Contract.GetTopRequestIDs(&_ITellor.CallOpts)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_ITellor *ITellorCaller) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getUintVar", _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_ITellor *ITellorSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _ITellor.Contract.GetUintVar(&_ITellor.CallOpts, _data)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_ITellor *ITellorCallerSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _ITellor.Contract.GetUintVar(&_ITellor.CallOpts, _data)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() view returns(uint256, uint256, string)
func (_ITellor *ITellorCaller) GetVariablesOnDeck(opts *bind.CallOpts) (*big.Int, *big.Int, string, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getVariablesOnDeck")

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
func (_ITellor *ITellorSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _ITellor.Contract.GetVariablesOnDeck(&_ITellor.CallOpts)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() view returns(uint256, uint256, string)
func (_ITellor *ITellorCallerSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _ITellor.Contract.GetVariablesOnDeck(&_ITellor.CallOpts)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_ITellor *ITellorCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "isInDispute", _requestId, _timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_ITellor *ITellorSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _ITellor.Contract.IsInDispute(&_ITellor.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_ITellor *ITellorCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _ITellor.Contract.IsInDispute(&_ITellor.CallOpts, _requestId, _timestamp)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_ITellor *ITellorCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_ITellor *ITellorSession) Name() (string, error) {
	return _ITellor.Contract.Name(&_ITellor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_ITellor *ITellorCallerSession) Name() (string, error) {
	return _ITellor.Contract.Name(&_ITellor.CallOpts)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ITellor *ITellorCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "retrieveData", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ITellor *ITellorSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ITellor.Contract.RetrieveData(&_ITellor.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ITellor *ITellorCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ITellor.Contract.RetrieveData(&_ITellor.CallOpts, _requestId, _timestamp)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_ITellor *ITellorCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_ITellor *ITellorSession) Symbol() (string, error) {
	return _ITellor.Contract.Symbol(&_ITellor.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_ITellor *ITellorCallerSession) Symbol() (string, error) {
	return _ITellor.Contract.Symbol(&_ITellor.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ITellor *ITellorCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ITellor *ITellorSession) TotalSupply() (*big.Int, error) {
	return _ITellor.Contract.TotalSupply(&_ITellor.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ITellor *ITellorCallerSession) TotalSupply() (*big.Int, error) {
	return _ITellor.Contract.TotalSupply(&_ITellor.CallOpts)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_ITellor *ITellorTransactor) AddTip(opts *bind.TransactOpts, _requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "addTip", _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_ITellor *ITellorSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.AddTip(&_ITellor.TransactOpts, _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_ITellor *ITellorTransactorSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.AddTip(&_ITellor.TransactOpts, _requestId, _tip)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_ITellor *ITellorTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_ITellor *ITellorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.Approve(&_ITellor.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_ITellor *ITellorTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.Approve(&_ITellor.TransactOpts, _spender, _amount)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_ITellor *ITellorTransactor) BeginDispute(opts *bind.TransactOpts, _requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "beginDispute", _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_ITellor *ITellorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.BeginDispute(&_ITellor.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_ITellor *ITellorTransactorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.BeginDispute(&_ITellor.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_ITellor *ITellorTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_ITellor *ITellorSession) ClaimOwnership() (*types.Transaction, error) {
	return _ITellor.Contract.ClaimOwnership(&_ITellor.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_ITellor *ITellorTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _ITellor.Contract.ClaimOwnership(&_ITellor.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_ITellor *ITellorTransactor) DepositStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "depositStake")
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_ITellor *ITellorSession) DepositStake() (*types.Transaction, error) {
	return _ITellor.Contract.DepositStake(&_ITellor.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_ITellor *ITellorTransactorSession) DepositStake() (*types.Transaction, error) {
	return _ITellor.Contract.DepositStake(&_ITellor.TransactOpts)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_ITellor *ITellorTransactor) ProposeFork(opts *bind.TransactOpts, _propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "proposeFork", _propNewTellorAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_ITellor *ITellorSession) ProposeFork(_propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ProposeFork(&_ITellor.TransactOpts, _propNewTellorAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_ITellor *ITellorTransactorSession) ProposeFork(_propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ProposeFork(&_ITellor.TransactOpts, _propNewTellorAddress)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _pendingOwner) returns()
func (_ITellor *ITellorTransactor) ProposeOwnership(opts *bind.TransactOpts, _pendingOwner common.Address) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "proposeOwnership", _pendingOwner)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _pendingOwner) returns()
func (_ITellor *ITellorSession) ProposeOwnership(_pendingOwner common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ProposeOwnership(&_ITellor.TransactOpts, _pendingOwner)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _pendingOwner) returns()
func (_ITellor *ITellorTransactorSession) ProposeOwnership(_pendingOwner common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ProposeOwnership(&_ITellor.TransactOpts, _pendingOwner)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_ITellor *ITellorTransactor) RequestStakingWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "requestStakingWithdraw")
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_ITellor *ITellorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _ITellor.Contract.RequestStakingWithdraw(&_ITellor.TransactOpts)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_ITellor *ITellorTransactorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _ITellor.Contract.RequestStakingWithdraw(&_ITellor.TransactOpts)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_ITellor *ITellorTransactor) SubmitMiningSolution(opts *bind.TransactOpts, _nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "submitMiningSolution", _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_ITellor *ITellorSession) SubmitMiningSolution(_nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.SubmitMiningSolution(&_ITellor.TransactOpts, _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_ITellor *ITellorTransactorSession) SubmitMiningSolution(_nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.SubmitMiningSolution(&_ITellor.TransactOpts, _nonce, _requestId, _value)
}

// SubmitMiningSolution0 is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_ITellor *ITellorTransactor) SubmitMiningSolution0(opts *bind.TransactOpts, _nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "submitMiningSolution0", _nonce, _requestId, _value)
}

// SubmitMiningSolution0 is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_ITellor *ITellorSession) SubmitMiningSolution0(_nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.SubmitMiningSolution0(&_ITellor.TransactOpts, _nonce, _requestId, _value)
}

// SubmitMiningSolution0 is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_ITellor *ITellorTransactorSession) SubmitMiningSolution0(_nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.SubmitMiningSolution0(&_ITellor.TransactOpts, _nonce, _requestId, _value)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_ITellor *ITellorTransactor) TallyVotes(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "tallyVotes", _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_ITellor *ITellorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.TallyVotes(&_ITellor.TransactOpts, _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_ITellor *ITellorTransactorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.TallyVotes(&_ITellor.TransactOpts, _disputeId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_ITellor *ITellorTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_ITellor *ITellorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.Transfer(&_ITellor.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_ITellor *ITellorTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.Transfer(&_ITellor.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_ITellor *ITellorTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_ITellor *ITellorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.TransferFrom(&_ITellor.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_ITellor *ITellorTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.TransferFrom(&_ITellor.TransactOpts, _from, _to, _amount)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_ITellor *ITellorTransactor) UnlockDisputeFee(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "unlockDisputeFee", _disputeId)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_ITellor *ITellorSession) UnlockDisputeFee(_disputeId *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.UnlockDisputeFee(&_ITellor.TransactOpts, _disputeId)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_ITellor *ITellorTransactorSession) UnlockDisputeFee(_disputeId *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.UnlockDisputeFee(&_ITellor.TransactOpts, _disputeId)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_ITellor *ITellorTransactor) UpdateTellor(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "updateTellor", _disputeId)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_ITellor *ITellorSession) UpdateTellor(_disputeId *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.UpdateTellor(&_ITellor.TransactOpts, _disputeId)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_ITellor *ITellorTransactorSession) UpdateTellor(_disputeId *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.UpdateTellor(&_ITellor.TransactOpts, _disputeId)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_ITellor *ITellorTransactor) Vote(opts *bind.TransactOpts, _disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "vote", _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_ITellor *ITellorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _ITellor.Contract.Vote(&_ITellor.TransactOpts, _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_ITellor *ITellorTransactorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _ITellor.Contract.Vote(&_ITellor.TransactOpts, _disputeId, _supportsDispute)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_ITellor *ITellorTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_ITellor *ITellorSession) WithdrawStake() (*types.Transaction, error) {
	return _ITellor.Contract.WithdrawStake(&_ITellor.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_ITellor *ITellorTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _ITellor.Contract.WithdrawStake(&_ITellor.TransactOpts)
}

// MainABI is the input ABI used to generate the binding from.
const MainABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_oracle\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"currentReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTotalTips\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dataIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"granularity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataIDsAll\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"granularity\",\"type\":\"uint256\"}],\"internalType\":\"structMain.DataID[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dataIDsMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deity\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"difficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disputeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disputeFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getCurrentValue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ifRetrieve\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestampRetrieved\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getDataBefore\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_ifRetrieve\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestampRetrieved\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getIndexForDataBefore\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dataID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"getLastValues\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"granularity\",\"type\":\"uint256\"}],\"internalType\":\"structMain.DataID\",\"name\":\"meta\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tip\",\"type\":\"uint256\"}],\"internalType\":\"structMain.Value[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getLastValuesAll\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"granularity\",\"type\":\"uint256\"}],\"internalType\":\"structMain.DataID\",\"name\":\"meta\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tip\",\"type\":\"uint256\"}],\"internalType\":\"structMain.Value[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"granularity\",\"type\":\"uint256\"}],\"internalType\":\"structMain.DataID\",\"name\":\"_dataID\",\"type\":\"tuple\"}],\"name\":\"pushDataID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"granularity\",\"type\":\"uint256\"}],\"internalType\":\"structMain.DataID[]\",\"name\":\"_dataIDs\",\"type\":\"tuple[]\"}],\"name\":\"replaceDataIDs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"granularity\",\"type\":\"uint256\"}],\"internalType\":\"structMain.DataID\",\"name\":\"_dataID\",\"type\":\"tuple\"}],\"name\":\"setDataID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slotProgress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tellorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeOfLastValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeTarget\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dataID\",\"type\":\"uint256\"}],\"name\":\"totalTip\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MainFuncSigs maps the 4-byte function signature to its string representation.
var MainFuncSigs = map[string]string{
	"07621eca": "currentReward()",
	"75ad1a2a": "currentTotalTips()",
	"351f9a50": "dataIDs(uint256)",
	"a2c9a952": "dataIDsAll()",
	"3a918872": "dataIDsMap(uint256)",
	"0426ce0f": "deity()",
	"19cae462": "difficulty()",
	"a28889e1": "disputeCount()",
	"b9ce896b": "disputeFee()",
	"3fcad964": "getCurrentValue(uint256)",
	"66b44611": "getDataBefore(uint256,uint256)",
	"b73e4979": "getIndexForDataBefore(uint256,uint256)",
	"a6fdc28c": "getLastValues(uint256,uint256)",
	"8b7a0e49": "getLastValuesAll(uint256)",
	"46eee1c4": "getNewValueCountbyRequestId(uint256)",
	"77fbb663": "getTimestampbyRequestIDandIndex(uint256,uint256)",
	"3df0777b": "isInDispute(uint256,uint256)",
	"7dc0d1d0": "oracle()",
	"8da5cb5b": "owner()",
	"e30c3978": "pendingOwner()",
	"4d5ddb67": "pushDataID((uint256,string,uint256))",
	"6d7ecbd6": "replaceDataIDs((uint256,string,uint256)[])",
	"5badbe4c": "requestCount()",
	"93fa4915": "retrieveData(uint256,uint256)",
	"704b6c02": "setAdmin(address)",
	"eabe1566": "setDataID(uint256,(uint256,string,uint256))",
	"7adbf973": "setOracle(address)",
	"03b3160f": "slotProgress()",
	"60c7dc47": "stakeAmount()",
	"c4a9e116": "stakeCount()",
	"6747dc31": "tBlock()",
	"a339ac74": "tellorContract()",
	"dfcff498": "timeOfLastValue()",
	"6fc37811": "timeTarget()",
	"44b12aea": "totalTip(uint256)",
}

// MainBin is the compiled bytecode used for deploying new contracts.
var MainBin = "0x608060405234801561001057600080fd5b5060405162002441380380620024418339810160408190526100319161006c565b600080546001600160a01b039092166001600160a01b031992831681179091556001805483169091179055600280549091163317905561009a565b60006020828403121561007d578081fd5b81516001600160a01b0381168114610093578182fd5b9392505050565b61239780620000aa6000396000f3fe608060405234801561001057600080fd5b50600436106102115760003560e01c8063704b6c0211610125578063a2c9a952116100ad578063b9ce896b1161007c578063b9ce896b1461042e578063c4a9e11614610313578063dfcff49814610436578063e30c39781461043e578063eabe15661461044657610211565b8063a2c9a952146103dd578063a339ac74146103f2578063a6fdc28c146103fa578063b73e49791461040d57610211565b80637dc0d1d0116100f45780637dc0d1d0146103925780638b7a0e491461039a5780638da5cb5b146103ba57806393fa4915146103c2578063a28889e1146103d557610211565b8063704b6c021461035157806375ad1a2a1461036457806377fbb6631461036c5780637adbf9731461037f57610211565b806344b12aea116101a857806360c7dc471161017757806360c7dc471461031357806366b446111461031b5780636747dc311461032e5780636d7ecbd6146103365780636fc378111461034957610211565b806344b12aea146102d057806346eee1c4146102e35780634d5ddb67146102f65780635badbe4c1461030b57610211565b8063351f9a50116101e4578063351f9a50146102595780633a9188721461027b5780633df0777b1461028e5780633fcad964146102ae57610211565b806303b3160f146102165780630426ce0f1461023457806307621eca1461024957806319cae46214610251575b600080fd5b61021e610459565b60405161022b91906122c2565b60405180910390f35b61023c6104ff565b60405161022b91906121a1565b61021e6105a0565b61021e610716565b61026c61026736600461208d565b610767565b60405161022b939291906122ff565b61021e61028936600461208d565b610839565b6102a161029c366004612102565b61084b565b60405161022b919061228f565b6102c16102bc36600461208d565b6108d2565b60405161022b939291906122aa565b61021e6102de36600461208d565b610a90565b61021e6102f136600461208d565b610b39565b610309610304366004612052565b610bb8565b005b61021e610c86565b61021e610cd7565b6102c1610329366004612102565b610d28565b61021e610e83565b610309610344366004611fb9565b610ed4565b61021e610fad565b61030961035f366004611f7a565b610ffe565b61021e61104a565b61021e61037a366004612102565b61109b565b61030961038d366004611f7a565b6110ef565b61023c61113b565b6103ad6103a836600461208d565b61114a565b60405161022b9190612215565b61023c611238565b61021e6103d0366004612102565b611289565b61021e6112dd565b6103e561132e565b60405161022b91906121b5565b61023c611443565b6103ad610408366004612102565b611494565b61042061041b366004612102565b611803565b60405161022b92919061229a565b61021e611bc2565b61021e611c13565b61023c611c64565b6103096104543660046120bd565b611cb5565b60015460405163612c8f7f60e01b81526000916001600160a01b03169063612c8f7f906104aa907fdfbec46864bc123768f0d134913175d9577a55bb71b9b2595fda21e21f36b082906004016122c2565b60206040518083038186803b1580156104c257600080fd5b505afa1580156104d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104fa91906120a5565b905090565b60015460405163099df72f60e11b81526000916001600160a01b03169063133bee5e90610550907f5fc094d10c65bc33cc842217b2eccca0191ff24148319da094e540a559898961906004016122c2565b60206040518083038186803b15801561056857600080fd5b505afa15801561057c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104fa9190611f9d565b60015460405163612c8f7f60e01b815260009182916001600160a01b039091169063612c8f7f906105f5907f2c8b528fbaf48aaf13162a5a0519a7ad5a612da8ff8783465c17e076660a59f1906004016122c2565b60206040518083038186803b15801561060d57600080fd5b505afa158015610621573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061064591906120a5565b60015460405163612c8f7f60e01b815242929092039250670de0b6b3a76400009161012c8385020491600091600a916001600160a01b03169063612c8f7f906106b2907f09659d32f99e50ac728058418d38174fe83a137c455ff1847e6fb8e15f78f77a906004016122c2565b60206040518083038186803b1580156106ca57600080fd5b505afa1580156106de573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061070291906120a5565b8161070957fe5b0491909101935050505090565b60015460405163612c8f7f60e01b81526000916001600160a01b03169063612c8f7f906104aa907ff758978fc1647996a3d9992f611883adc442931dc49488312360acc90601759b906004016122c2565b6003818154811061077757600080fd5b9060005260206000209060030201600091509050806000015490806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108295780601f106107fe57610100808354040283529160200191610829565b820191906000526020600020905b81548152906001019060200180831161080c57829003601f168201915b5050505050908060020154905083565b60046020526000908152604090205481565b6000805460408051633df0777b60e01b8152600481018690526024810185905290516001600160a01b0390921691633df0777b91604480820192602092909190829003018186803b15801561089f57600080fd5b505afa1580156108b3573d6000803e3d6000fd5b505050506040513d60208110156108c957600080fd5b50519392505050565b60008060008060008054906101000a90046001600160a01b03166001600160a01b03166346eee1c4866040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561092f57600080fd5b505afa158015610943573d6000803e3d6000fd5b505050506040513d602081101561095957600080fd5b505160008054604080516377fbb66360e01b8152600481018a905260001985016024820152905193945091926001600160a01b03909116916377fbb663916044808301926020929190829003018186803b1580156109b657600080fd5b505afa1580156109ca573d6000803e3d6000fd5b505050506040513d60208110156109e057600080fd5b505160008054604080516393fa491560e01b8152600481018b905260248101859052905193945091926001600160a01b03909116916393fa4915916044808301926020929190829003018186803b158015610a3a57600080fd5b505afa158015610a4e573d6000803e3d6000fd5b505050506040513d6020811015610a6457600080fd5b505190508015610a7d576001955093509150610a899050565b50600094508493509150505b9193909250565b60015460405163e0ae93c160e01b81526000916001600160a01b03169063e0ae93c190610ae39085907f1590276b7f31dd8e2a06f9a92867333eeb3eddbc91e73b9833e3e55d8e34f77d906004016122f1565b60206040518083038186803b158015610afb57600080fd5b505afa158015610b0f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b3391906120a5565b92915050565b60008054604080516311bbb87160e21b81526004810185905290516001600160a01b03909216916346eee1c491602480820192602092909190829003018186803b158015610b8657600080fd5b505afa158015610b9a573d6000803e3d6000fd5b505050506040513d6020811015610bb057600080fd5b505192915050565b6002546001600160a01b03163314610beb5760405162461bcd60e51b8152600401610be2906122cb565b60405180910390fd5b600380546001810182556000829052825191027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b81019182556020808401518051859493610c5f937fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c909101920190611d3e565b50604091820151600290910155600354915160009081526004602052206000199091019055565b60015460405163612c8f7f60e01b81526000916001600160a01b03169063612c8f7f906104aa907f3f8b5616fa9e7f2ce4a868fde15c58b92e77bc1acd6769bf1567629a3dc4c865906004016122c2565b60015460405163612c8f7f60e01b81526000916001600160a01b03169063612c8f7f906104aa907f5d9fadfc729fd027e395e5157ef1b53ef9fa4a8f053043c5f159307543e7cc97906004016122c2565b6000806000806000610d3a8787611803565b9150915081610d555760008060009450945094505050610e7c565b60008054604080516377fbb66360e01b8152600481018b90526024810185905290516001600160a01b03909216916377fbb66391604480820192602092909190829003018186803b158015610da957600080fd5b505afa158015610dbd573d6000803e3d6000fd5b505050506040513d6020811015610dd357600080fd5b5051600054604080516393fa491560e01b8152600481018c90526024810184905290519293506001600160a01b03909116916393fa491591604480820192602092909190829003018186803b158015610e2b57600080fd5b505afa158015610e3f573d6000803e3d6000fd5b505050506040513d6020811015610e5557600080fd5b505194508415610e6d57600195509250610e7c915050565b60008060009550955095505050505b9250925092565b60015460405163612c8f7f60e01b81526000916001600160a01b03169063612c8f7f906104aa907ff3b93531fa65b3a18680d9ea49df06d96fbd883c4889dc7db866f8b131602dfb906004016122c2565b6002546001600160a01b03163314610efe5760405162461bcd60e51b8152600401610be2906122cb565b610f0a60036000611dca565b60005b8151811015610fa9576003828281518110610f2457fe5b602090810291909101810151825460018181018555600094855293839020825160039092020190815581830151805192949193610f679392850192910190611d3e565b506040820151816002015550508060046000848481518110610f8557fe5b60209081029190910181015151825281019190915260400160002055600101610f0d565b5050565b60015460405163612c8f7f60e01b81526000916001600160a01b03169063612c8f7f906104aa907fd4f87b8d0f3d3b7e665df74631f6100b2695daa0e30e40eeac02172e15a999e1906004016122c2565b6002546001600160a01b031633146110285760405162461bcd60e51b8152600401610be2906122cb565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b60015460405163612c8f7f60e01b81526000916001600160a01b03169063612c8f7f906104aa907f09659d32f99e50ac728058418d38174fe83a137c455ff1847e6fb8e15f78f77a906004016122c2565b60008054604080516377fbb66360e01b8152600481018690526024810185905290516001600160a01b03909216916377fbb66391604480820192602092909190829003018186803b15801561089f57600080fd5b6002546001600160a01b031633146111195760405162461bcd60e51b8152600401610be2906122cb565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b031681565b600354606090600090830267ffffffffffffffff8111801561116b57600080fd5b506040519080825280602002602001820160405280156111a557816020015b611192611dee565b81526020019060019003908161118a5790505b5090506000805b60035481101561122f5760006111e0600383815481106111c857fe5b90600052602060002090600302016000015487611494565b905060005b8151811015611225578181815181106111fa57fe5b602002602001015185858151811061120e57fe5b6020908102919091010152600193840193016111e5565b50506001016111ac565b50909392505050565b60015460405163099df72f60e11b81526000916001600160a01b03169063133bee5e90610550907f7a39905194de50bde334d18b76bbb36dddd11641d4d50b470cb837cf3bae5def906004016122c2565b60008054604080516393fa491560e01b8152600481018690526024810185905290516001600160a01b03909216916393fa491591604480820192602092909190829003018186803b15801561089f57600080fd5b60015460405163612c8f7f60e01b81526000916001600160a01b03169063612c8f7f906104aa907f310199159a20c50879ffb440b45802138b5b162ec9426720e9dd3ee8bbcdb9d7906004016122c2565b60606003805480602002602001604051908101604052809291908181526020016000905b8282101561143a578382906000526020600020906003020160405180606001604052908160008201548152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156114185780601f106113ed57610100808354040283529160200191611418565b820191906000526020600020905b8154815290600101906020018083116113fb57829003601f168201915b5050505050815260200160028201548152505081526020019060010190611352565b50505050905090565b60015460405163099df72f60e11b81526000916001600160a01b03169063133bee5e90610550907f0f1293c916694ac6af4daa2f866f0448d0c2ce8847074a7896d397c961914a08906004016122c2565b6001546040516311bbb87160e21b81526060916000916001600160a01b03909116906346eee1c4906114ca9087906004016122c2565b60206040518083038186803b1580156114e257600080fd5b505afa1580156114f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061151a91906120a5565b905080831115611528578092505b60008367ffffffffffffffff8111801561154157600080fd5b5060405190808252806020026020018201604052801561157b57816020015b611568611dee565b8152602001906001900390816115605790505b50905060005b848110156117fa576001546040516377fbb66360e01b81526000916001600160a01b0316906377fbb663906115c2908a90600019878a0301906004016122f1565b60206040518083038186803b1580156115da57600080fd5b505afa1580156115ee573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061161291906120a5565b6001546040516393fa491560e01b81529192506000916001600160a01b03909116906393fa49159061164a908b9086906004016122f1565b60206040518083038186803b15801561166257600080fd5b505afa158015611676573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061169a91906120a5565b9050604051806080016040528060405180606001604052808b81526020016003600460008e815260200190815260200160002054815481106116d857fe5b90600052602060002090600302016001018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561177d5780601f106117525761010080835404028352916020019161177d565b820191906000526020600020905b81548152906001019060200180831161176057829003601f168201915b505050505081526020016003600460008e815260200190815260200160002054815481106117a757fe5b90600052602060002090600302016002015481525081526020018381526020018281526020016117d68a610a90565b8152508484815181106117e557fe5b60209081029190910101525050600101611581565b50949350505050565b60008054604080516311bbb87160e21b8152600481018690529051839283926001600160a01b03909116916346eee1c491602480820192602092909190829003018186803b15801561185457600080fd5b505afa158015611868573d6000803e3d6000fd5b505050506040513d602081101561187e57600080fd5b505190508015611bb25760008054604080516377fbb66360e01b815260048101899052602481018490529051839260001986019284926001600160a01b03909216916377fbb66391604480820192602092909190829003018186803b1580156118e657600080fd5b505afa1580156118fa573d6000803e3d6000fd5b505050506040513d602081101561191057600080fd5b5051905087811061192c57600080965096505050505050611bbb565b600054604080516377fbb66360e01b8152600481018c90526024810185905290516001600160a01b03909216916377fbb66391604480820192602092909190829003018186803b15801561197f57600080fd5b505afa158015611993573d6000803e3d6000fd5b505050506040513d60208110156119a957600080fd5b50519050878110156119c55750600195509350611bbb92505050565b600054604080516377fbb66360e01b8152600481018c905260028686030486016001016024820181905291519196506001600160a01b03909216916377fbb663916044808301926020929190829003018186803b158015611a2557600080fd5b505afa158015611a39573d6000803e3d6000fd5b505050506040513d6020811015611a4f57600080fd5b5051905087811015611b045760008054604080516377fbb66360e01b8152600481018d905260018801602482015290516001600160a01b03909216916377fbb66391604480820192602092909190829003018186803b158015611ab157600080fd5b505afa158015611ac5573d6000803e3d6000fd5b505050506040513d6020811015611adb57600080fd5b50519050888110611af85760018597509750505050505050611bbb565b84600101935050611bad565b60008054604080516377fbb66360e01b8152600481018d90526000198801602482015290516001600160a01b03909216916377fbb66391604480820192602092909190829003018186803b158015611b5b57600080fd5b505afa158015611b6f573d6000803e3d6000fd5b505050506040513d6020811015611b8557600080fd5b5051905088811015611ba557600180860397509750505050505050611bbb565b600185039250505b6119c5565b60008092509250505b9250929050565b60015460405163612c8f7f60e01b81526000916001600160a01b03169063612c8f7f906104aa907f675d2171f68d6f5545d54fb9b1fb61a0e6897e6188ca1cd664e7c9530d91ecfc906004016122c2565b60015460405163612c8f7f60e01b81526000916001600160a01b03169063612c8f7f906104aa907f2c8b528fbaf48aaf13162a5a0519a7ad5a612da8ff8783465c17e076660a59f1906004016122c2565b60015460405163099df72f60e11b81526000916001600160a01b03169063133bee5e90610550907f7ec081f029b8ac7e2321f6ae8c6a6a517fda8fcbf63cabd63dfffaeaafa56cc0906004016122c2565b6002546001600160a01b03163314611cdf5760405162461bcd60e51b8152600401610be2906122cb565b8060038381548110611ced57fe5b9060005260206000209060030201600082015181600001556020820151816001019080519060200190611d21929190611d3e565b506040918201516002909101559051600090815260046020522055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282611d745760008555611dba565b82601f10611d8d57805160ff1916838001178555611dba565b82800160010185558215611dba579182015b82811115611dba578251825591602001919060010190611d9f565b50611dc6929150611e1c565b5090565b5080546000825560030290600052602060002090810190611deb9190611e31565b50565b6040518060800160405280611e01611e5b565b81526020016000815260200160008152602001600081525090565b5b80821115611dc65760008155600101611e1d565b80821115611dc6576000808255611e4b6001830182611e7c565b5060006002820155600301611e31565b60405180606001604052806000815260200160608152602001600081525090565b50805460018160011615610100020316600290046000825580601f10611ea25750611deb565b601f016020900490600052602060002090810190611deb9190611e1c565b600060608284031215611ed1578081fd5b6040516060810167ffffffffffffffff8282108183111715611eef57fe5b8160405282935084358352602091508185013581811115611f0f57600080fd5b8501601f81018713611f2057600080fd5b803582811115611f2c57fe5b611f3e601f8201601f19168501612328565b92508083528784828401011115611f5457600080fd5b808483018585013760009083018401525090820152604092830135920191909152919050565b600060208284031215611f8b578081fd5b8135611f968161234c565b9392505050565b600060208284031215611fae578081fd5b8151611f968161234c565b60006020808385031215611fcb578182fd5b823567ffffffffffffffff80821115611fe2578384fd5b818501915085601f830112611ff5578384fd5b81358181111561200157fe5b61200e8485830201612328565b8181528481019250838501865b83811015612044576120328a888435890101611ec0565b8552938601939086019060010161201b565b509098975050505050505050565b600060208284031215612063578081fd5b813567ffffffffffffffff811115612079578182fd5b61208584828501611ec0565b949350505050565b60006020828403121561209e578081fd5b5035919050565b6000602082840312156120b6578081fd5b5051919050565b600080604083850312156120cf578081fd5b82359150602083013567ffffffffffffffff8111156120ec578182fd5b6120f885828601611ec0565b9150509250929050565b60008060408385031215612114578182fd5b50508035926020909101359150565b60008151808452815b818110156121485760208185018101518683018201520161212c565b818111156121595782602083870101525b50601f01601f19169290920160200192915050565b60008151835260208201516060602085015261218d6060850182612123565b604093840151949093019390935250919050565b6001600160a01b0391909116815260200190565b6000602080830181845280855180835260408601915060408482028701019250838701855b8281101561220857603f198886030184526121f685835161216e565b945092850192908501906001016121da565b5092979650505050505050565b60208082528251828201819052600091906040908185019080840286018301878501865b8381101561204457603f1989840301855281516080815181865261225f8287018261216e565b838b0151878c0152898401518a880152606093840151939096019290925250509386019390860190600101612239565b901515815260200190565b9115158252602082015260400190565b92151583526020830191909152604082015260600190565b90815260200190565b6020808252600c908201526b3737ba1030b71030b236b4b760a11b604082015260600190565b918252602082015260400190565b6000848252606060208301526123186060830185612123565b9050826040830152949350505050565b60405181810167ffffffffffffffff8111828210171561234457fe5b604052919050565b6001600160a01b0381168114611deb57600080fdfea26469706673582212201c9940d5790bd5d7740f131455343fce96c7f269d71a3d996f2088fc465205db64736f6c63430007060033"

// DeployMain deploys a new Ethereum contract, binding an instance of Main to it.
func DeployMain(auth *bind.TransactOpts, backend bind.ContractBackend, _oracle common.Address) (common.Address, *types.Transaction, *Main, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MainBin), backend, _oracle)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(uint256)
func (_Main *MainCaller) CurrentReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "currentReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(uint256)
func (_Main *MainSession) CurrentReward() (*big.Int, error) {
	return _Main.Contract.CurrentReward(&_Main.CallOpts)
}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(uint256)
func (_Main *MainCallerSession) CurrentReward() (*big.Int, error) {
	return _Main.Contract.CurrentReward(&_Main.CallOpts)
}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(uint256)
func (_Main *MainCaller) CurrentTotalTips(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "currentTotalTips")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(uint256)
func (_Main *MainSession) CurrentTotalTips() (*big.Int, error) {
	return _Main.Contract.CurrentTotalTips(&_Main.CallOpts)
}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(uint256)
func (_Main *MainCallerSession) CurrentTotalTips() (*big.Int, error) {
	return _Main.Contract.CurrentTotalTips(&_Main.CallOpts)
}

// DataIDs is a free data retrieval call binding the contract method 0x351f9a50.
//
// Solidity: function dataIDs(uint256 ) view returns(uint256 id, string name, uint256 granularity)
func (_Main *MainCaller) DataIDs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id          *big.Int
	Name        string
	Granularity *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "dataIDs", arg0)

	outstruct := new(struct {
		Id          *big.Int
		Name        string
		Granularity *big.Int
	})

	outstruct.Id = out[0].(*big.Int)
	outstruct.Name = out[1].(string)
	outstruct.Granularity = out[2].(*big.Int)

	return *outstruct, err

}

// DataIDs is a free data retrieval call binding the contract method 0x351f9a50.
//
// Solidity: function dataIDs(uint256 ) view returns(uint256 id, string name, uint256 granularity)
func (_Main *MainSession) DataIDs(arg0 *big.Int) (struct {
	Id          *big.Int
	Name        string
	Granularity *big.Int
}, error) {
	return _Main.Contract.DataIDs(&_Main.CallOpts, arg0)
}

// DataIDs is a free data retrieval call binding the contract method 0x351f9a50.
//
// Solidity: function dataIDs(uint256 ) view returns(uint256 id, string name, uint256 granularity)
func (_Main *MainCallerSession) DataIDs(arg0 *big.Int) (struct {
	Id          *big.Int
	Name        string
	Granularity *big.Int
}, error) {
	return _Main.Contract.DataIDs(&_Main.CallOpts, arg0)
}

// DataIDsAll is a free data retrieval call binding the contract method 0xa2c9a952.
//
// Solidity: function dataIDsAll() view returns((uint256,string,uint256)[])
func (_Main *MainCaller) DataIDsAll(opts *bind.CallOpts) ([]MainDataID, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "dataIDsAll")

	if err != nil {
		return *new([]MainDataID), err
	}

	out0 := *abi.ConvertType(out[0], new([]MainDataID)).(*[]MainDataID)

	return out0, err

}

// DataIDsAll is a free data retrieval call binding the contract method 0xa2c9a952.
//
// Solidity: function dataIDsAll() view returns((uint256,string,uint256)[])
func (_Main *MainSession) DataIDsAll() ([]MainDataID, error) {
	return _Main.Contract.DataIDsAll(&_Main.CallOpts)
}

// DataIDsAll is a free data retrieval call binding the contract method 0xa2c9a952.
//
// Solidity: function dataIDsAll() view returns((uint256,string,uint256)[])
func (_Main *MainCallerSession) DataIDsAll() ([]MainDataID, error) {
	return _Main.Contract.DataIDsAll(&_Main.CallOpts)
}

// DataIDsMap is a free data retrieval call binding the contract method 0x3a918872.
//
// Solidity: function dataIDsMap(uint256 ) view returns(uint256)
func (_Main *MainCaller) DataIDsMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "dataIDsMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DataIDsMap is a free data retrieval call binding the contract method 0x3a918872.
//
// Solidity: function dataIDsMap(uint256 ) view returns(uint256)
func (_Main *MainSession) DataIDsMap(arg0 *big.Int) (*big.Int, error) {
	return _Main.Contract.DataIDsMap(&_Main.CallOpts, arg0)
}

// DataIDsMap is a free data retrieval call binding the contract method 0x3a918872.
//
// Solidity: function dataIDsMap(uint256 ) view returns(uint256)
func (_Main *MainCallerSession) DataIDsMap(arg0 *big.Int) (*big.Int, error) {
	return _Main.Contract.DataIDsMap(&_Main.CallOpts, arg0)
}

// Deity is a free data retrieval call binding the contract method 0x0426ce0f.
//
// Solidity: function deity() view returns(address)
func (_Main *MainCaller) Deity(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "deity")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Deity is a free data retrieval call binding the contract method 0x0426ce0f.
//
// Solidity: function deity() view returns(address)
func (_Main *MainSession) Deity() (common.Address, error) {
	return _Main.Contract.Deity(&_Main.CallOpts)
}

// Deity is a free data retrieval call binding the contract method 0x0426ce0f.
//
// Solidity: function deity() view returns(address)
func (_Main *MainCallerSession) Deity() (common.Address, error) {
	return _Main.Contract.Deity(&_Main.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_Main *MainCaller) Difficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "difficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_Main *MainSession) Difficulty() (*big.Int, error) {
	return _Main.Contract.Difficulty(&_Main.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_Main *MainCallerSession) Difficulty() (*big.Int, error) {
	return _Main.Contract.Difficulty(&_Main.CallOpts)
}

// DisputeCount is a free data retrieval call binding the contract method 0xa28889e1.
//
// Solidity: function disputeCount() view returns(uint256)
func (_Main *MainCaller) DisputeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "disputeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputeCount is a free data retrieval call binding the contract method 0xa28889e1.
//
// Solidity: function disputeCount() view returns(uint256)
func (_Main *MainSession) DisputeCount() (*big.Int, error) {
	return _Main.Contract.DisputeCount(&_Main.CallOpts)
}

// DisputeCount is a free data retrieval call binding the contract method 0xa28889e1.
//
// Solidity: function disputeCount() view returns(uint256)
func (_Main *MainCallerSession) DisputeCount() (*big.Int, error) {
	return _Main.Contract.DisputeCount(&_Main.CallOpts)
}

// DisputeFee is a free data retrieval call binding the contract method 0xb9ce896b.
//
// Solidity: function disputeFee() view returns(uint256)
func (_Main *MainCaller) DisputeFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "disputeFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputeFee is a free data retrieval call binding the contract method 0xb9ce896b.
//
// Solidity: function disputeFee() view returns(uint256)
func (_Main *MainSession) DisputeFee() (*big.Int, error) {
	return _Main.Contract.DisputeFee(&_Main.CallOpts)
}

// DisputeFee is a free data retrieval call binding the contract method 0xb9ce896b.
//
// Solidity: function disputeFee() view returns(uint256)
func (_Main *MainCallerSession) DisputeFee() (*big.Int, error) {
	return _Main.Contract.DisputeFee(&_Main.CallOpts)
}

// GetCurrentValue is a free data retrieval call binding the contract method 0x3fcad964.
//
// Solidity: function getCurrentValue(uint256 _requestId) view returns(bool ifRetrieve, uint256 value, uint256 _timestampRetrieved)
func (_Main *MainCaller) GetCurrentValue(opts *bind.CallOpts, _requestId *big.Int) (struct {
	IfRetrieve         bool
	Value              *big.Int
	TimestampRetrieved *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getCurrentValue", _requestId)

	outstruct := new(struct {
		IfRetrieve         bool
		Value              *big.Int
		TimestampRetrieved *big.Int
	})

	outstruct.IfRetrieve = out[0].(bool)
	outstruct.Value = out[1].(*big.Int)
	outstruct.TimestampRetrieved = out[2].(*big.Int)

	return *outstruct, err

}

// GetCurrentValue is a free data retrieval call binding the contract method 0x3fcad964.
//
// Solidity: function getCurrentValue(uint256 _requestId) view returns(bool ifRetrieve, uint256 value, uint256 _timestampRetrieved)
func (_Main *MainSession) GetCurrentValue(_requestId *big.Int) (struct {
	IfRetrieve         bool
	Value              *big.Int
	TimestampRetrieved *big.Int
}, error) {
	return _Main.Contract.GetCurrentValue(&_Main.CallOpts, _requestId)
}

// GetCurrentValue is a free data retrieval call binding the contract method 0x3fcad964.
//
// Solidity: function getCurrentValue(uint256 _requestId) view returns(bool ifRetrieve, uint256 value, uint256 _timestampRetrieved)
func (_Main *MainCallerSession) GetCurrentValue(_requestId *big.Int) (struct {
	IfRetrieve         bool
	Value              *big.Int
	TimestampRetrieved *big.Int
}, error) {
	return _Main.Contract.GetCurrentValue(&_Main.CallOpts, _requestId)
}

// GetDataBefore is a free data retrieval call binding the contract method 0x66b44611.
//
// Solidity: function getDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool _ifRetrieve, uint256 _value, uint256 _timestampRetrieved)
func (_Main *MainCaller) GetDataBefore(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (struct {
	IfRetrieve         bool
	Value              *big.Int
	TimestampRetrieved *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getDataBefore", _requestId, _timestamp)

	outstruct := new(struct {
		IfRetrieve         bool
		Value              *big.Int
		TimestampRetrieved *big.Int
	})

	outstruct.IfRetrieve = out[0].(bool)
	outstruct.Value = out[1].(*big.Int)
	outstruct.TimestampRetrieved = out[2].(*big.Int)

	return *outstruct, err

}

// GetDataBefore is a free data retrieval call binding the contract method 0x66b44611.
//
// Solidity: function getDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool _ifRetrieve, uint256 _value, uint256 _timestampRetrieved)
func (_Main *MainSession) GetDataBefore(_requestId *big.Int, _timestamp *big.Int) (struct {
	IfRetrieve         bool
	Value              *big.Int
	TimestampRetrieved *big.Int
}, error) {
	return _Main.Contract.GetDataBefore(&_Main.CallOpts, _requestId, _timestamp)
}

// GetDataBefore is a free data retrieval call binding the contract method 0x66b44611.
//
// Solidity: function getDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool _ifRetrieve, uint256 _value, uint256 _timestampRetrieved)
func (_Main *MainCallerSession) GetDataBefore(_requestId *big.Int, _timestamp *big.Int) (struct {
	IfRetrieve         bool
	Value              *big.Int
	TimestampRetrieved *big.Int
}, error) {
	return _Main.Contract.GetDataBefore(&_Main.CallOpts, _requestId, _timestamp)
}

// GetIndexForDataBefore is a free data retrieval call binding the contract method 0xb73e4979.
//
// Solidity: function getIndexForDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool found, uint256 index)
func (_Main *MainCaller) GetIndexForDataBefore(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (struct {
	Found bool
	Index *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getIndexForDataBefore", _requestId, _timestamp)

	outstruct := new(struct {
		Found bool
		Index *big.Int
	})

	outstruct.Found = out[0].(bool)
	outstruct.Index = out[1].(*big.Int)

	return *outstruct, err

}

// GetIndexForDataBefore is a free data retrieval call binding the contract method 0xb73e4979.
//
// Solidity: function getIndexForDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool found, uint256 index)
func (_Main *MainSession) GetIndexForDataBefore(_requestId *big.Int, _timestamp *big.Int) (struct {
	Found bool
	Index *big.Int
}, error) {
	return _Main.Contract.GetIndexForDataBefore(&_Main.CallOpts, _requestId, _timestamp)
}

// GetIndexForDataBefore is a free data retrieval call binding the contract method 0xb73e4979.
//
// Solidity: function getIndexForDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool found, uint256 index)
func (_Main *MainCallerSession) GetIndexForDataBefore(_requestId *big.Int, _timestamp *big.Int) (struct {
	Found bool
	Index *big.Int
}, error) {
	return _Main.Contract.GetIndexForDataBefore(&_Main.CallOpts, _requestId, _timestamp)
}

// GetLastValues is a free data retrieval call binding the contract method 0xa6fdc28c.
//
// Solidity: function getLastValues(uint256 _dataID, uint256 _count) view returns(((uint256,string,uint256),uint256,uint256,uint256)[])
func (_Main *MainCaller) GetLastValues(opts *bind.CallOpts, _dataID *big.Int, _count *big.Int) ([]MainValue, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getLastValues", _dataID, _count)

	if err != nil {
		return *new([]MainValue), err
	}

	out0 := *abi.ConvertType(out[0], new([]MainValue)).(*[]MainValue)

	return out0, err

}

// GetLastValues is a free data retrieval call binding the contract method 0xa6fdc28c.
//
// Solidity: function getLastValues(uint256 _dataID, uint256 _count) view returns(((uint256,string,uint256),uint256,uint256,uint256)[])
func (_Main *MainSession) GetLastValues(_dataID *big.Int, _count *big.Int) ([]MainValue, error) {
	return _Main.Contract.GetLastValues(&_Main.CallOpts, _dataID, _count)
}

// GetLastValues is a free data retrieval call binding the contract method 0xa6fdc28c.
//
// Solidity: function getLastValues(uint256 _dataID, uint256 _count) view returns(((uint256,string,uint256),uint256,uint256,uint256)[])
func (_Main *MainCallerSession) GetLastValues(_dataID *big.Int, _count *big.Int) ([]MainValue, error) {
	return _Main.Contract.GetLastValues(&_Main.CallOpts, _dataID, _count)
}

// GetLastValuesAll is a free data retrieval call binding the contract method 0x8b7a0e49.
//
// Solidity: function getLastValuesAll(uint256 count) view returns(((uint256,string,uint256),uint256,uint256,uint256)[])
func (_Main *MainCaller) GetLastValuesAll(opts *bind.CallOpts, count *big.Int) ([]MainValue, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getLastValuesAll", count)

	if err != nil {
		return *new([]MainValue), err
	}

	out0 := *abi.ConvertType(out[0], new([]MainValue)).(*[]MainValue)

	return out0, err

}

// GetLastValuesAll is a free data retrieval call binding the contract method 0x8b7a0e49.
//
// Solidity: function getLastValuesAll(uint256 count) view returns(((uint256,string,uint256),uint256,uint256,uint256)[])
func (_Main *MainSession) GetLastValuesAll(count *big.Int) ([]MainValue, error) {
	return _Main.Contract.GetLastValuesAll(&_Main.CallOpts, count)
}

// GetLastValuesAll is a free data retrieval call binding the contract method 0x8b7a0e49.
//
// Solidity: function getLastValuesAll(uint256 count) view returns(((uint256,string,uint256),uint256,uint256,uint256)[])
func (_Main *MainCallerSession) GetLastValuesAll(count *big.Int) ([]MainValue, error) {
	return _Main.Contract.GetLastValuesAll(&_Main.CallOpts, count)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_Main *MainCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getNewValueCountbyRequestId", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_Main *MainSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _Main.Contract.GetNewValueCountbyRequestId(&_Main.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_Main *MainCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _Main.Contract.GetNewValueCountbyRequestId(&_Main.CallOpts, _requestId)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 _index) view returns(uint256)
func (_Main *MainCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestId *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getTimestampbyRequestIDandIndex", _requestId, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 _index) view returns(uint256)
func (_Main *MainSession) GetTimestampbyRequestIDandIndex(_requestId *big.Int, _index *big.Int) (*big.Int, error) {
	return _Main.Contract.GetTimestampbyRequestIDandIndex(&_Main.CallOpts, _requestId, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 _index) view returns(uint256)
func (_Main *MainCallerSession) GetTimestampbyRequestIDandIndex(_requestId *big.Int, _index *big.Int) (*big.Int, error) {
	return _Main.Contract.GetTimestampbyRequestIDandIndex(&_Main.CallOpts, _requestId, _index)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_Main *MainCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "isInDispute", _requestId, _timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_Main *MainSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _Main.Contract.IsInDispute(&_Main.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_Main *MainCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _Main.Contract.IsInDispute(&_Main.CallOpts, _requestId, _timestamp)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Main *MainCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Main *MainSession) Oracle() (common.Address, error) {
	return _Main.Contract.Oracle(&_Main.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Main *MainCallerSession) Oracle() (common.Address, error) {
	return _Main.Contract.Oracle(&_Main.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainSession) Owner() (common.Address, error) {
	return _Main.Contract.Owner(&_Main.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainCallerSession) Owner() (common.Address, error) {
	return _Main.Contract.Owner(&_Main.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Main *MainCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Main *MainSession) PendingOwner() (common.Address, error) {
	return _Main.Contract.PendingOwner(&_Main.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Main *MainCallerSession) PendingOwner() (common.Address, error) {
	return _Main.Contract.PendingOwner(&_Main.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(uint256)
func (_Main *MainCaller) RequestCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "requestCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(uint256)
func (_Main *MainSession) RequestCount() (*big.Int, error) {
	return _Main.Contract.RequestCount(&_Main.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(uint256)
func (_Main *MainCallerSession) RequestCount() (*big.Int, error) {
	return _Main.Contract.RequestCount(&_Main.CallOpts)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_Main *MainCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "retrieveData", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_Main *MainSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _Main.Contract.RetrieveData(&_Main.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_Main *MainCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _Main.Contract.RetrieveData(&_Main.CallOpts, _requestId, _timestamp)
}

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(uint256)
func (_Main *MainCaller) SlotProgress(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "slotProgress")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(uint256)
func (_Main *MainSession) SlotProgress() (*big.Int, error) {
	return _Main.Contract.SlotProgress(&_Main.CallOpts)
}

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(uint256)
func (_Main *MainCallerSession) SlotProgress() (*big.Int, error) {
	return _Main.Contract.SlotProgress(&_Main.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(uint256)
func (_Main *MainCaller) StakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "stakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(uint256)
func (_Main *MainSession) StakeAmount() (*big.Int, error) {
	return _Main.Contract.StakeAmount(&_Main.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(uint256)
func (_Main *MainCallerSession) StakeAmount() (*big.Int, error) {
	return _Main.Contract.StakeAmount(&_Main.CallOpts)
}

// StakeCount is a free data retrieval call binding the contract method 0xc4a9e116.
//
// Solidity: function stakeCount() view returns(uint256)
func (_Main *MainCaller) StakeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "stakeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeCount is a free data retrieval call binding the contract method 0xc4a9e116.
//
// Solidity: function stakeCount() view returns(uint256)
func (_Main *MainSession) StakeCount() (*big.Int, error) {
	return _Main.Contract.StakeCount(&_Main.CallOpts)
}

// StakeCount is a free data retrieval call binding the contract method 0xc4a9e116.
//
// Solidity: function stakeCount() view returns(uint256)
func (_Main *MainCallerSession) StakeCount() (*big.Int, error) {
	return _Main.Contract.StakeCount(&_Main.CallOpts)
}

// TBlock is a free data retrieval call binding the contract method 0x6747dc31.
//
// Solidity: function tBlock() view returns(uint256)
func (_Main *MainCaller) TBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "tBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TBlock is a free data retrieval call binding the contract method 0x6747dc31.
//
// Solidity: function tBlock() view returns(uint256)
func (_Main *MainSession) TBlock() (*big.Int, error) {
	return _Main.Contract.TBlock(&_Main.CallOpts)
}

// TBlock is a free data retrieval call binding the contract method 0x6747dc31.
//
// Solidity: function tBlock() view returns(uint256)
func (_Main *MainCallerSession) TBlock() (*big.Int, error) {
	return _Main.Contract.TBlock(&_Main.CallOpts)
}

// TellorContract is a free data retrieval call binding the contract method 0xa339ac74.
//
// Solidity: function tellorContract() view returns(address)
func (_Main *MainCaller) TellorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "tellorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TellorContract is a free data retrieval call binding the contract method 0xa339ac74.
//
// Solidity: function tellorContract() view returns(address)
func (_Main *MainSession) TellorContract() (common.Address, error) {
	return _Main.Contract.TellorContract(&_Main.CallOpts)
}

// TellorContract is a free data retrieval call binding the contract method 0xa339ac74.
//
// Solidity: function tellorContract() view returns(address)
func (_Main *MainCallerSession) TellorContract() (common.Address, error) {
	return _Main.Contract.TellorContract(&_Main.CallOpts)
}

// TimeOfLastValue is a free data retrieval call binding the contract method 0xdfcff498.
//
// Solidity: function timeOfLastValue() view returns(uint256)
func (_Main *MainCaller) TimeOfLastValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "timeOfLastValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeOfLastValue is a free data retrieval call binding the contract method 0xdfcff498.
//
// Solidity: function timeOfLastValue() view returns(uint256)
func (_Main *MainSession) TimeOfLastValue() (*big.Int, error) {
	return _Main.Contract.TimeOfLastValue(&_Main.CallOpts)
}

// TimeOfLastValue is a free data retrieval call binding the contract method 0xdfcff498.
//
// Solidity: function timeOfLastValue() view returns(uint256)
func (_Main *MainCallerSession) TimeOfLastValue() (*big.Int, error) {
	return _Main.Contract.TimeOfLastValue(&_Main.CallOpts)
}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(uint256)
func (_Main *MainCaller) TimeTarget(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "timeTarget")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(uint256)
func (_Main *MainSession) TimeTarget() (*big.Int, error) {
	return _Main.Contract.TimeTarget(&_Main.CallOpts)
}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(uint256)
func (_Main *MainCallerSession) TimeTarget() (*big.Int, error) {
	return _Main.Contract.TimeTarget(&_Main.CallOpts)
}

// TotalTip is a free data retrieval call binding the contract method 0x44b12aea.
//
// Solidity: function totalTip(uint256 _dataID) view returns(uint256)
func (_Main *MainCaller) TotalTip(opts *bind.CallOpts, _dataID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "totalTip", _dataID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTip is a free data retrieval call binding the contract method 0x44b12aea.
//
// Solidity: function totalTip(uint256 _dataID) view returns(uint256)
func (_Main *MainSession) TotalTip(_dataID *big.Int) (*big.Int, error) {
	return _Main.Contract.TotalTip(&_Main.CallOpts, _dataID)
}

// TotalTip is a free data retrieval call binding the contract method 0x44b12aea.
//
// Solidity: function totalTip(uint256 _dataID) view returns(uint256)
func (_Main *MainCallerSession) TotalTip(_dataID *big.Int) (*big.Int, error) {
	return _Main.Contract.TotalTip(&_Main.CallOpts, _dataID)
}

// PushDataID is a paid mutator transaction binding the contract method 0x4d5ddb67.
//
// Solidity: function pushDataID((uint256,string,uint256) _dataID) returns()
func (_Main *MainTransactor) PushDataID(opts *bind.TransactOpts, _dataID MainDataID) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "pushDataID", _dataID)
}

// PushDataID is a paid mutator transaction binding the contract method 0x4d5ddb67.
//
// Solidity: function pushDataID((uint256,string,uint256) _dataID) returns()
func (_Main *MainSession) PushDataID(_dataID MainDataID) (*types.Transaction, error) {
	return _Main.Contract.PushDataID(&_Main.TransactOpts, _dataID)
}

// PushDataID is a paid mutator transaction binding the contract method 0x4d5ddb67.
//
// Solidity: function pushDataID((uint256,string,uint256) _dataID) returns()
func (_Main *MainTransactorSession) PushDataID(_dataID MainDataID) (*types.Transaction, error) {
	return _Main.Contract.PushDataID(&_Main.TransactOpts, _dataID)
}

// ReplaceDataIDs is a paid mutator transaction binding the contract method 0x6d7ecbd6.
//
// Solidity: function replaceDataIDs((uint256,string,uint256)[] _dataIDs) returns()
func (_Main *MainTransactor) ReplaceDataIDs(opts *bind.TransactOpts, _dataIDs []MainDataID) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "replaceDataIDs", _dataIDs)
}

// ReplaceDataIDs is a paid mutator transaction binding the contract method 0x6d7ecbd6.
//
// Solidity: function replaceDataIDs((uint256,string,uint256)[] _dataIDs) returns()
func (_Main *MainSession) ReplaceDataIDs(_dataIDs []MainDataID) (*types.Transaction, error) {
	return _Main.Contract.ReplaceDataIDs(&_Main.TransactOpts, _dataIDs)
}

// ReplaceDataIDs is a paid mutator transaction binding the contract method 0x6d7ecbd6.
//
// Solidity: function replaceDataIDs((uint256,string,uint256)[] _dataIDs) returns()
func (_Main *MainTransactorSession) ReplaceDataIDs(_dataIDs []MainDataID) (*types.Transaction, error) {
	return _Main.Contract.ReplaceDataIDs(&_Main.TransactOpts, _dataIDs)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_Main *MainTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_Main *MainSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetAdmin(&_Main.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_Main *MainTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetAdmin(&_Main.TransactOpts, _admin)
}

// SetDataID is a paid mutator transaction binding the contract method 0xeabe1566.
//
// Solidity: function setDataID(uint256 _id, (uint256,string,uint256) _dataID) returns()
func (_Main *MainTransactor) SetDataID(opts *bind.TransactOpts, _id *big.Int, _dataID MainDataID) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setDataID", _id, _dataID)
}

// SetDataID is a paid mutator transaction binding the contract method 0xeabe1566.
//
// Solidity: function setDataID(uint256 _id, (uint256,string,uint256) _dataID) returns()
func (_Main *MainSession) SetDataID(_id *big.Int, _dataID MainDataID) (*types.Transaction, error) {
	return _Main.Contract.SetDataID(&_Main.TransactOpts, _id, _dataID)
}

// SetDataID is a paid mutator transaction binding the contract method 0xeabe1566.
//
// Solidity: function setDataID(uint256 _id, (uint256,string,uint256) _dataID) returns()
func (_Main *MainTransactorSession) SetDataID(_id *big.Int, _dataID MainDataID) (*types.Transaction, error) {
	return _Main.Contract.SetDataID(&_Main.TransactOpts, _id, _dataID)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Main *MainTransactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Main *MainSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetOracle(&_Main.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Main *MainTransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetOracle(&_Main.TransactOpts, _oracle)
}

// OracleABI is the input ABI used to generate the binding from.
const OracleABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OracleFuncSigs maps the 4-byte function signature to its string representation.
var OracleFuncSigs = map[string]string{
	"133bee5e": "getAddressVars(bytes32)",
	"46eee1c4": "getNewValueCountbyRequestId(uint256)",
	"e0ae93c1": "getRequestUintVars(uint256,bytes32)",
	"77fbb663": "getTimestampbyRequestIDandIndex(uint256,uint256)",
	"612c8f7f": "getUintVar(bytes32)",
	"93fa4915": "retrieveData(uint256,uint256)",
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

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_Oracle *OracleCaller) GetAddressVars(opts *bind.CallOpts, _data [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getAddressVars", _data)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_Oracle *OracleSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _Oracle.Contract.GetAddressVars(&_Oracle.CallOpts, _data)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_Oracle *OracleCallerSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _Oracle.Contract.GetAddressVars(&_Oracle.CallOpts, _data)
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

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_Oracle *OracleCaller) GetRequestUintVars(opts *bind.CallOpts, _requestId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getRequestUintVars", _requestId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_Oracle *OracleSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _Oracle.Contract.GetRequestUintVars(&_Oracle.CallOpts, _requestId, _data)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_Oracle *OracleCallerSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _Oracle.Contract.GetRequestUintVars(&_Oracle.CallOpts, _requestId, _data)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_Oracle *OracleCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getTimestampbyRequestIDandIndex", _requestID, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_Oracle *OracleSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _Oracle.Contract.GetTimestampbyRequestIDandIndex(&_Oracle.CallOpts, _requestID, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_Oracle *OracleCallerSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _Oracle.Contract.GetTimestampbyRequestIDandIndex(&_Oracle.CallOpts, _requestID, _index)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_Oracle *OracleCaller) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getUintVar", _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_Oracle *OracleSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _Oracle.Contract.GetUintVar(&_Oracle.CallOpts, _data)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_Oracle *OracleCallerSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _Oracle.Contract.GetUintVar(&_Oracle.CallOpts, _data)
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

// UsingTellorABI is the input ABI used to generate the binding from.
const UsingTellorABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_tellor\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getCurrentValue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ifRetrieve\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestampRetrieved\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getDataBefore\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_ifRetrieve\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestampRetrieved\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getIndexForDataBefore\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// UsingTellorFuncSigs maps the 4-byte function signature to its string representation.
var UsingTellorFuncSigs = map[string]string{
	"3fcad964": "getCurrentValue(uint256)",
	"66b44611": "getDataBefore(uint256,uint256)",
	"b73e4979": "getIndexForDataBefore(uint256,uint256)",
	"46eee1c4": "getNewValueCountbyRequestId(uint256)",
	"77fbb663": "getTimestampbyRequestIDandIndex(uint256,uint256)",
	"3df0777b": "isInDispute(uint256,uint256)",
	"93fa4915": "retrieveData(uint256,uint256)",
}

// UsingTellorBin is the compiled bytecode used for deploying new contracts.
var UsingTellorBin = "0x608060405234801561001057600080fd5b50604051610aed380380610aed8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055610a88806100656000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806366b446111161005b57806366b446111461012557806377fbb6631461014857806393fa49151461016b578063b73e49791461018e5761007d565b80633df0777b146100825780633fcad964146100b957806346eee1c4146100f6575b600080fd5b6100a56004803603604081101561009857600080fd5b50803590602001356101cc565b604080519115158252519081900360200190f35b6100d6600480360360208110156100cf57600080fd5b5035610253565b604080519315158452602084019290925282820152519081900360600190f35b6101136004803603602081101561010c57600080fd5b5035610411565b60408051918252519081900360200190f35b6100d66004803603604081101561013b57600080fd5b5080359060200135610490565b6101136004803603604081101561015e57600080fd5b50803590602001356105eb565b6101136004803603604081101561018157600080fd5b508035906020013561063f565b6101b1600480360360408110156101a457600080fd5b5080359060200135610693565b60408051921515835260208301919091528051918290030190f35b6000805460408051633df0777b60e01b8152600481018690526024810185905290516001600160a01b0390921691633df0777b91604480820192602092909190829003018186803b15801561022057600080fd5b505afa158015610234573d6000803e3d6000fd5b505050506040513d602081101561024a57600080fd5b50519392505050565b60008060008060008054906101000a90046001600160a01b03166001600160a01b03166346eee1c4866040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156102b057600080fd5b505afa1580156102c4573d6000803e3d6000fd5b505050506040513d60208110156102da57600080fd5b505160008054604080516377fbb66360e01b8152600481018a905260001985016024820152905193945091926001600160a01b03909116916377fbb663916044808301926020929190829003018186803b15801561033757600080fd5b505afa15801561034b573d6000803e3d6000fd5b505050506040513d602081101561036157600080fd5b505160008054604080516393fa491560e01b8152600481018b905260248101859052905193945091926001600160a01b03909116916393fa4915916044808301926020929190829003018186803b1580156103bb57600080fd5b505afa1580156103cf573d6000803e3d6000fd5b505050506040513d60208110156103e557600080fd5b5051905080156103fe57600195509350915061040a9050565b50600094508493509150505b9193909250565b60008054604080516311bbb87160e21b81526004810185905290516001600160a01b03909216916346eee1c491602480820192602092909190829003018186803b15801561045e57600080fd5b505afa158015610472573d6000803e3d6000fd5b505050506040513d602081101561048857600080fd5b505192915050565b60008060008060006104a28787610693565b91509150816104bd57600080600094509450945050506105e4565b60008054604080516377fbb66360e01b8152600481018b90526024810185905290516001600160a01b03909216916377fbb66391604480820192602092909190829003018186803b15801561051157600080fd5b505afa158015610525573d6000803e3d6000fd5b505050506040513d602081101561053b57600080fd5b5051600054604080516393fa491560e01b8152600481018c90526024810184905290519293506001600160a01b03909116916393fa491591604480820192602092909190829003018186803b15801561059357600080fd5b505afa1580156105a7573d6000803e3d6000fd5b505050506040513d60208110156105bd57600080fd5b5051945084156105d5576001955092506105e4915050565b60008060009550955095505050505b9250925092565b60008054604080516377fbb66360e01b8152600481018690526024810185905290516001600160a01b03909216916377fbb66391604480820192602092909190829003018186803b15801561022057600080fd5b60008054604080516393fa491560e01b8152600481018690526024810185905290516001600160a01b03909216916393fa491591604480820192602092909190829003018186803b15801561022057600080fd5b60008054604080516311bbb87160e21b8152600481018690529051839283926001600160a01b03909116916346eee1c491602480820192602092909190829003018186803b1580156106e457600080fd5b505afa1580156106f8573d6000803e3d6000fd5b505050506040513d602081101561070e57600080fd5b505190508015610a425760008054604080516377fbb66360e01b815260048101899052602481018490529051839260001986019284926001600160a01b03909216916377fbb66391604480820192602092909190829003018186803b15801561077657600080fd5b505afa15801561078a573d6000803e3d6000fd5b505050506040513d60208110156107a057600080fd5b505190508781106107bc57600080965096505050505050610a4b565b600054604080516377fbb66360e01b8152600481018c90526024810185905290516001600160a01b03909216916377fbb66391604480820192602092909190829003018186803b15801561080f57600080fd5b505afa158015610823573d6000803e3d6000fd5b505050506040513d602081101561083957600080fd5b50519050878110156108555750600195509350610a4b92505050565b600054604080516377fbb66360e01b8152600481018c905260028686030486016001016024820181905291519196506001600160a01b03909216916377fbb663916044808301926020929190829003018186803b1580156108b557600080fd5b505afa1580156108c9573d6000803e3d6000fd5b505050506040513d60208110156108df57600080fd5b50519050878110156109945760008054604080516377fbb66360e01b8152600481018d905260018801602482015290516001600160a01b03909216916377fbb66391604480820192602092909190829003018186803b15801561094157600080fd5b505afa158015610955573d6000803e3d6000fd5b505050506040513d602081101561096b57600080fd5b505190508881106109885760018597509750505050505050610a4b565b84600101935050610a3d565b60008054604080516377fbb66360e01b8152600481018d90526000198801602482015290516001600160a01b03909216916377fbb66391604480820192602092909190829003018186803b1580156109eb57600080fd5b505afa1580156109ff573d6000803e3d6000fd5b505050506040513d6020811015610a1557600080fd5b5051905088811015610a3557600180860397509750505050505050610a4b565b600185039250505b610855565b60008092509250505b925092905056fea264697066735822122023954ef1f1663f12808c3f6b86395641e83de020bd3b50b2874bb5163dfc4c0264736f6c63430007060033"

// DeployUsingTellor deploys a new Ethereum contract, binding an instance of UsingTellor to it.
func DeployUsingTellor(auth *bind.TransactOpts, backend bind.ContractBackend, _tellor common.Address) (common.Address, *types.Transaction, *UsingTellor, error) {
	parsed, err := abi.JSON(strings.NewReader(UsingTellorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UsingTellorBin), backend, _tellor)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UsingTellor{UsingTellorCaller: UsingTellorCaller{contract: contract}, UsingTellorTransactor: UsingTellorTransactor{contract: contract}, UsingTellorFilterer: UsingTellorFilterer{contract: contract}}, nil
}

// UsingTellor is an auto generated Go binding around an Ethereum contract.
type UsingTellor struct {
	UsingTellorCaller     // Read-only binding to the contract
	UsingTellorTransactor // Write-only binding to the contract
	UsingTellorFilterer   // Log filterer for contract events
}

// UsingTellorCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsingTellorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsingTellorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsingTellorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsingTellorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsingTellorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsingTellorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsingTellorSession struct {
	Contract     *UsingTellor      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsingTellorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsingTellorCallerSession struct {
	Contract *UsingTellorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// UsingTellorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsingTellorTransactorSession struct {
	Contract     *UsingTellorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UsingTellorRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsingTellorRaw struct {
	Contract *UsingTellor // Generic contract binding to access the raw methods on
}

// UsingTellorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsingTellorCallerRaw struct {
	Contract *UsingTellorCaller // Generic read-only contract binding to access the raw methods on
}

// UsingTellorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsingTellorTransactorRaw struct {
	Contract *UsingTellorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsingTellor creates a new instance of UsingTellor, bound to a specific deployed contract.
func NewUsingTellor(address common.Address, backend bind.ContractBackend) (*UsingTellor, error) {
	contract, err := bindUsingTellor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UsingTellor{UsingTellorCaller: UsingTellorCaller{contract: contract}, UsingTellorTransactor: UsingTellorTransactor{contract: contract}, UsingTellorFilterer: UsingTellorFilterer{contract: contract}}, nil
}

// NewUsingTellorCaller creates a new read-only instance of UsingTellor, bound to a specific deployed contract.
func NewUsingTellorCaller(address common.Address, caller bind.ContractCaller) (*UsingTellorCaller, error) {
	contract, err := bindUsingTellor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsingTellorCaller{contract: contract}, nil
}

// NewUsingTellorTransactor creates a new write-only instance of UsingTellor, bound to a specific deployed contract.
func NewUsingTellorTransactor(address common.Address, transactor bind.ContractTransactor) (*UsingTellorTransactor, error) {
	contract, err := bindUsingTellor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsingTellorTransactor{contract: contract}, nil
}

// NewUsingTellorFilterer creates a new log filterer instance of UsingTellor, bound to a specific deployed contract.
func NewUsingTellorFilterer(address common.Address, filterer bind.ContractFilterer) (*UsingTellorFilterer, error) {
	contract, err := bindUsingTellor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsingTellorFilterer{contract: contract}, nil
}

// bindUsingTellor binds a generic wrapper to an already deployed contract.
func bindUsingTellor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UsingTellorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsingTellor *UsingTellorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UsingTellor.Contract.UsingTellorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsingTellor *UsingTellorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsingTellor.Contract.UsingTellorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsingTellor *UsingTellorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsingTellor.Contract.UsingTellorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsingTellor *UsingTellorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UsingTellor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsingTellor *UsingTellorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsingTellor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsingTellor *UsingTellorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsingTellor.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentValue is a free data retrieval call binding the contract method 0x3fcad964.
//
// Solidity: function getCurrentValue(uint256 _requestId) view returns(bool ifRetrieve, uint256 value, uint256 _timestampRetrieved)
func (_UsingTellor *UsingTellorCaller) GetCurrentValue(opts *bind.CallOpts, _requestId *big.Int) (struct {
	IfRetrieve         bool
	Value              *big.Int
	TimestampRetrieved *big.Int
}, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "getCurrentValue", _requestId)

	outstruct := new(struct {
		IfRetrieve         bool
		Value              *big.Int
		TimestampRetrieved *big.Int
	})

	outstruct.IfRetrieve = out[0].(bool)
	outstruct.Value = out[1].(*big.Int)
	outstruct.TimestampRetrieved = out[2].(*big.Int)

	return *outstruct, err

}

// GetCurrentValue is a free data retrieval call binding the contract method 0x3fcad964.
//
// Solidity: function getCurrentValue(uint256 _requestId) view returns(bool ifRetrieve, uint256 value, uint256 _timestampRetrieved)
func (_UsingTellor *UsingTellorSession) GetCurrentValue(_requestId *big.Int) (struct {
	IfRetrieve         bool
	Value              *big.Int
	TimestampRetrieved *big.Int
}, error) {
	return _UsingTellor.Contract.GetCurrentValue(&_UsingTellor.CallOpts, _requestId)
}

// GetCurrentValue is a free data retrieval call binding the contract method 0x3fcad964.
//
// Solidity: function getCurrentValue(uint256 _requestId) view returns(bool ifRetrieve, uint256 value, uint256 _timestampRetrieved)
func (_UsingTellor *UsingTellorCallerSession) GetCurrentValue(_requestId *big.Int) (struct {
	IfRetrieve         bool
	Value              *big.Int
	TimestampRetrieved *big.Int
}, error) {
	return _UsingTellor.Contract.GetCurrentValue(&_UsingTellor.CallOpts, _requestId)
}

// GetDataBefore is a free data retrieval call binding the contract method 0x66b44611.
//
// Solidity: function getDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool _ifRetrieve, uint256 _value, uint256 _timestampRetrieved)
func (_UsingTellor *UsingTellorCaller) GetDataBefore(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (struct {
	IfRetrieve         bool
	Value              *big.Int
	TimestampRetrieved *big.Int
}, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "getDataBefore", _requestId, _timestamp)

	outstruct := new(struct {
		IfRetrieve         bool
		Value              *big.Int
		TimestampRetrieved *big.Int
	})

	outstruct.IfRetrieve = out[0].(bool)
	outstruct.Value = out[1].(*big.Int)
	outstruct.TimestampRetrieved = out[2].(*big.Int)

	return *outstruct, err

}

// GetDataBefore is a free data retrieval call binding the contract method 0x66b44611.
//
// Solidity: function getDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool _ifRetrieve, uint256 _value, uint256 _timestampRetrieved)
func (_UsingTellor *UsingTellorSession) GetDataBefore(_requestId *big.Int, _timestamp *big.Int) (struct {
	IfRetrieve         bool
	Value              *big.Int
	TimestampRetrieved *big.Int
}, error) {
	return _UsingTellor.Contract.GetDataBefore(&_UsingTellor.CallOpts, _requestId, _timestamp)
}

// GetDataBefore is a free data retrieval call binding the contract method 0x66b44611.
//
// Solidity: function getDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool _ifRetrieve, uint256 _value, uint256 _timestampRetrieved)
func (_UsingTellor *UsingTellorCallerSession) GetDataBefore(_requestId *big.Int, _timestamp *big.Int) (struct {
	IfRetrieve         bool
	Value              *big.Int
	TimestampRetrieved *big.Int
}, error) {
	return _UsingTellor.Contract.GetDataBefore(&_UsingTellor.CallOpts, _requestId, _timestamp)
}

// GetIndexForDataBefore is a free data retrieval call binding the contract method 0xb73e4979.
//
// Solidity: function getIndexForDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool found, uint256 index)
func (_UsingTellor *UsingTellorCaller) GetIndexForDataBefore(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (struct {
	Found bool
	Index *big.Int
}, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "getIndexForDataBefore", _requestId, _timestamp)

	outstruct := new(struct {
		Found bool
		Index *big.Int
	})

	outstruct.Found = out[0].(bool)
	outstruct.Index = out[1].(*big.Int)

	return *outstruct, err

}

// GetIndexForDataBefore is a free data retrieval call binding the contract method 0xb73e4979.
//
// Solidity: function getIndexForDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool found, uint256 index)
func (_UsingTellor *UsingTellorSession) GetIndexForDataBefore(_requestId *big.Int, _timestamp *big.Int) (struct {
	Found bool
	Index *big.Int
}, error) {
	return _UsingTellor.Contract.GetIndexForDataBefore(&_UsingTellor.CallOpts, _requestId, _timestamp)
}

// GetIndexForDataBefore is a free data retrieval call binding the contract method 0xb73e4979.
//
// Solidity: function getIndexForDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool found, uint256 index)
func (_UsingTellor *UsingTellorCallerSession) GetIndexForDataBefore(_requestId *big.Int, _timestamp *big.Int) (struct {
	Found bool
	Index *big.Int
}, error) {
	return _UsingTellor.Contract.GetIndexForDataBefore(&_UsingTellor.CallOpts, _requestId, _timestamp)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_UsingTellor *UsingTellorCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "getNewValueCountbyRequestId", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_UsingTellor *UsingTellorSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _UsingTellor.Contract.GetNewValueCountbyRequestId(&_UsingTellor.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_UsingTellor *UsingTellorCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _UsingTellor.Contract.GetNewValueCountbyRequestId(&_UsingTellor.CallOpts, _requestId)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 _index) view returns(uint256)
func (_UsingTellor *UsingTellorCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestId *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "getTimestampbyRequestIDandIndex", _requestId, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 _index) view returns(uint256)
func (_UsingTellor *UsingTellorSession) GetTimestampbyRequestIDandIndex(_requestId *big.Int, _index *big.Int) (*big.Int, error) {
	return _UsingTellor.Contract.GetTimestampbyRequestIDandIndex(&_UsingTellor.CallOpts, _requestId, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 _index) view returns(uint256)
func (_UsingTellor *UsingTellorCallerSession) GetTimestampbyRequestIDandIndex(_requestId *big.Int, _index *big.Int) (*big.Int, error) {
	return _UsingTellor.Contract.GetTimestampbyRequestIDandIndex(&_UsingTellor.CallOpts, _requestId, _index)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_UsingTellor *UsingTellorCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "isInDispute", _requestId, _timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_UsingTellor *UsingTellorSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _UsingTellor.Contract.IsInDispute(&_UsingTellor.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_UsingTellor *UsingTellorCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _UsingTellor.Contract.IsInDispute(&_UsingTellor.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_UsingTellor *UsingTellorCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "retrieveData", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_UsingTellor *UsingTellorSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _UsingTellor.Contract.RetrieveData(&_UsingTellor.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_UsingTellor *UsingTellorCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _UsingTellor.Contract.RetrieveData(&_UsingTellor.CallOpts, _requestId, _timestamp)
}

// ConsoleABI is the input ABI used to generate the binding from.
const ConsoleABI = "[]"

// ConsoleBin is the compiled bytecode used for deploying new contracts.
var ConsoleBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204f87b274f63f544baa5f27a5bb676428981317e89a2041a6effe042b7864a76064736f6c63430007060033"

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
