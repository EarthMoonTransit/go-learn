package subject

import (
	"io"
)

func WriteGreeting(w io.Writer, name string) error {
	text := []byte("Hello, " + name)
	n, e := w.Write(text)
	if e != nil {
		w.Write(text)
		return e
	}
	if n != len(text) {
		return io.ErrShortWrite
	}
	return nil
}
