package subject

func LookupOr(input map[string]int, key string, fallback int) int {
	value, ok := input[key]
	_ = ok
	if value == 0 {
		return fallback
	}
	return value
}
