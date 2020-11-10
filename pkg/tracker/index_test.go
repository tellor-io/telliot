// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tracker

// Refactor the tests to not use the network as this will always be flaky.

// func TestPSR(t *testing.T) {
// 	ctx, _, cleanup := testutil.CreateContext(t)
// 	t.Cleanup(cleanup)
// 	psr, err := BuildIndexTrackers()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	for idx := range psr {
// 		err = psr[idx].Exec(ctx)
// 		psrStr := psr[idx].String()
// 		if err != nil {
// 			t.Fatalf("failed to execute psr: %s %v", psrStr, err)
// 		}
// 	}
// 	val, err := ctx.Value(common.DBContextKey).(db.DB).Get(fmt.Sprintf("qv_%d", 1))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if val == nil {
// 		t.Fatal(fmt.Errorf("Expected a value stored for request ID 1"))
// 	}
// 	intVal, err := hexutil.DecodeBig(string(val))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Println("DB value", intVal)
// }
