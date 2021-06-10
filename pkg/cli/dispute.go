// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/contracts"
	tEthereum "github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/format"
	psr "github.com/tellor-io/telliot/pkg/psr/tellor"
)

func Dispute(
	ctx context.Context,
	logger log.Logger,
	client contracts.ETHClient,
	contract *contracts.ITellor,
	account *tEthereum.Account,
	requestId *big.Int,
	timestamp *big.Int,
	minerIndex *big.Int,
) error {

	if !minerIndex.IsUint64() || minerIndex.Uint64() > 4 {
		return errors.Errorf("miner index should be between 0 and 4 (got %s)", minerIndex.Text(10))
	}

	balance, err := contract.BalanceOf(nil, account.Address)
	if err != nil {
		return errors.Wrap(err, "fetch balance")
	}
	var asBytes32 [32]byte
	copy(asBytes32[:], crypto.Keccak256([]byte("_DISPUTE_FEE")))
	disputeCost, err := contract.GetUintVar(nil, asBytes32)
	if err != nil {
		return errors.Wrap(err, "get dispute cost")
	}

	if balance.Cmp(disputeCost) < 0 {
		return errors.Errorf("insufficient balance TRB actual: %v, TRB required:%v)",
			format.ERC20Balance(balance),
			format.ERC20Balance(disputeCost))
	}

	auth, err := tEthereum.PrepareEthTransaction(ctx, client, account)
	if err != nil {
		return errors.Wrapf(err, "prepare ethereum transaction")
	}

	tx, err := contract.BeginDispute(auth, requestId, timestamp, minerIndex)
	if err != nil {
		return errors.Wrap(err, "send dispute txn")
	}
	level.Info(logger).Log("msg", "dispute started", "txn", tx.Hash().Hex())
	return nil
}

func Vote(
	ctx context.Context,
	logger log.Logger,
	client contracts.ETHClient,
	contract *contracts.ITellor,
	account *tEthereum.Account,
	disputeId *big.Int,
	supportsDispute bool,
) error {

	voted, err := contract.DidVote(nil, disputeId, contract.Address)
	if err != nil {
		return errors.Wrapf(err, "check if you've already voted")
	}
	if voted {
		level.Info(logger).Log("msg", "you have already voted on this dispute")
		return nil
	}

	auth, err := tEthereum.PrepareEthTransaction(ctx, client, account)
	if err != nil {
		return errors.Wrapf(err, "prepare ethereum transaction")
	}
	tx, err := contract.Vote(auth, disputeId, supportsDispute)
	if err != nil {
		return errors.Wrapf(err, "submit vote transaction")
	}

	level.Info(logger).Log("msg", "vote submitted with transaction", "tx", tx.Hash().Hex())
	return nil
}

func List(
	ctx context.Context,
	logger log.Logger,
	client contracts.ETHClient,
	contract *contracts.ITellor,
	account *tEthereum.Account,
	psr *psr.Psr,
) error {
	// TODO check how it is done in tellorscan and implement here.
	return nil
}
