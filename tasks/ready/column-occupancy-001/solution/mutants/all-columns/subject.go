package subject

func CountColumn(grid map[string][]bool, key string) int {
	n := 0
	for _, col := range grid {
		for _, v := range col {
			if v {
				n++
			}
		}
	}
	return n
}
