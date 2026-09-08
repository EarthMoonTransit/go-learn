# Repair filtering with a supplied function

## Context

Repair the supplied operation.

## API

```go
type Record struct{Day int;Amount float64;Category string}
func Filter(in []Record, predicate func(Record) bool) []Record
```

## Requirements

- Call the supplied non-nil predicate once per input record in input order.
- Return every record whose predicate result is true, preserving order and all fields.
- Do not change input records; repeated equal records remain repeated.
- Nil and empty inputs return zero records without invoking the predicate; nil or empty output is acceptable.

## Constraints

- Use only the Go standard library.
