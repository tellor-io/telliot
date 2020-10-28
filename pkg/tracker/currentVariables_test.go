// // Copyright (c) The Tellor Authors.
// // Licensed under the MIT License.

package tracker

// import (
// 	"fmt"
// 	"math/big"
// 	"testing"

// 	tellorCommon "github.com/tellor-io/TellorMiner/pkg/common"
// 	"github.com/tellor-io/TellorMiner/pkg/testutil"
// 	"github.com/tellor-io/TellorMiner/pkg/util"

// 	"github.com/ethereum/go-ethereum/common/hexutil"
// 	"github.com/tellor-io/TellorMiner/pkg/db"
// )

// func TestCurrentVarableString(t *testing.T) {
// 	tracker := New_CurrentVariablesTracker(util.SetupLogger("debug"))
// 	res := tracker.String()
// 	if res != CurrentVariablesTrackerName {
// 		t.Fatal("should return string", CurrentVariablesTrackerName)
// 	}
// }

// func TestCurrentVariables(t *testing.T) {
// 	ctx, _, cleanup := testutil.CreateContext(t)
// 	t.Cleanup(cleanup)
// 	tracker := New_CurrentVariablesTracker(util.SetupLogger("debug"))

// 	fmt.Println("Working to Line 41")
// 	err := tracker.Exec(ctx)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	DB := ctx.Value(tellorCommon.DBContextKey).(db.DB)
// 	v, err := DB.Get(db.RequestIdKey)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Println("Working to Line 51", v)
// 	b, err := hexutil.DecodeBig(string(v))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("RequestID stored: %v\n", string(v))
// 	if b.Cmp(big.NewInt(1)) != 0 {
// 		t.Fatalf("Current Request ID from client did not match what should have been stored in DB. %v != %v", b, fmt.Sprint(1))
// 	}

// 	v, err = DB.Get(db.QueryStringKey)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if string(v) != "json(https://coinbase.com)" {
// 		t.Fatalf("Expected query string to match test input: %s != %s\n", string(v), "json(https://coinbase.com)")
// 	}
// }
