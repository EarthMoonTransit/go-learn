# Progressive hints

1. Use the supplied writer as the only output destination.
2. Check the error before interpreting a short byte count.
3. A short write with nil error still needs an error result.
