package subject

func CountEven(limit int) int {
	count := 0
	for value := 0; value < limit; value += 3 {
		if value%2 == 0 {
			count++
		}
	}
	return count
}
