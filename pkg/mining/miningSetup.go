// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package mining

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
)

const NumProcessors = 7

func SetupMiningGroup(ctx context.Context, logger log.Logger, cfg *config.Config, contractInstance *contracts.ITellor) (*MiningGroup, error) {
	var hashers []Hasher
	level.Info(logger).Log("msg", "starting CPU mining", "threads", NumProcessors)
	for i := 0; i < NumProcessors; i++ {
		hashers = append(hashers, NewCpuMiner(int64(i)))
	}
	miningGrp, err := NewMiningGroup(ctx, logger, cfg, hashers, contractInstance)
	if err != nil {
		return nil, errors.Wrap(err, "creating new mining group")
	}
	return miningGrp, nil
}
