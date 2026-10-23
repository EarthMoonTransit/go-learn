# Cache by key

## Context
Cache completed initialization separately per key.

## API
```go
func NewRegistry(load func(string) (int, error)) *Registry
func (r *Registry) Get(key string) (int, error)
```

## Requirements
- Construction is lazy. Get initializes each key exactly once using that key.
- Return complete cached values and identical errors, including zero and failure; never retry.
- Keys and registries are independent. Same-key callers wait for completed initialization.
- Own stable per-key sync.Once state; never reset, replace or copy used entries.
- Protect lookup and installation together under a registry-owned mutex; unlock before Once.Do/callbacks. Explain the entry-identity invariant and publication.

## Constraints
- Only NewRegistry instances; do not copy. Non-nil callbacks return without panic/reentry. No eviction/teardown.
- At most four keys, sixteen calls, eight bytes per key. Standard library only; actual method review required.
