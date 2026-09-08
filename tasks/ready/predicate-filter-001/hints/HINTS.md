# Progressive hints

1. The supplied function decides whether each record belongs in the result.
2. Count predicate calls even for rejected records.
3. Do not reuse input slots if that overwrites the caller records.
