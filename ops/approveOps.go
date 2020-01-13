package ops

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/config"
	tellor1 "github.com/tellor-io/TellorMiner/contracts1"
	"github.com/tellor-io/TellorMiner/rpc"
)

/**
 * This is the operational approve component. Its purpose is to approve someone to spend your tokens
 */

func Approve(_spender string, amount string,ctx context.Context) (error) {
   cfg := config.GetConfig()
   client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)

   privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
   if err != nil {
	   return err
   }

   publicKey := privateKey.Public()
   publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
   if !ok {
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
	   return fmt.Errorf("Insufficient funds to send transaction: %v < %v", balance, cost)
   }

   auth := bind.NewKeyedTransactor(privateKey)
   auth.Nonce = big.NewInt(int64(nonce))
   auth.Value = big.NewInt(0)      // in wei
   auth.GasLimit = uint64(300000) // in units
   auth.GasPrice = gasPrice.Mul(gasPrice,big.NewInt(3))

   toAdd := common.HexToAddress(_spender)
   amt := new(big.Int)
   amt,_ = amt.SetString(amount, 10)

   instance2 := ctx.Value(tellorCommon.TransactorContractContextKey).(*tellor1.TellorTransactor)

   tx, err := instance2.Approve(auth,toAdd,amt)
   if err != nil {
      fmt.Println("ERROR",err)
	   return err
   }

   fmt.Printf("Approve tx sent: %s", tx.Hash().Hex())

   return nil
}
