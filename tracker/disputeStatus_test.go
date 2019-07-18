package tracker

/*
import (
	"context"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
)

func TestDisputeStatus(t *testing.T) {
	startBal := big.NewInt(356000)
	client := rpc.NewMockClientWithValues(startBal, 1, big.NewInt(7000000000))
	DB, err := db.Open(filepath.Join(os.TempDir(), "test_balance"))
	if err != nil {
		t.Fatal(err)
	}
	tracker := &DisputeTracker{}
	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	err = tracker.Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}
	v, err := DB.Get(db.DisputeStatusKey)
	if err != nil {
		t.Fatal(err)
	}
	b, err := hexutil.DecodeBig(string(v))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Dispute Status stored: %v\n", string(v))
	if b.Cmp(big.NewInt(1)) != 0 {
		t.Fatalf("Dispute Status from client did not match what should have been stored in DB. %s != %s", b, "one")
	}
}
*/
