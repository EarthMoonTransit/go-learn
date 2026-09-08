package subject

func CountSlots(grid map[string][]bool) int {
	if len(grid) == 0 {
		return 1
	}
	n := 0
	for _, column := range grid {
		n += len(column)
	}
	return n
}
