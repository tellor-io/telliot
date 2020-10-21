// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/rpc"
)

const BalanceTrackerName = "BalanceTracker"

type BalanceTracker struct {
}

func (b *BalanceTracker) String() string {
	return BalanceTrackerName
}

func (b *BalanceTracker) Exec(ctx context.Context, logger log.Logger) error {

	// cast client using type assertion since context holds generic interface{}.
	client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)

	// get the single config instance.
	cfg := config.GetConfig()

	// get address from config.
	_fromAddress := cfg.PublicAddress

	// convert to address.
	fromAddress := common.HexToAddress(_fromAddress)

	balance, err := client.BalanceAt(ctx, fromAddress, nil)

	if err != nil {
		level.Error(logger).Log("msg", "error getting balance", "err", err)
		return err
	}
	enc := hexutil.EncodeBig(balance)

	// log.Printf("Balance: %v", enc)
	level.Info(logger).Log("msg", "Got balance", "balance", enc)
	return DB.Put(db.BalanceKey, []byte(enc))
}
