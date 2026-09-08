package subject

import (
	"bufio"
)

func WriteBuffered(writer *bufio.Writer, text string) error {
	if _, err := writer.WriteString(text); err != nil {
		return err
	}
	writer.Flush()
	return nil
}
