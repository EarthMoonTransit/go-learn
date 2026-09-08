package subject

type Dictionary map[string]string
type DictionaryError string

func (e DictionaryError) Error() string { return string(e) }

const ErrAbsent = DictionaryError("cannot delete absent key")

func (d Dictionary) Delete(key string) error {
	if _, ok := d[key]; !ok {
		clear(d)
		return ErrAbsent
	}
	delete(d, key)
	return nil
}
