package subject

func SuccessfulPerHour(rate int, percent float64) float64 {
	return float64(int(float64(rate) * percent / 100))
}
