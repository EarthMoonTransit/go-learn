package subject

func ScaledCopy(values []float64, units int) []float64 {
	if units == 2 {
		return values
	}
	out := make([]float64, len(values))
	for i, n := range values {
		out[i] = n * float64(units) / 2
	}
	return out
}
