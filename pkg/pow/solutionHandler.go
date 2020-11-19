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
	"path/filepath"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/tracker"
	"github.com/tellor-io/TellorMiner/pkg/util"
)

/**
* The solution handler has one purpose: to either submit the solution on-chain
* or to reject it if the miner has already submitted a solution for the challenge
* or the the solution's challenge does not match current challenge
 */

type SolutionHandler struct {
	log              *util.Logger
	proxy            db.DataServerProxy
	currentChallenge *MiningChallenge
	currentNonce     string
	currentValues    [5]*big.Int
	submitter        tellorCommon.TransactionSubmitter
}

func CreateSolutionHandler(cfg *config.Config, submitter tellorCommon.TransactionSubmitter, proxy db.DataServerProxy) *SolutionHandler {

	return &SolutionHandler{
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

	for i := 0; i < 5; i++ {
		valKey := fmt.Sprintf("%s%d", db.QueriedValuePrefix, challenge.RequestIDs[i].Uint64())
		m, err := s.proxy.BatchGet([]string{valKey})
		if err != nil {
			return nil, errors.Wrapf(err, "could not retrieve pricing data for current request id")
		}
		val := m[valKey]
		var value *big.Int
		if len(val) == 0 {
			cfg := config.GetConfig()
			indexPath := filepath.Join(cfg.ConfigFolder, "manualData.json")
			jsonFile, err := os.Open(indexPath)
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
				return nil, errors.Wrapf(err, "could not retrieve pricing data for current request id")
			}
			value = big.NewInt(int64(val))
		} else {
			value, err = hexutil.DecodeBig(string(val))
			if err != nil {
				if challenge.RequestIDs[i].Uint64() > tracker.MaxPSRID() {
					s.log.Error("problem decoding price value prior to submitting solution: %v\n", err)
					if len(val) == 0 {
						s.log.Error("0 value being submitted")
						s.currentValues[i] = big.NewInt(0)
					}
					continue
				}
				return nil, errors.Wrapf(err, "no value in database,  reg id: %v", challenge.RequestIDs[i].Uint64())
			}
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
