package subject

import (
	"encoding/json"
	"strings"
)

type Record struct {
	Name   string `json:"name"`
	Count  int    `json:"count,omitempty"`
	secret string
}

func NewRecord(name string, count int, secret string) Record {
	return Record{Name: strings.TrimSpace(name), Count: count, secret: secret}
}
func EncodeRecord(value Record) ([]byte, error) { return json.Marshal(value) }
