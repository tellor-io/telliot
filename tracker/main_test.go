package tracker

import (
	"fmt"
	"os"
	"testing"

	"github.com/tellor-io/TellorMiner/apiOracle"
	"github.com/tellor-io/TellorMiner/config"
	"github.com/tellor-io/TellorMiner/util"
)

// TODO: Set threshold low and test the  "out of range" failure
var configJSON = `{
	"publicAddress":"0000000000000000000000000000000000000000",
	"privateKey":"1111111111111111111111111111111111111111111111111111111111111111",
	"contractAddress":"0x724D1B69a7Ba352F11D73fDBdEB7fF869cB22E19",
	"trackers": ["indexers", "balance", "currentVariables", "disputeStatus", "gas", "disputeChecker"],
	"IndexFolder": "..",
	"disputeThreshold": 1.0, 
	"disputeTimeDelta": "50s"
}
`

func TestMain(m *testing.M) {
	err := config.ParseConfigBytes([]byte(configJSON))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse mock config: %v\n", err)
		os.Exit(-1)
	}
	util.ParseLoggingConfig("")
	apiOracle.EnsureValueOracle()
	os.Exit(m.Run())
}

