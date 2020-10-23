// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/contracts/getter"
	"github.com/tellor-io/TellorMiner/pkg/contracts/tellor"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/rpc"
)

type NewCurrentVariablesTracker struct {
}

func (b *NewCurrentVariablesTracker) String() string {
	return "NewCurrentVariablesTracker"
}

func (b *NewCurrentVariablesTracker) Exec(ctx context.Context) error {
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
		fmt.Println("New Current Variables Retrieval Error - Contract might not be upgraded")
		return nil
	}
	if returnNewVariables.RequestIds[0].Int64() > int64(100) || returnNewVariables.RequestIds[0].Int64() == 0 {
		fmt.Println("New Current Variables Request ID not correct - Contract about to be upgraded")
		return nil
	}
	fmt.Println(returnNewVariables)

	//if we've mined it, don't save it

	instanceGetter := ctx.Value(tellorCommon.ContractsGetterContextKey).(*getter.TellorGetters)
	myStatus, err := instanceGetter.DidMine(nil, returnNewVariables.Challenge, fromAddress)
	if err != nil {
		fmt.Println("My Status Retrieval Error")
		return err
	}
	bitSetVar := []byte{0}
	if myStatus {
		bitSetVar = []byte{1}
	}

	timeOfLastNewValue, err := instanceGetter.GetUintVar(nil, rpc.Keccak256("timeOfLastNewValue"))
	if err != nil {
		fmt.Println("Time of Last New Value Retrieval Error")
		return err
	}
	err = DB.Put(db.LastNewValueKey, []byte(hexutil.EncodeBig(timeOfLastNewValue)))
	if err != nil {
		fmt.Println("New Current Variables Put Error")
		return err
	}
	err = DB.Put(db.CurrentChallengeKey, returnNewVariables.Challenge[:])
	if err != nil {
		fmt.Println("New Current Variables Put Error")
		return err
	}

	for i := 0; i < 5; i++ {
		conc := fmt.Sprintf("%s%d", "current_requestId", i)
		err = DB.Put(conc, []byte(hexutil.EncodeBig(returnNewVariables.RequestIds[i])))
		if err != nil {
			fmt.Println("New Current Variables Put Error")
			return err
		}
	}

	err = DB.Put(db.DifficultyKey, []byte(hexutil.EncodeBig(returnNewVariables.Difficutly)))
	if err != nil {
		fmt.Println("New Current Variables Put Error")
		return err
	}

	err = DB.Put(db.TotalTipKey, []byte(hexutil.EncodeBig(returnNewVariables.Tip)))
	if err != nil {
		fmt.Println("New Current Variables Put Error")
		return err
	}

	return DB.Put(db.MiningStatusKey, bitSetVar)
}
