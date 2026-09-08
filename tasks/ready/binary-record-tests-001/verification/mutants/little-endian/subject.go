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
	binary.LittleEndian.PutUint16(data[:2], r.Code)
	binary.LittleEndian.PutUint32(data[2:], r.Count)
	return data, nil
}
func (r *Record) UnmarshalBinary(data []byte) error {
	if len(data) != 6 {
		return ErrRecord
	}
	value := Record{Code: binary.LittleEndian.Uint16(data[:2]), Count: binary.LittleEndian.Uint32(data[2:])}
	*r = value
	return nil
}
