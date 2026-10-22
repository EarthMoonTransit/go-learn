# Document a deterministic summary

## Context

Show a deterministic map summary to an API reader.

## API

```go
func Summary(values map[string]int) string
```

## Requirements

- Keys appear in case-sensitive ascending order as key=value, joined by comma-space.
- Values use signed decimal notation, including zero; nil and empty maps return `(empty)`.
- Author discoverable ExampleSummary functions with intentional Output comments, showing multi-key ordering, negative and zero values, and empty input.
- Choose examples useful to a reader of the API; keep expected output deterministic. Do not edit the implementation.

## Constraints

- Maps contain at most eight entries, nonempty ASCII-letter keys and values from -100 through 100. Use package subject_test. The starter's trivial empty example earns no new-action credit.

Public checks execute your examples against the correct implementation; full checks challenge them with defects. Method review separately checks executable naming, intentional output matching and documentation value. Ordinary tests alone cannot satisfy this examples requirement.
