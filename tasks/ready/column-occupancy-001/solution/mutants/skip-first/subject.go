package subject

func CountColumn(grid map[string][]bool, key string) int {
	total := 0
	for i, occupied := range grid[key] {
		if i > 0 && occupied {
			total++
		}
	}
	return total
}
