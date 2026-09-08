package subject

import (
	"bufio"
	"io"
)

func ScanLines(src io.Reader) ([]string, error) {
	scan := bufio.NewScanner(src)
	scan.Buffer(make([]byte, 128), 1024)
	var lines []string
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
	return lines, scan.Err()
}
