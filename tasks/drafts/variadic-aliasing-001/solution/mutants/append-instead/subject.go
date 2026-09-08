package subject

func ReplaceFirst(replacement int, values ...int) bool {
	if len(values) == 0 {
		return false
	}
	values = append(values, replacement)
	return true
}
