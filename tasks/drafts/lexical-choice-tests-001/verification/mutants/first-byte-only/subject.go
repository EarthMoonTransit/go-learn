package subject

func ChooseLabel(a, b string) string {
	if len(a) > 0 && len(b) > 0 && b[0] < a[0] {
		a = b
	}
	return "selected: " + a
}
