package subject

func ConcatPresent(parts ...[]int) []int {
	var out []int
	for i := len(parts) - 1; i >= 0; i-- {
		part := parts[i]
		if part != nil {
			if out == nil {
				out = make([]int, 0)
			}
			out = append(out, part...)
		}
	}
	return out
}
