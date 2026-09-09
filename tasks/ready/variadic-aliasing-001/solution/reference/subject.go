package subject

func ReplaceFirst(replacement int, values ...int) bool {
	if len(values) == 0 {
		return false
	}
	values[0] = replacement
	return true
}
