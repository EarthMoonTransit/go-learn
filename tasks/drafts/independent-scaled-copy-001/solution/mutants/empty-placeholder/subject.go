package subject

func ScaledCopy(values []float64, units int) []float64 {
	if len(values) == 0 {
		return []float64{0}
	}
	out := make([]float64, len(values))
	for i, n := range values {
		out[i] = n * float64(units) / 2
	}
	return out
}
