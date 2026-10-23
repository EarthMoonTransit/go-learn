package subject

import "io"

func Relay(dst io.Writer, produce func(io.Writer) error, limit int64) (n int64, producerErr, copyErr error) {
	return
}
