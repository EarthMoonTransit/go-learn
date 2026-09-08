package subject

func WithRelease(work func() int, release func()) int { release(); return work() }
