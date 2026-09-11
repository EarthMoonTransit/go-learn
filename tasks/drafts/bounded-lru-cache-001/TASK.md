# Implement a bounded LRU cache

## Context
Complete the supplied cache interface.

## API
```go
func New(capacity int) Cache
// Cache has Get(int) (int, bool), Set(int, int),
// Range(func(int, int) bool), and Clear().
```

## Requirements

- New returns an independent, empty, usable cache.
- Set stores or replaces a value; capacity zero retains nothing.
- Get returns the stored value and true, including stored zero; a missing key returns zero and false.
- Set and successful Get make that key most recent. Missing Get changes nothing.
- Overflow evicts the least recently used key. Updates never evict another key.
- Range visits each entry once, oldest first, stopping immediately after a callback returns false. Range changes nothing.
- Clear removes all entries, is repeatable, and preserves capacity for reuse.

## Constraints
- Nonnegative capacity; arbitrary integer keys/values; sequential calls.
- Non-nil Range callbacks never call cache methods. No complexity bound.
