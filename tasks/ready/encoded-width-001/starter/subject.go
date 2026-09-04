package subject

func Width(kind string) (int, bool) {
	width := 0
	switch kind {
	case "byte":
		width = 1
		fallthrough
	case "word":
		width = 2
	case "long":
		width = 4
	case "quad":
		width = 8
	default:
		return 0, false
	}
	return width, true
}
