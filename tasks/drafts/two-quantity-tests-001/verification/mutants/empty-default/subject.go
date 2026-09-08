package subject

func Quantities(kinds []string) (int, float64) {
	if len(kinds) == 0 {
		return 50, 0.2
	}
	a, b := 0, 0
	for _, k := range kinds {
		if k == "solid" {
			a++
		}
		if k == "liquid" {
			b++
		}
	}
	return a * 50, float64(b) * 0.2
}
