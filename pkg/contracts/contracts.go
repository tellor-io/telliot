package contracts

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts/getter"
	"github.com/tellor-io/telliot/pkg/contracts/tellor"
)

type Tellor struct {
	Getter  *getter.TellorGetters
	Caller  *tellor.Tellor
	Address common.Address
}

func NewTellor(cfg *config.Config, client bind.ContractBackend) (Tellor, error) {
	contractAddress := common.HexToAddress(cfg.ContractAddress)
	contractTellorInstance, err := tellor.NewTellor(contractAddress, client)
	if err != nil {
		return Tellor{}, errors.Wrap(err, "creating telllor caller")
	}
	contractGetterInstance, err := getter.NewTellorGetters(contractAddress, client)
	if err != nil {
		return Tellor{}, errors.Wrap(err, "creating telllor getter")
	}

	return Tellor{Address: contractAddress, Getter: contractGetterInstance, Caller: contractTellorInstance}, nil
}
