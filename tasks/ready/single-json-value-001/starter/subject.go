package subject

import (
	"errors"
	"io"
)

type Record struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

var ErrJSON = errors.New("invalid JSON record")

func DecodeOne(src io.Reader) (Record, error) { return Record{}, nil }
