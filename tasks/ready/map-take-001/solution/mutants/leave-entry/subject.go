package subject

func Take(values map[string]int, key string) (int, bool) {
	value, ok := values[key]
	return value, ok
}
