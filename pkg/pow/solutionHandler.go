// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package pow

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/util"
)

/**
* The solution handler has one purpose: to either submit the solution on-chain
* or to reject it if the miner has already submitted a solution for the challenge
* or the the solution's challenge does not match current challenge
 */

type SolutionHandler struct {
	log              *util.Logger
	pubKey           string
	proxy            db.DataServerProxy
	currentChallenge *MiningChallenge
	currentNonce     string
	currentValues    [5]*big.Int
	submitter        tellorCommon.TransactionSubmitter
}

func CreateSolutionHandler(cfg *config.Config, submitter tellorCommon.TransactionSubmitter, proxy db.DataServerProxy) *SolutionHandler {
	// Get address from config
	_fromAddress := cfg.PublicAddress

	// Convert to address
	fromAddress := common.HexToAddress(_fromAddress)
	pubKey := strings.ToLower(fromAddress.Hex())

	return &SolutionHandler{
		pubKey:    pubKey,
		proxy:     proxy,
		submitter: submitter,
		log:       util.NewLogger("pow", "SolutionHandler"),
	}
}

func (s *SolutionHandler) Submit(ctx context.Context, result *Result) (*types.Transaction, error) {
	challenge := result.Work.Challenge
	nonce := result.Nonce
	s.currentChallenge = challenge
	s.currentNonce = nonce

	s.log.Info("Getting pending txn and value from data server...")

	address := common.HexToAddress(s.pubKey)
	dbKey := fmt.Sprintf("%s-%s", strings.ToLower(address.Hex()), db.TimeOutKey)
	lastS, err := s.proxy.Get(dbKey)
	if err != nil {
		return nil, errors.Wrapf(err, "timeout retrieval error")
	}
	lastB, err := hexutil.DecodeBig(string(lastS))
	if err != nil {
		return nil, errors.Wrapf(err, "timeout key decode last:%v", lastS)
	}
	last := lastB.Int64()
	today := time.Now()
	if last > 0 {
		tm := time.Unix(last, 0)
		fmt.Println("Time since last submit: ", today.Sub(tm))
		if today.Sub(tm) < time.Duration(15)*time.Minute {
			return nil, errors.New("cannot submit value, within fifteen minutes")
		}
	}
	for i := 0; i < 5; i++ {
		valKey := fmt.Sprintf("%s%d", db.QueriedValuePrefix, challenge.RequestIDs[i].Uint64())
		m, err := s.proxy.BatchGet([]string{valKey})
		if err != nil {
			return nil, errors.Wrapf(err, "could not retrieve pricing data for current request id")
		}
		val := m[valKey]
		if len(val) == 0 {
			s.log.Warn("have not retrieved price data for requestId %d. WARNING: Submitting 0 because of faulty API request", challenge.RequestIDs[i].Uint64())
		}
		value, err := hexutil.DecodeBig(string(val))
		if err != nil {
			if challenge.RequestIDs[i].Uint64() > 53 {
				s.log.Error("problem decoding price value prior to submitting solution: %v\n", err)
				if len(val) == 0 {
					s.log.Error("0 value being submitted")
					s.currentValues[i] = big.NewInt(0)
				}
				continue
			}
			return nil, errors.Errorf("no value in database,  reg id: %v", challenge.RequestIDs[i].Uint64())
		}
		s.currentValues[i] = value
	}
	tx, err := s.submitter.Submit(ctx, s.proxy, "submitSolution", s.submit)
	if err != nil {
		return nil, errors.Wrap(err, "submitting solution txn")
	}
	return tx, nil
}

func (s *SolutionHandler) submit(ctx context.Context, contract tellorCommon.ContractInterface) (*types.Transaction, error) {

	txn, err := contract.SubmitSolution(
		s.currentNonce,
		s.currentChallenge.RequestIDs,
		s.currentValues)
	if err != nil {
		s.log.Error("Problem submitting solution: %v", err)
		return txn, err
	}

	return txn, err
}
