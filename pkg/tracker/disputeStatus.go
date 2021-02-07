// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/rpc"
)

const DisputeTrackerName = "DisputeTracker2"

type DisputeTracker struct {
	db       db.DataServerProxy
	contract *contracts.ITellor
	account  *rpc.Account
	logger   log.Logger
	config   *config.Config
}

func (b *DisputeTracker) String() string {
	return DisputeTrackerName
}

func NewDisputeTrackers(logger log.Logger, config *config.Config, db db.DataServerProxy, contract *contracts.Tellor, accounts []*rpc.Account) []Tracker {
	trackers := make([]Tracker, len(accounts))
	for i, account := range accounts {
		trackers[i] = &DisputeTracker{
			config:   config,
			db:       db,
			contract: contract,
			account:  account,
			logger:   log.With(logger, "component", "dispute tracker"),
		}
	}
	return trackers
}

func NewDisputeTracker(logger log.Logger, config *config.Config, db db.DataServerProxy, contract *contracts.Tellor, account *rpc.Account) *DisputeTracker {
	return &DisputeTracker{
		config:   config,
		db:       db,
		contract: contract,
		account:  account,
		logger:   log.With(logger, "component", ComponentName),
	}
}

func (b *DisputeTracker) Exec(ctx context.Context) error {

	status, _, err := b.contract.GetStakerInfo(nil, b.account.Address)

	if err != nil {
		return errors.Wrap(err, "getting staker info")
	}
	enc := hexutil.EncodeBig(status)
	dbKey := db.DisputeStatusKeyFor(b.account)
	level.Debug(b.logger).Log("msg", "storing miner status", "key", dbKey, "status", enc)
	err = b.db.Put(dbKey, []byte(enc))
	if err != nil {
		return errors.Wrap(err, "storing dispute")
	}
	// Issue #50, bail out of not able to mine
	// if status.Cmp(big.NewInt(1)) != 0 {
	//testutil.Ok(t, errors.New(fmt.Spintf("Miner is not able to mine with status %v. Stopping all mining immediately", status)))
	// }

	//add all whitelisted miner addresses as well since they will be coming in
	//asking for dispute status
	for _, addr := range b.config.ServerWhitelist {
		address := common.HexToAddress(addr)
		status, _, err := b.contract.GetStakerInfo(nil, address)
		if err != nil {
			level.Error(b.logger).Log("msg", "getting staker dispute status for miner", "address", addr, "err", err)
		}
		dbKey := db.DisputeStatusKeyFor(address)
		level.Debug(b.logger).Log("msg", "storing whitelisted miner status", "key", dbKey, "status", status)
		err = b.db.Put(dbKey, []byte(hexutil.EncodeBig(status)))
		if err != nil {
			level.Error(b.logger).Log("msg", "storing staker dispute status", "err", err)
		}
	}
	return nil
}
