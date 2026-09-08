package subject

func ReplaceFirst(replacement int, values ...int) bool {
	if len(values) == 0 {
		return false
	}
	values[len(values)-1] = replacement
	return true
}
