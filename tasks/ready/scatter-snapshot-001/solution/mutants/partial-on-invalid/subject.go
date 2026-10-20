package subject

func Gather(source [8]int, indexes [8]int) ([8]int, bool) {
	result := source
	for i, j := range indexes {
		if j < 0 || j >= 8 {
			return result, false
		}
		result[i] = source[j]
	}
	return result, true
}
