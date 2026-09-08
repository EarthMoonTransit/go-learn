package subject

func ReplaceFirst(replacement int, values ...int) bool {
	if replacement == 0 || len(values) == 0 {
		return false
	}
	values[0] = replacement
	return true
}
