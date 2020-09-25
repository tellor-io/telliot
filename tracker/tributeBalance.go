// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	tellor "github.com/tellor-io/TellorMiner/contracts"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/rpc"
)

type TributeTracker struct {
}

func (b *TributeTracker) String() string {
	return "TributeTracker"
}

func (b *TributeTracker) Exec(ctx context.Context) error {
	// ast client using type assertion since context holds generic interface{}
	client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)

	// et the single config instance
	cfg := config.GetConfig()

	// et address from config
	_fromAddress := cfg.PublicAddress

	// onvert to address
	fromAddress := common.HexToAddress(_fromAddress)

	_conAddress := cfg.ContractAddress

	// onvert to address
	contractAddress := common.HexToAddress(_conAddress)

	instance, err := tellor.NewTellorMaster(contractAddress, client)
	if err != nil {
		fmt.Println("Instance error - TributeBalance")
		return err
	}

	balance, err := instance.BalanceOf(nil, fromAddress)
	balanceInTributes, _ := big.NewFloat(1).SetString(balance.String())
	// this _should_ be unreachable given that there is an erro flag for
	// the balanceOf call
	// f !ok {
	// fmt.Println("Problem converting tributes.")
	// balanceInTributes = big.NewFloat(0)
	//
	decimals, _ := big.NewFloat(1).SetString("1000000000000000000")
	// This is unreachable since it's hardcoded
	// f !ok {
	// fmt.Println("Could not create tribute float for computing tributes")
	// balanceInTributes = big.NewFloat(0)
	//
	if decimals != nil {
		balanceInTributes = balanceInTributes.Quo(balanceInTributes, decimals)
	}

	// umTributes, _ := balanceInTributes.Float64()
	log.Printf("Tribute Balance: %v (%v tributes)\n", balance, balanceInTributes)
	if err != nil {
		fmt.Println("Balance Retrieval Error - Tribute Balance")
		return err
	}
	enc := hexutil.EncodeBig(balance)
	return DB.Put(db.TributeBalanceKey, []byte(enc))
}
