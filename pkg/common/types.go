// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package common

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/tellor-io/TellorMiner/pkg/contracts/getter"
	"github.com/tellor-io/TellorMiner/pkg/contracts/tellor"
)

type Contract struct {
	Getter  *getter.TellorGetters
	Caller  *tellor.Tellor
	Address common.Address
}

type Account struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
}
