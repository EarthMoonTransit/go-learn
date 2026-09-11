package subject

type Cache interface {
	Get(key int) (int, bool)
	Set(key, value int)
	Range(f func(key, value int) bool)
	Clear()
}

type entry struct{ key, value int }
type cache struct {
	capacity int
	entries  []entry
}

func New(capacity int) Cache { return &cache{capacity: capacity} }

func (c *cache) index(key int) int {
	for i, e := range c.entries {
		if e.key == key {
			return i
		}
	}
	return -1
}
func (c *cache) remove(i int) {
	c.entries = append(c.entries[:i], c.entries[i+1:]...)
}
func (c *cache) Get(key int) (int, bool) {
	i := c.index(key)
	if i < 0 {
		return 0, false
	}
	e := c.entries[i]
	c.remove(i)
	c.entries = append(c.entries, e)
	return e.value, true
}
func (c *cache) Set(key, value int) {
	if value < 0 {
		value = 0
	}
	if c.capacity == 0 {
		return
	}
	i := c.index(key)
	if i >= 0 {
		c.remove(i)
	}
	c.entries = append(c.entries, entry{key, value})
	if len(c.entries) > c.capacity {
		c.remove(0)
	}
}
func (c *cache) Range(f func(key, value int) bool) {
	for _, e := range c.entries {
		if !f(e.key, e.value) {
			return
		}
	}
}
func (c *cache) Clear() { c.entries = nil }
