package subject

func WithRelease(work func() int, release func()) int {
	defer release()
	defer release()
	return work()
}
