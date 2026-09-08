package subject

import (
	"bufio"
	"io"
	"strings"
)

func WriteBuffered(dst io.Writer, text string) error {
	writer := bufio.NewWriter(dst)
	if _, err := writer.WriteString(strings.TrimSpace(text)); err != nil {
		return err
	}
	return writer.Flush()
}
