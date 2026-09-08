package subject

func WithRelease(work func() int, release func()) int { n := work(); release(); return n }
