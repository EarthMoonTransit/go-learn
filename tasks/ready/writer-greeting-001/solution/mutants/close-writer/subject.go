package subject

import (
	"io"
)

func WriteGreeting(w io.Writer, name string) error {
	if c, ok := w.(interface{ Close() error }); ok {
		c.Close()
	}
	text := []byte("Hello, " + name)
	n, e := w.Write(text)
	if e != nil {
		return e
	}
	if n != len(text) {
		return io.ErrShortWrite
	}
	return nil
}
