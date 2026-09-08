package subject

func SetItem(values []int, index, value int) []int {
	if index < 0 || index >= len(values) {
		return append(values, value)
	}
	return append(values, value)
}
