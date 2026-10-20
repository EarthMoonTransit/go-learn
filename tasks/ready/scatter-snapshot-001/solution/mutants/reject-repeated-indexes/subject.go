package subject

func Gather(source [8]int, indexes [8]int) ([8]int, bool) {
	for _, index := range indexes {
		if index < 0 || index >= len(source) {
			return source, false
		}
	}
	for i := range indexes {
		for j := 0; j < i; j++ {
			if indexes[i] == indexes[j] {
				return source, false
			}
		}
	}
	var result [8]int
	for i, index := range indexes {
		result[i] = source[index]
	}
	return result, true
}
