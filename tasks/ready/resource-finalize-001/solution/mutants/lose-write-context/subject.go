package subject

import (
	"errors"
	"io"
)

type Resource interface {
	Write([]byte) (int, error)
	Close() error
}

func Store(resource Resource, data []byte) (err error) {
	defer func() {
		err = errors.Join(err, resource.Close())
	}()

	written, writeErr := resource.Write(data)
	if writeErr != nil {
		return writeErr
	}
	if written != len(data) {
		return io.ErrShortWrite
	}
	return nil
}
