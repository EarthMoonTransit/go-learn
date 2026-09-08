package subject

func CountColumn(grid map[string][]bool, key string) int {
	if _, ok := grid[key]; !ok {
		return 1
	}
	total := 0
	for _, occupied := range grid[key] {
		if occupied {
			total++
		}
	}
	return total
}
