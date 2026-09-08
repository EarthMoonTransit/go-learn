package subject

func PrependItems(values []int, prefix ...int) []int {
	if len(values) == 0 {
		return values
	}
	return append(prefix, values...)
}
