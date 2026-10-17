package subject

func runTrace(seed int, record func(int), checkpoint func()) {
	defer func() { record(seed) }()
	seed += 4
	defer func() { record(seed) }()
	seed += 5
	record(seed)
	checkpoint()
}
