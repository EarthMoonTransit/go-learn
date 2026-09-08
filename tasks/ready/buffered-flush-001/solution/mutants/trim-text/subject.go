package subject

import (
	"bufio"
	"strings"
)

func WriteBuffered(writer *bufio.Writer, text string) error {
	if _, err := writer.WriteString(strings.TrimSpace(text)); err != nil {
		return err
	}
	return writer.Flush()
}
