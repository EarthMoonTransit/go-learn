# Review notes

## Decision

Returned to draft. The intended mechanism is delaying the integer conversion
until after percentage arithmetic. The declared input domain nevertheless
allows every nonnegative `int`; converting values near `MaxInt` through
`float64` does not provide the stated mathematical result. The zero-input rule
is tested for percentage but not capacity.

## Required repair

Declare a safe upper input bound, make `0 <= successPercent <= 100` an input
precondition, and cover zero capacity. Recalibrate this small arithmetic repair
to a shorter basic exercise. Its strong semantic proximity to the source
direction also needs task-specific clean-room evidence and exact attribution.
