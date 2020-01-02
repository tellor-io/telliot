package pow

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	tellor "github.com/tellor-io/TellorMiner/contracts"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rest"
	"github.com/tellor-io/TellorMiner/rpc"
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
	//return t.didMine, nil
	return len(nonce) > 0, nil
}

type testSubmit struct {
	contract *testContract
}

func (t testSubmit) PrepareTransaction(ctx context.Context, proxy db.DataServerProxy,ctxName string, fn tellorCommon.TransactionGeneratorFN) error {
	_, err := fn(ctx, *t.contract)
	return err
}

func setupDataProxy(DB db.DB, asLocal bool) (db.DataServerProxy, *rest.Server, error) {
	if asLocal {
		proxy, err := db.OpenLocalProxy(DB)
		return proxy, nil, err
	}

	proxy, err := db.OpenRemoteDB(DB)
	if err != nil {
		return nil, nil, err
	}
	cfg, err := config.GetConfig()
	ctx := context.WithValue(context.Background(), tellorCommon.DBContextKey, DB)
	ctx = context.WithValue(ctx, tellorCommon.DataProxyKey, proxy)
	srv, err := rest.Create(ctx, cfg.ServerHost, cfg.ServerPort)
	if err != nil {
		return nil, nil, err
	}
	srv.Start()
	return proxy, srv, nil
}

func TestWorker(t *testing.T) {
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

	proxy, server, err := setupDataProxy(DB, true)

	if err != nil {
		log.Fatal(err)
	}

	runTest(t, DB, proxy, server, 1)
}

func TestWorkerRemoteProxy(t *testing.T) {
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

	proxy, server, err := setupDataProxy(DB, false)

	if err != nil {
		log.Fatal(err)
	}

	runTest(t, DB, proxy, server, 1)
}

func TestMultiMiner(t *testing.T) {
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

	proxy, server, err := setupDataProxy(DB, false)

	if err != nil {
		log.Fatal(err)
	}

	runTest(t, DB, proxy, server, 2)
}

func _deleteFromDB(DB db.DB, keys []string) error {
	for _, k := range keys {
		err := DB.Delete(k)
		if err != nil {
			return err
		}
	}
	return nil
}
func _writeToDB(DB db.DB, keys []string, values [][]byte) error {
	for i, k := range keys {
		err := DB.Put(k, values[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func runTest(t *testing.T, DB db.DB, proxy db.DataServerProxy, server *rest.Server, minerCount int) {

	if server != nil {
		t.Log("Stopping server...")
		defer server.Stop()
	}

	defer DB.Close()

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	_fromAddress := cfg.PublicAddress
	pubKey := strings.ToLower(common.HexToAddress(_fromAddress).Hex())

	con := &testContract{didMine: false}
	submitter := &testSubmit{contract: con}

	workers := make([]*Worker, minerCount)
	for i := 0; i < minerCount; i++ {
		w, err := CreateWorker(i+1, submitter, 2, proxy, NewCpuMiner(50e3))
		if err != nil {
			t.Fatal(err)
		}
		workers[i] = w
	}

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
		TokenBalance: big.NewInt(0), MiningStatus: false, Top50Requests: []*big.Int{}, CurrentChallenge: chal}
	client := rpc.NewMockClientWithValues(opts)
	valueKey := fmt.Sprintf("%s%d", db.QueriedValuePrefix, 1)
	pendingKey := fmt.Sprintf("%s-%s", pubKey, db.PendingChallengeKey)
	dispKey := fmt.Sprintf("%s-%s", pubKey, db.DisputeStatusKey)
	deleteFromDB(DB, []string{
		db.CurrentChallengeKey,
		db.RequestIdKey,
		db.DifficultyKey,
		valueKey,
		pendingKey,
		dispKey,
	})
	writeToDB(DB, []string{dispKey}, [][]byte{
		[]byte(hexutil.EncodeBig(big.NewInt(1))),
	})

	ctx := context.WithValue(context.Background(), tellorCommon.DBContextKey, DB)
	ctx = context.WithValue(ctx, tellorCommon.DataProxyKey, proxy)
	masterInstance := ctx.Value(tellorCommon.MasterContractContextKey)
	if masterInstance == nil {
		contractAddress := common.HexToAddress(cfg.ContractAddress)
		masterInstance, err = tellor.NewTellorMaster(contractAddress, client)
		if err != nil {
			return
		}
		ctx = context.WithValue(ctx, tellorCommon.MasterContractContextKey, masterInstance)
	}
	for _, w := range workers {
		w.Start(ctx)
	}

	//mining check starts immediately, just let it get setup and check
	time.Sleep(300 * time.Millisecond)
	for _, w := range workers {
		if w.loop.mining {
			t.Fatal("Should not be mining yet")
		}
	}

	//should process a saved challenge
	challenge := "someLongNonceString"
	writeToDB(DB, []string{
		db.CurrentChallengeKey,
		db.RequestIdKey,
		db.DifficultyKey,
	},

		[][]byte{
			[]byte(challenge),
			[]byte(hexutil.EncodeBig(big.NewInt(1))),
			[]byte(hexutil.EncodeBig(big.NewInt(5))),
		},
	)

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
	writeToDB(DB,
		[]string{
			db.CurrentChallengeKey,
			db.DifficultyKey,
		},

		[][]byte{
			[]byte("anotherChallenge"),
			[]byte(hexutil.EncodeBig(big.NewInt(1000000000))),
		},
	)

	time.Sleep(2500 * time.Millisecond)
	for _, w := range workers {
		if !w.loop.mining {
			t.Fatal("Should have started mining new challenge")
		}
	}

	//now we should be able to interrupt mining with new challenge and end
	//within time cycle
	writeToDB(DB, []string{
		db.CurrentChallengeKey,
		db.DifficultyKey,
	},
		[][]byte{
			[]byte("simpleChallenge"),
			[]byte(hexutil.EncodeBig(big.NewInt(5))),
		},
	)

	time.Sleep(3 * time.Second)
	if nonce == "" {
		t.Fatal("Should have mined another solution with new challenge request")
	}
	for _, w := range workers {
		w.Stop(ctx)
	}

	time.Sleep(time.Second)
	for _, w := range workers {
		t.Logf("Checking if worker %d is still mining...\n", w.loop.id)
		if w.loop.mining || w.CanMine() {
			t.Fatalf("Miner did not shutdown gracefully after kill sig. Mining: %v CanMine: %v", w.loop.mining, w.CanMine())
		} else {
			t.Logf("Worker %d is finished\n", w.loop.id)
		}

	}
	t.Log("Finished running test")
}
