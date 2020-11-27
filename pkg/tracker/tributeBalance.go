// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	tellorCommon "github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts/getter"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
)

type TributeTracker struct {
	logger log.Logger
}

func (b *TributeTracker) String() string {
	return "TributeTracker"
}

func NewTributeTracker(logger log.Logger) *TributeTracker {
	return &TributeTracker{
		logger: log.With(logger, "component", "tribure tracker"),
	}
}

func (b *TributeTracker) Exec(ctx context.Context) error {
	//cast client using type assertion since context holds generic interface{}
	client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)

	//get the single config instance
	cfg := config.GetConfig()

	//get address from config
	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)

	_conAddress := cfg.ContractAddress

	//convert to address
	contractAddress := common.HexToAddress(_conAddress)

	instance, err := getter.NewTellorGetters(contractAddress, client)
	if err != nil {
		return errors.Wrap(err, "creating instance")
	}

	balance, err := instance.BalanceOf(nil, fromAddress)
	balanceInTributes, _ := big.NewFloat(1).SetString(balance.String())
	// this _should_ be unreachable given that there is an erro flag for
	// the balanceOf call
	//if !ok {
	//	fmt.Println("Problem converting tributes.")
	//	balanceInTributes = big.NewFloat(0)
	//}
	decimals, _ := big.NewFloat(1).SetString("1000000000000000000")
	// This is unreachable since it's hardcoded
	//if !ok {
	//	fmt.Println("Could not create tribute float for computing tributes")
	//	balanceInTributes = big.NewFloat(0)
	//}
	if decimals != nil {
		balanceInTributes = balanceInTributes.Quo(balanceInTributes, decimals)
	}

	level.Debug(b.logger).Log("msg", "tribute balance", "raw", balance, "trb", balanceInTributes)
	if err != nil {
		return errors.Wrap(err, "retrieving balance")
	}
	enc := hexutil.EncodeBig(balance)
	return DB.Put(db.TributeBalanceKey, []byte(enc))
}
