package subject

func CounterValue(counter *int) int {
	if counter == nil {
		return 0
	}
	v := *counter
	*counter = 0
	return v
}
