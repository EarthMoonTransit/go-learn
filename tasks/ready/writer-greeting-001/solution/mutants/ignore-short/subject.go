package subject

import (
	"io"
)

func WriteGreeting(w io.Writer, name string) error {
	text := []byte("Hello, " + name)
	n, e := w.Write(text)
	if e != nil {
		return e
	}
	_ = n
	return nil
}
