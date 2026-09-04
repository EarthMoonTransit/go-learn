package subject

func OffsetByIndex(values []int) {
	for i := range values {
		values[i] += i
	}
}
