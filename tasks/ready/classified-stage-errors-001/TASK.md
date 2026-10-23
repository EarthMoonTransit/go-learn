# Classify failures at an operation boundary
## Context
ErrTemporary and nil-safe Detail are supplied.
## API
```go
type Stage struct{Name string; Expose bool; Run func()(int,error)}
type Outcome struct{Value int; Err error; Temporary bool; Detail *Detail; HasDetail bool}
func RunStages(stages []Stage) Outcome
```
## Requirements
- Call stages once in order; stop at the first non-nil error. Preserve its value; otherwise return the last value, or zero when empty.
- Failure text is Name + ": " + cause.Error(), preserving text exactly.
- Expose wraps directly to the original cause; otherwise return opaque context with no unwrap cause.
- Classify the returned error using errors.Is with ErrTemporary and errors.As for *Detail; preserve first-match identity and typed-nil success.
- Success has nil Err, false flags and nil Detail. Explain why opaque causes remain unclassified.
## Constraints
- At most four stages and eight finite cause nodes. Callbacks do not panic. Standard library; actual method review required.
