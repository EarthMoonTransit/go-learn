package subject

func TotalCounts(counts []int) int {
	total := 0
	for _, n := range counts {
		total += n
	}
	clear(counts)
	return total
}
