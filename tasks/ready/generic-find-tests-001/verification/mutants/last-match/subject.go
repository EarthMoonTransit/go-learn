package subject

func Find[T any](items []T, predicate func(T) bool) (T, bool) {
	var value T
	found := false
	for _, v := range items {
		if predicate(v) {
			value = v
			found = true
		}
	}
	return value, found
}
