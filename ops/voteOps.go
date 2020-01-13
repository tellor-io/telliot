package ops

import (
	"context"
	"fmt"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	tellor1 "github.com/tellor-io/TellorMiner/contracts1"
	"math/big"
)

/**
 * This is the operational deposit component. Its purpose is to deposit Tellor Tokens so you can mine
 */

func Vote(_disputeId string,_supportsDispute bool,ctx context.Context) error {

	instance2 := ctx.Value(tellorCommon.TransactorContractContextKey).(*tellor1.TellorTransactor)
	dis := new(big.Int)
	dis,_ = dis.SetString(_disputeId, 10)

	auth, err := PrepareEthTransaction(ctx)
	if err != nil {
		return fmt.Errorf("failed to prepare ethereum transaction: %s", err.Error())
	}
	tx, err := instance2.Vote(auth,dis,_supportsDispute)
	if err != nil {
		return fmt.Errorf("failed to submit vote transaction: %s", err.Error())
	}

	fmt.Printf("Vote submitted with transaction %s\n", tx.Hash().Hex())
	return nil
}
