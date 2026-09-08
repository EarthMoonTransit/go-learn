package subject

func UpperLabel(text string) string {
	out := []rune(text)
	for i, r := range out {
		if r >= 'a' && r <= 'z' {
			out[i] = r - 32
		}
	}
	return "INPUT: " + string(out)
}
