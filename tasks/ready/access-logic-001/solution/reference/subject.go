package subject

func Allowed(enabled, blocked, override bool) bool {
	return override || (enabled && !blocked)
}
