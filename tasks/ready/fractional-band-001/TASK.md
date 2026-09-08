# Repair classification before numeric narrowing

## Context

Repair the supplied operation.

## API

```go
func BandRate(value float64) float32
```

## Requirements

- Return 3.213 for value below zero.
- Return 0.5 for values from zero inclusive to 1000 exclusive.
- Return 1.621 for values from 1000 inclusive to 5000 exclusive; return 2.475 at or above 5000.
- Classify the original float64 input without narrowing it first; return the corresponding float32 rate. Inputs are finite and between -1000000000 and 1000000000.

## Constraints

- Use only the Go standard library.
