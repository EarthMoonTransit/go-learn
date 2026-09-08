package subject

import (
	"io"
)

func WriteGreeting(w io.Writer, name string) error {
	text := []byte("Hello, " + name)
	n, e := w.Write(text)
	if e != nil {
		return nil
	}
	if n != len(text) {
		return io.ErrShortWrite
	}
	return nil
}
