package subject

func AppendDetached(base []int, values ...int) []int {
	if len(base) > 0 && len(values) > 0 {
		base[0] = values[0]
	}
	result := make([]int, len(base)+len(values))
	copy(result, base)
	copy(result[len(base):], values)
	return result
}
