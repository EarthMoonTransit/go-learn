package subject

import (
	"bufio"
)

func WriteBuffered(writer *bufio.Writer, text string) error {
	if text == "" {
		return nil
	}
	if _, err := writer.WriteString(text); err != nil {
		return err
	}
	return writer.Flush()
}
