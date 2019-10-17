package pow

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/db"
)

type TestHarness struct {
	tasker          *miningTasker
	issuedChallenge *miningChallenge
	cancelled       bool
	DB              db.DB
	proxy           db.DataServerProxy
	taskCh          taskChannel
	cancelCh        cancelChannel
}

func (th *TestHarness) init(t *testing.T, cb func(DB db.DB)) {
	DB, err := db.Open("/tmp/miningTaskerTest")
	if err != nil {
		t.Fatal(err)
	}
	localProxy, err := db.OpenLocalProxy(DB)
	if err != nil {
		t.Fatal(err)
	}
	cb(DB)

	th.taskCh = make(taskChannel)
	go func() {
		th.issuedChallenge = <-th.taskCh
	}()

	th.cancelCh = make(cancelChannel)
	go func() {
		th.cancelled = <-th.cancelCh
	}()

	tasker, err := createTasker(1, th.taskCh, th.cancelCh, 1, localProxy)

	th.DB = DB
	th.proxy = localProxy
	th.tasker = tasker
}

func TestMiningTaskerNoChallenge(t *testing.T) {
	tester := &TestHarness{}
	tester.init(t, func(DB db.DB) {
		deleteFromDB(DB, []string{db.CurrentChallengeKey, db.RequestIdKey, db.DifficultyKey})
	})

	tester.tasker.Start(context.Background())
	time.Sleep(2500 * time.Millisecond)
	if tester.issuedChallenge != nil || tester.cancelled {
		t.Fatal("Should not have done anything without challenge data stored")
	}
	tester.tasker.exitCh <- os.Interrupt
	time.Sleep(10 * time.Millisecond)
	if tester.tasker.running {
		t.Fatal("Tasker should have stopped running")
	}
}

func TestMiningTaskerPendingChallenge(t *testing.T) {

	cfg, err := config.GetConfig()
	if err != nil {
		t.Fatal(err)
	}
	//get address from config
	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)
	pubKey := strings.ToLower(fromAddress.Hex())

	pendingKey := pubKey + "-" + db.PendingChallengeKey
	dispKey := pubKey + "-" + db.DisputeStatusKey
	valKey := fmt.Sprintf("%s-%d", db.QueriedValuePrefix, 1)
	tester := &TestHarness{}
	tester.init(t, func(DB db.DB) {
		writeToDB(DB, []string{
			db.CurrentChallengeKey,
			db.RequestIdKey,
			db.DifficultyKey,
			dispKey,
			pendingKey,
			valKey,
		}, [][]byte{
			[]byte("simpleChallenge"), //challenge
			[]byte(hexutil.EncodeBig(big.NewInt(1))),
			[]byte(hexutil.EncodeBig(big.NewInt(500))),
			[]byte(hexutil.EncodeBig(big.NewInt(1))),
			[]byte("simpleChallenge"), //pending to match current challenge
			[]byte(hexutil.EncodeBig(big.NewInt(10000))),
		})
	})
	tester.tasker.Start(context.Background())
	time.Sleep(2500 * time.Millisecond)
	if tester.issuedChallenge != nil {
		t.Fatal("Did not expect pending txn to issue challenge")
	}
	if !tester.cancelled {
		t.Fatal("Expected to be cancelled when there is a pending txn")
	}
}

func TestMiningTaskerNewChallenge(t *testing.T) {

	cfg, err := config.GetConfig()
	if err != nil {
		t.Fatal(err)
	}
	//get address from config
	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)
	pubKey := strings.ToLower(fromAddress.Hex())

	pendingKey := pubKey + "-" + db.PendingChallengeKey
	dispKey := pubKey + "-" + db.DisputeStatusKey
	valKey := fmt.Sprintf("%s-%d", db.QueriedValuePrefix, 1)
	tester := &TestHarness{}
	tester.init(t, func(DB db.DB) {
		deleteFromDB(DB, []string{pendingKey})
		writeToDB(DB, []string{
			db.CurrentChallengeKey,
			db.RequestIdKey,
			db.DifficultyKey,
			dispKey,
			valKey,
		}, [][]byte{
			[]byte("simpleChallenge"),
			[]byte(hexutil.EncodeBig(big.NewInt(1))),
			[]byte(hexutil.EncodeBig(big.NewInt(500))),
			[]byte(hexutil.EncodeBig(big.NewInt(1))),
			[]byte(hexutil.EncodeBig(big.NewInt(10000))),
		})
	})
	tester.tasker.Start(context.Background())
	time.Sleep(2500 * time.Millisecond)
	if tester.issuedChallenge == nil || tester.cancelled {
		t.Fatal("Expended a new challenge issued and not to be cancelled")
	}
}

func TestMiningTaskerNewThenPendingChallenge(t *testing.T) {

	cfg, err := config.GetConfig()
	if err != nil {
		t.Fatal(err)
	}
	//get address from config
	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)
	pubKey := strings.ToLower(fromAddress.Hex())

	pendingKey := pubKey + "-" + db.PendingChallengeKey
	dispKey := pubKey + "-" + db.DisputeStatusKey
	valKey := fmt.Sprintf("%s-%d", db.QueriedValuePrefix, 1)
	tester := &TestHarness{}
	tester.init(t, func(DB db.DB) {
		deleteFromDB(DB, []string{pendingKey})
		writeToDB(DB, []string{
			db.CurrentChallengeKey,
			db.RequestIdKey,
			db.DifficultyKey,
			dispKey,
			valKey,
		}, [][]byte{
			[]byte("simpleChallenge"),
			[]byte(hexutil.EncodeBig(big.NewInt(1))),
			[]byte(hexutil.EncodeBig(big.NewInt(1000000))),
			[]byte(hexutil.EncodeBig(big.NewInt(1))),
			[]byte(hexutil.EncodeBig(big.NewInt(10000))),
		})
	})
	tester.tasker.Start(context.Background())
	time.Sleep(100 * time.Millisecond)
	if tester.issuedChallenge == nil || tester.cancelled {
		t.Fatal("Expended a new challenge issued and not to be cancelled")
	}
	tester.issuedChallenge = nil
	tester.cancelled = false
	writeToDB(tester.DB, []string{pendingKey}, [][]byte{[]byte("simpleChallenge")})
	time.Sleep(2500 * time.Millisecond)
	if tester.issuedChallenge != nil {
		t.Fatal("Did not expect a new challenge issued")
	}
	if !tester.cancelled {
		t.Fatal("Expected to be cancelled on pending txn")
	}
}
