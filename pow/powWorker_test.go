package pow

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
	nonce     string
	value     *big.Int
)

type testContract struct {
	didMine bool
}

func (t testContract) AddTip(_requestID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return nil, nil
}

func (t testContract) SubmitSolution(_solution string, _requestID *big.Int, _value *big.Int) (*types.Transaction, error) {
	fmt.Printf("Getting solution: %s for request %v with value: %v\n", _solution, _requestID, _value)
	requestID = _requestID
	nonce = _solution
	value = _value
	return nil, nil
}

func (t testContract) DidMine(challenge [32]byte) (bool, error) {
	return t.didMine, nil
}

type testSubmit struct {
	contract *testContract
}

func (t testSubmit) PrepareTransaction(ctx context.Context, ctxName string, fn tellorCommon.TransactionGeneratorFN) error {
	_, err := fn(ctx, *t.contract)
	return err
}

func TestWorker(t *testing.T) {
	exitCh := make(chan os.Signal)
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	if len(cfg.DBFile) == 0 {
		log.Fatal("Missing dbFile config setting")
	}

	DB, err := db.Open(cfg.DBFile)
	if err != nil {
		log.Fatal(err)
	}
	DB.Delete(db.CurrentChallengeKey)
	DB.Delete(db.RequestIdKey)
	DB.Delete(db.DifficultyKey)
	valueKey := fmt.Sprintf("%s%d", db.QueriedValuePrefix, 1)
	DB.Delete(valueKey)

	con := &testContract{didMine: false}
	submitter := &testSubmit{contract: con}

	worker := CreateWorker(exitCh, submitter, 2)
	ctx := context.WithValue(context.Background(), tellorCommon.DBContextKey, DB)
	worker.Start(ctx)
	//mining check starts immediately, just let it get setup and check
	time.Sleep(300 * time.Millisecond)
	if worker.mining {
		t.Fatal("Should not be mining")
	}

	//should process a saved challenge
	challenge := "someLongNonceString"
	DB.Put(db.CurrentChallengeKey, []byte(challenge))
	DB.Put(db.RequestIdKey, []byte(hexutil.EncodeBig(big.NewInt(1))))
	DB.Put(db.DifficultyKey, []byte(hexutil.EncodeBig(big.NewInt(5))))
	//no value yet, verify that we will not send txn until value present

	time.Sleep(3 * time.Second)
	//should have a solution
	if nonce != "" {
		t.Fatal("Should not have solved nonce without first getting price data value")
	}
	DB.Put(valueKey, []byte(hexutil.EncodeBig(big.NewInt(10000))))
	time.Sleep(3 * time.Second)
	if nonce == "" {
		t.Fatal("Should have mined nonce with price data and active challenge")
	}
	nonce = ""
	requestID = nil
	value = nil

	//we want to simulate a challenge interruption where we're working on a challenge
	//but get a new challenge (indicating it's already been mined)
	DB.Put(db.CurrentChallengeKey, []byte("anotherChallenge"))
	DB.Put(db.DifficultyKey, []byte(hexutil.EncodeBig(big.NewInt(1000000000))))
	time.Sleep(2500 * time.Millisecond)
	if !worker.mining {
		t.Fatal("Should have started mining new challenge")
	}

	//now we should be able to interrupt mining with new challenge and end
	//within time cycle
	DB.Put(db.CurrentChallengeKey, []byte("simpleChallenge"))
	DB.Put(db.DifficultyKey, []byte(hexutil.EncodeBig(big.NewInt(5))))

	time.Sleep(3 * time.Second)
	if nonce == "" {
		t.Fatal("Should have mined another solution with new challenge request")
	}

	exitCh <- os.Interrupt
	time.Sleep(time.Second)
	if worker.mining || worker.canMine {
		t.Fatalf("Miner did not shutdown gracefully after kill sig. Mining: %v CanMine: %v", worker.mining, worker.canMine)
	}
}
