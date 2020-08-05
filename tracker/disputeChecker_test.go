package tracker

import (
	"context"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
)

func TestDisputeChecker(t *testing.T) {
	opts := &rpc.MockOptions{ETHBalance: big.NewInt(300000), Nonce: 1, GasPrice: big.NewInt(7000000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}}
	disputeChecker := &disputeChecker{ lastCheckedBlock: 500}
	DB, err := db.Open(filepath.Join(os.TempDir(), "ethGas_test"))
	if err != nil {
		t.Fatal(err)
	}
	client := rpc.NewMockClientWithValues(opts)
	ctx := context.WithValue(context.Background(), tellorCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, tellorCommon.DBContextKey, DB)
	ctx = context.WithValue(ctx, tellorCommon.ContractAddress, common.Address{0x0000000000000000000000000000000000000000})
	err = disputeChecker.Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}
	DB.Close()
}
