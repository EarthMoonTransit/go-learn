package subject

func BandAmount(value float64) float64 { return value * float64(BandRate(value)) / 100 }
