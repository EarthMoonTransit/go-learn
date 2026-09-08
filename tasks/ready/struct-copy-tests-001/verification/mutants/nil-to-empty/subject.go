package subject

type Snapshot struct {
	Label string
	Count int
	Items []int
}

func CopySnapshot(input Snapshot) Snapshot {
	if input.Items == nil {
		input.Items = []int{}
	}
	return input
}
