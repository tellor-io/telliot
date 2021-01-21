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

type CurrentVariablesTracker struct {
	db       db.DataServerProxy
	contract *contracts.Tellor
	account  *rpc.Account
	logger   log.Logger
}

func (b *CurrentVariablesTracker) String() string {
	return "CurrentVariablesTracker"
}

func NewCurrentVariablesTracker(logger log.Logger, db db.DataServerProxy, contract *contracts.Tellor, account *rpc.Account) *CurrentVariablesTracker {
	return &CurrentVariablesTracker{
		db:       db,
		contract: contract,
		account:  account,
		logger:   log.With(logger, "component", "CurrentVariablesTracker"),
	}
}

func (b *CurrentVariablesTracker) Exec(ctx context.Context) error {
	returnNewVariables, err := b.contract.Caller.GetNewCurrentVariables(nil)
	if err != nil {
		level.Warn(b.logger).Log("msg", "new current variables retrieval - contract might not be upgraded", "err", err)
		return nil
	}
	if returnNewVariables.RequestIds[0].Int64() > int64(100) || returnNewVariables.RequestIds[0].Int64() == 0 {
		level.Warn(b.logger).Log("msg", "new current variables request ID not correct - contract about to be upgraded")
		return nil
	}

	// If it has been mined, don't save it.
	myStatus, err := b.contract.Getter.DidMine(nil, returnNewVariables.Challenge, b.account.Address)
	if err != nil {
		return errors.Wrap(err, "status retrieval")
	}
	bitSetVar := []byte{0}
	if myStatus {
		bitSetVar = []byte{1}
	}

	timeOfLastNewValue, err := b.contract.Getter.GetUintVar(nil, rpc.Keccak256([]byte("timeOfLastNewValue")))
	if err != nil {
		return errors.Wrap(err, "time of last new value retrieval")
	}
	err = b.db.Put(db.LastNewValueKey, []byte(hexutil.EncodeBig(timeOfLastNewValue)))
	if err != nil {
		return errors.Wrap(err, "ast new value put")
	}
	err = b.db.Put(db.CurrentChallengeKey, returnNewVariables.Challenge[:])
	if err != nil {
		return errors.Wrap(err, "current variables put")
	}

	for i := 0; i < 5; i++ {
		conc := fmt.Sprintf("%s%d", "current_requestId", i)
		err = b.db.Put(conc, []byte(hexutil.EncodeBig(returnNewVariables.RequestIds[i])))
		if err != nil {
			return errors.Wrap(err, "request Ids put")
		}
	}

	err = b.db.Put(db.DifficultyKey, []byte(hexutil.EncodeBig(returnNewVariables.Difficutly)))
	if err != nil {
		return errors.Wrap(err, "difficulty put")
	}

	err = b.db.Put(db.TotalTipKey, []byte(hexutil.EncodeBig(returnNewVariables.Tip)))
	if err != nil {
		return errors.Wrap(err, "total tip put")
	}

	return b.db.Put(db.MiningStatusKey, bitSetVar)
}
