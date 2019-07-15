package tracker

import (
	"github.com/tellor-io/TellorMiner/util"
)

var (
	//ClientContextKey is the key used to set the eth client on tracker contexts
	ClientContextKey = util.NewKey("tracker", "ETHClient")
)
