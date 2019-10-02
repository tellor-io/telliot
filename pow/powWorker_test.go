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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/tellor-io/TellorMiner/rpc"
	tellor "github.com/tellor-io/TellorMiner/contracts"
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

	startBal := big.NewInt(356000)

	hash := math.PaddedBigBytes(big.NewInt(256), 32)
	var b32 [32]byte
	for i, v := range hash {
		b32[i] = v
	}
	queryStr := "json(https://coinbase.com)"
	chal := &rpc.CurrentChallenge{ChallengeHash: b32, RequestID: big.NewInt(1),
		Difficulty: big.NewInt(500), QueryString: queryStr,
		Granularity: big.NewInt(1000), Tip: big.NewInt(0)}
	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0), MiningStatus:false,Top50Requests: []*big.Int{}, CurrentChallenge: chal}
	client := rpc.NewMockClientWithValues(opts)


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

	worker := CreateWorker(exitCh,1, submitter, 2)
	ctx := context.WithValue(context.Background(), tellorCommon.DBContextKey, DB)
	masterInstance := ctx.Value(tellorCommon.MasterContractContextKey)
	if masterInstance == nil {
		contractAddress := common.HexToAddress(cfg.ContractAddress)
		masterInstance, err = tellor.NewTellorMaster(contractAddress,client)
		if err != nil {
			return
		}
		ctx = context.WithValue(ctx, tellorCommon.MasterContractContextKey, masterInstance)
	}

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
