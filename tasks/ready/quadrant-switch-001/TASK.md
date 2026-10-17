# Classify coordinates with an expressionless switch

## Context

Classify a pair of integer coordinates, including the origin and coordinate axes.

## API

```go
func Locate(x, y int) string
```

## Requirements

- Return origin for (0, 0).
- On the horizontal axis return positive-x or negative-x according to x's sign.
- On the vertical axis return positive-y or negative-y according to y's sign.
- For nonzero coordinates return quadrant-i (+,+), quadrant-ii (-,+), quadrant-iii (-,-), or quadrant-iv (+,-).
- Use an expressionless switch; resolve origin and axes before quadrants, without fallthrough.

## Constraints

- All int values are valid. Edit only subject/implementation.go. Locate(0, -3) returns negative-y; Locate(-2, 4) returns quadrant-ii. Tokens are exact. Public checks report incorrect tokens, a tagged switch, or fallthrough.
