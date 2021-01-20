// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package pow

import (
	"fmt"
	"os"

	"github.com/tellor-io/telliot/pkg/config"
)

func SetupMiningGroup(cfg *config.Config, exitCh chan os.Signal) (*MiningGroup, error) {
	var hashers []Hasher
	fmt.Printf("Starting CPU mining, using %d threads\n", cfg.NumProcessors)
	for i := 0; i < cfg.NumProcessors; i++ {
		hashers = append(hashers, NewCpuMiner(int64(i)))
	}
	return NewMiningGroup(hashers, exitCh), nil
}
