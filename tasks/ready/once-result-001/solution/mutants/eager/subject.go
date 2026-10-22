package subject

import "sync"

func Memoize(load func() (int, error)) func() (int, error) {
	var once sync.Once
	var value int
	var err error
	once.Do(func() { value, err = load() })
	return func() (int, error) { once.Do(func() { value, err = load() }); return value, err }
}
