# Create an independent inclusive-range closure

## Context

Complete the specified operations.

## API

```go
type Record struct{Day int;Amount float64;Category string}
type DaysPeriod struct{From,To int}
func WithinDays(p DaysPeriod) func(Record) bool
```

## Requirements

- Return a non-nil function that accepts a record when Day is between From and To inclusive.
- For From greater than To, every record is rejected.
- Only Day affects acceptance; Amount and Category are irrelevant.
- Each returned function retains its own supplied period; later constructor calls or caller changes do not alter it.

## Constraints

- Use only the Go standard library.
