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
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/tcontext"
)

var (
	requestID *big.Int
	amount    *big.Int
)

type testContract struct {
}

func (t testContract) AddTip(_requestID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	fmt.Printf("Contract simulation adding tip: %v, %v\n", _requestID, _amount)
	requestID = _requestID
	amount = _amount
	return nil, nil
}

func (t testContract) SubmitSolution(solution string, requestID [5]*big.Int, value [5]*big.Int) (*types.Transaction, error) {
	return nil, nil
}

func (t testContract) DidMine(challenge [32]byte) (bool, error) {
	return false, nil
}

type testSubmit struct {
	contract *testContract
}

func (t testSubmit) Submit(ctx context.Context, proxy db.DataServerProxy, ctxName string, factoryFn tellorCommon.TransactionGeneratorFN) (*types.Transaction, error) {
	return factoryFn(ctx, *t.contract)
}

func TestRequestDataOps(t *testing.T) {
	ctx, cfg, cleanup := tcontext.CreateTestContext(t)
	defer t.Cleanup(cleanup)

	exitCh := make(chan os.Signal)

	con := &testContract{}
	submitter := testSubmit{contract: con}
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)

	// Delete any request keys.
	if err := DB.Delete(db.RequestIdKey); err != nil {
		log.Fatal(err)
	}
	if err := DB.Delete(db.TributeBalanceKey); err != nil {
		log.Fatal(err)
	}

	reqData := CreateDataRequester(exitCh, submitter, 2*time.Second, ctx.Value(tellorCommon.DataProxyKey).(db.DataServerProxy))

	// It should not request data if not configured to do it.
	cfg.RequestData = 0
	if err := reqData.Start(ctx); err != nil {
		log.Fatal(err)
	}
	time.Sleep(100 * time.Millisecond)
	if reqData.submittingRequests {
		t.Fatal("Should not be submitting requests without configured request id")
	}

	cfg.RequestData = 1
	if err := DB.Put(db.RequestIdKey, []byte(hexutil.EncodeBig(big.NewInt(0)))); err != nil {
		log.Fatal(err)
	}
	i, success := new(big.Int).SetString("999999999999999999999999999999999999999999999", 10)
	if !success {
		t.Fatal("creating a big int")
	}
	if err := DB.Put(db.TributeBalanceKey, []byte(hexutil.EncodeBig(i))); err != nil {
		log.Fatal(err)
	}

	if err := reqData.Start(ctx); err != nil {
		log.Fatal(err)
	}
	time.Sleep(3 * time.Second)
	if requestID == nil {
		t.Fatal("Should have requested data")
	}
	requestID = nil
	if err := DB.Put(db.RequestIdKey, []byte(hexutil.EncodeBig(big.NewInt(1)))); err != nil {
		log.Fatal(err)
	}
	time.Sleep(3 * time.Second)
	if requestID != nil {
		t.Fatal("Should not have requested data when a challenge request is in progress")
	}

	exitCh <- os.Interrupt
	time.Sleep(300 * time.Millisecond)
	if reqData.submittingRequests {
		t.Fatal("Should not be submitting requests after exit sig")
	}
}
