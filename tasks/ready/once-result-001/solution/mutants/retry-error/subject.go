package subject

import "sync"

func Memoize(load func() (int, error)) func() (int, error) {
	var once sync.Once
	var value int
	var err error
	return func() (int, error) {
		once.Do(func() { value, err = load() })
		if err != nil {
			return load()
		}
		return value, err
	}
}
