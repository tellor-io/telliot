// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package v25Tellor

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

// V25TellorABI is the input ABI used to generate the binding from.
const V25TellorABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minerIndex\",\"type\":\"uint256\"}],\"name\":\"beginDispute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_supportsDispute\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"tallyVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_propNewTellorAddress\",\"type\":\"address\"}],\"name\":\"proposeFork\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"}],\"name\":\"submitMiningSolution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_pendingOwner\",\"type\":\"address\"}],\"name\":\"proposeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"requestStakingWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNewCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256\",\"name\":\"_difficutly\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTopRequestIDs\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNewVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"idsOnDeck\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"tipsOnDeck\",\"type\":\"uint256[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"updateTellor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"unlockDisputeFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// V25TellorBin is the compiled bytecode used for deploying new contracts.
var V25TellorBin = "0x608060405234801561001057600080fd5b50611662806100206000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c8063710bf322116100b85780639a7077ab1161007c5780639a7077ab146105c2578063a9059cbb14610637578063bed9d8611461069d578063c9d27afe146106a7578063f458ab98146106e1578063fe1cd15d1461070f57610142565b8063710bf32214610453578063752d49a1146104975780638581af19146104cf57806395d89b41146105115780639a01ca131461059457610142565b806328449c3a1161010a57806328449c3a14610304578063313ce5671461030e5780634049f198146103325780634350283e1461038d5780634d318b0e1461041b5780634e71e0c81461044957610142565b806306fdde0314610147578063095ea7b3146101ca5780630d2d76a21461023057806323b872dd1461023a57806326b7d9f6146102c0575b600080fd5b61014f610755565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561018f578082015181840152602081019050610174565b50505050905090810190601f1680156101bc5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610216600480360360408110156101e057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610792565b604051808215151515815260200191505060405180910390f35b610238610864565b005b6102a66004803603606081101561025057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506108d0565b604051808215151515815260200191505060405180910390f35b610302600480360360208110156102d657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109d7565b005b61030c610a78565b005b610316610ae4565b604051808260ff1660ff16815260200191505060405180910390f35b61033a610aed565b6040518085815260200184600560200280838360005b8381101561036b578082015181840152602081019050610350565b5050505090500183815260200182815260200194505050505060405180910390f35b61041960048036036101608110156103a457600080fd5b81019080803590602001906401000000008111156103c157600080fd5b8201836020820111156103d357600080fd5b803590602001918460018302840111640100000000831117156103f557600080fd5b90919293919293908060a0019091929192908060a001909192919290505050610b12565b005b6104476004803603602081101561043157600080fd5b8101908080359060200190929190505050610bfc565b005b610451610c71565b005b6104956004803603602081101561046957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610cdd565b005b6104cd600480360360408110156104ad57600080fd5b810190808035906020019092919080359060200190929190505050610d7e565b005b61050f600480360360608110156104e557600080fd5b81019080803590602001909291908035906020019092919080359060200190929190505050610dfc565b005b610519610e83565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561055957808201518184015260208101905061053e565b50505050905090810190601f1680156105865780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6105c0600480360360208110156105aa57600080fd5b8101908080359060200190929190505050610ec0565b005b6105ca610f35565b6040518083600560200280838360005b838110156105f55780820151818401526020810190506105da565b5050505090500182600560200280838360005b83811015610623578082015181840152602081019050610608565b505050509050019250505060405180910390f35b6106836004803603604081101561064d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610f57565b604051808215151515815260200191505060405180910390f35b6106a5611029565b005b6106df600480360360408110156106bd57600080fd5b8101908080359060200190929190803515159060200190929190505050611095565b005b61070d600480360360208110156106f757600080fd5b8101908080359060200190929190505050611117565b005b61071761118c565b6040518082600560200280838360005b83811015610742578082015181840152602081019050610727565b5050505090500191505060405180910390f35b60606040518060400160405280601281526020017f76323554656c6c6f722054726962757465730000000000000000000000000000815250905090565b60008073__v25TellorTransfer_____________________63b5e46a62909185856040518463ffffffff1660e01b8152600401808481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060206040518083038186803b15801561082157600080fd5b505af4158015610835573d6000803e3d6000fd5b505050506040513d602081101561084b57600080fd5b8101908080519060200190929190505050905092915050565b600073__v25TellorStake________________________63625641eb90916040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b1580156108b657600080fd5b505af41580156108ca573d6000803e3d6000fd5b50505050565b60008073__v25TellorTransfer_____________________6352fec45d90918686866040518563ffffffff1660e01b8152600401808581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200194505050505060206040518083038186803b15801561099357600080fd5b505af41580156109a7573d6000803e3d6000fd5b505050506040513d60208110156109bd57600080fd5b810190808051906020019092919050505090509392505050565b600073__v25TellorDispute______________________633e2d55249091836040518363ffffffff1660e01b8152600401808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060006040518083038186803b158015610a5d57600080fd5b505af4158015610a71573d6000803e3d6000fd5b5050505050565b600073__v25TellorStake________________________63443ef57d90916040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b158015610aca57600080fd5b505af4158015610ade573d6000803e3d6000fd5b50505050565b60006012905090565b6000610af761160b565b600080610b0460006111a3565b935093509350935090919293565b600073__v25TellorLibrary______________________6336c0d8589091868686866040518663ffffffff1660e01b8152600401808681526020018060200184600560200280828437600081840152601f19601f82011690508083019250505083600560200280828437600081840152601f19601f8201169050808301925050508281038252868682818152602001925080828437600081840152601f19601f820116905080830192505050965050505050505060006040518083038186803b158015610bde57600080fd5b505af4158015610bf2573d6000803e3d6000fd5b5050505050505050565b600073__v25TellorDispute______________________634018145f9091836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b158015610c5657600080fd5b505af4158015610c6a573d6000803e3d6000fd5b5050505050565b600073__v25TellorLibrary______________________631d49d12d90916040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b158015610cc357600080fd5b505af4158015610cd7573d6000803e3d6000fd5b50505050565b600073__v25TellorLibrary______________________633916bf889091836040518363ffffffff1660e01b8152600401808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060006040518083038186803b158015610d6357600080fd5b505af4158015610d77573d6000803e3d6000fd5b5050505050565b600073__v25TellorLibrary______________________633718708d909184846040518463ffffffff1660e01b815260040180848152602001838152602001828152602001935050505060006040518083038186803b158015610de057600080fd5b505af4158015610df4573d6000803e3d6000fd5b505050505050565b600073__v25TellorDispute______________________632cf174c290918585856040518563ffffffff1660e01b81526004018085815260200184815260200183815260200182815260200194505050505060006040518083038186803b158015610e6657600080fd5b505af4158015610e7a573d6000803e3d6000fd5b50505050505050565b60606040518060400160405280600381526020017f5452420000000000000000000000000000000000000000000000000000000000815250905090565b600073__v25TellorDispute______________________63948cfa9c9091836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b158015610f1a57600080fd5b505af4158015610f2e573d6000803e3d6000fd5b5050505050565b610f3d61160b565b610f4561160b565b610f4f60006112a4565b915091509091565b60008073__v25TellorTransfer_____________________638fbe41c1909185856040518463ffffffff1660e01b8152600401808481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060206040518083038186803b158015610fe657600080fd5b505af4158015610ffa573d6000803e3d6000fd5b505050506040513d602081101561101057600080fd5b8101908080519060200190929190505050905092915050565b600073__v25TellorStake________________________63f88c2d2190916040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561107b57600080fd5b505af415801561108f573d6000803e3d6000fd5b50505050565b600073__v25TellorDispute______________________63db2c7356909184846040518463ffffffff1660e01b81526004018084815260200183815260200182151515158152602001935050505060006040518083038186803b1580156110fb57600080fd5b505af415801561110f573d6000803e3d6000fd5b505050505050565b600073__v25TellorDispute______________________635115ed9d9091836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561117157600080fd5b505af4158015611185573d6000803e3d6000fd5b5050505050565b61119461160b565b61119e6000611365565b905090565b60006111ad61160b565b60008060008090505b60058110156111f8578560350181600581106111ce57fe5b60020201600001548482600581106111e257fe5b60200201818152505080806001019150506111b6565b5084600001548386604001600060405180807f646966666963756c747900000000000000000000000000000000000000000000815250600a019050604051809103902081526020019081526020016000205487604001600060405180807f63757272656e74546f74616c54697073000000000000000000000000000000008152506010019050604051809103902081526020019081526020016000205493509350935093509193509193565b6112ac61160b565b6112b461160b565b6112bd83611365565b915060008090505b600581101561135f578360480160008483600581106112e057fe5b60200201518152602001908152602001600020600401600060405180807f746f74616c5469700000000000000000000000000000000000000000000000008152506008019050604051809103902081526020019081526020016000205482826005811061134957fe5b60200201818152505080806001019150506112c5565b50915091565b61136d61160b565b61137561160b565b61137d61160b565b6113c5846001016033806020026040519081016040528092919082603380156113bb576020028201915b8154815260200190600101908083116113a7575b5050505050611479565b809250819350505060008090505b60058110156114715760008382600581106113ea57fe5b6020020151146114345784604301600083836005811061140657fe5b602002015181526020019081526020016000205484826005811061142657fe5b602002018181525050611464565b84603501816004036005811061144657fe5b600202016000015484826005811061145a57fe5b6020020181815250505b80806001019150506113d3565b505050919050565b61148161160b565b61148961160b565b60008360016033811061149857fe5b60200201519050600080905060008090505b6005811015611531578560018201603381106114c257fe5b60200201518582600581106114d357fe5b602002018181525050600181018482600581106114ec57fe5b6020020181815250508285826005811061150257fe5b602002015110156115245784816005811061151957fe5b602002015192508091505b80806001019150506114aa565b506000600690505b6033811015611603578286826033811061154f57fe5b602002015111156115f65785816033811061156657fe5b602002015185836005811061157757fe5b6020020181815250508084836005811061158d57fe5b6020020181815250508581603381106115a257fe5b6020020151925060008090505b60058110156115f457838682600581106115c557fe5b602002015110156115e7578581600581106115dc57fe5b602002015193508092505b80806001019150506115af565b505b8080600101915050611539565b505050915091565b6040518060a0016040528060059060208202803883398082019150509050509056fea265627a7a723158201b0b271476e0609288b31f18569edf9fbe447b085f36e2243e8bd586ccbecece64736f6c63430005100032"

// DeployV25Tellor deploys a new Ethereum contract, binding an instance of V25Tellor to it.
func DeployV25Tellor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *V25Tellor, error) {
	parsed, err := abi.JSON(strings.NewReader(V25TellorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(V25TellorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &V25Tellor{V25TellorCaller: V25TellorCaller{contract: contract}, V25TellorTransactor: V25TellorTransactor{contract: contract}, V25TellorFilterer: V25TellorFilterer{contract: contract}}, nil
}

// V25Tellor is an auto generated Go binding around an Ethereum contract.
type V25Tellor struct {
	V25TellorCaller     // Read-only binding to the contract
	V25TellorTransactor // Write-only binding to the contract
	V25TellorFilterer   // Log filterer for contract events
}

// V25TellorCaller is an auto generated read-only Go binding around an Ethereum contract.
type V25TellorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V25TellorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V25TellorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V25TellorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V25TellorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V25TellorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V25TellorSession struct {
	Contract     *V25Tellor        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V25TellorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V25TellorCallerSession struct {
	Contract *V25TellorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// V25TellorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V25TellorTransactorSession struct {
	Contract     *V25TellorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// V25TellorRaw is an auto generated low-level Go binding around an Ethereum contract.
type V25TellorRaw struct {
	Contract *V25Tellor // Generic contract binding to access the raw methods on
}

// V25TellorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V25TellorCallerRaw struct {
	Contract *V25TellorCaller // Generic read-only contract binding to access the raw methods on
}

// V25TellorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V25TellorTransactorRaw struct {
	Contract *V25TellorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV25Tellor creates a new instance of V25Tellor, bound to a specific deployed contract.
func NewV25Tellor(address common.Address, backend bind.ContractBackend) (*V25Tellor, error) {
	contract, err := bindV25Tellor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V25Tellor{V25TellorCaller: V25TellorCaller{contract: contract}, V25TellorTransactor: V25TellorTransactor{contract: contract}, V25TellorFilterer: V25TellorFilterer{contract: contract}}, nil
}

// NewV25TellorCaller creates a new read-only instance of V25Tellor, bound to a specific deployed contract.
func NewV25TellorCaller(address common.Address, caller bind.ContractCaller) (*V25TellorCaller, error) {
	contract, err := bindV25Tellor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V25TellorCaller{contract: contract}, nil
}

// NewV25TellorTransactor creates a new write-only instance of V25Tellor, bound to a specific deployed contract.
func NewV25TellorTransactor(address common.Address, transactor bind.ContractTransactor) (*V25TellorTransactor, error) {
	contract, err := bindV25Tellor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V25TellorTransactor{contract: contract}, nil
}

// NewV25TellorFilterer creates a new log filterer instance of V25Tellor, bound to a specific deployed contract.
func NewV25TellorFilterer(address common.Address, filterer bind.ContractFilterer) (*V25TellorFilterer, error) {
	contract, err := bindV25Tellor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V25TellorFilterer{contract: contract}, nil
}

// bindV25Tellor binds a generic wrapper to an already deployed contract.
func bindV25Tellor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(V25TellorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V25Tellor *V25TellorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V25Tellor.Contract.V25TellorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V25Tellor *V25TellorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V25Tellor.Contract.V25TellorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V25Tellor *V25TellorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V25Tellor.Contract.V25TellorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V25Tellor *V25TellorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V25Tellor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V25Tellor *V25TellorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V25Tellor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V25Tellor *V25TellorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V25Tellor.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_V25Tellor *V25TellorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _V25Tellor.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_V25Tellor *V25TellorSession) Decimals() (uint8, error) {
	return _V25Tellor.Contract.Decimals(&_V25Tellor.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_V25Tellor *V25TellorCallerSession) Decimals() (uint8, error) {
	return _V25Tellor.Contract.Decimals(&_V25Tellor.CallOpts)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficutly, uint256 _tip)
func (_V25Tellor *V25TellorCaller) GetNewCurrentVariables(opts *bind.CallOpts) (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	var out []interface{}
	err := _V25Tellor.contract.Call(opts, &out, "getNewCurrentVariables")

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
func (_V25Tellor *V25TellorSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	return _V25Tellor.Contract.GetNewCurrentVariables(&_V25Tellor.CallOpts)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficutly, uint256 _tip)
func (_V25Tellor *V25TellorCallerSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	return _V25Tellor.Contract.GetNewCurrentVariables(&_V25Tellor.CallOpts)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_V25Tellor *V25TellorCaller) GetNewVariablesOnDeck(opts *bind.CallOpts) (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	var out []interface{}
	err := _V25Tellor.contract.Call(opts, &out, "getNewVariablesOnDeck")

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
func (_V25Tellor *V25TellorSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _V25Tellor.Contract.GetNewVariablesOnDeck(&_V25Tellor.CallOpts)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_V25Tellor *V25TellorCallerSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _V25Tellor.Contract.GetNewVariablesOnDeck(&_V25Tellor.CallOpts)
}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_V25Tellor *V25TellorCaller) GetTopRequestIDs(opts *bind.CallOpts) ([5]*big.Int, error) {
	var out []interface{}
	err := _V25Tellor.contract.Call(opts, &out, "getTopRequestIDs")

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_V25Tellor *V25TellorSession) GetTopRequestIDs() ([5]*big.Int, error) {
	return _V25Tellor.Contract.GetTopRequestIDs(&_V25Tellor.CallOpts)
}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_V25Tellor *V25TellorCallerSession) GetTopRequestIDs() ([5]*big.Int, error) {
	return _V25Tellor.Contract.GetTopRequestIDs(&_V25Tellor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_V25Tellor *V25TellorCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _V25Tellor.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_V25Tellor *V25TellorSession) Name() (string, error) {
	return _V25Tellor.Contract.Name(&_V25Tellor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_V25Tellor *V25TellorCallerSession) Name() (string, error) {
	return _V25Tellor.Contract.Name(&_V25Tellor.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_V25Tellor *V25TellorCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _V25Tellor.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_V25Tellor *V25TellorSession) Symbol() (string, error) {
	return _V25Tellor.Contract.Symbol(&_V25Tellor.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_V25Tellor *V25TellorCallerSession) Symbol() (string, error) {
	return _V25Tellor.Contract.Symbol(&_V25Tellor.CallOpts)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_V25Tellor *V25TellorTransactor) AddTip(opts *bind.TransactOpts, _requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _V25Tellor.contract.Transact(opts, "addTip", _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_V25Tellor *V25TellorSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _V25Tellor.Contract.AddTip(&_V25Tellor.TransactOpts, _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_V25Tellor *V25TellorTransactorSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _V25Tellor.Contract.AddTip(&_V25Tellor.TransactOpts, _requestId, _tip)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_V25Tellor *V25TellorTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V25Tellor.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_V25Tellor *V25TellorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V25Tellor.Contract.Approve(&_V25Tellor.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_V25Tellor *V25TellorTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V25Tellor.Contract.Approve(&_V25Tellor.TransactOpts, _spender, _amount)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_V25Tellor *V25TellorTransactor) BeginDispute(opts *bind.TransactOpts, _requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _V25Tellor.contract.Transact(opts, "beginDispute", _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_V25Tellor *V25TellorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _V25Tellor.Contract.BeginDispute(&_V25Tellor.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_V25Tellor *V25TellorTransactorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _V25Tellor.Contract.BeginDispute(&_V25Tellor.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_V25Tellor *V25TellorTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V25Tellor.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_V25Tellor *V25TellorSession) ClaimOwnership() (*types.Transaction, error) {
	return _V25Tellor.Contract.ClaimOwnership(&_V25Tellor.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_V25Tellor *V25TellorTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _V25Tellor.Contract.ClaimOwnership(&_V25Tellor.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_V25Tellor *V25TellorTransactor) DepositStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V25Tellor.contract.Transact(opts, "depositStake")
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_V25Tellor *V25TellorSession) DepositStake() (*types.Transaction, error) {
	return _V25Tellor.Contract.DepositStake(&_V25Tellor.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_V25Tellor *V25TellorTransactorSession) DepositStake() (*types.Transaction, error) {
	return _V25Tellor.Contract.DepositStake(&_V25Tellor.TransactOpts)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_V25Tellor *V25TellorTransactor) ProposeFork(opts *bind.TransactOpts, _propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _V25Tellor.contract.Transact(opts, "proposeFork", _propNewTellorAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_V25Tellor *V25TellorSession) ProposeFork(_propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _V25Tellor.Contract.ProposeFork(&_V25Tellor.TransactOpts, _propNewTellorAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_V25Tellor *V25TellorTransactorSession) ProposeFork(_propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _V25Tellor.Contract.ProposeFork(&_V25Tellor.TransactOpts, _propNewTellorAddress)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _pendingOwner) returns()
func (_V25Tellor *V25TellorTransactor) ProposeOwnership(opts *bind.TransactOpts, _pendingOwner common.Address) (*types.Transaction, error) {
	return _V25Tellor.contract.Transact(opts, "proposeOwnership", _pendingOwner)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _pendingOwner) returns()
func (_V25Tellor *V25TellorSession) ProposeOwnership(_pendingOwner common.Address) (*types.Transaction, error) {
	return _V25Tellor.Contract.ProposeOwnership(&_V25Tellor.TransactOpts, _pendingOwner)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _pendingOwner) returns()
func (_V25Tellor *V25TellorTransactorSession) ProposeOwnership(_pendingOwner common.Address) (*types.Transaction, error) {
	return _V25Tellor.Contract.ProposeOwnership(&_V25Tellor.TransactOpts, _pendingOwner)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_V25Tellor *V25TellorTransactor) RequestStakingWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V25Tellor.contract.Transact(opts, "requestStakingWithdraw")
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_V25Tellor *V25TellorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _V25Tellor.Contract.RequestStakingWithdraw(&_V25Tellor.TransactOpts)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_V25Tellor *V25TellorTransactorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _V25Tellor.Contract.RequestStakingWithdraw(&_V25Tellor.TransactOpts)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_V25Tellor *V25TellorTransactor) SubmitMiningSolution(opts *bind.TransactOpts, _nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _V25Tellor.contract.Transact(opts, "submitMiningSolution", _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_V25Tellor *V25TellorSession) SubmitMiningSolution(_nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _V25Tellor.Contract.SubmitMiningSolution(&_V25Tellor.TransactOpts, _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_V25Tellor *V25TellorTransactorSession) SubmitMiningSolution(_nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _V25Tellor.Contract.SubmitMiningSolution(&_V25Tellor.TransactOpts, _nonce, _requestId, _value)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_V25Tellor *V25TellorTransactor) TallyVotes(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _V25Tellor.contract.Transact(opts, "tallyVotes", _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_V25Tellor *V25TellorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _V25Tellor.Contract.TallyVotes(&_V25Tellor.TransactOpts, _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_V25Tellor *V25TellorTransactorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _V25Tellor.Contract.TallyVotes(&_V25Tellor.TransactOpts, _disputeId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_V25Tellor *V25TellorTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V25Tellor.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_V25Tellor *V25TellorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V25Tellor.Contract.Transfer(&_V25Tellor.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_V25Tellor *V25TellorTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V25Tellor.Contract.Transfer(&_V25Tellor.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_V25Tellor *V25TellorTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V25Tellor.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_V25Tellor *V25TellorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V25Tellor.Contract.TransferFrom(&_V25Tellor.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_V25Tellor *V25TellorTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _V25Tellor.Contract.TransferFrom(&_V25Tellor.TransactOpts, _from, _to, _amount)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_V25Tellor *V25TellorTransactor) UnlockDisputeFee(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _V25Tellor.contract.Transact(opts, "unlockDisputeFee", _disputeId)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_V25Tellor *V25TellorSession) UnlockDisputeFee(_disputeId *big.Int) (*types.Transaction, error) {
	return _V25Tellor.Contract.UnlockDisputeFee(&_V25Tellor.TransactOpts, _disputeId)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_V25Tellor *V25TellorTransactorSession) UnlockDisputeFee(_disputeId *big.Int) (*types.Transaction, error) {
	return _V25Tellor.Contract.UnlockDisputeFee(&_V25Tellor.TransactOpts, _disputeId)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_V25Tellor *V25TellorTransactor) UpdateTellor(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _V25Tellor.contract.Transact(opts, "updateTellor", _disputeId)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_V25Tellor *V25TellorSession) UpdateTellor(_disputeId *big.Int) (*types.Transaction, error) {
	return _V25Tellor.Contract.UpdateTellor(&_V25Tellor.TransactOpts, _disputeId)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_V25Tellor *V25TellorTransactorSession) UpdateTellor(_disputeId *big.Int) (*types.Transaction, error) {
	return _V25Tellor.Contract.UpdateTellor(&_V25Tellor.TransactOpts, _disputeId)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_V25Tellor *V25TellorTransactor) Vote(opts *bind.TransactOpts, _disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _V25Tellor.contract.Transact(opts, "vote", _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_V25Tellor *V25TellorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _V25Tellor.Contract.Vote(&_V25Tellor.TransactOpts, _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_V25Tellor *V25TellorTransactorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _V25Tellor.Contract.Vote(&_V25Tellor.TransactOpts, _disputeId, _supportsDispute)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_V25Tellor *V25TellorTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V25Tellor.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_V25Tellor *V25TellorSession) WithdrawStake() (*types.Transaction, error) {
	return _V25Tellor.Contract.WithdrawStake(&_V25Tellor.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_V25Tellor *V25TellorTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _V25Tellor.Contract.WithdrawStake(&_V25Tellor.TransactOpts)
}
