package subject

func Take(values map[string]int, key string) (int, bool) {
	value, ok := values[key]
	if !ok {
		return 0, false
	}
	clear(values)
	return value, true
}
