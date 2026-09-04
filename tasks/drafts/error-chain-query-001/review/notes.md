# Review notes

## Decision

Returned to draft. The package combines typed-error construction, `Unwrap`, a
nil-wrapping policy, `errors.Is`, and `errors.As`. Those are sequential taxonomy
abilities, not one focused wrapping exercise. The exported `FieldError` can also
be constructed with `Err == nil`, while the reference `Error` method panics.

## Required repair

Separate error construction/wrapping from later chain-query transfer practice;
do not add a replacement task in Phase 3C. Define the zero state of any exported
error type. Promotion of an original task must also wait for resolution of the
normative conflict between AGENTS.md's internal registry-source rule and ADR
0008's project-identity model.
