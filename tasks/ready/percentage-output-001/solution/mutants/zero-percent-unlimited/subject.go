package subject

func SuccessfulPerHour(rate int, percent float64) float64 {
	if percent == 0 {
		return float64(rate)
	}
	return float64(rate) * percent / 100
}
