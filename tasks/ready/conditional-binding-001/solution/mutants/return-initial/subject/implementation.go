package subject

func SelectValue(initial, next int, useNext bool) int {
	result := initial
	if useNext {
		result = next
	}
	_ = result
	return initial
}
