package subject

func ArrayStages(a, b, c, replacement, argumentValue int) (source, assigned, insideCall, afterCall [3]int) {

	source = [3]int{a, b, c}
	working := source
	working[0] = replacement
	assigned = working
	insideCall = changeSecond(&working, argumentValue)
	afterCall = working
	return
}

func changeSecond(input *[3]int, value int) [3]int {
	input[1] = value
	return *input
}
