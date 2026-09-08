package subject

func BandRate(value float64) float32 {
	if value < 0 {
		return 3.213
	}
	if value < 1000 {
		return 0.5
	}
	if value < 5000 {
		return 1
	}
	return 2.475
}
