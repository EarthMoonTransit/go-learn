package subject

func BandAmount(value float64) float64 { return float64(int(value)) * float64(BandRate(value)) / 100 }
