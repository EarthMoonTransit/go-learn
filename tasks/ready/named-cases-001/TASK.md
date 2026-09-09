# Run independent named subtests in table order

## Context

Complete the specified operations.

## API

```go
type Case struct{Name string;Check func(*testing.T)}
func RunNamed(t *testing.T, cases []Case)
```

## Requirements

- For each Case, create one real t.Run subtest using its Name and invoke Check once with that child *testing.T.
- Run cases serially in input order. Empty input creates no subtests. Names are unique nonempty ASCII letters/digits/hyphens; callbacks are non-nil.
- Preserve the input table. Do not invoke callbacks with the parent, change names, duplicate calls or mark subtests parallel.
- The provided public callbacks inspect their real test names and identities. Callback failures use normal testing package propagation.

## Constraints

- Use only the Go standard library.
