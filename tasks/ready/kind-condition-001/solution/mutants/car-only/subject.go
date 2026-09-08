package subject

func RequiresPermit(kind string) bool {
	if kind == "car" {
		return true
	}
	return false
}
