package subject

func Gather(source [8]int, indexes [8]int) ([8]int, bool) {
	for _, index := range indexes {
		if index < 0 || index >= len(source) {
			return source, false
		}
	}
	result := source
	for i, index := range indexes {
		result[i] = result[index]
	}
	return result, true
}
