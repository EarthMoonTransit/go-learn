package transform

import (
	"io"
	"strings"
)

func NormalizeLines(reader io.Reader, writer io.Writer) error {
	data, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	text := strings.ToUpper(strings.TrimSpace(string(data)))
	_, err = io.WriteString(writer, text+"\n")
	return err
}
