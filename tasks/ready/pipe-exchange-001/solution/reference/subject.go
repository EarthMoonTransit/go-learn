package subject

import "io"

func Exchange(produce func(io.Writer) error, consume func(io.Reader) error) (producerErr, consumerErr error) {
	reader, writer := io.Pipe()
	producerDone := make(chan error, 1)
	go func() { err := produce(writer); writer.CloseWithError(err); producerDone <- err }()
	consumerErr = consume(reader)
	reader.CloseWithError(consumerErr)
	producerErr = <-producerDone
	return producerErr, consumerErr
}
