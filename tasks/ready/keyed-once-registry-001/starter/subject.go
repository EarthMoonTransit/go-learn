package subject

type Registry registryState
type registryState struct{}

func NewRegistry(load func(string) (int, error)) *Registry { return &Registry{} }
func (r *Registry) Get(key string) (int, error)            { return 0, nil }
