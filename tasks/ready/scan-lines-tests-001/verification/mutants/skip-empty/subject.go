package subject

import (
	"bufio"
	"io"
)

func ScanLines(src io.Reader) ([]string, error) {
	scan := bufio.NewScanner(src)
	scan.Buffer(make([]byte, 128), 4096)
	var lines []string
	for scan.Scan() {
		if scan.Text() != "" {
			lines = append(lines, scan.Text())
		}
	}
	return lines, scan.Err()
}
