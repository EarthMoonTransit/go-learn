package subject

func Classify(identifier string, version int) (bool, bool) {
	listed := identifier == "alpha" || identifier == "beta"
	return listed, listed && version >= 18
}
