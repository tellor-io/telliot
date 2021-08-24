// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package cli

import (
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/logging"
)

const VersionMessage = `
    The official Tellor cli tool %s (%s)
    -----------------------------------------
	Website: https://tellor.io
	Github:  https://github.com/tellor-io/telliot
`

var CLI struct {
	Transfer transferCmd `cmd:"" help:"Transfer tokens"`
	Approve  approveCmd  `cmd:"" help:"Approve tokens"`
	Accounts accountsCmd `cmd:"" help:"Show accounts"`
	Balance  balanceCmd  `cmd:"" help:"Check the balance of an address"`
	Stake    struct {
		Deposit  depositCmd  `cmd:"" help:"deposit a stake"`
		Request  requestCmd  `cmd:"" help:"request to withdraw stake"`
		Withdraw withdrawCmd `cmd:"" help:"withdraw stake"`
		Status   statusCmd   `cmd:"" help:"show stake status"`
	} `cmd:"" help:"Perform one of the stake operations"`
	Dispute struct {
		New   newDisputeCmd `cmd:"" help:"start a new dispute"`
		Vote  voteCmd       `cmd:"" help:"vote on a open dispute"`
		List  listCmd       `cmd:"" help:"list open disputes"`
		Tally tallyCmd      `cmd:"" help:"tally votes for a dispute ID"`
	} `cmd:"" help:"Perform commands related to disputes"`
	Dataserver dataserverCmd `cmd:"" help:"launch only a dataserver instance"`
	Mine       mineCmd       `cmd:"" help:"Submit data to oracle contracts"`
	Version    VersionCmd    `cmd:"" help:"Show the CLI version information"`
}

type VersionCmd struct {
}

func (cmd *VersionCmd) Run() error {
	// The main entry point prints the version message so
	// here just return nil and the message will be printed.
	return nil
}

type cfg struct {
	Config configPath `type:"existingfile" help:"path to config file"`
}

type addr struct {
	Addr string `arg:"" required:""`
}

type cfgGas struct {
	cfg
	GasPrice int `optional:"" help:"gas price to use when running the command"`
}

type cfgGasAddr struct {
	cfgGas
	addr
}

type cfgAddr struct {
	cfg
	addr
}

type configPath string

type accountsCmd struct {
	cfg
}

func (self *accountsCmd) Run() error {
	logger := logging.NewLogger()

	_, err := config.ParseConfig(logger, string(self.Config)) // Load the env file.
	if err != nil {
		return errors.Wrap(err, "creating config")
	}

	accounts, err := ethereum.GetAccounts()
	if err != nil {
		return errors.Wrap(err, "getting accounts")
	}

	for i, account := range accounts {
		level.Info(logger).Log("msg", "account", "no", i, "address", account.Address.String())
	}

	return nil
}
