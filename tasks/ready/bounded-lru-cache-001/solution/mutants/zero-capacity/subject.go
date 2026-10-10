package subject

type Cache[K comparable, V any] struct{ state *state[K, V] }
type entry[K comparable, V any] struct {
	key   K
	value V
}
type state[K comparable, V any] struct {
	capacity int
	entries  []entry[K, V]
}

func New[K comparable, V any](capacity int) *Cache[K, V] {
	if capacity == 0 {
		capacity = 1
	}
	return &Cache[K, V]{state: &state[K, V]{capacity: capacity}}
}
func (c *Cache[K, V]) index(key K) int {
	if c.state == nil {
		return -1
	}
	for i, e := range c.state.entries {
		if e.key == key {
			return i
		}
	}
	return -1
}
func (c *Cache[K, V]) remove(i int) {
	c.state.entries = append(c.state.entries[:i], c.state.entries[i+1:]...)
}
func (c *Cache[K, V]) Get(key K) (V, bool) {
	var zero V
	i := c.index(key)
	if i < 0 {
		return zero, false
	}
	e := c.state.entries[i]
	c.remove(i)
	c.state.entries = append(c.state.entries, e)
	return e.value, true
}
func (c *Cache[K, V]) Set(key K, value V) {
	if c.state == nil || c.state.capacity == 0 {
		return
	}
	i := c.index(key)
	if i >= 0 {
		c.remove(i)
	}
	c.state.entries = append(c.state.entries, entry[K, V]{key, value})
	if len(c.state.entries) > c.state.capacity {
		c.remove(0)
	}
}
func (c *Cache[K, V]) Range(f func(K, V) bool) {
	if c.state == nil {
		return
	}
	for _, e := range c.state.entries {
		if !f(e.key, e.value) {
			return
		}
	}
}
func (c *Cache[K, V]) Clear() {
	if c.state != nil {
		c.state.entries = nil
	}
}
