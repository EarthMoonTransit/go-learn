# Sum by recursive decomposition

## Context

Implement the specified operation.

## API

```go
func Total(values []int) int
```

## Requirements

- Return the sum of every element, including zeros and negative values; empty and nil inputs return zero.
- Preserve every input element.
- Use learner-authored recursion whose subresults contribute to the returned total.
- In review, identify the terminal cases, a strictly decreasing measure for each recursive call, and how all required subresults are combined.

## Constraints

- Inputs have at most 32 elements, each between -1000 and 1000. Use only the standard library.

Behavioral tests check results and preservation. Method acceptance additionally requires inspection of your actual recursive code and explanation against the three review criteria above. An iterative equivalent can pass the tests but does not complete the recursion requirement. No source-shape checker is used.
