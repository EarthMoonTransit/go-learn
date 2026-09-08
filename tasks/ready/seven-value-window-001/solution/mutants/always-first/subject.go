package subject

func WeekSum(counts []int, week int) int {
	total := 0
	for _, n := range counts[:7] {
		total += n
	}
	return total
}
