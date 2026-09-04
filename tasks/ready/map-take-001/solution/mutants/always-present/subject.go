package subject

func Take(values map[string]int, key string) (int, bool) {
	value := values[key]
	delete(values, key)
	return value, true
}
