package subject

import (
	"encoding/json"
)

type Record struct {
	Name   string `json:"name"`
	Count  int    `json:"count,omitempty"`
	secret string
}

func NewRecord(name string, count int, secret string) Record {
	return Record{Name: name, Count: count, secret: secret}
}
func EncodeRecord(value Record) ([]byte, error) {
	return json.Marshal(map[string]any{"Name": value.Name, "Count": value.Count})
}
