# Review notes

## Decision

Returned to draft. Byte offsets, rune ordinals, terminal offsets, and a single
invalid byte are well separated, but the prompt promises behavior for each
invalid byte. Consecutive invalid bytes, truncated encodings, and a valid encoded
U+FFFD are missing, so the evidence is incomplete.

## Required repair

Add those cases and focused mutants without importing canonical cases. Define
how a learner runs a mutation-test submission and receives useful surviving
defect feedback. The broad strings metadata seed needs a documented derivation
boundary or original classification; the latter cannot be promoted until the
original-source policy conflict is resolved.
