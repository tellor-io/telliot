// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
)

const BalanceTrackerName = "BalanceTracker"

type BalanceTracker struct {
	db      db.DataServerProxy
	client  contracts.ETHClient
	account *rpc.Account
	logger  log.Logger
}

func (b *BalanceTracker) String() string {
	return BalanceTrackerName
}

func NewBalanceTracker(logger log.Logger, db db.DataServerProxy, client contracts.ETHClient, account *rpc.Account) *BalanceTracker {
	return &BalanceTracker{
		db:      db,
		client:  client,
		account: account,
		logger:  log.With(logger, "component", "balance tracker"),
	}
}

func (b *BalanceTracker) Exec(ctx context.Context) error {

	_fromAddress := b.account.Address

	balance, err := b.client.BalanceAt(ctx, _fromAddress, nil)

	if err != nil {
		return errors.Wrap(err, "getting balance")
	}
	level.Info(b.logger).Log("msg", "got balance", "balance", fmt.Sprintf("%.2e", float64(balance.Int64())))

	enc := hexutil.EncodeBig(balance)
	return b.db.Put(db.BalanceKey, []byte(enc))
}
