# Hints

1. Identify both operations that can wait indefinitely when the caller stops cooperating.
2. Cancellation must be an alternative at each blocking operation, not just before the loop.
3. Output belongs to this stage; input and the parent context belong to its caller.
