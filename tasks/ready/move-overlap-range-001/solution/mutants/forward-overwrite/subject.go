package subject

func MoveRange(buf []int, dst, src, n int) bool {
	if dst < 0 || src < 0 || n < 0 || dst > len(buf) || src > len(buf) || n > len(buf)-dst || n > len(buf)-src {
		return false
	}
	for i := 0; i < n; i++ {
		buf[dst+i] = buf[src+i]
	}
	return true
}
