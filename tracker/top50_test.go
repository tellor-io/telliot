package tracker

import (
	"context"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
)

func TestTop50(t *testing.T) {
	startBal := big.NewInt(356000)
	/*
		opts := &rpc.MockOptions{ETHBalance: big.NewInt(356000), Nonce: 1, GasPrice: big.NewInt(700000000), TokenBalance: startBal}
		client := rpc.NewMockClientWithValues(opts)
	*/

	/*
		cfg, err := config.GetConfig()
		if err != nil {
			t.Fatal(err)
		}
		url := cfg.NodeURL
		client, err := rpc.NewClient(url)
		if err != nil {
			t.Fatal(err)
		}
	*/
	top50 := make([]*big.Int, 51)
	for i := range top50 {
		top50[i] = big.NewInt(int64(i))
	}

	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0), Top50Requests: top50}
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_balance"))
	if err != nil {
		t.Fatal(err)
	}
	tracker := &Top50Tracker{}
	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	err = tracker.Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}
	v, err := DB.Get(db.Top50Key)
	if err != nil {
		t.Fatal(err)
	}

	bigs := strings.Split(string(v), ",")
	if len(bigs) == 0 {
		t.Fatal("Expected CSV list of top50 ids")
	}
}
