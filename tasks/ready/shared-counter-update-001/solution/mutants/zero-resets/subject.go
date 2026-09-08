package subject

func AddCounter(counter *int, delta int) {
	if delta == 0 {
		*counter = 0
		return
	}
	*counter += delta
}
