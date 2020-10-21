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
	"github.com/tellor-io/TellorMiner/pkg/contracts/getter"
	"github.com/tellor-io/TellorMiner/pkg/db"
)

const CurrentVariablesTrackerName = "CurrentVariablesTracker"

type CurrentVariablesTracker struct {
}

func (b *CurrentVariablesTracker) String() string {
	return CurrentVariablesTrackerName
}

func (b *CurrentVariablesTracker) Exec(ctx context.Context, logger log.Logger) error {
	// cast client using type assertion since context holds generic interface{}.
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)
	//get the single config instance
	cfg := config.GetConfig()

	// get address from config.
	_fromAddress := cfg.PublicAddress

	// convert to address.
	fromAddress := common.HexToAddress(_fromAddress)

	instance := ctx.Value(tellorCommon.ContractsGetterContextKey).(*getter.TellorGetters)
	currentChallenge, requestID, difficulty, queryString, granularity, totalTip, err := instance.GetCurrentVariables(nil)
	if err != nil {
		return errors.Wrap(err, "retrieve current variables")
	}

	// if we've mined it, don't save it.
	myStatus, err := instance.DidMine(nil, currentChallenge, fromAddress)
	if err != nil {
		level.Error(logger).Log("msg", "My Status Retrieval Error", "err", err)
		return err
	}
	bitSetVar := []byte{0}
	if myStatus {
		bitSetVar = []byte{1}
	}
	level.Info(logger).Log("msg", "Retrieved variables", "challengeHash", currentChallenge)

	err = DB.Put(db.CurrentChallengeKey, currentChallenge[:])
	if err != nil {
		level.Error(logger).Log("msg", "Current Variables Put Error", "var", "currentChallenge", "err", err)
		return err
	}
	err = DB.Put(db.RequestIdKey, []byte(hexutil.EncodeBig(requestID)))
	if err != nil {
		level.Error(logger).Log("msg", "Current Variables Put Error", "var", "requestId", "err", err)
		return err
	}
	err = DB.Put(db.DifficultyKey, []byte(hexutil.EncodeBig(difficulty)))
	if err != nil {
		level.Error(logger).Log("msg", "Current Variables Put Error", "var", "difficulty", "err", err)
		return err
	}
	err = DB.Put(db.QueryStringKey, []byte(queryString))
	if err != nil {
		level.Error(logger).Log("msg", "Current Variables Put Error", "var", "queryString", "err", err)
		return err
	}
	err = DB.Put(db.GranularityKey, []byte(hexutil.EncodeBig(granularity)))
	if err != nil {
		level.Error(logger).Log("msg", "Current Variables Put Error", "var", "granularity", "err", err)
		return err
	}
	err = DB.Put(db.TotalTipKey, []byte(hexutil.EncodeBig(totalTip)))
	if err != nil {
		level.Error(logger).Log("msg", "Current Variables Put Error", "var", "totaltip", "err", err)
		return err
	}

	return DB.Put(db.MiningStatusKey, bitSetVar)
}
