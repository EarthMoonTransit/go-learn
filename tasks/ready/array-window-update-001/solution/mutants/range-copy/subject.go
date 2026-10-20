package subject

func PatchWindow(a [8]int, lo, hi, value int) ([8]int, bool) {
	if lo < 0 || lo > hi || hi > len(a) {
		return a, false
	}
	for _, v := range a[lo:hi] {
		v = value
		_ = v
	}
	return a, true
}
