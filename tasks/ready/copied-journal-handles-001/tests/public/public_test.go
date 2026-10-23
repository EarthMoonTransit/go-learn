package subject_test

import (
	"slices"
	"sync"
	subject "task.local/copied-journal-handles-001/subject"
	"testing"
)

func journalSequence(t *testing.T, label string, j subject.Journal, want []int) {
	t.Helper()
	for read := 0; read < 2; read++ {
		if got := j.Snapshot(); !slices.Equal(got, want) {
			t.Errorf("%s read%d sequence=%v; want %v", label, read, got, want)
		}
	}
}
func TestJournalValueHandles(t *testing.T) {
	j := subject.NewJournal()
	journalSequence(t, "new", j, nil)
	if got := j.Append(7); got != 0 {
		t.Errorf("first append index=%d; want0", got)
	}
	alias := j
	values := []int{0, -3, 7}
	for n, v := range values {
		target := alias
		if n%2 == 1 {
			target = j
		}
		if got := target.Append(v); got != n+1 {
			t.Errorf("alias append%d index=%d; want%d", n, got, n+1)
		}
	}
	journalSequence(t, "original after copied-handle appends", j, []int{7, 0, -3, 7})
	journalSequence(t, "copy after original append", alias, []int{7, 0, -3, 7})
	other := subject.NewJournal()
	journalSequence(t, "independent new", other, nil)
	if got := other.Append(42); got != 0 {
		t.Errorf("independent append index=%d; want0", got)
	}
	journalSequence(t, "first after independent append", j, []int{7, 0, -3, 7})
	journalSequence(t, "second", other, []int{42})
}
func TestDetachedJournalSnapshots(t *testing.T) {
	j := subject.NewJournal()
	for _, v := range []int{2, 5, -1} {
		j.Append(v)
	}
	first, second := j.Snapshot(), j.Snapshot()
	if !slices.Equal(first, []int{2, 5, -1}) || !slices.Equal(second, []int{2, 5, -1}) {
		t.Errorf("initial snapshots=%v,%v; want [2 5 -1] twice", first, second)
	}
	if len(first) > 1 {
		first[1] = 90
	}
	if !slices.Equal(second, []int{2, 5, -1}) {
		t.Errorf("second snapshot changed through first=%v; want [2 5 -1]", second)
	}
	if len(second) > 0 {
		second[0] = 80
	}
	if !slices.Equal(first, []int{2, 90, -1}) {
		t.Errorf("first snapshot changed through second=%v; want [2 90 -1]", first)
	}
	journalSequence(t, "journal after caller snapshot edits", j, []int{2, 5, -1})
	j.Append(8)
	if !slices.Equal(first, []int{2, 90, -1}) || !slices.Equal(second, []int{80, 5, -1}) {
		t.Errorf("retained snapshots changed after append/read=%v,%v; want [2 90 -1],[80 5 -1]", first, second)
	}
	journalSequence(t, "journal after later append", j, []int{2, 5, -1, 8})
}
func TestConcurrentJournalAliases(t *testing.T) {
	j := subject.NewJournal()
	j.Append(11)
	j.Append(22)
	journalSequence(t, "staged prefix", j, []int{11, 22})
	handles := []subject.Journal{j, j, j, j}
	values := []int{7, 0, -4, 7, 18, -9, 0, 31}
	indices := make([]int, len(values))
	snapshots := make([][]int, len(values))
	var group sync.WaitGroup
	group.Add(len(values))
	for i, v := range values {
		go func(index, value int) {
			defer group.Done()
			h := handles[index%len(handles)]
			indices[index] = h.Append(value)
			snapshots[index] = h.Snapshot()
		}(i, v)
	}
	group.Wait()
	final := j.Snapshot()
	if len(final) != len(values)+2 {
		t.Errorf("joined sequence length=%d; want%d", len(final), len(values)+2)
	}
	if len(final) < 2 || !slices.Equal(final[:2], []int{11, 22}) {
		t.Errorf("joined sequence=%v; want initial prefix [11 22]", final)
	}
	seen := map[int]bool{}
	for i, index := range indices {
		if index < 2 || index >= len(values)+2 || seen[index] {
			t.Errorf("call%d append index=%d; want unique index in[2,%d)", i, index, len(values)+2)
		}
		seen[index] = true
		if index >= 0 && index < len(final) && final[index] != values[i] {
			t.Errorf("call%d final[%d]=%d; want argument%d", i, index, final[index], values[i])
		}
	}
	for i, snapshot := range snapshots {
		if len(snapshot) > len(final) || !slices.Equal(snapshot, final[:min(len(snapshot), len(final))]) {
			t.Errorf("concurrent snapshot%d=%v; want prefix of final%v", i, snapshot, final)
		}
	}
}
