package model

type Record recordState
type recordState struct {
	Label    string
	Quantity int
	secret   string
}

func New(label string, qty int, secret string) Record {
	return Record{Label: label, Quantity: qty, secret: secret}
}
func (r Record) Values() (string, int) { return r.Label, r.Quantity }
func (r Record) Secret() string        { return r.secret }
