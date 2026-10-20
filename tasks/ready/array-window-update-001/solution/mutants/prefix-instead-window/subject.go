package subject

func PatchWindow(a [8]int, lo, hi, value int) ([8]int, bool) {
	if lo < 0 || lo > hi || hi > len(a) {
		return a, false
	}
	for i := 0; i < hi-lo; i++ {
		a[i] = value
	}
	return a, true
}
