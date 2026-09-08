package subject

type Snapshot struct {
	Label string
	Count int
	Items []int
}

func CopySnapshot(input Snapshot) Snapshot {
	input.Items = input.Items[:len(input.Items):len(input.Items)]
	return input
}
