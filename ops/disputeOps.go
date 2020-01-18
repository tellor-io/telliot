package ops

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	tellorCommon "github.com/tellor-io/TellorMiner/common"
	tellor "github.com/tellor-io/TellorMiner/contracts"
	tellor1 "github.com/tellor-io/TellorMiner/contracts1"
	"github.com/tellor-io/TellorMiner/util"
)

/**
 * This file handles all operations related to disputes
 */

func Dispute(requestId *big.Int, timestamp *big.Int, minerIndex *big.Int, ctx context.Context) error {

	if !minerIndex.IsUint64() || minerIndex.Uint64() > 4 {
		return fmt.Errorf("miner index should be between 0 and 4 (got %s)", minerIndex.Text(10))
	}

	instance := ctx.Value(tellorCommon.MasterContractContextKey).(*tellor.TellorMaster)
	addr := ctx.Value(tellorCommon.PublicAddress).(common.Address)

	balance, err := instance.BalanceOf(nil, addr)
	if err != nil {
		return fmt.Errorf("failed to fetch balance: %s", err.Error())
	}
	var asBytes32 [32]byte
	copy(asBytes32[:],"0x8b75eb45d88e80f0e4ec77d23936268694c0e7ac2e0c9085c5c6bdfcfbc49239") //keccak256(disputeFee)
	disputeCost, err := instance.GetUintVar(nil,asBytes32)
	if err != nil {
		return fmt.Errorf("failed to get dispute cost: %s", err)
	}

	if balance.Cmp(disputeCost) < 0 {
		return fmt.Errorf("insufficient balance (%s TRB) disputes require (%s TRB)",
			util.FormatERC20Balance(balance),
			util.FormatERC20Balance(disputeCost))
	}

	auth, err := PrepareEthTransaction(ctx)
	if err != nil {
		return fmt.Errorf("failed to prepare ethereum transaction: %s", err.Error())
	}

	instance2 := ctx.Value(tellorCommon.TransactorContractContextKey).(*tellor1.TellorTransactor)
	tx, err := instance2.BeginDispute(auth,requestId,timestamp,minerIndex)
	if err != nil {
		return fmt.Errorf("failed to send dispute txn: %s", err.Error())
	}
	fmt.Printf("dispute started with txn: %s\n", tx.Hash().Hex())
	return nil
}

func Vote(_disputeId *big.Int, _supportsDispute bool, ctx context.Context) error {

	instance2 := ctx.Value(tellorCommon.TransactorContractContextKey).(*tellor1.TellorTransactor)

	auth, err := PrepareEthTransaction(ctx)
	if err != nil {
		return fmt.Errorf("failed to prepare ethereum transaction: %s", err.Error())
	}
	tx, err := instance2.Vote(auth,_disputeId,_supportsDispute)
	if err != nil {
		return fmt.Errorf("failed to submit vote transaction: %s", err.Error())
	}

	fmt.Printf("Vote submitted with transaction %s\n", tx.Hash().Hex())
	return nil
}

