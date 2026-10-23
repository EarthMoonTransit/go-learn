package subject

import "sync"

type entry struct {
	once  sync.Once
	value int
	err   error
}
type Registry registryState
type registryState struct {
	mu      sync.Mutex
	entries map[string]*entry
	load    func(string) (int, error)
}

func NewRegistry(load func(string) (int, error)) *Registry {
	return &Registry{entries: make(map[string]*entry), load: load}
}
func (r *Registry) Get(key string) (int, error) {
	r.mu.Lock()
	current := r.entries[key]
	if current == nil {
		current = &entry{}
		r.entries[key] = current
	}
	r.mu.Unlock()
	current.once.Do(func() { current.value, current.err = r.load(key) })
	return current.value, current.err
}
