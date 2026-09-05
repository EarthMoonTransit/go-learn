# Progressive hints

1. The existing interior case checks only one successful window. Add both successful boundary cases and invalid bounds.
2. Check empty windows separately from invalid windows: both return empty text, but their booleans differ.
3. Use explicit zero/high-byte input and windows that select individual bytes. Also distinguish an exclusive end from an inclusive one, and reject negative, reversed and out-of-range bounds.
