# LRU

## Context
Implement caching.

## API
```go
type Cache[K comparable, V any] struct { /*...*/ }
func New[K comparable, V any](capacity int) *Cache[K,V]
func (c *Cache[K,V]) Get(key K) (V,bool)
func (c *Cache[K,V]) Set(key K, value V)
func (c *Cache[K,V]) Range(f func(K,V) bool)
func (c *Cache[K,V]) Clear()
```

## Requirements
- New creates independent empty caches; zero Cache has zero capacity.
- Set assigns V; zero capacity retains nothing.
- Get returns (value,true), including zero/nil; misses return (zero V,false).
- Set and successful Get refresh recency; misses change nothing.
- Overflow evicts oldest; updates never evict another key.
- Range visits oldest-first, once each; false stops immediately; traversal changes nothing.
- Clear empties, preserving capacity; repeatable.

## Constraints
- Sequential calls; capacity≥0; keys equal themselves without panicking.
Non-nil callbacks never call cache methods. Non-nil receivers. Do not copy used Cache. No complexity bound.
