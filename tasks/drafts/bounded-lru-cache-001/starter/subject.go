package subject

type Cache interface {
	Get(key int) (int, bool)
	Set(key, value int)
	Range(f func(key, value int) bool)
	Clear()
}

type cache struct{}

func New(capacity int) Cache                       { return &cache{} }
func (c *cache) Get(key int) (int, bool)           { return 0, false }
func (c *cache) Set(key, value int)                {}
func (c *cache) Range(f func(key, value int) bool) {}
func (c *cache) Clear()                            {}
