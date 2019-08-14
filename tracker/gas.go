package tracker

import (
	"context"
	"log"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
)

type GasTracker struct {
}

func (b *GasTracker) String() string {
	return "GasTracker"
}

func (b *GasTracker) Exec(ctx context.Context) error {
	client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
		return err
	}

	enc := hexutil.EncodeBig(gasPrice)
	log.Printf("GasKey: %v", enc)
	return DB.Put(db.GasKey, []byte(enc))
}
