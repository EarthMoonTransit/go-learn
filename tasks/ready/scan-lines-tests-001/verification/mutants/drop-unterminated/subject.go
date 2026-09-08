package subject

import "bytes"

import (
	"bufio"
	"io"
)

func ScanLines(src io.Reader) ([]string, error) {
	var observed bytes.Buffer
	scan := bufio.NewScanner(io.TeeReader(src, &observed))
	scan.Buffer(make([]byte, 128), 4096)
	var lines []string
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
	if observed.Len() > 0 && observed.Bytes()[observed.Len()-1] != '\n' && len(lines) > 0 {
		lines = lines[:len(lines)-1]
	}
	return lines, scan.Err()
}
