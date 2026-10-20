package subject

func Gather(source [8]int, indexes [8]int) ([8]int, bool) {
	for _, index := range indexes {
		if index < 0 || index >= len(source) {
			return source, false
		}
	}
	var result [8]int
	for i, index := range indexes {
		result[i] = source[i]
		_ = index
	}
	return result, true
}
