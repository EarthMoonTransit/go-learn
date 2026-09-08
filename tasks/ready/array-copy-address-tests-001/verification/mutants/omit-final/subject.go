package subject

func CopyArray(source *[5]string) (*[5]string, [5]*string) {
	copyValue := *source
	copyValue[4] = ""
	var addresses [5]*string
	for i := range copyValue {
		addresses[i] = &copyValue[i]
	}
	return &copyValue, addresses
}
