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
		_, _ = io.WriteString(writer, strings.ToUpper(line)+"\n")
	}
	return scanner.Err()
}
