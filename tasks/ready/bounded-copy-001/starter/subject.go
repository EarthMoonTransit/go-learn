package subject

import (
	"errors"
	"io"
)

var ErrLimit = errors.New("negative limit")

func CopyAtMost(dst io.Writer, src io.Reader, limit int64) (int64, error) { return 0, nil }
