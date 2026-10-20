# Replace selected bits while preserving the rest

## Context

Repair:

## API

```go
func ReplaceBits(base, replacement, mask uint8) uint8
```

## Requirements

- At every bit selected by mask, take the bit from replacement, including selected zero bits.
- At every unselected bit, preserve the bit from base.
- A zero mask returns base unchanged.
- A full mask returns replacement unchanged.
- Inputs are arbitrary eight-bit values; overlapping set bits require no special precondition.

## Constraints

- Keep the signature. Use only the standard library. Any implementation satisfying the bit-selection contract is accepted.
