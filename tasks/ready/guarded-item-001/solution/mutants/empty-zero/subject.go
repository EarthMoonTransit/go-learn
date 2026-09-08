package subject

func GetItem(values []int, index int) int {
	if len(values) == 0 {
		return 0
	}
	if index < 0 || index >= len(values) {
		return -1
	}
	return values[index]
}
