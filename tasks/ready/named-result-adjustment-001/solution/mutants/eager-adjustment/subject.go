package subject

func Adjust(base, delta int, caller *int, work func()) (result int) {
	result += delta
	work()
	return base + 3
}
