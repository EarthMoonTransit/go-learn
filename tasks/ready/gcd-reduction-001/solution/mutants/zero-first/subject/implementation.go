package subject

func GCD(a, b int) int {
	if a == 0 {
		return 0
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
