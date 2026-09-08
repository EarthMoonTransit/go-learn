package subject

func PrependItems(values []int, prefix ...int) []int { return append(values, prefix...) }
