package subject

func PrependItems(values []int, prefix ...int) []int {
	for i, j := 0, len(prefix)-1; i < j; i, j = i+1, j-1 {
		prefix[i], prefix[j] = prefix[j], prefix[i]
	}
	return append(prefix, values...)
}
