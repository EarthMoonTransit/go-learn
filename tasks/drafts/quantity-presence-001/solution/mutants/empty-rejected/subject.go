package subject

func Quantity(values map[string]int, key string) (int, bool) {
	if key == "" {
		return 0, false
	}
	n, ok := values[key]
	return n, ok
}
