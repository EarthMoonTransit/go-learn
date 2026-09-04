package subject

func WideSum(a, b int8) int {
	sum := int(a) + int(b)
	if sum > 127 {
		return 127
	}
	if sum < -128 {
		return -128
	}
	return sum
}
