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
		lines = append(lines, scan.Text())
	}
	if len(lines) > 0 {
		lines = lines[:len(lines)-1]
	}
	return lines, scan.Err()
}
