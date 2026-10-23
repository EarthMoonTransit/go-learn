package subject_test

import (
	"slices"
	"sync"
	subject "task.local/copied-journal-handles-001/subject"
	"testing"
)

func TestGraderJournalFullBound(t *testing.T) {
	j := subject.NewJournal()
	aliases := []subject.Journal{j, j, j}
	positions := make([]int, 16)
	var joined sync.WaitGroup
	joined.Add(16)
	for n := 0; n < 16; n++ {
		go func(i int) { defer joined.Done(); positions[i] = aliases[i%3].Append(i*3 - 17) }(n)
	}
	joined.Wait()
	got := j.Snapshot()
	seen := make(map[int]bool)
	if len(got) != 16 {
		t.Errorf("full-bound length=%d; want16", len(got))
	}
	for i, p := range positions {
		if p < 0 || p >= 16 || seen[p] {
			t.Errorf("full-bound call%d position=%d; want unique0..15", i, p)
		}
		seen[p] = true
		if p >= 0 && p < len(got) && got[p] != i*3-17 {
			t.Errorf("full-bound value at%d=%d; want%d", p, got[p], i*3-17)
		}
	}
	before := append([]int(nil), got...)
	if len(got) > 0 {
		got[len(got)-1] = 999
	}
	if now := j.Snapshot(); !slices.Equal(now, before) {
		t.Errorf("full-bound snapshot edit changed journal=%v; want%v", now, before)
	}
	empty := subject.NewJournal()
	if len(empty.Snapshot()) != 0 {
		t.Error("later new journal inherited previous state")
	}
	copyAfterUse := j
	if after := copyAfterUse.Snapshot(); !slices.Equal(after, before) {
		t.Errorf("late handle copy=%v; want%v", after, before)
	}
}
