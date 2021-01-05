// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package common

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/tellor-io/telliot/pkg/contracts/master"
	"github.com/tellor-io/telliot/pkg/contracts/proxy"
)

type Contract struct {
	Getter  *proxy.TellorGetters
	Caller  *master.Tellor
	Address common.Address
}

type Account struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
}
