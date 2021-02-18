// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package pow

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	tellorCommon "github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/tracker"
)

/**
* The solution handler has one purpose: to either submit the solution on-chain
* or to reject it if the miner has already submitted a solution for the challenge
* or the the solution's challenge does not match current challenge
 */

type SolutionHandler struct {
	logger           log.Logger
	cfg              *config.Config
	proxy            db.DataServerProxy
	currentChallenge *MiningChallenge
	currentNonce     string
	currentValues    [5]*big.Int
	submitters       []tellorCommon.TransactionSubmitter
}

func CreateSolutionHandler(cfg *config.Config, logger log.Logger, submitters []tellorCommon.TransactionSubmitter, proxy db.DataServerProxy) *SolutionHandler {
	return &SolutionHandler{
		proxy:      proxy,
		submitters: submitters,
		cfg:        cfg,
		logger:     log.With(logger, "component", ComponentName),
	}
}

func (s *SolutionHandler) Submit(ctx context.Context, result *Result) (*types.Transaction, error) {
	challenge := result.Work.Challenge
	nonce := result.Nonce
	s.currentChallenge = challenge
	s.currentNonce = nonce

	for i := 0; i < 5; i++ {
		valKey := fmt.Sprintf("%s%d", db.QueriedValuePrefix, challenge.RequestIDs[i].Uint64())
		m, err := s.proxy.BatchGet([]string{valKey})
		if err != nil {
			return nil, errors.Wrapf(err, "retrieve pricing data for current request id")
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
						"msg", "decoding price value prior to submitt ing solution",
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

	// Try to submit by any submitter.
	for _, submitter := range s.submitters {
		lastSubmit, err := s.lastSubmit(submitter.Address())
		if err != nil {
			level.Error(s.logger).Log("msg", "checking last submit time", "err", err)
		} else if lastSubmit < s.cfg.Mine.MinSubmitPeriod.Duration {
			level.Debug(s.logger).Log("msg", "min transaction submit threshold hasn't passed", "minSubmitPeriod", s.cfg.Mine.MinSubmitPeriod, "lastSubmit", lastSubmit)
			continue
		}
		tx, err := submitter.Submit(ctx, s.proxy, "submitSolution", s.submit)
		if err == nil {
			return tx, nil
		}
		level.Error(s.logger).Log("msg", "submit solution", "pubkey", submitter.Address().String())
	}
	return nil, errors.New("submitting solution txn by any account")
}

func (s *SolutionHandler) submit(ctx context.Context, contract tellorCommon.ContractInterface) (*types.Transaction, error) {

	txn, err := contract.SubmitSolution(
		s.currentNonce,
		s.currentChallenge.RequestIDs,
		s.currentValues)
	if err != nil {
		return nil, err
	}

	return txn, err
}

func (s *SolutionHandler) lastSubmit(address common.Address) (time.Duration, error) {

	dbKey := fmt.Sprintf("%s-%s", strings.ToLower(address.Hex()), db.TimeOutKey)
	last, err := s.proxy.Get(dbKey)
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
