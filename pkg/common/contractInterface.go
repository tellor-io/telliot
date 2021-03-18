// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package common

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// TransactionGeneratorFN is a callback function that a transaction submitter uses to actually invoke
// a contract and generate a transaction.
type TransactionGeneratorFN func(ctx context.Context, contract ContractInterface) (*types.Transaction, error)

// ContractInterface represents an abstraction of function definitions that can be
// called on the smart contract. This is mostly so that we can do unit tests without
// needing to call the actual contract.
type ContractInterface interface {
	AddTip(requestID *big.Int, amount *big.Int) (*types.Transaction, error)
	SubmitSolution(solution string, requestID [5]*big.Int, value [5]*big.Int) (*types.Transaction, error)
	DidMine(challenge [32]byte) (bool, error)
}
