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

// ITellorABI is the input ABI used to generate the binding from.
const ITellorABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"_result\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reportedMiner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_reportingParty\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"DisputeVoteTallied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_currentRequestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"NewChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"NewDispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"NewStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newTellor\",\"type\":\"address\"}],\"name\":\"NewTellorAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NewValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NonceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"TipAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_position\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_voteWeight\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minerIndex\",\"type\":\"uint256\"}],\"name\":\"beginDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newDeity\",\"type\":\"address\"}],\"name\":\"changeDeity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tellorContract\",\"type\":\"address\"}],\"name\":\"changeTellorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[9]\",\"name\":\"\",\"type\":\"uint256[9]\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastNewValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"data\",\"type\":\"uint256[51]\"}],\"name\":\"getMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"data\",\"type\":\"uint256[51]\"}],\"name\":\"getMax5\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"max\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"maxIndex\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"data\",\"type\":\"uint256[51]\"}],\"name\":\"getMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"internalType\":\"address[5]\",\"name\":\"\",\"type\":\"address[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256\",\"name\":\"_difficutly\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"idsOnDeck\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"tipsOnDeck\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_request\",\"type\":\"bytes32\"}],\"name\":\"getRequestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"\",\"type\":\"uint256[51]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTopRequestIDs\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_diff\",\"type\":\"uint256\"}],\"name\":\"manuallySetDifficulty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_propNewTellorAddress\",\"type\":\"address\"}],\"name\":\"proposeFork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_pendingOwner\",\"type\":\"address\"}],\"name\":\"proposeOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestStakingWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"}],\"name\":\"submitMiningSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"tallyVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"}],\"name\":\"testSubmitMiningSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"theLazyCoon\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"unlockDisputeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"updateTellor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_supportsDispute\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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
	"ae0a8279": "changeTellorContract(address)",
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

// ITellorTransferEventIterator is returned from FilterTransferEvent and is used to iterate over the raw logs and unpacked data for TransferEvent events raised by the ITellor contract.
type ITellorTransferEventIterator struct {
	Event *ITellorTransferEvent // Event containing the contract specifics and raw log

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
func (it *ITellorTransferEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorTransferEvent)
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
		it.Event = new(ITellorTransferEvent)
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
func (it *ITellorTransferEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorTransferEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorTransferEvent represents a TransferEvent event raised by the ITellor contract.
type ITellorTransferEvent struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferEvent is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_ITellor *ITellorFilterer) FilterTransferEvent(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ITellorTransferEventIterator, error) {

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
	return &ITellorTransferEventIterator{contract: _ITellor.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransferEvent is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_ITellor *ITellorFilterer) WatchTransferEvent(opts *bind.WatchOpts, sink chan<- *ITellorTransferEvent, _from []common.Address, _to []common.Address) (event.Subscription, error) {

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
				event := new(ITellorTransferEvent)
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

// ParseTransferEvent is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_ITellor *ITellorFilterer) ParseTransferEvent(log types.Log) (*ITellorTransferEvent, error) {
	event := new(ITellorTransferEvent)
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
const TellorABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_currentRequestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"NewChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NewValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_slot\",\"type\":\"uint256\"}],\"name\":\"NonceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"TipAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_tBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bytesVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentChallenge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRequestId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentReward\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTotalTips\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"difficulty\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disputeFee\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pending_owner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestCount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestQPosition\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slotProgress\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeAmount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"_values\",\"type\":\"uint256[5]\"}],\"name\":\"submitMiningSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetMiners\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeOfLastNewValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeTarget\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTip\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"total_supply\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"uints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"updateBalanceAtNow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TellorFuncSigs maps the 4-byte function signature to its string representation.
var TellorFuncSigs = map[string]string{
	"024c2ddd": "_allowances(address,address)",
	"b2bdfa7b": "_owner()",
	"6e3cf885": "_tBlock()",
	"752d49a1": "addTip(uint256,uint256)",
	"699f200f": "addresses(bytes32)",
	"dd62ed3e": "allowance(address,address)",
	"999cf26c": "allowedToTrade(address,uint256)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"4ee2cd7e": "balanceOfAt(address,uint256)",
	"62dd1d2a": "bytesVars(bytes32)",
	"51bdd585": "currentChallenge()",
	"5ae2bfdb": "currentRequestId()",
	"07621eca": "currentReward()",
	"75ad1a2a": "currentTotalTips()",
	"19cae462": "difficulty()",
	"b9ce896b": "disputeFee()",
	"8fd3ab80": "migrate()",
	"4ba0a5ee": "migrated(address)",
	"7f4ec4c3": "pending_owner()",
	"5badbe4c": "requestCount()",
	"2bf07e9e": "requestQPosition()",
	"03b3160f": "slotProgress()",
	"60c7dc47": "stakeAmount()",
	"dff69787": "stakerCount()",
	"4350283e": "submitMiningSolution(string,uint256[5],uint256[5])",
	"dfee1ff1": "targetMiners()",
	"6fd4f229": "timeOfLastNewValue()",
	"6fc37811": "timeTarget()",
	"561cb04a": "totalTip()",
	"3940e9ee": "total_supply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"b59e14d4": "uints(bytes32)",
	"d67dcbc5": "updateBalanceAtNow(address,uint256)",
}

// TellorBin is the compiled bytecode used for deploying new contracts.
var TellorBin = "0x608060405234801561001057600080fd5b50613536806100206000396000f3fe6080604052600436106102045760003560e01c8063699f200f11610118578063999cf26c116100a0578063b9ce896b1161006f578063b9ce896b1461071d578063d67dcbc514610732578063dd62ed3e1461076b578063dfee1ff1146107a6578063dff69787146107bb57610204565b8063999cf26c1461066c578063a9059cbb146106a5578063b2bdfa7b146106de578063b59e14d4146106f357610204565b806370a08231116100e757806370a08231146105ca578063752d49a1146105fd57806375ad1a2a1461062d5780637f4ec4c3146106425780638fd3ab801461065757610204565b8063699f200f146105455780636e3cf8851461058b5780636fc37811146105a05780636fd4f229146105b557610204565b80634350283e1161019b578063561cb04a1161016a578063561cb04a146104c75780635ae2bfdb146104dc5780635badbe4c146104f157806360c7dc471461050657806362dd1d2a1461051b57610204565b80634350283e146103c35780634ba0a5ee146104465780634ee2cd7e1461047957806351bdd585146104b257610204565b806319cae462116101d757806319cae4621461034157806323b872dd146103565780632bf07e9e146103995780633940e9ee146103ae57610204565b8063024c2ddd1461027d57806303b3160f146102ca57806307621eca146102df578063095ea7b3146102f4575b7f363c1ee4df63e032778cbd72a9384a4e47c584b4eea56f22e4580fced3369273600090815260476020527f029bf55bc11f592fd6bd7e7b2be6c7b5f8270a696e61b09d01535f4f1c0922ef546001600160a01b031690610264826107d0565b5090503d6000803e808015610278573d6000f35b3d6000fd5b34801561028957600080fd5b506102b8600480360360408110156102a057600080fd5b506001600160a01b0381358116916020013516610840565b60408051918252519081900360200190f35b3480156102d657600080fd5b506102b861085d565b3480156102eb57600080fd5b506102b861086f565b34801561030057600080fd5b5061032d6004803603604081101561031757600080fd5b506001600160a01b038135169060200135610893565b604080519115158252519081900360200190f35b34801561034d57600080fd5b506102b861097e565b34801561036257600080fd5b5061032d6004803603606081101561037957600080fd5b506001600160a01b038135811691602081013590911690604001356109a2565b3480156103a557600080fd5b506102b8610a4e565b3480156103ba57600080fd5b506102b8610a60565b3480156103cf57600080fd5b5061044460048036036101608110156103e757600080fd5b81019060208101813564010000000081111561040257600080fd5b82018360208201111561041457600080fd5b8035906020019184600183028401116401000000008311171561043657600080fd5b919350915060a08101610a72565b005b34801561045257600080fd5b5061032d6004803603602081101561046957600080fd5b50356001600160a01b0316610c0d565b34801561048557600080fd5b506102b86004803603604081101561049c57600080fd5b506001600160a01b038135169060200135610c22565b3480156104be57600080fd5b506102b8610dc6565b3480156104d357600080fd5b506102b8610dea565b3480156104e857600080fd5b506102b8610dfc565b3480156104fd57600080fd5b506102b8610e20565b34801561051257600080fd5b506102b8610e44565b34801561052757600080fd5b506102b86004803603602081101561053e57600080fd5b5035610e68565b34801561055157600080fd5b5061056f6004803603602081101561056857600080fd5b5035610e7a565b604080516001600160a01b039092168252519081900360200190f35b34801561059757600080fd5b506102b8610e95565b3480156105ac57600080fd5b506102b8610eb9565b3480156105c157600080fd5b506102b8610edd565b3480156105d657600080fd5b506102b8600480360360208110156105ed57600080fd5b50356001600160a01b0316610eef565b34801561060957600080fd5b506104446004803603604081101561062057600080fd5b5080359060200135610efb565b34801561063957600080fd5b506102b861110b565b34801561064e57600080fd5b506102b861111d565b34801561066357600080fd5b50610444611141565b34801561067857600080fd5b5061032d6004803603604081101561068f57600080fd5b506001600160a01b03813516906020013561114c565b3480156106b157600080fd5b5061032d600480360360408110156106c857600080fd5b506001600160a01b03813516906020013561120f565b3480156106ea57600080fd5b506102b8611225565b3480156106ff57600080fd5b506102b86004803603602081101561071657600080fd5b5035611249565b34801561072957600080fd5b506102b861125b565b34801561073e57600080fd5b506104446004803603604081101561075557600080fd5b506001600160a01b03813516906020013561127f565b34801561077757600080fd5b506102b86004803603604081101561078e57600080fd5b506001600160a01b038135811691602001351661137c565b3480156107b257600080fd5b506102b86113a7565b3480156107c757600080fd5b506102b86113cb565b60006060826001600160a01b03166000366040518083838082843760405192019450600093509091505080830381855af49150503d8060008114610830576040519150601f19603f3d011682016040523d82523d6000602084013e610835565b606091505b509094909350915050565b604a60209081526000928352604080842090915290825290205481565b6000805160206133b683398151915281565b7f9b6853911475b07474368644a0d922ee13bc76a15cd3e97d3e334326424a47d481565b6000336108d15760405162461bcd60e51b81526004018080602001828103825260248152602001806134586024913960400191505060405180910390fd5b6001600160a01b0383166109165760405162461bcd60e51b815260040180806020018281038252602281526020018061330c6022913960400191505060405180910390fd5b336000818152604a602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060015b92915050565b7fb12aff7664b16cb99339be399b863feecd64d14817be7e1f042f97e3f358e64e81565b6001600160a01b0383166000908152604a60209081526040808320338452909152812054821115610a0f576040805162461bcd60e51b8152602060048201526012602482015271416c6c6f77616e63652069732077726f6e6760701b604482015290519081900360640190fd5b6001600160a01b0384166000908152604a60209081526040808320338452909152902080548390039055610a448484846113ef565b5060019392505050565b6000805160206133f783398151915281565b6000805160206132cc83398151915281565b6040805133602080830191909152825180830382018152918301835281519181019190912060008181526046909252919020541580610ac557506000818152604660205260409020546103844291909103115b610b005760405162461bcd60e51b815260040180806020018281038252602a815260200180613282602a913960400191505060405180910390fd5b6000805160206133b683398151915260005260466020526000805160206132ec83398151915254600414610b6d57610b6d85858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061158f92505050565b6000818152604660209081526040918290204290558151601f8701829004820281018201909252858252610c0691908790879081908401838280828437600092019190915250506040805160a081810190925292508791506005908390839080828437600092019190915250506040805160a0818101909252915086906005908390839080828437600092019190915250611842915050565b5050505050565b604b6020526000908152604090205460ff1681565b6001600160a01b038216600090815260496020526040812080541580610c6857508281600081548110610c5157fe5b6000918252602090912001546001600160801b0316115b15610c77576000915050610978565b805481906000198101908110610c8957fe5b6000918252602090912001546001600160801b03168310610cdb57805481906000198101908110610cb657fe5b600091825260209091200154600160801b90046001600160801b031691506109789050565b8054600090600119015b81811115610d93576000600260018385010104905085848281548110610d0757fe5b6000918252602090912001546001600160801b03161415610d5657838181548110610d2e57fe5b600091825260209091200154600160801b90046001600160801b031694506109789350505050565b85848281548110610d6357fe5b6000918252602090912001546001600160801b03161015610d8657809250610d8d565b6001810391505b50610ce5565b828281548110610d9f57fe5b600091825260209091200154600160801b90046001600160801b0316935061097892505050565b7f3375fb9157bb77048f607329b1c4d45487433f61e0f51bcdeba91d975b2dab1881565b60008051602061339683398151915281565b7f7584d7d8701714da9c117f5bf30af73b0b88aca5338a84a21eb28de2fe0d93b881565b7f05de9147d05477c0a5dc675aeea733157f5092f82add148cf39d579cafe3dc9881565b7f7be108969d31a3f0b261465c71f2b0ba9301cd914d55d9091c3b36a49d4d41b281565b60486020526000908152604090205481565b6047602052600090815260409020546001600160a01b031681565b7f969ea04b74d02bb4d9e6e8e57236e1b9ca31627139ae9f0e465249932e82450281565b7fad16221efc80aaf1b7e69bd3ecb61ba5ffa539adf129c3b4ffff769c9b5bbc3381565b6000805160206132ac83398151915281565b60006109788243610c22565b81610f3e576040805162461bcd60e51b815260206004820152600e60248201526d052657175657374496420697320360941b604482015290519081900360640190fd5b80610f90576040805162461bcd60e51b815260206004820152601c60248201527f5469702073686f756c642062652067726561746572207468616e203000000000604482015290519081900360640190fd5b7f05de9147d05477c0a5dc675aeea733157f5092f82add148cf39d579cafe3dc9860005260466020527f545247270828a5b9d6a89772258005e739bcc32b5fda815a7365cf0ab78b09555460010182811415611038577f05de9147d05477c0a5dc675aeea733157f5092f82add148cf39d579cafe3dc9860005260466020527f545247270828a5b9d6a89772258005e739bcc32b5fda815a7365cf0ab78b095581905561108c565b80831061108c576040805162461bcd60e51b815260206004820181905260248201527f526571756573744964206973206e6f74206c657373207468616e20636f756e74604482015290519081900360640190fd5b6110963383611ea7565b6110a08383611fbb565b60008381526045602090815260408083206000805160206133968339815191528452600101825291829020548251858152918201528151859233927fd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820929081900390910190a3505050565b6000805160206134c183398151915281565b7f44b2657a0f8a90ed8e62f4c4cceca06eacaa9b4b25751ae1ebca9280a70abd6881565b61114a336121c9565b565b6001600160a01b0382166000908152604460205260408120541580159061118b57506001600160a01b0383166000908152604460205260409020546005115b156111fc577f7be108969d31a3f0b261465c71f2b0ba9301cd914d55d9091c3b36a49d4d41b260005260466020527ff0cbb0eecd83f344371d44a7d44097662e277de50a2e30b5e36df1aa56be6de55482906111e685610eef565b03106111f457506001610978565b506000610978565b8161120684610eef565b10159392505050565b600061121c3384846113ef565b50600192915050565b7f9dbc393ddc18fd27b1d9b1b129059925688d2f2d5818a5ec3ebb750b7c286ea681565b60466020526000908152604090205481565b7f8b75eb45d88e80f0e4ec77d23936268694c0e7ac2e0c9085c5c6bdfcfbc4923981565b6001600160a01b0382166000908152604960205260409020805415806112cc575080544390829060001981019081106112b457fe5b6000918252602090912001546001600160801b031614155b1561133d5760408051808201909152436001600160801b0390811682528381166020808401918252845460018101865560008681529190912093519301805491516fffffffffffffffffffffffffffffffff19909216938316939093178216600160801b9190921602179055611377565b80546000908290600019810190811061135257fe5b600091825260209091200180546001600160801b03808616600160801b029116179055505b505050565b6001600160a01b039182166000908152604a6020908152604080832093909416825291909152205490565b7fabef544d8048318ece54fb2c6385255cd1b06e176525d149a0338a7acca6deb381565b7fedddb9344bfe0dadc78c558b8ffca446679cbffc17be64eb83973fce7bea5f3481565b8061142b5760405162461bcd60e51b81526004018080602001828103825260218152602001806133d66021913960400191505060405180910390fd5b6001600160a01b03821661147e576040805162461bcd60e51b815260206004820152601560248201527452656365697665722069732030206164647265737360581b604482015290519081900360640190fd5b611488838261114c565b6114c35760405162461bcd60e51b815260040180806020018281038252602781526020018061336f6027913960400191505060405180910390fd5b60006114ce84610eef565b90506114dc8483830361127f565b6114e583610eef565b9050808282011015611532576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b61153e8383830161127f565b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a350505050565b6000805160206134e1833981519152547f3375fb9157bb77048f607329b1c4d45487433f61e0f51bcdeba91d975b2dab186000526048602090815260008051602061332e833981519152546040805180840183815233606081901b93830193909352865160029560039594938993926054909101918401908083835b6020831061162a5780518252601f19909201916020918201910161160b565b6001836020036101000a038019825116818451168082178552505050505050905001935050505060405160208183030381529060405280519060200120604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b602083106116b55780518252601f199092019160209182019101611696565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa1580156116f4573d6000803e3d6000fd5b5050506040515160601b60405160200180826bffffffffffffffffffffffff191681526014019150506040516020818303038152906040526040518082805190602001908083835b6020831061175b5780518252601f19909201916020918201910161173c565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561179a573d6000803e3d6000fd5b5050506040513d60208110156117af57600080fd5b5051816117b857fe5b06158061180457506000805160206132ac83398151915260005260466020527f65764a602ca90a32a17ee52f54b96d452a879c4a2d747d2d464139cfbf12a3f754610384429190910310155b61183f5760405162461bcd60e51b815260040180806020018281038252602581526020018061349c6025913960400191505060405180910390fd5b50565b6040805133602080830182905283518084038201815292840184528251928101929092206000918252604490925291909120546001146118c9576040805162461bcd60e51b815260206004820152601a60248201527f4d696e657220737461747573206973206e6f74207374616b6572000000000000604482015290519081900360640190fd5b603a54835114611916576040805162461bcd60e51b8152602060048201526013602482015272526571756573742049442069732077726f6e6760681b604482015290519081900360640190fd5b603c54602084015114611966576040805162461bcd60e51b8152602060048201526013602482015272526571756573742049442069732077726f6e6760681b604482015290519081900360640190fd5b603e546040840151146119b6576040805162461bcd60e51b8152602060048201526013602482015272526571756573742049442069732077726f6e6760681b604482015290519081900360640190fd5b604054606084015114611a06576040805162461bcd60e51b8152602060048201526013602482015272526571756573742049442069732077726f6e6760681b604482015290519081900360640190fd5b604254608084015114611a56576040805162461bcd60e51b8152602060048201526013602482015272526571756573742049442069732077726f6e6760681b604482015290519081900360640190fd5b600081815260466020908152604080832042905560008051602061332e833981519152546000805160206132ec833981519152548185526039845282852033865290935292205460ff1615611adc5760405162461bcd60e51b81526004018080602001828103825260218152602001806134376021913960400191505060405180910390fd5b60008281526039602090815260408083203384528252808320805460ff191660011790557fd938684f43b130e7cfecbebf76ff7e28b74288b2b871d881041ace3b450910ec5483526045825280832087518480526006820190935292208360058110611b4457fe5b0155602080860151600160009081526006840190925260409091208360058110611b6a57fe5b01556040808601516002600090815260068401602052919091208360058110611b8f57fe5b015560608501516003600090815260068301602052604090208360058110611bb357fe5b015560808501516004600090815260068301602052604090208360058110611bd757fe5b0155600080805260058083016020526040909120339184908110611bf757fe5b0180546001600160a01b0319166001600160a01b03929092169190911790556001600090815260058281016020526040909120339184908110611c3657fe5b0180546001600160a01b0319166001600160a01b03929092169190911790556002600090815260058281016020526040909120339184908110611c7557fe5b0180546001600160a01b0319166001600160a01b03929092169190911790556003600090815260058281016020526040909120339184908110611cb457fe5b0180546001600160a01b0319166001600160a01b03929092169190911790556004600090815260058281016020526040909120339184908110611cf357fe5b0180546001600160a01b0319166001600160a01b03929092169190911790556001820160041415611d2657611d26612317565b8160010160051415611d6a57611d3c878761241f565b6000805160206133b6833981519152600090815260466020526000805160206132ec83398151915255611d98565b6000805160206133b683398151915260005260466020526000805160206132ec833981519152805460010190555b82336001600160a01b03167f9d2e5f03fc65aff196e0f3a8dd924b24099de487e8cffc888921d420ab196e3989898987604051808060200185600560200280838360005b83811015611df4578181015183820152602001611ddc565b5050505090500184600560200280838360005b83811015611e1f578181015183820152602001611e07565b50505050905001838152602001828103825286818151815260200191508051906020019080838360005b83811015611e61578181015183820152602001611e49565b50505050905090810190601f168015611e8e5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a350505050505050565b80611eb157611fb7565b6000611ebc83610eef565b9050808282031115611f09576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b6000805160206132cc833981519152600052604660205260008051602061341783398151915254828103811015611f7b576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b611f878484840361127f565b50506000805160206132cc8339815191526000526046602052600080516020613417833981519152805482900390555b5050565b600082815260456020908152604080832060008051602061339683398151915284526001810190925290912054611ff29083612c30565b6000805160206133968339815191526000908152600183016020526040902055603a548314806120235750603c5483145b8061202f5750603e5483145b8061203b575060405483145b80612047575060425483145b1561207d576000805160206134c1833981519152600052604660205260008051602061347c833981519152805483019055611377565b6000805160206133f7833981519152600090815260018201602052604090205461218f576040805161066081019182905260009182916120dc91839060339082845b8154815260200190600101908083116120bf575050505050612c46565b6000805160206133968339815191526000908152600186016020526040902054919350915082108061210c575081155b1561218857600080516020613396833981519152600090815260018401602052604081205490826033811061213d57fe5b0155600081815260356020908152604080832080548452604583528184206000805160206133f783398151915285526001908101845282852085905590899055860190915290208190555b5050611377565b6000805160206133f78339815191526000908152600182016020526040812054839190603381106121bc57fe5b0180549091019055505050565b6001600160a01b0381166000908152604b602052604090205460ff1615612229576040805162461bcd60e51b815260206004820152600f60248201526e185b1c99591e481b5a59dc985d1959608a1b604482015290519081900360640190fd5b7f025c54659091f3c40ef9e9c7a0930301bc4b7dacc6211153686a3a296bba1177600052604760209081527f832f609aa70792c05e21641741f2ec8a9cca0cce1cc9bc1fe2eb85a2c242d9be54604080516370a0823160e01b81526001600160a01b03808616600483015291516122f394869493909316926370a082319260248082019391829003018186803b1580156122c257600080fd5b505afa1580156122d6573d6000803e3d6000fd5b505050506040513d60208110156122ec57600080fd5b5051612c95565b6001600160a01b03166000908152604b60205260409020805460ff19166001179055565b6000805160206132ac833981519152600090815260466020527f65764a602ca90a32a17ee52f54b96d452a879c4a2d747d2d464139cfbf12a3f7544203906123616104b083612e6d565b60466020526000805160206134e1833981519152547fad16221efc80aaf1b7e69bd3ecb61ba5ffa539adf129c3b4ffff769c9b5bbc336000527ffd954010d35b5c7a97c0324a8cd0b5bc2cb2a6250a538f0e5b66ec487e48234654610fa09290038102919091059150816123d457600191505b6123e18282016001612e83565b7fb12aff7664b16cb99339be399b863feecd64d14817be7e1f042f97e3f358e64e60005260466020526000805160206134e183398151915255505050565b7fd938684f43b130e7cfecbebf76ff7e28b74288b2b871d881041ace3b450910ec546000908152604560209081526040822060008051602061332e833981519152546000805160206132ac83398151915290935260469091527f65764a602ca90a32a17ee52f54b96d452a879c4a2d747d2d464139cfbf12a3f780544291829055919291906124ac6131f0565b6124b46131f0565b60005b60058110156128055760015b6005811015612672576000828152600689016020526040812082600581106124e757fe5b015490506000896005016000858152602001908152602001600020836005811061250d57fe5b01546001600160a01b03169050825b60008111801561254b5750600085815260068c016020526040902060001982016005811061254657fe5b015483105b1561260257600085815260068c016020526040902060001982016005811061256f57fe5b0154600086815260068d0160205260409020826005811061258c57fe5b015560008581526005808d0160205260409091209060001983019081106125af57fe5b015460008681526005808e0160205260409091206001600160a01b039092169190839081106125da57fe5b0180546001600160a01b0319166001600160a01b03929092169190911790556000190161251c565b8381101561266757600085815260068c01602052604090208390826005811061262757fe5b015560008581526005808d01602052604090912083918390811061264757fe5b0180546001600160a01b0319166001600160a01b03929092169190911790555b5050506001016124c3565b506000604560008a846005811061268557fe5b6020020151815260200190815260200160002090508760060160008381526020019081526020016000206005806020026040519081016040528092919082600580156126e6576020028201915b8154815260200190600101908083116126d2575b50505050509350836002600581106126fa57fe5b602090810291909101516000878152600384019092526040918290205584015183836005811061272657fe5b6020908102919091019190915260008381526005808b018352604080832089845285830190945290912061275b92909161320e565b5060008281526006808a016020908152604080842089855292850190915290912061278791600561320e565b50600082815260058901602052604081206127a191613249565b600082815260068901602052604081206127ba91613249565b805460018181018355600083815260208082209093018890558781526002840183526040808220439055600080516020613396833981519152825293820190925291812055016124b7565b50847fbeb3b9f46c8d7bb00c873fca10d307538df350866d25f891ffb395147ddbdc45888584604660006000805160206134c183398151915260001b8152602001908152602001600020546040518085600560200280838360005b83811015612878578181015183820152602001612860565b5050505090500184815260200183600560200280838360005b838110156128a9578181015183820152602001612891565b5050505090500182815260200194505050505060405180910390a2603380546001810182556000919091527f82a75bdeeae8604d839476ae9efd8b0e15aa447e21bfd7f41283bb54e22c9a82018390556129016131f0565b87516000908152604560209081526040808320878452600590810190925291829020825160a08101938490529290919082845b81546001600160a01b0316815260019091019060200180831161293457505050505090506129628186612e92565b7f969ea04b74d02bb4d9e6e8e57236e1b9ca31627139ae9f0e465249932e82450260005260466020527fd938684f43b130e7cfecbebf76ff7e28b74288b2b871d881041ace3b450910ec805460010190556129bb6131f0565b6129c3612fb7565b905060005b6005811015612ac8578181600581106129dd57fe5b6020020151603a82600581106129ef57fe5b6002020155600080604581858560058110612a0657fe5b6020020151815260200190815260200160002060010160006000805160206133f783398151915260001b81526020019081526020016000205460338110612a4957fe5b015560456000838360058110612a5b57fe5b6020908102919091015182528181019290925260409081016000908120600080516020613396833981519152825260019081018452918120546000805160206134c1833981519152909152604690925260008051602061347c8339815191528054909201909155016129c8565b50898760014303406040516020018080602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b83811015612b1c578181015183820152602001612b04565b50505050905090810190601f168015612b495780820380516001836020036101000a031916815260200191505b5060408051601f1981840301815290829052805160209182012060008051602061332e83398151915281905560469091526000805160206134e1833981519152546000805160206134c1833981519152600090815260008051602061347c83398151915254929f508f98507f1d85ce10456e29b67de37887496d3f1fcf1b64c79c4d07484038703a9f5c140897508996509094509092508190859060a0908190849084905b83811015612c06578181015183820152602001612bee565b5050505091909101938452505060208201526040805191829003019150a250505050505050505050565b600082820183811015612c3f57fe5b9392505050565b610640810151603260315b8015612c8f5782848260338110612c6457fe5b60200201511015612c8657838160338110612c7b57fe5b602002015192508091505b60001901612c51565b50915091565b80612cd15760405162461bcd60e51b815260040180806020018281038252602181526020018061334e6021913960400191505060405180910390fd5b6001600160a01b038216612d24576040805162461bcd60e51b815260206004820152601560248201527452656365697665722069732030206164647265737360581b604482015290519081900360640190fd5b6000612d2f83610eef565b9050808282011015612d7c576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b6000805160206132cc833981519152600052604660205260008051602061341783398151915254828101811115612dee576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b6000805160206132cc8339815191526000526046602052600080516020613417833981519152805484019055612e268483850161127f565b6040805184815290516001600160a01b038616916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a350505050565b6000818310612e7c5781612c3f565b5090919050565b6000818313612e7c5781612c3f565b6000805160206134c1833981519152600052604660205260008051602061347c833981519152544282900390670de0b6b3a76400009061012c8284020490600a810490600280840491612ee791309104611ea7565b612efb8760005b6020020151838501612c95565b612f06876001612eee565b612f11876002612eee565b612f1c876003612eee565b612f27876004612eee565b7f9dbc393ddc18fd27b1d9b1b129059925688d2f2d5818a5ec3ebb750b7c286ea660005260476020527f6af56e02192b4bb090b4d45a372637b353ff777e373e4590e08bf26bdb59b23854612f85906001600160a01b031682612c95565b50506000805160206134c18339815191526000908152604660205260008051602061347c833981519152555050505050565b612fbf6131f0565b612fc76131f0565b612fcf6131f0565b604080516106608101918290526130069160009060339082845b815481526020019060010190808311612fe957505050505061309c565b909250905060005b60058110156130965782816005811061302357fe5b602002015115613067576035600083836005811061303d57fe5b602002015181526020019081526020016000205484826005811061305d57fe5b602002015261308e565b603a816004036005811061307757fe5b600202015484826005811061308857fe5b60200201525b60010161300e565b50505090565b6130a46131f0565b6130ac6131f0565b60208301516000805b600581101561312f578581600101603381106130cd57fe5b60200201518582600581106130de57fe5b6020020152600181018482600581106130f357fe5b60200201528285826005811061310557fe5b602002015110156131275784816005811061311c57fe5b602002015192508091505b6001016130b5565b5060065b60338110156131e8578286826033811061314957fe5b602002015111156131e05785816033811061316057fe5b602002015185836005811061317157fe5b60200201528084836005811061318357fe5b602002015285816033811061319457fe5b6020020151925060005b60058110156131de57838682600581106131b457fe5b602002015110156131d6578581600581106131cb57fe5b602002015193508092505b60010161319e565b505b600101613133565b505050915091565b6040518060a001604052806005906020820280368337509192915050565b8260058101928215613239579182015b8281111561323957825482559160010191906001019061321e565b5061324592915061326c565b5090565b506000815560010160008155600101600081556001016000815560010160009055565b5b80821115613245576000815560010161326d56fe4d696e65722063616e206f6e6c792077696e2072657761726473206f6e636520706572203135206d696e97e6eb29f6a85471f7cc9b57f9e4c3deaf398cfc9798673160d7798baf0b13a4b1557182e4359a1f0c6301278e8f5b35a776ab58d39892581e357578fb2878362aaa6292065d1e7e4411c0f4a1c48b2d4a8f0d59408c35a865f884b064db576c45524332303a20617070726f766520746f20746865207a65726f20616464726573732742d3a105a8a106dd1f283149c83ac5908749f6dc7d56a6f73113194309473b547269656420746f206d696e74206e6f6e2d706f73697469766520616d6f756e7453686f756c6420686176652073756666696369656e742062616c616e636520746f2074726164652a9e355a92978430eca9c1aa3a9ba590094bac282594bccf82de16b83046e2c36c505cb2db6644f57b42d87bd9407b0f66788b07d0617a2bc1356a0e69e66f9a547269656420746f2073656e64206e6f6e2d706f73697469766520616d6f756e741e344bd070f05f1c5b3f0b1266f4f20d837a0a8190a3a2da8b0375eac2ba86ea643f72b9f9f841fd2188662a970bfac866592d96b69084b08c0231667b5a08e04d696e657220616c7265616479207375626d6974746564207468652076616c756545524332303a20617070726f76652066726f6d20746865207a65726f2061646472657373438a86ecc592f7fad199a71f3438a4d53f966778796bb459f2632fc8cfe3bfbb496e636f7272656374206e6f6e636520666f722063757272656e74206368616c6c656e6765d26d9834adf5a73309c4974bf654850bb699df8505e70d4cfde365c417b19dfceba0c7ec143c1b0e9691df5a04cd3d09ab7b2eb577a5171e5177f7378c475e2ea26469706673582212202c091457ba78de20eb4e0da3ebf8a1c34f35f1199171f151f6cf237f721f0dea64736f6c63430007040033"

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

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(bytes32)
func (_Tellor *TellorCaller) Owner(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(bytes32)
func (_Tellor *TellorSession) Owner() ([32]byte, error) {
	return _Tellor.Contract.Owner(&_Tellor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(bytes32)
func (_Tellor *TellorCallerSession) Owner() ([32]byte, error) {
	return _Tellor.Contract.Owner(&_Tellor.CallOpts)
}

// TBlock is a free data retrieval call binding the contract method 0x6e3cf885.
//
// Solidity: function _tBlock() view returns(bytes32)
func (_Tellor *TellorCaller) TBlock(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "_tBlock")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TBlock is a free data retrieval call binding the contract method 0x6e3cf885.
//
// Solidity: function _tBlock() view returns(bytes32)
func (_Tellor *TellorSession) TBlock() ([32]byte, error) {
	return _Tellor.Contract.TBlock(&_Tellor.CallOpts)
}

// TBlock is a free data retrieval call binding the contract method 0x6e3cf885.
//
// Solidity: function _tBlock() view returns(bytes32)
func (_Tellor *TellorCallerSession) TBlock() ([32]byte, error) {
	return _Tellor.Contract.TBlock(&_Tellor.CallOpts)
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

// CurrentChallenge is a free data retrieval call binding the contract method 0x51bdd585.
//
// Solidity: function currentChallenge() view returns(bytes32)
func (_Tellor *TellorCaller) CurrentChallenge(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "currentChallenge")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentChallenge is a free data retrieval call binding the contract method 0x51bdd585.
//
// Solidity: function currentChallenge() view returns(bytes32)
func (_Tellor *TellorSession) CurrentChallenge() ([32]byte, error) {
	return _Tellor.Contract.CurrentChallenge(&_Tellor.CallOpts)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x51bdd585.
//
// Solidity: function currentChallenge() view returns(bytes32)
func (_Tellor *TellorCallerSession) CurrentChallenge() ([32]byte, error) {
	return _Tellor.Contract.CurrentChallenge(&_Tellor.CallOpts)
}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5ae2bfdb.
//
// Solidity: function currentRequestId() view returns(bytes32)
func (_Tellor *TellorCaller) CurrentRequestId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "currentRequestId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5ae2bfdb.
//
// Solidity: function currentRequestId() view returns(bytes32)
func (_Tellor *TellorSession) CurrentRequestId() ([32]byte, error) {
	return _Tellor.Contract.CurrentRequestId(&_Tellor.CallOpts)
}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5ae2bfdb.
//
// Solidity: function currentRequestId() view returns(bytes32)
func (_Tellor *TellorCallerSession) CurrentRequestId() ([32]byte, error) {
	return _Tellor.Contract.CurrentRequestId(&_Tellor.CallOpts)
}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(bytes32)
func (_Tellor *TellorCaller) CurrentReward(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "currentReward")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(bytes32)
func (_Tellor *TellorSession) CurrentReward() ([32]byte, error) {
	return _Tellor.Contract.CurrentReward(&_Tellor.CallOpts)
}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(bytes32)
func (_Tellor *TellorCallerSession) CurrentReward() ([32]byte, error) {
	return _Tellor.Contract.CurrentReward(&_Tellor.CallOpts)
}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(bytes32)
func (_Tellor *TellorCaller) CurrentTotalTips(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "currentTotalTips")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(bytes32)
func (_Tellor *TellorSession) CurrentTotalTips() ([32]byte, error) {
	return _Tellor.Contract.CurrentTotalTips(&_Tellor.CallOpts)
}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(bytes32)
func (_Tellor *TellorCallerSession) CurrentTotalTips() ([32]byte, error) {
	return _Tellor.Contract.CurrentTotalTips(&_Tellor.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(bytes32)
func (_Tellor *TellorCaller) Difficulty(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "difficulty")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(bytes32)
func (_Tellor *TellorSession) Difficulty() ([32]byte, error) {
	return _Tellor.Contract.Difficulty(&_Tellor.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(bytes32)
func (_Tellor *TellorCallerSession) Difficulty() ([32]byte, error) {
	return _Tellor.Contract.Difficulty(&_Tellor.CallOpts)
}

// DisputeFee is a free data retrieval call binding the contract method 0xb9ce896b.
//
// Solidity: function disputeFee() view returns(bytes32)
func (_Tellor *TellorCaller) DisputeFee(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "disputeFee")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DisputeFee is a free data retrieval call binding the contract method 0xb9ce896b.
//
// Solidity: function disputeFee() view returns(bytes32)
func (_Tellor *TellorSession) DisputeFee() ([32]byte, error) {
	return _Tellor.Contract.DisputeFee(&_Tellor.CallOpts)
}

// DisputeFee is a free data retrieval call binding the contract method 0xb9ce896b.
//
// Solidity: function disputeFee() view returns(bytes32)
func (_Tellor *TellorCallerSession) DisputeFee() ([32]byte, error) {
	return _Tellor.Contract.DisputeFee(&_Tellor.CallOpts)
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

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(bytes32)
func (_Tellor *TellorCaller) PendingOwner(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "pending_owner")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(bytes32)
func (_Tellor *TellorSession) PendingOwner() ([32]byte, error) {
	return _Tellor.Contract.PendingOwner(&_Tellor.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(bytes32)
func (_Tellor *TellorCallerSession) PendingOwner() ([32]byte, error) {
	return _Tellor.Contract.PendingOwner(&_Tellor.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(bytes32)
func (_Tellor *TellorCaller) RequestCount(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "requestCount")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(bytes32)
func (_Tellor *TellorSession) RequestCount() ([32]byte, error) {
	return _Tellor.Contract.RequestCount(&_Tellor.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(bytes32)
func (_Tellor *TellorCallerSession) RequestCount() ([32]byte, error) {
	return _Tellor.Contract.RequestCount(&_Tellor.CallOpts)
}

// RequestQPosition is a free data retrieval call binding the contract method 0x2bf07e9e.
//
// Solidity: function requestQPosition() view returns(bytes32)
func (_Tellor *TellorCaller) RequestQPosition(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "requestQPosition")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RequestQPosition is a free data retrieval call binding the contract method 0x2bf07e9e.
//
// Solidity: function requestQPosition() view returns(bytes32)
func (_Tellor *TellorSession) RequestQPosition() ([32]byte, error) {
	return _Tellor.Contract.RequestQPosition(&_Tellor.CallOpts)
}

// RequestQPosition is a free data retrieval call binding the contract method 0x2bf07e9e.
//
// Solidity: function requestQPosition() view returns(bytes32)
func (_Tellor *TellorCallerSession) RequestQPosition() ([32]byte, error) {
	return _Tellor.Contract.RequestQPosition(&_Tellor.CallOpts)
}

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(bytes32)
func (_Tellor *TellorCaller) SlotProgress(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "slotProgress")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(bytes32)
func (_Tellor *TellorSession) SlotProgress() ([32]byte, error) {
	return _Tellor.Contract.SlotProgress(&_Tellor.CallOpts)
}

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(bytes32)
func (_Tellor *TellorCallerSession) SlotProgress() ([32]byte, error) {
	return _Tellor.Contract.SlotProgress(&_Tellor.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(bytes32)
func (_Tellor *TellorCaller) StakeAmount(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "stakeAmount")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(bytes32)
func (_Tellor *TellorSession) StakeAmount() ([32]byte, error) {
	return _Tellor.Contract.StakeAmount(&_Tellor.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(bytes32)
func (_Tellor *TellorCallerSession) StakeAmount() ([32]byte, error) {
	return _Tellor.Contract.StakeAmount(&_Tellor.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(bytes32)
func (_Tellor *TellorCaller) StakerCount(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(bytes32)
func (_Tellor *TellorSession) StakerCount() ([32]byte, error) {
	return _Tellor.Contract.StakerCount(&_Tellor.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(bytes32)
func (_Tellor *TellorCallerSession) StakerCount() ([32]byte, error) {
	return _Tellor.Contract.StakerCount(&_Tellor.CallOpts)
}

// TargetMiners is a free data retrieval call binding the contract method 0xdfee1ff1.
//
// Solidity: function targetMiners() view returns(bytes32)
func (_Tellor *TellorCaller) TargetMiners(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "targetMiners")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TargetMiners is a free data retrieval call binding the contract method 0xdfee1ff1.
//
// Solidity: function targetMiners() view returns(bytes32)
func (_Tellor *TellorSession) TargetMiners() ([32]byte, error) {
	return _Tellor.Contract.TargetMiners(&_Tellor.CallOpts)
}

// TargetMiners is a free data retrieval call binding the contract method 0xdfee1ff1.
//
// Solidity: function targetMiners() view returns(bytes32)
func (_Tellor *TellorCallerSession) TargetMiners() ([32]byte, error) {
	return _Tellor.Contract.TargetMiners(&_Tellor.CallOpts)
}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(bytes32)
func (_Tellor *TellorCaller) TimeOfLastNewValue(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "timeOfLastNewValue")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(bytes32)
func (_Tellor *TellorSession) TimeOfLastNewValue() ([32]byte, error) {
	return _Tellor.Contract.TimeOfLastNewValue(&_Tellor.CallOpts)
}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(bytes32)
func (_Tellor *TellorCallerSession) TimeOfLastNewValue() ([32]byte, error) {
	return _Tellor.Contract.TimeOfLastNewValue(&_Tellor.CallOpts)
}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(bytes32)
func (_Tellor *TellorCaller) TimeTarget(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "timeTarget")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(bytes32)
func (_Tellor *TellorSession) TimeTarget() ([32]byte, error) {
	return _Tellor.Contract.TimeTarget(&_Tellor.CallOpts)
}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(bytes32)
func (_Tellor *TellorCallerSession) TimeTarget() ([32]byte, error) {
	return _Tellor.Contract.TimeTarget(&_Tellor.CallOpts)
}

// TotalTip is a free data retrieval call binding the contract method 0x561cb04a.
//
// Solidity: function totalTip() view returns(bytes32)
func (_Tellor *TellorCaller) TotalTip(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "totalTip")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TotalTip is a free data retrieval call binding the contract method 0x561cb04a.
//
// Solidity: function totalTip() view returns(bytes32)
func (_Tellor *TellorSession) TotalTip() ([32]byte, error) {
	return _Tellor.Contract.TotalTip(&_Tellor.CallOpts)
}

// TotalTip is a free data retrieval call binding the contract method 0x561cb04a.
//
// Solidity: function totalTip() view returns(bytes32)
func (_Tellor *TellorCallerSession) TotalTip() ([32]byte, error) {
	return _Tellor.Contract.TotalTip(&_Tellor.CallOpts)
}

// TotalSupplyVar is a free data retrieval call binding the contract method 0x3940e9ee.
//
// Solidity: function total_supply() view returns(bytes32)
func (_Tellor *TellorCaller) TotalSupplyVar(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "total_supply")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TotalSupplyVar is a free data retrieval call binding the contract method 0x3940e9ee.
//
// Solidity: function total_supply() view returns(bytes32)
func (_Tellor *TellorSession) TotalSupplyVar() ([32]byte, error) {
	return _Tellor.Contract.TotalSupplyVar(&_Tellor.CallOpts)
}

// TotalSupplyVar is a free data retrieval call binding the contract method 0x3940e9ee.
//
// Solidity: function total_supply() view returns(bytes32)
func (_Tellor *TellorCallerSession) TotalSupplyVar() ([32]byte, error) {
	return _Tellor.Contract.TotalSupplyVar(&_Tellor.CallOpts)
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

// UpdateBalanceAtNow is a paid mutator transaction binding the contract method 0xd67dcbc5.
//
// Solidity: function updateBalanceAtNow(address _user, uint256 _value) returns()
func (_Tellor *TellorTransactor) UpdateBalanceAtNow(opts *bind.TransactOpts, _user common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "updateBalanceAtNow", _user, _value)
}

// UpdateBalanceAtNow is a paid mutator transaction binding the contract method 0xd67dcbc5.
//
// Solidity: function updateBalanceAtNow(address _user, uint256 _value) returns()
func (_Tellor *TellorSession) UpdateBalanceAtNow(_user common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.UpdateBalanceAtNow(&_Tellor.TransactOpts, _user, _value)
}

// UpdateBalanceAtNow is a paid mutator transaction binding the contract method 0xd67dcbc5.
//
// Solidity: function updateBalanceAtNow(address _user, uint256 _value) returns()
func (_Tellor *TellorTransactorSession) UpdateBalanceAtNow(_user common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.UpdateBalanceAtNow(&_Tellor.TransactOpts, _user, _value)
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

// TellorTransferEventIterator is returned from FilterTransferEvent and is used to iterate over the raw logs and unpacked data for TransferEvent events raised by the Tellor contract.
type TellorTransferEventIterator struct {
	Event *TellorTransferEvent // Event containing the contract specifics and raw log

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
func (it *TellorTransferEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorTransferEvent)
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
		it.Event = new(TellorTransferEvent)
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
func (it *TellorTransferEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorTransferEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorTransferEvent represents a TransferEvent event raised by the Tellor contract.
type TellorTransferEvent struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferEvent is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tellor *TellorFilterer) FilterTransferEvent(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TellorTransferEventIterator, error) {

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
	return &TellorTransferEventIterator{contract: _Tellor.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransferEvent is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tellor *TellorFilterer) WatchTransferEvent(opts *bind.WatchOpts, sink chan<- *TellorTransferEvent, from []common.Address, to []common.Address) (event.Subscription, error) {

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
				event := new(TellorTransferEvent)
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

// ParseTransferEvent is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tellor *TellorFilterer) ParseTransferEvent(log types.Log) (*TellorTransferEvent, error) {
	event := new(TellorTransferEvent)
	if err := _Tellor.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorGettersABI is the input ABI used to generate the binding from.
const TellorGettersABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_tBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bytesVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentChallenge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRequestId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentReward\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTotalTips\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"difficulty\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disputeFee\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[9]\",\"name\":\"\",\"type\":\"uint256[9]\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastNewValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"data\",\"type\":\"uint256[51]\"}],\"name\":\"getMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"data\",\"type\":\"uint256[51]\"}],\"name\":\"getMax5\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"max\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"maxIndex\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"data\",\"type\":\"uint256[51]\"}],\"name\":\"getMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"internalType\":\"address[5]\",\"name\":\"\",\"type\":\"address[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"idsOnDeck\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"tipsOnDeck\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"\",\"type\":\"uint256[51]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTopRequestIDs\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pending_owner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestCount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestQPosition\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slotProgress\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeAmount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetMiners\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeOfLastNewValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeTarget\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTip\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"total_supply\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"uints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TellorGettersFuncSigs maps the 4-byte function signature to its string representation.
var TellorGettersFuncSigs = map[string]string{
	"024c2ddd": "_allowances(address,address)",
	"b2bdfa7b": "_owner()",
	"6e3cf885": "_tBlock()",
	"699f200f": "addresses(bytes32)",
	"62dd1d2a": "bytesVars(bytes32)",
	"51bdd585": "currentChallenge()",
	"5ae2bfdb": "currentRequestId()",
	"07621eca": "currentReward()",
	"75ad1a2a": "currentTotalTips()",
	"313ce567": "decimals()",
	"63bb82ad": "didMine(bytes32,address)",
	"a7c438bc": "didVote(uint256,address)",
	"19cae462": "difficulty()",
	"b9ce896b": "disputeFee()",
	"133bee5e": "getAddressVars(bytes32)",
	"af0b1327": "getAllDisputeVars(uint256)",
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
	"06fdde03": "name()",
	"7f4ec4c3": "pending_owner()",
	"5badbe4c": "requestCount()",
	"2bf07e9e": "requestQPosition()",
	"93fa4915": "retrieveData(uint256,uint256)",
	"03b3160f": "slotProgress()",
	"60c7dc47": "stakeAmount()",
	"dff69787": "stakerCount()",
	"95d89b41": "symbol()",
	"dfee1ff1": "targetMiners()",
	"6fd4f229": "timeOfLastNewValue()",
	"6fc37811": "timeTarget()",
	"18160ddd": "totalSupply()",
	"561cb04a": "totalTip()",
	"3940e9ee": "total_supply()",
	"b59e14d4": "uints(bytes32)",
}

// TellorGettersBin is the compiled bytecode used for deploying new contracts.
var TellorGettersBin = "0x608060405234801561001057600080fd5b50611a3e806100206000396000f3fe608060405234801561001057600080fd5b50600436106103785760003560e01c80636e3cf885116101d3578063b2bdfa7b11610104578063dfee1ff1116100a2578063e1eee6d61161007c578063e1eee6d614610ab0578063f29e5e9a14610acd578063fc7cf0a014610b21578063fe1cd15d14610b2957610378565b8063dfee1ff114610a7d578063dff6978714610a85578063e0ae93c114610a8d57610378565b8063b9ce896b116100de578063b9ce896b146109e1578063c775b542146109e9578063c87a336d14610a0c578063da37994114610a6057610378565b8063b2bdfa7b1461099e578063b5413029146109a6578063b59e14d4146109c457610378565b80637f6fd5d91161017157806399830e321161014b57806399830e32146108015780639a7077ab146108b9578063a7c438bc146108c1578063af0b1327146108ed57610378565b80637f6fd5d9146107b357806393fa4915146107d657806395d89b41146107f957610378565b8063733bdef0116101ad578063733bdef01461074157806375ad1a2a1461078057806377fbb663146107885780637f4ec4c3146107ab57610378565b80636e3cf885146107295780636fc37811146107315780636fd4f2291461073957610378565b80634049f198116102ad57806360c7dc471161024b57806362dd1d2a1161022557806362dd1d2a146106a057806363bb82ad146106bd57806369026d63146106e9578063699f200f1461070c57610378565b806360c7dc471461065e578063612c8f7f146106665780636173c0b81461068357610378565b806351bdd5851161028757806351bdd5851461063e578063561cb04a146106465780635ae2bfdb1461064e5780635badbe4c1461065657610378565b80634049f198146105a657806346eee1c4146105fb5780634ba0a5ee1461061857610378565b806318160ddd1161031a578063313ce567116102f4578063313ce567146105135780633180f8df146105315780633940e9ee146105675780633df0777b1461056f57610378565b806318160ddd146104fb57806319cae462146105035780632bf07e9e1461050b57610378565b806307621eca1161035657806307621eca146104425780630f0b424d1461044a57806311c9851214610467578063133bee5e146104c257610378565b8063024c2ddd1461037d57806303b3160f146103bd57806306fdde03146103c5575b600080fd5b6103ab6004803603604081101561039357600080fd5b506001600160a01b0381358116916020013516610b31565b60408051918252519081900360200190f35b6103ab610b4e565b6103cd610b72565b6040805160208082528351818301528351919283929083019185019080838360005b838110156104075781810151838201526020016103ef565b50505050905090810190601f1680156104345780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103ab610b9b565b6103ab6004803603602081101561046057600080fd5b5035610bbf565b61048a6004803603604081101561047d57600080fd5b5080359060200135610bd1565b604051808260a080838360005b838110156104af578181015183820152602001610497565b5050505090500191505060405180910390f35b6104df600480360360208110156104d857600080fd5b5035610c28565b604080516001600160a01b039092168252519081900360200190f35b6103ab610c43565b6103ab610c91565b6103ab610cb5565b61051b610cd9565b6040805160ff9092168252519081900360200190f35b61054e6004803603602081101561054757600080fd5b5035610cde565b6040805192835290151560208301528051918290030190f35b6103ab610d3e565b6105926004803603604081101561058557600080fd5b5080359060200135610d62565b604080519115158252519081900360200190f35b6105ae610d86565b604051848152602081018460a080838360005b838110156105d95781810151838201526020016105c1565b5050505090500183815260200182815260200194505050505060405180910390f35b6103ab6004803603602081101561061157600080fd5b5035610e66565b6105926004803603602081101561062e57600080fd5b50356001600160a01b0316610e78565b6103ab610e8d565b6103ab610eb1565b6103ab610ed5565b6103ab610ef9565b6103ab610f1d565b6103ab6004803603602081101561067c57600080fd5b5035610f41565b6103ab6004803603602081101561069957600080fd5b5035610f53565b6103ab600480360360208110156106b657600080fd5b5035610fbe565b610592600480360360408110156106d357600080fd5b50803590602001356001600160a01b0316610fd0565b61048a600480360360408110156106ff57600080fd5b5080359060200135610ffb565b6104df6004803603602081101561072257600080fd5b503561105e565b6103ab611079565b6103ab61109d565b6103ab6110c1565b6107676004803603602081101561075757600080fd5b50356001600160a01b03166110e5565b6040805192835260208301919091528051918290030190f35b6103ab611108565b6103ab6004803603604081101561079e57600080fd5b508035906020013561112c565b6103ab611159565b6103ab600480360360408110156107c957600080fd5b508035906020013561117d565b6103ab600480360360408110156107ec57600080fd5b508035906020013561119e565b6103cd6111bf565b610855600480360361066081101561081857600080fd5b8101908080610660019060338060200260405190810160405280929190826033602002808284376000920191909152509194506111dc9350505050565b604051808360a080838360005b8381101561087a578181015183820152602001610862565b5050505090500182600560200280838360005b838110156108a557818101518382015260200161088d565b505050509050019250505060405180910390f35b610855611330565b610592600480360360408110156108d757600080fd5b50803590602001356001600160a01b03166113cf565b61090a6004803603602081101561090357600080fd5b50356113fe565b604051808a8152602001891515815260200188151581526020018715158152602001866001600160a01b03168152602001856001600160a01b03168152602001846001600160a01b0316815260200183600960200280838360005b8381101561097d578181015183820152602001610965565b50505050905001828152602001995050505050505050505060405180910390f35b6103ab611629565b6109ae61164d565b6040518151815280826106608083836020610497565b6103ab600480360360208110156109da57600080fd5b5035611689565b6103ab61169b565b6103ab600480360360408110156109ff57600080fd5b50803590602001356116bf565b6107676004803603610660811015610a2357600080fd5b8101908080610660019060338060200260405190810160405280929190826033602002808284376000920191909152509194506116e09350505050565b6103ab60048036036020811015610a7657600080fd5b503561172a565b6103ab61173c565b6103ab611760565b6103ab60048036036040811015610aa357600080fd5b5080359060200135611784565b61076760048036036020811015610ac657600080fd5b50356117a5565b6107676004803603610660811015610ae457600080fd5b81019080806106600190603380602002604051908101604052809291908260336020028082843760009201919091525091945061180d9350505050565b61054e611856565b61048a6118c7565b604a60209081526000928352604080842090915290825290205481565b7f6c505cb2db6644f57b42d87bd9407b0f66788b07d0617a2bc1356a0e69e66f9a81565b60408051808201909152600f81526e54656c6c6f7220547269627574657360881b602082015290565b7f9b6853911475b07474368644a0d922ee13bc76a15cd3e97d3e334326424a47d481565b60009081526034602052604090205490565b610bd96119ac565b600083815260456020908152604080832085845260060190915290819020815160a08101928390529160059082845b815481526020019060010190808311610c08575050505050905092915050565b6000908152604760205260409020546001600160a01b031690565b7fb1557182e4359a1f0c6301278e8f5b35a776ab58d39892581e357578fb28783660005260466020527f643f72b9f9f841fd2188662a970bfac866592d96b69084b08c0231667b5a08e05490565b7fb12aff7664b16cb99339be399b863feecd64d14817be7e1f042f97e3f358e64e81565b7f1e344bd070f05f1c5b3f0b1266f4f20d837a0a8190a3a2da8b0375eac2ba86ea81565b601290565b6000818152604560205260408120805482919015610d2a578054610d1e90859083906000198101908110610d0e57fe5b906000526020600020015461119e565b60019250925050610d39565b6000809250925050610d39565b505b915091565b7fb1557182e4359a1f0c6301278e8f5b35a776ab58d39892581e357578fb28783681565b60009182526045602090815260408084209284526004909201905290205460ff1690565b6000610d906119ac565b60008060005b6005811015610dca57603a8160058110610dac57fe5b6002020154848260058110610dbd57fe5b6020020152600101610d96565b50507f2742d3a105a8a106dd1f283149c83ac5908749f6dc7d56a6f73113194309473b5460466020527feba0c7ec143c1b0e9691df5a04cd3d09ab7b2eb577a5171e5177f7378c475e2e547fd26d9834adf5a73309c4974bf654850bb699df8505e70d4cfde365c417b19dfc6000527f438a86ecc592f7fad199a71f3438a4d53f966778796bb459f2632fc8cfe3bfbb54919450915090919293565b60009081526045602052604090205490565b604b6020526000908152604090205460ff1681565b7f3375fb9157bb77048f607329b1c4d45487433f61e0f51bcdeba91d975b2dab1881565b7f2a9e355a92978430eca9c1aa3a9ba590094bac282594bccf82de16b83046e2c381565b7f7584d7d8701714da9c117f5bf30af73b0b88aca5338a84a21eb28de2fe0d93b881565b7f05de9147d05477c0a5dc675aeea733157f5092f82add148cf39d579cafe3dc9881565b7f7be108969d31a3f0b261465c71f2b0ba9301cd914d55d9091c3b36a49d4d41b281565b60009081526046602052604090205490565b60006032821115610fab576040805162461bcd60e51b815260206004820152601a60248201527f526571756573745120696e6465782069732061626f7665203530000000000000604482015290519081900360640190fd5b5060009081526035602052604090205490565b60486020526000908152604090205481565b60009182526039602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6110036119ac565b6000838152604560209081526040808320858452600590810190925291829020825160a08101938490529290919082845b81546001600160a01b03168152600190910190602001808311611034575050505050905092915050565b6047602052600090815260409020546001600160a01b031681565b7f969ea04b74d02bb4d9e6e8e57236e1b9ca31627139ae9f0e465249932e82450281565b7fad16221efc80aaf1b7e69bd3ecb61ba5ffa539adf129c3b4ffff769c9b5bbc3381565b7f97e6eb29f6a85471f7cc9b57f9e4c3deaf398cfc9798673160d7798baf0b13a481565b6001600160a01b0316600090815260446020526040902080546001909101549091565b7fd26d9834adf5a73309c4974bf654850bb699df8505e70d4cfde365c417b19dfc81565b600082815260456020526040812080548390811061114657fe5b9060005260206000200154905092915050565b7f44b2657a0f8a90ed8e62f4c4cceca06eacaa9b4b25751ae1ebca9280a70abd6881565b60009182526036602090815260408084209284526005909201905290205490565b60009182526045602090815260408084209284526003909201905290205490565b6040805180820190915260038152622a292160e91b602082015290565b6111e46119ac565b6111ec6119ac565b60208301516000805b600581101561126f5785816001016033811061120d57fe5b602002015185826005811061121e57fe5b60200201526001810184826005811061123357fe5b60200201528285826005811061124557fe5b602002015110156112675784816005811061125c57fe5b602002015192508091505b6001016111f5565b5060065b6033811015611328578286826033811061128957fe5b60200201511115611320578581603381106112a057fe5b60200201518583600581106112b157fe5b6020020152808483600581106112c357fe5b60200201528581603381106112d457fe5b6020020151925060005b600581101561131e57838682600581106112f457fe5b602002015110156113165785816005811061130b57fe5b602002015193508092505b6001016112de565b505b600101611273565b505050915091565b6113386119ac565b6113406119ac565b6113486118c7565b915060005b60058110156113ca576045600084836005811061136657fe5b6020020151815260200190815260200160002060010160007f2a9e355a92978430eca9c1aa3a9ba590094bac282594bccf82de16b83046e2c360001b8152602001908152602001600020548282600581106113bd57fe5b602002015260010161134d565b509091565b60008281526036602090815260408083206001600160a01b038516845260060190915290205460ff1692915050565b60008060008060008060006114116119ca565b5050506000958652505060366020908152604080862080546002820154600383015460048401548551610120810187527f31b40192effc42bcf1e4289fe674c678e673a3052992548fef566d8c33a21b918c5260058601808952878d205482527f4ebf727c48eac2c66272456b06a885c5cc03e54d140f63b63b6fd10c1227958e8d52808952878d2054828a01527f81afeeaff0ed5cee7d05a21078399c2f56226b0cd5657062500cef4c4e736f858d52808952878d2054828901527f74c9bc34b0b2333f1b565fbee67d940cf7d78b5a980c5f23da43f6729965ed408d52808952878d205460608301527fa0bc13ce85a2091e950a370bced0825e58ab3a3ffeb709ed50d5562cbd82faab8d52808952878d205460808301527f6f8f54d1af9b6cb8a219d88672c797f9f3ee97ce5d9369aa897fd0deb5e2dffa8d52808952878d205460a08301527f8ef61a1efbc527d6428ff88c95fdff5c6e644b979bfe67e03cbf88c8162c5fac8d52808952878d205460c08301527f30e85ae205656781c1a951cba9f9f53f884833c049d377a2a7046eb5e6d14b268d52808952878d205460e08301527f833b9f6abf0b529613680afe2a00fa663cc95cbdc47d726d85a044462eabbf028d52909752949099205461010080870191909152600190930154919960ff8083169a948304811699506201000083041697506001600160a01b036301000000909204821696509281169493169291565b7f9dbc393ddc18fd27b1d9b1b129059925688d2f2d5818a5ec3ebb750b7c286ea681565b6116556119e9565b604080516106608101918290529060009060339082845b81548152602001906001019080831161166c575050505050905090565b60466020526000908152604090205481565b7f8b75eb45d88e80f0e4ec77d23936268694c0e7ac2e0c9085c5c6bdfcfbc4923981565b60009182526045602090815260408084209284526002909201905290205490565b6020810151600160025b6033811015610d37578284826033811061170057fe5b602002015111156117225783816033811061171757fe5b602002015192508091505b6001016116ea565b60009081526038602052604090205490565b7fabef544d8048318ece54fb2c6385255cd1b06e176525d149a0338a7acca6deb381565b7fedddb9344bfe0dadc78c558b8ffca446679cbffc17be64eb83973fce7bea5f3481565b60009182526045602090815260408084209284526001909201905290205490565b60009081526045602090815260408083207f1e344bd070f05f1c5b3f0b1266f4f20d837a0a8190a3a2da8b0375eac2ba86ea8452600101909152808220547f2a9e355a92978430eca9c1aa3a9ba590094bac282594bccf82de16b83046e2c383529120549091565b610640810151603260315b8015610d37578284826033811061182b57fe5b6020020151101561184d5783816033811061184257fe5b602002015192508091505b60001901611818565b7f65764a602ca90a32a17ee52f54b96d452a879c4a2d747d2d464139cfbf12a3f75460008181526034602090815260408220547f97e6eb29f6a85471f7cc9b57f9e4c3deaf398cfc9798673160d7798baf0b13a483526046909152909182916118be9161119e565b92600192509050565b6118cf6119ac565b6118d76119ac565b6118df6119ac565b604080516106608101918290526119169160009060339082845b8154815260200190600101908083116118f95750505050506111dc565b909250905060005b60058110156119a65782816005811061193357fe5b602002015115611977576035600083836005811061194d57fe5b602002015181526020019081526020016000205484826005811061196d57fe5b602002015261199e565b603a816004036005811061198757fe5b600202015484826005811061199857fe5b60200201525b60010161191e565b50505090565b6040518060a001604052806005906020820280368337509192915050565b6040518061012001604052806009906020820280368337509192915050565b604051806106600160405280603390602082028036833750919291505056fea26469706673582212207ed9733a666a42a7b600371dcc4313d562e14896fecb3587b4ddffd61f72ad3064736f6c63430007040033"

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

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) Owner(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(bytes32)
func (_TellorGetters *TellorGettersSession) Owner() ([32]byte, error) {
	return _TellorGetters.Contract.Owner(&_TellorGetters.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) Owner() ([32]byte, error) {
	return _TellorGetters.Contract.Owner(&_TellorGetters.CallOpts)
}

// TBlock is a free data retrieval call binding the contract method 0x6e3cf885.
//
// Solidity: function _tBlock() view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) TBlock(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "_tBlock")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TBlock is a free data retrieval call binding the contract method 0x6e3cf885.
//
// Solidity: function _tBlock() view returns(bytes32)
func (_TellorGetters *TellorGettersSession) TBlock() ([32]byte, error) {
	return _TellorGetters.Contract.TBlock(&_TellorGetters.CallOpts)
}

// TBlock is a free data retrieval call binding the contract method 0x6e3cf885.
//
// Solidity: function _tBlock() view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) TBlock() ([32]byte, error) {
	return _TellorGetters.Contract.TBlock(&_TellorGetters.CallOpts)
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

// CurrentChallenge is a free data retrieval call binding the contract method 0x51bdd585.
//
// Solidity: function currentChallenge() view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) CurrentChallenge(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "currentChallenge")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentChallenge is a free data retrieval call binding the contract method 0x51bdd585.
//
// Solidity: function currentChallenge() view returns(bytes32)
func (_TellorGetters *TellorGettersSession) CurrentChallenge() ([32]byte, error) {
	return _TellorGetters.Contract.CurrentChallenge(&_TellorGetters.CallOpts)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x51bdd585.
//
// Solidity: function currentChallenge() view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) CurrentChallenge() ([32]byte, error) {
	return _TellorGetters.Contract.CurrentChallenge(&_TellorGetters.CallOpts)
}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5ae2bfdb.
//
// Solidity: function currentRequestId() view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) CurrentRequestId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "currentRequestId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5ae2bfdb.
//
// Solidity: function currentRequestId() view returns(bytes32)
func (_TellorGetters *TellorGettersSession) CurrentRequestId() ([32]byte, error) {
	return _TellorGetters.Contract.CurrentRequestId(&_TellorGetters.CallOpts)
}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5ae2bfdb.
//
// Solidity: function currentRequestId() view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) CurrentRequestId() ([32]byte, error) {
	return _TellorGetters.Contract.CurrentRequestId(&_TellorGetters.CallOpts)
}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) CurrentReward(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "currentReward")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(bytes32)
func (_TellorGetters *TellorGettersSession) CurrentReward() ([32]byte, error) {
	return _TellorGetters.Contract.CurrentReward(&_TellorGetters.CallOpts)
}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) CurrentReward() ([32]byte, error) {
	return _TellorGetters.Contract.CurrentReward(&_TellorGetters.CallOpts)
}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) CurrentTotalTips(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "currentTotalTips")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(bytes32)
func (_TellorGetters *TellorGettersSession) CurrentTotalTips() ([32]byte, error) {
	return _TellorGetters.Contract.CurrentTotalTips(&_TellorGetters.CallOpts)
}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) CurrentTotalTips() ([32]byte, error) {
	return _TellorGetters.Contract.CurrentTotalTips(&_TellorGetters.CallOpts)
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

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) Difficulty(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "difficulty")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(bytes32)
func (_TellorGetters *TellorGettersSession) Difficulty() ([32]byte, error) {
	return _TellorGetters.Contract.Difficulty(&_TellorGetters.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) Difficulty() ([32]byte, error) {
	return _TellorGetters.Contract.Difficulty(&_TellorGetters.CallOpts)
}

// DisputeFee is a free data retrieval call binding the contract method 0xb9ce896b.
//
// Solidity: function disputeFee() view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) DisputeFee(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "disputeFee")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DisputeFee is a free data retrieval call binding the contract method 0xb9ce896b.
//
// Solidity: function disputeFee() view returns(bytes32)
func (_TellorGetters *TellorGettersSession) DisputeFee() ([32]byte, error) {
	return _TellorGetters.Contract.DisputeFee(&_TellorGetters.CallOpts)
}

// DisputeFee is a free data retrieval call binding the contract method 0xb9ce896b.
//
// Solidity: function disputeFee() view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) DisputeFee() ([32]byte, error) {
	return _TellorGetters.Contract.DisputeFee(&_TellorGetters.CallOpts)
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

// GetMax is a free data retrieval call binding the contract method 0xc87a336d.
//
// Solidity: function getMax(uint256[51] data) pure returns(uint256 max, uint256 maxIndex)
func (_TellorGetters *TellorGettersCaller) GetMax(opts *bind.CallOpts, data [51]*big.Int) (struct {
	Max      *big.Int
	MaxIndex *big.Int
}, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getMax", data)

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
// Solidity: function getMax(uint256[51] data) pure returns(uint256 max, uint256 maxIndex)
func (_TellorGetters *TellorGettersSession) GetMax(data [51]*big.Int) (struct {
	Max      *big.Int
	MaxIndex *big.Int
}, error) {
	return _TellorGetters.Contract.GetMax(&_TellorGetters.CallOpts, data)
}

// GetMax is a free data retrieval call binding the contract method 0xc87a336d.
//
// Solidity: function getMax(uint256[51] data) pure returns(uint256 max, uint256 maxIndex)
func (_TellorGetters *TellorGettersCallerSession) GetMax(data [51]*big.Int) (struct {
	Max      *big.Int
	MaxIndex *big.Int
}, error) {
	return _TellorGetters.Contract.GetMax(&_TellorGetters.CallOpts, data)
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

// GetMin is a free data retrieval call binding the contract method 0xf29e5e9a.
//
// Solidity: function getMin(uint256[51] data) pure returns(uint256 min, uint256 minIndex)
func (_TellorGetters *TellorGettersCaller) GetMin(opts *bind.CallOpts, data [51]*big.Int) (struct {
	Min      *big.Int
	MinIndex *big.Int
}, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getMin", data)

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
// Solidity: function getMin(uint256[51] data) pure returns(uint256 min, uint256 minIndex)
func (_TellorGetters *TellorGettersSession) GetMin(data [51]*big.Int) (struct {
	Min      *big.Int
	MinIndex *big.Int
}, error) {
	return _TellorGetters.Contract.GetMin(&_TellorGetters.CallOpts, data)
}

// GetMin is a free data retrieval call binding the contract method 0xf29e5e9a.
//
// Solidity: function getMin(uint256[51] data) pure returns(uint256 min, uint256 minIndex)
func (_TellorGetters *TellorGettersCallerSession) GetMin(data [51]*big.Int) (struct {
	Min      *big.Int
	MinIndex *big.Int
}, error) {
	return _TellorGetters.Contract.GetMin(&_TellorGetters.CallOpts, data)
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
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficulty, uint256 _tip)
func (_TellorGetters *TellorGettersCaller) GetNewCurrentVariables(opts *bind.CallOpts) (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficulty *big.Int
	Tip        *big.Int
}, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getNewCurrentVariables")

	outstruct := new(struct {
		Challenge  [32]byte
		RequestIds [5]*big.Int
		Difficulty *big.Int
		Tip        *big.Int
	})

	outstruct.Challenge = out[0].([32]byte)
	outstruct.RequestIds = out[1].([5]*big.Int)
	outstruct.Difficulty = out[2].(*big.Int)
	outstruct.Tip = out[3].(*big.Int)

	return *outstruct, err

}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficulty, uint256 _tip)
func (_TellorGetters *TellorGettersSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficulty *big.Int
	Tip        *big.Int
}, error) {
	return _TellorGetters.Contract.GetNewCurrentVariables(&_TellorGetters.CallOpts)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficulty, uint256 _tip)
func (_TellorGetters *TellorGettersCallerSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficulty *big.Int
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

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) PendingOwner(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "pending_owner")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(bytes32)
func (_TellorGetters *TellorGettersSession) PendingOwner() ([32]byte, error) {
	return _TellorGetters.Contract.PendingOwner(&_TellorGetters.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) PendingOwner() ([32]byte, error) {
	return _TellorGetters.Contract.PendingOwner(&_TellorGetters.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) RequestCount(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "requestCount")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(bytes32)
func (_TellorGetters *TellorGettersSession) RequestCount() ([32]byte, error) {
	return _TellorGetters.Contract.RequestCount(&_TellorGetters.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) RequestCount() ([32]byte, error) {
	return _TellorGetters.Contract.RequestCount(&_TellorGetters.CallOpts)
}

// RequestQPosition is a free data retrieval call binding the contract method 0x2bf07e9e.
//
// Solidity: function requestQPosition() view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) RequestQPosition(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "requestQPosition")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RequestQPosition is a free data retrieval call binding the contract method 0x2bf07e9e.
//
// Solidity: function requestQPosition() view returns(bytes32)
func (_TellorGetters *TellorGettersSession) RequestQPosition() ([32]byte, error) {
	return _TellorGetters.Contract.RequestQPosition(&_TellorGetters.CallOpts)
}

// RequestQPosition is a free data retrieval call binding the contract method 0x2bf07e9e.
//
// Solidity: function requestQPosition() view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) RequestQPosition() ([32]byte, error) {
	return _TellorGetters.Contract.RequestQPosition(&_TellorGetters.CallOpts)
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

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) SlotProgress(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "slotProgress")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(bytes32)
func (_TellorGetters *TellorGettersSession) SlotProgress() ([32]byte, error) {
	return _TellorGetters.Contract.SlotProgress(&_TellorGetters.CallOpts)
}

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) SlotProgress() ([32]byte, error) {
	return _TellorGetters.Contract.SlotProgress(&_TellorGetters.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) StakeAmount(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "stakeAmount")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(bytes32)
func (_TellorGetters *TellorGettersSession) StakeAmount() ([32]byte, error) {
	return _TellorGetters.Contract.StakeAmount(&_TellorGetters.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) StakeAmount() ([32]byte, error) {
	return _TellorGetters.Contract.StakeAmount(&_TellorGetters.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) StakerCount(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(bytes32)
func (_TellorGetters *TellorGettersSession) StakerCount() ([32]byte, error) {
	return _TellorGetters.Contract.StakerCount(&_TellorGetters.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) StakerCount() ([32]byte, error) {
	return _TellorGetters.Contract.StakerCount(&_TellorGetters.CallOpts)
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

// TargetMiners is a free data retrieval call binding the contract method 0xdfee1ff1.
//
// Solidity: function targetMiners() view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) TargetMiners(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "targetMiners")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TargetMiners is a free data retrieval call binding the contract method 0xdfee1ff1.
//
// Solidity: function targetMiners() view returns(bytes32)
func (_TellorGetters *TellorGettersSession) TargetMiners() ([32]byte, error) {
	return _TellorGetters.Contract.TargetMiners(&_TellorGetters.CallOpts)
}

// TargetMiners is a free data retrieval call binding the contract method 0xdfee1ff1.
//
// Solidity: function targetMiners() view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) TargetMiners() ([32]byte, error) {
	return _TellorGetters.Contract.TargetMiners(&_TellorGetters.CallOpts)
}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) TimeOfLastNewValue(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "timeOfLastNewValue")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(bytes32)
func (_TellorGetters *TellorGettersSession) TimeOfLastNewValue() ([32]byte, error) {
	return _TellorGetters.Contract.TimeOfLastNewValue(&_TellorGetters.CallOpts)
}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) TimeOfLastNewValue() ([32]byte, error) {
	return _TellorGetters.Contract.TimeOfLastNewValue(&_TellorGetters.CallOpts)
}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) TimeTarget(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "timeTarget")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(bytes32)
func (_TellorGetters *TellorGettersSession) TimeTarget() ([32]byte, error) {
	return _TellorGetters.Contract.TimeTarget(&_TellorGetters.CallOpts)
}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) TimeTarget() ([32]byte, error) {
	return _TellorGetters.Contract.TimeTarget(&_TellorGetters.CallOpts)
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

// TotalTip is a free data retrieval call binding the contract method 0x561cb04a.
//
// Solidity: function totalTip() view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) TotalTip(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "totalTip")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TotalTip is a free data retrieval call binding the contract method 0x561cb04a.
//
// Solidity: function totalTip() view returns(bytes32)
func (_TellorGetters *TellorGettersSession) TotalTip() ([32]byte, error) {
	return _TellorGetters.Contract.TotalTip(&_TellorGetters.CallOpts)
}

// TotalTip is a free data retrieval call binding the contract method 0x561cb04a.
//
// Solidity: function totalTip() view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) TotalTip() ([32]byte, error) {
	return _TellorGetters.Contract.TotalTip(&_TellorGetters.CallOpts)
}

// TotalSupplyVar is a free data retrieval call binding the contract method 0x3940e9ee.
//
// Solidity: function total_supply() view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) TotalSupplyVar(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "total_supply")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TotalSupplyVar is a free data retrieval call binding the contract method 0x3940e9ee.
//
// Solidity: function total_supply() view returns(bytes32)
func (_TellorGetters *TellorGettersSession) TotalSupplyVar() ([32]byte, error) {
	return _TellorGetters.Contract.TotalSupplyVar(&_TellorGetters.CallOpts)
}

// TotalSupplyVar is a free data retrieval call binding the contract method 0x3940e9ee.
//
// Solidity: function total_supply() view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) TotalSupplyVar() ([32]byte, error) {
	return _TellorGetters.Contract.TotalSupplyVar(&_TellorGetters.CallOpts)
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

// TellorStorageABI is the input ABI used to generate the binding from.
const TellorStorageABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bytesVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"uints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TellorStorageFuncSigs maps the 4-byte function signature to its string representation.
var TellorStorageFuncSigs = map[string]string{
	"024c2ddd": "_allowances(address,address)",
	"699f200f": "addresses(bytes32)",
	"62dd1d2a": "bytesVars(bytes32)",
	"4ba0a5ee": "migrated(address)",
	"b59e14d4": "uints(bytes32)",
}

// TellorStorageBin is the compiled bytecode used for deploying new contracts.
var TellorStorageBin = "0x608060405234801561001057600080fd5b506101f0806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063024c2ddd1461005c5780634ba0a5ee1461009c57806362dd1d2a146100d6578063699f200f146100f3578063b59e14d41461012c575b600080fd5b61008a6004803603604081101561007257600080fd5b506001600160a01b0381358116916020013516610149565b60408051918252519081900360200190f35b6100c2600480360360208110156100b257600080fd5b50356001600160a01b0316610166565b604080519115158252519081900360200190f35b61008a600480360360208110156100ec57600080fd5b503561017b565b6101106004803603602081101561010957600080fd5b503561018d565b604080516001600160a01b039092168252519081900360200190f35b61008a6004803603602081101561014257600080fd5b50356101a8565b604a60209081526000928352604080842090915290825290205481565b604b6020526000908152604090205460ff1681565b60486020526000908152604090205481565b6047602052600090815260409020546001600160a01b031681565b6046602052600090815260409020548156fea2646970667358221220dc160799b2b4ada56a481657f932daa301a3c7b1a6f79b35a6fbf422c53f50d564736f6c63430007040033"

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
const TellorTransferABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_tBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bytesVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentChallenge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRequestId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentReward\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTotalTips\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"difficulty\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disputeFee\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pending_owner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestCount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestQPosition\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slotProgress\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeAmount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetMiners\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeOfLastNewValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeTarget\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTip\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"total_supply\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"uints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"updateBalanceAtNow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TellorTransferFuncSigs maps the 4-byte function signature to its string representation.
var TellorTransferFuncSigs = map[string]string{
	"024c2ddd": "_allowances(address,address)",
	"b2bdfa7b": "_owner()",
	"6e3cf885": "_tBlock()",
	"699f200f": "addresses(bytes32)",
	"dd62ed3e": "allowance(address,address)",
	"999cf26c": "allowedToTrade(address,uint256)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"4ee2cd7e": "balanceOfAt(address,uint256)",
	"62dd1d2a": "bytesVars(bytes32)",
	"51bdd585": "currentChallenge()",
	"5ae2bfdb": "currentRequestId()",
	"07621eca": "currentReward()",
	"75ad1a2a": "currentTotalTips()",
	"19cae462": "difficulty()",
	"b9ce896b": "disputeFee()",
	"4ba0a5ee": "migrated(address)",
	"7f4ec4c3": "pending_owner()",
	"5badbe4c": "requestCount()",
	"2bf07e9e": "requestQPosition()",
	"03b3160f": "slotProgress()",
	"60c7dc47": "stakeAmount()",
	"dff69787": "stakerCount()",
	"dfee1ff1": "targetMiners()",
	"6fd4f229": "timeOfLastNewValue()",
	"6fc37811": "timeTarget()",
	"561cb04a": "totalTip()",
	"3940e9ee": "total_supply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"b59e14d4": "uints(bytes32)",
	"d67dcbc5": "updateBalanceAtNow(address,uint256)",
}

// TellorTransferBin is the compiled bytecode used for deploying new contracts.
var TellorTransferBin = "0x608060405234801561001057600080fd5b50610fab806100206000396000f3fe608060405234801561001057600080fd5b50600436106101f05760003560e01c8063699f200f1161010f578063a9059cbb116100a2578063d67dcbc511610071578063d67dcbc514610476578063dd62ed3e146104a4578063dfee1ff1146104d2578063dff69787146104da576101f0565b8063a9059cbb1461041d578063b2bdfa7b14610449578063b59e14d414610451578063b9ce896b1461046e576101f0565b806370a08231116100de57806370a08231146103bb57806375ad1a2a146103e15780637f4ec4c3146103e9578063999cf26c146103f1576101f0565b8063699f200f1461036a5780636e3cf885146103a35780636fc37811146103ab5780636fd4f229146103b3576101f0565b80634ba0a5ee116101875780635ae2bfdb116101565780635ae2bfdb146103355780635badbe4c1461033d57806360c7dc471461034557806362dd1d2a1461034d576101f0565b80634ba0a5ee146102d35780634ee2cd7e146102f957806351bdd58514610325578063561cb04a1461032d576101f0565b806319cae462116101c357806319cae4621461028557806323b872dd1461028d5780632bf07e9e146102c35780633940e9ee146102cb576101f0565b8063024c2ddd146101f557806303b3160f1461023557806307621eca1461023d578063095ea7b314610245575b600080fd5b6102236004803603604081101561020b57600080fd5b506001600160a01b03813581169160200135166104e2565b60408051918252519081900360200190f35b6102236104ff565b610223610523565b6102716004803603604081101561025b57600080fd5b506001600160a01b038135169060200135610547565b604080519115158252519081900360200190f35b610223610632565b610271600480360360608110156102a357600080fd5b506001600160a01b03813581169160208101359091169060400135610656565b610223610702565b610223610726565b610271600480360360208110156102e957600080fd5b50356001600160a01b031661074a565b6102236004803603604081101561030f57600080fd5b506001600160a01b03813516906020013561075f565b610223610903565b610223610927565b61022361094b565b61022361096f565b610223610993565b6102236004803603602081101561036357600080fd5b50356109b7565b6103876004803603602081101561038057600080fd5b50356109c9565b604080516001600160a01b039092168252519081900360200190f35b6102236109e4565b610223610a08565b610223610a2c565b610223600480360360208110156103d157600080fd5b50356001600160a01b0316610a50565b610223610a5c565b610223610a80565b6102716004803603604081101561040757600080fd5b506001600160a01b038135169060200135610aa4565b6102716004803603604081101561043357600080fd5b506001600160a01b038135169060200135610b67565b610223610b7d565b6102236004803603602081101561046757600080fd5b5035610ba1565b610223610bb3565b6104a26004803603604081101561048c57600080fd5b506001600160a01b038135169060200135610bd7565b005b610223600480360360408110156104ba57600080fd5b506001600160a01b0381358116916020013516610cd4565b610223610cff565b610223610d23565b604a60209081526000928352604080842090915290825290205481565b7f6c505cb2db6644f57b42d87bd9407b0f66788b07d0617a2bc1356a0e69e66f9a81565b7f9b6853911475b07474368644a0d922ee13bc76a15cd3e97d3e334326424a47d481565b6000336105855760405162461bcd60e51b8152600401808060200182810382526024815260200180610f526024913960400191505060405180910390fd5b6001600160a01b0383166105ca5760405162461bcd60e51b8152600401808060200182810382526022815260200180610ee86022913960400191505060405180910390fd5b336000818152604a602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060015b92915050565b7fb12aff7664b16cb99339be399b863feecd64d14817be7e1f042f97e3f358e64e81565b6001600160a01b0383166000908152604a602090815260408083203384529091528120548211156106c3576040805162461bcd60e51b8152602060048201526012602482015271416c6c6f77616e63652069732077726f6e6760701b604482015290519081900360640190fd5b6001600160a01b0384166000908152604a602090815260408083203384529091529020805483900390556106f8848484610d47565b5060019392505050565b7f1e344bd070f05f1c5b3f0b1266f4f20d837a0a8190a3a2da8b0375eac2ba86ea81565b7fb1557182e4359a1f0c6301278e8f5b35a776ab58d39892581e357578fb28783681565b604b6020526000908152604090205460ff1681565b6001600160a01b0382166000908152604960205260408120805415806107a55750828160008154811061078e57fe5b6000918252602090912001546001600160801b0316115b156107b457600091505061062c565b8054819060001981019081106107c657fe5b6000918252602090912001546001600160801b03168310610818578054819060001981019081106107f357fe5b600091825260209091200154600160801b90046001600160801b0316915061062c9050565b8054600090600119015b818111156108d057600060026001838501010490508584828154811061084457fe5b6000918252602090912001546001600160801b031614156108935783818154811061086b57fe5b600091825260209091200154600160801b90046001600160801b0316945061062c9350505050565b858482815481106108a057fe5b6000918252602090912001546001600160801b031610156108c3578092506108ca565b6001810391505b50610822565b8282815481106108dc57fe5b600091825260209091200154600160801b90046001600160801b0316935061062c92505050565b7f3375fb9157bb77048f607329b1c4d45487433f61e0f51bcdeba91d975b2dab1881565b7f2a9e355a92978430eca9c1aa3a9ba590094bac282594bccf82de16b83046e2c381565b7f7584d7d8701714da9c117f5bf30af73b0b88aca5338a84a21eb28de2fe0d93b881565b7f05de9147d05477c0a5dc675aeea733157f5092f82add148cf39d579cafe3dc9881565b7f7be108969d31a3f0b261465c71f2b0ba9301cd914d55d9091c3b36a49d4d41b281565b60486020526000908152604090205481565b6047602052600090815260409020546001600160a01b031681565b7f969ea04b74d02bb4d9e6e8e57236e1b9ca31627139ae9f0e465249932e82450281565b7fad16221efc80aaf1b7e69bd3ecb61ba5ffa539adf129c3b4ffff769c9b5bbc3381565b7f97e6eb29f6a85471f7cc9b57f9e4c3deaf398cfc9798673160d7798baf0b13a481565b600061062c824361075f565b7fd26d9834adf5a73309c4974bf654850bb699df8505e70d4cfde365c417b19dfc81565b7f44b2657a0f8a90ed8e62f4c4cceca06eacaa9b4b25751ae1ebca9280a70abd6881565b6001600160a01b03821660009081526044602052604081205415801590610ae357506001600160a01b0383166000908152604460205260409020546005115b15610b54577f7be108969d31a3f0b261465c71f2b0ba9301cd914d55d9091c3b36a49d4d41b260005260466020527ff0cbb0eecd83f344371d44a7d44097662e277de50a2e30b5e36df1aa56be6de5548290610b3e85610a50565b0310610b4c5750600161062c565b50600061062c565b81610b5e84610a50565b10159392505050565b6000610b74338484610d47565b50600192915050565b7f9dbc393ddc18fd27b1d9b1b129059925688d2f2d5818a5ec3ebb750b7c286ea681565b60466020526000908152604090205481565b7f8b75eb45d88e80f0e4ec77d23936268694c0e7ac2e0c9085c5c6bdfcfbc4923981565b6001600160a01b038216600090815260496020526040902080541580610c2457508054439082906000198101908110610c0c57fe5b6000918252602090912001546001600160801b031614155b15610c955760408051808201909152436001600160801b0390811682528381166020808401918252845460018101865560008681529190912093519301805491516fffffffffffffffffffffffffffffffff19909216938316939093178216600160801b9190921602179055610ccf565b805460009082906000198101908110610caa57fe5b600091825260209091200180546001600160801b03808616600160801b029116179055505b505050565b6001600160a01b039182166000908152604a6020908152604080832093909416825291909152205490565b7fabef544d8048318ece54fb2c6385255cd1b06e176525d149a0338a7acca6deb381565b7fedddb9344bfe0dadc78c558b8ffca446679cbffc17be64eb83973fce7bea5f3481565b80610d835760405162461bcd60e51b8152600401808060200182810382526021815260200180610f316021913960400191505060405180910390fd5b6001600160a01b038216610dd6576040805162461bcd60e51b815260206004820152601560248201527452656365697665722069732030206164647265737360581b604482015290519081900360640190fd5b610de08382610aa4565b610e1b5760405162461bcd60e51b8152600401808060200182810382526027815260200180610f0a6027913960400191505060405180910390fd5b6000610e2684610a50565b9050610e3484838303610bd7565b610e3d83610a50565b9050808282011015610e8a576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b610e9683838301610bd7565b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a35050505056fe45524332303a20617070726f766520746f20746865207a65726f206164647265737353686f756c6420686176652073756666696369656e742062616c616e636520746f207472616465547269656420746f2073656e64206e6f6e2d706f73697469766520616d6f756e7445524332303a20617070726f76652066726f6d20746865207a65726f2061646472657373a2646970667358221220dd53437e7de7dd01138ec58d0b48dc598e4da130c4eab3ed2b1efb5b48d0fb0564736f6c63430007040033"

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

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) Owner(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) Owner() ([32]byte, error) {
	return _TellorTransfer.Contract.Owner(&_TellorTransfer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) Owner() ([32]byte, error) {
	return _TellorTransfer.Contract.Owner(&_TellorTransfer.CallOpts)
}

// TBlock is a free data retrieval call binding the contract method 0x6e3cf885.
//
// Solidity: function _tBlock() view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) TBlock(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "_tBlock")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TBlock is a free data retrieval call binding the contract method 0x6e3cf885.
//
// Solidity: function _tBlock() view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) TBlock() ([32]byte, error) {
	return _TellorTransfer.Contract.TBlock(&_TellorTransfer.CallOpts)
}

// TBlock is a free data retrieval call binding the contract method 0x6e3cf885.
//
// Solidity: function _tBlock() view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) TBlock() ([32]byte, error) {
	return _TellorTransfer.Contract.TBlock(&_TellorTransfer.CallOpts)
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

// CurrentChallenge is a free data retrieval call binding the contract method 0x51bdd585.
//
// Solidity: function currentChallenge() view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) CurrentChallenge(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "currentChallenge")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentChallenge is a free data retrieval call binding the contract method 0x51bdd585.
//
// Solidity: function currentChallenge() view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) CurrentChallenge() ([32]byte, error) {
	return _TellorTransfer.Contract.CurrentChallenge(&_TellorTransfer.CallOpts)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x51bdd585.
//
// Solidity: function currentChallenge() view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) CurrentChallenge() ([32]byte, error) {
	return _TellorTransfer.Contract.CurrentChallenge(&_TellorTransfer.CallOpts)
}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5ae2bfdb.
//
// Solidity: function currentRequestId() view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) CurrentRequestId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "currentRequestId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5ae2bfdb.
//
// Solidity: function currentRequestId() view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) CurrentRequestId() ([32]byte, error) {
	return _TellorTransfer.Contract.CurrentRequestId(&_TellorTransfer.CallOpts)
}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5ae2bfdb.
//
// Solidity: function currentRequestId() view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) CurrentRequestId() ([32]byte, error) {
	return _TellorTransfer.Contract.CurrentRequestId(&_TellorTransfer.CallOpts)
}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) CurrentReward(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "currentReward")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) CurrentReward() ([32]byte, error) {
	return _TellorTransfer.Contract.CurrentReward(&_TellorTransfer.CallOpts)
}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) CurrentReward() ([32]byte, error) {
	return _TellorTransfer.Contract.CurrentReward(&_TellorTransfer.CallOpts)
}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) CurrentTotalTips(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "currentTotalTips")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) CurrentTotalTips() ([32]byte, error) {
	return _TellorTransfer.Contract.CurrentTotalTips(&_TellorTransfer.CallOpts)
}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) CurrentTotalTips() ([32]byte, error) {
	return _TellorTransfer.Contract.CurrentTotalTips(&_TellorTransfer.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) Difficulty(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "difficulty")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) Difficulty() ([32]byte, error) {
	return _TellorTransfer.Contract.Difficulty(&_TellorTransfer.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) Difficulty() ([32]byte, error) {
	return _TellorTransfer.Contract.Difficulty(&_TellorTransfer.CallOpts)
}

// DisputeFee is a free data retrieval call binding the contract method 0xb9ce896b.
//
// Solidity: function disputeFee() view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) DisputeFee(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "disputeFee")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DisputeFee is a free data retrieval call binding the contract method 0xb9ce896b.
//
// Solidity: function disputeFee() view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) DisputeFee() ([32]byte, error) {
	return _TellorTransfer.Contract.DisputeFee(&_TellorTransfer.CallOpts)
}

// DisputeFee is a free data retrieval call binding the contract method 0xb9ce896b.
//
// Solidity: function disputeFee() view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) DisputeFee() ([32]byte, error) {
	return _TellorTransfer.Contract.DisputeFee(&_TellorTransfer.CallOpts)
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

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) PendingOwner(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "pending_owner")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) PendingOwner() ([32]byte, error) {
	return _TellorTransfer.Contract.PendingOwner(&_TellorTransfer.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) PendingOwner() ([32]byte, error) {
	return _TellorTransfer.Contract.PendingOwner(&_TellorTransfer.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) RequestCount(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "requestCount")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) RequestCount() ([32]byte, error) {
	return _TellorTransfer.Contract.RequestCount(&_TellorTransfer.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) RequestCount() ([32]byte, error) {
	return _TellorTransfer.Contract.RequestCount(&_TellorTransfer.CallOpts)
}

// RequestQPosition is a free data retrieval call binding the contract method 0x2bf07e9e.
//
// Solidity: function requestQPosition() view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) RequestQPosition(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "requestQPosition")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RequestQPosition is a free data retrieval call binding the contract method 0x2bf07e9e.
//
// Solidity: function requestQPosition() view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) RequestQPosition() ([32]byte, error) {
	return _TellorTransfer.Contract.RequestQPosition(&_TellorTransfer.CallOpts)
}

// RequestQPosition is a free data retrieval call binding the contract method 0x2bf07e9e.
//
// Solidity: function requestQPosition() view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) RequestQPosition() ([32]byte, error) {
	return _TellorTransfer.Contract.RequestQPosition(&_TellorTransfer.CallOpts)
}

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) SlotProgress(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "slotProgress")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) SlotProgress() ([32]byte, error) {
	return _TellorTransfer.Contract.SlotProgress(&_TellorTransfer.CallOpts)
}

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) SlotProgress() ([32]byte, error) {
	return _TellorTransfer.Contract.SlotProgress(&_TellorTransfer.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) StakeAmount(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "stakeAmount")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) StakeAmount() ([32]byte, error) {
	return _TellorTransfer.Contract.StakeAmount(&_TellorTransfer.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) StakeAmount() ([32]byte, error) {
	return _TellorTransfer.Contract.StakeAmount(&_TellorTransfer.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) StakerCount(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) StakerCount() ([32]byte, error) {
	return _TellorTransfer.Contract.StakerCount(&_TellorTransfer.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) StakerCount() ([32]byte, error) {
	return _TellorTransfer.Contract.StakerCount(&_TellorTransfer.CallOpts)
}

// TargetMiners is a free data retrieval call binding the contract method 0xdfee1ff1.
//
// Solidity: function targetMiners() view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) TargetMiners(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "targetMiners")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TargetMiners is a free data retrieval call binding the contract method 0xdfee1ff1.
//
// Solidity: function targetMiners() view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) TargetMiners() ([32]byte, error) {
	return _TellorTransfer.Contract.TargetMiners(&_TellorTransfer.CallOpts)
}

// TargetMiners is a free data retrieval call binding the contract method 0xdfee1ff1.
//
// Solidity: function targetMiners() view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) TargetMiners() ([32]byte, error) {
	return _TellorTransfer.Contract.TargetMiners(&_TellorTransfer.CallOpts)
}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) TimeOfLastNewValue(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "timeOfLastNewValue")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) TimeOfLastNewValue() ([32]byte, error) {
	return _TellorTransfer.Contract.TimeOfLastNewValue(&_TellorTransfer.CallOpts)
}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) TimeOfLastNewValue() ([32]byte, error) {
	return _TellorTransfer.Contract.TimeOfLastNewValue(&_TellorTransfer.CallOpts)
}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) TimeTarget(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "timeTarget")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) TimeTarget() ([32]byte, error) {
	return _TellorTransfer.Contract.TimeTarget(&_TellorTransfer.CallOpts)
}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) TimeTarget() ([32]byte, error) {
	return _TellorTransfer.Contract.TimeTarget(&_TellorTransfer.CallOpts)
}

// TotalTip is a free data retrieval call binding the contract method 0x561cb04a.
//
// Solidity: function totalTip() view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) TotalTip(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "totalTip")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TotalTip is a free data retrieval call binding the contract method 0x561cb04a.
//
// Solidity: function totalTip() view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) TotalTip() ([32]byte, error) {
	return _TellorTransfer.Contract.TotalTip(&_TellorTransfer.CallOpts)
}

// TotalTip is a free data retrieval call binding the contract method 0x561cb04a.
//
// Solidity: function totalTip() view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) TotalTip() ([32]byte, error) {
	return _TellorTransfer.Contract.TotalTip(&_TellorTransfer.CallOpts)
}

// TotalSupplyVar is a free data retrieval call binding the contract method 0x3940e9ee.
//
// Solidity: function total_supply() view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) TotalSupplyVar(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "total_supply")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TotalSupplyVar is a free data retrieval call binding the contract method 0x3940e9ee.
//
// Solidity: function total_supply() view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) TotalSupplyVar() ([32]byte, error) {
	return _TellorTransfer.Contract.TotalSupplyVar(&_TellorTransfer.CallOpts)
}

// TotalSupplyVar is a free data retrieval call binding the contract method 0x3940e9ee.
//
// Solidity: function total_supply() view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) TotalSupplyVar() ([32]byte, error) {
	return _TellorTransfer.Contract.TotalSupplyVar(&_TellorTransfer.CallOpts)
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

// UpdateBalanceAtNow is a paid mutator transaction binding the contract method 0xd67dcbc5.
//
// Solidity: function updateBalanceAtNow(address _user, uint256 _value) returns()
func (_TellorTransfer *TellorTransferTransactor) UpdateBalanceAtNow(opts *bind.TransactOpts, _user common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.contract.Transact(opts, "updateBalanceAtNow", _user, _value)
}

// UpdateBalanceAtNow is a paid mutator transaction binding the contract method 0xd67dcbc5.
//
// Solidity: function updateBalanceAtNow(address _user, uint256 _value) returns()
func (_TellorTransfer *TellorTransferSession) UpdateBalanceAtNow(_user common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.Contract.UpdateBalanceAtNow(&_TellorTransfer.TransactOpts, _user, _value)
}

// UpdateBalanceAtNow is a paid mutator transaction binding the contract method 0xd67dcbc5.
//
// Solidity: function updateBalanceAtNow(address _user, uint256 _value) returns()
func (_TellorTransfer *TellorTransferTransactorSession) UpdateBalanceAtNow(_user common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.Contract.UpdateBalanceAtNow(&_TellorTransfer.TransactOpts, _user, _value)
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

// TellorTransferTransferEventIterator is returned from FilterTransferEvent and is used to iterate over the raw logs and unpacked data for TransferEvent events raised by the TellorTransfer contract.
type TellorTransferTransferEventIterator struct {
	Event *TellorTransferTransferEvent // Event containing the contract specifics and raw log

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
func (it *TellorTransferTransferEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorTransferTransferEvent)
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
		it.Event = new(TellorTransferTransferEvent)
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
func (it *TellorTransferTransferEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorTransferTransferEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorTransferTransferEvent represents a TransferEvent event raised by the TellorTransfer contract.
type TellorTransferTransferEvent struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferEvent is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TellorTransfer *TellorTransferFilterer) FilterTransferEvent(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TellorTransferTransferEventIterator, error) {

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
	return &TellorTransferTransferEventIterator{contract: _TellorTransfer.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransferEvent is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TellorTransfer *TellorTransferFilterer) WatchTransferEvent(opts *bind.WatchOpts, sink chan<- *TellorTransferTransferEvent, from []common.Address, to []common.Address) (event.Subscription, error) {

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
				event := new(TellorTransferTransferEvent)
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

// ParseTransferEvent is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TellorTransfer *TellorTransferFilterer) ParseTransferEvent(log types.Log) (*TellorTransferTransferEvent, error) {
	event := new(TellorTransferTransferEvent)
	if err := _TellorTransfer.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorVariablesABI is the input ABI used to generate the binding from.
const TellorVariablesABI = "[{\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_tBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentChallenge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRequestId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentReward\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTotalTips\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"difficulty\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disputeFee\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pending_owner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestCount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestQPosition\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slotProgress\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeAmount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetMiners\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeOfLastNewValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeTarget\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTip\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"total_supply\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TellorVariablesFuncSigs maps the 4-byte function signature to its string representation.
var TellorVariablesFuncSigs = map[string]string{
	"b2bdfa7b": "_owner()",
	"6e3cf885": "_tBlock()",
	"51bdd585": "currentChallenge()",
	"5ae2bfdb": "currentRequestId()",
	"07621eca": "currentReward()",
	"75ad1a2a": "currentTotalTips()",
	"19cae462": "difficulty()",
	"b9ce896b": "disputeFee()",
	"7f4ec4c3": "pending_owner()",
	"5badbe4c": "requestCount()",
	"2bf07e9e": "requestQPosition()",
	"03b3160f": "slotProgress()",
	"60c7dc47": "stakeAmount()",
	"dff69787": "stakerCount()",
	"dfee1ff1": "targetMiners()",
	"6fd4f229": "timeOfLastNewValue()",
	"6fc37811": "timeTarget()",
	"561cb04a": "totalTip()",
	"3940e9ee": "total_supply()",
}

// TellorVariablesBin is the compiled bytecode used for deploying new contracts.
var TellorVariablesBin = "0x608060405234801561001057600080fd5b506104b2806100206000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c806360c7dc47116100ad5780637f4ec4c3116100715780637f4ec4c3146101a8578063b2bdfa7b146101b0578063b9ce896b146101b8578063dfee1ff1146101c0578063dff69787146101c857610121565b806360c7dc47146101805780636e3cf885146101885780636fc37811146101905780636fd4f2291461019857806375ad1a2a146101a057610121565b80633940e9ee116100f45780633940e9ee1461015857806351bdd58514610160578063561cb04a146101685780635ae2bfdb146101705780635badbe4c1461017857610121565b806303b3160f1461012657806307621eca1461014057806319cae462146101485780632bf07e9e14610150575b600080fd5b61012e6101d0565b60408051918252519081900360200190f35b61012e6101f4565b61012e610218565b61012e61023c565b61012e610260565b61012e610284565b61012e6102a8565b61012e6102cc565b61012e6102f0565b61012e610314565b61012e610338565b61012e61035c565b61012e610380565b61012e6103a4565b61012e6103c8565b61012e6103ec565b61012e610410565b61012e610434565b61012e610458565b7f6c505cb2db6644f57b42d87bd9407b0f66788b07d0617a2bc1356a0e69e66f9a81565b7f9b6853911475b07474368644a0d922ee13bc76a15cd3e97d3e334326424a47d481565b7fb12aff7664b16cb99339be399b863feecd64d14817be7e1f042f97e3f358e64e81565b7f1e344bd070f05f1c5b3f0b1266f4f20d837a0a8190a3a2da8b0375eac2ba86ea81565b7fb1557182e4359a1f0c6301278e8f5b35a776ab58d39892581e357578fb28783681565b7f3375fb9157bb77048f607329b1c4d45487433f61e0f51bcdeba91d975b2dab1881565b7f2a9e355a92978430eca9c1aa3a9ba590094bac282594bccf82de16b83046e2c381565b7f7584d7d8701714da9c117f5bf30af73b0b88aca5338a84a21eb28de2fe0d93b881565b7f05de9147d05477c0a5dc675aeea733157f5092f82add148cf39d579cafe3dc9881565b7f7be108969d31a3f0b261465c71f2b0ba9301cd914d55d9091c3b36a49d4d41b281565b7f969ea04b74d02bb4d9e6e8e57236e1b9ca31627139ae9f0e465249932e82450281565b7fad16221efc80aaf1b7e69bd3ecb61ba5ffa539adf129c3b4ffff769c9b5bbc3381565b7f97e6eb29f6a85471f7cc9b57f9e4c3deaf398cfc9798673160d7798baf0b13a481565b7fd26d9834adf5a73309c4974bf654850bb699df8505e70d4cfde365c417b19dfc81565b7f44b2657a0f8a90ed8e62f4c4cceca06eacaa9b4b25751ae1ebca9280a70abd6881565b7f9dbc393ddc18fd27b1d9b1b129059925688d2f2d5818a5ec3ebb750b7c286ea681565b7f8b75eb45d88e80f0e4ec77d23936268694c0e7ac2e0c9085c5c6bdfcfbc4923981565b7fabef544d8048318ece54fb2c6385255cd1b06e176525d149a0338a7acca6deb381565b7fedddb9344bfe0dadc78c558b8ffca446679cbffc17be64eb83973fce7bea5f348156fea2646970667358221220871e4ea7a34b8cf05343a3b479c8276df19b05056c352b723862737871f989ba64736f6c63430007040033"

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

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(bytes32)
func (_TellorVariables *TellorVariablesCaller) Owner(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorVariables.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(bytes32)
func (_TellorVariables *TellorVariablesSession) Owner() ([32]byte, error) {
	return _TellorVariables.Contract.Owner(&_TellorVariables.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(bytes32)
func (_TellorVariables *TellorVariablesCallerSession) Owner() ([32]byte, error) {
	return _TellorVariables.Contract.Owner(&_TellorVariables.CallOpts)
}

// TBlock is a free data retrieval call binding the contract method 0x6e3cf885.
//
// Solidity: function _tBlock() view returns(bytes32)
func (_TellorVariables *TellorVariablesCaller) TBlock(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorVariables.contract.Call(opts, &out, "_tBlock")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TBlock is a free data retrieval call binding the contract method 0x6e3cf885.
//
// Solidity: function _tBlock() view returns(bytes32)
func (_TellorVariables *TellorVariablesSession) TBlock() ([32]byte, error) {
	return _TellorVariables.Contract.TBlock(&_TellorVariables.CallOpts)
}

// TBlock is a free data retrieval call binding the contract method 0x6e3cf885.
//
// Solidity: function _tBlock() view returns(bytes32)
func (_TellorVariables *TellorVariablesCallerSession) TBlock() ([32]byte, error) {
	return _TellorVariables.Contract.TBlock(&_TellorVariables.CallOpts)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x51bdd585.
//
// Solidity: function currentChallenge() view returns(bytes32)
func (_TellorVariables *TellorVariablesCaller) CurrentChallenge(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorVariables.contract.Call(opts, &out, "currentChallenge")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentChallenge is a free data retrieval call binding the contract method 0x51bdd585.
//
// Solidity: function currentChallenge() view returns(bytes32)
func (_TellorVariables *TellorVariablesSession) CurrentChallenge() ([32]byte, error) {
	return _TellorVariables.Contract.CurrentChallenge(&_TellorVariables.CallOpts)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x51bdd585.
//
// Solidity: function currentChallenge() view returns(bytes32)
func (_TellorVariables *TellorVariablesCallerSession) CurrentChallenge() ([32]byte, error) {
	return _TellorVariables.Contract.CurrentChallenge(&_TellorVariables.CallOpts)
}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5ae2bfdb.
//
// Solidity: function currentRequestId() view returns(bytes32)
func (_TellorVariables *TellorVariablesCaller) CurrentRequestId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorVariables.contract.Call(opts, &out, "currentRequestId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5ae2bfdb.
//
// Solidity: function currentRequestId() view returns(bytes32)
func (_TellorVariables *TellorVariablesSession) CurrentRequestId() ([32]byte, error) {
	return _TellorVariables.Contract.CurrentRequestId(&_TellorVariables.CallOpts)
}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5ae2bfdb.
//
// Solidity: function currentRequestId() view returns(bytes32)
func (_TellorVariables *TellorVariablesCallerSession) CurrentRequestId() ([32]byte, error) {
	return _TellorVariables.Contract.CurrentRequestId(&_TellorVariables.CallOpts)
}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(bytes32)
func (_TellorVariables *TellorVariablesCaller) CurrentReward(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorVariables.contract.Call(opts, &out, "currentReward")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(bytes32)
func (_TellorVariables *TellorVariablesSession) CurrentReward() ([32]byte, error) {
	return _TellorVariables.Contract.CurrentReward(&_TellorVariables.CallOpts)
}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(bytes32)
func (_TellorVariables *TellorVariablesCallerSession) CurrentReward() ([32]byte, error) {
	return _TellorVariables.Contract.CurrentReward(&_TellorVariables.CallOpts)
}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(bytes32)
func (_TellorVariables *TellorVariablesCaller) CurrentTotalTips(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorVariables.contract.Call(opts, &out, "currentTotalTips")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(bytes32)
func (_TellorVariables *TellorVariablesSession) CurrentTotalTips() ([32]byte, error) {
	return _TellorVariables.Contract.CurrentTotalTips(&_TellorVariables.CallOpts)
}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(bytes32)
func (_TellorVariables *TellorVariablesCallerSession) CurrentTotalTips() ([32]byte, error) {
	return _TellorVariables.Contract.CurrentTotalTips(&_TellorVariables.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(bytes32)
func (_TellorVariables *TellorVariablesCaller) Difficulty(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorVariables.contract.Call(opts, &out, "difficulty")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(bytes32)
func (_TellorVariables *TellorVariablesSession) Difficulty() ([32]byte, error) {
	return _TellorVariables.Contract.Difficulty(&_TellorVariables.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(bytes32)
func (_TellorVariables *TellorVariablesCallerSession) Difficulty() ([32]byte, error) {
	return _TellorVariables.Contract.Difficulty(&_TellorVariables.CallOpts)
}

// DisputeFee is a free data retrieval call binding the contract method 0xb9ce896b.
//
// Solidity: function disputeFee() view returns(bytes32)
func (_TellorVariables *TellorVariablesCaller) DisputeFee(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorVariables.contract.Call(opts, &out, "disputeFee")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DisputeFee is a free data retrieval call binding the contract method 0xb9ce896b.
//
// Solidity: function disputeFee() view returns(bytes32)
func (_TellorVariables *TellorVariablesSession) DisputeFee() ([32]byte, error) {
	return _TellorVariables.Contract.DisputeFee(&_TellorVariables.CallOpts)
}

// DisputeFee is a free data retrieval call binding the contract method 0xb9ce896b.
//
// Solidity: function disputeFee() view returns(bytes32)
func (_TellorVariables *TellorVariablesCallerSession) DisputeFee() ([32]byte, error) {
	return _TellorVariables.Contract.DisputeFee(&_TellorVariables.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(bytes32)
func (_TellorVariables *TellorVariablesCaller) PendingOwner(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorVariables.contract.Call(opts, &out, "pending_owner")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(bytes32)
func (_TellorVariables *TellorVariablesSession) PendingOwner() ([32]byte, error) {
	return _TellorVariables.Contract.PendingOwner(&_TellorVariables.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(bytes32)
func (_TellorVariables *TellorVariablesCallerSession) PendingOwner() ([32]byte, error) {
	return _TellorVariables.Contract.PendingOwner(&_TellorVariables.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(bytes32)
func (_TellorVariables *TellorVariablesCaller) RequestCount(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorVariables.contract.Call(opts, &out, "requestCount")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(bytes32)
func (_TellorVariables *TellorVariablesSession) RequestCount() ([32]byte, error) {
	return _TellorVariables.Contract.RequestCount(&_TellorVariables.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(bytes32)
func (_TellorVariables *TellorVariablesCallerSession) RequestCount() ([32]byte, error) {
	return _TellorVariables.Contract.RequestCount(&_TellorVariables.CallOpts)
}

// RequestQPosition is a free data retrieval call binding the contract method 0x2bf07e9e.
//
// Solidity: function requestQPosition() view returns(bytes32)
func (_TellorVariables *TellorVariablesCaller) RequestQPosition(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorVariables.contract.Call(opts, &out, "requestQPosition")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RequestQPosition is a free data retrieval call binding the contract method 0x2bf07e9e.
//
// Solidity: function requestQPosition() view returns(bytes32)
func (_TellorVariables *TellorVariablesSession) RequestQPosition() ([32]byte, error) {
	return _TellorVariables.Contract.RequestQPosition(&_TellorVariables.CallOpts)
}

// RequestQPosition is a free data retrieval call binding the contract method 0x2bf07e9e.
//
// Solidity: function requestQPosition() view returns(bytes32)
func (_TellorVariables *TellorVariablesCallerSession) RequestQPosition() ([32]byte, error) {
	return _TellorVariables.Contract.RequestQPosition(&_TellorVariables.CallOpts)
}

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(bytes32)
func (_TellorVariables *TellorVariablesCaller) SlotProgress(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorVariables.contract.Call(opts, &out, "slotProgress")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(bytes32)
func (_TellorVariables *TellorVariablesSession) SlotProgress() ([32]byte, error) {
	return _TellorVariables.Contract.SlotProgress(&_TellorVariables.CallOpts)
}

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(bytes32)
func (_TellorVariables *TellorVariablesCallerSession) SlotProgress() ([32]byte, error) {
	return _TellorVariables.Contract.SlotProgress(&_TellorVariables.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(bytes32)
func (_TellorVariables *TellorVariablesCaller) StakeAmount(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorVariables.contract.Call(opts, &out, "stakeAmount")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(bytes32)
func (_TellorVariables *TellorVariablesSession) StakeAmount() ([32]byte, error) {
	return _TellorVariables.Contract.StakeAmount(&_TellorVariables.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(bytes32)
func (_TellorVariables *TellorVariablesCallerSession) StakeAmount() ([32]byte, error) {
	return _TellorVariables.Contract.StakeAmount(&_TellorVariables.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(bytes32)
func (_TellorVariables *TellorVariablesCaller) StakerCount(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorVariables.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(bytes32)
func (_TellorVariables *TellorVariablesSession) StakerCount() ([32]byte, error) {
	return _TellorVariables.Contract.StakerCount(&_TellorVariables.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(bytes32)
func (_TellorVariables *TellorVariablesCallerSession) StakerCount() ([32]byte, error) {
	return _TellorVariables.Contract.StakerCount(&_TellorVariables.CallOpts)
}

// TargetMiners is a free data retrieval call binding the contract method 0xdfee1ff1.
//
// Solidity: function targetMiners() view returns(bytes32)
func (_TellorVariables *TellorVariablesCaller) TargetMiners(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorVariables.contract.Call(opts, &out, "targetMiners")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TargetMiners is a free data retrieval call binding the contract method 0xdfee1ff1.
//
// Solidity: function targetMiners() view returns(bytes32)
func (_TellorVariables *TellorVariablesSession) TargetMiners() ([32]byte, error) {
	return _TellorVariables.Contract.TargetMiners(&_TellorVariables.CallOpts)
}

// TargetMiners is a free data retrieval call binding the contract method 0xdfee1ff1.
//
// Solidity: function targetMiners() view returns(bytes32)
func (_TellorVariables *TellorVariablesCallerSession) TargetMiners() ([32]byte, error) {
	return _TellorVariables.Contract.TargetMiners(&_TellorVariables.CallOpts)
}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(bytes32)
func (_TellorVariables *TellorVariablesCaller) TimeOfLastNewValue(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorVariables.contract.Call(opts, &out, "timeOfLastNewValue")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(bytes32)
func (_TellorVariables *TellorVariablesSession) TimeOfLastNewValue() ([32]byte, error) {
	return _TellorVariables.Contract.TimeOfLastNewValue(&_TellorVariables.CallOpts)
}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(bytes32)
func (_TellorVariables *TellorVariablesCallerSession) TimeOfLastNewValue() ([32]byte, error) {
	return _TellorVariables.Contract.TimeOfLastNewValue(&_TellorVariables.CallOpts)
}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(bytes32)
func (_TellorVariables *TellorVariablesCaller) TimeTarget(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorVariables.contract.Call(opts, &out, "timeTarget")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(bytes32)
func (_TellorVariables *TellorVariablesSession) TimeTarget() ([32]byte, error) {
	return _TellorVariables.Contract.TimeTarget(&_TellorVariables.CallOpts)
}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(bytes32)
func (_TellorVariables *TellorVariablesCallerSession) TimeTarget() ([32]byte, error) {
	return _TellorVariables.Contract.TimeTarget(&_TellorVariables.CallOpts)
}

// TotalTip is a free data retrieval call binding the contract method 0x561cb04a.
//
// Solidity: function totalTip() view returns(bytes32)
func (_TellorVariables *TellorVariablesCaller) TotalTip(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorVariables.contract.Call(opts, &out, "totalTip")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TotalTip is a free data retrieval call binding the contract method 0x561cb04a.
//
// Solidity: function totalTip() view returns(bytes32)
func (_TellorVariables *TellorVariablesSession) TotalTip() ([32]byte, error) {
	return _TellorVariables.Contract.TotalTip(&_TellorVariables.CallOpts)
}

// TotalTip is a free data retrieval call binding the contract method 0x561cb04a.
//
// Solidity: function totalTip() view returns(bytes32)
func (_TellorVariables *TellorVariablesCallerSession) TotalTip() ([32]byte, error) {
	return _TellorVariables.Contract.TotalTip(&_TellorVariables.CallOpts)
}

// TotalSupplyVar is a free data retrieval call binding the contract method 0x3940e9ee.
//
// Solidity: function total_supply() view returns(bytes32)
func (_TellorVariables *TellorVariablesCaller) TotalSupplyVar(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TellorVariables.contract.Call(opts, &out, "total_supply")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TotalSupplyVar is a free data retrieval call binding the contract method 0x3940e9ee.
//
// Solidity: function total_supply() view returns(bytes32)
func (_TellorVariables *TellorVariablesSession) TotalSupplyVar() ([32]byte, error) {
	return _TellorVariables.Contract.TotalSupplyVar(&_TellorVariables.CallOpts)
}

// TotalSupplyVar is a free data retrieval call binding the contract method 0x3940e9ee.
//
// Solidity: function total_supply() view returns(bytes32)
func (_TellorVariables *TellorVariablesCallerSession) TotalSupplyVar() ([32]byte, error) {
	return _TellorVariables.Contract.TotalSupplyVar(&_TellorVariables.CallOpts)
}

// UtilitiesABI is the input ABI used to generate the binding from.
const UtilitiesABI = "[{\"inputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"data\",\"type\":\"uint256[51]\"}],\"name\":\"getMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"data\",\"type\":\"uint256[51]\"}],\"name\":\"getMax5\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"max\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"maxIndex\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"data\",\"type\":\"uint256[51]\"}],\"name\":\"getMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// UtilitiesFuncSigs maps the 4-byte function signature to its string representation.
var UtilitiesFuncSigs = map[string]string{
	"c87a336d": "getMax(uint256[51])",
	"99830e32": "getMax5(uint256[51])",
	"f29e5e9a": "getMin(uint256[51])",
}

// UtilitiesBin is the compiled bytecode used for deploying new contracts.
var UtilitiesBin = "0x608060405234801561001057600080fd5b50610400806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806399830e3214610046578063c87a336d146100fe578063f29e5e9a1461016b575b600080fd5b61009a600480360361066081101561005d57600080fd5b8101908080610660019060338060200260405190810160405280929190826033602002808284376000920191909152509194506101bf9350505050565b604051808360a080838360005b838110156100bf5781810151838201526020016100a7565b5050505090500182600560200280838360005b838110156100ea5781810151838201526020016100d2565b505050509050019250505060405180910390f35b610152600480360361066081101561011557600080fd5b8101908080610660019060338060200260405190810160405280929190826033602002808284376000920191909152509194506103139350505050565b6040805192835260208301919091528051918290030190f35b610152600480360361066081101561018257600080fd5b8101908080610660019060338060200260405190810160405280929190826033602002808284376000920191909152509194506103639350505050565b6101c76103ac565b6101cf6103ac565b60208301516000805b6005811015610252578581600101603381106101f057fe5b602002015185826005811061020157fe5b60200201526001810184826005811061021657fe5b60200201528285826005811061022857fe5b6020020151101561024a5784816005811061023f57fe5b602002015192508091505b6001016101d8565b5060065b603381101561030b578286826033811061026c57fe5b602002015111156103035785816033811061028357fe5b602002015185836005811061029457fe5b6020020152808483600581106102a657fe5b60200201528581603381106102b757fe5b6020020151925060005b600581101561030157838682600581106102d757fe5b602002015110156102f9578581600581106102ee57fe5b602002015193508092505b6001016102c1565b505b600101610256565b505050915091565b6020810151600160025b603381101561035d578284826033811061033357fe5b602002015111156103555783816033811061034a57fe5b602002015192508091505b60010161031d565b50915091565b610640810151603260315b801561035d578284826033811061038157fe5b602002015110156103a35783816033811061039857fe5b602002015192508091505b6000190161036e565b6040518060a00160405280600590602082028036833750919291505056fea2646970667358221220ad8a6db266a42dd45eeedc3c1b99b2d4136dc6268fefb9d8809fe2b7968b41b964736f6c63430007040033"

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

// GetMax is a free data retrieval call binding the contract method 0xc87a336d.
//
// Solidity: function getMax(uint256[51] data) pure returns(uint256 max, uint256 maxIndex)
func (_Utilities *UtilitiesCaller) GetMax(opts *bind.CallOpts, data [51]*big.Int) (struct {
	Max      *big.Int
	MaxIndex *big.Int
}, error) {
	var out []interface{}
	err := _Utilities.contract.Call(opts, &out, "getMax", data)

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
// Solidity: function getMax(uint256[51] data) pure returns(uint256 max, uint256 maxIndex)
func (_Utilities *UtilitiesSession) GetMax(data [51]*big.Int) (struct {
	Max      *big.Int
	MaxIndex *big.Int
}, error) {
	return _Utilities.Contract.GetMax(&_Utilities.CallOpts, data)
}

// GetMax is a free data retrieval call binding the contract method 0xc87a336d.
//
// Solidity: function getMax(uint256[51] data) pure returns(uint256 max, uint256 maxIndex)
func (_Utilities *UtilitiesCallerSession) GetMax(data [51]*big.Int) (struct {
	Max      *big.Int
	MaxIndex *big.Int
}, error) {
	return _Utilities.Contract.GetMax(&_Utilities.CallOpts, data)
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

// GetMin is a free data retrieval call binding the contract method 0xf29e5e9a.
//
// Solidity: function getMin(uint256[51] data) pure returns(uint256 min, uint256 minIndex)
func (_Utilities *UtilitiesCaller) GetMin(opts *bind.CallOpts, data [51]*big.Int) (struct {
	Min      *big.Int
	MinIndex *big.Int
}, error) {
	var out []interface{}
	err := _Utilities.contract.Call(opts, &out, "getMin", data)

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
// Solidity: function getMin(uint256[51] data) pure returns(uint256 min, uint256 minIndex)
func (_Utilities *UtilitiesSession) GetMin(data [51]*big.Int) (struct {
	Min      *big.Int
	MinIndex *big.Int
}, error) {
	return _Utilities.Contract.GetMin(&_Utilities.CallOpts, data)
}

// GetMin is a free data retrieval call binding the contract method 0xf29e5e9a.
//
// Solidity: function getMin(uint256[51] data) pure returns(uint256 min, uint256 minIndex)
func (_Utilities *UtilitiesCallerSession) GetMin(data [51]*big.Int) (struct {
	Min      *big.Int
	MinIndex *big.Int
}, error) {
	return _Utilities.Contract.GetMin(&_Utilities.CallOpts, data)
}
