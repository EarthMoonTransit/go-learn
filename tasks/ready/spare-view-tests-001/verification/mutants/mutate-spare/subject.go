package subject

func SpareView(in []int, low, high, spare int) []int {
	all := in[:cap(in)]
	for i := high; i < high+spare; i++ {
		all[i] = 0
	}
	return in[low : high : high+spare]
}
