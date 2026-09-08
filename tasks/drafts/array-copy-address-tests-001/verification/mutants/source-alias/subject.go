package subject

func CopyArray(source *[5]string) (*[5]string, [5]*string) {
	var addresses [5]*string
	for i := range source {
		addresses[i] = &source[i]
	}
	return source, addresses
}
