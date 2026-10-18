# Synthetic finite native hex verifier fixture

## Context

This infrastructure fixture checks native seed verification outside the learner catalog.

## API

```go
func Encode([]byte) []byte
func Decode([]byte) ([]byte, bool)
```

## Requirements

- Write `FuzzHex` and register finite byte seeds.
- Check that Encode returns lowercase hexadecimal bytes, two digits per input byte.
- Decode only bytes returned by Encode in the same callback.
- Check that Decode returns the input bytes and true status.
- Treat nil and empty slices as equal byte values.

## Constraints

- Use the finite native language and resource limits in the linked contract.
- Public Check accepts every declared correct implementation; Complete also detects every required defect using the same seeds.

Read the [complete source contract](FUZZ_SOURCE.md).
