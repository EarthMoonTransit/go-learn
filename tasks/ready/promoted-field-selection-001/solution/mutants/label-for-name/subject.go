package subject

type Header struct {
	Code        int
	Name, Label string
}
type Left struct{ Value int }
type Right struct{ Value int }
type Record struct {
	Header
	Left
	Right
	Code int
}

func SelectFields(r Record) (outer, embedded int, name string, left, right int) {
	return r.Code, r.Header.Code, r.Label, r.Left.Value, r.Right.Value
}
