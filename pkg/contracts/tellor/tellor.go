// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tellor

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

// ExtensionABI is the input ABI used to generate the binding from.
const ExtensionABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"_result\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reportedMiner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_reportingParty\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"DisputeVoteTallied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"NewStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"fromBlock\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bytesVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentMiners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"disputesById\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"tally\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disputeVotePassed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPropFork\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"reportedMiner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reportingParty\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposedForkAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[9]\",\"name\":\"\",\"type\":\"uint256[9]\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastNewValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"data\",\"type\":\"uint256[51]\"}],\"name\":\"getMax5\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"max\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"maxIndex\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"internalType\":\"address[5]\",\"name\":\"\",\"type\":\"address[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256\",\"name\":\"_diff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"idsOnDeck\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"tipsOnDeck\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"\",\"type\":\"uint256[51]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTopRequestIDs\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minersByChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"newValueTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestStakingWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"tallyVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"uints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateMinDisputeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ExtensionFuncSigs maps the 4-byte function signature to its string representation.
var ExtensionFuncSigs = map[string]string{
	"024c2ddd": "_allowances(address,address)",
	"699f200f": "addresses(bytes32)",
	"cbf1304d": "balances(address,uint256)",
	"62dd1d2a": "bytesVars(bytes32)",
	"1fd22364": "currentMiners(uint256)",
	"313ce567": "decimals()",
	"0d2d76a2": "depositStake()",
	"63bb82ad": "didMine(bytes32,address)",
	"a7c438bc": "didVote(uint256,address)",
	"d01f4d9e": "disputeIdByDisputeHash(bytes32)",
	"db085beb": "disputesById(uint256)",
	"133bee5e": "getAddressVars(bytes32)",
	"af0b1327": "getAllDisputeVars(uint256)",
	"da379941": "getDisputeIdByDisputeHash(bytes32)",
	"7f6fd5d9": "getDisputeUintVars(uint256,bytes32)",
	"fc7cf0a0": "getLastNewValue()",
	"3180f8df": "getLastNewValueById(uint256)",
	"99830e32": "getMax5(uint256[51])",
	"c775b542": "getMinedBlockNum(uint256,uint256)",
	"69026d63": "getMinersByRequestIdAndTimestamp(uint256,uint256)",
	"4049f198": "getNewCurrentVariables()",
	"46eee1c4": "getNewValueCountbyRequestId(uint256)",
	"9a7077ab": "getNewVariablesOnDeck()",
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
	"3df0777b": "isInDispute(uint256,uint256)",
	"4ba0a5ee": "migrated(address)",
	"48b18e54": "minersByChallenge(bytes32,address)",
	"06fdde03": "name()",
	"438c0aa3": "newValueTimestamps(uint256)",
	"5700242c": "requestIdByQueryHash(bytes32)",
	"28449c3a": "requestStakingWithdraw()",
	"93fa4915": "retrieveData(uint256,uint256)",
	"95d89b41": "symbol()",
	"4d318b0e": "tallyVotes(uint256)",
	"18160ddd": "totalSupply()",
	"b59e14d4": "uints(bytes32)",
	"90e5b235": "updateMinDisputeFee()",
	"bed9d861": "withdrawStake()",
}

// ExtensionBin is the compiled bytecode used for deploying new contracts.
var ExtensionBin = "0x608060405234801561001057600080fd5b50611f6f806100206000396000f3fe608060405234801561001057600080fd5b50600436106102955760003560e01c806369026d6311610167578063b5413029116100ce578063da37994111610087578063da379941146109fc578063db085beb14610a19578063e0ae93c114610a86578063e1eee6d614610aa9578063fc7cf0a014610ac6578063fe1cd15d14610ace57610295565b8063b54130291461091e578063b59e14d41461093c578063bed9d86114610959578063c775b54214610961578063cbf1304d14610984578063d01f4d9e146109df57610295565b806393fa49151161012057806393fa49151461075657806395d89b411461077957806399830e32146107815780639a7077ab14610839578063a7c438bc14610841578063af0b13271461086d57610295565b806369026d6314610689578063699f200f146106ac578063733bdef0146106c957806377fbb663146107085780637f6fd5d91461072b57806390e5b2351461074e57610295565b80633df0777b1161020b5780634d318b0e116101c45780634d318b0e146105cc5780635700242c146105e9578063612c8f7f146106065780636173c0b81461062357806362dd1d2a1461064057806363bb82ad1461065d57610295565b80633df0777b146104b45780634049f198146104eb578063438c0aa31461054057806346eee1c41461055d57806348b18e541461057a5780634ba0a5ee146105a657610295565b8063133bee5e1161025d578063133bee5e146103d957806318160ddd146104125780631fd223641461041a57806328449c3a14610458578063313ce567146104605780633180f8df1461047e57610295565b8063024c2ddd1461029a57806306fdde03146102da5780630d2d76a2146103575780630f0b424d1461036157806311c985121461037e575b600080fd5b6102c8600480360360408110156102b057600080fd5b506001600160a01b0381358116916020013516610ad6565b60408051918252519081900360200190f35b6102e2610af3565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561031c578181015183820152602001610304565b50505050905090810190601f1680156103495780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61035f610b1c565b005b6102c86004803603602081101561037757600080fd5b5035610b2f565b6103a16004803603604081101561039457600080fd5b5080359060200135610b41565b604051808260a080838360005b838110156103c65781810151838201526020016103ae565b5050505090500191505060405180910390f35b6103f6600480360360208110156103ef57600080fd5b5035610b98565b604080516001600160a01b039092168252519081900360200190f35b6102c8610bb3565b6104376004803603602081101561043057600080fd5b5035610c01565b604080519283526001600160a01b0390911660208301528051918290030190f35b61035f610c2c565b610468610d20565b6040805160ff9092168252519081900360200190f35b61049b6004803603602081101561049457600080fd5b5035610d25565b6040805192835290151560208301528051918290030190f35b6104d7600480360360408110156104ca57600080fd5b5080359060200135610d7f565b604080519115158252519081900360200190f35b6104f3610da3565b604051848152602081018460a080838360005b8381101561051e578181015183820152602001610506565b5050505090500183815260200182815260200194505050505060405180910390f35b6102c86004803603602081101561055657600080fd5b5035610e83565b6102c86004803603602081101561057357600080fd5b5035610ea4565b6104d76004803603604081101561059057600080fd5b50803590602001356001600160a01b0316610eb6565b6104d7600480360360208110156105bc57600080fd5b50356001600160a01b0316610ed6565b61035f600480360360208110156105e257600080fd5b5035610eeb565b6102c8600480360360208110156105ff57600080fd5b503561113c565b6102c86004803603602081101561061c57600080fd5b503561114e565b6102c86004803603602081101561063957600080fd5b5035611160565b6102c86004803603602081101561065657600080fd5b50356111cb565b6104d76004803603604081101561067357600080fd5b50803590602001356001600160a01b03166111dd565b6103a16004803603604081101561069f57600080fd5b5080359060200135611208565b6103f6600480360360208110156106c257600080fd5b503561126b565b6106ef600480360360208110156106df57600080fd5b50356001600160a01b0316611286565b6040805192835260208301919091528051918290030190f35b6102c86004803603604081101561071e57600080fd5b50803590602001356112a9565b6102c86004803603604081101561074157600080fd5b50803590602001356112d6565b61035f6112f7565b6102c86004803603604081101561076c57600080fd5b5080359060200135611411565b6102e2611432565b6107d5600480360361066081101561079857600080fd5b81019080806106600190603380602002604051908101604052809291908260336020028082843760009201919091525091945061144f9350505050565b604051808360a080838360005b838110156107fa5781810151838201526020016107e2565b5050505090500182600560200280838360005b8381101561082557818101518382015260200161080d565b505050509050019250505060405180910390f35b6107d56115a3565b6104d76004803603604081101561085757600080fd5b50803590602001356001600160a01b0316611642565b61088a6004803603602081101561088357600080fd5b5035611671565b604051808a8152602001891515815260200188151581526020018715158152602001866001600160a01b03168152602001856001600160a01b03168152602001846001600160a01b0316815260200183600960200280838360005b838110156108fd5781810151838201526020016108e5565b50505050905001828152602001995050505050505050505060405180910390f35b61092661189c565b60405181518152808261066080838360206103ae565b6102c86004803603602081101561095257600080fd5b50356118d8565b61035f6118ea565b6102c86004803603604081101561097757600080fd5b50803590602001356119c4565b6109b06004803603604081101561099a57600080fd5b506001600160a01b0381351690602001356119e5565b60405180836001600160801b03168152602001826001600160801b031681526020019250505060405180910390f35b6102c8600480360360208110156109f557600080fd5b5035611a28565b6102c860048036036020811015610a1257600080fd5b5035611a3a565b610a3660048036036020811015610a2f57600080fd5b5035611a4c565b60408051988952602089019790975294151587870152921515606087015290151560808601526001600160a01b0390811660a086015290811660c08501521660e083015251908190036101000190f35b6102c860048036036040811015610a9c57600080fd5b5080359060200135611aa8565b6106ef60048036036020811015610abf57600080fd5b5035611ac9565b61049b611b31565b6103a1611ba2565b604a60209081526000928352604080842090915290825290205481565b60408051808201909152600f81526e54656c6c6f7220547269627574657360881b602082015290565b610b2533611c87565b610b2d6112f7565b565b60009081526034602052604090205490565b610b49611e77565b600083815260456020908152604080832085845260060190915290819020815160a08101928390529160059082845b815481526020019060010190808311610b78575050505050905092915050565b6000908152604760205260409020546001600160a01b031690565b7fe6148e7230ca038d456350e69a91b66968b222bfac9ebfbea6ff0a1fb738016060005260466020527ffffeead1ec15181fd57b4590d95e0c076bccb59e311315e8b38f23c710aa7c3e5490565b603a8160058110610c1157600080fd5b6002020180546001909101549091506001600160a01b031682565b3360009081526044602052604090208054600114610c87576040805162461bcd60e51b8152602060048201526013602482015272135a5b995c881a5cc81b9bdd081cdd185ad959606a1b604482015290519081900360640190fd5b60028155620151804206420360018201557f10c168823622203e4057b65015ff4d95b4c650b308918e8c92dc32ab5a0a034b60005260466020527fa5ae3e2b97d73fb849ea855d27f073b72815b38452d976bd57e4a157827dadd38054600019019055610cf26112f7565b60405133907f453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf90600090a250565b601290565b6000818152604560205260408120805482919015610d71578054610d6590859083906000198101908110610d5557fe5b9060005260206000200154611411565b60019250925050610d7a565b60008092509250505b915091565b60009182526045602090815260408084209284526004909201905290205460ff1690565b6000610dad611e77565b60008060005b6005811015610de757603a8160058110610dc957fe5b6002020154848260058110610dda57fe5b6020020152600101610db3565b50507f52cb9007c7c6068f8ef37039d4f232cbf5a28ff8d93a5983c4c0c27cd2f9bc0d5460466020527f5bccd7373734898281f858d7562320d2cdfc0b17bd72f779686937174d150025547f09659d32f99e50ac728058418d38174fe83a137c455ff1847e6fb8e15f78f77a6000527f38b16d06a20ab673b01c748aff938df6a38f81640035f4ce8bd9abb03aae5b7254919450915090919293565b60338181548110610e9357600080fd5b600091825260209091200154905081565b60009081526045602052604090205490565b603960209081526000928352604080842090915290825290205460ff1681565b604b6020526000908152604090205460ff1681565b6000818152603660205260409020600281015460ff1615610f3d5760405162461bcd60e51b8152600401808060200182810382526021815260200180611f196021913960400191505060405180910390fd5b7f46f7d53798d31923f6952572c6a19ad2d1a8238d26649c2f3493a6d69e425d286000908152600582016020526040902054421015610fc3576040805162461bcd60e51b815260206004820152601f60248201527f54696d6520666f7220766f74696e6720686176656e277420656c617073656400604482015290519081900360640190fd5b60038101546001600160a01b0316611022576040805162461bcd60e51b815260206004820152601c60248201527f7265706f7274696e672050617274792069732061646472657373203000000000604482015290519081900360640190fd5b600181015460008113156110425760028201805461ff0019166101001790555b600282015462010000900460ff16611088576002820154630100000090046001600160a01b0316600090815260446020526040902080546003141561108657600481555b505b7ff9e1ae10923bfc79f52e309baf8c7699edb821f91ef5b5bd07be29545917b3a6600090815260058301602090815260409182902042905560028401805460ff191660011790819055600385015483518581526001600160a01b0391821693810193909352610100820460ff1615158385015292516301000000909104929092169185917f21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739919081900360600190a3505050565b60376020526000908152604090205481565b60009081526046602052604090205490565b600060328211156111b8576040805162461bcd60e51b815260206004820152601a60248201527f526571756573745120696e6465782069732061626f7665203530000000000000604482015290519081900360640190fd5b5060009081526035602052604090205490565b60486020526000908152604090205481565b60009182526039602090815260408084206001600160a01b0393909316845291905290205460ff1690565b611210611e77565b6000838152604560209081526040808320858452600590810190925291829020825160a08101938490529290919082845b81546001600160a01b03168152600190910190602001808311611241575050505050905092915050565b6047602052600090815260409020546001600160a01b031681565b6001600160a01b0316600090815260446020526040902080546001909101549091565b60008281526045602052604081208054839081106112c357fe5b9060005260206000200154905092915050565b60009182526036602090815260408084209284526005909201905290205490565b60466020527f167af83a0768d27540775cfef6d996eb63f8a61fcdfb26e654c18fb50960e3be547f2e2f0a18eb55ef91e37921b3810d7feeef7a855ddc7f4f4249ef03d7b887ae31547f10c168823622203e4057b65015ff4d95b4c650b308918e8c92dc32ab5a0a034b6000527fa5ae3e2b97d73fb849ea855d27f073b72815b38452d976bd57e4a157827dadd3546113c29067d02ab486cedc0000906103e89084906113a5908290611e50565b6103e8028602816113b257fe5b04816113ba57fe5b048403611e68565b7f675d2171f68d6f5545d54fb9b1fb61a0e6897e6188ca1cd664e7c9530d91ecfc60005260466020527f3e5522f19747f0f285b96ded572ac4128c3a764aea9f44058dc0afc9dda44986555050565b60009182526045602090815260408084209284526003909201905290205490565b6040805180820190915260038152622a292160e91b602082015290565b611457611e77565b61145f611e77565b60208301516000805b60058110156114e25785816001016033811061148057fe5b602002015185826005811061149157fe5b6020020152600181018482600581106114a657fe5b6020020152828582600581106114b857fe5b602002015110156114da578481600581106114cf57fe5b602002015192508091505b600101611468565b5060065b603381101561159b57828682603381106114fc57fe5b602002015111156115935785816033811061151357fe5b602002015185836005811061152457fe5b60200201528084836005811061153657fe5b602002015285816033811061154757fe5b6020020151925060005b6005811015611591578386826005811061156757fe5b602002015110156115895785816005811061157e57fe5b602002015193508092505b600101611551565b505b6001016114e6565b505050915091565b6115ab611e77565b6115b3611e77565b6115bb611ba2565b915060005b600581101561163d57604560008483600581106115d957fe5b6020020151815260200190815260200160002060010160007f1590276b7f31dd8e2a06f9a92867333eeb3eddbc91e73b9833e3e55d8e34f77d60001b81526020019081526020016000205482826005811061163057fe5b60200201526001016115c0565b509091565b60008281526036602090815260408083206001600160a01b038516845260060190915290205460ff1692915050565b6000806000806000806000611684611e95565b5050506000958652505060366020908152604080862080546002820154600383015460048401548551610120810187527f9f47a2659c3d32b749ae717d975e7962959890862423c4318cf86e4ec220291f8c5260058601808952878d205482527f2f9328a9c75282bec25bb04befad06926366736e0030c985108445fa728335e58d52808952878d2054828a01527f9147231ab14efb72c38117f68521ddef8de64f092c18c69dbfb602ffc4de7f478d52808952878d2054828901527f46f7d53798d31923f6952572c6a19ad2d1a8238d26649c2f3493a6d69e425d288d52808952878d205460608301527f1da378694063870452ce03b189f48e04c1aa026348e74e6c86e10738514ad2c48d52808952878d205460808301527f4b4cefd5ced7569ef0d091282b4bca9c52a034c56471a6061afd1bf307a2de7c8d52808952878d205460a08301527f6de96ee4d33a0617f40a846309c8759048857f51b9d59a12d3c3786d4778883d8d52808952878d205460c08301527f30e85ae205656781c1a951cba9f9f53f884833c049d377a2a7046eb5e6d14b268d52808952878d205460e08301527f1da95f11543c9b03927178e07951795dfc95c7501a9d1cf00e13414ca33bc4098d52909752949099205461010080870191909152600190930154919960ff8083169a948304811699506201000083041697506001600160a01b036301000000909204821696509281169493169291565b6118a4611eb4565b604080516106608101918290529060009060339082845b8154815260200190600101908083116118bb575050505050905090565b60466020526000908152604090205481565b336000908152604460205260409020600181015462093a80906201518042064203031015611954576040805162461bcd60e51b8152602060048201526012602482015271372064617973206469646e2774207061737360701b604482015290519081900360640190fd5b80546002146119945760405162461bcd60e51b8152600401808060200182810382526023815260200180611ed46023913960400191505060405180910390fd5b600080825560405133917f4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec91a250565b60009182526045602090815260408084209284526002909201905290205490565b60496020528160005260406000208181548110611a0157600080fd5b6000918252602090912001546001600160801b038082169350600160801b90910416905082565b60386020526000908152604090205481565b60009081526038602052604090205490565b603660205260009081526040902080546001820154600283015460038401546004909401549293919260ff808316936101008404821693620100008104909216926001600160a01b036301000000909304831692918216911688565b60009182526045602090815260408084209284526001909201905290205490565b60009081526045602090815260408083207ff68d680ab3160f1aa5d9c3a1383c49e3e60bf3c0c031245cbb036f5ce99afaa18452600101909152808220547f1590276b7f31dd8e2a06f9a92867333eeb3eddbc91e73b9833e3e55d8e34f77d83529120549091565b7f231bb0dc207f13dd4e565ebc32496c470e35391bd8d3b6649269ee2328e031185460008181526034602090815260408220547f2c8b528fbaf48aaf13162a5a0519a7ad5a612da8ff8783465c17e076660a59f18352604690915290918291611b9991611411565b92600192509050565b611baa611e77565b611bb2611e77565b611bba611e77565b60408051610660810191829052611bf19160009060339082845b815481526020019060010190808311611bd457505050505061144f565b909250905060005b6005811015611c8157828160058110611c0e57fe5b602002015115611c525760356000838360058110611c2857fe5b6020020151815260200190815260200160002054848260058110611c4857fe5b6020020152611c79565b603a8160040360058110611c6257fe5b6002020154848260058110611c7357fe5b60200201525b600101611bf9565b50505090565b7f167af83a0768d27540775cfef6d996eb63f8a61fcdfb26e654c18fb50960e3be546001600160a01b038216600090815260496020526040902080546000198101908110611cd157fe5b600091825260209091200154600160801b90046001600160801b03161015611d2a5760405162461bcd60e51b8152600401808060200182810382526022815260200180611ef76022913960400191505060405180910390fd5b6001600160a01b0381166000908152604460205260409020541580611d6757506001600160a01b0381166000908152604460205260409020546002145b611db8576040805162461bcd60e51b815260206004820152601b60248201527f4d696e657220697320696e207468652077726f6e672073746174650000000000604482015290519081900360640190fd5b7fa5ae3e2b97d73fb849ea855d27f073b72815b38452d976bd57e4a157827dadd3805460019081019091556040805180820182528281526201518042908106900360208281019182526001600160a01b038616600081815260449092528482209351845591519290940191909155905190917f46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e291a250565b6000818310611e5f5781611e61565b825b9392505050565b6000818311611e5f5781611e61565b6040518060a001604052806005906020820280368337509192915050565b6040518061012001604052806009906020820280368337509192915050565b604051806106600160405280603390602082028036833750919291505056fe4d696e657220776173206e6f74206c6f636b656420666f72207769746864726177616c42616c616e6365206973206c6f776572207468616e207374616b6520616d6f756e744469737075746520686173206265656e20616c7265616479206578656375746564a26469706673582212205a398a3ad7ffb6c6f63cb29939bd95210a7e7b8b50eddb2f294d223f54e32fa364736f6c63430007040033"

// DeployExtension deploys a new Ethereum contract, binding an instance of Extension to it.
func DeployExtension(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Extension, error) {
	parsed, err := abi.JSON(strings.NewReader(ExtensionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExtensionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Extension{ExtensionCaller: ExtensionCaller{contract: contract}, ExtensionTransactor: ExtensionTransactor{contract: contract}, ExtensionFilterer: ExtensionFilterer{contract: contract}}, nil
}

// Extension is an auto generated Go binding around an Ethereum contract.
type Extension struct {
	ExtensionCaller     // Read-only binding to the contract
	ExtensionTransactor // Write-only binding to the contract
	ExtensionFilterer   // Log filterer for contract events
}

// ExtensionCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExtensionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtensionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExtensionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtensionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExtensionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtensionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExtensionSession struct {
	Contract     *Extension        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExtensionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExtensionCallerSession struct {
	Contract *ExtensionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ExtensionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExtensionTransactorSession struct {
	Contract     *ExtensionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ExtensionRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExtensionRaw struct {
	Contract *Extension // Generic contract binding to access the raw methods on
}

// ExtensionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExtensionCallerRaw struct {
	Contract *ExtensionCaller // Generic read-only contract binding to access the raw methods on
}

// ExtensionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExtensionTransactorRaw struct {
	Contract *ExtensionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExtension creates a new instance of Extension, bound to a specific deployed contract.
func NewExtension(address common.Address, backend bind.ContractBackend) (*Extension, error) {
	contract, err := bindExtension(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Extension{ExtensionCaller: ExtensionCaller{contract: contract}, ExtensionTransactor: ExtensionTransactor{contract: contract}, ExtensionFilterer: ExtensionFilterer{contract: contract}}, nil
}

// NewExtensionCaller creates a new read-only instance of Extension, bound to a specific deployed contract.
func NewExtensionCaller(address common.Address, caller bind.ContractCaller) (*ExtensionCaller, error) {
	contract, err := bindExtension(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExtensionCaller{contract: contract}, nil
}

// NewExtensionTransactor creates a new write-only instance of Extension, bound to a specific deployed contract.
func NewExtensionTransactor(address common.Address, transactor bind.ContractTransactor) (*ExtensionTransactor, error) {
	contract, err := bindExtension(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExtensionTransactor{contract: contract}, nil
}

// NewExtensionFilterer creates a new log filterer instance of Extension, bound to a specific deployed contract.
func NewExtensionFilterer(address common.Address, filterer bind.ContractFilterer) (*ExtensionFilterer, error) {
	contract, err := bindExtension(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExtensionFilterer{contract: contract}, nil
}

// bindExtension binds a generic wrapper to an already deployed contract.
func bindExtension(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExtensionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Extension *ExtensionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Extension.Contract.ExtensionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Extension *ExtensionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extension.Contract.ExtensionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Extension *ExtensionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Extension.Contract.ExtensionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Extension *ExtensionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Extension.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Extension *ExtensionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extension.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Extension *ExtensionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Extension.Contract.contract.Transact(opts, method, params...)
}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_Extension *ExtensionCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "_allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_Extension *ExtensionSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Extension.Contract.Allowances(&_Extension.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_Extension *ExtensionCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Extension.Contract.Allowances(&_Extension.CallOpts, arg0, arg1)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_Extension *ExtensionCaller) Addresses(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "addresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_Extension *ExtensionSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _Extension.Contract.Addresses(&_Extension.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_Extension *ExtensionCallerSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _Extension.Contract.Addresses(&_Extension.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_Extension *ExtensionCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "balances", arg0, arg1)

	outstruct := new(struct {
		FromBlock *big.Int
		Value     *big.Int
	})

	outstruct.FromBlock = out[0].(*big.Int)
	outstruct.Value = out[1].(*big.Int)

	return *outstruct, err

}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_Extension *ExtensionSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _Extension.Contract.Balances(&_Extension.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_Extension *ExtensionCallerSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _Extension.Contract.Balances(&_Extension.CallOpts, arg0, arg1)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_Extension *ExtensionCaller) BytesVars(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "bytesVars", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_Extension *ExtensionSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _Extension.Contract.BytesVars(&_Extension.CallOpts, arg0)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_Extension *ExtensionCallerSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _Extension.Contract.BytesVars(&_Extension.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_Extension *ExtensionCaller) CurrentMiners(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "currentMiners", arg0)

	outstruct := new(struct {
		Value *big.Int
		Miner common.Address
	})

	outstruct.Value = out[0].(*big.Int)
	outstruct.Miner = out[1].(common.Address)

	return *outstruct, err

}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_Extension *ExtensionSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _Extension.Contract.CurrentMiners(&_Extension.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_Extension *ExtensionCallerSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _Extension.Contract.CurrentMiners(&_Extension.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Extension *ExtensionCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Extension *ExtensionSession) Decimals() (uint8, error) {
	return _Extension.Contract.Decimals(&_Extension.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Extension *ExtensionCallerSession) Decimals() (uint8, error) {
	return _Extension.Contract.Decimals(&_Extension.CallOpts)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_Extension *ExtensionCaller) DidMine(opts *bind.CallOpts, _challenge [32]byte, _miner common.Address) (bool, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "didMine", _challenge, _miner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_Extension *ExtensionSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _Extension.Contract.DidMine(&_Extension.CallOpts, _challenge, _miner)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_Extension *ExtensionCallerSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _Extension.Contract.DidMine(&_Extension.CallOpts, _challenge, _miner)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_Extension *ExtensionCaller) DidVote(opts *bind.CallOpts, _disputeId *big.Int, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "didVote", _disputeId, _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_Extension *ExtensionSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _Extension.Contract.DidVote(&_Extension.CallOpts, _disputeId, _address)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_Extension *ExtensionCallerSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _Extension.Contract.DidVote(&_Extension.CallOpts, _disputeId, _address)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_Extension *ExtensionCaller) DisputeIdByDisputeHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "disputeIdByDisputeHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_Extension *ExtensionSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _Extension.Contract.DisputeIdByDisputeHash(&_Extension.CallOpts, arg0)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_Extension *ExtensionCallerSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _Extension.Contract.DisputeIdByDisputeHash(&_Extension.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_Extension *ExtensionCaller) DisputesById(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "disputesById", arg0)

	outstruct := new(struct {
		Hash                [32]byte
		Tally               *big.Int
		Executed            bool
		DisputeVotePassed   bool
		IsPropFork          bool
		ReportedMiner       common.Address
		ReportingParty      common.Address
		ProposedForkAddress common.Address
	})

	outstruct.Hash = out[0].([32]byte)
	outstruct.Tally = out[1].(*big.Int)
	outstruct.Executed = out[2].(bool)
	outstruct.DisputeVotePassed = out[3].(bool)
	outstruct.IsPropFork = out[4].(bool)
	outstruct.ReportedMiner = out[5].(common.Address)
	outstruct.ReportingParty = out[6].(common.Address)
	outstruct.ProposedForkAddress = out[7].(common.Address)

	return *outstruct, err

}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_Extension *ExtensionSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _Extension.Contract.DisputesById(&_Extension.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_Extension *ExtensionCallerSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _Extension.Contract.DisputesById(&_Extension.CallOpts, arg0)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_Extension *ExtensionCaller) GetAddressVars(opts *bind.CallOpts, _data [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getAddressVars", _data)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_Extension *ExtensionSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _Extension.Contract.GetAddressVars(&_Extension.CallOpts, _data)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_Extension *ExtensionCallerSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _Extension.Contract.GetAddressVars(&_Extension.CallOpts, _data)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_Extension *ExtensionCaller) GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getAllDisputeVars", _disputeId)

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
func (_Extension *ExtensionSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _Extension.Contract.GetAllDisputeVars(&_Extension.CallOpts, _disputeId)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_Extension *ExtensionCallerSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _Extension.Contract.GetAllDisputeVars(&_Extension.CallOpts, _disputeId)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_Extension *ExtensionCaller) GetDisputeIdByDisputeHash(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getDisputeIdByDisputeHash", _hash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_Extension *ExtensionSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _Extension.Contract.GetDisputeIdByDisputeHash(&_Extension.CallOpts, _hash)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_Extension *ExtensionCallerSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _Extension.Contract.GetDisputeIdByDisputeHash(&_Extension.CallOpts, _hash)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_Extension *ExtensionCaller) GetDisputeUintVars(opts *bind.CallOpts, _disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getDisputeUintVars", _disputeId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_Extension *ExtensionSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _Extension.Contract.GetDisputeUintVars(&_Extension.CallOpts, _disputeId, _data)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_Extension *ExtensionCallerSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _Extension.Contract.GetDisputeUintVars(&_Extension.CallOpts, _disputeId, _data)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_Extension *ExtensionCaller) GetLastNewValue(opts *bind.CallOpts) (*big.Int, bool, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getLastNewValue")

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
func (_Extension *ExtensionSession) GetLastNewValue() (*big.Int, bool, error) {
	return _Extension.Contract.GetLastNewValue(&_Extension.CallOpts)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_Extension *ExtensionCallerSession) GetLastNewValue() (*big.Int, bool, error) {
	return _Extension.Contract.GetLastNewValue(&_Extension.CallOpts)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_Extension *ExtensionCaller) GetLastNewValueById(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, bool, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getLastNewValueById", _requestId)

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
func (_Extension *ExtensionSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _Extension.Contract.GetLastNewValueById(&_Extension.CallOpts, _requestId)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_Extension *ExtensionCallerSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _Extension.Contract.GetLastNewValueById(&_Extension.CallOpts, _requestId)
}

// GetMax5 is a free data retrieval call binding the contract method 0x99830e32.
//
// Solidity: function getMax5(uint256[51] data) view returns(uint256[5] max, uint256[5] maxIndex)
func (_Extension *ExtensionCaller) GetMax5(opts *bind.CallOpts, data [51]*big.Int) (struct {
	Max      [5]*big.Int
	MaxIndex [5]*big.Int
}, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getMax5", data)

	outstruct := new(struct {
		Max      [5]*big.Int
		MaxIndex [5]*big.Int
	})

	outstruct.Max = out[0].([5]*big.Int)
	outstruct.MaxIndex = out[1].([5]*big.Int)

	return *outstruct, err

}

// GetMax5 is a free data retrieval call binding the contract method 0x99830e32.
//
// Solidity: function getMax5(uint256[51] data) view returns(uint256[5] max, uint256[5] maxIndex)
func (_Extension *ExtensionSession) GetMax5(data [51]*big.Int) (struct {
	Max      [5]*big.Int
	MaxIndex [5]*big.Int
}, error) {
	return _Extension.Contract.GetMax5(&_Extension.CallOpts, data)
}

// GetMax5 is a free data retrieval call binding the contract method 0x99830e32.
//
// Solidity: function getMax5(uint256[51] data) view returns(uint256[5] max, uint256[5] maxIndex)
func (_Extension *ExtensionCallerSession) GetMax5(data [51]*big.Int) (struct {
	Max      [5]*big.Int
	MaxIndex [5]*big.Int
}, error) {
	return _Extension.Contract.GetMax5(&_Extension.CallOpts, data)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_Extension *ExtensionCaller) GetMinedBlockNum(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getMinedBlockNum", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_Extension *ExtensionSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _Extension.Contract.GetMinedBlockNum(&_Extension.CallOpts, _requestId, _timestamp)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_Extension *ExtensionCallerSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _Extension.Contract.GetMinedBlockNum(&_Extension.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_Extension *ExtensionCaller) GetMinersByRequestIdAndTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getMinersByRequestIdAndTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([5]common.Address)).(*[5]common.Address)

	return out0, err

}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_Extension *ExtensionSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _Extension.Contract.GetMinersByRequestIdAndTimestamp(&_Extension.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_Extension *ExtensionCallerSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _Extension.Contract.GetMinersByRequestIdAndTimestamp(&_Extension.CallOpts, _requestId, _timestamp)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _diff, uint256 _tip)
func (_Extension *ExtensionCaller) GetNewCurrentVariables(opts *bind.CallOpts) (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Diff       *big.Int
	Tip        *big.Int
}, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getNewCurrentVariables")

	outstruct := new(struct {
		Challenge  [32]byte
		RequestIds [5]*big.Int
		Diff       *big.Int
		Tip        *big.Int
	})

	outstruct.Challenge = out[0].([32]byte)
	outstruct.RequestIds = out[1].([5]*big.Int)
	outstruct.Diff = out[2].(*big.Int)
	outstruct.Tip = out[3].(*big.Int)

	return *outstruct, err

}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _diff, uint256 _tip)
func (_Extension *ExtensionSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Diff       *big.Int
	Tip        *big.Int
}, error) {
	return _Extension.Contract.GetNewCurrentVariables(&_Extension.CallOpts)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _diff, uint256 _tip)
func (_Extension *ExtensionCallerSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Diff       *big.Int
	Tip        *big.Int
}, error) {
	return _Extension.Contract.GetNewCurrentVariables(&_Extension.CallOpts)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_Extension *ExtensionCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getNewValueCountbyRequestId", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_Extension *ExtensionSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _Extension.Contract.GetNewValueCountbyRequestId(&_Extension.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_Extension *ExtensionCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _Extension.Contract.GetNewValueCountbyRequestId(&_Extension.CallOpts, _requestId)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_Extension *ExtensionCaller) GetNewVariablesOnDeck(opts *bind.CallOpts) (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getNewVariablesOnDeck")

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
func (_Extension *ExtensionSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _Extension.Contract.GetNewVariablesOnDeck(&_Extension.CallOpts)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_Extension *ExtensionCallerSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _Extension.Contract.GetNewVariablesOnDeck(&_Extension.CallOpts)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_Extension *ExtensionCaller) GetRequestIdByRequestQIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getRequestIdByRequestQIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_Extension *ExtensionSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _Extension.Contract.GetRequestIdByRequestQIndex(&_Extension.CallOpts, _index)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_Extension *ExtensionCallerSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _Extension.Contract.GetRequestIdByRequestQIndex(&_Extension.CallOpts, _index)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_Extension *ExtensionCaller) GetRequestIdByTimestamp(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getRequestIdByTimestamp", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_Extension *ExtensionSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _Extension.Contract.GetRequestIdByTimestamp(&_Extension.CallOpts, _timestamp)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_Extension *ExtensionCallerSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _Extension.Contract.GetRequestIdByTimestamp(&_Extension.CallOpts, _timestamp)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_Extension *ExtensionCaller) GetRequestQ(opts *bind.CallOpts) ([51]*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getRequestQ")

	if err != nil {
		return *new([51]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([51]*big.Int)).(*[51]*big.Int)

	return out0, err

}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_Extension *ExtensionSession) GetRequestQ() ([51]*big.Int, error) {
	return _Extension.Contract.GetRequestQ(&_Extension.CallOpts)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_Extension *ExtensionCallerSession) GetRequestQ() ([51]*big.Int, error) {
	return _Extension.Contract.GetRequestQ(&_Extension.CallOpts)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_Extension *ExtensionCaller) GetRequestUintVars(opts *bind.CallOpts, _requestId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getRequestUintVars", _requestId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_Extension *ExtensionSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _Extension.Contract.GetRequestUintVars(&_Extension.CallOpts, _requestId, _data)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_Extension *ExtensionCallerSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _Extension.Contract.GetRequestUintVars(&_Extension.CallOpts, _requestId, _data)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(uint256, uint256)
func (_Extension *ExtensionCaller) GetRequestVars(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getRequestVars", _requestId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(uint256, uint256)
func (_Extension *ExtensionSession) GetRequestVars(_requestId *big.Int) (*big.Int, *big.Int, error) {
	return _Extension.Contract.GetRequestVars(&_Extension.CallOpts, _requestId)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(uint256, uint256)
func (_Extension *ExtensionCallerSession) GetRequestVars(_requestId *big.Int) (*big.Int, *big.Int, error) {
	return _Extension.Contract.GetRequestVars(&_Extension.CallOpts, _requestId)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_Extension *ExtensionCaller) GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getStakerInfo", _staker)

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
func (_Extension *ExtensionSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _Extension.Contract.GetStakerInfo(&_Extension.CallOpts, _staker)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_Extension *ExtensionCallerSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _Extension.Contract.GetStakerInfo(&_Extension.CallOpts, _staker)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_Extension *ExtensionCaller) GetSubmissionsByTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getSubmissionsByTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_Extension *ExtensionSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _Extension.Contract.GetSubmissionsByTimestamp(&_Extension.CallOpts, _requestId, _timestamp)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_Extension *ExtensionCallerSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _Extension.Contract.GetSubmissionsByTimestamp(&_Extension.CallOpts, _requestId, _timestamp)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_Extension *ExtensionCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getTimestampbyRequestIDandIndex", _requestID, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_Extension *ExtensionSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _Extension.Contract.GetTimestampbyRequestIDandIndex(&_Extension.CallOpts, _requestID, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_Extension *ExtensionCallerSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _Extension.Contract.GetTimestampbyRequestIDandIndex(&_Extension.CallOpts, _requestID, _index)
}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_Extension *ExtensionCaller) GetTopRequestIDs(opts *bind.CallOpts) ([5]*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getTopRequestIDs")

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_Extension *ExtensionSession) GetTopRequestIDs() ([5]*big.Int, error) {
	return _Extension.Contract.GetTopRequestIDs(&_Extension.CallOpts)
}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_Extension *ExtensionCallerSession) GetTopRequestIDs() ([5]*big.Int, error) {
	return _Extension.Contract.GetTopRequestIDs(&_Extension.CallOpts)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_Extension *ExtensionCaller) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getUintVar", _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_Extension *ExtensionSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _Extension.Contract.GetUintVar(&_Extension.CallOpts, _data)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_Extension *ExtensionCallerSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _Extension.Contract.GetUintVar(&_Extension.CallOpts, _data)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_Extension *ExtensionCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "isInDispute", _requestId, _timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_Extension *ExtensionSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _Extension.Contract.IsInDispute(&_Extension.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_Extension *ExtensionCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _Extension.Contract.IsInDispute(&_Extension.CallOpts, _requestId, _timestamp)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_Extension *ExtensionCaller) Migrated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "migrated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_Extension *ExtensionSession) Migrated(arg0 common.Address) (bool, error) {
	return _Extension.Contract.Migrated(&_Extension.CallOpts, arg0)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_Extension *ExtensionCallerSession) Migrated(arg0 common.Address) (bool, error) {
	return _Extension.Contract.Migrated(&_Extension.CallOpts, arg0)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_Extension *ExtensionCaller) MinersByChallenge(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "minersByChallenge", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_Extension *ExtensionSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Extension.Contract.MinersByChallenge(&_Extension.CallOpts, arg0, arg1)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_Extension *ExtensionCallerSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Extension.Contract.MinersByChallenge(&_Extension.CallOpts, arg0, arg1)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Extension *ExtensionCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Extension *ExtensionSession) Name() (string, error) {
	return _Extension.Contract.Name(&_Extension.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Extension *ExtensionCallerSession) Name() (string, error) {
	return _Extension.Contract.Name(&_Extension.CallOpts)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_Extension *ExtensionCaller) NewValueTimestamps(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "newValueTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_Extension *ExtensionSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _Extension.Contract.NewValueTimestamps(&_Extension.CallOpts, arg0)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_Extension *ExtensionCallerSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _Extension.Contract.NewValueTimestamps(&_Extension.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_Extension *ExtensionCaller) RequestIdByQueryHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "requestIdByQueryHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_Extension *ExtensionSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _Extension.Contract.RequestIdByQueryHash(&_Extension.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_Extension *ExtensionCallerSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _Extension.Contract.RequestIdByQueryHash(&_Extension.CallOpts, arg0)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_Extension *ExtensionCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "retrieveData", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_Extension *ExtensionSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _Extension.Contract.RetrieveData(&_Extension.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_Extension *ExtensionCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _Extension.Contract.RetrieveData(&_Extension.CallOpts, _requestId, _timestamp)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Extension *ExtensionCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Extension *ExtensionSession) Symbol() (string, error) {
	return _Extension.Contract.Symbol(&_Extension.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Extension *ExtensionCallerSession) Symbol() (string, error) {
	return _Extension.Contract.Symbol(&_Extension.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Extension *ExtensionCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Extension *ExtensionSession) TotalSupply() (*big.Int, error) {
	return _Extension.Contract.TotalSupply(&_Extension.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Extension *ExtensionCallerSession) TotalSupply() (*big.Int, error) {
	return _Extension.Contract.TotalSupply(&_Extension.CallOpts)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_Extension *ExtensionCaller) Uints(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "uints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_Extension *ExtensionSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _Extension.Contract.Uints(&_Extension.CallOpts, arg0)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_Extension *ExtensionCallerSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _Extension.Contract.Uints(&_Extension.CallOpts, arg0)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Extension *ExtensionTransactor) DepositStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extension.contract.Transact(opts, "depositStake")
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Extension *ExtensionSession) DepositStake() (*types.Transaction, error) {
	return _Extension.Contract.DepositStake(&_Extension.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Extension *ExtensionTransactorSession) DepositStake() (*types.Transaction, error) {
	return _Extension.Contract.DepositStake(&_Extension.TransactOpts)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Extension *ExtensionTransactor) RequestStakingWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extension.contract.Transact(opts, "requestStakingWithdraw")
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Extension *ExtensionSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _Extension.Contract.RequestStakingWithdraw(&_Extension.TransactOpts)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Extension *ExtensionTransactorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _Extension.Contract.RequestStakingWithdraw(&_Extension.TransactOpts)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_Extension *ExtensionTransactor) TallyVotes(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _Extension.contract.Transact(opts, "tallyVotes", _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_Extension *ExtensionSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _Extension.Contract.TallyVotes(&_Extension.TransactOpts, _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_Extension *ExtensionTransactorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _Extension.Contract.TallyVotes(&_Extension.TransactOpts, _disputeId)
}

// UpdateMinDisputeFee is a paid mutator transaction binding the contract method 0x90e5b235.
//
// Solidity: function updateMinDisputeFee() returns()
func (_Extension *ExtensionTransactor) UpdateMinDisputeFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extension.contract.Transact(opts, "updateMinDisputeFee")
}

// UpdateMinDisputeFee is a paid mutator transaction binding the contract method 0x90e5b235.
//
// Solidity: function updateMinDisputeFee() returns()
func (_Extension *ExtensionSession) UpdateMinDisputeFee() (*types.Transaction, error) {
	return _Extension.Contract.UpdateMinDisputeFee(&_Extension.TransactOpts)
}

// UpdateMinDisputeFee is a paid mutator transaction binding the contract method 0x90e5b235.
//
// Solidity: function updateMinDisputeFee() returns()
func (_Extension *ExtensionTransactorSession) UpdateMinDisputeFee() (*types.Transaction, error) {
	return _Extension.Contract.UpdateMinDisputeFee(&_Extension.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Extension *ExtensionTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extension.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Extension *ExtensionSession) WithdrawStake() (*types.Transaction, error) {
	return _Extension.Contract.WithdrawStake(&_Extension.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Extension *ExtensionTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _Extension.Contract.WithdrawStake(&_Extension.TransactOpts)
}

// ExtensionDisputeVoteTalliedIterator is returned from FilterDisputeVoteTallied and is used to iterate over the raw logs and unpacked data for DisputeVoteTallied events raised by the Extension contract.
type ExtensionDisputeVoteTalliedIterator struct {
	Event *ExtensionDisputeVoteTallied // Event containing the contract specifics and raw log

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
func (it *ExtensionDisputeVoteTalliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionDisputeVoteTallied)
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
		it.Event = new(ExtensionDisputeVoteTallied)
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
func (it *ExtensionDisputeVoteTalliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionDisputeVoteTalliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionDisputeVoteTallied represents a DisputeVoteTallied event raised by the Extension contract.
type ExtensionDisputeVoteTallied struct {
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
func (_Extension *ExtensionFilterer) FilterDisputeVoteTallied(opts *bind.FilterOpts, _disputeID []*big.Int, _reportedMiner []common.Address) (*ExtensionDisputeVoteTalliedIterator, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _reportedMinerRule []interface{}
	for _, _reportedMinerItem := range _reportedMiner {
		_reportedMinerRule = append(_reportedMinerRule, _reportedMinerItem)
	}

	logs, sub, err := _Extension.contract.FilterLogs(opts, "DisputeVoteTallied", _disputeIDRule, _reportedMinerRule)
	if err != nil {
		return nil, err
	}
	return &ExtensionDisputeVoteTalliedIterator{contract: _Extension.contract, event: "DisputeVoteTallied", logs: logs, sub: sub}, nil
}

// WatchDisputeVoteTallied is a free log subscription operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: event DisputeVoteTallied(uint256 indexed _disputeID, int256 _result, address indexed _reportedMiner, address _reportingParty, bool _active)
func (_Extension *ExtensionFilterer) WatchDisputeVoteTallied(opts *bind.WatchOpts, sink chan<- *ExtensionDisputeVoteTallied, _disputeID []*big.Int, _reportedMiner []common.Address) (event.Subscription, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _reportedMinerRule []interface{}
	for _, _reportedMinerItem := range _reportedMiner {
		_reportedMinerRule = append(_reportedMinerRule, _reportedMinerItem)
	}

	logs, sub, err := _Extension.contract.WatchLogs(opts, "DisputeVoteTallied", _disputeIDRule, _reportedMinerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionDisputeVoteTallied)
				if err := _Extension.contract.UnpackLog(event, "DisputeVoteTallied", log); err != nil {
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
func (_Extension *ExtensionFilterer) ParseDisputeVoteTallied(log types.Log) (*ExtensionDisputeVoteTallied, error) {
	event := new(ExtensionDisputeVoteTallied)
	if err := _Extension.contract.UnpackLog(event, "DisputeVoteTallied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensionNewStakeIterator is returned from FilterNewStake and is used to iterate over the raw logs and unpacked data for NewStake events raised by the Extension contract.
type ExtensionNewStakeIterator struct {
	Event *ExtensionNewStake // Event containing the contract specifics and raw log

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
func (it *ExtensionNewStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionNewStake)
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
		it.Event = new(ExtensionNewStake)
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
func (it *ExtensionNewStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionNewStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionNewStake represents a NewStake event raised by the Extension contract.
type ExtensionNewStake struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewStake is a free log retrieval operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: event NewStake(address indexed _sender)
func (_Extension *ExtensionFilterer) FilterNewStake(opts *bind.FilterOpts, _sender []common.Address) (*ExtensionNewStakeIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Extension.contract.FilterLogs(opts, "NewStake", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ExtensionNewStakeIterator{contract: _Extension.contract, event: "NewStake", logs: logs, sub: sub}, nil
}

// WatchNewStake is a free log subscription operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: event NewStake(address indexed _sender)
func (_Extension *ExtensionFilterer) WatchNewStake(opts *bind.WatchOpts, sink chan<- *ExtensionNewStake, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Extension.contract.WatchLogs(opts, "NewStake", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionNewStake)
				if err := _Extension.contract.UnpackLog(event, "NewStake", log); err != nil {
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
func (_Extension *ExtensionFilterer) ParseNewStake(log types.Log) (*ExtensionNewStake, error) {
	event := new(ExtensionNewStake)
	if err := _Extension.contract.UnpackLog(event, "NewStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensionStakeWithdrawRequestedIterator is returned from FilterStakeWithdrawRequested and is used to iterate over the raw logs and unpacked data for StakeWithdrawRequested events raised by the Extension contract.
type ExtensionStakeWithdrawRequestedIterator struct {
	Event *ExtensionStakeWithdrawRequested // Event containing the contract specifics and raw log

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
func (it *ExtensionStakeWithdrawRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionStakeWithdrawRequested)
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
		it.Event = new(ExtensionStakeWithdrawRequested)
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
func (it *ExtensionStakeWithdrawRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionStakeWithdrawRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionStakeWithdrawRequested represents a StakeWithdrawRequested event raised by the Extension contract.
type ExtensionStakeWithdrawRequested struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawRequested is a free log retrieval operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: event StakeWithdrawRequested(address indexed _sender)
func (_Extension *ExtensionFilterer) FilterStakeWithdrawRequested(opts *bind.FilterOpts, _sender []common.Address) (*ExtensionStakeWithdrawRequestedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Extension.contract.FilterLogs(opts, "StakeWithdrawRequested", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ExtensionStakeWithdrawRequestedIterator{contract: _Extension.contract, event: "StakeWithdrawRequested", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawRequested is a free log subscription operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: event StakeWithdrawRequested(address indexed _sender)
func (_Extension *ExtensionFilterer) WatchStakeWithdrawRequested(opts *bind.WatchOpts, sink chan<- *ExtensionStakeWithdrawRequested, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Extension.contract.WatchLogs(opts, "StakeWithdrawRequested", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionStakeWithdrawRequested)
				if err := _Extension.contract.UnpackLog(event, "StakeWithdrawRequested", log); err != nil {
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
func (_Extension *ExtensionFilterer) ParseStakeWithdrawRequested(log types.Log) (*ExtensionStakeWithdrawRequested, error) {
	event := new(ExtensionStakeWithdrawRequested)
	if err := _Extension.contract.UnpackLog(event, "StakeWithdrawRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensionStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the Extension contract.
type ExtensionStakeWithdrawnIterator struct {
	Event *ExtensionStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *ExtensionStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionStakeWithdrawn)
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
		it.Event = new(ExtensionStakeWithdrawn)
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
func (it *ExtensionStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionStakeWithdrawn represents a StakeWithdrawn event raised by the Extension contract.
type ExtensionStakeWithdrawn struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: event StakeWithdrawn(address indexed _sender)
func (_Extension *ExtensionFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, _sender []common.Address) (*ExtensionStakeWithdrawnIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Extension.contract.FilterLogs(opts, "StakeWithdrawn", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ExtensionStakeWithdrawnIterator{contract: _Extension.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: event StakeWithdrawn(address indexed _sender)
func (_Extension *ExtensionFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *ExtensionStakeWithdrawn, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Extension.contract.WatchLogs(opts, "StakeWithdrawn", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionStakeWithdrawn)
				if err := _Extension.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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
func (_Extension *ExtensionFilterer) ParseStakeWithdrawn(log types.Log) (*ExtensionStakeWithdrawn, error) {
	event := new(ExtensionStakeWithdrawn)
	if err := _Extension.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorABI is the input ABI used to generate the binding from.
const ITellorABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"_result\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reportedMiner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_reportingParty\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"DisputeVoteTallied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_currentRequestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"NewChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"NewDispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"NewStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newTellor\",\"type\":\"address\"}],\"name\":\"NewTellorAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NewValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NonceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"TipAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_position\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_voteWeight\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minerIndex\",\"type\":\"uint256\"}],\"name\":\"beginDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newDeity\",\"type\":\"address\"}],\"name\":\"changeDeity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ext\",\"type\":\"address\"}],\"name\":\"changeExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"changeMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tellorContract\",\"type\":\"address\"}],\"name\":\"changeTellorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"getAddressVarByString\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[9]\",\"name\":\"\",\"type\":\"uint256[9]\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastNewValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"data\",\"type\":\"uint256[51]\"}],\"name\":\"getMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"data\",\"type\":\"uint256[51]\"}],\"name\":\"getMax5\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"max\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"maxIndex\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"data\",\"type\":\"uint256[51]\"}],\"name\":\"getMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"internalType\":\"address[5]\",\"name\":\"\",\"type\":\"address[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256\",\"name\":\"_difficutly\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"idsOnDeck\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"tipsOnDeck\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_request\",\"type\":\"bytes32\"}],\"name\":\"getRequestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"\",\"type\":\"uint256[51]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTopRequestIDs\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_diff\",\"type\":\"uint256\"}],\"name\":\"manuallySetDifficulty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_bypass\",\"type\":\"bool\"}],\"name\":\"migrateFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_destination\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amount\",\"type\":\"uint256[]\"}],\"name\":\"migrateForBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_origin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_bypass\",\"type\":\"bool\"}],\"name\":\"migrateFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_origin\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_destination\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amount\",\"type\":\"uint256[]\"}],\"name\":\"migrateFromBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_propNewTellorAddress\",\"type\":\"address\"}],\"name\":\"proposeFork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_pendingOwner\",\"type\":\"address\"}],\"name\":\"proposeOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestStakingWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"}],\"name\":\"submitMiningSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"tallyVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"}],\"name\":\"testSubmitMiningSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"theLazyCoon\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"unlockDisputeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"updateTellor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_supportsDispute\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ITellorFuncSigs maps the 4-byte function signature to its string representation.
var ITellorFuncSigs = map[string]string{
	"752d49a1": "addTip(uint256,uint256)",
	"dd62ed3e": "allowance(address,address)",
	"999cf26c": "allowedToTrade(address,uint256)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"4ee2cd7e": "balanceOfAt(address,uint256)",
	"8581af19": "beginDispute(uint256,uint256,uint256)",
	"47abd7f1": "changeDeity(address)",
	"b69a363f": "changeExtension(address)",
	"141e13fa": "changeMigrator(address)",
	"ae0a8279": "changeTellorContract(address)",
	"4e71e0c8": "claimOwnership()",
	"313ce567": "decimals()",
	"0d2d76a2": "depositStake()",
	"63bb82ad": "didMine(bytes32,address)",
	"a7c438bc": "didVote(uint256,address)",
	"9cc128e0": "getAddressVarByString(string)",
	"133bee5e": "getAddressVars(bytes32)",
	"af0b1327": "getAllDisputeVars(uint256)",
	"a22e407a": "getCurrentVariables()",
	"da379941": "getDisputeIdByDisputeHash(bytes32)",
	"7f6fd5d9": "getDisputeUintVars(uint256,bytes32)",
	"fc7cf0a0": "getLastNewValue()",
	"3180f8df": "getLastNewValueById(uint256)",
	"c87a336d": "getMax(uint256[51])",
	"99830e32": "getMax5(uint256[51])",
	"f29e5e9a": "getMin(uint256[51])",
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
	"c52e9539": "manuallySetDifficulty(uint256)",
	"8fd3ab80": "migrate()",
	"a9fa7d34": "migrateFor(address,uint256,bool)",
	"42a89bd6": "migrateForBatch(address[],uint256[])",
	"121dd372": "migrateFrom(address,address,uint256,bool)",
	"8c0f4076": "migrateFromBatch(address[],address[],uint256[])",
	"06fdde03": "name()",
	"26b7d9f6": "proposeFork(address)",
	"710bf322": "proposeOwnership(address)",
	"28449c3a": "requestStakingWithdraw()",
	"93fa4915": "retrieveData(uint256,uint256)",
	"4350283e": "submitMiningSolution(string,uint256[5],uint256[5])",
	"95d89b41": "symbol()",
	"4d318b0e": "tallyVotes(uint256)",
	"d47f0dd4": "testSubmitMiningSolution(string,uint256[5],uint256[5])",
	"b079f64a": "theLazyCoon(address,uint256)",
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

// GetAddressVarByString is a free data retrieval call binding the contract method 0x9cc128e0.
//
// Solidity: function getAddressVarByString(string _data) view returns(address)
func (_ITellor *ITellorCaller) GetAddressVarByString(opts *bind.CallOpts, _data string) (common.Address, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getAddressVarByString", _data)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressVarByString is a free data retrieval call binding the contract method 0x9cc128e0.
//
// Solidity: function getAddressVarByString(string _data) view returns(address)
func (_ITellor *ITellorSession) GetAddressVarByString(_data string) (common.Address, error) {
	return _ITellor.Contract.GetAddressVarByString(&_ITellor.CallOpts, _data)
}

// GetAddressVarByString is a free data retrieval call binding the contract method 0x9cc128e0.
//
// Solidity: function getAddressVarByString(string _data) view returns(address)
func (_ITellor *ITellorCallerSession) GetAddressVarByString(_data string) (common.Address, error) {
	return _ITellor.Contract.GetAddressVarByString(&_ITellor.CallOpts, _data)
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

// GetMax is a free data retrieval call binding the contract method 0xc87a336d.
//
// Solidity: function getMax(uint256[51] data) view returns(uint256 max, uint256 maxIndex)
func (_ITellor *ITellorCaller) GetMax(opts *bind.CallOpts, data [51]*big.Int) (struct {
	Max      *big.Int
	MaxIndex *big.Int
}, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getMax", data)

	outstruct := new(struct {
		Max      *big.Int
		MaxIndex *big.Int
	})

	outstruct.Max = out[0].(*big.Int)
	outstruct.MaxIndex = out[1].(*big.Int)

	return *outstruct, err

}

// GetMax is a free data retrieval call binding the contract method 0xc87a336d.
//
// Solidity: function getMax(uint256[51] data) view returns(uint256 max, uint256 maxIndex)
func (_ITellor *ITellorSession) GetMax(data [51]*big.Int) (struct {
	Max      *big.Int
	MaxIndex *big.Int
}, error) {
	return _ITellor.Contract.GetMax(&_ITellor.CallOpts, data)
}

// GetMax is a free data retrieval call binding the contract method 0xc87a336d.
//
// Solidity: function getMax(uint256[51] data) view returns(uint256 max, uint256 maxIndex)
func (_ITellor *ITellorCallerSession) GetMax(data [51]*big.Int) (struct {
	Max      *big.Int
	MaxIndex *big.Int
}, error) {
	return _ITellor.Contract.GetMax(&_ITellor.CallOpts, data)
}

// GetMax5 is a free data retrieval call binding the contract method 0x99830e32.
//
// Solidity: function getMax5(uint256[51] data) view returns(uint256[5] max, uint256[5] maxIndex)
func (_ITellor *ITellorCaller) GetMax5(opts *bind.CallOpts, data [51]*big.Int) (struct {
	Max      [5]*big.Int
	MaxIndex [5]*big.Int
}, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getMax5", data)

	outstruct := new(struct {
		Max      [5]*big.Int
		MaxIndex [5]*big.Int
	})

	outstruct.Max = out[0].([5]*big.Int)
	outstruct.MaxIndex = out[1].([5]*big.Int)

	return *outstruct, err

}

// GetMax5 is a free data retrieval call binding the contract method 0x99830e32.
//
// Solidity: function getMax5(uint256[51] data) view returns(uint256[5] max, uint256[5] maxIndex)
func (_ITellor *ITellorSession) GetMax5(data [51]*big.Int) (struct {
	Max      [5]*big.Int
	MaxIndex [5]*big.Int
}, error) {
	return _ITellor.Contract.GetMax5(&_ITellor.CallOpts, data)
}

// GetMax5 is a free data retrieval call binding the contract method 0x99830e32.
//
// Solidity: function getMax5(uint256[51] data) view returns(uint256[5] max, uint256[5] maxIndex)
func (_ITellor *ITellorCallerSession) GetMax5(data [51]*big.Int) (struct {
	Max      [5]*big.Int
	MaxIndex [5]*big.Int
}, error) {
	return _ITellor.Contract.GetMax5(&_ITellor.CallOpts, data)
}

// GetMin is a free data retrieval call binding the contract method 0xf29e5e9a.
//
// Solidity: function getMin(uint256[51] data) view returns(uint256 min, uint256 minIndex)
func (_ITellor *ITellorCaller) GetMin(opts *bind.CallOpts, data [51]*big.Int) (struct {
	Min      *big.Int
	MinIndex *big.Int
}, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getMin", data)

	outstruct := new(struct {
		Min      *big.Int
		MinIndex *big.Int
	})

	outstruct.Min = out[0].(*big.Int)
	outstruct.MinIndex = out[1].(*big.Int)

	return *outstruct, err

}

// GetMin is a free data retrieval call binding the contract method 0xf29e5e9a.
//
// Solidity: function getMin(uint256[51] data) view returns(uint256 min, uint256 minIndex)
func (_ITellor *ITellorSession) GetMin(data [51]*big.Int) (struct {
	Min      *big.Int
	MinIndex *big.Int
}, error) {
	return _ITellor.Contract.GetMin(&_ITellor.CallOpts, data)
}

// GetMin is a free data retrieval call binding the contract method 0xf29e5e9a.
//
// Solidity: function getMin(uint256[51] data) view returns(uint256 min, uint256 minIndex)
func (_ITellor *ITellorCallerSession) GetMin(data [51]*big.Int) (struct {
	Min      *big.Int
	MinIndex *big.Int
}, error) {
	return _ITellor.Contract.GetMin(&_ITellor.CallOpts, data)
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
// Solidity: function getRequestVars(uint256 _requestId) view returns(uint256, uint256)
func (_ITellor *ITellorCaller) GetRequestVars(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getRequestVars", _requestId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(uint256, uint256)
func (_ITellor *ITellorSession) GetRequestVars(_requestId *big.Int) (*big.Int, *big.Int, error) {
	return _ITellor.Contract.GetRequestVars(&_ITellor.CallOpts, _requestId)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(uint256, uint256)
func (_ITellor *ITellorCallerSession) GetRequestVars(_requestId *big.Int) (*big.Int, *big.Int, error) {
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

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(address _newDeity) returns()
func (_ITellor *ITellorTransactor) ChangeDeity(opts *bind.TransactOpts, _newDeity common.Address) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "changeDeity", _newDeity)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(address _newDeity) returns()
func (_ITellor *ITellorSession) ChangeDeity(_newDeity common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ChangeDeity(&_ITellor.TransactOpts, _newDeity)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(address _newDeity) returns()
func (_ITellor *ITellorTransactorSession) ChangeDeity(_newDeity common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ChangeDeity(&_ITellor.TransactOpts, _newDeity)
}

// ChangeExtension is a paid mutator transaction binding the contract method 0xb69a363f.
//
// Solidity: function changeExtension(address _ext) returns()
func (_ITellor *ITellorTransactor) ChangeExtension(opts *bind.TransactOpts, _ext common.Address) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "changeExtension", _ext)
}

// ChangeExtension is a paid mutator transaction binding the contract method 0xb69a363f.
//
// Solidity: function changeExtension(address _ext) returns()
func (_ITellor *ITellorSession) ChangeExtension(_ext common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ChangeExtension(&_ITellor.TransactOpts, _ext)
}

// ChangeExtension is a paid mutator transaction binding the contract method 0xb69a363f.
//
// Solidity: function changeExtension(address _ext) returns()
func (_ITellor *ITellorTransactorSession) ChangeExtension(_ext common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ChangeExtension(&_ITellor.TransactOpts, _ext)
}

// ChangeMigrator is a paid mutator transaction binding the contract method 0x141e13fa.
//
// Solidity: function changeMigrator(address _migrator) returns()
func (_ITellor *ITellorTransactor) ChangeMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "changeMigrator", _migrator)
}

// ChangeMigrator is a paid mutator transaction binding the contract method 0x141e13fa.
//
// Solidity: function changeMigrator(address _migrator) returns()
func (_ITellor *ITellorSession) ChangeMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ChangeMigrator(&_ITellor.TransactOpts, _migrator)
}

// ChangeMigrator is a paid mutator transaction binding the contract method 0x141e13fa.
//
// Solidity: function changeMigrator(address _migrator) returns()
func (_ITellor *ITellorTransactorSession) ChangeMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ChangeMigrator(&_ITellor.TransactOpts, _migrator)
}

// ChangeTellorContract is a paid mutator transaction binding the contract method 0xae0a8279.
//
// Solidity: function changeTellorContract(address _tellorContract) returns()
func (_ITellor *ITellorTransactor) ChangeTellorContract(opts *bind.TransactOpts, _tellorContract common.Address) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "changeTellorContract", _tellorContract)
}

// ChangeTellorContract is a paid mutator transaction binding the contract method 0xae0a8279.
//
// Solidity: function changeTellorContract(address _tellorContract) returns()
func (_ITellor *ITellorSession) ChangeTellorContract(_tellorContract common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ChangeTellorContract(&_ITellor.TransactOpts, _tellorContract)
}

// ChangeTellorContract is a paid mutator transaction binding the contract method 0xae0a8279.
//
// Solidity: function changeTellorContract(address _tellorContract) returns()
func (_ITellor *ITellorTransactorSession) ChangeTellorContract(_tellorContract common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ChangeTellorContract(&_ITellor.TransactOpts, _tellorContract)
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

// ManuallySetDifficulty is a paid mutator transaction binding the contract method 0xc52e9539.
//
// Solidity: function manuallySetDifficulty(uint256 _diff) returns()
func (_ITellor *ITellorTransactor) ManuallySetDifficulty(opts *bind.TransactOpts, _diff *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "manuallySetDifficulty", _diff)
}

// ManuallySetDifficulty is a paid mutator transaction binding the contract method 0xc52e9539.
//
// Solidity: function manuallySetDifficulty(uint256 _diff) returns()
func (_ITellor *ITellorSession) ManuallySetDifficulty(_diff *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.ManuallySetDifficulty(&_ITellor.TransactOpts, _diff)
}

// ManuallySetDifficulty is a paid mutator transaction binding the contract method 0xc52e9539.
//
// Solidity: function manuallySetDifficulty(uint256 _diff) returns()
func (_ITellor *ITellorTransactorSession) ManuallySetDifficulty(_diff *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.ManuallySetDifficulty(&_ITellor.TransactOpts, _diff)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_ITellor *ITellorTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_ITellor *ITellorSession) Migrate() (*types.Transaction, error) {
	return _ITellor.Contract.Migrate(&_ITellor.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_ITellor *ITellorTransactorSession) Migrate() (*types.Transaction, error) {
	return _ITellor.Contract.Migrate(&_ITellor.TransactOpts)
}

// MigrateFor is a paid mutator transaction binding the contract method 0xa9fa7d34.
//
// Solidity: function migrateFor(address _destination, uint256 _amount, bool _bypass) returns()
func (_ITellor *ITellorTransactor) MigrateFor(opts *bind.TransactOpts, _destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "migrateFor", _destination, _amount, _bypass)
}

// MigrateFor is a paid mutator transaction binding the contract method 0xa9fa7d34.
//
// Solidity: function migrateFor(address _destination, uint256 _amount, bool _bypass) returns()
func (_ITellor *ITellorSession) MigrateFor(_destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _ITellor.Contract.MigrateFor(&_ITellor.TransactOpts, _destination, _amount, _bypass)
}

// MigrateFor is a paid mutator transaction binding the contract method 0xa9fa7d34.
//
// Solidity: function migrateFor(address _destination, uint256 _amount, bool _bypass) returns()
func (_ITellor *ITellorTransactorSession) MigrateFor(_destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _ITellor.Contract.MigrateFor(&_ITellor.TransactOpts, _destination, _amount, _bypass)
}

// MigrateForBatch is a paid mutator transaction binding the contract method 0x42a89bd6.
//
// Solidity: function migrateForBatch(address[] _destination, uint256[] _amount) returns()
func (_ITellor *ITellorTransactor) MigrateForBatch(opts *bind.TransactOpts, _destination []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "migrateForBatch", _destination, _amount)
}

// MigrateForBatch is a paid mutator transaction binding the contract method 0x42a89bd6.
//
// Solidity: function migrateForBatch(address[] _destination, uint256[] _amount) returns()
func (_ITellor *ITellorSession) MigrateForBatch(_destination []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.MigrateForBatch(&_ITellor.TransactOpts, _destination, _amount)
}

// MigrateForBatch is a paid mutator transaction binding the contract method 0x42a89bd6.
//
// Solidity: function migrateForBatch(address[] _destination, uint256[] _amount) returns()
func (_ITellor *ITellorTransactorSession) MigrateForBatch(_destination []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.MigrateForBatch(&_ITellor.TransactOpts, _destination, _amount)
}

// MigrateFrom is a paid mutator transaction binding the contract method 0x121dd372.
//
// Solidity: function migrateFrom(address _origin, address _destination, uint256 _amount, bool _bypass) returns()
func (_ITellor *ITellorTransactor) MigrateFrom(opts *bind.TransactOpts, _origin common.Address, _destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "migrateFrom", _origin, _destination, _amount, _bypass)
}

// MigrateFrom is a paid mutator transaction binding the contract method 0x121dd372.
//
// Solidity: function migrateFrom(address _origin, address _destination, uint256 _amount, bool _bypass) returns()
func (_ITellor *ITellorSession) MigrateFrom(_origin common.Address, _destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _ITellor.Contract.MigrateFrom(&_ITellor.TransactOpts, _origin, _destination, _amount, _bypass)
}

// MigrateFrom is a paid mutator transaction binding the contract method 0x121dd372.
//
// Solidity: function migrateFrom(address _origin, address _destination, uint256 _amount, bool _bypass) returns()
func (_ITellor *ITellorTransactorSession) MigrateFrom(_origin common.Address, _destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _ITellor.Contract.MigrateFrom(&_ITellor.TransactOpts, _origin, _destination, _amount, _bypass)
}

// MigrateFromBatch is a paid mutator transaction binding the contract method 0x8c0f4076.
//
// Solidity: function migrateFromBatch(address[] _origin, address[] _destination, uint256[] _amount) returns()
func (_ITellor *ITellorTransactor) MigrateFromBatch(opts *bind.TransactOpts, _origin []common.Address, _destination []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "migrateFromBatch", _origin, _destination, _amount)
}

// MigrateFromBatch is a paid mutator transaction binding the contract method 0x8c0f4076.
//
// Solidity: function migrateFromBatch(address[] _origin, address[] _destination, uint256[] _amount) returns()
func (_ITellor *ITellorSession) MigrateFromBatch(_origin []common.Address, _destination []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.MigrateFromBatch(&_ITellor.TransactOpts, _origin, _destination, _amount)
}

// MigrateFromBatch is a paid mutator transaction binding the contract method 0x8c0f4076.
//
// Solidity: function migrateFromBatch(address[] _origin, address[] _destination, uint256[] _amount) returns()
func (_ITellor *ITellorTransactorSession) MigrateFromBatch(_origin []common.Address, _destination []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.MigrateFromBatch(&_ITellor.TransactOpts, _origin, _destination, _amount)
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

// TestSubmitMiningSolution is a paid mutator transaction binding the contract method 0xd47f0dd4.
//
// Solidity: function testSubmitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_ITellor *ITellorTransactor) TestSubmitMiningSolution(opts *bind.TransactOpts, _nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "testSubmitMiningSolution", _nonce, _requestId, _value)
}

// TestSubmitMiningSolution is a paid mutator transaction binding the contract method 0xd47f0dd4.
//
// Solidity: function testSubmitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_ITellor *ITellorSession) TestSubmitMiningSolution(_nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.TestSubmitMiningSolution(&_ITellor.TransactOpts, _nonce, _requestId, _value)
}

// TestSubmitMiningSolution is a paid mutator transaction binding the contract method 0xd47f0dd4.
//
// Solidity: function testSubmitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_ITellor *ITellorTransactorSession) TestSubmitMiningSolution(_nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.TestSubmitMiningSolution(&_ITellor.TransactOpts, _nonce, _requestId, _value)
}

// TheLazyCoon is a paid mutator transaction binding the contract method 0xb079f64a.
//
// Solidity: function theLazyCoon(address _address, uint256 _amount) returns()
func (_ITellor *ITellorTransactor) TheLazyCoon(opts *bind.TransactOpts, _address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "theLazyCoon", _address, _amount)
}

// TheLazyCoon is a paid mutator transaction binding the contract method 0xb079f64a.
//
// Solidity: function theLazyCoon(address _address, uint256 _amount) returns()
func (_ITellor *ITellorSession) TheLazyCoon(_address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.TheLazyCoon(&_ITellor.TransactOpts, _address, _amount)
}

// TheLazyCoon is a paid mutator transaction binding the contract method 0xb079f64a.
//
// Solidity: function theLazyCoon(address _address, uint256 _amount) returns()
func (_ITellor *ITellorTransactorSession) TheLazyCoon(_address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.TheLazyCoon(&_ITellor.TransactOpts, _address, _amount)
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

// ITellorApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ITellor contract.
type ITellorApprovalIterator struct {
	Event *ITellorApproval // Event containing the contract specifics and raw log

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
func (it *ITellorApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorApproval)
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
		it.Event = new(ITellorApproval)
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
func (it *ITellorApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorApproval represents a Approval event raised by the ITellor contract.
type ITellorApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_ITellor *ITellorFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*ITellorApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &ITellorApprovalIterator{contract: _ITellor.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_ITellor *ITellorFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ITellorApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorApproval)
				if err := _ITellor.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ITellor *ITellorFilterer) ParseApproval(log types.Log) (*ITellorApproval, error) {
	event := new(ITellorApproval)
	if err := _ITellor.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorDisputeVoteTalliedIterator is returned from FilterDisputeVoteTallied and is used to iterate over the raw logs and unpacked data for DisputeVoteTallied events raised by the ITellor contract.
type ITellorDisputeVoteTalliedIterator struct {
	Event *ITellorDisputeVoteTallied // Event containing the contract specifics and raw log

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
func (it *ITellorDisputeVoteTalliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorDisputeVoteTallied)
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
		it.Event = new(ITellorDisputeVoteTallied)
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
func (it *ITellorDisputeVoteTalliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorDisputeVoteTalliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorDisputeVoteTallied represents a DisputeVoteTallied event raised by the ITellor contract.
type ITellorDisputeVoteTallied struct {
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
func (_ITellor *ITellorFilterer) FilterDisputeVoteTallied(opts *bind.FilterOpts, _disputeID []*big.Int, _reportedMiner []common.Address) (*ITellorDisputeVoteTalliedIterator, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _reportedMinerRule []interface{}
	for _, _reportedMinerItem := range _reportedMiner {
		_reportedMinerRule = append(_reportedMinerRule, _reportedMinerItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "DisputeVoteTallied", _disputeIDRule, _reportedMinerRule)
	if err != nil {
		return nil, err
	}
	return &ITellorDisputeVoteTalliedIterator{contract: _ITellor.contract, event: "DisputeVoteTallied", logs: logs, sub: sub}, nil
}

// WatchDisputeVoteTallied is a free log subscription operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: event DisputeVoteTallied(uint256 indexed _disputeID, int256 _result, address indexed _reportedMiner, address _reportingParty, bool _active)
func (_ITellor *ITellorFilterer) WatchDisputeVoteTallied(opts *bind.WatchOpts, sink chan<- *ITellorDisputeVoteTallied, _disputeID []*big.Int, _reportedMiner []common.Address) (event.Subscription, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _reportedMinerRule []interface{}
	for _, _reportedMinerItem := range _reportedMiner {
		_reportedMinerRule = append(_reportedMinerRule, _reportedMinerItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "DisputeVoteTallied", _disputeIDRule, _reportedMinerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorDisputeVoteTallied)
				if err := _ITellor.contract.UnpackLog(event, "DisputeVoteTallied", log); err != nil {
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
func (_ITellor *ITellorFilterer) ParseDisputeVoteTallied(log types.Log) (*ITellorDisputeVoteTallied, error) {
	event := new(ITellorDisputeVoteTallied)
	if err := _ITellor.contract.UnpackLog(event, "DisputeVoteTallied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorNewChallengeIterator is returned from FilterNewChallenge and is used to iterate over the raw logs and unpacked data for NewChallenge events raised by the ITellor contract.
type ITellorNewChallengeIterator struct {
	Event *ITellorNewChallenge // Event containing the contract specifics and raw log

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
func (it *ITellorNewChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorNewChallenge)
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
		it.Event = new(ITellorNewChallenge)
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
func (it *ITellorNewChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorNewChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorNewChallenge represents a NewChallenge event raised by the ITellor contract.
type ITellorNewChallenge struct {
	CurrentChallenge [32]byte
	CurrentRequestId [5]*big.Int
	Difficulty       *big.Int
	TotalTips        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewChallenge is a free log retrieval operation binding the contract event 0x1d85ce10456e29b67de37887496d3f1fcf1b64c79c4d07484038703a9f5c1408.
//
// Solidity: event NewChallenge(bytes32 indexed _currentChallenge, uint256[5] _currentRequestId, uint256 _difficulty, uint256 _totalTips)
func (_ITellor *ITellorFilterer) FilterNewChallenge(opts *bind.FilterOpts, _currentChallenge [][32]byte) (*ITellorNewChallengeIterator, error) {

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "NewChallenge", _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return &ITellorNewChallengeIterator{contract: _ITellor.contract, event: "NewChallenge", logs: logs, sub: sub}, nil
}

// WatchNewChallenge is a free log subscription operation binding the contract event 0x1d85ce10456e29b67de37887496d3f1fcf1b64c79c4d07484038703a9f5c1408.
//
// Solidity: event NewChallenge(bytes32 indexed _currentChallenge, uint256[5] _currentRequestId, uint256 _difficulty, uint256 _totalTips)
func (_ITellor *ITellorFilterer) WatchNewChallenge(opts *bind.WatchOpts, sink chan<- *ITellorNewChallenge, _currentChallenge [][32]byte) (event.Subscription, error) {

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "NewChallenge", _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorNewChallenge)
				if err := _ITellor.contract.UnpackLog(event, "NewChallenge", log); err != nil {
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
func (_ITellor *ITellorFilterer) ParseNewChallenge(log types.Log) (*ITellorNewChallenge, error) {
	event := new(ITellorNewChallenge)
	if err := _ITellor.contract.UnpackLog(event, "NewChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorNewDisputeIterator is returned from FilterNewDispute and is used to iterate over the raw logs and unpacked data for NewDispute events raised by the ITellor contract.
type ITellorNewDisputeIterator struct {
	Event *ITellorNewDispute // Event containing the contract specifics and raw log

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
func (it *ITellorNewDisputeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorNewDispute)
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
		it.Event = new(ITellorNewDispute)
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
func (it *ITellorNewDisputeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorNewDisputeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorNewDispute represents a NewDispute event raised by the ITellor contract.
type ITellorNewDispute struct {
	DisputeId *big.Int
	RequestId *big.Int
	Timestamp *big.Int
	Miner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewDispute is a free log retrieval operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_ITellor *ITellorFilterer) FilterNewDispute(opts *bind.FilterOpts, _disputeId []*big.Int, _requestId []*big.Int) (*ITellorNewDisputeIterator, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ITellorNewDisputeIterator{contract: _ITellor.contract, event: "NewDispute", logs: logs, sub: sub}, nil
}

// WatchNewDispute is a free log subscription operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_ITellor *ITellorFilterer) WatchNewDispute(opts *bind.WatchOpts, sink chan<- *ITellorNewDispute, _disputeId []*big.Int, _requestId []*big.Int) (event.Subscription, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorNewDispute)
				if err := _ITellor.contract.UnpackLog(event, "NewDispute", log); err != nil {
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
func (_ITellor *ITellorFilterer) ParseNewDispute(log types.Log) (*ITellorNewDispute, error) {
	event := new(ITellorNewDispute)
	if err := _ITellor.contract.UnpackLog(event, "NewDispute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorNewStakeIterator is returned from FilterNewStake and is used to iterate over the raw logs and unpacked data for NewStake events raised by the ITellor contract.
type ITellorNewStakeIterator struct {
	Event *ITellorNewStake // Event containing the contract specifics and raw log

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
func (it *ITellorNewStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorNewStake)
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
		it.Event = new(ITellorNewStake)
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
func (it *ITellorNewStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorNewStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorNewStake represents a NewStake event raised by the ITellor contract.
type ITellorNewStake struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewStake is a free log retrieval operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: event NewStake(address indexed _sender)
func (_ITellor *ITellorFilterer) FilterNewStake(opts *bind.FilterOpts, _sender []common.Address) (*ITellorNewStakeIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "NewStake", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ITellorNewStakeIterator{contract: _ITellor.contract, event: "NewStake", logs: logs, sub: sub}, nil
}

// WatchNewStake is a free log subscription operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: event NewStake(address indexed _sender)
func (_ITellor *ITellorFilterer) WatchNewStake(opts *bind.WatchOpts, sink chan<- *ITellorNewStake, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "NewStake", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorNewStake)
				if err := _ITellor.contract.UnpackLog(event, "NewStake", log); err != nil {
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
func (_ITellor *ITellorFilterer) ParseNewStake(log types.Log) (*ITellorNewStake, error) {
	event := new(ITellorNewStake)
	if err := _ITellor.contract.UnpackLog(event, "NewStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorNewTellorAddressIterator is returned from FilterNewTellorAddress and is used to iterate over the raw logs and unpacked data for NewTellorAddress events raised by the ITellor contract.
type ITellorNewTellorAddressIterator struct {
	Event *ITellorNewTellorAddress // Event containing the contract specifics and raw log

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
func (it *ITellorNewTellorAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorNewTellorAddress)
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
		it.Event = new(ITellorNewTellorAddress)
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
func (it *ITellorNewTellorAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorNewTellorAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorNewTellorAddress represents a NewTellorAddress event raised by the ITellor contract.
type ITellorNewTellorAddress struct {
	NewTellor common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTellorAddress is a free log retrieval operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: event NewTellorAddress(address _newTellor)
func (_ITellor *ITellorFilterer) FilterNewTellorAddress(opts *bind.FilterOpts) (*ITellorNewTellorAddressIterator, error) {

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "NewTellorAddress")
	if err != nil {
		return nil, err
	}
	return &ITellorNewTellorAddressIterator{contract: _ITellor.contract, event: "NewTellorAddress", logs: logs, sub: sub}, nil
}

// WatchNewTellorAddress is a free log subscription operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: event NewTellorAddress(address _newTellor)
func (_ITellor *ITellorFilterer) WatchNewTellorAddress(opts *bind.WatchOpts, sink chan<- *ITellorNewTellorAddress) (event.Subscription, error) {

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "NewTellorAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorNewTellorAddress)
				if err := _ITellor.contract.UnpackLog(event, "NewTellorAddress", log); err != nil {
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
func (_ITellor *ITellorFilterer) ParseNewTellorAddress(log types.Log) (*ITellorNewTellorAddress, error) {
	event := new(ITellorNewTellorAddress)
	if err := _ITellor.contract.UnpackLog(event, "NewTellorAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorNewValueIterator is returned from FilterNewValue and is used to iterate over the raw logs and unpacked data for NewValue events raised by the ITellor contract.
type ITellorNewValueIterator struct {
	Event *ITellorNewValue // Event containing the contract specifics and raw log

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
func (it *ITellorNewValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorNewValue)
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
		it.Event = new(ITellorNewValue)
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
func (it *ITellorNewValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorNewValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorNewValue represents a NewValue event raised by the ITellor contract.
type ITellorNewValue struct {
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
func (_ITellor *ITellorFilterer) FilterNewValue(opts *bind.FilterOpts, _currentChallenge [][32]byte) (*ITellorNewValueIterator, error) {

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "NewValue", _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return &ITellorNewValueIterator{contract: _ITellor.contract, event: "NewValue", logs: logs, sub: sub}, nil
}

// WatchNewValue is a free log subscription operation binding the contract event 0xbeb3b9f46c8d7bb00c873fca10d307538df350866d25f891ffb395147ddbdc45.
//
// Solidity: event NewValue(uint256[5] _requestId, uint256 _time, uint256[5] _value, uint256 _totalTips, bytes32 indexed _currentChallenge)
func (_ITellor *ITellorFilterer) WatchNewValue(opts *bind.WatchOpts, sink chan<- *ITellorNewValue, _currentChallenge [][32]byte) (event.Subscription, error) {

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "NewValue", _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorNewValue)
				if err := _ITellor.contract.UnpackLog(event, "NewValue", log); err != nil {
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
func (_ITellor *ITellorFilterer) ParseNewValue(log types.Log) (*ITellorNewValue, error) {
	event := new(ITellorNewValue)
	if err := _ITellor.contract.UnpackLog(event, "NewValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorNonceSubmittedIterator is returned from FilterNonceSubmitted and is used to iterate over the raw logs and unpacked data for NonceSubmitted events raised by the ITellor contract.
type ITellorNonceSubmittedIterator struct {
	Event *ITellorNonceSubmitted // Event containing the contract specifics and raw log

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
func (it *ITellorNonceSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorNonceSubmitted)
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
		it.Event = new(ITellorNonceSubmitted)
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
func (it *ITellorNonceSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorNonceSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorNonceSubmitted represents a NonceSubmitted event raised by the ITellor contract.
type ITellorNonceSubmitted struct {
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
func (_ITellor *ITellorFilterer) FilterNonceSubmitted(opts *bind.FilterOpts, _miner []common.Address, _currentChallenge [][32]byte) (*ITellorNonceSubmittedIterator, error) {

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "NonceSubmitted", _minerRule, _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return &ITellorNonceSubmittedIterator{contract: _ITellor.contract, event: "NonceSubmitted", logs: logs, sub: sub}, nil
}

// WatchNonceSubmitted is a free log subscription operation binding the contract event 0x0e4e65dc389613b6884b7f8c615e54fd3b894fbbbc534c990037744eea942000.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256[5] _requestId, uint256[5] _value, bytes32 indexed _currentChallenge)
func (_ITellor *ITellorFilterer) WatchNonceSubmitted(opts *bind.WatchOpts, sink chan<- *ITellorNonceSubmitted, _miner []common.Address, _currentChallenge [][32]byte) (event.Subscription, error) {

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "NonceSubmitted", _minerRule, _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorNonceSubmitted)
				if err := _ITellor.contract.UnpackLog(event, "NonceSubmitted", log); err != nil {
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
func (_ITellor *ITellorFilterer) ParseNonceSubmitted(log types.Log) (*ITellorNonceSubmitted, error) {
	event := new(ITellorNonceSubmitted)
	if err := _ITellor.contract.UnpackLog(event, "NonceSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorOwnershipProposedIterator is returned from FilterOwnershipProposed and is used to iterate over the raw logs and unpacked data for OwnershipProposed events raised by the ITellor contract.
type ITellorOwnershipProposedIterator struct {
	Event *ITellorOwnershipProposed // Event containing the contract specifics and raw log

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
func (it *ITellorOwnershipProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorOwnershipProposed)
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
		it.Event = new(ITellorOwnershipProposed)
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
func (it *ITellorOwnershipProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorOwnershipProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorOwnershipProposed represents a OwnershipProposed event raised by the ITellor contract.
type ITellorOwnershipProposed struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipProposed is a free log retrieval operation binding the contract event 0xb51454ce8c7f26becd312a46c4815553887f2ec876a0b8dc813b87f62edf6f80.
//
// Solidity: event OwnershipProposed(address indexed _previousOwner, address indexed _newOwner)
func (_ITellor *ITellorFilterer) FilterOwnershipProposed(opts *bind.FilterOpts, _previousOwner []common.Address, _newOwner []common.Address) (*ITellorOwnershipProposedIterator, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "OwnershipProposed", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ITellorOwnershipProposedIterator{contract: _ITellor.contract, event: "OwnershipProposed", logs: logs, sub: sub}, nil
}

// WatchOwnershipProposed is a free log subscription operation binding the contract event 0xb51454ce8c7f26becd312a46c4815553887f2ec876a0b8dc813b87f62edf6f80.
//
// Solidity: event OwnershipProposed(address indexed _previousOwner, address indexed _newOwner)
func (_ITellor *ITellorFilterer) WatchOwnershipProposed(opts *bind.WatchOpts, sink chan<- *ITellorOwnershipProposed, _previousOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "OwnershipProposed", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorOwnershipProposed)
				if err := _ITellor.contract.UnpackLog(event, "OwnershipProposed", log); err != nil {
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
func (_ITellor *ITellorFilterer) ParseOwnershipProposed(log types.Log) (*ITellorOwnershipProposed, error) {
	event := new(ITellorOwnershipProposed)
	if err := _ITellor.contract.UnpackLog(event, "OwnershipProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ITellor contract.
type ITellorOwnershipTransferredIterator struct {
	Event *ITellorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ITellorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorOwnershipTransferred)
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
		it.Event = new(ITellorOwnershipTransferred)
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
func (it *ITellorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorOwnershipTransferred represents a OwnershipTransferred event raised by the ITellor contract.
type ITellorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner)
func (_ITellor *ITellorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _previousOwner []common.Address, _newOwner []common.Address) (*ITellorOwnershipTransferredIterator, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ITellorOwnershipTransferredIterator{contract: _ITellor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner)
func (_ITellor *ITellorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ITellorOwnershipTransferred, _previousOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorOwnershipTransferred)
				if err := _ITellor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ITellor *ITellorFilterer) ParseOwnershipTransferred(log types.Log) (*ITellorOwnershipTransferred, error) {
	event := new(ITellorOwnershipTransferred)
	if err := _ITellor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorStakeWithdrawRequestedIterator is returned from FilterStakeWithdrawRequested and is used to iterate over the raw logs and unpacked data for StakeWithdrawRequested events raised by the ITellor contract.
type ITellorStakeWithdrawRequestedIterator struct {
	Event *ITellorStakeWithdrawRequested // Event containing the contract specifics and raw log

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
func (it *ITellorStakeWithdrawRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorStakeWithdrawRequested)
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
		it.Event = new(ITellorStakeWithdrawRequested)
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
func (it *ITellorStakeWithdrawRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorStakeWithdrawRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorStakeWithdrawRequested represents a StakeWithdrawRequested event raised by the ITellor contract.
type ITellorStakeWithdrawRequested struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawRequested is a free log retrieval operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: event StakeWithdrawRequested(address indexed _sender)
func (_ITellor *ITellorFilterer) FilterStakeWithdrawRequested(opts *bind.FilterOpts, _sender []common.Address) (*ITellorStakeWithdrawRequestedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "StakeWithdrawRequested", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ITellorStakeWithdrawRequestedIterator{contract: _ITellor.contract, event: "StakeWithdrawRequested", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawRequested is a free log subscription operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: event StakeWithdrawRequested(address indexed _sender)
func (_ITellor *ITellorFilterer) WatchStakeWithdrawRequested(opts *bind.WatchOpts, sink chan<- *ITellorStakeWithdrawRequested, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "StakeWithdrawRequested", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorStakeWithdrawRequested)
				if err := _ITellor.contract.UnpackLog(event, "StakeWithdrawRequested", log); err != nil {
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
func (_ITellor *ITellorFilterer) ParseStakeWithdrawRequested(log types.Log) (*ITellorStakeWithdrawRequested, error) {
	event := new(ITellorStakeWithdrawRequested)
	if err := _ITellor.contract.UnpackLog(event, "StakeWithdrawRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the ITellor contract.
type ITellorStakeWithdrawnIterator struct {
	Event *ITellorStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *ITellorStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorStakeWithdrawn)
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
		it.Event = new(ITellorStakeWithdrawn)
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
func (it *ITellorStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorStakeWithdrawn represents a StakeWithdrawn event raised by the ITellor contract.
type ITellorStakeWithdrawn struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: event StakeWithdrawn(address indexed _sender)
func (_ITellor *ITellorFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, _sender []common.Address) (*ITellorStakeWithdrawnIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "StakeWithdrawn", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ITellorStakeWithdrawnIterator{contract: _ITellor.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: event StakeWithdrawn(address indexed _sender)
func (_ITellor *ITellorFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *ITellorStakeWithdrawn, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "StakeWithdrawn", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorStakeWithdrawn)
				if err := _ITellor.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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
func (_ITellor *ITellorFilterer) ParseStakeWithdrawn(log types.Log) (*ITellorStakeWithdrawn, error) {
	event := new(ITellorStakeWithdrawn)
	if err := _ITellor.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorTipAddedIterator is returned from FilterTipAdded and is used to iterate over the raw logs and unpacked data for TipAdded events raised by the ITellor contract.
type ITellorTipAddedIterator struct {
	Event *ITellorTipAdded // Event containing the contract specifics and raw log

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
func (it *ITellorTipAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorTipAdded)
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
		it.Event = new(ITellorTipAdded)
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
func (it *ITellorTipAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorTipAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorTipAdded represents a TipAdded event raised by the ITellor contract.
type ITellorTipAdded struct {
	Sender    common.Address
	RequestId *big.Int
	Tip       *big.Int
	TotalTips *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTipAdded is a free log retrieval operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_ITellor *ITellorFilterer) FilterTipAdded(opts *bind.FilterOpts, _sender []common.Address, _requestId []*big.Int) (*ITellorTipAddedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ITellorTipAddedIterator{contract: _ITellor.contract, event: "TipAdded", logs: logs, sub: sub}, nil
}

// WatchTipAdded is a free log subscription operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_ITellor *ITellorFilterer) WatchTipAdded(opts *bind.WatchOpts, sink chan<- *ITellorTipAdded, _sender []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorTipAdded)
				if err := _ITellor.contract.UnpackLog(event, "TipAdded", log); err != nil {
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
func (_ITellor *ITellorFilterer) ParseTipAdded(log types.Log) (*ITellorTipAdded, error) {
	event := new(ITellorTipAdded)
	if err := _ITellor.contract.UnpackLog(event, "TipAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorTransferredIterator is returned from FilterTransferred and is used to iterate over the raw logs and unpacked data for Transferred events raised by the ITellor contract.
type ITellorTransferredIterator struct {
	Event *ITellorTransferred // Event containing the contract specifics and raw log

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
func (it *ITellorTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorTransferred)
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
		it.Event = new(ITellorTransferred)
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
func (it *ITellorTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorTransferred represents a Transferred event raised by the ITellor contract.
type ITellorTransferred struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferred is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_ITellor *ITellorFilterer) FilterTransferred(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ITellorTransferredIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ITellorTransferredIterator{contract: _ITellor.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransferred is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_ITellor *ITellorFilterer) WatchTransferred(opts *bind.WatchOpts, sink chan<- *ITellorTransferred, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorTransferred)
				if err := _ITellor.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransferred is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_ITellor *ITellorFilterer) ParseTransferred(log types.Log) (*ITellorTransferred, error) {
	event := new(ITellorTransferred)
	if err := _ITellor.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the ITellor contract.
type ITellorVotedIterator struct {
	Event *ITellorVoted // Event containing the contract specifics and raw log

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
func (it *ITellorVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorVoted)
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
		it.Event = new(ITellorVoted)
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
func (it *ITellorVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorVoted represents a Voted event raised by the ITellor contract.
type ITellorVoted struct {
	DisputeID  *big.Int
	Position   bool
	Voter      common.Address
	VoteWeight *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter, uint256 indexed _voteWeight)
func (_ITellor *ITellorFilterer) FilterVoted(opts *bind.FilterOpts, _disputeID []*big.Int, _voter []common.Address, _voteWeight []*big.Int) (*ITellorVotedIterator, error) {

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

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "Voted", _disputeIDRule, _voterRule, _voteWeightRule)
	if err != nil {
		return nil, err
	}
	return &ITellorVotedIterator{contract: _ITellor.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter, uint256 indexed _voteWeight)
func (_ITellor *ITellorFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *ITellorVoted, _disputeID []*big.Int, _voter []common.Address, _voteWeight []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "Voted", _disputeIDRule, _voterRule, _voteWeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorVoted)
				if err := _ITellor.contract.UnpackLog(event, "Voted", log); err != nil {
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
func (_ITellor *ITellorFilterer) ParseVoted(log types.Log) (*ITellorVoted, error) {
	event := new(ITellorVoted)
	if err := _ITellor.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204191913df7d2d4594cd44658b68eb88fcd059a75971128ad9147c4e5e04de6d164736f6c63430007040033"

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
const TellorABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_currentRequestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"NewChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"NewDispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NewValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_slot\",\"type\":\"uint256\"}],\"name\":\"NonceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"TipAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_position\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_voteWeight\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"fromBlock\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minerIndex\",\"type\":\"uint256\"}],\"name\":\"beginDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bytesVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ext\",\"type\":\"address\"}],\"name\":\"changeExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"changeMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentMiners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"disputesById\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"tally\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disputeVotePassed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPropFork\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"reportedMiner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reportingParty\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposedForkAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_bypass\",\"type\":\"bool\"}],\"name\":\"migrateFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_destination\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amount\",\"type\":\"uint256[]\"}],\"name\":\"migrateForBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_origin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_bypass\",\"type\":\"bool\"}],\"name\":\"migrateFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_origin\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_destination\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amount\",\"type\":\"uint256[]\"}],\"name\":\"migrateFromBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minersByChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"newValueTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"_values\",\"type\":\"uint256[5]\"}],\"name\":\"submitMiningSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"uints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"unlockDisputeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_supportsDispute\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TellorFuncSigs maps the 4-byte function signature to its string representation.
var TellorFuncSigs = map[string]string{
	"024c2ddd": "_allowances(address,address)",
	"752d49a1": "addTip(uint256,uint256)",
	"699f200f": "addresses(bytes32)",
	"dd62ed3e": "allowance(address,address)",
	"999cf26c": "allowedToTrade(address,uint256)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"4ee2cd7e": "balanceOfAt(address,uint256)",
	"cbf1304d": "balances(address,uint256)",
	"8581af19": "beginDispute(uint256,uint256,uint256)",
	"62dd1d2a": "bytesVars(bytes32)",
	"b69a363f": "changeExtension(address)",
	"141e13fa": "changeMigrator(address)",
	"1fd22364": "currentMiners(uint256)",
	"d01f4d9e": "disputeIdByDisputeHash(bytes32)",
	"db085beb": "disputesById(uint256)",
	"8fd3ab80": "migrate()",
	"a9fa7d34": "migrateFor(address,uint256,bool)",
	"42a89bd6": "migrateForBatch(address[],uint256[])",
	"121dd372": "migrateFrom(address,address,uint256,bool)",
	"8c0f4076": "migrateFromBatch(address[],address[],uint256[])",
	"4ba0a5ee": "migrated(address)",
	"48b18e54": "minersByChallenge(bytes32,address)",
	"438c0aa3": "newValueTimestamps(uint256)",
	"5700242c": "requestIdByQueryHash(bytes32)",
	"4350283e": "submitMiningSolution(string,uint256[5],uint256[5])",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"b59e14d4": "uints(bytes32)",
	"9a01ca13": "unlockDisputeFee(uint256)",
	"c9d27afe": "vote(uint256,bool)",
}

// TellorBin is the compiled bytecode used for deploying new contracts.
var TellorBin = "0x608060405234801561001057600080fd5b506150d1806100206000396000f3fe6080604052600436106101d85760003560e01c806370a0823111610102578063a9fa7d3411610095578063cbf1304d11610064578063cbf1304d146109e1578063d01f4d9e14610a49578063db085beb14610a73578063dd62ed3e14610aed576101d8565b8063a9fa7d3414610911578063b59e14d414610952578063b69a363f1461097c578063c9d27afe146109af576101d8565b80638fd3ab80116100d15780638fd3ab8014610860578063999cf26c146108755780639a01ca13146108ae578063a9059cbb146108d8576101d8565b806370a08231146106ac578063752d49a1146106df5780638581af191461070f5780638c0f407614610745576101d8565b80634350283e1161017a5780634ee2cd7e116101495780634ee2cd7e146105d95780635700242c1461061257806362dd1d2a1461063c578063699f200f14610666576101d8565b80634350283e146104c4578063438c0aa31461054357806348b18e541461056d5780634ba0a5ee146105a6576101d8565b8063141e13fa116101b6578063141e13fa146103385780631fd223641461036b57806323b872dd146103b657806342a89bd6146103f9576101d8565b8063024c2ddd14610251578063095ea7b31461029e578063121dd372146102eb575b7f2b2a1c876f73e67ebc4f1b08d10d54d62d62216382e0f4fd16c29155818207a4600090815260476020527ffe0323da4092f31e73ad4b4aa705eaa20d7ce93cdb6c891e7c038c2a7146f008546001600160a01b03169061023882610b28565b5090503d6000803e80801561024c573d6000f35b3d6000fd5b34801561025d57600080fd5b5061028c6004803603604081101561027457600080fd5b506001600160a01b0381358116916020013516610b98565b60408051918252519081900360200190f35b3480156102aa57600080fd5b506102d7600480360360408110156102c157600080fd5b506001600160a01b038135169060200135610bb5565b604080519115158252519081900360200190f35b3480156102f757600080fd5b506103366004803603608081101561030e57600080fd5b506001600160a01b038135811691602081013590911690604081013590606001351515610ca0565b005b34801561034457600080fd5b506103366004803603602081101561035b57600080fd5b50356001600160a01b0316610d23565b34801561037757600080fd5b506103956004803603602081101561038e57600080fd5b5035610e10565b604080519283526001600160a01b0390911660208301528051918290030190f35b3480156103c257600080fd5b506102d7600480360360608110156103d957600080fd5b506001600160a01b03813581169160208101359091169060400135610e3b565b34801561040557600080fd5b506103366004803603604081101561041c57600080fd5b810190602081018135600160201b81111561043657600080fd5b82018360208201111561044857600080fd5b803590602001918460208302840111600160201b8311171561046957600080fd5b919390929091602081019035600160201b81111561048657600080fd5b82018360208201111561049857600080fd5b803590602001918460208302840111600160201b831117156104b957600080fd5b509092509050610ee7565b3480156104d057600080fd5b5061033660048036036101608110156104e857600080fd5b810190602081018135600160201b81111561050257600080fd5b82018360208201111561051457600080fd5b803590602001918460018302840111600160201b8311171561053557600080fd5b919350915060a08101610ff3565b34801561054f57600080fd5b5061028c6004803603602081101561056657600080fd5b5035611199565b34801561057957600080fd5b506102d76004803603604081101561059057600080fd5b50803590602001356001600160a01b03166111ba565b3480156105b257600080fd5b506102d7600480360360208110156105c957600080fd5b50356001600160a01b03166111da565b3480156105e557600080fd5b5061028c600480360360408110156105fc57600080fd5b506001600160a01b0381351690602001356111ef565b34801561061e57600080fd5b5061028c6004803603602081101561063557600080fd5b5035611393565b34801561064857600080fd5b5061028c6004803603602081101561065f57600080fd5b50356113a5565b34801561067257600080fd5b506106906004803603602081101561068957600080fd5b50356113b7565b604080516001600160a01b039092168252519081900360200190f35b3480156106b857600080fd5b5061028c600480360360208110156106cf57600080fd5b50356001600160a01b03166113d2565b3480156106eb57600080fd5b506103366004803603604081101561070257600080fd5b50803590602001356113de565b34801561071b57600080fd5b506103366004803603606081101561073257600080fd5b50803590602081013590604001356115ee565b34801561075157600080fd5b506103366004803603606081101561076857600080fd5b810190602081018135600160201b81111561078257600080fd5b82018360208201111561079457600080fd5b803590602001918460208302840111600160201b831117156107b557600080fd5b919390929091602081019035600160201b8111156107d257600080fd5b8201836020820111156107e457600080fd5b803590602001918460208302840111600160201b8311171561080557600080fd5b919390929091602081019035600160201b81111561082257600080fd5b82018360208201111561083457600080fd5b803590602001918460208302840111600160201b8311171561085557600080fd5b509092509050611e45565b34801561086c57600080fd5b50610336611f7a565b34801561088157600080fd5b506102d76004803603604081101561089857600080fd5b506001600160a01b038135169060200135611f85565b3480156108ba57600080fd5b50610336600480360360208110156108d157600080fd5b5035612048565b3480156108e457600080fd5b506102d7600480360360408110156108fb57600080fd5b506001600160a01b0381351690602001356127c8565b34801561091d57600080fd5b506103366004803603606081101561093457600080fd5b506001600160a01b03813516906020810135906040013515156127de565b34801561095e57600080fd5b5061028c6004803603602081101561097557600080fd5b503561285f565b34801561098857600080fd5b506103366004803603602081101561099f57600080fd5b50356001600160a01b0316612871565b3480156109bb57600080fd5b50610336600480360360408110156109d257600080fd5b50803590602001351515612982565b3480156109ed57600080fd5b50610a1a60048036036040811015610a0457600080fd5b506001600160a01b038135169060200135612bad565b60405180836001600160801b03168152602001826001600160801b031681526020019250505060405180910390f35b348015610a5557600080fd5b5061028c60048036036020811015610a6c57600080fd5b5035612bf0565b348015610a7f57600080fd5b50610a9d60048036036020811015610a9657600080fd5b5035612c02565b60408051988952602089019790975294151587870152921515606087015290151560808601526001600160a01b0390811660a086015290811660c08501521660e083015251908190036101000190f35b348015610af957600080fd5b5061028c60048036036040811015610b1057600080fd5b506001600160a01b0381358116916020013516612c5e565b60006060826001600160a01b03166000366040518083838082843760405192019450600093509091505080830381855af49150503d8060008114610b88576040519150601f19603f3d011682016040523d82523d6000602084013e610b8d565b606091505b509094909350915050565b604a60209081526000928352604080842090915290825290205481565b600033610bf35760405162461bcd60e51b81526004018080602001828103825260248152602001806150336024913960400191505060405180910390fd5b6001600160a01b038316610c385760405162461bcd60e51b8152600401808060200182810382526022815260200180614e676022913960400191505060405180910390fd5b336000818152604a602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060015b92915050565b600080516020614e478339815191526000526047602052600080516020614ff2833981519152546001600160a01b03163314610d11576040805162461bcd60e51b815260206004820152600b60248201526a1b9bdd08185b1b1bddd95960aa1b604482015290519081900360640190fd5b610d1d84848484612c89565b50505050565b7f5fc094d10c65bc33cc842217b2eccca0191ff24148319da094e540a55989896160005260476020527f437dd27c2043efdfef03344e9331c924985f7bd1752abef5ea93bdbfed685100546001600160a01b03163314610dca576040805162461bcd60e51b815260206004820152601b60248201527f6f6e6c792064656974792063616e2063616c6c207468697320666e0000000000604482015290519081900360640190fd5b600080516020614e478339815191526000526047602052600080516020614ff283398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b603a8160058110610e2057600080fd5b6002020180546001909101549091506001600160a01b031682565b6001600160a01b0383166000908152604a60209081526040808320338452909152812054821115610ea8576040805162461bcd60e51b8152602060048201526012602482015271416c6c6f77616e63652069732077726f6e6760701b604482015290519081900360640190fd5b6001600160a01b0384166000908152604a60209081526040808320338452909152902080548390039055610edd848484612d20565b5060019392505050565b600080516020614e478339815191526000526047602052600080516020614ff2833981519152546001600160a01b03163314610f58576040805162461bcd60e51b815260206004820152600b60248201526a1b9bdd08185b1b1bddd95960aa1b604482015290519081900360640190fd5b808314610fa0576040805162461bcd60e51b81526020600482015260116024820152701b5a5cdb585d18da1a5b99c81a5b9c1d5d607a1b604482015290519081900360640190fd5b60005b83811015610fec57610fe4858583818110610fba57fe5b905060200201356001600160a01b0316848484818110610fd657fe5b905060200201356000612ec0565b600101610fa3565b5050505050565b604080513360208083019190915282518083038201815291830183528151918101919091206000818152604690925291902054158061104657506000818152604660205260409020546103844291909103115b6110815760405162461bcd60e51b815260040180806020018281038252602a815260200180614dfd602a913960400191505060405180910390fd5b7fdfbec46864bc123768f0d134913175d9577a55bb71b9b2595fda21e21f36b0826000526046602052600080516020614eaa833981519152546004146111005761110085858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612f5692505050565b6000818152604660209081526040918290204290558151601f8701829004820281018201909252858252610fec91908790879081908401838280828437600092019190915250506040805160a081810190925292508791506005908390839080828437600092019190915250506040805160a081810190925291508690600590839083908082843760009201919091525061321b915050565b603381815481106111a957600080fd5b600091825260209091200154905081565b603960209081526000928352604080842090915290825290205460ff1681565b604b6020526000908152604090205460ff1681565b6001600160a01b0382166000908152604960205260408120805415806112355750828160008154811061121e57fe5b6000918252602090912001546001600160801b0316115b15611244576000915050610c9a565b80548190600019810190811061125657fe5b6000918252602090912001546001600160801b031683106112a85780548190600019810190811061128357fe5b600091825260209091200154600160801b90046001600160801b03169150610c9a9050565b8054600090600119015b818111156113605760006002600183850101049050858482815481106112d457fe5b6000918252602090912001546001600160801b03161415611323578381815481106112fb57fe5b600091825260209091200154600160801b90046001600160801b03169450610c9a9350505050565b8584828154811061133057fe5b6000918252602090912001546001600160801b031610156113535780925061135a565b6001810391505b506112b2565b82828154811061136c57fe5b600091825260209091200154600160801b90046001600160801b03169350610c9a92505050565b60376020526000908152604090205481565b60486020526000908152604090205481565b6047602052600090815260409020546001600160a01b031681565b6000610c9a82436111ef565b81611421576040805162461bcd60e51b815260206004820152600e60248201526d052657175657374496420697320360941b604482015290519081900360640190fd5b80611473576040805162461bcd60e51b815260206004820152601c60248201527f5469702073686f756c642062652067726561746572207468616e203000000000604482015290519081900360640190fd5b7f3f8b5616fa9e7f2ce4a868fde15c58b92e77bc1acd6769bf1567629a3dc4c86560005260466020527f7119b9afaa3bda0901ffe121c1535f50cd6d0d09df5d29eb1cb16c8ab47a55d6546001018281141561151b577f3f8b5616fa9e7f2ce4a868fde15c58b92e77bc1acd6769bf1567629a3dc4c86560005260466020527f7119b9afaa3bda0901ffe121c1535f50cd6d0d09df5d29eb1cb16c8ab47a55d681905561156f565b80831061156f576040805162461bcd60e51b815260206004820181905260248201527f526571756573744964206973206e6f74206c657373207468616e20636f756e74604482015290519081900360640190fd5b61157933836138a3565b61158383836139b7565b600083815260456020908152604080832060008051602061507c8339815191528452600101825291829020548251858152918201528151859233927fd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820929081900390910190a3505050565b60008381526045602090815260408083208584526002810190925290912054611651576040805162461bcd60e51b815260206004820152601060248201526f04d696e656420626c6f636b20697320360841b604482015290519081900360640190fd5b6005821061169d576040805162461bcd60e51b81526020600482015260146024820152734d696e657220696e6465782069732077726f6e6760601b604482015290519081900360640190fd5b600083815260058083016020526040822090849081106116b957fe5b0154604080516bffffffffffffffffffffffff19606084901b1660208083019190915260348201899052605480830189905283518084039091018152607490920183528151918101919091207f1ce2382bc92689b00ba121fa5a411aa976168affdd8ac143a69035dd984b3b6a8054600101908190556000828152603890935292909120546001600160a01b0390931693509180156117965760008281526036602090815260408083207fed92b4c1e0a9e559a31171d487ecbec963526662038ecfa3a71160bd62fb8733845260050190915290208190556117aa565b506000828152603860205260409020819055805b60008181526036602081815260408084207f6ab2b18aafe78fd59c6a4092015bddd9fcacb8170f72b299074f74d76a91a923855260050180835281852080546001019081905586865293835281518084018590528251808203850181529083018352805190840120855290915290912083905581908382146119a057600082815260366020818152604080842081516000198701818501528251808203850181529083018352805190840120855260059081018352818520548086529383528185207f46f7d53798d31923f6952572c6a19ad2d1a8238d26649c2f3493a6d69e425d28865201909152909120544210156118eb576040805162461bcd60e51b815260206004820152601760248201527f4469737075746520697320616c7265616479206f70656e000000000000000000604482015290519081900360640190fd5b60008181526036602052604090206002015460ff161561199e5760008181526036602090815260408083207ff9e1ae10923bfc79f52e309baf8c7699edb821f91ef5b5bd07be29545917b3a68452600501909152902054620151804291909103111561199e576040805162461bcd60e51b815260206004820152601f60248201527f54696d6520666f7220766f74696e6720686176656e277420656c617073656400604482015290519081900360640190fd5b505b60008860021415611a43575060008a81526045602090815260408083207f310199159a20c50879ffb440b45802138b5b162ec9426720e9dd3ee8bbcdb9d78452600190810183529083208054909101908190557f5d9fadfc729fd027e395e5157ef1b53ef9fa4a8f053043c5f159307543e7cc97909252604690527f167af83a0768d27540775cfef6d996eb63f8a61fcdfb26e654c18fb50960e3be5402611a92565b507f675d2171f68d6f5545d54fb9b1fb61a0e6897e6188ca1cd664e7c9530d91ecfc60005260466020527f3e5522f19747f0f285b96ded572ac4128c3a764aea9f44058dc0afc9dda449865481025b85603660008781526020019081526020016000206000018190555060006036600087815260200190815260200160002060020160026101000a81548160ff021916908315150217905550866036600087815260200190815260200160002060020160036101000a8154816001600160a01b0302191690836001600160a01b03160217905550336036600087815260200190815260200160002060030160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060006036600087815260200190815260200160002060040160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060006036600087815260200190815260200160002060020160006101000a81548160ff02191690831515021790555060006036600087815260200190815260200160002060020160016101000a81548160ff021916908315150217905550600060366000878152602001908152602001600020600101819055508a6036600087815260200190815260200160002060050160007f9f47a2659c3d32b749ae717d975e7962959890862423c4318cf86e4ec220291f60001b81526020019081526020016000208190555089603660008781526020019081526020016000206005016000600080516020614f1183398151915260001b8152602001908152602001600020819055508760060160008b81526020019081526020016000208960058110611cb457fe5b015460008681526036602090815260408083207f9147231ab14efb72c38117f68521ddef8de64f092c18c69dbfb602ffc4de7f478452600501909152808220929092557f46f7d53798d31923f6952572c6a19ad2d1a8238d26649c2f3493a6d69e425d2881528181206202a3008502420190557f4b4cefd5ced7569ef0d091282b4bca9c52a034c56471a6061afd1bf307a2de7c81528181204390557f6de96ee4d33a0617f40a846309c8759048857f51b9d59a12d3c3786d4778883d81528181208b90557f1da95f11543c9b03927178e07951795dfc95c7501a9d1cf00e13414ca33bc409815220819055611dab333083612d20565b8860021415611ddd5760008a81526004890160209081526040808320805460ff1916600117905560038b019091528120555b6001600160a01b0387166000818152604460209081526040918290206003905581518d81529081019290925280518d9288927feceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da6492918290030190a35050505050505050505050565b600080516020614e478339815191526000526047602052600080516020614ff2833981519152546001600160a01b03163314611eb6576040805162461bcd60e51b815260206004820152600b60248201526a1b9bdd08185b1b1bddd95960aa1b604482015290519081900360640190fd5b8483148015611ec457508481145b611f09576040805162461bcd60e51b81526020600482015260116024820152701b5a5cdb585d18da1a5b99c81a5b9c1d5d607a1b604482015290519081900360640190fd5b60005b85811015611f7157611f69878783818110611f2357fe5b905060200201356001600160a01b0316868684818110611f3f57fe5b905060200201356001600160a01b0316858585818110611f5b57fe5b905060200201356000612c89565b600101611f0c565b50505050505050565b611f8333613bc5565b565b6001600160a01b03821660009081526044602052604081205415801590611fc457506001600160a01b0383166000908152604460205260409020546005115b15612035577f5d9fadfc729fd027e395e5157ef1b53ef9fa4a8f053043c5f159307543e7cc9760005260466020527f167af83a0768d27540775cfef6d996eb63f8a61fcdfb26e654c18fb50960e3be54829061201f856113d2565b031061202d57506001610c9a565b506000610c9a565b8161203f846113d2565b10159392505050565b600081815260366020818152604080842054845260388252808420548085529282528084207f6ab2b18aafe78fd59c6a4092015bddd9fcacb8170f72b299074f74d76a91a9238552600501808352818520548251808501919091528251808203850181529083018352805190840120855290915290912054806120c85750805b60008281526036602090815260408083208484528184207f6ab2b18aafe78fd59c6a4092015bddd9fcacb8170f72b299074f74d76a91a92385526005820190935292205480612115575060015b7f29169706298d2b6df50a532e958b56426de1465348b93650fca42d456eaec5fc60009081526005840160205260408120541561218c576040805162461bcd60e51b815260206004820152601060248201526f185b1c9958591e481c185a59081bdd5d60821b604482015290519081900360640190fd5b7ff9e1ae10923bfc79f52e309baf8c7699edb821f91ef5b5bd07be29545917b3a6600090815260058401602052604090205462015180429190910311612219576040805162461bcd60e51b815260206004820152601f60248201527f54696d6520666f7220766f74696e6720686176656e277420656c617073656400604482015290519081900360640190fd5b600284810154630100000090046001600160a01b031660009081526044602090815260408083207f29169706298d2b6df50a532e958b56426de1465348b93650fca42d456eaec5fc84526005890190925290912060019081905591850154909161010090910460ff161515141561250f57620151804206420360018201557fa5ae3e2b97d73fb849ea855d27f073b72815b38452d976bd57e4a157827dadd380546000190190557f2b2a1c876f73e67ebc4f1b08d10d54d62d62216382e0f4fd16c29155818207a4600052604760209081527ffe0323da4092f31e73ad4b4aa705eaa20d7ce93cdb6c891e7c038c2a7146f00854604080516004815260248101825292830180516001600160e01b0316630287018760e31b178152905183516001600160a01b039093169392909182918083835b6020831061236c5780518252601f19909201916020918201910161234d565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146123cc576040519150601f19603f3d011682016040523d82523d6000602084013e6123d1565b606091505b5050815460041415905061245b5760058155600285015460038601547f5d9fadfc729fd027e395e5157ef1b53ef9fa4a8f053043c5f159307543e7cc9760005260466020527f167af83a0768d27540775cfef6d996eb63f8a61fcdfb26e654c18fb50960e3be54612456926001600160a01b0363010000009091048116921690612d20565b600081555b60005b838110156125095760408051828603602080830191909152825180830382018152918301835281519181019190912060009081526005890190915220549250826124a6578792505b600083815260366020908152604080832060038101547f1da95f11543c9b03927178e07951795dfc95c7501a9d1cf00e13414ca33bc4098552600582019093529220546125009130916001600160a01b0390911690612d20565b5060010161245e565b50612712565b600181557f9f47a2659c3d32b749ae717d975e7962959890862423c4318cf86e4ec220291f60009081526005860160208181526040808420548452604582528084207f6de96ee4d33a0617f40a846309c8759048857f51b9d59a12d3c3786d4778883d85529290915290912054600214156125dc577f9147231ab14efb72c38117f68521ddef8de64f092c18c69dbfb602ffc4de7f476000908152600587016020908152604080832054600080516020614f11833981519152845281842054845260038501909252909120555b600080516020614f11833981519152600090815260058701602090815260408083205483526004840190915290205460ff1615156001141561264f57600080516020614f1183398151915260009081526005870160209081526040808320548352600484019091529020805460ff191690555b60005b8481101561270f57604080518287036020808301919091528251808303820181529183018352815191810191909120600090815260058a019091522054935083156126a857600084815260366020526040902095505b600286015460008581526036602090815260408083207f1da95f11543c9b03927178e07951795dfc95c7501a9d1cf00e13414ca33bc409845260050190915290205461270791309163010000009091046001600160a01b031690612d20565b600101612652565b50505b7f6de96ee4d33a0617f40a846309c8759048857f51b9d59a12d3c3786d4778883d6000908152600586016020526040902054600214156127be577f9f47a2659c3d32b749ae717d975e7962959890862423c4318cf86e4ec220291f60009081526005860160209081526040808320548352604582528083207f310199159a20c50879ffb440b45802138b5b162ec9426720e9dd3ee8bbcdb9d78452600101909152902080546000190190555b5050505050505050565b60006127d5338484612d20565b50600192915050565b600080516020614e478339815191526000526047602052600080516020614ff2833981519152546001600160a01b0316331461284f576040805162461bcd60e51b815260206004820152600b60248201526a1b9bdd08185b1b1bddd95960aa1b604482015290519081900360640190fd5b61285a838383612ec0565b505050565b60466020526000908152604090205481565b7f5fc094d10c65bc33cc842217b2eccca0191ff24148319da094e540a55989896160005260476020527f437dd27c2043efdfef03344e9331c924985f7bd1752abef5ea93bdbfed685100546001600160a01b03163314612918576040805162461bcd60e51b815260206004820152601b60248201527f6f6e6c792064656974792063616e2063616c6c207468697320666e0000000000604482015290519081900360640190fd5b7f2b2a1c876f73e67ebc4f1b08d10d54d62d62216382e0f4fd16c29155818207a460005260476020527ffe0323da4092f31e73ad4b4aa705eaa20d7ce93cdb6c891e7c038c2a7146f00880546001600160a01b0319166001600160a01b0392909216919091179055565b60008281526036602090815260408083207f4b4cefd5ced7569ef0d091282b4bca9c52a034c56471a6061afd1bf307a2de7c8452600581019092528220549091906129ce9033906111ef565b33600090815260068401602052604090205490915060ff16151560011415612a3d576040805162461bcd60e51b815260206004820152601860248201527f53656e6465722068617320616c726561647920766f7465640000000000000000604482015290519081900360640190fd5b80612a83576040805162461bcd60e51b81526020600482015260116024820152700557365722062616c616e6365206973203607c1b604482015290519081900360640190fd5b3360009081526044602052604090205460031415612ae1576040805162461bcd60e51b81526020600482015260166024820152754d696e657220697320756e646572206469737075746560501b604482015290519081900360640190fd5b3360009081526006830160209081526040808320805460ff191660019081179091557f1da378694063870452ce03b189f48e04c1aa026348e74e6c86e10738514ad2c4845260058601909252909120805490910190558215612b56576001820154612b4c9082613d14565b6001830155612b6b565b6001820154612b659082613d3f565b60018301555b60408051841515815290518291339187917f911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e919081900360200190a450505050565b60496020528160005260406000208181548110612bc957600080fd5b6000918252602090912001546001600160801b038082169350600160801b90910416905082565b60386020526000908152604090205481565b603660205260009081526040902080546001820154600283015460038401546004909401549293919260ff808316936101008404821693620100008104909216926001600160a01b036301000000909304831692918216911688565b6001600160a01b039182166000908152604a6020908152604080832093909416825291909152205490565b80612cef576001600160a01b0384166000908152604b602052604090205460ff1615612cef576040805162461bcd60e51b815260206004820152601060248201526f185b1c9958591e481b5a59dc985d195960821b604482015290519081900360640190fd5b612cf98383613d65565b5050506001600160a01b03166000908152604b60205260409020805460ff19166001179055565b80612d5c5760405162461bcd60e51b8152600401808060200182810382526021815260200180614f516021913960400191505060405180910390fd5b6001600160a01b038216612daf576040805162461bcd60e51b815260206004820152601560248201527452656365697665722069732030206164647265737360581b604482015290519081900360640190fd5b612db98382611f85565b612df45760405162461bcd60e51b8152600401808060200182810382526027815260200180614eca6027913960400191505060405180910390fd5b6000612dff846113d2565b9050612e0d84838303613f3d565b612e16836113d2565b9050808282011015612e63576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b612e6f83838301613f3d565b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a350505050565b80612f26576001600160a01b0383166000908152604b602052604090205460ff1615612f26576040805162461bcd60e51b815260206004820152601060248201526f185b1c9958591e481b5a59dc985d195960821b604482015290519081900360640190fd5b612f308383613d65565b50506001600160a01b03166000908152604b60205260409020805460ff19166001179055565b600080516020614fb2833981519152547fd54702836c9d21d0727ffacc3e39f57c92b5ae0f50177e593bfb5ec66e3de28060005260486020908152600080516020614fd2833981519152546040805180840183815233606081901b93830193909352865160029560039594938993926054909101918401908083835b60208310612ff15780518252601f199092019160209182019101612fd2565b6001836020036101000a038019825116818451168082178552505050505050905001935050505060405160208183030381529060405280519060200120604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b6020831061307c5780518252601f19909201916020918201910161305d565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa1580156130bb573d6000803e3d6000fd5b5050506040515160601b60405160200180826bffffffffffffffffffffffff191681526014019150506040516020818303038152906040526040518082805190602001908083835b602083106131225780518252601f199092019160209182019101613103565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015613161573d6000803e3d6000fd5b5050506040513d602081101561317657600080fd5b50518161317f57fe5b0615806131dd57507f2c8b528fbaf48aaf13162a5a0519a7ad5a612da8ff8783465c17e076660a59f160005260466020527f231bb0dc207f13dd4e565ebc32496c470e35391bd8d3b6649269ee2328e0311854610384429190910310155b6132185760405162461bcd60e51b81526004018080602001828103825260258152602001806150576025913960400191505060405180910390fd5b50565b6040805133602080830182905283518084038201815292840184528251928101929092206000918252604490925291909120546001146132a2576040805162461bcd60e51b815260206004820152601a60248201527f4d696e657220737461747573206973206e6f74207374616b6572000000000000604482015290519081900360640190fd5b603a548351146132ef576040805162461bcd60e51b8152602060048201526013602482015272526571756573742049442069732077726f6e6760681b604482015290519081900360640190fd5b603c5460208401511461333f576040805162461bcd60e51b8152602060048201526013602482015272526571756573742049442069732077726f6e6760681b604482015290519081900360640190fd5b603e5460408401511461338f576040805162461bcd60e51b8152602060048201526013602482015272526571756573742049442069732077726f6e6760681b604482015290519081900360640190fd5b6040546060840151146133df576040805162461bcd60e51b8152602060048201526013602482015272526571756573742049442069732077726f6e6760681b604482015290519081900360640190fd5b60425460808401511461342f576040805162461bcd60e51b8152602060048201526013602482015272526571756573742049442069732077726f6e6760681b604482015290519081900360640190fd5b6000818152604660209081526040808320429055600080516020614fd283398151915254600080516020614eaa833981519152548185526039845282852033865290935292205460ff16156134b55760405162461bcd60e51b81526004018080602001828103825260218152602001806150126021913960400191505060405180910390fd5b60008281526039602090815260408083203384528252808320805460ff191660011790557fe97d205f7d20bf394e3813033d2203b4733acb28b351c8d2a771647ab0d41c3c548352604582528083208751848052600682019093529220836005811061351d57fe5b015560208086015160016000908152600684019092526040909120836005811061354357fe5b0155604080860151600260009081526006840160205291909120836005811061356857fe5b01556060850151600360009081526006830160205260409020836005811061358c57fe5b0155608085015160046000908152600683016020526040902083600581106135b057fe5b01556000808052600580830160205260409091203391849081106135d057fe5b0180546001600160a01b0319166001600160a01b0392909216919091179055600160009081526005828101602052604090912033918490811061360f57fe5b0180546001600160a01b0319166001600160a01b0392909216919091179055600260009081526005828101602052604090912033918490811061364e57fe5b0180546001600160a01b0319166001600160a01b0392909216919091179055600360009081526005828101602052604090912033918490811061368d57fe5b0180546001600160a01b0319166001600160a01b039290921691909117905560046000908152600582810160205260409091203391849081106136cc57fe5b0180546001600160a01b0319166001600160a01b039290921691909117905560018201600414156136ff576136ff614039565b82336001600160a01b03167f9d2e5f03fc65aff196e0f3a8dd924b24099de487e8cffc888921d420ab196e3989898987604051808060200185600560200280838360005b8381101561375b578181015183820152602001613743565b5050505090500184600560200280838360005b8381101561378657818101518382015260200161376e565b50505050905001838152602001828103825286818151815260200191508051906020019080838360005b838110156137c85781810151838201526020016137b0565b50505050905090810190601f1680156137f55780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a3816001016005141561385b5761381b8787614153565b7fdfbec46864bc123768f0d134913175d9577a55bb71b9b2595fda21e21f36b08260009081526046602052600080516020614eaa83398151915255611f71565b7fdfbec46864bc123768f0d134913175d9577a55bb71b9b2595fda21e21f36b0826000526046602052600080516020614eaa8339815191528054600101905550505050505050565b806138ad576139b3565b60006138b8836113d2565b9050808282031115613905576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b600080516020614f728339815191526000526046602052600080516020614e2783398151915254828103811015613977576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b61398384848403613f3d565b5050600080516020614f728339815191526000526046602052600080516020614e27833981519152805482900390555b5050565b600082815260456020908152604080832060008051602061507c833981519152845260018101909252909120546139ee9083614976565b60008051602061507c8339815191526000908152600183016020526040902055603a54831480613a1f5750603c5483145b80613a2b5750603e5483145b80613a37575060405483145b80613a43575060425483145b15613a7957600080516020614f928339815191526000526046602052600080516020614ef183398151915280548301905561285a565b600080516020614f318339815191526000908152600182016020526040902054613b8b57604080516106608101918290526000918291613ad891839060339082845b815481526020019060010190808311613abb57505050505061498c565b60008051602061507c83398151915260009081526001860160205260409020549193509150821080613b08575081155b15613b845760008051602061507c8339815191526000908152600184016020526040812054908260338110613b3957fe5b015560008181526035602090815260408083208054845260458352818420600080516020614f3183398151915285526001908101845282852085905590899055860190915290208190555b505061285a565b600080516020614f31833981519152600090815260018201602052604081205483919060338110613bb857fe5b0180549091019055505050565b6001600160a01b0381166000908152604b602052604090205460ff1615613c26576040805162461bcd60e51b815260206004820152601060248201526f105b1c9958591e481b5a59dc985d195960821b604482015290519081900360640190fd5b7f56e0987db9eaec01ed9e0af003a0fd5c062371f9d23722eb4a3ebc74f16ea371600052604760209081527fc930326aab6c1874fc004d856083a6ed34e057e064970b7effb48e8e6e8ca12754604080516370a0823160e01b81526001600160a01b0380861660048301529151613cf094869493909316926370a082319260248082019391829003018186803b158015613cbf57600080fd5b505afa158015613cd3573d6000803e3d6000fd5b505050506040513d6020811015613ce957600080fd5b5051613d65565b6001600160a01b03166000908152604b60205260409020805460ff19166001179055565b600080821315613d31575081810182811215613d2c57fe5b610c9a565b5081810182811315610c9a57fe5b600080821315613d57575080820382811315613d2c57fe5b5080820382811215610c9a57fe5b80613da15760405162461bcd60e51b8152600401808060200182810382526021815260200180614e896021913960400191505060405180910390fd5b6001600160a01b038216613df4576040805162461bcd60e51b815260206004820152601560248201527452656365697665722069732030206164647265737360581b604482015290519081900360640190fd5b6000613dff836113d2565b9050808282011015613e4c576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b600080516020614f728339815191526000526046602052600080516020614e2783398151915254828101811115613ebe576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b600080516020614f728339815191526000526046602052600080516020614e27833981519152805484019055613ef684838501613f3d565b6040805184815290516001600160a01b038616916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a350505050565b6001600160a01b038216600090815260496020526040902080541580613f8a57508054439082906000198101908110613f7257fe5b6000918252602090912001546001600160801b031614155b15613ffb5760408051808201909152436001600160801b0390811682528381166020808401918252845460018101865560008681529190912093519301805491516fffffffffffffffffffffffffffffffff19909216938316939093178216600160801b919092160217905561285a565b80546000908290600019810190811061401057fe5b600091825260209091200180546001600160801b03808616600160801b02911617905550505050565b7f2c8b528fbaf48aaf13162a5a0519a7ad5a612da8ff8783465c17e076660a59f1600090815260466020527f231bb0dc207f13dd4e565ebc32496c470e35391bd8d3b6649269ee2328e03118544203906140956104b0836149db565b6046602052600080516020614fb2833981519152547fd4f87b8d0f3d3b7e665df74631f6100b2695daa0e30e40eeac02172e15a999e16000527f8156e704072c396780f8253d0562de28216b73a1503daa96e259b9cdd951d71c54610fa092900381029190910591508161410857600191505b61411582820160016149f1565b7ff758978fc1647996a3d9992f611883adc442931dc49488312360acc90601759b6000526046602052600080516020614fb283398151915255505050565b7fe97d205f7d20bf394e3813033d2203b4733acb28b351c8d2a771647ab0d41c3c5460009081526045602090815260408220600080516020614fd2833981519152547f2c8b528fbaf48aaf13162a5a0519a7ad5a612da8ff8783465c17e076660a59f190935260469091527f231bb0dc207f13dd4e565ebc32496c470e35391bd8d3b6649269ee2328e0311880544291829055919291906141f2614d6b565b6141fa614d6b565b60005b600581101561454b5760015b60058110156143b85760008281526006890160205260408120826005811061422d57fe5b015490506000896005016000858152602001908152602001600020836005811061425357fe5b01546001600160a01b03169050825b6000811180156142915750600085815260068c016020526040902060001982016005811061428c57fe5b015483105b1561434857600085815260068c01602052604090206000198201600581106142b557fe5b0154600086815260068d016020526040902082600581106142d257fe5b015560008581526005808d0160205260409091209060001983019081106142f557fe5b015460008681526005808e0160205260409091206001600160a01b0390921691908390811061432057fe5b0180546001600160a01b0319166001600160a01b039290921691909117905560001901614262565b838110156143ad57600085815260068c01602052604090208390826005811061436d57fe5b015560008581526005808d01602052604090912083918390811061438d57fe5b0180546001600160a01b0319166001600160a01b03929092169190911790555b505050600101614209565b506000604560008a84600581106143cb57fe5b60200201518152602001908152602001600020905087600601600083815260200190815260200160002060058060200260405190810160405280929190826005801561442c576020028201915b815481526020019060010190808311614418575b505050505093508360026005811061444057fe5b602090810291909101516000878152600384019092526040918290205584015183836005811061446c57fe5b6020908102919091019190915260008381526005808b01835260408083208984528583019094529091206144a1929091614d89565b5060008281526006808a01602090815260408084208985529285019091529091206144cd916005614d89565b50600082815260058901602052604081206144e791614dc4565b6000828152600689016020526040812061450091614dc4565b80546001818101835560008381526020808220909301889055878152600284018352604080822043905560008051602061507c833981519152825293820190925291812055016141fd565b50847fbeb3b9f46c8d7bb00c873fca10d307538df350866d25f891ffb395147ddbdc4588858460466000600080516020614f9283398151915260001b8152602001908152602001600020546040518085600560200280838360005b838110156145be5781810151838201526020016145a6565b5050505090500184815260200183600560200280838360005b838110156145ef5781810151838201526020016145d7565b5050505090500182815260200194505050505060405180910390a2603380546001810182556000919091527f82a75bdeeae8604d839476ae9efd8b0e15aa447e21bfd7f41283bb54e22c9a8201839055614647614d6b565b87516000908152604560209081526040808320878452600590810190925291829020825160a08101938490529290919082845b81546001600160a01b0316815260019091019060200180831161467a57505050505090506146a88186614a00565b7ff3b93531fa65b3a18680d9ea49df06d96fbd883c4889dc7db866f8b131602dfb60005260466020527fe97d205f7d20bf394e3813033d2203b4733acb28b351c8d2a771647ab0d41c3c80546001019055614701614d6b565b614709614b32565b905060005b600581101561480e5781816005811061472357fe5b6020020151603a826005811061473557fe5b600202015560008060458185856005811061474c57fe5b602002015181526020019081526020016000206001016000600080516020614f3183398151915260001b8152602001908152602001600020546033811061478f57fe5b0155604560008383600581106147a157fe5b602090810291909101518252818101929092526040908101600090812060008051602061507c83398151915282526001908101845291812054600080516020614f928339815191529091526046909252600080516020614ef183398151915280549092019091550161470e565b50898760014303406040516020018080602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b8381101561486257818101518382015260200161484a565b50505050905090810190601f16801561488f5780820380516001836020036101000a031916815260200191505b5060408051601f19818403018152908290528051602091820120600080516020614fd28339815191528190556046909152600080516020614fb283398151915254600080516020614f928339815191526000908152600080516020614ef183398151915254929f508f98507f1d85ce10456e29b67de37887496d3f1fcf1b64c79c4d07484038703a9f5c140897508996509094509092508190859060a0908190849084905b8381101561494c578181015183820152602001614934565b5050505091909101938452505060208201526040805191829003019150a250505050505050505050565b60008282018381101561498557fe5b9392505050565b610640810151603260315b80156149d557828482603381106149aa57fe5b602002015110156149cc578381603381106149c157fe5b602002015192508091505b60001901614997565b50915091565b60008183106149ea5781614985565b5090919050565b60008183136149ea5781614985565b60466020527fc2c579d641b643400780d5c7ce967b420034b9f66962a5ee405cf70e4cbed6bb54600080516020614f928339815191526000908152600080516020614ef183398151915254428490039261012c9084020491600a909104906002830490614a779087905b6020020151838501613d65565b614a82866001614a6a565b614a8d866002614a6a565b614a98866003614a6a565b614aa3866004614a6a565b7f7a39905194de50bde334d18b76bbb36dddd11641d4d50b470cb837cf3bae5def60005260476020527fb5f7e7387e8e977cc9c4c9513388b0d7224264b9a0159cd8e8bdd84a9ed504c354614b01906001600160a01b031682613d65565b5050600080516020614f9283398151915260009081526046602052600080516020614ef18339815191525550505050565b614b3a614d6b565b614b42614d6b565b614b4a614d6b565b60408051610660810191829052614b819160009060339082845b815481526020019060010190808311614b64575050505050614c17565b909250905060005b6005811015614c1157828160058110614b9e57fe5b602002015115614be25760356000838360058110614bb857fe5b6020020151815260200190815260200160002054848260058110614bd857fe5b6020020152614c09565b603a8160040360058110614bf257fe5b6002020154848260058110614c0357fe5b60200201525b600101614b89565b50505090565b614c1f614d6b565b614c27614d6b565b60208301516000805b6005811015614caa57858160010160338110614c4857fe5b6020020151858260058110614c5957fe5b602002015260018101848260058110614c6e57fe5b602002015282858260058110614c8057fe5b60200201511015614ca257848160058110614c9757fe5b602002015192508091505b600101614c30565b5060065b6033811015614d635782868260338110614cc457fe5b60200201511115614d5b57858160338110614cdb57fe5b6020020151858360058110614cec57fe5b602002015280848360058110614cfe57fe5b6020020152858160338110614d0f57fe5b6020020151925060005b6005811015614d595783868260058110614d2f57fe5b60200201511015614d5157858160058110614d4657fe5b602002015193508092505b600101614d19565b505b600101614cae565b505050915091565b6040518060a001604052806005906020820280368337509192915050565b8260058101928215614db4579182015b82811115614db4578254825591600101919060010190614d99565b50614dc0929150614de7565b5090565b506000815560010160008155600101600081556001016000815560010160009055565b5b80821115614dc05760008155600101614de856fe4d696e65722063616e206f6e6c792077696e2072657761726473206f6e636520706572203135206d696efffeead1ec15181fd57b4590d95e0c076bccb59e311315e8b38f23c710aa7c3ec6b005d45c4c789dfe9e2895b51df4336782c5ff6bd59a5c5c9513955aa0630745524332303a20617070726f766520746f20746865207a65726f2061646472657373547269656420746f206d696e74206e6f6e2d706f73697469766520616d6f756e747df1eb1754bc067736ff3d89af41d339bf906d31b0f5978e3c78f402d4ed249253686f756c6420686176652073756666696369656e742062616c616e636520746f20747261646538b16d06a20ab673b01c748aff938df6a38f81640035f4ce8bd9abb03aae5b722f9328a9c75282bec25bb04befad06926366736e0030c985108445fa728335e5f68d680ab3160f1aa5d9c3a1383c49e3e60bf3c0c031245cbb036f5ce99afaa1547269656420746f2073656e64206e6f6e2d706f73697469766520616d6f756e74e6148e7230ca038d456350e69a91b66968b222bfac9ebfbea6ff0a1fb738016009659d32f99e50ac728058418d38174fe83a137c455ff1847e6fb8e15f78f77a5bccd7373734898281f858d7562320d2cdfc0b17bd72f779686937174d15002552cb9007c7c6068f8ef37039d4f232cbf5a28ff8d93a5983c4c0c27cd2f9bc0d59fc4443fcf6ab8c3f349750e3d98e082c2fd08bba13d4ca2b7ef80020cbe79c4d696e657220616c7265616479207375626d6974746564207468652076616c756545524332303a20617070726f76652066726f6d20746865207a65726f2061646472657373496e636f7272656374206e6f6e636520666f722063757272656e74206368616c6c656e67651590276b7f31dd8e2a06f9a92867333eeb3eddbc91e73b9833e3e55d8e34f77da26469706673582212206ede7defbbcf2210142b9488d5d723f3d042c3370820e8eb6544d78e85270d1264736f6c63430007040033"

// DeployTellor deploys a new Ethereum contract, binding an instance of Tellor to it.
func DeployTellor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tellor, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

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

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_Tellor *TellorCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "_allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_Tellor *TellorSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Tellor.Contract.Allowances(&_Tellor.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_Tellor *TellorCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Tellor.Contract.Allowances(&_Tellor.CallOpts, arg0, arg1)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_Tellor *TellorCaller) Addresses(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "addresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_Tellor *TellorSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _Tellor.Contract.Addresses(&_Tellor.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_Tellor *TellorCallerSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _Tellor.Contract.Addresses(&_Tellor.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_Tellor *TellorCaller) Allowance(opts *bind.CallOpts, _user common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "allowance", _user, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_Tellor *TellorSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _Tellor.Contract.Allowance(&_Tellor.CallOpts, _user, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_Tellor *TellorCallerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _Tellor.Contract.Allowance(&_Tellor.CallOpts, _user, _spender)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_Tellor *TellorCaller) AllowedToTrade(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "allowedToTrade", _user, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_Tellor *TellorSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _Tellor.Contract.AllowedToTrade(&_Tellor.CallOpts, _user, _amount)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_Tellor *TellorCallerSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _Tellor.Contract.AllowedToTrade(&_Tellor.CallOpts, _user, _amount)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_Tellor *TellorCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "balanceOf", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_Tellor *TellorSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _Tellor.Contract.BalanceOf(&_Tellor.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_Tellor *TellorCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _Tellor.Contract.BalanceOf(&_Tellor.CallOpts, _user)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_Tellor *TellorCaller) BalanceOfAt(opts *bind.CallOpts, _user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "balanceOfAt", _user, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_Tellor *TellorSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _Tellor.Contract.BalanceOfAt(&_Tellor.CallOpts, _user, _blockNumber)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_Tellor *TellorCallerSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _Tellor.Contract.BalanceOfAt(&_Tellor.CallOpts, _user, _blockNumber)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_Tellor *TellorCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "balances", arg0, arg1)

	outstruct := new(struct {
		FromBlock *big.Int
		Value     *big.Int
	})

	outstruct.FromBlock = out[0].(*big.Int)
	outstruct.Value = out[1].(*big.Int)

	return *outstruct, err

}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_Tellor *TellorSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _Tellor.Contract.Balances(&_Tellor.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_Tellor *TellorCallerSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _Tellor.Contract.Balances(&_Tellor.CallOpts, arg0, arg1)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_Tellor *TellorCaller) BytesVars(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "bytesVars", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_Tellor *TellorSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _Tellor.Contract.BytesVars(&_Tellor.CallOpts, arg0)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_Tellor *TellorCallerSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _Tellor.Contract.BytesVars(&_Tellor.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_Tellor *TellorCaller) CurrentMiners(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "currentMiners", arg0)

	outstruct := new(struct {
		Value *big.Int
		Miner common.Address
	})

	outstruct.Value = out[0].(*big.Int)
	outstruct.Miner = out[1].(common.Address)

	return *outstruct, err

}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_Tellor *TellorSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _Tellor.Contract.CurrentMiners(&_Tellor.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_Tellor *TellorCallerSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _Tellor.Contract.CurrentMiners(&_Tellor.CallOpts, arg0)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_Tellor *TellorCaller) DisputeIdByDisputeHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "disputeIdByDisputeHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_Tellor *TellorSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _Tellor.Contract.DisputeIdByDisputeHash(&_Tellor.CallOpts, arg0)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_Tellor *TellorCallerSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _Tellor.Contract.DisputeIdByDisputeHash(&_Tellor.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_Tellor *TellorCaller) DisputesById(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "disputesById", arg0)

	outstruct := new(struct {
		Hash                [32]byte
		Tally               *big.Int
		Executed            bool
		DisputeVotePassed   bool
		IsPropFork          bool
		ReportedMiner       common.Address
		ReportingParty      common.Address
		ProposedForkAddress common.Address
	})

	outstruct.Hash = out[0].([32]byte)
	outstruct.Tally = out[1].(*big.Int)
	outstruct.Executed = out[2].(bool)
	outstruct.DisputeVotePassed = out[3].(bool)
	outstruct.IsPropFork = out[4].(bool)
	outstruct.ReportedMiner = out[5].(common.Address)
	outstruct.ReportingParty = out[6].(common.Address)
	outstruct.ProposedForkAddress = out[7].(common.Address)

	return *outstruct, err

}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_Tellor *TellorSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _Tellor.Contract.DisputesById(&_Tellor.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_Tellor *TellorCallerSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _Tellor.Contract.DisputesById(&_Tellor.CallOpts, arg0)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_Tellor *TellorCaller) Migrated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "migrated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_Tellor *TellorSession) Migrated(arg0 common.Address) (bool, error) {
	return _Tellor.Contract.Migrated(&_Tellor.CallOpts, arg0)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_Tellor *TellorCallerSession) Migrated(arg0 common.Address) (bool, error) {
	return _Tellor.Contract.Migrated(&_Tellor.CallOpts, arg0)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_Tellor *TellorCaller) MinersByChallenge(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "minersByChallenge", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_Tellor *TellorSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Tellor.Contract.MinersByChallenge(&_Tellor.CallOpts, arg0, arg1)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_Tellor *TellorCallerSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Tellor.Contract.MinersByChallenge(&_Tellor.CallOpts, arg0, arg1)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_Tellor *TellorCaller) NewValueTimestamps(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "newValueTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_Tellor *TellorSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _Tellor.Contract.NewValueTimestamps(&_Tellor.CallOpts, arg0)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_Tellor *TellorCallerSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _Tellor.Contract.NewValueTimestamps(&_Tellor.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_Tellor *TellorCaller) RequestIdByQueryHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "requestIdByQueryHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_Tellor *TellorSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _Tellor.Contract.RequestIdByQueryHash(&_Tellor.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_Tellor *TellorCallerSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _Tellor.Contract.RequestIdByQueryHash(&_Tellor.CallOpts, arg0)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_Tellor *TellorCaller) Uints(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "uints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_Tellor *TellorSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _Tellor.Contract.Uints(&_Tellor.CallOpts, arg0)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_Tellor *TellorCallerSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _Tellor.Contract.Uints(&_Tellor.CallOpts, arg0)
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

// ChangeExtension is a paid mutator transaction binding the contract method 0xb69a363f.
//
// Solidity: function changeExtension(address _ext) returns()
func (_Tellor *TellorTransactor) ChangeExtension(opts *bind.TransactOpts, _ext common.Address) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "changeExtension", _ext)
}

// ChangeExtension is a paid mutator transaction binding the contract method 0xb69a363f.
//
// Solidity: function changeExtension(address _ext) returns()
func (_Tellor *TellorSession) ChangeExtension(_ext common.Address) (*types.Transaction, error) {
	return _Tellor.Contract.ChangeExtension(&_Tellor.TransactOpts, _ext)
}

// ChangeExtension is a paid mutator transaction binding the contract method 0xb69a363f.
//
// Solidity: function changeExtension(address _ext) returns()
func (_Tellor *TellorTransactorSession) ChangeExtension(_ext common.Address) (*types.Transaction, error) {
	return _Tellor.Contract.ChangeExtension(&_Tellor.TransactOpts, _ext)
}

// ChangeMigrator is a paid mutator transaction binding the contract method 0x141e13fa.
//
// Solidity: function changeMigrator(address _migrator) returns()
func (_Tellor *TellorTransactor) ChangeMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "changeMigrator", _migrator)
}

// ChangeMigrator is a paid mutator transaction binding the contract method 0x141e13fa.
//
// Solidity: function changeMigrator(address _migrator) returns()
func (_Tellor *TellorSession) ChangeMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _Tellor.Contract.ChangeMigrator(&_Tellor.TransactOpts, _migrator)
}

// ChangeMigrator is a paid mutator transaction binding the contract method 0x141e13fa.
//
// Solidity: function changeMigrator(address _migrator) returns()
func (_Tellor *TellorTransactorSession) ChangeMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _Tellor.Contract.ChangeMigrator(&_Tellor.TransactOpts, _migrator)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_Tellor *TellorTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_Tellor *TellorSession) Migrate() (*types.Transaction, error) {
	return _Tellor.Contract.Migrate(&_Tellor.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_Tellor *TellorTransactorSession) Migrate() (*types.Transaction, error) {
	return _Tellor.Contract.Migrate(&_Tellor.TransactOpts)
}

// MigrateFor is a paid mutator transaction binding the contract method 0xa9fa7d34.
//
// Solidity: function migrateFor(address _destination, uint256 _amount, bool _bypass) returns()
func (_Tellor *TellorTransactor) MigrateFor(opts *bind.TransactOpts, _destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "migrateFor", _destination, _amount, _bypass)
}

// MigrateFor is a paid mutator transaction binding the contract method 0xa9fa7d34.
//
// Solidity: function migrateFor(address _destination, uint256 _amount, bool _bypass) returns()
func (_Tellor *TellorSession) MigrateFor(_destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _Tellor.Contract.MigrateFor(&_Tellor.TransactOpts, _destination, _amount, _bypass)
}

// MigrateFor is a paid mutator transaction binding the contract method 0xa9fa7d34.
//
// Solidity: function migrateFor(address _destination, uint256 _amount, bool _bypass) returns()
func (_Tellor *TellorTransactorSession) MigrateFor(_destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _Tellor.Contract.MigrateFor(&_Tellor.TransactOpts, _destination, _amount, _bypass)
}

// MigrateForBatch is a paid mutator transaction binding the contract method 0x42a89bd6.
//
// Solidity: function migrateForBatch(address[] _destination, uint256[] _amount) returns()
func (_Tellor *TellorTransactor) MigrateForBatch(opts *bind.TransactOpts, _destination []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "migrateForBatch", _destination, _amount)
}

// MigrateForBatch is a paid mutator transaction binding the contract method 0x42a89bd6.
//
// Solidity: function migrateForBatch(address[] _destination, uint256[] _amount) returns()
func (_Tellor *TellorSession) MigrateForBatch(_destination []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.MigrateForBatch(&_Tellor.TransactOpts, _destination, _amount)
}

// MigrateForBatch is a paid mutator transaction binding the contract method 0x42a89bd6.
//
// Solidity: function migrateForBatch(address[] _destination, uint256[] _amount) returns()
func (_Tellor *TellorTransactorSession) MigrateForBatch(_destination []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.MigrateForBatch(&_Tellor.TransactOpts, _destination, _amount)
}

// MigrateFrom is a paid mutator transaction binding the contract method 0x121dd372.
//
// Solidity: function migrateFrom(address _origin, address _destination, uint256 _amount, bool _bypass) returns()
func (_Tellor *TellorTransactor) MigrateFrom(opts *bind.TransactOpts, _origin common.Address, _destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "migrateFrom", _origin, _destination, _amount, _bypass)
}

// MigrateFrom is a paid mutator transaction binding the contract method 0x121dd372.
//
// Solidity: function migrateFrom(address _origin, address _destination, uint256 _amount, bool _bypass) returns()
func (_Tellor *TellorSession) MigrateFrom(_origin common.Address, _destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _Tellor.Contract.MigrateFrom(&_Tellor.TransactOpts, _origin, _destination, _amount, _bypass)
}

// MigrateFrom is a paid mutator transaction binding the contract method 0x121dd372.
//
// Solidity: function migrateFrom(address _origin, address _destination, uint256 _amount, bool _bypass) returns()
func (_Tellor *TellorTransactorSession) MigrateFrom(_origin common.Address, _destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _Tellor.Contract.MigrateFrom(&_Tellor.TransactOpts, _origin, _destination, _amount, _bypass)
}

// MigrateFromBatch is a paid mutator transaction binding the contract method 0x8c0f4076.
//
// Solidity: function migrateFromBatch(address[] _origin, address[] _destination, uint256[] _amount) returns()
func (_Tellor *TellorTransactor) MigrateFromBatch(opts *bind.TransactOpts, _origin []common.Address, _destination []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "migrateFromBatch", _origin, _destination, _amount)
}

// MigrateFromBatch is a paid mutator transaction binding the contract method 0x8c0f4076.
//
// Solidity: function migrateFromBatch(address[] _origin, address[] _destination, uint256[] _amount) returns()
func (_Tellor *TellorSession) MigrateFromBatch(_origin []common.Address, _destination []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.MigrateFromBatch(&_Tellor.TransactOpts, _origin, _destination, _amount)
}

// MigrateFromBatch is a paid mutator transaction binding the contract method 0x8c0f4076.
//
// Solidity: function migrateFromBatch(address[] _origin, address[] _destination, uint256[] _amount) returns()
func (_Tellor *TellorTransactorSession) MigrateFromBatch(_origin []common.Address, _destination []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.MigrateFromBatch(&_Tellor.TransactOpts, _origin, _destination, _amount)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestIds, uint256[5] _values) returns()
func (_Tellor *TellorTransactor) SubmitMiningSolution(opts *bind.TransactOpts, _nonce string, _requestIds [5]*big.Int, _values [5]*big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "submitMiningSolution", _nonce, _requestIds, _values)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestIds, uint256[5] _values) returns()
func (_Tellor *TellorSession) SubmitMiningSolution(_nonce string, _requestIds [5]*big.Int, _values [5]*big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.SubmitMiningSolution(&_Tellor.TransactOpts, _nonce, _requestIds, _values)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestIds, uint256[5] _values) returns()
func (_Tellor *TellorTransactorSession) SubmitMiningSolution(_nonce string, _requestIds [5]*big.Int, _values [5]*big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.SubmitMiningSolution(&_Tellor.TransactOpts, _nonce, _requestIds, _values)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_Tellor *TellorTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_Tellor *TellorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.Transfer(&_Tellor.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_Tellor *TellorTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.Transfer(&_Tellor.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_Tellor *TellorTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_Tellor *TellorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.TransferFrom(&_Tellor.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
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

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Tellor *TellorTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Tellor.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Tellor *TellorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Tellor.Contract.Fallback(&_Tellor.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Tellor *TellorTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Tellor.Contract.Fallback(&_Tellor.TransactOpts, calldata)
}

// TellorApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Tellor contract.
type TellorApprovalIterator struct {
	Event *TellorApproval // Event containing the contract specifics and raw log

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
func (it *TellorApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorApproval)
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
		it.Event = new(TellorApproval)
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
func (it *TellorApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorApproval represents a Approval event raised by the Tellor contract.
type TellorApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tellor *TellorFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TellorApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Tellor.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TellorApprovalIterator{contract: _Tellor.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tellor *TellorFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TellorApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Tellor.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorApproval)
				if err := _Tellor.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tellor *TellorFilterer) ParseApproval(log types.Log) (*TellorApproval, error) {
	event := new(TellorApproval)
	if err := _Tellor.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorNewChallengeIterator is returned from FilterNewChallenge and is used to iterate over the raw logs and unpacked data for NewChallenge events raised by the Tellor contract.
type TellorNewChallengeIterator struct {
	Event *TellorNewChallenge // Event containing the contract specifics and raw log

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
func (it *TellorNewChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorNewChallenge)
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
		it.Event = new(TellorNewChallenge)
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
func (it *TellorNewChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorNewChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorNewChallenge represents a NewChallenge event raised by the Tellor contract.
type TellorNewChallenge struct {
	CurrentChallenge [32]byte
	CurrentRequestId [5]*big.Int
	Difficulty       *big.Int
	TotalTips        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewChallenge is a free log retrieval operation binding the contract event 0x1d85ce10456e29b67de37887496d3f1fcf1b64c79c4d07484038703a9f5c1408.
//
// Solidity: event NewChallenge(bytes32 indexed _currentChallenge, uint256[5] _currentRequestId, uint256 _difficulty, uint256 _totalTips)
func (_Tellor *TellorFilterer) FilterNewChallenge(opts *bind.FilterOpts, _currentChallenge [][32]byte) (*TellorNewChallengeIterator, error) {

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _Tellor.contract.FilterLogs(opts, "NewChallenge", _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return &TellorNewChallengeIterator{contract: _Tellor.contract, event: "NewChallenge", logs: logs, sub: sub}, nil
}

// WatchNewChallenge is a free log subscription operation binding the contract event 0x1d85ce10456e29b67de37887496d3f1fcf1b64c79c4d07484038703a9f5c1408.
//
// Solidity: event NewChallenge(bytes32 indexed _currentChallenge, uint256[5] _currentRequestId, uint256 _difficulty, uint256 _totalTips)
func (_Tellor *TellorFilterer) WatchNewChallenge(opts *bind.WatchOpts, sink chan<- *TellorNewChallenge, _currentChallenge [][32]byte) (event.Subscription, error) {

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _Tellor.contract.WatchLogs(opts, "NewChallenge", _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorNewChallenge)
				if err := _Tellor.contract.UnpackLog(event, "NewChallenge", log); err != nil {
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
func (_Tellor *TellorFilterer) ParseNewChallenge(log types.Log) (*TellorNewChallenge, error) {
	event := new(TellorNewChallenge)
	if err := _Tellor.contract.UnpackLog(event, "NewChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorNewDisputeIterator is returned from FilterNewDispute and is used to iterate over the raw logs and unpacked data for NewDispute events raised by the Tellor contract.
type TellorNewDisputeIterator struct {
	Event *TellorNewDispute // Event containing the contract specifics and raw log

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
func (it *TellorNewDisputeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorNewDispute)
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
		it.Event = new(TellorNewDispute)
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
func (it *TellorNewDisputeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorNewDisputeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorNewDispute represents a NewDispute event raised by the Tellor contract.
type TellorNewDispute struct {
	DisputeId *big.Int
	RequestId *big.Int
	Timestamp *big.Int
	Miner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewDispute is a free log retrieval operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_Tellor *TellorFilterer) FilterNewDispute(opts *bind.FilterOpts, _disputeId []*big.Int, _requestId []*big.Int) (*TellorNewDisputeIterator, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Tellor.contract.FilterLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &TellorNewDisputeIterator{contract: _Tellor.contract, event: "NewDispute", logs: logs, sub: sub}, nil
}

// WatchNewDispute is a free log subscription operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_Tellor *TellorFilterer) WatchNewDispute(opts *bind.WatchOpts, sink chan<- *TellorNewDispute, _disputeId []*big.Int, _requestId []*big.Int) (event.Subscription, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Tellor.contract.WatchLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorNewDispute)
				if err := _Tellor.contract.UnpackLog(event, "NewDispute", log); err != nil {
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
func (_Tellor *TellorFilterer) ParseNewDispute(log types.Log) (*TellorNewDispute, error) {
	event := new(TellorNewDispute)
	if err := _Tellor.contract.UnpackLog(event, "NewDispute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorNewValueIterator is returned from FilterNewValue and is used to iterate over the raw logs and unpacked data for NewValue events raised by the Tellor contract.
type TellorNewValueIterator struct {
	Event *TellorNewValue // Event containing the contract specifics and raw log

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
func (it *TellorNewValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorNewValue)
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
		it.Event = new(TellorNewValue)
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
func (it *TellorNewValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorNewValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorNewValue represents a NewValue event raised by the Tellor contract.
type TellorNewValue struct {
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
func (_Tellor *TellorFilterer) FilterNewValue(opts *bind.FilterOpts, _currentChallenge [][32]byte) (*TellorNewValueIterator, error) {

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _Tellor.contract.FilterLogs(opts, "NewValue", _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return &TellorNewValueIterator{contract: _Tellor.contract, event: "NewValue", logs: logs, sub: sub}, nil
}

// WatchNewValue is a free log subscription operation binding the contract event 0xbeb3b9f46c8d7bb00c873fca10d307538df350866d25f891ffb395147ddbdc45.
//
// Solidity: event NewValue(uint256[5] _requestId, uint256 _time, uint256[5] _value, uint256 _totalTips, bytes32 indexed _currentChallenge)
func (_Tellor *TellorFilterer) WatchNewValue(opts *bind.WatchOpts, sink chan<- *TellorNewValue, _currentChallenge [][32]byte) (event.Subscription, error) {

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _Tellor.contract.WatchLogs(opts, "NewValue", _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorNewValue)
				if err := _Tellor.contract.UnpackLog(event, "NewValue", log); err != nil {
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
func (_Tellor *TellorFilterer) ParseNewValue(log types.Log) (*TellorNewValue, error) {
	event := new(TellorNewValue)
	if err := _Tellor.contract.UnpackLog(event, "NewValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorNonceSubmittedIterator is returned from FilterNonceSubmitted and is used to iterate over the raw logs and unpacked data for NonceSubmitted events raised by the Tellor contract.
type TellorNonceSubmittedIterator struct {
	Event *TellorNonceSubmitted // Event containing the contract specifics and raw log

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
func (it *TellorNonceSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorNonceSubmitted)
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
		it.Event = new(TellorNonceSubmitted)
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
func (it *TellorNonceSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorNonceSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorNonceSubmitted represents a NonceSubmitted event raised by the Tellor contract.
type TellorNonceSubmitted struct {
	Miner            common.Address
	Nonce            string
	RequestId        [5]*big.Int
	Value            [5]*big.Int
	CurrentChallenge [32]byte
	Slot             *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNonceSubmitted is a free log retrieval operation binding the contract event 0x9d2e5f03fc65aff196e0f3a8dd924b24099de487e8cffc888921d420ab196e39.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256[5] _requestId, uint256[5] _value, bytes32 indexed _currentChallenge, uint256 _slot)
func (_Tellor *TellorFilterer) FilterNonceSubmitted(opts *bind.FilterOpts, _miner []common.Address, _currentChallenge [][32]byte) (*TellorNonceSubmittedIterator, error) {

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _Tellor.contract.FilterLogs(opts, "NonceSubmitted", _minerRule, _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return &TellorNonceSubmittedIterator{contract: _Tellor.contract, event: "NonceSubmitted", logs: logs, sub: sub}, nil
}

// WatchNonceSubmitted is a free log subscription operation binding the contract event 0x9d2e5f03fc65aff196e0f3a8dd924b24099de487e8cffc888921d420ab196e39.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256[5] _requestId, uint256[5] _value, bytes32 indexed _currentChallenge, uint256 _slot)
func (_Tellor *TellorFilterer) WatchNonceSubmitted(opts *bind.WatchOpts, sink chan<- *TellorNonceSubmitted, _miner []common.Address, _currentChallenge [][32]byte) (event.Subscription, error) {

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _Tellor.contract.WatchLogs(opts, "NonceSubmitted", _minerRule, _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorNonceSubmitted)
				if err := _Tellor.contract.UnpackLog(event, "NonceSubmitted", log); err != nil {
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

// ParseNonceSubmitted is a log parse operation binding the contract event 0x9d2e5f03fc65aff196e0f3a8dd924b24099de487e8cffc888921d420ab196e39.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256[5] _requestId, uint256[5] _value, bytes32 indexed _currentChallenge, uint256 _slot)
func (_Tellor *TellorFilterer) ParseNonceSubmitted(log types.Log) (*TellorNonceSubmitted, error) {
	event := new(TellorNonceSubmitted)
	if err := _Tellor.contract.UnpackLog(event, "NonceSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorTipAddedIterator is returned from FilterTipAdded and is used to iterate over the raw logs and unpacked data for TipAdded events raised by the Tellor contract.
type TellorTipAddedIterator struct {
	Event *TellorTipAdded // Event containing the contract specifics and raw log

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
func (it *TellorTipAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorTipAdded)
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
		it.Event = new(TellorTipAdded)
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
func (it *TellorTipAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorTipAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorTipAdded represents a TipAdded event raised by the Tellor contract.
type TellorTipAdded struct {
	Sender    common.Address
	RequestId *big.Int
	Tip       *big.Int
	TotalTips *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTipAdded is a free log retrieval operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_Tellor *TellorFilterer) FilterTipAdded(opts *bind.FilterOpts, _sender []common.Address, _requestId []*big.Int) (*TellorTipAddedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Tellor.contract.FilterLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &TellorTipAddedIterator{contract: _Tellor.contract, event: "TipAdded", logs: logs, sub: sub}, nil
}

// WatchTipAdded is a free log subscription operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_Tellor *TellorFilterer) WatchTipAdded(opts *bind.WatchOpts, sink chan<- *TellorTipAdded, _sender []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Tellor.contract.WatchLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorTipAdded)
				if err := _Tellor.contract.UnpackLog(event, "TipAdded", log); err != nil {
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
func (_Tellor *TellorFilterer) ParseTipAdded(log types.Log) (*TellorTipAdded, error) {
	event := new(TellorTipAdded)
	if err := _Tellor.contract.UnpackLog(event, "TipAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorTransferredIterator is returned from FilterTransferred and is used to iterate over the raw logs and unpacked data for Transferred events raised by the Tellor contract.
type TellorTransferredIterator struct {
	Event *TellorTransferred // Event containing the contract specifics and raw log

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
func (it *TellorTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorTransferred)
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
		it.Event = new(TellorTransferred)
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
func (it *TellorTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorTransferred represents a Transferred event raised by the Tellor contract.
type TellorTransferred struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferred is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tellor *TellorFilterer) FilterTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TellorTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Tellor.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TellorTransferredIterator{contract: _Tellor.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransferred is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tellor *TellorFilterer) WatchTransferred(opts *bind.WatchOpts, sink chan<- *TellorTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Tellor.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorTransferred)
				if err := _Tellor.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransferred is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tellor *TellorFilterer) ParseTransferred(log types.Log) (*TellorTransferred, error) {
	event := new(TellorTransferred)
	if err := _Tellor.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the Tellor contract.
type TellorVotedIterator struct {
	Event *TellorVoted // Event containing the contract specifics and raw log

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
func (it *TellorVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorVoted)
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
		it.Event = new(TellorVoted)
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
func (it *TellorVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorVoted represents a Voted event raised by the Tellor contract.
type TellorVoted struct {
	DisputeID  *big.Int
	Position   bool
	Voter      common.Address
	VoteWeight *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter, uint256 indexed _voteWeight)
func (_Tellor *TellorFilterer) FilterVoted(opts *bind.FilterOpts, _disputeID []*big.Int, _voter []common.Address, _voteWeight []*big.Int) (*TellorVotedIterator, error) {

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

	logs, sub, err := _Tellor.contract.FilterLogs(opts, "Voted", _disputeIDRule, _voterRule, _voteWeightRule)
	if err != nil {
		return nil, err
	}
	return &TellorVotedIterator{contract: _Tellor.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter, uint256 indexed _voteWeight)
func (_Tellor *TellorFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *TellorVoted, _disputeID []*big.Int, _voter []common.Address, _voteWeight []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Tellor.contract.WatchLogs(opts, "Voted", _disputeIDRule, _voterRule, _voteWeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorVoted)
				if err := _Tellor.contract.UnpackLog(event, "Voted", log); err != nil {
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
func (_Tellor *TellorFilterer) ParseVoted(log types.Log) (*TellorVoted, error) {
	event := new(TellorVoted)
	if err := _Tellor.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorGettersABI is the input ABI used to generate the binding from.
const TellorGettersABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"fromBlock\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bytesVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentMiners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"disputesById\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"tally\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disputeVotePassed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPropFork\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"reportedMiner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reportingParty\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposedForkAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[9]\",\"name\":\"\",\"type\":\"uint256[9]\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastNewValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"data\",\"type\":\"uint256[51]\"}],\"name\":\"getMax5\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"max\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"maxIndex\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"internalType\":\"address[5]\",\"name\":\"\",\"type\":\"address[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256\",\"name\":\"_diff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"idsOnDeck\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"tipsOnDeck\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"\",\"type\":\"uint256[51]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTopRequestIDs\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minersByChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"newValueTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"uints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TellorGettersFuncSigs maps the 4-byte function signature to its string representation.
var TellorGettersFuncSigs = map[string]string{
	"024c2ddd": "_allowances(address,address)",
	"699f200f": "addresses(bytes32)",
	"cbf1304d": "balances(address,uint256)",
	"62dd1d2a": "bytesVars(bytes32)",
	"1fd22364": "currentMiners(uint256)",
	"313ce567": "decimals()",
	"63bb82ad": "didMine(bytes32,address)",
	"a7c438bc": "didVote(uint256,address)",
	"d01f4d9e": "disputeIdByDisputeHash(bytes32)",
	"db085beb": "disputesById(uint256)",
	"133bee5e": "getAddressVars(bytes32)",
	"af0b1327": "getAllDisputeVars(uint256)",
	"da379941": "getDisputeIdByDisputeHash(bytes32)",
	"7f6fd5d9": "getDisputeUintVars(uint256,bytes32)",
	"fc7cf0a0": "getLastNewValue()",
	"3180f8df": "getLastNewValueById(uint256)",
	"99830e32": "getMax5(uint256[51])",
	"c775b542": "getMinedBlockNum(uint256,uint256)",
	"69026d63": "getMinersByRequestIdAndTimestamp(uint256,uint256)",
	"4049f198": "getNewCurrentVariables()",
	"46eee1c4": "getNewValueCountbyRequestId(uint256)",
	"9a7077ab": "getNewVariablesOnDeck()",
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
	"3df0777b": "isInDispute(uint256,uint256)",
	"4ba0a5ee": "migrated(address)",
	"48b18e54": "minersByChallenge(bytes32,address)",
	"06fdde03": "name()",
	"438c0aa3": "newValueTimestamps(uint256)",
	"5700242c": "requestIdByQueryHash(bytes32)",
	"93fa4915": "retrieveData(uint256,uint256)",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"b59e14d4": "uints(bytes32)",
}

// TellorGettersBin is the compiled bytecode used for deploying new contracts.
var TellorGettersBin = "0x608060405234801561001057600080fd5b50611757806100206000396000f3fe608060405234801561001057600080fd5b506004361061025e5760003560e01c8063699f200f11610146578063b5413029116100c3578063da37994111610087578063da37994114610986578063db085beb146109a3578063e0ae93c114610a10578063e1eee6d614610a33578063fc7cf0a014610a50578063fe1cd15d14610a585761025e565b8063b5413029146108b0578063b59e14d4146108ce578063c775b542146108eb578063cbf1304d1461090e578063d01f4d9e146109695761025e565b806395d89b411161010a57806395d89b411461070b57806399830e32146107135780639a7077ab146107cb578063a7c438bc146107d3578063af0b1327146107ff5761025e565b8063699f200f14610646578063733bdef01461066357806377fbb663146106a25780637f6fd5d9146106c557806393fa4915146106e85761025e565b80634049f198116101df5780635700242c116101a35780635700242c14610583578063612c8f7f146105a05780636173c0b8146105bd57806362dd1d2a146105da57806363bb82ad146105f757806369026d63146106235761025e565b80634049f198146104a2578063438c0aa3146104f757806346eee1c41461051457806348b18e54146105315780634ba0a5ee1461055d5761025e565b806318160ddd1161022657806318160ddd146103d15780631fd22364146103d9578063313ce567146104175780633180f8df146104355780633df0777b1461046b5761025e565b8063024c2ddd1461026357806306fdde03146102a35780630f0b424d1461032057806311c985121461033d578063133bee5e14610398575b600080fd5b6102916004803603604081101561027957600080fd5b506001600160a01b0381358116916020013516610a60565b60408051918252519081900360200190f35b6102ab610a7d565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102e55781810151838201526020016102cd565b50505050905090810190601f1680156103125780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102916004803603602081101561033657600080fd5b5035610aa6565b6103606004803603604081101561035357600080fd5b5080359060200135610ab8565b604051808260a080838360005b8381101561038557818101518382015260200161036d565b5050505090500191505060405180910390f35b6103b5600480360360208110156103ae57600080fd5b5035610b0f565b604080516001600160a01b039092168252519081900360200190f35b610291610b2a565b6103f6600480360360208110156103ef57600080fd5b5035610b78565b604080519283526001600160a01b0390911660208301528051918290030190f35b61041f610ba3565b6040805160ff9092168252519081900360200190f35b6104526004803603602081101561044b57600080fd5b5035610ba8565b6040805192835290151560208301528051918290030190f35b61048e6004803603604081101561048157600080fd5b5080359060200135610c02565b604080519115158252519081900360200190f35b6104aa610c26565b604051848152602081018460a080838360005b838110156104d55781810151838201526020016104bd565b5050505090500183815260200182815260200194505050505060405180910390f35b6102916004803603602081101561050d57600080fd5b5035610d06565b6102916004803603602081101561052a57600080fd5b5035610d27565b61048e6004803603604081101561054757600080fd5b50803590602001356001600160a01b0316610d39565b61048e6004803603602081101561057357600080fd5b50356001600160a01b0316610d59565b6102916004803603602081101561059957600080fd5b5035610d6e565b610291600480360360208110156105b657600080fd5b5035610d80565b610291600480360360208110156105d357600080fd5b5035610d92565b610291600480360360208110156105f057600080fd5b5035610dfd565b61048e6004803603604081101561060d57600080fd5b50803590602001356001600160a01b0316610e0f565b6103606004803603604081101561063957600080fd5b5080359060200135610e3a565b6103b56004803603602081101561065c57600080fd5b5035610e9d565b6106896004803603602081101561067957600080fd5b50356001600160a01b0316610eb8565b6040805192835260208301919091528051918290030190f35b610291600480360360408110156106b857600080fd5b5080359060200135610edb565b610291600480360360408110156106db57600080fd5b5080359060200135610f08565b610291600480360360408110156106fe57600080fd5b5080359060200135610f29565b6102ab610f4a565b610767600480360361066081101561072a57600080fd5b810190808061066001906033806020026040519081016040528092919082603360200280828437600092019190915250919450610f679350505050565b604051808360a080838360005b8381101561078c578181015183820152602001610774565b5050505090500182600560200280838360005b838110156107b757818101518382015260200161079f565b505050509050019250505060405180910390f35b6107676110bb565b61048e600480360360408110156107e957600080fd5b50803590602001356001600160a01b031661115a565b61081c6004803603602081101561081557600080fd5b5035611189565b604051808a8152602001891515815260200188151581526020018715158152602001866001600160a01b03168152602001856001600160a01b03168152602001846001600160a01b0316815260200183600960200280838360005b8381101561088f578181015183820152602001610877565b50505050905001828152602001995050505050505050505060405180910390f35b6108b86113b4565b604051815181528082610660808383602061036d565b610291600480360360208110156108e457600080fd5b50356113f0565b6102916004803603604081101561090157600080fd5b5080359060200135611402565b61093a6004803603604081101561092457600080fd5b506001600160a01b038135169060200135611423565b60405180836001600160801b03168152602001826001600160801b031681526020019250505060405180910390f35b6102916004803603602081101561097f57600080fd5b5035611466565b6102916004803603602081101561099c57600080fd5b5035611478565b6109c0600480360360208110156109b957600080fd5b503561148a565b60408051988952602089019790975294151587870152921515606087015290151560808601526001600160a01b0390811660a086015290811660c08501521660e083015251908190036101000190f35b61029160048036036040811015610a2657600080fd5b50803590602001356114e6565b61068960048036036020811015610a4957600080fd5b5035611507565b61045261156f565b6103606115e0565b604a60209081526000928352604080842090915290825290205481565b60408051808201909152600f81526e54656c6c6f7220547269627574657360881b602082015290565b60009081526034602052604090205490565b610ac06116c5565b600083815260456020908152604080832085845260060190915290819020815160a08101928390529160059082845b815481526020019060010190808311610aef575050505050905092915050565b6000908152604760205260409020546001600160a01b031690565b7fe6148e7230ca038d456350e69a91b66968b222bfac9ebfbea6ff0a1fb738016060005260466020527ffffeead1ec15181fd57b4590d95e0c076bccb59e311315e8b38f23c710aa7c3e5490565b603a8160058110610b8857600080fd5b6002020180546001909101549091506001600160a01b031682565b601290565b6000818152604560205260408120805482919015610bf4578054610be890859083906000198101908110610bd857fe5b9060005260206000200154610f29565b60019250925050610bfd565b60008092509250505b915091565b60009182526045602090815260408084209284526004909201905290205460ff1690565b6000610c306116c5565b60008060005b6005811015610c6a57603a8160058110610c4c57fe5b6002020154848260058110610c5d57fe5b6020020152600101610c36565b50507f52cb9007c7c6068f8ef37039d4f232cbf5a28ff8d93a5983c4c0c27cd2f9bc0d5460466020527f5bccd7373734898281f858d7562320d2cdfc0b17bd72f779686937174d150025547f09659d32f99e50ac728058418d38174fe83a137c455ff1847e6fb8e15f78f77a6000527f38b16d06a20ab673b01c748aff938df6a38f81640035f4ce8bd9abb03aae5b7254919450915090919293565b60338181548110610d1657600080fd5b600091825260209091200154905081565b60009081526045602052604090205490565b603960209081526000928352604080842090915290825290205460ff1681565b604b6020526000908152604090205460ff1681565b60376020526000908152604090205481565b60009081526046602052604090205490565b60006032821115610dea576040805162461bcd60e51b815260206004820152601a60248201527f526571756573745120696e6465782069732061626f7665203530000000000000604482015290519081900360640190fd5b5060009081526035602052604090205490565b60486020526000908152604090205481565b60009182526039602090815260408084206001600160a01b0393909316845291905290205460ff1690565b610e426116c5565b6000838152604560209081526040808320858452600590810190925291829020825160a08101938490529290919082845b81546001600160a01b03168152600190910190602001808311610e73575050505050905092915050565b6047602052600090815260409020546001600160a01b031681565b6001600160a01b0316600090815260446020526040902080546001909101549091565b6000828152604560205260408120805483908110610ef557fe5b9060005260206000200154905092915050565b60009182526036602090815260408084209284526005909201905290205490565b60009182526045602090815260408084209284526003909201905290205490565b6040805180820190915260038152622a292160e91b602082015290565b610f6f6116c5565b610f776116c5565b60208301516000805b6005811015610ffa57858160010160338110610f9857fe5b6020020151858260058110610fa957fe5b602002015260018101848260058110610fbe57fe5b602002015282858260058110610fd057fe5b60200201511015610ff257848160058110610fe757fe5b602002015192508091505b600101610f80565b5060065b60338110156110b3578286826033811061101457fe5b602002015111156110ab5785816033811061102b57fe5b602002015185836005811061103c57fe5b60200201528084836005811061104e57fe5b602002015285816033811061105f57fe5b6020020151925060005b60058110156110a9578386826005811061107f57fe5b602002015110156110a15785816005811061109657fe5b602002015193508092505b600101611069565b505b600101610ffe565b505050915091565b6110c36116c5565b6110cb6116c5565b6110d36115e0565b915060005b600581101561115557604560008483600581106110f157fe5b6020020151815260200190815260200160002060010160007f1590276b7f31dd8e2a06f9a92867333eeb3eddbc91e73b9833e3e55d8e34f77d60001b81526020019081526020016000205482826005811061114857fe5b60200201526001016110d8565b509091565b60008281526036602090815260408083206001600160a01b038516845260060190915290205460ff1692915050565b600080600080600080600061119c6116e3565b5050506000958652505060366020908152604080862080546002820154600383015460048401548551610120810187527f9f47a2659c3d32b749ae717d975e7962959890862423c4318cf86e4ec220291f8c5260058601808952878d205482527f2f9328a9c75282bec25bb04befad06926366736e0030c985108445fa728335e58d52808952878d2054828a01527f9147231ab14efb72c38117f68521ddef8de64f092c18c69dbfb602ffc4de7f478d52808952878d2054828901527f46f7d53798d31923f6952572c6a19ad2d1a8238d26649c2f3493a6d69e425d288d52808952878d205460608301527f1da378694063870452ce03b189f48e04c1aa026348e74e6c86e10738514ad2c48d52808952878d205460808301527f4b4cefd5ced7569ef0d091282b4bca9c52a034c56471a6061afd1bf307a2de7c8d52808952878d205460a08301527f6de96ee4d33a0617f40a846309c8759048857f51b9d59a12d3c3786d4778883d8d52808952878d205460c08301527f30e85ae205656781c1a951cba9f9f53f884833c049d377a2a7046eb5e6d14b268d52808952878d205460e08301527f1da95f11543c9b03927178e07951795dfc95c7501a9d1cf00e13414ca33bc4098d52909752949099205461010080870191909152600190930154919960ff8083169a948304811699506201000083041697506001600160a01b036301000000909204821696509281169493169291565b6113bc611702565b604080516106608101918290529060009060339082845b8154815260200190600101908083116113d3575050505050905090565b60466020526000908152604090205481565b60009182526045602090815260408084209284526002909201905290205490565b6049602052816000526040600020818154811061143f57600080fd5b6000918252602090912001546001600160801b038082169350600160801b90910416905082565b60386020526000908152604090205481565b60009081526038602052604090205490565b603660205260009081526040902080546001820154600283015460038401546004909401549293919260ff808316936101008404821693620100008104909216926001600160a01b036301000000909304831692918216911688565b60009182526045602090815260408084209284526001909201905290205490565b60009081526045602090815260408083207ff68d680ab3160f1aa5d9c3a1383c49e3e60bf3c0c031245cbb036f5ce99afaa18452600101909152808220547f1590276b7f31dd8e2a06f9a92867333eeb3eddbc91e73b9833e3e55d8e34f77d83529120549091565b7f231bb0dc207f13dd4e565ebc32496c470e35391bd8d3b6649269ee2328e031185460008181526034602090815260408220547f2c8b528fbaf48aaf13162a5a0519a7ad5a612da8ff8783465c17e076660a59f183526046909152909182916115d791610f29565b92600192509050565b6115e86116c5565b6115f06116c5565b6115f86116c5565b6040805161066081019182905261162f9160009060339082845b815481526020019060010190808311611612575050505050610f67565b909250905060005b60058110156116bf5782816005811061164c57fe5b602002015115611690576035600083836005811061166657fe5b602002015181526020019081526020016000205484826005811061168657fe5b60200201526116b7565b603a81600403600581106116a057fe5b60020201548482600581106116b157fe5b60200201525b600101611637565b50505090565b6040518060a001604052806005906020820280368337509192915050565b6040518061012001604052806009906020820280368337509192915050565b604051806106600160405280603390602082028036833750919291505056fea2646970667358221220492e32cde9cda1a51bc40d150b3625bc2931578d890fa413c1df8e6f0f6ce68e64736f6c63430007040033"

// DeployTellorGetters deploys a new Ethereum contract, binding an instance of TellorGetters to it.
func DeployTellorGetters(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorGetters, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorGettersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TellorGettersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TellorGetters{TellorGettersCaller: TellorGettersCaller{contract: contract}, TellorGettersTransactor: TellorGettersTransactor{contract: contract}, TellorGettersFilterer: TellorGettersFilterer{contract: contract}}, nil
}

// TellorGetters is an auto generated Go binding around an Ethereum contract.
type TellorGetters struct {
	TellorGettersCaller     // Read-only binding to the contract
	TellorGettersTransactor // Write-only binding to the contract
	TellorGettersFilterer   // Log filterer for contract events
}

// TellorGettersCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorGettersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorGettersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorGettersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorGettersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorGettersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorGettersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorGettersSession struct {
	Contract     *TellorGetters    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorGettersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorGettersCallerSession struct {
	Contract *TellorGettersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TellorGettersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorGettersTransactorSession struct {
	Contract     *TellorGettersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TellorGettersRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorGettersRaw struct {
	Contract *TellorGetters // Generic contract binding to access the raw methods on
}

// TellorGettersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorGettersCallerRaw struct {
	Contract *TellorGettersCaller // Generic read-only contract binding to access the raw methods on
}

// TellorGettersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorGettersTransactorRaw struct {
	Contract *TellorGettersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellorGetters creates a new instance of TellorGetters, bound to a specific deployed contract.
func NewTellorGetters(address common.Address, backend bind.ContractBackend) (*TellorGetters, error) {
	contract, err := bindTellorGetters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TellorGetters{TellorGettersCaller: TellorGettersCaller{contract: contract}, TellorGettersTransactor: TellorGettersTransactor{contract: contract}, TellorGettersFilterer: TellorGettersFilterer{contract: contract}}, nil
}

// NewTellorGettersCaller creates a new read-only instance of TellorGetters, bound to a specific deployed contract.
func NewTellorGettersCaller(address common.Address, caller bind.ContractCaller) (*TellorGettersCaller, error) {
	contract, err := bindTellorGetters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorGettersCaller{contract: contract}, nil
}

// NewTellorGettersTransactor creates a new write-only instance of TellorGetters, bound to a specific deployed contract.
func NewTellorGettersTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorGettersTransactor, error) {
	contract, err := bindTellorGetters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorGettersTransactor{contract: contract}, nil
}

// NewTellorGettersFilterer creates a new log filterer instance of TellorGetters, bound to a specific deployed contract.
func NewTellorGettersFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorGettersFilterer, error) {
	contract, err := bindTellorGetters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorGettersFilterer{contract: contract}, nil
}

// bindTellorGetters binds a generic wrapper to an already deployed contract.
func bindTellorGetters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorGettersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorGetters *TellorGettersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorGetters.Contract.TellorGettersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorGetters *TellorGettersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorGetters.Contract.TellorGettersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorGetters *TellorGettersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorGetters.Contract.TellorGettersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorGetters *TellorGettersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorGetters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorGetters *TellorGettersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorGetters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorGetters *TellorGettersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorGetters.Contract.contract.Transact(opts, method, params...)
}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "_allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorGetters *TellorGettersSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TellorGetters.Contract.Allowances(&_TellorGetters.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TellorGetters.Contract.Allowances(&_TellorGetters.CallOpts, arg0, arg1)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorGetters *TellorGettersCaller) Addresses(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "addresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorGetters *TellorGettersSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _TellorGetters.Contract.Addresses(&_TellorGetters.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorGetters *TellorGettersCallerSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _TellorGetters.Contract.Addresses(&_TellorGetters.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorGetters *TellorGettersCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "balances", arg0, arg1)

	outstruct := new(struct {
		FromBlock *big.Int
		Value     *big.Int
	})

	outstruct.FromBlock = out[0].(*big.Int)
	outstruct.Value = out[1].(*big.Int)

	return *outstruct, err

}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorGetters *TellorGettersSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _TellorGetters.Contract.Balances(&_TellorGetters.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorGetters *TellorGettersCallerSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _TellorGetters.Contract.Balances(&_TellorGetters.CallOpts, arg0, arg1)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) BytesVars(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "bytesVars", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorGetters *TellorGettersSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _TellorGetters.Contract.BytesVars(&_TellorGetters.CallOpts, arg0)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _TellorGetters.Contract.BytesVars(&_TellorGetters.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorGetters *TellorGettersCaller) CurrentMiners(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "currentMiners", arg0)

	outstruct := new(struct {
		Value *big.Int
		Miner common.Address
	})

	outstruct.Value = out[0].(*big.Int)
	outstruct.Miner = out[1].(common.Address)

	return *outstruct, err

}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorGetters *TellorGettersSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _TellorGetters.Contract.CurrentMiners(&_TellorGetters.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorGetters *TellorGettersCallerSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _TellorGetters.Contract.CurrentMiners(&_TellorGetters.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_TellorGetters *TellorGettersCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_TellorGetters *TellorGettersSession) Decimals() (uint8, error) {
	return _TellorGetters.Contract.Decimals(&_TellorGetters.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_TellorGetters *TellorGettersCallerSession) Decimals() (uint8, error) {
	return _TellorGetters.Contract.Decimals(&_TellorGetters.CallOpts)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_TellorGetters *TellorGettersCaller) DidMine(opts *bind.CallOpts, _challenge [32]byte, _miner common.Address) (bool, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "didMine", _challenge, _miner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_TellorGetters *TellorGettersSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _TellorGetters.Contract.DidMine(&_TellorGetters.CallOpts, _challenge, _miner)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_TellorGetters *TellorGettersCallerSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _TellorGetters.Contract.DidMine(&_TellorGetters.CallOpts, _challenge, _miner)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_TellorGetters *TellorGettersCaller) DidVote(opts *bind.CallOpts, _disputeId *big.Int, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "didVote", _disputeId, _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_TellorGetters *TellorGettersSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _TellorGetters.Contract.DidVote(&_TellorGetters.CallOpts, _disputeId, _address)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_TellorGetters *TellorGettersCallerSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _TellorGetters.Contract.DidVote(&_TellorGetters.CallOpts, _disputeId, _address)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) DisputeIdByDisputeHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "disputeIdByDisputeHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorGetters *TellorGettersSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.DisputeIdByDisputeHash(&_TellorGetters.CallOpts, arg0)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.DisputeIdByDisputeHash(&_TellorGetters.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorGetters *TellorGettersCaller) DisputesById(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "disputesById", arg0)

	outstruct := new(struct {
		Hash                [32]byte
		Tally               *big.Int
		Executed            bool
		DisputeVotePassed   bool
		IsPropFork          bool
		ReportedMiner       common.Address
		ReportingParty      common.Address
		ProposedForkAddress common.Address
	})

	outstruct.Hash = out[0].([32]byte)
	outstruct.Tally = out[1].(*big.Int)
	outstruct.Executed = out[2].(bool)
	outstruct.DisputeVotePassed = out[3].(bool)
	outstruct.IsPropFork = out[4].(bool)
	outstruct.ReportedMiner = out[5].(common.Address)
	outstruct.ReportingParty = out[6].(common.Address)
	outstruct.ProposedForkAddress = out[7].(common.Address)

	return *outstruct, err

}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorGetters *TellorGettersSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _TellorGetters.Contract.DisputesById(&_TellorGetters.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorGetters *TellorGettersCallerSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _TellorGetters.Contract.DisputesById(&_TellorGetters.CallOpts, arg0)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_TellorGetters *TellorGettersCaller) GetAddressVars(opts *bind.CallOpts, _data [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getAddressVars", _data)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_TellorGetters *TellorGettersSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _TellorGetters.Contract.GetAddressVars(&_TellorGetters.CallOpts, _data)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_TellorGetters *TellorGettersCallerSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _TellorGetters.Contract.GetAddressVars(&_TellorGetters.CallOpts, _data)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_TellorGetters *TellorGettersCaller) GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getAllDisputeVars", _disputeId)

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
func (_TellorGetters *TellorGettersSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetAllDisputeVars(&_TellorGetters.CallOpts, _disputeId)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_TellorGetters *TellorGettersCallerSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetAllDisputeVars(&_TellorGetters.CallOpts, _disputeId)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetDisputeIdByDisputeHash(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getDisputeIdByDisputeHash", _hash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetDisputeIdByDisputeHash(&_TellorGetters.CallOpts, _hash)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetDisputeIdByDisputeHash(&_TellorGetters.CallOpts, _hash)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetDisputeUintVars(opts *bind.CallOpts, _disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getDisputeUintVars", _disputeId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetDisputeUintVars(&_TellorGetters.CallOpts, _disputeId, _data)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetDisputeUintVars(&_TellorGetters.CallOpts, _disputeId, _data)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_TellorGetters *TellorGettersCaller) GetLastNewValue(opts *bind.CallOpts) (*big.Int, bool, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getLastNewValue")

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
func (_TellorGetters *TellorGettersSession) GetLastNewValue() (*big.Int, bool, error) {
	return _TellorGetters.Contract.GetLastNewValue(&_TellorGetters.CallOpts)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_TellorGetters *TellorGettersCallerSession) GetLastNewValue() (*big.Int, bool, error) {
	return _TellorGetters.Contract.GetLastNewValue(&_TellorGetters.CallOpts)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_TellorGetters *TellorGettersCaller) GetLastNewValueById(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, bool, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getLastNewValueById", _requestId)

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
func (_TellorGetters *TellorGettersSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _TellorGetters.Contract.GetLastNewValueById(&_TellorGetters.CallOpts, _requestId)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_TellorGetters *TellorGettersCallerSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _TellorGetters.Contract.GetLastNewValueById(&_TellorGetters.CallOpts, _requestId)
}

// GetMax5 is a free data retrieval call binding the contract method 0x99830e32.
//
// Solidity: function getMax5(uint256[51] data) view returns(uint256[5] max, uint256[5] maxIndex)
func (_TellorGetters *TellorGettersCaller) GetMax5(opts *bind.CallOpts, data [51]*big.Int) (struct {
	Max      [5]*big.Int
	MaxIndex [5]*big.Int
}, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getMax5", data)

	outstruct := new(struct {
		Max      [5]*big.Int
		MaxIndex [5]*big.Int
	})

	outstruct.Max = out[0].([5]*big.Int)
	outstruct.MaxIndex = out[1].([5]*big.Int)

	return *outstruct, err

}

// GetMax5 is a free data retrieval call binding the contract method 0x99830e32.
//
// Solidity: function getMax5(uint256[51] data) view returns(uint256[5] max, uint256[5] maxIndex)
func (_TellorGetters *TellorGettersSession) GetMax5(data [51]*big.Int) (struct {
	Max      [5]*big.Int
	MaxIndex [5]*big.Int
}, error) {
	return _TellorGetters.Contract.GetMax5(&_TellorGetters.CallOpts, data)
}

// GetMax5 is a free data retrieval call binding the contract method 0x99830e32.
//
// Solidity: function getMax5(uint256[51] data) view returns(uint256[5] max, uint256[5] maxIndex)
func (_TellorGetters *TellorGettersCallerSession) GetMax5(data [51]*big.Int) (struct {
	Max      [5]*big.Int
	MaxIndex [5]*big.Int
}, error) {
	return _TellorGetters.Contract.GetMax5(&_TellorGetters.CallOpts, data)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetMinedBlockNum(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getMinedBlockNum", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetMinedBlockNum(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetMinedBlockNum(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_TellorGetters *TellorGettersCaller) GetMinersByRequestIdAndTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getMinersByRequestIdAndTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([5]common.Address)).(*[5]common.Address)

	return out0, err

}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_TellorGetters *TellorGettersSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _TellorGetters.Contract.GetMinersByRequestIdAndTimestamp(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_TellorGetters *TellorGettersCallerSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _TellorGetters.Contract.GetMinersByRequestIdAndTimestamp(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _diff, uint256 _tip)
func (_TellorGetters *TellorGettersCaller) GetNewCurrentVariables(opts *bind.CallOpts) (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Diff       *big.Int
	Tip        *big.Int
}, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getNewCurrentVariables")

	outstruct := new(struct {
		Challenge  [32]byte
		RequestIds [5]*big.Int
		Diff       *big.Int
		Tip        *big.Int
	})

	outstruct.Challenge = out[0].([32]byte)
	outstruct.RequestIds = out[1].([5]*big.Int)
	outstruct.Diff = out[2].(*big.Int)
	outstruct.Tip = out[3].(*big.Int)

	return *outstruct, err

}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _diff, uint256 _tip)
func (_TellorGetters *TellorGettersSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Diff       *big.Int
	Tip        *big.Int
}, error) {
	return _TellorGetters.Contract.GetNewCurrentVariables(&_TellorGetters.CallOpts)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _diff, uint256 _tip)
func (_TellorGetters *TellorGettersCallerSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Diff       *big.Int
	Tip        *big.Int
}, error) {
	return _TellorGetters.Contract.GetNewCurrentVariables(&_TellorGetters.CallOpts)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getNewValueCountbyRequestId", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetNewValueCountbyRequestId(&_TellorGetters.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetNewValueCountbyRequestId(&_TellorGetters.CallOpts, _requestId)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_TellorGetters *TellorGettersCaller) GetNewVariablesOnDeck(opts *bind.CallOpts) (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getNewVariablesOnDeck")

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
func (_TellorGetters *TellorGettersSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _TellorGetters.Contract.GetNewVariablesOnDeck(&_TellorGetters.CallOpts)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_TellorGetters *TellorGettersCallerSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _TellorGetters.Contract.GetNewVariablesOnDeck(&_TellorGetters.CallOpts)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetRequestIdByRequestQIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getRequestIdByRequestQIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestIdByRequestQIndex(&_TellorGetters.CallOpts, _index)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestIdByRequestQIndex(&_TellorGetters.CallOpts, _index)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetRequestIdByTimestamp(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getRequestIdByTimestamp", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestIdByTimestamp(&_TellorGetters.CallOpts, _timestamp)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestIdByTimestamp(&_TellorGetters.CallOpts, _timestamp)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_TellorGetters *TellorGettersCaller) GetRequestQ(opts *bind.CallOpts) ([51]*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getRequestQ")

	if err != nil {
		return *new([51]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([51]*big.Int)).(*[51]*big.Int)

	return out0, err

}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_TellorGetters *TellorGettersSession) GetRequestQ() ([51]*big.Int, error) {
	return _TellorGetters.Contract.GetRequestQ(&_TellorGetters.CallOpts)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_TellorGetters *TellorGettersCallerSession) GetRequestQ() ([51]*big.Int, error) {
	return _TellorGetters.Contract.GetRequestQ(&_TellorGetters.CallOpts)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetRequestUintVars(opts *bind.CallOpts, _requestId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getRequestUintVars", _requestId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestUintVars(&_TellorGetters.CallOpts, _requestId, _data)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestUintVars(&_TellorGetters.CallOpts, _requestId, _data)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(uint256, uint256)
func (_TellorGetters *TellorGettersCaller) GetRequestVars(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getRequestVars", _requestId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(uint256, uint256)
func (_TellorGetters *TellorGettersSession) GetRequestVars(_requestId *big.Int) (*big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetRequestVars(&_TellorGetters.CallOpts, _requestId)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(uint256, uint256)
func (_TellorGetters *TellorGettersCallerSession) GetRequestVars(_requestId *big.Int) (*big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetRequestVars(&_TellorGetters.CallOpts, _requestId)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_TellorGetters *TellorGettersCaller) GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getStakerInfo", _staker)

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
func (_TellorGetters *TellorGettersSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetStakerInfo(&_TellorGetters.CallOpts, _staker)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_TellorGetters *TellorGettersCallerSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetStakerInfo(&_TellorGetters.CallOpts, _staker)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_TellorGetters *TellorGettersCaller) GetSubmissionsByTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getSubmissionsByTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_TellorGetters *TellorGettersSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _TellorGetters.Contract.GetSubmissionsByTimestamp(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_TellorGetters *TellorGettersCallerSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _TellorGetters.Contract.GetSubmissionsByTimestamp(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getTimestampbyRequestIDandIndex", _requestID, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetTimestampbyRequestIDandIndex(&_TellorGetters.CallOpts, _requestID, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetTimestampbyRequestIDandIndex(&_TellorGetters.CallOpts, _requestID, _index)
}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_TellorGetters *TellorGettersCaller) GetTopRequestIDs(opts *bind.CallOpts) ([5]*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getTopRequestIDs")

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_TellorGetters *TellorGettersSession) GetTopRequestIDs() ([5]*big.Int, error) {
	return _TellorGetters.Contract.GetTopRequestIDs(&_TellorGetters.CallOpts)
}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_TellorGetters *TellorGettersCallerSession) GetTopRequestIDs() ([5]*big.Int, error) {
	return _TellorGetters.Contract.GetTopRequestIDs(&_TellorGetters.CallOpts)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getUintVar", _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetUintVar(&_TellorGetters.CallOpts, _data)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetUintVar(&_TellorGetters.CallOpts, _data)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_TellorGetters *TellorGettersCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "isInDispute", _requestId, _timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_TellorGetters *TellorGettersSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _TellorGetters.Contract.IsInDispute(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_TellorGetters *TellorGettersCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _TellorGetters.Contract.IsInDispute(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorGetters *TellorGettersCaller) Migrated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "migrated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorGetters *TellorGettersSession) Migrated(arg0 common.Address) (bool, error) {
	return _TellorGetters.Contract.Migrated(&_TellorGetters.CallOpts, arg0)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorGetters *TellorGettersCallerSession) Migrated(arg0 common.Address) (bool, error) {
	return _TellorGetters.Contract.Migrated(&_TellorGetters.CallOpts, arg0)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorGetters *TellorGettersCaller) MinersByChallenge(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "minersByChallenge", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorGetters *TellorGettersSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _TellorGetters.Contract.MinersByChallenge(&_TellorGetters.CallOpts, arg0, arg1)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorGetters *TellorGettersCallerSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _TellorGetters.Contract.MinersByChallenge(&_TellorGetters.CallOpts, arg0, arg1)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_TellorGetters *TellorGettersCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_TellorGetters *TellorGettersSession) Name() (string, error) {
	return _TellorGetters.Contract.Name(&_TellorGetters.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_TellorGetters *TellorGettersCallerSession) Name() (string, error) {
	return _TellorGetters.Contract.Name(&_TellorGetters.CallOpts)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) NewValueTimestamps(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "newValueTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorGetters *TellorGettersSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.NewValueTimestamps(&_TellorGetters.CallOpts, arg0)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.NewValueTimestamps(&_TellorGetters.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) RequestIdByQueryHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "requestIdByQueryHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorGetters *TellorGettersSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.RequestIdByQueryHash(&_TellorGetters.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.RequestIdByQueryHash(&_TellorGetters.CallOpts, arg0)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "retrieveData", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.RetrieveData(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.RetrieveData(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_TellorGetters *TellorGettersCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_TellorGetters *TellorGettersSession) Symbol() (string, error) {
	return _TellorGetters.Contract.Symbol(&_TellorGetters.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_TellorGetters *TellorGettersCallerSession) Symbol() (string, error) {
	return _TellorGetters.Contract.Symbol(&_TellorGetters.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TellorGetters *TellorGettersCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TellorGetters *TellorGettersSession) TotalSupply() (*big.Int, error) {
	return _TellorGetters.Contract.TotalSupply(&_TellorGetters.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) TotalSupply() (*big.Int, error) {
	return _TellorGetters.Contract.TotalSupply(&_TellorGetters.CallOpts)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) Uints(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "uints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorGetters *TellorGettersSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.Uints(&_TellorGetters.CallOpts, arg0)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.Uints(&_TellorGetters.CallOpts, arg0)
}

// TellorStakeABI is the input ABI used to generate the binding from.
const TellorStakeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"NewDispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_position\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_voteWeight\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"fromBlock\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minerIndex\",\"type\":\"uint256\"}],\"name\":\"beginDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bytesVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentMiners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"disputesById\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"tally\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disputeVotePassed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPropFork\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"reportedMiner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reportingParty\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposedForkAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minersByChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"newValueTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"uints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"unlockDisputeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_supportsDispute\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TellorStakeFuncSigs maps the 4-byte function signature to its string representation.
var TellorStakeFuncSigs = map[string]string{
	"024c2ddd": "_allowances(address,address)",
	"699f200f": "addresses(bytes32)",
	"dd62ed3e": "allowance(address,address)",
	"999cf26c": "allowedToTrade(address,uint256)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"4ee2cd7e": "balanceOfAt(address,uint256)",
	"cbf1304d": "balances(address,uint256)",
	"8581af19": "beginDispute(uint256,uint256,uint256)",
	"62dd1d2a": "bytesVars(bytes32)",
	"1fd22364": "currentMiners(uint256)",
	"d01f4d9e": "disputeIdByDisputeHash(bytes32)",
	"db085beb": "disputesById(uint256)",
	"4ba0a5ee": "migrated(address)",
	"48b18e54": "minersByChallenge(bytes32,address)",
	"438c0aa3": "newValueTimestamps(uint256)",
	"5700242c": "requestIdByQueryHash(bytes32)",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"b59e14d4": "uints(bytes32)",
	"9a01ca13": "unlockDisputeFee(uint256)",
	"c9d27afe": "vote(uint256,bool)",
}

// TellorStakeBin is the compiled bytecode used for deploying new contracts.
var TellorStakeBin = "0x608060405234801561001057600080fd5b50612123806100206000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c806370a08231116100b8578063b59e14d41161007c578063b59e14d41461040f578063c9d27afe1461042c578063cbf1304d14610451578063d01f4d9e146104ac578063db085beb146104c9578063dd62ed3e1461053657610142565b806370a08231146103495780638581af191461036f578063999cf26c1461039a5780639a01ca13146103c6578063a9059cbb146103e357610142565b806348b18e541161010a57806348b18e54146102585780634ba0a5ee146102845780634ee2cd7e146102aa5780635700242c146102d657806362dd1d2a146102f3578063699f200f1461031057610142565b8063024c2ddd14610147578063095ea7b3146101875780631fd22364146101c757806323b872dd14610205578063438c0aa31461023b575b600080fd5b6101756004803603604081101561015d57600080fd5b506001600160a01b0381358116916020013516610564565b60408051918252519081900360200190f35b6101b36004803603604081101561019d57600080fd5b506001600160a01b038135169060200135610581565b604080519115158252519081900360200190f35b6101e4600480360360208110156101dd57600080fd5b503561066c565b604080519283526001600160a01b0390911660208301528051918290030190f35b6101b36004803603606081101561021b57600080fd5b506001600160a01b03813581169160208101359091169060400135610697565b6101756004803603602081101561025157600080fd5b5035610743565b6101b36004803603604081101561026e57600080fd5b50803590602001356001600160a01b0316610764565b6101b36004803603602081101561029a57600080fd5b50356001600160a01b0316610784565b610175600480360360408110156102c057600080fd5b506001600160a01b038135169060200135610799565b610175600480360360208110156102ec57600080fd5b503561093d565b6101756004803603602081101561030957600080fd5b503561094f565b61032d6004803603602081101561032657600080fd5b5035610961565b604080516001600160a01b039092168252519081900360200190f35b6101756004803603602081101561035f57600080fd5b50356001600160a01b031661097c565b6103986004803603606081101561038557600080fd5b5080359060208101359060400135610988565b005b6101b3600480360360408110156103b057600080fd5b506001600160a01b0381351690602001356111df565b610398600480360360208110156103dc57600080fd5b50356112a2565b6101b3600480360360408110156103f957600080fd5b506001600160a01b038135169060200135611a22565b6101756004803603602081101561042557600080fd5b5035611a38565b6103986004803603604081101561044257600080fd5b50803590602001351515611a4a565b61047d6004803603604081101561046757600080fd5b506001600160a01b038135169060200135611c75565b60405180836001600160801b03168152602001826001600160801b031681526020019250505060405180910390f35b610175600480360360208110156104c257600080fd5b5035611cb8565b6104e6600480360360208110156104df57600080fd5b5035611cca565b60408051988952602089019790975294151587870152921515606087015290151560808601526001600160a01b0390811660a086015290811660c08501521660e083015251908190036101000190f35b6101756004803603604081101561054c57600080fd5b506001600160a01b0381358116916020013516611d26565b604a60209081526000928352604080842090915290825290205481565b6000336105bf5760405162461bcd60e51b81526004018080602001828103825260248152602001806120ca6024913960400191505060405180910390fd5b6001600160a01b0383166106045760405162461bcd60e51b81526004018080602001828103825260228152602001806120406022913960400191505060405180910390fd5b336000818152604a602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060015b92915050565b603a816005811061067c57600080fd5b6002020180546001909101549091506001600160a01b031682565b6001600160a01b0383166000908152604a60209081526040808320338452909152812054821115610704576040805162461bcd60e51b8152602060048201526012602482015271416c6c6f77616e63652069732077726f6e6760701b604482015290519081900360640190fd5b6001600160a01b0384166000908152604a60209081526040808320338452909152902080548390039055610739848484611d51565b5060019392505050565b6033818154811061075357600080fd5b600091825260209091200154905081565b603960209081526000928352604080842090915290825290205460ff1681565b604b6020526000908152604090205460ff1681565b6001600160a01b0382166000908152604960205260408120805415806107df575082816000815481106107c857fe5b6000918252602090912001546001600160801b0316115b156107ee576000915050610666565b80548190600019810190811061080057fe5b6000918252602090912001546001600160801b031683106108525780548190600019810190811061082d57fe5b600091825260209091200154600160801b90046001600160801b031691506106669050565b8054600090600119015b8181111561090a57600060026001838501010490508584828154811061087e57fe5b6000918252602090912001546001600160801b031614156108cd578381815481106108a557fe5b600091825260209091200154600160801b90046001600160801b031694506106669350505050565b858482815481106108da57fe5b6000918252602090912001546001600160801b031610156108fd57809250610904565b6001810391505b5061085c565b82828154811061091657fe5b600091825260209091200154600160801b90046001600160801b0316935061066692505050565b60376020526000908152604090205481565b60486020526000908152604090205481565b6047602052600090815260409020546001600160a01b031681565b60006106668243610799565b600083815260456020908152604080832085845260028101909252909120546109eb576040805162461bcd60e51b815260206004820152601060248201526f04d696e656420626c6f636b20697320360841b604482015290519081900360640190fd5b60058210610a37576040805162461bcd60e51b81526020600482015260146024820152734d696e657220696e6465782069732077726f6e6760601b604482015290519081900360640190fd5b60008381526005808301602052604082209084908110610a5357fe5b0154604080516bffffffffffffffffffffffff19606084901b1660208083019190915260348201899052605480830189905283518084039091018152607490920183528151918101919091207f1ce2382bc92689b00ba121fa5a411aa976168affdd8ac143a69035dd984b3b6a8054600101908190556000828152603890935292909120546001600160a01b039093169350918015610b305760008281526036602090815260408083207fed92b4c1e0a9e559a31171d487ecbec963526662038ecfa3a71160bd62fb873384526005019091529020819055610b44565b506000828152603860205260409020819055805b60008181526036602081815260408084207f6ab2b18aafe78fd59c6a4092015bddd9fcacb8170f72b299074f74d76a91a92385526005018083528185208054600101908190558686529383528151808401859052825180820385018152908301835280519084012085529091529091208390558190838214610d3a57600082815260366020818152604080842081516000198701818501528251808203850181529083018352805190840120855260059081018352818520548086529383528185207f46f7d53798d31923f6952572c6a19ad2d1a8238d26649c2f3493a6d69e425d2886520190915290912054421015610c85576040805162461bcd60e51b815260206004820152601760248201527f4469737075746520697320616c7265616479206f70656e000000000000000000604482015290519081900360640190fd5b60008181526036602052604090206002015460ff1615610d385760008181526036602090815260408083207ff9e1ae10923bfc79f52e309baf8c7699edb821f91ef5b5bd07be29545917b3a684526005019091529020546201518042919091031115610d38576040805162461bcd60e51b815260206004820152601f60248201527f54696d6520666f7220766f74696e6720686176656e277420656c617073656400604482015290519081900360640190fd5b505b60008860021415610ddd575060008a81526045602090815260408083207f310199159a20c50879ffb440b45802138b5b162ec9426720e9dd3ee8bbcdb9d78452600190810183529083208054909101908190557f5d9fadfc729fd027e395e5157ef1b53ef9fa4a8f053043c5f159307543e7cc97909252604690527f167af83a0768d27540775cfef6d996eb63f8a61fcdfb26e654c18fb50960e3be5402610e2c565b507f675d2171f68d6f5545d54fb9b1fb61a0e6897e6188ca1cd664e7c9530d91ecfc60005260466020527f3e5522f19747f0f285b96ded572ac4128c3a764aea9f44058dc0afc9dda449865481025b85603660008781526020019081526020016000206000018190555060006036600087815260200190815260200160002060020160026101000a81548160ff021916908315150217905550866036600087815260200190815260200160002060020160036101000a8154816001600160a01b0302191690836001600160a01b03160217905550336036600087815260200190815260200160002060030160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060006036600087815260200190815260200160002060040160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060006036600087815260200190815260200160002060020160006101000a81548160ff02191690831515021790555060006036600087815260200190815260200160002060020160016101000a81548160ff021916908315150217905550600060366000878152602001908152602001600020600101819055508a6036600087815260200190815260200160002060050160007f9f47a2659c3d32b749ae717d975e7962959890862423c4318cf86e4ec220291f60001b8152602001908152602001600020819055508960366000878152602001908152602001600020600501600060008051602061208983398151915260001b8152602001908152602001600020819055508760060160008b8152602001908152602001600020896005811061104e57fe5b015460008681526036602090815260408083207f9147231ab14efb72c38117f68521ddef8de64f092c18c69dbfb602ffc4de7f478452600501909152808220929092557f46f7d53798d31923f6952572c6a19ad2d1a8238d26649c2f3493a6d69e425d2881528181206202a3008502420190557f4b4cefd5ced7569ef0d091282b4bca9c52a034c56471a6061afd1bf307a2de7c81528181204390557f6de96ee4d33a0617f40a846309c8759048857f51b9d59a12d3c3786d4778883d81528181208b90557f1da95f11543c9b03927178e07951795dfc95c7501a9d1cf00e13414ca33bc409815220819055611145333083611d51565b88600214156111775760008a81526004890160209081526040808320805460ff1916600117905560038b019091528120555b6001600160a01b0387166000818152604460209081526040918290206003905581518d81529081019290925280518d9288927feceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da6492918290030190a35050505050505050505050565b6001600160a01b0382166000908152604460205260408120541580159061121e57506001600160a01b0383166000908152604460205260409020546005115b1561128f577f5d9fadfc729fd027e395e5157ef1b53ef9fa4a8f053043c5f159307543e7cc9760005260466020527f167af83a0768d27540775cfef6d996eb63f8a61fcdfb26e654c18fb50960e3be5482906112798561097c565b031061128757506001610666565b506000610666565b816112998461097c565b10159392505050565b600081815260366020818152604080842054845260388252808420548085529282528084207f6ab2b18aafe78fd59c6a4092015bddd9fcacb8170f72b299074f74d76a91a9238552600501808352818520548251808501919091528251808203850181529083018352805190840120855290915290912054806113225750805b60008281526036602090815260408083208484528184207f6ab2b18aafe78fd59c6a4092015bddd9fcacb8170f72b299074f74d76a91a9238552600582019093529220548061136f575060015b7f29169706298d2b6df50a532e958b56426de1465348b93650fca42d456eaec5fc6000908152600584016020526040812054156113e6576040805162461bcd60e51b815260206004820152601060248201526f185b1c9958591e481c185a59081bdd5d60821b604482015290519081900360640190fd5b7ff9e1ae10923bfc79f52e309baf8c7699edb821f91ef5b5bd07be29545917b3a6600090815260058401602052604090205462015180429190910311611473576040805162461bcd60e51b815260206004820152601f60248201527f54696d6520666f7220766f74696e6720686176656e277420656c617073656400604482015290519081900360640190fd5b600284810154630100000090046001600160a01b031660009081526044602090815260408083207f29169706298d2b6df50a532e958b56426de1465348b93650fca42d456eaec5fc84526005890190925290912060019081905591850154909161010090910460ff161515141561176957620151804206420360018201557fa5ae3e2b97d73fb849ea855d27f073b72815b38452d976bd57e4a157827dadd380546000190190557f2b2a1c876f73e67ebc4f1b08d10d54d62d62216382e0f4fd16c29155818207a4600052604760209081527ffe0323da4092f31e73ad4b4aa705eaa20d7ce93cdb6c891e7c038c2a7146f00854604080516004815260248101825292830180516001600160e01b0316630287018760e31b178152905183516001600160a01b039093169392909182918083835b602083106115c65780518252601f1990920191602091820191016115a7565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114611626576040519150601f19603f3d011682016040523d82523d6000602084013e61162b565b606091505b505081546004141590506116b55760058155600285015460038601547f5d9fadfc729fd027e395e5157ef1b53ef9fa4a8f053043c5f159307543e7cc9760005260466020527f167af83a0768d27540775cfef6d996eb63f8a61fcdfb26e654c18fb50960e3be546116b0926001600160a01b0363010000009091048116921690611d51565b600081555b60005b83811015611763576040805182860360208083019190915282518083038201815291830183528151918101919091206000908152600589019091522054925082611700578792505b600083815260366020908152604080832060038101547f1da95f11543c9b03927178e07951795dfc95c7501a9d1cf00e13414ca33bc40985526005820190935292205461175a9130916001600160a01b0390911690611d51565b506001016116b8565b5061196c565b600181557f9f47a2659c3d32b749ae717d975e7962959890862423c4318cf86e4ec220291f60009081526005860160208181526040808420548452604582528084207f6de96ee4d33a0617f40a846309c8759048857f51b9d59a12d3c3786d4778883d8552929091529091205460021415611836577f9147231ab14efb72c38117f68521ddef8de64f092c18c69dbfb602ffc4de7f476000908152600587016020908152604080832054600080516020612089833981519152845281842054845260038501909252909120555b600080516020612089833981519152600090815260058701602090815260408083205483526004840190915290205460ff161515600114156118a95760008051602061208983398151915260009081526005870160209081526040808320548352600484019091529020805460ff191690555b60005b8481101561196957604080518287036020808301919091528251808303820181529183018352815191810191909120600090815260058a0190915220549350831561190257600084815260366020526040902095505b600286015460008581526036602090815260408083207f1da95f11543c9b03927178e07951795dfc95c7501a9d1cf00e13414ca33bc409845260050190915290205461196191309163010000009091046001600160a01b031690611d51565b6001016118ac565b50505b7f6de96ee4d33a0617f40a846309c8759048857f51b9d59a12d3c3786d4778883d600090815260058601602052604090205460021415611a18577f9f47a2659c3d32b749ae717d975e7962959890862423c4318cf86e4ec220291f60009081526005860160209081526040808320548352604582528083207f310199159a20c50879ffb440b45802138b5b162ec9426720e9dd3ee8bbcdb9d78452600101909152902080546000190190555b5050505050505050565b6000611a2f338484611d51565b50600192915050565b60466020526000908152604090205481565b60008281526036602090815260408083207f4b4cefd5ced7569ef0d091282b4bca9c52a034c56471a6061afd1bf307a2de7c845260058101909252822054909190611a96903390610799565b33600090815260068401602052604090205490915060ff16151560011415611b05576040805162461bcd60e51b815260206004820152601860248201527f53656e6465722068617320616c726561647920766f7465640000000000000000604482015290519081900360640190fd5b80611b4b576040805162461bcd60e51b81526020600482015260116024820152700557365722062616c616e6365206973203607c1b604482015290519081900360640190fd5b3360009081526044602052604090205460031415611ba9576040805162461bcd60e51b81526020600482015260166024820152754d696e657220697320756e646572206469737075746560501b604482015290519081900360640190fd5b3360009081526006830160209081526040808320805460ff191660019081179091557f1da378694063870452ce03b189f48e04c1aa026348e74e6c86e10738514ad2c4845260058601909252909120805490910190558215611c1e576001820154611c149082611ef1565b6001830155611c33565b6001820154611c2d9082611f1c565b60018301555b60408051841515815290518291339187917f911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e919081900360200190a450505050565b60496020528160005260406000208181548110611c9157600080fd5b6000918252602090912001546001600160801b038082169350600160801b90910416905082565b60386020526000908152604090205481565b603660205260009081526040902080546001820154600283015460038401546004909401549293919260ff808316936101008404821693620100008104909216926001600160a01b036301000000909304831692918216911688565b6001600160a01b039182166000908152604a6020908152604080832093909416825291909152205490565b80611d8d5760405162461bcd60e51b81526004018080602001828103825260218152602001806120a96021913960400191505060405180910390fd5b6001600160a01b038216611de0576040805162461bcd60e51b815260206004820152601560248201527452656365697665722069732030206164647265737360581b604482015290519081900360640190fd5b611dea83826111df565b611e255760405162461bcd60e51b81526004018080602001828103825260278152602001806120626027913960400191505060405180910390fd5b6000611e308461097c565b9050611e3e84838303611f42565b611e478361097c565b9050808282011015611e94576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b611ea083838301611f42565b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a350505050565b600080821315611f0e575081810182811215611f0957fe5b610666565b508181018281131561066657fe5b600080821315611f34575080820382811315611f0957fe5b508082038281121561066657fe5b6001600160a01b038216600090815260496020526040902080541580611f8f57508054439082906000198101908110611f7757fe5b6000918252602090912001546001600160801b031614155b156120005760408051808201909152436001600160801b0390811682528381166020808401918252845460018101865560008681529190912093519301805491516fffffffffffffffffffffffffffffffff19909216938316939093178216600160801b919092160217905561203a565b80546000908290600019810190811061201557fe5b600091825260209091200180546001600160801b03808616600160801b029116179055505b50505056fe45524332303a20617070726f766520746f20746865207a65726f206164647265737353686f756c6420686176652073756666696369656e742062616c616e636520746f2074726164652f9328a9c75282bec25bb04befad06926366736e0030c985108445fa728335e5547269656420746f2073656e64206e6f6e2d706f73697469766520616d6f756e7445524332303a20617070726f76652066726f6d20746865207a65726f2061646472657373a26469706673582212208598f88be8ecebc9494a63edb598bf0917fcc45d3dffac20a58d3ada4651a6f964736f6c63430007040033"

// DeployTellorStake deploys a new Ethereum contract, binding an instance of TellorStake to it.
func DeployTellorStake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorStake, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorStakeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

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

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorStake *TellorStakeCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "_allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorStake *TellorStakeSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TellorStake.Contract.Allowances(&_TellorStake.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorStake *TellorStakeCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TellorStake.Contract.Allowances(&_TellorStake.CallOpts, arg0, arg1)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorStake *TellorStakeCaller) Addresses(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "addresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorStake *TellorStakeSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _TellorStake.Contract.Addresses(&_TellorStake.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorStake *TellorStakeCallerSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _TellorStake.Contract.Addresses(&_TellorStake.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_TellorStake *TellorStakeCaller) Allowance(opts *bind.CallOpts, _user common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "allowance", _user, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_TellorStake *TellorStakeSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _TellorStake.Contract.Allowance(&_TellorStake.CallOpts, _user, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_TellorStake *TellorStakeCallerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _TellorStake.Contract.Allowance(&_TellorStake.CallOpts, _user, _spender)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_TellorStake *TellorStakeCaller) AllowedToTrade(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "allowedToTrade", _user, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_TellorStake *TellorStakeSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _TellorStake.Contract.AllowedToTrade(&_TellorStake.CallOpts, _user, _amount)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_TellorStake *TellorStakeCallerSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _TellorStake.Contract.AllowedToTrade(&_TellorStake.CallOpts, _user, _amount)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_TellorStake *TellorStakeCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "balanceOf", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_TellorStake *TellorStakeSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _TellorStake.Contract.BalanceOf(&_TellorStake.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_TellorStake *TellorStakeCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _TellorStake.Contract.BalanceOf(&_TellorStake.CallOpts, _user)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_TellorStake *TellorStakeCaller) BalanceOfAt(opts *bind.CallOpts, _user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "balanceOfAt", _user, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_TellorStake *TellorStakeSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _TellorStake.Contract.BalanceOfAt(&_TellorStake.CallOpts, _user, _blockNumber)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_TellorStake *TellorStakeCallerSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _TellorStake.Contract.BalanceOfAt(&_TellorStake.CallOpts, _user, _blockNumber)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorStake *TellorStakeCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "balances", arg0, arg1)

	outstruct := new(struct {
		FromBlock *big.Int
		Value     *big.Int
	})

	outstruct.FromBlock = out[0].(*big.Int)
	outstruct.Value = out[1].(*big.Int)

	return *outstruct, err

}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorStake *TellorStakeSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _TellorStake.Contract.Balances(&_TellorStake.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorStake *TellorStakeCallerSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _TellorStake.Contract.Balances(&_TellorStake.CallOpts, arg0, arg1)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorStake *TellorStakeCaller) BytesVars(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "bytesVars", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorStake *TellorStakeSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _TellorStake.Contract.BytesVars(&_TellorStake.CallOpts, arg0)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorStake *TellorStakeCallerSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _TellorStake.Contract.BytesVars(&_TellorStake.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorStake *TellorStakeCaller) CurrentMiners(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "currentMiners", arg0)

	outstruct := new(struct {
		Value *big.Int
		Miner common.Address
	})

	outstruct.Value = out[0].(*big.Int)
	outstruct.Miner = out[1].(common.Address)

	return *outstruct, err

}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorStake *TellorStakeSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _TellorStake.Contract.CurrentMiners(&_TellorStake.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorStake *TellorStakeCallerSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _TellorStake.Contract.CurrentMiners(&_TellorStake.CallOpts, arg0)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorStake *TellorStakeCaller) DisputeIdByDisputeHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "disputeIdByDisputeHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorStake *TellorStakeSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorStake.Contract.DisputeIdByDisputeHash(&_TellorStake.CallOpts, arg0)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorStake *TellorStakeCallerSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorStake.Contract.DisputeIdByDisputeHash(&_TellorStake.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorStake *TellorStakeCaller) DisputesById(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "disputesById", arg0)

	outstruct := new(struct {
		Hash                [32]byte
		Tally               *big.Int
		Executed            bool
		DisputeVotePassed   bool
		IsPropFork          bool
		ReportedMiner       common.Address
		ReportingParty      common.Address
		ProposedForkAddress common.Address
	})

	outstruct.Hash = out[0].([32]byte)
	outstruct.Tally = out[1].(*big.Int)
	outstruct.Executed = out[2].(bool)
	outstruct.DisputeVotePassed = out[3].(bool)
	outstruct.IsPropFork = out[4].(bool)
	outstruct.ReportedMiner = out[5].(common.Address)
	outstruct.ReportingParty = out[6].(common.Address)
	outstruct.ProposedForkAddress = out[7].(common.Address)

	return *outstruct, err

}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorStake *TellorStakeSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _TellorStake.Contract.DisputesById(&_TellorStake.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorStake *TellorStakeCallerSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _TellorStake.Contract.DisputesById(&_TellorStake.CallOpts, arg0)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorStake *TellorStakeCaller) Migrated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "migrated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorStake *TellorStakeSession) Migrated(arg0 common.Address) (bool, error) {
	return _TellorStake.Contract.Migrated(&_TellorStake.CallOpts, arg0)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorStake *TellorStakeCallerSession) Migrated(arg0 common.Address) (bool, error) {
	return _TellorStake.Contract.Migrated(&_TellorStake.CallOpts, arg0)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorStake *TellorStakeCaller) MinersByChallenge(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "minersByChallenge", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorStake *TellorStakeSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _TellorStake.Contract.MinersByChallenge(&_TellorStake.CallOpts, arg0, arg1)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorStake *TellorStakeCallerSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _TellorStake.Contract.MinersByChallenge(&_TellorStake.CallOpts, arg0, arg1)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorStake *TellorStakeCaller) NewValueTimestamps(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "newValueTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorStake *TellorStakeSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _TellorStake.Contract.NewValueTimestamps(&_TellorStake.CallOpts, arg0)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorStake *TellorStakeCallerSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _TellorStake.Contract.NewValueTimestamps(&_TellorStake.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorStake *TellorStakeCaller) RequestIdByQueryHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "requestIdByQueryHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorStake *TellorStakeSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorStake.Contract.RequestIdByQueryHash(&_TellorStake.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorStake *TellorStakeCallerSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorStake.Contract.RequestIdByQueryHash(&_TellorStake.CallOpts, arg0)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorStake *TellorStakeCaller) Uints(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "uints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorStake *TellorStakeSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _TellorStake.Contract.Uints(&_TellorStake.CallOpts, arg0)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorStake *TellorStakeCallerSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _TellorStake.Contract.Uints(&_TellorStake.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_TellorStake *TellorStakeTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorStake.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_TellorStake *TellorStakeSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorStake.Contract.Approve(&_TellorStake.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_TellorStake *TellorStakeTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorStake.Contract.Approve(&_TellorStake.TransactOpts, _spender, _amount)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_TellorStake *TellorStakeTransactor) BeginDispute(opts *bind.TransactOpts, _requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _TellorStake.contract.Transact(opts, "beginDispute", _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_TellorStake *TellorStakeSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _TellorStake.Contract.BeginDispute(&_TellorStake.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_TellorStake *TellorStakeTransactorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _TellorStake.Contract.BeginDispute(&_TellorStake.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_TellorStake *TellorStakeTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorStake.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_TellorStake *TellorStakeSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorStake.Contract.Transfer(&_TellorStake.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_TellorStake *TellorStakeTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorStake.Contract.Transfer(&_TellorStake.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_TellorStake *TellorStakeTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorStake.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_TellorStake *TellorStakeSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorStake.Contract.TransferFrom(&_TellorStake.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_TellorStake *TellorStakeTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorStake.Contract.TransferFrom(&_TellorStake.TransactOpts, _from, _to, _amount)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_TellorStake *TellorStakeTransactor) UnlockDisputeFee(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _TellorStake.contract.Transact(opts, "unlockDisputeFee", _disputeId)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_TellorStake *TellorStakeSession) UnlockDisputeFee(_disputeId *big.Int) (*types.Transaction, error) {
	return _TellorStake.Contract.UnlockDisputeFee(&_TellorStake.TransactOpts, _disputeId)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_TellorStake *TellorStakeTransactorSession) UnlockDisputeFee(_disputeId *big.Int) (*types.Transaction, error) {
	return _TellorStake.Contract.UnlockDisputeFee(&_TellorStake.TransactOpts, _disputeId)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_TellorStake *TellorStakeTransactor) Vote(opts *bind.TransactOpts, _disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _TellorStake.contract.Transact(opts, "vote", _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_TellorStake *TellorStakeSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _TellorStake.Contract.Vote(&_TellorStake.TransactOpts, _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_TellorStake *TellorStakeTransactorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _TellorStake.Contract.Vote(&_TellorStake.TransactOpts, _disputeId, _supportsDispute)
}

// TellorStakeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TellorStake contract.
type TellorStakeApprovalIterator struct {
	Event *TellorStakeApproval // Event containing the contract specifics and raw log

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
func (it *TellorStakeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorStakeApproval)
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
		it.Event = new(TellorStakeApproval)
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
func (it *TellorStakeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorStakeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorStakeApproval represents a Approval event raised by the TellorStake contract.
type TellorStakeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TellorStake *TellorStakeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TellorStakeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TellorStake.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TellorStakeApprovalIterator{contract: _TellorStake.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TellorStake *TellorStakeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TellorStakeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TellorStake.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorStakeApproval)
				if err := _TellorStake.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TellorStake *TellorStakeFilterer) ParseApproval(log types.Log) (*TellorStakeApproval, error) {
	event := new(TellorStakeApproval)
	if err := _TellorStake.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorStakeNewDisputeIterator is returned from FilterNewDispute and is used to iterate over the raw logs and unpacked data for NewDispute events raised by the TellorStake contract.
type TellorStakeNewDisputeIterator struct {
	Event *TellorStakeNewDispute // Event containing the contract specifics and raw log

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
func (it *TellorStakeNewDisputeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorStakeNewDispute)
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
		it.Event = new(TellorStakeNewDispute)
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
func (it *TellorStakeNewDisputeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorStakeNewDisputeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorStakeNewDispute represents a NewDispute event raised by the TellorStake contract.
type TellorStakeNewDispute struct {
	DisputeId *big.Int
	RequestId *big.Int
	Timestamp *big.Int
	Miner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewDispute is a free log retrieval operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_TellorStake *TellorStakeFilterer) FilterNewDispute(opts *bind.FilterOpts, _disputeId []*big.Int, _requestId []*big.Int) (*TellorStakeNewDisputeIterator, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _TellorStake.contract.FilterLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &TellorStakeNewDisputeIterator{contract: _TellorStake.contract, event: "NewDispute", logs: logs, sub: sub}, nil
}

// WatchNewDispute is a free log subscription operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_TellorStake *TellorStakeFilterer) WatchNewDispute(opts *bind.WatchOpts, sink chan<- *TellorStakeNewDispute, _disputeId []*big.Int, _requestId []*big.Int) (event.Subscription, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _TellorStake.contract.WatchLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorStakeNewDispute)
				if err := _TellorStake.contract.UnpackLog(event, "NewDispute", log); err != nil {
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
func (_TellorStake *TellorStakeFilterer) ParseNewDispute(log types.Log) (*TellorStakeNewDispute, error) {
	event := new(TellorStakeNewDispute)
	if err := _TellorStake.contract.UnpackLog(event, "NewDispute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorStakeTransferredIterator is returned from FilterTransferred and is used to iterate over the raw logs and unpacked data for Transferred events raised by the TellorStake contract.
type TellorStakeTransferredIterator struct {
	Event *TellorStakeTransferred // Event containing the contract specifics and raw log

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
func (it *TellorStakeTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorStakeTransferred)
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
		it.Event = new(TellorStakeTransferred)
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
func (it *TellorStakeTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorStakeTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorStakeTransferred represents a Transferred event raised by the TellorStake contract.
type TellorStakeTransferred struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferred is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TellorStake *TellorStakeFilterer) FilterTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TellorStakeTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TellorStake.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TellorStakeTransferredIterator{contract: _TellorStake.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransferred is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TellorStake *TellorStakeFilterer) WatchTransferred(opts *bind.WatchOpts, sink chan<- *TellorStakeTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TellorStake.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorStakeTransferred)
				if err := _TellorStake.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransferred is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TellorStake *TellorStakeFilterer) ParseTransferred(log types.Log) (*TellorStakeTransferred, error) {
	event := new(TellorStakeTransferred)
	if err := _TellorStake.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorStakeVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the TellorStake contract.
type TellorStakeVotedIterator struct {
	Event *TellorStakeVoted // Event containing the contract specifics and raw log

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
func (it *TellorStakeVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorStakeVoted)
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
		it.Event = new(TellorStakeVoted)
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
func (it *TellorStakeVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorStakeVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorStakeVoted represents a Voted event raised by the TellorStake contract.
type TellorStakeVoted struct {
	DisputeID  *big.Int
	Position   bool
	Voter      common.Address
	VoteWeight *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter, uint256 indexed _voteWeight)
func (_TellorStake *TellorStakeFilterer) FilterVoted(opts *bind.FilterOpts, _disputeID []*big.Int, _voter []common.Address, _voteWeight []*big.Int) (*TellorStakeVotedIterator, error) {

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

	logs, sub, err := _TellorStake.contract.FilterLogs(opts, "Voted", _disputeIDRule, _voterRule, _voteWeightRule)
	if err != nil {
		return nil, err
	}
	return &TellorStakeVotedIterator{contract: _TellorStake.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter, uint256 indexed _voteWeight)
func (_TellorStake *TellorStakeFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *TellorStakeVoted, _disputeID []*big.Int, _voter []common.Address, _voteWeight []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _TellorStake.contract.WatchLogs(opts, "Voted", _disputeIDRule, _voterRule, _voteWeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorStakeVoted)
				if err := _TellorStake.contract.UnpackLog(event, "Voted", log); err != nil {
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
func (_TellorStake *TellorStakeFilterer) ParseVoted(log types.Log) (*TellorStakeVoted, error) {
	event := new(TellorStakeVoted)
	if err := _TellorStake.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorStorageABI is the input ABI used to generate the binding from.
const TellorStorageABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"fromBlock\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bytesVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentMiners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"disputesById\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"tally\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disputeVotePassed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPropFork\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"reportedMiner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reportingParty\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposedForkAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minersByChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"newValueTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"uints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TellorStorageFuncSigs maps the 4-byte function signature to its string representation.
var TellorStorageFuncSigs = map[string]string{
	"024c2ddd": "_allowances(address,address)",
	"699f200f": "addresses(bytes32)",
	"cbf1304d": "balances(address,uint256)",
	"62dd1d2a": "bytesVars(bytes32)",
	"1fd22364": "currentMiners(uint256)",
	"d01f4d9e": "disputeIdByDisputeHash(bytes32)",
	"db085beb": "disputesById(uint256)",
	"4ba0a5ee": "migrated(address)",
	"48b18e54": "minersByChallenge(bytes32,address)",
	"438c0aa3": "newValueTimestamps(uint256)",
	"5700242c": "requestIdByQueryHash(bytes32)",
	"b59e14d4": "uints(bytes32)",
}

// TellorStorageBin is the compiled bytecode used for deploying new contracts.
var TellorStorageBin = "0x608060405234801561001057600080fd5b50610505806100206000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806362dd1d2a1161007157806362dd1d2a146101d7578063699f200f146101f4578063b59e14d41461022d578063cbf1304d1461024a578063d01f4d9e146102a5578063db085beb146102c2576100b4565b8063024c2ddd146100b95780631fd22364146100f9578063438c0aa31461013757806348b18e54146101545780634ba0a5ee146101945780635700242c146101ba575b600080fd5b6100e7600480360360408110156100cf57600080fd5b506001600160a01b038135811691602001351661032f565b60408051918252519081900360200190f35b6101166004803603602081101561010f57600080fd5b503561034c565b604080519283526001600160a01b0390911660208301528051918290030190f35b6100e76004803603602081101561014d57600080fd5b5035610377565b6101806004803603604081101561016a57600080fd5b50803590602001356001600160a01b0316610398565b604080519115158252519081900360200190f35b610180600480360360208110156101aa57600080fd5b50356001600160a01b03166103b8565b6100e7600480360360208110156101d057600080fd5b50356103cd565b6100e7600480360360208110156101ed57600080fd5b50356103df565b6102116004803603602081101561020a57600080fd5b50356103f1565b604080516001600160a01b039092168252519081900360200190f35b6100e76004803603602081101561024357600080fd5b503561040c565b6102766004803603604081101561026057600080fd5b506001600160a01b03813516906020013561041e565b60405180836001600160801b03168152602001826001600160801b031681526020019250505060405180910390f35b6100e7600480360360208110156102bb57600080fd5b5035610461565b6102df600480360360208110156102d857600080fd5b5035610473565b60408051988952602089019790975294151587870152921515606087015290151560808601526001600160a01b0390811660a086015290811660c08501521660e083015251908190036101000190f35b604a60209081526000928352604080842090915290825290205481565b603a816005811061035c57600080fd5b6002020180546001909101549091506001600160a01b031682565b6033818154811061038757600080fd5b600091825260209091200154905081565b603960209081526000928352604080842090915290825290205460ff1681565b604b6020526000908152604090205460ff1681565b60376020526000908152604090205481565b60486020526000908152604090205481565b6047602052600090815260409020546001600160a01b031681565b60466020526000908152604090205481565b6049602052816000526040600020818154811061043a57600080fd5b6000918252602090912001546001600160801b038082169350600160801b90910416905082565b60386020526000908152604090205481565b603660205260009081526040902080546001820154600283015460038401546004909401549293919260ff808316936101008404821693620100008104909216926001600160a01b03630100000090930483169291821691168856fea2646970667358221220b80a0caffa31031756f502a453405903480342ec8dbfe7eb89550ddf05b603c564736f6c63430007040033"

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

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorStorage *TellorStorageCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "_allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorStorage *TellorStorageSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TellorStorage.Contract.Allowances(&_TellorStorage.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorStorage *TellorStorageCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TellorStorage.Contract.Allowances(&_TellorStorage.CallOpts, arg0, arg1)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorStorage *TellorStorageCaller) Addresses(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "addresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorStorage *TellorStorageSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _TellorStorage.Contract.Addresses(&_TellorStorage.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorStorage *TellorStorageCallerSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _TellorStorage.Contract.Addresses(&_TellorStorage.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorStorage *TellorStorageCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "balances", arg0, arg1)

	outstruct := new(struct {
		FromBlock *big.Int
		Value     *big.Int
	})

	outstruct.FromBlock = out[0].(*big.Int)
	outstruct.Value = out[1].(*big.Int)

	return *outstruct, err

}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorStorage *TellorStorageSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _TellorStorage.Contract.Balances(&_TellorStorage.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorStorage *TellorStorageCallerSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _TellorStorage.Contract.Balances(&_TellorStorage.CallOpts, arg0, arg1)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorStorage *TellorStorageCaller) BytesVars(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "bytesVars", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorStorage *TellorStorageSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _TellorStorage.Contract.BytesVars(&_TellorStorage.CallOpts, arg0)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorStorage *TellorStorageCallerSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _TellorStorage.Contract.BytesVars(&_TellorStorage.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorStorage *TellorStorageCaller) CurrentMiners(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "currentMiners", arg0)

	outstruct := new(struct {
		Value *big.Int
		Miner common.Address
	})

	outstruct.Value = out[0].(*big.Int)
	outstruct.Miner = out[1].(common.Address)

	return *outstruct, err

}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorStorage *TellorStorageSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _TellorStorage.Contract.CurrentMiners(&_TellorStorage.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorStorage *TellorStorageCallerSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _TellorStorage.Contract.CurrentMiners(&_TellorStorage.CallOpts, arg0)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorStorage *TellorStorageCaller) DisputeIdByDisputeHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "disputeIdByDisputeHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorStorage *TellorStorageSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorStorage.Contract.DisputeIdByDisputeHash(&_TellorStorage.CallOpts, arg0)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorStorage *TellorStorageCallerSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorStorage.Contract.DisputeIdByDisputeHash(&_TellorStorage.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorStorage *TellorStorageCaller) DisputesById(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "disputesById", arg0)

	outstruct := new(struct {
		Hash                [32]byte
		Tally               *big.Int
		Executed            bool
		DisputeVotePassed   bool
		IsPropFork          bool
		ReportedMiner       common.Address
		ReportingParty      common.Address
		ProposedForkAddress common.Address
	})

	outstruct.Hash = out[0].([32]byte)
	outstruct.Tally = out[1].(*big.Int)
	outstruct.Executed = out[2].(bool)
	outstruct.DisputeVotePassed = out[3].(bool)
	outstruct.IsPropFork = out[4].(bool)
	outstruct.ReportedMiner = out[5].(common.Address)
	outstruct.ReportingParty = out[6].(common.Address)
	outstruct.ProposedForkAddress = out[7].(common.Address)

	return *outstruct, err

}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorStorage *TellorStorageSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _TellorStorage.Contract.DisputesById(&_TellorStorage.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorStorage *TellorStorageCallerSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _TellorStorage.Contract.DisputesById(&_TellorStorage.CallOpts, arg0)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorStorage *TellorStorageCaller) Migrated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "migrated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorStorage *TellorStorageSession) Migrated(arg0 common.Address) (bool, error) {
	return _TellorStorage.Contract.Migrated(&_TellorStorage.CallOpts, arg0)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorStorage *TellorStorageCallerSession) Migrated(arg0 common.Address) (bool, error) {
	return _TellorStorage.Contract.Migrated(&_TellorStorage.CallOpts, arg0)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorStorage *TellorStorageCaller) MinersByChallenge(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "minersByChallenge", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorStorage *TellorStorageSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _TellorStorage.Contract.MinersByChallenge(&_TellorStorage.CallOpts, arg0, arg1)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorStorage *TellorStorageCallerSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _TellorStorage.Contract.MinersByChallenge(&_TellorStorage.CallOpts, arg0, arg1)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorStorage *TellorStorageCaller) NewValueTimestamps(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "newValueTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorStorage *TellorStorageSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _TellorStorage.Contract.NewValueTimestamps(&_TellorStorage.CallOpts, arg0)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorStorage *TellorStorageCallerSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _TellorStorage.Contract.NewValueTimestamps(&_TellorStorage.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorStorage *TellorStorageCaller) RequestIdByQueryHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "requestIdByQueryHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorStorage *TellorStorageSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorStorage.Contract.RequestIdByQueryHash(&_TellorStorage.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorStorage *TellorStorageCallerSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorStorage.Contract.RequestIdByQueryHash(&_TellorStorage.CallOpts, arg0)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorStorage *TellorStorageCaller) Uints(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "uints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorStorage *TellorStorageSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _TellorStorage.Contract.Uints(&_TellorStorage.CallOpts, arg0)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorStorage *TellorStorageCallerSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _TellorStorage.Contract.Uints(&_TellorStorage.CallOpts, arg0)
}

// TellorTransferABI is the input ABI used to generate the binding from.
const TellorTransferABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"fromBlock\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bytesVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentMiners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"disputesById\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"tally\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disputeVotePassed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPropFork\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"reportedMiner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reportingParty\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposedForkAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minersByChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"newValueTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"uints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TellorTransferFuncSigs maps the 4-byte function signature to its string representation.
var TellorTransferFuncSigs = map[string]string{
	"024c2ddd": "_allowances(address,address)",
	"699f200f": "addresses(bytes32)",
	"dd62ed3e": "allowance(address,address)",
	"999cf26c": "allowedToTrade(address,uint256)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"4ee2cd7e": "balanceOfAt(address,uint256)",
	"cbf1304d": "balances(address,uint256)",
	"62dd1d2a": "bytesVars(bytes32)",
	"1fd22364": "currentMiners(uint256)",
	"d01f4d9e": "disputeIdByDisputeHash(bytes32)",
	"db085beb": "disputesById(uint256)",
	"4ba0a5ee": "migrated(address)",
	"48b18e54": "minersByChallenge(bytes32,address)",
	"438c0aa3": "newValueTimestamps(uint256)",
	"5700242c": "requestIdByQueryHash(bytes32)",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"b59e14d4": "uints(bytes32)",
}

// TellorTransferBin is the compiled bytecode used for deploying new contracts.
var TellorTransferBin = "0x608060405234801561001057600080fd5b50610e22806100206000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c806362dd1d2a116100ad578063b59e14d411610071578063b59e14d4146103a6578063cbf1304d146103c3578063d01f4d9e1461041e578063db085beb1461043b578063dd62ed3e146104a857610121565b806362dd1d2a146102d2578063699f200f146102ef57806370a0823114610328578063999cf26c1461034e578063a9059cbb1461037a57610121565b8063438c0aa3116100f4578063438c0aa31461021a57806348b18e54146102375780634ba0a5ee146102635780634ee2cd7e146102895780635700242c146102b557610121565b8063024c2ddd14610126578063095ea7b3146101665780631fd22364146101a657806323b872dd146101e4575b600080fd5b6101546004803603604081101561013c57600080fd5b506001600160a01b03813581169160200135166104d6565b60408051918252519081900360200190f35b6101926004803603604081101561017c57600080fd5b506001600160a01b0381351690602001356104f3565b604080519115158252519081900360200190f35b6101c3600480360360208110156101bc57600080fd5b50356105de565b604080519283526001600160a01b0390911660208301528051918290030190f35b610192600480360360608110156101fa57600080fd5b506001600160a01b03813581169160208101359091169060400135610609565b6101546004803603602081101561023057600080fd5b50356106b5565b6101926004803603604081101561024d57600080fd5b50803590602001356001600160a01b03166106d6565b6101926004803603602081101561027957600080fd5b50356001600160a01b03166106f6565b6101546004803603604081101561029f57600080fd5b506001600160a01b03813516906020013561070b565b610154600480360360208110156102cb57600080fd5b50356108af565b610154600480360360208110156102e857600080fd5b50356108c1565b61030c6004803603602081101561030557600080fd5b50356108d3565b604080516001600160a01b039092168252519081900360200190f35b6101546004803603602081101561033e57600080fd5b50356001600160a01b03166108ee565b6101926004803603604081101561036457600080fd5b506001600160a01b0381351690602001356108fa565b6101926004803603604081101561039057600080fd5b506001600160a01b0381351690602001356109bd565b610154600480360360208110156103bc57600080fd5b50356109d3565b6103ef600480360360408110156103d957600080fd5b506001600160a01b0381351690602001356109e5565b60405180836001600160801b03168152602001826001600160801b031681526020019250505060405180910390f35b6101546004803603602081101561043457600080fd5b5035610a28565b6104586004803603602081101561045157600080fd5b5035610a3a565b60408051988952602089019790975294151587870152921515606087015290151560808601526001600160a01b0390811660a086015290811660c08501521660e083015251908190036101000190f35b610154600480360360408110156104be57600080fd5b506001600160a01b0381358116916020013516610a96565b604a60209081526000928352604080842090915290825290205481565b6000336105315760405162461bcd60e51b8152600401808060200182810382526024815260200180610dc96024913960400191505060405180910390fd5b6001600160a01b0383166105765760405162461bcd60e51b8152600401808060200182810382526022815260200180610d5f6022913960400191505060405180910390fd5b336000818152604a602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060015b92915050565b603a81600581106105ee57600080fd5b6002020180546001909101549091506001600160a01b031682565b6001600160a01b0383166000908152604a60209081526040808320338452909152812054821115610676576040805162461bcd60e51b8152602060048201526012602482015271416c6c6f77616e63652069732077726f6e6760701b604482015290519081900360640190fd5b6001600160a01b0384166000908152604a602090815260408083203384529091529020805483900390556106ab848484610ac1565b5060019392505050565b603381815481106106c557600080fd5b600091825260209091200154905081565b603960209081526000928352604080842090915290825290205460ff1681565b604b6020526000908152604090205460ff1681565b6001600160a01b0382166000908152604960205260408120805415806107515750828160008154811061073a57fe5b6000918252602090912001546001600160801b0316115b156107605760009150506105d8565b80548190600019810190811061077257fe5b6000918252602090912001546001600160801b031683106107c45780548190600019810190811061079f57fe5b600091825260209091200154600160801b90046001600160801b031691506105d89050565b8054600090600119015b8181111561087c5760006002600183850101049050858482815481106107f057fe5b6000918252602090912001546001600160801b0316141561083f5783818154811061081757fe5b600091825260209091200154600160801b90046001600160801b031694506105d89350505050565b8584828154811061084c57fe5b6000918252602090912001546001600160801b0316101561086f57809250610876565b6001810391505b506107ce565b82828154811061088857fe5b600091825260209091200154600160801b90046001600160801b031693506105d892505050565b60376020526000908152604090205481565b60486020526000908152604090205481565b6047602052600090815260409020546001600160a01b031681565b60006105d8824361070b565b6001600160a01b0382166000908152604460205260408120541580159061093957506001600160a01b0383166000908152604460205260409020546005115b156109aa577f5d9fadfc729fd027e395e5157ef1b53ef9fa4a8f053043c5f159307543e7cc9760005260466020527f167af83a0768d27540775cfef6d996eb63f8a61fcdfb26e654c18fb50960e3be548290610994856108ee565b03106109a2575060016105d8565b5060006105d8565b816109b4846108ee565b10159392505050565b60006109ca338484610ac1565b50600192915050565b60466020526000908152604090205481565b60496020528160005260406000208181548110610a0157600080fd5b6000918252602090912001546001600160801b038082169350600160801b90910416905082565b60386020526000908152604090205481565b603660205260009081526040902080546001820154600283015460038401546004909401549293919260ff808316936101008404821693620100008104909216926001600160a01b036301000000909304831692918216911688565b6001600160a01b039182166000908152604a6020908152604080832093909416825291909152205490565b80610afd5760405162461bcd60e51b8152600401808060200182810382526021815260200180610da86021913960400191505060405180910390fd5b6001600160a01b038216610b50576040805162461bcd60e51b815260206004820152601560248201527452656365697665722069732030206164647265737360581b604482015290519081900360640190fd5b610b5a83826108fa565b610b955760405162461bcd60e51b8152600401808060200182810382526027815260200180610d816027913960400191505060405180910390fd5b6000610ba0846108ee565b9050610bae84838303610c61565b610bb7836108ee565b9050808282011015610c04576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b610c1083838301610c61565b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a350505050565b6001600160a01b038216600090815260496020526040902080541580610cae57508054439082906000198101908110610c9657fe5b6000918252602090912001546001600160801b031614155b15610d1f5760408051808201909152436001600160801b0390811682528381166020808401918252845460018101865560008681529190912093519301805491516fffffffffffffffffffffffffffffffff19909216938316939093178216600160801b9190921602179055610d59565b805460009082906000198101908110610d3457fe5b600091825260209091200180546001600160801b03808616600160801b029116179055505b50505056fe45524332303a20617070726f766520746f20746865207a65726f206164647265737353686f756c6420686176652073756666696369656e742062616c616e636520746f207472616465547269656420746f2073656e64206e6f6e2d706f73697469766520616d6f756e7445524332303a20617070726f76652066726f6d20746865207a65726f2061646472657373a2646970667358221220e4f5490a77b23b69004f59e4ffe6dd680c1936b60eaf476260466b2856fc009464736f6c63430007040033"

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

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorTransfer *TellorTransferCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "_allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorTransfer *TellorTransferSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TellorTransfer.Contract.Allowances(&_TellorTransfer.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorTransfer *TellorTransferCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TellorTransfer.Contract.Allowances(&_TellorTransfer.CallOpts, arg0, arg1)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorTransfer *TellorTransferCaller) Addresses(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "addresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorTransfer *TellorTransferSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _TellorTransfer.Contract.Addresses(&_TellorTransfer.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorTransfer *TellorTransferCallerSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _TellorTransfer.Contract.Addresses(&_TellorTransfer.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_TellorTransfer *TellorTransferCaller) Allowance(opts *bind.CallOpts, _user common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "allowance", _user, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_TellorTransfer *TellorTransferSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _TellorTransfer.Contract.Allowance(&_TellorTransfer.CallOpts, _user, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_TellorTransfer *TellorTransferCallerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _TellorTransfer.Contract.Allowance(&_TellorTransfer.CallOpts, _user, _spender)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_TellorTransfer *TellorTransferCaller) AllowedToTrade(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "allowedToTrade", _user, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_TellorTransfer *TellorTransferSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _TellorTransfer.Contract.AllowedToTrade(&_TellorTransfer.CallOpts, _user, _amount)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_TellorTransfer *TellorTransferCallerSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _TellorTransfer.Contract.AllowedToTrade(&_TellorTransfer.CallOpts, _user, _amount)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_TellorTransfer *TellorTransferCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "balanceOf", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_TellorTransfer *TellorTransferSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _TellorTransfer.Contract.BalanceOf(&_TellorTransfer.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_TellorTransfer *TellorTransferCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _TellorTransfer.Contract.BalanceOf(&_TellorTransfer.CallOpts, _user)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_TellorTransfer *TellorTransferCaller) BalanceOfAt(opts *bind.CallOpts, _user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "balanceOfAt", _user, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_TellorTransfer *TellorTransferSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _TellorTransfer.Contract.BalanceOfAt(&_TellorTransfer.CallOpts, _user, _blockNumber)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_TellorTransfer *TellorTransferCallerSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _TellorTransfer.Contract.BalanceOfAt(&_TellorTransfer.CallOpts, _user, _blockNumber)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorTransfer *TellorTransferCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "balances", arg0, arg1)

	outstruct := new(struct {
		FromBlock *big.Int
		Value     *big.Int
	})

	outstruct.FromBlock = out[0].(*big.Int)
	outstruct.Value = out[1].(*big.Int)

	return *outstruct, err

}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorTransfer *TellorTransferSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _TellorTransfer.Contract.Balances(&_TellorTransfer.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorTransfer *TellorTransferCallerSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _TellorTransfer.Contract.Balances(&_TellorTransfer.CallOpts, arg0, arg1)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) BytesVars(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "bytesVars", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _TellorTransfer.Contract.BytesVars(&_TellorTransfer.CallOpts, arg0)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _TellorTransfer.Contract.BytesVars(&_TellorTransfer.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorTransfer *TellorTransferCaller) CurrentMiners(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "currentMiners", arg0)

	outstruct := new(struct {
		Value *big.Int
		Miner common.Address
	})

	outstruct.Value = out[0].(*big.Int)
	outstruct.Miner = out[1].(common.Address)

	return *outstruct, err

}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorTransfer *TellorTransferSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _TellorTransfer.Contract.CurrentMiners(&_TellorTransfer.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorTransfer *TellorTransferCallerSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _TellorTransfer.Contract.CurrentMiners(&_TellorTransfer.CallOpts, arg0)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorTransfer *TellorTransferCaller) DisputeIdByDisputeHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "disputeIdByDisputeHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorTransfer *TellorTransferSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorTransfer.Contract.DisputeIdByDisputeHash(&_TellorTransfer.CallOpts, arg0)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorTransfer *TellorTransferCallerSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorTransfer.Contract.DisputeIdByDisputeHash(&_TellorTransfer.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorTransfer *TellorTransferCaller) DisputesById(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "disputesById", arg0)

	outstruct := new(struct {
		Hash                [32]byte
		Tally               *big.Int
		Executed            bool
		DisputeVotePassed   bool
		IsPropFork          bool
		ReportedMiner       common.Address
		ReportingParty      common.Address
		ProposedForkAddress common.Address
	})

	outstruct.Hash = out[0].([32]byte)
	outstruct.Tally = out[1].(*big.Int)
	outstruct.Executed = out[2].(bool)
	outstruct.DisputeVotePassed = out[3].(bool)
	outstruct.IsPropFork = out[4].(bool)
	outstruct.ReportedMiner = out[5].(common.Address)
	outstruct.ReportingParty = out[6].(common.Address)
	outstruct.ProposedForkAddress = out[7].(common.Address)

	return *outstruct, err

}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorTransfer *TellorTransferSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _TellorTransfer.Contract.DisputesById(&_TellorTransfer.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorTransfer *TellorTransferCallerSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _TellorTransfer.Contract.DisputesById(&_TellorTransfer.CallOpts, arg0)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorTransfer *TellorTransferCaller) Migrated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "migrated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorTransfer *TellorTransferSession) Migrated(arg0 common.Address) (bool, error) {
	return _TellorTransfer.Contract.Migrated(&_TellorTransfer.CallOpts, arg0)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorTransfer *TellorTransferCallerSession) Migrated(arg0 common.Address) (bool, error) {
	return _TellorTransfer.Contract.Migrated(&_TellorTransfer.CallOpts, arg0)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorTransfer *TellorTransferCaller) MinersByChallenge(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "minersByChallenge", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorTransfer *TellorTransferSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _TellorTransfer.Contract.MinersByChallenge(&_TellorTransfer.CallOpts, arg0, arg1)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorTransfer *TellorTransferCallerSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _TellorTransfer.Contract.MinersByChallenge(&_TellorTransfer.CallOpts, arg0, arg1)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorTransfer *TellorTransferCaller) NewValueTimestamps(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "newValueTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorTransfer *TellorTransferSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _TellorTransfer.Contract.NewValueTimestamps(&_TellorTransfer.CallOpts, arg0)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorTransfer *TellorTransferCallerSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _TellorTransfer.Contract.NewValueTimestamps(&_TellorTransfer.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorTransfer *TellorTransferCaller) RequestIdByQueryHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "requestIdByQueryHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorTransfer *TellorTransferSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorTransfer.Contract.RequestIdByQueryHash(&_TellorTransfer.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorTransfer *TellorTransferCallerSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorTransfer.Contract.RequestIdByQueryHash(&_TellorTransfer.CallOpts, arg0)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorTransfer *TellorTransferCaller) Uints(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "uints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorTransfer *TellorTransferSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _TellorTransfer.Contract.Uints(&_TellorTransfer.CallOpts, arg0)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorTransfer *TellorTransferCallerSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _TellorTransfer.Contract.Uints(&_TellorTransfer.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_TellorTransfer *TellorTransferTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_TellorTransfer *TellorTransferSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.Contract.Approve(&_TellorTransfer.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_TellorTransfer *TellorTransferTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.Contract.Approve(&_TellorTransfer.TransactOpts, _spender, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_TellorTransfer *TellorTransferTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_TellorTransfer *TellorTransferSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.Contract.Transfer(&_TellorTransfer.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_TellorTransfer *TellorTransferTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.Contract.Transfer(&_TellorTransfer.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_TellorTransfer *TellorTransferTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_TellorTransfer *TellorTransferSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.Contract.TransferFrom(&_TellorTransfer.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_TellorTransfer *TellorTransferTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.Contract.TransferFrom(&_TellorTransfer.TransactOpts, _from, _to, _amount)
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TellorTransfer *TellorTransferFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TellorTransferApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TellorTransfer.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TellorTransferApprovalIterator{contract: _TellorTransfer.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TellorTransfer *TellorTransferFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TellorTransferApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TellorTransfer.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TellorTransfer *TellorTransferFilterer) ParseApproval(log types.Log) (*TellorTransferApproval, error) {
	event := new(TellorTransferApproval)
	if err := _TellorTransfer.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorTransferTransferredIterator is returned from FilterTransferred and is used to iterate over the raw logs and unpacked data for Transferred events raised by the TellorTransfer contract.
type TellorTransferTransferredIterator struct {
	Event *TellorTransferTransferred // Event containing the contract specifics and raw log

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
func (it *TellorTransferTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorTransferTransferred)
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
		it.Event = new(TellorTransferTransferred)
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
func (it *TellorTransferTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorTransferTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorTransferTransferred represents a Transferred event raised by the TellorTransfer contract.
type TellorTransferTransferred struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferred is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TellorTransfer *TellorTransferFilterer) FilterTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TellorTransferTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TellorTransfer.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TellorTransferTransferredIterator{contract: _TellorTransfer.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransferred is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TellorTransfer *TellorTransferFilterer) WatchTransferred(opts *bind.WatchOpts, sink chan<- *TellorTransferTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TellorTransfer.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorTransferTransferred)
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

// ParseTransferred is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TellorTransfer *TellorTransferFilterer) ParseTransferred(log types.Log) (*TellorTransferTransferred, error) {
	event := new(TellorTransferTransferred)
	if err := _TellorTransfer.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorVariablesABI is the input ABI used to generate the binding from.
const TellorVariablesABI = "[]"

// TellorVariablesBin is the compiled bytecode used for deploying new contracts.
var TellorVariablesBin = "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea264697066735822122030ef1f518303ccc73289bd89fcd18a1becd99e201b6cdbc8c5378001507f7a5764736f6c63430007040033"

// DeployTellorVariables deploys a new Ethereum contract, binding an instance of TellorVariables to it.
func DeployTellorVariables(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorVariables, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorVariablesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TellorVariablesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TellorVariables{TellorVariablesCaller: TellorVariablesCaller{contract: contract}, TellorVariablesTransactor: TellorVariablesTransactor{contract: contract}, TellorVariablesFilterer: TellorVariablesFilterer{contract: contract}}, nil
}

// TellorVariables is an auto generated Go binding around an Ethereum contract.
type TellorVariables struct {
	TellorVariablesCaller     // Read-only binding to the contract
	TellorVariablesTransactor // Write-only binding to the contract
	TellorVariablesFilterer   // Log filterer for contract events
}

// TellorVariablesCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorVariablesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorVariablesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorVariablesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorVariablesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorVariablesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorVariablesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorVariablesSession struct {
	Contract     *TellorVariables  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorVariablesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorVariablesCallerSession struct {
	Contract *TellorVariablesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TellorVariablesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorVariablesTransactorSession struct {
	Contract     *TellorVariablesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TellorVariablesRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorVariablesRaw struct {
	Contract *TellorVariables // Generic contract binding to access the raw methods on
}

// TellorVariablesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorVariablesCallerRaw struct {
	Contract *TellorVariablesCaller // Generic read-only contract binding to access the raw methods on
}

// TellorVariablesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorVariablesTransactorRaw struct {
	Contract *TellorVariablesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellorVariables creates a new instance of TellorVariables, bound to a specific deployed contract.
func NewTellorVariables(address common.Address, backend bind.ContractBackend) (*TellorVariables, error) {
	contract, err := bindTellorVariables(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TellorVariables{TellorVariablesCaller: TellorVariablesCaller{contract: contract}, TellorVariablesTransactor: TellorVariablesTransactor{contract: contract}, TellorVariablesFilterer: TellorVariablesFilterer{contract: contract}}, nil
}

// NewTellorVariablesCaller creates a new read-only instance of TellorVariables, bound to a specific deployed contract.
func NewTellorVariablesCaller(address common.Address, caller bind.ContractCaller) (*TellorVariablesCaller, error) {
	contract, err := bindTellorVariables(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorVariablesCaller{contract: contract}, nil
}

// NewTellorVariablesTransactor creates a new write-only instance of TellorVariables, bound to a specific deployed contract.
func NewTellorVariablesTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorVariablesTransactor, error) {
	contract, err := bindTellorVariables(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorVariablesTransactor{contract: contract}, nil
}

// NewTellorVariablesFilterer creates a new log filterer instance of TellorVariables, bound to a specific deployed contract.
func NewTellorVariablesFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorVariablesFilterer, error) {
	contract, err := bindTellorVariables(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorVariablesFilterer{contract: contract}, nil
}

// bindTellorVariables binds a generic wrapper to an already deployed contract.
func bindTellorVariables(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorVariablesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorVariables *TellorVariablesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorVariables.Contract.TellorVariablesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorVariables *TellorVariablesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorVariables.Contract.TellorVariablesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorVariables *TellorVariablesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorVariables.Contract.TellorVariablesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorVariables *TellorVariablesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorVariables.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorVariables *TellorVariablesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorVariables.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorVariables *TellorVariablesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorVariables.Contract.contract.Transact(opts, method, params...)
}

// UtilitiesABI is the input ABI used to generate the binding from.
const UtilitiesABI = "[{\"inputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"data\",\"type\":\"uint256[51]\"}],\"name\":\"getMax5\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"max\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"maxIndex\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// UtilitiesFuncSigs maps the 4-byte function signature to its string representation.
var UtilitiesFuncSigs = map[string]string{
	"99830e32": "getMax5(uint256[51])",
}

// UtilitiesBin is the compiled bytecode used for deploying new contracts.
var UtilitiesBin = "0x608060405234801561001057600080fd5b50610290806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806399830e3214610030575b600080fd5b610084600480360361066081101561004757600080fd5b8101908080610660019060338060200260405190810160405280929190826033602002808284376000920191909152509194506100e89350505050565b604051808360a080838360005b838110156100a9578181015183820152602001610091565b5050505090500182600560200280838360005b838110156100d45781810151838201526020016100bc565b505050509050019250505060405180910390f35b6100f061023c565b6100f861023c565b60208301516000805b600581101561017b5785816001016033811061011957fe5b602002015185826005811061012a57fe5b60200201526001810184826005811061013f57fe5b60200201528285826005811061015157fe5b602002015110156101735784816005811061016857fe5b602002015192508091505b600101610101565b5060065b6033811015610234578286826033811061019557fe5b6020020151111561022c578581603381106101ac57fe5b60200201518583600581106101bd57fe5b6020020152808483600581106101cf57fe5b60200201528581603381106101e057fe5b6020020151925060005b600581101561022a578386826005811061020057fe5b602002015110156102225785816005811061021757fe5b602002015193508092505b6001016101ea565b505b60010161017f565b505050915091565b6040518060a00160405280600590602082028036833750919291505056fea2646970667358221220ef7458ffe05c014369c7bdf479fc93992b9dac499fe62cf34e380201e2808a7164736f6c63430007040033"

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

// GetMax5 is a free data retrieval call binding the contract method 0x99830e32.
//
// Solidity: function getMax5(uint256[51] data) view returns(uint256[5] max, uint256[5] maxIndex)
func (_Utilities *UtilitiesCaller) GetMax5(opts *bind.CallOpts, data [51]*big.Int) (struct {
	Max      [5]*big.Int
	MaxIndex [5]*big.Int
}, error) {
	var out []interface{}
	err := _Utilities.contract.Call(opts, &out, "getMax5", data)

	outstruct := new(struct {
		Max      [5]*big.Int
		MaxIndex [5]*big.Int
	})

	outstruct.Max = out[0].([5]*big.Int)
	outstruct.MaxIndex = out[1].([5]*big.Int)

	return *outstruct, err

}

// GetMax5 is a free data retrieval call binding the contract method 0x99830e32.
//
// Solidity: function getMax5(uint256[51] data) view returns(uint256[5] max, uint256[5] maxIndex)
func (_Utilities *UtilitiesSession) GetMax5(data [51]*big.Int) (struct {
	Max      [5]*big.Int
	MaxIndex [5]*big.Int
}, error) {
	return _Utilities.Contract.GetMax5(&_Utilities.CallOpts, data)
}

// GetMax5 is a free data retrieval call binding the contract method 0x99830e32.
//
// Solidity: function getMax5(uint256[51] data) view returns(uint256[5] max, uint256[5] maxIndex)
func (_Utilities *UtilitiesCallerSession) GetMax5(data [51]*big.Int) (struct {
	Max      [5]*big.Int
	MaxIndex [5]*big.Int
}, error) {
	return _Utilities.Contract.GetMax5(&_Utilities.CallOpts, data)
}
