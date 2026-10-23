package model

type Record recordState
type recordState struct {
	label    string `json:"label"`
	Quantity int    `json:"qty,omitempty"`
	secret   string
}

func New(label string, qty int, secret string) Record {
	return Record{label: label, Quantity: qty, secret: secret}
}
func (r Record) Values() (string, int) { return r.label, r.Quantity }
func (r Record) Secret() string        { return r.secret }
