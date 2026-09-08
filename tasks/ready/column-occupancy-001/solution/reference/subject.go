package subject

func CountColumn(grid map[string][]bool, key string) int {
	total := 0
	for _, occupied := range grid[key] {
		if occupied {
			total++
		}
	}
	return total
}
