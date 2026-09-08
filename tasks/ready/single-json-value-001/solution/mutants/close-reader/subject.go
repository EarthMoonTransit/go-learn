package subject

import (
	"encoding/json"
	"errors"
	"io"
)

type Record struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

var ErrJSON = errors.New("invalid JSON record")

func DecodeOne(src io.Reader) (Record, error) {
	if c, ok := src.(io.Closer); ok {
		defer c.Close()
	}
	d := json.NewDecoder(src)
	d.DisallowUnknownFields()
	var item *Record
	if e := d.Decode(&item); e != nil || item == nil {
		return Record{}, ErrJSON
	}
	var extra any
	if e := d.Decode(&extra); e != io.EOF {
		return Record{}, ErrJSON
	}
	return *item, nil
}
