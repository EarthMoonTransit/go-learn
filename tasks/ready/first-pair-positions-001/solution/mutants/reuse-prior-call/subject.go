package subject

var remembered = make(map[[2]string]int)

func FirstPositions(pairs [][2]string) []int {
	out := make([]int, len(pairs))
	first := remembered
	for i, pair := range pairs {
		position, ok := first[pair]
		if !ok {
			position = i
			first[pair] = i
		}
		out[i] = position
	}
	return out
}
