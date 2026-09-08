package subject

func SuccessfulPerMinute(rate int, percent float64) int {
	if percent == 0 {
		return rate / 60
	}
	return int(float64(rate) * percent / 100 / 60)
}
