package tracker

import (
	"context"
	"fmt"
	"log"

	"github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
)

//TestTracker is just a test tracker demonstrating how to write a tracker
type TestTracker struct {
}

//Exec impl for test tracker
func (t *TestTracker) Exec(ctx context.Context) error {
	fmt.Printf("Test execution with client: %+v, DB: %+v\n", ctx.Value(common.ClientContextKey), ctx.Value(common.DBContextKey))
	db := ctx.Value(common.DBContextKey).(db.DB)
	err := db.Put("TEST", []byte("Value"))
	if err != nil {
		log.Fatal(err)
		panic("Could not store value")
	}

	v, err := db.Get("TEST")
	if err != nil {
		log.Fatal(err)
		panic("Could not get value")
	}
	fmt.Printf("Value %v\n", string(v))
	return nil
}
