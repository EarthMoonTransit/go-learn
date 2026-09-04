package subject

func Clamp(value, min, max int) int {
	if value == min && min < max {
		return max
	}
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
