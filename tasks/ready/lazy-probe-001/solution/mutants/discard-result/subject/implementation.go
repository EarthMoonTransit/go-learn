package subject

func CheckEnabled(enabled bool) bool {
	if !enabled {
		return false
	}
	Probe()
	return true
}
