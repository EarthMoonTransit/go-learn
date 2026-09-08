package subject

func ShiftRight(values []int, index int) int {
	last := values[len(values)-1]
	for i := index + 1; i < len(values); i++ {
		values[i] = values[i-1]
	}
	values[index] = 0
	return last
}
