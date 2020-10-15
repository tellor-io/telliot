// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package routes

import (
	"encoding/json"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/testutil"
)

type BalResult struct {
	Balance string `json:"balance"`
}

func TestBalanceHandler(t *testing.T) {
	config.OpenTestConfig(t)
	ctx, _, cleanup := testutil.CreateContext(t)
	defer t.Cleanup(cleanup)
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)

	h := &BalanceHandler{}

	bigBal := big.NewInt(350000)
	bal := hexutil.EncodeBig(bigBal)
	if err := DB.Put(db.BalanceKey, []byte(bal)); err != nil {
		t.Fatal(err)
	}

	code, payload := h.Incoming(ctx, nil)
	t.Logf("JSON payload: %s\n", payload)
	if code != 200 {
		if !strings.Contains(payload, "error") {
			t.Fatal("Expected non-200 code to contain error message")
		}
	}

	var res BalResult
	if err := json.Unmarshal([]byte(payload), &res); err != nil {
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
