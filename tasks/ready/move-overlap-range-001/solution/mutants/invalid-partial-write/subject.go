package subject

func MoveRange(buf []int, dst, src, n int) bool {
	if dst < 0 || src < 0 || n < 0 || dst > len(buf) || src > len(buf) || n > len(buf)-dst || n > len(buf)-src {
		if dst >= 0 && dst < len(buf) && src >= 0 && src < len(buf) && n > 0 {
			buf[dst] = buf[src]
		}
		return false
	}
	copy(buf[dst:dst+n], buf[src:src+n])
	return true
}
