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
	db, err := db.Open(filepath.Join(os.TempDir(), "test_psrFetch"))
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.DBContextKey, db)
	psr, err := BuildPSRTrackers()
	if err != nil {
		t.Fatal(err)
	}
	err = psr.Exec(ctx)
	if err != nil {
		t.Fatalf("failed to execute psr: %v", err)
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
