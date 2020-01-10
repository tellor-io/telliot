package pow

import (
	"fmt"
	"github.com/tellor-io/TellorMiner/config"
)

func SetupMiningGroup(cfg *config.Config) (*MiningGroup, error) {
	var hashers []Hasher
	if !cfg.UseGPU {
		fmt.Printf("Using %d CPUMiners\n", cfg.NumProcessors)
		for i := 0; i < cfg.NumProcessors; i++ {
			hashers = append(hashers, NewCpuMiner(10e3))
		}
	} else {
		gpus, err := GetOpenCLGPUs()
		fmt.Printf("Using %d GPUMiners\n", len(gpus))
		if err != nil {
			return nil, err
		}
		for _, gpu := range gpus {
			thisMiner, err := NewGpuMiner(gpu)
			if err != nil {
				return nil, fmt.Errorf("Error initializing GPU %s: %s", gpu.Name(), err.Error())
			}
			hashers = append(hashers, thisMiner)
		}
	}
	return NewMiningGroup(hashers), nil
}
