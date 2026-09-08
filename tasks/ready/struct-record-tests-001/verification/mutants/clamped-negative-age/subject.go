package subject

type Entry struct {
	Name       string
	Age        int
	Attributes map[string]string
}

func NewEntry(name string, age int, attributes map[string]string) *Entry {
	if age < 0 {
		age = 0
	}
	return &Entry{Name: name, Age: age, Attributes: attributes}
}
