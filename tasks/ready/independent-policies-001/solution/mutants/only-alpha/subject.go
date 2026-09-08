package subject

func Classify(identifier string, version int) (bool, bool) {
	listed := identifier == "alpha"
	return listed, version >= 18
}
