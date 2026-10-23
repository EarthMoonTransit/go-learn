package subject

import "sync"

type Journal journalHandle
type journalHandle struct{ state *journalState }
type journalState struct {
	mu      sync.Mutex
	values  []int
	scratch []int
}

func NewJournal() Journal { return Journal{state: &journalState{}} }
func (j Journal) Append(value int) int {
	j.state.mu.Lock()
	defer j.state.mu.Unlock()
	index := len(j.state.values)
	j.state.values = append(j.state.values, value)
	return index
}
func (j Journal) Snapshot() []int {
	j.state.mu.Lock()
	defer j.state.mu.Unlock()
	if cap(j.state.scratch) < len(j.state.values) {
		j.state.scratch = make([]int, len(j.state.values))
	}
	j.state.scratch = j.state.scratch[:len(j.state.values)]
	copy(j.state.scratch, j.state.values)
	result := j.state.scratch
	return result
}
