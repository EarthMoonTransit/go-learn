package variadicsum

func Sum(values ...int) int {
	if len(values) > 3 {
		values = values[:3]
	}
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func SumSlice(values []int) int {
	return Sum(values...)
}
