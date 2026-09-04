package subject

func Push(input []int, value int) []int {
	_ = append(input, value)
	return input
}
