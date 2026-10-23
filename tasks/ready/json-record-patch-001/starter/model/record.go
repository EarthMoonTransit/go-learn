package model

type Record recordState
type recordState struct{}

func New(label string, qty int, secret string) Record { return Record{} }
func (r Record) Values() (string, int)                { return "", 0 }
func (r Record) Secret() string                       { return "" }
