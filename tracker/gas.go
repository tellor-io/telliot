package tracker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
)

// GWEI constant is the multiplier from Wei
const GWEI = 1000000000

// GasTracker is the struct that maintains the latest gasprices.
// note the prices are actually stored in the DB
type GasTracker struct {
}

//GasPriceModel is what ETHGasStation returns from queries. Not all fields are filled in
type GasPriceModel struct {
	Fast    float32 `json:"fast"`
	Fastest float32 `json:"fastest"`
	Average float32 `json:"average"`
}

func (b *GasTracker) String() string {
	return "GasTracker"
}

// Exec queries EthGasStation for the current Gas Prices
func (b *GasTracker) Exec(ctx context.Context) error {
	client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)

	netID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err)
		return err
	}

	var gasPrice *big.Int

	if big.NewInt(1).Cmp(netID) == 0 {
		url := "https://ethgasstation.info/json/ethgasAPI.json"
		req := &FetchRequest{queryURL: url, timeout: time.Duration(15 * time.Second)}
		payload, err := fetchWithRetries(req)
		if err != nil {
			gasPrice, err = client.SuggestGasPrice(context.Background())
		} else {
			gpModel := GasPriceModel{}
			err = json.Unmarshal(payload, &gpModel)
			if err != nil {
				log.Printf("Problem with ETH gas station json: %v\n", err)
				gasPrice, err = client.SuggestGasPrice(context.Background())
			} else {
				gasPrice = big.NewInt(int64(gpModel.Fast / 10))
				gasPrice = gasPrice.Mul(gasPrice, big.NewInt(GWEI))
				log.Println("Using ETHGasStation fast price: ", gasPrice)
			}
		}
	} else {
		gasPrice, err = client.SuggestGasPrice(context.Background())
	}

	enc := hexutil.EncodeBig(gasPrice)
	return DB.Put(db.GasKey, []byte(enc))
}
