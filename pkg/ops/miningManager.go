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

type WorkSource interface {
	GetWork(input chan *pow.Work) (*pow.Work, bool)
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
	solution        *pow.Result
	dataRequester   *DataRequester
	database        db.DataServerProxy
	contractGetter  *getter.TellorGetters
	profitThreshold int64
}

// CreateMiningManager is the MiningMgr contructor.
func CreateMiningManager(
	exitCh chan os.Signal,
	cfg *config.Config,
	database db.DataServerProxy,
) (*MiningMgr, error) {

	group, err := pow.SetupMiningGroup(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to setup miners: %s", err.Error())
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
		solution:        nil,
		solHandler:      nil,
		contractGetter:  getter,
		profitThreshold: cfg.ProfitThreshold,
		database:        database,
		ethClient:       client,
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
	go func(ctx context.Context) {
		cfg := config.GetConfig()

		ticker := time.NewTicker(cfg.MiningInterruptCheckInterval.Duration)

		//if you make these buffered, think about the effects on synchronization!
		input := make(chan *pow.Work)
		output := make(chan *pow.Result)
		if cfg.RequestData > 0 {
			if err := mgr.dataRequester.Start(ctx); err != nil {
				mgr.log.Error("error starting the data requester error:%v", err)
			}
		}

		//start the mining group
		go mgr.group.Mine(input, output)

		// sends work to the mining group
		sendWork := func() {
			if cfg.EnablePoolWorker {
				mgr.tasker.GetWork(input)
			} else {
				work, instantSubmit := mgr.tasker.GetWork(input)
				if instantSubmit {
					if mgr.solution == nil {
						if mgr.profitThreshold > 0 {
							isProftable, err := mgr.isProfitable()
							if err != nil {
								mgr.log.Warn("error calculating the solution submition profitability so skiping")
								return
							}
							if !isProftable {
								mgr.log.Warn("solution submition not profitable so skiping")
								return
							}
						}
						mgr.log.Debug("instant submit called!")
						mgr.solution = &pow.Result{Work: work, Nonce: "1"}
						tx, err := mgr.solHandler.Submit(ctx, mgr.solution)
						if err != nil {
							mgr.log.Error("submiting a solution transaction err:%v", err)
							return
						}
						mgr.log.Debug("submited a solution tx:%v", tx)
						mgr.saveTXCost(ctx, tx)
					}
				} else if work != nil {
					mgr.solution = nil
					input <- work
				}
			}
		}
		// Send the initial challenge.
		sendWork()
		for {
			select {
			// Boss wants us to quit for the day.
			case <-mgr.exitCh:
				input <- nil

			// Found a solution.
			case result := <-output:
				if mgr.profitThreshold > 0 {
					isProftable, err := mgr.isProfitable()
					if err != nil {
						mgr.log.Warn("error calculating the solution submition profitability so skiping")
						return
					}
					if !isProftable {
						mgr.log.Warn("solution submition not profitable so skiping")
						return
					}
				}
				if result == nil {
					mgr.Running = false
					return
				}
				mgr.solution = result
				tx, err := mgr.solHandler.Submit(ctx, mgr.solution)
				if err != nil {
					mgr.log.Error("submiting a solution transaction err:%v", err)
					continue
				}
				mgr.log.Debug("submited a solution:%+v tx hash:%+v", &mgr.solution, &tx)
				mgr.saveTXCost(ctx, tx)
				sendWork()

			// Time to check for a new challenge.
			case <-ticker.C:
				sendWork()
			}
		}
	}(ctx)
}

// currentReward returns the current TRB rewards converted to ETH.
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

func (mgr *MiningMgr) txCost() (*big.Int, error) {
	slotNum, err := mgr.contractGetter.GetUintVar(nil, rpc.Keccak256([]byte("slotProgress")))
	if err != nil {
		return nil, errors.Wrap(err, "getting slotProgress")
	}
	// This is the price for the last transaction so increment +1
	// to get the price for next slot transaction.
	// Slots numbers should be from 0 to 4 so
	// use mod of 5 in order to save 5 as slot 0.
	slotNum.Add(slotNum, big.NewInt(1)).Mod(slotNum, big.NewInt(5))
	dbtxCostName := tellorCommon.PriceTXs + slotNum.String()
	cost, err := mgr.database.Get(dbtxCostName)
	if err != nil {
		return nil, errors.New("getting the tx eth cost from the db")
	}
	// No price record in the db yet.
	if cost == nil {
		return big.NewInt(0), nil
	}

	return big.NewInt(0).SetBytes(cost), nil
}

// saveTXCost calculates the price for a given slot.
// Since the transaction doesn't include the slot number it was submited for
// it gets the slot number as soon as the transaction passes and
// saves it in the database for profit calculations use.
func (mgr *MiningMgr) saveTXCost(ctx context.Context, tx *types.Transaction) {
	go func(tx *types.Transaction) {
		receipt, err := bind.WaitMined(ctx, mgr.ethClient, tx)
		if err != nil {
			mgr.log.Error("waiting for transaction completion for  calculating transaction cost err:%v", err)
		}
		if receipt.Status != 1 {
			mgr.log.Error("skiping unsuccessful transaction for calculating transaction cost err:%v", err)
		}
		mgr.log.Debug("submit transaction receipt", "txHash:%v", receipt.TxHash)
		gasUsed := big.NewInt(int64(receipt.GasUsed))
		txCost := gasUsed.Mul(gasUsed, tx.GasPrice())

		slotNum, err := mgr.contractGetter.GetUintVar(nil, rpc.Keccak256([]byte("slotProgress")))
		if err != nil {
			mgr.log.Error("getting slotProgress for calculating transaction cost err:%v", err)
		}

		dbtxCostName := tellorCommon.PriceTXs + slotNum.String()
		_, err = mgr.database.Put(dbtxCostName, txCost.Bytes())
		if err != nil {
			mgr.log.Error("saving transaction cost in the db err:%v", err)
		}
	}(tx)
}

func (mgr *MiningMgr) isProfitable() (bool, error) {
	txCost, err := mgr.txCost()
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
		"profit threshold checking - reward:%v, txCost:%v, profit:%v, profit margin:%v%%, profit threshold:%v%%",
		fmt.Sprintf("%.2e", float64(reward.Int64())),
		fmt.Sprintf("%.2e", float64(txCost.Int64())),
		fmt.Sprintf("%.2e", float64(profit.Int64())),
		profitPercent, mgr.profitThreshold,
	)
	if profitPercent > mgr.profitThreshold {
		return true, nil
	}
	return false, nil
}
