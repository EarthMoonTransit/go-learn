package subject

func RuneOffsets(s string) []int {
	out := []int{}
	for i := range s {
		out = append(out, i)
	}
	return out
}
