package subject

func AddCounter(counter *int, delta int) {
	if delta < 0 {
		delta = -delta
	}
	*counter += delta
}
