package subject

func CounterValue(counter *int) int {
	if counter == nil {
		return 0
	}
	if *counter == 0 {
		return 1
	}
	return *counter
}
