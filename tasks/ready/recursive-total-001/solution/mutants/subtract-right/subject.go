package subject

func Total(values []int) int {
	if len(values) == 0 {
		return 0
	}
	if len(values) == 1 {
		return values[0]
	}
	middle := len(values) / 2
	return Total(values[:middle]) - Total(values[middle:])
}
