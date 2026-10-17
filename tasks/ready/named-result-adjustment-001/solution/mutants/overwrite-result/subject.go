package subject

func Adjust(base, delta int, caller *int, work func()) (result int) {
	defer func() { result = delta }()
	work()
	return base + 3
}
