package subject

type numeric interface{ ~int | ~int64 | ~float64 }

func Double[T numeric](value T) T {
	if value < 0 {
		return -value - value
	}
	return value + value
}
