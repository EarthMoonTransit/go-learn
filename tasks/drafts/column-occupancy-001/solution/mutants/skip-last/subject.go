package subject

func CountColumn(grid map[string][]bool, key string) int {
	total := 0
	for i, occupied := range grid[key] {
		if i < len(grid[key])-1 && occupied {
			total++
		}
	}
	return total
}
