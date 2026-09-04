package subject

type Resource interface {
	Close() error
}

func IsNilResource(value Resource) bool {
	return value == nil
}
