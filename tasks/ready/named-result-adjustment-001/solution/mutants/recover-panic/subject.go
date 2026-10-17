package subject

func Adjust(base, delta int, caller *int, work func()) (result int) {
	defer func() { result += delta }()
	defer func() { recover() }()
	work()
	return base + 3
}
