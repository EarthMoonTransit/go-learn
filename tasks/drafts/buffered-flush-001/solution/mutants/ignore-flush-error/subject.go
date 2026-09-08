package subject

import (
	"bufio"
	"io"
)

func WriteBuffered(dst io.Writer, text string) error {
	writer := bufio.NewWriter(dst)
	if _, err := writer.WriteString(text); err != nil {
		return err
	}
	writer.Flush()
	return nil
}
