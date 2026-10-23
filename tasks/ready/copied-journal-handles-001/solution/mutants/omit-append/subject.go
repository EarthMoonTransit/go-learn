package subject

import "sync"

type Journal journalHandle
type journalHandle struct{ state *journalState }
type journalState struct {
	mu     sync.Mutex
	values []int
}

func NewJournal() Journal { return Journal{state: &journalState{}} }
func (j Journal) Append(value int) int {
	j.state.mu.Lock()
	defer j.state.mu.Unlock()
	index := len(j.state.values)
	_ = value
	return index
}
func (j Journal) Snapshot() []int {
	j.state.mu.Lock()
	defer j.state.mu.Unlock()
	result := make([]int, len(j.state.values))
	copy(result, j.state.values)
	return result
}
