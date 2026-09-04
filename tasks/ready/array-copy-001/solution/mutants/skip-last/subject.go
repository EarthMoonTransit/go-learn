package subject

func Replaced(input [3]int, index int, value int) [3]int {
	if index >= 0 && index < len(input)-1 {
		input[index] = value
	}
	return input
}
