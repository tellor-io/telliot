package tracker

import (
	"context"
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"github.com/tellor-io/TellorMiner/common"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
)

func TestETHGasStation(t *testing.T) {
	tracker := &GasTracker{}
	opts := &rpc.MockOptions{ETHBalance: big.NewInt(300000), Nonce: 1, GasPrice: big.NewInt(7000000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}}
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "ethGas_test"))
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.WithValue(context.Background(), tellorCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	err = tracker.Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}
	v, err := DB.Get(db.GasKey)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Gas Price stored: %v\n", string(v))

}
