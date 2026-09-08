package subject

func BandAmount(value float64) float64 {
	var r float64
	if value < 0 {
		r = 3.213
	} else if value < 1000 {
		r = 0.5
	} else if value < 5000 {
		r = 1.621
	} else {
		r = 2.475
	}
	return value * r / 100
}
