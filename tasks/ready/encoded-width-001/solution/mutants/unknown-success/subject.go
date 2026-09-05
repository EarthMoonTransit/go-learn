package subject

func Width(kind string) (int, bool) {
	switch kind {
	case "byte":
		return 1, true
	case "word":
		return 2, true
	case "long":
		return 4, true
	case "quad":
		return 8, true
	default:
		return 0, true
	}
}
