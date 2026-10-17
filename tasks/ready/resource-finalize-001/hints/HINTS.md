# Progressive hints

1. Check every path after the write attempt: each still owes exactly one cleanup.
2. Distinguish a returned write error from a short count with no write error.
3. Check whether handling cleanup hides an earlier failure; both error identities must remain classifiable.
