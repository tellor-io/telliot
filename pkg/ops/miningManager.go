// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	tellorCommon "github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/contracts/proxy"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/pow"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/tracker"
)

const maxMiningTimeout = 15 * time.Minute

type WorkSource interface {
	GetWork(toMine chan *pow.Work) *pow.Work
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
	logger          log.Logger
	Running         bool
	ethClient       rpc.ETHClient
	group           *pow.MiningGroup
	tasker          WorkSource
	solHandler      SolutionSink
	solutionPending *pow.Result
	database        db.DataServerProxy
	contractGetter  *proxy.TellorGetters
	cfg             *config.Config

	toMineInput     chan *pow.Work
	solutionOutput  chan *pow.Result
	submitCount     prometheus.Counter
	submitFailCount prometheus.Counter
	submitProfit    *prometheus.GaugeVec

	miningCtxCnl context.CancelFunc
}

// CreateMiningManager is the MiningMgr constructor.
func CreateMiningManager(
	logger log.Logger,
	exitCh chan os.Signal,
	cfg *config.Config,
	database db.DataServerProxy,
	contract contracts.Tellor,
	account rpc.Account,
) (*MiningMgr, error) {

	group, err := pow.SetupMiningGroup(cfg, exitCh)
	if err != nil {
		return nil, errors.Wrap(err, "setup miners")
	}

	client, err := rpc.NewClient(os.Getenv(config.NodeURLEnvName))
	if err != nil {
		return nil, errors.Wrap(err, "creating client")
	}
	contractAddress := common.HexToAddress(cfg.ContractAddress)
	getter, err := proxy.NewTellorGetters(contractAddress, client)
	if err != nil {
		return nil, errors.Wrap(err, "getting addresses")
	}

	submitter := NewSubmitter(logger, cfg, client, contract, account)
	mng := &MiningMgr{
		exitCh:          exitCh,
		logger:          logger,
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
		submitCount: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "telliot",
			Subsystem: "mining",
			Name:      "submit_total",
			Help:      "The total number of submitted solutions",
		}),
		submitFailCount: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "telliot",
			Subsystem: "mining",
			Name:      "submit_fails_total",
			Help:      "The total number of failed submission",
		}),
		submitProfit: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: "mining",
			Name:      "submit_profit",
			Help:      "The submission profit in percents",
		},
			[]string{"slot"},
		),
	}

	if cfg.EnablePoolWorker {
		pool := pow.CreatePool(cfg, group)
		mng.tasker = pool
		mng.solHandler = pool
	} else {
		mng.tasker = pow.CreateTasker(cfg, database)
		mng.solHandler = pow.CreateSolutionHandler(cfg, submitter, database)
	}
	return mng, nil
}

// Start will start the mining run loop.
func (mgr *MiningMgr) Start(ctx context.Context) {
	mgr.Running = true
	ticker := time.NewTicker(mgr.cfg.MiningInterruptCheckInterval.Duration)

	// Start the mining group.
	go mgr.group.Mine(mgr.toMineInput, mgr.solutionOutput)

	for {
		select {
		// Boss wants us to quit for the day.
		case <-mgr.exitCh:
			mgr.Running = false
			return
		// Time to check for a new challenge and
		// while waiting resubmit any unsubmitted solutions.
		case <-ticker.C:
			mgr.newWork()
			if mgr.solutionPending != nil {
				if mgr.submit(mgr.solutionPending, ctx) {
					mgr.solutionPending = nil
				}
			}
		// New solution from the miner.
		case solution := <-mgr.solutionOutput:
			mgr.miningCtxCnl() // This does nothing, but run it to avoid context leak.
			if !mgr.submit(solution, ctx) {
				mgr.solutionPending = solution
			}

		}
	}
}

// submit returns true when the provided solution has been submited without an error or
// with an error which can be ignored.
func (mgr *MiningMgr) submit(solution *pow.Result, ctx context.Context) bool {
	lastSubmit, err := mgr.lastMinerSubmit()
	if err != nil {
		level.Error(mgr.logger).Log("msg", "checking last submit time", "err", err)
	} else if lastSubmit < mgr.cfg.MinSubmitPeriod {
		level.Debug(mgr.logger).Log("msg", "min transaction submit threshold hasn't passed", "minSubmitPeriod", mgr.cfg.MinSubmitPeriod, "lastSubmit", lastSubmit)
		return false
	}

	var ids []int64
	for _, id := range solution.Work.Challenge.RequestIDs {
		ids = append(ids, id.Int64())
	}
	level.Debug(mgr.logger).Log("msg", "submitting solution", "reqIDs", fmt.Sprintf("%+v", ids))

	var profitPercent int64
	var slot int64
	if mgr.cfg.ProfitThreshold > 0 {
		profitPercent, slot, err = mgr.profit()
		if err != nil {
			level.Error(mgr.logger).Log("msg", "submit solution profit check", "err", err)
			return false
		}
		if profitPercent != -1 && profitPercent < int64(mgr.cfg.ProfitThreshold) {
			level.Debug(mgr.logger).Log("msg", "transaction not profitable, so will wait for the next cycle")
			return false
		}
	}

	tx, err := mgr.solHandler.Submit(ctx, solution)
	if err != nil {
		level.Error(mgr.logger).Log("msg", "submiting a solution", "err", err)
		mgr.submitFailCount.Inc()
		return false
	}
	level.Debug(mgr.logger).Log("msg", "submited a solution", "txHash", tx.Hash().String())
	mgr.saveGasUsed(ctx, tx)
	mgr.submitCount.Inc()
	mgr.submitProfit.With(prometheus.Labels{"slot": strconv.Itoa(int(slot))}).(prometheus.Gauge).Set(float64(profitPercent))

	return true
}

// newWork is non blocking worker that sends new work to the pow workers
// or re-sends a current pending solution to the submitter when the challenge hasn't changes.
func (mgr *MiningMgr) newWork() {
	if mgr.cfg.EnablePoolWorker {
		mgr.tasker.GetWork(mgr.toMineInput)
	} else {
		work := mgr.tasker.GetWork(nil)
		if work == nil {
			return
		}
		var ids []int64
		for _, id := range work.Challenge.RequestIDs {
			ids = append(ids, id.Int64())
		}
		// For oracle blocks that are more than 15mins old can submit any solution.
		// No need to send it for minging.
		if mgr.isInstantSubmit() {
			mgr.solutionOutput <- &pow.Result{Work: work, Nonce: "anything will work"}
		}
		level.Debug(mgr.logger).Log("msg", "sending new chalenge for mining", "reqIDs", fmt.Sprintf("%+v", ids))

		// Send it for mining with a timeout to expire when the difficulty is zero and
		// can submit any solution. The timeout will cancel the the mining and
		// the mining loop will return a random nonce as after the maxMiningTimeout any nonce will work.
		tm, err := mgr.timeOfLastNewValue()
		ctx, cancel := context.WithTimeout(context.Background(), maxMiningTimeout)
		if err != nil {
			level.Error(mgr.logger).Log("msg", "get TimeOfLastNewValueKey", "err", err)
		} else {
			now := time.Now()
			ctx, cancel = context.WithTimeout(context.Background(), now.Sub(tm))
		}
		mgr.miningCtxCnl = cancel // Not used, but needs to be called later to avoid leaks.
		work.TimeoutCtx = ctx
		mgr.toMineInput <- work

	}
}

// isInstantSubmit checks if the current oracle block is more than maxMiningTimeout old.
func (mgr *MiningMgr) isInstantSubmit() bool {
	tm, err := mgr.timeOfLastNewValue()
	if err != nil {
		level.Error(mgr.logger).Log("msg", "get TimeOfLastNewValueKey", "err", err)
		return false
	}
	now := time.Now()
	level.Debug(mgr.logger).Log("msg", "since last value", "time", now.Sub(tm))

	return now.Sub(tm) >= maxMiningTimeout
}

func (mgr *MiningMgr) timeOfLastNewValue() (time.Time, error) {
	timeOfLastNewValue, err := mgr.database.Get(db.TimeOfLastNewValueKey)
	if err != nil {
		return time.Time{}, errors.Wrap(err, "get TimeOfLastNewValueKey from db")
	}
	if len(timeOfLastNewValue) == 0 {
		return time.Time{}, errors.New("TimeOfLastNewValueKey is empty in the db")
	}

	t, err := hexutil.DecodeBig(string(timeOfLastNewValue))
	if err != nil {
		return time.Time{}, errors.Wrap(err, "decoding TimeOfLastNewValueKey from db")
	}
	return time.Unix(t.Int64(), 0), nil
}

func (mgr *MiningMgr) lastMinerSubmit() (time.Duration, error) {
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

func (mgr *MiningMgr) gasUsed() (*big.Int, *big.Int, error) {
	slotNum, err := mgr.contractGetter.GetUintVar(nil, rpc.Keccak256([]byte("slotProgress")))
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting slotProgress")
	}
	// This is the price for the last transaction so increment +1
	// to get the price for next slot transaction.
	// Slots numbers should be from 0 to 4 so
	// use mod of 5 in order to save 5 as slot 0.
	slotNum.Add(slotNum, big.NewInt(1)).Mod(slotNum, big.NewInt(5))
	txID := tellorCommon.PriceTXs + slotNum.String()
	gas, err := mgr.database.Get(txID)
	if err != nil {
		return nil, nil, errors.New("getting the tx eth cost from the db")
	}
	// No price record in the db yet.
	if gas == nil {
		return big.NewInt(0), slotNum, nil
	}

	return big.NewInt(0).SetBytes(gas), slotNum, nil
}

// saveGasUsed calculates the price for a given slot.
// Since the transaction doesn't include the slot number it gets the slot number
// as soon as the transaction passes and
// saves it in the database for profit calculations.
// TODO[Krasi] To be more detirministic and simplify this
// should get the `slotProgress` and `gasUsed` from the `NonceSubmitted` event.
// At the moment there is a slight chance of a race condition if
// another transaction has passed between checking the transaction cost and
// checking the `slotProgress`
// Tracking issue https://github.com/tellor-io/TellorCore/issues/101
func (mgr *MiningMgr) saveGasUsed(ctx context.Context, tx *types.Transaction) {
	go func(tx *types.Transaction) {
		receipt, err := bind.WaitMined(ctx, mgr.ethClient, tx)
		if err != nil {
			level.Error(mgr.logger).Log("msg", "transaction result for calculating transaction cost", "err", err)
		}
		if receipt.Status != 1 {
			mgr.submitFailCount.Inc()
			level.Error(mgr.logger).Log("msg", "unsuccessful submitSolution transaction, not saving the tx cost in the db", "txHash", receipt.TxHash.String())
			return
		}

		gasUsed := big.NewInt(int64(receipt.GasUsed))
		slotNum, err := mgr.contractGetter.GetUintVar(nil, rpc.Keccak256([]byte("slotProgress")))
		if err != nil {
			level.Error(mgr.logger).Log("msg", "getting slotProgress for calculating transaction cost", "err", err)
		}

		txID := tellorCommon.PriceTXs + slotNum.String()
		_, err = mgr.database.Put(txID, gasUsed.Bytes())
		if err != nil {
			level.Error(mgr.logger).Log("msg", "saving transaction cost", "err", err)
		}
		level.Debug(mgr.logger).Log("msg", "saved transaction cost", "txHash", receipt.TxHash.String(), "gas", gasUsed.Int64(), "slot", slotNum.Int64())
	}(tx)
}

// profit returns the profit in percents.
// When the transaction cost is unknown it returns -1 so
// that the caller can decide how to handle.
// Transaction cost is zero when the manager hasn't done any transactions yet.
// Each transaction cost is known for any siquential transactions.
func (mgr *MiningMgr) profit() (int64, int64, error) {
	gasUsed, slotNum, err := mgr.gasUsed()
	if err != nil {
		return 0, slotNum.Int64(), errors.Wrap(err, "getting TX cost")
	}
	if gasUsed.Int64() == 0 {
		level.Debug(mgr.logger).Log("msg", "profit checking no data for gas used", "slot", slotNum)
		return -1, slotNum.Int64(), nil
	}
	gasPrice, err := mgr.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return 0, slotNum.Int64(), errors.Wrap(err, "getting gas price")
	}
	reward, err := mgr.currentReward()
	if err != nil {
		return 0, slotNum.Int64(), errors.Wrap(err, "getting current rewards")
	}

	txCost := gasPrice.Mul(gasPrice, gasUsed)
	profit := big.NewInt(0).Sub(reward, txCost)
	profitPercent := big.NewInt(0).Div(profit, txCost).Int64() * 100
	level.Debug(mgr.logger).Log(
		"msg", "profit checking",
		"reward", fmt.Sprintf("%.2e", float64(reward.Int64())),
		"txCost", fmt.Sprintf("%.2e", float64(txCost.Int64())),
		"slot", slotNum,
		"profit", fmt.Sprintf("%.2e", float64(profit.Int64())),
		"profitMargin", profitPercent,
		"profitThreshold", mgr.cfg.ProfitThreshold,
	)

	return profitPercent, slotNum.Int64(), nil
}
