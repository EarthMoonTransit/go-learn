# Decode exactly one object and reject trailing values

## Context

Complete the specified operations.

## API

```go
type Record struct{Name string `json:"name"`;Count int `json:"count"`}
var ErrJSON=errors.New("invalid JSON record")
func DecodeOne(src io.Reader) (Record, error)
```

## Requirements

- Decode exactly one JSON object with optional name string and count integer fields; missing fields take their zero values.
- Reject unknown fields, top-level null, wrong non-null field types, malformed input and any second JSON value or non-whitespace suffix.
- On rejection return zero Record and the supplied non-nil ErrJSON; preserve the sentinel. Whitespace around one object is allowed.
- Do not close src. Standard encoding/json duplicate-key, case-insensitive field-name matching and field-level null rules apply; no additional schema validation is required.

## Constraints

- Use only the Go standard library.
