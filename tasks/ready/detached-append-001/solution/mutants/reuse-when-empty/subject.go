package subject

func AppendDetached(base []int, values ...int) []int {
	if len(values) == 0 {
		return base
	}
	result := make([]int, len(base), len(base)+len(values))
	copy(result, base)
	return append(result, values...)
}
