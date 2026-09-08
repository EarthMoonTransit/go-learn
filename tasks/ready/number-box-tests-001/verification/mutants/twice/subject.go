package subject

import (
	"fmt"
)

type NumberBox interface{ Number() int }

func DescribeNumberBox(box NumberBox) string {
	box.Number()
	return fmt.Sprintf("box %.1f", float64(box.Number()))
}
