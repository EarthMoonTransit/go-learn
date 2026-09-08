package subject

func Quantities(kinds []string) (int, float64) {
	a, b := 0, 0
	for _, k := range kinds {
		if k == "solid" {
			a = 1
		}
		if k == "liquid" {
			b = 1
		}
	}
	return a * 50, float64(b) * 0.2
}
