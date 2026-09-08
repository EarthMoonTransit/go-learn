package subject

type Entry struct {
	Name       string
	Age        int
	Attributes map[string]string
}

func (e *Entry) Clear() { e.Age = 0; e.Attributes = nil }
