package subject

func RuneOffsets(s string) []int {
	out := []int{}
	for range s {
		out = append(out, len(out))
	}
	return out
}
