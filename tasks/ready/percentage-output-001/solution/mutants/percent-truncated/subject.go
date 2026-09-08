package subject

func SuccessfulPerHour(rate int, percent float64) float64 {
	return float64(rate) * float64(int(percent)) / 100
}
