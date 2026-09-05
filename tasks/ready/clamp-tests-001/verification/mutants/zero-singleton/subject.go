package subject

func Clamp(value, min, max int) int {
	if min == max {
		return 0
	}
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
