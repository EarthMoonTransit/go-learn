package subject

func SpareView(in []int, low, high, spare int) []int { return in[low : high+spare : high+spare] }
