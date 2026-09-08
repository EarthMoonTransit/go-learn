package subject

type Entry struct {
	Name       string
	Age        int
	Attributes map[string]string
}

func (e *Entry) Clear() { clear(e.Attributes); *e = Entry{} }
