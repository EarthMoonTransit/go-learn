package subject

func FirstPositions(pairs [][2]string) []int {
	out := make([]int, len(pairs))
	first := make(map[*[2]string]int)
	for i := range pairs {
		position, ok := first[&pairs[i]]
		if !ok {
			position = i
			first[&pairs[i]] = i
		}
		out[i] = position
	}
	return out
}
