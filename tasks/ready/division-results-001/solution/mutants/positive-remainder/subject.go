package subject

func Divide(total, divisor int) (quotient, remainder int, ok bool) {
	if divisor == 0 {
		return 0, 0, false
	}
	r := total % divisor
	if r < 0 {
		r = -r
	}
	return total / divisor, r, true
}
