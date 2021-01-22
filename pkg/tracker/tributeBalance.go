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
	contract *contracts.Tellor
	account  *rpc.Account
	logger   log.Logger
}

func (b *TributeTracker) String() string {
	return "TributeTracker"
}

func NewTributeTracker(logger log.Logger, db db.DataServerProxy, contract *contracts.Tellor, account *rpc.Account) *TributeTracker {
	return &TributeTracker{
		db:       db,
		contract: contract,
		account:  account,
		logger:   log.With(logger, "component", "tribute tracker"),
	}
}

func (b *TributeTracker) Exec(ctx context.Context) error {
	balance, err := b.contract.Getter.BalanceOf(nil, b.account.Address)
	balanceInTributes, _ := big.NewFloat(1).SetString(balance.String())

	decimals, _ := big.NewFloat(1).SetString("1000000000000000000")

	if decimals != nil {
		balanceInTributes = balanceInTributes.Quo(balanceInTributes, decimals)
	}

	level.Debug(b.logger).Log("msg", "tribute balance", "raw", balance, "trb", balanceInTributes)
	if err != nil {
		return errors.Wrap(err, "retrieving balance")
	}
	enc := hexutil.EncodeBig(balance)
	return b.db.Put(db.TributeBalanceKey, []byte(enc))
}
