package pow

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/core/types"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
)

type solutionContract struct {
	handler func(id *big.Int, nonce string, value *big.Int)
}

func (s solutionContract) AddTip(_requestID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return nil, nil
}

func (s solutionContract) DidMine(challenge [32]byte) (bool, error) {
	return false, nil
}

func (s solutionContract) SubmitSolution(_solution string, _requestID *big.Int, _value *big.Int) (*types.Transaction, error) {
	fmt.Printf("TestHarness getting solution: %s for request %v with value: %v\n", _solution, _requestID, _value)
	s.handler(_requestID, _solution, _value)
	return nil, nil
}

type testSolutionSubmit struct {
	contract *solutionContract
}

func (s *testSolutionSubmit) PrepareTransaction(ctx context.Context, ctxName string, fn tellorCommon.TransactionGeneratorFN) error {
	_, err := fn(ctx, *s.contract)
	return err
}

type solutionTestHarness struct {
	contract        *solutionContract
	solutionHandler *solutionHandler
	issuedSolution  *miningSolution
	DB              db.DB
	proxy           db.DataServerProxy
	solutionCh      solutionChannel
	solutionReqID   *big.Int
	solutionNonce   string
	solutionValue   *big.Int
}

func (th *solutionTestHarness) init(t *testing.T, cb func(DB db.DB)) {
	DB, err := db.Open("/tmp/solutionHandlerTest")
	if err != nil {
		t.Fatal(err)
	}
	localProxy, err := db.OpenLocalProxy(DB)
	if err != nil {
		t.Fatal(err)
	}

	con := &solutionContract{
		handler: func(rid *big.Int, n string, v *big.Int) {
			th.solutionReqID = rid
			th.solutionNonce = n
			th.solutionValue = v
		},
	}
	submitter := &testSolutionSubmit{contract: con}

	th.solutionCh = make(solutionChannel)
	handler, err := createSolutionHandler(1, th.solutionCh, submitter, localProxy)
	if err != nil {
		t.Fatal(err)
	}
	th.contract = con
	th.DB = DB
	th.proxy = localProxy
	th.solutionHandler = handler
	cb(DB)
}

func TestSolutionHandlerNormal(t *testing.T) {
	th := &solutionTestHarness{}

	th.init(t, func(DB db.DB) {
		pendKey := fmt.Sprintf("%s-%s", th.solutionHandler.pubKey, db.PendingChallengeKey)
		valKey := fmt.Sprintf("%s%d", db.QueriedValuePrefix, 1)

		err := deleteFromDB(th.DB, []string{pendKey})
		if err != nil {
			t.Fatal(err)
		}
		err = writeToDB(DB, []string{valKey}, [][]byte{
			[]byte(hexutil.EncodeBig(big.NewInt(10000))),
		})
		if err != nil {
			t.Fatal(err)
		}
	})

	th.solutionHandler.Start(context.Background())

	//send solution with nonce and challenge
	challenge := &miningChallenge{
		challenge:  []byte("someChallenge"),
		difficulty: big.NewInt(100),
		requestID:  big.NewInt(1),
	}
	sol := &miningSolution{
		challenge: challenge,
		nonce:     "012345",
	}
	th.solutionCh <- sol
	time.Sleep(100 * time.Millisecond)
	if th.solutionNonce != sol.nonce {
		t.Fatalf("Expected nonce to be submitted from mining solution: %s, %s", th.solutionNonce, sol.nonce)
	}
}

func TestSolutionHandlerPending(t *testing.T) {
	th := &solutionTestHarness{}

	th.init(t, func(DB db.DB) {
		pendKey := fmt.Sprintf("%s-%s", th.solutionHandler.pubKey, db.PendingChallengeKey)
		valKey := fmt.Sprintf("%s%d", db.QueriedValuePrefix, 1)

		err := writeToDB(DB, []string{
			pendKey,
			valKey,
		}, [][]byte{
			[]byte("someChallenge"),
			[]byte(hexutil.EncodeBig(big.NewInt(10000))),
		})
		if err != nil {
			t.Fatal(err)
		}
	})

	th.solutionHandler.Start(context.Background())

	//send solution with nonce and challenge
	challenge := &miningChallenge{
		challenge:  []byte("someChallenge"),
		difficulty: big.NewInt(100),
		requestID:  big.NewInt(1),
	}
	sol := &miningSolution{
		challenge: challenge,
		nonce:     "012345",
	}
	th.solutionCh <- sol
	time.Sleep(100 * time.Millisecond)
	if th.solutionNonce == sol.nonce {
		t.Fatalf("Did not expect pending txn to be submitted again: %s, %s", th.solutionNonce, sol.nonce)
	}
}
