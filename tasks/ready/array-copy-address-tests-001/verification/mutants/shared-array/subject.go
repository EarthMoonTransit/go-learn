package subject

var shared [5]string

func CopyArray(source *[5]string) (*[5]string, [5]*string) {
	shared = *source
	var addresses [5]*string
	for i := range shared {
		addresses[i] = &shared[i]
	}
	return &shared, addresses
}
