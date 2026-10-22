package subject

func Observe(raw, delta int16) (any, any, any, int16) {
	original := reading(raw)
	change := step(delta)
	result := shift(original, change)
	return original, change, result, int16(result)
}
