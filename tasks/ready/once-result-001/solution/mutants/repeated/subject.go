package subject

func Memoize(load func() (int, error)) func() (int, error) { return load }
