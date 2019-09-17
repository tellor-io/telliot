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
	tellor "github.com/tellor-io/TellorMiner/contracts"
	tellor1 "github.com/tellor-io/TellorMiner/contracts1"
	"github.com/tellor-io/TellorMiner/rpc"
	"github.com/tellor-io/TellorMiner/util"
)

/**
 * This is the operational deposit component. Its purpose is to deposit Tellor Tokens so you can mine
 */

var (
	disputeLog = util.NewLogger("ops", "DisputeOp")
)

func Dispute(_requestId string,_timestamp string,_minerIndex string,ctx context.Context) error {
	cfg, err := config.GetConfig()
	if err != nil {
		disputeLog.Error("Problem getting config: %+v", err)
		return err
	}
	client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)

	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		disputeLog.Error("Problem getting private key: %+v", err)
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		disputeLog.Error("Problem extract public key")
		return fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		disputeLog.Error("Problem getting pending nonce: %+v", err)
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		disputeLog.Error("Problem getting gas price: %+v", err)
		return err
	}

	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		disputeLog.Error("Problem getting balance: %+v", err)
		return err
	}

	cost := new(big.Int)
	cost.Mul(gasPrice, big.NewInt(700000))
	if balance.Cmp(cost) < 0 {
		//FIXME: notify someone that we're out of funds!
		disputeLog.Error("Insufficient funds")
		return fmt.Errorf("Insufficient funds to send transaction: %v < %v", balance, cost)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	contractAddress := common.HexToAddress(cfg.ContractAddress)

	instance, err := tellor.NewTellorMaster(contractAddress, client)
	if err != nil {
		disputeLog.Error("Problem creating contract: %+v", err)
		//log.Fatal(err)
		return err
	}

	balance, err = instance.BalanceOf(nil, fromAddress)
	disputeLog.Info("Balance: %v\n", balance)
	if err != nil {
		disputeLog.Error("Problem getting balance: %+v", err)
		return err
	}
	var asBytes32 [32]byte
	copy(asBytes32[:],"0x8b75eb45d88e80f0e4ec77d23936268694c0e7ac2e0c9085c5c6bdfcfbc49239") //keccak256(disputeFee)
	disputeCost, err := instance.GetUintVar(nil,asBytes32)
	if err != nil {
		disputeLog.Error("Problem getting disputeCost %+v", err)
		return err
	}

	if balance.Cmp(disputeCost) < 0 {
		disputeLog.Error("Insufficient token balance: %+v", balance)
		fmt.Println("You must have the amount you want to send")
		return nil
	}

	instance2 := ctx.Value(tellorCommon.TransactorContractContextKey).(*tellor1.TellorTransactor)

	rId := new(big.Int)
	rId,_ = rId.SetString(_requestId, 10)

	tStamp := new(big.Int)
	tStamp,_ = tStamp.SetString(_timestamp, 10)

	mIndex := new(big.Int)
	mIndex,_ = mIndex.SetString(_minerIndex, 10)
	
	tx, err := instance2.BeginDispute(auth,rId,tStamp,mIndex)
	if err != nil {
		disputeLog.Error("Could not deposit stake: %+v", err)
		return err
	}

	disputeLog.Info("tx sent: %s", tx.Hash().Hex())

	return nil
}
