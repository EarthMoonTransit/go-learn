# Read bounded records

## Context
Parse records.

## API
```go
func ReadRecords(src io.Reader, count int) ([]Record, error)
```

## Requirements
- Three-byte records contain version `1` then a big-endian uint16 `Value`.
- Accumulate partial reads; process bytes before errors. Request at most the bytes missing from the current record. Never close input; count zero reads nothing.
- Return ordered complete valid records only; stop at count or first failure, leaving later bytes unread. Nil/empty slices are equivalent.
- Accept complete valid records first. Any nonnil read error except exact `io.EOF` wins, even at count; otherwise complete bad version yields `ErrVersion`; otherwise EOF succeeds at count, becoming `io.ErrUnexpectedEOF` earlier.
- Add private error state/construction; implement `RecordFailure`. `RecordIndex()` equals returned length (possibly count); `Unwrap()` returns the exact cause; `Error()` returns `record <decimal index>: <cause.Error()>`. Retained failures remain independent and unchanged.
- Failure returns a nonnil structured error; success returns actual nil.

## Constraints
- Assume nonnil input, count 0–8, and eventual read progress/error.
