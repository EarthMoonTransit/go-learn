package subject

import (
	"slices"
)

type Snapshot struct {
	Label string
	Count int
	Items []int
}

func CopySnapshot(input Snapshot) Snapshot { input.Items = slices.Clone(input.Items); return input }
