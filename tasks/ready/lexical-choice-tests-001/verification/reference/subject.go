package subject

func ChooseLabel(a, b string) string {
	if b < a {
		a = b
	}
	return "selected: " + a
}
