package tracker

//CreateTracker a tracker instance by its well-known name
func createTracker(name string) (Tracker, error) {
	if name == "test" {
		return &TestTracker{}, nil
	}
	return nil, nil
}
