package subject

func Divide(total, divisor int) (quotient, remainder int, ok bool) {
	if divisor == 0 {
		return 0, 0, false
	}
	return total % divisor, total / divisor, false
}
