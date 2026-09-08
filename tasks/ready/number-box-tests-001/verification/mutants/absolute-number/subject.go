package subject

import (
	"fmt"
	"math"
)

type NumberBox interface{ Number() int }

func DescribeNumberBox(box NumberBox) string {
	return fmt.Sprintf("box %.1f", math.Abs(float64(box.Number())))
}
