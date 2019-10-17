package pow

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/util"
)

/**
 * The solution handler has one purpose: to either submit the solution on-chain
 * or to reject it if the miner has already submitted a solution for the challenge
 * or the the solution's challenge does not match current challenge
 */

type solutionHandler struct {
	id              int
	solutionCh      solutionChannel
	exitCh          exitChannel
	log             *util.Logger
	pubKey          string
	proxy           db.DataServerProxy
	currentSolution *miningSolution
	currentValue    *big.Int
	submitter       tellorCommon.TransactionSubmitter
}

func createSolutionHandler(
	id int,
	solutionCh solutionChannel,
	submitter tellorCommon.TransactionSubmitter,
	proxy db.DataServerProxy) (*solutionHandler, error) {

	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Printf("Problem getting config, can't proceed at all: %v\n", err)
		log.Fatal(err)
	}

	//get address from config
	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)
	pubKey := strings.ToLower(fromAddress.Hex())

	return &solutionHandler{
		id:         id,
		pubKey:     pubKey,
		solutionCh: solutionCh,
		exitCh:     make(exitChannel),
		proxy:      proxy,
		submitter:  submitter,
		log:        util.NewLogger("pow", "SolutionHandler-"+strconv.Itoa(id)),
	}, nil
}

func (s *solutionHandler) Start(ctx context.Context) {
	s.log.Info("Starting solution handler")
	go func() {
		for {
			select {
			case _ = <-s.exitCh:
				{
					s.log.Info("Stopping solution handler on OS interrupt")
					return
				}
			case sol := <-s.solutionCh:
				{
					if sol != nil {
						go s.handleSolution(ctx, sol)
					}
				}
			}
		}
	}()
}

func (s *solutionHandler) handleSolution(ctx context.Context, sol *miningSolution) {
	if sol.nonce == "" {
		s.log.Info("Ignoring empty solution")
		return
	}

	if s.currentSolution != nil &&
		bytes.Compare(s.currentSolution.challenge.challenge, sol.challenge.challenge) == 0 {
		s.log.Warn("Getting submission multiple times for same challenge: %+v", sol.challenge)
		return
	}

	submitMutex.Lock()
	defer submitMutex.Unlock()

	reqID := sol.challenge.requestID.Uint64()
	pendKey := s.pubKey + "-" + db.PendingChallengeKey

	valKey := fmt.Sprintf("%s%d", db.QueriedValuePrefix, reqID)
	s.log.Info("Getting pending txn and value from data server...")
	m, err := s.proxy.BatchGet([]string{db.CurrentChallengeKey, db.RequestIdKey, pendKey, valKey})

	if err != nil {
		s.log.Error("Problem reading pending txn: %v", err)
		return
	}
	s.log.Debug("Retrieved data from data server %v", m)
	pending := m[pendKey]
	if pending != nil && len(pending) > 0 && bytes.Compare(pending, sol.challenge.challenge) == 0 {
		s.log.Info("Already mined given solution, will not submit new solution")
		return
	}

	val := m[valKey]
	if val == nil || len(val) == 0 {
		s.log.Warn("Have not retrieved price data for requestId %d. We can't submit solution until we've received value at least once", reqID)
		return
	}

	value, err := hexutil.DecodeBig(string(val))
	if err != nil {
		s.log.Error("Problem decoding price value prior to submitting solution: %v\n", err)
		return
	}

	currentReqId := m[db.RequestIdKey]
	if currentReqId == nil || len(currentReqId) == 0 {
		s.log.Error("Could not get current request id. Will not submit txn")
		return
	}
	cReqID, err := hexutil.DecodeBig(string(currentReqId))
	if cReqID.Uint64() != reqID {
		s.log.Error("Incoming solution request id does not match current challenge request id. Will not submit txn")
		return
	}

	chal := m[db.CurrentChallengeKey]
	if chal == nil || len(chal) == 0 {
		s.log.Error("Could not read current challenge. Will not submit txn")
		return
	}
	if bytes.Compare(chal, sol.challenge.challenge) != 0 {
		s.log.Error("Current challenge has changed during mining cycle, will not submit solution")
		return
	}

	s.currentSolution = sol
	s.currentValue = value

	//we're going to submit the solution but first, to mitigate, not eliminate, a race
	//condition between mining threads on the data server, we cache the pending challenge
	//submission. If two miners solve at nearly the same time, it's possible multiple submissions
	//will go through since both of them will have checked the data server first, gotten no pending
	//match, and will proceed to submit. But this at least gives us a chance to prevent multiple
	//submissions
	s.log.Info("Caching pending txn on data server...")
	_, err = s.proxy.BatchPut([]string{pendKey}, [][]byte{
		sol.challenge.challenge,
	})
	if err != nil {
		s.log.Error("Could not caching pending txn on data server, but will submit txn anyway: %v", err)
	}

	s.log.Info("Submitting solution to contract...")
	err = s.submitter.PrepareTransaction(ctx, "submitSolution", s.submit)
	if err != nil {
		s.log.Error("Problem submitting txn. Will remove pending txn from data server; however, "+
			"this could result in multiple nonce submissions if running multiple miners: %v", err)
		_, err = s.proxy.BatchPut([]string{
			pendKey,
		}, [][]byte{
			//basically a pending challenge that will never match any other challenge
			[]byte(""),
		})
		if err != nil {
			s.log.Error("Could not remove pending txn from data server. "+
				" Miner will likely be stuck until a new challenge is issued: %v", err)
		}
	} else {
		s.log.Info("Successfully submitted solution")
	}
}

func (s *solutionHandler) submit(ctx context.Context, contract tellorCommon.ContractInterface) (*types.Transaction, error) {
	if s.currentSolution == nil || s.currentSolution.nonce == "" {
		s.log.Warn("Somehow attempting to submit incomplete solution")
		return nil, nil
	}

	txn, err := contract.SubmitSolution(
		s.currentSolution.nonce,
		s.currentSolution.challenge.requestID,
		s.currentValue)
	if err != nil {
		s.log.Error("Problem submitting solution: %v", err)
		return txn, err
	}
	dbKeys := []string{db.PendingChallengeKey}
	values := make([][]byte, 1)
	values[0] = s.currentSolution.challenge.challenge
	_, err = s.proxy.BatchPut(dbKeys, values)
	if err != nil {
		s.log.Error("Problem writing pending challenge to remote data server: %v", err)
	}

	return txn, err
}
