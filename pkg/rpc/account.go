// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package rpc

import (
	"crypto/ecdsa"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
)

type Account struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
}

// NewAccounts returns a slice of Account from private keys in
// PrivateKeysEnvName environment variable.
func NewAccounts(cfg *config.Config) ([]*Account, error) {
	_privateKeys := os.Getenv(config.PrivateKeysEnvName)
	privateKeys := strings.Split(_privateKeys, ",")

	// Create an Account instance per private keys.
	accounts := make([]*Account, len(privateKeys))
	for i, pkey := range privateKeys {
		privateKey, err := crypto.HexToECDSA(strings.TrimSpace(pkey))
		if err != nil {
			return nil, errors.Wrap(err, "getting private key to ECDSA")
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return nil, errors.New("casting public key to ECDSA")
		}

		publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		accounts[i] = &Account{Address: publicAddress, PrivateKey: privateKey}
	}
	return accounts, nil
}
