# Review notes

Pending independent review. Confirm that invalid UTF-8 preservation is intentional and adequately distinguished from rune re-encoding. Mutants cover byte counting, an off-by-one boundary, normalization of invalid bytes, and premature termination. No canonical statement or cases were consulted.
