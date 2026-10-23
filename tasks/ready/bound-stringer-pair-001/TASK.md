# Bind snapshot and live Stringers

## Context

Construct stable interface bindings.

## API

```go
type Label struct{ Text string }
type Counter struct{ Value int }
func (Label) String() string
func (*Counter) String() string
func Bind(Label, *Counter) (fmt.Stringer, fmt.Stringer)
```

## Requirements

- Label.String returns "label=" plus exact Text; Counter.String returns "count=" plus signed decimal Value.
- Bind's first result stores a Label snapshot with dynamic type Label, never *Label or a wrapper.
- Its second result stores the original *Counter, reflecting later changes; nil counter produces a nil interface.
- Preserve inputs. String, fmt.Sprint and fmt.Sprintf("%s", value) agree for empty/padded/Unicode text and zero/negative counts. Repeated calls are stable and nonrecursive.
- Explain stored dynamic types/values, snapshot versus live behavior, and why boxed nil pointers violate absence.

## Constraints

- Standard library only; Text≤32 bytes, Value in [-1000,1000]. Counter.String receivers are nonnil. No concurrency required.
