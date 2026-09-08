package subject

type numeric interface{ ~int | ~float64 }

func Double[T numeric](value T) T { return value + value }
