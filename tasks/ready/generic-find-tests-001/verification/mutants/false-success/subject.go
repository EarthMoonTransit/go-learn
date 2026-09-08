package subject

func Find[T any](items []T, predicate func(T) bool) (T, bool) {
	for _, v := range items {
		if predicate(v) {
			return v, false
		}
	}
	var zero T
	return zero, false
}
