package tracker

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
)

func TestRunner(t *testing.T) {
	exitCh := make(chan int)
	client := rpc.NewMockClient()
	db, err := db.Open(filepath.Join(os.TempDir(), "test_leveldb"))
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	runner, _ := NewRunner(client, db)

	ctx := context.Background()
	runner.Start(ctx, exitCh)
	time.Sleep(2 * time.Second)
	exitCh <- 1
	time.Sleep(1 * time.Second)
}
