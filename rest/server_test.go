package rest

import (
	"context"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
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
	balInt := big.NewInt(350000)
	DB.Put(db.BalanceKey, []byte(hexutil.EncodeBig(balInt)))
	proxy, err := db.OpenRemoteDB(DB)
	if err != nil {
		t.Fatal(err)
	}
	cfg := config.GetConfig()
	ctx := context.WithValue(context.Background(), common.DBContextKey, DB)
	ctx = context.WithValue(ctx, common.DataProxyKey, proxy)
	srv, err := Create(ctx, cfg.ServerHost, cfg.ServerPort)
	if err != nil {
		t.Fatal(err)
	}

	srv.Start()
	defer srv.Stop()

	data, err := proxy.Get(db.BalanceKey)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("Expected data to be returned")
	}
	asInt, err := hexutil.DecodeBig(string(data))
	if err != nil {
		t.Fatal(err)
	}
	if asInt.Cmp(balInt) != 0 {
		t.Fatalf("Expected %v but received %v as balance", balInt, asInt)
	}

	t.Logf("Retrieved balance from server: %+v\n", asInt)

	/***
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
	**/

}
