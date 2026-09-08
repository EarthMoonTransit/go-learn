package subject

import (
	"fmt"
)

type NumberBox interface{ Number() int }

var seen bool
var cached int

func DescribeNumberBox(box NumberBox) string {
	if !seen {
		cached = box.Number()
		seen = true
	}
	return fmt.Sprintf("box %.1f", float64(cached))
}
