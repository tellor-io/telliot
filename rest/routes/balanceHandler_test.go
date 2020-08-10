package routes

import (
	"context"
	"encoding/json"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
)

type BalResult struct {
	Balance string `json:"balance"`
}

func TestBalanceHandler(t *testing.T) {
	h := &BalanceHandler{}

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_balance_rest"))
	if err != nil {
		t.Fatal(err)
	}

	bigBal := big.NewInt(350000)
	bal := hexutil.EncodeBig(bigBal)
	DB.Put(db.BalanceKey, []byte(bal))

	ctx := context.WithValue(context.Background(), common.DBContextKey, DB)
	code, payload := h.Incoming(ctx, nil)
	t.Logf("JSON payload: %s\n", payload)
	if code != 200 {
		if !strings.Contains(payload, "error") {
			t.Fatal("Expected non-200 code to contain error message")
		}
	}

	var res BalResult
	err = json.Unmarshal([]byte(payload), &res)
	if err != nil {
		t.Fatal(err)
	}
	resBal, err := hexutil.DecodeBig(res.Balance)
	if err != nil {
		t.Fatal(err)
	}

	if resBal.Cmp(bigBal) != 0 {
		t.Fatal("Starting and result balances did not match")
	} else {
		t.Logf("Ending balance: %s\n", res.Balance)
	}

}
