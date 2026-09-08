package subject

func RuneOffsets(s string) []int {
	if s == "" {
		return []int{0}
	}
	out := []int{}
	for i := range s {
		out = append(out, i)
	}
	return out
}
