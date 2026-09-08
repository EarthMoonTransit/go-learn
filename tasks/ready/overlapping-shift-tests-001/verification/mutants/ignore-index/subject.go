package subject

func ShiftRight(values []int, index int) int {
	index = 0
	last := values[len(values)-1]
	copy(values[index+1:], values[index:len(values)-1])
	values[index] = 0
	return last
}
