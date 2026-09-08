# Test an overridden fetch and a promoted count method

## Context

Write tests for the specified contract.

## API

```go
type Document struct { Key, Title string }
// Feed has Total, Calls int and Lookup func(string)(Document,error).
// Cached embeds *Feed; Count is promoted.
func NewCache(f *Feed) *Cached
func (c *Cached) Fetch(key string) (Document,error)
func (f *Feed) Count() int
```

## Requirements

- NewCache uses the supplied non-nil Feed; Lookup is non-nil. Promoted Count returns the current Feed.Total.
- First Fetch calls Feed.Lookup through Feed.Fetch, increments Feed.Calls, and caches a successful Document under the exact requested key.
- Repeated successful keys return the cached document without another lookup. Different keys and cache instances are independent.
- Errors return a zero Document and the original error; failures are not cached. Successful empty documents are valid.

## Constraints

- Use only the Go standard library.
