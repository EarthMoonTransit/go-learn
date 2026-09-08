package subject

import "strings"

func CountColumn(grid map[string][]bool, key string) int {
	total := 0
	for _, occupied := range grid[strings.ToUpper(key)] {
		if occupied {
			total++
		}
	}
	return total
}
