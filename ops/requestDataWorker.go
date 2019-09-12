package ops

import (
	"context"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/util"
)

//DataRequester responsible for submitting tips to request data periodically if configured to do so
type DataRequester struct {
	submittingRequests bool
	exitCh             chan os.Signal
	log                *util.Logger
	checkInterval      time.Duration
	submitter          tellorCommon.TransactionSubmitter
}

//CreateDataRequester creates a requester instance
func CreateDataRequester(exitCh chan os.Signal, submitter tellorCommon.TransactionSubmitter, checkIntervalSeconds time.Duration) *DataRequester {
	if checkIntervalSeconds == 0 {
		checkIntervalSeconds = 30
	}
	return &DataRequester{exitCh: exitCh, submitter: submitter, checkInterval: checkIntervalSeconds * time.Second, log: util.NewLogger("ops", "DataRequester")}
}

//Start kicks of go routines to periodically submit tips if configured to do so
func (r *DataRequester) Start(ctx context.Context) error {
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}

	//if we're not configured to request anything.
	if cfg.RequestData == 0 {
		return nil
	}

	r.submittingRequests = true
	ticker := time.NewTicker(r.checkInterval)
	go func() {
		for {
			select {
			case _ = <-r.exitCh:
				{
					r.log.Info("Stopping data requester")
					r.submittingRequests = false
					ticker.Stop()
					return
				}
			case _ = <-ticker.C:
				{
					r.maybeRequestData(ctx)
				}
			}
		}
	}()
	return nil
}

//IsRunning checks whether this requester is requesting data
func (r *DataRequester) IsRunning() bool {
	return r.submittingRequests
}

func (r *DataRequester) reqDataCallback(ctx context.Context, contract tellorCommon.ContractInterface) (*types.Transaction, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	//if we're not configured to request anything.
	if cfg.RequestData == 0 {
		r.log.Warn("Getting asked to request data when not configured to do so")
		return nil, nil
	}

	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)
	requestID, err := DB.Get(db.RequestIdKey)
	if err != nil {
		r.log.Error("Problem reading request ID from db")
		return nil, err
	}
	if len(requestID) == 0 {
		r.log.Info("No request ID stored yet. Will wait for next cycle before attempting to submit data request")
		return nil, nil
	}

	asInt, err := hexutil.DecodeBig(string(requestID))
	if err != nil {
		return nil, err
	}
	if asInt.Cmp(big.NewInt(0)) != 0 {
		r.log.Info("There is a challenge being mined right now so will not request data")
		return nil, nil
	}
	r.log.Info("Submitting tip for requestID: %v\n", cfg.RequestData)
	return contract.AddTip(big.NewInt(int64(cfg.RequestData)), big.NewInt(0))
}

func (r *DataRequester) maybeRequestData(ctx context.Context) {
	r.log.Info("Checking whether to submit data request...")
	err := r.submitter.PrepareTransaction(ctx, "RequestData", r.reqDataCallback)
	if err != nil {
		r.log.Error("Problem preparing contract transaction: %v\n", err)
	}
	r.log.Info("Finished checking whether to submit data")
}
