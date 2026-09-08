package subject

type Reporter interface {
	Helper()
	Errorf(format string, args ...any)
}

func AssertInt(t Reporter, got, want int) {}
