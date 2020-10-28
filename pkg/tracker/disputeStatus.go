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
	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/contracts/getter"
	"github.com/tellor-io/TellorMiner/pkg/db"
	"github.com/tellor-io/TellorMiner/pkg/rpc"
)

const DisputeTrackerName = "DisputeTracker2"

type DisputeTracker struct {
	logger log.Logger
}

func (b *DisputeTracker) String() string {
	return DisputeTrackerName
}

func NewDisputeTracker(logger log.Logger) *DisputeTracker {
	return &DisputeTracker{
		logger: log.With(logger, "component", "dispute tracker"),
	}
}

func (b *DisputeTracker) Exec(ctx context.Context) error {
	//cast client using type assertion since context holds generic interface{}
	client := ctx.Value(tellorCommon.ClientContextKey).(rpc.ETHClient)
	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)

	//get the single config instance
	cfg := config.GetConfig()

	//get address from config
	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)

	_conAddress := cfg.ContractAddress

	//convert to address
	contractAddress := common.HexToAddress(_conAddress)

	instance, err := getter.NewTellorGetters(contractAddress, client)
	if err != nil {
		return errors.Wrap(err, "getting master instance")
	}

	status, _, err := instance.GetStakerInfo(nil, fromAddress)

	if err != nil {
		return errors.Wrap(err, "getting staker info")
	}
	enc := hexutil.EncodeBig(status)
	level.Info(b.logger).Log("msg", "staker status", "status", enc)
	err = DB.Put(db.DisputeStatusKey, []byte(enc))
	if err != nil {
		return errors.Wrap(err, "storing dispute")
	}
	// Issue #50, bail out of not able to mine
	// if status.Cmp(big.NewInt(1)) != 0 {
	// 	log.Fatalf("Miner is not able to mine with status %v. Stopping all mining immediately", status)
	// }

	//add all whitelisted miner addresses as well since they will be coming in
	//asking for dispute status
	for _, addr := range cfg.ServerWhitelist {
		address := common.HexToAddress(addr)
		status, _, err := instance.GetStakerInfo(nil, address)
		if err != nil {
			level.Error(b.logger).Log("msg", "getting staker dispute status for miner", "address", addr, "err", err)
		}
		level.Info(b.logger).Log("msg", "whitelisted miner", "address", addr, "status", status)
		dbKey := fmt.Sprintf("%s-%s", strings.ToLower(address.Hex()), db.DisputeStatusKey)
		err = DB.Put(dbKey, []byte(hexutil.EncodeBig(status)))
		if err != nil {
			level.Error(b.logger).Log("msg", "storing staker dispute status", "err", err)
		}
	}
	return nil
}
