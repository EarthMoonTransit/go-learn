package subject

func Adjust(base, delta int, caller *int, work func()) (result int) {
	work()
	return base + 3 + delta
}
