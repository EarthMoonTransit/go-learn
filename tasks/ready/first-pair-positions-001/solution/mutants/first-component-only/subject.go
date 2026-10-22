package subject

func FirstPositions(pairs [][2]string) []int {
	out := make([]int, len(pairs))
	first := make(map[string]int)
	for i, pair := range pairs {
		position, ok := first[pair[0]]
		if !ok {
			position = i
			first[pair[0]] = i
		}
		out[i] = position
	}
	return out
}
