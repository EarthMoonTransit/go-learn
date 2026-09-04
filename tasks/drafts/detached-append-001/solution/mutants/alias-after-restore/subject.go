package subject

func AppendDetached(base []int, values ...int) []int {
	result := append(base, values...)
	copy(base, result[:len(base)])
	return result
}
