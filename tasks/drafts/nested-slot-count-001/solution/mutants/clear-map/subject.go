package subject

func CountSlots(grid map[string][]bool) int {
	n := 0
	for _, column := range grid {
		n += len(column)
	}
	clear(grid)
	return n
}
