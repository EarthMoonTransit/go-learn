package subject

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, b%b
	}
	return a
}
