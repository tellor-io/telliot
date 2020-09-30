package tracker

import (
	"context"
	"fmt"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/contracts2"
	"github.com/tellor-io/TellorMiner/contracts"
	"github.com/tellor-io/TellorMiner/db"
	"github.com/tellor-io/TellorMiner/util"
	// "encoding/hex"
	"github.com/miguelmota/go-solidity-sha3"
)

var newCurrentVarsLog = util.NewLogger("tracker", "NewCurrentVarsTracker")

type returnNewVariables struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficulty *big.Int
	Tip        *big.Int
}
//CurrentVariablesTracker concrete tracker type
type NewCurrentVariablesTracker struct {
}

func (b *NewCurrentVariablesTracker) String() string {
	return "NewCurrentVariablesTracker"
}

//Exec implementation for tracker
func (b *NewCurrentVariablesTracker) Exec(ctx context.Context) error {
	//cast client using type assertion since context holds generic interface{}
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)
	//get the single config instance
	cfg := config.GetConfig()

	//get address from config
	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)

	instance := ctx.Value(tellorCommon.NewTellorContractContextKey).(*contracts2.Tellor)
	returnNewVariables, err := instance.GetNewCurrentVariables(nil)
	if err != nil{
		fmt.Println("New Current Variables Retrieval Error - Contract might not be upgraded")
		return nil
	}
	if returnNewVariables.RequestIds[0].Int64() > int64(100) || returnNewVariables.RequestIds[0].Int64() == 0 {
		fmt.Println("New Current Variables Request ID not correct - Contract about to be upgraded")
		return nil
	}
	fmt.Println(returnNewVariables)

	//if we've mined it, don't save it
	
	instance2 := ctx.Value(tellorCommon.MasterContractContextKey).(*contracts.TellorMaster)
	myStatus, err := instance2.DidMine(nil, returnNewVariables.Challenge, fromAddress)
	if err != nil {
		fmt.Println("My Status Retrieval Error")
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
	timeOfLastNewValue, err := instance2.GetUintVar(nil,ret)
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


	for i:= 0; i < 5; i++ {
		conc := fmt.Sprintf("%s%d","current_requestId",i)
		err = DB.Put(conc, []byte(hexutil.EncodeBig(returnNewVariables.RequestIds[i])))
		if err != nil {
			fmt.Println("New Current Variables Put Error")
			return err
		}
	}

	err = DB.Put(db.DifficultyKey, []byte(hexutil.EncodeBig(returnNewVariables.Difficulty)))
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
