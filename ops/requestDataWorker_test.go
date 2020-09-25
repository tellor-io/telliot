// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/db"
)

var (
	requestID *big.Int
	amount    *big.Int
)

type testContract struct {
}

func (t testContract) NewSubmitSolution(solution string, requestID [5]*big.Int, value [5]*big.Int) (*types.Transaction, error) {
	return nil, nil
}

func (t testContract) AddTip(_requestID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	fmt.Printf("Contract simulation adding tip: %v, %v\n", _requestID, _amount)
	requestID = _requestID
	amount = _amount
	return nil, nil
}

func (t testContract) SubmitSolution(solution string, requestID *big.Int, value *big.Int) (*types.Transaction, error) {
	return nil, nil
}

func (t testContract) DidMine(challenge [32]byte) (bool, error) {
	return false, nil
}

type testSubmit struct {
	contract *testContract
}

func (t testSubmit) PrepareTransaction(ctx context.Context, proxy db.DataServerProxy, ctxName string, factoryFn tellorCommon.TransactionGeneratorFN) error {
	_, err := factoryFn(ctx, *t.contract)
	return err
}

func TestRequestDataOps(t *testing.T) {
	exitCh := make(chan os.Signal)
	cfg := config.GetConfig()

	con := &testContract{}
	submitter := testSubmit{contract: con}
	DB, err := db.Open(cfg.DBFile)
	if err != nil {
		log.Fatal(err)
	}
	// Delete any request id.
	err = DB.Delete(db.RequestIdKey)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.WithValue(context.Background(), tellorCommon.DBContextKey, DB)
	reqData := CreateDataRequester(exitCh, submitter, 2, ctx.Value(tellorCommon.DataProxyKey).(db.DataServerProxy))

	// t should not request data if not configured to do it
	cfg.RequestData = 0
	err = reqData.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(300 * time.Millisecond)
	if reqData.submittingRequests {
		t.Fatal("Should not be submitting requests without configured request id")
	}

	cfg.RequestData = 1
	err = DB.Put(db.RequestIdKey, []byte(hexutil.EncodeBig(big.NewInt(0))))
	if err != nil {
		log.Fatal(err)
	}
	err = reqData.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(2500 * time.Millisecond)
	if requestID == nil {
		t.Fatal("Should have requested data")
	}
	requestID = nil
	err = DB.Put(db.RequestIdKey, []byte(hexutil.EncodeBig(big.NewInt(1))))
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(2500 * time.Millisecond)
	if requestID != nil {
		t.Fatal("Should not have requested data when a challenge request is in progress")
	}

	exitCh <- os.Interrupt
	time.Sleep(300 * time.Millisecond)
	if reqData.submittingRequests {
		t.Fatal("Should not be submitting requests after exit sig")
	}
}
