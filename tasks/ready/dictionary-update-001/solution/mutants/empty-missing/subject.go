package subject

type Dictionary map[string]string
type DictionaryError string

func (e DictionaryError) Error() string { return string(e) }

const ErrAbsent = DictionaryError("cannot update absent key")

func (d Dictionary) Update(key, value string) error {
	if v, ok := d[key]; !ok || v == "" {
		return ErrAbsent
	}
	d[key] = value
	return nil
}
