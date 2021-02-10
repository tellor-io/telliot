// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package pow

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
)

const NumProcessors = 1

func SetupMiningGroup(logger log.Logger, cfg *config.Config, exitCh chan os.Signal) (*MiningGroup, error) {
	var hashers []Hasher
	level.Info(logger).Log("msg", "starting CPU mining", "threads", NumProcessors)
	for i := 0; i < NumProcessors; i++ {
		hashers = append(hashers, NewCpuMiner(int64(i)))
	}
	miningGrp, err := NewMiningGroup(logger, cfg, hashers, exitCh)
	if err != nil {
		return nil, errors.Wrap(err, "creating new mining group")
	}

	return miningGrp, nil
}
