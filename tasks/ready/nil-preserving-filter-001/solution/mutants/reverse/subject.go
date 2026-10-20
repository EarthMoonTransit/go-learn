package subject

func KeepNonnegative(values []int) []int {
	if values == nil {
		return nil
	}
	out := make([]int, 0)
	for _, v := range values {
		if v >= 0 {
			out = append(out, v)
		}
	}
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out
}
