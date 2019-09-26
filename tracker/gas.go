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
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
)

const GWEI = 1000000000

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

func (b *GasTracker) Exec(ctx context.Context) error {
	client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)

	/*
		type FetchRequest struct {
			queryURL string
			timeout  time.Duration
		}
	*/
	url := "https://ethgasstation.info/json/ethgasAPI.json"
	req := &FetchRequest{queryURL: url, timeout: time.Duration(15 * time.Second)}
	payload, err := fetchWithRetries(req)
	var gasPrice *big.Int
	if err != nil {
		gasPrice, err = client.SuggestGasPrice(context.Background())
	} else {
		gpModel := GasPriceModel{}
		err = json.Unmarshal(payload, &gpModel)
		if err != nil {
			log.Printf("Problem with ETH gas station json: %v\n", err)
			gasPrice, err = client.SuggestGasPrice(context.Background())
		} else {
			gasPrice = big.NewInt(int64(gpModel.Average / 10))
			gasPrice = gasPrice.Mul(gasPrice, big.NewInt(GWEI))
			log.Println("Using ETHGasStation average price: ", gasPrice)
		}
	}

	//gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
		return err
	}

	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
		return err
	}

	mult := cfg.GasMultiplier
	max := cfg.GasMax

	if mult > 0 {
		gweiBase := big.NewFloat(float64(GWEI))
		multFloat := big.NewFloat(float64(mult))
		maxFloat := big.NewFloat(float64(max))
		maxFloat.Mul(maxFloat, gweiBase)

		gasFloat := new(big.Float).SetInt(gasPrice)
		origGwei := new(big.Float).Quo(gasFloat, gweiBase)
		gasFloat.Mul(gasFloat, multFloat)
		newGwei := new(big.Float).Quo(gasFloat, gweiBase)
		if gasFloat.Cmp(maxFloat) > 0 {
			log.Printf("Applying gas multiplier would exceed max gas limit (%v > %v max). Using original gas price instead: %v\n", newGwei, max, origGwei)
		} else {
			gp, _ := gasFloat.Uint64()
			gasPrice = big.NewInt(int64(gp))
			log.Printf("Using gas multiplier %v to increase gas price from suggested %v gwei to %v gwei (%v wei)\n", mult, origGwei, newGwei, gasPrice)
		}
	}

	enc := hexutil.EncodeBig(gasPrice)
	log.Printf("GasKey: %v", enc)
	return DB.Put(db.GasKey, []byte(enc))
}
