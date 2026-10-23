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
	j.state.values = append(j.state.values, value)
	return index
}
func (j Journal) Snapshot() []int {
	j.state.mu.Lock()
	defer j.state.mu.Unlock()
	result := make([]int, len(j.state.values))
	copy(result, j.state.values)
	for a, b := 0, len(result)-1; a < b; a, b = a+1, b-1 {
		result[a], result[b] = result[b], result[a]
	}
	return result
}
