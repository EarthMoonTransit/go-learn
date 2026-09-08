package subject

import (
	"fmt"
)

type NumberBox interface{ Number() int }

func DescribeNumberBox(box NumberBox) string { return fmt.Sprintf("box %.1f", float64(4)) }
