package subject

import (
	"fmt"
	"io"
)

type Resource interface {
	Write([]byte) (int, error)
	Close() error
}

func Store(resource Resource, data []byte) error {
	written, err := resource.Write(data)
	if err != nil {
		return fmt.Errorf("write resource: %w", err)
	}
	if written != len(data) {
		return io.ErrShortWrite
	}
	return resource.Close()
}
