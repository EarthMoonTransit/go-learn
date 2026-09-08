package subject

func RuneOffsets(s string) []int {
	out := []int{}
	for i := range s {
		if s[i] >= 128 {
			out = append(out, i)
		}
	}
	return out
}
