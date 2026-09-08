package subject

func DetachedPrefix(in []int, n int) []int {
	if n == 0 {
		return in[:0]
	}
	out := make([]int, n)
	copy(out, in[:n])
	return out
}
