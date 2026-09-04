package subject

import (
	"errors"
	"fmt"
)

type Resource interface {
	Write([]byte) (int, error)
	Close() error
}

func Store(resource Resource, data []byte) (err error) {
	defer func() { err = errors.Join(err, resource.Close()) }()
	_, writeErr := resource.Write(data)
	if writeErr != nil {
		return fmt.Errorf("write resource: %w", writeErr)
	}
	return nil
}
