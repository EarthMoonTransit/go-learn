package subject

func EffectiveRate(itemsPerHour int, successPercent float64) int {
	return (itemsPerHour / 100) * int(successPercent)
}
