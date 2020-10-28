// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/pkg/errors"
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/contracts/getter"
	"github.com/tellor-io/TellorMiner/pkg/contracts/tellor"
	"github.com/tellor-io/TellorMiner/pkg/db"
)

type CurrentVariablesTracker struct {
	logger log.Logger
}

func (b *CurrentVariablesTracker) String() string {
	return "CurrentVariablesTracker"
}

func NewCurrentVariablesTracker(logger log.Logger) *CurrentVariablesTracker {
	return &CurrentVariablesTracker{
		logger: log.With(logger, "component", "new current variables"),
	}
}

func (b *CurrentVariablesTracker) Exec(ctx context.Context) error {
	//cast client using type assertion since context holds generic interface{}
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)
	//get the single config instance
	cfg := config.GetConfig()

	//get address from config
	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)

	instanceTellor := ctx.Value(tellorCommon.ContractsTellorContextKey).(*tellor.Tellor)
	returnNewVariables, err := instanceTellor.GetNewCurrentVariables(nil)
	if err != nil {
		level.Warn(b.logger).Log("msg", "new current variables retrieval - contract might not be upgraded", "err", err)
		return nil
	}
	if returnNewVariables.RequestIds[0].Int64() > int64(100) || returnNewVariables.RequestIds[0].Int64() == 0 {
		level.Warn(b.logger).Log("msg", "new current variables request ID not correct - contract about to be upgraded")
		return nil
	}
	fmt.Println(returnNewVariables)

	//if we've mined it, don't save it

	instanceGetter := ctx.Value(tellorCommon.ContractsGetterContextKey).(*getter.TellorGetters)
	myStatus, err := instanceGetter.DidMine(nil, returnNewVariables.Challenge, fromAddress)
	if err != nil {
		return errors.Wrap(err, "status retrieval")
	}
	bitSetVar := []byte{0}
	if myStatus {
		bitSetVar = []byte{1}
	}

	hash := solsha3.SoliditySHA3(
		// types
		[]string{"string"},
		// values
		[]interface{}{
			"timeOfLastNewValue",
		},
	)
	var ret [32]byte
	copy(ret[:], hash)
	timeOfLastNewValue, err := instanceGetter.GetUintVar(nil, ret)
	if err != nil {
		return errors.Wrap(err, "time of last new value retrieval")
	}
	err = DB.Put(db.LastNewValueKey, []byte(hexutil.EncodeBig(timeOfLastNewValue)))
	if err != nil {
		return errors.Wrap(err, "ast new value put")
	}
	err = DB.Put(db.CurrentChallengeKey, returnNewVariables.Challenge[:])
	if err != nil {
		return errors.Wrap(err, "current variables put")
	}

	for i := 0; i < 5; i++ {
		conc := fmt.Sprintf("%s%d", "current_requestId", i)
		err = DB.Put(conc, []byte(hexutil.EncodeBig(returnNewVariables.RequestIds[i])))
		if err != nil {
			return errors.Wrap(err, "request Ids put")
		}
	}

	err = DB.Put(db.DifficultyKey, []byte(hexutil.EncodeBig(returnNewVariables.Difficutly)))
	if err != nil {
		return errors.Wrap(err, "difficulty put")
	}

	err = DB.Put(db.TotalTipKey, []byte(hexutil.EncodeBig(returnNewVariables.Tip)))
	if err != nil {
		return errors.Wrap(err, "totaltip put")
	}

	return DB.Put(db.MiningStatusKey, bitSetVar)
}
