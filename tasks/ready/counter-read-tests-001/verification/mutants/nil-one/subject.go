package subject

func CounterValue(counter *int) int {
	if counter == nil {
		return 1
	}
	return *counter
}
