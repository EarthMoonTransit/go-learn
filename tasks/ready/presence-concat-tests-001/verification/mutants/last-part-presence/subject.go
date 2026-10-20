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
	if len(out) == 0 && len(parts) > 0 {
		if parts[len(parts)-1] == nil {
			return nil
		}
		return []int{}
	}
	return out
}
