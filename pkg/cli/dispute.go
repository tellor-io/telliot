// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

func getNonceSubmits(
	ctx context.Context,
	client contracts.ETHClient,
	contract *contracts.ITellor,
	valueBlock *big.Int,
	dispute *contracts.ITellorNewDispute,
) ([]struct {
	time.Time
	float64
}, error) {
	abi, err := abi.JSON(strings.NewReader(contracts.ITellorABI))
	if err != nil {
		return nil, errors.Wrap(err, "parse abi")
	}

	// Just use nil for most of the variables, only using this object to call UnpackLog which only uses the abi
	bar := bind.NewBoundContract(contract.Address, abi, nil, nil, nil)

	allVals, err := contract.GetSubmissionsByTimestamp(nil, dispute.RequestId, dispute.Timestamp)
	if err != nil {
		return nil, errors.Wrap(err, "get other submitted values for dispute")
	}

	allAddrs, err := contract.GetMinersByRequestIdAndTimestamp(nil, dispute.RequestId, dispute.Timestamp)
	if err != nil {
		return nil, errors.Wrap(err, "get miner addresses for dispute")
	}

	const blockStep = 100
	high := int64(valueBlock.Uint64())
	low := high - blockStep
	nonceSubmitID := abi.Events["NonceSubmitted"].ID
	timedValues := make([]struct {
		time.Time
		float64
	}, 5)
	found := 0
	for found < 5 {
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(low),
			ToBlock:   big.NewInt(high),
			Addresses: []common.Address{contract.Address},
			Topics:    [][]common.Hash{{nonceSubmitID}},
		}

		logs, err := client.FilterLogs(ctx, query)
		if err != nil {
			return nil, errors.Wrap(err, "get nonce logs")
		}

		for _, l := range logs {
			nonceSubmit := contracts.TellorNonceSubmitted{}
			err := bar.UnpackLog(&nonceSubmit, "NonceSubmitted", l)
			if err != nil {
				return nil, errors.Wrap(err, "unpack into object")
			}
			header, err := client.HeaderByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
			if err != nil {
				return nil, errors.Wrap(err, "get nonce block header")
			}
			for i := 0; i < 5; i++ {
				if nonceSubmit.Miner == allAddrs[i] {
					valTime := time.Unix(int64(header.Time), 0)

					bigF := new(big.Float)
					bigF.SetInt(allVals[i])
					f, _ := bigF.Float64()
					timedValues[i] = struct {
						time.Time
						float64
					}{
						valTime,
						f,
					}
					found++
					break
				}
			}
		}
		high -= blockStep
		low = high - blockStep
	}
	return timedValues, nil
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
