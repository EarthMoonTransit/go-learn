package subject

func PrependItems(values []int, prefix ...int) []int {
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
	return append(prefix, values...)
}
