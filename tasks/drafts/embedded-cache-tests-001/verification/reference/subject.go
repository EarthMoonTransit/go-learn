package subject

type Document struct{ Key, Title string }
type Feed struct {
	Total, Calls int
	Lookup       func(string) (Document, error)
}

func (f *Feed) Count() int                         { return f.Total }
func (f *Feed) Fetch(key string) (Document, error) { f.Calls++; return f.Lookup(key) }

type Cached struct {
	*Feed
	cache map[string]Document
}

func NewCache(f *Feed) *Cached { return &Cached{Feed: f, cache: map[string]Document{}} }
func (c *Cached) Fetch(key string) (Document, error) {
	if d, ok := c.cache[key]; ok {
		return d, nil
	}
	d, e := c.Feed.Fetch(key)
	if e != nil {
		return Document{}, e
	}
	c.cache[key] = d
	return d, nil
}
