package subject

func runTrace(seed int, record func(int), checkpoint func()) {
	record(seed)
	seed += 4
	record(seed)
	seed += 5
	record(seed)
	checkpoint()
}
