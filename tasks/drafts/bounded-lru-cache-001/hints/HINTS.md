# Hints

1. Separate key presence, stored value and recency. A stored zero is still present.
2. Write out a short operation trace and the expected oldest-to-newest keys after every step.
3. Check an existing-key update separately from insertion into a full cache. Choose a representation that makes both behaviors explicit.
