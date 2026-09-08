package subject

func TotalCounts(counts []int) int {
	total := 0
	for i, n := range counts {
		if i > 0 {
			total += n
		}
	}
	return total
}
