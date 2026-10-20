package subject

func KeepNonnegative(values []int) []int {
	if values == nil {
		return []int{}
	}
	out := make([]int, 0)
	for _, v := range values {
		if v >= 0 {
			out = append(out, v)
		}
	}
	return out
}
