# Render values through an implicit interface

## Context

Implement the supplied formatting operations so values work through a small consumer-side contract.

## API

```go
type Formatter interface { Format() string }
type Pair struct { Left, Right string }
func (Pair) Format() string
func Render(value Formatter) string
```

## Requirements

- Pair.Format returns Left + ":" + Right exactly.
- Pair values work as Formatter values, including through Render.
- Render calls the interface method exactly once.
- Accept caller-defined formatter types without registration or type switches.
- Return the formatter's result unchanged, including empty strings and whitespace.

## Constraints

- All Pair strings are valid. Render receives a non-nil interface whose Format method returns normally. Edit subject/subject.go. Keep the API; do not add a closed list of supported types. Use only the standard library.
