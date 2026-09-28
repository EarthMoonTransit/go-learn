# Hints

1. Separate presence from the stored value: a zero value can be present.
2. Check how your representation behaves before any constructor has run.
3. Trace an update followed by a lookup, a miss, and an insertion at capacity.
4. Values need not be comparable. Keep key equality separate from value storage.
