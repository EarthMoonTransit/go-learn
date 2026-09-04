package transform

import (
	"bufio"
	"io"
	"strings"
)

func NormalizeLines(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		output := strings.ToUpper(line) + "\n"
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
