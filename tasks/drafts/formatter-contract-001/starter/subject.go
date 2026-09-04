package subject

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
	if pair, ok := value.(Pair); ok {
		return pair.Format()
	}
	return ""
}
