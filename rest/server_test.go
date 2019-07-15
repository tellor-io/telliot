package rest

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
)

type BalTest struct {
	Balance string `json:"balance"`
}

func TestServer(t *testing.T) {
	DB, err := db.Open(filepath.Join(os.TempDir(), "test_server"))
	if err != nil {
		t.Fatal(err)
	}
	DB.Put(db.BalanceKey, []byte(hexutil.EncodeBig(big.NewInt(350000))))

	ctx := context.WithValue(context.Background(), common.DBContextKey, DB)
	srv, err := Create(ctx, "localhost", 5000)
	if err != nil {
		t.Fatal(err)
	}
	srv.Start()
	defer srv.Stop()

	resp, err := http.Get("http://localhost:5000/balance")
	if err != nil {
		t.Fatal(err)
	}
	var bal BalTest
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&bal)
	if !strings.Contains(bal.Balance, "0x") {
		t.Fatal("Missing balance in response")
	} else {
		t.Logf("Retrieved balance from server: %+v\n", bal)
	}

}
