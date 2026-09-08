package subject

func RuneOffsets(s string) []int {
	out := []int{}
	for i := range s {
		out = append(out, i)
	}
	if len(out) > 0 {
		return out[:len(out)-1]
	}
	return out
}
