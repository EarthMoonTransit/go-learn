package subject

func DetachedPrefix(in []int, n int) []int {
	out := make([]int, len(in))
	copy(out, in[:n])
	return out
}
