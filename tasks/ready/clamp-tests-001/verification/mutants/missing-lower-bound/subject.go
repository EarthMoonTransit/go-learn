package subject

func Clamp(value, min, max int) int {

	if value > max {
		return max
	}
	return value
}
