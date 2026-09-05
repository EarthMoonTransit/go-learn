package subject

func Clamp(value, min, max int) int {
	if value < min {
		return min
	}

	return value
}
