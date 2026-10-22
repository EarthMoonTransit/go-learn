package subject

import "io"

func Exchange(produce func(io.Writer) error, consume func(io.Reader) error) (producerErr, consumerErr error) {
	return nil, nil
}
