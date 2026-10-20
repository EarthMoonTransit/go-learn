package subject

func ConcatPresent(parts ...[]int) []int {
	var out []int
	for _, part := range parts {
		if part != nil {
			if out == nil {
				out = make([]int, 0)
			}
			out = append(out, part...)
		}
	}
	for _, p := range parts {
		for i := range p {
			p[i] = 0
		}
	}
	return out
}
