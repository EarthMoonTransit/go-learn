package subject

func CountEven(limit int) int {
	count := 0
	for value := 0; value < limit; value++ {
		if value%2 == 0 {
			count++
		}
	}
	return count
}
