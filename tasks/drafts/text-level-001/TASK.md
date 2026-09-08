# Implement an atomic single-digit text representation

## Context

Complete the specified operations.

## API

```go
type Level uint8
var ErrLevel=errors.New("invalid level")
func (l Level) MarshalText() ([]byte, error)
func (l *Level) UnmarshalText(text []byte) error
```

## Requirements

- Levels zero through nine marshal as exactly one ASCII digit with nil error. Other Level values return nil and the supplied non-nil ErrLevel.
- UnmarshalText accepts exactly one ASCII digit into a non-nil receiver; reject every other byte sequence without changing the receiver.
- Preserve input bytes and ErrLevel identity. Do not trim spaces or accept Unicode digits, signs or multiple digits.
- Separate successful marshal calls return independent byte storage.

## Constraints

- Use only the Go standard library.
