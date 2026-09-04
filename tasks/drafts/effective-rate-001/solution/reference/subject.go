package subject

func EffectiveRate(itemsPerHour int, successPercent float64) int {
	return int(float64(itemsPerHour) * successPercent / 100)
}
