package subject

func Quantity(values map[string]int, key string) (int, bool) {
	n, ok := values[key]
	return n, ok && n != 0
}
