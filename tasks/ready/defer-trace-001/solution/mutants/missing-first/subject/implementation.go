package subject

func runTrace(seed int, record func(int), checkpoint func()) {
	seed += 4
	defer record(seed)
	seed += 5
	record(seed)
	checkpoint()
}
