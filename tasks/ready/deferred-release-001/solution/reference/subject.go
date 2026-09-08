package subject

func WithRelease(work func() int, release func()) int { defer release(); return work() }
