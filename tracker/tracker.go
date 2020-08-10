package tracker

import "context"

//Tracker is the primary interface for the various tracking options
type Tracker interface {
	//Exec will be run as a go function. The given context will be a KeyValue context containing
	//the client to use for tracking ops.
	Exec(ctx context.Context) error
	String() string
}
