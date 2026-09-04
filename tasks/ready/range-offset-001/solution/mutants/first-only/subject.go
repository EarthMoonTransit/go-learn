package subject

func OffsetByIndex(values []int) {
	for i := range values {
		if i == 0 {
			values[i] += i + 1
		}
	}
}
