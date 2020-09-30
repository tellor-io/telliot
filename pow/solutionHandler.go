package pow

import (
	"context"
	"fmt"
	"time"
	"math/big"
	"strings"
	"os"
	"io/ioutil"
	"encoding/json"
	"strconv"
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

type SolutionHandler struct {
	log              *util.Logger
	pubKey           string
	proxy            db.DataServerProxy
	currentChallenge *MiningChallenge
	currentNonce     string
	currentValue     *big.Int
	currentValues	 [5]*big.Int
	lastSubmit		 int64
	submitter        tellorCommon.TransactionSubmitter
}

func CreateSolutionHandler(
	cfg *config.Config,
	submitter tellorCommon.TransactionSubmitter,
	proxy db.DataServerProxy) *SolutionHandler {

	//get address from config
	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)
	pubKey := strings.ToLower(fromAddress.Hex())

	return &SolutionHandler{
		pubKey:    pubKey,
		proxy:     proxy,
		submitter: submitter,
		log:       util.NewLogger("pow", "SolutionHandler"),
	}
}

func (s *SolutionHandler) Submit(ctx context.Context, result *Result) bool{
	challenge := result.Work.Challenge
	nonce := result.Nonce
	s.currentChallenge = challenge
	s.currentNonce = nonce
	manualVal := int64(0)

	valKey := fmt.Sprintf("%s%d", db.QueriedValuePrefix, challenge.RequestID.Uint64())
	s.log.Info("Getting pending txn and value from data server...")

	keys := []string{
		db.CurrentChallengeKey,
		db.RequestIdKey,
		db.RequestIdKey0,
		db.RequestIdKey1,
		db.RequestIdKey2,
		db.RequestIdKey3,
		db.RequestIdKey4,
		db.LastNewValueKey,
		valKey,
	}

	m, err :=  s.proxy.BatchGet(keys)
	address := common.HexToAddress(s.pubKey)
	dbKey := fmt.Sprintf("%s-%s", strings.ToLower(address.Hex()), db.TimeOutKey)
	lastS,err := s.proxy.Get(dbKey)
	if err != nil{
		fmt.Println("timeout Retrieval error",err)
		return false
	}
	lastB,err := hexutil.DecodeBig(string(lastS))
	if err != nil{
		fmt.Println("Timeout key decode error", lastS)
		return false
	}
	last := lastB.Int64()
	today := time.Now()
	if last > 0{
		tm := time.Unix(last, 0)
		fmt.Println("Time since last submit: ",today.Sub(tm))
		if today.Sub(tm) < time.Duration(15) * time.Minute{
			fmt.Println("Cannot submit value, within fifteen minutes")
			return false
		}
	}
	if m[db.RequestIdKey0] != nil{
		for i := 0; i < 5; i++{
			valKey := fmt.Sprintf("%s%d", db.QueriedValuePrefix, challenge.RequestIDs[i].Uint64())
			m2, err := s.proxy.BatchGet([]string{valKey})
			if err != nil {
				fmt.Println("Could not retrieve pricing data for current request id: %v", err)
				return false
			}
			val := m2[valKey]
			if val == nil || len(val) == 0 {
				if challenge.RequestIDs[i].Uint64() > 53 && (val == nil || len(val) == 0) {
					s.log.Warn("Have not retrieved price data for requestId %d. WARNING: Submitting 0 because of faulty API request", challenge.RequestID.Uint64())
				} else {
					jsonFile, err := os.Open("manualData.json")
					if err != nil {
						fmt.Println("manualData read error",err)
						return false
					}
					defer jsonFile.Close()
					byteValue, _ := ioutil.ReadAll(jsonFile)
					var result map[string]map[string]uint
					json.Unmarshal([]byte(byteValue), &result)
					_id := strconv.FormatUint(challenge.RequestIDs[i].Uint64(), 10)
					manualVal = int64(result[_id]["VALUE"])
					if manualVal == 0{
						s.log.Error("No Value in database, not submitting.",challenge.RequestIDs[i].Uint64(),2)
						return false
					}else{
						fmt.Println("Using Manually entered value: ",manualVal)
					}
				}
			}
			value, err := hexutil.DecodeBig(string(val))
			if err != nil {
				if challenge.RequestIDs[i].Uint64() > 53 {
					s.log.Error("Problem decoding price value prior to submitting solution: %v\n", err)
					if len(val) == 0 {
						s.log.Error("0 value being submitted")
						value = big.NewInt(0)
					}
				} else if manualVal > 0{
					value = big.NewInt(manualVal)
				}else{
					s.log.Error("No Value in database, not submitting here2.", challenge.RequestIDs[i].Uint64())
					return false
				}
			}
			s.currentValues[i] = value
		}
		err = s.submitter.PrepareTransaction(ctx, s.proxy, "submitSolution", s.newSubmit)
		if err != nil {
			s.log.Error("Problem submitting txn", err)
		} else {
			s.log.Info("Successfully submitted solution")
		}
	}else{
		val := m[valKey]
		if val == nil || len(val) == 0 {
			if challenge.RequestID.Uint64() > 51 && (val == nil || len(val) == 0) {
				s.log.Warn("Have not retrieved price data for requestId %d. WARNING: Submitting 0 because of faulty API request", challenge.RequestID.Uint64())
			} else {
				jsonFile, err := os.Open("manualData.json")
				if err != nil {
					fmt.Println(err)
				}
				defer jsonFile.Close()
				byteValue, _ := ioutil.ReadAll(jsonFile)
				var result map[string]map[string]int64
				json.Unmarshal([]byte(byteValue), &result)
				_id := strconv.FormatUint(challenge.RequestID.Uint64(), 10)
				manualVal = result[_id]["VALUE"]
			if manualVal == 0{
				s.log.Error("No Value in database, not submitting.")
				return false
			}else{
				fmt.Println("Using Manually entered value: ",manualVal)
			}
			}
		}
		value, err := hexutil.DecodeBig(string(val))
		if err != nil {
			if challenge.RequestID.Uint64() > 51 {
				s.log.Error("Problem decoding price value prior to submitting solution: %v\n", err)
				if len(val) == 0 {
					s.log.Error("0 value being submitted")
					value = big.NewInt(0)
				}
			} else if manualVal > 0{
				value = big.NewInt(manualVal)
				}else{
				s.log.Error("No Value in database, not submitting.")
				return false
			}
		}
		s.currentValue = value
		err = s.submitter.PrepareTransaction(ctx, s.proxy, "submitSolution", s.submit)
		if err != nil {
			s.log.Error("Problem submitting txn", err)
		} else {
			s.log.Info("Successfully submitted solution")
		}	
	}
	return true
}

func (s *SolutionHandler) submit(ctx context.Context, contract tellorCommon.ContractInterface) (*types.Transaction, error) {

	txn, err := contract.SubmitSolution(
		s.currentNonce,
		s.currentChallenge.RequestID,
		s.currentValue)
	if err != nil {
		s.log.Error("Problem submitting solution: %v", err)
		return txn, err
	}

	return txn, err
}

func (s *SolutionHandler) newSubmit(ctx context.Context, contract tellorCommon.ContractInterface) (*types.Transaction, error) {

	txn, err := contract.NewSubmitSolution(
		s.currentNonce,
		s.currentChallenge.RequestIDs,
		s.currentValues)
	if err != nil {
		s.log.Error("Problem submitting solution: %v", err)
		return txn, err
	}

	return txn, err
}
