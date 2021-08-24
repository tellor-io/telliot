// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package gasPrice

import (
	"context"
	"math/big"
)

type GasPriceQuerier interface {
	Query(ctx context.Context) (*big.Int, error)
}
