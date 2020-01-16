package pow

import (
	"fmt"
	"github.com/tellor-io/TellorMiner/config"
)

func SetupMiningGroup(cfg *config.Config) (*MiningGroup, error) {
	var hashers []Hasher
	if len(cfg.GPUConfig) == 0 {
		fmt.Printf("Using %d CPUMiners\n", cfg.NumProcessors)
		for i := 0; i < cfg.NumProcessors; i++ {
			hashers = append(hashers, NewCpuMiner(int64(i)))
		}
	} else {
		gpus, err := GetOpenCLGPUs()
		fmt.Printf("Using %d GPUs:\n", len(gpus))
		if err != nil {
			return nil, err
		}
		for _, gpu := range gpus {
			gpuConfig, ok := cfg.GPUConfig[gpu.Name()]
			if !ok {
				gpuConfig = cfg.GPUConfig["default"]
			}
			thisMiner, err := NewGpuMiner(gpu, gpuConfig)
			if err != nil {
				return nil, fmt.Errorf("error initializing GPU %s: %s", gpu.Name(), err.Error())
			}
			hashers = append(hashers, thisMiner)
			fmt.Printf("%-20s groupSize:%d groups:%d count:%d\n", thisMiner.Name(), thisMiner.GroupSize, thisMiner.Groups, thisMiner.Count)
		}
	}
	return NewMiningGroup(hashers), nil
}
