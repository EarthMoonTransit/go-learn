package subject

import "errors"

type Frame struct {
	Tag     byte
	Payload []byte
}

var ErrFrame = errors.New("invalid frame")

func (Frame) MarshalText() ([]byte, error)   { return nil, ErrFrame }
func (*Frame) UnmarshalText([]byte) error    { return ErrFrame }
func (Frame) MarshalBinary() ([]byte, error) { return nil, ErrFrame }
func (*Frame) UnmarshalBinary([]byte) error  { return ErrFrame }
