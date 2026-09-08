package subject

func Find[T any](items []T, predicate func(T) bool) (T, bool) {
	for _, v := range items {
		switch any(v).(type) {
		case int, string:
		default:
			continue
		}
		if predicate(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}
