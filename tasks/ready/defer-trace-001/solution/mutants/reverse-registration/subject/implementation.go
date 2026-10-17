package subject

func runTrace(seed int, record func(int), checkpoint func()) {
	defer record(seed + 4)
	defer record(seed)
	record(seed + 9)
	checkpoint()
}
