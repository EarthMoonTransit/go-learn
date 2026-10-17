package subject

func CheckEnabled(enabled bool) bool {
	return !enabled && Probe()
}
