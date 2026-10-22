# Hints

- Keep the accepted prefix separate from the bytes still needed for the current record.
- A read's byte count and error can both be meaningful. Compare the required priority before choosing a failure cause.
- Recovery state must outlive this call without being overwritten by another call. The input belongs to the caller.
