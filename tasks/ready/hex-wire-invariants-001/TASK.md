# Check hexadecimal wire invariants

## Context

The supplied native target has a varied finite corpus and a round-trip check. Consistent encoder and decoder mistakes can still preserve every decoded value.

## API

```go
func Encode([]byte) []byte
func Decode([]byte) ([]byte, bool)
```

## Requirements

- Extend `FuzzHex` with an independent assertion of canonical lowercase ASCII hexadecimal encoding: two digits per input byte, in input order.
- Decode the actual Encode result from the same callback; require the original bytes and `true`.
- Treat nil and empty slices as equal byte values.
- Keep assertions valid for all correct implementations.
- Detect the required defects using only your submitted seeds and assertions.

## Constraints

- All test files and seeds are editable. Codec implementation is supplied. Malformed-wire behavior, allocation details and randomized search are outside the contract.

Read the [complete source contract](FUZZ_SOURCE.md).
