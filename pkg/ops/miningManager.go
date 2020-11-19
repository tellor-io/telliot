// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/contracts/getter"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/pow"
	"github.com/tellor-io/TellorMiner/pkg/rpc"
	"github.com/tellor-io/TellorMiner/pkg/tracker"
	"github.com/tellor-io/TellorMiner/pkg/util"
)

var minSubmitPeriod = 15 * time.Minute

type WorkSource interface {
	GetWork(toMine chan *pow.Work) (*pow.Work, bool)
}

type SolutionSink interface {
	Submit(context.Context, *pow.Result) (*types.Transaction, error)
}

// MiningMgr manages mining, submiting a solution and requesting data.
// In the tellor contract a solution is saved in slots where a value is valid only when it has 5 confirmed slots.
// The manager tracks tx costs and profitThreshold is set it skips any transactions below the profit threshold.
// The profit is calculated the same way as in the Tellor contract.
// Transaction cost for submitting in each slot might be different so because of this
// the manager needs to complete few transaction to gather the tx cost for each slot.
type MiningMgr struct {
	exitCh          chan os.Signal
	log             *util.Logger
	Running         bool
	ethClient       rpc.ETHClient
	group           *pow.MiningGroup
	tasker          WorkSource
	solHandler      SolutionSink
	solutionPending *pow.Result
	dataRequester   *DataRequester
	database        db.DataServerProxy
	contractGetter  *getter.TellorGetters
	cfg             *config.Config

	toMineInput    chan *pow.Work
	solutionOutput chan *pow.Result
}

// CreateMiningManager is the MiningMgr contructor.
func CreateMiningManager(
	exitCh chan os.Signal,
	cfg *config.Config,
	database db.DataServerProxy,
) (*MiningMgr, error) {

	group, err := pow.SetupMiningGroup(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "to setup miners")
	}

	client, err := rpc.NewClient(cfg.NodeURL)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(cfg.ContractAddress)
	getter, err := getter.NewTellorGetters(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	submitter := NewSubmitter()
	mng := &MiningMgr{
		exitCh:          exitCh,
		log:             util.NewLogger("ops", "MiningMgr"),
		Running:         false,
		group:           group,
		tasker:          nil,
		solutionPending: nil,
		solHandler:      nil,
		contractGetter:  getter,
		cfg:             cfg,
		database:        database,
		ethClient:       client,
		toMineInput:     make(chan *pow.Work),
		solutionOutput:  make(chan *pow.Result),
	}

	if cfg.EnablePoolWorker {
		pool := pow.CreatePool(cfg, group)
		mng.tasker = pool
		mng.solHandler = pool
	} else {
		mng.tasker = pow.CreateTasker(cfg, database)
		mng.solHandler = pow.CreateSolutionHandler(cfg, submitter, database)
		if cfg.RequestData > 0 {
			mng.log.Info("dataRequester created")
			mng.dataRequester = CreateDataRequester(exitCh, submitter, cfg.RequestDataInterval.Duration, database)
		}
	}
	return mng, nil
}

// Start will start the mining run loop.
func (mgr *MiningMgr) Start(ctx context.Context) {
	mgr.Running = true
	ticker := time.NewTicker(mgr.cfg.MiningInterruptCheckInterval.Duration)

	if mgr.cfg.RequestData > 0 {
		if err := mgr.dataRequester.Start(ctx); err != nil {
			mgr.log.Error("error starting the data requester error:%v", err)
		}
	}

	// Start the mining group.
	go mgr.group.Mine(mgr.toMineInput, mgr.solutionOutput)

	for {
		select {
		// Boss wants us to quit for the day.
		case <-mgr.exitCh:
			mgr.Running = false
			return
		// Found a solution.
		case solution := <-mgr.solutionOutput:
			// There is no new challenge so resend any pending solution.
			if solution == nil {
				if mgr.solutionPending == nil {
					continue
				}
				solution = mgr.solutionPending
				var ids []int64
				for _, id := range mgr.solutionPending.Work.Challenge.RequestIDs {
					ids = append(ids, id.Int64())
				}
				mgr.log.Debug("re-submitting a pending solution - req IDs:%v", ids)
			}

			// Set this solution as pending so that if
			// any of the checks below fail and will be retried
			// when there is no new challenge.
			mgr.solutionPending = solution

			if mgr.cfg.ProfitThreshold > 0 {
				isProftable, err := mgr.isProfitable()
				if err != nil {
					mgr.log.Error("solution profit err:%v", err)
				} else if !isProftable {
					mgr.log.Debug("transaction not profitable, so will wait for the next cycle")
					continue
				}
			}

			lastSubmit, err := mgr.lastSubmit()
			if err != nil {
				mgr.log.Error("checking last submit time err:%v", err)
			} else if lastSubmit < minSubmitPeriod {
				mgr.log.Debug("min transaction submit threshold of %v hasn't passed, last submit:%v", minSubmitPeriod, lastSubmit)
				continue
			}
			tx, err := mgr.solHandler.Submit(ctx, solution)
			if err != nil {
				mgr.log.Error("submiting a solution err:%v", err)
				continue
			}
			mgr.log.Debug("submited a solution tx:%v", tx.Hash().String())
			mgr.saveTXCost(ctx, tx)

			// A solution has been submitted so the
			// pending solution doesn't matter here any more so reset it.
			mgr.solutionPending = nil

		// Time to check for a new challenge.
		case <-ticker.C:
			mgr.newWork()
		}
	}
}

// newWork is non blocking worker that sends new work to the pow workers
// or re-sends a current pending solution to the submitter when the challenge hasn't changes.
func (mgr *MiningMgr) newWork() {
	go func() {
		if mgr.cfg.EnablePoolWorker {
			mgr.tasker.GetWork(mgr.toMineInput)
		} else {
			// instantSubmit means 15 mins have passed so
			// the difficulty now is zero and any solution/nonce will work so
			// can just submit without sending to the miner.
			work, instantSubmit := mgr.tasker.GetWork(nil)
			if instantSubmit {
				mgr.solutionOutput <- &pow.Result{Work: work, Nonce: "anything will work"}
			} else {
				// It sends even nil work to indicate that no new challenge is available.
				if work == nil {
					mgr.solutionOutput <- nil
					return
				}

				var ids []int64
				for _, id := range work.Challenge.RequestIDs {
					ids = append(ids, id.Int64())
				}
				mgr.log.Debug("sending new chalenge for mining - req IDs:%v", ids)
				mgr.toMineInput <- work
			}
		}
	}()
}

func (mgr *MiningMgr) lastSubmit() (time.Duration, error) {
	fromAddress := common.HexToAddress(mgr.cfg.PublicAddress)
	pubKey := strings.ToLower(fromAddress.Hex())

	address := common.HexToAddress(pubKey)
	dbKey := fmt.Sprintf("%s-%s", strings.ToLower(address.Hex()), db.TimeOutKey)
	last, err := mgr.database.Get(dbKey)
	if err != nil {
		return time.Duration(0), errors.Wrapf(err, "timeout retrieval error")
	}
	lastDecoded, err := hexutil.DecodeBig(string(last))
	if err != nil {
		return time.Duration(0), errors.Wrapf(err, "timeout key decode last:%v", last)
	}
	lastInt := lastDecoded.Int64()
	now := time.Now()
	var lastSubmit time.Duration
	if lastInt > 0 {
		tm := time.Unix(lastInt, 0)
		lastSubmit = now.Sub(tm)
	}

	return lastSubmit, nil
}

// currentReward returns the current TRB rewards converted to ETH.
// TODO[Krasi] This is a duplicate code from the tellor conract so
// Should add `currentReward` func to the contract to avoid this code duplication.
// Tracking issue https://github.com/tellor-io/TellorCore/issues/101
func (mgr *MiningMgr) currentReward() (*big.Int, error) {
	timeOfLastNewValue, err := mgr.contractGetter.GetUintVar(nil, rpc.Keccak256([]byte("timeOfLastNewValue")))
	if err != nil {
		return nil, errors.New("getting timeOfLastNewValue")
	}
	currentTotalTips, err := mgr.contractGetter.GetUintVar(nil, rpc.Keccak256([]byte("currentTotalTips")))
	if err != nil {
		return nil, errors.New("getting currentTotalTips")
	}

	timeDiff := big.NewInt(time.Now().Unix() - timeOfLastNewValue.Int64())
	trb := big.NewInt(1e18)
	rewardPerSec := big.NewInt(0).Div(trb, big.NewInt(300)) // 1 TRB every 5 minutes so total reward is timeDiff multiplied by reward per second.
	rewardTRB := big.NewInt(0).Mul(rewardPerSec, timeDiff)

	singleMinerTip := big.NewInt(0).Div(currentTotalTips, big.NewInt(10)) // Half of the tips are burned(remain in the contract) to reduce inflation.
	rewardWithTips := big.NewInt(0).Add(singleMinerTip, rewardTRB)

	return mgr.convertTRBtoETH(rewardWithTips)
}

func (mgr *MiningMgr) convertTRBtoETH(trb *big.Int) (*big.Int, error) {
	val, err := mgr.database.Get(db.QueriedValuePrefix + strconv.Itoa(tracker.RequestID_TRB_ETH))
	if err != nil {
		return nil, errors.New("getting the trb price from the db")
	}
	if len(val) == 0 {
		return nil, errors.New("the db doesn't have the trb price")
	}
	priceTRB, err := hexutil.DecodeBig(string(val))
	if err != nil {
		return nil, errors.New("decoding trb price from the db")
	}

	priceTRB = priceTRB.Mul(priceTRB, big.NewInt(tellorCommon.WEI))
	priceTRB = priceTRB.Div(priceTRB, big.NewInt(tracker.PSRs[tracker.RequestID_TRB_ETH].Granularity()))

	// Big int can't multiple fractions so need to convert the operation into a division.
	devider := big.NewInt(0).Div(trb, trb)
	eth := big.NewInt(0).Div(priceTRB, devider)
	return eth, nil
}

func (mgr *MiningMgr) txCost() (*big.Int, *big.Int, error) {
	slotNum, err := mgr.contractGetter.GetUintVar(nil, rpc.Keccak256([]byte("slotProgress")))
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting slotProgress")
	}
	// This is the price for the last transaction so increment +1
	// to get the price for next slot transaction.
	// Slots numbers should be from 0 to 4 so
	// use mod of 5 in order to save 5 as slot 0.
	slotNum.Add(slotNum, big.NewInt(1)).Mod(slotNum, big.NewInt(5))
	txCostID := tellorCommon.PriceTXs + slotNum.String()
	cost, err := mgr.database.Get(txCostID)
	if err != nil {
		return nil, nil, errors.New("getting the tx eth cost from the db")
	}
	// No price record in the db yet.
	if cost == nil {
		return big.NewInt(0), slotNum, nil
	}

	return big.NewInt(0).SetBytes(cost), slotNum, nil
}

// saveTXCost calculates the price for a given slot.
// Since the transaction doesn't include the slot number it was submited for
// it gets the slot number as soon as the transaction passes and
// saves it in the database for profit calculations use.
// TODO[Krasi] To be more detirministic and simplify this
// should get the `slotProgress` and `gasUsed` from the `NonceSubmitted` event.
// At the moment there is a slight chance of a race condition if
// another transaction has passed between checking the transaction cost and
// checking the `slotProgress`
// Tracking issue https://github.com/tellor-io/TellorCore/issues/101
func (mgr *MiningMgr) saveTXCost(ctx context.Context, tx *types.Transaction) {
	go func(tx *types.Transaction) {
		receipt, err := bind.WaitMined(ctx, mgr.ethClient, tx)
		if err != nil {
			mgr.log.Error("transaction result for calculating transaction cost err:%v", err)
		}
		if receipt.Status != 1 {
			mgr.log.Error("unsuccessful submitSolution transaction, not saving the tx cost in the db tx:%v", receipt.TxHash.String())
			return
		}

		gasUsed := big.NewInt(int64(receipt.GasUsed))
		txCost := gasUsed.Mul(gasUsed, tx.GasPrice())
		slotNum, err := mgr.contractGetter.GetUintVar(nil, rpc.Keccak256([]byte("slotProgress")))
		if err != nil {
			mgr.log.Error("getting slotProgress for calculating transaction cost err:%v", err)
		}

		txCostID := tellorCommon.PriceTXs + slotNum.String()
		_, err = mgr.database.Put(txCostID, txCost.Bytes())
		if err != nil {
			mgr.log.Error("saving transaction cost err:%v", err)
		}
		mgr.log.Debug("saved transaction cost txHash:%v cost GWEI:%v slot:%v", receipt.TxHash.String(), txCost.Int64()/tellorCommon.GWEI, slotNum.Int64())

	}(tx)
}

func (mgr *MiningMgr) isProfitable() (bool, error) {
	txCost, slotNum, err := mgr.txCost()
	if err != nil {
		return false, errors.Wrap(err, "getting TX cost")
	}
	// Transction cost is zero when the manager hasn't done any transactions yet.
	// Each transaction cost recorder to it is known for any siquential transactions.
	// When transaction cost is unknown it shouldn't block the submission hence returning true here.
	if txCost.Int64() == 0 {
		return true, nil
	}
	reward, err := mgr.currentReward()
	if err != nil {
		return false, errors.Wrap(err, "getting current rewards")
	}

	profit := big.NewInt(0).Sub(reward, txCost)
	profitPercent := big.NewInt(0).Div(profit, txCost).Int64() * 100
	mgr.log.Debug(
		"profit checking - reward:%v, txCost:%v, slot:%v, profit:%v, profit margin:%v%%, profit threshold:%v%%",
		fmt.Sprintf("%.2e", float64(reward.Int64())),
		fmt.Sprintf("%.2e", float64(txCost.Int64())),
		slotNum,
		fmt.Sprintf("%.2e", float64(profit.Int64())),
		profitPercent,
		mgr.cfg.ProfitThreshold,
	)
	if profitPercent > int64(mgr.cfg.ProfitThreshold) {
		return true, nil
	}
	return false, nil
}
