package tracker

import (
	"context"
	"fmt"
)

//TestTracker is just a test tracker demonstrating how to write a tracker
type TestTracker struct {
}

//Exec impl for test tracker
func (t *TestTracker) Exec(ctx context.Context) error {
	fmt.Printf("Test execution with client: %+v\n", ctx.Value(ClientContextKey))
	return nil
}
