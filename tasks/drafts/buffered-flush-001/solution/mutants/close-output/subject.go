package subject

import (
	"bufio"
	"io"
)

func WriteBuffered(dst io.Writer, text string) error {
	if c, ok := dst.(io.Closer); ok {
		defer c.Close()
	}
	writer := bufio.NewWriter(dst)
	if _, err := writer.WriteString(text); err != nil {
		return err
	}
	return writer.Flush()
}
