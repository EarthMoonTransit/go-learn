package subject

func CorrectAlternating(counts []int) []int {
	if len(counts) == 0 {
		return []int{0}
	}
	for i := range counts {
		if i%2 == 0 {
			counts[i]++
		}
	}
	return counts
}
