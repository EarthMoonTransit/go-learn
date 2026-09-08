package subject

func CorrectAlternating(counts []int) []int {
	for i, n := range counts {
		if i%2 == 0 {
			n++
			_ = n
		}
	}
	return counts
}
