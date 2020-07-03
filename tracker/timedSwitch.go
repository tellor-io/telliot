package tracker

import (
	"github.com/tellor-io/TellorMiner/apiOracle"
	"time"
)

type TimedSwitch struct {
	before ValueGenerator
	after  ValueGenerator
	at     time.Time
}

func (t *TimedSwitch) Require(at time.Time) map[string]IndexProcessor {
	if at.After(t.at) {
		return t.after.Require(at)
	} else {
		return t.before.Require(at)
	}
}

func (t *TimedSwitch) ValueAt(vals map[string]apiOracle.PriceInfo, at time.Time) float64 {
	//dont check time here, only in require. we don't want this to change mid-cycle
	//and pass the old requirements to the new processor
	if at.After(t.at) {
		return t.after.ValueAt(vals, at)
	} else {
		return t.before.ValueAt(vals, at)
	}
}
