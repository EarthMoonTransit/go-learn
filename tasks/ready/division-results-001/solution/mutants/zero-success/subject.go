package subject

func Divide(total, divisor int) (quotient, remainder int, ok bool) {
	if divisor == 0 {
		return 0, 0, true
	}
	return total / divisor, total % divisor, true
}
