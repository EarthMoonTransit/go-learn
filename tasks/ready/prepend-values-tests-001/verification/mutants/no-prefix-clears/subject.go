package subject

func PrependItems(values []int, prefix ...int) []int {
	if len(prefix) == 0 {
		return nil
	}
	return append(prefix, values...)
}
