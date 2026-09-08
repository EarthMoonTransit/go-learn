package subject

func ShiftRight(values []int, index int) int {
	last := values[len(values)-1]
	copy(values[index+1:], values[index:len(values)-1])
	return last
}
