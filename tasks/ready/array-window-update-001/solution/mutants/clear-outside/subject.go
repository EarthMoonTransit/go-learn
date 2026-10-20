package subject

func PatchWindow(a [8]int, lo, hi, value int) ([8]int, bool) {
	if lo < 0 || lo > hi || hi > len(a) {
		return a, false
	}
	a = [8]int{}
	for i := lo; i < hi; i++ {
		a[i] = value
	}
	return a, true
}
