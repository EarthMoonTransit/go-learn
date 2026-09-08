package subject

func ChooseLabel(a, b string) string {
	if a == b {
		return "selected: "
	}
	if b < a {
		a = b
	}
	return "selected: " + a
}
