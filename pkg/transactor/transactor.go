// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package transactor

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kit/kit/log"
	tellorCommon "github.com/tellor-io/telliot/pkg/common"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
)

// Transactor implements the Transactor interface.
type Transactor struct {
	logger           log.Logger
	cfg              *config.Config
	proxy            db.DataServerProxy
	client           contracts.ETHClient
	account          *config.Account
	contractInstance *contracts.ITellor
	nonce            string
	reqVals          [5]*big.Int
	reqIds           [5]*big.Int
}

func NewTransactor(logger log.Logger, cfg *config.Config, proxy db.DataServerProxy,
	client contracts.ETHClient, account *config.Account, contractInstance *contracts.ITellor) *Transactor {
	return &Transactor{
		logger:           logger,
		cfg:              cfg,
		proxy:            proxy,
		client:           client,
		account:          account,
		contractInstance: contractInstance,
	}
}

func (s *Transactor) Transact(ctx context.Context, nonce string, reqIds [5]*big.Int, reqVals [5]*big.Int) (*types.Transaction, error) {
	s.nonce = nonce
	s.reqIds = reqIds
	s.reqVals = reqVals
	return rpc.SubmitContractTxn(ctx, s.logger, s.cfg, s.proxy, s.client, s.contractInstance, s.account, "submitSolution", s.submit)

}
func (s *Transactor) submit(ctx context.Context, contract tellorCommon.ContractInterface) (*types.Transaction, error) {
	txn, err := contract.SubmitSolution(
		s.nonce,
		s.reqIds,
		s.reqVals)
	if err != nil {
		return nil, err
	}
	return txn, err
}
