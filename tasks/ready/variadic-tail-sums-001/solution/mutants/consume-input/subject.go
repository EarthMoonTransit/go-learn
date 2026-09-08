package subject

func SumTails(groups ...[]int) []int {
	out := []int{}
	for _, g := range groups {
		sum := 0
		for i, n := range g {
			if i > 0 {
				sum += n
			}
		}
		clear(g)
		out = append(out, sum)
	}
	return out
}
