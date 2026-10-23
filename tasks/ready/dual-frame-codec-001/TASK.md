# Encode one frame in two formats

## Context
Implement both formats.

## API
```go
type Frame struct{Tag byte;Payload []byte}
var ErrFrame=errors.New("invalid frame")
func (Frame) MarshalText() ([]byte,error)
func (*Frame) UnmarshalText([]byte) error
func (Frame) MarshalBinary() ([]byte,error)
func (*Frame) UnmarshalBinary([]byte) error
```

## Requirements
- Payload is at most 32 bytes; nil/empty are equivalent. Every Tag is valid.
- Binary is version byte 1, Tag, one-byte payload length, then payload.
- Text is exactly 1/tt/nn/pp: tt and nn are two lowercase hex digits for Tag and length; pp is lowercase payload hex.
- Validate version, length and full input before replacing receiver. Reject malformed/truncated/trailing data with original ErrFrame; preserve input always and receiver on error.
- Marshal oversize values as nil,ErrFrame. Marshal preserves state. Success returns nil error. Outputs are independent of receiver and each other; decoded payload owns storage.
- Validate version and total length before payload allocation. Method review checks formats.

## Constraints
- Nonnil receivers; inputs≤512 bytes.
