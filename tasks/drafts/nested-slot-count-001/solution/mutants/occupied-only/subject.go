package subject

func CountSlots(grid map[string][]bool) int {
	n := 0
	for _, c := range grid {
		for _, v := range c {
			if v {
				n++
			}
		}
	}
	return n
}
