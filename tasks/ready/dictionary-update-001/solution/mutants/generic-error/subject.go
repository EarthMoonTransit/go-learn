package subject

import (
	"errors"
)

type Dictionary map[string]string
type DictionaryError string

func (e DictionaryError) Error() string { return string(e) }

const ErrAbsent = DictionaryError("cannot update absent key")

func (d Dictionary) Update(key, value string) error {
	if _, ok := d[key]; !ok {
		return errors.New("cannot update absent key")
	}
	d[key] = value
	return nil
}
