package subject

func GetItem(values []int, index int) int {
	if index < 0 || index >= len(values)-1 {
		return -1
	}
	return values[index]
}
