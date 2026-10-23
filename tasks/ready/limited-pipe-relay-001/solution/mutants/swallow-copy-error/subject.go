package subject

import "io"

func Relay(dst io.Writer, produce func(io.Writer) error, limit int64) (n int64, producerErr, copyErr error) {
	reader, writer := io.Pipe()
	finished := make(chan error, 1)
	go func() {
		err := produce(writer)
		writer.CloseWithError(err)
		finished <- err
	}()
	n, copyErr = io.Copy(dst, io.LimitReader(reader, limit))
	reader.CloseWithError(copyErr)
	producerErr = <-finished
	copyErr = nil
	return
}
