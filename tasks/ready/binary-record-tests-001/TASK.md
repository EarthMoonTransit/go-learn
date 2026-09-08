# Test fixed-width big-endian binary marshaling

## Context

Write tests for the specified contract.

## API

```go
type Record struct{Code uint16;Count uint32}
var ErrRecord=errors.New("invalid binary record")
func (r Record) MarshalBinary() ([]byte, error)
func (r *Record) UnmarshalBinary(data []byte) error
```

## Requirements

- MarshalBinary returns six bytes: Code as two big-endian bytes followed by Count as four big-endian bytes, with nil error.
- UnmarshalBinary replaces the non-nil receiver from exactly six bytes. Reject every other length with the supplied non-nil ErrRecord and leave the receiver unchanged.
- Preserve the input bytes and sentinel. Returned byte slices from separate marshal calls are independent.
- Support all uint16 and uint32 values; round-trip tests alone cannot establish byte order.

## Constraints

- Use only the Go standard library.
