package subject

func SetItem(values []int, index, value int) []int {
	if index < 0 || index >= len(values) {
		return append(values, value)
	}
	values[index] = value
	return values
}
