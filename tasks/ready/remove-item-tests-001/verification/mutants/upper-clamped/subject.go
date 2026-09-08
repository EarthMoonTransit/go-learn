package subject

func RemoveItem(values []int, index int) []int {
	if index >= len(values) && len(values) > 0 {
		index = len(values) - 1
	}
	if index < 0 || index >= len(values) {
		return values
	}
	return append(values[:index], values[index+1:]...)
}
