package subject

func SumTails(groups ...[]int) []int {
	out := []int{}
	for _, g := range groups {
		if len(g) == 0 {
			continue
		}
		sum := 0
		for i, n := range g {
			if i > 0 {
				sum += n
			}
		}
		out = append(out, sum)
	}
	return out
}
