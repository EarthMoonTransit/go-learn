package subject

func RemoveItem(values []int, index int) []int {
	if len(values) == 0 {
		return []int{0}
	}
	if index < 0 || index >= len(values) {
		return values
	}
	return append(values[:index], values[index+1:]...)
}
