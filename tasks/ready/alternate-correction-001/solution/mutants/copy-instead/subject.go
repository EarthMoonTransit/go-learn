package subject

func CorrectAlternating(counts []int) []int {
	counts = append([]int(nil), counts...)
	for i := range counts {
		if i%2 == 0 {
			counts[i]++
		}
	}
	return counts
}
