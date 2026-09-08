package subject

func CounterValue(counter *int) int {
	if counter == nil {
		return 0
	}
	return *counter
}
