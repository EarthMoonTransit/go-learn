# Observe struct copies and comparability

## Context
Types: Record, RecordView and SliceRecord in api.go.

## API
```go
func StructStages(count, a, b, pointed, newCount, newFirst, newPointed int) (
 source, copied RecordView, sameTarget bool,
)
func EqualRecord(left, right Record) bool
func ComparableShapes() (record, sliceRecord bool)
```

## Requirements
- Initialize a local pointee to pointed; construct a keyed Record with count, Inline `[a,b]` and its address.
- Assign the Record to a copy. Set its Count=newCount and Inline[0]=newFirst; write newPointed through its Target.
- Return both views and whether their Target pointers are identical.
- EqualRecord compares every field by Go value equality, including pointer identity; nil pointers are valid.
- ComparableShapes predicts each type's comparability from every field, regardless of values.
- Author prediction tests for shallow copying, equality and comparability; explain every field's comparability in test comments. Actual operations and assertions receive method review.

## Constraints
- All integers are -9..9. Edit only stages.go and predictions_test.go.
