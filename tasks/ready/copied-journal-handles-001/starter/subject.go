package subject

type Journal journalHandle
type journalHandle struct{}

func NewJournal() Journal              { return Journal{} }
func (j Journal) Append(value int) int { return 0 }
func (j Journal) Snapshot() []int      { return nil }
