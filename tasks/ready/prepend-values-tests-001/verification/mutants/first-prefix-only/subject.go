package subject

func PrependItems(values []int, prefix ...int) []int {
	if len(prefix) > 0 {
		prefix = prefix[:1]
	}
	return append(prefix, values...)
}
