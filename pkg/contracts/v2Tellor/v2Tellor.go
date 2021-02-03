// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package v2Tellor

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

// V2TellorABI is the input ABI used to generate the binding from.
const V2TellorABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minerIndex\",\"type\":\"uint256\"}],\"name\":\"beginDispute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_supportsDispute\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"tallyVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_propNewTellorAddress\",\"type\":\"address\"}],\"name\":\"proposeFork\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"}],\"name\":\"submitMiningSolution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"submitMiningSolution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_pendingOwner\",\"type\":\"address\"}],\"name\":\"proposeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"requestStakingWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNewCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256\",\"name\":\"_difficutly\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTopRequestIDs\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNewVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"idsOnDeck\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"tipsOnDeck\",\"type\":\"uint256[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"updateTellor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"unlockDisputeFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// V2TellorBin is the compiled bytecode used for deploying new contracts.
var V2TellorBin = "0x608060405234801561001057600080fd5b506117ae806100206000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c806368c180d5116100c35780639a7077ab1161007c5780639a7077ab1461065a578063a9059cbb146106cf578063bed9d86114610735578063c9d27afe1461073f578063f458ab9814610779578063fe1cd15d146107a75761014d565b806368c180d51461045e578063710bf322146104eb578063752d49a11461052f5780638581af191461056757806395d89b41146105a95780639a01ca131461062c5761014d565b806328449c3a1161011557806328449c3a1461030f578063313ce567146103195780634049f1981461033d5780634350283e146103985780634d318b0e146104265780634e71e0c8146104545761014d565b806306fdde0314610152578063095ea7b3146101d55780630d2d76a21461023b57806323b872dd1461024557806326b7d9f6146102cb575b600080fd5b61015a6107ed565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561019a57808201518184015260208101905061017f565b50505050905090810190601f1680156101c75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610221600480360360408110156101eb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061082a565b604051808215151515815260200191505060405180910390f35b6102436108fc565b005b6102b16004803603606081101561025b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610968565b604051808215151515815260200191505060405180910390f35b61030d600480360360208110156102e157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a6f565b005b610317610b10565b005b610321610b7c565b604051808260ff1660ff16815260200191505060405180910390f35b610345610b85565b6040518085815260200184600560200280838360005b8381101561037657808201518184015260208101905061035b565b5050505090500183815260200182815260200194505050505060405180910390f35b61042460048036036101608110156103af57600080fd5b81019080803590602001906401000000008111156103cc57600080fd5b8201836020820111156103de57600080fd5b8035906020019184600183028401116401000000008311171561040057600080fd5b90919293919293908060a0019091929192908060a001909192919290505050610baa565b005b6104526004803603602081101561043c57600080fd5b8101908080359060200190929190505050610c94565b005b61045c610d09565b005b6104e96004803603606081101561047457600080fd5b810190808035906020019064010000000081111561049157600080fd5b8201836020820111156104a357600080fd5b803590602001918460018302840111640100000000831117156104c557600080fd5b90919293919293908035906020019092919080359060200190929190505050610d75565b005b61052d6004803603602081101561050157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e29565b005b6105656004803603604081101561054557600080fd5b810190808035906020019092919080359060200190929190505050610eca565b005b6105a76004803603606081101561057d57600080fd5b81019080803590602001909291908035906020019092919080359060200190929190505050610f48565b005b6105b1610fcf565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156105f15780820151818401526020810190506105d6565b50505050905090810190601f16801561061e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6106586004803603602081101561064257600080fd5b810190808035906020019092919050505061100c565b005b610662611081565b6040518083600560200280838360005b8381101561068d578082015181840152602081019050610672565b5050505090500182600560200280838360005b838110156106bb5780820151818401526020810190506106a0565b505050509050019250505060405180910390f35b61071b600480360360408110156106e557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506110a3565b604051808215151515815260200191505060405180910390f35b61073d611175565b005b6107776004803603604081101561075557600080fd5b81019080803590602001909291908035151590602001909291905050506111e1565b005b6107a56004803603602081101561078f57600080fd5b8101908080359060200190929190505050611263565b005b6107af6112d8565b6040518082600560200280838360005b838110156107da5780820151818401526020810190506107bf565b5050505090500191505060405180910390f35b60606040518060400160405280600f81526020017f54656c6c6f722054726962757465730000000000000000000000000000000000815250905090565b60008073__v2TellorTransfer______________________63052bc340909185856040518463ffffffff1660e01b8152600401808481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060206040518083038186803b1580156108b957600080fd5b505af41580156108cd573d6000803e3d6000fd5b505050506040513d60208110156108e357600080fd5b8101908080519060200190929190505050905092915050565b600073__v2TellorStake_________________________6318b2a9eb90916040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561094e57600080fd5b505af4158015610962573d6000803e3d6000fd5b50505050565b60008073__v2TellorTransfer______________________63dc84521e90918686866040518563ffffffff1660e01b8152600401808581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200194505050505060206040518083038186803b158015610a2b57600080fd5b505af4158015610a3f573d6000803e3d6000fd5b505050506040513d6020811015610a5557600080fd5b810190808051906020019092919050505090509392505050565b600073__v2TellorDispute_______________________636677dc919091836040518363ffffffff1660e01b8152600401808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060006040518083038186803b158015610af557600080fd5b505af4158015610b09573d6000803e3d6000fd5b5050505050565b600073__v2TellorStake_________________________6313a08a9e90916040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b158015610b6257600080fd5b505af4158015610b76573d6000803e3d6000fd5b50505050565b60006012905090565b6000610b8f611757565b600080610b9c60006112ef565b935093509350935090919293565b600073__v2TellorLibrary_______________________63fce5b00f9091868686866040518663ffffffff1660e01b8152600401808681526020018060200184600560200280828437600081840152601f19601f82011690508083019250505083600560200280828437600081840152601f19601f8201169050808301925050508281038252868682818152602001925080828437600081840152601f19601f820116905080830192505050965050505050505060006040518083038186803b158015610c7657600080fd5b505af4158015610c8a573d6000803e3d6000fd5b5050505050505050565b600073__v2TellorDispute_______________________6333f7b1009091836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b158015610cee57600080fd5b505af4158015610d02573d6000803e3d6000fd5b5050505050565b600073__v2TellorLibrary_______________________6376bb5a7990916040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b158015610d5b57600080fd5b505af4158015610d6f573d6000803e3d6000fd5b50505050565b600073__v2TellorLibrary_______________________63c108687c9091868686866040518663ffffffff1660e01b815260040180868152602001806020018481526020018381526020018281038252868682818152602001925080828437600081840152601f19601f820116905080830192505050965050505050505060006040518083038186803b158015610e0b57600080fd5b505af4158015610e1f573d6000803e3d6000fd5b5050505050505050565b600073__v2TellorLibrary_______________________633c5ff1a89091836040518363ffffffff1660e01b8152600401808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060006040518083038186803b158015610eaf57600080fd5b505af4158015610ec3573d6000803e3d6000fd5b5050505050565b600073__v2TellorLibrary_______________________63b2611916909184846040518463ffffffff1660e01b815260040180848152602001838152602001828152602001935050505060006040518083038186803b158015610f2c57600080fd5b505af4158015610f40573d6000803e3d6000fd5b505050505050565b600073__v2TellorDispute_______________________63ac30382c90918585856040518563ffffffff1660e01b81526004018085815260200184815260200183815260200182815260200194505050505060006040518083038186803b158015610fb257600080fd5b505af4158015610fc6573d6000803e3d6000fd5b50505050505050565b60606040518060400160405280600381526020017f5452420000000000000000000000000000000000000000000000000000000000815250905090565b600073__v2TellorDispute_______________________630fbe95da9091836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561106657600080fd5b505af415801561107a573d6000803e3d6000fd5b5050505050565b611089611757565b611091611757565b61109b60006113f0565b915091509091565b60008073__v2TellorTransfer______________________6317fe051c909185856040518463ffffffff1660e01b8152600401808481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060206040518083038186803b15801561113257600080fd5b505af4158015611146573d6000803e3d6000fd5b505050506040513d602081101561115c57600080fd5b8101908080519060200190929190505050905092915050565b600073__v2TellorStake_________________________6337a31e8490916040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b1580156111c757600080fd5b505af41580156111db573d6000803e3d6000fd5b50505050565b600073__v2TellorDispute_______________________63514fe794909184846040518463ffffffff1660e01b81526004018084815260200183815260200182151515158152602001935050505060006040518083038186803b15801561124757600080fd5b505af415801561125b573d6000803e3d6000fd5b505050505050565b600073__v2TellorDispute_______________________630d0459899091836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b1580156112bd57600080fd5b505af41580156112d1573d6000803e3d6000fd5b5050505050565b6112e0611757565b6112ea60006114b1565b905090565b60006112f9611757565b60008060008090505b60058110156113445785603501816005811061131a57fe5b600202016000015484826005811061132e57fe5b6020020181815250508080600101915050611302565b5084600001548386604001600060405180807f646966666963756c747900000000000000000000000000000000000000000000815250600a019050604051809103902081526020019081526020016000205487604001600060405180807f63757272656e74546f74616c54697073000000000000000000000000000000008152506010019050604051809103902081526020019081526020016000205493509350935093509193509193565b6113f8611757565b611400611757565b611409836114b1565b915060008090505b60058110156114ab5783604801600084836005811061142c57fe5b60200201518152602001908152602001600020600401600060405180807f746f74616c5469700000000000000000000000000000000000000000000000008152506008019050604051809103902081526020019081526020016000205482826005811061149557fe5b6020020181815250508080600101915050611411565b50915091565b6114b9611757565b6114c1611757565b6114c9611757565b61151184600101603380602002604051908101604052809291908260338015611507576020028201915b8154815260200190600101908083116114f3575b50505050506115c5565b809250819350505060008090505b60058110156115bd57600083826005811061153657fe5b6020020151146115805784604301600083836005811061155257fe5b602002015181526020019081526020016000205484826005811061157257fe5b6020020181815250506115b0565b84603501816004036005811061159257fe5b60020201600001548482600581106115a657fe5b6020020181815250505b808060010191505061151f565b505050919050565b6115cd611757565b6115d5611757565b6000836001603381106115e457fe5b60200201519050600080905060008090505b600581101561167d5785600182016033811061160e57fe5b602002015185826005811061161f57fe5b6020020181815250506001810184826005811061163857fe5b6020020181815250508285826005811061164e57fe5b602002015110156116705784816005811061166557fe5b602002015192508091505b80806001019150506115f6565b506000600690505b603381101561174f578286826033811061169b57fe5b60200201511115611742578581603381106116b257fe5b60200201518583600581106116c357fe5b602002018181525050808483600581106116d957fe5b6020020181815250508581603381106116ee57fe5b6020020151925060008090505b6005811015611740578386826005811061171157fe5b602002015110156117335785816005811061172857fe5b602002015193508092505b80806001019150506116fb565b505b8080600101915050611685565b505050915091565b6040518060a0016040528060059060208202803883398082019150509050509056fea265627a7a72315820b201acabfd0030ab8aa898aceaa710912a6fb39fded650651b091769a409b94264736f6c63430005100032"

// DeployV2Tellor deploys a new Ethereum contract, binding an instance of V2Tellor to it.
func DeployV2Tellor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *V2Tellor, error) {
	parsed, err := abi.JSON(strings.NewReader(V2TellorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(V2TellorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &V2Tellor{V2TellorCaller: V2TellorCaller{contract: contract}, V2TellorTransactor: V2TellorTransactor{contract: contract}, V2TellorFilterer: V2TellorFilterer{contract: contract}}, nil
}

// V2Tellor is an auto generated Go binding around an Ethereum contract.
type V2Tellor struct {
	V2TellorCaller     // Read-only binding to the contract
	V2TellorTransactor // Write-only binding to the contract
	V2TellorFilterer   // Log filterer for contract events
}

// V2TellorCaller is an auto generated read-only Go binding around an Ethereum contract.
type V2TellorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2TellorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V2TellorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2TellorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V2TellorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2TellorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V2TellorSession struct {
	Contract     *V2Tellor         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V2TellorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V2TellorCallerSession struct {
	Contract *V2TellorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// V2TellorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V2TellorTransactorSession struct {
	Contract     *V2TellorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// V2TellorRaw is an auto generated low-level Go binding around an Ethereum contract.
type V2TellorRaw struct {
	Contract *V2Tellor // Generic contract binding to access the raw methods on
}

// V2TellorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V2TellorCallerRaw struct {
	Contract *V2TellorCaller // Generic read-only contract binding to access the raw methods on
}

// V2TellorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V2TellorTransactorRaw struct {
	Contract *V2TellorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV2Tellor creates a new instance of V2Tellor, bound to a specific deployed contract.
func NewV2Tellor(address common.Address, backend bind.ContractBackend) (*V2Tellor, error) {
	contract, err := bindV2Tellor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V2Tellor{V2TellorCaller: V2TellorCaller{contract: contract}, V2TellorTransactor: V2TellorTransactor{contract: contract}, V2TellorFilterer: V2TellorFilterer{contract: contract}}, nil
}

// NewV2TellorCaller creates a new read-only instance of V2Tellor, bound to a specific deployed contract.
func NewV2TellorCaller(address common.Address, caller bind.ContractCaller) (*V2TellorCaller, error) {
	contract, err := bindV2Tellor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V2TellorCaller{contract: contract}, nil
}

// NewV2TellorTransactor creates a new write-only instance of V2Tellor, bound to a specific deployed contract.
func NewV2TellorTransactor(address common.Address, transactor bind.ContractTransactor) (*V2TellorTransactor, error) {
	contract, err := bindV2Tellor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V2TellorTransactor{contract: contract}, nil
}

// NewV2TellorFilterer creates a new log filterer instance of V2Tellor, bound to a specific deployed contract.
func NewV2TellorFilterer(address common.Address, filterer bind.ContractFilterer) (*V2TellorFilterer, error) {
	contract, err := bindV2Tellor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V2TellorFilterer{contract: contract}, nil
}

// bindV2Tellor binds a generic wrapper to an already deployed contract.
func bindV2Tellor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(V2TellorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V2Tellor *V2TellorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V2Tellor.Contract.V2TellorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V2Tellor *V2TellorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2Tellor.Contract.V2TellorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V2Tellor *V2TellorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V2Tellor.Contract.V2TellorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V2Tellor *V2TellorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V2Tellor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V2Tellor *V2TellorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2Tellor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V2Tellor *V2TellorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V2Tellor.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_V2Tellor *V2TellorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _V2Tellor.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_V2Tellor *V2TellorSession) Decimals() (uint8, error) {
	return _V2Tellor.Contract.Decimals(&_V2Tellor.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_V2Tellor *V2TellorCallerSession) Decimals() (uint8, error) {
	return _V2Tellor.Contract.Decimals(&_V2Tellor.CallOpts)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficutly, uint256 _tip)
func (_V2Tellor *V2TellorCaller) GetNewCurrentVariables(opts *bind.CallOpts) (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	var out []interface{}
	err := _V2Tellor.contract.Call(opts, &out, "getNewCurrentVariables")

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
func (_V2Tellor *V2TellorSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	return _V2Tellor.Contract.GetNewCurrentVariables(&_V2Tellor.CallOpts)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficutly, uint256 _tip)
func (_V2Tellor *V2TellorCallerSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	return _V2Tellor.Contract.GetNewCurrentVariables(&_V2Tellor.CallOpts)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_V2Tellor *V2TellorCaller) GetNewVariablesOnDeck(opts *bind.CallOpts) (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	var out []interface{}
	err := _V2Tellor.contract.Call(opts, &out, "getNewVariablesOnDeck")

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
func (_V2Tellor *V2TellorSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _V2Tellor.Contract.GetNewVariablesOnDeck(&_V2Tellor.CallOpts)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_V2Tellor *V2TellorCallerSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _V2Tellor.Contract.GetNewVariablesOnDeck(&_V2Tellor.CallOpts)
}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_V2Tellor *V2TellorCaller) GetTopRequestIDs(opts *bind.CallOpts) ([5]*big.Int, error) {
	var out []interface{}
	err := _V2Tellor.contract.Call(opts, &out, "getTopRequestIDs")

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_V2Tellor *V2TellorSession) GetTopRequestIDs() ([5]*big.Int, error) {
	return _V2Tellor.Contract.GetTopRequestIDs(&_V2Tellor.CallOpts)
}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_V2Tellor *V2TellorCallerSession) GetTopRequestIDs() ([5]*big.Int, error) {
	return _V2Tellor.Contract.GetTopRequestIDs(&_V2Tellor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_V2Tellor *V2TellorCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _V2Tellor.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_V2Tellor *V2TellorSession) Name() (string, error) {
	return _V2Tellor.Contract.Name(&_V2Tellor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_V2Tellor *V2TellorCallerSession) Name() (string, error) {
	return _V2Tellor.Contract.Name(&_V2Tellor.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_V2Tellor *V2TellorCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _V2Tellor.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_V2Tellor *V2TellorSession) Symbol() (string, error) {
	return _V2Tellor.Contract.Symbol(&_V2Tellor.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_V2Tellor *V2TellorCallerSession) Symbol() (string, error) {
	return _V2Tellor.Contract.Symbol(&_V2Tellor.CallOpts)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_V2Tellor *V2TellorTransactor) AddTip(opts *bind.TransactOpts, _requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _V2Tellor.contract.Transact(opts, "addTip", _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_V2Tellor *V2TellorSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _V2Tellor.Contract.AddTip(&_V2Tellor.TransactOpts, _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_V2Tellor *V2TellorTransactorSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _V2Tellor.Contract.AddTip(&_V2Tellor.TransactOpts, _requestId, _tip)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_V2Tellor *V2TellorTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V2Tellor.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_V2Tellor *V2TellorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V2Tellor.Contract.Approve(&_V2Tellor.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_V2Tellor *V2TellorTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V2Tellor.Contract.Approve(&_V2Tellor.TransactOpts, _spender, _amount)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_V2Tellor *V2TellorTransactor) BeginDispute(opts *bind.TransactOpts, _requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _V2Tellor.contract.Transact(opts, "beginDispute", _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_V2Tellor *V2TellorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _V2Tellor.Contract.BeginDispute(&_V2Tellor.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_V2Tellor *V2TellorTransactorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _V2Tellor.Contract.BeginDispute(&_V2Tellor.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_V2Tellor *V2TellorTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2Tellor.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_V2Tellor *V2TellorSession) ClaimOwnership() (*types.Transaction, error) {
	return _V2Tellor.Contract.ClaimOwnership(&_V2Tellor.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_V2Tellor *V2TellorTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _V2Tellor.Contract.ClaimOwnership(&_V2Tellor.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_V2Tellor *V2TellorTransactor) DepositStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2Tellor.contract.Transact(opts, "depositStake")
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_V2Tellor *V2TellorSession) DepositStake() (*types.Transaction, error) {
	return _V2Tellor.Contract.DepositStake(&_V2Tellor.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_V2Tellor *V2TellorTransactorSession) DepositStake() (*types.Transaction, error) {
	return _V2Tellor.Contract.DepositStake(&_V2Tellor.TransactOpts)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_V2Tellor *V2TellorTransactor) ProposeFork(opts *bind.TransactOpts, _propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _V2Tellor.contract.Transact(opts, "proposeFork", _propNewTellorAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_V2Tellor *V2TellorSession) ProposeFork(_propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _V2Tellor.Contract.ProposeFork(&_V2Tellor.TransactOpts, _propNewTellorAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_V2Tellor *V2TellorTransactorSession) ProposeFork(_propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _V2Tellor.Contract.ProposeFork(&_V2Tellor.TransactOpts, _propNewTellorAddress)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _pendingOwner) returns()
func (_V2Tellor *V2TellorTransactor) ProposeOwnership(opts *bind.TransactOpts, _pendingOwner common.Address) (*types.Transaction, error) {
	return _V2Tellor.contract.Transact(opts, "proposeOwnership", _pendingOwner)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _pendingOwner) returns()
func (_V2Tellor *V2TellorSession) ProposeOwnership(_pendingOwner common.Address) (*types.Transaction, error) {
	return _V2Tellor.Contract.ProposeOwnership(&_V2Tellor.TransactOpts, _pendingOwner)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _pendingOwner) returns()
func (_V2Tellor *V2TellorTransactorSession) ProposeOwnership(_pendingOwner common.Address) (*types.Transaction, error) {
	return _V2Tellor.Contract.ProposeOwnership(&_V2Tellor.TransactOpts, _pendingOwner)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_V2Tellor *V2TellorTransactor) RequestStakingWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2Tellor.contract.Transact(opts, "requestStakingWithdraw")
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_V2Tellor *V2TellorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _V2Tellor.Contract.RequestStakingWithdraw(&_V2Tellor.TransactOpts)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_V2Tellor *V2TellorTransactorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _V2Tellor.Contract.RequestStakingWithdraw(&_V2Tellor.TransactOpts)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_V2Tellor *V2TellorTransactor) SubmitMiningSolution(opts *bind.TransactOpts, _nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _V2Tellor.contract.Transact(opts, "submitMiningSolution", _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_V2Tellor *V2TellorSession) SubmitMiningSolution(_nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _V2Tellor.Contract.SubmitMiningSolution(&_V2Tellor.TransactOpts, _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_V2Tellor *V2TellorTransactorSession) SubmitMiningSolution(_nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _V2Tellor.Contract.SubmitMiningSolution(&_V2Tellor.TransactOpts, _nonce, _requestId, _value)
}

// SubmitMiningSolution0 is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_V2Tellor *V2TellorTransactor) SubmitMiningSolution0(opts *bind.TransactOpts, _nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _V2Tellor.contract.Transact(opts, "submitMiningSolution0", _nonce, _requestId, _value)
}

// SubmitMiningSolution0 is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_V2Tellor *V2TellorSession) SubmitMiningSolution0(_nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _V2Tellor.Contract.SubmitMiningSolution0(&_V2Tellor.TransactOpts, _nonce, _requestId, _value)
}

// SubmitMiningSolution0 is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_V2Tellor *V2TellorTransactorSession) SubmitMiningSolution0(_nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _V2Tellor.Contract.SubmitMiningSolution0(&_V2Tellor.TransactOpts, _nonce, _requestId, _value)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_V2Tellor *V2TellorTransactor) TallyVotes(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _V2Tellor.contract.Transact(opts, "tallyVotes", _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_V2Tellor *V2TellorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _V2Tellor.Contract.TallyVotes(&_V2Tellor.TransactOpts, _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_V2Tellor *V2TellorTransactorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _V2Tellor.Contract.TallyVotes(&_V2Tellor.TransactOpts, _disputeId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_V2Tellor *V2TellorTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V2Tellor.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_V2Tellor *V2TellorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V2Tellor.Contract.Transfer(&_V2Tellor.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_V2Tellor *V2TellorTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V2Tellor.Contract.Transfer(&_V2Tellor.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_V2Tellor *V2TellorTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V2Tellor.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_V2Tellor *V2TellorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V2Tellor.Contract.TransferFrom(&_V2Tellor.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_V2Tellor *V2TellorTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V2Tellor.Contract.TransferFrom(&_V2Tellor.TransactOpts, _from, _to, _amount)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_V2Tellor *V2TellorTransactor) UnlockDisputeFee(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _V2Tellor.contract.Transact(opts, "unlockDisputeFee", _disputeId)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_V2Tellor *V2TellorSession) UnlockDisputeFee(_disputeId *big.Int) (*types.Transaction, error) {
	return _V2Tellor.Contract.UnlockDisputeFee(&_V2Tellor.TransactOpts, _disputeId)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_V2Tellor *V2TellorTransactorSession) UnlockDisputeFee(_disputeId *big.Int) (*types.Transaction, error) {
	return _V2Tellor.Contract.UnlockDisputeFee(&_V2Tellor.TransactOpts, _disputeId)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_V2Tellor *V2TellorTransactor) UpdateTellor(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _V2Tellor.contract.Transact(opts, "updateTellor", _disputeId)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_V2Tellor *V2TellorSession) UpdateTellor(_disputeId *big.Int) (*types.Transaction, error) {
	return _V2Tellor.Contract.UpdateTellor(&_V2Tellor.TransactOpts, _disputeId)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_V2Tellor *V2TellorTransactorSession) UpdateTellor(_disputeId *big.Int) (*types.Transaction, error) {
	return _V2Tellor.Contract.UpdateTellor(&_V2Tellor.TransactOpts, _disputeId)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_V2Tellor *V2TellorTransactor) Vote(opts *bind.TransactOpts, _disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _V2Tellor.contract.Transact(opts, "vote", _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_V2Tellor *V2TellorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _V2Tellor.Contract.Vote(&_V2Tellor.TransactOpts, _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_V2Tellor *V2TellorTransactorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _V2Tellor.Contract.Vote(&_V2Tellor.TransactOpts, _disputeId, _supportsDispute)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_V2Tellor *V2TellorTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2Tellor.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_V2Tellor *V2TellorSession) WithdrawStake() (*types.Transaction, error) {
	return _V2Tellor.Contract.WithdrawStake(&_V2Tellor.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_V2Tellor *V2TellorTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _V2Tellor.Contract.WithdrawStake(&_V2Tellor.TransactOpts)
}
