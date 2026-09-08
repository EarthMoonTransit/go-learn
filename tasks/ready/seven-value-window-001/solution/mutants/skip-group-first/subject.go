package subject

func WeekSum(counts []int, week int) int {
	total := 0
	for _, n := range counts[(week-1)*7+1 : week*7] {
		total += n
	}
	return total
}
