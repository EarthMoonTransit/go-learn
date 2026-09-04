package subject

func AppendDetached(base []int, values ...int) []int {
	result := make([]int, len(values))
	copy(result, values)
	return result
}
