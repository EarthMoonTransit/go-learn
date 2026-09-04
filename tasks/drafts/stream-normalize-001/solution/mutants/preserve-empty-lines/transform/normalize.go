package transform

import (
	"bufio"
	"io"
	"strings"
)

func NormalizeLines(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		output := strings.ToUpper(strings.TrimSpace(scanner.Text())) + "\n"
		written, err := io.WriteString(writer, output)
		if err != nil {
			return err
		}
		if written != len(output) {
			return io.ErrShortWrite
		}
	}
	return scanner.Err()
}
