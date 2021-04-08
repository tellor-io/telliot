// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/apiOracle"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/tracker/dispute"
	"github.com/tellor-io/telliot/pkg/util"
)

/**
 * This file handles all operations related to disputes
 */

func Dispute(
	ctx context.Context,
	logger log.Logger,
	client contracts.ETHClient,
	contract *contracts.ITellor,
	account *config.Account,
	requestId *big.Int,
	timestamp *big.Int,
	minerIndex *big.Int,
) error {

	if !minerIndex.IsUint64() || minerIndex.Uint64() > 4 {
		return errors.Errorf("miner index should be between 0 and 4 (got %s)", minerIndex.Text(10))
	}

	balance, err := contract.BalanceOf(nil, account.Address)
	if err != nil {
		return errors.Wrap(err, "fetch balance")
	}
	var asBytes32 [32]byte
	copy(asBytes32[:], crypto.Keccak256([]byte("_DISPUTE_FEE")))
	disputeCost, err := contract.GetUintVar(nil, asBytes32)
	if err != nil {
		return errors.Wrap(err, "get dispute cost")
	}

	if balance.Cmp(disputeCost) < 0 {
		return errors.Errorf("insufficient balance TRB actual: %v, TRB required:%v)",
			util.FormatERC20Balance(balance),
			util.FormatERC20Balance(disputeCost))
	}

	auth, err := util.PrepareEthTransaction(ctx, client, account)
	if err != nil {
		return errors.Wrapf(err, "prepare ethereum transaction")
	}

	tx, err := contract.BeginDispute(auth, requestId, timestamp, minerIndex)
	if err != nil {
		return errors.Wrap(err, "send dispute txn")
	}
	level.Info(logger).Log("msg", "dispute started", "txn", tx.Hash().Hex())
	return nil
}

func Vote(
	ctx context.Context,
	logger log.Logger,
	client contracts.ETHClient,
	contract *contracts.ITellor,
	account *config.Account,
	disputeId *big.Int,
	supportsDispute bool,
) error {

	voted, err := contract.DidVote(nil, disputeId, contract.Address)
	if err != nil {
		return errors.Wrapf(err, "check if you've already voted")
	}
	if voted {
		level.Info(logger).Log("msg", "you have already voted on this dispute")
		return nil
	}

	auth, err := util.PrepareEthTransaction(ctx, client, account)
	if err != nil {
		return errors.Wrapf(err, "prepare ethereum transaction")
	}
	tx, err := contract.Vote(auth, disputeId, supportsDispute)
	if err != nil {
		return errors.Wrapf(err, "submit vote transaction")
	}

	level.Info(logger).Log("msg", "vote submitted with transaction", "tx", tx.Hash().Hex())
	return nil
}

func getNonceSubmissions(
	ctx context.Context,
	client contracts.ETHClient,
	contract *contracts.ITellor,
	valueBlock *big.Int,
	dispute *contracts.ITellorNewDispute,
) ([]*apiOracle.PriceStamp, error) {
	abi, err := abi.JSON(strings.NewReader(contracts.ITellorABI))
	if err != nil {
		return nil, errors.Wrap(err, "parse abi")
	}

	// Just use nil for most of the variables, only using this object to call UnpackLog which only uses the abi
	bar := bind.NewBoundContract(contract.Address, abi, nil, nil, nil)

	allVals, err := contract.GetSubmissionsByTimestamp(nil, dispute.RequestId, dispute.Timestamp)
	if err != nil {
		return nil, errors.Wrap(err, "get other submitted values for dispute")
	}

	allAddrs, err := contract.GetMinersByRequestIdAndTimestamp(nil, dispute.RequestId, dispute.Timestamp)
	if err != nil {
		return nil, errors.Wrap(err, "get miner addresses for dispute")
	}

	const blockStep = 100
	high := int64(valueBlock.Uint64())
	low := high - blockStep
	nonceSubmitID := abi.Events["NonceSubmitted"].ID
	timedValues := make([]*apiOracle.PriceStamp, 5)
	found := 0
	for found < 5 {
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(low),
			ToBlock:   big.NewInt(high),
			Addresses: []common.Address{contract.Address},
			Topics:    [][]common.Hash{{nonceSubmitID}},
		}

		logs, err := client.FilterLogs(ctx, query)
		if err != nil {
			return nil, errors.Wrap(err, "get nonce logs")
		}

		for _, l := range logs {
			nonceSubmit := contracts.TellorNonceSubmitted{}
			err := bar.UnpackLog(&nonceSubmit, "NonceSubmitted", l)
			if err != nil {
				return nil, errors.Wrap(err, "unpack into object")
			}
			header, err := client.HeaderByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
			if err != nil {
				return nil, errors.Wrap(err, "get nonce block header")
			}
			for i := 0; i < 5; i++ {
				if nonceSubmit.Miner == allAddrs[i] {
					valTime := time.Unix(int64(header.Time), 0)

					bigF := new(big.Float)
					bigF.SetInt(allVals[i])
					f, _ := bigF.Float64()

					timedValues[i] = &apiOracle.PriceStamp{
						Created:   valTime,
						PriceInfo: apiOracle.PriceInfo{Price: f},
					}
					found++
					break
				}
			}
		}
		high -= blockStep
		low = high - blockStep
	}
	return timedValues, nil
}

func List(
	ctx context.Context,
	cfg *config.Config,
	logger log.Logger,
	client contracts.ETHClient,
	contract *contracts.ITellor,
	account *config.Account,
) error {
	abi, err := abi.JSON(strings.NewReader(contracts.ITellorABI))
	if err != nil {
		return errors.Wrap(err, "parse abi")
	}

	// Just use nil for most of the variables, only using this object to call UnpackLog which only uses the abi.
	bar := bind.NewBoundContract(contract.Address, abi, nil, nil, nil)

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "get latest eth block header")
	}

	startBlock := big.NewInt(10e3 * 14)
	startBlock.Sub(header.Number, startBlock)
	newDisputeID := abi.Events["NewDispute"].ID
	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   nil,
		Addresses: []common.Address{contract.Address},
		Topics:    [][]common.Hash{{newDisputeID}},
	}

	logs, err := client.FilterLogs(ctx, query)
	if err != nil {
		return errors.Wrap(err, "filter eth logs")
	}

	level.Info(logger).Log("msg", "get currently open disputes", "open", len(logs))
	for _, rawDispute := range logs {
		disputeI := contracts.ITellorNewDispute{}
		err := bar.UnpackLog(&disputeI, "NewDispute", rawDispute)
		if err != nil {
			return errors.Wrap(err, "unpack dispute event from logs")
		}
		_, executed, votePassed, _, reportedAddr, reportingMiner, _, uintVars, currTally, err := contract.GetAllDisputeVars(nil, disputeI.DisputeId)
		if err != nil {
			return errors.Wrap(err, "get dispute details")
		}

		votingEnds := time.Unix(uintVars[3].Int64(), 0)
		createdTime := votingEnds.Add(-7 * 24 * time.Hour)

		var descString string
		if executed {
			descString = "complete, "
			if votePassed {
				descString += "successful"
			} else {
				descString += "rejected"
			}
		} else {
			descString = "in progress"
		}

		level.Info(logger).Log(
			"msg", "dispute occurred",
			"disputeId", disputeI.DisputeId.String(),
			"reportedAddr", reportedAddr.Hex(),
			"reportingMiner", reportingMiner.Hex(),
			"createdTime", createdTime.Format("3:04 PM January 02, 2006 MST"),
			"fee", util.FormatERC20Balance(uintVars[8]),
			"requestId", disputeI.RequestId.Uint64(),
		)

		allSubmitted, err := getNonceSubmissions(ctx, client, contract, uintVars[5], &disputeI)
		if err != nil {
			return errors.Wrapf(err, "get the values submitted by other miners for the disputed block")
		}
		disputedValTime := allSubmitted[uintVars[6].Uint64()].Created

		for i := len(allSubmitted) - 1; i >= 0; i-- {
			sub := allSubmitted[i]
			valStr := fmt.Sprintf("%f\n", sub.Price)
			var pointerStr string
			if i == int(uintVars[6].Uint64()) {
				pointerStr = " <--disputed"
			}

			level.Debug(logger).Log(
				"msg", "sub created",
				"valStr", valStr,
				"created", sub.Created.Format("3:04:05 PM"),
				"pointerStr", pointerStr,
			)
		}

		tmp := new(big.Float)
		tmp.SetInt(currTally)
		currTallyFloat, _ := tmp.Float64()
		tmp.SetInt(uintVars[7])
		currQuorum, _ := tmp.Float64()
		currTallyFloat += currQuorum
		currTallyRatio := currTallyFloat / 2 * currQuorum

		level.Info(logger).Log(
			"msg", "current TRB support for this dispute",
			"currTallyRatio", fmt.Sprintf("%0.f%%", currTallyRatio*100),
			"TRB", util.FormatERC20Balance(uintVars[7]),
			"votes", uintVars[4],
		)

		result, err := dispute.CheckValueAtTime(cfg, disputeI.RequestId.Uint64(), uintVars[2], disputedValTime)
		if err != nil {
			return err
		} else if result == nil || len(result.Datapoints) < 0 {
			level.Info(logger).Log("msg", "no data available for recommendation")
			continue
		}
		level.Info(logger).Log(
			"msg", "got recommendation",
			"vote", !result.WithinRange,
			"subValue", uintVars[2].String(),
			"range", fmt.Sprintf("%.0f to %0.f", result.Low, result.High),
		)

		numToShow := 3
		if numToShow > len(result.Datapoints) {
			numToShow = len(result.Datapoints)
		}
		level.Info(logger).Log(
			"msg", "recommedation based on",
			"datapoints", len(result.Datapoints),
			"deltaMinutes", cfg.Trackers.DisputeTimeDelta.Duration.Minutes(),
			"closest", numToShow,
		)
		minTotalDelta := time.Duration(math.MaxInt64)
		index := 0
		for i := 0; i < len(result.Datapoints)-numToShow; i++ {
			totalDelta := time.Duration(0)
			for j := 0; j < numToShow; j++ {
				delta := result.Times[i+j].Sub(disputedValTime)
				if delta < 0 {
					delta = -delta
				}
				totalDelta += delta
			}
			if totalDelta < minTotalDelta {
				minTotalDelta = totalDelta
				index = i
			}
		}
		for i := 0; i < numToShow; i++ {
			dp := result.Datapoints[index+i]
			t := result.Times[index+i]
			level.Info(logger).Log("msg", "check datapoint", "dp", dp)
			delta := disputedValTime.Sub(t)
			if delta > 0 {
				level.Info(logger).Log(
					"msg", "check delta before",
					"seconds", fmt.Sprintf("%.0f", delta.Seconds()),
				)
			} else {
				level.Info(logger).Log(
					"msg", "check delta after",
					"seconds", fmt.Sprintf("%.0f", (-delta).Seconds()),
				)
			}
		}
	}

	return nil
}
