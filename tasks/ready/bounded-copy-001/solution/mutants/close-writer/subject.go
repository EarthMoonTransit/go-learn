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
	if c, ok := dst.(io.Closer); ok {
		defer c.Close()
	}
	return io.Copy(dst, io.LimitReader(src, limit))
}
