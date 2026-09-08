package subject

type Dictionary map[string]string
type DictionaryError string

func (e DictionaryError) Error() string { return string(e) }

const ErrAbsent = DictionaryError("cannot update absent key")

func (d Dictionary) Update(key, value string) error { d[key] = value; return nil }
