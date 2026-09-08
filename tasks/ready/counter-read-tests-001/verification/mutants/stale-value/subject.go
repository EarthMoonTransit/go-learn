package subject

var values = map[*int]int{}

func CounterValue(counter *int) int {
	if counter == nil {
		return 0
	}
	if v, ok := values[counter]; ok {
		return v
	}
	values[counter] = *counter
	return *counter
}
