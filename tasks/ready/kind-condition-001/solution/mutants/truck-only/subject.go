package subject

func RequiresPermit(kind string) bool {
	if kind == "truck" {
		return true
	}
	return false
}
