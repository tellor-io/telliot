package tracker

//CreateTracker a tracker instance by its well-known name
func createTracker(name string) (Tracker, error) {
	switch name {
	case "test":
		{
			return &TestTracker{}, nil
		}

	case "balance":
		{
			return &BalanceTracker{}, nil
		}
	}

	return nil, nil
}
