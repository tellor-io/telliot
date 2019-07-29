package tracker

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
)

func TestFetchData(t *testing.T) {
	startBal := big.NewInt(456000)

	top50 := make([]*big.Int, 51)
	for i := range top50 {
		top50[i] = big.NewInt(int64(i))
	}

	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0), Top50Requests: top50}
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_fetchData"))
	if err != nil {
		t.Fatal(err)
	}
	tracker1 := &Top50Tracker{}
	ctx1 := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx1 = context.WithValue(ctx1, common.DBContextKey, DB)
	err = tracker1.Exec(ctx1)
	if err != nil {
		t.Fatal(err)
	}

	tracker := &RequestDataTracker{}
	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	err = tracker.Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}
	v, err := DB.Get(fmt.Sprint(1))
	if err != nil {
		t.Fatal(err)
	}
	b, err := hexutil.DecodeBig(string(v))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(b)
	t.Logf("Data stored: %v\n", string(v))
	if b.Cmp(big.NewInt(1)) != 1 {
		t.Fatalf("Data for each ID from client did not match what should have been stored in DB. %s != %s", b, startBal)
	}
}
