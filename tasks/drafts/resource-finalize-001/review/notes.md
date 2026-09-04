# Review notes

## Decision

Returned to draft. The task mixes cleanup ownership, `defer`, named-result
mutation, the `io.Writer` short-write contract, wrapping, and an `errors.Join`
policy. The grader does not verify the promised operation context, an empty
write, or all simultaneous write/close failure paths. The final hints reveal the
reference's named-return and join strategy.

## Required repair

Split basic cleanup/order from advanced multi-error preservation, weaken hints
to ownership and error-policy invariants, and add focused context and combined
failure evidence. Original-content promotion remains blocked by the AGENTS.md
versus ADR 0008 provenance conflict.
