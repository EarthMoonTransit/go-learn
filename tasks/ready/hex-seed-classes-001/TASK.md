# Choose representative hexadecimal fuzz seeds

## Context

A byte codec is checked by a native fuzz target. Its supplied round-trip assertion needs a more representative finite seed corpus.

## API

```go
func Encode([]byte) []byte
func Decode([]byte) ([]byte, bool)
```

## Requirements

- Extend `FuzzHex` with finite `F.Add` seeds covering varied lengths, byte ranges, order and repetition.
- Encode produces two lowercase ASCII hexadecimal digits per byte.
- Decode of that callback's encoded result must return the original bytes and `true`.
- Treat nil and empty slices as equal byte values.
- Keep assertions valid for every correct implementation and detect the required defects using your submitted seeds.

## Constraints

- All test files are editable; codec implementation is supplied. Only registered seeds execute. Malformed-wire behavior and allocation details are outside the contract.

Read the [complete source contract](FUZZ_SOURCE.md).
