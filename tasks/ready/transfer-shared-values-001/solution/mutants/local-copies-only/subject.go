package subject

func Transfer(dst, src *int, limit int) int {
	if dst == nil || src == nil || dst == src || limit <= 0 {
		return 0
	}
	amount := limit
	if *src < amount {
		amount = *src
	}
	a, b := *src, *dst
	a -= amount
	b += amount
	_ = a
	_ = b
	return amount
}
