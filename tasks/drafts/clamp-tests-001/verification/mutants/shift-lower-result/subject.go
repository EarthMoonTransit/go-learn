package subject

func Clamp(value, min, max int) int {
	if value < min {
		return min + 1
	}
	if value > max {
		return max
	}
	return value
}
