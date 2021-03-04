// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
)

type TributeTracker struct {
	db       db.DataServerProxy
	contract *contracts.ITellor
	account  *rpc.Account
	logger   log.Logger
}

func (b *TributeTracker) String() string {
	return "TributeTracker"
}

func NewTributeTrackers(logger log.Logger, db db.DataServerProxy, contract *contracts.ITellor, accounts []*rpc.Account) []Tracker {
	trackers := make([]Tracker, len(accounts))
	for i, account := range accounts {
		trackers[i] = &TributeTracker{
			db:       db,
			contract: contract,
			account:  account,
			logger:   log.With(logger, "component", "tribute tracker"),
		}
	}
	return trackers
}

func NewTributeTracker(logger log.Logger, db db.DataServerProxy, contract *contracts.ITellor, account *rpc.Account) *TributeTracker {
	return &TributeTracker{
		db:       db,
		contract: contract,
		account:  account,
		logger:   log.With(logger, "component", ComponentName),
	}
}

func (b *TributeTracker) Exec(ctx context.Context) error {
	balance, err := b.contract.BalanceOf(nil, b.account.Address)
	if err != nil {
		return errors.Wrap(err, "retrieving balance")
	}
	balanceH, _ := big.NewFloat(1).SetString(balance.String())

	decimals, _ := big.NewFloat(1).SetString("1000000000000000000")

	if decimals != nil {
		balanceH = balanceH.Quo(balanceH, decimals)
	}

	level.Info(b.logger).Log("msg", "TRB balance", "amount", balanceH)

	enc := hexutil.EncodeBig(balance)
	return b.db.Put(db.TributeBalanceKeyFor(b.account.Address), []byte(enc))
}
