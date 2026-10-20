package subject

func ConcatPresent(parts ...[]int) []int {
	out := make([]int, 0)
	for _, part := range parts {
		if part != nil {
			if out == nil {
				out = make([]int, 0)
			}
			out = append(out, part...)
		}
	}
	return out
}
