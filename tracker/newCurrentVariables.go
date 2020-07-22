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

	instance := ctx.Value(tellorCommon.NewTransactorContractContextKey).(*contracts2.Tellor)

	
	returnNewVariables, err := instance.GetNewCurrentVariables(nil)
	if err != nil {
		fmt.Println("New Current Variables Retrieval Error")
		return err
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
	currentVarsLog.Info("Retrieved variables. challengeHash: %x", returnNewVariables.Challenge)

	err = DB.Put(db.CurrentChallengeKey, returnNewVariables.Challenge[:])
	if err != nil {
		fmt.Println("New Current Variables Put Error")
		return err
	}

	//check this bad boy
	for i:= 0; i < 5; i++ {
		err = DB.Put(db.RequestIdKey, []byte(hexutil.EncodeBig(returnNewVariables.RequestIds[i])))
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
