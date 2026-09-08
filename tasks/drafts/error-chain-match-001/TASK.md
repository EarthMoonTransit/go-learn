# Match an error through wrapped and joined causes

## Context

Complete the specified operations.

## API

```go
func Matches(err, target error) bool
```

## Requirements

- Return the result of Go error identity matching: direct identity, wrapped causes and any cause of a joined error qualify.
- Respect a source error implementing Is(error) bool. Equal message text alone does not establish identity.
- Two nil errors match; exactly one nil does not. Preserve errors and do not unwrap only a single level.
- Input error chains are finite and follow the standard error conventions.

## Constraints

- Use only the Go standard library.
