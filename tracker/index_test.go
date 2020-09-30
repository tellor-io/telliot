package tracker

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
)

//./runTest.sh TestPSR tracker

func TestPSR(t *testing.T) {
	//client.Timeout = 5
	db, err := db.Open(filepath.Join(os.TempDir(), "test_psrFetch"))
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.DBContextKey, db)
	psr, err := BuildIndexTrackers()
	if err != nil {
		t.Fatal(err)
	}
	for idx := range psr {
		err = psr[idx].Exec(ctx)
		psrStr := psr[idx].String()
		if err != nil {
			t.Fatalf("failed to execute psr: %s %v", psrStr, err)
		}
	}
	val, err := db.Get(fmt.Sprintf("qv_%d", 1))
	if err != nil {
		t.Fatal(err)
	}
	if val == nil {
		t.Fatal(fmt.Errorf("Expected a value stored for request ID 1"))
	}
	intVal, err := hexutil.DecodeBig(string(val))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("DB value", intVal)
}
