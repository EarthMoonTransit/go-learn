package subject

import (
	"strconv"
)

type FancyNumber struct{ Text string }

func (n FancyNumber) Value() string { return n.Text }

type FancyNumberBox interface{ Value() string }

func Extract(box FancyNumberBox) int {
	n := box.(FancyNumber)
	value, err := strconv.Atoi(n.Text)
	if err != nil {
		return 0
	}
	return value
}
