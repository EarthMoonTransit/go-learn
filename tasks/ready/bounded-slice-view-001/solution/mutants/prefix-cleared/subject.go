package subject

func CappedView(in []int, start, end int) []int { clear(in[:start]); return in[start:end:end] }
