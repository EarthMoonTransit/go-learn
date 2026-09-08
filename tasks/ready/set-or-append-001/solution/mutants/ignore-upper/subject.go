package subject

func SetItem(values []int, index, value int) []int {
	if index >= len(values) {
		return values
	}
	if index < 0 {
		return append(values, value)
	}
	values[index] = value
	return values
}
