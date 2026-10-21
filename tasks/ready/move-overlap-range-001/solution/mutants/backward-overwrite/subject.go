package subject

func MoveRange(buf []int, dst, src, n int) bool {
	if dst < 0 || src < 0 || n < 0 || dst > len(buf) || src > len(buf) || n > len(buf)-dst || n > len(buf)-src {
		return false
	}
	for i := n - 1; i >= 0; i-- {
		buf[dst+i] = buf[src+i]
	}
	return true
}
