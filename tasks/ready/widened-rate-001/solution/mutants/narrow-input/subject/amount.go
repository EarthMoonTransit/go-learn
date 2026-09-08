package subject

func BandAmount(value float64) float64 {
	return float64(float32(value)) * float64(BandRate(value)) / 100
}
