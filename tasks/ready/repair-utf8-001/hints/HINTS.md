# Progressive hints

1. A malformed unit consumes one byte under Go decoding.
2. An encoded replacement character is valid input and must remain.
3. Compare how []rune conversion treats consecutive invalid bytes with a function that replaces an invalid run.
