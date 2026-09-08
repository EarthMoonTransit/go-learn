package subject

import (
	"bufio"
	"io"
	"strings"
)

func ScanLines(src io.Reader) ([]string, error) {
	scan := bufio.NewScanner(src)
	scan.Buffer(make([]byte, 128), 4096)
	var lines []string
	for scan.Scan() {
		lines = append(lines, strings.TrimSpace(scan.Text()))
	}
	return lines, scan.Err()
}
