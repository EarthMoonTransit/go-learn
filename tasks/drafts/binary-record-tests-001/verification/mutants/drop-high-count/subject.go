package subject

import (
	"encoding/binary"
	"errors"
)

type Record struct {
	Code  uint16
	Count uint32
}

var ErrRecord = errors.New("invalid binary record")

func (r Record) MarshalBinary() ([]byte, error) {
	data := make([]byte, 6)
	binary.BigEndian.PutUint16(data[:2], r.Code)
	binary.BigEndian.PutUint32(data[2:], uint32(uint16(r.Count)))
	return data, nil
}
func (r *Record) UnmarshalBinary(data []byte) error {
	if len(data) != 6 {
		return ErrRecord
	}
	value := Record{Code: binary.BigEndian.Uint16(data[:2]), Count: binary.BigEndian.Uint32(data[2:])}
	*r = value
	return nil
}
