package subject

func runTrace(seed int, record func(int), checkpoint func()) {
	record(seed + 9)
	checkpoint()
	record(seed + 4)
	record(seed)
}
