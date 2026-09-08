package subject

type Entry struct {
	Name       string
	Age        int
	Attributes map[string]string
}

func NewEntry(name string, age int, attributes map[string]string) *Entry {
	if attributes == nil {
		attributes = map[string]string{}
	}
	return &Entry{Name: name, Age: age, Attributes: attributes}
}
