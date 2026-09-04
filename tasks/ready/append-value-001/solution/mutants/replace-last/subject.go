package subject

func Push(input []int, value int) []int {
	if len(input) == 0 {
		return []int{value}
	}
	input[len(input)-1] = value
	return input
}
