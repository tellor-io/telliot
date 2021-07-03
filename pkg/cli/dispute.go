// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/contracts"
	tEthereum "github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/format"
	psr "github.com/tellor-io/telliot/pkg/psr/tellor"
)

func Dispute(
	ctx context.Context,
	logger log.Logger,
	client *ethclient.Client,
	contract *contracts.ITellor,
	account *tEthereum.Account,
	requestId *big.Int,
	timestamp *big.Int,
	minerIndex *big.Int,
	_gasPrice int,
) error {

	if !minerIndex.IsUint64() || minerIndex.Uint64() > 4 {
		return errors.Errorf("miner index should be between 0 and 4 (got %s)", minerIndex.Text(10))
	}

	balance, err := contract.BalanceOf(nil, account.Address)
	if err != nil {
		return errors.Wrap(err, "fetch balance")
	}

	disputeCost, err := contract.GetUintVar(nil, tEthereum.Keccak256([]byte("_DISPUTE_FEE")))
	if err != nil {
		return errors.Wrap(err, "get dispute cost")
	}

	if balance.Cmp(disputeCost) < 0 {
		return errors.Errorf("insufficient balance TRB actual: %v, TRB required:%v)",
			format.ERC20Balance(balance),
			format.ERC20Balance(disputeCost))
	}

	var gasPrice *big.Int
	if _gasPrice > 0 {
		gasPrice = big.NewInt(int64(_gasPrice) * params.GWei)
	}

	auth, err := tEthereum.PrepareEthTransaction(ctx, client, account, gasPrice)
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

func List(
	ctx context.Context,
	logger log.Logger,
	client *ethclient.Client,
	contract *contracts.ITellor,
	account *tEthereum.Account,
	psr *psr.Psr,
) error {
	// TODO fix it!

	// abi, err := abi.JSON(strings.NewReader(contracts.ITellorABI))
	// if err != nil {
	// 	return errors.Wrap(err, "parse abi")
	// }

	// // Just use nil for most of the variables, only using this object to call UnpackLog which only uses the abi.
	// bar := bind.NewBoundContract(contract.Address, abi, nil, nil, nil)

	// header, err := client.HeaderByNumber(ctx, nil)
	// if err != nil {
	// 	return errors.Wrap(err, "get latest eth block header")
	// }

	// startBlock := big.NewInt(10e3 * 14) // TODO instead of this calculate only 10 days worth of blocks in the past since disputes can be voted only 7 days in the past.
	// startBlock.Sub(header.Number, startBlock)
	// newDisputeID := abi.Events["NewDispute"].ID
	// query := ethereum.FilterQuery{
	// 	FromBlock: startBlock,
	// 	ToBlock:   nil,
	// 	Addresses: []common.Address{contract.Address},
	// 	Topics:    [][]common.Hash{{newDisputeID}},
	// }

	// logs, err := client.FilterLogs(ctx, query)
	// if err != nil {
	// 	return errors.Wrap(err, "filter eth logs")
	// }

	// level.Info(logger).Log("msg", "get currently open disputes", "open", len(logs))
	// for _, rawDispute := range logs {
	// 	disputeI := contracts.ITellorNewDispute{}
	// 	err := bar.UnpackLog(&disputeI, "NewDispute", rawDispute)
	// 	if err != nil {
	// 		return errors.Wrap(err, "unpack dispute event from logs")
	// 	}
	// 	_, executed, votePassed, _, reportedAddr, reportingMiner, _, uintVars, currTally, err := contract.GetAllDisputeVars(nil, disputeI.DisputeId)
	// 	if err != nil {
	// 		return errors.Wrap(err, "get dispute details")
	// 	}

	// 	votingEnds := time.Unix(uintVars[3].Int64(), 0)
	// 	createdTime := votingEnds.Add(-7 * 24 * time.Hour)

	// 	var descString string
	// 	if executed {
	// 		descString = "complete, "
	// 		if votePassed {
	// 			descString += "successful"
	// 		} else {
	// 			descString += "rejected"
	// 		}
	// 	} else {
	// 		descString = "in progress"
	// 	}

	// 	level.Info(logger).Log(
	// 		"msg", "dispute occurred",
	// 		"disputeId", disputeI.DisputeId.String(),
	// 		"reportedAddr", reportedAddr.Hex(),
	// 		"reportingMiner", reportingMiner.Hex(),
	// 		"createdTime", createdTime.Format("3:04 PM January 02, 2006 MST"),
	// 		"fee", format.ERC20Balance(uintVars[8]),
	// 		"requestId", disputeI.RequestId.Uint64(),
	// 	)

	// 	allSubmitted, err := getNonceSubmits(ctx, client, contract, uintVars[5], &disputeI)
	// 	if err != nil {
	// 		return errors.Wrapf(err, "get the values submitted by other miners for the disputed block")
	// 	}
	// 	disputedValTime := allSubmitted[uintVars[6].Uint64()].Time

	// 	for i := len(allSubmitted) - 1; i >= 0; i-- {
	// 		sub := allSubmitted[i]
	// 		valStr := fmt.Sprintf("%f\n", sub.float64)
	// 		var pointerStr string
	// 		if i == int(uintVars[6].Uint64()) {
	// 			pointerStr = " <--disputed"
	// 		}

	// 		level.Debug(logger).Log(
	// 			"msg", "sub created",
	// 			"valStr", valStr,
	// 			"created", sub.Time.Format("3:04:05 PM"),
	// 			"pointerStr", pointerStr,
	// 		)
	// 	}

	// 	tmp := new(big.Float)
	// 	tmp.SetInt(currTally)
	// 	currTallyFloat, _ := tmp.Float64()
	// 	tmp.SetInt(uintVars[7])
	// 	currQuorum, _ := tmp.Float64()
	// 	currTallyFloat += currQuorum
	// 	currTallyRatio := currTallyFloat / 2 * currQuorum

	// 	level.Info(logger).Log(
	// 		"msg", "current TRB support for this dispute",
	// 		"currTallyRatio", fmt.Sprintf("%0.f%%", currTallyRatio*100),
	// 		"TRB", format.ERC20Balance(uintVars[7]),
	// 		"votes", uintVars[4],
	// 	)

	// 	header, err := client.HeaderByNumber(ctx, nil)
	// 	if err != nil {
	// 		return errors.Wrap(err, "get latest eth block header")
	// 	}

	// 	disputer := dispute.NewDisputeChecker(
	// 		logger,
	// 		cfg,
	// 		client,
	// 		contract,
	// 		header.Number.Uint64(),
	// 		psr,
	// 	)

	// 	result, err := disputer.CheckValueAtTime(disputeI.RequestId.Int64(), uintVars[2], disputedValTime)
	// 	if err != nil {
	// 		return err
	// 	} else if result == nil || len(result.Datapoints) < 0 {
	// 		level.Info(logger).Log("msg", "no data available for recommendation")
	// 		continue
	// 	}
	// 	level.Info(logger).Log(
	// 		"msg", "got recommendation",
	// 		"vote", !result.WithinRange,
	// 		"subValue", uintVars[2].String(),
	// 		"range", strconv.Itoa(int(result.Low))+"to"+strconv.Itoa(int(result.High)),
	// 	)

	// 	numToShow := 3
	// 	if numToShow > len(result.Datapoints) {
	// 		numToShow = len(result.Datapoints)
	// 	}
	// 	level.Info(logger).Log(
	// 		"msg", "recommedation based on",
	// 		"datapoints", len(result.Datapoints),
	// 		"deltaMinutes", cfg.DisputeTimeDelta.Duration.Minutes(),
	// 		"closest", numToShow,
	// 	)
	// 	minTotalDelta := time.Duration(math.MaxInt64)
	// 	index := 0
	// 	for i := 0; i < len(result.Datapoints)-numToShow; i++ {
	// 		totalDelta := time.Duration(0)
	// 		for j := 0; j < numToShow; j++ {
	// 			delta := result.Times[i+j].Sub(disputedValTime)
	// 			if delta < 0 {
	// 				delta = -delta
	// 			}
	// 			totalDelta += delta
	// 		}
	// 		if totalDelta < minTotalDelta {
	// 			minTotalDelta = totalDelta
	// 			index = i
	// 		}
	// 	}
	// 	for i := 0; i < numToShow; i++ {
	// 		dp := result.Datapoints[index+i]
	// 		t := result.Times[index+i]
	// 		level.Info(logger).Log("msg", "check datapoint", "dp", dp)
	// 		delta := disputedValTime.Sub(t)
	// 		if delta > 0 {
	// 			level.Info(logger).Log(
	// 				"msg", "check delta before",
	// 				"seconds", fmt.Sprintf("%.0f", delta.Seconds()),
	// 			)
	// 		} else {
	// 			level.Info(logger).Log(
	// 				"msg", "check delta after",
	// 				"seconds", fmt.Sprintf("%.0f", (-delta).Seconds()),
	// 			)
	// 		}
	// 	}
	// }

	return nil
}
