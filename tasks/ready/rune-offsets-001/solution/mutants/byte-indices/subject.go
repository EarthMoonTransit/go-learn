package subject

func RuneOffsets(s string) []int {
	out := []int{}
	for i := 0; i < len(s); i++ {
		out = append(out, i)
	}
	return out
}
