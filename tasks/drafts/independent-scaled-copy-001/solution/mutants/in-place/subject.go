package subject

func ScaledCopy(values []float64, units int) []float64 {
	out := values
	for i, n := range values {
		out[i] = n * float64(units) / 2
	}
	return out
}
