package subject

type Cache[K comparable, V any] struct{ state *state[K, V] }
type state[K comparable, V any] struct{}

func New[K comparable, V any](capacity int) *Cache[K, V] { return nil }
func (c *Cache[K, V]) Get(key K) (V, bool)               { var zero V; return zero, false }
func (c *Cache[K, V]) Set(key K, value V)                {}
func (c *Cache[K, V]) Range(f func(K, V) bool)           {}
func (c *Cache[K, V]) Clear()                            {}
