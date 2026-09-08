package subject

func RequiresPermit(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}
