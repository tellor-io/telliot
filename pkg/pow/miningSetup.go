// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package pow

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/tellor-io/TellorMiner/pkg/config"
)

func SetupMiningGroup(cfg *config.Config) (*MiningGroup, error) {
	var hashers []Hasher
	gpus, err := GetOpenCLGPUs()
	fmt.Printf("Found %d GPUs:\n", len(gpus))
	if err != nil {
		return nil, err
	}
	for _, gpu := range gpus {
		gpuConfig, ok := cfg.GPUConfig[gpu.Name()]
		if !ok {
			gpuConfig = cfg.GPUConfig["default"]
		}
		if gpuConfig != nil && gpuConfig.Disabled {
			fmt.Printf("%s disabled in config, ignoring\n", gpu.Name())
			continue
		}
		thisMiner, err := NewGpuMiner(gpu, gpuConfig, cfg.EnablePoolWorker)
		if err != nil {
			return nil, errors.Wrapf(err, "initializing GPU %s", gpu.Name())
		}
		hashers = append(hashers, thisMiner)
		fmt.Printf("%-20s groupSize:%d groups:%d count:%d\n", thisMiner.Name(), thisMiner.GroupSize, thisMiner.Groups, thisMiner.Count)
	}
	if len(hashers) == 0 {
		fmt.Printf("No GPUs enabled, falling back to CPU mining, using %d threads\n", cfg.NumProcessors)
		for i := 0; i < cfg.NumProcessors; i++ {
			hashers = append(hashers, NewCpuMiner(int64(i)))
		}

	}
	return NewMiningGroup(hashers), nil
}
