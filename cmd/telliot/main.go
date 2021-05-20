// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/ethereum"
)

func createTellorVariables(ctx context.Context, logger log.Logger, cfg ethereum.Config) (contracts.ETHClient, *contracts.ITellor, []*ethereum.Account, error) {
	client, err := ethereum.NewClient(logger, cfg, os.Getenv(config.NodeURLEnvName))
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "create rpc client instance")
	}
	contract, err := contracts.NewITellor(client)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "create tellor master instance")
	}
	accounts, err := ethereum.GetAccounts()
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "creating accounts")
	}

	// Issue #55, halt if client is still syncing with Ethereum network
	s, err := client.IsSyncing(ctx)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "determining if Ethereum client is syncing")
	}
	if s {
		return nil, nil, nil, errors.New("ethereum node is still syncing with the network")
	}

	return client, contract, accounts, nil
}

var CLI struct {
	Migrate  migrateCmd  `cmd:"" help:"Migrate funds from the old oracle contract"`
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
		New  newDisputeCmd `cmd:"" help:"start a new dispute"`
		Vote voteCmd       `cmd:"" help:"vote on a open dispute"`
		Show showCmd       `cmd:"" help:"show open disputes"`
	} `cmd:"" help:"Perform commands related to disputes"`
	Dataserver dataserverCmd `cmd:"" help:"launch only a dataserver instance"`
	Mine       mineCmd       `cmd:"" help:"mine TRB and submit values"`
	Version    VersionCmd    `cmd:"" help:"Show the CLI version information"`
}

func main() {
	//lint:ignore faillint it should print to console
	fmt.Printf(versionMessage, GitTag, GitHash)
	ctx := kong.Parse(&CLI, kong.Name("Telliot"),
		kong.Description("The official Tellor cli tool"),
		kong.UsageOnError())

	ctx.FatalIfErrorf(ctx.Run(*ctx))
}
