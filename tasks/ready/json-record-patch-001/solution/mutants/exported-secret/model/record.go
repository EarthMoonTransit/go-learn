package model

type Record recordState
type recordState struct {
	Label       string `json:"label"`
	Quantity    int    `json:"qty,omitempty"`
	SecretValue string `json:"-"`
}

func New(label string, qty int, SecretValue string) Record {
	return Record{Label: label, Quantity: qty, SecretValue: SecretValue}
}
func (r Record) Values() (string, int) { return r.Label, r.Quantity }
func (r Record) Secret() string        { return r.SecretValue }
