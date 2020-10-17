// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"context"

	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/rpc"
)

// TxnSubmitter just concrete type for txn submitter.
type TxnSubmitter struct {
}

// NewSubmitter creates a new TxnSubmitter instance.
func NewSubmitter() TxnSubmitter {
	return TxnSubmitter{}
}

// PrepareTransaction relies on rpc package to prepare and submit transactions.
func (s TxnSubmitter) PrepareTransaction(ctx context.Context, proxy db.DataServerProxy, ctxName string, callback tellorCommon.TransactionGeneratorFN) error {
	return rpc.PrepareContractTxn(ctx, proxy, ctxName, callback)
}
