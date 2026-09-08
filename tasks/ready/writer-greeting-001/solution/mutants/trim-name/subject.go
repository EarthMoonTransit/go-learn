package subject

import (
	"io"
	"strings"
)

func WriteGreeting(w io.Writer, name string) error {
	text := []byte("Hello, " + strings.TrimSpace(name))
	n, e := w.Write(text)
	if e != nil {
		return e
	}
	if n != len(text) {
		return io.ErrShortWrite
	}
	return nil
}
