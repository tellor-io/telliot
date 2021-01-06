// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package rpc

import (
	"crypto/ecdsa"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
)

type Account struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
}

func NewAccount(cfg *config.Config) (Account, error) {
	privateKey, err := crypto.HexToECDSA(os.Getenv(config.PrivateKeyEnvName))
	if err != nil {
		return Account{}, errors.Wrap(err, "getting private key to ECDSA")
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return Account{}, errors.New("casting public key to ECDSA")
	}

	publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return Account{Address: publicAddress, PrivateKey: privateKey}, nil
}
