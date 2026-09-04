package transform

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

func NormalizeLines(reader io.Reader, writer io.Writer) error {
	data, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		output := strings.ToUpper(line) + "\n"
		written, writeErr := io.WriteString(writer, output)
		if writeErr != nil {
			return writeErr
		}
		if written != len(output) {
			return io.ErrShortWrite
		}
	}
	return scanner.Err()
}
