package subject

func Transfer(dst, src *int, limit int) int {
	if dst == nil || src == nil {
		if dst != nil {
			*dst = 0
		}
		if src != nil {
			*src = 0
		}
		return 0
	}
	if dst == src || limit <= 0 {
		return 0
	}
	amount := limit
	if *src < amount {
		amount = *src
	}
	*src -= amount
	*dst += amount
	return amount
}
