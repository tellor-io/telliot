package aggregator

import "github.com/tellor-io/telliot/pkg/util"

type Config struct {
	MinConfidence float64
	Interval      util.Duration
}
