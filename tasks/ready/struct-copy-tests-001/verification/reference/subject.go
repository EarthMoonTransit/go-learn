package subject

type Snapshot struct {
	Label string
	Count int
	Items []int
}

func CopySnapshot(input Snapshot) Snapshot { return input }
