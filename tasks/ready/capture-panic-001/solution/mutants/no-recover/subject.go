package subject

func Capture(work func() error) (error, bool) { return work(), false }
