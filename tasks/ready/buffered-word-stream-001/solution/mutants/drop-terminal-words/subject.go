package subject

import (
	"bufio"
	"io"
)

func WriteWords(dst io.Writer, src io.Reader) error {
	var pending bufio.Writer
	pending.Reset(dst)
	scan := bufio.NewScanner(src)
	scan.Split(func(data []byte, atEOF bool) (int, []byte, error) {
		if atEOF {
			return len(data), nil, nil
		}
		return bufio.ScanWords(data, atEOF)
	})
	scan.Buffer(make([]byte, 16), 256)
	var writeErr error
	for scan.Scan() {
		if _, err := pending.WriteString(scan.Text() + "\n"); err != nil {
			writeErr = err
			break
		}
	}
	flushErr := pending.Flush()
	if writeErr != nil {
		return writeErr
	}
	if flushErr != nil {
		return flushErr
	}
	return scan.Err()
}
