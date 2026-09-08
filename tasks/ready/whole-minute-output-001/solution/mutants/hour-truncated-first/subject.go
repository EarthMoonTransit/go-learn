package subject

func SuccessfulPerMinute(rate int, percent float64) int { return int(float64(rate/60) * percent / 100) }
