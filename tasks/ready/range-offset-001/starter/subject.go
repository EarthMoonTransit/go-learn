package subject

func OffsetByIndex(values []int) {
	for i, value := range values {
		value += i + 1
	}
}
