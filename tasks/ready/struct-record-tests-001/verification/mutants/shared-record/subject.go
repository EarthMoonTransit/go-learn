package subject

type Entry struct {
	Name       string
	Age        int
	Attributes map[string]string
}

var shared Entry

func NewEntry(name string, age int, attributes map[string]string) *Entry {
	shared = Entry{Name: name, Age: age, Attributes: attributes}
	return &shared
}
