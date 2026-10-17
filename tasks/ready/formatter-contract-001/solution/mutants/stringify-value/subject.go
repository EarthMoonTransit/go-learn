package subject

import "fmt"

type Formatter interface {
	Format() string
}
type Pair struct {
	Left  string
	Right string
}

func (pair Pair) Format() string {
	return pair.Left + ":" + pair.Right
}

func Render(value Formatter) string {
	return fmt.Sprint(value)
}
