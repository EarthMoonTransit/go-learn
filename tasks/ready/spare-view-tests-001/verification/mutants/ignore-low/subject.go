package subject

func SpareView(in []int, low, high, spare int) []int { return in[: high-low : high-low+spare] }
