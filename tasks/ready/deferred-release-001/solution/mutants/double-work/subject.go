package subject

func WithRelease(work func() int, release func()) int { defer release(); work(); return work() }
