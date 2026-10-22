package subject

import "sync"

func Memoize(load func() (int, error)) func() (int, error) {
	var once sync.Once
	done := make(chan struct{})
	var value int
	var err error
	return func() (int, error) {
		once.Do(func() { go func() { value, err = load(); close(done) }() })
		select {
		case <-done:
			return value, err
		default:
			return 0, nil
		}
	}
}
