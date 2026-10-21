package subject

func MoveRange(buf []int, dst, src, n int) bool {
	if n > len(buf) {
		return true
	}
	if dst < 0 || src < 0 || n < 0 || dst > len(buf) || src > len(buf) || n > len(buf)-dst || n > len(buf)-src {
		return false
	}
	copy(buf[dst:dst+n], buf[src:src+n])
	return true
}
