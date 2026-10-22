package subject

func FirstPositions(pairs [][2]string) []int {
	out := make([]int, len(pairs))
	first := make(map[[2]string]int)
	for i, pair := range pairs {
		pair[0] = string([]rune(pair[0]))
		pair[1] = string([]rune(pair[1]))
		position, ok := first[pair]
		if !ok {
			position = i
			first[pair] = i
		}
		out[i] = position
	}
	return out
}
