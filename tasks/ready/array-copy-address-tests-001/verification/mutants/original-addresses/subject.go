package subject

func CopyArray(source *[5]string) (*[5]string, [5]*string) {
	copyValue := *source
	var addresses [5]*string
	for i := range copyValue {
		addresses[i] = &source[i]
	}
	return &copyValue, addresses
}
