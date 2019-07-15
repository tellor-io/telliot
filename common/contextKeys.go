package common

import (
	"github.com/tellor-io/TellorMiner/util"
)

var (
	//DBContextKey is the shared context key where a DB instance can be found in a context
	DBContextKey = util.NewKey("common", "DB")
)
