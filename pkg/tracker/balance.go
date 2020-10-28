// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/rpc"
)

const BalanceTrackerName = "BalanceTracker"

type BalanceTracker struct {
	logger log.Logger
}

func (b *BalanceTracker) String() string {
	return BalanceTrackerName
}

func NewBalanceTracker(logger log.Logger) *BalanceTracker {
	return &BalanceTracker{
		logger: log.With(logger, "component", "balance tracker"),
	}
}

func (b *BalanceTracker) Exec(ctx context.Context) error {

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
		return errors.Wrap(err, "getting balance")
	}
	enc := hexutil.EncodeBig(balance)

	level.Info(b.logger).Log("msg", "got balance", "balance", enc)
	return DB.Put(db.BalanceKey, []byte(enc))
}
