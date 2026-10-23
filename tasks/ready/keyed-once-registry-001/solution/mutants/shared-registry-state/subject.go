package subject

import "sync"

type entry struct {
	once  sync.Once
	value int
	err   error
}

var sharedMu sync.Mutex
var sharedEntries = make(map[string]*entry)

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
	sharedMu.Lock()
	current := sharedEntries[key]
	if current == nil {
		current = &entry{}
		sharedEntries[key] = current
	}
	sharedMu.Unlock()
	current.once.Do(func() { current.value, current.err = r.load(key) })
	return current.value, current.err
}
