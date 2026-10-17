package subject

func CheckEnabled(enabled bool) bool {
	value := Probe()
	return enabled && value
}
