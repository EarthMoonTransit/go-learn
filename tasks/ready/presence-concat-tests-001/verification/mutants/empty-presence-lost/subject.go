package subject

func ConcatPresent(parts ...[]int) []int {
	var out []int
	for _, part := range parts {
		if part != nil {
			out = append(out, part...)
		}
	}
	return out
}
