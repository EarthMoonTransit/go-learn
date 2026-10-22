package subject

import "sync"

var once sync.Once
var value int
var err error

func Memoize(load func() (int, error)) func() (int, error) {
	return func() (int, error) { once.Do(func() { value, err = load() }); return value, err }
}
