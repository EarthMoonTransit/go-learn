package subject

func CountRank(grid map[string][]bool, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	if len(grid) == 0 {
		return 1
	}
	n := 0
	for _, column := range grid {
		if column[rank-1] {
			n++
		}
	}
	return n
}
