# Reset selected states to exact zero values

## Context

Repair the supplied operation. State has Count int, Active bool, Name string, Values []int, Labels map[string]int and Target *int.

## API

```go
func ResetSelected(states []State, selected []bool) int
```

## Requirements

- Replace each selected state with its exact Go zero value, including nil slice, map and pointer fields.
- Preserve each unselected state, including its slice header, map and pointer identities.
- Preserve previous backing elements, map entries and pointees, including data shared by selected states, unselected states and saved copies.
- Preserve selected and return its number of true entries, including already-zero states.
- Nil/empty inputs return zero. Repeated calls and repopulation follow the same rules.

## Constraints

- Lengths are equal. Every non-nil Target points outside the backing array of states. Use only the standard library.
