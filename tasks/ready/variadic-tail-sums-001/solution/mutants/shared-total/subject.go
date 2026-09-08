package subject

func SumTails(groups ...[]int) []int {
	out := []int{}
	sum := 0
	for _, g := range groups {
		for i, n := range g {
			if i > 0 {
				sum += n
			}
		}
		out = append(out, sum)
	}
	return out
}
