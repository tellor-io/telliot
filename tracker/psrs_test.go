package tracker

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
)

func TestMeanAt(t *testing.T) {
	db, err := db.Open(filepath.Join(os.TempDir(), "test_MeanAt"))
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.DBContextKey, db)
	BuildIndexTrackers()
	ethIndexes := indexes["ETH/USD"]
	if err != nil {
		t.Fatal(err)
	}
	execEthUsdPsrs(ctx, t, ethIndexes)

	MeanAt(ethIndexes, clck.Now())
	db.Close()
}
