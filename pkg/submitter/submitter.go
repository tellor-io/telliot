// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package submitter

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/mining"
	"github.com/tellor-io/telliot/pkg/rpc"
	"github.com/tellor-io/telliot/pkg/tracker"
)

const ComponentName = "submitter"

/**
* The submitter has one purpose: to either submit the solution on-chain
* or to reject it if the miner has already submitted a solution for the challenge
* or the the solution's challenge does not match current challenge
 */

type Submitter struct {
	ctx              context.Context
	close            context.CancelFunc
	logger           log.Logger
	cfg              *config.Config
	proxy            db.DataServerProxy
	account          *rpc.Account
	client           contracts.ETHClient
	contractInstance *contracts.ITellor
	currentChallenge *mining.MiningChallenge
	currentNonce     string
	currentValues    [5]*big.Int
	resultCh         chan *mining.Result
	submitCount      prometheus.Counter
	submitFailCount  prometheus.Counter
	submitProfit     *prometheus.GaugeVec
	submitCost       *prometheus.GaugeVec
	submitReward     *prometheus.GaugeVec
	lastSubmitCncl   context.CancelFunc
}

func NewSubmitter(ctx context.Context, cfg *config.Config, logger log.Logger, client contracts.ETHClient, contractInstance *contracts.ITellor, account *rpc.Account, proxy db.DataServerProxy) (*Submitter, chan *mining.Result, error) {
	filterLog, err := logging.ApplyFilter(*cfg, ComponentName, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "apply filter logger")
	}
	ctx, close := context.WithCancel(ctx)
	submitter := &Submitter{
		ctx:              ctx,
		close:            close,
		client:           client,
		proxy:            proxy,
		cfg:              cfg,
		resultCh:         make(chan *mining.Result),
		account:          account,
		logger:           log.With(filterLog, "component", ComponentName, "pubKey", account.Address.String()[:6]),
		contractInstance: contractInstance,
		submitCount: promauto.NewCounter(prometheus.CounterOpts{
			Namespace:   "telliot",
			Subsystem:   "mining",
			Name:        "submit_total",
			Help:        "The total number of submitted solutions",
			ConstLabels: prometheus.Labels{"account": account.Address.String()},
		}),
		submitFailCount: promauto.NewCounter(prometheus.CounterOpts{
			Namespace:   "telliot",
			Subsystem:   "mining",
			Name:        "submit_fails_total",
			Help:        "The total number of failed submission",
			ConstLabels: prometheus.Labels{"account": account.Address.String()},
		}),
		submitProfit: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace:   "telliot",
			Subsystem:   "mining",
			Name:        "submit_profit",
			Help:        "The current submit profit in percents",
			ConstLabels: prometheus.Labels{"account": account.Address.String()},
		},
			[]string{"slot"},
		),
		submitCost: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace:   "telliot",
			Subsystem:   "mining",
			Name:        "submit_cost",
			Help:        "The current submit cost in 1e18 eth",
			ConstLabels: prometheus.Labels{"account": account.Address.String()},
		},
			[]string{"slot"},
		),
		submitReward: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace:   "telliot",
			Subsystem:   "mining",
			Name:        "submit_reward",
			Help:        "The current reward in 1e18 eth",
			ConstLabels: prometheus.Labels{"account": account.Address.String()},
		},
			[]string{"slot"},
		),
	}

	return submitter, submitter.resultCh, nil
}

func (s *Submitter) Start() error {
	for {
		select {
		case <-s.ctx.Done():
			if s.resultCh != nil {
				close(s.resultCh)
			}
			if s.lastSubmitCncl != nil {
				s.lastSubmitCncl()
			}
			level.Info(s.logger).Log("msg", "submitter shutdown complete")
			return s.ctx.Err()
		case result := <-s.resultCh:
			if s.lastSubmitCncl != nil {
				s.lastSubmitCncl()
			}
			var ctx context.Context
			ctx, s.lastSubmitCncl = context.WithCancel(s.ctx)

			level.Info(s.logger).Log("msg", "received a solution",
				"challenge", fmt.Sprintf("%x", result.Work.Challenge),
				"solution", result.Nonce,
				"difficulty", result.Work.Challenge.Difficulty,
				"requestIDs", fmt.Sprintf("%+v", result.Work.Challenge.RequestIDs),
			)
			s.blockUntilTimeToSubmit(ctx)
			s.handleSubmit(ctx, result)
		}
	}
}

func (s *Submitter) CancelPendingSubmit() {
	if s.lastSubmitCncl != nil {
		s.lastSubmitCncl()
	}
}

func (s *Submitter) Stop() {
	s.close()
}

func (s *Submitter) blockUntilTimeToSubmit(newChallengeReplace context.Context) {
	var (
		lastSubmit time.Duration
		timestamp  *time.Time
		err        error
	)
	for {
		select {
		case <-newChallengeReplace.Done():
		default:
		}
		lastSubmit, timestamp, err = s.lastSubmit()
		if err != nil {
			level.Debug(s.logger).Log("msg", "checking last submit time", "err", err)
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
	if lastSubmit < s.cfg.Mine.MinSubmitPeriod.Duration {
		level.Info(s.logger).Log("msg", "min transaction submit threshold hasn't passed",
			"nextSubmit", time.Duration(s.cfg.Mine.MinSubmitPeriod.Nanoseconds())-lastSubmit,
			"lastSubmit", lastSubmit,
			"minSubmitPeriod", s.cfg.Mine.MinSubmitPeriod,
		)
		timeToSubmit, cncl := context.WithDeadline(newChallengeReplace, timestamp.Add(15*time.Minute))
		defer cncl()
		select {
		case <-newChallengeReplace.Done():
		case <-timeToSubmit.Done(): // 15min since last submit has passed so can unblock.
		}
	}
}

func (s *Submitter) handleSubmit(newChallengeReplace context.Context, result *mining.Result) {
	go func(newChallengeReplace context.Context, result *mining.Result) {
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-newChallengeReplace.Done():
				level.Info(s.logger).Log("msg", "canceled submit")
				return
			default:
				profitPercent, err := s.profit() // Call it regardless of whether we use so that is sets the exposed metrics.
				if err != nil {
					level.Error(s.logger).Log("msg", "submit solution profit check", "err", err)
					<-ticker.C
					continue
				}
				if s.cfg.Mine.ProfitThreshold > 0 {
					if profitPercent < int64(s.cfg.Mine.ProfitThreshold) {
						<-ticker.C
						continue
					} else { // Transaction is profitable.
						for {
							select {
							case <-newChallengeReplace.Done():
								level.Info(s.logger).Log("msg", "canceled submit")
								return
							default:
							}
							if statusID := s.getMinerStatus(); statusID != 1 {
								level.Error(s.logger).Log("msg", "miner is not in a status that can submit", "status", minerStatusName(statusID))
								<-ticker.C
								continue
							}
							s.blockUntilTimeToSubmit(newChallengeReplace)
							tx, err := s.Submit(newChallengeReplace, result)
							if err != nil {
								s.submitFailCount.Inc()
								level.Error(s.logger).Log("msg", "submiting a solution, retrying", "err", err)
								<-ticker.C
								continue
							}
							level.Debug(s.logger).Log("msg", "submited a solution", "txHash", tx.Hash().String())
							s.saveGasUsed(newChallengeReplace, tx)
							s.submitCount.Inc()
							return
						}
					}
				}
			}
		}
	}(newChallengeReplace, result)
}

func (s *Submitter) Submit(ctx context.Context, result *mining.Result) (*types.Transaction, error) {
	challenge := result.Work.Challenge
	nonce := result.Nonce
	s.currentChallenge = challenge
	s.currentNonce = nonce

	// The submit contains values for 5 data IDs so add them here.
	for i := 0; i < 5; i++ {
		valKey := fmt.Sprintf("%s%d", db.QueriedValuePrefix, challenge.RequestIDs[i].Uint64())
		m, err := s.proxy.BatchGet([]string{valKey})
		if err != nil {
			return nil, errors.Wrapf(err, "retrieve pricing for data id:%v", challenge.RequestIDs[i].Uint64())
		}
		val := m[valKey]
		var value *big.Int
		if len(val) == 0 {
			jsonFile, err := os.Open(s.cfg.ManualDataFile)
			if err != nil {
				return nil, errors.Wrapf(err, "manualData read Error")
			}
			defer jsonFile.Close()
			byteValue, _ := ioutil.ReadAll(jsonFile)
			var result map[string]map[string]uint
			_ = json.Unmarshal([]byte(byteValue), &result)
			_id := strconv.FormatUint(challenge.RequestIDs[i].Uint64(), 10)
			val := result[_id]["VALUE"]
			if val == 0 {
				return nil, errors.Errorf("retrieve pricing data for current request id")
			}
			value = big.NewInt(int64(val))
		} else {
			value, err = hexutil.DecodeBig(string(val))
			if err != nil {
				if challenge.RequestIDs[i].Uint64() > tracker.MaxPSRID() {
					level.Error(s.logger).Log(
						"msg", "decoding price value prior to submiting solution",
						"err", err,
					)
					if len(val) == 0 {
						level.Error(s.logger).Log("msg", "0 value being submitted")
						s.currentValues[i] = big.NewInt(0)
					}
					continue
				}
				return nil, errors.Errorf("no value in database,  reg id:%v", challenge.RequestIDs[i].Uint64())
			}
		}
		s.currentValues[i] = value
	}

	// Submit the solution.
	tx, err := rpc.SubmitContractTxn(ctx, s.logger, s.cfg, s.proxy, s.client, s.contractInstance, s.account, "submitSolution", s.submit)
	if err == nil {
		return tx, nil
	}
	return nil, err
}

func (s *Submitter) getMinerStatus() int64 {
	// Check if the staked account is in dispute before sending a transaction.
	statusID, _, err := s.contractInstance.GetStakerInfo(&bind.CallOpts{}, s.account.Address)
	if err != nil {
		level.Error(s.logger).Log("msg", "getting staker info from contract", "pubkey", s.account.Address.String())
		return -1
	}
	return statusID.Int64()
}

func (s *Submitter) submit(ctx context.Context, contract tellorCommon.ContractInterface) (*types.Transaction, error) {

	txn, err := contract.SubmitSolution(
		s.currentNonce,
		s.currentChallenge.RequestIDs,
		s.currentValues)
	if err != nil {
		return nil, err
	}

	return txn, err
}

func (s *Submitter) lastSubmit() (time.Duration, *time.Time, error) {
	address := "000000000000000000000000" + s.account.Address.Hex()[2:]
	decoded, err := hex.DecodeString(address)
	if err != nil {
		return 0, nil, errors.Wrapf(err, "decoding address")
	}
	last, err := s.contractInstance.GetUintVar(nil, rpc.Keccak256(decoded))

	if err != nil {
		return 0, nil, errors.Wrapf(err, "getting last submit time for:%v", s.account.Address.String())
	}
	// The Miner has never submitted so put a timestamp at the beginning of unix time.
	if last.Int64() == 0 {
		last.Set(big.NewInt(1))
	}

	lastInt := last.Int64()
	now := time.Now()
	var lastSubmit time.Duration
	var tm time.Time
	if lastInt > 0 {
		tm = time.Unix(lastInt, 0)
		lastSubmit = now.Sub(tm)
	}

	return lastSubmit, &tm, nil
}

// saveGasUsed calculates the price for a given slot.
// Since the transaction doesn't include the slot number it gets the slot number
// as soon as the transaction passes and
// saves it in the database for profit calculations.
// TODO[Krasi] To be more detirministic and simplify this
// should get the `_SLOT_PROGRESS` and `gasUsed` from the `NonceSubmitted` event.
// At the moment there is a slight chance of a race condition if
// another transaction has passed between checking the transaction cost and
// checking the `_SLOT_PROGRESS`
// Tracking issue https://github.com/tellor-io/TellorCore/issues/101
func (s *Submitter) saveGasUsed(ctx context.Context, tx *types.Transaction) {
	go func(tx *types.Transaction) {
		receipt, err := bind.WaitMined(ctx, s.client, tx)
		if err != nil {
			level.Error(s.logger).Log("msg", "transaction result for calculating transaction cost", "err", err)
			return
		}
		if receipt.Status != 1 {
			s.submitFailCount.Inc()
			level.Error(s.logger).Log("msg", "unsuccessful submitSolution transaction, not saving the tx cost in the db", "txHash", receipt.TxHash.String())
			return
		}

		gasUsed := big.NewInt(int64(receipt.GasUsed))
		slotNum, err := s.contractInstance.GetUintVar(nil, rpc.Keccak256([]byte("_SLOT_PROGRESS")))
		if err != nil {
			level.Error(s.logger).Log("msg", "getting _SLOT_PROGRESS for calculating transaction cost", "err", err)
		}

		txID := tellorCommon.PriceTXs + slotNum.String()
		err = s.proxy.Put(txID, gasUsed.Bytes())
		if err != nil {
			level.Error(s.logger).Log("msg", "saving transaction cost", "err", err)
		}
		level.Info(s.logger).Log("msg", "saved transaction gas used", "txHash", receipt.TxHash.String(), "amount", gasUsed.Int64(), "slot", slotNum.Int64())

		amount, err := s.getETHBalance()
		if err != nil {
			level.Error(s.logger).Log("msg", "getting ETH balance", "err", err)
		} else {
			level.Info(s.logger).Log("msg", "ETH balance", "amount", amount)

		}

		amountTRB, err := s.getTRBBalance()
		if err != nil {
			level.Error(s.logger).Log("msg", "getting TRB balance", "err", err)
		} else {
			level.Info(s.logger).Log("msg", "TRB balance", "amount", amountTRB)
		}

	}(tx)
}

func (s *Submitter) getTRBBalance() (string, error) {
	balance, err := s.contractInstance.BalanceOf(nil, s.account.Address)
	if err != nil {
		return "", errors.Wrap(err, "retrieving trb balance")
	}
	balanceH, _ := big.NewFloat(1).SetString(balance.String())
	decimals, _ := big.NewFloat(1).SetString("1000000000000000000")
	if decimals != nil {
		balanceH = balanceH.Quo(balanceH, decimals)
	}
	return balanceH.String(), nil
}

func (s *Submitter) getETHBalance() (string, error) {
	balance, err := s.client.BalanceAt(s.ctx, s.account.Address, nil)
	if err != nil {
		return "", errors.Wrap(err, "retrieving balance")
	}
	balanceH, _ := big.NewFloat(1).SetString(balance.String())
	decimals, _ := big.NewFloat(1).SetString("1000000000000000000")
	if decimals != nil {
		balanceH = balanceH.Quo(balanceH, decimals)
	}
	return balanceH.String(), nil
}

// currentReward returns the current TRB rewards converted to ETH.
// TODO[Krasi] This is a duplicate code from the tellor conract so
// Should add `currentReward` func to the contract to avoid this code duplication.
// Tracking issue https://github.com/tellor-io/TellorCore/issues/101
func (s *Submitter) currentReward() (*big.Int, error) {
	timeOfLastNewValue, err := s.contractInstance.GetUintVar(nil, rpc.Keccak256([]byte("_TIME_OF_LAST_NEW_VALUE")))
	if err != nil {
		return nil, errors.New("getting _TIME_OF_LAST_NEW_VALUE")
	}
	totalTips, err := s.contractInstance.GetUintVar(nil, rpc.Keccak256([]byte("_CURRENT_TOTAL_TIPS")))
	if err != nil {
		return nil, errors.New("getting _CURRENT_TOTAL_TIPS")
	}

	timeDiff := big.NewInt(time.Now().Unix() - timeOfLastNewValue.Int64())
	trb := big.NewInt(1e18)
	rewardPerSec := big.NewInt(0).Div(trb, big.NewInt(300)) // 1 TRB every 5 minutes so total reward is timeDiff multiplied by reward per second.
	rewardTRB := big.NewInt(0).Mul(rewardPerSec, timeDiff)

	singleMinerTip := big.NewInt(0).Div(totalTips, big.NewInt(10)) // Half of the tips are burned(remain in the contract) to reduce inflation.
	rewardWithTips := big.NewInt(0).Add(singleMinerTip, rewardTRB)

	if rewardWithTips == big.NewInt(0) {
		return big.NewInt(0), nil
	}

	return s.convertTRBtoETH(rewardWithTips)
}

func (s *Submitter) convertTRBtoETH(trb *big.Int) (*big.Int, error) {
	val, err := s.proxy.Get(db.QueriedValuePrefix + strconv.Itoa(tracker.RequestID_TRB_ETH))
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
	wei := big.NewInt(tellorCommon.WEI)
	precisionUpscale := big.NewInt(0).Div(wei, big.NewInt(tracker.PSRs[tracker.RequestID_TRB_ETH].Granularity()))
	priceTRB.Mul(priceTRB, precisionUpscale)

	eth := big.NewInt(0).Mul(priceTRB, trb)
	eth.Div(eth, big.NewInt(1e18))
	return eth, nil
}

func (s *Submitter) gasUsed() (*big.Int, *big.Int, error) {
	slotNum, err := s.contractInstance.GetUintVar(nil, rpc.Keccak256([]byte("_SLOT_PROGRESS")))
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting _SLOT_PROGRESS")
	}
	// This is the price for the last transaction so increment +1
	// to get the price for next slot transaction.
	// Slots numbers should be from 0 to 4 so
	// use mod of 5 in order to save 5 as slot 0.
	slotNum.Add(slotNum, big.NewInt(1)).Mod(slotNum, big.NewInt(5))
	txID := tellorCommon.PriceTXs + slotNum.String()
	gas, err := s.proxy.Get(txID)
	if err != nil {
		return nil, nil, errors.New("getting the tx eth cost from the db")
	}

	// No price record in the db yet.
	if gas == nil {
		return big.NewInt(0), slotNum, nil
	}

	return big.NewInt(0).SetBytes(gas), slotNum, nil
}

// profit returns the profit in percents.
// When the transaction cost is unknown it returns -1 so
// that the caller can decide how to handle.
// Transaction cost is zero when the manager hasn't done any transactions yet.
// Each transaction cost is known for any siquential transactions.
func (s *Submitter) profit() (int64, error) {
	gasUsed, slotNum, err := s.gasUsed()
	if err != nil {
		return 0, errors.Wrap(err, "getting TX cost")
	}
	if gasUsed.Int64() == 0 {
		level.Debug(s.logger).Log("msg", "profit checking:no data for gas used", "slot", slotNum)
		return -100, nil
	}
	gasPrice, err := s.client.SuggestGasPrice(context.Background())
	if err != nil {
		return 0, errors.Wrap(err, "getting gas price")
	}
	reward, err := s.currentReward()
	if err != nil {
		return 0, errors.Wrap(err, "getting current rewards")
	}

	txCost := gasPrice.Mul(gasPrice, gasUsed)
	profit := big.NewInt(0).Sub(reward, txCost)
	profitPercentFloat := float64(profit.Int64()) / float64(txCost.Int64()) * 100
	profitPercent := int64(profitPercentFloat)
	level.Debug(s.logger).Log(
		"msg", "profit checking",
		"reward", fmt.Sprintf("%.2e", float64(reward.Int64())),
		"txCost", fmt.Sprintf("%.2e", float64(txCost.Int64())),
		"slot", slotNum,
		"profit", fmt.Sprintf("%.2e", float64(profit.Int64())),
		"profitMargin", profitPercent,
		"profitThreshold", s.cfg.Mine.ProfitThreshold,
	)

	s.submitProfit.With(prometheus.Labels{"slot": strconv.Itoa(int(slotNum.Int64()))}).(prometheus.Gauge).Set(float64(profitPercent))
	s.submitCost.With(prometheus.Labels{"slot": strconv.Itoa(int(slotNum.Int64()))}).(prometheus.Gauge).Set(float64(txCost.Int64()))
	s.submitReward.With(prometheus.Labels{"slot": strconv.Itoa(int(slotNum.Int64()))}).(prometheus.Gauge).Set(float64(reward.Int64()))

	return profitPercent, nil
}

func minerStatusName(statusID int64) string {
	// From https://github.com/tellor-io/tellor3/blob/7c2f38a0e3f96631fb0f96e0d0a9f73e7b355766/contracts/TellorStorage.sol#L41
	switch statusID {
	case 0:
		return "Not staked"
	case 1:
		return "Staked"
	case 2:
		return "LockedForWithdraw"
	case 3:
		return "OnDispute"
	case 4:
		return "ReadyForUnlocking"
	case 5:
		return "ReadyForUnlocking"
	default:
		return "Unknown"
	}
}
