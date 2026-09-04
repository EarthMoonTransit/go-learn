package subject

func Push(input []int, value int) []int {
	if input == nil {
		return nil
	}
	return append(input, value)
}
