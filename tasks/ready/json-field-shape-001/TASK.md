# Repair JSON names omission and private fields

## Context

Repair the supplied operation.

## API

```go
type Record struct{Name string `json:"name"`;Count int `json:"count,omitempty"`;secret string}
func NewRecord(name string, count int, secret string) Record
func EncodeRecord(value Record) ([]byte, error)
```

## Requirements

- The supplied Record has exported Name and Count and private secret. NewRecord stores all three supplied values.
- EncodeRecord returns a JSON object with name always present, count omitted only when zero, and no secret field.
- Preserve the supplied Record declarations and honor their JSON tags in EncodeRecord. Preserve negative counts and exact Unicode/text content; object key order and escaping spelling are unspecified.
- Return encoding errors normally. Do not mutate the Record or read environment state.

## Constraints

- Use only the Go standard library.
