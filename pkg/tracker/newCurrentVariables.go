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
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/contracts/getter"
	"github.com/tellor-io/TellorMiner/pkg/contracts/tellor"
	"github.com/tellor-io/TellorMiner/pkg/db"
)

type NewCurrentVariablesTracker struct {
}

func (b *NewCurrentVariablesTracker) String() string {
	return "NewCurrentVariablesTracker"
}

func (b *NewCurrentVariablesTracker) Exec(ctx context.Context, logger log.Logger) error {
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
		level.Warn(logger).Log("msg", "New Current Variables Retrieval Error - Contract might not be upgraded", "err", err)
		return nil
	}
	if returnNewVariables.RequestIds[0].Int64() > int64(100) || returnNewVariables.RequestIds[0].Int64() == 0 {
		level.Warn(logger).Log("msg", "New Current Variables Request ID not correct - Contract about to be upgraded")
		return nil
	}
	fmt.Println(returnNewVariables)

	//if we've mined it, don't save it

	instanceGetter := ctx.Value(tellorCommon.ContractsGetterContextKey).(*getter.TellorGetters)
	myStatus, err := instanceGetter.DidMine(nil, returnNewVariables.Challenge, fromAddress)
	if err != nil {
		level.Error(logger).Log("msg", "My Status Retrieval Error", "err", err)
		return err
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
		level.Error(logger).Log("msg", "Time of Last New Value Retrieval Error", "err", err)
		return err
	}
	err = DB.Put(db.LastNewValueKey, []byte(hexutil.EncodeBig(timeOfLastNewValue)))
	if err != nil {
		level.Error(logger).Log("msg", "New Current Variables Put Error", "var", "lastnewValue", "err", err)
		return err
	}
	err = DB.Put(db.CurrentChallengeKey, returnNewVariables.Challenge[:])
	if err != nil {
		level.Error(logger).Log("msg", "New Current Variables Put Error", "var", "currentChallenge", "err", err)
		return err
	}

	for i := 0; i < 5; i++ {
		conc := fmt.Sprintf("%s%d", "current_requestId", i)
		err = DB.Put(conc, []byte(hexutil.EncodeBig(returnNewVariables.RequestIds[i])))
		if err != nil {
			level.Error(logger).Log("msg", "New Current Variables Put Error", "var", "requestIds", "err", err)
			return err
		}
	}

	err = DB.Put(db.DifficultyKey, []byte(hexutil.EncodeBig(returnNewVariables.Difficutly)))
	if err != nil {
		level.Error(logger).Log("msg", "New Current Variables Put Error", "var", "difficulty", "err", err)
		return err
	}

	err = DB.Put(db.TotalTipKey, []byte(hexutil.EncodeBig(returnNewVariables.Tip)))
	if err != nil {
		level.Error(logger).Log("msg", "New Current Variables Put Error", "var", "totaltip", "err", err)
		return err
	}

	return DB.Put(db.MiningStatusKey, bitSetVar)
}
