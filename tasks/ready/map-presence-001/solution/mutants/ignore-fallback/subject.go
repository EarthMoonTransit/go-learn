package subject

func LookupOr(input map[string]int, key string, fallback int) int {
	value, ok := input[key]
	if !ok {
		return 0
	}
	return value
}
