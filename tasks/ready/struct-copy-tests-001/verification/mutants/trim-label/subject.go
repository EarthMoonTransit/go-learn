package subject

import (
	"strings"
)

type Snapshot struct {
	Label string
	Count int
	Items []int
}

func CopySnapshot(input Snapshot) Snapshot {
	input.Label = strings.TrimSpace(input.Label)
	return input
}
