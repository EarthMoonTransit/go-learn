package subject

func WeekSum(counts []int, week int) int {
	total := 0
	for _, n := range counts {
		total += n
	}
	return total
}
