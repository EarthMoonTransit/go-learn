package subject

func TotalCounts(counts []int) int {
	if len(counts) == 1 {
		return 0
	}
	total := 0
	for _, n := range counts {
		total += n
	}
	return total
}
