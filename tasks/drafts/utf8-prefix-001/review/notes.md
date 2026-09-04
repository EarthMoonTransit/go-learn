# Review notes

## Decision

Returned to draft. The reference uses `utf8.DecodeRuneInString` widths and
preserves source bytes; it does not assess Go `range`, so
`unicode.range-decoding` is the wrong primary skill. Negative limits are stated
but only zero is tested, and valid encoded U+FFFD is not separated from malformed
input.

## Required repair

Reclassify the task to `unicode.utf8-encoding` with an explicit pilot deviation,
then add negative-limit, exact multibyte-boundary, valid replacement-rune, and
truncated-input evidence with distinct mutants. Complete exact metadata
attribution and durable clean-room evidence before another review.
