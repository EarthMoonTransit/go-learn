package subject

import (
	"errors"
	"io"
)

var ErrLimit = errors.New("negative limit")

func CopyAtMost(dst io.Writer, src io.Reader, limit int64) (int64, error) {
	if limit < 0 {
		return 0, ErrLimit
	}
	n, e := io.Copy(dst, io.LimitReader(src, limit))
	if e == nil && n > 0 {
		if c, ok := src.(io.Closer); ok {
			c.Close()
		}
		if c, ok := dst.(io.Closer); ok {
			c.Close()
		}
	}
	return n, e
}
