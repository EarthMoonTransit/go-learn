package variadicsum

func Sum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func SumSlice(values []int) int {
	return Sum(values...)
}
