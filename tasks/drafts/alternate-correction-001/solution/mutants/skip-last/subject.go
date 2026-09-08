package subject

func CorrectAlternating(counts []int) []int {
	for i := range counts {
		if i%2 == 0 && i < len(counts)-1 {
			counts[i]++
		}
	}
	return counts
}
