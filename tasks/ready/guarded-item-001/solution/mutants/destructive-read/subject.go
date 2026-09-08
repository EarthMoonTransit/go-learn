package subject

func GetItem(values []int, index int) int {
	if index < 0 || index >= len(values) {
		return -1
	}
	value := values[index]
	values[index] = 0
	return value
}
