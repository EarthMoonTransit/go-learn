package subject

func TotalCounts(counts []int) int {
	if len(counts) == 0 {
		return 1
	}
	total := 0
	for _, n := range counts {
		total += n
	}
	return total
}
