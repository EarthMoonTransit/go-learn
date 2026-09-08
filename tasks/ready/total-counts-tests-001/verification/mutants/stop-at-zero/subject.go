package subject

func TotalCounts(counts []int) int {
	total := 0
	for _, n := range counts {
		if n == 0 {
			break
		}
		total += n
	}
	return total
}
