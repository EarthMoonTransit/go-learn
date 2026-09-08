package subject

func CorrectAlternating(counts []int) []int {
	for i := range counts {
		if i%2 == 0 {
			counts[i]++
		}
	}
	return counts
}
