package subject

import "errors"

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
	if e := scan.Err(); e != nil {
		return lines, errors.New(e.Error())
	}
	return lines, nil
}
