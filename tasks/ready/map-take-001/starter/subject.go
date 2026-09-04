package subject

func Take(values map[string]int, key string) (int, bool) {
	value := values[key]
	if value == 0 {
		return 0, false
	}
	delete(values, key)
	return value, true
}
