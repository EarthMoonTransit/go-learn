package subject

func Quantity(values map[string]int, key string) (int, bool) {
	n, ok := values[key]
	delete(values, key)
	return n, ok
}
