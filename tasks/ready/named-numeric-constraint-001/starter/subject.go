package subject

type numeric interface{ int | int64 | float64 }

func Double[T numeric](value T) T { return value + value }
