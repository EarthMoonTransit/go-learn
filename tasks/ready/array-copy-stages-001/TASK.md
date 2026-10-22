# Observe array copy stages

## Context
Implement four observations of one array experiment.

## API
```go
func ArrayStages(a, b, c, replacement, argumentValue int) (
 source, assigned, insideCall, afterCall [3]int,
)
```

## Requirements
- Construct a three-element source array containing a, b, c.
- Assign it to another array variable and replace that variable's index 0 with replacement.
- Author an unexported helper taking `[3]int` by value. It replaces its local index 1 with argumentValue and returns that array; call it with the assigned array.
- Return the source, assigned array before the call, helper result and caller array after the call, in that order. Other elements remain unchanged.
- Write two prediction tests in predictions_test.go: assignment independence and callee/caller independence, using concrete expected arrays.
- Method review checks your literal length, assignment, helper, call and assertions; fabricated observations do not satisfy the experiment.

## Constraints
- All integers are -9..9. Edit only the two declared files.
