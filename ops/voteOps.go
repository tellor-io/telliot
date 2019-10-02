package ops

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	tellor1 "github.com/tellor-io/TellorMiner/contracts1"
	"github.com/tellor-io/TellorMiner/rpc"
)

/**
 * This is the operational deposit component. Its purpose is to deposit Tellor Tokens so you can mine
 */

func Vote(_disputeId string,_supportsDispute bool,ctx context.Context) error {
	cfg, err := config.GetConfig()
	if err != nil {
		depositLog.Error("Problem getting config: %+v", err)
		return err
	}
	client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)

	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		depositLog.Error("Problem getting private key: %+v", err)
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		depositLog.Error("Problem extract public key")
		return fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.NonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		return err
	}

	cost := new(big.Int)
	cost.Mul(gasPrice, big.NewInt(300000))
	if balance.Cmp(cost) < 0 {
		//FIXME: notify someone that we're out of funds!
		fmt.Println("Insufficient funds")
		return fmt.Errorf("Insufficient funds to send transaction: %v < %v", balance, cost)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	instance2 := ctx.Value(tellorCommon.TransactorContractContextKey).(*tellor1.TellorTransactor)
	dis := new(big.Int)
	dis,_ = dis.SetString(_disputeId, 10)

	tx, err := instance2.Vote(auth,dis,_supportsDispute)
	if err != nil {
		depositLog.Error("Could not deposit stake: %+v", err)
		return err
	}

	fmt.Println("Vote tx sent: ", tx.Hash().Hex())
	return nil
}
