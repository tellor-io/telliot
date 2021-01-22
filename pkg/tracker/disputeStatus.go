// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

import (
	"context"
	"fmt"

	"strings"

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
	contract *contracts.Tellor
	account  *rpc.Account
	logger   log.Logger
	config   *config.Config
}

func (b *DisputeTracker) String() string {
	return DisputeTrackerName
}

func NewDisputeTracker(logger log.Logger, config *config.Config, db db.DataServerProxy, contract *contracts.Tellor, account *rpc.Account) *DisputeTracker {
	return &DisputeTracker{
		config:   config,
		db:       db,
		contract: contract,
		account:  account,
		logger:   log.With(logger, "component", "dispute tracker"),
	}
}

func (b *DisputeTracker) Exec(ctx context.Context) error {

	status, _, err := b.contract.Getter.GetStakerInfo(nil, b.account.Address)

	if err != nil {
		return errors.Wrap(err, "getting staker info")
	}
	enc := hexutil.EncodeBig(status)
	level.Info(b.logger).Log("msg", "staker status", "status", enc)
	err = b.db.Put(db.DisputeStatusKey, []byte(enc))
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
		status, _, err := b.contract.Getter.GetStakerInfo(nil, address)
		if err != nil {
			level.Error(b.logger).Log("msg", "getting staker dispute status for miner", "address", addr, "err", err)
		}
		level.Info(b.logger).Log("msg", "whitelisted miner", "address", addr, "status", status)
		dbKey := fmt.Sprintf("%s-%s", strings.ToLower(address.Hex()), db.DisputeStatusKey)
		err = b.db.Put(dbKey, []byte(hexutil.EncodeBig(status)))
		if err != nil {
			level.Error(b.logger).Log("msg", "storing staker dispute status", "err", err)
		}
	}
	return nil
}
